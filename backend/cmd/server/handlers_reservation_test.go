package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jym/lincle/internal/repository"
	"github.com/jym/lincle/internal/repository/mock"
)

// ── CreateReservation ──

func TestCreateReservation_Success_Returns201(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetChatRoomForReservationFn: func(ctx context.Context, chatRoomID, userID string) (*repository.ChatRoomReservationInfo, error) {
			return &repository.ChatRoomReservationInfo{
				ListingID: "listing-1",
				SellerID:  "seller-1",
				BuyerID:   "buyer-1",
			}, nil
		},
		CountActiveReservationsFn: func(ctx context.Context, listingID string) (int, error) {
			return 0, nil // no active reservations
		},
		CreateReservationFn: func(ctx context.Context, params *repository.CreateReservationParams) error {
			if params.ProposerUserID != "buyer-1" {
				t.Errorf("proposer = %q, want %q", params.ProposerUserID, "buyer-1")
			}
			if params.CounterpartID != "seller-1" {
				t.Errorf("counterpart = %q, want %q", params.CounterpartID, "seller-1")
			}
			if params.ListingID != "listing-1" {
				t.Errorf("listingID = %q, want %q", params.ListingID, "listing-1")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/chats/:chatId/reservations", authMiddleware("buyer-1", "user"), handleCreateReservation(mockRepo))

	body := `{
		"scheduledAt":"2026-03-20T14:00:00Z",
		"meetingType":"in_game",
		"serverId":"server-1",
		"meetingPointText":"기란마을 분수대",
		"noteToCounterparty":"14시에 만나요"
	}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/chats/chat-1/reservations", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	var resp struct {
		ReservationID string `json:"reservationId"`
		Status        string `json:"status"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ReservationID == "" {
		t.Error("expected non-empty reservationId")
	}
	if resp.Status != "proposed" {
		t.Errorf("status = %q, want %q", resp.Status, "proposed")
	}
}

func TestCreateReservation_ActiveConflict_Returns409(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetChatRoomForReservationFn: func(ctx context.Context, chatRoomID, userID string) (*repository.ChatRoomReservationInfo, error) {
			return &repository.ChatRoomReservationInfo{
				ListingID: "listing-1",
				SellerID:  "seller-1",
				BuyerID:   "buyer-1",
			}, nil
		},
		CountActiveReservationsFn: func(ctx context.Context, listingID string) (int, error) {
			return 1, nil // active reservation exists
		},
	}

	r := setupRouter()
	r.POST("/api/v1/chats/:chatId/reservations", authMiddleware("buyer-1", "user"), handleCreateReservation(mockRepo))

	body := `{
		"scheduledAt":"2026-03-20T14:00:00Z",
		"meetingType":"in_game"
	}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/chats/chat-1/reservations", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusConflict, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "CONFLICT" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "CONFLICT")
	}
}

// ── ConfirmReservation ──

func TestConfirmReservation_Counterpart_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetReservationForConfirmFn: func(ctx context.Context, reservationID string) (*repository.ReservationConfirmInfo, error) {
			return &repository.ReservationConfirmInfo{
				CounterpartUserID: "seller-1",
				ListingID:         "listing-1",
				ChatRoomID:        "chat-1",
			}, nil
		},
		ConfirmReservationFn: func(ctx context.Context, params *repository.ConfirmReservationParams) error {
			if params.ReservationID != "res-1" {
				t.Errorf("reservationID = %q, want %q", params.ReservationID, "res-1")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/reservations/:resId/confirm", authMiddleware("seller-1", "user"), handleConfirmReservation(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reservations/res-1/confirm", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		ReservationID string `json:"reservationId"`
		Status        string `json:"status"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ReservationID != "res-1" {
		t.Errorf("reservationId = %q, want %q", resp.ReservationID, "res-1")
	}
	if resp.Status != "confirmed" {
		t.Errorf("status = %q, want %q", resp.Status, "confirmed")
	}
}

func TestConfirmReservation_Proposer_Returns403(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetReservationForConfirmFn: func(ctx context.Context, reservationID string) (*repository.ReservationConfirmInfo, error) {
			return &repository.ReservationConfirmInfo{
				CounterpartUserID: "seller-1", // only seller-1 can confirm
				ListingID:         "listing-1",
				ChatRoomID:        "chat-1",
			}, nil
		},
	}

	r := setupRouter()
	// Authenticate as buyer-1 (the proposer, NOT the counterpart)
	r.POST("/api/v1/reservations/:resId/confirm", authMiddleware("buyer-1", "user"), handleConfirmReservation(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reservations/res-1/confirm", nil)
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

// ── CancelReservation ──

func TestCancelReservation_Proposer_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetReservationForCancelFn: func(ctx context.Context, reservationID string) (*repository.ReservationCancelInfo, error) {
			return &repository.ReservationCancelInfo{
				ListingID:     "listing-1",
				ChatRoomID:    "chat-1",
				ProposerID:    "buyer-1",
				CounterpartID: "seller-1",
			}, nil
		},
		CancelReservationFn: func(ctx context.Context, params *repository.CancelReservationParams) error {
			if params.ReservationID != "res-1" {
				t.Errorf("reservationID = %q, want %q", params.ReservationID, "res-1")
			}
			if params.ReasonCode != "changed_mind" {
				t.Errorf("reasonCode = %q, want %q", params.ReasonCode, "changed_mind")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/reservations/:resId/cancel", authMiddleware("buyer-1", "user"), handleCancelReservation(mockRepo))

	body := `{"reasonCode":"changed_mind"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reservations/res-1/cancel", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		ReservationID string `json:"reservationId"`
		Status        string `json:"status"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ReservationID != "res-1" {
		t.Errorf("reservationId = %q, want %q", resp.ReservationID, "res-1")
	}
	if resp.Status != "cancelled" {
		t.Errorf("status = %q, want %q", resp.Status, "cancelled")
	}
}

func TestCancelReservation_NonParticipant_Returns403(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetReservationForCancelFn: func(ctx context.Context, reservationID string) (*repository.ReservationCancelInfo, error) {
			return &repository.ReservationCancelInfo{
				ListingID:     "listing-1",
				ChatRoomID:    "chat-1",
				ProposerID:    "buyer-1",
				CounterpartID: "seller-1",
			}, nil
		},
	}

	r := setupRouter()
	// Authenticate as user-other (neither proposer nor counterpart)
	r.POST("/api/v1/reservations/:resId/cancel", authMiddleware("user-other", "user"), handleCancelReservation(mockRepo))

	body := `{"reasonCode":"spam"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reservations/res-1/cancel", strings.NewReader(body))
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
