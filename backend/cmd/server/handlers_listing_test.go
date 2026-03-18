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

// ── ChangeListingStatus ──

func TestChangeListingStatus_InvalidTransition_Returns422(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1",
				Status:       "available",
			}, nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/status", authMiddleware("user-1", "user"), handleChangeListingStatus(mockRepo))

	// available -> completed is not allowed (must go through reserved -> pending_trade)
	body := `{"action":"complete"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/status", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusUnprocessableEntity, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "INVALID_TRANSITION" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "INVALID_TRANSITION")
	}
}

func TestChangeListingStatus_ValidTransition_Returns200(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1",
				Status:       "available",
			}, nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/status", authMiddleware("user-1", "user"), handleChangeListingStatus(mockRepo))

	// available -> reserved is allowed
	body := `{"action":"reserve"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/status", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		ListingID string `json:"listingId"`
		Status    string `json:"status"`
		UpdatedAt string `json:"updatedAt"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ListingID != "listing-1" {
		t.Errorf("listingId = %q, want %q", resp.ListingID, "listing-1")
	}
	if resp.Status != "reserved" {
		t.Errorf("status = %q, want %q", resp.Status, "reserved")
	}
	if resp.UpdatedAt == "" {
		t.Error("expected non-empty updatedAt")
	}
}

func TestChangeListingStatus_NonOwner_Returns403(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1", // owned by user-1
				Status:       "available",
			}, nil
		},
	}

	r := setupRouter()
	// Authenticate as user-2 (not the owner)
	r.POST("/api/v1/listings/:id/status", authMiddleware("user-2", "user"), handleChangeListingStatus(mockRepo))

	body := `{"action":"reserve"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/status", strings.NewReader(body))
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

// ── UpdateListing ──

func TestUpdateListing_CompletedStatus_Returns403(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1",
				Status:       "completed",
			}, nil
		},
	}

	r := setupRouter()
	r.PATCH("/api/v1/listings/:id", authMiddleware("user-1", "user"), handleUpdateListing(mockRepo))

	body := `{"title":"new title"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/api/v1/listings/listing-1", strings.NewReader(body))
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

func TestUpdateListing_CancelledStatus_Returns403(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1",
				Status:       "cancelled",
			}, nil
		},
	}

	r := setupRouter()
	r.PATCH("/api/v1/listings/:id", authMiddleware("user-1", "user"), handleUpdateListing(mockRepo))

	body := `{"title":"new title"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/api/v1/listings/listing-1", strings.NewReader(body))
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

func TestUpdateListing_NonOwner_Returns403(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1", // owned by user-1
				Status:       "available",
			}, nil
		},
	}

	r := setupRouter()
	// Authenticate as user-2 (not the owner)
	r.PATCH("/api/v1/listings/:id", authMiddleware("user-2", "user"), handleUpdateListing(mockRepo))

	body := `{"title":"new title"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/api/v1/listings/listing-1", strings.NewReader(body))
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

func TestUpdateListing_NoFields_Returns400(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1",
				Status:       "available",
			}, nil
		},
	}

	r := setupRouter()
	r.PATCH("/api/v1/listings/:id", authMiddleware("user-1", "user"), handleUpdateListing(mockRepo))

	body := `{}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/api/v1/listings/listing-1", strings.NewReader(body))
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

func TestChangeListingStatus_InvalidAction_Returns400(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return &repository.ListingOwnerStatus{
				AuthorUserID: "user-1",
				Status:       "available",
			}, nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/status", authMiddleware("user-1", "user"), handleChangeListingStatus(mockRepo))

	body := `{"action":"explode"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/status", strings.NewReader(body))
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

func TestChangeListingStatus_DeletedListing_Returns404(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingOwnerAndStatusFn: func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
			return nil, nil // listing not found
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/status", authMiddleware("user-1", "user"), handleChangeListingStatus(mockRepo))

	body := `{"action":"reserve"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/status", strings.NewReader(body))
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

// ── CreateListing ──

func TestCreateListing_MissingFields_Returns400(t *testing.T) {
	mockRepo := &mock.MockListingRepo{}

	r := setupRouter()
	r.POST("/api/v1/listings", authMiddleware("user-1", "user"), handleCreateListing(mockRepo))

	// Missing required fields (title, itemName, etc.)
	body := `{"listingType":"sell"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings", strings.NewReader(body))
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

func TestCreateListing_PriceRequired_ForNonOffer(t *testing.T) {
	mockRepo := &mock.MockListingRepo{}

	r := setupRouter()
	r.POST("/api/v1/listings", authMiddleware("user-1", "user"), handleCreateListing(mockRepo))

	// priceType=fixed but no priceAmount
	body := `{
		"listingType":"sell",
		"serverId":"server-1",
		"categoryId":"cat-1",
		"itemName":"진명황의 집행검",
		"title":"팝니다",
		"description":"좋은 아이템입니다 상태 좋아요 빠르게 거래 원합니다",
		"priceType":"fixed",
		"quantity":1,
		"tradeMethod":"in_game"
	}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings", strings.NewReader(body))
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
	if !strings.Contains(resp.Error.Message, "가격") {
		t.Errorf("error.message = %q, want to contain '가격'", resp.Error.Message)
	}
}
