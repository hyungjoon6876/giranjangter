package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
	"github.com/jym/lincle/internal/repository/mock"
)

const testSecret = "test-jwt-secret-key-for-unit-tests"

func newTestAuth() *middleware.AuthMiddleware {
	return middleware.NewAuthMiddleware(testSecret, 15*time.Minute, 7*24*time.Hour)
}

func devConfig() *config.Config {
	return &config.Config{
		Env:           "development",
		JWTRefreshTTL: 7 * 24 * time.Hour,
	}
}

func prodConfig() *config.Config {
	return &config.Config{
		Env:             "production",
		GoogleClientIDs: []string{"real-client-id"},
		JWTRefreshTTL:   7 * 24 * time.Hour,
	}
}

// setupRouter creates a Gin test router with optional auth middleware applied.
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

// setupAuthRouter wires auth routes for testing.
func setupAuthRouter(repo repository.AuthRepo, auth *middleware.AuthMiddleware, cfg *config.Config) *gin.Engine {
	r := setupRouter()
	r.POST("/api/v1/auth/login", handleLogin(repo, auth, cfg))
	r.POST("/api/v1/auth/refresh", handleRefresh(repo, auth, cfg))
	r.GET("/api/v1/me", auth.RequireAuth(), handleGetMe(repo))
	return r
}

type authResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	User         struct {
		UserID        string `json:"userId"`
		Nickname      string `json:"nickname"`
		Role          string `json:"role"`
		AccountStatus string `json:"accountStatus"`
		IsNewUser     bool   `json:"isNewUser"`
	} `json:"user"`
}

type meResponse struct {
	UserID              string  `json:"userId"`
	Role                string  `json:"role"`
	AccountStatus       string  `json:"accountStatus"`
	Nickname            string  `json:"nickname"`
	AvatarURL           *string `json:"avatarUrl"`
	Introduction        *string `json:"introduction"`
	PrimaryServerID     *string `json:"primaryServerId"`
	CompletedTradeCount int     `json:"completedTradeCount"`
	PositiveReviewCount int     `json:"positiveReviewCount"`
	ResponseBadge       string  `json:"responseBadge"`
	TrustBadge          string  `json:"trustBadge"`
	AlignmentScore      int     `json:"alignmentScore"`
	AlignmentGrade      string  `json:"alignmentGrade"`
}

type errResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// ── Login ──

func TestLogin_NewUser_Returns201(t *testing.T) {
	auth := newTestAuth()
	cfg := devConfig()

	var createdUserID, createdProvider, createdProviderKey, createdNickname string
	mockRepo := &mock.MockAuthRepo{
		FindUserByProviderFn: func(ctx context.Context, provider, key string) (*repository.UserWithNickname, error) {
			return nil, nil // user not found
		},
		CreateUserWithProfileFn: func(ctx context.Context, userID, provider, providerKey, nickname string) error {
			createdUserID = userID
			createdProvider = provider
			createdProviderKey = providerKey
			createdNickname = nickname
			return nil
		},
	}

	r := setupAuthRouter(mockRepo, auth, cfg)

	body := `{"provider":"google","providerToken":"new-user-token-123"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	var resp authResponse
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.AccessToken == "" {
		t.Error("expected non-empty accessToken")
	}
	if resp.RefreshToken == "" {
		t.Error("expected non-empty refreshToken")
	}
	if resp.ExpiresIn != 900 {
		t.Errorf("expiresIn = %d, want 900", resp.ExpiresIn)
	}
	if !resp.User.IsNewUser {
		t.Error("expected user.isNewUser = true")
	}
	if resp.User.Role != "user" {
		t.Errorf("user.role = %q, want %q", resp.User.Role, "user")
	}
	if resp.User.AccountStatus != "active" {
		t.Errorf("user.accountStatus = %q, want %q", resp.User.AccountStatus, "active")
	}
	if !strings.HasPrefix(resp.User.Nickname, "유저_") {
		t.Errorf("user.nickname = %q, want prefix '유저_'", resp.User.Nickname)
	}

	// Verify repo was called correctly
	if createdUserID == "" {
		t.Error("CreateUserWithProfile was not called")
	}
	if createdProvider != "google" {
		t.Errorf("provider = %q, want %q", createdProvider, "google")
	}
	if createdProviderKey != "new-user-token-123" {
		t.Errorf("providerKey = %q, want %q", createdProviderKey, "new-user-token-123")
	}
	if !strings.HasPrefix(createdNickname, "유저_") {
		t.Errorf("nickname = %q, want prefix '유저_'", createdNickname)
	}
}

func TestLogin_ExistingUser_Returns200(t *testing.T) {
	auth := newTestAuth()
	cfg := devConfig()

	mockRepo := &mock.MockAuthRepo{
		FindUserByProviderFn: func(ctx context.Context, provider, key string) (*repository.UserWithNickname, error) {
			return &repository.UserWithNickname{
				UserID:   "existing-user-1",
				Role:     "user",
				Nickname: "기존유저",
			}, nil
		},
	}

	r := setupAuthRouter(mockRepo, auth, cfg)

	body := `{"provider":"google","providerToken":"existing-key"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp authResponse
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.User.IsNewUser {
		t.Error("expected user.isNewUser = false")
	}
	if resp.User.UserID != "existing-user-1" {
		t.Errorf("user.userId = %q, want %q", resp.User.UserID, "existing-user-1")
	}
	if resp.User.Nickname != "기존유저" {
		t.Errorf("user.nickname = %q, want %q", resp.User.Nickname, "기존유저")
	}
}

func TestLogin_DevMode_AcceptsAnyToken(t *testing.T) {
	auth := newTestAuth()
	cfg := devConfig()

	mockRepo := &mock.MockAuthRepo{
		FindUserByProviderFn: func(ctx context.Context, provider, key string) (*repository.UserWithNickname, error) {
			return nil, nil
		},
	}

	r := setupAuthRouter(mockRepo, auth, cfg)

	// Test with google provider
	body := `{"provider":"google","providerToken":"literally-anything"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("google: status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}

	// Test with non-google provider in dev mode
	body = `{"provider":"dev","providerToken":"dev-key"}`
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("dev: status = %d, want %d; body = %s", w.Code, http.StatusCreated, w.Body.String())
	}
}

func TestLogin_ProductionMode_InvalidToken_Returns401(t *testing.T) {
	auth := newTestAuth()
	cfg := prodConfig()

	var createCalled bool
	mockRepo := &mock.MockAuthRepo{
		CreateUserWithProfileFn: func(ctx context.Context, userID, provider, providerKey, nickname string) error {
			createCalled = true
			return nil
		},
	}

	r := setupAuthRouter(mockRepo, auth, cfg)

	body := `{"provider":"google","providerToken":"invalid-garbage-token"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusUnauthorized, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "UNAUTHORIZED" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "UNAUTHORIZED")
	}

	if createCalled {
		t.Error("CreateUserWithProfile should not have been called")
	}
}

// ── Refresh ──

func TestRefresh_ValidToken_Returns200(t *testing.T) {
	auth := newTestAuth()
	cfg := devConfig()

	_, refreshToken, err := auth.GenerateTokens("user-1", "user")
	if err != nil {
		t.Fatalf("GenerateTokens failed: %v", err)
	}

	mockRepo := &mock.MockAuthRepo{
		FindRefreshTokenFn: func(ctx context.Context, tokenHash string) (string, error) {
			return "token-row-id", nil
		},
		GetAccountStatusFn: func(ctx context.Context, userID string) (string, error) {
			return "active", nil
		},
	}

	r := setupAuthRouter(mockRepo, auth, cfg)

	body := `{"refreshToken":"` + refreshToken + `"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/auth/refresh", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.AccessToken == "" {
		t.Error("expected non-empty accessToken")
	}
	if resp.RefreshToken == "" {
		t.Error("expected non-empty refreshToken")
	}
	if resp.ExpiresIn != 900 {
		t.Errorf("expiresIn = %d, want 900", resp.ExpiresIn)
	}

	// Verify the new access token is valid and contains correct claims
	claims, err := auth.ParseToken(resp.AccessToken)
	if err != nil {
		t.Fatalf("new access token is not parseable: %v", err)
	}
	if claims.UserID != "user-1" {
		t.Errorf("claims.UserID = %q, want %q", claims.UserID, "user-1")
	}
	if claims.Role != "user" {
		t.Errorf("claims.Role = %q, want %q", claims.Role, "user")
	}
}

func TestRefresh_InvalidToken_Returns401(t *testing.T) {
	auth := newTestAuth()
	cfg := devConfig()

	mockRepo := &mock.MockAuthRepo{}

	r := setupAuthRouter(mockRepo, auth, cfg)

	body := `{"refreshToken":"this.is.not.a.valid.jwt"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/auth/refresh", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusUnauthorized, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "UNAUTHORIZED" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "UNAUTHORIZED")
	}
}

// ── GetMe ──

func TestGetMe_Authenticated_Returns200(t *testing.T) {
	auth := newTestAuth()

	mockRepo := &mock.MockAuthRepo{
		GetUserProfileFn: func(ctx context.Context, userID string) (*repository.FullUserProfile, error) {
			return &repository.FullUserProfile{
				UserID:         userID,
				Role:           "user",
				AccountStatus:  "active",
				Nickname:       "테스트유저",
				TradeCount:     5,
				ReviewCount:    3,
				ResponseBadge:  "fast",
				TrustBadge:     "trusted",
				AlignmentScore: 50,
				AlignmentGrade: "neutral",
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/me", auth.RequireAuth(), handleGetMe(mockRepo))

	accessToken, _, err := auth.GenerateTokens("user-1", "user")
	if err != nil {
		t.Fatalf("GenerateTokens failed: %v", err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/me", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp meResponse
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.UserID != "user-1" {
		t.Errorf("userId = %q, want %q", resp.UserID, "user-1")
	}
	if resp.Role != "user" {
		t.Errorf("role = %q, want %q", resp.Role, "user")
	}
	if resp.AccountStatus != "active" {
		t.Errorf("accountStatus = %q, want %q", resp.AccountStatus, "active")
	}
	if resp.Nickname != "테스트유저" {
		t.Errorf("nickname = %q, want %q", resp.Nickname, "테스트유저")
	}
	if resp.CompletedTradeCount != 5 {
		t.Errorf("completedTradeCount = %d, want %d", resp.CompletedTradeCount, 5)
	}
	if resp.PositiveReviewCount != 3 {
		t.Errorf("positiveReviewCount = %d, want %d", resp.PositiveReviewCount, 3)
	}
	if resp.AlignmentScore != 50 {
		t.Errorf("alignmentScore = %d, want %d", resp.AlignmentScore, 50)
	}
	if resp.AlignmentGrade != "neutral" {
		t.Errorf("alignmentGrade = %q, want %q", resp.AlignmentGrade, "neutral")
	}
}

func TestGetMe_Unauthenticated_Returns401(t *testing.T) {
	auth := newTestAuth()
	mockRepo := &mock.MockAuthRepo{}

	r := setupRouter()
	r.GET("/api/v1/me", auth.RequireAuth(), handleGetMe(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusUnauthorized, w.Body.String())
	}

	var resp errResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error.Code != "UNAUTHORIZED" {
		t.Errorf("error.code = %q, want %q", resp.Error.Code, "UNAUTHORIZED")
	}
}
