package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/domain"
	"github.com/jym/lincle/internal/guard"
	"github.com/jym/lincle/internal/middleware"
)

type createListingRequest struct {
	ListingType  string  `json:"listingType" binding:"required,oneof=sell buy"`
	ServerID     string  `json:"serverId" binding:"required"`
	CategoryID   string  `json:"categoryId" binding:"required"`
	ItemName     string  `json:"itemName" binding:"required,min=1,max=100"`
	Title        string  `json:"title" binding:"required,min=2,max=100"`
	Description  string  `json:"description" binding:"required,min=10,max=2000"`
	PriceType    string  `json:"priceType" binding:"required,oneof=fixed negotiable offer"`
	PriceAmount  *int64  `json:"priceAmount"`
	Quantity     int     `json:"quantity" binding:"required,min=1"`
	Enhancement  *int    `json:"enhancementLevel"`
	OptionsText  *string `json:"optionsText"`
	TradeMethod  string  `json:"tradeMethod" binding:"required,oneof=in_game offline_pc_bang either"`
	MeetingArea  *string `json:"preferredMeetingAreaText"`
	TimeText     *string `json:"availableTimeText"`
	ImageIDs     []string `json:"imageIds"`
}

func handleCreateListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		var req createListingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		// Price validation
		if req.PriceType != "offer" && (req.PriceAmount == nil || *req.PriceAmount <= 0) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": "가격을 입력해주세요."},
			})
			return
		}

		id := uuid.New().String()
		now := time.Now().UTC()

		_, err := db.Exec(`
			INSERT INTO listings (id, listing_type, author_user_id, server_id, category_id,
				item_name, title, description, price_type, price_amount, quantity,
				enhancement_level, options_text, trade_method,
				preferred_meeting_area_text, available_time_text,
				status, visibility, last_activity_at, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, 'available', 'public', $17, $18, $19)`,
			id, req.ListingType, userID, req.ServerID, req.CategoryID,
			req.ItemName, req.Title, req.Description, req.PriceType, req.PriceAmount, req.Quantity,
			req.Enhancement, req.OptionsText, req.TradeMethod,
			req.MeetingArea, req.TimeText,
			now, now, now,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": fmt.Sprintf("매물 등록 실패: %v", err)},
			})
			return
		}

		// Record status history
		db.Exec(`INSERT INTO status_history (id, entity_type, entity_id, to_status, changed_by_user_id, created_at)
			VALUES ($1, 'listing', $2, 'available', $3, $4)`, uuid.New().String(), id, userID, now)

		c.JSON(http.StatusCreated, gin.H{
			"listingId":  id,
			"status":     "available",
			"visibility": "public",
			"createdAt":  now.Format(time.RFC3339),
		})
	}
}

func handleListListings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Query("q")
		serverID := c.Query("serverId")
		categoryID := c.Query("categoryId")
		listingType := c.Query("listingType")
		status := c.DefaultQuery("status", "available")
		sort := c.DefaultQuery("sort", "recent")
		limitStr := c.DefaultQuery("limit", "20")
		cursor := c.Query("cursor")

		limit, _ := strconv.Atoi(limitStr)
		if limit <= 0 || limit > 100 {
			limit = 20
		}

		query := `SELECT l.id, l.listing_type, l.title, l.item_name,
			l.price_type, l.price_amount, l.enhancement_level,
			l.server_id, s.name as server_name,
			l.status, l.trade_method, l.view_count, l.favorite_count, l.chat_count,
			l.last_activity_at, l.created_at,
			p.user_id as author_id, p.nickname, p.trust_badge, p.response_badge,
			im.icon_id
			FROM listings l
			JOIN servers s ON l.server_id = s.id
			JOIN user_profiles p ON l.author_user_id = p.user_id
			LEFT JOIN item_master im ON im.name = l.item_name
			WHERE l.deleted_at IS NULL AND l.visibility = 'public'`

		args := []interface{}{}
		paramIdx := 1

		if status != "" {
			query += fmt.Sprintf(" AND l.status = $%d", paramIdx)
			args = append(args, status)
			paramIdx++
		}
		if serverID != "" {
			query += fmt.Sprintf(" AND l.server_id = $%d", paramIdx)
			args = append(args, serverID)
			paramIdx++
		}
		if categoryID != "" {
			query += fmt.Sprintf(" AND (l.category_id = $%d OR l.category_id IN (SELECT id FROM categories WHERE parent_id = $%d))", paramIdx, paramIdx+1)
			args = append(args, categoryID, categoryID)
			paramIdx += 2
		}
		if listingType != "" {
			query += fmt.Sprintf(" AND l.listing_type = $%d", paramIdx)
			args = append(args, listingType)
			paramIdx++
		}
		if q != "" {
			query += fmt.Sprintf(" AND (l.title LIKE $%d OR l.item_name LIKE $%d)", paramIdx, paramIdx+1)
			pattern := "%" + q + "%"
			args = append(args, pattern, pattern)
			paramIdx += 2
		}
		if cursor != "" {
			query += fmt.Sprintf(" AND l.created_at < $%d", paramIdx)
			args = append(args, cursor)
			paramIdx++
		}

		switch sort {
		case "price_asc":
			query += " ORDER BY l.price_amount ASC, l.created_at DESC"
		case "price_desc":
			query += " ORDER BY l.price_amount DESC, l.created_at DESC"
		case "popular":
			query += " ORDER BY l.favorite_count DESC, l.created_at DESC"
		default:
			query += " ORDER BY l.last_activity_at DESC"
		}

		query += fmt.Sprintf(" LIMIT %d", limit+1)

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()},
			})
			return
		}
		defer rows.Close()

		type listingItem struct {
			ListingID       string  `json:"listingId"`
			ListingType     string  `json:"listingType"`
			Title           string  `json:"title"`
			ItemName        string  `json:"itemName"`
			PriceType       string  `json:"priceType"`
			PriceAmount     *int64  `json:"priceAmount"`
			EnhancementLvl  *int    `json:"enhancementLevel"`
			ServerID        string  `json:"serverId"`
			ServerName      string  `json:"serverName"`
			Status          string  `json:"status"`
			TradeMethod     string  `json:"tradeMethod"`
			ViewCount       int     `json:"viewCount"`
			FavoriteCount   int     `json:"favoriteCount"`
			ChatCount       int     `json:"chatCount"`
			LastActivityAt  string  `json:"lastActivityAt"`
			CreatedAt       string  `json:"createdAt"`
			Author          gin.H   `json:"author"`
			IconURL         *string `json:"iconUrl"`
		}

		var items []listingItem
		for rows.Next() {
			var item listingItem
			var authorID, nickname, trustBadge, responseBadge string
			var lastActivity, created time.Time
			var iconID *string

			err := rows.Scan(&item.ListingID, &item.ListingType, &item.Title, &item.ItemName,
				&item.PriceType, &item.PriceAmount, &item.EnhancementLvl,
				&item.ServerID, &item.ServerName,
				&item.Status, &item.TradeMethod, &item.ViewCount, &item.FavoriteCount, &item.ChatCount,
				&lastActivity, &created,
				&authorID, &nickname, &trustBadge, &responseBadge,
				&iconID)
			if err != nil {
				continue
			}
			item.LastActivityAt = lastActivity.Format(time.RFC3339)
			item.CreatedAt = created.Format(time.RFC3339)
			item.Author = gin.H{
				"userId":        authorID,
				"nickname":      nickname,
				"trustBadge":    trustBadge,
				"responseBadge": responseBadge,
			}
			if iconID != nil {
				url := "/static/icons/" + *iconID + ".png"
				item.IconURL = &url
			}
			items = append(items, item)
		}

		hasMore := len(items) > limit
		if hasMore {
			items = items[:limit]
		}

		var nextCursor *string
		if hasMore && len(items) > 0 {
			c := items[len(items)-1].CreatedAt
			nextCursor = &c
		}

		c.JSON(http.StatusOK, gin.H{
			"data": items,
			"cursor": gin.H{
				"next":    nextCursor,
				"hasMore": hasMore,
			},
		})
	}
}

func handleGetListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)

		// Increment view count
		db.Exec("UPDATE listings SET view_count = view_count + 1 WHERE id = $1", id)

		var l struct {
			ID              string  `json:"listingId"`
			ListingType     string  `json:"listingType"`
			Title           string  `json:"title"`
			ItemName        string  `json:"itemName"`
			Description     string  `json:"description"`
			PriceType       string  `json:"priceType"`
			PriceAmount     *int64  `json:"priceAmount"`
			Quantity        int     `json:"quantity"`
			Enhancement     *int    `json:"enhancementLevel"`
			OptionsText     *string `json:"optionsText"`
			ServerID        string  `json:"serverId"`
			ServerName      string  `json:"serverName"`
			CategoryID      string  `json:"categoryId"`
			CategoryName    string  `json:"categoryName"`
			Status          string  `json:"status"`
			Visibility      string  `json:"visibility"`
			TradeMethod     string  `json:"tradeMethod"`
			MeetingArea     *string `json:"preferredMeetingAreaText"`
			TimeText        *string `json:"availableTimeText"`
			AuthorID        string  `json:"-"`
			Nickname        string  `json:"-"`
			TrustBadge      string  `json:"-"`
			ResponseBadge   string  `json:"-"`
			TradeCount      int     `json:"-"`
			ViewCount       int     `json:"viewCount"`
			FavoriteCount   int     `json:"favoriteCount"`
			ChatCount       int     `json:"chatCount"`
			ReservedChatID  *string `json:"reservedChatRoomId"`
			LastActivityAt  time.Time `json:"-"`
			CreatedAt       time.Time `json:"-"`
			UpdatedAt       time.Time `json:"-"`
		}

		var iconID *string
		err := db.QueryRow(`
			SELECT l.id, l.listing_type, l.title, l.item_name, l.description,
				l.price_type, l.price_amount, l.quantity, l.enhancement_level, l.options_text,
				l.server_id, s.name, l.category_id, c.name,
				l.status, l.visibility, l.trade_method,
				l.preferred_meeting_area_text, l.available_time_text,
				l.author_user_id, p.nickname, p.trust_badge, p.response_badge, p.completed_trade_count,
				l.view_count, l.favorite_count, l.chat_count,
				l.reserved_chat_room_id, l.last_activity_at, l.created_at, l.updated_at,
				im.icon_id
			FROM listings l
			JOIN servers s ON l.server_id = s.id
			JOIN categories c ON l.category_id = c.id
			JOIN user_profiles p ON l.author_user_id = p.user_id
			LEFT JOIN item_master im ON im.name = l.item_name
			WHERE l.id = $1 AND l.deleted_at IS NULL`, id,
		).Scan(&l.ID, &l.ListingType, &l.Title, &l.ItemName, &l.Description,
			&l.PriceType, &l.PriceAmount, &l.Quantity, &l.Enhancement, &l.OptionsText,
			&l.ServerID, &l.ServerName, &l.CategoryID, &l.CategoryName,
			&l.Status, &l.Visibility, &l.TradeMethod,
			&l.MeetingArea, &l.TimeText,
			&l.AuthorID, &l.Nickname, &l.TrustBadge, &l.ResponseBadge, &l.TradeCount,
			&l.ViewCount, &l.FavoriteCount, &l.ChatCount,
			&l.ReservedChatID, &l.LastActivityAt, &l.CreatedAt, &l.UpdatedAt,
			&iconID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."},
			})
			return
		}

		isOwner := userID == l.AuthorID
		var isFavorited bool
		if userID != "" {
			db.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = $1 AND listing_id = $2", userID, id).Scan(&isFavorited)
		}

		// Available actions
		var actions []string
		if isOwner {
			actions = []string{"edit", "change_status"}
		} else if userID != "" {
			switch domain.ListingStatus(l.Status) {
			case domain.ListingAvailable, domain.ListingReserved:
				actions = []string{"start_chat", "favorite", "report"}
			default:
				actions = []string{"report"}
			}
		}

		var iconURL *string
		if iconID != nil {
			u := "/static/icons/" + *iconID + ".png"
			iconURL = &u
		}

		c.JSON(http.StatusOK, gin.H{
			"listingId":                l.ID,
			"listingType":              l.ListingType,
			"title":                    l.Title,
			"itemName":                 l.ItemName,
			"description":              l.Description,
			"priceType":                l.PriceType,
			"priceAmount":              l.PriceAmount,
			"quantity":                 l.Quantity,
			"enhancementLevel":         l.Enhancement,
			"optionsText":              l.OptionsText,
			"serverId":                 l.ServerID,
			"serverName":               l.ServerName,
			"categoryId":               l.CategoryID,
			"categoryName":             l.CategoryName,
			"status":                   l.Status,
			"visibility":               l.Visibility,
			"tradeMethod":              l.TradeMethod,
			"preferredMeetingAreaText":  l.MeetingArea,
			"availableTimeText":         l.TimeText,
			"iconUrl":                  iconURL,
			"author": gin.H{
				"userId":              l.AuthorID,
				"nickname":            l.Nickname,
				"trustBadge":          l.TrustBadge,
				"responseBadge":       l.ResponseBadge,
				"completedTradeCount": l.TradeCount,
			},
			"viewCount":           l.ViewCount,
			"favoriteCount":       l.FavoriteCount,
			"chatCount":           l.ChatCount,
			"isFavorited":         isFavorited,
			"isOwner":             isOwner,
			"availableActions":    actions,
			"reservedChatRoomId":  l.ReservedChatID,
			"lastActivityAt":      l.LastActivityAt.Format(time.RFC3339),
			"createdAt":           l.CreatedAt.Format(time.RFC3339),
			"updatedAt":           l.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func handleUpdateListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)

		// Verify ownership
		var authorID, status string
		err := db.QueryRow("SELECT author_user_id, status FROM listings WHERE id = $1 AND deleted_at IS NULL", id).Scan(&authorID, &status)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		if authorID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "본인 매물만 수정할 수 있습니다."}})
			return
		}
		if status == "completed" || status == "cancelled" {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "종결된 매물은 수정할 수 없습니다."}})
			return
		}

		var req struct {
			Title       *string `json:"title"`
			Description *string `json:"description"`
			PriceType   *string `json:"priceType"`
			PriceAmount *int64  `json:"priceAmount"`
			Quantity    *int    `json:"quantity"`
			Enhancement *int    `json:"enhancementLevel"`
			OptionsText *string `json:"optionsText"`
			TradeMethod *string `json:"tradeMethod"`
			MeetingArea *string `json:"preferredMeetingAreaText"`
			TimeText    *string `json:"availableTimeText"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		// Build SET clause dynamically with parameterized values only
		setClauses := []string{}
		args := []interface{}{}
		paramIdx := 1

		type field struct {
			col string
			val interface{}
			set bool
		}
		fields := []field{
			{"title", req.Title, req.Title != nil},
			{"description", req.Description, req.Description != nil},
			{"price_type", req.PriceType, req.PriceType != nil},
			{"price_amount", req.PriceAmount, req.PriceAmount != nil},
			{"quantity", req.Quantity, req.Quantity != nil},
			{"enhancement_level", req.Enhancement, req.Enhancement != nil},
			{"options_text", req.OptionsText, req.OptionsText != nil},
			{"trade_method", req.TradeMethod, req.TradeMethod != nil},
			{"preferred_meeting_area_text", req.MeetingArea, req.MeetingArea != nil},
			{"available_time_text", req.TimeText, req.TimeText != nil},
		}
		for _, f := range fields {
			if f.set {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", f.col, paramIdx))
				args = append(args, f.val)
				paramIdx++
			}
		}

		if len(setClauses) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "수정할 필드가 없습니다."}})
			return
		}

		query := "UPDATE listings SET " + strings.Join(setClauses, ", ") + fmt.Sprintf(", updated_at = NOW() WHERE id = $%d", paramIdx)
		args = append(args, id)

		if _, err := db.Exec(query, args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 수정 실패"}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "수정 완료"})
	}
}

type changeStatusRequest struct {
	Action     string  `json:"action" binding:"required"`
	ReasonCode *string `json:"reasonCode"`
}

func handleChangeListingStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)

		var req changeStatusRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		var authorID string
		var currentStatus domain.ListingStatus
		err := db.QueryRow("SELECT author_user_id, status FROM listings WHERE id = $1 AND deleted_at IS NULL", id).Scan(&authorID, &currentStatus)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		if authorID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "본인 매물만 상태를 변경할 수 있습니다."}})
			return
		}

		actionToStatus := map[string]domain.ListingStatus{
			"reserve":     domain.ListingReserved,
			"unreserve":   domain.ListingAvailable,
			"start_trade": domain.ListingPendingTrade,
			"complete":    domain.ListingCompleted,
			"cancel":      domain.ListingCancelled,
			"reopen":      domain.ListingAvailable,
		}

		targetStatus, ok := actionToStatus[req.Action]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "잘못된 액션입니다."}})
			return
		}

		if err := guard.ValidateListingTransition(currentStatus, targetStatus); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": gin.H{"code": "INVALID_TRANSITION", "message": err.Error()}})
			return
		}

		now := time.Now().UTC()
		db.Exec("UPDATE listings SET status = $1, updated_at = $2, last_activity_at = $3 WHERE id = $4", targetStatus, now, now, id)
		db.Exec(`INSERT INTO status_history (id, entity_type, entity_id, from_status, to_status, changed_by_user_id, reason_code, created_at)
			VALUES ($1, 'listing', $2, $3, $4, $5, $6, $7)`,
			uuid.New().String(), id, currentStatus, targetStatus, userID, req.ReasonCode, now)

		c.JSON(http.StatusOK, gin.H{
			"listingId": id,
			"status":    targetStatus,
			"updatedAt": now.Format(time.RFC3339),
		})
	}
}

func handleFavoriteListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)

		_, err := db.Exec("INSERT INTO favorites (id, user_id, listing_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
			uuid.New().String(), userID, id)
		if err == nil {
			db.Exec("UPDATE listings SET favorite_count = favorite_count + 1 WHERE id = $1", id)
		}
		c.Status(http.StatusNoContent)
	}
}

func handleUnfavoriteListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)

		res, _ := db.Exec("DELETE FROM favorites WHERE user_id = $1 AND listing_id = $2", userID, id)
		if n, _ := res.RowsAffected(); n > 0 {
			db.Exec("UPDATE listings SET favorite_count = GREATEST(0, favorite_count - 1) WHERE id = $1", id)
		}
		c.Status(http.StatusNoContent)
	}
}

func handleMyListings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		status := c.Query("status")

		query := `SELECT id, listing_type, title, item_name, price_type, price_amount,
			status, view_count, favorite_count, chat_count, created_at
			FROM listings WHERE author_user_id = $1 AND deleted_at IS NULL`
		args := []interface{}{userID}
		paramIdx := 2

		if status != "" {
			query += fmt.Sprintf(" AND status = $%d", paramIdx)
			args = append(args, status)
			paramIdx++
		}
		query += " ORDER BY created_at DESC LIMIT 50"

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": err.Error()}})
			return
		}
		defer rows.Close()

		var items []gin.H
		for rows.Next() {
			var id, lt, title, itemName, pt, st string
			var price *int64
			var vc, fc, cc int
			var created time.Time
			rows.Scan(&id, &lt, &title, &itemName, &pt, &price, &st, &vc, &fc, &cc, &created)
			items = append(items, gin.H{
				"listingId": id, "listingType": lt, "title": title, "itemName": itemName,
				"priceType": pt, "priceAmount": price, "status": st,
				"viewCount": vc, "favoriteCount": fc, "chatCount": cc,
				"createdAt": created.Format(time.RFC3339),
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}
