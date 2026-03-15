package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/middleware"
)

func handleCreateReport(db *sql.DB) gin.HandlerFunc {
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
		if _, err := db.Exec("INSERT INTO reports (id, reporter_user_id, target_type, target_id, report_type, description, status) VALUES ($1, $2, $3, $4, $5, $6, 'submitted')",
			reportID, userID, req.TargetType, req.TargetID, req.ReportType, req.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "신고 접수 실패"}})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"reportId": reportID, "status": "submitted"})
	}
}

func handleMyReports(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		rows, err := db.Query("SELECT id, target_type, target_id, report_type, status, created_at FROM reports WHERE reporter_user_id = $1 ORDER BY created_at DESC LIMIT 50", userID)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		defer rows.Close()

		var reports []gin.H
		for rows.Next() {
			var id, tt, tid, rt, st string
			var created time.Time
			if err := rows.Scan(&id, &tt, &tid, &rt, &st, &created); err != nil {
				continue
			}
			reports = append(reports, gin.H{"reportId": id, "targetType": tt, "targetId": tid, "reportType": rt, "status": st, "createdAt": created.Format(time.RFC3339)})
		}
		c.JSON(http.StatusOK, gin.H{"data": reports})
	}
}
