# 린클 기술 스택 & 아키텍처

> 확정 기술 스택 + 개인 NAS 배포 기준 아키텍처 (PostgreSQL 기반)
>
> 작성일: 2026-03-14
> 상태: 확정

---

## 1. 기술 스택 확정

### 1.1 Backend

| 항목 | 선택 | 근거 |
|------|------|------|
| **언어** | Go 1.22+ | 성능 최상위, 메모리 ~30MB, NAS 단일 바이너리 배포 |
| **웹 프레임워크** | Gin 또는 Fiber | Gin: 안정성/생태계, Fiber: Express 스타일로 빠른 개발 |
| **DB** | PostgreSQL 16 | ACID, JSONB, advisory lock, ENUM, 복잡한 JOIN. 거래 정합성 핵심 |
| **ORM/쿼리** | sqlc | SQL 직접 작성 + Go 타입 안전 코드 자동 생성. 성능 최적, SQL 완전 제어 |
| **캐시** | Redis 7 | 세션, 읽음 커서, rate limit, SSE 팬아웃, 알림 dedup |
| **실시간 (채팅)** | SSE (Server-Sent Events) | goroutine 기반으로 수천 동시 접속 가능. WebSocket보다 인프라 단순 |
| **인증** | JWT (access + refresh token) | 모바일 앱 친화, 세션리스 |
| **파일 저장** | 로컬 디스크 (NAS) → S3 (확장) | NAS 스토리지 직접 활용. 확장 시 S3-compatible 전환 |
| **API 문서** | OpenAPI 3.0 + swaggo | Go 코드에서 API 문서 자동 생성 |

### 1.2 Frontend

| 항목 | 선택 | 근거 |
|------|------|------|
| **프레임워크** | Flutter 3.x (Dart) | iOS + Android + Web 단일 코드베이스 |
| **상태관리** | Riverpod 2.0 | 타입 안전, 의존성 주입 내장, 테스트 용이 |
| **HTTP 클라이언트** | Dio | 인터셉터, 재시도, 토큰 갱신 지원 |
| **라우팅** | GoRouter | 딥링크 지원, 선언적 라우팅 |
| **채팅 UI** | 커스텀 (ListView + StreamBuilder) | 시스템 카드/예약 카드 혼합 필요 → 라이브러리보다 커스텀이 유리 |
| **로컬 저장** | Hive 또는 SharedPreferences | 토큰, 드래프트, 설정 저장 |
| **이미지** | cached_network_image | 캐싱 + 로딩 상태 |
| **코드 생성** | freezed + json_serializable | 불변 모델 + JSON 직렬화 자동 생성 |

### 1.3 인프라 / 배포

| 항목 | 선택 | 근거 |
|------|------|------|
| **배포 환경** | 개인 NAS (2GB RAM) | Docker Compose로 서비스 관리 |
| **컨테이너** | Docker + Docker Compose | NAS에서 직접 실행 |
| **리버스 프록시** | Caddy 또는 Nginx | 자동 HTTPS, 작은 메모리 풋프린트 |
| **CI/CD** | GitHub Actions → NAS SSH 배포 | 빌드는 GitHub, 바이너리만 NAS에 전송 |
| **도메인/SSL** | Cloudflare (무료) + Let's Encrypt | DNS + CDN + 무료 SSL |
| **모니터링** | 내장 /health + 로그 파일 | Grafana/Prometheus는 메모리 과다 → 로그 기반 모니터링 |
| **백업** | SQLite 파일 복사 (cron) | NAS 내 별도 볼륨 또는 외부 스토리지 |

### 1.4 관리자 도구

| 항목 | 선택 | 근거 |
|------|------|------|
| **백오피스** | Flutter Web (동일 코드베이스) 또는 별도 경량 웹 | Flutter Web으로 관리자 화면 포함 시 코드 재사용 극대화 |
| **대안** | Go 템플릿 기반 SSR 어드민 | 메모리 추가 없음, 별도 프론트 빌드 불필요 |

---

## 2. 메모리 예산 (NAS)

```
┌──────────────────────────────────────────┐
│              NAS 리소스 배분              │
├──────────────────────────────────────────┤
│ NAS OS + 시스템 프로세스    ~500MB        │
│ Docker Engine               ~100MB       │
│ Go 서버 (lincle-api)        ~30-50MB     │
│ PostgreSQL 16               ~200-300MB   │
│ Redis 7                     ~50-100MB    │
│ Caddy (리버스 프록시)        ~30MB        │
│ ─────────────────────────────            │
│ 사용 합계                   ~910MB-1.1GB  │
│                                          │
│ PostgreSQL 튜닝 권장:                     │
│   shared_buffers = 256MB                 │
│   work_mem = 8MB                         │
│   effective_cache_size = 1GB             │
│   max_connections = 50                   │
└──────────────────────────────────────────┘
```

→ NAS에 여유가 있으므로 PostgreSQL + Redis 풀 스택 운영 가능.

---

## 3. 아키텍처 다이어그램

```
                    ┌─────────────┐
                    │  Cloudflare │
                    │  (DNS/CDN)  │
                    └──────┬──────┘
                           │ HTTPS
                    ┌──────┴──────┐
                    │    Caddy    │
                    │ (리버스 프록시)│
                    └──────┬──────┘
                           │
              ┌────────────┴────────────┐
              │                         │
       ┌──────┴──────┐          ┌──────┴──────┐
       │  lincle-api │          │ Static Files│
       │   (Go 서버)  │          │ (Flutter Web)│
       │   :8080     │          │   :8081     │
       └──────┬──────┘          └─────────────┘
              │
       ┌──────┴──────┐     ┌─────────────┐
       │ PostgreSQL  │     │    Redis    │
       │   :5432     │     │   :6379    │
       └──────┬──────┘     └─────────────┘
              │
       ┌──────┴──────┐
       │  NAS 디스크   │
       │ (이미지/첨부) │
       └─────────────┘


 [모바일 앱]                        [웹 브라우저]
 Flutter iOS/Android                Flutter Web
      │                                  │
      └──────── HTTPS API ───────────────┘
                    │
              lincle-api (Go)
                    │
              ┌─────┴─────┐
              │ SSE Stream │  ← 채팅 실시간 수신
              └────────────┘
```

---

## 4. 프로젝트 구조

```
lincle/
├── docs/                          # 문서
│   ├── TECH_STACK_ARCHITECTURE.md  # 이 문서
│   └── ...
├── backend/                       # Go 백엔드
│   ├── cmd/
│   │   └── server/
│   │       └── main.go            # 엔트리포인트
│   ├── internal/
│   │   ├── config/                # 환경설정
│   │   ├── middleware/            # JWT, CORS, 로깅
│   │   ├── domain/                # 도메인 모델/엔티티
│   │   │   ├── listing/
│   │   │   ├── chat/
│   │   │   ├── reservation/
│   │   │   ├── review/
│   │   │   ├── report/
│   │   │   └── user/
│   │   ├── handler/               # HTTP 핸들러 (컨트롤러)
│   │   ├── service/               # 비즈니스 로직
│   │   ├── repository/            # DB 접근 계층
│   │   ├── guard/                 # 상태 전이 guard, 권한 체크
│   │   └── event/                 # 도메인 이벤트 발행
│   ├── db/
│   │   ├── migrations/            # SQL migration 파일
│   │   └── seed/                  # 시드 데이터
│   ├── api/
│   │   └── openapi.yaml           # API 스펙
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── frontend/                      # Flutter 프론트엔드
│   ├── lib/
│   │   ├── main.dart
│   │   ├── app/                   # 앱 설정, 라우팅
│   │   ├── features/              # 기능별 모듈
│   │   │   ├── auth/
│   │   │   ├── listing/
│   │   │   ├── chat/
│   │   │   ├── reservation/
│   │   │   ├── review/
│   │   │   ├── report/
│   │   │   ├── notification/
│   │   │   ├── profile/
│   │   │   └── admin/
│   │   ├── shared/                # 공통 위젯, 유틸
│   │   └── models/                # 데이터 모델 (freezed)
│   ├── pubspec.yaml
│   └── web/                       # Flutter Web 설정
├── docker-compose.yml             # NAS 배포용
├── Caddyfile                      # 리버스 프록시 설정
└── .env.example                   # 환경변수 템플릿
```

---

## 5. 핵심 기술 결정

### 5.1 PostgreSQL을 DB로 선택한 이유

| 장점 | 설명 |
|------|------|
| ACID 트랜잭션 | 거래 상태 전이의 정합성 보장 |
| ENUM 타입 | 상태값을 DB 레벨에서 강제 |
| JSONB | 메타데이터, 예약 부가정보 유연하게 저장 |
| Advisory Lock | 예약 확정 동시성 제어에 최적 |
| Foreign Key | 엔티티 관계 무결성 DB 레벨 보장 |
| 복잡한 쿼리 | 검색/필터/집계/서브쿼리 자유로움 |
| 생태계 | 백업, 모니터링, 마이그레이션 도구 풍부 |

| NAS 운영 팁 | 설명 |
|------------|------|
| 튜닝 | shared_buffers=256MB, work_mem=8MB |
| 백업 | pg_dump cron 스케줄 (일 1회) |
| 볼륨 | Docker volume으로 데이터 영속화 |
| 커넥션 | max_connections=50 (Go 커넥션풀 10-20) |

### 5.2 Redis 활용 전략

| 용도 | 키 패턴 | TTL |
|------|---------|-----|
| JWT refresh token | `auth:refresh:{userId}` | 30일 |
| 읽음 커서 캐시 | `chat:read:{chatRoomId}:{userId}` | 영구 (DB 동기) |
| Rate limit | `rate:{userId}:{action}` | 1분~1시간 |
| SSE 구독 관리 | `sse:sub:{userId}` | 연결 해제 시 삭제 |
| 알림 dedup | `notif:dedup:{eventKey}` | 5분 |
| 온라인 상태 | `presence:{userId}` | 5분 (heartbeat) |

→ Redis는 휘발성 데이터 전용. 재시작해도 서비스 영향 없도록 설계.

### 5.4 채팅 실시간 전략 (SSE + Go)

```go
// 핵심 구조
type SSEBroker struct {
    clients   map[string]chan SSEEvent  // userId → event channel
    mu        sync.RWMutex
}

// 메시지 전송 시:
// 1. DB에 메시지 저장
// 2. SSEBroker에서 상대방 채널 조회
// 3. 채널이 있으면 즉시 push (온라인)
// 4. 채널이 없으면 푸시 알림 발송 (오프라인)

// 클라이언트 재연결 시:
// GET /sse/connect?lastEventId=xxx
// → lastEventId 이후 누락 메시지를 DB에서 조회해서 먼저 전송
// → 이후 실시간 스트림 연결
```

- goroutine 당 메모리 ~2-8KB → 동시 1,000명 접속해도 ~8MB
- 단일 서버이므로 Redis Pub/Sub 불필요

### 5.5 Flutter ↔ Go API 타입 공유

Go와 Dart는 타입을 직접 공유할 수 없으므로 **OpenAPI codegen**으로 해결:

```
openapi.yaml (API 스펙)
    │
    ├──→ Go: oapi-codegen → 서버 핸들러 인터페이스 + 타입 생성
    │
    └──→ Dart: openapi-generator → API 클라이언트 + 모델 자동 생성
```

→ API 스펙을 Single Source of Truth로 유지하면 타입 불일치 방지

---

## 6. 확장 경로

### 6.1 NAS → 클라우드 전환

현재 NAS 아키텍처가 클라우드로 전환될 때 변경 최소화를 위한 설계:

| 계층 | NAS (현재) | 클라우드 (확장) | 변경 범위 |
|------|-----------|---------------|----------|
| DB | PostgreSQL (NAS Docker) | PostgreSQL (RDS/Supabase) | 접속 URL만 변경 |
| 캐시 | Redis (NAS Docker) | Redis (ElastiCache) | 접속 URL만 변경 |
| 파일 | NAS 디스크 | S3/R2 | storage 인터페이스 구현체 교체 |
| 배포 | Docker on NAS | Fly.io / Railway | Dockerfile 동일, 인프라만 변경 |
| 도메인 | DDNS + Cloudflare | 고정 IP + Cloudflare | DNS 설정만 변경 |
| SSL | Let's Encrypt (Caddy) | 동일 | 변경 없음 |
| 모니터링 | 로그 파일 | Sentry + Grafana | 로거 어댑터 추가 |

**핵심 원칙**: 모든 외부 의존성은 **인터페이스로 추상화**
```go
// 예시: storage 인터페이스
type FileStorage interface {
    Upload(ctx context.Context, key string, data io.Reader) (string, error)
    GetURL(ctx context.Context, key string) (string, error)
    Delete(ctx context.Context, key string) error
}

// NAS 구현
type LocalFileStorage struct { basePath string }

// S3 구현 (확장 시)
type S3FileStorage struct { client *s3.Client; bucket string }
```

### 6.2 트래픽 증가 시 확장 단계

```
Stage 1: NAS 단일 서버 (현재)
├── DAU < 2,000
├── PostgreSQL + Redis on Docker
└── 총 메모리 ~1-1.5GB

Stage 2: 외부 서버 전환
├── DAU 2,000-10,000
├── NAS 한계 도달 시
├── VPS/클라우드 + managed DB
├── CDN 이미지 서빙
└── 총 비용 ~10-30만원/월

Stage 4: 수평 확장
├── DAU 10,000+
├── 다중 인스턴스 + 로드밸런서
├── Redis Pub/Sub (SSE 팬아웃)
├── 이미지 CDN 완전 분리
└── 총 비용 ~50-100만원/월
```

---

## 7. 개발 환경 설정

### 7.1 필수 도구

```bash
# Go
go install golang.org/dl/go1.22@latest

# Flutter
flutter --version  # 3.x

# Docker
docker --version
docker compose version

# API 코드 생성
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Dart API 코드 생성
dart pub global activate openapi_generator_cli
```

### 7.2 로컬 개발 시작

```bash
# 백엔드
cd backend
cp .env.example .env
go run cmd/server/main.go
# → http://localhost:8080

# 프론트엔드
cd frontend
flutter pub get
flutter run -d chrome     # 웹
flutter run -d ios         # iOS 시뮬레이터
flutter run -d android     # Android 에뮬레이터
```

### 7.3 NAS 배포

```yaml
# docker-compose.yml
version: '3.8'
services:
  postgres:
    image: postgres:16-alpine
    ports:
      - "5432:5432"
    volumes:
      - pg_data:/var/lib/postgresql/data
    environment:
      POSTGRES_DB: lincle
      POSTGRES_USER: lincle
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    command: >
      postgres
        -c shared_buffers=256MB
        -c work_mem=8MB
        -c effective_cache_size=1GB
        -c max_connections=50
        -c random_page_cost=1.1
        -c log_min_duration_statement=1000
    restart: unless-stopped
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U lincle"]
      interval: 10s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --maxmemory 64mb --maxmemory-policy allkeys-lru
    restart: unless-stopped

  lincle-api:
    build: ./backend
    ports:
      - "8080:8080"
    volumes:
      - ./uploads:/app/uploads
    environment:
      - ENV=production
      - DATABASE_URL=postgres://lincle:${POSTGRES_PASSWORD}@postgres:5432/lincle?sslmode=disable
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=${JWT_SECRET}
      - UPLOAD_DIR=/app/uploads
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_started
    restart: unless-stopped

  caddy:
    image: caddy:2-alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile
      - ./frontend/build/web:/srv
      - caddy_data:/data
    restart: unless-stopped

volumes:
  pg_data:
  redis_data:
  caddy_data:
```

```
# Caddyfile
lincle.yourdomain.com {
    handle /api/* {
        reverse_proxy lincle-api:8080
    }
    handle /sse/* {
        reverse_proxy lincle-api:8080
    }
    handle {
        root * /srv
        try_files {path} /index.html
        file_server
    }
}
```

---

## 8. .env.example

```bash
# Server
ENV=development
PORT=8080
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8081

# Database
DATABASE_URL=postgres://lincle:lincle_dev@localhost:5432/lincle?sslmode=disable
POSTGRES_PASSWORD=lincle_dev

# Redis
REDIS_URL=redis://localhost:6379

# JWT
JWT_SECRET=change-me-in-production
JWT_ACCESS_TTL=15m
JWT_REFRESH_TTL=720h

# File Upload
UPLOAD_DIR=./uploads
MAX_UPLOAD_SIZE=10485760  # 10MB

# Push Notification (FCM)
FCM_SERVER_KEY=

# External (확장용)
# S3_ENDPOINT=
# S3_BUCKET=
# S3_ACCESS_KEY=
# S3_SECRET_KEY=
```
