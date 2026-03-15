package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/middleware"
)

func handleListNotifications(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		rows, err := db.Query("SELECT id, type, title, body, reference_type, reference_id, deep_link, is_read, created_at FROM notifications WHERE user_id = $1 ORDER BY created_at DESC LIMIT 50", userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		defer rows.Close()

		var notifs []gin.H
		for rows.Next() {
			var id, nType, title, body string
			var refType, refID, deepLink *string
			var isRead bool
			var created time.Time
			if err := rows.Scan(&id, &nType, &title, &body, &refType, &refID, &deepLink, &isRead, &created); err != nil {
				continue
			}
			notifs = append(notifs, gin.H{
				"notificationId": id, "type": nType, "title": title, "body": body,
				"referenceType": refType, "referenceId": refID, "deepLink": deepLink,
				"isRead": isRead, "createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": notifs})
	}
}

func handleReadNotifications(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		var req struct {
			NotificationIDs []string `json:"notificationIds" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}
		placeholders := make([]string, len(req.NotificationIDs))
		args := make([]interface{}, 0, len(req.NotificationIDs)+1)
		args = append(args, userID)
		for i, nid := range req.NotificationIDs {
			placeholders[i] = fmt.Sprintf("$%d", i+2)
			args = append(args, nid)
		}
		db.Exec("UPDATE notifications SET is_read = true WHERE user_id = $1 AND id IN ("+strings.Join(placeholders, ",")+")", args...)
		c.Status(http.StatusNoContent)
	}
}
