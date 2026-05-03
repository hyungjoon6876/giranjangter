# Ralph Dual-Loop Progress

> iter당 한 줄 로그. 형식: `iter <N> | bugs +<n>/-<m>=<open> | improvs +<k>=<total> | <event>`
> 예: `iter 12 | bugs +2/-1=3 | improvs +1=18 | merged auto/bug-12-0007`

iter 1 | bugs +3/-1=2 | improvs +25=25 | dry-run; merged auto/bug-1-0003 (web/components/listing/listing-filters.tsx); bug-0001/0002 left open (color-contrast, will be reprocessed when / cycles back)
iter 2 | bugs +1/-1=2 | improvs +16=41 | merged auto/bug-2-0004 (web/app/listings/page.tsx new); /listings prod 404 → redirect("/"); cached sample_listing_id b155f6a7…
iter 3 | bugs +2/-3=1 | improvs +19=60 | merged auto/bug-3-0005 (shared/design-tokens.json textDim #5A4E54→#8C8085 4.84:1+) + auto/bug-3-0006 (header.tsx login link 30→34px); shared token fix also resolved carried-over bug-0001
iter 4 | bugs +1/-1=1 | improvs +15=75 | merged auto/bug-4-0007 (web/app/login/page.tsx '로그인 없이 둘러보기' 18→44px tap target)
iter 5 | bugs +0/-0=1 | improvs +15=90 | /profile clean (cons_no_new=1); prod stale text-dim contrast = bug-0005 root cause (deploy pending); chat: redirect param drop + anon-invisible chat CTA + image attach missing
iter 6 | bugs +0/-0=1 | improvs +11=101 | /chats anon→login redirect, all clean (cons_no_new=2); page queue 1바퀴 완료 6/6; reservation: confirm/cancel UI 부재, /reservations 404, KST→UTC timezone bug
iter 7 | bugs +0/-1=0 | improvs +16=117 | merged auto/bug-7-0002 (web/lib/utils.ts STATUS_COLORS.available #059669→#10B981 contrast 4.41→7.06); 🎉 BUG LOOP DONE (cons=3, pages=6/6, 0 open); review: ReviewModal dead code, rating binary only, /reviews 404, no edit/delete/report
iter 8 | bugs SKIPPED (loop done) | improvs +12=129 | report: imp-0118 chat report 100% silent fail (target_type 불일치 chat_room vs backend oneof), no my-reports page, no evidence upload, no SLA notification
iter 9 | bugs SKIPPED | improvs +8=137 | notification: imp-0130 TS Notification type 불일치 (silent broken render) + imp-0131 backend INSERT 0건 (feature 사실상 dead code) + 인앱 채널만 (no web push/email)
iter 10 | bugs SKIPPED | improvs +18=155 | profile: /profile/<userId> 인덱스 404 (only /reviews), trustBadge raw 렌더, logout 확인 없음, blocked-list/settings/characters 화면 부재
iter 11 | bugs SKIPPED | improvs +15=170 | auth: oauth provider 다양화, 계정 삭제 self-serve, device session mgmt, multi-tab logout sync, T&C consent tracking; 🎯 areas 11/11 완주
iter 12 | bugs SKIPPED | improvs +18=188 | home round2: filter URL state 부재, twitter:image localhost, 계정 카테고리 ToS 위반 위험, Korean 만/억 가격 포맷, line-clamp/preconnect/manifest service-worker
iter 13 | bugs SKIPPED | improvs +12=200 | listings round2: SEO sitemap/robots, JSON-LD ItemList, price-drop watch, saved searches, listing comparison, RSS/Atom feed; ⭐ 200 milestone
iter 14 | bugs SKIPPED | improvs +12=212 | listing_detail round2: JSON-LD Product/Offer + BreadcrumbList, item DB link, stale listing warning, counter-offer form, newcomer safe-trade warning, KakaoTalk share
iter 15 | bugs SKIPPED | improvs +15=227 | listing_create round2: bulk paste-to-create, image OCR 강화 detect, scheduled publish, KRW unit helper, duplication/templates; ⚠️ image contract bug (web 'images' vs backend 'imageIds')
iter 16 | bugs SKIPPED | improvs +14=241 | chat round2: in-room search, input draft 저장, Web Notification, URL scam warning, chat export 분쟁 증거, scroll-anchor "↓N new" pill
iter 17 | bugs SKIPPED | improvs +10=251 | reservation round2: imp-0242/0243 contract & cache bugs, reschedule state machine, .ics export, no-show flow, conflict preflight, expires_at sweep
iter 18 | bugs SKIPPED | improvs +8=259 | review round2: ⚠️ imp-0252 권한 체크 누락 (3rd party 후기 작성), imp-0254 LIMIT 50 silent 누락, verified-purchase 배지 부재, 후기 모더레이션 0건; 16→8 saturation 시작
