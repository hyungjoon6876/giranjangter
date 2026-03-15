# 하네스 평가 리포트
> 생성일: 2026-03-15 | 레포지토리: lincle

## 종합 하네스 점수: 3.7 / 10.0

## 등급 요약

| 레이어 | 등급 | 점수 | 핵심 이슈 |
|--------|------|------|-----------|
| 1. 에이전트 운영 | Missing | 0 | CLAUDE.md, AGENTS.md 전무 — 에이전트 진입점 없음 |
| 2. 레포지토리 지식 | Adequate | 7 | 풍부한 기술 문서, 단 루트 진입점 없음 |
| 3. 워크플로우 | Missing | 0 | CI/CD, 브랜치 전략, 배포 프로세스 미문서화 |
| 4. 코딩 컨벤션 | Weak | 4 | 코드에 일관된 패턴 존재하나 문서화 안됨 |
| 5. 테스팅 | Missing | 0 | 테스팅 전략·구조·커버리지 요구사항 전무 |
| 6. 보안 및 시크릿 | Adequate | 7 | .env.example + RBAC 문서화, 에이전트 보안 규칙 미비 |
| 7. 의존성 관리 | Weak | 4 | Lock 파일 존재, 정책/자동 업데이트 없음 |
| 8. 에러 핸들링/로깅 | Weak | 4 | 코드 패턴 일관적이나 미문서화 |
| 9. 문서 일관성 | Adequate | 7 | docs/ 내 중복 없으나 교차 참조 체계 부재 |
| 10. 하네스 안정성 | Weak | 4 | 루트 네비게이션 없음, 고아 문서 존재 |
| **종합** | | **3.7** | |

---

## Part 1: 레이어별 평가

### 레이어 1: 에이전트 운영 레이어
**등급: Missing (0)**

- **근거**: CLAUDE.md, AGENTS.md 파일 부재. `.claude/` 디렉토리 존재하나 .md 파일 없음
- **발견 사항**:
  - 에이전트 행동 규칙 없음
  - 태스크 실행 루프(계획→구현→검증→문서화) 미정의
  - 자율성 경계 미정의 — 에이전트가 무엇을 할 수 있고 없는지 알 수 없음
  - 불명확 지시 시 폴백 규칙 없음
- **위반 원칙**: 충분한 컨텍스트(1), 안전한 기본값(8), 발견 가능성(6)

---

### 레이어 2: 레포지토리 지식 레이어
**등급: Adequate (7)**

- **근거**: `docs/TECH_STACK_ARCHITECTURE.md` (489줄), `docs/STATE_SEQUENCE_DIAGRAMS.md` (503줄), `docs/OPENAPI_DRAFT.md` (795줄), `docs/RBAC_ACTION_MATRIX.md` (505줄), `docs/STARTER_DDL.md` (1025줄), `docs/EVENT_CATALOG.md` (196줄)
- **발견 사항**:
  - 기술 스택, 인프라, 데이터베이스, API, 상태머신, 이벤트 등 상세 문서 존재
  - 루트 레벨 ARCHITECTURE.md 또는 README.md 없어 진입점이 불명확
  - docs/ 내 문서들이 상세하지만 에이전트가 "어디서부터 읽어야 하는지" 가이드 없음
  - 모듈/패키지 경계 설명이 TECH_STACK_ARCHITECTURE.md에 일부 포함
- **위반 원칙**: 발견 가능성(6), 점진적 공개(5)

---

### 레이어 3: 워크플로우 레이어
**등급: Missing (0)**

- **근거**: `.github/workflows/` 미존재, CONTRIBUTING.md 미존재, Makefile 미존재
- **발견 사항**:
  - CI/CD 파이프라인 없음
  - 브랜치 전략 미문서화
  - PR 프로세스 미정의
  - 배포 프로세스: docker-compose.yml과 Caddyfile 존재하나 실행 절차 미문서화
  - `tasks/IMPLEMENTATION_PLAN.md` (66KB)에 구현 로드맵 있으나 개발 워크플로우는 아님
- **위반 원칙**: 충분한 컨텍스트(1), 구체적 지시(4)

---

### 레이어 4: 코딩 컨벤션 레이어
**등급: Weak (4)**

- **근거**: `frontend/analysis_options.yaml` (기본 flutter_lints), 코드 샘플링 결과
- **발견 사항**:
  - **미문서화 관례 (코드에서 발견됨)**:
    - Go: Handler Factory 패턴 (`handleXxx(db *sql.DB) gin.HandlerFunc`)
    - Go: 파일명 `handlers_xxx.go`, `xxx_guard.go` (snake_case + domain prefix)
    - Go: 구조화 에러 응답 `gin.H{"error": gin.H{"code": "...", "message": "..."}}`
    - Dart: Feature-driven 디렉토리 구조 (`lib/features/xxx/`)
    - Dart: ConsumerStatefulWidget + Riverpod 패턴
    - Dart: Provider 접미사 컨벤션 (`xxxProvider`)
  - analysis_options.yaml는 기본값만 사용 — 커스텀 규칙 없음
  - Go 린터 설정 없음 (golangci-lint 등)
  - .editorconfig 없음
- **위반 원칙**: 구체적 지시(4), 단일 진실 원천(2)

---

### 레이어 5: 테스팅 레이어
**등급: Missing (0)**

- **근거**: `frontend/test/` 디렉토리 존재하나 테스트 전략/구조 문서 없음
- **발견 사항**:
  - 테스팅 전략 문서 없음
  - 테스트 유형(단위/통합/E2E) 정의 없음
  - 커버리지 요구사항 없음
  - 모킹/픽스처 컨벤션 없음
  - Go 테스트 설정 미확인 (테스트 파일 여부 불명)
  - CI에 테스트 단계 없음 (CI 자체 부재)
- **위반 원칙**: 충분한 컨텍스트(1), 구체적 지시(4)

---

### 레이어 6: 보안 및 시크릿 레이어
**등급: Adequate (7)**

- **근거**: `.env.example`, `.gitignore` (다중 레벨), `docs/RBAC_ACTION_MATRIX.md`
- **발견 사항**:
  - `.env.example` 존재: ENV, PORT, DATABASE_URL, JWT 설정, CORS 등 문서화
  - `.gitignore`에서 .env, 키스토어 파일 등 적절히 제외
  - RBAC 매트릭스로 인가 체계 상세 문서화 (40+ 액션 코드, 6개 역할)
  - JWT 인증 구현 확인 (middleware/auth.go)
  - BUT: 에이전트 전용 보안 규칙 없음 ("시크릿을 코드에 넣지 마라" 등)
  - 의존성 취약점 스캐닝 없음 (Dependabot 등)
- **위반 원칙**: 안전한 기본값(8)

---

### 레이어 7: 의존성 관리 레이어
**등급: Weak (4)**

- **근거**: `backend/go.mod`, `backend/go.sum`, `frontend/pubspec.yaml`
- **발견 사항**:
  - Go: go.mod에 의존성 명시, go.sum으로 체크섬 고정
  - Flutter: pubspec.yaml에 의존성 명시 (^버전 범위)
  - BUT: 의존성 정책 문서 없음
  - Lock 파일 커밋 정책 미문서화
  - 자동 의존성 업데이트 없음 (Dependabot/Renovate 미설정)
  - 승인/금지 의존성 목록 없음
  - Go 1.25.7 vs Dockerfile의 Go 1.22-alpine — **버전 불일치 감지**
- **위반 원칙**: 코드-문서 일관성(3), 구체적 지시(4)

---

### 레이어 8: 에러 핸들링 및 로깅 레이어
**등급: Weak (4)**

- **근거**: 코드 샘플링 (handlers_*.go, api_client.dart)
- **발견 사항**:
  - **미문서화 관례 (코드에서 발견됨)**:
    - Go: 구조화 에러 `gin.H{"error": gin.H{"code": "...", "message": "..."}}` 일관 사용
    - Go: HTTP 상태 코드 매핑 (400=검증, 401=인증, 404=미발견, 500=서버)
    - Go: DB 트랜잭션 롤백 패턴
    - Dart: try-catch + setState 패턴
    - Dart: Dio 인터셉터로 401 자동 토큰 갱신
  - 에러 핸들링 패턴 문서 없음
  - 로깅 프레임워크/컨벤션 없음 (fmt.Printf? log? zerolog?)
  - 로그 레벨 가이드라인 없음
- **위반 원칙**: 구체적 지시(4), 단일 진실 원천(2)

---

## Part 2: 종합 분석

### 문서 일관성 (레이어 9)
**등급: Adequate (7)**

- docs/ 내 7개 문서는 중복 없이 각자 명확한 주제를 다룸
- 직접적 충돌 발견되지 않음
- BUT: 문서 간 교차 참조 체계 부재 (각 문서가 독립적으로 존재)
- `docs/prd-chat-sync-notes.md`가 "PRD.md"를 참조하나 해당 파일 미존재 → **깨진 참조**
- 루트 README.md 없어 docs/ 문서들의 권한 체계 불명확
- TECH_STACK_ARCHITECTURE.md가 PostgreSQL 배포를 설명하나 현재 코드는 SQLite 사용 → **잠재적 불일치** (마이그레이션 계획일 수 있음)

### 하네스 안정성 (레이어 10)
**등급: Weak (4)**

- 루트 레벨 진입 문서(README.md) 없음
- docs/ 파일명은 자기설명적이나 ALL_CAPS 사용이 비표준
- 고아 문서: 모든 docs/ 파일이 다른 문서에서 참조되지 않음 (인덱스 없음)
- `prd-chat-sync-notes.md` → PRD.md 깨진 참조
- TECH_STACK_ARCHITECTURE.md ↔ 실제 코드 간 DB 불일치 (PostgreSQL vs SQLite)
- 문서 계층 구조: flat docs/ + tasks/ — 계층은 단순하나 네비게이션 없음

### 하네스 약점 (우선순위순)

| # | 이슈 | 심각도 | 레이어 |
|---|------|--------|--------|
| 1 | CLAUDE.md 없음 — 에이전트 진입점 부재 | Critical | L1 |
| 2 | AGENTS.md 없음 — 에이전트 역할/경계 미정의 | Critical | L1 |
| 3 | 루트 README.md 없음 — 프로젝트 개요 부재 | High | L2, L10 |
| 4 | CI/CD 없음 — 자동화 테스트/배포 불가 | High | L3 |
| 5 | 테스팅 전략 없음 — 품질 보증 기준 부재 | High | L5 |
| 6 | 코딩 컨벤션 미문서화 — 코드 일관성 보장 불가 | Medium | L4 |
| 7 | 에러 핸들링 패턴 미문서화 | Medium | L8 |
| 8 | 의존성 정책 없음 | Medium | L7 |
| 9 | docs/ 네비게이션 인덱스 없음 | Low | L10 |
| 10 | Go 버전 불일치 (go.mod vs Dockerfile) | Low | L7, L9 |
| 11 | PRD.md 깨진 참조 | Low | L9 |

### 재설계 필요성 판단

**점진적 개선 가능**. 기존 docs/ 문서는 품질이 높고 잘 구조화되어 있다. 기존 문서를 건드리지 않고 에이전트 운영 레이어(CLAUDE.md, AGENTS.md)와 루트 진입점을 추가하면 점수가 크게 향상된다. 구조 재편은 불필요하다.

### 권고사항 (초기화 단계 우선순위)

**P0 (필수) — 에이전트 진입점:**
1. `CLAUDE.md` 생성 — 에이전트 운영 규칙 + docs/ 참조 맵 + 안전한 기본값
2. `AGENTS.md` 생성 — 에이전트 역할, 경계, 태스크 루프

**P1 (중요) — 핵심 운영 지식:**
3. `ARCHITECTURE.md` 생성 — 시스템 구조 요약 (TECH_STACK_ARCHITECTURE.md 참조)
4. `docs/conventions.md` 생성 — 코드 패턴 문서화 (코드 샘플링 기반)
5. `docs/testing.md` 생성 — 테스팅 전략 정의
6. `docs/core-beliefs.md` 생성 — 팀 핵심 엔지니어링 원칙

**P2 (보조) — 보충 컨텍스트:**
7. `docs/workflows.md` 생성 — 개발/배포 워크플로우
8. `docs/error-handling.md` 생성 — 에러/로깅 컨벤션
9. docs/ 교차 참조 체계 구축 (CLAUDE.md에서 문서 맵 제공)
