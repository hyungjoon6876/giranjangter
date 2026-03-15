package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/event"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func main() {
	cfg := config.Load()

	// Database
	db, err := repository.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	// Seed data (dev only)
	if cfg.IsDev() {
		if err := repository.SeedDB(db); err != nil {
			log.Printf("seed warning: %v", err)
		}
	}

	// Cleanup expired refresh tokens on startup
	if result, err := db.Exec("DELETE FROM refresh_tokens WHERE expires_at < NOW()"); err != nil {
		log.Printf("refresh token cleanup warning: %v", err)
	} else if n, _ := result.RowsAffected(); n > 0 {
		log.Printf("cleaned up %d expired refresh tokens", n)
	}

	// Periodic cleanup of expired refresh tokens (every 24 hours)
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			result, err := db.Exec("DELETE FROM refresh_tokens WHERE expires_at < NOW()")
			if err != nil {
				log.Printf("refresh token cleanup error: %v", err)
				continue
			}
			if n, _ := result.RowsAffected(); n > 0 {
				log.Printf("cleaned up %d expired refresh tokens", n)
			}
		}
	}()

	// SSE Broker
	sseBroker := event.NewBroker()

	// Auth middleware
	auth := middleware.NewAuthMiddleware(cfg.JWTSecret, cfg.JWTAccessTTL, cfg.JWTRefreshTTL)

	// Router
	if !cfg.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// CORS
	r.Use(corsMiddleware(cfg))

	// Static files (uploads + icons)
	r.Static("/uploads", cfg.UploadDir)
	r.Static("/static", "./static")

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.1.0",
			"sse":     sseBroker.OnlineCount(),
		})
	})

	// API v1
	v1 := r.Group("/api/v1")
	{
		// Public routes
		v1.GET("/servers", listServers(db))
		v1.GET("/categories", listCategories(db))
		v1.GET("/items/search", searchItems(db))

		// Auth routes
		v1.POST("/auth/login", handleLogin(db, auth, cfg))
		v1.POST("/auth/refresh", handleRefresh(db, auth, cfg))

		// Read-only routes (JWT only, no DB check — restricted users can read)
		readOnly := v1.Group("")
		readOnly.Use(auth.RequireAuth())
		{
			readOnly.GET("/me", handleGetMe(db))
			readOnly.GET("/me/listings", handleMyListings(db))
			readOnly.GET("/chats", handleListChats(db))
			readOnly.GET("/chats/:chatId/messages", handleListMessages(db))
			readOnly.GET("/me/trades", handleMyTrades(db))
			readOnly.GET("/notifications", handleListNotifications(db))
			readOnly.GET("/users/:userId/reviews", handleGetUserReviews(db))
			readOnly.GET("/me/reports", handleMyReports(db))
			readOnly.GET("/sse/connect", handleSSEConnect(sseBroker))
		}

		// Write operations — DB status check + restricted 사용자 차단
		write := v1.Group("")
		write.Use(auth.RequireAuthWithDB(db))
		write.Use(middleware.RejectIfRestricted())
		{
			// Auth
			write.POST("/auth/logout", handleLogout(db))

			write.PATCH("/me/profile", handleUpdateProfile(db))

			// Listings
			write.POST("/listings", handleCreateListing(db))
			write.PATCH("/listings/:id", handleUpdateListing(db))
			write.POST("/listings/:id/status", handleChangeListingStatus(db))
			write.POST("/listings/:id/favorite", handleFavoriteListing(db))
			write.DELETE("/listings/:id/favorite", handleUnfavoriteListing(db))

			// Chat
			write.POST("/listings/:id/chats", handleCreateChat(db))
			write.POST("/chats/:chatId/messages", handleSendMessage(db, sseBroker))
			write.POST("/chats/:chatId/read", handleMarkRead(db))

			// Reservations
			write.POST("/chats/:chatId/reservations", handleCreateReservation(db))
			write.POST("/reservations/:resId/confirm", handleConfirmReservation(db))
			write.POST("/reservations/:resId/cancel", handleCancelReservation(db))

			// Trade completion
			write.POST("/listings/:id/complete", handleCompleteTrade(db))
			write.POST("/trade-completions/:compId/confirm", handleConfirmCompletion(db))

			// Reviews
			write.POST("/trade-completions/:compId/reviews", handleCreateReview(db))

			// Reports
			write.POST("/reports", handleCreateReport(db))

			// Notifications
			write.POST("/notifications/read", handleReadNotifications(db))

			// Upload
			write.POST("/uploads/images", handleUploadImage(cfg, db))

			// Block
			write.POST("/users/:userId/block", handleBlockUser(db))
			write.DELETE("/users/:userId/block", handleUnblockUser(db))
		}

		// Public listing routes (optional auth for favorited status)
		v1.GET("/listings", auth.OptionalAuth(), handleListListings(db))
		v1.GET("/listings/:id", auth.OptionalAuth(), handleGetListing(db))

		// Admin routes
		admin := v1.Group("/admin")
		admin.Use(auth.RequireAuth(), middleware.RequireRole("moderator", "admin"))
		{
			admin.GET("/reports", handleAdminListReports(db))
			admin.GET("/reports/:reportId", handleAdminGetReport(db))
			admin.POST("/reports/:reportId/actions", handleAdminReportAction(db))
			admin.POST("/listings/:id/hide", handleAdminHideListing(db))
			admin.POST("/users/:userId/restrict", handleAdminRestrictUser(db))

			// Dashboard & user management
			admin.GET("/dashboard/stats", handleAdminDashboardStats(db))
			admin.GET("/users", handleAdminListUsers(db))
			admin.GET("/users/:userId", handleAdminGetUser(db))
			admin.GET("/users/:userId/moderation-history", handleAdminUserModerationHistory(db))

			// Audit, chat inspection, trades, listings
			admin.GET("/audit-logs", handleAdminListAuditLogs(db))
			admin.GET("/chats/:chatId/messages", handleAdminChatMessages(db))
			admin.GET("/trades", handleAdminListTrades(db))
			admin.GET("/listings", handleAdminListAllListings(db))
			admin.POST("/listings/:id/restore", handleAdminRestoreListing(db))
			admin.PATCH("/reports/:reportId", handleAdminUpdateReportStatus(db))
		}
	}

	log.Printf("lincle-api starting on :%s (env=%s)", cfg.Port, cfg.Env)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func corsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		for _, allowed := range cfg.AllowedOrigins {
			if origin == allowed {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}
		if cfg.IsDev() {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
