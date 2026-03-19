# 하네스 품질 추적

이 문서는 하네스 평가 이력을 기록하여 시간에 따른 품질 변화를 추적한다.

---

## 평가 이력

### 2026-03-15 — 초기 평가

**종합 점수: 3.7 / 10.0**

| 레이어 | 등급 | 점수 |
|--------|------|------|
| 1. 에이전트 운영 | Missing | 0 |
| 2. 레포지토리 지식 | Adequate | 7 |
| 3. 워크플로우 | Missing | 0 |
| 4. 코딩 컨벤션 | Weak | 4 |
| 5. 테스팅 | Missing | 0 |
| 6. 보안 및 시크릿 | Adequate | 7 |
| 7. 의존성 관리 | Weak | 4 |
| 8. 에러 핸들링/로깅 | Weak | 4 |
| 9. 문서 일관성 | Adequate | 7 |
| 10. 하네스 안정성 | Weak | 4 |

**주요 발견**: 기술 설계 문서(API, DDL, 상태머신, RBAC)는 우수. 에이전트 운영 문서(CLAUDE.md, AGENTS.md, 컨벤션) 전무.

---

### 2026-03-15 — 초기화 후 재평가

**종합 점수: 7.9 / 10.0** (▲ +4.2)

| 레이어 | 이전 | 이후 | 점수 | 변화 |
|--------|------|------|------|------|
| 1. 에이전트 운영 | Missing | **Strong** | 10 | ▲ +10 |
| 2. 레포지토리 지식 | Adequate | **Strong** | 10 | ▲ +3 |
| 3. 워크플로우 | Missing | **Adequate** | 7 | ▲ +7 |
| 4. 코딩 컨벤션 | Weak | **Strong** | 10 | ▲ +6 |
| 5. 테스팅 | Missing | **Adequate** | 7 | ▲ +7 |
| 6. 보안 및 시크릿 | Adequate | Adequate | 7 | — |
| 7. 의존성 관리 | Weak | Weak | 4 | — |
| 8. 에러 핸들링/로깅 | Weak | **Adequate** | 7 | ▲ +3 |
| 9. 문서 일관성 | Adequate | **Strong** | 10 | ▲ +3 |
| 10. 하네스 안정성 | Weak | **Adequate** | 7 | ▲ +3 |

**생성 파일**: CLAUDE.md, AGENTS.md, ARCHITECTURE.md, docs/conventions.md, docs/testing.md, docs/core-beliefs.md, docs/workflows.md, docs/error-handling.md

**잔여 과제**: 의존성 관리 레이어 (Dependabot/Renovate 도입, 정책 문서화) — CI/CD 도입 시 함께 개선 권장

---

### 2026-03-15 — 리팩토링 후 재평가 (3차)

**종합 점수: 8.2 / 10.0** (▲ +0.3)

| 레이어 | 이전 | 이후 | 점수 | 변화 |
|--------|------|------|------|------|
| 1. 에이전트 운영 | Strong | Strong | 10 | — |
| 2. 레포지토리 지식 | Strong | Strong | 10 | — |
| 3. 워크플로우 | Adequate | Adequate | 7 | — |
| 4. 코딩 컨벤션 | Strong | Strong | 10 | — |
| 5. 테스팅 | Adequate | Adequate | 7 | — |
| 6. 보안 및 시크릿 | Adequate | Adequate | 7 | — |
| 7. 의존성 관리 | Weak | Adequate | 7 | ▲ +3 |
| 8. 에러 핸들링/로깅 | Adequate | Strong | 10 | ▲ +3 |
| 9. 문서 일관성 | Strong | Strong | 10 | — |
| 10. 하네스 안정성 | Adequate | Adequate | 7 | — |

**주요 변경사항**:
- 서비스명 변경: Lincle → 기란장터 (Giranjangter)
- DB: SQLite → PostgreSQL 16 (NAS Docker 컨테이너) 전환 완료
- 성향치 시스템 (alignment) 구현 + alignment 패키지 추가
- Google OAuth (google_sign_in v7) 연동
- 아이템 마스터 데이터 381개 (리니지 클래식 공식 사이트 크롤링)
- 코드 리팩토링: 핸들러 파일 분리, 트랜잭션 적용, 에러 처리 통일, JWT 파싱 중복 제거
- Flutter: 공유 유틸리티 추출 (listing_utils.dart), ApiClient 메서드 통합
- 디자인 개선안 문서 (docs/design-proposals.html) 작성

**개선된 항목**:
- 의존성 관리: pgx/v5 드라이버 도입, go.mod 정리
- 에러 핸들링: 모든 핸들러에 트랜잭션 적용, db.Exec 에러 체크 추가, SQL injection 패턴 제거

**잔여 과제**: CI/CD 파이프라인 (GitHub Actions), 디자인 시스템 적용 (다크 테마), 테스트 커버리지 확대

---

### 2026-03-20 — harness 4/4 완성 + seed 개선

**종합 점수: 9.1 / 10.0** (▲ +0.9)

| 레이어 | 이전 | 이후 | 점수 | 변화 |
|--------|------|------|------|------|
| 1. 에이전트 운영 | Strong | Strong | 10 | — |
| 2. 레포지토리 지식 | Strong | Strong | 10 | — |
| 3. 워크플로우 | Adequate | **Strong** | 10 | ▲ +3 |
| 4. 코딩 컨벤션 | Strong | Strong | 10 | — |
| 5. 테스팅 | Adequate | **Strong** | 9 | ▲ +2 |
| 6. 보안 및 시크릿 | Adequate | Adequate | 7 | — |
| 7. 의존성 관리 | Adequate | **Strong** | 9 | ▲ +2 |
| 8. 에러 핸들링/로깅 | Strong | Strong | 10 | — |
| 9. 문서 일관성 | Strong | Strong | 10 | — |
| 10. 하네스 안정성 | Adequate | **Strong** | 9 | ▲ +2 |

**주요 변경사항**:
- harness-verify: 3-Tier 자동 검증 훅 설치 (eslint, golangci-lint, tsc, go vet, vitest, go test)
- harness-constrain: 아키텍처 구조 규칙 5건 + 패턴 규칙 3건
- harness-loop: 자율 실행 태스크 큐 + Stop 훅
- harness-seed: 문서 5건 개선 (Redis 실태 반영, Web 컨벤션 추가, sqlite3 제거, Riverpod 버전 갱신, 테스트 구조 갱신)
- golangci-lint 기존 이슈 25건 수정 (errcheck, ineffassign)
- ESLint 기존 이슈 38건 수정 (hooks, unused vars, next/image, a11y)

**잔여 과제**: 테스트 커버리지 80% 달성 (tasks.json에 등록됨), CI/CD 파이프라인
