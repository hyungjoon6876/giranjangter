package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
)

func handleUploadImage(cfg *config.Config, repo repository.UploadRepo) gin.HandlerFunc {
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

		// Peek first 512 bytes for content type detection
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "파일 열기 실패"},
			})
			return
		}
		defer src.Close()

		buf := make([]byte, 512)
		n, _ := src.Read(buf)
		contentType := http.DetectContentType(buf[:n])
		allowedTypes := map[string]bool{
			"image/jpeg": true, "image/png": true, "image/gif": true, "image/webp": true,
		}
		if !allowedTypes[contentType] {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": "이미지 파일만 업로드할 수 있습니다."},
			})
			return
		}
		// Reset file position for saving
		if _, err := src.Seek(0, 0); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "파일 처리 실패"},
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

		dstPath := filepath.Join(uploadPath, filename)
		dstFile, err := os.Create(dstPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "파일 저장 실패"},
			})
			return
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "파일 저장 실패"},
			})
			return
		}

		url := fmt.Sprintf("/uploads/images/%s", filename)

		userID := middleware.GetUserID(c)
		if err := repo.InsertImage(c.Request.Context(), &repository.InsertImageParams{
			ID:          imageID,
			UserID:      userID,
			Filename:    filename,
			URL:         url,
			ContentType: contentType,
			SizeBytes:   file.Size,
		}); err != nil {
			log.Printf("error: image DB insert failed: %v", err)
			os.Remove(dstPath) // cleanup orphaned file
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "이미지 등록 실패"},
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"imageId":      imageID,
			"url":          url,
			"thumbnailUrl": url, // TODO: generate actual thumbnail
			"sizeBytes":    file.Size,
		})
	}
}
