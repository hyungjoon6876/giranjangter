package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/config"
)

func TestCorsMiddleware_AllowedOrigin_SetsHeader(t *testing.T) {
	cfg := &config.Config{
		Env:            "production",
		AllowedOrigins: []string{"https://giranjt.com", "https://www.giranjt.com"},
	}
	r := setupRouter()
	r.Use(corsMiddleware(cfg))
	r.GET("/test", func(c *gin.Context) { c.Status(200) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://giranjt.com")
	r.ServeHTTP(w, req)

	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "https://giranjt.com" {
		t.Errorf("ACAO = %q, want %q", got, "https://giranjt.com")
	}
}

func TestCorsMiddleware_UnknownOrigin_NoACAO(t *testing.T) {
	cfg := &config.Config{
		Env:            "production",
		AllowedOrigins: []string{"https://giranjt.com"},
	}
	r := setupRouter()
	r.Use(corsMiddleware(cfg))
	r.GET("/test", func(c *gin.Context) { c.Status(200) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://evil.com")
	r.ServeHTTP(w, req)

	if got := w.Header().Get("Access-Control-Allow-Origin"); got == "https://evil.com" {
		t.Errorf("should not allow unknown origin, got %q", got)
	}
}

func TestCorsMiddleware_DevMode_AllowsAll(t *testing.T) {
	cfg := &config.Config{
		Env:            "development",
		AllowedOrigins: []string{},
	}
	r := setupRouter()
	r.Use(corsMiddleware(cfg))
	r.GET("/test", func(c *gin.Context) { c.Status(200) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	r.ServeHTTP(w, req)

	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("dev mode ACAO = %q, want *", got)
	}
}

func TestCorsMiddleware_OPTIONS_Returns204(t *testing.T) {
	cfg := &config.Config{
		Env:            "development",
		AllowedOrigins: []string{},
	}
	r := setupRouter()
	r.Use(corsMiddleware(cfg))
	r.GET("/test", func(c *gin.Context) { c.Status(200) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest("OPTIONS", "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204 for OPTIONS, got %d", w.Code)
	}
}

func TestCorsMiddleware_SetsRequiredHeaders(t *testing.T) {
	cfg := &config.Config{Env: "development"}
	r := setupRouter()
	r.Use(corsMiddleware(cfg))
	r.GET("/test", func(c *gin.Context) { c.Status(200) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)

	if got := w.Header().Get("Access-Control-Allow-Methods"); got == "" {
		t.Error("missing Access-Control-Allow-Methods")
	}
	if got := w.Header().Get("Access-Control-Allow-Headers"); got == "" {
		t.Error("missing Access-Control-Allow-Headers")
	}
	if got := w.Header().Get("Access-Control-Max-Age"); got != "86400" {
		t.Errorf("Max-Age = %q, want 86400", got)
	}
}
