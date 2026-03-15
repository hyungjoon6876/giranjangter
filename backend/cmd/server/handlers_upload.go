package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/config"
)

func handleUploadImage(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": "파일을 선택해주세요."},
			})
			return
		}

		// Size check
		if file.Size > cfg.MaxUploadSize {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": "파일 크기가 10MB를 초과합니다."},
			})
			return
		}

		// Extension check
		ext := strings.ToLower(filepath.Ext(file.Filename))
		allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
		if !allowedExts[ext] {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": "지원하지 않는 파일 형식입니다. (jpg, png, webp)"},
			})
			return
		}

		// Generate unique filename
		imageID := uuid.New().String()
		filename := imageID + ext
		uploadPath := filepath.Join(cfg.UploadDir, "images")

		if err := os.MkdirAll(uploadPath, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "업로드 디렉토리 생성 실패"},
			})
			return
		}

		dst := filepath.Join(uploadPath, filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "파일 저장 실패"},
			})
			return
		}

		url := fmt.Sprintf("/uploads/images/%s", filename)

		c.JSON(http.StatusCreated, gin.H{
			"imageId":      imageID,
			"url":          url,
			"thumbnailUrl": url, // TODO: generate actual thumbnail
			"sizeBytes":    file.Size,
		})
	}
}
