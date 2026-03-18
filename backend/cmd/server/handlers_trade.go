package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func handleCompleteTrade(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req struct {
			ReservationID  string  `json:"reservationId" binding:"required"`
			CompletionNote *string `json:"completionNote"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		participants, err := repo.GetConfirmedReservation(ctx, req.ReservationID, listingID)
		if err != nil || participants == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "확정된 예약을 찾을 수 없습니다."}})
			return
		}

		if userID != participants.ProposerID && userID != participants.CounterpartID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "거래 참여자만 완료 요청할 수 있습니다."}})
			return
		}

		counterpart := participants.CounterpartID
		if userID == participants.CounterpartID {
			counterpart = participants.ProposerID
		}

		compID := uuid.New().String()
		now := time.Now().UTC()
		expiresAt := now.Add(48 * time.Hour)

		if err := repo.CreateTradeCompletion(ctx, &repository.CreateTradeCompletionParams{
			ID:             compID,
			ListingID:      listingID,
			ReservationID:  req.ReservationID,
			RequestedByID:  userID,
			CounterpartID:  counterpart,
			CompletionNote: req.CompletionNote,
			AutoConfirmAt:  expiresAt,
			Now:            now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "거래 완료 요청 실패"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"completionId":     compID,
			"completionStatus": "pending_confirm",
			"expiresAt":        expiresAt.Format(time.RFC3339),
		})
	}
}

func handleConfirmCompletion(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		compID := c.Param("compId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		info, err := repo.GetPendingCompletion(ctx, compID)
		if err != nil || info == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "완료 요청을 찾을 수 없습니다."}})
			return
		}
		if info.CounterpartUserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "거래 상대방만 확인할 수 있습니다."}})
			return
		}

		now := time.Now().UTC()

		if err := repo.ConfirmCompletion(ctx, &repository.ConfirmCompletionParams{
			CompletionID:  compID,
			ReservationID: info.ReservationID,
			ListingID:     info.ListingID,
			RequestedByID: info.RequestedByUserID,
			CounterpartID: info.CounterpartUserID,
			Now:           now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "거래 확정 실패"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"completionId": compID, "completionStatus": "confirmed"})
	}
}

func handleMyTrades(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		items, err := repo.ListMyTrades(c.Request.Context(), userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var trades []gin.H
		for _, item := range items {
			trades = append(trades, gin.H{
				"chatRoomId": item.ChatRoomID, "listingId": item.ListingID, "listingTitle": item.ListingTitle,
				"tradeStatus": item.ListingStatus, "chatStatus": item.ChatStatus,
				"counterparty": gin.H{"userId": item.CounterpartID, "nickname": item.CounterpartNick},
				"updatedAt": item.UpdatedAt.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": trades})
	}
}
