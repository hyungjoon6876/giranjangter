# Ralph Dual-Loop Orchestrator

이 파일은 Ralph가 매 iteration마다 다시 읽는 메인 프롬프트다. 매번 같은 지시를 받지만 `tasks/*.json`과 `tasks/*.md`의 누적 상태가 다르므로 매 iter 다른 행동을 한다.

## 0. 매 iteration 시작 시 실행할 절차

```
1. tasks/bug-state.json, tasks/improve-state.json 읽기
2. tasks/bug-found.md, tasks/improvements.md 읽기 (Dedup용)
3. 종료 조건 검사 (§5)
   • 둘 다 충족: <promise>DUAL_LOOP_COMPLETE</promise> 출력하고 종료
   • 미충족: 다음 단계
4. Subagent dispatch:
   • bug_done && imp_done: 도달 불가 (step 3에서 종료됨)
   • bug_done만: improve-suggest만 dispatch (bug-detect-fix skip)
   • imp_done만: bug-detect-fix만 dispatch (improve-suggest skip)
   • 둘 다 미done: 두 subagent 병렬 dispatch (Agent tool 2개 동시 호출)
5. subagent 결과(들) 받아서 §6 절차로 state 갱신
6. tasks/ralph-progress.md에 한 줄 추가
7. 메시지 종료 (Ralph가 stop hook 통해 다음 iter 진입)
```

종료된 loop를 skip하는 이유: 검사·제안에 비용(~3분)이 들고, 한 번 종료 조건 충족한 loop는 추가 작업이 거의 의미 없다. 단, bug_done 후에도 improve의 제안이 코드를 건드리지는 않으므로 새 버그가 생길 수 없다.

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

페이지 큐의 `/listings/:id`는 placeholder. 다음과 같이 치환:
1. 시작 시 또는 `bug-state.sample_listing_id`가 null이면 `/listings`에 접근해 첫 매물의 id를 추출
2. 추출 성공 시 `bug-state.sample_listing_id`에 캐시
3. 매 iter `/listings/:id`가 선택되면 캐시된 id로 치환
4. 추출 실패 시 그 iter는 페이지 인덱스를 +1 더 진행 (skip)

## 3. bug-detect-fix subagent 프롬프트 템플릿

다음 템플릿의 `{PAGE}`, `{ITER}`, `{SAMPLE_ID}`를 메인이 치환해서 Agent tool로 호출한다.

```
당신은 giranjt.com의 단일 페이지에 대한 객관적 버그 탐지·수정·머지를 담당한다.

## 입력
- 대상 페이지 경로: {PAGE}
- 뷰포트 3종: 375x667, 768x1024, 1280x800
- production URL base: https://giranjt.com
- local URL base: http://localhost:3000
- 기존 미해결 버그: tasks/bug-found.md (이미 있는 (page, rule_id, selector)는 dedup)
- sample_listing_id: {SAMPLE_ID} (페이지가 /listings/:id면 이걸로 치환)
- 현재 iteration: {ITER}

## 검사 기준 (모두 도구가 자동 판정 — 주관적 판단 금지)

### A군 — 기능적 결함 (Playwright runtime)
- A-1: pageerror 이벤트 또는 console.error 발생
- A-2: response.status() ≥ 400 (의도적 401 보호 엔드포인트는 evidence에 표시)
- A-3: <img> 로드 후 naturalWidth === 0
- A-4: <body> 텍스트 길이 < 50자 또는 document.title 비어있음
- A-5: 인터랙티브 요소 클릭 후 2s 내 DOM mutation, 새 네트워크 요청, 네비게이션 모두 없음

### B군 — 시각/UX 결함 (axe-core + DOM 측정)
- B-1: documentElement.scrollWidth > clientWidth + 5
- B-2: scrollWidth>clientWidth+2 AND overflow≠visible AND no text-overflow:ellipsis
- B-3: 375px 뷰포트, 인터랙티브 요소 BoundingRect width<32 또는 height<32
- B-4: axe-core color-contrast violation (WCAG AA)
- B-5: axe-core label/button-name/link-name/image-alt violations

axe-core 주입 (browser_evaluate 안에서):
  const s = document.createElement('script');
  s.src = 'https://cdnjs.cloudflare.com/ajax/libs/axe-core/4.8.4/axe.min.js';
  document.head.appendChild(s);
  await new Promise(r => s.onload = r);
  const r = await axe.run({runOnly: ['color-contrast','label','button-name','link-name','image-alt']});
  return r.violations;

normalized_selector 규칙: data-testid 우선 → 없으면 id → 없으면 class chain (nth-child 인덱스 제외)

## 단계별 동작

### 1) Detect
{PAGE}를 production URL에 3 뷰포트에서 로드, A1-A5+B1-B5 결과 수집.

### 2) Dedup + 처리 대상 결정

tasks/bug-found.md 읽고 (page, rule_id, normalized_selector) 키로 비교.

- **new_count**: 이번에 detect로 새로 발견된 항목 중 bug-found.md에 없는 개수 (consecutive_no_new 판정에만 사용)
- **처리 대상**:
  - 이번에 새로 발견된 모든 항목 (위 dedup으로 새 entry 추가)
  - PLUS bug-found.md에서 `page == {PAGE}` AND `status == open` 인 기존 항목 (이전 iter에서 처리 못 한 것)
- 처리 대상은 모두 이 iter 안에서 순차 처리한다. **deferral 금지**.
- `status == blocked`인 항목은 재처리하지 않는다 (사람 손 필요로 분류되어 needs-human.md로 이미 라우팅됨).

### 3) 처리 대상이 0건이면
바로 결과 반환:
  { "page_visited": "{PAGE_RESOLVED}", "new_count": 0, "fixed_ids": [], "blocked_ids": [], "files_touched": [] }

### 4) 처리 대상이 있으면 각각 순차:

a) bug-found.md에 status: open으로 추가 (id: bug-NNNN, found_at_iter: {ITER})

b) Localize:
   - selector → DOM 컨텍스트 추출 (textContent, aria-label, data-testid, class)
   - Grep으로 web/, frontend/ 소스 검색 (data-testid 우선, 없으면 한국어 텍스트, 클래스명)
   - 매칭 0건 또는 5건 이상 → 자동 fix 안 함, blocked + needs-human

c) bug-state.files_modified_count[<file>] >= 3인 파일 → blocked + needs-human

d) Branch + Fix (코드 파일만 수정. tasks/* 상태 파일은 main에서만 수정):
   - git checkout -b auto/bug-{ITER}-{ID}
   - Edit 도구로 코드 수정 (web/, backend/, frontend/, admin/, shared/ 안)
   - tasks/* 파일은 이 단계에서 건드리지 않는다

e) Verify Tier 1+2+3 (실패 시 어느 단계에서든 → 브랜치 폐기 + needs-human):
   Tier 1 lint:
     web/ → cd web && npx eslint <변경파일>
     backend/ → cd backend && golangci-lint run ./... (없으면 go vet ./...)
     frontend/ → cd frontend && flutter analyze <변경파일>
   Tier 2 typecheck:
     web/ → cd web && npx tsc --noEmit
     backend/ → cd backend && go vet ./...
   Tier 3 test:
     web/ → cd web && npx vitest run
     backend/ → cd backend && go test ./...
     frontend/ → cd frontend && flutter test

f) Re-detect on local:
   http://localhost:3000{PAGE_RESOLVED}에 같은 시나리오 재실행, 같은 (rule_id, selector) violation 사라짐 확인.
   dev server unreachable이면 1회 재시도, 그래도 실패 → needs-human.

g) Code Review:
   Agent(subagent_type="pr-review-toolkit:code-reviewer", ...) 호출, diff + 버그 설명 + fix 의도 전달.
   APPROVED 또는 REJECTED 출력 파싱.
   REJECTED → 브랜치 폐기 + needs-human (review_feedback 포함)

h) Merge:
   - git checkout main
   - 코드 파일만 stage: `git add <code-files>` (tasks/* 제외)
   - git merge --squash auto/bug-{ITER}-{ID}
   - git commit -m "fix(auto): <evidence 한 줄> [bug-{ID}]"
   - git branch -D auto/bug-{ITER}-{ID}
   - 그 다음 main 작업 트리에서 (uncommitted) tasks/* 갱신:
     • bug-found.md status를 fixed로 변경 + fixed_at_iter 기록
     • bug-fixed.md에 항목 추가
   - bug-state.files_modified_count[<file>] += 1 — 메인이 §6에서 처리

블록(localize/verify/review 실패) 시:
   - git checkout main && git branch -D auto/bug-{ITER}-{ID}
   - bug-found.md에서 해당 항목 status를 blocked로 변경 + blocked_reason 기록 (uncommitted)
   - needs-human.md에 항목 추가 (uncommitted)

### 5) 결과 반환

JSON으로 메인에게 반환:
{
  "page_visited": "<actual URL path>",
  "new_count": <int>,
  "fixed_ids": [...],
  "blocked_ids": [...],
  "files_touched": [...]
}

## 절대 금지
- E2E 테스트 추가/수정 (검증에만 활용)
- DB 마이그레이션
- package.json 의존성 추가
- main에 직접 커밋 (반드시 auto/bug-* 브랜치 경유)
- production URL에 쓰기 요청 (POST/PUT/DELETE/PATCH)
- 로그인 시드 계정 생성 — 익명/비로그인 시나리오만
- tasks/improvements.md, tasks/improve-state.json 수정
```

## 4. improve-suggest subagent 프롬프트 템플릿

다음 템플릿의 `{AREA}`, `{ITER}`, `{NEXT_ID_NUM}`을 메인이 치환해서 Agent tool로 호출한다.

```
당신은 giranjt.com의 한 feature 영역을 깊게 들여다보고 5관점 개선 제안을 생성한다. 코드 수정은 절대 하지 않는다.

## 입력
- 대상 영역: {AREA}
- production URL base: https://giranjt.com
- 기존 제안: tasks/improvements.md (이미 있는 (area, type, target)은 dedup)
- 현재 iteration: {ITER}
- 다음 id 시작 번호: imp-{NEXT_ID_NUM}

## 영역 → URL 매핑
- home: /
- listings: /listings
- listing_detail: /listings/<목록 첫 매물 id>
- listing_create: /listings/new (또는 등록 진입점)
- chat: /chat
- reservation: /reservations
- review: /reviews
- report: 신고 모달 진입 플로우 (매물 상세에서)
- notification: /notifications 또는 알림 드롭다운
- profile: /profile
- auth: /login

## 5관점

| Type | 관점 | 신호 예시 |
|------|------|-----------|
| ux | 플로우 단순화, 피드백 부족, 마찰점 | 클릭 수 많음, 로딩 인디케이터 없음, 에러 후 회복 어려움 |
| feature | 없는 기능, 추정 사용자 니즈 | 비교, 즐겨찾기, 필터 부재 |
| performance | 체감 속도, 로딩 패턴 | 스켈레톤 부재, 이미지 lazy 미적용 |
| content | 한국어 자연스러움, 빈상태 메시지 | 직역 어색함, "데이터 없음" 같은 무미건조 문구 |
| a11y_mobile | 키보드/스크린리더/한손 | tab focus 불가, 모바일 reach 어려운 위치 |

## 동작

1) Playwright로 {AREA} URL에 진입 (375x667 + 1280x800 두 뷰포트)
2) 다양한 시나리오 시도: 정상 입력, 빈 입력, 매우 긴 입력, 스크롤 끝까지, 모바일 한손 터치
3) 5관점에서 떠오르는 제안 모두 나열
4) tasks/improvements.md 읽고 (area, type, target) 키로 dedup
5) 새 제안만 추출

target은 영역 내 구체 요소/플로우 명사 (예: "send_button", "filter_chip", "empty_state_message", "submit_flow").

## 출력 형식 — tasks/improvements.md에 append

각 새 제안마다 다음 YAML 블록:

- id: imp-{ID}
  found_at_iter: {ITER}
  area: {AREA}
  type: ux|feature|performance|content|a11y_mobile
  target: <target_noun>
  problem: "관찰된 문제 (1문장, 사실 기반)"
  proposal: "구체적 개선안 (1-2문장)"
  effort: trivial|small|medium|large
  impact: low|medium|high
  evidence: "Playwright 관찰 사실 또는 viewport+요소 위치"

effort 가이드:
- trivial: 한 줄 카피, 한 줄 CSS
- small: 한 컴포넌트 내 변경, 한 hook 추가
- medium: 새 컴포넌트 1-3개, API 1개 추가
- large: 신규 화면, 데이터 모델 변경

impact 가이드:
- high: 핵심 플로우(거래 성사) 직접 개선
- medium: 사용자 만족도 개선
- low: nice-to-have

## 결과 반환

{
  "area_visited": "{AREA}",
  "new_count": <int>,
  "new_ids": ["imp-NNNN", ...]
}

## 절대 금지
- tasks/improvements.md 외 어떤 파일도 수정/생성 금지
- 코드, 설정, 문서 어떤 것도 건드리지 않음
- production URL에 쓰기 요청 (POST/PUT/DELETE/PATCH)
- 로그인 필요한 화면은 비로그인 상태로만 관찰
- bug-state.json, bug-found.md, bug-fixed.md, needs-human.md 수정
```

## 5. 종료 조건 검사

매 iter 처음에 검사. 충족하면 즉시 `<promise>DUAL_LOOP_COMPLETE</promise>` 출력 후 종료.

```
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
   - files_modified_count[<file>] += 1 for each file in subagent.files_touched

b) `improve-state.json`:
   - iteration += 1
   - last_area_index = (last_area_index + 1) % 11
   - areas_visited에 area 추가 (없으면), areas_visited_count 갱신
   - if subagent.new_count == 0: consecutive_no_new += 1
     else: consecutive_no_new = 0
   - total_unique_suggestions += subagent.new_count

c) `tasks/ralph-progress.md`에 한 줄 append:
   `iter {N} | bugs +{found}/-{fixed}={open} | improvs +{new}={total} | {events}`
   events 예: "merged auto/bug-12-0007", "blocked bug-12-0008 localize-failed", "skip /listings/:id (no sample)"

d) **Iter end commit** — 모든 tasks/* 변경(JSON state + 누적 .md들)을 한 커밋으로 묶음:
   ```bash
   git add tasks/
   git commit -m "chore(ralph): iter {N} state — bugs +{found}/-{fixed}, improvs +{new}"
   ```
   이렇게 하면 코드 fix 커밋(`fix(auto): ...`)과 상태 커밋(`chore(ralph): iter N state`)이 분리되어 history가 깔끔하다.
   상태 변경이 없으면(아무 일도 일어나지 않으면) 빈 커밋 만들지 말고 skip.

## 7. 안전장치 / 동시성

- 두 subagent는 같은 message에서 Agent tool 2개 동시 호출 (병렬)
- bug-found.md, bug-fixed.md, needs-human.md는 bug-detect-fix subagent가 main 작업 트리에서 수정 (uncommitted)
- bug-state.json은 메인이 §6에서 갱신
- improvements.md는 improve-suggest subagent가 main 작업 트리에서 수정 (uncommitted)
- improve-state.json은 메인이 §6에서 갱신
- ralph-progress.md는 메인이 §6에서 추가
- **모든 tasks/* 변경의 커밋은 메인이 §6 d)에서 한 번에 묶음** (subagent는 tasks/* 커밋하지 않음)
- 코드 fix 커밋은 별도 (auto/bug-* 브랜치 squash merge로 main에 들어옴)
- 매 iter 메인은 코드 파일 직접 수정 금지 (subagent에게 위임)
- 같은 파일 3회 이상 자동수정 차단 (files_modified_count로 체크)

## 8. 시작 시 dirty 상태 회복

만약 이전 iter가 도중에 죽어서 다음 시작 시 git이 dirty면:
1. `auto/bug-*` 브랜치에 있으면 main으로 돌아오고 그 브랜치 삭제
   - `git checkout main && git branch -D <auto/bug-*>` (작업 손실됨, 재시도 가능)
2. uncommitted change 있으면 stash drop (자동 fix 도중 죽은 경우 안전)
   - `git stash && git stash drop`
3. needs-human.md에 "iter {N-1}: dirty recovery skipped" 한 줄 기록
4. 그 다음 정상 절차로 진입
