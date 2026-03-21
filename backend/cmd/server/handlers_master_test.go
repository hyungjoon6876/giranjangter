package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jym/lincle/internal/repository"
	"github.com/jym/lincle/internal/repository/mock"
)

func TestListServers_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockMasterRepo{
		ListServersFn: func(ctx context.Context) ([]repository.ServerItem, error) {
			return []repository.ServerItem{
				{ServerID: "server-1", ServerName: "데포로쥬"},
				{ServerID: "server-2", ServerName: "케레니스"},
				{ServerID: "server-3", ServerName: "살바도르"},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/servers", listServers(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/servers", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			ServerID   string `json:"serverId"`
			ServerName string `json:"serverName"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 3 {
		t.Errorf("len(data) = %d, want 3", len(resp.Data))
	}
	if resp.Data[0].ServerID != "server-1" {
		t.Errorf("data[0].serverId = %q, want %q", resp.Data[0].ServerID, "server-1")
	}
	if resp.Data[0].ServerName != "데포로쥬" {
		t.Errorf("data[0].serverName = %q, want %q", resp.Data[0].ServerName, "데포로쥬")
	}
}

func TestListServers_EmptyList_Returns200(t *testing.T) {
	mockRepo := &mock.MockMasterRepo{
		ListServersFn: func(ctx context.Context) ([]repository.ServerItem, error) {
			return []repository.ServerItem{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/servers", listServers(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/servers", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []interface{} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 0 {
		t.Errorf("len(data) = %d, want 0", len(resp.Data))
	}
}

func TestListCategories_Success_Returns200(t *testing.T) {
	mockRepo := &mock.MockMasterRepo{
		ListCategoriesFn: func(ctx context.Context) ([]repository.CategoryItem, error) {
			return []repository.CategoryItem{
				{CategoryID: "weapon", CategoryName: "무기", ParentID: nil},
				{CategoryID: "sword", CategoryName: "검", ParentID: strPtr("weapon")},
				{CategoryID: "armor", CategoryName: "방어구", ParentID: nil},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/categories", listCategories(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/categories", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			CategoryID   string  `json:"categoryId"`
			CategoryName string  `json:"categoryName"`
			ParentID     *string `json:"parentId"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 3 {
		t.Errorf("len(data) = %d, want 3", len(resp.Data))
	}
	if resp.Data[0].CategoryID != "weapon" {
		t.Errorf("data[0].categoryId = %q, want %q", resp.Data[0].CategoryID, "weapon")
	}
	if resp.Data[0].CategoryName != "무기" {
		t.Errorf("data[0].categoryName = %q, want %q", resp.Data[0].CategoryName, "무기")
	}
	if resp.Data[0].ParentID != nil {
		t.Errorf("data[0].parentId = %v, want nil", resp.Data[0].ParentID)
	}
	if resp.Data[1].ParentID == nil || *resp.Data[1].ParentID != "weapon" {
		t.Errorf("data[1].parentId = %v, want %q", resp.Data[1].ParentID, "weapon")
	}
}

func TestSearchItems_WithQuery_Returns200(t *testing.T) {
	mockRepo := &mock.MockMasterRepo{
		SearchItemsFn: func(ctx context.Context, query string, categoryID *string) ([]repository.ItemSearchResult, error) {
			return []repository.ItemSearchResult{
				{
					ID:             "item-1",
					Name:           "그라나도 에스파다",
					CategoryID:     "sword",
					IconID:         strPtr("sword_icon"),
					SubCategory:    "검",
					OptionText:     strPtr("STR+5"),
					IsEnchantable:  1,
					SafeEnchantLvl: 3,
					MaxEnchantLvl:  9,
				},
				{
					ID:             "item-2",
					Name:           "엘리트 그라나도",
					CategoryID:     "sword",
					IconID:         nil,
					SubCategory:    "검",
					OptionText:     nil,
					IsEnchantable:  0,
					SafeEnchantLvl: 0,
					MaxEnchantLvl:  0,
				},
			}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/items/search", searchItems(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/items/search?q=그라나도", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			ID               string  `json:"id"`
			Name             string  `json:"name"`
			CategoryID       string  `json:"categoryId"`
			IconURL          *string `json:"iconUrl"`
			SubCategory      string  `json:"subCategory"`
			OptionText       *string `json:"optionText"`
			IsEnchantable    bool    `json:"isEnchantable"`
			SafeEnchantLevel int     `json:"safeEnchantLevel"`
			MaxEnchantLevel  int     `json:"maxEnchantLevel"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 2 {
		t.Errorf("len(data) = %d, want 2", len(resp.Data))
	}
	if resp.Data[0].ID != "item-1" {
		t.Errorf("data[0].id = %q, want %q", resp.Data[0].ID, "item-1")
	}
	if resp.Data[0].Name != "그라나도 에스파다" {
		t.Errorf("data[0].name = %q, want %q", resp.Data[0].Name, "그라나도 에스파다")
	}
	if resp.Data[0].IconURL == nil || *resp.Data[0].IconURL != "/static/icons/sword_icon.png" {
		t.Errorf("data[0].iconUrl = %v, want %q", resp.Data[0].IconURL, "/static/icons/sword_icon.png")
	}
	if !resp.Data[0].IsEnchantable {
		t.Error("data[0].isEnchantable = false, want true")
	}
	if resp.Data[0].SafeEnchantLevel != 3 {
		t.Errorf("data[0].safeEnchantLevel = %d, want 3", resp.Data[0].SafeEnchantLevel)
	}

	if resp.Data[1].IconURL != nil {
		t.Errorf("data[1].iconUrl = %v, want nil", resp.Data[1].IconURL)
	}
	if resp.Data[1].IsEnchantable {
		t.Error("data[1].isEnchantable = true, want false")
	}
}

func TestSearchItems_WithCategoryFilter_Returns200(t *testing.T) {
	mockRepo := &mock.MockMasterRepo{
		SearchItemsFn: func(ctx context.Context, query string, categoryID *string) ([]repository.ItemSearchResult, error) {
			if categoryID != nil && *categoryID == "sword" {
				return []repository.ItemSearchResult{
					{
						ID:          "item-sword",
						Name:        "검류 아이템",
						CategoryID:  "sword",
						SubCategory: "검",
					},
				}, nil
			}
			return []repository.ItemSearchResult{}, nil
		},
	}

	r := setupRouter()
	r.GET("/api/v1/items/search", searchItems(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/items/search?categoryId=sword", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			CategoryID  string `json:"categoryId"`
			SubCategory string `json:"subCategory"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 1 {
		t.Errorf("len(data) = %d, want 1", len(resp.Data))
	}
	if resp.Data[0].ID != "item-sword" {
		t.Errorf("data[0].id = %q, want %q", resp.Data[0].ID, "item-sword")
	}
}

func TestSearchItems_EmptyQuery_Returns200EmptyArray(t *testing.T) {
	mockRepo := &mock.MockMasterRepo{}

	r := setupRouter()
	r.GET("/api/v1/items/search", searchItems(mockRepo))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/items/search", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp struct {
		Data []interface{} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp.Data) != 0 {
		t.Errorf("len(data) = %d, want 0", len(resp.Data))
	}
}