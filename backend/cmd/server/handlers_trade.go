package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/alignment"
	"github.com/jym/lincle/internal/domain"
	"github.com/jym/lincle/internal/middleware"
)

func handleCompleteTrade(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		userID := middleware.GetUserID(c)
		var req struct {
			ReservationID  string  `json:"reservationId" binding:"required"`
			CompletionNote *string `json:"completionNote"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		var proposer, resCp string
		err = tx.QueryRow("SELECT proposer_user_id, counterpart_user_id FROM reservations WHERE id = $1 AND listing_id = $2 AND status = 'confirmed'", req.ReservationID, listingID).Scan(&proposer, &resCp)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "확정된 예약을 찾을 수 없습니다."}})
			return
		}

		if userID != proposer && userID != resCp {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "거래 참여자만 완료 요청할 수 있습니다."}})
			return
		}

		counterpart := resCp
		if userID == resCp {
			counterpart = proposer
		}

		compID := uuid.New().String()
		now := time.Now().UTC()
		expiresAt := now.Add(48 * time.Hour)

		_, err = tx.Exec(`INSERT INTO trade_completions (id, listing_id, reservation_id, requested_by_user_id, counterpart_user_id, status, completion_note, auto_confirm_at, created_at)
			VALUES ($1, $2, $3, $4, $5, 'pending_confirm', $6, $7, $8)`,
			compID, listingID, req.ReservationID, userID, counterpart, req.CompletionNote, expiresAt, now)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "거래 완료 요청 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"completionId":     compID,
			"completionStatus": "pending_confirm",
			"expiresAt":        expiresAt.Format(time.RFC3339),
		})
	}
}

func handleConfirmCompletion(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		compID := c.Param("compId")
		userID := middleware.GetUserID(c)

		var counterpart, listingID, resID, requestedBy string
		err := db.QueryRow("SELECT counterpart_user_id, listing_id, reservation_id, requested_by_user_id FROM trade_completions WHERE id = $1 AND status = 'pending_confirm'", compID).Scan(&counterpart, &listingID, &resID, &requestedBy)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "완료 요청을 찾을 수 없습니다."}})
			return
		}
		if counterpart != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "거래 상대방만 확인할 수 있습니다."}})
			return
		}

		now := time.Now().UTC()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("UPDATE trade_completions SET status = 'confirmed', confirmed_at = $1 WHERE id = $2", now, compID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "거래 확정 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE reservations SET status = 'fulfilled' WHERE id = $1", resID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "예약 상태 업데이트 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE listings SET status = 'completed', updated_at = $1, last_activity_at = $2 WHERE id = $3", now, now, listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 상태 업데이트 실패"}})
			return
		}

		// 양측 성향치 +5 (거래 완료)
		if err := alignment.Change(tx, requestedBy, domain.AlignmentTradeConfirmed, "trade_confirmed", "completion", compID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "성향치 업데이트 실패"}})
			return
		}
		if err := alignment.Change(tx, counterpart, domain.AlignmentTradeConfirmed, "trade_confirmed", "completion", compID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "성향치 업데이트 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"completionId": compID, "completionStatus": "confirmed"})
	}
}

func handleMyTrades(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		rows, err := db.Query(`
			SELECT cr.id, cr.listing_id, l.title, l.status,
				CASE WHEN cr.seller_user_id = $1 THEN cr.buyer_user_id ELSE cr.seller_user_id END as cp_id,
				p.nickname, cr.chat_status, cr.updated_at
			FROM chat_rooms cr
			JOIN listings l ON cr.listing_id = l.id
			JOIN user_profiles p ON p.user_id = CASE WHEN cr.seller_user_id = $2 THEN cr.buyer_user_id ELSE cr.seller_user_id END
			WHERE (cr.seller_user_id = $3 OR cr.buyer_user_id = $4)
			ORDER BY cr.updated_at DESC LIMIT 50`, userID, userID, userID, userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		defer rows.Close()

		var trades []gin.H
		for rows.Next() {
			var chatID, listingID, title, listingStatus, cpID, cpNick, chatStatus string
			var updated time.Time
			if err := rows.Scan(&chatID, &listingID, &title, &listingStatus, &cpID, &cpNick, &chatStatus, &updated); err != nil {
				continue
			}
			trades = append(trades, gin.H{
				"chatRoomId": chatID, "listingId": listingID, "listingTitle": title,
				"tradeStatus": listingStatus, "chatStatus": chatStatus,
				"counterparty": gin.H{"userId": cpID, "nickname": cpNick},
				"updatedAt": updated.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": trades})
	}
}
