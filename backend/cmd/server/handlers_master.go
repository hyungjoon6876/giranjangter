package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listServers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name FROM servers WHERE is_active = 1 ORDER BY sort_order")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
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
