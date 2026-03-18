package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/oauth"
	"github.com/jym/lincle/internal/repository"
)

// hashToken creates a SHA-256 hash of a token string for secure DB storage.
func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

type loginRequest struct {
	Provider      string `json:"provider" binding:"required"`
	ProviderToken string `json:"providerToken" binding:"required"`
}

type refreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func handleLogin(repo repository.AuthRepo, auth *middleware.AuthMiddleware, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		var providerKey string

		switch req.Provider {
		case "google":
			if cfg.IsDev() && len(cfg.GoogleClientIDs) == 0 {
				// Dev mode without Google credentials: use token as-is
				providerKey = req.ProviderToken
				log.Println("[auth] dev mode: skipping Google token verification")
			} else {
				// Production: verify Google ID token
				info, err := oauth.VerifyGoogleIDToken(req.ProviderToken, cfg.GoogleClientIDs)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": gin.H{"code": "UNAUTHORIZED", "message": "Google 인증에 실패했습니다."},
					})
					return
				}
				providerKey = info.Sub // Google's stable user ID
			}
		default:
			if cfg.IsDev() {
				// Dev mode: allow any provider with token as key
				providerKey = req.ProviderToken
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": gin.H{"code": "VALIDATION_ERROR", "message": "지원하지 않는 로그인 방식입니다."},
				})
				return
			}
		}

		ctx := c.Request.Context()
		var userID, role, nickname string
		var isNew bool

		existing, err := repo.FindUserByProvider(ctx, req.Provider, providerKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}

		if existing == nil {
			userID = uuid.New().String()
			role = "user"
			nickname = "유저_" + userID[:8]
			isNew = true

			if err := repo.CreateUserWithProfile(ctx, userID, req.Provider, providerKey, nickname); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{"code": "INTERNAL_ERROR", "message": "계정 생성 실패"},
				})
				return
			}
		} else {
			userID = existing.UserID
			role = existing.Role
			nickname = existing.Nickname
		}

		repo.UpdateLastLogin(ctx, userID, time.Now().UTC())

		accessToken, refreshToken, err := auth.GenerateTokens(userID, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "토큰 생성 실패"},
			})
			return
		}

		// Store refresh token hash in DB for server-side management
		if err := repo.StoreRefreshToken(ctx, uuid.New().String(), userID, hashToken(refreshToken), time.Now().Add(cfg.JWTRefreshTTL)); err != nil {
			log.Printf("[auth] failed to store refresh token: %v", err)
		}

		status := http.StatusOK
		if isNew {
			status = http.StatusCreated
		}

		c.JSON(status, gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"expiresIn":    900,
			"user": gin.H{
				"userId":        userID,
				"nickname":      nickname,
				"role":          role,
				"accountStatus": "active",
				"isNewUser":     isNew,
			},
		})
	}
}

func handleRefresh(repo repository.AuthRepo, auth *middleware.AuthMiddleware, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req refreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		claims, err := auth.ParseToken(req.RefreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "유효하지 않은 리프레시 토큰"},
			})
			return
		}

		ctx := c.Request.Context()

		// Verify refresh token exists in DB and hasn't expired
		tokenID, err := repo.FindRefreshToken(ctx, hashToken(req.RefreshToken))
		if err != nil || tokenID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "유효하지 않은 리프레시 토큰"},
			})
			return
		}

		// Check account status FIRST (before deleting anything)
		accountStatus, err := repo.GetAccountStatus(ctx, claims.UserID)
		if err != nil || accountStatus == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "사용자를 찾을 수 없습니다."},
			})
			return
		}
		if accountStatus == "suspended" || accountStatus == "withdrawn" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": gin.H{"code": "FORBIDDEN", "message": "계정이 비활성 상태입니다."},
			})
			return
		}

		accessToken, refreshToken, err := auth.GenerateTokens(claims.UserID, claims.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "토큰 생성 실패"},
			})
			return
		}

		// Rotate: delete old token and store new one atomically
		if err := repo.RotateRefreshToken(ctx, tokenID, uuid.New().String(), claims.UserID, hashToken(refreshToken), time.Now().Add(cfg.JWTRefreshTTL)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"expiresIn":    900,
		})
	}
}

func handleGetMe(repo repository.AuthRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		profile, err := repo.GetUserProfile(c.Request.Context(), userID)
		if err != nil || profile == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{"code": "NOT_FOUND", "message": "사용자를 찾을 수 없습니다."},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"userId":              profile.UserID,
			"role":                profile.Role,
			"accountStatus":       profile.AccountStatus,
			"nickname":            profile.Nickname,
			"avatarUrl":           profile.AvatarURL,
			"introduction":        profile.Introduction,
			"primaryServerId":     profile.PrimaryServerID,
			"completedTradeCount": profile.TradeCount,
			"positiveReviewCount": profile.ReviewCount,
			"responseBadge":       profile.ResponseBadge,
			"trustBadge":          profile.TrustBadge,
			"alignmentScore":      profile.AlignmentScore,
			"alignmentGrade":      profile.AlignmentGrade,
		})
	}
}

func handleUpdateProfile(repo repository.AuthRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		var req struct {
			Nickname      *string `json:"nickname"`
			Introduction  *string `json:"introduction"`
			PrimaryServer *string `json:"primaryServerId"`
			AvatarURL     *string `json:"avatarUrl"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		fields := repository.ProfileUpdateFields{
			Nickname:      req.Nickname,
			Introduction:  req.Introduction,
			PrimaryServer: req.PrimaryServer,
			AvatarURL:     req.AvatarURL,
		}

		if fields.Nickname == nil && fields.Introduction == nil && fields.PrimaryServer == nil && fields.AvatarURL == nil {
			handleGetMe(repo)(c)
			return
		}

		if err := repo.UpdateProfile(c.Request.Context(), userID, fields); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "프로필 수정 실패"}})
			return
		}

		handleGetMe(repo)(c)
	}
}

func handleLogout(repo repository.AuthRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "인증이 필요합니다."},
			})
			return
		}
		// Delete all refresh tokens for this user (logout from all devices)
		repo.DeleteRefreshTokensByUser(c.Request.Context(), userID)
		c.Status(http.StatusNoContent)
	}
}
