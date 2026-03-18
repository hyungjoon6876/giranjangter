package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/repository"
)

func listServers(repo repository.MasterRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := repo.ListServers(c.Request.Context())
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var servers []gin.H
		for _, item := range items {
			servers = append(servers, gin.H{"serverId": item.ServerID, "serverName": item.ServerName})
		}
		c.JSON(http.StatusOK, gin.H{"data": servers})
	}
}

func listCategories(repo repository.MasterRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := repo.ListCategories(c.Request.Context())
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var categories []gin.H
		for _, item := range items {
			categories = append(categories, gin.H{"categoryId": item.CategoryID, "categoryName": item.CategoryName, "parentId": item.ParentID})
		}
		c.JSON(http.StatusOK, gin.H{"data": categories})
	}
}

func searchItems(repo repository.MasterRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.TrimSpace(c.Query("q"))
		categoryID := strings.TrimSpace(c.Query("categoryId"))

		// Return empty array if query is empty
		if query == "" {
			c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
			return
		}

		var catPtr *string
		if categoryID != "" {
			catPtr = &categoryID
		}

		results, err := repo.SearchItems(c.Request.Context(), query, catPtr)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var items []gin.H
		for _, r := range results {
			var iconURL *string
			if r.IconID != nil {
				u := "/static/icons/" + *r.IconID + ".png"
				iconURL = &u
			}
			items = append(items, gin.H{
				"id":         r.ID,
				"name":       r.Name,
				"categoryId": r.CategoryID,
				"iconUrl":    iconURL,
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}
