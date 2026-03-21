package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jym/lincle/internal/domain"
	"github.com/jym/lincle/internal/guard"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
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

func handleCreateListing(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req createListingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
			})
			return
		}

		// Image ownership check
		for _, imgID := range req.ImageIDs {
			owned, err := repo.CheckImageOwnership(ctx, imgID, userID)
			if err != nil || !owned {
				c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "본인이 업로드한 이미지만 사용할 수 있습니다."}})
				return
			}
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

		if err := repo.InsertListing(ctx, &repository.InsertListingParams{
			ID:           id,
			ListingType:  req.ListingType,
			AuthorUserID: userID,
			ServerID:     req.ServerID,
			CategoryID:   req.CategoryID,
			ItemName:     req.ItemName,
			Title:        req.Title,
			Description:  req.Description,
			PriceType:    req.PriceType,
			PriceAmount:  req.PriceAmount,
			Quantity:     req.Quantity,
			Enhancement:  req.Enhancement,
			OptionsText:  req.OptionsText,
			TradeMethod:  req.TradeMethod,
			MeetingArea:  req.MeetingArea,
			TimeText:     req.TimeText,
			Now:          now,
		}); err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "매물 등록에 실패했습니다."},
			})
			return
		}

		// Record status history
		repo.InsertStatusHistory(ctx, &repository.InsertStatusHistoryParams{
			ID:            uuid.New().String(),
			EntityType:    "listing",
			EntityID:      id,
			ToStatus:      "available",
			ChangedByUser: userID,
			CreatedAt:     now,
		})

		c.JSON(http.StatusCreated, gin.H{
			"listingId":  id,
			"status":     "available",
			"visibility": "public",
			"createdAt":  now.Format(time.RFC3339),
		})
	}
}

func handleListListings(repo repository.ListingRepo) gin.HandlerFunc {
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

		items, err := repo.ListListings(c.Request.Context(), repository.ListingFilter{
			Query:       q,
			ServerID:    serverID,
			CategoryID:  categoryID,
			ListingType: listingType,
			Status:      status,
			Sort:        sort,
			Limit:       limit + 1,
			Cursor:      cursor,
		})
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."},
			})
			return
		}

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

		var result []listingItem
		for _, item := range items {
			li := listingItem{
				ListingID:      item.ListingID,
				ListingType:    item.ListingType,
				Title:          item.Title,
				ItemName:       item.ItemName,
				PriceType:      item.PriceType,
				PriceAmount:    item.PriceAmount,
				EnhancementLvl: item.EnhancementLvl,
				ServerID:       item.ServerID,
				ServerName:     item.ServerName,
				Status:         item.Status,
				TradeMethod:    item.TradeMethod,
				ViewCount:      item.ViewCount,
				FavoriteCount:  item.FavoriteCount,
				ChatCount:      item.ChatCount,
				LastActivityAt: item.LastActivityAt.Format(time.RFC3339),
				CreatedAt:      item.CreatedAt.Format(time.RFC3339Nano),
				Author: gin.H{
					"userId":        item.AuthorID,
					"nickname":      item.AuthorNickname,
					"trustBadge":    item.TrustBadge,
					"responseBadge": item.ResponseBadge,
				},
			}
			if item.IconID != nil {
				url := "/static/icons/" + *item.IconID + ".png"
				li.IconURL = &url
			}
			result = append(result, li)
		}

		hasMore := len(result) > limit
		if hasMore {
			result = result[:limit]
		}

		var nextCursor *string
		if hasMore && len(result) > 0 {
			cr := result[len(result)-1].CreatedAt
			nextCursor = &cr
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
			"cursor": gin.H{
				"next":    nextCursor,
				"hasMore": hasMore,
			},
		})
	}
}

func handleGetListing(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		// Increment view count
		repo.IncrementViewCount(ctx, id)

		detail, err := repo.GetListing(ctx, id)
		if err != nil || detail == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."},
			})
			return
		}

		isOwner := userID == detail.AuthorID
		var isFavorited bool
		if userID != "" {
			isFavorited, _ = repo.IsFavorited(ctx, userID, id)
		}

		// Available actions
		var actions []string
		if isOwner {
			actions = []string{"edit", "change_status"}
		} else if userID != "" {
			switch domain.ListingStatus(detail.Status) {
			case domain.ListingAvailable, domain.ListingReserved:
				actions = []string{"start_chat", "favorite", "report"}
			default:
				actions = []string{"report"}
			}
		}

		var iconURL *string
		if detail.IconID != nil {
			u := "/static/icons/" + *detail.IconID + ".png"
			iconURL = &u
		}

		c.JSON(http.StatusOK, gin.H{
			"listingId":                detail.ID,
			"listingType":              detail.ListingType,
			"title":                    detail.Title,
			"itemName":                 detail.ItemName,
			"description":              detail.Description,
			"priceType":                detail.PriceType,
			"priceAmount":              detail.PriceAmount,
			"quantity":                 detail.Quantity,
			"enhancementLevel":         detail.Enhancement,
			"optionsText":              detail.OptionsText,
			"serverId":                 detail.ServerID,
			"serverName":               detail.ServerName,
			"categoryId":               detail.CategoryID,
			"categoryName":             detail.CategoryName,
			"status":                   detail.Status,
			"visibility":               detail.Visibility,
			"tradeMethod":              detail.TradeMethod,
			"preferredMeetingAreaText":  detail.MeetingArea,
			"availableTimeText":         detail.TimeText,
			"iconUrl":                  iconURL,
			"author": gin.H{
				"userId":              detail.AuthorID,
				"nickname":            detail.AuthorNickname,
				"trustBadge":          detail.TrustBadge,
				"responseBadge":       detail.ResponseBadge,
				"completedTradeCount": detail.TradeCount,
			},
			"viewCount":           detail.ViewCount,
			"favoriteCount":       detail.FavoriteCount,
			"chatCount":           detail.ChatCount,
			"isFavorited":         isFavorited,
			"isOwner":             isOwner,
			"availableActions":    actions,
			"reservedChatRoomId":  detail.ReservedChatID,
			"lastActivityAt":      detail.LastActivityAt.Format(time.RFC3339),
			"createdAt":           detail.CreatedAt.Format(time.RFC3339),
			"updatedAt":           detail.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func handleUpdateListing(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		// Verify ownership
		owner, err := repo.GetListingOwnerAndStatus(ctx, id)
		if err != nil || owner == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		if owner.AuthorUserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "본인 매물만 수정할 수 있습니다."}})
			return
		}
		if owner.Status == "completed" || owner.Status == "cancelled" {
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

		fields := repository.ListingUpdateFields{
			Title:       req.Title,
			Description: req.Description,
			PriceType:   req.PriceType,
			PriceAmount: req.PriceAmount,
			Quantity:    req.Quantity,
			Enhancement: req.Enhancement,
			OptionsText: req.OptionsText,
			TradeMethod: req.TradeMethod,
			MeetingArea: req.MeetingArea,
			TimeText:    req.TimeText,
		}

		// Check if any fields are set
		if req.Title == nil && req.Description == nil && req.PriceType == nil && req.PriceAmount == nil &&
			req.Quantity == nil && req.Enhancement == nil && req.OptionsText == nil && req.TradeMethod == nil &&
			req.MeetingArea == nil && req.TimeText == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": "수정할 필드가 없습니다."}})
			return
		}

		if err := repo.UpdateListing(ctx, id, fields); err != nil {
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

func handleChangeListingStatus(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		var req changeStatusRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		owner, err := repo.GetListingOwnerAndStatus(ctx, id)
		if err != nil || owner == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		if owner.AuthorUserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "본인 매물만 상태를 변경할 수 있습니다."}})
			return
		}

		currentStatus := domain.ListingStatus(owner.Status)

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
		repo.UpdateListingStatus(ctx, id, targetStatus, now)

		fromStr := string(currentStatus)
		repo.InsertStatusHistory(ctx, &repository.InsertStatusHistoryParams{
			ID:            uuid.New().String(),
			EntityType:    "listing",
			EntityID:      id,
			FromStatus:    &fromStr,
			ToStatus:      string(targetStatus),
			ChangedByUser: userID,
			ReasonCode:    req.ReasonCode,
			CreatedAt:     now,
		})

		c.JSON(http.StatusOK, gin.H{
			"listingId": id,
			"status":    targetStatus,
			"updatedAt": now.Format(time.RFC3339),
		})
	}
}

func handleFavoriteListing(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)
		ctx := c.Request.Context()

		// Verify listing exists
		exists, err := repo.ListingExists(ctx, id)
		if err != nil || !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
			return
		}
		repo.AddFavorite(ctx, uuid.New().String(), userID, id)
		c.Status(http.StatusNoContent)
	}
}

func handleUnfavoriteListing(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := middleware.GetUserID(c)
		repo.RemoveFavorite(c.Request.Context(), userID, id)
		c.Status(http.StatusNoContent)
	}
}

func handleMyListings(repo repository.ListingRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		status := c.Query("status")

		var statusPtr *string
		if status != "" {
			statusPtr = &status
		}

		items, err := repo.ListMyListings(c.Request.Context(), userID, statusPtr)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류가 발생했습니다."}})
			return
		}

		var result []gin.H
		for _, item := range items {
			entry := gin.H{
				"listingId": item.ListingID, "listingType": item.ListingType, "title": item.Title, "itemName": item.ItemName,
				"priceType": item.PriceType, "priceAmount": item.PriceAmount, "status": item.Status,
				"enhancementLevel": item.EnhancementLvl, "serverName": item.ServerName,
				"viewCount": item.ViewCount, "favoriteCount": item.FavoriteCount, "chatCount": item.ChatCount,
				"createdAt": item.CreatedAt.Format(time.RFC3339),
				"author": gin.H{"userId": item.AuthorID, "nickname": item.AuthorNickname},
			}
			if item.IconID != nil {
				entry["iconUrl"] = "/static/icons/" + *item.IconID + ".png"
			}
			result = append(result, entry)
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
