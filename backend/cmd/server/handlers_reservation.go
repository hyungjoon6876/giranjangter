package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/middleware"
)

func handleCreateReservation(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)
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

		var listingID, sellerID, buyerID string
		err := db.QueryRow("SELECT listing_id, seller_user_id, buyer_user_id FROM chat_rooms WHERE id = $1", chatID).Scan(&listingID, &sellerID, &buyerID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "채팅방을 찾을 수 없습니다."}})
			return
		}

		var activeCount int
		db.QueryRow("SELECT COUNT(*) FROM reservations WHERE listing_id = $1 AND status IN ('proposed','confirmed')", listingID).Scan(&activeCount)
		if activeCount > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": gin.H{"code": "CONFLICT", "message": "이미 활성 예약이 존재합니다."}})
			return
		}

		counterpart := buyerID
		if userID == buyerID {
			counterpart = sellerID
		}

		resID := uuid.New().String()
		now := time.Now().UTC()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec(`INSERT INTO reservations (id, listing_id, chat_room_id, proposer_user_id, counterpart_user_id, status, scheduled_at, meeting_type, server_id, meeting_point_text, note_to_counterparty, expires_at, created_at)
			VALUES ($1, $2, $3, $4, $5, 'proposed', $6, $7, $8, $9, $10, $11, $12)`,
			resID, listingID, chatID, userID, counterpart, req.ScheduledAt, req.MeetingType, req.ServerID, req.MeetingPoint, req.Note, req.ExpiresAt, now); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 생성 실패"}})
			return
		}

		meta, _ := json.Marshal(gin.H{"eventType": "reservation_proposed", "reservationId": resID})
		if _, err := tx.Exec("INSERT INTO chat_messages (id, chat_room_id, message_type, body_text, metadata_json, sent_at) VALUES ($1, $2, 'system', '예약이 제안되었습니다.', $3, $4)",
			uuid.New().String(), chatID, string(meta), now); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "시스템 메시지 생성 실패"}})
			return
		}

		if _, err := tx.Exec("UPDATE chat_rooms SET chat_status = 'reservation_proposed', updated_at = $1 WHERE id = $2", now, chatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "채팅 상태 업데이트 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"reservationId": resID, "status": "proposed"})
	}
}

func handleConfirmReservation(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		resID := c.Param("resId")
		userID := middleware.GetUserID(c)

		var counterpart, listingID, chatRoomID string
		err := db.QueryRow("SELECT counterpart_user_id, listing_id, chat_room_id FROM reservations WHERE id = $1 AND status = 'proposed'", resID).Scan(&counterpart, &listingID, &chatRoomID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "예약을 찾을 수 없습니다."}})
			return
		}
		if counterpart != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "예약 상대방만 확정할 수 있습니다."}})
			return
		}

		now := time.Now().UTC()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("UPDATE reservations SET status = 'confirmed', confirmed_at = $1 WHERE id = $2", now, resID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 확정 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE listings SET status = 'reserved', reserved_chat_room_id = $1, updated_at = $2, last_activity_at = $3 WHERE id = $4", chatRoomID, now, now, listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 상태 업데이트 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE chat_rooms SET chat_status = 'reservation_confirmed', updated_at = $1 WHERE id = $2", now, chatRoomID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "채팅 상태 업데이트 실패"}})
			return
		}
		if _, err := tx.Exec(`INSERT INTO status_history (id, entity_type, entity_id, from_status, to_status, changed_by_user_id, created_at) VALUES ($1, 'listing', $2, 'available', 'reserved', $3, $4)`,
			uuid.New().String(), listingID, userID, now); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "이력 기록 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"reservationId": resID, "status": "confirmed"})
	}
}

func handleCancelReservation(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		resID := c.Param("resId")
		userID := middleware.GetUserID(c)
		var req struct {
			ReasonCode string `json:"reasonCode"`
		}
		c.ShouldBindJSON(&req)

		var listingID, chatRoomID, proposer, counterpart string
		err := db.QueryRow("SELECT listing_id, chat_room_id, proposer_user_id, counterpart_user_id FROM reservations WHERE id = $1 AND status IN ('proposed','confirmed')", resID).Scan(&listingID, &chatRoomID, &proposer, &counterpart)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "예약을 찾을 수 없습니다."}})
			return
		}
		if userID != proposer && userID != counterpart {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "예약 참여자만 취소할 수 있습니다."}})
			return
		}

		now := time.Now().UTC()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("UPDATE reservations SET status = 'cancelled', cancelled_at = $1, cancellation_reason_code = $2 WHERE id = $3", now, req.ReasonCode, resID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 취소 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE listings SET status = 'available', reserved_chat_room_id = NULL, updated_at = $1, last_activity_at = $2 WHERE id = $3", now, now, listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 상태 업데이트 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE chat_rooms SET chat_status = 'open', updated_at = $1 WHERE id = $2", now, chatRoomID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "채팅 상태 업데이트 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"reservationId": resID, "status": "cancelled"})
	}
}
