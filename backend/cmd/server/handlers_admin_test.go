package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ---------- ShouldBindJSON validation tests (no DB needed) ----------

func TestAdminReportAction_MissingFields_Returns400(t *testing.T) {
	r := setupRouter()
	// Pass nil db — handler returns 400 before any DB call
	r.POST("/api/v1/admin/reports/:reportId/actions", authMiddleware("admin-1", "admin"), handleAdminReportAction(nil))

	body := `{}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/admin/reports/r1/actions", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAdminReportAction_InvalidActionCode_Returns400(t *testing.T) {
	r := setupRouter()
	r.POST("/api/v1/admin/reports/:reportId/actions", authMiddleware("admin-1", "admin"), handleAdminReportAction(nil))

	body := `{"actionCode":"invalid_code","targetUserId":"u1"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/admin/reports/r1/actions", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAdminRestrictUser_MissingFields_Returns400(t *testing.T) {
	r := setupRouter()
	r.POST("/api/v1/admin/users/:userId/restrict", authMiddleware("admin-1", "admin"), handleAdminRestrictUser(nil))

	body := `{}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/admin/users/u1/restrict", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAdminRestrictUser_InvalidScope_Returns400(t *testing.T) {
	r := setupRouter()
	r.POST("/api/v1/admin/users/:userId/restrict", authMiddleware("admin-1", "admin"), handleAdminRestrictUser(nil))

	body := `{"restrictionScope":"invalid","reasonCode":"spam"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/admin/users/u1/restrict", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAdminUpdateReportStatus_MissingFields_Returns400(t *testing.T) {
	r := setupRouter()
	r.PATCH("/api/v1/admin/reports/:reportId/status", authMiddleware("admin-1", "admin"), handleAdminUpdateReportStatus(nil))

	body := `{}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/api/v1/admin/reports/r1/status", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAdminUpdateReportStatus_InvalidStatus_Returns400(t *testing.T) {
	r := setupRouter()
	r.PATCH("/api/v1/admin/reports/:reportId/status", authMiddleware("admin-1", "admin"), handleAdminUpdateReportStatus(nil))

	body := `{"status":"deleted"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/api/v1/admin/reports/r1/status", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

// ---------- Utility function tests ----------

func TestItoa(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{42, "42"},
		{100, "100"},
	}
	for _, tc := range tests {
		got := itoa(tc.input)
		if got != tc.expected {
			t.Errorf("itoa(%d) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestNullStr_Valid(t *testing.T) {
	ns := sql.NullString{String: "hello", Valid: true}
	result := nullStr(ns)
	if result != "hello" {
		t.Errorf("nullStr(valid) = %v, want %q", result, "hello")
	}
}

func TestNullStr_Invalid(t *testing.T) {
	ns := sql.NullString{Valid: false}
	result := nullStr(ns)
	if result != nil {
		t.Errorf("nullStr(invalid) = %v, want nil", result)
	}
}

// ---------- Error response structure tests ----------

func TestAdminReportAction_ErrorFormat(t *testing.T) {
	r := setupRouter()
	r.POST("/api/v1/admin/reports/:reportId/actions", authMiddleware("admin-1", "admin"), handleAdminReportAction(nil))

	body := `{}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/admin/reports/r1/actions", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	var resp struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Error.Code != "VALIDATION_ERROR" {
		t.Errorf("error code = %q, want VALIDATION_ERROR", resp.Error.Code)
	}
}
