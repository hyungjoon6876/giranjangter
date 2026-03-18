package mock

import (
	"context"

	"github.com/jym/lincle/internal/repository"
)

// MockReservationRepo is a function-field based mock for repository.ReservationRepo.
type MockReservationRepo struct {
	GetChatRoomForReservationFn func(ctx context.Context, chatRoomID, userID string) (*repository.ChatRoomReservationInfo, error)
	CountActiveReservationsFn   func(ctx context.Context, listingID string) (int, error)
	CreateReservationFn         func(ctx context.Context, params *repository.CreateReservationParams) error
	GetReservationForConfirmFn  func(ctx context.Context, reservationID string) (*repository.ReservationConfirmInfo, error)
	ConfirmReservationFn        func(ctx context.Context, params *repository.ConfirmReservationParams) error
	GetReservationForCancelFn   func(ctx context.Context, reservationID string) (*repository.ReservationCancelInfo, error)
	CancelReservationFn         func(ctx context.Context, params *repository.CancelReservationParams) error
	GetConfirmedReservationFn   func(ctx context.Context, reservationID, listingID string) (*repository.ReservationParticipants, error)
	CreateTradeCompletionFn     func(ctx context.Context, params *repository.CreateTradeCompletionParams) error
	GetPendingCompletionFn      func(ctx context.Context, completionID string) (*repository.PendingCompletionInfo, error)
	ConfirmCompletionFn         func(ctx context.Context, params *repository.ConfirmCompletionParams) error
	ListMyTradesFn              func(ctx context.Context, userID string) ([]repository.MyTradeItem, error)
	GetCompletionForReviewFn    func(ctx context.Context, completionID string) (*repository.CompletionReviewInfo, error)
	CreateReviewFn              func(ctx context.Context, params *repository.CreateReviewParams) error
	ListUserReviewsFn           func(ctx context.Context, targetUserID string) ([]repository.UserReviewItem, error)
	CreateReportFn              func(ctx context.Context, params *repository.CreateReportParams) error
	ListMyReportsFn             func(ctx context.Context, userID string) ([]repository.MyReportItem, error)
	ListNotificationsFn         func(ctx context.Context, userID string) ([]repository.NotificationItem, error)
	MarkNotificationsReadFn     func(ctx context.Context, userID string, notificationIDs []string) error
}

func (m *MockReservationRepo) GetChatRoomForReservation(ctx context.Context, chatRoomID, userID string) (*repository.ChatRoomReservationInfo, error) {
	if m.GetChatRoomForReservationFn != nil {
		return m.GetChatRoomForReservationFn(ctx, chatRoomID, userID)
	}
	return nil, nil
}

func (m *MockReservationRepo) CountActiveReservations(ctx context.Context, listingID string) (int, error) {
	if m.CountActiveReservationsFn != nil {
		return m.CountActiveReservationsFn(ctx, listingID)
	}
	return 0, nil
}

func (m *MockReservationRepo) CreateReservation(ctx context.Context, params *repository.CreateReservationParams) error {
	if m.CreateReservationFn != nil {
		return m.CreateReservationFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) GetReservationForConfirm(ctx context.Context, reservationID string) (*repository.ReservationConfirmInfo, error) {
	if m.GetReservationForConfirmFn != nil {
		return m.GetReservationForConfirmFn(ctx, reservationID)
	}
	return nil, nil
}

func (m *MockReservationRepo) ConfirmReservation(ctx context.Context, params *repository.ConfirmReservationParams) error {
	if m.ConfirmReservationFn != nil {
		return m.ConfirmReservationFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) GetReservationForCancel(ctx context.Context, reservationID string) (*repository.ReservationCancelInfo, error) {
	if m.GetReservationForCancelFn != nil {
		return m.GetReservationForCancelFn(ctx, reservationID)
	}
	return nil, nil
}

func (m *MockReservationRepo) CancelReservation(ctx context.Context, params *repository.CancelReservationParams) error {
	if m.CancelReservationFn != nil {
		return m.CancelReservationFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) GetConfirmedReservation(ctx context.Context, reservationID, listingID string) (*repository.ReservationParticipants, error) {
	if m.GetConfirmedReservationFn != nil {
		return m.GetConfirmedReservationFn(ctx, reservationID, listingID)
	}
	return nil, nil
}

func (m *MockReservationRepo) CreateTradeCompletion(ctx context.Context, params *repository.CreateTradeCompletionParams) error {
	if m.CreateTradeCompletionFn != nil {
		return m.CreateTradeCompletionFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) GetPendingCompletion(ctx context.Context, completionID string) (*repository.PendingCompletionInfo, error) {
	if m.GetPendingCompletionFn != nil {
		return m.GetPendingCompletionFn(ctx, completionID)
	}
	return nil, nil
}

func (m *MockReservationRepo) ConfirmCompletion(ctx context.Context, params *repository.ConfirmCompletionParams) error {
	if m.ConfirmCompletionFn != nil {
		return m.ConfirmCompletionFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) ListMyTrades(ctx context.Context, userID string) ([]repository.MyTradeItem, error) {
	if m.ListMyTradesFn != nil {
		return m.ListMyTradesFn(ctx, userID)
	}
	return nil, nil
}

func (m *MockReservationRepo) GetCompletionForReview(ctx context.Context, completionID string) (*repository.CompletionReviewInfo, error) {
	if m.GetCompletionForReviewFn != nil {
		return m.GetCompletionForReviewFn(ctx, completionID)
	}
	return nil, nil
}

func (m *MockReservationRepo) CreateReview(ctx context.Context, params *repository.CreateReviewParams) error {
	if m.CreateReviewFn != nil {
		return m.CreateReviewFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) ListUserReviews(ctx context.Context, targetUserID string) ([]repository.UserReviewItem, error) {
	if m.ListUserReviewsFn != nil {
		return m.ListUserReviewsFn(ctx, targetUserID)
	}
	return nil, nil
}

func (m *MockReservationRepo) CreateReport(ctx context.Context, params *repository.CreateReportParams) error {
	if m.CreateReportFn != nil {
		return m.CreateReportFn(ctx, params)
	}
	return nil
}

func (m *MockReservationRepo) ListMyReports(ctx context.Context, userID string) ([]repository.MyReportItem, error) {
	if m.ListMyReportsFn != nil {
		return m.ListMyReportsFn(ctx, userID)
	}
	return nil, nil
}

func (m *MockReservationRepo) ListNotifications(ctx context.Context, userID string) ([]repository.NotificationItem, error) {
	if m.ListNotificationsFn != nil {
		return m.ListNotificationsFn(ctx, userID)
	}
	return nil, nil
}

func (m *MockReservationRepo) MarkNotificationsRead(ctx context.Context, userID string, notificationIDs []string) error {
	if m.MarkNotificationsReadFn != nil {
		return m.MarkNotificationsReadFn(ctx, userID, notificationIDs)
	}
	return nil
}

// Ensure compile-time interface satisfaction.
var _ repository.ReservationRepo = (*MockReservationRepo)(nil)
