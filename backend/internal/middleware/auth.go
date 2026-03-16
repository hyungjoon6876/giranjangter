package middleware

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	secret    []byte
	accessTTL time.Duration
	refreshTTL time.Duration
}

func NewAuthMiddleware(secret string, accessTTL, refreshTTL time.Duration) *AuthMiddleware {
	return &AuthMiddleware{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

// GenerateTokens creates access and refresh tokens.
func (a *AuthMiddleware) GenerateTokens(userID, role string) (accessToken, refreshToken string, err error) {
	now := time.Now()

	accessClaims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(a.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   userID,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(a.secret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(a.refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   userID,
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(a.secret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ParseToken validates and parses a JWT token.
func (a *AuthMiddleware) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return a.secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

// extractAndValidateToken extracts the Bearer token from the Authorization header
// (or from a "token" query parameter for SSE/EventSource connections),
// parses it, and returns the claims if valid. On failure it sends an error response
// and returns nil.
func (a *AuthMiddleware) extractAndValidateToken(c *gin.Context) *Claims {
	header := c.GetHeader("Authorization")
	var tokenStr string

	if header != "" {
		tokenStr = strings.TrimPrefix(header, "Bearer ")
		if tokenStr == header {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "잘못된 인증 형식입니다."},
			})
			return nil
		}
	} else if qToken := c.Query("token"); qToken != "" {
		tokenStr = qToken
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": gin.H{"code": "UNAUTHORIZED", "message": "인증이 필요합니다."},
		})
		return nil
	}

	claims, err := a.ParseToken(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": gin.H{"code": "UNAUTHORIZED", "message": "유효하지 않은 토큰입니다."},
		})
		return nil
	}

	return claims
}

// RequireAuth is a Gin middleware that validates JWT tokens.
func (a *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := a.extractAndValidateToken(c)
		if claims == nil {
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}

// RequireAuthWithDB checks JWT + account status from DB.
func (a *AuthMiddleware) RequireAuthWithDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := a.extractAndValidateToken(c)
		if claims == nil {
			return
		}

		// DB에서 계정 상태 확인
		var accountStatus, role string
		err := db.QueryRow("SELECT account_status, role FROM users WHERE id = $1", claims.UserID).Scan(&accountStatus, &role)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "사용자를 찾을 수 없습니다."},
			})
			return
		}

		switch accountStatus {
		case "suspended":
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": gin.H{"code": "ACCOUNT_SUSPENDED", "message": "계정이 정지되었습니다. 고객센터에 문의해주세요."},
			})
			return
		case "withdrawn":
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": gin.H{"code": "ACCOUNT_WITHDRAWN", "message": "탈퇴한 계정입니다."},
			})
			return
		case "restricted":
			// restricted는 일부 기능만 제한 — context에 상태를 넣어서 핸들러에서 판단
			c.Set("accountRestricted", true)
		}

		c.Set("userId", claims.UserID)
		c.Set("userRole", role) // DB의 최신 role 사용
		c.Set("accountStatus", accountStatus)
		c.Next()
	}
}

// OptionalAuth extracts user info if token is present, but doesn't block.
func (a *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.Next()
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		if tokenStr == header {
			c.Next()
			return
		}

		claims, err := a.ParseToken(tokenStr)
		if err == nil {
			c.Set("userId", claims.UserID)
			c.Set("userRole", claims.Role)
		}
		c.Next()
	}
}

// RequireRole ensures the authenticated user has the required role.
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": gin.H{"code": "FORBIDDEN", "message": "권한이 없습니다."},
			})
			return
		}
		userRole := role.(string)
		for _, r := range roles {
			if userRole == r {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": gin.H{"code": "FORBIDDEN", "message": "권한이 없습니다."},
		})
	}
}

// RejectIfRestricted blocks restricted users from write operations.
func RejectIfRestricted() gin.HandlerFunc {
	return func(c *gin.Context) {
		if restricted, _ := c.Get("accountRestricted"); restricted == true {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": gin.H{
					"code":    "ACCOUNT_RESTRICTED",
					"message": "계정에 이용 제한이 적용되어 있습니다. 자세한 내용은 알림을 확인해주세요.",
				},
			})
			return
		}
		c.Next()
	}
}

// GetUserID extracts the authenticated user's ID from the context.
func GetUserID(c *gin.Context) string {
	if v, ok := c.Get("userId"); ok {
		return v.(string)
	}
	return ""
}
