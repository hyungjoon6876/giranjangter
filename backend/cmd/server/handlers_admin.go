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

func handleAdminListReports(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, reporter_user_id, target_type, target_id, report_type, status, created_at FROM reports ORDER BY created_at DESC LIMIT 50")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
			return
		}
		defer rows.Close()

		var reports []gin.H
		for rows.Next() {
			var id, reporter, tt, tid, rt, st string
			var created time.Time
			if err := rows.Scan(&id, &reporter, &tt, &tid, &rt, &st, &created); err != nil {
				continue
			}
			reports = append(reports, gin.H{
				"reportId": id, "reporterUserId": reporter, "targetType": tt, "targetId": tid,
				"reportType": rt, "status": st, "createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": reports})
	}
}

func handleAdminGetReport(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		reportID := c.Param("reportId")
		var r struct {
			ID          string `json:"reportId"`
			Reporter    string `json:"reporterUserId"`
			TargetType  string `json:"targetType"`
			TargetID    string `json:"targetId"`
			ReportType  string `json:"reportType"`
			Description string `json:"description"`
			Status      string `json:"status"`
		}
		err := db.QueryRow("SELECT id, reporter_user_id, target_type, target_id, report_type, description, status FROM reports WHERE id = $1", reportID).Scan(&r.ID, &r.Reporter, &r.TargetType, &r.TargetID, &r.ReportType, &r.Description, &r.Status)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "신고를 찾을 수 없습니다."}})
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

func handleAdminReportAction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		reportID := c.Param("reportId")
		adminID := middleware.GetUserID(c)
		var req struct {
			ActionCode       string  `json:"actionCode" binding:"required"`
			TargetUserID     string  `json:"targetUserId" binding:"required"`
			Memo             *string `json:"memo"`
			RestrictionScope *string `json:"restrictionScope"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		actionID := uuid.New().String()
		if _, err := tx.Exec("INSERT INTO moderation_actions (id, report_id, actor_user_id, target_user_id, action_code, restriction_scope, memo) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			actionID, reportID, adminID, req.TargetUserID, req.ActionCode, req.RestrictionScope, req.Memo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "조치 기록 실패"}})
			return
		}
		if _, err := tx.Exec("UPDATE reports SET status = 'resolved', updated_at = NOW() WHERE id = $1", reportID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "신고 상태 업데이트 실패"}})
			return
		}
		if _, err := tx.Exec("INSERT INTO audit_logs (id, actor_id, actor_role, action, target_type, target_id, details_json) VALUES ($1, $2, 'admin', $3, 'user', $4, $5)",
			uuid.New().String(), adminID, req.ActionCode, req.TargetUserID, req.Memo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "감사 로그 기록 실패"}})
			return
		}

		// 신고 처리 확정 시 대상자 성향치 -20
		if err := alignment.Change(tx, req.TargetUserID, domain.AlignmentReportConfirmed, "report_confirmed", "report", reportID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "성향치 업데이트 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"actionId": actionID, "status": "resolved"})
	}
}

func handleAdminHideListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		adminID := middleware.GetUserID(c)

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("UPDATE listings SET visibility = 'hidden', updated_at = NOW() WHERE id = $1", listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 숨김 실패"}})
			return
		}
		if _, err := tx.Exec("INSERT INTO audit_logs (id, actor_id, actor_role, action, target_type, target_id) VALUES ($1, $2, 'admin', 'listing.moderation.hide', 'listing', $3)",
			uuid.New().String(), adminID, listingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "감사 로그 기록 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"listingId": listingID, "visibility": "hidden"})
	}
}

func handleAdminRestrictUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetUserID := c.Param("userId")
		adminID := middleware.GetUserID(c)
		var req struct {
			RestrictionScope string  `json:"restrictionScope" binding:"required"`
			DurationDays     *int    `json:"durationDays"`
			ReasonCode       string  `json:"reasonCode" binding:"required"`
			Memo             *string `json:"memo"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer tx.Rollback()

		if _, err := tx.Exec("UPDATE users SET account_status = 'restricted' WHERE id = $1", targetUserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "계정 제한 실패"}})
			return
		}
		actionID := uuid.New().String()
		if _, err := tx.Exec("INSERT INTO moderation_actions (id, actor_user_id, target_user_id, action_code, restriction_scope, duration_days, memo) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			actionID, adminID, targetUserID, "user.restriction."+req.RestrictionScope, req.RestrictionScope, req.DurationDays, req.Memo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "조치 기록 실패"}})
			return
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"actionId": actionID, "targetUserId": targetUserID, "restrictionScope": req.RestrictionScope})
	}
}
