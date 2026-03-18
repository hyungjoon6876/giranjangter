package mock

import (
	"context"
	"time"

	"github.com/jym/lincle/internal/repository"
)

// MockAuthRepo is a function-field based mock for repository.AuthRepo.
type MockAuthRepo struct {
	FindUserByProviderFn      func(ctx context.Context, provider, providerKey string) (*repository.UserWithNickname, error)
	CreateUserWithProfileFn   func(ctx context.Context, userID, provider, providerKey, nickname string) error
	UpdateLastLoginFn         func(ctx context.Context, userID string, at time.Time) error
	StoreRefreshTokenFn       func(ctx context.Context, id, userID, tokenHash string, expiresAt time.Time) error
	FindRefreshTokenFn        func(ctx context.Context, tokenHash string) (string, error)
	GetAccountStatusFn        func(ctx context.Context, userID string) (string, error)
	RotateRefreshTokenFn      func(ctx context.Context, oldTokenID, newTokenID, userID, newTokenHash string, expiresAt time.Time) error
	DeleteRefreshTokensByUserFn func(ctx context.Context, userID string) error
	GetUserProfileFn          func(ctx context.Context, userID string) (*repository.FullUserProfile, error)
	UpdateProfileFn           func(ctx context.Context, userID string, fields repository.ProfileUpdateFields) error
}

func (m *MockAuthRepo) FindUserByProvider(ctx context.Context, provider, providerKey string) (*repository.UserWithNickname, error) {
	if m.FindUserByProviderFn != nil {
		return m.FindUserByProviderFn(ctx, provider, providerKey)
	}
	return nil, nil
}

func (m *MockAuthRepo) CreateUserWithProfile(ctx context.Context, userID, provider, providerKey, nickname string) error {
	if m.CreateUserWithProfileFn != nil {
		return m.CreateUserWithProfileFn(ctx, userID, provider, providerKey, nickname)
	}
	return nil
}

func (m *MockAuthRepo) UpdateLastLogin(ctx context.Context, userID string, at time.Time) error {
	if m.UpdateLastLoginFn != nil {
		return m.UpdateLastLoginFn(ctx, userID, at)
	}
	return nil
}

func (m *MockAuthRepo) StoreRefreshToken(ctx context.Context, id, userID, tokenHash string, expiresAt time.Time) error {
	if m.StoreRefreshTokenFn != nil {
		return m.StoreRefreshTokenFn(ctx, id, userID, tokenHash, expiresAt)
	}
	return nil
}

func (m *MockAuthRepo) FindRefreshToken(ctx context.Context, tokenHash string) (string, error) {
	if m.FindRefreshTokenFn != nil {
		return m.FindRefreshTokenFn(ctx, tokenHash)
	}
	return "", nil
}

func (m *MockAuthRepo) GetAccountStatus(ctx context.Context, userID string) (string, error) {
	if m.GetAccountStatusFn != nil {
		return m.GetAccountStatusFn(ctx, userID)
	}
	return "", nil
}

func (m *MockAuthRepo) RotateRefreshToken(ctx context.Context, oldTokenID, newTokenID, userID, newTokenHash string, expiresAt time.Time) error {
	if m.RotateRefreshTokenFn != nil {
		return m.RotateRefreshTokenFn(ctx, oldTokenID, newTokenID, userID, newTokenHash, expiresAt)
	}
	return nil
}

func (m *MockAuthRepo) DeleteRefreshTokensByUser(ctx context.Context, userID string) error {
	if m.DeleteRefreshTokensByUserFn != nil {
		return m.DeleteRefreshTokensByUserFn(ctx, userID)
	}
	return nil
}

func (m *MockAuthRepo) GetUserProfile(ctx context.Context, userID string) (*repository.FullUserProfile, error) {
	if m.GetUserProfileFn != nil {
		return m.GetUserProfileFn(ctx, userID)
	}
	return nil, nil
}

func (m *MockAuthRepo) UpdateProfile(ctx context.Context, userID string, fields repository.ProfileUpdateFields) error {
	if m.UpdateProfileFn != nil {
		return m.UpdateProfileFn(ctx, userID, fields)
	}
	return nil
}

// Ensure compile-time interface satisfaction.
var _ repository.AuthRepo = (*MockAuthRepo)(nil)
