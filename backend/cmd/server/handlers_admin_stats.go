package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handleAdminDashboardStats(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		var stats struct {
			TotalUsers       int `json:"totalUsers"`
			NewUsersToday    int `json:"newUsersToday"`
			ActiveListings   int `json:"activeListings"`
			NewListingsToday int `json:"newListingsToday"`
			PendingReports   int `json:"pendingReports"`
			TradesToday      int `json:"tradesToday"`
			ActiveChats      int `json:"activeChats"`
			RestrictedUsers  int `json:"restrictedUsers"`
		}

		_ = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&stats.TotalUsers)
		_ = db.QueryRow("SELECT COUNT(*) FROM users WHERE created_at >= $1", today).Scan(&stats.NewUsersToday)
		_ = db.QueryRow("SELECT COUNT(*) FROM listings WHERE status = 'available' AND visibility = 'public'").Scan(&stats.ActiveListings)
		_ = db.QueryRow("SELECT COUNT(*) FROM listings WHERE created_at >= $1", today).Scan(&stats.NewListingsToday)
		_ = db.QueryRow("SELECT COUNT(*) FROM reports WHERE status = 'submitted'").Scan(&stats.PendingReports)
		_ = db.QueryRow("SELECT COUNT(*) FROM trade_completions WHERE confirmed_at >= $1", today).Scan(&stats.TradesToday)
		_ = db.QueryRow("SELECT COUNT(*) FROM chat_rooms WHERE chat_status = 'open'").Scan(&stats.ActiveChats)
		_ = db.QueryRow("SELECT COUNT(*) FROM users WHERE account_status IN ('restricted', 'suspended')").Scan(&stats.RestrictedUsers)

		c.JSON(http.StatusOK, stats)
	}
}
