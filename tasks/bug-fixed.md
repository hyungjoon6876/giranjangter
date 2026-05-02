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
