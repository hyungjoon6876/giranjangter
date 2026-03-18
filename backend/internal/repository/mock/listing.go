package mock

import (
	"context"
	"time"

	"github.com/jym/lincle/internal/domain"
	"github.com/jym/lincle/internal/repository"
)

// MockListingRepo is a function-field based mock for repository.ListingRepo.
type MockListingRepo struct {
	CheckImageOwnershipFn    func(ctx context.Context, imageID, userID string) (bool, error)
	InsertListingFn          func(ctx context.Context, listing *repository.InsertListingParams) error
	InsertStatusHistoryFn    func(ctx context.Context, params *repository.InsertStatusHistoryParams) error
	ListListingsFn           func(ctx context.Context, filter repository.ListingFilter) ([]repository.ListingListItem, error)
	GetListingFn             func(ctx context.Context, listingID string) (*repository.ListingDetail, error)
	IncrementViewCountFn     func(ctx context.Context, listingID string) error
	IsFavoritedFn            func(ctx context.Context, userID, listingID string) (bool, error)
	GetListingOwnerAndStatusFn func(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error)
	UpdateListingFn          func(ctx context.Context, listingID string, fields repository.ListingUpdateFields) error
	UpdateListingStatusFn    func(ctx context.Context, listingID string, status domain.ListingStatus, now time.Time) error
	ListingExistsFn          func(ctx context.Context, listingID string) (bool, error)
	AddFavoriteFn            func(ctx context.Context, id, userID, listingID string) error
	RemoveFavoriteFn         func(ctx context.Context, userID, listingID string) error
	ListMyListingsFn         func(ctx context.Context, userID string, status *string) ([]repository.MyListingItem, error)
}

func (m *MockListingRepo) CheckImageOwnership(ctx context.Context, imageID, userID string) (bool, error) {
	if m.CheckImageOwnershipFn != nil {
		return m.CheckImageOwnershipFn(ctx, imageID, userID)
	}
	return false, nil
}

func (m *MockListingRepo) InsertListing(ctx context.Context, listing *repository.InsertListingParams) error {
	if m.InsertListingFn != nil {
		return m.InsertListingFn(ctx, listing)
	}
	return nil
}

func (m *MockListingRepo) InsertStatusHistory(ctx context.Context, params *repository.InsertStatusHistoryParams) error {
	if m.InsertStatusHistoryFn != nil {
		return m.InsertStatusHistoryFn(ctx, params)
	}
	return nil
}

func (m *MockListingRepo) ListListings(ctx context.Context, filter repository.ListingFilter) ([]repository.ListingListItem, error) {
	if m.ListListingsFn != nil {
		return m.ListListingsFn(ctx, filter)
	}
	return nil, nil
}

func (m *MockListingRepo) GetListing(ctx context.Context, listingID string) (*repository.ListingDetail, error) {
	if m.GetListingFn != nil {
		return m.GetListingFn(ctx, listingID)
	}
	return nil, nil
}

func (m *MockListingRepo) IncrementViewCount(ctx context.Context, listingID string) error {
	if m.IncrementViewCountFn != nil {
		return m.IncrementViewCountFn(ctx, listingID)
	}
	return nil
}

func (m *MockListingRepo) IsFavorited(ctx context.Context, userID, listingID string) (bool, error) {
	if m.IsFavoritedFn != nil {
		return m.IsFavoritedFn(ctx, userID, listingID)
	}
	return false, nil
}

func (m *MockListingRepo) GetListingOwnerAndStatus(ctx context.Context, listingID string) (*repository.ListingOwnerStatus, error) {
	if m.GetListingOwnerAndStatusFn != nil {
		return m.GetListingOwnerAndStatusFn(ctx, listingID)
	}
	return nil, nil
}

func (m *MockListingRepo) UpdateListing(ctx context.Context, listingID string, fields repository.ListingUpdateFields) error {
	if m.UpdateListingFn != nil {
		return m.UpdateListingFn(ctx, listingID, fields)
	}
	return nil
}

func (m *MockListingRepo) UpdateListingStatus(ctx context.Context, listingID string, status domain.ListingStatus, now time.Time) error {
	if m.UpdateListingStatusFn != nil {
		return m.UpdateListingStatusFn(ctx, listingID, status, now)
	}
	return nil
}

func (m *MockListingRepo) ListingExists(ctx context.Context, listingID string) (bool, error) {
	if m.ListingExistsFn != nil {
		return m.ListingExistsFn(ctx, listingID)
	}
	return false, nil
}

func (m *MockListingRepo) AddFavorite(ctx context.Context, id, userID, listingID string) error {
	if m.AddFavoriteFn != nil {
		return m.AddFavoriteFn(ctx, id, userID, listingID)
	}
	return nil
}

func (m *MockListingRepo) RemoveFavorite(ctx context.Context, userID, listingID string) error {
	if m.RemoveFavoriteFn != nil {
		return m.RemoveFavoriteFn(ctx, userID, listingID)
	}
	return nil
}

func (m *MockListingRepo) ListMyListings(ctx context.Context, userID string, status *string) ([]repository.MyListingItem, error) {
	if m.ListMyListingsFn != nil {
		return m.ListMyListingsFn(ctx, userID, status)
	}
	return nil, nil
}

// Ensure compile-time interface satisfaction.
var _ repository.ListingRepo = (*MockListingRepo)(nil)
