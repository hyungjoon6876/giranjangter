package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jym/lincle/internal/repository"
	"github.com/jym/lincle/internal/repository/mock"
)

func TestListNotifications_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListNotificationsFn: func(ctx context.Context, userID string) ([]repository.NotificationItem, error) {
			return []repository.NotificationItem{
				{
					NotificationID: "notif-1",
					Type:           "trade_completion",
					Title:          "거래 완료 요청",
					Body:           "거래가 완료되었습니다",
					ReferenceType:  strPtr("completion"),
					ReferenceID:    strPtr("completion-1"),
					DeepLink:       strPtr("/trades/completion-1"),
					IsRead:         false,
					CreatedAt:      time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
				},
				{
					NotificationID: "notif-2",
					Type:           "chat_message",
					Title:          "새 메시지",
					Body:           "안녕하세요!",
					ReferenceType:  strPtr("chat"),
					ReferenceID:    strPtr("chat-1"),
					DeepLink:       strPtr("/chats/chat-1"),
					IsRead:         true,
					CreatedAt:      time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/notifications", authMiddleware("user-1", "user"), handleListNotifications(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/notifications", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			NotificationID string  `json:"notificationId"`
			Type           string  `json:"type"`
			Title          string  `json:"title"`
			Body           string  `json:"body"`
			ReferenceType  *string `json:"referenceType"`
			ReferenceID    *string `json:"referenceId"`
			DeepLink       *string `json:"deepLink"`
			IsRead         bool    `json:"isRead"`
			CreatedAt      string  `json:"createdAt"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 2 {
		t.Errorf("len(data) = %d, want 2", len(resp.Data))
	}
	if resp.Data[0].NotificationID != "notif-1" {
		t.Errorf("data[0].notificationId = %q, want %q", resp.Data[0].NotificationID, "notif-1")
	}
	if resp.Data[0].Type != "trade_completion" {
		t.Errorf("data[0].type = %q, want %q", resp.Data[0].Type, "trade_completion")
	}
	if resp.Data[0].Title != "거래 완료 요청" {
		t.Errorf("data[0].title = %q, want %q", resp.Data[0].Title, "거래 완료 요청")
	}
	if resp.Data[0].IsRead {
		t.Error("data[0].isRead = true, want false")
	}
	if resp.Data[1].IsRead == false {
		t.Error("data[1].isRead = false, want true")
	}
}

func TestListNotifications_EmptyList_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListNotificationsFn: func(ctx context.Context, userID string) ([]repository.NotificationItem, error) {
			return []repository.NotificationItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/notifications", authMiddleware("user-1", "user"), handleListNotifications(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/notifications", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []interface{} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 0 {
		t.Errorf("len(data) = %d, want 0", len(resp.Data))
	}
}

func TestReadNotifications_Success_Returns204(t *testing.T) {
	var markedUserID string
	var markedIDs []string

	mockRepo := &mock.MockReservationRepo{
		MarkNotificationsReadFn: func(ctx context.Context, userID string, notificationIDs []string) error {
			markedUserID = userID
			markedIDs = notificationIDs
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/notifications/read", authMiddleware("user-1", "user"), handleReadNotifications(mockRepo))

	body := `{"notificationIds":["notif-1","notif-2"]}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/notifications/read", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusNoContent, w.Body.String())
	}

	if markedUserID != "user-1" {
		t.Errorf("markedUserID = %q, want %q", markedUserID, "user-1")
	}
	if len(markedIDs) != 2 {
		t.Errorf("len(markedIDs) = %d, want 2", len(markedIDs))
	}
	if markedIDs[0] != "notif-1" || markedIDs[1] != "notif-2" {
		t.Errorf("markedIDs = %v, want [notif-1 notif-2]", markedIDs)
	}
}

func TestReadNotifications_MissingIDs_Returns400(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{}

	r := setupRouter()
	r.POST("/api/v1/notifications/read", authMiddleware("user-1", "user"), handleReadNotifications(mockRepo))

	body := `{}` // missing notificationIds
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/notifications/read", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
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