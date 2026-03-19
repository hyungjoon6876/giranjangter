package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handleAdminListUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Query("q")
		status := c.Query("status")
		limit := 50

		query := `SELECT u.id, up.nickname, u.account_status, u.role,
			up.completed_trade_count, up.alignment_score, up.alignment_grade,
			u.last_login_at, u.created_at
			FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id
			WHERE 1=1`
		args := []interface{}{}
		argN := 1

		if q != "" {
			query += ` AND (up.nickname ILIKE $` + itoa(argN) + `)`
			args = append(args, "%"+q+"%")
			argN++
		}
		if status != "" {
			query += ` AND u.account_status = $` + itoa(argN)
			args = append(args, status)
			argN++
		}
		_ = argN
		query += ` ORDER BY u.created_at DESC LIMIT ` + itoa(limit)

		rows, err := db.Query(query, args...)
		if err != nil {
			log.Printf("admin list users error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer func() { _ = rows.Close() }()

		var users []gin.H
		for rows.Next() {
			var id, accountStatus, role string
			var nickname, grade sql.NullString
			var tradeCount sql.NullInt64
			var score sql.NullInt64
			var lastLogin *time.Time
			var created time.Time
			if err := rows.Scan(&id, &nickname, &accountStatus, &role, &tradeCount, &score, &grade, &lastLogin, &created); err != nil {
				continue
			}
			u := gin.H{
				"userId":              id,
				"nickname":            nickname.String,
				"accountStatus":       accountStatus,
				"role":                role,
				"completedTradeCount": tradeCount.Int64,
				"alignmentScore":      score.Int64,
				"alignmentGrade":      grade.String,
				"createdAt":           created.Format(time.RFC3339),
			}
			if lastLogin != nil {
				u["lastLoginAt"] = lastLogin.Format(time.RFC3339)
			}
			users = append(users, u)
		}
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func handleAdminGetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")

		var id, accountStatus, role string
		var nickname, intro, server, grade, trustBadge sql.NullString
		var tradeCount, reviewPos sql.NullInt64
		var score sql.NullInt64
		err := db.QueryRow(`SELECT u.id, up.nickname, u.account_status, u.role,
			up.introduction, up.primary_server_id, up.completed_trade_count,
			up.positive_review_count, up.alignment_score, up.alignment_grade, up.trust_badge
			FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.id = $1`, userID).
			Scan(&id, &nickname, &accountStatus, &role, &intro, &server,
				&tradeCount, &reviewPos, &score, &grade, &trustBadge)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "사용자를 찾을 수 없습니다."}})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userId":              id,
			"nickname":            nickname.String,
			"accountStatus":       accountStatus,
			"role":                role,
			"introduction":        nullStr(intro),
			"primaryServerId":     nullStr(server),
			"completedTradeCount": tradeCount.Int64,
			"positiveReviewCount": reviewPos.Int64,
			"alignmentScore":      score.Int64,
			"alignmentGrade":      grade.String,
			"trustBadge":          nullStr(trustBadge),
		})
	}
}

func handleAdminUserModerationHistory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")
		rows, err := db.Query(`SELECT id, report_id, actor_user_id, action_code, restriction_scope, duration_days, memo, created_at
			FROM moderation_actions WHERE target_user_id = $1 ORDER BY created_at DESC LIMIT 50`, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer func() { _ = rows.Close() }()

		var actions []gin.H
		for rows.Next() {
			var id, actorID, actionCode string
			var reportID, scope, memo sql.NullString
			var duration sql.NullInt64
			var created time.Time
			if err := rows.Scan(&id, &reportID, &actorID, &actionCode, &scope, &duration, &memo, &created); err != nil {
				continue
			}
			a := gin.H{
				"actionId":         id,
				"reportId":         nullStr(reportID),
				"actorUserId":      actorID,
				"actionCode":       actionCode,
				"restrictionScope": nullStr(scope),
				"memo":             nullStr(memo),
				"createdAt":        created.Format(time.RFC3339),
			}
			if duration.Valid {
				a["durationDays"] = duration.Int64
			}
			actions = append(actions, a)
		}
		c.JSON(http.StatusOK, gin.H{"data": actions})
	}
}

// itoa converts an int to its string representation for query building.
func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}

// nullStr returns nil for invalid NullString, or the string value.
func nullStr(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}
