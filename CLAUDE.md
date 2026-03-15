# 기란장터 (Giranjangter)

리니지 클래식 아이템 거래 중개 플랫폼 — 무료, 커뮤니티 기반

## 기술 스택

- **Backend**: Go 1.25 + Gin (REST API, SSE) + pgx/v5
- **Frontend**: Flutter 3.11 + Riverpod (Flutter Web 우선, 이후 모바일)
- **Database**: PostgreSQL 16 (NAS Docker 컨테이너)
- **Infra**: Docker + Caddy reverse proxy, NAS 배포 (2GB RAM)
- **Auth**: JWT + Google OAuth (google_sign_in v7)

## 프로젝트 구조

```
backend/
  cmd/server/       # HTTP 핸들러 (handlers_*.go)
  internal/         # config, domain, guard, middleware, repository, event, oauth, alignment
  db/migrations/    # SQL 마이그레이션
frontend/
  lib/features/     # 기능별 모듈 (auth, listing, chat, reservation, review, report, notification, admin)
  lib/shared/       # api/, providers/, theme/, widgets/, models/
  lib/app/          # 라우터 설정 (GoRouter)
docs/               # 기술 설계 문서 (API, DDL, RBAC, 상태머신, 이벤트)
tasks/              # 구현 계획 및 태스크 추적
```

## 핵심 규칙

- 상태 전이는 반드시 Guard를 통과해야 한다 (`internal/guard/`)
- RBAC 매트릭스를 따른다 → `docs/RBAC_ACTION_MATRIX.md`
- API 응답은 구조화 에러 형식: `{"error": {"code": "...", "message": "..."}}`
- 시크릿은 환경변수로만 관리 — .env 파일은 절대 커밋하지 않는다

## 코딩 컨벤션

전체 가이드: [docs/conventions.md](docs/conventions.md)

- **Go**: Handler Factory 패턴 `handleXxx(db) gin.HandlerFunc`, 파일명 `handlers_xxx.go`
- **Dart**: Feature-driven 구조 `lib/features/xxx/`, ConsumerStatefulWidget + Riverpod
- **네이밍**: Go — snake_case 파일 + camelCase 함수, Dart — snake_case 파일 + PascalCase 클래스

## 테스트

```bash
# Backend
cd backend && go test ./...

# Frontend
cd frontend && flutter test
```

전체 테스팅 가이드: [docs/testing.md](docs/testing.md)

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
