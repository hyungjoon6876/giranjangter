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

// ── ListListings ──

func TestListListings_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		ListListingsFn: func(ctx context.Context, filter repository.ListingFilter) ([]repository.ListingListItem, error) {
			return []repository.ListingListItem{
				{
					ListingID:       "listing-1",
					ListingType:     "sell",
					Title:           "아이템 판매",
					ItemName:        "진명황의 집행검",
					PriceType:       "fixed",
					PriceAmount:     int64Ptr(50000),
					EnhancementLvl:  intPtr(5),
					ServerID:        "server-1",
					ServerName:      "데포로쥬",
					Status:          "available",
					TradeMethod:     "in_game",
					ViewCount:       10,
					FavoriteCount:   2,
					ChatCount:       3,
					AuthorID:        "user-1",
					AuthorNickname:  "판매자",
					TrustBadge:      "trusted",
					ResponseBadge:   "fast",
					IconID:          strPtr("sword_icon"),
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/listings", handleListListings(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/listings?q=진명황&serverId=server-1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []map[string]interface{} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 1 {
		t.Errorf("len(data) = %d, want 1", len(resp.Data))
	}
	if resp.Data[0]["listingId"] != "listing-1" {
		t.Errorf("listingId = %v, want %q", resp.Data[0]["listingId"], "listing-1")
	}
}

func TestListListings_EmptyResults_Returns200(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		ListListingsFn: func(ctx context.Context, filter repository.ListingFilter) ([]repository.ListingListItem, error) {
			return []repository.ListingListItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/listings", handleListListings(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/listings", nil)
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

// ── GetListing ──

func TestGetListing_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingFn: func(ctx context.Context, listingID string) (*repository.ListingDetail, error) {
			return &repository.ListingDetail{
				ID:              "listing-1",
				ListingType:     "sell",
				Title:           "아이템 판매",
				ItemName:        "진명황의 집행검",
				Description:     "좋은 아이템입니다",
				PriceType:       "fixed",
				PriceAmount:     int64Ptr(50000),
				Quantity:        1,
				Enhancement:     intPtr(5),
				ServerID:        "server-1",
				ServerName:      "데포로쥬",
				CategoryID:      "cat-1",
				CategoryName:    "무기",
				Status:          "available",
				TradeMethod:     "in_game",
				AuthorID:        "user-1",
				AuthorNickname:  "판매자",
				TrustBadge:      "trusted",
				ResponseBadge:   "fast",
				ViewCount:       10,
				FavoriteCount:   2,
				ChatCount:       3,
			}, nil
		},
		IncrementViewCountFn: func(ctx context.Context, listingID string) error {
			return nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/listings/:id", handleGetListing(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/listings/listing-1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp["listingId"] != "listing-1" {
		t.Errorf("listingId = %v, want %q", resp["listingId"], "listing-1")
	}
	if resp["title"] != "아이템 판매" {
		t.Errorf("title = %v, want %q", resp["title"], "아이템 판매")
	}
}

func TestGetListing_NotFound_Returns404(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		GetListingFn: func(ctx context.Context, listingID string) (*repository.ListingDetail, error) {
			return nil, nil // listing not found
		},
	}

	r := setupRouter()
	r.GET("/api/v1/listings/:id", handleGetListing(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/listings/nonexistent", nil)
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

// ── FavoriteListing ──

func TestFavoriteListing_Success_Returns204(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		ListingExistsFn: func(ctx context.Context, listingID string) (bool, error) {
			return true, nil
		},
		AddFavoriteFn: func(ctx context.Context, id, userID, listingID string) error {
			if userID != "user-1" {
				t.Errorf("userID = %q, want %q", userID, "user-1")
			}
			if listingID != "listing-1" {
				t.Errorf("listingID = %q, want %q", listingID, "listing-1")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/favorite", authMiddleware("user-1", "user"), handleFavoriteListing(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/listing-1/favorite", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusNoContent, w.Body.String())
	}
}

func TestFavoriteListing_NotFound_Returns404(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		ListingExistsFn: func(ctx context.Context, listingID string) (bool, error) {
			return false, nil // listing doesn't exist
		},
	}

	r := setupRouter()
	r.POST("/api/v1/listings/:id/favorite", authMiddleware("user-1", "user"), handleFavoriteListing(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/listings/nonexistent/favorite", nil)
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

// ── UnfavoriteListing ──

func TestUnfavoriteListing_Success_Returns204(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		RemoveFavoriteFn: func(ctx context.Context, userID, listingID string) error {
			if userID != "user-1" {
				t.Errorf("userID = %q, want %q", userID, "user-1")
			}
			if listingID != "listing-1" {
				t.Errorf("listingID = %q, want %q", listingID, "listing-1")
			}
			return nil
		},
	}

	r := setupRouter()
	r.DELETE("/api/v1/listings/:id/favorite", authMiddleware("user-1", "user"), handleUnfavoriteListing(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/api/v1/listings/listing-1/favorite", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusNoContent, w.Body.String())
	}
}

// ── MyListings ──

func TestMyListings_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		ListMyListingsFn: func(ctx context.Context, userID string, status *string) ([]repository.MyListingItem, error) {
			return []repository.MyListingItem{
				{
					ListingID:     "listing-1",
					ListingType:   "sell",
					Title:         "아이템 판매",
					ItemName:      "진명황의 집행검",
					PriceType:     "fixed",
					PriceAmount:   int64Ptr(50000),
					Status:        "available",
					ViewCount:     10,
					FavoriteCount: 2,
					ChatCount:     3,
				},
				{
					ListingID:     "listing-2",
					ListingType:   "buy",
					Title:         "아이템 구매",
					ItemName:      "다른 아이템",
					PriceType:     "offer",
					PriceAmount:   nil,
					Status:        "sold",
					ViewCount:     5,
					FavoriteCount: 1,
					ChatCount:     2,
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/my/listings", authMiddleware("user-1", "user"), handleMyListings(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/my/listings?status=available", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []map[string]interface{} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 2 {
		t.Errorf("len(data) = %d, want 2", len(resp.Data))
	}
	if resp.Data[0]["listingId"] != "listing-1" {
		t.Errorf("listingId = %v, want %q", resp.Data[0]["listingId"], "listing-1")
	}
	if resp.Data[0]["title"] != "아이템 판매" {
		t.Errorf("title = %v, want %q", resp.Data[0]["title"], "아이템 판매")
	}
}

func TestMyListings_EmptyList_Returns200(t *testing.T) {
	mockRepo := &mock.MockListingRepo{
		ListMyListingsFn: func(ctx context.Context, userID string, status *string) ([]repository.MyListingItem, error) {
			return []repository.MyListingItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/my/listings", authMiddleware("user-1", "user"), handleMyListings(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/my/listings", nil)
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

// Helper functions
func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}
