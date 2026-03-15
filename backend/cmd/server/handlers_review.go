package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/alignment"
	"github.com/jym/lincle/internal/domain"
	"github.com/jym/lincle/internal/middleware"
)

func handleCreateReview(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		compID := c.Param("compId")
		userID := middleware.GetUserID(c)
		var req struct {
			Rating  string  `json:"rating" binding:"required,oneof=positive negative"`
			Comment *string `json:"comment"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		var status, reqBy, counterpart string
		err := db.QueryRow("SELECT status, requested_by_user_id, counterpart_user_id FROM trade_completions WHERE id = $1", compID).Scan(&status, &reqBy, &counterpart)
		if err != nil || status != "confirmed" {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "완료된 거래에만 후기를 작성할 수 있습니다."}})
			return
		}

		targetUser := counterpart
		if userID == counterpart {
			targetUser = reqBy
		}

		reviewID := uuid.New().String()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("INSERT INTO reviews (id, completion_id, reviewer_user_id, target_user_id, rating, comment) VALUES ($1, $2, $3, $4, $5, $6)",
			reviewID, compID, userID, targetUser, req.Rating, req.Comment); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": gin.H{"code": "CONFLICT", "message": "이미 후기를 작성했습니다."}})
			return
		}

		if req.Rating == "positive" {
			if _, err := tx.Exec("UPDATE user_profiles SET positive_review_count = positive_review_count + 1 WHERE user_id = $1", targetUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "프로필 업데이트 실패"}})
				return
			}
			// 대상자: 긍정 후기 받음 +3
			if err := alignment.Change(tx, targetUser, domain.AlignmentPositiveReview, "positive_review_received", "review", reviewID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "성향치 업데이트 실패"}})
				return
			}
		} else {
			// 대상자: 부정 후기 받음 -10
			if err := alignment.Change(tx, targetUser, domain.AlignmentNegativeReview, "negative_review_received", "review", reviewID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "성향치 업데이트 실패"}})
				return
			}
		}

		if _, err := tx.Exec("UPDATE user_profiles SET completed_trade_count = completed_trade_count + 1 WHERE user_id = $1", userID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "프로필 업데이트 실패"}})
			return
		}

		// 작성자: 후기 작성 인센티브 +1
		if err := alignment.Change(tx, userID, domain.AlignmentReviewWritten, "review_written", "review", reviewID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "성향치 업데이트 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"reviewId": reviewID})
	}
}

func handleGetUserReviews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetUserID := c.Param("userId")
		rows, err := db.Query(`SELECT r.id, r.rating, r.comment, r.created_at, p.nickname
			FROM reviews r JOIN user_profiles p ON r.reviewer_user_id = p.user_id
			WHERE r.target_user_id = $1 ORDER BY r.created_at DESC LIMIT 50`, targetUserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
			return
		}
		defer rows.Close()

		var reviews []gin.H
		for rows.Next() {
			var id, rating, nick string
			var comment *string
			var created time.Time
			if err := rows.Scan(&id, &rating, &comment, &created, &nick); err != nil {
				continue
			}
			reviews = append(reviews, gin.H{"reviewId": id, "rating": rating, "comment": comment, "reviewerNickname": nick, "createdAt": created.Format(time.RFC3339)})
		}
		c.JSON(http.StatusOK, gin.H{"data": reviews})
	}
}
