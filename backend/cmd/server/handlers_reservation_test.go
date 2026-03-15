package main

import (
	"testing"
)

// Reservation handler tests require a database connection.
// The handlers use *sql.DB directly (not interfaces), so they cannot be
// unit-tested with mocks without refactoring.
//
// These stubs document the critical test scenarios for integration testing.
// Reservation flows involve multi-table transactions (reservations, chat_rooms,
// chat_messages, listings, status_history), so integration tests are essential.

// ── CreateReservation ──

func TestCreateReservation_Success_Returns201(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with chat_room, listing, and users")

	// Test plan:
	// 1. Seed a listing (id="listing-1", status="available") owned by "seller-1"
	// 2. Seed a chat_room (id="chat-1") between "seller-1" and "buyer-1" for "listing-1"
	// 3. Authenticate as "buyer-1"
	// 4. POST /api/v1/chats/chat-1/reservations with:
	//    {
	//      "scheduledAt": "2026-03-20T14:00:00Z",
	//      "meetingType": "in_game",
	//      "serverId": "server-1",
	//      "meetingPointText": "기란마을 분수대",
	//      "noteToCounterparty": "14시에 만나요"
	//    }
	// 5. Expect 201 with:
	//    - reservationId (non-empty UUID)
	//    - status = "proposed"
	// 6. Verify in DB:
	//    - reservations row: proposer=buyer-1, counterpart=seller-1, status="proposed"
	//    - chat_messages row: message_type="system", metadata contains "reservation_proposed"
	//    - chat_rooms.chat_status updated to "reservation_proposed"
}

func TestCreateReservation_ActiveConflict_Returns409(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with existing active reservation")

	// Test plan:
	// 1. Seed listing "listing-1" (status="available")
	// 2. Seed chat_room "chat-1" for listing-1 between seller-1 and buyer-1
	// 3. Seed an existing reservation for listing-1 with status="proposed" (or "confirmed")
	// 4. Authenticate as "buyer-1"
	// 5. POST /api/v1/chats/chat-1/reservations with valid reservation data
	// 6. Expect 409 with:
	//    {"error": {"code": "CONFLICT", "message": "이미 활성 예약이 존재합니다."}}
	// 7. Verify no new reservation row created in DB
	//
	// This enforces the business rule: only one active (proposed/confirmed)
	// reservation per listing at a time.
}

// ── ConfirmReservation ──

func TestConfirmReservation_Counterpart_Returns200(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with proposed reservation")

	// Test plan:
	// 1. Seed listing "listing-1" (status="available")
	// 2. Seed chat_room "chat-1" for listing-1
	// 3. Seed reservation "res-1": proposer=buyer-1, counterpart=seller-1, status="proposed"
	// 4. Authenticate as "seller-1" (the counterpart — only they can confirm)
	// 5. POST /api/v1/reservations/res-1/confirm
	// 6. Expect 200 with:
	//    - reservationId = "res-1"
	//    - status = "confirmed"
	// 7. Verify in DB:
	//    - reservations.status = "confirmed", confirmed_at is set
	//    - listings.status = "reserved", reserved_chat_room_id = "chat-1"
	//    - chat_rooms.chat_status = "reservation_confirmed"
	//    - status_history row: entity_type="listing", from_status="available", to_status="reserved"
}

func TestConfirmReservation_Proposer_Returns403(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with proposed reservation")

	// Test plan:
	// 1. Seed reservation "res-1": proposer=buyer-1, counterpart=seller-1, status="proposed"
	// 2. Authenticate as "buyer-1" (the proposer — they CANNOT confirm their own reservation)
	// 3. POST /api/v1/reservations/res-1/confirm
	// 4. Expect 403 with:
	//    {"error": {"code": "FORBIDDEN", "message": "예약 상대방만 확정할 수 있습니다."}}
	// 5. Verify reservation status unchanged in DB (still "proposed")
	//
	// This enforces the business rule: only the counterpart can confirm.
	// The proposer made the offer; the other party must accept it.
}

// ── CancelReservation ──

func TestCancelReservation_Proposer_Returns200(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with proposed reservation")

	// Test plan:
	// 1. Seed listing "listing-1" (status="reserved")
	// 2. Seed chat_room "chat-1" with chat_status="reservation_proposed"
	// 3. Seed reservation "res-1": proposer=buyer-1, counterpart=seller-1, status="proposed"
	// 4. Authenticate as "buyer-1" (the proposer — they CAN cancel)
	// 5. POST /api/v1/reservations/res-1/cancel with {"reasonCode": "changed_mind"}
	// 6. Expect 200 with:
	//    - reservationId = "res-1"
	//    - status = "cancelled"
	// 7. Verify in DB:
	//    - reservations.status = "cancelled", cancelled_at is set, cancellation_reason_code = "changed_mind"
	//    - listings.status reverted to "available", reserved_chat_room_id = NULL
	//    - chat_rooms.chat_status reverted to "open"
}

func TestCancelReservation_NonParticipant_Returns403(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with reservation and unrelated user")

	// Test plan:
	// 1. Seed reservation "res-1": proposer=buyer-1, counterpart=seller-1, status="proposed"
	// 2. Authenticate as "user-other" (neither proposer nor counterpart)
	// 3. POST /api/v1/reservations/res-1/cancel with {"reasonCode": "spam"}
	// 4. Expect 403 with:
	//    {"error": {"code": "FORBIDDEN", "message": "예약 참여자만 취소할 수 있습니다."}}
	// 5. Verify reservation status unchanged in DB
	//
	// This enforces the business rule: only the proposer or counterpart
	// (i.e., the two participants of the reservation) can cancel it.
}
