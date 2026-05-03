# Bug Found Log

> Loop 1이 발견했지만 아직 처리 중이거나 처리 대기인 버그. 형식: YAML 블록.
> Dedup key: (page, rule_id, normalized_selector)

- id: bug-0001
  found_at_iter: 1
  page: /
  rule_id: B-4-color-contrast
  normalized_selector: ".text-text-dim"
  viewport: 1280x800
  evidence: "axe-core color-contrast: foreground #5a4e54 on background #17131a = 2.31:1 (need 4.5:1). Affects 6+ DOM nodes including footer nav links, copyright, contact email."
  source_root_cause: "shared/design-tokens.json — colors.textDim = #5A4E54 (too dark vs bg #17131A)"
  status: open

- id: bug-0002
  found_at_iter: 1
  page: /
  rule_id: B-4-color-contrast
  normalized_selector: "span.inline-block.px-2.py-0\\.5.rounded[style*='rgb(5, 150, 105)']"
  viewport: 1280x800
  evidence: "axe-core color-contrast: 판매중 status badge (color #059669 on tinted bg rgba(5,150,105,0.1) over page bg #17131A) computed contrast 4.41:1 (need 4.5:1)."
  source_root_cause: "web/lib/utils.ts — STATUS_COLORS.available = '#059669' (foreground darker than needed for tinted bg)"
  status: open

- id: bug-0003
  found_at_iter: 1
  page: /
  rule_id: B-3-tap-target
  normalized_selector: "button.bg-medium.text-text-secondary[class*='py-1.5']"
  viewport: 375x667
  evidence: "20+ server filter chip buttons measure 30px tall (text-sm 12px + py-1.5 padding 6px*2 = 30px). WCAG 2.5.5 minimum 24px, our project rule ≥32px."
  source_root_cause: "web/components/listing/listing-filters.tsx:29 — chipBase 'py-1.5' produces 30px height"
  status: fixed
  fixed_at_iter: 1

- id: bug-0004
  found_at_iter: 2
  page: /listings
  rule_id: A-2-http-error
  normalized_selector: "document"
  viewport: 1280x800
  evidence: "GET https://giranjt.com/listings returns HTTP/2 404 (verified via curl, x-nextjs-cache: HIT). Page renders Next.js default '404 — This page could not be found.' UI. /listings has no page.tsx; only /listings/[id] subdirectory exists. Users following '/listings' (canonical browse URL, expected from /listings/[id] hierarchy) hit a dead end."
  source_root_cause: "web/app/listings/ has no page.tsx — listings browse is rendered by app/page.tsx (root /), so /listings is unrouted"
  status: fixed
  fixed_at_iter: 2
