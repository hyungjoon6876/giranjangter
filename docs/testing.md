# 테스팅 전략

## 실행 명령

```bash
# Backend 전체 테스트
cd backend && go test ./...

# Backend 특정 패키지
cd backend && go test ./internal/guard/...

# Web 유닛/컴포넌트 테스트
cd web && npx vitest run

# Web 감시 모드
cd web && npx vitest

# Web E2E 테스트 (배포된 서버 대상)
cd web && npx playwright test

# Web E2E UI 모드
cd web && npx playwright test --ui

# Web E2E 커스텀 서버 지정
E2E_BASE_URL=http://localhost:3000 npx playwright test

# Flutter 전체 테스트
cd frontend && flutter test

# Flutter 특정 파일
cd frontend && flutter test test/xxx_test.dart
```

## 테스트 유형

| 유형 | Backend (Go) | Web (Next.js) | Flutter | 용도 |
|------|-------------|---------------|---------|------|
| 단위 테스트 | `*_test.go` (같은 패키지) | `__tests__/**/*.test.{ts,tsx}` | `test/*_test.dart` | 개별 함수/클래스 검증 |
| 컴포넌트 테스트 | N/A | `__tests__/components/*.test.tsx` | `test/widget/*_test.dart` | UI 컴포넌트 렌더링 검증 |
| E2E 테스트 | N/A | `e2e/*.spec.ts` (Playwright) | N/A | 사용자 시나리오 검증 |
| 핸들러 테스트 | Mock repository 기반 | N/A | N/A | 핸들러 HTTP 동작 검증 |

## 현재 테스트 현황

### Backend (Go) — 68 통과, 0 skip

**Guard/Domain/Infra 테스트:**
- `internal/guard/listing_guard_test.go` — 상태 전이 검증 (listing + reservation)
- `internal/middleware/auth_test.go` — JWT 생성/파싱, RequireAuth/OptionalAuth, 역할 체크
- `internal/domain/alignment_test.go` — 성향 등급 계산
- `internal/event/broker_test.go` — SSE 브로커 구독/메시징 (5건)

**핸들러 테스트 (Mock Repository 기반):**
- `cmd/server/handlers_auth_test.go` (8건) — 로그인 (신규/기존/dev/prod 모드), 토큰 갱신, 프로필 조회
- `cmd/server/handlers_chat_test.go` (7건) — 채팅방 생성/충돌/인증/권한, 메시지 조회/전송/중복
- `cmd/server/handlers_listing_test.go` (11건) — 상태 변경 (유효/무효/비소유자), 수정 권한, 입력 검증
- `cmd/server/handlers_reservation_test.go` (6건) — 예약 생성/충돌, 확인/취소 권한

### Web (Next.js) — 가장 활발 (23파일 116테스트 전통과)

**컴포넌트 테스트 (18개):** `web/__tests__/components/`
- badge, bottom-nav-badge, chat-input, chat-list-item, chat-message, chat-panel
- empty-state, error-state, header-badge, listing-card, listing-filters
- loading, modal, report-modal, reservation-modal, review-modal, skeleton, toast

**라이브러리/훅 테스트 (4개):** `web/__tests__/lib/`
- `api-client.test.ts` — URL 구성, 파라미터 처리
- `api-client-advanced.test.ts` — 토큰 관리, Authorization 헤더, HTTP 상태 처리
- `use-sse.test.ts` — EventSource 모킹, 지수 백오프, 재연결
- `utils.test.ts` — formatPrice, formatTimeAgo, statusLabel, statusColor

**페이지 테스트 (1개):** `web/__tests__/app/`
- `login.test.tsx` — 로그인 페이지 렌더링

**E2E 테스트 (5개):** `web/e2e/`
- `auth-guard.spec.ts` — 인증 보호 (리다이렉트, 비보호 페이지 접근)
- `home.spec.ts` — 홈페이지 (히어로, 필터, 정렬, 빈 상태, 검색)
- `listing-detail.spec.ts` — 거래글 상세 (접근성 랜드마크, 즐겨찾기)
- `login.spec.ts` — 로그인 (Google 버튼, 둘러보기 내비게이션)
- `navigation.spec.ts` — 내비게이션/접근성 (skip-to-content, aria-current)

### Flutter — 51 통과, 1 skip

**유닛 테스트:**
- `test/api_client_test.dart` (11통과, 1skip*) — Dio 설정, baseUrl, 헤더, 타임아웃, 토큰 저장/로드/삭제
- `test/utils_test.dart` (31통과) — formatPrice, statusLabel, statusColor, chatStatusLabel, formatTimeAgo

**위젯 테스트 (Mock 주입 기반):**
- `test/widget_test.dart` (1통과) — 앱 스모크 테스트
- `test/widget/login_screen_test.dart` (4통과) — 로고, 서브타이틀, 둘러보기, Google 로그인 버튼
- `test/widget/listing_list_screen_test.dart` (4통과) — AppBar 로고, FAB, 검색 아이콘, 빈 상태

> *1 skip: `createChat validateStatus` — Dio mock 필요한 설계 의도 테스트. 기능 동작과 무관.

## 테스트 설정

### Web — Vitest

`web/vitest.config.ts`:
- 환경: `jsdom`
- 셋업: `vitest.setup.ts` (`@testing-library/jest-dom/vitest` 매처)
- 경로 별칭: `@` → 프로젝트 루트
- 포함 패턴: `__tests__/**/*.test.{ts,tsx}`

### Web — Playwright

`web/playwright.config.ts`:
- 대상 서버: `E2E_BASE_URL` 환경변수 (기본값: `http://192.168.50.222:18090`)
- 타임아웃: 30초
- 재시도: 1회
- 스크린샷: 실패 시만
- 트레이스: 첫 번째 재시도 시
- 브라우저: Chromium만

> E2E는 배포된 NAS 서버를 대상으로 실행. 로컬에서 돌리려면 `E2E_BASE_URL=http://localhost:3000`으로 지정.

## 파일 네이밍

- **Go**: `<원본파일>_test.go` — 같은 디렉토리에 위치 (예: `listing_guard_test.go`)
- **TSX/TS**: `<원본파일>.test.tsx` — `__tests__/` 하위에 구조 미러링 (예: `modal.test.tsx`)
- **E2E**: `<기능>.spec.ts` — `e2e/` 디렉토리 (예: `home.spec.ts`)
- **Dart**: `<원본파일>_test.dart` — `test/` 디렉토리에 미러링

## 테스트 우선 대상

다음 영역은 반드시 테스트를 작성한다:

1. **Guard (상태 전이)** — 허용/금지 전이를 모두 검증
2. **Middleware (인증/인가)** — JWT 검증, 역할 체크
3. **핸들러 입력 검증** — 필수 필드 누락, 잘못된 형식
4. **도메인 로직** — 비즈니스 규칙 (예: 자기 거래글에 예약 불가)
5. **UI 컴포넌트** — 렌더링, 접근성 속성, 사용자 인터랙션
6. **API 클라이언트** — 요청 구성, 토큰 관리, 에러 처리
7. **E2E 핵심 흐름** — 인증 가드, 내비게이션, 주요 페이지 렌더링

## 테스트 작성 원칙

- 테스트 함수명: `Test<함수명>_<시나리오>` (Go) / `describe > it('설명')` (Vitest) / `<설명>` (Dart)
- 하나의 테스트는 하나의 동작만 검증
- 테스트 데이터는 테스트 함수 내에서 생성 (전역 상태 공유 금지)
- DB 테스트는 트랜잭션으로 감싸서 격리

## 모킹 컨벤션

### Go
- Repository 인터페이스 + 함수 필드 Mock 패턴 (`internal/repository/mock/`)
- `newTestAuth()` 헬퍼로 테스트용 AuthMiddleware 생성
- `authMiddleware(userID, role)` 헬퍼로 인증된 요청 시뮬레이션
- `httptest.NewRecorder()` + `gin.TestMode`로 HTTP 핸들러 테스트

### Web (Vitest)
- `vi.fn()` / `vi.spyOn()` — 함수/메서드 모킹
- `global.fetch = vi.fn()` — HTTP 요청 모킹
- EventSource 클래스 모킹 — SSE 훅 테스트
- `localStorage` 모킹 — 토큰 저장 테스트
- `vi.useFakeTimers()` — 타이머 의존 로직 (지수 백오프 등)
- `QueryClient` + `QueryClientProvider` 래퍼 — TanStack Query 훅 테스트

### Dart
- `SharedPreferences.setMockInitialValues({})` — 로컬 저장소 모킹
- `IApiClient` 인터페이스 + `MockApiClient` (`test/helpers/mock_api_client.dart`)
- `AuthService` 인터페이스 + `MockAuthService` (`test/helpers/mock_auth_service.dart`)
- Riverpod `ProviderScope.overrides`로 Provider 교체
- Conditional import로 `google_sign_in_web` (dart:ui_web) 플랫폼 격리

## CI/CD 테스트 자동화

현재 CI/CD 파이프라인 없음. 테스트는 배포 전 수동으로 실행한다:

```bash
# 배포 전 필수 체크
cd web && npx vitest run        # 유닛 테스트 통과
cd web && npx next build        # 빌드 성공
cd web && npx playwright test   # E2E 통과 (선택)
cd backend && go test ./...     # 백엔드 테스트 통과
```
