# 린클 거래소 MVP Claude Code 개발 실행 마스터 스펙 v2

부제: PRD 재서술 문서가 아니라, Claude Code가 실제 저장소에서 **문서 해석 → 실행계획 수립 → 코드 수정 → 테스트 → 상태 업데이트**까지 연속 수행하도록 만드는 상세 개발 실행 지시서

버전: v2  
작성 목적: 다운로드 이슈 대응 재배포 + Claude Code 실개발 관점 세부화  
권장 위치: `docs/lincle/CLAUDE_CODE_EXECUTION_MASTER_SPEC_v2.md`

---

## 1. 이 문서의 역할

이 문서는 다음 세 가지를 동시에 만족시키기 위해 작성한다.

- PRD와 기술 설계 문서를 **개발 착수용 실행계획**으로 변환한다.
- Claude Code가 저장소 안에서 **실제로 손을 움직이게 하는 작업 규칙**을 제공한다.
- 세션이 바뀌어도 사람이 이어받을 수 있도록 **산출물, 상태, 리스크, 결정사항**을 문서로 남기게 한다.

이 문서는 다음을 목적으로 하지 않는다.

- PRD를 보기 좋게 다시 요약하는 것
- 기술 스택을 새로 제안하는 것
- 아키텍처를 전면 재설계하는 것
- 문서 없는 정책을 상상으로 확정하는 것

핵심은 한 문장으로 정리된다.

> **Claude Code는 설명자나 조언자가 아니라, 현재 저장소 기준으로 MVP를 가장 빨리 전진시키는 구현 담당자여야 한다.**

---

## 2. 언제 이 문서를 쓰는가

이 문서는 아래 세 가지 상황에서 사용한다.

### 2.1 Claude Code 세션 시작 지시문으로 사용
- 이 문서 전체를 Claude Code에 넣는다.
- 이어서 `PRD.md`, `TECH_STACK_ARCHITECTURE.md` 및 companion docs를 읽게 한다.
- 첫 응답 뒤 곧바로 코드와 문서를 수정하게 한다.

### 2.2 저장소 내부 실행 기준서로 사용
아래 경로 중 하나를 권장한다.

- `docs/lincle/CLAUDE_CODE_EXECUTION_MASTER_SPEC_v2.md`
- `docs/lincle/IMPLEMENTATION_MASTER_SPEC_v2.md`
- `CLAUDE.md` 내부의 린클 전용 섹션

### 2.3 사람 검토 기준서로 사용
Claude Code가 생성한 실행계획, 상태 문서, 코드 수정, 테스트 결과를 리뷰할 때 이 문서를 체크리스트로 쓴다.

---

## 3. 소스 오브 트루스와 우선순위

Claude Code는 아래 문서를 우선순위대로 읽고 충돌을 정리해야 한다.

### 3.1 필수 입력 문서
1. `docs/lincle/PRD.md`
2. `docs/lincle/TECH_STACK_ARCHITECTURE.md`

### 3.2 가능하면 함께 읽을 문서
1. `docs/lincle/OPENAPI_DRAFT.md`
2. `docs/lincle/STARTER_DDL.md`
3. `docs/lincle/RBAC_ACTION_MATRIX.md`
4. `docs/lincle/RETENTION_CONFIG_APPENDIX.md`
5. `docs/lincle/EVENT_CATALOG.md`
6. `docs/lincle/STATE_SEQUENCE_DIAGRAMS.md`

### 3.3 저장소에서 반드시 확인할 증거
- 디렉터리 구조
- 백엔드/프론트엔드 프레임워크
- 현재 라우팅 구조
- 인증/세션 처리 방식
- migration 체계
- schema/model/entity 정의
- config/env 체계
- 테스트 러너와 lint/typecheck 체계
- Docker / compose / CI
- seed / fixture / demo data
- 관리자 화면 또는 운영 도구의 기존 흔적
- logging / event / audit / notification 관련 코드

### 3.4 충돌 시 판단 순서
1. 최신 의사결정 문서가 명확하면 문서를 우선한다.
2. 이미 동작하는 코드가 기준선으로 더 신뢰할 만하면 최소 변경으로 수렴한다.
3. 어느 쪽도 확신이 없으면 `SPEC_GAP_LOG.md`에 기록하고, MVP에 더 안전한 방향으로 임시 수렴한다.
4. 미확정 이슈가 있다고 전체 진행을 멈추지 말고, 독립적으로 진행 가능한 작업을 먼저 전진시킨다.

---

## 4. 절대 규칙

Claude Code는 아래 규칙을 어기면 안 된다.

### 4.1 PRD 재작성 금지
- PRD를 길게 다시 설명하지 않는다.
- 중복 요약에 시간을 쓰지 않는다.
- 구현과 연결되는 정보만 추출한다.

### 4.2 기술 스택 재설계 금지
- 프레임워크 교체, ORM 교체, 상태관리 교체, 인증체계 교체 같은 대공사는 제안만 하고 실행하지 않는다.
- 현재 저장소 기준으로 **최소 변경**으로 MVP를 전진시킨다.

### 4.3 과도한 상상 금지
- 결제, 정산, 법무, 본인인증, 보존기간, 신고처리 정책 같은 고위험 규칙을 문서 없이 추정 확정하지 않는다.
- 확정이 필요한 항목은 `확인 필요`로 남기고, 안전한 placeholder 구조만 구현한다.

### 4.4 계획으로 끝내지 말 것
- 첫 응답 뒤 바로 파일 수정에 들어간다.
- 문서 생성만 하고 멈추지 않는다.
- 가능한 범위는 코드, 테스트, 상태 문서까지 전부 남긴다.

### 4.5 증거 없이 완료 선언 금지
다음이 없으면 완료가 아니다.
- 수정 파일 목록
- 테스트 또는 검증 흔적
- 남은 리스크
- 다음 시작점
- 업데이트된 상태 문서

### 4.6 작은 작업 묶음 유지
- 하나의 패치는 하나의 목적을 가진다.
- unrelated refactor는 하지 않는다.
- 큰 리팩터링보다 작은 vertical slice를 우선한다.

### 4.7 서버에서 권한 강제
- UI 숨김만으로 끝내지 않는다.
- 소유자/상대방/관리자/일반 사용자 권한을 서버에서 강제한다.
- 상태 전이는 반드시 백엔드에서 검증한다.

### 4.8 문서와 코드 동기화
다음 중 하나를 바꾸면 관련 문서도 같이 바꾼다.
- API
- DB schema/migration
- 상태 enum/transition
- RBAC action code
- config key / retention / feature flag
- 관리자 처리 규칙
- 이벤트명 / audit field

---

## 5. Claude Code의 기본 동작 모드

Claude Code는 아래 순서로 움직여야 한다.

### 5.1 세션 시작 루프
1. 저장소와 입력 문서를 스캔한다.
2. 현재 상태를 10줄 내외로 요약한다.
3. 바로 구현 가능한 범위와 선행 확정 항목을 분리한다.
4. `docs/lincle/MVP_EXECUTION_PLAN_v1.md`를 생성 또는 갱신한다.
5. `docs/lincle/BUILD_STATUS.md`, `DECISION_LOG.md`, `SPEC_GAP_LOG.md`, `TEST_RUN_REPORT.md`를 생성 또는 갱신한다.
6. Phase 0부터 실제 구현에 착수한다.
7. 각 작업 묶음이 끝날 때마다 테스트를 실행하고 상태를 업데이트한다.
8. 막히는 항목은 gap으로 남기고 다음 독립 작업으로 넘어간다.

### 5.2 중단 가능한 조건
아래 경우만 세션 속도를 늦출 수 있다.
- 외부 비밀값이나 계정이 없어서 실행 자체가 불가능한 경우
- 필수 외부 서비스가 없고 로컬 mock도 만들 수 없는 경우
- 법무/정산/신원인증처럼 고위험 정책이 핵심 로직에 직접 박혀야 하는데 문서가 전혀 없는 경우
- 현재 repo baseline이 깨져 있어 unrelated failure가 먼저 해결되어야 하는 경우

위 경우에도 반드시 해야 할 일:
- 무엇이 막혔는지 `SPEC_GAP_LOG.md`에 남긴다.
- 우회 가능한 범위가 있으면 계속 진행한다.
- 중단 범위를 최소화한다.

### 5.3 세션 첫 응답 형식
Claude Code의 첫 응답은 아래 구조를 따라야 한다.

1. 현재 프로젝트 상태 요약
2. 지금 바로 구현 가능한 범위
3. 선행 확정이 필요한 항목
4. 이번 세션 우선 작업 3~7개
5. 바로 착수할 첫 구현 묶음
6. 생성/갱신할 상태 문서 목록

첫 응답은 승인 요청 문구가 아니라 **착수 선언**이어야 한다.

---

## 6. 저장소 스캔 체크리스트

Claude Code는 초기 15~30분 안에 아래를 파악해야 한다.

### 6.1 구조 파악
- 루트 앱 구조와 패키지 경계
- `apps/`, `packages/`, `services/`, `web/`, `mobile/`, `backend/` 유사 디렉터리
- 실행 엔트리
- 환경별 설정 파일
- 공통 라이브러리 위치

### 6.2 백엔드 파악
- API 라우터/컨트롤러/서비스/리포지토리 구조
- ORM 또는 쿼리 빌더 종류
- migration 생성/실행 방법
- 인증 미들웨어
- RBAC 또는 policy guard
- 에러 응답 포맷
- pagination 표준
- 이벤트/알림 발행 구조
- audit/logging 삽입 포인트

### 6.3 프론트엔드 파악
- 라우팅 구조
- auth guard
- API client
- 폼/검증 라이브러리
- 전역 상태 또는 서버 상태 관리
- 목록/상세/모달/토스트 같은 공통 UI
- 관리자 화면 존재 여부

### 6.4 테스트/운영 파악
- lint / typecheck / unit / integration / e2e 명령
- dev / test / prod 환경변수 템플릿
- docker-compose / devcontainer
- CI job
- seed / fixture / mock
- health check / readiness
- 로그 필드와 correlation id

### 6.5 스캔 예시 명령
실제 명령은 저장소 기술 스택에 맞춰 바꾸되, 아래 수준의 증거를 남긴다.

```text
pwd
ls -la
tree -L 3
rg "router|controller|service|migration|schema|entity|model" .
rg "RBAC|role|permission|policy|guard" .
rg "event|audit|logger|notification" .
rg "transaction|order|reservation|chat|message|report|review" .
cat package.json
cat pnpm-workspace.yaml
cat pyproject.toml
cat docker-compose.yml
cat README.md
```

스캔 결과는 `BUILD_STATUS.md`의 `현재 기준선 파악` 섹션에 짧게 남긴다.

---

## 7. 반드시 남겨야 하는 산출물

아래 문서는 문서만 보기 위한 문서가 아니라 **다음 세션과 사람 리뷰를 위한 운영 흔적**이다.

### 7.1 필수 문서
- `docs/lincle/MVP_EXECUTION_PLAN_v1.md`
- `docs/lincle/BUILD_STATUS.md`
- `docs/lincle/DECISION_LOG.md`
- `docs/lincle/SPEC_GAP_LOG.md`
- `docs/lincle/TEST_RUN_REPORT.md`

### 7.2 상황에 따라 추가
- `docs/lincle/API_CHANGELOG.md`
- `docs/lincle/ADMIN_RUNBOOK.md`
- `docs/lincle/OBSERVABILITY_CHECKLIST.md`
- `docs/lincle/RELEASE_CHECKLIST.md`
- `docs/lincle/RELEASE_BLOCKERS.md`

### 7.3 코드 산출물 예시
- OpenAPI 수정
- migration 파일
- schema/entity/model
- API route/controller/service
- state enum 및 transition guard
- RBAC action constants / policy mapping
- config registry / env example
- frontend route/page/component/form
- admin list/detail/action 화면
- logging / audit / event hooks
- tests / fixtures / seed
- runbook / release checklist

---

## 8. 작업 우선순위 프레임

우선순위는 아래 순서를 기본으로 한다.

### 8.1 재작업 비용이 큰 계약 축
- OpenAPI
- DB migration
- 상태머신
- RBAC/action code
- config key / retention / feature flag
- 이벤트명 / 감사 필드

### 8.2 핵심 사용자 가치가 보이는 얇은 vertical slice
- 계정 진입
- 상품 등록
- 상품 목록/상세
- 권한 확인

### 8.3 거래 성립을 위한 커뮤니케이션
- 1:1 채팅

### 8.4 거래 핵심 제어
- 거래 상태 전이
- 경쟁 조건 방지
- 이벤트 / audit

### 8.5 출시 리스크 완화
- 리뷰
- 신고
- 운영/관리자 도구
- 모니터링/로그

### 8.6 안정화
- 회귀 테스트
- seed/demo data
- release checklist
- runbook

---

## 9. Pre-build freeze 상세 지시

이 섹션은 “구현보다 먼저 얼려야 하는 것”을 정의한다. 전부 100% 완벽해야만 다음으로 가는 것은 아니다. 하지만 **최소 실행 가능한 초안**은 반드시 있어야 한다.

### 9.1 OpenAPI 최종화
왜 먼저 필요한가:
- 프론트엔드와 백엔드의 계약 기준선이다.
- 필드명, nullable, 에러 포맷이 흔들리면 재작업이 크다.

반드시 정리할 것:
- endpoint 목록
- request/response shape
- auth 필요 여부
- 권한 조건
- pagination 방식
- 공통 에러 포맷
- 상태 변경 API의 요청/응답

완료 조건:
- 핵심 MVP 도메인에 대해 “없어서 구현을 못 하는 endpoint”가 남지 않는다.
- 최소한 Phase 1~3의 endpoint는 초안이 정리된다.
- 문서와 실제 handler 시그니처의 차이가 있으면 changelog 또는 gap log에 남긴다.

### 9.2 DB migration 초안 확정
왜 먼저 필요한가:
- 엔티티 관계, nullable, 인덱스, soft delete 정책이 이후 모든 기능에 영향을 준다.

반드시 정리할 것:
- 핵심 테이블과 관계
- user / product / conversation / message / transaction / review / report / admin action
- timestamp / soft delete / audit 필드
- unique constraint / foreign key / index 후보
- migration 생성/rollback 기준

완료 조건:
- 신규 환경에서 from-scratch 적용이 가능하다.
- 최소 주요 엔티티의 필드명이 API와 충돌하지 않는다.
- seed 또는 demo data 전략의 방향이 있다.

### 9.3 상태머신 freeze
왜 먼저 필요한가:
- 거래 흐름은 이 시스템의 핵심이며, 상태 전이가 흔들리면 프론트/백엔드/운영이 동시에 깨진다.

반드시 정리할 것:
- 상태 enum
- 허용 전이
- 금지 전이
- 상태별 허용 액션
- 누가 어떤 전이를 트리거할 수 있는지
- 취소/완료/분쟁 진입의 정책 placeholder

완료 조건:
- 상태 전이 표 또는 상수 파일이 존재한다.
- 서버 guard가 이를 참조하도록 설계된다.
- 대표 금지 전이 테스트를 작성할 수 있다.

### 9.4 RBAC / action code freeze
왜 먼저 필요한가:
- 사용자/관리자/소유자/상대방 권한이 혼재된 도메인이므로 권한 체계가 흐리면 보안과 운영 리스크가 커진다.

반드시 정리할 것:
- role 정의
- action code 목록
- 자원별 소유권 규칙
- 관리자 기능 접근 규칙
- 거래/채팅/신고/리뷰 액션별 권한 기준

완료 조건:
- action constants가 코드로 존재한다.
- 최소한 서버 guard 또는 policy check에 연결된다.
- 관리자 외 접근 금지 엔드포인트가 명시된다.

### 9.5 config / retention / feature flag freeze
왜 먼저 필요한가:
- 환경 변수와 정책 기본값이 뒤늦게 바뀌면 운영 및 테스트가 불안정해진다.

반드시 정리할 것:
- env key 이름
- 기본값
- secret / non-secret 구분
- retention 관련 값
- feature flag 필요 여부
- 개발/테스트/운영에서 달라지는 값

완료 조건:
- `.env.example` 또는 config registry가 존재한다.
- 이름과 기본값이 문서에 정리된다.
- 임시값이라도 명시적으로 관리된다.

### 9.6 이벤트 / audit field 초안
왜 먼저 필요한가:
- 거래, 신고, 관리자 조치, 채팅의 추적 가능성은 출시 리스크와 직결된다.

반드시 정리할 것:
- 핵심 이벤트명
- actor id / target id / action code / timestamp
- request id / correlation id 필요 여부
- 관리자 조치 로그 필드
- 신고/거래 상태 변경 로그 포인트

완료 조건:
- 최소 이벤트 카탈로그나 상수 목록이 존재한다.
- 로그 또는 audit trail에 필요한 필드가 정리된다.

### 9.7 owner / approver 지정
왜 먼저 필요한가:
- 기술 결정이 코드에서 살아도 사람이 승인 주체를 모르면 다시 흔들린다.

반드시 정리할 것:
- OpenAPI owner
- DB schema owner
- 거래 상태 owner
- RBAC owner
- 운영 정책 owner
- 승인 필요 여부

완료 조건:
- 문서상 최소 이름 또는 역할 단위 owner가 적혀 있다.
- 사람이 없는 경우 `가정 owner` 또는 `확인 필요`로라도 남긴다.

---

## 10. Phase 0. 계약/기반 확정

### 목표
재작업 위험이 큰 계약 계층과 공통 기반을 얼린다. 이후의 모든 기능 구현이 이 기준선 위에서 움직이게 만든다.

### 포함 범위
- OpenAPI 초안 정리
- 핵심 migration 초안
- 상태 enum / transition 정의
- RBAC/action code 정의
- env/config key 초안
- audit/event 포인트 정의
- 에러 응답 및 pagination 기준 정리
- repo 기동/테스트 기준 확인

### 선행조건
- 필수 문서 읽기 완료
- 현재 repo 구조 파악
- 기존 auth / DB / routing 방식 파악

### 백엔드 작업
- 핵심 도메인 엔티티 정리
- migration 파일 또는 DDL 초안 작성
- 공통 에러 코드와 응답 포맷 정리
- 상태 guard 구조 도입
- action constants 도입
- audit hook / event emit placeholder 삽입

### 프론트엔드 작업
- 라우팅 큰 그림 확정
- auth guard / role gate 방식 정리
- 공통 API client 에러 처리 방식 확인
- form/list/detail/status badge의 최소 공통 컴포넌트 초안

### 운영/관리 작업
- 관리자 조회가 필요한 리소스 정의
- 감사 추적 포인트 확인
- 운영자 계정/역할 정책 placeholder 기록

### 테스트/검증
- repo 기동
- migration apply / rollback
- lint / typecheck
- 상태 전이 금지 케이스 최소 테스트
- 권한 없는 접근 차단 테스트

### 완료 기준
- 핵심 계약 축에 대한 초안이 모두 존재한다.
- 다음 Phase 기능 구현자가 필드/상태/권한 때문에 계속 멈추지 않는다.
- 문서와 코드의 차이가 gap log에 드러난다.

### 대표 산출물
- OpenAPI 업데이트
- migration 초안
- state constants / transition map
- RBAC constants
- config registry / `.env.example`
- `DECISION_LOG.md` 초기 기록
- `SPEC_GAP_LOG.md` 초기 항목

---

## 11. Phase 1. 계정/상품/탐색

### 목표
사용자가 계정 맥락 안에서 상품을 등록하고, 목록에서 찾고, 상세를 보고, 소유자 권한을 확인할 수 있는 MVP의 첫 visible slice를 만든다.

### 포함 범위
- 프로필 조회/수정 최소 API
- 상품 등록
- 상품 수정
- 상품 목록
- 상품 상세
- 검색/필터 최소 기능
- 소유자 권한 체크
- 기본 상태/카테고리/가격 노출

### 포함하지 않을 수 있는 것
- 복잡한 추천
- 고급 검색 랭킹
- 이미지 업로드의 완전한 미디어 파이프라인
- SEO / 성능 최적화 대공사

### 백엔드 작업
- product create/update/detail/list API
- validation rules
- ownership check
- soft delete 또는 hide 정책 반영
- 검색/정렬/필터 파라미터 정리
- 미디어 필드는 placeholder 또는 연결부까지만 처리

### 프론트엔드 작업
- 로그인 이후 진입 경로
- 상품 등록 폼
- 상품 목록
- 상품 상세
- 빈 상태 / 오류 상태 / 권한 오류
- 상태 badge / 카테고리 / 가격 / 작성자 표시

### 운영/관리 작업
- 부적절 상품 hide 또는 moderation placeholder
- 신고 진입점과 연결 가능성 확보

### 테스트/검증
- 본인 상품만 수정 가능
- 잘못된 데이터 validation
- 목록 pagination / 빈 결과 / 잘못된 id
- 숨김/삭제 상품 노출 정책
- API와 화면 간 필드 매핑 점검

### 완료 기준
- 사용자가 최소한 상품을 등록하고 다시 찾고 상세를 볼 수 있다.
- 소유자 권한과 비소유자 권한이 서버에서 구분된다.
- 상품 도메인의 OpenAPI와 실제 응답이 맞는다.

---

## 12. Phase 2. 1:1 채팅

### 목표
거래 전 조율을 위한 대화 흐름을 제공한다. 이 Phase는 이후 거래/신고/감사 추적의 맥락을 형성한다.

### 포함 범위
- 대화방 생성
- 대화방 목록
- 메시지 조회
- 메시지 전송
- 참여자 권한 검증
- 상품 맥락 표시
- 최소 알림 이벤트 또는 placeholder

### 핵심 정책 초안
- 동일 상품/동일 상대 기준 중복 방 허용 여부
- 방 생성 주체
- 읽음 처리 최소 범위
- 삭제/차단 사용자 처리

### 백엔드 작업
- conversation / message 모델
- 참여자 검증
- 중복 방지 규칙
- message pagination
- 최소 읽음 처리 또는 unread count placeholder
- 이벤트 hook

### 프론트엔드 작업
- 채팅방 목록
- 채팅 상세
- 입력창/전송
- 상품 컨텍스트 표시
- 상대방 표시
- 에러/빈 상태 처리

### 운영/관리 작업
- 채팅 신고로 이어질 수 있도록 식별자와 audit 포인트 확보
- retention 정책 반영 위치 명시

### 테스트/검증
- 제3자 접근 차단
- 동일 사용자 간 중복 방 정책
- 긴 대화 pagination
- validation 실패
- 차단/삭제 케이스 placeholder 동작

### 완료 기준
- 사용자가 거래 전 대화를 시작하고 이어갈 수 있다.
- 대화 기록이 신고/감사/거래 맥락과 연결 가능한 구조를 가진다.

---

## 13. Phase 3. 거래 상태머신

### 목표
린클 거래소 MVP의 핵심 제어 로직을 서버 주도 상태 전이로 구현한다.

### 포함 범위
- 거래 생성 또는 예약 시작
- 진행 중 상태
- 완료
- 취소
- 분쟁 진입 placeholder 또는 최소 구현
- 상태별 CTA와 권한 분기
- 이벤트 / audit trail

### 핵심 원칙
- 임의 상태 점프 금지
- 상태 전이 검증은 서버에서 먼저 수행
- 권한 검증은 role + ownership + 상대방 맥락까지 고려
- 이중 요청과 경쟁 조건을 방지
- 거래 상태 변화는 반드시 traceable해야 함

### 백엔드 작업
- transaction / reservation 도메인 모델
- 상태 enum
- transition guard
- 상태 변경 API
- idempotency 또는 중복 요청 방지 전략
- audit/event 기록

### 프론트엔드 작업
- 상태별 CTA
- 허용되지 않은 액션 비노출
- 현재 상태 / 상태 이력 표시
- 충돌/실패 응답 처리
- 거래 상세 화면 최소 구성

### 운영/관리 작업
- 이상 상태 거래 조회
- 수동 개입이 필요한 포인트 식별
- 분쟁 후보 상태 표시

### 테스트/검증
- 허용 전이
- 금지 전이
- 중복 요청
- 경쟁 조건
- 권한 없는 사용자 호출
- 취소 후 완료 같은 금지 시나리오
- 거래 완료 후 리뷰 가능 조건

### 완료 기준
- 거래 흐름이 서버에서 일관되게 통제된다.
- UI가 상태를 기준으로 동작한다.
- 대표 금지 전이 테스트가 통과한다.

---

## 14. Phase 4. 리뷰/신고/운영도구

### 목표
출시 가능한 신뢰/안전 장치를 만든다. MVP라도 신고 처리와 관리자 조회가 없으면 운영 리스크가 너무 크다.

### 포함 범위
- 거래 완료 후 리뷰
- 상품/사용자/채팅/거래 신고
- 최소 관리자 조회
- 관리자 상태 변경
- moderation action log

### 백엔드 작업
- review eligibility check
- review create/read 최소 기능
- 중복 리뷰 방지
- report create/status change
- target linkage 구조
- admin query / action API
- moderation audit log

### 프론트엔드 작업
- 리뷰 작성/조회 UI
- 신고 폼
- 관리자 목록/상세 최소 화면
- 처리 상태 표시

### 운영/관리 작업
- 신고 상태 코드 정의
- 처리 메모 필드
- 조치 기록 필드
- 우선순위 또는 최소 SLA placeholder

### 테스트/검증
- 거래 완료 전 리뷰 작성 차단
- 동일 거래 중복 리뷰 차단
- 신고 대상 존재성 검증
- 관리자 외 관리자 기능 접근 차단
- moderation log 누락 여부

### 완료 기준
- 사용자가 리뷰와 신고를 남길 수 있다.
- 운영자가 최소 조회와 상태 변경을 할 수 있다.
- 관련 조치 흔적이 남는다.

---

## 15. Phase 5. 안정화 / QA / 출시준비

### 목표
시연, 내부 배포, 제한적 초기 출시가 가능한 수준까지 품질을 끌어올린다.

### 포함 범위
- 회귀 테스트
- seed/demo data
- release checklist
- runbook
- health check
- 에러 처리 정리
- 기본 abuse guard
- 로그 필드 정리
- 모바일/반응형 최소 점검
- 관리자 계정 초기화 절차

### 백엔드 작업
- 공통 에러 처리 통일
- health / readiness 정리
- 로그 필드 표준화
- rate limit 또는 abuse placeholder
- integration test 정리
- migration from-scratch 재검증

### 프론트엔드 작업
- 로딩/오류/빈 상태 마감
- 폼 오류 메시지 정리
- 모바일 최소 대응
- 권한 오류 UX 정리

### 운영/관리 작업
- 관리자 계정 생성 절차
- release checklist
- 운영 runbook
- 사고 대응 메모
- 로그/감사 추적 조회 절차

### 테스트/검증
- 핵심 사용자 시나리오 E2E
- 권한 회귀
- 상태 전이 회귀
- 알림/이벤트 누락
- 데이터 정합성
- 신규 환경 기동

### 완료 기준
- 새 환경에서 기동 가능
- 핵심 흐름이 적어도 내부 데모 수준으로 안정적
- 알려진 리스크가 문서화됨
- 릴리스 blocker가 정리됨

---

## 16. Epic 기준 백로그 구조

Claude Code는 아래 구조로 Epic → Feature → Task를 쪼개야 한다.

### 16.1 인증/계정
목적:
- 사용자 식별과 권한 부여, 프로필 관리

주요 Feature:
- auth context
- 세션 또는 토큰 검증
- 프로필 조회/수정
- 계정 상태 처리

핵심 Task 예시:
- auth middleware 연결
- current user context helper
- profile endpoint
- account status field 반영
- 미인증 접근 차단 테스트

QA 체크:
- 미인증 접근
- role mismatch
- 정지/탈퇴 사용자 처리

### 16.2 상품 등록/수정/조회
목적:
- 거래 대상을 만들고 노출한다.

주요 Feature:
- create/update/detail/list
- ownership
- status exposure
- validation

핵심 Task 예시:
- product schema
- create/update handler
- owner-only update guard
- detail/list serializer
- status badge mapping

QA 체크:
- validation
- 권한 없는 수정
- soft delete / hide 정책

### 16.3 검색/탐색
목적:
- 발견성과 탐색 경험을 만든다.

주요 Feature:
- filter/sort/search
- pagination
- query normalization

핵심 Task 예시:
- list query parser
- pagination util
- filter mapping
- empty state UI

QA 체크:
- 조합 필터
- 빈 결과
- invalid query
- 성능 기초

### 16.4 채팅
목적:
- 거래 전 합의와 문의를 가능하게 한다.

주요 Feature:
- conversation create
- message send/read
- participant guard
- unread placeholder

핵심 Task 예시:
- conversation model
- duplicate prevention
- message pagination
- chat detail UI

QA 체크:
- 제3자 접근 차단
- 중복 방
- 긴 목록 pagination

### 16.5 거래/예약/완료
목적:
- 실제 거래 상태를 제어한다.

주요 Feature:
- transaction create
- state transition
- idempotency
- audit / event

핵심 Task 예시:
- state enum
- guard function
- transition API
- status history append
- audit record insertion

QA 체크:
- 금지 전이
- 중복 요청
- 동시성
- 권한 검증

### 16.6 리뷰
목적:
- 거래 품질 신호를 남긴다.

주요 Feature:
- eligibility check
- create/read
- duplicate prevention

핵심 Task 예시:
- review schema
- canReview helper
- review form UI
- aggregate exposure placeholder

QA 체크:
- 거래 미완료 리뷰 차단
- 중복 리뷰 차단

### 16.7 신고/분쟁
목적:
- 안전과 운영 대응을 가능하게 한다.

주요 Feature:
- report create
- target linkage
- admin status change
- moderation note

핵심 Task 예시:
- report target typing
- create API
- admin list/detail
- action log

QA 체크:
- 대상 무결성
- 관리자 권한
- 로그 누락

### 16.8 알림
목적:
- 핵심 상태 변화를 사용자에게 알린다.

주요 Feature:
- event hook
- notification dispatch placeholder
- unread or badge indicator placeholder

핵심 Task 예시:
- event constants
- emitter hook
- notification adapter interface
- UI indicator placeholder

QA 체크:
- 중복 발송
- 잘못된 수신자
- 누락

### 16.9 관리자/모더레이션
목적:
- 최소 운영 가능 상태를 만든다.

주요 Feature:
- admin auth
- list/detail/action
- audit visibility

핵심 Task 예시:
- admin route guard
- admin list query
- status change action
- decision memo field

QA 체크:
- 관리자 외 접근 차단
- action log 남음 여부

### 16.10 공통 인프라/관측성
목적:
- 디버깅, 운영, 안정화를 가능하게 한다.

주요 Feature:
- structured logging
- request id
- health check
- config registry
- demo seed

핵심 Task 예시:
- logger wrapper
- correlation id middleware
- health endpoint
- env example
- seed script

QA 체크:
- 민감정보 로그 노출 금지
- 기동 가능성
- runbook 반영

---

## 17. Claude Code 세션 운영 프로토콜

### 17.1 작업 시작 전
항상 아래를 짧게 남긴다.
- 이번 묶음 목표
- 수정 예정 파일/영역
- 선행 확인 포인트
- 예상 테스트 범위

### 17.2 작업 완료 후
항상 아래를 남긴다.
- 완료한 작업
- 수정한 파일
- 실행한 테스트
- 실패/보류 항목
- 남은 리스크
- 다음 묶음 제안

### 17.3 문서 업데이트 주기
아래는 작업 묶음 단위로 갱신한다.
- `BUILD_STATUS.md`: 매 작업 묶음 후
- `DECISION_LOG.md`: 결정 발생 시 즉시
- `SPEC_GAP_LOG.md`: 충돌/미확정 발견 시 즉시
- `TEST_RUN_REPORT.md`: 테스트 실행 후 즉시
- `MVP_EXECUTION_PLAN_v1.md`: phase 또는 우선순위 변동 시

### 17.4 중간 보고 예시
```text
이번 묶음 목표:
- 거래 상태 enum 및 transition guard 도입
- 거래 상태 변경 API 권한 검증 추가
- 대표 금지 전이 테스트 3건 추가

수정 예정 영역:
- backend transaction domain
- policy guard
- tests

검증 계획:
- lint
- unit test
- integration test 일부
```

작업 후:
```text
완료:
- transaction status enum 추가
- transition guard 연결
- 완료/취소 금지 전이 테스트 추가

수정 파일:
- ...
- ...

검증:
- lint 통과
- unit test 통과
- integration test 1건 실패 (원인 기록)

남은 리스크:
- 분쟁 진입 상태 코드 확인 필요
- notification event naming 미확정

다음 묶음:
- 거래 상세 응답에 상태 이력 노출
- 프론트 CTA 분기 반영
```

---

## 18. Definition of Ready / Definition of Done

### 18.1 작업 시작 전 Ready 기준
한 작업 묶음은 아래가 맞으면 시작한다.
- 목표가 한 문장으로 설명된다.
- 관련 endpoint 또는 화면이 식별된다.
- 수정 대상 파일 범위가 대략 보인다.
- 선행 정책 중 미확정이 있으면 기록되었다.
- 테스트 포인트가 최소 하나 이상 보인다.

### 18.2 작업 완료 Done 기준
아래가 없으면 Done이 아니다.
- 코드 변경이 실제 존재한다.
- 관련 문서가 갱신되었다.
- 최소 검증이 실행되었다.
- 남은 리스크가 명시되었다.
- 다음 작업 시작점이 적혔다.

### 18.3 Phase 완료 기준 공통
각 Phase는 아래를 만족해야 넘어간다.
- 대표 사용자 흐름이 끊기지 않는다.
- 대표 권한 실패 케이스가 막힌다.
- 상태/필드/권한의 큰 충돌이 정리된다.
- `BUILD_STATUS.md`에 완료/미완료가 분리 기록된다.

---

## 19. 품질 게이트

### 19.1 필수 기술 검증
- lint
- typecheck 또는 compile
- unit test
- integration test
- migration apply
- from-scratch boot 확인

### 19.2 도메인 검증
- 인증되지 않은 사용자 접근 차단
- 권한 없는 사용자 액션 차단
- 금지 상태 전이 차단
- soft delete / hidden resource 노출 정책 확인
- 신고/리뷰 대상 무결성 확인
- 관리자 기능 비관리자 접근 차단

### 19.3 운영 검증
- 로그 필드 일관성
- request id / actor id / target id 추적 가능성
- 관리자 조치 흔적
- config 기본값 문서화
- seed/demo data 준비

### 19.4 릴리스 blocker 예시
- 신규 환경에서 기동 불가
- migration 실패
- OpenAPI와 실제 응답이 상이
- 금지 상태 전이 허용
- 비관리자 관리자 페이지 접근 가능
- 채팅 또는 거래가 제3자에게 노출
- 신고/관리자 조치 audit trail 부재

---

## 20. 상태 문서 템플릿

아래 템플릿은 Claude Code가 실제로 만들어야 하는 최소 구조다.

### 20.1 `MVP_EXECUTION_PLAN_v1.md`
```text
# 린클 거래소 MVP 실행계획 v1

## 목표
## 현재 기준선
## Pre-build freeze 항목 현황
## Phase별 우선순위
## 이번 세션 / 이번 주 집중 항목
## 병렬 가능 작업
## 선행 확정 필요 작업
## 주요 리스크
## 다음 체크포인트
```

### 20.2 `BUILD_STATUS.md`
```text
# BUILD STATUS

## 현재 기준선 파악
## 완료된 항목
## 진행 중 항목
## 아직 시작 안 한 항목
## 이번 세션 수정 파일
## 테스트 실행 결과 요약
## 현재 blocker / 리스크
## 다음 세션 시작점
```

### 20.3 `DECISION_LOG.md`
```text
# DECISION LOG

- [결정 ID]
  - 날짜:
  - 주제:
  - 결정 내용:
  - 근거:
  - 영향 범위:
  - 관련 파일:
  - 후속 작업:
```

### 20.4 `SPEC_GAP_LOG.md`
```text
# SPEC GAP LOG

- [GAP ID]
  - 발견 날짜:
  - 항목:
  - 문서 내용:
  - 코드 현실:
  - 영향:
  - 임시 처리:
  - 확인 필요:
```

### 20.5 `TEST_RUN_REPORT.md`
```text
# TEST RUN REPORT

## 실행한 명령
## 통과
## 실패
## 스킵
## 원인 분석
## 다음 조치
```

### 20.6 `RELEASE_BLOCKERS.md`
```text
# RELEASE BLOCKERS

- [BLOCKER ID]
  - 증상:
  - 영향:
  - 재현 방법:
  - 우회 가능 여부:
  - 담당:
  - 해결 조건:
```

---

## 21. 권장 폴더 구조

가능하면 아래처럼 정리한다.

```text
docs/
  lincle/
    PRD.md
    TECH_STACK_ARCHITECTURE.md
    OPENAPI_DRAFT.md
    STARTER_DDL.md
    RBAC_ACTION_MATRIX.md
    RETENTION_CONFIG_APPENDIX.md
    EVENT_CATALOG.md
    STATE_SEQUENCE_DIAGRAMS.md
    CLAUDE_CODE_EXECUTION_MASTER_SPEC_v2.md
    MVP_EXECUTION_PLAN_v1.md
    BUILD_STATUS.md
    DECISION_LOG.md
    SPEC_GAP_LOG.md
    TEST_RUN_REPORT.md
    API_CHANGELOG.md
    RELEASE_CHECKLIST.md
    RELEASE_BLOCKERS.md
```

---

## 22. 4주 기준 권장 로드맵

이 일정은 실제 인원과 저장소 상태에 따라 조정 가능하지만, Claude Code가 어떤 순서로 repo를 전진시켜야 하는지 기준선을 제공한다.

### Week 1
목표:
- Phase 0 완료에 최대한 가깝게 접근
- 상품 등록/목록/상세의 첫 vertical slice 시작

핵심 산출물:
- OpenAPI 초안 정리
- migration 1차
- 상태 enum / RBAC constants
- 상품 create/list/detail 최소 구현
- 상태 문서 초기화

주간 의사결정 포인트:
- 상태머신 정의 초안
- action code freeze
- 핵심 엔티티 필드명 확정

### Week 2
목표:
- 상품/탐색 slice를 마감하고 채팅 도입
- 거래 상태머신의 첫 서버 구현 시작

핵심 산출물:
- 상품 수정/권한 체크
- 검색/필터 최소 기능
- conversation/message 기본 흐름
- transaction skeleton
- 대표 권한 테스트

주간 의사결정 포인트:
- 채팅 중복 방 규칙
- 거래 생성 시점
- cancellation policy placeholder

### Week 3
목표:
- 거래 흐름을 서버 주도 상태머신으로 연결
- 리뷰/신고/관리자 최소 도구 시작

핵심 산출물:
- 상태 변경 API
- 금지 전이 테스트
- 리뷰 eligibility
- 신고 생성
- 관리자 조회/상태 변경 최소 기능

주간 의사결정 포인트:
- 분쟁 상태 placeholder 범위
- moderation status code
- audit log 필수 필드

### Week 4
목표:
- 회귀, seed, runbook, release checklist 정리
- 내부 데모 또는 제한 배포 가능 수준까지 안정화

핵심 산출물:
- 핵심 시나리오 E2E 또는 통합 검증
- seed/demo data
- release blockers 정리
- 운영 runbook
- observability / log / audit 점검

주간 의사결정 포인트:
- 출시 제외 범위 확정
- known issues 정리
- 초기 운영 대응 흐름 확인

---

## 23. 첫 세션에서 바로 해야 할 12개 액션

1. `docs/lincle` 실행 문서들의 존재 여부를 확인한다.
2. repo 구조와 실행 명령을 스캔한다.
3. OpenAPI 관련 파일과 실제 API handler 차이를 찾는다.
4. migration 체계와 현재 엔티티를 매핑한다.
5. 거래 상태를 표현하는 enum/상수/문서가 이미 있는지 찾는다.
6. 권한 처리 코드와 관리자 라우트를 찾는다.
7. `MVP_EXECUTION_PLAN_v1.md` 초안을 작성한다.
8. `BUILD_STATUS.md`에 현재 기준선을 기록한다.
9. `SPEC_GAP_LOG.md`에 가장 큰 충돌 3~5개를 먼저 적는다.
10. Phase 0의 가장 작은 구현 묶음 1개를 바로 선택한다.
11. 코드 수정 후 lint/test를 실행한다.
12. 테스트 결과와 다음 시작점을 문서에 남긴다.

---

## 24. 전형적인 실패 패턴과 교정 원칙

### 실패 패턴 1. 문서 요약만 길고 코드가 없음
교정:
- 첫 응답 다음 바로 파일 수정
- 상태 문서는 구현 흔적을 보조하는 수준으로 유지

### 실패 패턴 2. 리팩터링을 너무 크게 벌림
교정:
- 한 패치는 한 목적
- MVP와 무관한 구조개선은 backlog로 내린다

### 실패 패턴 3. 서버 권한 검증 없이 UI만 만듦
교정:
- UI 작업 전후로 guard/policy를 같이 구현
- 권한 실패 테스트를 항상 한 건 이상 붙인다

### 실패 패턴 4. 상태머신 없이 if문으로 누더기 처리
교정:
- 상태 enum과 transition map부터 만든다
- 금지 전이 테스트를 대표 시나리오로 둔다

### 실패 패턴 5. 문서와 코드가 계속 어긋남
교정:
- API/DB/상태/RBAC 변경 시 문서를 즉시 갱신
- changelog와 gap log를 병행

### 실패 패턴 6. 테스트를 안 돌리고 완료 선언
교정:
- 최소 lint + 타입검사 + 대표 테스트 1개 이상
- 실패면 실패 원인을 기록하고 다음 액션을 적는다

### 실패 패턴 7. blocker 하나로 전체 중단
교정:
- blocker 범위를 분리
- 독립 작업으로 계속 전진
- gap과 blocker를 문서로 남긴다

---

## 25. Claude Code에 바로 넣는 상세 시작 프롬프트

아래 프롬프트는 이 문서의 요약본이 아니라, 실제 세션 시작용 실행 지시문이다.

```text
너는 이 저장소의 테크 리드이자 hands-on 구현 담당자다.

내가 제공하는 문서는 “린클 거래소”의 PRD와 기술 설계 문서이며,
너의 목표는 이 문서들을 바탕으로 현재 repo에서 MVP 개발을 실제로 전진시키는 것이다.

중요:
- PRD를 다시 길게 요약하지 마라.
- 기술 스택을 새로 갈아엎는 제안을 하지 마라.
- 문서에 없는 정책을 과도하게 상상하지 마라.
- 대신, 저장소를 스캔하고, 실행계획을 세우고, 바로 파일 수정과 테스트, 상태 업데이트까지 진행하라.
- 구현 가능한 것은 승인 대기 없이 진행하고, 미확정 고위험 정책만 “확인 필요”로 남겨라.
- 목표는 설명이 아니라 repo를 실제로 전진시키는 것이다.

반드시 먼저 읽을 문서:
1. docs/lincle/PRD.md
2. docs/lincle/TECH_STACK_ARCHITECTURE.md

가능하면 함께 읽을 문서:
3. docs/lincle/OPENAPI_DRAFT.md
4. docs/lincle/STARTER_DDL.md
5. docs/lincle/RBAC_ACTION_MATRIX.md
6. docs/lincle/RETENTION_CONFIG_APPENDIX.md
7. docs/lincle/EVENT_CATALOG.md
8. docs/lincle/STATE_SEQUENCE_DIAGRAMS.md

너의 절대 규칙:
- 계획 제시로 끝내지 말 것
- 첫 응답 뒤 바로 코드/문서 수정에 들어갈 것
- 한 번에 하나의 목적을 가진 작은 작업 묶음으로 진행할 것
- API/DB/상태/RBAC/config 변경 시 관련 문서를 함께 갱신할 것
- 테스트 없이 완료 선언하지 말 것
- 문서와 코드 충돌 시 SPEC_GAP_LOG에 남기고 MVP에 더 안전한 방향으로 계속 진행할 것
- blocker 하나 때문에 전체 작업을 멈추지 말 것

세션 시작 직후 해야 할 일:
1. 저장소 구조와 실행 명령을 스캔하라.
2. 현재 구현 수준과 입력 문서 대응 상태를 요약하라.
3. 아래 문서를 생성 또는 갱신하라.
   - docs/lincle/MVP_EXECUTION_PLAN_v1.md
   - docs/lincle/BUILD_STATUS.md
   - docs/lincle/DECISION_LOG.md
   - docs/lincle/SPEC_GAP_LOG.md
   - docs/lincle/TEST_RUN_REPORT.md
4. Phase 0(계약/기반 확정)부터 실제 구현에 착수하라.

Phase 0에서 우선 정리할 것:
- OpenAPI 초안
- DB migration 초안
- 상태 enum / transition
- RBAC/action code
- config key / retention 기본값
- event/audit field
- 공통 에러 응답 및 pagination 기준

그 다음 구현 우선순위:
- Phase 1: 계정/상품/탐색
- Phase 2: 1:1 채팅
- Phase 3: 거래 상태머신
- Phase 4: 리뷰/신고/운영도구
- Phase 5: 안정화/QA/출시준비

각 작업 묶음마다 반드시 수행:
- 관련 코드 수정
- 테스트 추가 또는 갱신
- 상태 문서 갱신
- 수정 파일 목록 기록
- 남은 리스크와 다음 시작점 기록

세션 첫 응답 형식:
1. 현재 프로젝트 상태 요약
2. 지금 바로 구현 가능한 범위
3. 선행 확정이 필요한 항목
4. 이번 세션 우선 작업 목록
5. 바로 착수할 첫 구현 묶음
6. 생성/갱신할 상태 문서 목록

중요:
- 첫 응답 뒤에는 실제 파일 수정에 바로 들어가라.
- 법무/정산/본인인증 같은 고위험 정책은 문서가 없으면 확정하지 말고 확인 필요로 남겨라.
- 하지만 구현 가능한 범위는 멈추지 말고 계속 전진시켜라.
- 추천만 하고 끝내는 것은 실패다.
- 오늘 세션 안에서 문서 생성/갱신, 실제 코드 수정, 테스트 실행, 상태 업데이트까지 반드시 남겨라.
```

---

## 26. 최종 체크리스트

Claude Code가 이 문서를 제대로 따른다면 아래 결과가 남아야 한다.

- PRD 재요약이 아니라 실행계획과 코드 수정이 나온다.
- Phase 0의 계약 계층이 먼저 정리된다.
- 상품/탐색/채팅/거래/리뷰/신고/운영도구가 단계적으로 전진한다.
- 상태 문서가 남아 다음 세션 이어받기가 쉽다.
- 미확정 이슈와 blocker가 숨겨지지 않는다.
- 테스트와 검증 흔적이 남는다.
- “무엇을 바꿨고, 왜 바꿨고, 다음에 무엇을 해야 하는지”가 한 번에 보인다.

문서의 성공 기준은 단순하다.

> **Claude Code가 더 이상 계획서 작성 도구가 아니라, 린클 거래소 MVP를 실제로 조립하고 검증하는 실행 엔진처럼 움직이면 성공이다.**