package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/event"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
	"github.com/jym/lincle/internal/repository/mock"
)

// authMiddleware injects userID and userRole into the Gin context (simulates RequireAuth).
func authMiddleware(userID, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userId", userID)
		c.Set("userRole", role)
		c.Next()
	}
}

func TestCreateChat_Returns201_NewChatRoom(t *testing.T) {
	mockRepo := &mock.MockChatRepo{
		GetListingAuthorFn: func(ctx context.Context, listingID string) (string, error) {
			return "seller-1", nil
		},
		FindExistingChatRoomFn: func(ctx context.Context, listingID, sellerID, buyerID string) (string, error) {
			return "", nil // no existing chat
		},
		CreateChatRoomFn: func(ctx context.Context, params *repository.CreateChatRoomParams) error {
			if params.SellerID != "seller-1" {
				t.Errorf("sellerID = %q, want %q", params.SellerID, "seller-1")
			}
			if params.BuyerID != "buyer-1" {
				t.Errorf("buyerID = %q, want %q", params.BuyerID, "buyer-1")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/chats", authMiddleware("buyer-1", "user"), handleCreateChat(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/chats", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	var resp struct {
		ChatRoomID   string `json:"chatRoomId"`
		ListingID    string `json:"listingId"`
		SellerUserID string `json:"sellerUserId"`
		BuyerUserID  string `json:"buyerUserId"`
		ChatStatus   string `json:"chatStatus"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ChatRoomID == "" {
		t.Error("expected non-empty chatRoomId")
	}
	if resp.ListingID != "listing-1" {
		t.Errorf("listingId = %q, want %q", resp.ListingID, "listing-1")
	}
	if resp.SellerUserID != "seller-1" {
		t.Errorf("sellerUserId = %q, want %q", resp.SellerUserID, "seller-1")
	}
	if resp.BuyerUserID != "buyer-1" {
		t.Errorf("buyerUserId = %q, want %q", resp.BuyerUserID, "buyer-1")
	}
	if resp.ChatStatus != "open" {
		t.Errorf("chatStatus = %q, want %q", resp.ChatStatus, "open")
	}
}

func TestCreateChat_Returns409_ExistingChatRoom(t *testing.T) {
	mockRepo := &mock.MockChatRepo{
		GetListingAuthorFn: func(ctx context.Context, listingID string) (string, error) {
			return "seller-1", nil
		},
		FindExistingChatRoomFn: func(ctx context.Context, listingID, sellerID, buyerID string) (string, error) {
			return "existing-chat-id", nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/chats", authMiddleware("buyer-1", "user"), handleCreateChat(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/chats", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusConflict, w.Body.String())
	}

	var resp struct {
		ChatRoomID string `json:"chatRoomId"`
		Message    string `json:"message"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ChatRoomID != "existing-chat-id" {
		t.Errorf("chatRoomId = %q, want %q", resp.ChatRoomID, "existing-chat-id")
	}
}

func TestCreateChat_Returns401_Unauthorized(t *testing.T) {
	auth := newTestAuth()
	mockRepo := &mock.MockChatRepo{}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/chats", auth.RequireAuth(), handleCreateChat(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/chats", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusUnauthorized, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "UNAUTHORIZED" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "UNAUTHORIZED")
	}
}

func TestCreateChat_Returns400_OwnListing(t *testing.T) {
	mockRepo := &mock.MockChatRepo{
		GetListingAuthorFn: func(ctx context.Context, listingID string) (string, error) {
			return "seller-1", nil // listing owned by seller-1
		},
	}

	r := setupRouter()
	// Authenticate as "seller-1" (the owner)
	r.POST("/api/v1/listings/:id/chats", authMiddleware("seller-1", "user"), handleCreateChat(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/chats", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusBadRequest, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "VALIDATION_ERROR" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "VALIDATION_ERROR")
	}
}

func TestCreateChat_Returns404_DeletedListing(t *testing.T) {
	mockRepo := &mock.MockChatRepo{
		GetListingAuthorFn: func(ctx context.Context, listingID string) (string, error) {
			return "", nil // listing not found (soft-deleted)
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/chats", authMiddleware("buyer-1", "user"), handleCreateChat(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/chats", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusNotFound, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "NOT_FOUND" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "NOT_FOUND")
	}
}

func TestListMessages_Returns403_NonParticipant(t *testing.T) {
	mockRepo := &mock.MockChatRepo{
		IsChatParticipantFn: func(ctx context.Context, chatRoomID, userID string) (bool, error) {
			return false, nil // not a participant
		},
	}

	r := setupRouter()
	r.GET("/api/v1/chats/:chatId/messages", authMiddleware("user-other", "user"), handleListMessages(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/chats/chat-1/messages", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusForbidden, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "FORBIDDEN" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "FORBIDDEN")
	}
}

func TestSendMessage_DedupByClientMessageID(t *testing.T) {
	mockRepo := &mock.MockChatRepo{
		GetChatRoomParticipantsFn: func(ctx context.Context, chatRoomID, userID string) (*repository.ChatParticipants, error) {
			return &repository.ChatParticipants{
				SellerID: "seller-1",
				BuyerID:  "buyer-1",
			}, nil
		},
		CheckDuplicateMessageFn: func(ctx context.Context, clientMessageID string) (bool, error) {
			if clientMessageID == "dedup-1" {
				return true, nil // duplicate found
			}
			return false, nil
		},
	}

	broker := event.NewBroker()
	r := setupRouter()
	r.POST("/api/v1/chats/:chatId/messages",
		authMiddleware("buyer-1", "user"),
		handleSendMessage(mockRepo, broker),
	)

	body := `{"messageType":"text","bodyText":"hello","clientMessageId":"dedup-1"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/chats/chat-1/messages", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Dedup returns 200 (not 201)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Message string `json:"message"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Message == "" {
		t.Error("expected dedup message in response")
	}

	// Ensure the middleware helper is in scope — suppress unused import
	_ = middleware.GetUserID
}
