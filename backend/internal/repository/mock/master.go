package mock

import (
	"context"

	"github.com/jym/lincle/internal/repository"
)

// MockMasterRepo is a function-field based mock for repository.MasterRepo.
type MockMasterRepo struct {
	ListServersFn    func(ctx context.Context) ([]repository.ServerItem, error)
	ListCategoriesFn func(ctx context.Context) ([]repository.CategoryItem, error)
	SearchItemsFn    func(ctx context.Context, query string, categoryID *string) ([]repository.ItemSearchResult, error)
}

func (m *MockMasterRepo) ListServers(ctx context.Context) ([]repository.ServerItem, error) {
	if m.ListServersFn != nil {
		return m.ListServersFn(ctx)
	}
	return nil, nil
}

func (m *MockMasterRepo) ListCategories(ctx context.Context) ([]repository.CategoryItem, error) {
	if m.ListCategoriesFn != nil {
		return m.ListCategoriesFn(ctx)
	}
	return nil, nil
}

func (m *MockMasterRepo) SearchItems(ctx context.Context, query string, categoryID *string) ([]repository.ItemSearchResult, error) {
	if m.SearchItemsFn != nil {
		return m.SearchItemsFn(ctx, query, categoryID)
	}
	return nil, nil
}

// Ensure compile-time interface satisfaction.
var _ repository.MasterRepo = (*MockMasterRepo)(nil)
