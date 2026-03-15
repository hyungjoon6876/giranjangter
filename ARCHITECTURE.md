# 아키텍처

Lincle는 Go REST API 백엔드 + Flutter 크로스플랫폼 프론트엔드로 구성된 P2P 거래 플랫폼이다.

상세 기술 스택 및 인프라 설계: [docs/TECH_STACK_ARCHITECTURE.md](docs/TECH_STACK_ARCHITECTURE.md)

## 시스템 구성도

```
[Flutter App] ──HTTP/SSE──▶ [Caddy Reverse Proxy] ──▶ [Go API Server] ──▶ [SQLite/PostgreSQL]
  (iOS/Android/Web)            (TLS, Static Files)     (Gin, JWT Auth)
```

## 백엔드 모듈 맵

```
backend/
├── cmd/server/              # 진입점 + HTTP 핸들러
│   ├── main.go              # 라우터 설정, 미들웨어 체인
│   ├── handlers_auth.go     # 인증 (로그인, 회원가입, OAuth)
│   ├── handlers_listing.go  # 거래글 CRUD + 상태 전이
│   ├── handlers_chat.go     # 채팅방 + 메시지
│   ├── handlers_reservation.go  # 예약
│   ├── handlers_trade.go    # 거래 완료
│   ├── handlers_review.go   # 리뷰
│   ├── handlers_report.go   # 신고
│   └── handlers_admin.go    # 관리자 기능
├── internal/
│   ├── config/              # 환경변수 로딩
│   ├── domain/              # 도메인 모델 + ENUM 정의
│   ├── guard/               # 상태 전이 유효성 검증
│   ├── middleware/          # JWT 인증, CORS, 역할 체크
│   ├── repository/          # DB 초기화 + 시딩
│   ├── event/               # SSE 이벤트 브로커
│   ├── oauth/               # Google OAuth 연동
│   └── alignment/           # 유저 평판 (성향) 시스템
└── db/
    ├── migrations/          # SQL 마이그레이션 파일 (순서 보장)
    └── seed/                # 테스트 시드 데이터
```

## 프론트엔드 모듈 맵

```
frontend/lib/
├── app/
│   └── router.dart          # GoRouter 라우팅 설정
├── features/                # 기능별 독립 모듈
│   ├── auth/                # 로그인, 회원가입
│   ├── listing/             # 거래글 목록, 상세, 작성
│   ├── chat/                # 채팅
│   ├── reservation/         # 예약
│   ├── review/              # 리뷰
│   ├── report/              # 신고
│   ├── notification/        # 알림
│   ├── profile/             # 프로필
│   └── admin/               # 관리자
├── shared/
│   ├── api/                 # ApiClient (Dio + 인터셉터)
│   ├── providers/           # 전역 Riverpod Provider
│   ├── theme/               # Material 3 테마
│   ├── widgets/             # 공용 위젯
│   └── models/              # 공유 데이터 모델
└── main.dart                # 앱 진입점
```

## 핵심 데이터 흐름

### 거래글 상태 전이

```
available → reserved → pending_trade → completed
    │           │            │
    └───────────┴────────────┴──→ cancelled
```

상세 전이 규칙: [docs/STATE_SEQUENCE_DIAGRAMS.md](docs/STATE_SEQUENCE_DIAGRAMS.md)

### 요청 흐름

```
Client Request
  → Caddy (TLS 종료, 라우팅)
    → Gin Router
      → Auth Middleware (JWT 검증)
        → Role Middleware (RBAC 체크)
          → Handler (비즈니스 로직)
            → Guard (상태 전이 검증)
              → DB (SQL 실행)
                → JSON Response
```

### 실시간 통신

```
Client ──SSE 연결──▶ Event Broker ◀── Handler (이벤트 발행)
```

## 의존성 방향

```
handlers → middleware → guard → domain
    │          │                   ▲
    └──────────┴───── repository ──┘
                         │
                    event broker
```

- `domain`은 다른 패키지에 의존하지 않는다
- `guard`는 `domain`만 의존한다
- `handlers`는 모든 내부 패키지를 사용할 수 있다
