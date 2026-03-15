package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func listServers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name FROM servers WHERE is_active = 1 ORDER BY sort_order")
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		defer rows.Close()

		var servers []gin.H
		for rows.Next() {
			var id, name string
			rows.Scan(&id, &name)
			servers = append(servers, gin.H{"serverId": id, "serverName": name})
		}
		c.JSON(http.StatusOK, gin.H{"data": servers})
	}
}

func listCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, parent_id FROM categories ORDER BY parent_id NULLS FIRST, sort_order")
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		defer rows.Close()

		var categories []gin.H
		for rows.Next() {
			var id, name string
			var parentID *string
			rows.Scan(&id, &name, &parentID)
			categories = append(categories, gin.H{"categoryId": id, "categoryName": name, "parentId": parentID})
		}
		c.JSON(http.StatusOK, gin.H{"data": categories})
	}
}

func searchItems(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.TrimSpace(c.Query("q"))
		categoryID := strings.TrimSpace(c.Query("categoryId"))

		// Return empty array if query is empty
		if query == "" {
			c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
			return
		}

		var rows *sql.Rows
		var err error
		var args []interface{}

		if categoryID == "" {
			// Search without category filter
			sqlQuery := "SELECT id, name, category_id, icon_id FROM item_master WHERE name ILIKE $1 ORDER BY name LIMIT 20"
			args = []interface{}{"%" + query + "%"}
			rows, err = db.Query(sqlQuery, args...)
		} else {
			// Search with category filter
			sqlQuery := "SELECT id, name, category_id, icon_id FROM item_master WHERE name ILIKE $1 AND category_id = $2 ORDER BY name LIMIT 20"
			args = []interface{}{"%" + query + "%", categoryID}
			rows, err = db.Query(sqlQuery, args...)
		}

		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}
		defer rows.Close()

		var items []gin.H
		for rows.Next() {
			var id, name, categoryID string
			var iconID sql.NullString
			rows.Scan(&id, &name, &categoryID, &iconID)

			// Build iconUrl if icon_id is not null
			var iconURL *string
			if iconID.Valid {
				u := "/static/icons/" + iconID.String + ".png"
				iconURL = &u
			}

			items = append(items, gin.H{
				"id":         id,
				"name":       name,
				"categoryId": categoryID,
				"iconUrl":    iconURL,
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}
