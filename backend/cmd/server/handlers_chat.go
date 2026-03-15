package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/event"
	"github.com/jym/lincle/internal/middleware"
)

func handleCreateChat(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		userID := middleware.GetUserID(c)

		var authorID string
		err := db.QueryRow("SELECT author_user_id FROM listings WHERE id = $1 AND deleted_at IS NULL", listingID).Scan(&authorID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		if authorID == userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "본인 매물에는 채팅을 시작할 수 없습니다."}})
			return
		}

		// Check existing chat
		var existingID string
		err = db.QueryRow(
			"SELECT id FROM chat_rooms WHERE listing_id = $1 AND ((seller_user_id = $2 AND buyer_user_id = $3) OR (seller_user_id = $4 AND buyer_user_id = $5))",
			listingID, authorID, userID, userID, authorID,
		).Scan(&existingID)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"chatRoomId": existingID, "message": "이미 채팅방이 존재합니다."})
			return
		}

		chatID := uuid.New().String()
		now := time.Now().UTC()
		_, err = db.Exec(
			"INSERT INTO chat_rooms (id, listing_id, seller_user_id, buyer_user_id, chat_status, created_at, updated_at) VALUES ($1, $2, $3, $4, 'open', $5, $6)",
			chatID, listingID, authorID, userID, now, now,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
			return
		}

		db.Exec("UPDATE listings SET chat_count = chat_count + 1 WHERE id = $1", listingID)

		c.JSON(http.StatusCreated, gin.H{
			"chatRoomId":   chatID,
			"listingId":    listingID,
			"sellerUserId": authorID,
			"buyerUserId":  userID,
			"chatStatus":   "open",
			"createdAt":    now.Format(time.RFC3339),
		})
	}
}

func handleListChats(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		rows, err := db.Query(`
			SELECT cr.id, cr.listing_id, l.title, cr.chat_status, cr.last_message_at, cr.updated_at,
				CASE WHEN cr.seller_user_id = $1 THEN cr.buyer_user_id ELSE cr.seller_user_id END as counterpart_id,
				p.nickname, p.trust_badge
			FROM chat_rooms cr
			JOIN listings l ON cr.listing_id = l.id
			JOIN user_profiles p ON p.user_id = CASE WHEN cr.seller_user_id = $2 THEN cr.buyer_user_id ELSE cr.seller_user_id END
			WHERE cr.seller_user_id = $3 OR cr.buyer_user_id = $4
			ORDER BY COALESCE(cr.last_message_at, cr.created_at) DESC
			LIMIT 50`, userID, userID, userID, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
			return
		}
		defer rows.Close()

		var chats []gin.H
		for rows.Next() {
			var chatID, listingID, title, status, cpID, cpNick, cpBadge string
			var lastMsg, updated *time.Time
			if err := rows.Scan(&chatID, &listingID, &title, &status, &lastMsg, &updated, &cpID, &cpNick, &cpBadge); err != nil {
				continue
			}
			chats = append(chats, gin.H{
				"chatRoomId": chatID, "listingId": listingID, "listingTitle": title,
				"chatStatus": status,
				"counterparty": gin.H{"userId": cpID, "nickname": cpNick, "trustBadge": cpBadge},
				"updatedAt": updated,
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": chats})
	}
}

func handleListMessages(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)

		// Verify participant
		var count int
		db.QueryRow("SELECT COUNT(*) FROM chat_rooms WHERE id = $1 AND (seller_user_id = $2 OR buyer_user_id = $3)", chatID, userID, userID).Scan(&count)
		if count == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "채팅 참여자만 메시지를 조회할 수 있습니다."}})
			return
		}

		rows, err := db.Query(
			"SELECT id, sender_user_id, message_type, body_text, metadata_json, sent_at FROM chat_messages WHERE chat_room_id = $1 AND deleted_at IS NULL ORDER BY sent_at DESC LIMIT 50",
			chatID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
			return
		}
		defer rows.Close()

		var msgs []gin.H
		for rows.Next() {
			var id string
			var sender *string
			var msgType string
			var body, meta *string
			var sentAt time.Time
			if err := rows.Scan(&id, &sender, &msgType, &body, &meta, &sentAt); err != nil {
				continue
			}
			msgs = append(msgs, gin.H{
				"messageId": id, "senderUserId": sender, "messageType": msgType,
				"bodyText": body, "metadataJson": meta, "sentAt": sentAt.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": msgs})
	}
}

func handleSendMessage(db *sql.DB, broker *event.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)

		var req struct {
			MessageType     string  `json:"messageType" binding:"required"`
			BodyText        *string `json:"bodyText"`
			ClientMessageID *string `json:"clientMessageId"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		// Verify participant and get counterpart
		var sellerID, buyerID string
		err := db.QueryRow("SELECT seller_user_id, buyer_user_id FROM chat_rooms WHERE id = $1 AND (seller_user_id = $2 OR buyer_user_id = $3)", chatID, userID, userID).Scan(&sellerID, &buyerID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "채팅 참여자만 메시지를 보낼 수 있습니다."}})
			return
		}

		// Dedup check
		if req.ClientMessageID != nil {
			var existing int
			db.QueryRow("SELECT COUNT(*) FROM chat_messages WHERE client_message_id = $1", *req.ClientMessageID).Scan(&existing)
			if existing > 0 {
				c.JSON(http.StatusOK, gin.H{"message": "이미 전송된 메시지입니다."})
				return
			}
		}

		msgID := uuid.New().String()
		now := time.Now().UTC()
		if _, err := db.Exec(
			"INSERT INTO chat_messages (id, chat_room_id, sender_user_id, message_type, body_text, client_message_id, sent_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			msgID, chatID, userID, req.MessageType, req.BodyText, req.ClientMessageID, now,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "메시지 저장 실패"}})
			return
		}
		db.Exec("UPDATE chat_rooms SET last_message_at = $1, updated_at = $2 WHERE id = $3", now, now, chatID)

		msgPayload := gin.H{
			"messageId":    msgID,
			"chatRoomId":   chatID,
			"senderUserId": userID,
			"messageType":  req.MessageType,
			"bodyText":     req.BodyText,
			"sentAt":       now.Format(time.RFC3339),
		}

		// SSE broadcast to counterpart
		counterpart := buyerID
		if userID == buyerID {
			counterpart = sellerID
		}
		broker.SendToUser(counterpart, event.SSEEvent{
			EventType: "new_message",
			Data:      msgPayload,
		})

		c.JSON(http.StatusCreated, msgPayload)
	}
}

func handleMarkRead(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)
		var req struct {
			LastReadMessageID string `json:"lastReadMessageId" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}
		db.Exec(`INSERT INTO chat_read_cursors (chat_room_id, user_id, last_read_message_id, updated_at)
			VALUES ($1, $2, $3, NOW())
			ON CONFLICT(chat_room_id, user_id) DO UPDATE SET last_read_message_id = $4, updated_at = NOW()`,
			chatID, userID, req.LastReadMessageID, req.LastReadMessageID)
		c.Status(http.StatusNoContent)
	}
}

func handleSSEConnect(broker *event.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("X-Accel-Buffering", "no")

		events, cleanup := broker.Subscribe(userID)
		defer cleanup()

		c.SSEvent("connected", gin.H{"status": "connected", "userId": userID})
		c.Writer.Flush()

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		clientGone := c.Request.Context().Done()
		for {
			select {
			case evt, ok := <-events:
				if !ok {
					return
				}
				c.SSEvent(evt.EventType, evt.Data)
				c.Writer.Flush()
			case <-ticker.C:
				c.SSEvent("heartbeat", gin.H{"time": time.Now().UTC().Format(time.RFC3339)})
				c.Writer.Flush()
			case <-clientGone:
				return
			}
		}
	}
}

func handleBlockUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetID := c.Param("userId")
		userID := middleware.GetUserID(c)
		db.Exec("INSERT INTO block_relations (id, blocker_user_id, blocked_user_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
			uuid.New().String(), userID, targetID)
		c.Status(http.StatusNoContent)
	}
}

func handleUnblockUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetID := c.Param("userId")
		userID := middleware.GetUserID(c)
		db.Exec("DELETE FROM block_relations WHERE blocker_user_id = $1 AND blocked_user_id = $2", userID, targetID)
		c.Status(http.StatusNoContent)
	}
}
