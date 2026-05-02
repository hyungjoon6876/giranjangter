# Ralph Dual-Loop Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ralph Loop 1개를 띄워, 매 iteration에서 두 개의 병렬 subagent (bug-detect-fix, improve-suggest)를 dispatch하여 giranjt.com에 대한 자동 버그 수정과 개선 제안 백로그 생성을 동시에 수행한다.

**Architecture:** 단일 Claude Code 세션 + `/ralph-loop` + 매 iter마다 `Agent` 도구로 두 subagent 병렬 dispatch. 상태는 `tasks/*.json`/`*.md` 파일로 영속화. 자동 fix는 `auto/bug-*` 브랜치 → review subagent 게이트 → main squash 머지. 배포는 사람.

**Tech Stack:** Ralph Loop (`ralph-loop:ralph-loop`), Playwright MCP (`mcp__plugin_playwright_playwright__*`), axe-core, code-reviewer subagent (`pr-review-toolkit:code-reviewer`), git, bash.

**Spec:** `docs/superpowers/specs/2026-05-02-ralph-dual-loop-design.md`

---

## File Structure

생성/수정될 파일:

```
tasks/                                  # 새로 생성 (gitignore 안 됨, 추적함)
├── ralph-dual.md                       # Ralph가 매 iter 다시 읽는 메인 프롬프트
├── bug-state.json                      # Loop 1 카운터/체크포인트
├── bug-found.md                        # 미해결 버그 누적
├── bug-fixed.md                        # 자동 머지 완료 로그
├── needs-human.md                      # 사람 손 필요 큐
├── improve-state.json                  # Loop 2 카운터/체크포인트
├── improvements.md                     # 개선 제안 백로그
└── ralph-progress.md                   # iter별 한 줄 로그

scripts/
└── ralph-preflight.sh                  # 시작 전 환경 점검

docs/superpowers/specs/2026-05-02-ralph-dual-loop-design.md   # 이미 존재
docs/superpowers/plans/2026-05-02-ralph-dual-loop-plan.md     # 이 파일
```

각 파일의 책임:
- `ralph-dual.md`: 단일 진실 원천(SOT). 매 iter 메인 컨텍스트가 이 파일을 읽고 그대로 실행
- `*-state.json`: 결정적 카운터·인덱스 (수동 편집 가능, 외부 분석 가능한 JSON)
- `*-found.md`/`*-fixed.md`/`improvements.md`/`needs-human.md`: 사람이 읽는 누적 로그 (YAML 블록)
- `ralph-progress.md`: 사람이 읽는 진행 상황 추적용
- `ralph-preflight.sh`: 루프 시작 전 환경 검증 (실패시 루프 시작 금지)

---

## Task 1: tasks 디렉토리와 초기 상태 파일 생성

**Files:**
- Create: `tasks/bug-state.json`
- Create: `tasks/improve-state.json`
- Create: `tasks/bug-found.md`
- Create: `tasks/bug-fixed.md`
- Create: `tasks/needs-human.md`
- Create: `tasks/improvements.md`
- Create: `tasks/ralph-progress.md`

- [ ] **Step 1: 디렉토리 생성**

```bash
mkdir -p tasks
```

- [ ] **Step 2: bug-state.json 초기 상태 작성**

`tasks/bug-state.json`:

```json
{
  "iteration": 0,
  "consecutive_no_new": 0,
  "last_page_index": -1,
  "page_queue": ["/", "/listings", "/listings/:id", "/login", "/profile", "/chat"],
  "pages_visited": [],
  "pages_visited_count": 0,
  "sample_listing_id": null,
  "files_modified_count": {},
  "totals": {
    "found": 0,
    "fixed": 0,
    "blocked": 0
  }
}
```

설명:
- `last_page_index = -1` → 다음 iter는 인덱스 0(`/`)부터 시작
- `pages_visited` 배열로 어떤 페이지를 방문했는지 추적, `pages_visited_count`는 unique count
- `sample_listing_id`는 매 iter `/listings` 방문 시 동적 갱신 (placeholder 치환용)
- `files_modified_count`는 같은 파일 3회 이상 수정 차단을 위한 카운터 (key=경로, value=count)

- [ ] **Step 3: improve-state.json 초기 상태 작성**

`tasks/improve-state.json`:

```json
{
  "iteration": 0,
  "consecutive_no_new": 0,
  "last_area_index": -1,
  "area_queue": ["home", "listings", "listing_detail", "listing_create", "chat", "reservation", "review", "report", "notification", "profile", "auth"],
  "areas_visited": [],
  "areas_visited_count": 0,
  "total_unique_suggestions": 0
}
```

- [ ] **Step 4: 빈 누적 로그 파일들 작성**

`tasks/bug-found.md`:

```markdown
# Bug Found Log

> Loop 1이 발견했지만 아직 처리 중이거나 처리 대기인 버그. 형식: YAML 블록.
> Dedup key: (page, rule_id, normalized_selector)

```

`tasks/bug-fixed.md`:

```markdown
# Bug Fixed Log

> Loop 1이 자동으로 발견 → 수정 → 리뷰 → main 머지 완료한 버그.

```

`tasks/needs-human.md`:

```markdown
# Needs Human Queue

> 자동 처리 실패 또는 안전상 자동 수정하지 않은 항목. 사용자가 직접 처리한다.

```

`tasks/improvements.md`:

```markdown
# Improvement Suggestions Backlog

> Loop 2가 수집한 개선 제안. 코드 수정 안 됨. 우선순위 정렬 후 별도로 implement한다.
> Dedup key: (area, type, target)

```

`tasks/ralph-progress.md`:

```markdown
# Ralph Dual-Loop Progress

> iter당 한 줄 로그. 형식: `iter <N> | bugs +<n>/-<m>=<open> | improvs +<k>=<total> | <event>`
> 예: `iter 12 | bugs +2/-1=3 | improvs +1=18 | merged auto/bug-12-0007`

```

- [ ] **Step 5: 커밋**

```bash
git add tasks/
git commit -m "$(cat <<'EOF'
feat(ralph): scaffold dual-loop state files

8 state/log files to drive the Ralph dual-loop orchestration.
- *-state.json: 결정적 카운터
- *-found/fixed.md, needs-human.md, improvements.md: 누적 로그
- ralph-progress.md: iter별 진행 추적

EOF
)"
```

---

## Task 2: Pre-flight 스크립트 작성

**Files:**
- Create: `scripts/ralph-preflight.sh`

- [ ] **Step 1: 스크립트 작성**

`scripts/ralph-preflight.sh`:

```bash
#!/bin/bash
# Ralph Dual-Loop preflight check
# Usage: bash scripts/ralph-preflight.sh
# Exit 0 = OK to start. Non-zero = fix the issue first.

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$REPO_ROOT"

fail() { echo "FAIL: $1" >&2; exit 1; }
ok()   { echo "OK:   $1"; }

# 1. Git: clean working tree + on main
if [ -n "$(git status --porcelain)" ]; then
  fail "working tree not clean. commit/stash first."
fi
ok "git clean"

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$CURRENT_BRANCH" != "main" ]; then
  fail "not on main (current: $CURRENT_BRANCH). switch with: git checkout main"
fi
ok "on main"

# 2. State files exist
for f in \
  tasks/bug-state.json \
  tasks/improve-state.json \
  tasks/bug-found.md \
  tasks/bug-fixed.md \
  tasks/needs-human.md \
  tasks/improvements.md \
  tasks/ralph-progress.md \
  tasks/ralph-dual.md; do
  if [ ! -f "$f" ]; then
    fail "missing state file: $f"
  fi
done
ok "state files present"

# 3. State files are valid JSON
for j in tasks/bug-state.json tasks/improve-state.json; do
  if ! python3 -c "import json; json.load(open('$j'))" 2>/dev/null; then
    fail "invalid JSON: $j"
  fi
done
ok "state JSON valid"

# 4. Web dev server reachable on :3000
if ! curl -sf -m 3 http://localhost:3000 >/dev/null 2>&1; then
  fail "web dev server not running on :3000. start in another terminal: cd web && npm run dev"
fi
ok "web dev server up"

# 5. Production reachable
if ! curl -sf -m 5 https://giranjt.com/ >/dev/null 2>&1; then
  fail "https://giranjt.com/ not reachable"
fi
ok "production reachable"

# 6. Playwright present in web/
if ! (cd web && [ -f node_modules/.bin/playwright ]); then
  fail "playwright not installed in web/. run: cd web && npm install"
fi
ok "playwright installed"

echo "=== Preflight passed ==="
```

- [ ] **Step 2: 실행 권한 부여 + 파일 시스템 확인**

```bash
chmod +x scripts/ralph-preflight.sh
ls -l scripts/ralph-preflight.sh
```
Expected: 실행 권한(`-rwxr-xr-x`)이 부여된 파일이 보임

- [ ] **Step 3: 의도적 실패 케이스로 검증 (state 파일 없는 상태)**

이 단계는 Task 1이 끝났으므로 state 파일이 모두 있는 상태에서 실행한다. ralph-dual.md만 아직 없을 것이므로 그 부분에서 실패해야 정상.

```bash
bash scripts/ralph-preflight.sh
```
Expected: `FAIL: missing state file: tasks/ralph-dual.md` (Task 3에서 만들 파일)

- [ ] **Step 4: 커밋**

```bash
git add scripts/ralph-preflight.sh
git commit -m "$(cat <<'EOF'
feat(ralph): add preflight check script

Verifies git clean+on-main, state files present and valid JSON,
local web dev server reachable, production reachable, playwright installed.
Run before /ralph-loop start.

EOF
)"
```

---

## Task 3: 메인 Ralph 프롬프트 작성 (`tasks/ralph-dual.md`)

**Files:**
- Create: `tasks/ralph-dual.md`

이 파일이 Ralph가 매 iter 읽는 단일 프롬프트다. 길지만 한 번에 다 작성한다.

- [ ] **Step 1: 메인 프롬프트 파일 작성**

`tasks/ralph-dual.md`:

````markdown
# Ralph Dual-Loop Orchestrator

이 파일은 Ralph가 매 iteration마다 다시 읽는 메인 프롬프트다. 매번 같은 지시를 받지만 `tasks/*.json`과 `tasks/*.md`의 누적 상태가 다르므로 매 iter 다른 행동을 한다.

## 0. 매 iteration 시작 시 실행할 절차

```
1. tasks/bug-state.json, tasks/improve-state.json 읽기
2. tasks/bug-found.md, tasks/improvements.md 읽기 (Dedup용)
3. 종료 조건 검사 (§5)
   • 충족: <promise>DUAL_LOOP_COMPLETE</promise> 출력하고 종료
   • 미충족: 다음 단계
4. 두 subagent를 병렬 dispatch (§3, §4) — Agent tool 두 번을 같은 메시지에서 호출
5. 두 subagent 결과 받아서 §6 절차로 state 갱신
6. tasks/ralph-progress.md에 한 줄 추가
7. 메시지 종료 (Ralph가 stop hook 통해 다음 iter 진입)
```

## 1. 사용 도구 / Subagent

| 용도 | 도구 |
|------|------|
| 헤드리스 브라우저 | `mcp__plugin_playwright_playwright__*` MCP |
| 버그 탐지·수정 subagent | `Agent(subagent_type="general-purpose")` + 아래 §3 프롬프트 |
| 개선 제안 subagent | `Agent(subagent_type="general-purpose")` + 아래 §4 프롬프트 |
| 코드 리뷰 게이트 | `Agent(subagent_type="pr-review-toolkit:code-reviewer")` (bug-detect-fix subagent가 내부에서 호출) |
| 코드 검색·수정 | Read, Edit, Grep, Bash |

## 2. 라운드로빈 인덱스 진행 규칙

매 iter마다:
- `bug-state.last_page_index = (last_page_index + 1) % len(page_queue)`
- `improve-state.last_area_index = (last_area_index + 1) % len(area_queue)`

선택된 페이지/영역을 해당 subagent에 전달한다.

## 3. bug-detect-fix subagent 프롬프트 템플릿

```
당신은 giranjt.com의 단일 페이지에 대한 객관적 버그 탐지·수정·머지를 담당한다.

## 입력
- 대상 페이지 경로: {PAGE}
- 뷰포트 3종: 375x667, 768x1024, 1280x800
- production URL base: https://giranjt.com
- local URL base: http://localhost:3000
- 기존 미해결 버그: tasks/bug-found.md (이미 있는 (page, rule_id, selector)는 dedup)
- sample_listing_id: {SAMPLE_ID} (페이지가 /listings/:id면 이걸로 치환)

## 검사 기준 (모두 도구가 자동 판정 — 주관적 판단 금지)

### A군 — 기능적 결함 (Playwright runtime)
- A-1: pageerror 이벤트 또는 console.error 발생
- A-2: 모든 response.status() ≥ 400 수집 (의도적 401 보호 엔드포인트는 evidence에 표시)
- A-3: <img> 로드 후 naturalWidth === 0
- A-4: <body> 텍스트 길이 < 50자 또는 document.title 비어있음
- A-5: 인터랙티브 요소 클릭 후 2s 내 DOM mutation, 새 네트워크 요청, 네비게이션 모두 없음

### B군 — 시각/UX 결함 (axe-core + DOM 측정)
- B-1: documentElement.scrollWidth > clientWidth + 5
- B-2: scrollWidth>clientWidth+2 AND overflow≠visible AND no text-overflow:ellipsis
- B-3: 375px 뷰포트, 인터랙티브 요소 BoundingRect width<32 또는 height<32
- B-4: axe-core color-contrast violation (WCAG AA)
- B-5: axe-core label/button-name/link-name/image-alt violations

axe-core 주입 방법: browser_evaluate 안에서
  const s = document.createElement('script');
  s.src = 'https://cdnjs.cloudflare.com/ajax/libs/axe-core/4.8.4/axe.min.js';
  document.head.appendChild(s);
  await new Promise(r => s.onload = r);
  const r = await axe.run({runOnly: ['color-contrast','label','button-name','link-name','image-alt']});
  return r.violations;

## 단계별 동작

### 1) Detect
- {PAGE}를 production URL에 3 뷰포트에서 로드
- A1-A5, B1-B5 결과 수집
- normalized_selector 생성 규칙: data-testid 우선, 없으면 id, 없으면 class chain (nth-child 제외)

### 2) Dedup
- tasks/bug-found.md 읽고 (page, rule_id, normalized_selector) 키로 비교
- 새 위반만 처리 대상으로 선정

### 3) 처리할 새 버그가 0건이면
- 결과: { "new_count": 0, "fixed_ids": [], "blocked_ids": [] }
- 종료 (consecutive_no_new가 메인에서 +1 됨)

### 4) 처리할 새 버그가 있으면 각 버그에 대해 순차:

a) bug-found.md에 status: open으로 추가 (id: bug-NNNN, found_at_iter: {ITER})

b) Localize:
   - selector → DOM 컨텍스트 추출 (textContent, aria-label, data-testid, class)
   - Grep으로 web/, frontend/ 소스 검색 (data-testid 또는 한국어 텍스트 우선)
   - 매칭 0건 또는 5건 이상이면:
     • 자동 fix 시도 안 함
     • bug-found.md status를 blocked로 표시, blocked_reason 기록
     • needs-human.md에 추가 (page, rule, selector, evidence, dom_html, grep_candidates)
     • 다음 버그로

c) 같은 파일이 bug-state.files_modified_count에서 이미 3회 이상 → blocked + needs-human

d) Branch + Fix:
   - git checkout -b auto/bug-{ITER}-{ID}
   - 코드 수정 (Edit 도구)

e) Verify Tier 1+2+3 (실패 시 어느 단계에서든 → 브랜치 폐기 + needs-human):
   - Tier 1 lint:
     • 수정 대상이 web/이면: cd web && npx eslint <변경파일>
     • backend/이면: cd backend && golangci-lint run <변경파일> (또는 ./...)
     • frontend/이면: cd frontend && flutter analyze <변경파일>
   - Tier 2 typecheck:
     • web/: cd web && npx tsc --noEmit
     • backend/: cd backend && go vet ./...
   - Tier 3 test:
     • web/: cd web && npx vitest run
     • backend/: cd backend && go test ./...
     • frontend/: cd frontend && flutter test

f) Re-detect on local:
   - http://localhost:3000{PAGE}에 대해 같은 시나리오 재실행
   - 같은 (rule_id, normalized_selector)의 violation이 사라졌는지 확인
   - dev server unreachable이면 1회 재시도, 그래도 실패면 needs-human

g) Code Review:
   - Agent(subagent_type="pr-review-toolkit:code-reviewer", ...) 호출
   - 입력: 변경된 diff + 버그 설명 + fix 의도
   - 출력 파싱: APPROVED 또는 REJECTED
   - REJECTED → 브랜치 폐기 + needs-human (review_feedback 포함)

h) Merge:
   - git checkout main
   - git merge --squash auto/bug-{ITER}-{ID}
   - git commit -m "fix(auto): <bug evidence 한 줄> [bug-{ID}]"
   - git branch -D auto/bug-{ITER}-{ID}
   - bug-found.md status를 fixed로 변경 + fixed_at_iter 기록
   - bug-fixed.md에 항목 추가
   - bug-state.files_modified_count[<file>] += 1

### 5) 결과 반환

JSON으로 메인에게 반환:
```json
{
  "page_visited": "{PAGE_RESOLVED}",
  "new_count": 3,
  "fixed_ids": ["bug-0042", "bug-0043"],
  "blocked_ids": ["bug-0044"],
  "files_touched": ["web/components/listing/Card.tsx"]
}
```

## 금지
- E2E 테스트 추가/수정
- DB 마이그레이션
- package.json 의존성 추가
- main에 직접 커밋 (반드시 auto/bug-* 브랜치 경유)
- production URL에 쓰기 요청
- 로그인 필요한 페이지에서 시드 계정 만들지 말고 익명/비로그인 시나리오만
```

## 4. improve-suggest subagent 프롬프트 템플릿

```
당신은 giranjt.com의 한 feature 영역을 깊게 들여다보고 5관점 개선 제안을 생성한다. 코드 수정은 절대 하지 않는다.

## 입력
- 대상 영역: {AREA} (home/listings/listing_detail/listing_create/chat/reservation/review/report/notification/profile/auth 중 하나)
- production URL base: https://giranjt.com
- 기존 제안: tasks/improvements.md (이미 있는 (area, type, target)은 dedup)

## 영역 → URL 매핑
- home: /
- listings: /listings
- listing_detail: /listings/{first_id} (목록에서 첫 매물 추출)
- listing_create: /listings/new (또는 등록 진입점)
- chat: /chat
- reservation: /reservations
- review: /reviews
- report: 신고 모달 진입 플로우
- notification: /notifications 또는 알림 드롭다운
- profile: /profile
- auth: /login

## 5관점

각 관점마다 가능한 한 발견되는 만큼 제안 생성:

| Type | 관점 | 예시 신호 |
|------|------|-----------|
| ux | 플로우 단순화, 피드백 부족, 마찰점 | 클릭 수가 많음, 로딩 인디케이터 없음, 에러 후 회복 어려움 |
| feature | 없는 기능, 추정 사용자 니즈 | 비교 기능, 즐겨찾기, 필터 부재 |
| performance | 체감 속도, 로딩 패턴 | LCP 느림(목측), 스켈레톤 부재, 이미지 lazy 미적용 |
| content | 한국어 자연스러움, 빈상태 메시지 | 직역 어색함, "데이터 없음" 같은 무미건조 문구 |
| a11y_mobile | 키보드/스크린리더/한손 | tab focus 불가, 모바일에서 reach 어려운 위치 |

## 동작

1) Playwright로 {AREA} URL에 진입 (375x667 + 1280x800 두 뷰포트)
2) 다양한 시나리오 시도: 정상 입력, 빈 입력, 매우 긴 입력, 스크롤 끝까지, 모바일 한손 터치 영역
3) 5관점에서 떠오르는 제안 모두 나열
4) tasks/improvements.md 읽고 (area, type, target) 키로 dedup
5) 새 제안만 추출 (target은 영역 내 구체 요소/플로우 명사. 예: "send_button", "filter_chip", "empty_state_message")

## 출력 형식

각 새 제안을 다음 YAML 블록으로 tasks/improvements.md에 append:

```yaml
- id: imp-{NEXT_NUM}
  found_at_iter: {ITER}
  area: {AREA}
  type: ux|feature|performance|content|a11y_mobile
  target: {target_noun}
  problem: "관찰된 문제 (1문장, 사실 기반)"
  proposal: "구체적 개선안 (1-2문장)"
  effort: trivial|small|medium|large
  impact: low|medium|high
  evidence: "Playwright 관찰 사실 또는 viewport+요소 위치"
```

effort 가이드:
- trivial: 한 줄 카피 변경, 한 줄 CSS
- small: 한 컴포넌트 내 변경, 한 hook 추가
- medium: 새 컴포넌트 1-3개 추가, API 1개 추가
- large: 신규 화면, 데이터 모델 변경

impact 가이드:
- high: 핵심 플로우(거래 성사) 직접 개선
- medium: 사용자 만족도 개선
- low: nice-to-have

## 결과 반환

```json
{
  "area_visited": "{AREA}",
  "new_count": 4,
  "new_ids": ["imp-0067", "imp-0068", "imp-0069", "imp-0070"]
}
```

## 절대 금지
- tasks/improvements.md 외 어떤 파일도 수정하지 않음
- 코드, 설정, 문서 어떤 것도 건드리지 않음
- production URL에 쓰기 요청 (POST/PUT/DELETE) 보내지 않음
- 로그인 필요한 화면은 비로그인 상태로만 관찰
```

## 5. 종료 조건 검사

매 iter 처음에 검사. 충족하면 즉시 `<promise>DUAL_LOOP_COMPLETE</promise>` 출력 후 종료.

```python
# pseudo-code
bug_done = (bug_state.consecutive_no_new >= 3) and (bug_state.pages_visited_count == 6)
imp_done = (improve_state.consecutive_no_new >= 3) and (improve_state.areas_visited_count == 11)
hard_cap = (bug_state.iteration >= 1000)

if (bug_done and imp_done) or hard_cap:
    print("<promise>DUAL_LOOP_COMPLETE</promise>")
    exit
```

## 6. State 갱신 절차 (subagent들 끝난 뒤 메인이 직접)

a) `bug-state.json`:
   - iteration += 1
   - last_page_index = (last_page_index + 1) % 6
   - pages_visited에 page_resolved 추가 (없으면), pages_visited_count 갱신
   - sample_listing_id 갱신 (subagent가 추출했다면)
   - if subagent.new_count == 0: consecutive_no_new += 1
     else: consecutive_no_new = 0
   - totals.found += subagent.new_count
   - totals.fixed += len(subagent.fixed_ids)
   - totals.blocked += len(subagent.blocked_ids)

b) `improve-state.json`:
   - iteration += 1
   - last_area_index = (last_area_index + 1) % 11
   - areas_visited에 area 추가 (없으면), areas_visited_count 갱신
   - if subagent.new_count == 0: consecutive_no_new += 1
     else: consecutive_no_new = 0
   - total_unique_suggestions += subagent.new_count

c) `tasks/ralph-progress.md`에 한 줄 append:
   ```
   iter {N} | bugs +{found}/-{fixed}={open} | improvs +{new}={total} | {events}
   ```
   events 예: "merged auto/bug-12-0007", "blocked bug-12-0008 localize-failed"

## 7. 안전장치

- 메시지 내에 모든 subagent 호출은 같은 message에서 병렬로 (Agent tool 2개 동시 호출)
- bug-found.md, bug-fixed.md, needs-human.md, bug-state.json은 bug-detect-fix subagent만 쓴다
- improvements.md, improve-state.json은 improve-suggest subagent만 쓴다
- ralph-progress.md는 메인이 직접 쓴다
- 각 iter에서 메인은 코드 파일 직접 수정 금지 (subagent에게 위임)

## 8. 시작 시 dirty 상태 회복

만약 이전 iter가 도중에 죽어서 다음 시작 시 git이 dirty면:
1. `auto/bug-*` 브랜치에 있으면 main으로 돌아오고 그 브랜치 삭제
2. uncommitted change 있으면 stash drop (자동 fix 도중 죽은 경우 안전)
3. needs-human.md에 "iter {N-1}: dirty recovery skipped" 한 줄 기록
````

- [ ] **Step 2: 파일이 정상 작성되었는지 확인**

```bash
wc -l tasks/ralph-dual.md
head -20 tasks/ralph-dual.md
```
Expected: 200줄 이상, 첫 줄이 `# Ralph Dual-Loop Orchestrator`

- [ ] **Step 3: 커밋**

```bash
git add tasks/ralph-dual.md
git commit -m "$(cat <<'EOF'
feat(ralph): add main orchestrator prompt

매 iter Ralph가 읽는 단일 SOT. 8개 섹션:
0. 매 iter 절차
1. 사용 도구
2. 라운드로빈 진행
3. bug-detect-fix subagent 프롬프트
4. improve-suggest subagent 프롬프트
5. 종료 조건
6. state 갱신
7. 안전장치
8. dirty 복구

EOF
)"
```

---

## Task 4: Pre-flight 1차 통과 확인

- [ ] **Step 1: 모든 state 파일 + ralph-dual.md + 스크립트가 갖춰졌는지 확인**

```bash
ls -la tasks/ scripts/
```
Expected: 8개 tasks/ 파일 + scripts/ralph-preflight.sh

- [ ] **Step 2: 사용자에게 web dev server 시작 요청**

사용자가 별도 터미널에서:
```
cd web && npm run dev
```

확인:
```bash
curl -sf http://localhost:3000 -o /dev/null && echo "dev server up" || echo "dev server down"
```
Expected: "dev server up"

위 단계는 사용자 액션이 필요하므로, 메인 에이전트가 사용자에게 명시적으로 요청한다.

- [ ] **Step 3: Pre-flight 실행**

```bash
bash scripts/ralph-preflight.sh
```
Expected: 모든 OK 라인 출력 + `=== Preflight passed ===`

실패 시: 실패한 항목 보고 + 사용자에게 수정 방법 안내 후 다시 실행.

---

## Task 5: 1회 dry-run 검증 (Ralph 시작 전)

루프 진입 전, 메인 프롬프트의 절차를 1회만 수동 실행하여 동작이 spec과 맞는지 확인한다.

- [ ] **Step 1: 메인이 직접 1 iter 시뮬레이션**

다음을 직접 실행:

a) `tasks/bug-state.json`, `tasks/improve-state.json` 읽기
b) 종료 조건 검사 (당연히 미충족)
c) 인덱스 진행: `last_page_index = 0` → 페이지 `/`, `last_area_index = 0` → 영역 `home`
d) 두 subagent를 동시에 dispatch:
   - bug-detect-fix subagent에 PAGE=`/`, SAMPLE_ID=null 전달
   - improve-suggest subagent에 AREA=`home` 전달
e) 결과 받아서 state 갱신
f) `tasks/ralph-progress.md`에 한 줄 추가

- [ ] **Step 2: 결과물 확인**

```bash
cat tasks/bug-state.json | python3 -m json.tool
cat tasks/improve-state.json | python3 -m json.tool
cat tasks/ralph-progress.md
ls tasks/bug-found.md tasks/improvements.md
git log --oneline -5
```

확인 포인트:
- iteration이 0 → 1로 증가
- pages_visited에 `/` 추가, pages_visited_count = 1
- areas_visited에 `home` 추가, areas_visited_count = 1
- ralph-progress.md에 `iter 1 | ...` 한 줄 있음
- 새 버그 발견 시 bug-found.md에 항목, fix 머지되었으면 git log에 `fix(auto): ...` 커밋

- [ ] **Step 3: 문제 발견 시 ralph-dual.md 수정 후 재커밋**

dry-run에서 발견된 결함은 메인 프롬프트 (`tasks/ralph-dual.md`) 수정으로 해결. 같은 메시지에서 수정 + 재커밋.

```bash
git add tasks/ralph-dual.md
git commit -m "fix(ralph): <발견된 문제 요약>"
```

dry-run이 spec과 일치하면 다음 단계로.

---

## Task 6: Ralph Loop 본 실행

- [ ] **Step 1: state 파일을 dry-run 직전 상태로 되돌릴지 결정**

dry-run 결과를 보존하고 거기서 이어서 가도 OK (실제로는 그게 더 효율적). 또는 깨끗한 0 상태로 리셋:

리셋이 필요하면:
```bash
git checkout HEAD -- tasks/bug-state.json tasks/improve-state.json tasks/ralph-progress.md
# bug-found.md / improvements.md / bug-fixed.md / needs-human.md는 dry-run 결과 유지
```

리셋 안 하면 바로 다음 단계로.

- [ ] **Step 2: Ralph Loop 시작**

`/ralph-loop` 슬래시 커맨드를 다음과 같이 호출:

```
/ralph-loop "{tasks/ralph-dual.md 전체 내용을 인라인으로}" --max-iterations 1000 --completion-promise "DUAL_LOOP_COMPLETE"
```

또는 간결히:

```
/ralph-loop "tasks/ralph-dual.md 파일을 읽어서 그 안의 모든 지시를 정확히 따라라." --max-iterations 1000 --completion-promise "DUAL_LOOP_COMPLETE"
```

후자를 사용한다 (프롬프트가 너무 길면 슬래시 커맨드에 문제 가능). Ralph가 매 iter `tasks/ralph-dual.md`를 다시 읽어서 동작.

- [ ] **Step 3: 첫 3 iter 모니터링**

처음 3 iter 동안 다음을 확인:
- iter마다 `ralph-progress.md`에 한 줄씩 추가되는지
- bug-state.json/improve-state.json의 iteration 카운터가 증가하는지
- 의도하지 않은 main 직접 커밋이 없는지 (`git log --oneline main` 비교)
- needs-human.md에 합리적인 항목만 들어가는지

문제 발견 시:
- `/cancel-ralph`로 즉시 중단
- ralph-dual.md 수정 + 커밋
- 다시 시작 (Step 2부터)

- [ ] **Step 4: 자동 종료 또는 하드캡 도달 시 결과 확인**

루프 종료 후:
```bash
cat tasks/ralph-progress.md | tail -30
git log --oneline --all | head -50
wc -l tasks/bug-found.md tasks/bug-fixed.md tasks/needs-human.md tasks/improvements.md
```

확인 포인트:
- `<promise>DUAL_LOOP_COMPLETE</promise>` 출력 (정상 종료)
- 또는 iteration이 1000에 도달 (하드캡)
- bug-fixed.md에 머지된 fix 목록
- needs-human.md에 사람 손 필요한 항목 큐
- improvements.md에 우선순위 정렬 가능한 백로그

- [ ] **Step 5: 사용자에게 최종 보고**

다음 형식으로 사용자에게 보고:
- 총 iter 수, 종료 사유
- 자동 머지된 fix 개수 + 주요 카테고리
- needs-human 큐의 항목 수 + 카테고리
- improvements 백로그 항목 수 + impact별 분포
- 다음 액션 제안:
  1. `tasks/needs-human.md` 처리
  2. `tasks/improvements.md`에서 high impact 항목 골라 implement
  3. `deploy/deploy.sh`로 누적된 자동 fix 배포

---

## Self-Review

### Spec coverage

| spec 요구 | 구현 task |
|----------|-----------|
| 단일 Ralph Loop + 병렬 subagent 2개 | Task 3 §1, §3, §4 |
| Pre-flight setup 5단계 | Task 2 + Task 4 |
| 페이지 큐 6개 (라운드로빈) | Task 1 Step 2, Task 3 §3 |
| 영역 큐 11개 (라운드로빈) | Task 1 Step 3, Task 3 §4 |
| 검사 기준 A1-A5, B1-B5 | Task 3 §3 (subagent 프롬프트) |
| Dedup 키 정규화 | Task 3 §3 (selector normalization) |
| Verification gate 5단계 | Task 3 §3 (Verify Tier1+2+3 → re-detect → review) |
| 같은 파일 3회 이상 차단 | Task 1 Step 2 (files_modified_count), Task 3 §3 (4-c) |
| Localize 실패 → needs-human | Task 3 §3 (4-b) |
| 종료조건 (3-consec + 전체 방문) | Task 3 §5, Task 1 (pages_visited / areas_visited) |
| 자동 배포 없음 | Task 3 §3 ("금지" 섹션) |
| 자동 머지는 squash | Task 3 §3 (4-h) |
| improve subagent는 improvements.md만 쓴다 | Task 3 §4 ("절대 금지" 섹션) |
| dirty 복구 | Task 3 §8 |

✅ 커버리지 OK.

### Placeholder scan

- "TBD", "TODO", "implement later" 검색 → 없음
- "Add appropriate error handling" → 없음 (모든 에러 처리 명시: localize 실패 → needs-human, dev server down → 1회 재시도 후 needs-human, review reject → 브랜치 폐기 + needs-human)
- "Similar to Task N" → 없음
- 모든 step에 실제 명령/코드/파일 내용 포함

✅ Placeholder 없음.

### Type consistency

- `bug-found.md` 항목 스키마와 subagent 프롬프트의 status 값 (`open|fixing|fixed|blocked`) 일치 — Task 1 / Task 3
- `improvements.md` 항목의 type enum (`ux|feature|performance|content|a11y_mobile`) — Task 1 / Task 3 동일
- `branch` 명명: `auto/bug-{ITER}-{ID}` 일관 — Task 3 §3 (4-d, 4-h)
- `consecutive_no_new`, `pages_visited_count`, `areas_visited_count` 필드명 — spec과 plan 모두 동일
- `files_modified_count`는 plan에서 처음 도입 (spec의 "같은 파일 3회 이상 차단" 구현)

✅ 일관성 OK.

---

## Out of Scope (이 plan에서 다루지 않음)

- Ralph Loop의 stop hook 내부 동작 수정 (플러그인 자체)
- 별도 worktree 사용 (단일 main 브랜치 + auto/bug-* 임시 브랜치로 충분)
- E2E 테스트 추가 (검증 단계에서만 활용)
- 자동 배포 파이프라인
