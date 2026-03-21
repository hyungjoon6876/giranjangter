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

func TestCompleteTrade_Success_Returns201(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetConfirmedReservationFn: func(ctx context.Context, reservationID, listingID string) (*repository.ReservationParticipants, error) {
			return &repository.ReservationParticipants{
				ProposerID:    "user-1",
				CounterpartID: "user-2",
			}, nil
		},
		CreateTradeCompletionFn: func(ctx context.Context, params *repository.CreateTradeCompletionParams) error {
			if params.ListingID != "listing-1" {
				t.Errorf("listingID = %q, want %q", params.ListingID, "listing-1")
			}
			if params.ReservationID != "reservation-1" {
				t.Errorf("reservationID = %q, want %q", params.ReservationID, "reservation-1")
			}
			if params.RequestedByID != "user-1" {
				t.Errorf("requestedByID = %q, want %q", params.RequestedByID, "user-1")
			}
			if params.CounterpartID != "user-2" {
				t.Errorf("counterpartID = %q, want %q", params.CounterpartID, "user-2")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/complete", authMiddleware("user-1", "user"), handleCompleteTrade(mockRepo))

	body := `{"reservationId":"reservation-1","completionNote":"거래 완료했습니다"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/complete", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	var resp struct {
		CompletionID     string `json:"completionId"`
		CompletionStatus string `json:"completionStatus"`
		ExpiresAt        string `json:"expiresAt"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.CompletionID == "" {
		t.Error("expected non-empty completionId")
	}
	if resp.CompletionStatus != "pending_confirm" {
		t.Errorf("completionStatus = %q, want %q", resp.CompletionStatus, "pending_confirm")
	}
	if resp.ExpiresAt == "" {
		t.Error("expected non-empty expiresAt")
	}
}

func TestCompleteTrade_NotFound_Returns404(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetConfirmedReservationFn: func(ctx context.Context, reservationID, listingID string) (*repository.ReservationParticipants, error) {
			return nil, nil // reservation not found
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/complete", authMiddleware("user-1", "user"), handleCompleteTrade(mockRepo))

	body := `{"reservationId":"nonexistent"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/complete", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
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

func TestCompleteTrade_NonParticipant_Returns403(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetConfirmedReservationFn: func(ctx context.Context, reservationID, listingID string) (*repository.ReservationParticipants, error) {
			return &repository.ReservationParticipants{
				ProposerID:    "user-1",
				CounterpartID: "user-2",
			}, nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/complete", authMiddleware("user-other", "user"), handleCompleteTrade(mockRepo))

	body := `{"reservationId":"reservation-1"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/complete", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
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

func TestCompleteTrade_MissingFields_Returns400(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/complete", authMiddleware("user-1", "user"), handleCompleteTrade(mockRepo))

	body := `{}` // missing reservationId
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/complete", strings.NewReader(body))
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

func TestConfirmCompletion_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetPendingCompletionFn: func(ctx context.Context, completionID string) (*repository.PendingCompletionInfo, error) {
			return &repository.PendingCompletionInfo{
				CounterpartUserID: "user-2",
				ListingID:         "listing-1",
				ReservationID:     "reservation-1",
				RequestedByUserID: "user-1",
			}, nil
		},
		ConfirmCompletionFn: func(ctx context.Context, params *repository.ConfirmCompletionParams) error {
			if params.CompletionID != "completion-1" {
				t.Errorf("completionID = %q, want %q", params.CompletionID, "completion-1")
			}
			if params.CounterpartID != "user-2" {
				t.Errorf("counterpartID = %q, want %q", params.CounterpartID, "user-2")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/confirm", authMiddleware("user-2", "user"), handleConfirmCompletion(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/completion-1/confirm", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		CompletionID     string `json:"completionId"`
		CompletionStatus string `json:"completionStatus"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.CompletionID != "completion-1" {
		t.Errorf("completionId = %q, want %q", resp.CompletionID, "completion-1")
	}
	if resp.CompletionStatus != "confirmed" {
		t.Errorf("completionStatus = %q, want %q", resp.CompletionStatus, "confirmed")
	}
}

func TestConfirmCompletion_NotFound_Returns404(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetPendingCompletionFn: func(ctx context.Context, completionID string) (*repository.PendingCompletionInfo, error) {
			return nil, nil // completion not found
		},
	}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/confirm", authMiddleware("user-2", "user"), handleConfirmCompletion(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/nonexistent/confirm", nil)
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

func TestConfirmCompletion_NonCounterpart_Returns403(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetPendingCompletionFn: func(ctx context.Context, completionID string) (*repository.PendingCompletionInfo, error) {
			return &repository.PendingCompletionInfo{
				CounterpartUserID: "user-2",
				ListingID:         "listing-1",
				ReservationID:     "reservation-1",
				RequestedByUserID: "user-1",
			}, nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/confirm", authMiddleware("user-other", "user"), handleConfirmCompletion(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/completion-1/confirm", nil)
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

func TestMyTrades_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListMyTradesFn: func(ctx context.Context, userID string) ([]repository.MyTradeItem, error) {
			return []repository.MyTradeItem{
				{
					ChatRoomID:      "chat-1",
					ListingID:       "listing-1",
					ListingTitle:    "아이템 판매",
					ListingStatus:   "sold",
					CounterpartID:   "user-2",
					CounterpartNick: "구매자",
					ChatStatus:      "closed",
					UpdatedAt:       time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
				},
				{
					ChatRoomID:      "chat-2",
					ListingID:       "listing-2",
					ListingTitle:    "다른 아이템",
					ListingStatus:   "completed",
					CounterpartID:   "user-3",
					CounterpartNick: "판매자",
					ChatStatus:      "closed",
					UpdatedAt:       time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/my/trades", authMiddleware("user-1", "user"), handleMyTrades(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/my/trades", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			ChatRoomID   string `json:"chatRoomId"`
			ListingID    string `json:"listingId"`
			ListingTitle string `json:"listingTitle"`
			TradeStatus  string `json:"tradeStatus"`
			ChatStatus   string `json:"chatStatus"`
			Counterparty struct {
				UserID   string `json:"userId"`
				Nickname string `json:"nickname"`
			} `json:"counterparty"`
			UpdatedAt string `json:"updatedAt"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 2 {
		t.Errorf("len(data) = %d, want 2", len(resp.Data))
	}
	if resp.Data[0].ChatRoomID != "chat-1" {
		t.Errorf("data[0].chatRoomId = %q, want %q", resp.Data[0].ChatRoomID, "chat-1")
	}
	if resp.Data[0].ListingTitle != "아이템 판매" {
		t.Errorf("data[0].listingTitle = %q, want %q", resp.Data[0].ListingTitle, "아이템 판매")
	}
	if resp.Data[0].TradeStatus != "sold" {
		t.Errorf("data[0].tradeStatus = %q, want %q", resp.Data[0].TradeStatus, "sold")
	}
	if resp.Data[0].Counterparty.UserID != "user-2" {
		t.Errorf("data[0].counterparty.userId = %q, want %q", resp.Data[0].Counterparty.UserID, "user-2")
	}
}

func TestMyTrades_EmptyList_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListMyTradesFn: func(ctx context.Context, userID string) ([]repository.MyTradeItem, error) {
			return []repository.MyTradeItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/my/trades", authMiddleware("user-1", "user"), handleMyTrades(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/my/trades", nil)
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