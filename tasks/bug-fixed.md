# Bug Fixed Log

> Loop 1이 자동으로 발견 → 수정 → 리뷰 → main 머지 완료한 버그.

- id: bug-0003
  fixed_at_iter: 1
  page: /
  rule_id: B-3-tap-target
  fix_summary: "ListingFilters 서버/카테고리 필터 칩 py-1.5(30px) → py-2 + min-h-[32px] (34px)로 키워 WCAG 2.5.5 tap target 충족"
  files_touched: ["web/components/listing/listing-filters.tsx"]
  verification: "localhost:3000 재탐지: 전체 칩 30px → 34px 확인. 244 vitest tests pass, eslint clean, tsc clean (preexisting 무관 에러 제외)."
  squash_commit: "5d78bba"

- id: bug-0004
  fixed_at_iter: 2
  page: /listings
  rule_id: A-2-http-error
  fix_summary: "web/app/listings/page.tsx 신규 추가 — server-side redirect('/')로 canonical browse URL의 404를 제거"
  files_touched: ["web/app/listings/page.tsx"]
  verification: "curl http://localhost:3000/listings: 307 Location: / 확인 (이전 production은 404). next build: /listings ○ Static prerendered. 244 vitest tests pass, eslint clean."
  squash_commit: "75209c9"

- id: bug-0005
  fixed_at_iter: 3
  page: /listings/:id
  rule_id: B-4-color-contrast
  fix_summary: "shared/design-tokens.json textDim #5A4E54 → #8C8085. axe-core 9개 violation 해결, 동일 root token 사용하는 모든 페이지에 영향. bug-0001 (page=/, 같은 .text-text-dim selector) 도 함께 해결."
  files_touched: ["shared/design-tokens.json"]
  verification: "localhost:3000/ axe color-contrast 0 violations 확인. 244 vitest tests pass, eslint clean. 새 색상: 5.28:1 on #08080C, 4.84:1 on #17131A (PASS WCAG AA 4.5:1)."
  squash_commit: "0808a83"
  resolved_alongside: ["bug-0001"]

- id: bug-0006
  fixed_at_iter: 3
  page: /listings/:id
  rule_id: B-3-tap-target
  fix_summary: "web/components/layout/header.tsx — 로그인 Link py-1.5 → py-2 + min-h-[32px] inline-flex items-center. 30px → 34px로 키워 WCAG 2.5.5 / 프로젝트 ≥32px 룰 충족."
  files_touched: ["web/components/layout/header.tsx"]
  verification: "localhost:3000 375x667 재탐지: 로그인 link height 30 → 34px. 244 vitest pass, eslint clean."
  squash_commit: "03a7162"

- id: bug-0007
  fixed_at_iter: 4
  page: /login
  rule_id: B-3-tap-target
  fix_summary: "web/app/login/page.tsx — '로그인 없이 둘러보기' button에 px-4 py-2 min-h-[44px] 추가. 패딩 없던 18px 높이 → 44px로 키워 WCAG 2.5.5 minimum 24px / 프로젝트 ≥32px 충족 (Google 권장 44px touch target)."
  files_touched: ["web/app/login/page.tsx"]
  verification: "localhost:3000/login 375x667 재탐지: button 103x18 → 132x44 확인. 244 vitest pass, eslint clean. 사전존재 type 에러는 변경 무관 (test fixture 타입)."
  squash_commit: "65bff90"
