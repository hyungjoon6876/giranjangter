# Test Infrastructure Improvement Design

**Date**: 2026-03-18
**Goal**: Backend 32 skip + Flutter 10 skip 해소

## 1. Backend — Repository Interface Extraction

### Problem
핸들러가 `*sql.DB`를 직접 사용하여 SQL 쿼리를 실행. Mock 불가능하여 32개 핸들러 테스트가 stub 상태.

### Design

**인터페이스 정의** (`backend/internal/repository/repository.go`):

도메인별 4개 인터페이스로 분리:

```go
type AuthRepo interface {
    FindUserByGoogleID(ctx context.Context, googleID string) (*User, error)
    CreateUser(ctx context.Context, googleID, email, nickname, profileURL string) (*User, error)
    FindUserByID(ctx context.Context, userID string) (*User, error)
    UpdateUser(ctx context.Context, userID string, nickname, profileURL *string) error
    SaveRefreshToken(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error
    DeleteRefreshToken(ctx context.Context, tokenHash string) error
    ValidateRefreshToken(ctx context.Context, tokenHash string) (string, error)
}

type ListingRepo interface {
    CreateListing(ctx context.Context, listing *CreateListingInput) (*Listing, error)
    GetListing(ctx context.Context, id string) (*Listing, error)
    UpdateListing(ctx context.Context, id string, fields map[string]interface{}) error
    ChangeListingStatus(ctx context.Context, id, newStatus string) error
    ListListings(ctx context.Context, filters ListingFilters) ([]Listing, string, error)
}

type ChatRepo interface {
    CreateChat(ctx context.Context, listingID, buyerID string) (*ChatRoom, error)
    FindExistingChat(ctx context.Context, listingID, buyerID string) (*ChatRoom, error)
    GetChat(ctx context.Context, chatID string) (*ChatRoom, error)
    ListChats(ctx context.Context, userID string) ([]ChatRoom, error)
    ListMessages(ctx context.Context, chatID string, cursor *string) ([]Message, string, error)
    SendMessage(ctx context.Context, chatID, senderID, text, clientMsgID string) (*Message, error)
}

type ReservationRepo interface {
    CreateReservation(ctx context.Context, chatID string, data *ReservationInput) (*Reservation, error)
    GetActiveReservation(ctx context.Context, listingID string) (*Reservation, error)
    ConfirmReservation(ctx context.Context, reservationID, userID string) error
    CancelReservation(ctx context.Context, reservationID, userID string) error
}
```

**구현체** (`backend/internal/repository/postgres.go`):

기존 핸들러 내 SQL 로직을 이동:

```go
type PostgresAuthRepo struct{ db *sql.DB }
type PostgresListingRepo struct{ db *sql.DB }
type PostgresChatRepo struct{ db *sql.DB }
type PostgresReservationRepo struct{ db *sql.DB }
```

**핸들러 시그니처 변경**:

```go
// Before
func handleLogin(db *sql.DB, auth *AuthMiddleware, cfg *Config) gin.HandlerFunc

// After
func handleLogin(repo repository.AuthRepo, auth *AuthMiddleware, cfg *Config) gin.HandlerFunc
```

**main.go 변경**:

```go
authRepo := repository.NewPostgresAuthRepo(db)
listingRepo := repository.NewPostgresListingRepo(db)
chatRepo := repository.NewPostgresChatRepo(db)
reservationRepo := repository.NewPostgresReservationRepo(db)

v1.POST("/auth/login", handleLogin(authRepo, auth, cfg))
```

**Mock 구현** (`backend/internal/repository/mock/mock.go`):

```go
type MockAuthRepo struct {
    FindUserByGoogleIDFn func(ctx context.Context, googleID string) (*repository.User, error)
    CreateUserFn         func(ctx context.Context, ...) (*repository.User, error)
    // ...
}

func (m *MockAuthRepo) FindUserByGoogleID(ctx context.Context, googleID string) (*repository.User, error) {
    return m.FindUserByGoogleIDFn(ctx, googleID)
}
```

**테스트 패턴**:

```go
func TestLogin_NewUser_Returns201(t *testing.T) {
    mock := &mock.MockAuthRepo{
        FindUserByGoogleIDFn: func(ctx context.Context, googleID string) (*repository.User, error) {
            return nil, sql.ErrNoRows // 새 유저
        },
        CreateUserFn: func(ctx context.Context, ...) (*repository.User, error) {
            return &repository.User{ID: "test-id", ...}, nil
        },
    }
    auth := newTestAuth()
    cfg := &config.Config{Env: "development"}

    w := httptest.NewRecorder()
    c, engine := gin.CreateTestContext(w)
    engine.POST("/api/v1/auth/login", handleLogin(mock, auth, cfg))
    // ... assert 201
}
```

### Domain Types

repository 패키지에 필요한 도메인 타입:

```go
type User struct {
    ID, GoogleID, Email, Nickname, ProfileImageURL, Role string
    CreatedAt time.Time
}

type Listing struct {
    ID, Title, SellerID, ServerID, Status string
    Price *int64
    // ... 필요한 필드
}

type ChatRoom struct { ... }
type Message struct { ... }
type Reservation struct { ... }
```

> 기존 `internal/domain/` 타입을 재사용하거나 repository 전용 타입 정의. 핸들러 내 인라인 struct를 공식 타입으로 승격.

---

## 2. Flutter — ApiClient Interface Extraction

### Problem
ApiClient가 concrete 클래스로 Dio를 직접 생성. 위젯 테스트에서 mock 불가 → 4건 skip.

### Design

**인터페이스** (`lib/shared/api/api_client_interface.dart`):

```dart
abstract class IApiClient {
  bool get isLoggedIn;
  String get staticBaseUrl;

  // Auth
  Future<Map<String, dynamic>> login(String provider, String token);
  Future<Map<String, dynamic>> getMe();
  Future<void> loadTokens();
  Future<void> saveTokens(String access, String refresh);
  Future<void> clearTokens();

  // Listings
  Future<Map<String, dynamic>> getListings({...});
  Future<Map<String, dynamic>> getListing(String id);
  Future<Map<String, dynamic>> createListing(Map<String, dynamic> data);
  Future<void> favoriteListing(String id);
  Future<void> unfavoriteListing(String id);

  // Chat
  Future<Map<String, dynamic>> createChat(String listingId);
  Future<Map<String, dynamic>> getChats();
  Future<Map<String, dynamic>> getMessages(String chatId, {String? cursor});
  Future<Map<String, dynamic>> sendMessage(String chatId, String text);

  // My data
  Future<Map<String, dynamic>> getMyListings({String? status});
  Future<Map<String, dynamic>> getMyTrades();
  Future<Map<String, dynamic>> getNotifications();
  Future<void> markNotificationsRead(List<String> ids);
  Future<Map<String, dynamic>> getMyReports();
  Future<Map<String, dynamic>> getUserReviews(String userId);

  // Reservation/Trade/Review
  Future<Map<String, dynamic>> createReservation(String chatId, Map<String, dynamic> data);
  Future<Map<String, dynamic>> completeTrade(String listingId, Map<String, dynamic> data);
  Future<Map<String, dynamic>> createReview(String completionId, Map<String, dynamic> data);

  // Other
  Future<Map<String, dynamic>> createReport(Map<String, dynamic> data);
  Future<List<dynamic>> getServers();
  Future<List<dynamic>> getCategories();
  Future<List<dynamic>> searchItems(String query, {String? categoryId});
}
```

**구현체 변경** (`lib/shared/api/api_client.dart`):

```dart
class ApiClient implements IApiClient {
  // 기존 코드 그대로, implements 추가만
}
```

**Provider 변경** (`lib/shared/providers/app_providers.dart`):

```dart
final apiClientProvider = Provider<IApiClient>((ref) => ApiClient());
```

**Mock** (`test/helpers/mock_api_client.dart`):

```dart
class MockApiClient implements IApiClient {
  @override bool get isLoggedIn => false;
  @override String get staticBaseUrl => 'http://test';

  @override
  Future<List<dynamic>> getServers() async => [
    {'id': 'test_server', 'name': '테스트서버'}
  ];

  @override
  Future<Map<String, dynamic>> getListings({...}) async => {
    'data': [], 'next_cursor': null
  };

  // ... 나머지 메서드는 기본 빈 응답 반환
}
```

**위젯 테스트 패턴**:

```dart
testWidgets('renders app bar with logo', (tester) async {
  await tester.pumpWidget(
    ProviderScope(
      overrides: [
        apiClientProvider.overrideWithValue(MockApiClient()),
      ],
      child: MaterialApp(home: ListingListScreen()),
    ),
  );
  expect(find.byType(Image), findsOneWidget);
});
```

---

## 3. Flutter — Google Sign-In Platform Abstraction

### Problem
LoginScreen이 `google_sign_in` 패키지를 직접 import. 이 패키지가 `dart:ui_web` (web 전용)에 의존하여 일반 테스트 러너에서 실행 불가 → 6건 skip.

### Design

**인터페이스** (`lib/shared/api/auth_service.dart`):

```dart
abstract class AuthService {
  Future<AuthResult?> signInWithGoogle();
  Future<void> signOut();
}

class AuthResult {
  final String idToken;
  final String? email;
  final String? displayName;
  AuthResult({required this.idToken, this.email, this.displayName});
}
```

**구현체** (`lib/shared/api/google_auth_service.dart`):

```dart
import 'package:google_sign_in/google_sign_in.dart';

class GoogleAuthService implements AuthService {
  final GoogleSignIn _googleSignIn;

  GoogleAuthService({String? clientId})
      : _googleSignIn = GoogleSignIn(clientId: clientId, scopes: ['email']);

  @override
  Future<AuthResult?> signInWithGoogle() async {
    final account = await _googleSignIn.signIn();
    if (account == null) return null;
    final auth = await account.authentication;
    return AuthResult(
      idToken: auth.idToken!,
      email: account.email,
      displayName: account.displayName,
    );
  }

  @override
  Future<void> signOut() => _googleSignIn.signOut();
}
```

**Provider**:

```dart
final authServiceProvider = Provider<AuthService>((ref) {
  final cfg = ref.watch(configProvider);
  return GoogleAuthService(clientId: cfg.googleClientId);
});
```

**LoginScreen 변경**:

```dart
// Before: import 'package:google_sign_in/google_sign_in.dart'; 직접 사용
// After: ref.read(authServiceProvider).signInWithGoogle() 사용
```

**Mock** (`test/helpers/mock_auth_service.dart`):

```dart
class MockAuthService implements AuthService {
  bool signInCalled = false;

  @override
  Future<AuthResult?> signInWithGoogle() async {
    signInCalled = true;
    return AuthResult(idToken: 'test-token', email: 'test@test.com');
  }

  @override
  Future<void> signOut() async {}
}
```

**테스트 패턴**:

```dart
testWidgets('renders logo', (tester) async {
  await tester.pumpWidget(
    ProviderScope(
      overrides: [
        apiClientProvider.overrideWithValue(MockApiClient()),
        authServiceProvider.overrideWithValue(MockAuthService()),
      ],
      child: MaterialApp(home: LoginScreen()),
    ),
  );
  expect(find.byType(Image), findsOneWidget);
});
```

---

## 영향 범위

| 영역 | 파일 수 | 변경 유형 |
|------|---------|----------|
| Backend repository 인터페이스 | ~3 신규 | 인터페이스 + 구현체 + mock |
| Backend 핸들러 시그니처 | ~5 수정 | `*sql.DB` → `repository.XxxRepo` |
| Backend main.go | 1 수정 | repo 생성 + 주입 |
| Backend 테스트 | ~4 수정 | stub → 실제 테스트 구현 |
| Flutter ApiClient 인터페이스 | ~2 신규 | interface + mock |
| Flutter AuthService | ~3 신규 | interface + impl + mock |
| Flutter providers | 1 수정 | 타입 변경 |
| Flutter LoginScreen | 1 수정 | google_sign_in 직접 사용 → AuthService |
| Flutter 테스트 | ~3 수정 | skip 제거 + mock 주입 |

## 성공 기준

- `go test ./...` — 68건 전체 PASS (skip 0)
- `flutter test` — 52건 전체 PASS (skip 0)
- 기존 `npx vitest run` — 116건 변함없이 PASS
- 앱 기능 동작에 영향 없음 (인터페이스만 추출, 로직 변경 없음)
