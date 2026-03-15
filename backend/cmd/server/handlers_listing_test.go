package main

import (
	"testing"
)

// Listing handler tests require a database connection.
// The handlers use *sql.DB directly (not interfaces), so they cannot be
// unit-tested with mocks without refactoring.
//
// These stubs document the critical test scenarios for integration testing.
// The state transition guard logic is independently tested in
// internal/guard/listing_guard_test.go.

func TestChangeListingStatus_InvalidTransition_Returns422(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing in specific status")

	// Test plan:
	// 1. Seed a listing with status "available", owned by "user-1"
	// 2. Authenticate as "user-1"
	// 3. POST /api/v1/listings/:id/status with {"action": "complete"}
	//    (available -> completed is not allowed; must go through reserved -> pending_trade first)
	// 4. Expect 422 with {"error": {"code": "INVALID_TRANSITION", ...}}
	// 5. Verify listing status unchanged in DB
}

func TestChangeListingStatus_ValidTransition_Returns200(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing")

	// Test plan:
	// 1. Seed a listing with status "available", owned by "user-1"
	// 2. Authenticate as "user-1"
	// 3. POST /api/v1/listings/:id/status with {"action": "reserve"}
	//    (available -> reserved is allowed)
	// 4. Expect 200 with {"listingId": ..., "status": "reserved", "updatedAt": ...}
	// 5. Verify listing.status = "reserved" in DB
	// 6. Verify status_history row created
}

func TestChangeListingStatus_NonOwner_Returns403(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing owned by different user")

	// Test plan:
	// 1. Seed a listing owned by "user-1"
	// 2. Authenticate as "user-2" (not the owner)
	// 3. POST /api/v1/listings/:id/status with {"action": "reserve"}
	// 4. Expect 403 with {"error": {"code": "FORBIDDEN", "message": "본인 매물만 상태를 변경할 수 있습니다."}}
}

func TestUpdateListing_CompletedStatus_Returns403(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with completed listing")

	// Test plan:
	// 1. Seed a listing with status "completed", owned by "user-1"
	// 2. Authenticate as "user-1"
	// 3. PATCH /api/v1/listings/:id with {"title": "new title"}
	// 4. Expect 403 with {"error": {"code": "FORBIDDEN", "message": "종결된 매물은 수정할 수 없습니다."}}
	// 5. Verify listing.title unchanged in DB
}

func TestUpdateListing_CancelledStatus_Returns403(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with cancelled listing")

	// Test plan:
	// 1. Seed a listing with status "cancelled", owned by "user-1"
	// 2. Authenticate as "user-1"
	// 3. PATCH /api/v1/listings/:id with {"title": "new title"}
	// 4. Expect 403 with {"error": {"code": "FORBIDDEN", "message": "종결된 매물은 수정할 수 없습니다."}}
}

func TestUpdateListing_NonOwner_Returns403(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing owned by different user")

	// Test plan:
	// 1. Seed a listing owned by "user-1"
	// 2. Authenticate as "user-2"
	// 3. PATCH /api/v1/listings/:id with {"title": "new title"}
	// 4. Expect 403 with {"error": {"code": "FORBIDDEN", "message": "본인 매물만 수정할 수 있습니다."}}
}

func TestUpdateListing_NoFields_Returns400(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing")

	// Test plan:
	// 1. Seed a listing with status "available", owned by "user-1"
	// 2. Authenticate as "user-1"
	// 3. PATCH /api/v1/listings/:id with {} (empty body)
	// 4. Expect 400 with {"error": {"code": "VALIDATION_ERROR", "message": "수정할 필드가 없습니다."}}
}

func TestChangeListingStatus_InvalidAction_Returns400(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with listing")

	// Test plan:
	// 1. Seed a listing owned by "user-1"
	// 2. Authenticate as "user-1"
	// 3. POST /api/v1/listings/:id/status with {"action": "explode"}
	// 4. Expect 400 with {"error": {"code": "VALIDATION_ERROR", "message": "잘못된 액션입니다."}}
}

func TestChangeListingStatus_DeletedListing_Returns404(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with soft-deleted listing")

	// Test plan:
	// 1. Seed a listing with deleted_at set
	// 2. Authenticate as the listing owner
	// 3. POST /api/v1/listings/:id/status with {"action": "reserve"}
	// 4. Expect 404 with {"error": {"code": "NOT_FOUND", ...}}
}

func TestCreateListing_MissingFields_Returns400(t *testing.T) {
	t.Skip("requires integration test setup: needs DB")

	// Test plan:
	// 1. Authenticate as "user-1"
	// 2. POST /api/v1/listings with incomplete body (missing required fields)
	// 3. Expect 400 with {"error": {"code": "VALIDATION_ERROR", ...}}
}

func TestCreateListing_PriceRequired_ForNonOffer(t *testing.T) {
	t.Skip("requires integration test setup: needs DB")

	// Test plan:
	// 1. Authenticate as "user-1"
	// 2. POST /api/v1/listings with priceType "fixed" but no priceAmount
	// 3. Expect 400 with {"error": {"code": "VALIDATION_ERROR", "message": "가격을 입력해주세요."}}
}
