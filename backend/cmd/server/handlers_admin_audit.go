package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/middleware"
)

func handleAdminListAuditLogs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT id, actor_id, actor_role, action, target_type, target_id, details_json, ip_address, created_at
			FROM audit_logs ORDER BY created_at DESC LIMIT 100`)
		if err != nil {
			log.Printf("audit log error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var logs []gin.H
		for rows.Next() {
			var id, action, targetType, targetID string
			var actorID, actorRole, details, ip sql.NullString
			var created time.Time
			if err := rows.Scan(&id, &actorID, &actorRole, &action, &targetType, &targetID, &details, &ip, &created); err != nil {
				continue
			}
			logs = append(logs, gin.H{
				"logId":      id,
				"actorId":    nullStr(actorID),
				"actorRole":  nullStr(actorRole),
				"action":     action,
				"targetType": targetType,
				"targetId":   targetID,
				"details":    nullStr(details),
				"ipAddress":  nullStr(ip),
				"createdAt":  created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": logs})
	}
}

func handleAdminChatMessages(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")

		// Verify chat exists
		var exists bool
		db.QueryRow("SELECT EXISTS(SELECT 1 FROM chat_rooms WHERE id = $1)", chatID).Scan(&exists)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "채팅방을 찾을 수 없습니다."}})
			return
		}

		rows, err := db.Query(`SELECT id, sender_user_id, message_type, body_text, metadata_json, sent_at
			FROM chat_messages WHERE chat_room_id = $1 AND deleted_at IS NULL
			ORDER BY sent_at ASC LIMIT 200`, chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var messages []gin.H
		for rows.Next() {
			var id, msgType string
			var senderID, bodyText, metadata sql.NullString
			var sentAt time.Time
			if err := rows.Scan(&id, &senderID, &msgType, &bodyText, &metadata, &sentAt); err != nil {
				continue
			}
			messages = append(messages, gin.H{
				"messageId":    id,
				"senderUserId": nullStr(senderID),
				"messageType":  msgType,
				"bodyText":     nullStr(bodyText),
				"metadataJson": nullStr(metadata),
				"sentAt":       sentAt.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": messages})
	}
}

func handleAdminListTrades(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT tc.id, tc.listing_id, l.title, tc.status,
			tc.requested_by_user_id, tc.counterpart_user_id, tc.auto_confirm_at, tc.created_at
			FROM trade_completions tc
			LEFT JOIN listings l ON tc.listing_id = l.id
			ORDER BY tc.created_at DESC LIMIT 50`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var trades []gin.H
		for rows.Next() {
			var id, listingID, status, requestedBy, counterpart string
			var title sql.NullString
			var autoConfirm *time.Time
			var created time.Time
			if err := rows.Scan(&id, &listingID, &title, &status, &requestedBy, &counterpart, &autoConfirm, &created); err != nil {
				continue
			}
			t := gin.H{
				"completionId":       id,
				"listingId":          listingID,
				"listingTitle":       title.String,
				"status":             status,
				"requestedByUserId":  requestedBy,
				"counterpartUserId":  counterpart,
				"createdAt":          created.Format(time.RFC3339),
			}
			if autoConfirm != nil {
				t["autoConfirmAt"] = autoConfirm.Format(time.RFC3339)
			}
			trades = append(trades, t)
		}
		c.JSON(http.StatusOK, gin.H{"data": trades})
	}
}

func handleAdminListAllListings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")
		visibility := c.Query("visibility")

		query := `SELECT l.id, l.title, l.item_name, l.status, l.visibility, l.listing_type,
			up.nickname as author_nickname, l.created_at
			FROM listings l
			LEFT JOIN user_profiles up ON l.author_user_id = up.user_id
			WHERE l.deleted_at IS NULL`
		args := []interface{}{}
		argN := 1

		if status != "" {
			query += fmt.Sprintf(" AND l.status = $%d", argN)
			args = append(args, status)
			argN++
		}
		if visibility != "" {
			query += fmt.Sprintf(" AND l.visibility = $%d", argN)
			args = append(args, visibility)
			argN++
		}
		_ = argN
		query += " ORDER BY l.created_at DESC LIMIT 50"

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var listings []gin.H
		for rows.Next() {
			var id, title, itemName, st, vis, lt string
			var author sql.NullString
			var created time.Time
			if err := rows.Scan(&id, &title, &itemName, &st, &vis, &lt, &author, &created); err != nil {
				continue
			}
			listings = append(listings, gin.H{
				"listingId":      id,
				"title":          title,
				"itemName":       itemName,
				"status":         st,
				"visibility":     vis,
				"listingType":    lt,
				"authorNickname": author.String,
				"createdAt":      created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": listings})
	}
}

func handleAdminRestoreListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		adminID := middleware.GetUserID(c)

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("UPDATE listings SET visibility = 'public', updated_at = NOW() WHERE id = $1", listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 복원 실패"}})
			return
		}
		if _, err := tx.Exec(`INSERT INTO audit_logs (id, actor_id, actor_role, action, target_type, target_id)
			VALUES ($1, $2, 'admin', 'listing.moderation.restore', 'listing', $3)`,
			uuid.New().String(), adminID, listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "감사 로그 기록 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"listingId": listingID, "visibility": "public"})
	}
}

func handleAdminUpdateReportStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		reportID := c.Param("reportId")
		var req struct {
			Status string `json:"status" binding:"required,oneof=submitted assigned resolved"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}
		if _, err := db.Exec("UPDATE reports SET status = $1, updated_at = NOW() WHERE id = $2", req.Status, reportID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "신고 상태 업데이트 실패"}})
			return
		}
		c.JSON(http.StatusOK, gin.H{"reportId": reportID, "status": req.Status})
	}
}
