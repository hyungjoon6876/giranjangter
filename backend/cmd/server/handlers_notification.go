package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func handleListNotifications(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		items, err := repo.ListNotifications(c.Request.Context(), userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var notifs []gin.H
		for _, item := range items {
			notifs = append(notifs, gin.H{
				"notificationId": item.NotificationID, "type": item.Type, "title": item.Title, "body": item.Body,
				"referenceType": item.ReferenceType, "referenceId": item.ReferenceID, "deepLink": item.DeepLink,
				"isRead": item.IsRead, "createdAt": item.CreatedAt.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": notifs})
	}
}

func handleReadNotifications(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		var req struct {
			NotificationIDs []string `json:"notificationIds" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}
		repo.MarkNotificationsRead(c.Request.Context(), userID, req.NotificationIDs)
		c.Status(http.StatusNoContent)
	}
}
