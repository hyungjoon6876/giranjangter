package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func handleCreateReservation(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req struct {
			ScheduledAt  string  `json:"scheduledAt" binding:"required"`
			MeetingType  string  `json:"meetingType" binding:"required"`
			ServerID     *string `json:"serverId"`
			MeetingPoint *string `json:"meetingPointText"`
			Note         *string `json:"noteToCounterparty"`
			ExpiresAt    *string `json:"expiresAt"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		info, err := repo.GetChatRoomForReservation(ctx, chatID, userID)
		if err != nil || info == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "채팅방을 찾을 수 없습니다."}})
			return
		}

		activeCount, err := repo.CountActiveReservations(ctx, info.ListingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		if activeCount > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": gin.H{"code": "CONFLICT", "message": "이미 활성 예약이 존재합니다."}})
			return
		}

		counterpart := info.BuyerID
		if userID == info.BuyerID {
			counterpart = info.SellerID
		}

		resID := uuid.New().String()
		now := time.Now().UTC()
		meta, _ := json.Marshal(gin.H{"eventType": "reservation_proposed", "reservationId": resID})

		if err := repo.CreateReservation(ctx, &repository.CreateReservationParams{
			ReservationID:   resID,
			ListingID:       info.ListingID,
			ChatRoomID:      chatID,
			ProposerUserID:  userID,
			CounterpartID:   counterpart,
			ScheduledAt:     req.ScheduledAt,
			MeetingType:     req.MeetingType,
			ServerID:        req.ServerID,
			MeetingPoint:    req.MeetingPoint,
			Note:            req.Note,
			ExpiresAt:       req.ExpiresAt,
			SystemMessageID: uuid.New().String(),
			SystemMetaJSON:  string(meta),
			Now:             now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 생성 실패"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"reservationId": resID, "status": "proposed"})
	}
}

func handleConfirmReservation(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		resID := c.Param("resId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		info, err := repo.GetReservationForConfirm(ctx, resID)
		if err != nil || info == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "예약을 찾을 수 없습니다."}})
			return
		}
		if info.CounterpartUserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "예약 상대방만 확정할 수 있습니다."}})
			return
		}

		now := time.Now().UTC()

		if err := repo.ConfirmReservation(ctx, &repository.ConfirmReservationParams{
			ReservationID:   resID,
			ListingID:       info.ListingID,
			ChatRoomID:      info.ChatRoomID,
			UserID:          userID,
			StatusHistoryID: uuid.New().String(),
			Now:             now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 확정 실패"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"reservationId": resID, "status": "confirmed"})
	}
}

func handleCancelReservation(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		resID := c.Param("resId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req struct {
			ReasonCode string `json:"reasonCode"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		info, err := repo.GetReservationForCancel(ctx, resID)
		if err != nil || info == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "예약을 찾을 수 없습니다."}})
			return
		}
		if userID != info.ProposerID && userID != info.CounterpartID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "예약 참여자만 취소할 수 있습니다."}})
			return
		}

		now := time.Now().UTC()

		if err := repo.CancelReservation(ctx, &repository.CancelReservationParams{
			ReservationID: resID,
			ListingID:     info.ListingID,
			ChatRoomID:    info.ChatRoomID,
			ReasonCode:    req.ReasonCode,
			Now:           now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 취소 실패"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"reservationId": resID, "status": "cancelled"})
	}
}
