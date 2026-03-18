# Test Infrastructure Improvement Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Backend 32 skip + Flutter 10 skip → 전체 0 skip. Repository 인터페이스 추출 + Flutter 추상화.

**Architecture:** Backend 핸들러의 `*sql.DB` 직접 사용을 도메인별 Repository 인터페이스로 교체. Flutter는 ApiClient/AuthService 인터페이스 추출 후 Riverpod override로 Mock 주입.

**Tech Stack:** Go (database/sql, gin, httptest), Dart (flutter_riverpod, flutter_test)

---

## File Map

### Backend — New Files
| File | Responsibility |
|------|---------------|
| `backend/internal/repository/interfaces.go` | AuthRepo, ListingRepo, ChatRepo, ReservationRepo 인터페이스 |
| `backend/internal/repository/postgres_auth.go` | PostgresAuthRepo 구현체 |
| `backend/internal/repository/postgres_listing.go` | PostgresListingRepo 구현체 |
| `backend/internal/repository/postgres_chat.go` | PostgresChatRepo 구현체 |
| `backend/internal/repository/postgres_reservation.go` | PostgresReservationRepo 구현체 |
| `backend/internal/repository/mock/auth.go` | MockAuthRepo |
| `backend/internal/repository/mock/listing.go` | MockListingRepo |
| `backend/internal/repository/mock/chat.go` | MockChatRepo |
| `backend/internal/repository/mock/reservation.go` | MockReservationRepo |

### Backend — Modified Files
| File | Change |
|------|--------|
| `backend/cmd/server/handlers_auth.go` | `*sql.DB` → `repository.AuthRepo` |
| `backend/cmd/server/handlers_listing.go` | `*sql.DB` → `repository.ListingRepo` |
| `backend/cmd/server/handlers_chat.go` | `*sql.DB` → `repository.ChatRepo` |
| `backend/cmd/server/handlers_reservation.go` | `*sql.DB` → `repository.ReservationRepo` |
| `backend/cmd/server/main.go` | Repo 생성 + 주입 |
| `backend/cmd/server/handlers_auth_test.go` | stub → 실제 테스트 |
| `backend/cmd/server/handlers_chat_test.go` | stub → 실제 테스트 |
| `backend/cmd/server/handlers_listing_test.go` | stub → 실제 테스트 |
| `backend/cmd/server/handlers_reservation_test.go` | stub → 실제 테스트 |

### Flutter — New Files
| File | Responsibility |
|------|---------------|
| `frontend/lib/shared/api/api_client_interface.dart` | IApiClient abstract class |
| `frontend/lib/shared/api/auth_service.dart` | AuthService abstract class + AuthResult |
| `frontend/lib/shared/api/google_auth_service.dart` | GoogleAuthService implements AuthService |
| `frontend/test/helpers/mock_api_client.dart` | MockApiClient implements IApiClient |
| `frontend/test/helpers/mock_auth_service.dart` | MockAuthService implements AuthService |

### Flutter — Modified Files
| File | Change |
|------|--------|
| `frontend/lib/shared/api/api_client.dart` | `implements IApiClient` 추가 |
| `frontend/lib/shared/providers/app_providers.dart` | Provider 타입을 인터페이스로, authServiceProvider 추가 |
| `frontend/lib/features/auth/login_screen.dart` | GoogleSignIn 직접 사용 → AuthService |
| `frontend/test/widget_test.dart` | skip 제거, mock 주입 |
| `frontend/test/widget/login_screen_test.dart` | skip 제거, mock 주입 |
| `frontend/test/widget/listing_list_screen_test.dart` | skip 제거, mock 주입 |

---

## Task 1: Backend — Repository 인터페이스 정의

**Files:**
- Create: `backend/internal/repository/interfaces.go`

이 태스크는 다른 모든 Backend 태스크의 기반. 핸들러 코드에서 사용하는 DB 호출을 분석하여 인터페이스를 정의한다.

- [ ] **Step 1: interfaces.go 작성**

```go
package repository

import (
	"context"
	"time"

	"github.com/jym/lincle/internal/domain"
)

// AuthRepo handles user authentication and token management.
type AuthRepo interface {
	// FindUserByProvider returns a user by login provider and provider key.
	// Returns (nil, nil) if not found.
	FindUserByProvider(ctx context.Context, provider, providerKey string) (*AuthUser, error)
	// CreateUser creates a new user + profile. Returns the new user ID.
	CreateUser(ctx context.Context, provider, providerKey, nickname string) (string, error)
	// UpdateLastLogin updates the user's last_login_at timestamp.
	UpdateLastLogin(ctx context.Context, userID string, at time.Time) error
	// SaveRefreshToken stores a hashed refresh token.
	SaveRefreshToken(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error
	// ValidateRefreshToken checks if a refresh token hash is valid. Returns user ID.
	ValidateRefreshToken(ctx context.Context, tokenHash string) (string, error)
	// DeleteRefreshToken deletes a refresh token by its DB row ID.
	DeleteRefreshToken(ctx context.Context, tokenID string) error
	// DeleteAllRefreshTokens deletes all refresh tokens for a user (logout).
	DeleteAllRefreshTokens(ctx context.Context, userID string) error
	// GetUserAccountStatus returns the account_status for a user.
	GetUserAccountStatus(ctx context.Context, userID string) (string, error)
	// GetUserProfile returns the full user profile.
	GetUserProfile(ctx context.Context, userID string) (*UserProfile, error)
	// UpdateProfile updates the user's profile fields.
	UpdateProfile(ctx context.Context, userID string, fields map[string]interface{}) error
}

// AuthUser is a minimal user struct for auth operations.
type AuthUser struct {
	ID       string
	Role     string
	Nickname string
}

// UserProfile is the full profile for /me endpoint.
type UserProfile struct {
	ID                   string
	Role                 string
	AccountStatus        string
	Nickname             string
	AvatarURL            *string
	Introduction         *string
	PrimaryServerID      *string
	CompletedTradeCount  int
	PositiveReviewCount  int
	ResponseBadge        *string
	TrustBadge           *string
	AlignmentScore       int
	AlignmentGrade       string
}

// ListingRepo handles listing CRUD and queries.
type ListingRepo interface {
	CreateListing(ctx context.Context, input *CreateListingInput) (string, error)
	GetListing(ctx context.Context, id string) (*ListingDetail, error)
	IncrementViewCount(ctx context.Context, id string) error
	CheckFavorited(ctx context.Context, userID, listingID string) (bool, error)
	ListListings(ctx context.Context, filters ListingFilters) ([]ListingSummary, error)
	UpdateListing(ctx context.Context, id string, fields map[string]interface{}) error
	GetListingOwnerAndStatus(ctx context.Context, id string) (ownerID, status string, err error)
	ChangeListingStatus(ctx context.Context, id, fromStatus, toStatus, changedBy, reasonCode string) error
	FavoriteListing(ctx context.Context, userID, listingID string) error
	UnfavoriteListing(ctx context.Context, userID, listingID string) error
	ListMyListings(ctx context.Context, userID string, status *string) ([]ListingSummary, error)
	CheckImageOwnership(ctx context.Context, imageID, userID string) (bool, error)
}

type CreateListingInput struct {
	ListingType          string
	AuthorUserID         string
	ServerID             string
	CategoryID           string
	ItemName             string
	Title                string
	Description          string
	PriceType            string
	PriceAmount          *int64
	Quantity             int
	EnhancementLevel     *int
	OptionsText          *string
	TradeMethod          string
	PreferredMeetingArea *string
	AvailableTimeText    *string
	ImageIDs             []string
}

type ListingFilters struct {
	Status      *string
	ServerID    *string
	CategoryID  *string
	ListingType *string
	Query       *string
	Sort        string
	Cursor      *string
	UserID      *string // for favorite check
}

type ListingSummary struct {
	ID            string
	ListingType   string
	Title         string
	ItemName      string
	PriceType     string
	PriceAmount   *int64
	Status        string
	ServerName    string
	SellerNickname string
	ViewCount     int
	FavoriteCount int
	ChatCount     int
	IsFavorited   bool
	CreatedAt     time.Time
}

type ListingDetail struct {
	ListingSummary
	Description          string
	AuthorUserID         string
	ServerID             string
	CategoryID           string
	CategoryName         string
	Quantity             int
	EnhancementLevel     *int
	OptionsText          *string
	TradeMethod          string
	PreferredMeetingArea *string
	AvailableTimeText    *string
	IconID               *string
	UpdatedAt            time.Time
}

// ChatRepo handles chat rooms and messages.
type ChatRepo interface {
	GetListingOwner(ctx context.Context, listingID string) (string, error)
	FindExistingChat(ctx context.Context, listingID, userA, userB string) (string, error)
	CreateChat(ctx context.Context, listingID, sellerID, buyerID string) (string, error)
	IncrementChatCount(ctx context.Context, listingID string) error
	ListChats(ctx context.Context, userID string) ([]ChatSummary, error)
	CheckChatParticipant(ctx context.Context, chatID, userID string) (bool, error)
	ListMessages(ctx context.Context, chatID string, cursor *string) ([]domain.ChatMessage, error)
	GetChatParticipants(ctx context.Context, chatID, userID string) (sellerID, buyerID string, err error)
	CheckMessageDedup(ctx context.Context, clientMsgID string) (bool, error)
	SendMessage(ctx context.Context, chatID, senderID, msgType, body, clientMsgID string) (*domain.ChatMessage, error)
	MarkRead(ctx context.Context, chatID, userID, lastReadMsgID string) error
}

type ChatSummary struct {
	ID                 string
	ListingID          string
	ListingTitle       string
	CounterpartID      string
	CounterpartNickname string
	ChatStatus         string
	LastMessageAt      *time.Time
}

// ReservationRepo handles reservation lifecycle.
type ReservationRepo interface {
	GetChatForReservation(ctx context.Context, chatID, userID string) (*ChatReservationInfo, error)
	CheckActiveReservation(ctx context.Context, listingID string) (bool, error)
	CreateReservation(ctx context.Context, input *CreateReservationInput) (string, error)
	GetReservationForConfirm(ctx context.Context, resID string) (*ReservationInfo, error)
	ConfirmReservation(ctx context.Context, resID, listingID, chatID, confirmedBy string) error
	GetReservationForCancel(ctx context.Context, resID string) (*ReservationCancelInfo, error)
	CancelReservation(ctx context.Context, resID, listingID, chatID, reason string) error
}

type ChatReservationInfo struct {
	ListingID string
	SellerID  string
	BuyerID   string
}

type CreateReservationInput struct {
	ListingID        string
	ChatRoomID       string
	ProposerID       string
	CounterpartID    string
	ScheduledAt      time.Time
	MeetingType      string
	ServerID         *string
	MeetingPointText *string
	Note             *string
	ExpiresAt        time.Time
}

type ReservationInfo struct {
	CounterpartUserID string
	ListingID         string
	ChatRoomID        string
}

type ReservationCancelInfo struct {
	ListingID   string
	ChatRoomID  string
	ProposerID  string
	CounterpartID string
}
```

- [ ] **Step 2: 컴파일 확인**

Run: `cd backend && go build ./internal/repository/...`
Expected: SUCCESS (인터페이스만 정의, 구현 없음)

- [ ] **Step 3: Commit**

```bash
git add backend/internal/repository/interfaces.go
git commit -m "feat(backend): repository 인터페이스 정의 — auth, listing, chat, reservation"
```

---

## Task 2: Backend — PostgreSQL Repository 구현체

**Files:**
- Create: `backend/internal/repository/postgres_auth.go`
- Create: `backend/internal/repository/postgres_listing.go`
- Create: `backend/internal/repository/postgres_chat.go`
- Create: `backend/internal/repository/postgres_reservation.go`

각 핸들러 파일에서 SQL 쿼리를 그대로 추출하여 구현체로 이동. 로직 변경 없이 SQL을 1:1 이동하는 기계적 작업.

- [ ] **Step 1: postgres_auth.go 작성**

`handlers_auth.go`에서 SQL 쿼리를 추출:
- `FindUserByProvider` ← Line 79-82의 SELECT
- `CreateUser` ← Line 99-111의 INSERT users + INSERT user_profiles (트랜잭션)
- `UpdateLastLogin` ← Line 132의 UPDATE
- `SaveRefreshToken` ← Line 143-145의 INSERT
- `ValidateRefreshToken` ← Line 190-192의 SELECT
- `DeleteRefreshToken` ← Line 227의 DELETE
- `DeleteAllRefreshTokens` ← Line 376의 DELETE
- `GetUserAccountStatus` ← Line 203의 SELECT
- `GetUserProfile` ← Line 288-294의 SELECT
- `UpdateProfile` ← Line 332-356의 dynamic UPDATE

구조:
```go
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PostgresAuthRepo struct {
	db *sql.DB
}

func NewPostgresAuthRepo(db *sql.DB) *PostgresAuthRepo {
	return &PostgresAuthRepo{db: db}
}

// 각 메서드는 handlers_auth.go에서 해당 SQL을 그대로 복사.
// sql.ErrNoRows일 때 FindUserByProvider는 (nil, nil) 반환.
```

- [ ] **Step 2: postgres_listing.go 작성**

`handlers_listing.go`에서 SQL 추출. 가장 큰 파일 — `ListListings`의 동적 쿼리 빌더 포함.

- [ ] **Step 3: postgres_chat.go 작성**

`handlers_chat.go`에서 SQL 추출. `CreateChat`은 트랜잭션 (INSERT chat_rooms + UPDATE listings.chat_count).

- [ ] **Step 4: postgres_reservation.go 작성**

`handlers_reservation.go`에서 SQL 추출. `CreateReservation`, `ConfirmReservation`, `CancelReservation` 모두 복합 트랜잭션.

- [ ] **Step 5: 컴파일 확인**

Run: `cd backend && go build ./internal/repository/...`
Expected: SUCCESS

- [ ] **Step 6: Commit**

```bash
git add backend/internal/repository/postgres_*.go
git commit -m "feat(backend): PostgreSQL repository 구현체 — 핸들러에서 SQL 추출"
```

---

## Task 3: Backend — Mock Repository 구현

**Files:**
- Create: `backend/internal/repository/mock/auth.go`
- Create: `backend/internal/repository/mock/listing.go`
- Create: `backend/internal/repository/mock/chat.go`
- Create: `backend/internal/repository/mock/reservation.go`

함수 필드 기반 mock 패턴. 외부 라이브러리 없이 Go 표준만 사용.

- [ ] **Step 1: mock/auth.go 작성**

```go
package mock

import (
	"context"
	"time"

	"github.com/jym/lincle/internal/repository"
)

type MockAuthRepo struct {
	FindUserByProviderFn    func(ctx context.Context, provider, providerKey string) (*repository.AuthUser, error)
	CreateUserFn            func(ctx context.Context, provider, providerKey, nickname string) (string, error)
	UpdateLastLoginFn       func(ctx context.Context, userID string, at time.Time) error
	SaveRefreshTokenFn      func(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error
	ValidateRefreshTokenFn  func(ctx context.Context, tokenHash string) (string, error)
	DeleteRefreshTokenFn    func(ctx context.Context, tokenID string) error
	DeleteAllRefreshTokensFn func(ctx context.Context, userID string) error
	GetUserAccountStatusFn  func(ctx context.Context, userID string) (string, error)
	GetUserProfileFn        func(ctx context.Context, userID string) (*repository.UserProfile, error)
	UpdateProfileFn         func(ctx context.Context, userID string, fields map[string]interface{}) error
}

func (m *MockAuthRepo) FindUserByProvider(ctx context.Context, provider, providerKey string) (*repository.AuthUser, error) {
	if m.FindUserByProviderFn != nil {
		return m.FindUserByProviderFn(ctx, provider, providerKey)
	}
	return nil, nil
}

// ... 나머지 메서드도 동일 패턴
```

- [ ] **Step 2: mock/listing.go, mock/chat.go, mock/reservation.go 작성**

동일한 함수 필드 패턴. 각 인터페이스의 모든 메서드에 대해 `XxxFn` 필드 + delegation 메서드.

- [ ] **Step 3: 컴파일 확인**

Run: `cd backend && go build ./internal/repository/mock/...`
Expected: SUCCESS

- [ ] **Step 4: Commit**

```bash
git add backend/internal/repository/mock/
git commit -m "feat(backend): mock repository 구현 — 함수 필드 기반 테스트용"
```

---

## Task 4: Backend — 핸들러 리팩터링 + main.go 배선

**Files:**
- Modify: `backend/cmd/server/handlers_auth.go`
- Modify: `backend/cmd/server/handlers_listing.go`
- Modify: `backend/cmd/server/handlers_chat.go`
- Modify: `backend/cmd/server/handlers_reservation.go`
- Modify: `backend/cmd/server/main.go`

핸들러 함수 시그니처를 `*sql.DB` → Repository 인터페이스로 변경하고, 핸들러 내부에서 직접 SQL 호출하던 부분을 repository 메서드 호출로 교체.

- [ ] **Step 1: handlers_auth.go 리팩터링**

시그니처 변경:
```go
// Before
func handleLogin(db *sql.DB, auth *middleware.AuthMiddleware, cfg *config.Config) gin.HandlerFunc

// After
func handleLogin(repo repository.AuthRepo, auth *middleware.AuthMiddleware, cfg *config.Config) gin.HandlerFunc
```

핸들러 내부 변경 (handleLogin 예시):
```go
// Before
var userID, role, nickname string
err := db.QueryRow("SELECT u.id, u.role, ...").Scan(&userID, &role, &nickname)

// After
user, err := repo.FindUserByProvider(c.Request.Context(), "google", claims.Sub)
if err != nil { ... }
if user == nil {
    // new user
    userID, err := repo.CreateUser(c.Request.Context(), "google", claims.Sub, nickname)
    ...
}
```

모든 auth 핸들러에 동일 적용: handleRefresh, handleGetMe, handleUpdateProfile, handleLogout.

- [ ] **Step 2: handlers_listing.go 리팩터링**

```go
// Before
func handleCreateListing(db *sql.DB) gin.HandlerFunc
func handleListListings(db *sql.DB) gin.HandlerFunc
// ...

// After
func handleCreateListing(repo repository.ListingRepo) gin.HandlerFunc
func handleListListings(repo repository.ListingRepo) gin.HandlerFunc
// ...
```

- [ ] **Step 3: handlers_chat.go 리팩터링**

```go
// Before
func handleCreateChat(db *sql.DB) gin.HandlerFunc

// After
func handleCreateChat(repo repository.ChatRepo) gin.HandlerFunc
```

`handleSendMessage`은 `ChatRepo` + `*event.Broker` 유지.

- [ ] **Step 4: handlers_reservation.go 리팩터링**

```go
// Before
func handleCreateReservation(db *sql.DB) gin.HandlerFunc

// After
func handleCreateReservation(repo repository.ReservationRepo) gin.HandlerFunc
```

- [ ] **Step 5: main.go 배선 변경**

```go
// Before
v1.POST("/auth/login", handleLogin(db, auth, cfg))

// After
authRepo := repository.NewPostgresAuthRepo(db)
listingRepo := repository.NewPostgresListingRepo(db)
chatRepo := repository.NewPostgresChatRepo(db)
reservationRepo := repository.NewPostgresReservationRepo(db)

v1.POST("/auth/login", handleLogin(authRepo, auth, cfg))
write.POST("/listings", handleCreateListing(listingRepo))
// ... 모든 라우트 업데이트
```

주의: 일부 핸들러(handleCompleteTrade, handleCreateReport 등)는 아직 repository 미추출. 이들은 `db *sql.DB`를 그대로 유지. 단계적으로 전환.

- [ ] **Step 6: 컴파일 확인**

Run: `cd backend && go build ./cmd/server/...`
Expected: SUCCESS

- [ ] **Step 7: 기존 테스트 통과 확인**

Run: `cd backend && go test ./...`
Expected: 기존 guard/middleware/domain/event 테스트 전부 PASS. 핸들러 테스트는 여전히 SKIP (다음 태스크에서 구현).

- [ ] **Step 8: Commit**

```bash
git add backend/cmd/server/ backend/internal/repository/
git commit -m "refactor(backend): 핸들러 *sql.DB → repository 인터페이스 전환"
```

---

## Task 5: Backend — 핸들러 테스트 구현

**Files:**
- Modify: `backend/cmd/server/handlers_auth_test.go`
- Modify: `backend/cmd/server/handlers_chat_test.go`
- Modify: `backend/cmd/server/handlers_listing_test.go`
- Modify: `backend/cmd/server/handlers_reservation_test.go`

기존 stub 테스트의 `t.Skip()` 제거하고 mock repository로 실제 테스트 구현.

- [ ] **Step 1: 테스트 헬퍼 함수 작성**

`handlers_auth_test.go` 상단에 공통 헬퍼:

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jym/lincle/internal/config"
	"github.com/jym/lincle/internal/middleware"
	"github.com/jym/lincle/internal/repository"
	"github.com/jym/lincle/internal/repository/mock"
)

const testSecret = "test-jwt-secret-key-for-unit-tests"

func newTestAuth() *middleware.AuthMiddleware {
	return middleware.NewAuthMiddleware(testSecret, 15*time.Minute, 7*24*time.Hour)
}

func newDevConfig() *config.Config {
	return &config.Config{Env: "development", JWTSecret: testSecret}
}

func setupRouter(handlers ...func(*gin.Engine)) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	for _, h := range handlers {
		h(r)
	}
	return r
}
```

- [ ] **Step 2: handlers_auth_test.go 구현 (8건)**

```go
func TestLogin_NewUser_Returns201(t *testing.T) {
	auth := newTestAuth()
	cfg := newDevConfig()
	mockRepo := &mock.MockAuthRepo{
		FindUserByProviderFn: func(ctx context.Context, provider, key string) (*repository.AuthUser, error) {
			return nil, nil // 유저 없음
		},
		CreateUserFn: func(ctx context.Context, provider, key, nickname string) (string, error) {
			return "new-user-id", nil
		},
		UpdateLastLoginFn:  func(ctx context.Context, userID string, at time.Time) error { return nil },
		SaveRefreshTokenFn: func(ctx context.Context, userID, hash string, exp time.Time) error { return nil },
	}

	r := setupRouter(func(e *gin.Engine) {
		e.POST("/api/v1/auth/login", handleLogin(mockRepo, auth, cfg))
	})

	body := `{"provider":"google","token":"dev-test-token"}`
	req := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}
	var resp map[string]interface{}
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["is_new_user"] != true {
		t.Errorf("expected is_new_user=true")
	}
}
```

나머지 7건도 동일 패턴. 각 테스트는 mock의 함수 필드를 시나리오에 맞게 설정.

- [ ] **Step 3: handlers_listing_test.go 구현 (10건)**

핵심 테스트:
- `TestChangeListingStatus_InvalidTransition_Returns422` — mock이 현재 status 반환 → guard가 거부
- `TestCreateListing_MissingFields_Returns400` — 요청 body 불완전
- `TestUpdateListing_NonOwner_Returns403` — mock의 ownerID ≠ 요청자 ID

- [ ] **Step 4: handlers_chat_test.go 구현 (7건)**

핵심 테스트:
- `TestCreateChat_Returns400_OwnListing` — mock의 listingOwner == 요청자
- `TestSendMessage_DedupByClientMessageID` — mock의 CheckMessageDedup이 true 반환

- [ ] **Step 5: handlers_reservation_test.go 구현 (7건)**

핵심 테스트:
- `TestConfirmReservation_Proposer_Returns403` — mock의 counterpartID ≠ 요청자
- `TestCreateReservation_ActiveConflict_Returns409` — mock의 CheckActiveReservation이 true

- [ ] **Step 6: 전체 테스트 실행**

Run: `cd backend && go test -v ./...`
Expected: 68건 전체 PASS, 0 SKIP

- [ ] **Step 7: Commit**

```bash
git add backend/cmd/server/*_test.go
git commit -m "test(backend): 핸들러 테스트 32건 구현 — mock repository 기반"
```

---

## Task 6: Flutter — IApiClient 인터페이스 + Mock

**Files:**
- Create: `frontend/lib/shared/api/api_client_interface.dart`
- Create: `frontend/test/helpers/mock_api_client.dart`
- Modify: `frontend/lib/shared/api/api_client.dart`
- Modify: `frontend/lib/shared/providers/app_providers.dart`

- [ ] **Step 1: api_client_interface.dart 작성**

```dart
abstract class IApiClient {
  bool get isLoggedIn;
  String get staticBaseUrl;

  Future<void> loadTokens();
  Future<void> saveTokens(String access, String refresh);
  Future<void> clearTokens();

  Future<Map<String, dynamic>> login(String provider, String token);
  Future<Map<String, dynamic>> getMe();

  Future<Map<String, dynamic>> getListings({
    String? serverId, String? categoryId, String? q,
    String sort = 'recent', String? cursor,
  });
  Future<Map<String, dynamic>> getListing(String id);
  Future<Map<String, dynamic>> createListing(Map<String, dynamic> data);
  Future<void> favoriteListing(String id);
  Future<void> unfavoriteListing(String id);

  Future<Map<String, dynamic>> createChat(String listingId);
  Future<Map<String, dynamic>> getChats();
  Future<Map<String, dynamic>> getMessages(String chatId, {String? cursor});
  Future<Map<String, dynamic>> sendMessage(String chatId, String text);

  Future<Map<String, dynamic>> getMyListings({String? status});
  Future<Map<String, dynamic>> getMyTrades();
  Future<Map<String, dynamic>> getNotifications();
  Future<void> markNotificationsRead(List<String> ids);
  Future<Map<String, dynamic>> getMyReports();
  Future<Map<String, dynamic>> getUserReviews(String userId);

  Future<Map<String, dynamic>> createReservation(String chatId, Map<String, dynamic> data);
  Future<Map<String, dynamic>> completeTrade(String listingId, Map<String, dynamic> data);
  Future<Map<String, dynamic>> createReview(String completionId, Map<String, dynamic> data);

  Future<Map<String, dynamic>> createReport(Map<String, dynamic> data);
  Future<List<dynamic>> getServers();
  Future<List<dynamic>> getCategories();
  Future<List<dynamic>> searchItems(String query, {String? categoryId});
}
```

- [ ] **Step 2: api_client.dart에 `implements IApiClient` 추가**

```dart
// Before
class ApiClient {

// After
import 'api_client_interface.dart';
class ApiClient implements IApiClient {
```

기존 코드 변경 없음. `implements` 키워드만 추가.

- [ ] **Step 3: app_providers.dart 타입 변경**

```dart
// Before
final apiClientProvider = Provider<ApiClient>((ref) => ApiClient());

// After
import '../api/api_client_interface.dart';
final apiClientProvider = Provider<IApiClient>((ref) => ApiClient());
```

- [ ] **Step 4: mock_api_client.dart 작성**

```dart
import 'package:lincle/shared/api/api_client_interface.dart';

class MockApiClient implements IApiClient {
  @override bool get isLoggedIn => false;
  @override String get staticBaseUrl => 'http://test';

  @override Future<void> loadTokens() async {}
  @override Future<void> saveTokens(String a, String r) async {}
  @override Future<void> clearTokens() async {}

  @override Future<Map<String, dynamic>> login(String p, String t) async => {'access_token': 'test', 'refresh_token': 'test'};
  @override Future<Map<String, dynamic>> getMe() async => {'id': 'test', 'nickname': 'tester'};

  @override Future<Map<String, dynamic>> getListings({String? serverId, String? categoryId, String? q, String sort = 'recent', String? cursor}) async => {'data': [], 'next_cursor': null};
  @override Future<Map<String, dynamic>> getListing(String id) async => {};
  @override Future<Map<String, dynamic>> createListing(Map<String, dynamic> data) async => {};
  @override Future<void> favoriteListing(String id) async {}
  @override Future<void> unfavoriteListing(String id) async {}

  @override Future<Map<String, dynamic>> createChat(String listingId) async => {};
  @override Future<Map<String, dynamic>> getChats() async => {'data': []};
  @override Future<Map<String, dynamic>> getMessages(String chatId, {String? cursor}) async => {'data': []};
  @override Future<Map<String, dynamic>> sendMessage(String chatId, String text) async => {};

  @override Future<Map<String, dynamic>> getMyListings({String? status}) async => {'data': []};
  @override Future<Map<String, dynamic>> getMyTrades() async => {'data': []};
  @override Future<Map<String, dynamic>> getNotifications() async => {'data': []};
  @override Future<void> markNotificationsRead(List<String> ids) async {}
  @override Future<Map<String, dynamic>> getMyReports() async => {'data': []};
  @override Future<Map<String, dynamic>> getUserReviews(String userId) async => {'data': []};

  @override Future<Map<String, dynamic>> createReservation(String chatId, Map<String, dynamic> data) async => {};
  @override Future<Map<String, dynamic>> completeTrade(String listingId, Map<String, dynamic> data) async => {};
  @override Future<Map<String, dynamic>> createReview(String completionId, Map<String, dynamic> data) async => {};

  @override Future<Map<String, dynamic>> createReport(Map<String, dynamic> data) async => {};
  @override Future<List<dynamic>> getServers() async => [{'id': 'test_server', 'name': '테스트서버'}];
  @override Future<List<dynamic>> getCategories() async => [{'id': 'weapon', 'name': '무기'}];
  @override Future<List<dynamic>> searchItems(String query, {String? categoryId}) async => [];
}
```

- [ ] **Step 5: 컴파일 확인**

Run: `cd frontend && flutter analyze --no-fatal-infos`
Expected: No errors

- [ ] **Step 6: 기존 테스트 통과 확인**

Run: `cd frontend && flutter test test/api_client_test.dart test/utils_test.dart`
Expected: 모두 PASS (기존 테스트 깨지지 않음)

- [ ] **Step 7: Commit**

```bash
git add frontend/lib/shared/api/api_client_interface.dart frontend/lib/shared/api/api_client.dart frontend/lib/shared/providers/app_providers.dart frontend/test/helpers/mock_api_client.dart
git commit -m "feat(flutter): IApiClient 인터페이스 추출 + MockApiClient 구현"
```

---

## Task 7: Flutter — AuthService 추상화 + LoginScreen 리팩터링

**Files:**
- Create: `frontend/lib/shared/api/auth_service.dart`
- Create: `frontend/lib/shared/api/google_auth_service.dart`
- Create: `frontend/test/helpers/mock_auth_service.dart`
- Modify: `frontend/lib/features/auth/login_screen.dart`
- Modify: `frontend/lib/shared/providers/app_providers.dart`

- [ ] **Step 1: auth_service.dart 작성**

```dart
class AuthResult {
  final String idToken;
  final String? email;
  final String? displayName;
  AuthResult({required this.idToken, this.email, this.displayName});
}

abstract class AuthService {
  Future<void> initialize({String? clientId});
  Stream<AuthEvent> get authEvents;
  Future<void> authenticate();
  Future<void> signOut();
  bool get supportsAuthenticate;
  Widget renderSignInButton(AuthButtonConfig config);
}

enum AuthEventType { signIn, signOut, error }

class AuthEvent {
  final AuthEventType type;
  final AuthResult? result;
  final String? error;
  AuthEvent.signIn(this.result) : type = AuthEventType.signIn, error = null;
  AuthEvent.signOut() : type = AuthEventType.signOut, result = null, error = null;
  AuthEvent.error(this.error) : type = AuthEventType.error, result = null;
}

class AuthButtonConfig {
  final double width;
  AuthButtonConfig({this.width = 300});
}
```

- [ ] **Step 2: google_auth_service.dart 작성**

`login_screen.dart`의 GoogleSignIn 로직을 이동:
- `initialize()` — `GoogleSignIn.instance.initialize(clientId: clientId)`
- `authEvents` — `GoogleSignIn.instance.authenticationEvents`를 변환
- `authenticate()` — `GoogleSignIn.instance.authenticate()`
- `renderSignInButton()` — `web.renderButton(...)` 래핑

이 파일만 `google_sign_in`, `google_sign_in_web`을 import.

- [ ] **Step 3: mock_auth_service.dart 작성**

```dart
import 'dart:async';
import 'package:flutter/widgets.dart';
import 'package:lincle/shared/api/auth_service.dart';

class MockAuthService implements AuthService {
  final _controller = StreamController<AuthEvent>.broadcast();
  bool initializeCalled = false;

  @override
  Future<void> initialize({String? clientId}) async {
    initializeCalled = true;
  }

  @override Stream<AuthEvent> get authEvents => _controller.stream;
  @override Future<void> authenticate() async {}
  @override Future<void> signOut() async {}
  @override bool get supportsAuthenticate => false;
  @override Widget renderSignInButton(AuthButtonConfig config) => const SizedBox.shrink();

  void emitSignIn(AuthResult result) => _controller.add(AuthEvent.signIn(result));
  void dispose() => _controller.close();
}
```

- [ ] **Step 4: app_providers.dart에 authServiceProvider 추가**

```dart
final authServiceProvider = Provider<AuthService>((ref) {
  return GoogleAuthService();
});
```

- [ ] **Step 5: login_screen.dart 리팩터링**

```dart
// Before
import 'package:google_sign_in/google_sign_in.dart';
import 'package:google_sign_in_web/web_only.dart' as web;

// After
import '../../shared/api/auth_service.dart';
// google_sign_in import 제거
```

핸들러 내부:
```dart
// Before
final signIn = GoogleSignIn.instance;
await signIn.initialize(clientId: _clientId);
_authSub = signIn.authenticationEvents.listen(...)

// After
final authService = ref.read(authServiceProvider);
await authService.initialize(clientId: _clientId);
_authSub = authService.authEvents.listen(...)
```

버튼 렌더링:
```dart
// Before
web.renderButton(configuration: web.GSIButtonConfiguration(...))

// After
authService.renderSignInButton(AuthButtonConfig(width: 300))
```

- [ ] **Step 6: 컴파일 확인**

Run: `cd frontend && flutter analyze --no-fatal-infos`
Expected: No errors

- [ ] **Step 7: Commit**

```bash
git add frontend/lib/shared/api/auth_service.dart frontend/lib/shared/api/google_auth_service.dart frontend/lib/features/auth/login_screen.dart frontend/lib/shared/providers/app_providers.dart frontend/test/helpers/mock_auth_service.dart
git commit -m "feat(flutter): AuthService 추상화 — google_sign_in 의존성 격리"
```

---

## Task 8: Flutter — 위젯 테스트 활성화

**Files:**
- Modify: `frontend/test/widget_test.dart`
- Modify: `frontend/test/widget/login_screen_test.dart`
- Modify: `frontend/test/widget/listing_list_screen_test.dart`

- [ ] **Step 1: listing_list_screen_test.dart — skip 제거 + mock 주입**

```dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/shared/providers/app_providers.dart';
import 'package:lincle/features/listing/listing_list_screen.dart';
import '../../helpers/mock_api_client.dart';

void main() {
  group('ListingListScreen', () {
    Widget buildTestWidget() {
      return ProviderScope(
        overrides: [
          apiClientProvider.overrideWithValue(MockApiClient()),
        ],
        child: const MaterialApp(home: ListingListScreen()),
      );
    }

    testWidgets('renders app bar with logo image', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      expect(find.byType(Image), findsOneWidget);
    });

    testWidgets('shows FAB with "매물 등록" label', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      expect(find.text('매물 등록'), findsOneWidget);
      expect(find.byType(FloatingActionButton), findsOneWidget);
    });

    testWidgets('shows search icon in app bar', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      expect(find.byIcon(Icons.search), findsOneWidget);
    });

    testWidgets('shows loading indicator initially', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      // Before pumpAndSettle — loading state
      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });
  });
}
```

- [ ] **Step 2: login_screen_test.dart — skip 제거 + mock 주입**

```dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/shared/providers/app_providers.dart';
import 'package:lincle/features/auth/login_screen.dart';
import '../../helpers/mock_api_client.dart';
import '../../helpers/mock_auth_service.dart';

void main() {
  group('LoginScreen', () {
    Widget buildTestWidget() {
      return ProviderScope(
        overrides: [
          apiClientProvider.overrideWithValue(MockApiClient()),
          authServiceProvider.overrideWithValue(MockAuthService()),
        ],
        child: const MaterialApp(home: LoginScreen()),
      );
    }

    testWidgets('renders logo image', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      expect(find.byType(Image), findsOneWidget);
    });

    testWidgets('renders subtitle', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      expect(find.text('리니지 클래식 거래 플랫폼'), findsOneWidget);
    });

    testWidgets('shows 둘러보기 button', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      expect(find.text('둘러보기'), findsOneWidget);
    });

    testWidgets('Google login button hidden when no client ID', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();
      // MockAuthService.renderSignInButton returns SizedBox.shrink
      // Google button should not be visible
      expect(find.text('Google로 로그인'), findsNothing);
    });
  });
}
```

- [ ] **Step 3: widget_test.dart — skip 제거**

```dart
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/main.dart';
import 'package:lincle/shared/providers/app_providers.dart';
import 'helpers/mock_api_client.dart';
import 'helpers/mock_auth_service.dart';

void main() {
  testWidgets('App launches without error', (tester) async {
    await tester.pumpWidget(
      ProviderScope(
        overrides: [
          apiClientProvider.overrideWithValue(MockApiClient()),
          authServiceProvider.overrideWithValue(MockAuthService()),
        ],
        child: const MyApp(),
      ),
    );
    expect(find.byType(MaterialApp), findsOneWidget);
  });
}
```

주의: `main.dart`가 `google_sign_in`을 transitively import하지 않도록 확인 필요. LoginScreen이 AuthService를 사용하면 해결됨.

- [ ] **Step 4: 전체 Flutter 테스트 실행**

Run: `cd frontend && flutter test`
Expected: 52건 전체 PASS, 0 SKIP

- [ ] **Step 5: Commit**

```bash
git add frontend/test/
git commit -m "test(flutter): 위젯 테스트 10건 활성화 — mock 주입으로 skip 해소"
```

---

## Task 9: 문서 업데이트 + 최종 검증

**Files:**
- Modify: `docs/testing.md`

- [ ] **Step 1: 전체 테스트 실행**

```bash
cd backend && go test -v ./... 2>&1 | grep -c "PASS:"
cd web && npx vitest run
cd frontend && flutter test
```

Expected:
- Backend: 68 PASS, 0 SKIP
- Web: 23 files, 116 tests PASS
- Flutter: 52 PASS, 0 SKIP

- [ ] **Step 2: docs/testing.md 업데이트**

"현재 테스트 현황" 섹션의 skip 수치를 0으로 갱신.
모킹 컨벤션 섹션에 repository mock 패턴 추가.

- [ ] **Step 3: Commit**

```bash
git add docs/testing.md
git commit -m "docs: 테스트 문서 업데이트 — skip 0건 달성 반영"
```

---

## Execution Order & Dependencies

```
Task 1 (interfaces) ──┬── Task 2 (postgres impl)
                       ├── Task 3 (mock impl)
                       └────────┬── Task 4 (handler refactor + main.go)
                                └── Task 5 (handler tests)

Task 6 (IApiClient + mock) ──── Task 8 (listing widget tests)

Task 7 (AuthService + LoginScreen) ── Task 8 (login/app widget tests)

Task 9 (docs + final verify) — after all above
```

**Parallelizable groups:**
- Group A: Tasks 1→2→3→4→5 (Backend, sequential)
- Group B: Task 6 (Flutter ApiClient, independent)
- Group C: Task 7 (Flutter AuthService, independent)
- Group D: Task 8 (Flutter tests, depends on B+C)
- Group E: Task 9 (Final, depends on A+D)
