# Improvement Suggestions Backlog

> Loop 2가 수집한 개선 제안. 코드 수정 안 됨. 우선순위 정렬 후 별도로 implement한다.
> Dedup key: (area, type, target)

- id: imp-0001
  found_at_iter: 1
  area: home
  type: ux
  target: server_filter_chip_group
  problem: "375px 모바일 뷰포트에서 28개의 서버 필터 칩이 6+ 줄로 줄바꿈되어 매물 카드 도달 전 약 200px 세로 스크롤을 강요한다."
  proposal: "기본 상태에서는 인기 서버 6-8개만 노출하고 '더 보기' 토글로 나머지를 펼치거나, 서버 검색 입력 + 드롭다운 패턴으로 전환한다. 최근 선택한 서버를 localStorage로 기억해 우선 노출한다."
  effort: medium
  impact: high
  evidence: "Playwright 375x667 viewport: aria-label='서버 필터' group 안에 28개 button (각 width 46-130px, height 30px). 서버 칩 그룹이 매물 첫 카드 시작점을 화면 밖으로 밀어냄."

- id: imp-0002
  found_at_iter: 1
  area: home
  type: a11y_mobile
  target: filter_chip_tap_target
  problem: "서버/카테고리 필터 칩 39개가 모두 height 30-31px로 WCAG 2.5.5 (44x44 권장, 최소 32x32)와 Apple HIG 44pt를 위반한다."
  proposal: "필터 칩의 최소 height를 44px (또는 padding-y 12px+) 로 키우고, 시각적 크기와 별개로 hit-target만 키우려면 ::before 의사 요소로 inflate한다."
  effort: trivial
  impact: medium
  evidence: "Playwright getBoundingClientRect: 서버 칩 28개 모두 h=30, 카테고리 칩 11개 모두 h=31, 374x667 viewport에서."

- id: imp-0003
  found_at_iter: 1
  area: home
  type: feature
  target: favorite_button
  problem: "매물 카드에 '찜 0' 통계는 표시되지만 비로그인/로그인 상태 모두에서 찜하기 트리거 UI가 카드/상세 어디에도 보이지 않아 통계의 의미가 불분명하다."
  proposal: "카드 우상단에 하트 아이콘 토글을 추가하고 비로그인 클릭 시 로그인 모달을 띄운다. 로그인 사용자에게는 즐겨찾기 목록 페이지를 제공한다."
  effort: medium
  impact: high
  evidence: "DOM 쿼리: button[aria-label*=찜] = 0개, [aria-label*=즐겨찾] = 0개. 카드는 텍스트 '찜 0'만 노출하고 인터랙션 없음."

- id: imp-0004
  found_at_iter: 1
  area: home
  type: feature
  target: price_range_filter
  problem: "정렬은 가격순 옵션이 있지만 가격 범위 필터(min/max)가 없어 '300만원 이하 무기' 같은 핵심 거래 의도를 표현할 수 없다. 1원~43,680,742원으로 가격 분포가 광범위하다."
  proposal: "카테고리 칩 아래에 가격 범위 슬라이더 또는 min/max 입력 + 빠른 프리셋(1만/10만/100만/1000만)을 추가한다. URL 쿼리 파라미터로 공유 가능하게 한다."
  effort: medium
  impact: high
  evidence: "Playwright: select[aria-label='정렬 방식'] options = ['최신순','가격 낮은순','가격 높은순','인기순']. [aria-label*=가격] 필터 그룹 = 0개. 노출 가격 범위: 1원, 10원, 200,000원, 43,680,742원."

- id: imp-0005
  found_at_iter: 1
  area: home
  type: feature
  target: enhancement_level_filter
  problem: "리니지 클래식의 핵심 가치 차원인 강화 수치(+0~+10)로 매물을 필터링하지 못한다. 카드에는 '+10', '+3', '+1' 같은 강화 표시가 있어 사용자가 분명히 신경 쓴다."
  proposal: "카테고리 필터 줄에 '+5 이상' / '+7 이상' / '+10' 같은 강화 단계 칩을 추가하거나, 카테고리=무기/방어구 선택 시 동적으로 노출한다."
  effort: medium
  impact: medium
  evidence: "Playwright DOM: 카드에 +10, +3, +1 강화 표시 노출. [aria-label*=강화] 필터 = 0개."

- id: imp-0006
  found_at_iter: 1
  area: home
  type: ux
  target: empty_state_filtered_results
  problem: "서버 필터 클릭 후 결과 0건이면 'h2: 해당 서버에 매물이 없습니다' 텍스트만 표시되고 회복 동작(필터 초기화 / 다른 서버 보기 / 거래 알림 받기)이 없어 막다른 골목이 된다."
  proposal: "빈 상태에 '필터 초기화' 버튼, '인접 서버 보기' 추천(같은 서버군), '이 조건으로 매물 알림 받기' CTA를 추가한다."
  effort: small
  impact: medium
  evidence: "Playwright: 데포로쥬 서버 클릭 후 cardCount=0, h2='해당 서버에 매물이 없습니다' 단독 표시, 보조 액션 0개."

- id: imp-0007
  found_at_iter: 1
  area: home
  type: performance
  target: listing_grid_skeleton
  problem: "API 응답 도착 전 로딩 인디케이터/스켈레톤이 없어 Slow 3G 환경이나 NAS 콜드 스타트 시 빈 화면이 깜빡인다."
  proposal: "TanStack Query loading 상태에 카드 형태의 스켈레톤(2-3개) 또는 shimmer를 노출한다. 헤더/필터는 즉시 렌더, 그리드만 스켈레톤 처리."
  effort: small
  impact: medium
  evidence: "Playwright DOM 쿼리 결과: [class*=skeleton]=0, [data-testid*=skeleton]=0, [role=progressbar]=0. /api/v1/listings 응답 200 OK 동안 그리드 영역이 빈 상태."

- id: imp-0008
  found_at_iter: 1
  area: home
  type: performance
  target: hero_logo_image
  problem: "동일한 logo.png가 헤더(loading=auto)와 히어로(loading=lazy)에서 두 번 다운로드되며, 히어로 이미지가 LCP 후보임에도 lazy로 지정되어 있다."
  proposal: "히어로 이미지를 next/image의 priority(loading=eager + fetchpriority=high)로 설정하고, 동일 자산이라면 한 번만 받도록 srcset 캐시 키를 통일한다."
  effort: trivial
  impact: medium
  evidence: "Playwright img 검사: src=/_next/image?url=%2Flogo.png&w=256&q=75 (loading=auto, header), w=384&q=75 (loading=lazy, hero). 두 개 모두 alt='기란JT'."

- id: imp-0009
  found_at_iter: 1
  area: home
  type: a11y_mobile
  target: viewport_meta_zoom
  problem: "viewport 메타에 maximum-scale=1이 포함되어 핀치 줌이 차단되며, 이는 WCAG 1.4.4 (텍스트 크기 조정 200%) 위반이다."
  proposal: "<meta name='viewport' content='width=device-width, initial-scale=1'> 로 변경한다. 입력 포커스 시 iOS 자동 줌 방지를 원하면 input의 font-size를 16px 이상으로 키운다."
  effort: trivial
  impact: medium
  evidence: "Playwright: meta[name=viewport].content = 'width=device-width, initial-scale=1, maximum-scale=1'."

- id: imp-0010
  found_at_iter: 1
  area: home
  type: a11y_mobile
  target: search_input_mobile_keyboard
  problem: "두 개의 검색 input에 inputmode/enterkeyhint/autocomplete 속성이 없어 모바일 키보드가 일반 텍스트 모드로 뜨고 엔터키 라벨이 '확인' 등 비특화로 노출된다."
  proposal: "input에 enterkeyhint='search', autocomplete='off', spellcheck='false' 를 추가한다. iOS Safari에서 검색 라벨 키보드가 표시된다."
  effort: trivial
  impact: low
  evidence: "Playwright input[type=search] 2개 모두: inputMode='', enterkeyhint=null, autocomplete=''."

- id: imp-0011
  found_at_iter: 1
  area: home
  type: content
  target: hero_secondary_cta_label
  problem: "히어로의 두 번째 CTA 라벨이 '시작하기'인데 비로그인 사용자에게는 /login으로 이동시키는 일반적인 'Get Started' 직역으로 의도가 모호하다."
  proposal: "라벨을 '구글로 3초 가입' 또는 '로그인하고 거래 시작'처럼 행동과 보상을 명시적으로 표현한다. 이미 로그인된 사용자에게는 '내 매물 보기'로 분기."
  effort: trivial
  impact: medium
  evidence: "Playwright: 히어로 CTA 텍스트 '시작하기' (href=/login). 비로그인 anonymous에서 동일 라벨."

- id: imp-0012
  found_at_iter: 1
  area: home
  type: content
  target: empty_state_message_tone
  problem: "서버 필터 적용 후 빈 결과 메시지 '해당 서버에 매물이 없습니다'가 사실 통보형이며 사용자가 다음 행동을 떠올리기 어렵다."
  proposal: "'아직 [서버명]에 매물이 없어요. 첫 등록자가 되어 보세요.' 처럼 서버명을 인용하고 행동 유도 카피로 바꾼다. 옆에 '매물 등록' 버튼을 함께 둔다."
  effort: trivial
  impact: low
  evidence: "Playwright: 데포로쥬 필터 클릭 후 h2 텍스트='해당 서버에 매물이 없습니다' (서버명 미인용, CTA 동반 없음)."

- id: imp-0013
  found_at_iter: 1
  area: home
  type: ux
  target: dual_create_listing_entry
  problem: "모바일 375px에서 우하단 floating + 버튼(56x56)과 하단 네비 '등록' 탭이 모두 /create로 이동해 동일 동작을 두 번 노출한다. 둘이 시각적으로 가까워(y=531/611) 어느 것이 권장 경로인지 혼란."
  proposal: "FAB을 제거하거나(하단 탭이 이미 있음), FAB을 유지하려면 하단 탭의 '등록'을 다른 작업(예: 빠른 거래 검색 토글)으로 재할당한다."
  effort: small
  impact: low
  evidence: "Playwright: a[href='/create'] position:fixed FAB at (288,531) 56x56, nav[aria-label='하단 메뉴']에도 '등록' 항목 존재 at y=611."

- id: imp-0014
  found_at_iter: 1
  area: home
  type: feature
  target: pagination_or_infinite_scroll
  problem: "매물이 늘어났을 때(현재 4개 → N개)의 분량 제어 패턴이 없다. /api/v1/listings?limit=20 호출은 있으나 더 불러오기 트리거가 페이지에 없다."
  proposal: "Intersection Observer 기반 무한 스크롤 또는 명시적 '더 보기' 버튼을 추가한다. URL 쿼리 ?page=2를 지원해 SEO/뒤로가기와 호환되게 한다."
  effort: medium
  impact: medium
  evidence: "Playwright: 스크롤 끝(scrollY=935, scrollHeight=1602)에서 listings=4 그대로, [aria-label*=페이지]=0, '더보기' button=0, 네트워크에 후속 listings 호출 없음."

- id: imp-0015
  found_at_iter: 1
  area: home
  type: a11y_mobile
  target: heading_outline
  problem: "페이지에 h1이 0개이고 첫 헤딩이 h2('매물 목록')이다. 스크린리더의 랜드마크/헤딩 점프가 어색하고 SEO 페이지 주제가 약하다."
  proposal: "히어로 영역의 '리니지 클래식 아이템 거래, 안전하고 무료' 카피를 h1으로 마크업하거나, 시각적으로 숨긴 sr-only h1('기란JT 메인')을 추가한다."
  effort: trivial
  impact: medium
  evidence: "Playwright: querySelectorAll('h1').length = 0, h2.length = 1."

- id: imp-0016
  found_at_iter: 1
  area: home
  type: performance
  target: og_image_url
  problem: "meta[property='og:image'] 값이 'http://localhost:3000/images/og-image.png'로 설정되어 있어 SNS 공유 시 미리보기 이미지가 깨진다."
  proposal: "Next.js metadata에서 절대 URL을 NEXT_PUBLIC_SITE_URL 또는 process.env.VERCEL_URL 기반으로 빌드한다. 프로덕션에서는 https://giranjt.com/images/og-image.png 가 되도록 보장."
  effort: trivial
  impact: high
  evidence: "Playwright: meta[property='og:image'].content = 'http://localhost:3000/images/og-image.png' (production https://giranjt.com 도메인에서)."

- id: imp-0017
  found_at_iter: 1
  area: home
  type: performance
  target: structured_data_jsonld
  problem: "JSON-LD 구조화 데이터가 0개이다. 마켓플레이스 홈에 ItemList/Product 스키마가 없어 구글 SERP의 리치 결과(가격/평점) 노출 기회를 놓친다."
  proposal: "홈 페이지 SSR 응답에 schema.org ItemList + 각 매물의 Offer Product를 JSON-LD로 주입한다. 매물 상세에는 Product+AggregateOffer 스키마."
  effort: small
  impact: medium
  evidence: "Playwright: querySelectorAll('script[type=application/ld+json]').length = 0."

- id: imp-0018
  found_at_iter: 1
  area: home
  type: a11y_mobile
  target: footer_text_contrast
  problem: "푸터 카피라이트 텍스트의 명도 대비가 2.32:1로 WCAG AA(4.5:1) 미달이다. 색상 rgb(90,78,84) on rgb(23,19,26)."
  proposal: "푸터 텍스트 색상을 AppColors.taupe 등 5:1 이상 명도 대비를 보장하는 톤으로 올린다(예: rgb(163,134,118)). 또는 폰트 크기를 14px 이상으로 키워 large text 임계 적용."
  effort: trivial
  impact: medium
  evidence: "Playwright getComputedStyle: footer p color=rgb(90,78,84), bg=rgb(23,19,26), 12px, ratio=2.32."

- id: imp-0019
  found_at_iter: 1
  area: home
  type: content
  target: listing_card_low_quality_titles
  problem: "현재 표시되는 4개 매물 중 3개가 '111233', '123ㄱ1ㄱㅈㄷㄹ', 'ㅇㄴㄹ2ㅈ +3' 같은 의미 없는 자판 입력이다. 서비스 첫인상에서 '진짜 거래 가능한 곳인가' 신뢰가 깨진다."
  proposal: "관리자 모더레이션 큐 또는 자동 휴리스틱(자판 한글 의미 검출, 길이 제한)을 도입해 노이즈 매물을 비공개 처리한다. 단기적으로는 시드 더미 매물을 우선 노출되지 않도록 정렬에 demote."
  effort: small
  impact: high
  evidence: "Playwright snapshot: 매물 카드 제목 '팔로우/머꼬', '123ㄱ1ㄱㅈㄷㄹ/ㅇㄴㄹ2ㅈ +3', '111233/1234 +1' (4개 중 3개가 더미/난수형)."

- id: imp-0020
  found_at_iter: 1
  area: home
  type: ux
  target: filter_active_state_indicator
  problem: "서버/카테고리 필터를 적용해도 페이지 상단 또는 결과 섹션에 '현재 적용된 필터' 칩 요약이 없어 어떤 조건이 켜져 있는지 알기 위해 다시 스크롤해야 한다."
  proposal: "결과 섹션 위에 '데포로쥬 × | 무기 ×' 같은 활성 필터 칩 요약을 노출하고 각 칩에 X 버튼으로 개별 해제, '모두 지우기' 버튼을 함께 둔다."
  effort: small
  impact: medium
  evidence: "Playwright: aria-pressed=true 칩이 필터 그룹 안에만 있고, 결과 섹션('4개 매물') 주변에 적용된 필터 요약 표시가 없음."

- id: imp-0021
  found_at_iter: 1
  area: home
  type: performance
  target: api_request_dedup
  problem: "초기 로드 후 헤더 링크 prefetch로 같은 listing 상세 RSC 요청이 중복 발생한다(동일 매물 id에 _rsc=1r34m, p37cr 두 번 ×4매물 = 8회)."
  proposal: "Next.js Link prefetch 정책을 viewport 진입 시점으로 늦추거나(prefetch={null}), getServerSideProps→getStaticProps로 정적화 가능 페이지를 분리한다. RSC payload에 캐시 헤더를 부여."
  effort: medium
  impact: low
  evidence: "Playwright network: /listings/<id>?_rsc=1r34m + /listings/<id>?_rsc=p37cr 가 4개 매물에 대해 모두 발생, 총 ~16개 prefetch RSC 요청."

- id: imp-0022
  found_at_iter: 1
  area: home
  type: feature
  target: deal_alert_subscription
  problem: "특정 서버/아이템/가격대 조건에 맞는 매물이 등록될 때 알림 받는 'Saved Search' 기능이 없다. 거래 매칭 활성도가 핵심 KPI지만 사용자는 매번 수동으로 새로고침해야 한다."
  proposal: "현재 적용한 필터 조합을 '내 알림' 으로 저장하고, 새 매물 등록 시 SSE/푸시/이메일로 알린다. 알림 페이지는 이미 존재(/notifications)하므로 채널 추가."
  effort: large
  impact: high
  evidence: "Playwright: 헤더에 알림 벨이 있고 /notifications 페이지가 존재하나, 홈 화면 어디에도 'Saved Search/조건 알림' UI 없음. /api/v1/listings는 polling 없이 1회만 호출."

- id: imp-0023
  found_at_iter: 1
  area: home
  type: a11y_mobile
  target: skip_link_visibility
  problem: "skip link 'a[href=#main-content]'은 존재하지만 폭/높이가 1×1px로 키보드 포커스 시에도 화면에 거의 보이지 않아 실효성이 낮다."
  proposal: "skip link에 :focus 시 sr-only를 해제하고 top:0, padding:8px 16px, background:contrast 색상으로 가시화한다. WAI-ARIA 권장 패턴."
  effort: trivial
  impact: low
  evidence: "Playwright getBoundingClientRect: a[href='#main-content'] w=1, h=1 (default state)."

- id: imp-0024
  found_at_iter: 1
  area: home
  type: feature
  target: card_seller_trust_signal
  problem: "매물 카드의 판매자가 '유저_da1e387a' 같은 자동 ID로만 표시되어 신뢰도 시그널(거래 횟수, 평점, 가입 후 기간)이 전혀 없다."
  proposal: "카드 하단에 '⭐ 4.8 · 거래 12회' 또는 가입 일자/배지(인증 사용자/단골)을 추가한다. 0건 사용자는 '신규' 칩으로 정직하게 표기."
  effort: medium
  impact: high
  evidence: "Playwright: 카드 내 판매자 표시 텍스트 '유저_da1e387a', '유저_874fb902' (UUID prefix). 평점/거래수/배지 마크업 0개."

- id: imp-0025
  found_at_iter: 1
  area: home
  type: ux
  target: card_visual_hierarchy
  problem: "카드 내 '판매', '판매중', 서버명 칩 3개가 동일 톤·크기로 일렬 노출되고, 가격(200,000원)이 본문보다 작은 행에 들어가 시선 흐름이 분산된다."
  proposal: "가격을 카드 우측 상단의 가장 큰 텍스트로 승격시키고, 상태 칩은 작은 단일 dot+text로 축약한다(• 판매중). 서버명은 부제목으로 분리."
  effort: small
  impact: medium
  evidence: "Playwright snapshot: 카드 구조 [판매 | 판매중 | 서버명] 한 행, h3 제목, 아이템 행, '200,000원' (heading 아래 별도 작은 행), 통계 행 (조회/찜/채팅), 사용자 행. 가격이 시각적 위계 3-4단계 아래."

- id: imp-0026
  found_at_iter: 2
  area: listings
  type: ux
  target: listings_index_route_404
  problem: "매물 목록의 자연스러운 URL인 https://giranjt.com/listings 가 404를 반환한다. /api/v1/listings, /listings/<id>, footer 'a[href=/]' 등 다른 모든 곳이 'listings'를 마켓 목록 개념으로 쓰는데 정작 /listings 자체가 페이지를 가지지 않아 사용자가 추측한 URL, 외부 공유 링크, SEO 인덱싱 모두 깨진다."
  proposal: "web/app/listings/page.tsx 를 추가해 홈 페이지의 매물 그리드 섹션을 그대로 또는 hero 없이 노출하거나, next.config.ts redirects 로 /listings → / (또는 /#listings) 로 308 리다이렉트시킨다. 동시에 헤더 '마켓' 링크의 href 도 /listings 로 통일."
  effort: small
  impact: high
  evidence: "Playwright: GET https://giranjt.com/listings → Next.js 404 페이지 (h1='404', h2='This page could not be found.', body 텍스트 187자). web/app/listings/ 디렉토리에는 [id]/page.tsx 만 존재, page.tsx 인덱스 라우트 없음."

- id: imp-0027
  found_at_iter: 2
  area: listings
  type: ux
  target: filter_url_state_sync
  problem: "서버 칩, 카테고리 칩, 정렬 select, 검색어가 모두 useState 로컬 상태로만 관리되어 (1) 필터 적용 후 URL 이 변하지 않아 공유 불가, (2) 사용자가 ?server=zillian 같은 URL 을 직접 만들어도 무시되고 '전체'/recent 로 초기화, (3) 카드 클릭 후 뒤로가기 시 필터가 사라진다."
  proposal: "Next.js useSearchParams + router.replace 패턴(혹은 nuqs 라이브러리)으로 server/category/q/sort 4개 상태를 URL 쿼리에 양방향 동기화한다. 초기 마운트 시 URL → state 복원, state 변경 시 history.replaceState 로 URL 갱신."
  effort: medium
  impact: high
  evidence: "Playwright 클릭 테스트: 질리언 칩 클릭 후 location.href 'https://giranjt.com/' 그대로(filterChangedURL=false). 정렬 'price_asc' 변경 후 URL 동일(sortChangedURL=false). 직접 입력 https://giranjt.com/?server=zillian&sort=price_asc 진입 시 activeChips=['전체','전체'], sortValue='recent' 로 무시됨. 코드 web/app/page.tsx L18-21: serverId/categoryId/search/sort 모두 useState."

- id: imp-0028
  found_at_iter: 2
  area: listings
  type: ux
  target: header_search_input_no_action
  problem: "헤더의 '매물 검색' input(input[aria-label=매물 검색], 모바일/데스크탑 2개)에 텍스트 입력 후 Enter 를 눌러도 아무 일도 일어나지 않는다. URL 도 카드 결과도 바뀌지 않고 form 으로 감싸여 있지도 않아 사용자는 검색이 동작하지 않는다고 느낀다."
  proposal: "헤더 input 을 form 으로 감싸 onSubmit 시 router.push(`/?q=${encoded}`) 로 라우팅하고, 홈 페이지 mount 시 useSearchParams().get('q') 를 search state 초기값으로 사용한다. 동시에 input 에 enterkeyhint='search', form action='/' 부여."
  effort: small
  impact: high
  evidence: "Playwright: input[aria-label=매물 검색] 에 '도리깨' 입력 후 keydown Enter 디스패치 → location.href 미변경, cardsAfterSearch=4(필터 안 됨), search.form=null(form 미감쌈)."

- id: imp-0029
  found_at_iter: 2
  area: listings
  type: ux
  target: scroll_restore_back_navigation
  problem: "홈 매물 그리드에서 카드 클릭 → /listings/<id> 진입 → 뒤로가기 시 필터/정렬/검색/스크롤 위치가 모두 초기화되어, 매물 비교 탐색이 불가능하다. 실제 거래 탐색 패턴은 '여러 매물 미리보기 → 뒤로 → 다음 매물' 반복."
  proposal: "(1) 필터/정렬/검색을 URL 쿼리로 옮긴다(imp-0027 의존). (2) Next.js Link scroll prop 과 history scrollRestoration='manual' + sessionStorage 로 스크롤 위치 저장·복원. (3) 카드를 modal/route intercepting (parallel routes) 로 띄워 BFCache 활용도 가능."
  effort: medium
  impact: high
  evidence: "Playwright: 필터 적용 후 카드 클릭 시뮬 → 뒤로가기 시 URL 에 필터 정보 없음. 페이지가 'use client' + useState 로만 상태 관리하므로 뒤로가기 = 컴포넌트 재마운트 = 상태 초기화."

- id: imp-0030
  found_at_iter: 2
  area: listings
  type: feature
  target: listing_card_image_placeholder
  problem: "현재 매물 4개 중 3개가 iconUrl=null 이라 카드에 이미지 영역이 아예 그려지지 않는다(<a> 안 <img>=0). 시각 비교가 불가능하고 카드 높이도 168px ~ 190px 으로 불균일해져 그리드가 들쭉날쭉한다."
  proposal: "카드 상단에 고정 비율 영역(예: aspect-square 64x64 또는 cover 4:3) 을 두고 iconUrl 이 null 일 때 기본 placeholder(아이템 카테고리 SVG, 또는 아이템명 첫 글자 monogram)를 렌더한다. lib/static-icons.ts 에 카테고리별 fallback 매핑."
  effort: small
  impact: medium
  evidence: "Playwright: a[href^=/listings/] 4개 중 1개만 img 보유(도리깨 64x64), 나머지 3개 placeholderRects=[null,null,null]. cardHeights=[190,168,168,168] (도리깨만 아이콘 행 추가로 +22px)."

- id: imp-0031
  found_at_iter: 2
  area: listings
  type: feature
  target: listing_count_total
  problem: "결과 영역에 '4개 매물'이라는 현재 페이지에 '로드된' 개수만 표시될 뿐, 필터 조건의 전체 개수, 또는 '24시간 내 신규 N건' 같은 활동성 시그널이 없다. 사용자는 작은 마켓이 활성인지 죽었는지 판단 불가."
  proposal: "API 응답에 total/last24h 카운트 추가하고, 헤더 영역에 '활성 매물 N건 · 오늘 +M' 라벨을 노출한다. 0건일 때는 '오늘은 신규 등록이 없어요. 알림 받기' CTA 로 전환."
  effort: medium
  impact: medium
  evidence: "Playwright: 결과 라벨 '4개 매물'(data.length 그대로). API 응답 cursor.hasMore=false, next=null 만 존재, total/lastActivity 집계 필드 없음. 카드의 createdAt 모두 2026-03-17~21 (한 달 전)."

- id: imp-0032
  found_at_iter: 2
  area: listings
  type: feature
  target: keyword_filter_match_highlight
  problem: "검색을 하더라도(파라미터로 잘 전달되었다고 가정해도) 카드 제목/아이템명에서 매칭된 키워드가 강조되지 않아 사용자가 왜 이 결과가 나왔는지 추측해야 한다."
  proposal: "검색어가 있을 때 카드 title/itemName 의 일치 부분을 <mark> 또는 text-gold 로 하이라이트한다. ListingCard 컴포넌트에서 search 컨텍스트를 받고 String.prototype.split + map 으로 단순 구현."
  effort: small
  impact: low
  evidence: "Playwright: 검색 input 에 도리깨 입력해도 카드 텍스트 그대로 '도리깨 +10 판매합니다', mark/em/strong 없음. 검색 매칭 시각화 0개."

- id: imp-0033
  found_at_iter: 2
  area: listings
  type: feature
  target: card_relative_time_semantic
  problem: "카드 푸터의 '1개월 전' 라벨이 일반 <span> 으로 렌더되어 (1) <time datetime> 시맨틱이 빠져 스크린리더가 정확한 날짜를 읽지 못하고, (2) hover 로 정확한 등록일 툴팁을 보여주지 못하며, (3) SEO 의 'datePublished' 신호도 없다."
  proposal: "ListingCard 의 시간 표시를 <time dateTime={createdAt} title={fullDateStr}>{relative}</time> 로 변경. format은 lib/date-utils 의 toRelative + toISO 헬퍼로 통일."
  effort: trivial
  impact: low
  evidence: "Playwright: document.querySelectorAll('time').length=0, 모든 카드의 '1개월 전'이 div/span 텍스트. monthEls 4개 모두 단순 텍스트 노드 부모."

- id: imp-0034
  found_at_iter: 2
  area: listings
  type: performance
  target: above_fold_image_eager
  problem: "1280px 데스크탑에서 첫 번째 매물 카드 아이콘 이미지(loading='lazy', fetchPriority='auto')가 above-the-fold 임에도 lazy 로 표시되어 브라우저 우선순위 큐에서 늦게 로드된다."
  proposal: "ListingGrid 또는 ListingCard 컴포넌트에서 index < 4(또는 viewport 안에 들어올 갯수) 카드의 next/image 에 priority 또는 loading='eager' + fetchPriority='high' 부여. 모바일은 첫 1-2개만."
  effort: trivial
  impact: medium
  evidence: "Playwright: 첫 카드 img loading='lazy', fetchPriority='auto'. 하지만 모바일 375px 에서 첫 카드 y=791.5(scroll 필요), 데스크탑 1280px 에서는 첫 카드 above-fold 가능."

- id: imp-0035
  found_at_iter: 2
  area: listings
  type: performance
  target: listing_grid_pagination
  problem: "API /api/v1/listings?limit=20 응답에 cursor.hasMore=false, next=null 로 페이지네이션 메타가 있는데, 프론트엔드에는 useListings 훅이 1회만 호출되고 후속 페이지/cursor 추적이 없다. 매물이 50건이 넘어가면 첫 20개만 보이고 나머지는 영영 도달 불가."
  proposal: "useListings 훅을 useInfiniteQuery 로 전환해 next cursor 를 사용하고, ListingGrid 끝에 IntersectionObserver sentinel + '더 보기' 폴백 버튼을 둔다. ?page= 또는 ?cursor= 쿼리도 함께 지원해 SEO/뒤로가기와 호환."
  effort: medium
  impact: high
  evidence: "Playwright network: /api/v1/listings?limit=20 1회만 호출, scroll 끝(scrollY=935, scrollHeight=1602)에 도달해도 후속 listings 호출 없음. 응답 JSON 의 cursor 필드는 존재(hasMore=false 지만 구조상 페이지네이션 가능)."

- id: imp-0036
  found_at_iter: 2
  area: listings
  type: performance
  target: listings_seo_metadata
  problem: "/listings 가 404 라 인덱싱 자체가 불가능하고, 매물 카드 href 는 UUID(/listings/b155f6a7-a80b-4db5-ac3f-a080b3ad656d)로 사람도 검색엔진도 의미를 읽을 수 없다. document.title 도 모든 페이지가 '기란JT — 리니지 클래식 거래' 로 동일."
  proposal: "(1) /listings 라우트 추가(imp-0026). (2) 카드 href 를 /listings/<slug>-<short_id> 형식(예: dorigae-plus10-zillian-b155f6)으로 변경, generateMetadata 에서 listing.title 을 title 로 사용. (3) /listings 에 ItemList JSON-LD."
  effort: medium
  impact: medium
  evidence: "Playwright: card hrefs 4개 모두 /listings/<UUID> 형식, 슬러그 없음. /listings 페이지 title='기란JT — 리니지 클래식 거래'(글로벌 default), description='리니지 클래식 아이템 거래 중개 플랫폼'(전체 사이트 공통)."

- id: imp-0037
  found_at_iter: 2
  area: listings
  type: content
  target: empty_state_active_filters_summary
  problem: "검색이나 필터 조합으로 결과가 0건일 때 EmptyState 가 어떤 조건이 적용됐는지(서버=질리언 + 카테고리=무기 + 검색=도리깨)를 한 줄로 요약해주지 않는다. 사용자는 결과 0건의 원인을 추측해야 한다."
  proposal: "EmptyState description 을 동적으로 'X 서버 · Y 카테고리 · \"Z\" 검색에 해당하는 매물이 없습니다' 로 구성하고, 그 아래 '필터 초기화' 보조 버튼을 둔다(actionLabel='필터 초기화', secondaryAction='매물 등록')."
  effort: small
  impact: medium
  evidence: "Code web/app/page.tsx L116-122: serverId 일 때 description='다른 서버를 선택하거나 첫 매물을 등록해보세요!' 고정 문자열, 적용된 필터값 미인용. EmptyState 컴포넌트에 actionLabel 1개만 받음."

- id: imp-0038
  found_at_iter: 2
  area: listings
  type: content
  target: noise_listings_first_page
  problem: "비로그인 첫 방문자가 보는 4개 매물 중 3개가 '111233', '123ㄱ1ㄱㅈㄷㄹ', 'ㅇㄴㄹ2ㅈ +3' 같은 자판 노이즈이고 createdAt 이 2026-03-17~21로 6주째 정체된 매물이다. 신뢰도가 무너진다."
  proposal: "(1) 정렬 기본값을 'popular' 또는 'recent_active' 로 두고 viewCount<5 + favoriteCount=0 + chatCount=0 + age>30d 인 매물은 첫 페이지에서 demote. (2) 자판 노이즈 정규식 휴리스틱으로 admin 모더레이션 큐 자동 제출."
  effort: small
  impact: high
  evidence: "Playwright: 카드 4개 텍스트 '도리깨 +10 판매합니다', '팔로우/머꼬', '123ㄱ1ㄱㅈㄷㄹ/ㅇㄴㄹ2ㅈ +3', '111233/1234'. 시간 라벨 모두 '1개월 전'. /api/v1/listings 응답: 4매물 중 3매물의 viewCount<=4, favoriteCount=0~1, chatCount=0~1."

- id: imp-0039
  found_at_iter: 2
  area: listings
  type: a11y_mobile
  target: filter_section_above_fold
  problem: "375x667 모바일에서 첫 매물 카드의 y=791.5px(viewport 의 119%)에서 시작한다. hero(보이는 사용자 미인증) + 28개 서버 칩 + 11개 카테고리 칩 + 결과 헤더가 모두 그리드 위에 배치되어 매물 한 개를 보려면 1.2 화면 분량 스크롤이 강제된다."
  proposal: "(1) 모바일에서 필터 영역을 collapsed by default '필터 ▾' 버튼으로 접고 활성 필터 칩 요약만 노출. (2) hero 를 1줄(매물 둘러보기 CTA)로 축소. (3) sticky '필터/정렬' 바로 스크롤해도 항상 접근 가능."
  effort: medium
  impact: high
  evidence: "Playwright 375x667: filterRect.bottom=564, sortRect.y=755.5, firstCard y=791.5, 결과 첫 카드의 viewport 비율 119%. scrollHeight=1602(viewport 의 2.4 배)."

- id: imp-0040
  found_at_iter: 2
  area: listings
  type: a11y_mobile
  target: sort_select_filter_keyboard
  problem: "정렬 select 가 결과 헤더 안 우측에 있고, 모든 필터 칩이 button[aria-pressed] 패턴인데 키보드 사용자가 '정렬 우선' 흐름을 따르려면 28개 서버 칩 + 11개 카테고리 칩을 거쳐야 도달한다. 또한 select 대신 칩 그룹과 일관된 단일 패턴이 더 학습 효율적이다."
  proposal: "(1) 정렬 select 를 결과 영역 위로 이동시키고 페이지 본문 첫 인터랙션이 되게 한다. 또는 (2) 정렬을 [최신순 | 가격 ↑ | 가격 ↓ | 인기순] segmented control(button[role=radio] group)으로 바꿔 칩 패턴과 통일."
  effort: small
  impact: low
  evidence: "Playwright 375x667: 정렬 select 위치 (236, 755.5) 92x28, 화면 우측 끝. 같은 페이지에 button[aria-pressed] 칩 39개와 select 1개가 혼재(인터랙션 패턴 비일관)."

- id: imp-0041
  found_at_iter: 2
  area: listings
  type: a11y_mobile
  target: card_aria_label_punctuation
  problem: "카드 aria-label='도리깨 +10, 200,000원, 질리언' 형식이 가격을 가운데 두어 스크린리더가 '아이템→가격→서버' 순으로 읽는데, 한국어 사용자에겐 '서버→아이템→가격' 순 정보 우선순위가 자연스럽다. 또한 status('판매중')와 listingType('판매')가 라벨에 빠져 매수/매도 구분 불가."
  proposal: "aria-label 를 '[질리언] 도리깨 +10 판매중, 200,000원' 또는 '판매: 도리깨 +10 (질리언) - 200,000원' 형식으로 재구성하고, 컴포넌트 prop 에 listingType/status 가 노출되도록 한다."
  effort: trivial
  impact: low
  evidence: "Playwright: 카드 4개 aria-label 모두 '<itemName>[ +<level>], <price>원, <serverName>' 패턴. listingType('판매') status('판매중') 누락."

- id: imp-0042
  found_at_iter: 3
  area: listing_detail
  type: ux
  target: anonymous_user_action_bar_invisible
  problem: "비로그인 방문자에게 채팅하기/찜/공유/신고 액션 바가 통째로 숨겨진다. 백엔드 응답이 availableActions=null 이면 web/app/listings/[id]/page.tsx L207 조건 `actions.length > 0` 이 false 가 되어 sticky 액션 바 자체가 렌더되지 않는다. 결과적으로 매물 상세 페이지에 '판매자에게 문의', '관심 등록', '공유' 같은 핵심 거래 진입점이 단 하나도 보이지 않아 익명 사용자는 매물을 발견해도 즉시 다음 행동으로 이어질 수 없다."
  proposal: "익명 사용자도 액션 바를 보이게 하되, 클릭 시 useAuthGuard 의 requireAuth 가 로그인 모달을 띄우는 패턴으로 통일한다. handleChat/toggleFav 가 이미 requireAuth 를 거치므로 actions.length>0 조건을 제거하거나, availableActions 가 null 이면 기본값 ['favorite','start_chat'] 을 사용한다. 공유/신고는 인증 없이도 동작 가능."
  effort: small
  impact: high
  evidence: "Curl https://giranjt.com/api/v1/listings/b155f6a7-a80b-4db5-ac3f-a080b3ad656d (anon) → availableActions=null. Playwright 비로그인 375x667: button 0개(in main), 인터랙티브 요소는 'a[href=/profile/.../reviews]' 1개와 skip-link 만 노출. role=toolbar 미존재. web/app/listings/[id]/page.tsx L62, L207 조건 확인."

- id: imp-0043
  found_at_iter: 3
  area: listing_detail
  type: feature
  target: server_category_name_display
  problem: "API 응답에 serverName='질리언', categoryName='둔기' 가 포함되어 있는데 매물 상세 페이지 어디에도 표시되지 않는다. 사용자는 매물 검색에서 가장 핵심적인 두 차원(서버/카테고리)을 매물을 클릭한 뒤 알 길이 없어, '질리언 서버 무기' 매물인지 다른 서버인지 매번 메모리에 의존해야 한다."
  proposal: "h1 제목 위 또는 거래 정보 카드에 '#질리언 · 둔기' 형태의 메타 칩을 노출한다. 각 칩을 클릭하면 /?server=zillian 또는 /?category=weapon_mace 로 동일 조건의 다른 매물 목록으로 이동(필터 URL 동기화 imp-0027 의존). breadcrumb 형식도 가능: 마켓 > 질리언 > 둔기 > 도리깨 +10."
  effort: small
  impact: high
  evidence: "API 응답 JSON: serverName='질리언', serverId='zillian', categoryName='둔기', categoryId='weapon_mace'. Playwright 페이지 텍스트 검색: '질리언' 0건, '둔기' 0건. h1 위 영역 '판매 / 판매중 / 무관' 칩 3개와 stats(조회/관심/채팅/시간) 만 표시."

- id: imp-0044
  found_at_iter: 3
  area: listing_detail
  type: a11y_mobile
  target: stats_text_color_contrast
  problem: "조회/관심/채팅/시간 4개 stat 텍스트의 명도 대비가 2.52:1 로 WCAG AA(4.5:1) 미달이다. 색상 #5a4e54 (text-text-dim) on #08080c, font-size 12px. axe-core color-contrast violation serious 1건 발생."
  proposal: "Tailwind 토큰 'text-text-dim' (rgb(90,78,84)) 를 'text-text-secondary' (taupe rgb(163,134,118)) 로 변경하거나, 디자인 토큰의 text-dim 값 자체를 명도 50%+ 로 올린다. 동시에 stats 폰트를 13-14px 로 키워 large-text 임계 적용 가능. shared/design-tokens.json + web/app/globals.css 의 --color-text-dim 변수."
  effort: trivial
  impact: medium
  evidence: "Playwright + axe-core: '조회 15', '관심 0', '채팅 0', '1개월 전' 4개 span 모두 color=#5a4e54, bg=#08080c, ratio=2.52, 12px. web/app/listings/[id]/page.tsx L113 'text-text-dim' 클래스 사용."

- id: imp-0045
  found_at_iter: 3
  area: listing_detail
  type: feature
  target: image_gallery_zoom_lightbox
  problem: "API 의 listing.images 배열을 그리드로 렌더하지만(page.tsx L121-143) 클릭/탭 시 확대(lightbox/modal) 하는 인터랙션이 없다. iconUrl(64x64) fallback 도 모니터 화면에서는 작은 썸네일이라 +10 강화 아이템의 디테일을 확인할 수 없다. 거래에서 이미지는 신뢰의 1차 시그널인데 작은 썸네일만 제공하면 의심이 커진다."
  proposal: "이미지 클릭 시 fullscreen lightbox(예: yet-another-react-lightbox 또는 자체 dialog)를 열어 핀치줌/스와이프 가능하게 한다. iconUrl 만 있는 경우에도 카드를 키워 200x200 이상으로 노출. 키보드 ESC 닫기, Tab trap, role=dialog aria-modal=true 적용."
  effort: medium
  impact: high
  evidence: "Playwright: main img 1개(item icon 64x64), tabIndex=-1, role=null, aria-label=null, parentTag='DIV' (a/button 안에 없음 → isClickable=false), cursor='auto'. 코드 page.tsx L122-141 grid 만 그리고 onClick 미구현."

- id: imp-0046
  found_at_iter: 3
  area: listing_detail
  type: feature
  target: similar_listings_section
  problem: "매물 상세 하단에 '비슷한 매물 / 같은 서버 매물 / 같은 판매자의 다른 매물' 섹션이 없다. 사용자는 마음에 안 들면 뒤로가기 → 다시 목록 탐색 루프로 돌아가야 하고, 거래 매칭 컨버전이 떨어진다. 데스크탑(1280px) 에서 contentMaxWidth=896 으로 좌우 약 184px 씩 비어있어 우측 사이드바 또는 하단 추천 섹션 둘 중 어디든 공간이 충분하다."
  proposal: "(1) 같은 server+category 의 매물 6개를 useSimilarListings(serverId, categoryId, excludeId) 훅으로 가져와 하단 ListingGrid 로 노출. (2) 같은 판매자의 다른 매물 섹션 추가(useSellerListings(userId, excludeId, limit=4)). 각 카드는 listings 페이지와 동일 ListingCard 재사용. SSE 또는 React Query staleTime 5분."
  effort: medium
  impact: high
  evidence: "Playwright main 안 h2/h3/h4/role=region 검색: 본문 내 섹션 0개(footer 의 '서비스/정보' 만 검출). desktopUsesRightColumn=false, contentMaxWidth=896, mainWidth=1265. API 응답에 serverId/categoryId/author.userId 모두 존재."

- id: imp-0047
  found_at_iter: 3
  area: listing_detail
  type: ux
  target: noise_description_unfit_for_share
  problem: "description 이 'tgewetwewefwef'(14자, 자판 노이즈) 인데 매물 상세 본문 영역의 80% 를 차지하고 og:description 으로도 노출되어, 공유 링크/SEO/사용자 첫인상 모두에서 신뢰가 무너진다. 매물 등록 시 자판 노이즈 검출 / 최소 길이 / 가이드 카피 도움말이 없어 보인다."
  proposal: "(1) 매물 등록 폼(/create)에 description 최소 20자 + 자판 노이즈 정규식 검증을 추가하고, placeholder 에 '예: +10 도리깨 판매합니다. 거래 시 디스코드로 연락 주세요.' 같은 예시 도움말. (2) 노이즈 description 이 노출된 매물은 admin 모더레이션 큐에 자동 제출. (3) 짧은 description 일 때 itemName/serverName/enhancementLevel 로 자동 fallback 보조 카피 생성."
  effort: small
  impact: medium
  evidence: "Playwright + API: description='tgewetwewefwef'(14자), descriptionInfo.isBoilerplate=true(영문 자판 연속). page.tsx L181-183 description 그대로 출력, og:description='무료 커뮤니티 기반 리니지 클래식 아이템 거래 중개'(글로벌, 매물별 미설정)."

- id: imp-0048
  found_at_iter: 3
  area: listing_detail
  type: performance
  target: page_meta_per_listing
  problem: "document.title 이 '기란JT — 리니지 클래식 거래'(글로벌 default) 그대로이고 og:title/og:description/og:image/og:url 도 모두 사이트 전역 값이라 카카오톡/Slack/Twitter 에 매물 링크를 공유하면 똑같은 일반 카드만 노출된다. 매물별 메타데이터가 없어 SEO 인덱스 가치도 0 에 가깝다."
  proposal: "web/app/listings/[id]/page.tsx 를 server component 로 분리(또는 generateMetadata 함수 추가)해서 listing.title, itemName+가격, og:image=listing.images[0].url 또는 iconUrl 절대화 URL, og:type='product', og:url=canonical 을 동적 주입한다. JSON-LD Product schema(name, image, offers.price, offers.priceCurrency='KRW') 도 함께."
  effort: medium
  impact: high
  evidence: "Playwright: document.title='기란JT — 리니지 클래식 거래' (전 페이지 동일), og:title='기란JT — 리니지 클래식 거래 플랫폼', og:image='http://localhost:3000/images/og-image.png' (잘못됨, imp-0016 과 동일), og:url 미설정, JSON-LD count=0. 코드 page.tsx 'use client' 로 metadata 정적."

- id: imp-0049
  found_at_iter: 3
  area: listing_detail
  type: a11y_mobile
  target: time_relative_semantic
  problem: "stats 바의 '1개월 전' 라벨이 단순 <span> 으로 렌더되어 (1) 스크린리더가 정확한 등록 일자를 읽지 못하고, (2) hover 툴팁으로 정확한 일자가 안 보이며, (3) SEO datePublished 가 없다. 또한 '1개월 전' 자체가 매물의 신선도 시그널인데 정확한 일자가 없어 사용자가 노쇠한 매물인지 판단 어렵다."
  proposal: "page.tsx L117 의 formatTimeAgo 호출을 <time dateTime={l.createdAt} title={fullDateFmt(l.createdAt)}>{formatTimeAgo(l.createdAt)}</time> 로 감싼다. 1개월 이상 된 매물에는 '1개월 전 (활동 없음)' 같은 보조 시그널을 lastActivityAt 기반으로 추가."
  effort: trivial
  impact: low
  evidence: "Playwright: timeInfo={text:'1개월 전', tag:'SPAN', hasTimeTag:false, hasDateTime:null, title:null}. API: createdAt='2026-03-21T01:58:50Z', lastActivityAt='2026-03-21T01:58:50Z'(동일=한 달간 활동 0건)."

- id: imp-0050
  found_at_iter: 3
  area: listing_detail
  type: feature
  target: price_negotiation_signal
  problem: "API priceType='fixed'/'negotiable' 가 있고 코드 L174-176 에 협상 가능 분기가 존재하나, fixed 인 경우 '정찰가/협의 불가' 같은 명시적 시그널이 없다. 사용자는 클릭 후 채팅 시작해야 협상 가능 여부를 알 수 있다. 또한 negotiable 매물은 입찰/제시가 입력 폼이 없어 결국 모든 협상이 채팅으로만 일어남."
  proposal: "(1) priceType='fixed' 일 때 가격 옆에 '정찰' 칩, 'negotiable' 일 때 '협의 가능' 칩을 명시적으로 노출. (2) negotiable 인 경우 '제시 가격 입력 → 채팅 자동 생성' 빠른 액션 추가(handleChat 의 createChat mutate 에 initialMessage 옵션). (3) 비슷한 매물의 평균/중앙값 가격 ('이 카테고리 평균 18만원') 표시로 협상 기준 제공."
  effort: medium
  impact: medium
  evidence: "Playwright body 텍스트 검색: '협의|네고|가격 조정|할인 가능|정찰' 0건. API priceType='fixed', priceAmount=200000. 코드 page.tsx L174 negotiable 분기는 있으나 fixed UI 분기 없음."

- id: imp-0051
  found_at_iter: 3
  area: listing_detail
  type: a11y_mobile
  target: status_badge_size_mobile
  problem: "375x667 모바일에서 '판매', '판매중' 배지가 font-size 11px / height 20.5px 로 매우 작아 한손 터치 또는 시각 장애 사용자가 식별하기 어렵다. WCAG 2.5.5 권장 44x44, 최소 32x32 모두 미달. 또한 두 배지가 가로 정렬이라 같은 정보 라인을 차지해 시각적 위계가 약하다."
  proposal: "배지의 padding-y 를 4-6px, font-size 를 13-14px 로 키운다. 모바일에서 '판매중'(상태)는 색상으로만 구분되는 게 아니라 '판매 가능 · 즉시 거래' 같은 문구로 보강하거나 우측에 작은 dot 인디케이터 추가. ui/badge 컴포넌트의 sm/md/lg variant 추가."
  effort: trivial
  impact: low
  evidence: "Playwright 375x667: 판매 배지 36.25x20.5 11px, 판매중 배지 46.375x20.5 11px, 둘 다 padding=2px 8px. WCAG hit-target 추천 32x32 미달."

- id: imp-0052
  found_at_iter: 3
  area: listing_detail
  type: ux
  target: redundant_trade_method_display
  problem: "거래 방식 '무관' 이 두 곳에 동시 표시된다: (1) 페이지 헤더 우측 상단(L102-105 작은 회색 텍스트), (2) 거래 정보 dl 카드 첫 행. 같은 정보가 시각적으로 분리된 두 위치에 있어 사용자가 두 번 읽어야 하고, 헤더 우측의 '무관' 은 컨텍스트 없이 떠 있어 의미 파악이 어렵다."
  proposal: "헤더의 tradeMethod 텍스트(L103) 를 제거하고 거래 정보 카드에 단일화하거나, 반대로 헤더에 라벨 형식('거래: 무관')으로 명확하게 두고 카드에서 제거한다. 두 위치 다 유지하려면 헤더는 우상단 floating chip 으로 격상('🤝 무관')."
  effort: trivial
  impact: low
  evidence: "Playwright: '무관' 텍스트 2회 등장 — (1) 우상단 ml-auto span, color=rgb(163,134,118), 12px (page.tsx L103), (2) dl 첫 행 dt='거래 방식' dd='무관' (page.tsx L189). 시각적으로 같은 화면에 동시 노출."

- id: imp-0053
  found_at_iter: 3
  area: listing_detail
  type: ux
  target: desktop_layout_wasted_gutter
  problem: "1280x800 데스크탑에서 main 이 1265px 인데 콘텐츠는 max-w-4xl(896px)로 가운데 정렬되어 좌우 184px+ 씩 비어 있다. 하나뿐인 컬럼을 위에서 아래로 스크롤하는 구조라 데스크탑의 가로 화면 이점을 전혀 활용하지 못하고, 본문 길이도 짧아 페이지 하단까지 스크롤할 동기가 없다."
  proposal: "lg: 이상에서 2 컬럼 레이아웃으로 전환(왼쪽 60%: 이미지·제목·설명, 오른쪽 40%: 가격·액션·판매자·거래정보 카드, sticky top). lg:grid lg:grid-cols-[1fr_320px] lg:gap-8 패턴. 또는 본문 컬럼은 그대로 두고 우측에 '비슷한 매물' 사이드바(imp-0046 시너지)."
  effort: medium
  impact: medium
  evidence: "Playwright 1280x800: main width=1265, .max-w-4xl=896, 좌우 gutter=(1265-896)/2=184.5px. desktopUsesRightColumn=false (aside 없음, grid-cols 분할 없음). 페이지 scrollHeight=954, viewport 800 → 1.2 화면 분량."

- id: imp-0054
  found_at_iter: 3
  area: listing_detail
  type: content
  target: empty_optional_fields_no_indication
  problem: "API 의 preferredMeetingAreaText, availableTimeText, optionsText 가 모두 null 인데 코드(L166-168, L190-194)는 truthy check 후 미렌더하므로 사용자는 '이 매물의 거래 시간/접선 장소/옵션이 없다' 인지 '입력자가 빠뜨렸다' 인지 알 수 없다. 거래 의사결정 시 이런 정보가 거의 핵심이다."
  proposal: "필수가 아닌 필드라도 비어있으면 '거래 가능 시간 미입력 — 채팅으로 문의' 같은 dim 톤 placeholder 행을 노출하고, '문의' 버튼을 인라인으로 둔다. 또는 InfoRow 에 emptyFallback prop 추가해 ('판매자에게 문의하기' 링크) 일관 처리."
  effort: small
  impact: medium
  evidence: "API 응답: preferredMeetingAreaText=null, availableTimeText=null, optionsText=null. Playwright dl 안 항목 = ['거래 방식','수량'] 2 row 만 표시(다른 4 필드 모두 hidden). 사용자가 누락 여부 판단 불가."

- id: imp-0055
  found_at_iter: 3
  area: listing_detail
  type: a11y_mobile
  target: report_button_anonymous_unreachable
  problem: "신고 버튼이 액션 바 안(page.tsx L263)에만 존재하며, 액션 바는 actions.length>0 이어야 렌더된다. 비로그인 사용자가 노이즈/사기성 매물을 발견해도 신고할 진입점이 페이지 어디에도 없어 어뷰즈 모더레이션이 익명 사용자에게는 동작하지 않는다."
  proposal: "신고 버튼은 인증 게이팅 없이 항상 렌더하고, 신고 모달의 reporter_id 가 null 이면 '신고 사유와 함께 이메일/디스코드 ID 를 남기실 수 있나요?' 익명 신고 폼으로 분기한다. 또는 페이지 우상단(h1 옆)에 ⋯ 메뉴를 두고 '신고' / '공유' 를 분리해 액션 바와 무관하게 항상 접근 가능하게 만든다."
  effort: small
  impact: high
  evidence: "Playwright 비로그인 375x667: '신고' / '차단' 텍스트 button 0개. 코드 page.tsx L263 신고 버튼은 L207 actions.length>0 가드 안에만 위치. 익명 응답 availableActions=null → 액션 바 미렌더."

- id: imp-0056
  found_at_iter: 3
  area: listing_detail
  type: feature
  target: copy_link_share_anonymous
  problem: "공유 버튼(handleShare)도 액션 바 안에 있어 익명 사용자에게 보이지 않는다. 그러나 navigator.share / navigator.clipboard 동작에 인증이 필요 없어 공유는 익명 사용자에게도 가장 자연스러운 first action 인데 진입점이 없다."
  proposal: "공유 버튼을 액션 바와 별개로 페이지 헤더(h1 옆 또는 stats 바 우측)에 항상 노출. clipboard 복사 후 toast 'URL 이 복사되었습니다' 노출(이미 구현). 모바일은 navigator.share 우선, 데스크탑은 clipboard fallback."
  effort: trivial
  impact: medium
  evidence: "Playwright 비로그인: button[aria-label=공유] 0개, '공유' 텍스트 0개. 코드 handleShare(L74-94) 는 인증 무관하게 동작 가능하나 L257 트리거 button 이 액션 바 안."

- id: imp-0057
  found_at_iter: 3
  area: listing_detail
  type: ux
  target: back_to_listings_navigation
  problem: "매물 상세에서 마켓(목록)으로 돌아가는 명시적 백 버튼/breadcrumb 이 없다. 사용자는 브라우저 백 또는 하단 탭의 '마켓' 아이콘에 의존해야 하며, 둘 다 이전 필터/스크롤 위치를 보존하지 않는다(imp-0029 와 연동)."
  proposal: "h1 위에 breadcrumb '마켓 > 질리언 > 둔기 > 도리깨 +10' 또는 좌상단 '←' 백 버튼을 router.back() 으로 추가. breadcrumb 의 각 칩은 해당 필터 조건 URL 로 이동(server/category 칩). aria-label='이전 페이지로 돌아가기'."
  effort: small
  impact: medium
  evidence: "Playwright: main [aria-label*=뒤로], a[aria-label*=뒤로], [class*=breadcrumb] 0개. hasBackButton=false. 코드 page.tsx 에 router.back 호출 0건, Breadcrumb 컴포넌트 미존재."

- id: imp-0058
  found_at_iter: 3
  area: listing_detail
  type: performance
  target: hero_icon_image_priority
  problem: "본문의 item icon Image(64x64)가 above-the-fold 영역에 있는데 loading='lazy', fetchPriority='auto' 로 설정되어 LCP 후보임에도 브라우저 우선순위가 낮다. 데스크탑/모바일 모두 페이지 진입 직후 시각적 첫 인상에 영향을 준다."
  proposal: "page.tsx L148-155 의 Image 에 priority prop 또는 fetchPriority='high' loading='eager' 부여. images[0] 갤러리 이미지에도 동일 처리. unoptimized 플래그가 있어도 loading 속성은 적용된다."
  effort: trivial
  impact: low
  evidence: "Playwright: main img(item icon) src=/static/icons/42.png, naturalW=64, loading='lazy', fetchPriority='auto', rect.y=206.5(viewport 667 의 31%). 코드 page.tsx L148-155 Image 에 priority/loading 속성 미설정."

- id: imp-0059
  found_at_iter: 3
  area: listing_detail
  type: feature
  target: enhancement_level_visual_emphasis
  problem: "리니지 클래식의 핵심 가치 차원인 강화 +10 이 본문에서 'text-gold font-bold ml-2'(L160) 형식의 작은 인라인 텍스트로만 표시된다. itemName 옆에 살짝 노출되지만 시각 위계상 가격(48px font-display)과 비교하면 거의 보이지 않는다. 같은 도리깨여도 +0 과 +10 은 시장가가 수십 배 차이나는데 지금은 강조가 부족하다."
  proposal: "강화 레벨을 별도 큰 칩(48x48 원형, 골드 그라디언트, +10 표기)으로 디자인하거나, h1 제목 줄에 '도리깨 [+10]' 처럼 시각적 강조 칩으로 격상. 또는 가격 옆에 '+10 강화 매물' 라벨 추가. enhancementLevel >=7 일 때 골드, >=5 일 때 cream, 0 일 때 dim 색상."
  effort: small
  impact: medium
  evidence: "Playwright: '+10' 텍스트가 main h1 도 main h2 도 아닌 L160 span(text-gold font-bold ml-2 inline) 으로 itemName='도리깨' 우측에 위치. h1='도리깨 +10 판매합니다' 안에는 +10 이 텍스트로 들어가지만 별도 시각 강조 없음."

- id: imp-0060
  found_at_iter: 3
  area: listing_detail
  type: a11y_mobile
  target: tab_navigation_dead_end
  problem: "비로그인 사용자가 키보드로 탐색하면 main 영역의 tabbable 요소가 단 1개(판매자 프로필 링크)뿐이라, 본문 정보 어디에서도 키보드 사용자가 행동할 수 없다. 액션 바 미렌더 + 이미지 비클릭 + 신고/공유 버튼 부재가 합쳐진 결과."
  proposal: "imp-0042/imp-0045/imp-0055/imp-0056 와 함께 해결되어야 하는 메타 이슈. 최소 보장으로 (1) 페이지 진입 시 h1 에 tabIndex='-1' 부여 + skip-to-content 가 main 안 첫 헤딩에 포커스, (2) 액션 바 익명 노출, (3) 이미지 클릭/포커스 가능. WAI-ARIA APG dialog/disclosure 패턴 적용."
  effort: small
  impact: medium
  evidence: "Playwright 비로그인 375x667 main 안 tabbable: [{tag:'A', text:'유유저_da1e387a거래 0회 · newcomer리뷰 보기 ›', tabIndex:0}] 1개. h1 tabIndex=-1 미설정, button 0개, image 비클릭."

- id: imp-0061
  found_at_iter: 4
  area: listing_create
  type: ux
  target: listings_new_404_handling
  problem: "사용자가 자연스럽게 추측한 URL https://giranjt.com/listings/new 이 동적 라우트 [id]=new 로 매칭되어 GET /api/v1/listings/new 가 실패하고 '⚠️ 매물을 불러올 수 없습니다 / 네트워크 연결을 확인해주세요 / 다시 시도' 화면이 노출된다. 실제로는 매물 등록의 잘못된 진입 시도인데 '네트워크 오류'로 안내해 사용자가 원인을 파악할 수 없다."
  proposal: "(1) web/app/listings/new/page.tsx 를 추가해 /create 로 308 리다이렉트하거나, (2) next.config.ts redirects 에 source:'/listings/new' destination:'/create' permanent:true 를 추가한다. (3) /listings/[id]/page.tsx 의 ID validator 가 'new' 같은 reserved word 를 감지하면 /create 로 라우팅. (4) 네트워크 에러 화면을 '잘못된 매물 ID 또는 삭제된 매물입니다' 와 분기."
  effort: trivial
  impact: high
  evidence: "Playwright + curl: GET /listings/new 200 OK 로 [id]=new dynamic route 매칭, 본문에 '매물을 불러올 수 없습니다 네트워크 연결을 확인해주세요 다시 시도'. web/app/listings/ 디렉토리에 new/ 또는 new.tsx 미존재, [id]/page.tsx 만 존재. is404=true(에러 마커 검출)."

- id: imp-0062
  found_at_iter: 4
  area: listing_create
  type: ux
  target: login_gate_context_message
  problem: "비로그인 사용자가 /create 직접 진입(또는 FAB '+ 매물 등록' 클릭) 시 /login 으로 보내지지만, 로그인 페이지 main 의 카피는 '리니지 클래식 거래 플랫폼' + '로그인 없이 둘러보기' 단 2줄로 일반 로그인 화면과 100% 동일하다. 사용자가 '왜 여기로 왔는지', '로그인하면 어디로 돌아가는지' 알 수 없어 의도와 신뢰가 끊긴다. 또한 redirect 쿼리스트링이 직접 URL 진입(server-side)에서는 보존되지 않아 'Direct hard reload' 시나리오에서 redirect=/create 가 사라진다."
  proposal: "(1) login 페이지 상단에 redirect 파라미터 기반 컨텍스트 카피 노출: '매물 등록을 위해 로그인이 필요해요. Google 계정 1초 로그인 후 바로 등록 화면으로 돌아갑니다.' (2) /create page.tsx 의 useEffect 에서 router.push 대신 router.replace(`/login?redirect=${encodeURIComponent(pathname)}`) 를 호출해 redirect param 을 항상 보존하고 history pollution 방지. (3) middleware.ts 에서 server-side 401 redirect 시에도 redirect 쿼리 부착."
  effort: small
  impact: high
  evidence: "Playwright: 직접 https://giranjt.com/create → 최종 URL https://giranjt.com/login (redirect param 없음, redirectParam=null). 모바일 FAB 클릭 시는 /login?redirect=%2Fcreate 로 정상 보존. login main text='리니지 클래식 거래 플랫폼로그인 없이 둘러보기' (redirect 컨텍스트 0). use-auth-guard.ts L20-21: pathname encode + push, 직접 hard load 시는 client-side 라 toast 만 뜨고 사라짐."

- id: imp-0063
  found_at_iter: 4
  area: listing_create
  type: ux
  target: redundant_register_entry_points
  problem: "데스크탑 1280px 에서 매물 등록 진입점이 (1) 헤더 nav '매물 등록' 링크 (292,15), (2) 페이지 우상단 '+ 매물 등록' 버튼 (1151,273), (3) 푸터 '서비스/매물 등록' (970,1025) 3곳에 동시 존재한다. 모바일 375px 에서는 (1) 우하단 FAB '+'(288,531), (2) 하단 탭 '등록'(180,613) 2곳 — 둘 다 비슷한 시간에 화면에 보이고(거리 25.5px) 어느 것이 1차 경로인지 모호하다. 모든 경로가 비로그인 시 동일한 /login (메시지 차별화 없음)으로 보내지므로 정보 노이즈."
  proposal: "(1) 데스크탑은 헤더 nav '매물 등록' 만 1차 경로로 두고, 페이지 우상단 '+ 매물 등록' 버튼은 (a) 비로그인일 땐 hidden, (b) 로그인 시에만 노출. 푸터는 '내 매물' 같이 차별화. (2) 모바일은 FAB 또는 하단 탭 둘 중 하나만. FAB 유지 시 하단 탭 '등록' 자리를 '내 거래' 로 재할당해 거래 활동 중심. (3) FAB '+' 글자에 'aria-label=매물 등록' 외에도 호버/포커스 툴팁 '매물 등록하기' 추가."
  effort: small
  impact: medium
  evidence: "Playwright 1280x800: a[href='/create'] 5개(visible 3개), 375x667: visible 2개(FAB 56x56 at 288,531; 하단탭 90x55 at 180,613, fabBottom→navTop 거리 25.5px). FAB textContent='+', hasIcon=false (SVG 아이콘 없이 단순 '+' 글자, font-size 32px)."

- id: imp-0064
  found_at_iter: 4
  area: listing_create
  type: feature
  target: anonymous_inline_form_preview
  problem: "비로그인 사용자가 '+ 매물 등록' 클릭 시 즉시 /login 으로 보내져, 등록 화면이 어떻게 생겼는지 / 어떤 정보가 필요한지 / 양식을 채우는 데 얼마나 걸리는지 미리 알 수 없다. 이는 신규 사용자 conversion 의 큰 마찰점이다(가입 후 양식 보고 '생각보다 길다' 이탈)."
  proposal: "비로그인 상태에서도 /create 페이지를 'preview 모드'로 렌더한다. 입력 가능하지만 (1) 모든 input/select 가 disabled 또는 read-only, (2) 상단에 '미리보기 — 등록은 로그인 후 가능합니다' 배너, (3) '등록하기' 버튼이 'Google 로그인 후 등록하기' 로 라벨 변경. 양식을 sessionStorage 에 임시 저장해 로그인 후 자동 복원."
  effort: medium
  impact: high
  evidence: "코드 web/app/create/page.tsx L43-47: useEffect 에서 isLoggedIn 가드, L60 if (!isLoggedIn) return null. 비로그인 사용자는 페이지 콘텐츠를 0% 본 채 /login 으로 즉시 redirect."

- id: imp-0065
  found_at_iter: 4
  area: listing_create
  type: feature
  target: form_draft_autosave
  problem: "매물 등록 폼은 useState 8개 필드(listingType, serverId, categoryId, itemName, title, description, priceType, priceAmount, tradeMethod) + images 배열 + selectedItem + enhancementLevel 을 메모리에만 보관한다. 사용자가 실수로 뒤로가기/탭 닫기/세션 만료 시 입력한 모든 정보가 사라진다. 이미지 업로드(최대 5장 × 10MB)는 다시 업로드해야 한다."
  proposal: "useEffect 에서 form/images/selectedItem/enhancementLevel 변경 시마다 디바운스 1초 후 localStorage 'create_listing_draft' 에 저장. 페이지 마운트 시 draft 가 있으면 '이전 작성 중이던 내용이 있어요. 이어서 작성하시겠어요? [예/지우기]' 토스트 액션. 등록 성공 또는 명시적 '폼 지우기' 시 draft 삭제. images 는 imageId 만 보관(서버 업로드 결과)."
  effort: small
  impact: high
  evidence: "코드 web/app/create/page.tsx L26-41: useState 만 사용, localStorage/sessionStorage 호출 0건. handleSubmit L88 에서 mutate 후 router.push('/'), 실패 시 toast 만 뜨고 form 데이터 그대로 유지(이건 정상)지만 앱/브라우저 종료 시 손실."

- id: imp-0066
  found_at_iter: 4
  area: listing_create
  type: feature
  target: price_market_reference
  problem: "가격 입력 시 '이 아이템의 평균/중앙값 시세' 정보가 없어 신규 사용자가 너무 비싸거나 너무 싸게 책정할 위험이 크다. 결국 거래 매칭 효율과 사용자 만족도가 떨어진다. priceType='offer'(제안받음) 도 가격 옆에 disabled input 만 보일 뿐 가이드가 없다."
  proposal: "selectedItem + enhancementLevel + serverId 가 채워지면 priceAmount 입력 옆에 '동일 조건 시세 200,000원 ~ 350,000원 (평균 270,000원, 최근 거래 12건)' 노출. API 추가: GET /api/v1/listings/price-stats?itemId=&enchant=&serverId=. 데이터 부족 시 '이 조건 거래 데이터가 아직 부족해요' 안내."
  effort: medium
  impact: high
  evidence: "코드 web/app/create/page.tsx L196-207: priceAmount input 단독, 시세/가이드 컴포넌트 없음. 같은 페이지에 selectedItem.optionText, categoryPath 는 표시되나 시세 시그널 0개."

- id: imp-0067
  found_at_iter: 4
  area: listing_create
  type: feature
  target: meeting_area_time_input
  problem: "매물 상세에서 노출되는 '거래 가능 시간 / 접선 장소 / 옵션' 필드(API: preferredMeetingAreaText, availableTimeText, optionsText) 가 등록 폼에 입력 UI 가 없다. 결국 모든 매물의 해당 필드가 null 이 되고(imp-0054 와 직결), 거래 의사결정에 필수 정보가 채팅으로만 전달돼 효율이 떨어진다."
  proposal: "거래 섹션에 (1) '거래 가능 시간' textarea (placeholder='예: 평일 저녁 7-10시, 주말 종일'), (2) '접선 장소(오프라인)' input (tradeMethod 가 in_game 이 아닐 때만 노출), (3) '아이템 옵션' textarea (selectedItem.optionText 를 기본값으로 prefill, 사용자가 추가 옵션 적을 수 있게). 모두 optional 이지만 미입력 시 '미입력 항목이 있어요. 채팅 응답이 늦어질 수 있어요' 안내."
  effort: small
  impact: high
  evidence: "코드 web/app/create/page.tsx form state: listingType/serverId/categoryId/itemName/title/description/priceType/priceAmount/tradeMethod 9개. preferredMeetingAreaText/availableTimeText/optionsText 입력 UI 0개. handleSubmit L90-102 의 data payload 에도 미포함. 결과로 imp-0054 (상세 페이지에 거래시간/장소 모두 null) 발생."

- id: imp-0068
  found_at_iter: 4
  area: listing_create
  type: ux
  target: server_select_no_recent
  problem: "서버 선택이 native <select> 한 개로 28개+ 옵션을 알파벳/추가 순서로 노출한다. 사용자는 보통 자기 메인 서버 한두 개만 거래하므로 매번 같은 서버를 찾기 위해 드롭다운을 끝까지 스크롤한다. 또한 native select 는 모바일에서 OS 기본 휠/리스트 UI 라 한국어 검색이 불편하다."
  proposal: "(1) 최근 등록 시 사용한 서버를 localStorage 'last_servers' 에 보관(최대 3개), 드롭다운 상단 'optgroup label=최근' 으로 prepend. (2) 또는 서버 선택을 칩 그룹(홈과 일관)으로 변경 + 검색 input 결합. (3) 단일 서버 사용자(>80%)를 위해 선택 후 다음 등록부터 default 자동 적용."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx L159-173: <select id=serverId> + servers.map(<option>). localStorage 호출 0건. 사용자별 server preference state/cookie 미존재."

- id: imp-0069
  found_at_iter: 4
  area: listing_create
  type: ux
  target: title_autogen_lock_unclear
  problem: "제목이 selectedItem 변경 시 자동 생성되는데(L50-58), 사용자가 한 번이라도 직접 수정하면 titleAutoGenerated=false 가 되어 이후 자동생성이 멈춘다. 그러나 사용자에게 이 잠금 상태를 표시하는 시그널이 없다. 다시 자동생성으로 돌아갈 방법(reset 버튼)도 없어 itemName 을 바꿔도 제목이 그대로라 혼란."
  proposal: "input 우측에 '🔄 자동 생성' 토글 칩을 두고 titleAutoGenerated=true 일 때 활성, 사용자가 직접 수정하면 비활성으로 자동 변경. 사용자가 다시 칩을 누르면 즉시 selectedItem+enhancementLevel 기반으로 재생성. placeholder 도 ('아이템을 선택하면 자동 생성됩니다' → '아이템 선택 시 자동 생성. 직접 수정 시 잠금')."
  effort: trivial
  impact: low
  evidence: "코드 web/app/create/page.tsx L28: titleAutoGenerated state. L222-223 onChange 에서 setTitleAutoGenerated(false), 다시 true 로 만드는 UI 0건. placeholder='아이템을 선택하면 자동 생성됩니다'(L224) 만 표시."

- id: imp-0070
  found_at_iter: 4
  area: listing_create
  type: a11y_mobile
  target: fab_no_label_visual
  problem: "모바일 FAB 이 단순 '+' 글자(font-size 32px) 만 노출하고 SVG 아이콘이나 텍스트 라벨이 없다. aria-label='매물 등록' 은 있으나 시각적으로 '+' 만 보여 사용자가 '뭘 더하는 건지' 추측해야 한다. 56x56 hit-target 은 OK 지만 '+' 글자는 컨텍스트 약함."
  proposal: "(1) FAB 에 SVG plus 아이콘 + tooltip 'title=매물 등록' 부착, (2) 또는 expandable FAB 패턴: 평상시 56x56 '+', 호버/롱프레스 시 'plus 매물 등록' label 이 좌측으로 슬라이드 노출. (3) 첫 방문자에게는 1회 'fade-in' 방식 'tap to register' 라벨 노출(localStorage 'fab_seen' 으로 한 번만)."
  effort: trivial
  impact: low
  evidence: "코드 web/app/page.tsx L158-164: <Link href=/create aria-label='매물 등록' class='...text-2xl...'>+</Link>. textContent='+', hasIcon=false(SVG/img 없음), title null."

- id: imp-0071
  found_at_iter: 4
  area: listing_create
  type: a11y_mobile
  target: form_field_validation_inline
  problem: "폼 필드는 native required + minLength(title>=2, description>=10) 를 사용하지만 (1) 검증 실패 시 브라우저 기본 popup 만 노출되고, (2) 한국어 OS 에서는 '이 입력란을 작성하세요' 같은 직역 메시지, (3) 시각적 에러 상태(red border, error text) 가 없어 어느 필드를 봐야 할지 모호하다. priceAmount 도 number type 인데 음수/소수 검증이 클라이언트에 없다."
  proposal: "(1) react-hook-form + zod 또는 useState validation 으로 모든 필드의 한국어 에러 메시지 inline 노출 (예: '제목을 2자 이상 적어주세요', '설명은 의미 있는 거래 정보를 10자 이상 적어주세요'). (2) Submit 시 첫 invalid 필드로 자동 스크롤 + focus. (3) priceAmount 양의 정수만, 최소 100원 이상."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx: required, minLength={2}, minLength={10}, type='number' 만 사용. errorState/errorMessage/aria-invalid/aria-describedby 0개. handleSubmit L88-109 에서 try/catch 후 generic '등록에 실패했습니다' 토스트만."

- id: imp-0072
  found_at_iter: 4
  area: listing_create
  type: performance
  target: form_submit_no_progress
  problem: "createListing.isPending=true 일 때 submit 버튼이 disabled + 라벨 '등록 중...' 으로 바뀌지만, 이미지 5장(최대 50MB) 을 업로드한 매물 등록 시에도 같은 1초 미만 텍스트만 표시된다. 사용자는 진행률을 볼 수 없어 느린 네트워크에서 '멈춘 건가' 의심하고 다시 클릭(이미 disabled 라 안 눌리지만 인지 부하)."
  proposal: "(1) submit 시 fullscreen overlay 'spinner + 매물 등록 중...' (aria-live=polite), (2) 이미지 업로드는 ImageUpload 컴포넌트가 이미 처리하지만 등록 단계 자체에서 추가 진행률을 표시. (3) 5초 이상 걸리면 '응답이 늦어요. 네트워크를 확인해주세요' 안내. (4) 성공 시 router.push 전 짧은 success 토스트 또는 confetti 마이크로 인터랙션."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx L271-277: button disabled={createListing.isPending}, 라벨 '등록 중...' 단 1줄. progress bar/spinner overlay 0개. handleSubmit 의 await mutateAsync 동안 나머지 form 은 그대로 표시(다른 필드 수정 가능 → 데이터 정합성 잠재 이슈)."

- id: imp-0073
  found_at_iter: 4
  area: listing_create
  type: content
  target: description_placeholder_guidance
  problem: "description textarea 의 placeholder 가 '아이템을 선택하면 옵션이 자동 입력됩니다' 단 1줄이라, 사용자에게 어떤 내용을 적어야 거래가 잘 성사되는지 가이드가 없다. 결과적으로 imp-0047 의 'tgewetwewefwef' 같은 자판 노이즈 description 이 양산된다."
  proposal: "placeholder 를 다단계 예시로 풍부화: '예시:\\n[아이템 옵션]\\n공격력 +5, 명중 +1\\n\\n[거래 조건]\\n질리언 서버 인게임 거래만, 평일 저녁 시세 협상 가능. 디스코드 ID: example#1234'. 또는 placeholder 옆 '거래 잘 되는 설명 예시 보기 ▾' disclosure 로 모범 답안 노출. 입력 글자수 카운터(0/2000) 함께."
  effort: trivial
  impact: medium
  evidence: "코드 web/app/create/page.tsx L237-245: textarea placeholder='아이템을 선택하면 옵션이 자동 입력됩니다', minLength=10. 글자수 카운터/예시 disclosure/모범 카피 0개. imp-0047 의 description='tgewetwewefwef' 같은 자판 노이즈가 첫 페이지 노출 매물 4개 중 1개에 이미 존재."

- id: imp-0074
  found_at_iter: 4
  area: listing_create
  type: feature
  target: image_upload_camera_capture
  problem: "이미지 업로드 input 이 type='file' multiple 만 가지고 있어 모바일에서도 무조건 사진첩(또는 일반 파일 선택) 만 열린다. 거래용 인증샷(예: '인벤토리 보유 증거')은 즉석 촬영이 자연스러운데 'capture' 속성 부재로 카메라 직접 진입이 막혔다. 또한 이미지 순서 변경(드래그) 도 없어 첫 이미지(='대표') 를 다시 업로드해야 바꿀 수 있다."
  proposal: "(1) input 에 capture='environment' 추가 또는 별도 '📷 촬영' 버튼 분리(모바일 only). (2) 이미지 그리드에 react-dnd 또는 @dnd-kit 으로 드래그 순서 변경, 첫 번째가 자동 '대표'. (3) 각 이미지 우상단에 '대표로 지정' 버튼."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/image-upload.tsx L98-103: input type='file' accept multiple 만 존재. capture 속성 0건. handleRemove(L84) 만 있고 reorder 함수 0개. 그리드 L119: grid-cols-5 정적 노출, draggable 속성 0개."

- id: imp-0075
  found_at_iter: 4
  area: listing_create
  type: a11y_mobile
  target: image_upload_drop_zone_keyboard
  problem: "이미지 업로드 드롭존(div.border-dashed)이 onClick 핸들러로 inputRef 를 클릭하지만 키보드 접근이 불가능하다. div 라 tabIndex 없고 role 없고, Enter/Space 키로 동작하지 않는다. 또한 onDrop 만 있고 ondragenter/leave 시각 피드백(active border color) 도 없어 드롭 가능 영역인지 인식 어렵다."
  proposal: "(1) div 를 button type='button' 으로 변경하거나 role='button' tabIndex='0' + onKeyDown(Enter/Space) 핸들러. aria-label='이미지 업로드, 클릭 또는 드래그'. (2) onDragEnter 시 border-gold 강조 + onDragLeave 복구. (3) input 에 직접 label 연결(htmlFor)도 가능."
  effort: trivial
  impact: medium
  evidence: "코드 web/components/forms/image-upload.tsx L90-95: <div onDrop onDragOver onClick>...</div>. tabIndex/role/aria-label 모두 미설정, onKeyDown 핸들러 0건, onDragEnter/Leave 핸들러 0건."

- id: imp-0076
  found_at_iter: 5
  area: chat
  type: ux
  target: chat_redirect_param_loss_on_direct_load
  problem: "익명 사용자가 https://giranjt.com/chats 또는 /chats/abc-123 으로 직접 진입(외부 링크/북마크/공유)하면 ChatsPage / ChatDetailPage 의 useEffect 가 router.replace('/login') 만 호출하고 redirect 쿼리 파라미터를 부착하지 않아, 로그인 후 사용자는 의도한 채팅이 아닌 / 로 떨어진다. login page.tsx L31-35 는 ?redirect= 를 처리할 준비가 되어 있는데도 호출처에서 누락된다."
  proposal: "(1) web/app/chats/page.tsx L31, web/app/chats/[id]/page.tsx L97 의 router.replace('/login') 을 router.replace(`/login?redirect=${encodeURIComponent(pathname)}`) 로 변경. (2) 또는 useAuthGuard 의 requireAuth 패턴을 useEffect 내부에서도 재사용해 모든 게이팅 페이지에서 redirect 보존을 일관 처리. (3) /create 페이지에도 같은 패턴 점검(imp-0062 와 함께)."
  effort: trivial
  impact: high
  evidence: "Playwright 익명 375x667: https://giranjt.com/chats → 최종 URL https://giranjt.com/login (redirect param=null). https://giranjt.com/chats/abc-123 → 동일 결과. 코드 web/app/chats/page.tsx L31 router.replace('/login'), web/app/chats/[id]/page.tsx L97 동일 패턴. login page.tsx L31-35 는 searchParams.get('redirect') 처리 코드 보유."

- id: imp-0077
  found_at_iter: 5
  area: chat
  type: content
  target: chat_login_gate_context_message
  problem: "익명 사용자가 채팅 페이지로 진입해 /login 에 도달했을 때 로그인 화면 카피가 '리니지 클래식 거래 플랫폼' + '로그인 없이 둘러보기' 단 2줄로 일반 로그인 페이지와 100% 동일하다. 사용자는 '왜 채팅으로 못 갔는지', '로그인하면 채팅으로 돌아가는지' 알 수 없어 toast(2-3초 후 사라짐) 외에는 컨텍스트 시그널이 없다. imp-0062 와 동일한 패턴이 chat 영역에서도 발생."
  proposal: "login 페이지에 redirect 파라미터 패턴별 컨텍스트 카피 추가. /chats 또는 /chats/* 일 때 '판매자와 채팅하려면 로그인이 필요해요. Google 로그인 후 바로 채팅으로 돌아갑니다.' 노출. redirect 매핑 테이블(/create→매물 등록, /chats→채팅, /profile→내 프로필)을 컴포넌트 상수로 두고 분기."
  effort: trivial
  impact: medium
  evidence: "Playwright: https://giranjt.com/chats → /login bodyText='본문으로 건너뛰기마켓채팅매물 등록로그인리니지 클래식 거래 플랫폼로그인 없이 둘러보기...' (chat 컨텍스트 카피 0). 코드 web/app/login/page.tsx L89-115 은 redirect 변수 정의는 있으나 컨텍스트 메시지 분기 없음."

- id: imp-0078
  found_at_iter: 5
  area: chat
  type: feature
  target: anonymous_chat_intent_capture
  problem: "비로그인 사용자가 매물 상세에서 '채팅하기' 버튼을 보지 못한다(actions.length>0 가드, 익명 응답에 availableActions=null). 따라서 '이 판매자에게 문의하고 싶다' 의도가 있어도 1) 이 매물에 채팅 기능이 있는지 자체를 발견 못 하고, 2) 로그인 후 같은 매물로 돌아갈 경로가 끊긴다. 거래 conversion 의 핵심 경로가 막혀 있다."
  proposal: "(1) 매물 상세에 '채팅하기' 버튼을 익명에게도 항상 노출하되, 클릭 시 'Google 로그인 후 채팅 시작하기' 라벨 + onClick 에서 /login?redirect=/listings/<id>?action=chat 으로 이동. (2) 로그인 콜백에서 action=chat 이면 자동으로 createChat → /chats/<id> 진입. (3) 또는 channeling page chat preview: 익명에게는 read-only system message '판매자: 안녕하세요! 로그인 후 메시지를 보내주세요' 만 노출하는 limited UI."
  effort: medium
  impact: high
  evidence: "Playwright 익명 https://giranjt.com/listings/b155f6a7-...: button 0개(액션 바 미렌더). 코드 web/app/listings/[id]/page.tsx L207: actions.length>0 가드, L247 actions.includes('start_chat') 이중 가드. 익명 API 응답 availableActions=null."

- id: imp-0079
  found_at_iter: 5
  area: chat
  type: ux
  target: chat_message_length_no_limit
  problem: "ChatInput textarea 에 maxLength 가 설정되어 있지 않아 사용자가 수만 글자를 입력해도 클라이언트 단에서 차단되지 않는다. 서버 거부(413/422) 시 optimistic 메시지가 'failed' 상태로 남고 재전송 시도해도 같은 에러 반복, 사용자는 어디서 잘라야 할지 모른다. 이미 onSend(text.trim()) 로 공백만 막을 뿐 길이 가드 없음."
  proposal: "(1) ChatInput textarea 에 maxLength={2000} (또는 백엔드 검증 한도와 동일 값) 부여. (2) 입력 글자수가 1800 이상이면 우하단에 'XXX/2000' counter 표시(text-text-secondary, >1950 시 text-danger). (3) 붙여넣기 시 즉시 잘라내고 toast '메시지는 2000자까지 입력할 수 있어요'. (4) 같은 한도를 backend domain validation 과 일치하게 lib/constants.ts 로 분리."
  effort: trivial
  impact: low
  evidence: "코드 web/components/chat/chat-input.tsx L52-64: textarea props={value, onChange, onKeyDown, placeholder, rows, aria-label, className, disabled} — maxLength 미설정. submit() L33-38 의 trimmed.length 가드 0건. lib/constants.ts 또는 도메인 상수 미존재."

- id: imp-0080
  found_at_iter: 5
  area: chat
  type: feature
  target: image_attachment_in_chat
  problem: "채팅 messageType 은 'text' | 'system' | 'reservation_card' 3종뿐이고, 이미지/파일 첨부 진입점이 ChatInput 에 없다. 거래 협상에서는 인벤토리 인증샷, 옵션 캡처, 입금 영수증 같은 이미지 공유가 핵심이다. 현재는 외부 디스코드/카톡으로 우회해야 하므로 plat lock-in 이 약하고 분쟁 시 증거 보존도 어렵다."
  proposal: "(1) ChatInput 에 클립 아이콘 IconButton 추가 → input[type=file accept='image/*' capture='environment'] 트리거. (2) 백엔드에 messageType='image' 추가 + S3/storage upload 후 metadataJson={url,width,height} 형태 저장. (3) ChatMessage 에 image 분기로 next/image 렌더, 클릭 시 lightbox. (4) 단, MVP 수준에서는 이미지 1장/메시지 + 5MB 제한으로 시작."
  effort: large
  impact: high
  evidence: "코드 web/lib/types.ts L83 messageType: 'text' | 'system' | 'reservation_card'. ChatInput L52-73 에는 파일 input 0개. 첨부 버튼/아이콘 0개. 백엔드 send-message 핸들러 'text' messageType 만 검증."

- id: imp-0081
  found_at_iter: 5
  area: chat
  type: ux
  target: chat_list_search_filter
  problem: "채팅 목록(/chats)은 ChatRoom[] 을 단순 시간순으로 나열한다. 거래가 활발한 사용자는 수십~수백 개 채팅방을 갖게 되는데, '활성 매물별', '예약 단계별', '미답장만', '특정 매물 제목' 같은 필터/검색이 없어 원하는 대화를 찾기 어렵다."
  proposal: "(1) 페이지 상단에 검색 input '대화 검색 (상대방, 매물 제목)' + 필터 칩 [전체 / 안 읽음 / 예약 진행 / 거래 완료]. (2) 검색은 chats.filter(c => c.counterparty.nickname.includes(q) || c.listingTitle.includes(q)) 클라이언트 필터. (3) 빈 상태 분기 '검색 결과가 없어요. 다른 키워드로 시도해보세요'."
  effort: small
  impact: medium
  evidence: "코드 web/app/chats/page.tsx: 검색 input 0개, 필터 칩 0개. chats.map 순회만 존재 (L88-94). EmptyState 는 chats.length===0 케이스만 처리."

- id: imp-0082
  found_at_iter: 5
  area: chat
  type: feature
  target: typing_indicator
  problem: "현재 SSE 로 메시지 도착 이벤트만 처리하고, 상대방이 '입력 중'이라는 시그널이 없다. 협상 도중 응답 대기 시간이 길게 느껴지고, 사용자는 상대가 떠났는지 답을 작성 중인지 알 수 없어 추가 메시지를 보내거나 이탈한다."
  proposal: "(1) 백엔드 SSE 이벤트 타입 'typing' 추가, 클라이언트는 onChange debounce 1.5s 동안 'typing' 이벤트 publish, 받는 쪽은 활성 chat 헤더 아래에 '상대방이 입력 중...' indicator(animate-pulse 3 dots). (2) 마지막 typing 이벤트 후 3초 무음이면 indicator 자동 숨김. (3) 모바일은 chat header 가 좁아 ListingInfoCard 우측 또는 메시지 영역 하단에 인라인."
  effort: medium
  impact: medium
  evidence: "코드 web/lib/hooks/use-sse.ts: typing 이벤트 핸들러 없음. ChatPanel L141-155 은 connection status 만 노출. backend internal/event 카탈로그에도 typing 이벤트 미존재(추정). 사용자 활동 시그널은 unreadCount 와 lastMessage 단 2종."

- id: imp-0083
  found_at_iter: 5
  area: chat
  type: a11y_mobile
  target: chat_input_mobile_keyboard_attrs
  problem: "ChatInput textarea 에 enterkeyhint, autocomplete, autocapitalize, spellcheck, inputmode 속성이 없어 모바일 IME 키보드가 'send' 키 라벨도 없고 자동 대문자 변환이 영문 입력 시 어색하게 동작할 수 있다. 또한 Enter 가 '전송' 단축키인데 OS 기본 'return' 키와 시각적 차별화가 없어 학습이 어렵다."
  proposal: "textarea 에 enterkeyhint='send', autocomplete='off', autocapitalize='sentences'(한국어 환경에선 무영향이지만 영어/숫자 혼용 시 도움), spellcheck='false'(거래 도메인 단어가 사전에 없어 빨간 밑줄 노이즈) 추가. placeholder 옆에 작은 hint '⏎ 전송 / Shift+⏎ 줄바꿈' 1회 노출(localStorage).'first_chat_input_hint_shown' 으로 한 번만)."
  effort: trivial
  impact: low
  evidence: "코드 web/components/chat/chat-input.tsx L52-64: textarea props 에 enterkeyhint/autocomplete/autocapitalize/spellcheck/inputmode 0건. Enter→submit 동작은 handleKeyDown L45-50 에 구현되어 있지만 시각적 hint 0건."

- id: imp-0084
  found_at_iter: 5
  area: chat
  type: ux
  target: chat_height_calc_dynamic_viewport
  problem: "ChatDetailPage 의 컨테이너가 h-[calc(100vh-120px)] 로 fixed 120px(헤더+상태표시줄?) 을 가정한다. 모바일 Safari 의 동적 주소창(축소/확장으로 100vh 가 변동) 또는 데스크탑에서 헤더 높이가 64px 인 점을 고려할 때 일부 환경에서 메시지 영역이 짤리거나 ChatInput 이 키보드 위로 올라오면서 bottom 이 잘린다. iOS Safari 의 visualViewport 변화 시 자동 스크롤도 미적용."
  proposal: "(1) Tailwind v4 또는 CSS dvh 사용: h-[calc(100dvh-64px)]. (2) ChatInput 컨테이너에 padding-bottom: env(safe-area-inset-bottom) 부여(iOS 홈 인디케이터 영역). (3) keyboard 등장 시 visualViewport API 로 bottomRef.scrollIntoView 재실행. (4) 데스크탑 Header 64px 와 일치하는 변수 사용."
  effort: small
  impact: medium
  evidence: "코드 web/app/chats/[id]/page.tsx L125: className='flex flex-col h-[calc(100vh-120px)]'. ChatPanel L91 도 비슷한 vh 사용. dvh/svh 없음, env(safe-area-inset-*) 0건, visualViewport 핸들러 0건. Header 실제 높이는 components/layout/header.tsx L22 'h-16'(64px) 로 120 과 불일치."

- id: imp-0085
  found_at_iter: 5
  area: chat
  type: content
  target: chat_empty_state_action_label
  problem: "EmptyState 의 actionLabel='매물 둘러보기' actionHref='/' 가 home 으로 이동시키는데, 사용자 문맥은 '채팅이 없어서 채팅을 시작할 매물을 찾고 싶은 상태'이므로 '/' 보다는 '/listings'(매물 목록 직진) 또는 다른 매물 권유가 더 정확하다. 또한 '채팅이 없습니다' 자체가 사실 통보형 카피로 마찰점."
  proposal: "title='첫 거래 채팅을 시작해보세요', description='관심 매물에서 채팅하기 버튼을 누르면 판매자와 바로 대화할 수 있어요', actionLabel='관심 매물 보기' actionHref='/listings?sort=popular'. 추가로 '⭐ 찜한 매물에서 채팅 시작' 보조 버튼(자기 favorites 가 있을 때만)."
  effort: trivial
  impact: low
  evidence: "코드 web/app/chats/page.tsx L42: <EmptyState title='채팅이 없습니다' description='매물에서 채팅을 시작해보세요' actionLabel='매물 둘러보기' actionHref='/' />. '/' 와 '/listings' 모두 매물 페이지지만 / 는 hero+navigation, /listings 는 그리드 직진."

- id: imp-0086
  found_at_iter: 5
  area: chat
  type: feature
  target: deal_complete_review_prompt
  problem: "거래 완료(chatStatus='deal_completed') 후 자동으로 리뷰 작성 유도가 채팅방 안에 없다. 사용자는 별도 프로필 → 거래 내역 → 리뷰 진입을 해야 하므로 리뷰 작성률이 떨어지고, 이는 newcomer 신뢰 시그널 약화로 이어진다(매물 상세에 '거래 0회 · newcomer' 표기됨)."
  proposal: "(1) chatStatus 가 'deal_completed' 로 전환되면 ChatMessage 위에 system message 카드 '🎉 거래가 완료되었습니다 — 상대방을 평가해주세요'와 [별점 컴팩트 위젯 + '리뷰 쓰기' 버튼]. (2) 리뷰 안 쓴 채 24시간 지나면 알림 발송('아직 리뷰를 안 남기셨어요'). (3) 리뷰 작성 후 카드는 '리뷰 완료 · 5점' 으로 collapse."
  effort: medium
  impact: high
  evidence: "코드 web/app/chats/[id]/page.tsx 와 ChatPanel 모두 reservation 모달은 있으나 review 모달/링크 0개. ChatListItem L29-37 은 deal_completed 라벨만 노출. lib/hooks/use-chats.ts 에는 리뷰 mutation 없음."

- id: imp-0087
  found_at_iter: 5
  area: chat
  type: a11y_mobile
  target: reservation_button_tap_target
  problem: "ChatDetailPage 상단의 '예약 제안'(L130-135) / '신고'(L136-141) 버튼이 px-3 py-1.5 + text-xs 로 약 28x28~32x32 크기로 추정된다. WCAG 2.5.5(44x44 권장, 최소 32x32)와 Apple HIG 44pt 미달 가능성 높음. 모바일에서 두 버튼이 가로로 붙어 있어 오탭 시 신고 모달이 잘못 열릴 수 있다."
  proposal: "(1) 두 버튼 모두 min-h-[44px] px-4 py-2 text-sm 으로 격상. (2) '예약 제안' 은 primary 액션, '신고' 는 dropdown ⋯ 메뉴 안으로 이동해 시각/터치 분리. (3) 또는 신고 버튼을 chat header 우측 ⋯ 메뉴로 이동시켜 actionable + dangerous 액션을 분리(휴대폰 한손 reach 영역 고려)."
  effort: trivial
  impact: medium
  evidence: "코드 web/app/chats/[id]/page.tsx L130-141: className='px-3 py-1.5 text-xs ...'. computed height ≈ 28px(text-xs ≈ 12px line-height + py 12px). gap-2 (L129) 로 두 버튼 간 8px 간격."

- id: imp-0088
  found_at_iter: 5
  area: chat
  type: performance
  target: chat_list_polling_when_unfocused
  problem: "useMessages 가 sseConnected=false 시 refetchInterval 5000ms 로 폴링하는데, 탭이 background 상태(document.visibilityState='hidden')일 때도 동일 주기로 호출된다. 다중 채팅방을 띄워둔 사용자가 다른 일을 하는 동안 불필요한 네트워크/배터리/NAS CPU 소비. 또한 useChats(채팅 목록) 도 refetch 정책이 query default 라 stale time 동안 같은 호출 반복 가능."
  proposal: "(1) useMessages 에 refetchIntervalInBackground: false 명시. (2) document.visibilityState 변경 시 refetchInterval 동적 조정(visible 5s, hidden 30s 또는 0). (3) useChats 에 staleTime 30s 부여하고 SSE event 'chat_updated' 수신 시 invalidate. (4) page.tsx 가 unmount 되면 enabled=false 로 폴링 중단(이미 됨)."
  effort: small
  impact: low
  evidence: "코드 web/lib/hooks/use-chats.ts L23: refetchInterval: sseConnected ? false : 5_000. refetchIntervalInBackground 미설정(기본 false 지만 명시 권장). visibilitychange 핸들러 0건. useChats L7-12 는 staleTime/gcTime 미설정."

- id: imp-0089
  found_at_iter: 5
  area: chat
  type: ux
  target: failed_message_persistence
  problem: "useSendMessage onError 에서 optimistic 메시지를 status='failed' 로 표시하지만, 페이지 새로고침/탭 닫기 시 React Query cache 가 사라지고 failed 메시지도 함께 사라진다. 사용자는 '내가 보낸 줄 알았는데' 인지 부조화에 빠지고, 분쟁 시 발신 시도 자체를 증명할 수 없다."
  proposal: "(1) onError 시 localStorage 'chat_failed_<chatId>' 에 {clientMessageId, text, sentAt} 직렬화. (2) 채팅방 진입 시 localStorage 의 failed 메시지를 cache 에 prepend(중복 clientMessageId 제거). (3) 재전송 성공 시 또는 사용자가 '실패 메시지 삭제' 클릭 시 localStorage 에서 제거. (4) 30일 자동 만료."
  effort: small
  impact: medium
  evidence: "코드 web/lib/hooks/use-chats.ts L92-108: onError 에서 setQueryData 만 호출, localStorage/sessionStorage 0건. ChatMessage L100-104 의 onRetry 도 메모리 cache 의존. 새로고침 시 useInfiniteQuery 가 fresh 호출 → server 응답에는 failed 메시지 없음."

- id: imp-0090
  found_at_iter: 5
  area: chat
  type: feature
  target: quick_reply_templates
  problem: "거래 채팅에는 반복되는 정형 질문(예: '직거래 가능하신가요?', '시세 협의 가능하신가요?', '언제 시간 되시나요?')이 많은데 매번 타이핑해야 한다. 모바일 키보드 입력은 PC 보다 마찰이 크고 신규 사용자는 어떤 질문이 적절한지도 모른다."
  proposal: "(1) ChatInput 위에 호러즈ontal scrollable 칩 그룹 [거래 가능?] [시세 협의?] [지금 시간 됨?] [어디서 만날까요?] [강화 옵션?]. 칩 클릭 시 텍스트가 input 에 prefill 되고 사용자가 수정 후 전송. (2) 한 채팅방에서 사용한 칩은 24시간 hide(중복 메시지 방지). (3) 백엔드 분석으로 빈도 높은 표현을 자동 갱신."
  effort: medium
  impact: medium
  evidence: "코드 web/components/chat/chat-input.tsx 와 chat-panel.tsx 어디에도 quick reply / template 컴포넌트 0개. 사용자가 첫 채팅 진입 시 빈 textarea 외에 입력 보조 0건."

