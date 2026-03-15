# 코딩 컨벤션

## Go (Backend)

### 파일 네이밍
- 핸들러: `handlers_<도메인>.go` (예: `handlers_listing.go`, `handlers_auth.go`)
- 가드: `<도메인>_guard.go` (예: `listing_guard.go`)
- 미들웨어: 기능 단위 파일 (예: `auth.go`, `cors.go`)
- 도메인 모델: `models.go` (통합), 또는 `<도메인>.go` (분리 시)

### 함수/타입 네이밍
- 핸들러 함수: `handleXxx` (unexported) — `func handleCreateListing(db *sql.DB) gin.HandlerFunc`
- 요청 DTO: unexported struct — `type createListingRequest struct`
- 도메인 ENUM: `type XxxStatus string` + const 블록 — `ListingAvailable`, `ListingReserved`
- 전이 맵: `AllowedXxxTransitions` (exported)

### 핸들러 패턴 (Handler Factory)

```go
func handleCreateListing(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := middleware.GetUserID(c)

        var req createListingRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
            })
            return
        }

        // 비즈니스 로직
        // ...

        c.JSON(http.StatusCreated, result)
    }
}
```

### 미들웨어 체인

```go
// 읽기 전용 (인증 불필요)
v1.GET("/listings", handleListListings(db))

// 쓰기 (인증 + 제한 체크)
write := v1.Group("")
write.Use(auth.RequireAuthWithDB(db))
write.Use(middleware.RejectIfRestricted())
{
    write.POST("/listings", handleCreateListing(db))
}
```

### 에러 응답 형식

```go
c.JSON(http.StatusXxx, gin.H{
    "error": gin.H{
        "code":    "ERROR_CODE",    // 대문자 스네이크
        "message": "사람이 읽을 수 있는 메시지",
    },
})
```

주요 에러 코드: `VALIDATION_ERROR`, `UNAUTHORIZED`, `FORBIDDEN`, `NOT_FOUND`, `CONFLICT`, `INTERNAL_ERROR`

---

## Dart/Flutter (Frontend)

### 파일 네이밍
- 화면: `<기능>_<유형>_screen.dart` (예: `listing_list_screen.dart`)
- 위젯: `<이름>_widget.dart` 또는 용도별 (예: `server_filter_widget.dart`)
- 모델: `<이름>.dart` (예: `listing.dart`)
- Provider: `<이름>_provider.dart` 또는 통합 파일

### 클래스/변수 네이밍
- 클래스: PascalCase — `ListingListScreen`, `ApiClient`
- 변수/메서드: camelCase — `_loadListings`, `_selectedServer`
- Provider: `xxxProvider` — `serversProvider`, `currentUserProvider`
- private 멤버: `_` 접두사

### 위젯 패턴 (Riverpod)

```dart
class ListingListScreen extends ConsumerStatefulWidget {
  const ListingListScreen({super.key});

  @override
  ConsumerState<ListingListScreen> createState() => _ListingListScreenState();
}

class _ListingListScreenState extends ConsumerState<ListingListScreen> {
  @override
  Widget build(BuildContext context) {
    final servers = ref.watch(serversProvider);
    // ...
  }
}
```

### Provider 정의

```dart
final apiClientProvider = Provider<ApiClient>((ref) => ApiClient());

final serversProvider = FutureProvider<List<dynamic>>((ref) async {
  final api = ref.watch(apiClientProvider);
  return api.getServers();
});
```

### 디렉토리 구조 규칙
- 새 기능은 `lib/features/<기능명>/` 하위에 생성
- 2개 이상 기능에서 공유하는 코드는 `lib/shared/`로 이동
- 기능 내부 구조는 자유 (screen, widget, model 등)

---

## 공통

### 임포트 정렬
- **Go**: 표준 라이브러리 → 외부 패키지 → 내부 패키지 (빈 줄 구분)
- **Dart**: dart: → package: → relative imports

### 코드 생성 (Flutter)
- freezed, json_serializable 변경 후: `dart run build_runner build --delete-conflicting-outputs`
- 생성 파일 (`*.g.dart`, `*.freezed.dart`)은 직접 수정 금지
