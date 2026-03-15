package main

import (
	"testing"
)

// Chat handler tests require a database connection.
// The handlers use *sql.DB directly (not interfaces), so they cannot be
// unit-tested with mocks without refactoring.
//
// These stubs document the critical test scenarios for integration testing.

func TestCreateChat_Returns201_NewChatRoom(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listings and user_profiles tables")

	// Test plan:
	// 1. Seed a listing owned by "seller-1"
	// 2. Authenticate as "buyer-1"
	// 3. POST /api/v1/listings/:id/chats
	// 4. Expect 201 with chatRoomId, sellerUserId, buyerUserId, chatStatus="open"
	// 5. Verify chat_rooms row created in DB
	// 6. Verify listings.chat_count incremented
}

func TestCreateChat_Returns409_ExistingChatRoom(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with existing chat_room for listing+user pair")

	// Test plan:
	// 1. Seed a listing owned by "seller-1"
	// 2. Seed a chat_room between "seller-1" and "buyer-1" for this listing
	// 3. Authenticate as "buyer-1"
	// 4. POST /api/v1/listings/:id/chats (same listing)
	// 5. Expect 409 with {"chatRoomId": "<existing-id>", "message": "이미 채팅방이 존재합니다."}
	// 6. Verify no new chat_room row created
}

func TestCreateChat_Returns401_Unauthorized(t *testing.T) {
	t.Skip("requires integration test setup: validates auth middleware integration")

	// Test plan:
	// 1. POST /api/v1/listings/:id/chats without Authorization header
	// 2. Expect 401 with {"error": {"code": "UNAUTHORIZED", ...}}
	//
	// Note: The auth middleware itself is unit-tested in middleware/auth_test.go.
	// This test verifies the middleware is correctly wired to the route.
}

func TestCreateChat_Returns400_OwnListing(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing")

	// Test plan:
	// 1. Seed a listing owned by "seller-1"
	// 2. Authenticate as "seller-1" (the listing owner)
	// 3. POST /api/v1/listings/:id/chats
	// 4. Expect 400 with {"error": {"code": "VALIDATION_ERROR", "message": "본인 매물에는 채팅을 시작할 수 없습니다."}}
}

func TestCreateChat_Returns404_DeletedListing(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with soft-deleted listing")

	// Test plan:
	// 1. Seed a listing with deleted_at set
	// 2. Authenticate as "buyer-1"
	// 3. POST /api/v1/listings/:id/chats
	// 4. Expect 404 with {"error": {"code": "NOT_FOUND", ...}}
}

func TestListMessages_Returns403_NonParticipant(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with chat_room")

	// Test plan:
	// 1. Seed a chat_room between "seller-1" and "buyer-1"
	// 2. Authenticate as "user-other" (not a participant)
	// 3. GET /api/v1/chats/:chatId/messages
	// 4. Expect 403 with {"error": {"code": "FORBIDDEN", ...}}
}

func TestSendMessage_DedupByClientMessageID(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with chat_room and messages")

	// Test plan:
	// 1. Seed chat_room and an existing message with client_message_id "dedup-1"
	// 2. Authenticate as participant
	// 3. POST /api/v1/chats/:chatId/messages with clientMessageId "dedup-1"
	// 4. Expect 200 (not 201) with dedup message
	// 5. Verify no duplicate row created
}
