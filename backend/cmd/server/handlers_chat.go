package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/event"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func handleCreateChat(repo repository.ChatRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		authorID, err := repo.GetListingAuthor(ctx, listingID)
		if err != nil || authorID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		if authorID == userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "본인 매물에는 채팅을 시작할 수 없습니다."}})
			return
		}

		// Check existing chat
		existingID, err := repo.FindExistingChatRoom(ctx, listingID, authorID, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		if existingID != "" {
			c.JSON(http.StatusConflict, gin.H{"chatRoomId": existingID, "message": "이미 채팅방이 존재합니다."})
			return
		}

		chatID := uuid.New().String()
		now := time.Now().UTC()
		if err := repo.CreateChatRoom(ctx, &repository.CreateChatRoomParams{
			ID:        chatID,
			ListingID: listingID,
			SellerID:  authorID,
			BuyerID:   userID,
			Now:       now,
		}); err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "채팅방 생성에 실패했습니다."}})
			return
		}

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

func handleListChats(repo repository.ChatRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		items, err := repo.ListChatRooms(c.Request.Context(), userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var chats []gin.H
		for _, item := range items {
			chats = append(chats, gin.H{
				"chatRoomId": item.ChatRoomID, "listingId": item.ListingID, "listingTitle": item.ListingTitle,
				"chatStatus": item.ChatStatus,
				"counterparty": gin.H{"userId": item.CounterpartID, "nickname": item.CounterpartNick, "trustBadge": item.CounterpartBadge},
				"updatedAt": item.UpdatedAt,
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": chats})
	}
}

func handleListMessages(repo repository.ChatRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)
		cursor := c.Query("cursor")
		limitStr := c.DefaultQuery("limit", "50")
		ctx := c.Request.Context()

		limit, _ := strconv.Atoi(limitStr)
		if limit <= 0 || limit > 100 {
			limit = 50
		}

		// Verify participant
		isParticipant, err := repo.IsChatParticipant(ctx, chatID, userID)
		if err != nil || !isParticipant {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "채팅 참여자만 메시지를 조회할 수 있습니다."}})
			return
		}

		items, err := repo.ListMessages(ctx, chatID, limit+1, cursor)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var msgs []gin.H
		for _, item := range items {
			msgs = append(msgs, gin.H{
				"messageId": item.MessageID, "senderUserId": item.SenderUserID, "messageType": item.MessageType,
				"bodyText": item.BodyText, "metadataJson": item.MetadataJSON, "sentAt": item.SentAt.Format(time.RFC3339Nano),
			})
		}

		hasMore := len(msgs) > limit
		if hasMore {
			msgs = msgs[:limit]
		}

		var nextCursor *string
		if hasMore && len(msgs) > 0 {
			last := msgs[len(msgs)-1]["sentAt"].(string)
			nextCursor = &last
		}

		c.JSON(http.StatusOK, gin.H{
			"data": msgs,
			"cursor": gin.H{
				"next":    nextCursor,
				"hasMore": hasMore,
			},
		})
	}
}

func handleSendMessage(repo repository.ChatRepo, broker *event.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req struct {
			MessageType     string  `json:"messageType" binding:"required,oneof=text image"`
			BodyText        *string `json:"bodyText"`
			ClientMessageID *string `json:"clientMessageId"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		// Verify participant and get counterpart
		participants, err := repo.GetChatRoomParticipants(ctx, chatID, userID)
		if err != nil || participants == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "채팅 참여자만 메시지를 보낼 수 있습니다."}})
			return
		}

		// Dedup check
		if req.ClientMessageID != nil {
			isDup, _ := repo.CheckDuplicateMessage(ctx, *req.ClientMessageID)
			if isDup {
				c.JSON(http.StatusOK, gin.H{"message": "이미 전송된 메시지입니다."})
				return
			}
		}

		msgID := uuid.New().String()
		now := time.Now().UTC()
		if err := repo.InsertMessage(ctx, &repository.InsertMessageParams{
			ID:              msgID,
			ChatRoomID:      chatID,
			SenderUserID:    userID,
			MessageType:     req.MessageType,
			BodyText:        req.BodyText,
			ClientMessageID: req.ClientMessageID,
			Now:             now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "메시지 저장 실패"}})
			return
		}

		msgPayload := gin.H{
			"messageId":    msgID,
			"chatRoomId":   chatID,
			"senderUserId": userID,
			"messageType":  req.MessageType,
			"bodyText":     req.BodyText,
			"sentAt":       now.Format(time.RFC3339Nano),
		}

		// SSE broadcast to counterpart
		counterpart := participants.BuyerID
		if userID == participants.BuyerID {
			counterpart = participants.SellerID
		}
		broker.SendToUser(counterpart, event.SSEEvent{
			EventType: "new_message",
			Data:      msgPayload,
		})

		c.JSON(http.StatusCreated, msgPayload)
	}
}

func handleMarkRead(repo repository.ChatRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req struct {
			LastReadMessageID string `json:"lastReadMessageId" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		isParticipant, err := repo.IsChatParticipant(ctx, chatID, userID)
		if err != nil || !isParticipant {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "채팅 참여자만 읽음 처리할 수 있습니다."}})
			return
		}

		repo.UpsertReadCursor(ctx, chatID, userID, req.LastReadMessageID)
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

func handleBlockUser(repo repository.ChatRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetID := c.Param("userId")
		userID := middleware.GetUserID(c)
		repo.BlockUser(c.Request.Context(), uuid.New().String(), userID, targetID)
		c.Status(http.StatusNoContent)
	}
}

func handleUnblockUser(repo repository.ChatRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetID := c.Param("userId")
		userID := middleware.GetUserID(c)
		repo.UnblockUser(c.Request.Context(), userID, targetID)
		c.Status(http.StatusNoContent)
	}
}
