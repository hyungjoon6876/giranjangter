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

func TestCreateReview_Success_Returns201(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetCompletionForReviewFn: func(ctx context.Context, completionID string) (*repository.CompletionReviewInfo, error) {
			return &repository.CompletionReviewInfo{
				Status:            "confirmed",
				RequestedByUserID: "user-1",
				CounterpartUserID: "user-2",
			}, nil
		},
		CreateReviewFn: func(ctx context.Context, params *repository.CreateReviewParams) error {
			if params.CompletionID != "completion-1" {
				t.Errorf("completionID = %q, want %q", params.CompletionID, "completion-1")
			}
			if params.ReviewerID != "user-1" {
				t.Errorf("reviewerID = %q, want %q", params.ReviewerID, "user-1")
			}
			if params.TargetUserID != "user-2" {
				t.Errorf("targetUserID = %q, want %q", params.TargetUserID, "user-2")
			}
			if params.Rating != "positive" {
				t.Errorf("rating = %q, want %q", params.Rating, "positive")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/reviews", authMiddleware("user-1", "user"), handleCreateReview(mockRepo))

	body := `{"rating":"positive","comment":"좋은 거래였습니다"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/completion-1/reviews", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	var resp struct {
		ReviewID string `json:"reviewId"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ReviewID == "" {
		t.Error("expected non-empty reviewId")
	}
}

func TestCreateReview_NotConfirmed_Returns400(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetCompletionForReviewFn: func(ctx context.Context, completionID string) (*repository.CompletionReviewInfo, error) {
			return &repository.CompletionReviewInfo{
				Status:            "pending_confirm", // not confirmed
				RequestedByUserID: "user-1",
				CounterpartUserID: "user-2",
			}, nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/reviews", authMiddleware("user-1", "user"), handleCreateReview(mockRepo))

	body := `{"rating":"positive","comment":"좋은 거래였습니다"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/completion-1/reviews", strings.NewReader(body))
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

func TestCreateReview_DuplicateReview_Returns409(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		GetCompletionForReviewFn: func(ctx context.Context, completionID string) (*repository.CompletionReviewInfo, error) {
			return &repository.CompletionReviewInfo{
				Status:            "confirmed",
				RequestedByUserID: "user-1",
				CounterpartUserID: "user-2",
			}, nil
		},
		CreateReviewFn: func(ctx context.Context, params *repository.CreateReviewParams) error {
			// Simulate constraint violation error
			return &mockConstraintError{}
		},
	}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/reviews", authMiddleware("user-1", "user"), handleCreateReview(mockRepo))

	body := `{"rating":"positive","comment":"좋은 거래였습니다"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/completion-1/reviews", strings.NewReader(body))
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

func TestCreateReview_MissingFields_Returns400(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{}

	r := setupRouter()
	r.POST("/api/v1/completions/:compId/reviews", authMiddleware("user-1", "user"), handleCreateReview(mockRepo))

	body := `{}` // missing required fields
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/completions/completion-1/reviews", strings.NewReader(body))
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

func TestGetUserReviews_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListUserReviewsFn: func(ctx context.Context, targetUserID string) ([]repository.UserReviewItem, error) {
			return []repository.UserReviewItem{
				{
					ReviewID:         "review-1",
					Rating:           "positive",
					Comment:          strPtr("좋은 거래였습니다"),
					CreatedAt:        time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
					ReviewerNickname: "구매자1",
				},
				{
					ReviewID:         "review-2",
					Rating:           "negative",
					Comment:          strPtr("응답이 늦었습니다"),
					CreatedAt:        time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
					ReviewerNickname: "구매자2",
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/users/:userId/reviews", handleGetUserReviews(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/users/user-1/reviews", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			ReviewID         string  `json:"reviewId"`
			Rating           string  `json:"rating"`
			Comment          *string `json:"comment"`
			ReviewerNickname string  `json:"reviewerNickname"`
			CreatedAt        string  `json:"createdAt"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 2 {
		t.Errorf("len(data) = %d, want 2", len(resp.Data))
	}
	if resp.Data[0].ReviewID != "review-1" {
		t.Errorf("data[0].reviewId = %q, want %q", resp.Data[0].ReviewID, "review-1")
	}
	if resp.Data[0].Rating != "positive" {
		t.Errorf("data[0].rating = %q, want %q", resp.Data[0].Rating, "positive")
	}
	if resp.Data[0].Comment == nil || *resp.Data[0].Comment != "좋은 거래였습니다" {
		t.Errorf("data[0].comment = %v, want %q", resp.Data[0].Comment, "좋은 거래였습니다")
	}
}

func TestGetUserReviews_EmptyList_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListUserReviewsFn: func(ctx context.Context, targetUserID string) ([]repository.UserReviewItem, error) {
			return []repository.UserReviewItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/users/:userId/reviews", handleGetUserReviews(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/users/user-1/reviews", nil)
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

// Helper types
type mockConstraintError struct{}

func (e *mockConstraintError) Error() string {
	return "constraint violation"
}