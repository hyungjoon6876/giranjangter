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

func TestCreateReport_Success_Returns201(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		CreateReportFn: func(ctx context.Context, params *repository.CreateReportParams) error {
			if params.ReporterID != "user-1" {
				t.Errorf("reporterID = %q, want %q", params.ReporterID, "user-1")
			}
			if params.TargetType != "user" {
				t.Errorf("targetType = %q, want %q", params.TargetType, "user")
			}
			if params.TargetID != "user-bad" {
				t.Errorf("targetID = %q, want %q", params.TargetID, "user-bad")
			}
			if params.ReportType != "harassment" {
				t.Errorf("reportType = %q, want %q", params.ReportType, "harassment")
			}
			if params.Description != "사기 의심됩니다" {
				t.Errorf("description = %q, want %q", params.Description, "사기 의심됩니다")
			}
			return nil
		},
	}

	r := setupRouter()
	r.POST("/api/v1/reports", authMiddleware("user-1", "user"), handleCreateReport(mockRepo))

	body := `{"targetType":"user","targetId":"user-bad","reportType":"harassment","description":"사기 의심됩니다"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reports", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	var resp struct {
		ReportID string `json:"reportId"`
		Status   string `json:"status"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.ReportID == "" {
		t.Error("expected non-empty reportId")
	}
	if resp.Status != "submitted" {
		t.Errorf("status = %q, want %q", resp.Status, "submitted")
	}
}

func TestCreateReport_MissingFields_Returns400(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{}

	r := setupRouter()
	r.POST("/api/v1/reports", authMiddleware("user-1", "user"), handleCreateReport(mockRepo))

	body := `{"targetType":"user"}` // missing required fields
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reports", strings.NewReader(body))
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

func TestCreateReport_InvalidTargetType_Returns400(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{}

	r := setupRouter()
	r.POST("/api/v1/reports", authMiddleware("user-1", "user"), handleCreateReport(mockRepo))

	body := `{"targetType":"invalid","targetId":"user-bad","reportType":"harassment","description":"사기 의심됩니다"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/reports", strings.NewReader(body))
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

func TestMyReports_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListMyReportsFn: func(ctx context.Context, userID string) ([]repository.MyReportItem, error) {
			return []repository.MyReportItem{
				{
					ReportID:   "report-1",
					TargetType: "user",
					TargetID:   "user-bad",
					ReportType: "harassment",
					Status:     "pending",
					CreatedAt:  time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
				},
				{
					ReportID:   "report-2",
					TargetType: "listing",
					TargetID:   "listing-fake",
					ReportType: "fake_listing",
					Status:     "resolved",
					CreatedAt:  time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/my/reports", authMiddleware("user-1", "user"), handleMyReports(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/my/reports", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			ReportID   string `json:"reportId"`
			TargetType string `json:"targetType"`
			TargetID   string `json:"targetId"`
			ReportType string `json:"reportType"`
			Status     string `json:"status"`
			CreatedAt  string `json:"createdAt"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 2 {
		t.Errorf("len(data) = %d, want 2", len(resp.Data))
	}
	if resp.Data[0].ReportID != "report-1" {
		t.Errorf("data[0].reportId = %q, want %q", resp.Data[0].ReportID, "report-1")
	}
	if resp.Data[0].TargetType != "user" {
		t.Errorf("data[0].targetType = %q, want %q", resp.Data[0].TargetType, "user")
	}
	if resp.Data[0].ReportType != "harassment" {
		t.Errorf("data[0].reportType = %q, want %q", resp.Data[0].ReportType, "harassment")
	}
	if resp.Data[0].Status != "pending" {
		t.Errorf("data[0].status = %q, want %q", resp.Data[0].Status, "pending")
	}
}

func TestMyReports_EmptyList_Returns200(t *testing.T) {
	mockRepo := &mock.MockReservationRepo{
		ListMyReportsFn: func(ctx context.Context, userID string) ([]repository.MyReportItem, error) {
			return []repository.MyReportItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/my/reports", authMiddleware("user-1", "user"), handleMyReports(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/my/reports", nil)
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