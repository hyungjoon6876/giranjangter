# Ralph Dual-Loop Progress

> iter당 한 줄 로그. 형식: `iter <N> | bugs +<n>/-<m>=<open> | improvs +<k>=<total> | <event>`
> 예: `iter 12 | bugs +2/-1=3 | improvs +1=18 | merged auto/bug-12-0007`

iter 1 | bugs +3/-1=2 | improvs +25=25 | dry-run; merged auto/bug-1-0003 (web/components/listing/listing-filters.tsx); bug-0001/0002 left open (color-contrast, will be reprocessed when / cycles back)
iter 2 | bugs +1/-1=2 | improvs +16=41 | merged auto/bug-2-0004 (web/app/listings/page.tsx new); /listings prod 404 → redirect("/"); cached sample_listing_id b155f6a7…
iter 3 | bugs +2/-3=1 | improvs +19=60 | merged auto/bug-3-0005 (shared/design-tokens.json textDim #5A4E54→#8C8085 4.84:1+) + auto/bug-3-0006 (header.tsx login link 30→34px); shared token fix also resolved carried-over bug-0001
iter 4 | bugs +1/-1=1 | improvs +15=75 | merged auto/bug-4-0007 (web/app/login/page.tsx '로그인 없이 둘러보기' 18→44px tap target)
