package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/oauth"
)

type loginRequest struct {
	Provider      string `json:"provider" binding:"required"`
	ProviderToken string `json:"providerToken" binding:"required"`
}

type refreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func handleLogin(db *sql.DB, auth *middleware.AuthMiddleware, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		var providerKey string

		switch req.Provider {
		case "google":
			if cfg.IsDev() && len(cfg.GoogleClientIDs) == 0 {
				// Dev mode without Google credentials: use token as-is
				providerKey = req.ProviderToken
				log.Println("[auth] dev mode: skipping Google token verification")
			} else {
				// Production: verify Google ID token
				info, err := oauth.VerifyGoogleIDToken(req.ProviderToken, cfg.GoogleClientIDs)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": gin.H{"code": "UNAUTHORIZED", "message": "Google 인증에 실패했습니다."},
					})
					return
				}
				providerKey = info.Sub // Google's stable user ID
			}
		default:
			if cfg.IsDev() {
				// Dev mode: allow any provider with token as key
				providerKey = req.ProviderToken
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": gin.H{"code": "VALIDATION_ERROR", "message": "지원하지 않는 로그인 방식입니다."},
				})
				return
			}
		}

		var userID, role, nickname string
		var isNew bool

		err := db.QueryRow(
			"SELECT u.id, u.role, COALESCE(p.nickname, '') FROM users u LEFT JOIN user_profiles p ON u.id = p.user_id WHERE u.login_provider = $1 AND u.login_provider_user_key = $2",
			req.Provider, providerKey,
		).Scan(&userID, &role, &nickname)

		if err == sql.ErrNoRows {
			userID = uuid.New().String()
			role = "user"
			nickname = "유저_" + userID[:8]
			isNew = true

			tx, err := db.Begin()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
				})
				return
			}
			defer tx.Rollback()

			if _, err = tx.Exec(
				"INSERT INTO users (id, login_provider, login_provider_user_key, account_status, role) VALUES ($1, $2, $3, 'active', 'user')",
				userID, req.Provider, providerKey,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{"code": "INTERNAL_ERROR", "message": "계정 생성 실패"},
				})
				return
			}

			if _, err = tx.Exec(
				"INSERT INTO user_profiles (user_id, nickname) VALUES ($1, $2)",
				userID, nickname,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{"code": "INTERNAL_ERROR", "message": "프로필 생성 실패"},
				})
				return
			}

			if err := tx.Commit(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
				})
				return
			}
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}

		db.Exec("UPDATE users SET last_login_at = $1 WHERE id = $2", time.Now().UTC(), userID)

		accessToken, refreshToken, err := auth.GenerateTokens(userID, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "토큰 생성 실패"},
			})
			return
		}

		// Store refresh token hash in DB for server-side management
		tokenHash := sha256.Sum256([]byte(refreshToken))
		hashStr := hex.EncodeToString(tokenHash[:])
		if _, err := db.Exec(
			"INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, $4)",
			uuid.New().String(), userID, hashStr, time.Now().Add(cfg.JWTRefreshTTL),
		); err != nil {
			log.Printf("[auth] failed to store refresh token: %v", err)
		}

		status := http.StatusOK
		if isNew {
			status = http.StatusCreated
		}

		c.JSON(status, gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"expiresIn":    900,
			"user": gin.H{
				"userId":        userID,
				"nickname":      nickname,
				"role":          role,
				"accountStatus": "active",
				"isNewUser":     isNew,
			},
		})
	}
}

func handleRefresh(db *sql.DB, auth *middleware.AuthMiddleware, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req refreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		claims, err := auth.ParseToken(req.RefreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "유효하지 않은 리프레시 토큰"},
			})
			return
		}

		// Verify refresh token exists in DB and hasn't expired
		tokenHash := sha256.Sum256([]byte(req.RefreshToken))
		hashStr := hex.EncodeToString(tokenHash[:])
		var tokenID string
		err = db.QueryRow(
			"SELECT id FROM refresh_tokens WHERE token_hash = $1 AND expires_at > NOW()",
			hashStr,
		).Scan(&tokenID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "유효하지 않은 리프레시 토큰"},
			})
			return
		}

		// Check account status FIRST (before deleting anything)
		var accountStatus string
		err = db.QueryRow("SELECT account_status FROM users WHERE id = $1", claims.UserID).Scan(&accountStatus)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "사용자를 찾을 수 없습니다."},
			})
			return
		}
		if accountStatus == "suspended" || accountStatus == "withdrawn" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": gin.H{"code": "FORBIDDEN", "message": "계정이 비활성 상태입니다."},
			})
			return
		}

		// Transaction for token rotation
		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("DELETE FROM refresh_tokens WHERE id = $1", tokenID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}

		accessToken, refreshToken, err := auth.GenerateTokens(claims.UserID, claims.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "토큰 생성 실패"},
			})
			return
		}

		// Store new refresh token hash in DB
		newTokenHash := sha256.Sum256([]byte(refreshToken))
		newHashStr := hex.EncodeToString(newTokenHash[:])
		if _, err := tx.Exec(
			"INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, $4)",
			uuid.New().String(), claims.UserID, newHashStr, time.Now().Add(cfg.JWTRefreshTTL),
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"expiresIn":    900,
		})
	}
}

func handleGetMe(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		var u struct {
			ID             string  `json:"userId"`
			Role           string  `json:"role"`
			AccountStatus  string  `json:"accountStatus"`
			Nickname       string  `json:"nickname"`
			AvatarURL      *string `json:"avatarUrl"`
			Introduction   *string `json:"introduction"`
			PrimaryServer  *string `json:"primaryServerId"`
			TradeCount     int     `json:"completedTradeCount"`
			ReviewCount    int     `json:"positiveReviewCount"`
			ResponseBadge  string  `json:"responseBadge"`
			TrustBadge     string  `json:"trustBadge"`
			AlignmentScore int     `json:"alignmentScore"`
			AlignmentGrade string  `json:"alignmentGrade"`
		}

		err := db.QueryRow(`
			SELECT u.id, u.role, u.account_status,
				p.nickname, p.avatar_url, p.introduction, p.primary_server_id,
				p.completed_trade_count, p.positive_review_count, p.response_badge, p.trust_badge,
				p.alignment_score, p.alignment_grade
			FROM users u JOIN user_profiles p ON u.id = p.user_id
			WHERE u.id = $1`, userID,
		).Scan(&u.ID, &u.Role, &u.AccountStatus,
			&u.Nickname, &u.AvatarURL, &u.Introduction, &u.PrimaryServer,
			&u.TradeCount, &u.ReviewCount, &u.ResponseBadge, &u.TrustBadge,
			&u.AlignmentScore, &u.AlignmentGrade)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{"code": "NOT_FOUND", "message": "사용자를 찾을 수 없습니다."},
			})
			return
		}

		c.JSON(http.StatusOK, u)
	}
}

func handleUpdateProfile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		var req struct {
			Nickname      *string `json:"nickname"`
			Introduction  *string `json:"introduction"`
			PrimaryServer *string `json:"primaryServerId"`
			AvatarURL     *string `json:"avatarUrl"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		setClauses := []string{}
		args := []interface{}{}
		paramIdx := 1
		if req.Nickname != nil {
			setClauses = append(setClauses, fmt.Sprintf("nickname = $%d", paramIdx))
			args = append(args, *req.Nickname)
			paramIdx++
		}
		if req.Introduction != nil {
			setClauses = append(setClauses, fmt.Sprintf("introduction = $%d", paramIdx))
			args = append(args, *req.Introduction)
			paramIdx++
		}
		if req.PrimaryServer != nil {
			setClauses = append(setClauses, fmt.Sprintf("primary_server_id = $%d", paramIdx))
			args = append(args, *req.PrimaryServer)
			paramIdx++
		}
		if req.AvatarURL != nil {
			setClauses = append(setClauses, fmt.Sprintf("avatar_url = $%d", paramIdx))
			args = append(args, *req.AvatarURL)
			paramIdx++
		}
		if len(setClauses) == 0 {
			handleGetMe(db)(c)
			return
		}
		args = append(args, userID)
		query := "UPDATE user_profiles SET " + strings.Join(setClauses, ", ") + fmt.Sprintf(", updated_at = NOW() WHERE user_id = $%d", paramIdx)
		if _, err := db.Exec(query, args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "프로필 수정 실패"}})
			return
		}

		handleGetMe(db)(c)
	}
}

func handleLogout(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"code": "UNAUTHORIZED", "message": "인증이 필요합니다."},
			})
			return
		}
		// Delete all refresh tokens for this user (logout from all devices)
		db.Exec("DELETE FROM refresh_tokens WHERE user_id = $1", userID)
		c.Status(http.StatusNoContent)
	}
}
