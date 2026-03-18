package mock

import (
	"context"

	"github.com/jym/lincle/internal/repository"
)

// MockUploadRepo is a function-field based mock for repository.UploadRepo.
type MockUploadRepo struct {
	InsertImageFn func(ctx context.Context, params *repository.InsertImageParams) error
}

func (m *MockUploadRepo) InsertImage(ctx context.Context, params *repository.InsertImageParams) error {
	if m.InsertImageFn != nil {
		return m.InsertImageFn(ctx, params)
	}
	return nil
}

// Ensure compile-time interface satisfaction.
var _ repository.UploadRepo = (*MockUploadRepo)(nil)
