# Ralph Dual-Loop: Bug-Fix + Improvement-Suggest for giranjt.com

**Date**: 2026-05-02
**Target**: https://giranjt.com/ (production)
**Repo**: lincle (이 repo)

## 목적

giranjt.com에 대해 두 개의 자동화된 발견·수정·제안 사이클을 단일 Ralph Loop 안에서 병렬로 돌린다.

- **Loop 1 (bug-detect-fix)**: 객관적 결정성 기준의 버그를 헤드리스 브라우저로 찾아 자동 수정·리뷰·머지한다. 배포는 사람이 한다.
- **Loop 2 (improve-suggest)**: 주관적 개선점을 구조화된 백로그로 누적한다. 코드 수정은 하지 않는다.

## 상위 결정 (확정)

| # | 결정 | 사유 |
|---|------|------|
| 1 | 단일 Ralph Loop 세션 + 매 iteration 내부에서 병렬 subagent 2개 | Ralph Loop는 단일 세션 도구. State/git 충돌 회피, dedup 일관성 |
| 2 | Loop 1 자동화 Level 2 + 자동 리뷰·머지 (배포는 사람) | "수정해줘" 요청 충족. 잘못된 fix는 main에 들어가기 전 리뷰 게이트가 차단 |
| 3 | Bug 검사 = Scope B (객관적 측정만) | "3회 연속 0건" 종료 조건이 결정적이어야 의미 있음 |
| 4 | Improve = Option I (수집만, 구현 없음) | 개선은 본질적으로 우선순위·전략 판단 필요. 자동 구현은 노이즈 |

## 아키텍처

```
┌──────────────────────────────────────────────────────────────┐
│  /ralph-loop "<MAIN>" --max-iterations 1000 \                │
│              --completion-promise "DUAL_LOOP_COMPLETE"       │
│                                                              │
│  매 iteration:                                               │
│    1. tasks/*.json, tasks/*.md 상태 로드                     │
│    2. 두 작업을 병렬 subagent로 dispatch                     │
│         ┌─────────────────────┐  ┌──────────────────────┐    │
│         │ bug-detect-fix      │∥ │ improve-suggest      │    │
│         │ (general-purpose)   │  │ (general-purpose)    │    │
│         └─────────────────────┘  └──────────────────────┘    │
│    3. 결과 병합 + state 갱신 (메인이 직접 파일 수정)         │
│    4. 종료조건 충족시 <promise>DUAL_LOOP_COMPLETE</promise>  │
└──────────────────────────────────────────────────────────────┘
```

## 상태 파일 구조

```
tasks/
  ralph-dual.md           # 메인 프롬프트 (Ralph가 매 iter 다시 읽음)
  bug-state.json          # Loop 1 카운터/체크포인트
  bug-found.md            # 미해결 버그 목록 (key=page+rule+selector)
  bug-fixed.md            # 머지 완료된 버그 로그
  needs-human.md          # 자동수정 또는 리뷰 실패 → 사람 손 필요
  improve-state.json      # Loop 2 카운터/체크포인트
  improvements.md         # 제안 백로그 (YAML 블록 누적)
  ralph-progress.md       # iter별 한 줄 진행 로그
```

### bug-state.json 스키마

```json
{
  "iteration": 42,
  "consecutive_no_new": 1,
  "last_page_index": 3,
  "page_queue": ["/", "/listings", "/listings/:id", "/login", "/profile", "/chat"],
  "pages_visited_count": 6,
  "totals": {
    "found": 18,
    "fixed": 14,
    "blocked": 4
  }
}
```

`page_queue`의 `/listings/:id`는 placeholder. 실제 detect 시점에 `/listings`에서 첫 번째 매물 id를 추출해서 동적으로 치환한다 (placeholder 자체를 fetch하지 않는다).

### improve-state.json 스키마

```json
{
  "iteration": 42,
  "consecutive_no_new": 0,
  "last_area_index": 5,
  "area_queue": ["home", "listings", "listing_detail", "listing_create",
                 "chat", "reservation", "review", "report",
                 "notification", "profile", "auth"],
  "areas_visited_count": 7,
  "total_unique_suggestions": 67
}
```

### bug-found.md 항목 스키마

```yaml
- id: bug-0007
  found_at_iter: 12
  page: /listings
  viewport: 375x667
  rule_id: color-contrast            # axe-core rule ID 또는 A1-A5/B1-B5 ID
  element_selector: "button.filter-toggle"
  evidence: "contrast 3.2:1 (요구 4.5:1)"
  status: open|fixing|fixed|blocked
  branch: auto/bug-12-0007
  fixed_at_iter: 13                  # 적용시
  blocked_reason: "DB schema 변경 필요" # blocked시
```

### improvements.md 항목 스키마

```yaml
- id: imp-0042
  found_at_iter: 18
  area: chat
  type: ux                           # ux | feature | performance | content | a11y_mobile
  target: send_button                # 영역 내 구체 요소/플로우
  problem: "메시지 전송 후 textarea focus가 풀려 연속 채팅이 불편"
  proposal: "전송 후 focus() 유지 + 스크롤 하단 자동 이동"
  effort: trivial                    # trivial | small | medium | large
  impact: medium                     # low | medium | high
  evidence: "Playwright trace: focus가 body로 이동 (iter 18)"
```

## Pre-flight Setup (Ralph 시작 전 1회)

루프가 매 iteration에서 의존하는 외부 상태:

1. **로컬 web dev server**: `http://localhost:3000`에서 fix 검증용 재실행
   - 시작: `cd web && npm run dev` (백그라운드)
   - 헬스체크: `curl -sf http://localhost:3000` 200 확인
   - 죽으면 자동 재시작 (메인 오케스트레이터가 매 iter 헬스체크)
2. **Playwright + axe-core 의존성 확인**: `cd web && npx playwright --version`, `cd web && npx axe --version`
3. **초기 state 파일 생성**: 위 8개 파일이 없으면 빈 상태로 생성
4. **샘플 매물 id 캐시**: `/listings`에서 첫 매물 id를 한 번 추출하여 `bug-state.json.sample_listing_id`에 저장 (매 iter 갱신)
5. **Git 상태 확인**: working tree clean, current branch == main

위 5단계 중 하나라도 실패하면 루프 시작하지 않고 사용자에게 보고한다.

### Playwright 도구 선택

- **헤드리스 브라우저 자동화**: `mcp__plugin_playwright_playwright__*` MCP 도구 사용
  - subagent가 직접 호출 (`browser_navigate`, `browser_evaluate`, `browser_take_screenshot`)
  - axe-core는 `browser_evaluate`로 페이지에 주입 후 결과 수신
- **로컬 검증**: 같은 MCP 도구로 `http://localhost:3000` 시나리오 재실행
- **이유**: 새 프로세스/브라우저 세션 띄우는 비용 없음. iter간 page state 격리.

## Loop 1 — bug-detect-fix subagent

### 검사 기준

**A군 — 기능적 결함 (Playwright runtime)**

| ID | 기준 | 판정 |
|----|------|------|
| A-1 | JS 콘솔 error / 잡히지 않은 예외 | `pageerror` 이벤트, `console.error` 발생 |
| A-2 | 네트워크 4xx/5xx | response.status() ≥ 400 (의도된 401 보호 엔드포인트 제외) |
| A-3 | 깨진 이미지 | `<img>` 로드 후 `naturalWidth === 0` |
| A-4 | 빈/깨진 페이지 | `<body>` 텍스트 < 50자 또는 title 비어있음 |
| A-5 | 죽은 버튼/링크 | 클릭 후 2s 내 DOM 변화·네트워크 요청·네비게이션 모두 없음 |

**B군 — 시각/UX 결함 (객관적 측정)**

| ID | 기준 | 판정 |
|----|------|------|
| B-1 | 가로 오버플로 | `documentElement.scrollWidth > clientWidth + 5px` |
| B-2 | 텍스트 잘림 (인디케이터 없음) | `scrollWidth>clientWidth+2` AND `overflow≠visible` AND no `text-overflow:ellipsis` |
| B-3 | 모바일 탭 타겟 미달 | 375px 뷰포트, 인터랙티브 요소 BoundingRect < 32×32 px |
| B-4 | 색 대비 위반 (WCAG AA) | axe-core `color-contrast` rule violation |
| B-5 | 라벨/이름 누락 | axe-core `label`, `button-name`, `link-name`, `image-alt` violations |

### 라운드로빈 페이지 큐

```
/ → /listings → /listings/<sample-id> → /login → /profile → /chat → (반복)
```

- 각 iteration마다 1 페이지를 3 뷰포트(375 / 768 / 1280)에서 검사
- `<sample-id>`는 매 iter `/listings`에서 첫 번째 매물 id를 추출

### Dedup 키

```
(page_path, rule_id, normalized_selector)
```

- `normalized_selector`: nth-child 인덱스 제거, data-testid 우선
- 같은 키로 이미 `bug-found.md`에 있으면 새 발견 카운트하지 않음

### 단계별 동작

```
1. Detect:
   • Playwright로 페이지 로드 (production URL)
   • A1-A5 체크 + axe-core 주입 후 B1-B5 위반 수집
2. Dedup:
   • bug-found.md 로딩 → (page,rule,selector) 키 비교
   • 새 버그만 처리 대상
3. Localize:
   • selector → DOM 컨텍스트 (textContent, aria-label, data-testid) 추출
   • Grep으로 web/, frontend/ 소스에서 매칭되는 텍스트/속성 검색
   • 매칭 0건 또는 5건 이상이면 needs-human.md로 라우팅 (불확실)
4. Branch:
   • git checkout -b auto/bug-<iter>-<bug-id>
5. Fix:
   • 코드 수정 (Edit 도구)
6. Verify Tier 1+2+3:
   • lint (web: eslint, backend: golangci-lint, frontend: flutter analyze)
   • typecheck (web: tsc --noEmit, backend: go vet)
   • test (web: vitest run, backend: go test ./..., frontend: flutter test)
7. Re-detect on local:
   • Playwright로 http://localhost:3000 에서 같은 시나리오 재실행
   • 같은 (rule_id, selector) violation 사라짐 확인
8. Review:
   • pr-review-toolkit:code-reviewer subagent에 diff 전달
   • 승인/거부 판정 수신
9. Merge or Discard:
   • 승인 → main으로 squash merge, bug-fixed.md 추가
   • 거부 → 브랜치 폐기, needs-human.md 기록
```

**Step 3 Localize 실패 시**: 코드 위치를 결정적으로 매핑할 수 없으면 자동 fix를 시도하지 않고 `needs-human.md`에 다음 정보를 적재:
- bug id, page, rule, selector, evidence
- DOM 컨텍스트 (HTML 스니펫)
- Grep 매칭 후보 목록 (있으면)

### Step 7 (Re-detect on local) 상세

- Pre-flight에서 띄운 `http://localhost:3000` web dev server 사용
- backend가 필요한 fix는 backend dev server도 띄워둠 (선택, 대부분 web 단독)
- Playwright MCP로 production과 동일한 시나리오를 `http://localhost:3000`에 재실행
- 버그가 사라졌으면 fix 검증 통과
- dev server가 죽어있으면 (헬스체크 실패) 자동 재시작 후 1회 재시도, 그래도 실패면 needs-human

### Verification Gate (머지 전 must-pass)

```
□ Tier 1 lint pass
□ Tier 2 typecheck pass
□ Tier 3 test pass
□ Local re-detect: 같은 룰+셀렉터 violation 없음
□ Code review subagent: APPROVED
```

하나라도 실패 → 브랜치 폐기 + `needs-human.md`에 기록

## Loop 2 — improve-suggest subagent

### 영역 라운드로빈 큐

```
home → listings → listing_detail → listing_create → chat
     → reservation → review → report → notification → profile → auth → (반복)
```

11개 영역. 매 iteration 1개 영역에 집중.

### 5관점 생성

각 영역에서 다음 5관점으로 제안 생성:

| Type | 관점 |
|------|------|
| `ux` | 플로우 단순화, 피드백 부족, 마찰점 |
| `feature` | 현재 없는 기능, 추정 사용자 니즈 |
| `performance` | 체감 속도, 로딩 인디케이터 |
| `content` | 한국어 자연스러움, 빈 상태 메시지 |
| `a11y_mobile` | 탭 순서, 한손 사용성, 키보드 접근성 |

### Dedup 키

```
(area, type, target)
```

같은 트리플의 새 제안은 카운트하지 않음.

### 동작

```
1. Pick area (라운드로빈)
2. Explore via Playwright:
   • 다양한 시나리오 시도 (정상/빈/긴 입력, 모바일 뷰포트)
3. Generate suggestions (5관점)
4. Dedup → 새 제안만 추출
5. Append to improvements.md (YAML 블록)
```

**제약**: improve-suggest subagent는 `tasks/improvements.md` 외 파일 수정 금지.

## 종료 조건

| 조건 | 정의 |
|------|------|
| Bug loop done | `bug-state.consecutive_no_new ≥ 3` AND `pages_visited_count == 6` (모든 페이지 1회 이상 방문) |
| Improve loop done | `improve-state.consecutive_no_new ≥ 3` AND `areas_visited_count == 11` (모든 영역 1회 이상 방문) |
| Both done | 위 둘 모두 충족 → `<promise>DUAL_LOOP_COMPLETE</promise>` |
| Hard cap | iteration 1000 도달 시 강제 종료 |

전체 페이지/영역을 한 바퀴 돌기 전에는 종료하지 않는다 (한 페이지에서만 0건이라도 다른 페이지에 버그가 있을 수 있음). 페이지 6개·영역 11개 = 최소 11 iter 이상 보장.

## 안전장치

| 위험 | 완화 |
|------|------|
| 자동 fix가 잘못된 코드 생성 | 독립 review subagent + Tier 1/2/3 게이트 + local re-detect |
| Dedup 실패 (같은 버그가 다른 표현으로 신규로 잡힘) | `(page, rule_id, normalized_selector)` 정규화 키 |
| 무한 진동 (fix가 새 버그 유발) | iter별 변경 파일 추적, 같은 파일 3회 이상 수정 시 자동 차단 → needs-human |
| Playwright 비용 폭발 | 매 iter 1페이지×3뷰포트로 제한 |
| 운영 사이트 부하 | production URL은 읽기만 / fix 검증은 로컬 빌드 |
| 두 subagent 파일 충돌 | improve-suggest는 `tasks/improvements.md` 외 쓰기 금지 |
| 자동 머지 폭주 | main에 squash merge 1건/버그, deploy는 사람 |
| 잘못된 페이지 큐 (sample-id가 invalid) | 매 iter `/listings`에서 첫 매물 id를 동적 추출 |

## 동시성 / 파일 책임 분리

| 파일 | 쓰기 권한 |
|------|-----------|
| `tasks/bug-state.json` | bug-detect-fix only |
| `tasks/bug-found.md` | bug-detect-fix only |
| `tasks/bug-fixed.md` | bug-detect-fix only |
| `tasks/needs-human.md` | bug-detect-fix only |
| `tasks/improve-state.json` | improve-suggest only |
| `tasks/improvements.md` | improve-suggest only |
| `tasks/ralph-progress.md` | 메인 (병합 후 한 줄 추가) |
| `backend/`, `web/`, `frontend/`, `admin/` | bug-detect-fix only |

## 범위 외 (Out of Scope)

- ❌ 자동 배포 (사용자가 `deploy/deploy.sh` 실행)
- ❌ DB 스키마 변경 수반 fix → `needs-human.md`로 라우팅
- ❌ E2E 테스트 자동 추가/수정 (검증에만 사용)
- ❌ Loop 2의 자동 구현 (백로그 생성만)
- ❌ 두 번째 Claude Code 세션 운영 (단일 세션 안에서 모든 처리)

## 결과물

- 머지된 자동 fix 커밋들 (squash, main에 누적)
- `tasks/bug-fixed.md`: 자동 처리된 버그 목록
- `tasks/needs-human.md`: 사람 손이 필요한 항목 큐
- `tasks/improvements.md`: 우선순위 정렬 가능한 개선 백로그
- `tasks/ralph-progress.md`: iteration별 진행 로그

루프 종료 후 사용자가 할 일:
1. `tasks/needs-human.md` 처리
2. `tasks/improvements.md`에서 implement할 항목 골라 `/feature-dev` 등으로 처리
3. `deploy/deploy.sh`로 누적된 자동 fix 배포

## 다음 단계

이 spec 승인 → `superpowers:writing-plans` 스킬로 단계별 실행 plan 작성 → Ralph Loop 시작.
