package mock

import (
	"context"

	"github.com/jym/lincle/internal/repository"
)

// MockChatRepo is a function-field based mock for repository.ChatRepo.
type MockChatRepo struct {
	GetListingAuthorFn        func(ctx context.Context, listingID string) (string, error)
	FindExistingChatRoomFn    func(ctx context.Context, listingID, sellerID, buyerID string) (string, error)
	CreateChatRoomFn          func(ctx context.Context, params *repository.CreateChatRoomParams) error
	ListChatRoomsFn           func(ctx context.Context, userID string) ([]repository.ChatRoomListItem, error)
	IsChatParticipantFn       func(ctx context.Context, chatRoomID, userID string) (bool, error)
	ListMessagesFn            func(ctx context.Context, chatRoomID string, limit int, cursor string) ([]repository.ChatMessageItem, error)
	GetChatRoomParticipantsFn func(ctx context.Context, chatRoomID, userID string) (*repository.ChatParticipants, error)
	CheckDuplicateMessageFn   func(ctx context.Context, clientMessageID string) (bool, error)
	InsertMessageFn           func(ctx context.Context, params *repository.InsertMessageParams) error
	UpsertReadCursorFn        func(ctx context.Context, chatRoomID, userID, lastReadMessageID string) error
	BlockUserFn               func(ctx context.Context, id, blockerID, blockedID string) error
	UnblockUserFn             func(ctx context.Context, blockerID, blockedID string) error
}

func (m *MockChatRepo) GetListingAuthor(ctx context.Context, listingID string) (string, error) {
	if m.GetListingAuthorFn != nil {
		return m.GetListingAuthorFn(ctx, listingID)
	}
	return "", nil
}

func (m *MockChatRepo) FindExistingChatRoom(ctx context.Context, listingID, sellerID, buyerID string) (string, error) {
	if m.FindExistingChatRoomFn != nil {
		return m.FindExistingChatRoomFn(ctx, listingID, sellerID, buyerID)
	}
	return "", nil
}

func (m *MockChatRepo) CreateChatRoom(ctx context.Context, params *repository.CreateChatRoomParams) error {
	if m.CreateChatRoomFn != nil {
		return m.CreateChatRoomFn(ctx, params)
	}
	return nil
}

func (m *MockChatRepo) ListChatRooms(ctx context.Context, userID string) ([]repository.ChatRoomListItem, error) {
	if m.ListChatRoomsFn != nil {
		return m.ListChatRoomsFn(ctx, userID)
	}
	return nil, nil
}

func (m *MockChatRepo) IsChatParticipant(ctx context.Context, chatRoomID, userID string) (bool, error) {
	if m.IsChatParticipantFn != nil {
		return m.IsChatParticipantFn(ctx, chatRoomID, userID)
	}
	return false, nil
}

func (m *MockChatRepo) ListMessages(ctx context.Context, chatRoomID string, limit int, cursor string) ([]repository.ChatMessageItem, error) {
	if m.ListMessagesFn != nil {
		return m.ListMessagesFn(ctx, chatRoomID, limit, cursor)
	}
	return nil, nil
}

func (m *MockChatRepo) GetChatRoomParticipants(ctx context.Context, chatRoomID, userID string) (*repository.ChatParticipants, error) {
	if m.GetChatRoomParticipantsFn != nil {
		return m.GetChatRoomParticipantsFn(ctx, chatRoomID, userID)
	}
	return nil, nil
}

func (m *MockChatRepo) CheckDuplicateMessage(ctx context.Context, clientMessageID string) (bool, error) {
	if m.CheckDuplicateMessageFn != nil {
		return m.CheckDuplicateMessageFn(ctx, clientMessageID)
	}
	return false, nil
}

func (m *MockChatRepo) InsertMessage(ctx context.Context, params *repository.InsertMessageParams) error {
	if m.InsertMessageFn != nil {
		return m.InsertMessageFn(ctx, params)
	}
	return nil
}

func (m *MockChatRepo) UpsertReadCursor(ctx context.Context, chatRoomID, userID, lastReadMessageID string) error {
	if m.UpsertReadCursorFn != nil {
		return m.UpsertReadCursorFn(ctx, chatRoomID, userID, lastReadMessageID)
	}
	return nil
}

func (m *MockChatRepo) BlockUser(ctx context.Context, id, blockerID, blockedID string) error {
	if m.BlockUserFn != nil {
		return m.BlockUserFn(ctx, id, blockerID, blockedID)
	}
	return nil
}

func (m *MockChatRepo) UnblockUser(ctx context.Context, blockerID, blockedID string) error {
	if m.UnblockUserFn != nil {
		return m.UnblockUserFn(ctx, blockerID, blockedID)
	}
	return nil
}

// Ensure compile-time interface satisfaction.
var _ repository.ChatRepo = (*MockChatRepo)(nil)
