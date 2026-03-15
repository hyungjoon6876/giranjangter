package main

import (
	"testing"
)

// Auth handler tests require a database connection and config/middleware setup.
// The handlers use *sql.DB directly (not interfaces), so they cannot be
// unit-tested with mocks without refactoring.
//
// These stubs document the critical test scenarios for integration testing.
// The JWT middleware itself is independently tested in
// internal/middleware/auth_test.go.

// ── Login ──

func TestLogin_NewUser_Returns201(t *testing.T) {
	t.Skip("requires integration test setup: needs DB + AuthMiddleware + Config")

	// Test plan:
	// 1. Set up Config with IsDev()=true (no GoogleClientIDs)
	// 2. Create AuthMiddleware with test JWT secret
	// 3. POST /api/v1/auth/login with {"provider":"google","providerToken":"new-user-token-123"}
	// 4. Expect 201 (StatusCreated)
	// 5. Verify response body contains:
	//    - accessToken (non-empty string)
	//    - refreshToken (non-empty string)
	//    - expiresIn = 900
	//    - user.isNewUser = true
	//    - user.role = "user"
	//    - user.accountStatus = "active"
	//    - user.nickname starts with "유저_"
	// 6. Verify users row created in DB with login_provider="google"
	// 7. Verify user_profiles row created in DB
}

func TestLogin_ExistingUser_Returns200(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with seeded user")

	// Test plan:
	// 1. Set up Config with IsDev()=true
	// 2. Seed a user with login_provider="google", login_provider_user_key="existing-key"
	// 3. POST /api/v1/auth/login with {"provider":"google","providerToken":"existing-key"}
	// 4. Expect 200 (StatusOK) — not 201
	// 5. Verify response body contains:
	//    - user.isNewUser = false
	//    - user.userId matches seeded user's ID
	//    - user.nickname matches seeded profile
	// 6. Verify last_login_at updated in DB
}

func TestLogin_DevMode_AcceptsAnyToken(t *testing.T) {
	t.Skip("requires integration test setup: needs DB + dev Config")

	// Test plan:
	// 1. Set up Config with GIN_MODE="debug" (IsDev()=true) and empty GoogleClientIDs
	// 2. POST /api/v1/auth/login with {"provider":"google","providerToken":"literally-anything"}
	// 3. Expect success (201 for new user)
	// 4. Verify providerToken is used directly as login_provider_user_key
	//    (no Google verification call made)
	// 5. Also test with non-google provider:
	//    POST with {"provider":"dev","providerToken":"dev-key"}
	//    Expect success in dev mode (provider falls to default branch)
}

func TestLogin_ProductionMode_InvalidToken_Returns401(t *testing.T) {
	t.Skip("requires integration test setup: needs Config with GoogleClientIDs set")

	// Test plan:
	// 1. Set up Config with GIN_MODE="release" and GoogleClientIDs=["real-client-id"]
	// 2. POST /api/v1/auth/login with {"provider":"google","providerToken":"invalid-garbage-token"}
	// 3. Expect 401 with:
	//    {"error": {"code": "UNAUTHORIZED", "message": "Google 인증에 실패했습니다."}}
	// 4. Verify no user row created in DB
	//
	// Note: This test depends on oauth.VerifyGoogleIDToken rejecting the token.
	// In production mode, the handler does NOT skip Google verification.
}

// ── Refresh ──

func TestRefresh_ValidToken_Returns200(t *testing.T) {
	t.Skip("requires integration test setup: needs AuthMiddleware with test JWT secret")

	// Test plan:
	// 1. Create AuthMiddleware with known JWT secret
	// 2. Generate a valid refresh token via auth.GenerateTokens("user-1", "user")
	// 3. POST /api/v1/auth/refresh with {"refreshToken": "<valid-refresh-token>"}
	// 4. Expect 200 with:
	//    - accessToken (non-empty, different from original)
	//    - refreshToken (non-empty, different from original)
	//    - expiresIn = 900
	// 5. Verify new access token is parseable and contains correct userID/role claims
}

func TestRefresh_InvalidToken_Returns401(t *testing.T) {
	t.Skip("requires integration test setup: needs AuthMiddleware")

	// Test plan:
	// 1. Create AuthMiddleware with known JWT secret
	// 2. POST /api/v1/auth/refresh with {"refreshToken": "this.is.not.a.valid.jwt"}
	// 3. Expect 401 with:
	//    {"error": {"code": "UNAUTHORIZED", "message": "유효하지 않은 리프레시 토큰"}}
	// 4. Also test with expired refresh token:
	//    - Generate token, then manipulate exp claim to be in the past
	//    - Expect same 401 response
}

// ── GetMe ──

func TestGetMe_Authenticated_Returns200(t *testing.T) {
	t.Skip("requires integration test setup: needs DB with user + profile + auth middleware wired")

	// Test plan:
	// 1. Seed user (id="user-1", role="user", account_status="active")
	// 2. Seed user_profile (nickname="테스트유저", alignment_score=50, alignment_grade="neutral", etc.)
	// 3. Generate valid access token for "user-1"
	// 4. GET /api/v1/me with Authorization: Bearer <token>
	// 5. Expect 200 with complete profile:
	//    - userId = "user-1"
	//    - role = "user"
	//    - accountStatus = "active"
	//    - nickname = "테스트유저"
	//    - completedTradeCount, positiveReviewCount (integers)
	//    - responseBadge, trustBadge (strings)
	//    - alignmentScore = 50
	//    - alignmentGrade = "neutral"
}

func TestGetMe_Unauthenticated_Returns401(t *testing.T) {
	t.Skip("requires integration test setup: validates auth middleware is wired to /me route")

	// Test plan:
	// 1. GET /api/v1/me without Authorization header
	// 2. Expect 401 with {"error": {"code": "UNAUTHORIZED", ...}}
	//
	// Note: The auth middleware itself is unit-tested in middleware/auth_test.go.
	// This test verifies the middleware is correctly applied to the /me route.
}
