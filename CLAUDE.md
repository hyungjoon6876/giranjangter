# 기란장터 (Giranjangter)

리니지 클래식 아이템 거래 중개 플랫폼 — 무료, 커뮤니티 기반

## 기술 스택

- **Backend**: Go 1.25 + Gin (REST API, SSE) + pgx/v5
- **Web Frontend**: Next.js 16 + React 19 + TailwindCSS 3.4 + TanStack Query
- **Flutter Frontend**: Flutter 3.11 + Riverpod (모바일 앱)
- **Database**: PostgreSQL 16 (NAS Docker 컨테이너)
- **Infra**: Docker Compose + Caddy reverse proxy, NAS 배포 (2GB RAM)
- **Auth**: JWT + Google OAuth (Google Identity Services)

## 프로젝트 구조

```
backend/
  cmd/server/       # HTTP 핸들러 (handlers_*.go)
  internal/         # config, domain, guard, middleware, repository, event, oauth, alignment
  db/migrations/    # SQL 마이그레이션
web/                # Next.js 웹 프론트엔드 (메인)
  app/              # App Router 페이지
  components/       # UI 컴포넌트 (layout, listing, chat, ui, forms)
  lib/              # API 클라이언트, hooks, providers, types
  e2e/              # Playwright E2E 테스트
frontend/           # Flutter 모바일 앱
  lib/features/     # 기능별 모듈
  lib/shared/       # api/, providers/, theme/, widgets/, models/
admin/              # Next.js 운영 대시보드
shared/             # 공유 리소스 (design-tokens.json)
docs/               # 기술 설계 문서
deploy/             # NAS 배포 스크립트
```

## 핵심 규칙

- 상태 전이는 반드시 Guard를 통과해야 한다 (`internal/guard/`)
- RBAC 매트릭스를 따른다 → `docs/RBAC_ACTION_MATRIX.md`
- API 응답은 구조화 에러 형식: `{"error": {"code": "...", "message": "..."}}`
- 시크릿은 환경변수로만 관리 — .env 파일은 절대 커밋하지 않는다

## 코딩 컨벤션

전체 가이드: [docs/conventions.md](docs/conventions.md)

- **Go**: Handler Factory 패턴 `handleXxx(db) gin.HandlerFunc`, 파일명 `handlers_xxx.go`
- **Next.js/React**: App Router, `"use client"` 명시, TanStack Query로 서버 상태 관리
- **Dart**: Feature-driven 구조 `lib/features/xxx/`, ConsumerStatefulWidget + Riverpod
- **네이밍**: Go — snake_case 파일 + camelCase 함수, TSX — kebab-case 파일 + PascalCase 컴포넌트, Dart — snake_case 파일 + PascalCase 클래스
- **API 응답 null 방어**: `data?.data?.length` 패턴 사용 (API가 data: null 반환 가능)

## 테스트

```bash
# Backend
cd backend && go test ./...

# Web (유닛 + 컴포넌트)
cd web && npx vitest run

# Web (E2E — 배포된 서버 대상)
cd web && npx playwright test

# Flutter
cd frontend && flutter test
```

전체 테스팅 가이드: [docs/testing.md](docs/testing.md)

## 배포

### 필수 규칙: 무중단 롤링 배포

**절대 `docker-compose up -d`로 전체 서비스를 동시에 재시작하지 않는다.** 서비스가 내려가는 다운타임이 발생한다.

**정석 배포 절차** (`deploy/deploy.sh` 참조):

```
1. 이미지 빌드 (서비스 유지한 채로)
   docker-compose build --no-cache lincle-api lincle-web

2. Caddy 먼저 업데이트 (재시도 설정 반영)
   docker-compose up -d --no-deps caddy

3. API 롤링 재시작 → 헬스체크 통과 대기
   docker-compose up -d --no-deps lincle-api
   # /health 200 확인될 때까지 대기

4. Web 롤링 재시작 → 헬스체크 통과 대기
   docker-compose up -d --no-deps lincle-web
   # / 200 확인될 때까지 대기
```

핵심: **`--no-deps`로 서비스를 하나씩 재시작**. Caddy가 `lb_try_duration`으로 upstream 재시도하므로 짧은 재시작 동안 요청이 유실되지 않는다.

### 배포 대상

- **NAS**: `jym-nas` (192.168.50.222:9922)
- **포트**: 18090 (Caddy) → 8080 (API), 3000 (Web)
- **docker-compose 경로**: `/volume1/docker/lincle-deploy/`
- **sudo 필요**: `printf "password\n" | sudo -S docker-compose ...`

### 배포 전 체크리스트

1. **반드시 main 브랜치 기준으로 배포한다**
2. `cd web && npx vitest run` — 유닛 테스트 통과
3. `cd web && npx next build` — 빌드 성공
4. `cd web && npx playwright test` — E2E 테스트 통과 (선택)
5. `deploy/deploy.sh` 실행 또는 위 롤링 절차 수행

## 문서 참조 맵

| 문서 | 내용 |
|------|------|
| [docs/TECH_STACK_ARCHITECTURE.md](docs/TECH_STACK_ARCHITECTURE.md) | 기술 스택 + 인프라 설계 |
| [docs/OPENAPI_DRAFT.md](docs/OPENAPI_DRAFT.md) | MVP API 명세 (39 엔드포인트) |
| [docs/STARTER_DDL.md](docs/STARTER_DDL.md) | PostgreSQL 스키마 + 시드 데이터 |
| [docs/STATE_SEQUENCE_DIAGRAMS.md](docs/STATE_SEQUENCE_DIAGRAMS.md) | 상태머신 전이 규칙 |
| [docs/RBAC_ACTION_MATRIX.md](docs/RBAC_ACTION_MATRIX.md) | 역할별 권한 매트릭스 |
| [docs/EVENT_CATALOG.md](docs/EVENT_CATALOG.md) | 도메인 이벤트 카탈로그 |
| [docs/conventions.md](docs/conventions.md) | 코딩 컨벤션 |
| [docs/testing.md](docs/testing.md) | 테스팅 전략 |

## 판단이 어려울 때

- 지시가 충돌하면: 사용자에게 물어본다
- 문서의 파일 경로가 깨져있으면: 알리고, 추측하지 않는다
- 범위가 불확실하면: 적게 하고, 확인하고, 확장한다
- 상태 전이가 정의에 없으면: 임의로 추가하지 않고 확인한다
- DB 스키마 변경이 필요하면: 마이그레이션 파일로 관리한다
