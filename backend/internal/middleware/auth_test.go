package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "test-jwt-secret-key-for-unit-tests"

func newTestAuth() *AuthMiddleware {
	return NewAuthMiddleware(testSecret, 15*time.Minute, 7*24*time.Hour)
}

// ─── GenerateTokens + ParseToken round-trip ─────────

func TestGenerateTokens_ParseToken_RoundTrip(t *testing.T) {
	auth := newTestAuth()

	tests := []struct {
		name   string
		userID string
		role   string
	}{
		{"regular user", "user-123", "user"},
		{"moderator", "mod-456", "moderator"},
		{"admin", "admin-789", "admin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken, refreshToken, err := auth.GenerateTokens(tt.userID, tt.role)
			if err != nil {
				t.Fatalf("GenerateTokens failed: %v", err)
			}
			if accessToken == "" || refreshToken == "" {
				t.Fatal("expected non-empty tokens")
			}

			// Parse access token
			claims, err := auth.ParseToken(accessToken)
			if err != nil {
				t.Fatalf("ParseToken(access) failed: %v", err)
			}
			if claims.UserID != tt.userID {
				t.Errorf("UserID = %q, want %q", claims.UserID, tt.userID)
			}
			if claims.Role != tt.role {
				t.Errorf("Role = %q, want %q", claims.Role, tt.role)
			}
			if claims.Subject != tt.userID {
				t.Errorf("Subject = %q, want %q", claims.Subject, tt.userID)
			}

			// Parse refresh token
			rClaims, err := auth.ParseToken(refreshToken)
			if err != nil {
				t.Fatalf("ParseToken(refresh) failed: %v", err)
			}
			if rClaims.UserID != tt.userID {
				t.Errorf("refresh UserID = %q, want %q", rClaims.UserID, tt.userID)
			}
		})
	}
}

// ─── ParseToken with expired token ──────────────────

func TestParseToken_Expired(t *testing.T) {
	auth := newTestAuth()

	// Create a token that expired 1 hour ago
	now := time.Now().Add(-2 * time.Hour)
	claims := Claims{
		UserID: "user-expired",
		Role:   "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(1 * time.Hour)), // expired 1h ago
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   "user-expired",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(testSecret))
	if err != nil {
		t.Fatalf("failed to create expired token: %v", err)
	}

	_, err = auth.ParseToken(tokenStr)
	if err == nil {
		t.Error("expected error for expired token, got nil")
	}
}

// ─── ParseToken with malformed/invalid token ────────

func TestParseToken_Invalid(t *testing.T) {
	auth := newTestAuth()

	tests := []struct {
		name  string
		token string
	}{
		{"empty string", ""},
		{"garbage", "not-a-jwt-at-all"},
		{"truncated", "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiJ0ZXN0In0"},
		{"wrong secret", ""},
	}

	// Generate a token with a different secret for the "wrong secret" case
	wrongAuth := NewAuthMiddleware("wrong-secret", 15*time.Minute, 7*24*time.Hour)
	wrongToken, _, _ := wrongAuth.GenerateTokens("user-wrong", "user")
	tests[3].token = wrongToken

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := auth.ParseToken(tt.token)
			if err == nil {
				t.Error("expected error for invalid token, got nil")
			}
		})
	}
}

// ─── RequireAuth middleware tests ────────────────────

func setupGinTest() (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	return engine, w
}

type errorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func TestRequireAuth_NoHeader_Returns401(t *testing.T) {
	auth := newTestAuth()
	engine, _ := setupGinTest()

	engine.GET("/test", auth.RequireAuth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", w.Code, http.StatusUnauthorized)
	}

	var resp errorResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Error.Code != "UNAUTHORIZED" {
		t.Errorf("error code = %q, want %q", resp.Error.Code, "UNAUTHORIZED")
	}
}

func TestRequireAuth_InvalidToken_Returns401(t *testing.T) {
	auth := newTestAuth()
	engine, _ := setupGinTest()

	engine.GET("/test", auth.RequireAuth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	tests := []struct {
		name   string
		header string
	}{
		{"no Bearer prefix", "some-token-value"},
		{"malformed token", "Bearer not-a-real-jwt"},
		{"empty Bearer", "Bearer "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/test", nil)
			req.Header.Set("Authorization", tt.header)
			engine.ServeHTTP(w, req)

			if w.Code != http.StatusUnauthorized {
				t.Errorf("status = %d, want %d", w.Code, http.StatusUnauthorized)
			}
		})
	}
}

func TestRequireAuth_ValidToken_SetsContext(t *testing.T) {
	auth := newTestAuth()
	engine, _ := setupGinTest()

	var capturedUserID, capturedRole string
	engine.GET("/test", auth.RequireAuth(), func(c *gin.Context) {
		capturedUserID = GetUserID(c)
		if role, exists := c.Get("userRole"); exists {
			capturedRole = role.(string)
		}
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	accessToken, _, err := auth.GenerateTokens("user-42", "moderator")
	if err != nil {
		t.Fatalf("GenerateTokens failed: %v", err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if capturedUserID != "user-42" {
		t.Errorf("userId = %q, want %q", capturedUserID, "user-42")
	}
	if capturedRole != "moderator" {
		t.Errorf("userRole = %q, want %q", capturedRole, "moderator")
	}
}

// ─── OptionalAuth middleware tests ───────────────────

func TestOptionalAuth_NoHeader_PassesThrough(t *testing.T) {
	auth := newTestAuth()
	engine, _ := setupGinTest()

	var capturedUserID string
	engine.GET("/test", auth.OptionalAuth(), func(c *gin.Context) {
		capturedUserID = GetUserID(c)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if capturedUserID != "" {
		t.Errorf("userId = %q, want empty", capturedUserID)
	}
}

func TestOptionalAuth_ValidToken_SetsContext(t *testing.T) {
	auth := newTestAuth()
	engine, _ := setupGinTest()

	var capturedUserID string
	engine.GET("/test", auth.OptionalAuth(), func(c *gin.Context) {
		capturedUserID = GetUserID(c)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	accessToken, _, _ := auth.GenerateTokens("user-opt", "user")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if capturedUserID != "user-opt" {
		t.Errorf("userId = %q, want %q", capturedUserID, "user-opt")
	}
}

func TestOptionalAuth_InvalidToken_PassesThrough(t *testing.T) {
	auth := newTestAuth()
	engine, _ := setupGinTest()

	var capturedUserID string
	engine.GET("/test", auth.OptionalAuth(), func(c *gin.Context) {
		capturedUserID = GetUserID(c)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer bad-token")
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if capturedUserID != "" {
		t.Errorf("userId = %q, want empty", capturedUserID)
	}
}

// ─── RequireRole middleware tests ────────────────────

func TestRequireRole_MatchingRole_Passes(t *testing.T) {
	engine, _ := setupGinTest()

	engine.GET("/admin", func(c *gin.Context) {
		c.Set("userRole", "admin")
		c.Next()
	}, RequireRole("moderator", "admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/admin", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestRequireRole_NonMatchingRole_Returns403(t *testing.T) {
	engine, _ := setupGinTest()

	engine.GET("/admin", func(c *gin.Context) {
		c.Set("userRole", "user")
		c.Next()
	}, RequireRole("moderator", "admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/admin", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
	}
}

func TestRequireRole_NoRole_Returns403(t *testing.T) {
	engine, _ := setupGinTest()

	engine.GET("/admin", RequireRole("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/admin", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
	}
}

// ─── RejectIfRestricted middleware tests ─────────────

func TestRejectIfRestricted_NotRestricted_Passes(t *testing.T) {
	engine, _ := setupGinTest()

	engine.GET("/write", RejectIfRestricted(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/write", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestRejectIfRestricted_Restricted_Returns403(t *testing.T) {
	engine, _ := setupGinTest()

	engine.GET("/write", func(c *gin.Context) {
		c.Set("accountRestricted", true)
		c.Next()
	}, RejectIfRestricted(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/write", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
	}

	var resp errorResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Error.Code != "ACCOUNT_RESTRICTED" {
		t.Errorf("error code = %q, want %q", resp.Error.Code, "ACCOUNT_RESTRICTED")
	}
}

// ─── GetUserID helper ───────────────────────────────

func TestGetUserID_Set(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", "user-abc")

	if got := GetUserID(c); got != "user-abc" {
		t.Errorf("GetUserID = %q, want %q", got, "user-abc")
	}
}

func TestGetUserID_NotSet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	if got := GetUserID(c); got != "" {
		t.Errorf("GetUserID = %q, want empty", got)
	}
}
