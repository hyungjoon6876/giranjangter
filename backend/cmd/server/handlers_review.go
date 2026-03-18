package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func handleCreateReview(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		compID := c.Param("compId")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req struct {
			Rating  string  `json:"rating" binding:"required,oneof=positive negative"`
			Comment *string `json:"comment"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		info, err := repo.GetCompletionForReview(ctx, compID)
		if err != nil || info == nil || info.Status != "confirmed" {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "완료된 거래에만 후기를 작성할 수 있습니다."}})
			return
		}

		targetUser := info.CounterpartUserID
		if userID == info.CounterpartUserID {
			targetUser = info.RequestedByUserID
		}

		reviewID := uuid.New().String()

		if err := repo.CreateReview(ctx, &repository.CreateReviewParams{
			ReviewID:     reviewID,
			CompletionID: compID,
			ReviewerID:   userID,
			TargetUserID: targetUser,
			Rating:       req.Rating,
			Comment:      req.Comment,
		}); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": gin.H{"code": "CONFLICT", "message": "이미 후기를 작성했습니다."}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"reviewId": reviewID})
	}
}

func handleGetUserReviews(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetUserID := c.Param("userId")

		items, err := repo.ListUserReviews(c.Request.Context(), targetUserID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var reviews []gin.H
		for _, item := range items {
			reviews = append(reviews, gin.H{"reviewId": item.ReviewID, "rating": item.Rating, "comment": item.Comment, "reviewerNickname": item.ReviewerNickname, "createdAt": item.CreatedAt.Format(time.RFC3339)})
		}
		c.JSON(http.StatusOK, gin.H{"data": reviews})
	}
}
