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

func handleCreateReport(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		var req struct {
			TargetType  string `json:"targetType" binding:"required,oneof=user listing message"`
			TargetID    string `json:"targetId" binding:"required"`
			ReportType  string `json:"reportType" binding:"required,oneof=fake_listing scam_suspicion no_show harassment spam prohibited_item privacy_exposure other"`
			Description string `json:"description" binding:"required,min=1,max=2000"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		reportID := uuid.New().String()
		if err := repo.CreateReport(c.Request.Context(), &repository.CreateReportParams{
			ID:          reportID,
			ReporterID:  userID,
			TargetType:  req.TargetType,
			TargetID:    req.TargetID,
			ReportType:  req.ReportType,
			Description: req.Description,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "신고 접수 실패"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"reportId": reportID, "status": "submitted"})
	}
}

func handleMyReports(repo repository.ReservationRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		items, err := repo.ListMyReports(c.Request.Context(), userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var reports []gin.H
		for _, item := range items {
			reports = append(reports, gin.H{"reportId": item.ReportID, "targetType": item.TargetType, "targetId": item.TargetID, "reportType": item.ReportType, "status": item.Status, "createdAt": item.CreatedAt.Format(time.RFC3339)})
		}
		c.JSON(http.StatusOK, gin.H{"data": reports})
	}
}
