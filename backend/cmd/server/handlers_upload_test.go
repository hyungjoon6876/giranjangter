package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/repository/mock"
)

func TestUploadImage_NoFile_Returns400(t *testing.T) {
	cfg := &config.Config{Env: "development", MaxUploadSize: 10 << 20, UploadDir: t.TempDir()}
	mockRepo := &mock.MockUploadRepo{}
	r := setupRouter()
	r.POST("/api/v1/images", authMiddleware("user-1", "user"), handleUploadImage(cfg, mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/images", nil)
	req.Header.Set("Content-Type", "multipart/form-data")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestUploadImage_TooLarge_Returns400(t *testing.T) {
	cfg := &config.Config{Env: "development", MaxUploadSize: 100, UploadDir: t.TempDir()} // 100 bytes max
	mockRepo := &mock.MockUploadRepo{}
	r := setupRouter()
	r.POST("/api/v1/images", authMiddleware("user-1", "user"), handleUploadImage(cfg, mockRepo))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "big.jpg")
	part.Write(make([]byte, 200)) // 200 bytes > 100 max
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/images", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestUploadImage_InvalidExtension_Returns400(t *testing.T) {
	cfg := &config.Config{Env: "development", MaxUploadSize: 10 << 20, UploadDir: t.TempDir()}
	mockRepo := &mock.MockUploadRepo{}
	r := setupRouter()
	r.POST("/api/v1/images", authMiddleware("user-1", "user"), handleUploadImage(cfg, mockRepo))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "hack.exe")
	part.Write([]byte("not an image"))
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/images", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestUploadImage_InvalidContentType_Returns400(t *testing.T) {
	cfg := &config.Config{Env: "development", MaxUploadSize: 10 << 20, UploadDir: t.TempDir()}
	mockRepo := &mock.MockUploadRepo{}
	r := setupRouter()
	r.POST("/api/v1/images", authMiddleware("user-1", "user"), handleUploadImage(cfg, mockRepo))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "fake.jpg")
	part.Write([]byte("this is not JPEG data, just text content for testing purposes"))
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/images", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestUploadImage_ValidJPEG_Returns201(t *testing.T) {
	cfg := &config.Config{Env: "development", MaxUploadSize: 10 << 20, UploadDir: t.TempDir()}
	mockRepo := &mock.MockUploadRepo{}
	r := setupRouter()
	r.POST("/api/v1/images", authMiddleware("user-1", "user"), handleUploadImage(cfg, mockRepo))

	// Minimal valid JPEG: SOI marker + enough data
	jpegData := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00}
	jpegData = append(jpegData, make([]byte, 100)...)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "photo.jpg")
	part.Write(jpegData)
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/images", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}
}
