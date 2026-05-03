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

- id: imp-0091
  found_at_iter: 6
  area: reservation
  type: feature
  target: reservation_card_action_buttons
  problem: "ReservationCardMessage(web/components/chat/reservation-card-message.tsx)는 예약 정보를 readonly 표시만 하고 [확정][거절][취소] 버튼이 0개다. 백엔드는 POST /reservations/:id/confirm, /cancel 엔드포인트와 GetReservationForConfirm/Cancel 핸들러를 모두 구현했지만 web 클라이언트(api-client.ts L249-258)에는 createReservation 만 있어 web 사용자는 받은 예약을 확정/거절할 방법이 채팅방 안에 전혀 없다. 즉 거래 핵심 플로우 proposed→confirmed 가 web 에서 작동 불가능."
  proposal: "(1) api-client.ts 에 confirmReservation(reservationId) / cancelReservation(reservationId, reasonCode) 추가. (2) ReservationCardMessage 에 status='proposed' 이고 !isMine 일 때 [확정하기][거절하기] 버튼 2개, isMine 이거나 status='confirmed' 일 때 [예약 취소] 버튼 1개를 카드 하단에 노출. (3) 클릭 시 confirm dialog → mutation → invalidate ['messages', chatId] + ['chats']. (4) 상태별 카드 색상 변경(proposed=gold border, confirmed=green border, cancelled=opacity-50 strikethrough). (5) metadataJson 에 reservationId 가 없으면 버튼 비노출 fallback."
  effort: medium
  impact: high
  evidence: "코드 web/components/chat/reservation-card-message.tsx L8-28: button/onClick 0건, status 분기 0건. web/lib/api-client.ts grep 'confirmReservation|cancelReservation' = 0건. backend cmd/server/handlers_reservation.go L82-155 에는 handleConfirmReservation/handleCancelReservation 모두 구현. docs/STATE_SEQUENCE_DIAGRAMS.md L130 'proposed→confirmed by counterparty' 핵심 전이가 web 에서 트리거 불가."

- id: imp-0092
  found_at_iter: 6
  area: reservation
  type: ux
  target: reservations_route_404
  problem: "https://giranjt.com/reservations 직접 진입(외부 링크/북마크/알림 deepLink)이 404 'This page could not be found'를 반환한다. web/app 트리에 reservations 디렉토리가 존재하지 않는다. OPENAPI_DRAFT.md L687-693 의 알림 카탈로그에는 deepLink='/trades/01CHAT...' 만 있어 /reservations 경로 자체가 미정의이지만, 사용자/SEO/외부 공유 측면에서 자연스러운 진입 경로가 막혀있고 anonymous 가 redirect 받아 /login 으로 가는 것도 아니라 의도가 더 모호하다."
  proposal: "(1) web/app/reservations/page.tsx 신규 추가, useAuthGuard 로 게이팅 후 useMyTrades() 데이터 중 chatStatus 가 reservation_proposed/reservation_confirmed 인 거래만 필터링해 '예약 카드 리스트' 형태로 노출(제목, 일시, 접선 방식, 상대방, [채팅방 이동] 버튼). (2) 또는 최소한 /reservations → /profile/trades?filter=reservation 으로 redirect 처리. (3) 알림 deepLink 도 /reservations/<reservationId> 로 통일하여 깊이있는 진입 가능."
  effort: medium
  impact: high
  evidence: "Playwright 익명 https://giranjt.com/reservations 375x667 → 페이지 본문 '404 / This page could not be found' bodyTextLen=60. find web/app -type d -name 'reservation*' 결과 0건(reservations 디렉토리 부재). docs/OPENAPI_DRAFT.md L692 의 deepLink='/trades/01CHAT...' 와 라우트 미스매치."

- id: imp-0093
  found_at_iter: 6
  area: reservation
  type: ux
  target: reservation_modal_datetime_validation
  problem: "ReservationModal(web/components/forms/reservation-modal.tsx L45-46)의 date input 은 min 속성이 없어 사용자가 과거 날짜(2020-01-01 등)를 선택할 수 있다. 또한 Date+Time 을 'T...:00Z' 로 단순 결합해 한국 사용자 입력(KST 의도)이 UTC 로 간주되어 실제로는 9시간 늦은 시각으로 저장된다. backend handlers_reservation.go L21 ScheduledAt 은 string 그대로 전달되어 DB 저장 시 KST↔UTC 차이를 알 수 없다."
  proposal: "(1) date input 에 min={new Date().toISOString().slice(0,10)} 추가, max=오늘+30일(reservation_form_sheet.dart 와 동일 정책). (2) submit 시 KST 로컬 Date 를 명시적으로 UTC 로 변환: new Date(`${date}T${time}:00+09:00`).toISOString(). (3) 입력값이 60분 이내 임박이면 컨펌 dialog '15분 뒤 약속이에요. 정말 제안하시겠어요?'. (4) ReservationCard 표시 시에도 toLocaleString('ko-KR', { timeZone: 'Asia/Seoul' }) 로 KST 노출."
  effort: small
  impact: high
  evidence: "코드 web/components/forms/reservation-modal.tsx L25 `${form.scheduledDate}T${form.scheduledTime}:00Z` (Z=UTC suffix). L45 input[type=date] 에 min/max 속성 0건. frontend/lib/features/reservation/reservation_form_sheet.dart L54-55 firstDate=DateTime.now(), lastDate=+30days, .toUtc() 정상 변환과 비대칭."

- id: imp-0094
  found_at_iter: 6
  area: reservation
  type: feature
  target: reservation_meeting_type_either
  problem: "ReservationModal 의 meetingType 드롭다운(L48-51)은 '인게임' / 'PC방/오프라인' 2종만 노출한다. 그러나 Flutter ReservationFormSheet L82 와 backend(meeting_type 컬럼)는 'either'(무관) 옵션을 지원하고 OPENAPI_DRAFT.md L524 도 server_id 옵션을 가정한다. web 사용자는 '협의 가능' 의도를 표현 못 해 협상 마찰 증가."
  proposal: "(1) <option value='either'>협의(무관)</option> 추가. (2) 'in_game' 선택 시 server_id 입력(서버 드롭다운 — listing.serverName 자동 prefill), 'offline_pc_bang' 선택 시 meetingPointText placeholder='예: 강남역 PC방 던파' 강제, 'either' 선택 시 두 입력 모두 hide. (3) form layout 을 conditional rendering 으로 정리해 사용자가 불필요한 필드를 안 보게 한다."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/reservation-modal.tsx L48-51: option=['in_game','offline_pc_bang'] 2종. frontend/lib/features/reservation/reservation_form_sheet.dart L79-83: 3종(in_game, offline_pc_bang, either). backend/cmd/server/handlers_reservation.go L21 MeetingType 자유 string. docs/OPENAPI_DRAFT.md L525 server_id 필드 정의."

- id: imp-0095
  found_at_iter: 6
  area: reservation
  type: ux
  target: reservation_proposal_pre_validation
  problem: "예약 제안 submit 시 backend가 'CONFLICT: 이미 활성 예약이 존재합니다'(handlers_reservation.go L44-47)를 반환할 수 있는데, web 모달은 이 에러를 일반 toast '예약 제안에 실패했습니다'(reservation-modal.tsx L33)로만 처리한다. 사용자는 왜 실패했는지, 기존 예약이 어떤 상태인지, 어떻게 해결할지 알 수 없어 같은 시도를 반복하다 이탈한다."
  proposal: "(1) reservation-modal 진입 전 useChats 데이터에서 activeChat.reservationStatus 또는 chatStatus 를 확인해 이미 'reservation_proposed/confirmed' 면 모달 대신 '이미 예약이 진행 중이에요. 기존 예약을 먼저 처리해주세요' alert + [기존 예약 보기] 버튼 노출. (2) submit 실패 시 error.code 별 분기: CONFLICT='이미 활성 예약이 있어요. 기존 예약을 취소한 후 다시 제안해주세요', VALIDATION_ERROR='입력값을 확인해주세요'. (3) backend 응답 body 의 error.code 를 apiClient 가 throw 하도록 통일."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/reservation-modal.tsx L32-34 catch 블록: addToast('error', '예약 제안에 실패했습니다') — error 객체 무시. backend handlers_reservation.go L45 'CONFLICT' code 정의되어 있으나 web 에서 사용 안 됨. 채팅방 진입 전 사전 가드 0건."

- id: imp-0096
  found_at_iter: 6
  area: reservation
  type: a11y_mobile
  target: reservation_modal_form_field_44px
  problem: "ReservationModal 의 input/select(reservation-modal.tsx L39 inputClass='py-2.5')는 약 38-40px 높이로 WCAG 2.5.5 권장 44px 미달이다. 모바일 375x667 에서 date+time input 이 grid-cols-2 로 좁게 배치되어 추가로 한 손 조작 시 터치 정확도가 떨어진다. submit 버튼은 py-3(약 44px+) 로 OK 지만 폼 필드들은 작아 균형이 나쁨."
  proposal: "(1) inputClass 의 py-2.5 → py-3 또는 min-h-[44px], placeholder/value font-size 를 16px(text-base) 로 키워 iOS 자동 줌 방지. (2) date+time grid 를 모바일에서 grid-cols-1 stacked 로 변경(sm:grid-cols-2 로 데스크탑만 분리). (3) 라벨 가시화: 현재 aria-label 만 있는데 visible label 추가('거래 일시', '접선 방식') — 폼 길이가 짧으니 라벨 토글 비용이 적다."
  effort: trivial
  impact: medium
  evidence: "코드 web/components/forms/reservation-modal.tsx L39 inputClass: 'px-3 py-2.5 text-sm' (text-sm=14px+py-2.5=10px*2+content). L44 'grid grid-cols-2 gap-3' 모바일에서도 강제 2열. submit button L54 'py-3' 만 충분한 높이."

- id: imp-0097
  found_at_iter: 6
  area: reservation
  type: content
  target: reservation_card_message_status_label
  problem: "ReservationCardMessage 하단 카피(L23)가 isMine 따라 '예약 제안을 보냈습니다' / '예약 제안을 받았습니다' 2종만 분기한다. 그러나 reservation 상태는 proposed/confirmed/expired/cancelled/fulfilled 5종(STATE_SEQUENCE_DIAGRAMS.md L93-97)이고 chatStatus 도 같이 변동한다. 사용자는 카드를 봐도 '이게 아직 진행 중인지, 확정됐는지, 만료됐는지' 한눈에 모른다."
  proposal: "(1) message.metadataJson.status (또는 backend 가 system message 갱신) 기반으로 카피 분기: proposed='상대방의 응답을 기다리는 중', confirmed='✅ 약속이 확정되었어요. 늦지 않게 만나요', cancelled='❌ 예약이 취소되었습니다 — 사유: <reasonCode 한글 매핑>', expired='⏰ 응답 시간이 지나 자동 만료되었어요', fulfilled='🎉 거래가 완료되었습니다'. (2) 카드 색상도 함께 변경(border-gold/green/red/text-dim). (3) 만료까지 남은 시간을 라이브 카운터로 표시('23분 남음')."
  effort: small
  impact: high
  evidence: "코드 web/components/chat/reservation-card-message.tsx L22-24: 단순 isMine 분기 2가지. status 분기 0건. metadataJson.status / metadataJson.reservationStatus 참조 0건. docs/STATE_SEQUENCE_DIAGRAMS.md L93-97 5상태 정의."

- id: imp-0098
  found_at_iter: 6
  area: reservation
  type: feature
  target: reservation_quick_time_presets
  problem: "예약 모달의 date+time 입력은 빈 input 두 개로 시작해 사용자가 키보드/스피너로 한 자리씩 입력해야 한다. 거래 합의는 보통 '오늘 저녁', '내일 점심', '주말'처럼 자연어 표현으로 이뤄지는데 이를 폼 필드 입력으로 번역하는 마찰이 크다. 모바일에서 native datepicker 가 뜨더라도 두 단계(date→time)가 별도라 시간 더 걸린다."
  proposal: "(1) 모달 상단에 빠른 프리셋 칩 그룹 [오늘 저녁(20:00)] [내일 점심(12:00)] [내일 저녁(20:00)] [주말 오후(토 14:00)] [직접 선택]. 칩 클릭 시 form.scheduledDate/Time 자동 채움. (2) 사용자 마지막 선택 5개를 localStorage 'reservation_recent_times' 에 저장해 다음 모달 열 때 우선 노출. (3) 기본 선택은 '내일 저녁' (Flutter 기본값과 일치)."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/reservation-modal.tsx L17 useState 초기값 scheduledDate=''/scheduledTime=''(prefill 0). 빠른 선택 칩/프리셋 0개. frontend reservation_form_sheet.dart L17-18 은 +1day 14:00 prefill — web 과 비대칭."

- id: imp-0099
  found_at_iter: 6
  area: reservation
  type: ux
  target: reservation_completion_flow_visibility
  problem: "예약이 confirmed 된 후 '거래 완료'를 트리거할 UI가 web 어디에도 없다. backend는 trade_completions 테이블과 handleCreateTradeCompletion(handlers_reservation.go 추정)을 갖고 있고 OPENAPI_DRAFT.md L588-603 도 정의되어 있지만, web/lib/api-client.ts L261 의 completeTrade 는 listingId 기반 단일 호출이고 채팅방/예약 카드에서 호출하는 버튼이 없다. 사용자는 '약속 시간이 지났는데 어떻게 거래 완료를 처리하지?' 막힌다."
  proposal: "(1) ReservationCardMessage 가 status='confirmed' 이고 scheduledAt 이 지난 시점이면 [거래 완료 처리] 버튼 노출. 클릭 시 완료 모달 → completion_note + [완료 요청] → backend 가 24h auto-confirm. (2) 동시에 채팅방 헤더 '예약 제안' 버튼이 있던 자리를 [거래 완료 처리]로 동적 교체. (3) 이미 자동확정 대기 중이면 카드에 '상대방의 확정 대기 중 (자동확정 23h 남음)' 카운트다운. (4) 완료 후 [리뷰 쓰기] 자동 prompt(imp-0086 과 연계)."
  effort: medium
  impact: high
  evidence: "코드 web/components/chat/reservation-card-message.tsx 와 chats/[id]/page.tsx 어디에도 completeTrade/거래완료 트리거 button 0개. web/lib/api-client.ts L261-269 completeTrade 정의는 있으나 호출처 grep 0건. docs/OPENAPI_DRAFT.md L588 'POST /listings/{listingId}/trade-completions' 엔드포인트와 web UI 비대칭."

- id: imp-0100
  found_at_iter: 6
  area: reservation
  type: performance
  target: reservation_modal_lazy_mount
  problem: "ReservationModal 은 ChatDetailPage(L11 import + L176 mount)에 항상 마운트되어 open=false 시에도 useState/Modal portal/form fields 가 메모리에 유지된다. ReportModal 도 동일 패턴(L182). 채팅방 자체가 SSE 폴링/스크롤 등 리소스를 많이 쓰는 페이지라 사용 빈도 낮은 모달들이 항상 React tree 에 있는 것은 의미 있는 누적 비용."
  proposal: "(1) {reservationOpen && <ReservationModal ... />} 형태로 conditional render 변경 → open=false 시 unmount. (2) 또는 dynamic import: const ReservationModal = dynamic(() => import('@/components/forms/reservation-modal').then(m => m.ReservationModal), { ssr: false }) 로 첫 클릭 시점에 코드도 함께 로드(번들 분할). (3) Modal 컴포넌트 자체에 open=false 시 children 을 return null 하는 패턴이 이미 있으면 dynamic import 만 적용."
  effort: trivial
  impact: low
  evidence: "코드 web/app/chats/[id]/page.tsx L176-188: <ReservationModal open={reservationOpen} ...> + <ReportModal open={reportOpen} ...> 항상 마운트. dynamic import 0건, conditional render 0건. ReservationModal 본체 useState L17 form 객체가 매 마운트 init."

- id: imp-0101
  found_at_iter: 6
  area: reservation
  type: a11y_mobile
  target: reservation_modal_focus_trap_and_escape
  problem: "Modal 컴포넌트(@/components/ui/modal) 가 focus trap / aria-modal / Escape 닫기 / overlay 클릭 닫기를 모두 구현했는지 reservation-modal 사용처에서 확인 불가. 일반적으로 form 입력 후 'Tab' 으로 모달 외부로 빠져나가거나, 모바일에서 backdrop 클릭만으로 닫혀 입력 데이터가 사라지는 사고(예약 제안은 한 번 닫히면 데이터 손실)가 흔하다. 또한 sheet 열림 시 body 스크롤 lock 도 없으면 배경이 함께 스크롤된다."
  proposal: "(1) Modal 에 role='dialog' aria-modal='true' aria-labelledby='reservation-title' 명시. (2) 첫 렌더 시 첫 input 자동 focus, focus trap(예: focus-trap-react 또는 직접 onKeyDown Tab 가드). (3) Escape 키 → onClose 호출. (4) 모바일 backdrop tap 으로 닫기 전 form 변경값이 있으면 confirm '입력 중인 예약 정보가 사라집니다. 닫으시겠어요?'. (5) document.body.style.overflow='hidden' on open / 복구 on close."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/reservation-modal.tsx L42 <Modal title='예약 제안'> — props 에 aria-labelledby/initialFocus/escapeClose 0건 전달. Modal 컴포넌트 자체 정의(@/components/ui/modal) 미확인이지만 form 변경값 보존 로직은 호출처에 0건. 모바일에서 backdrop tap 시 onClose 직행."

- id: imp-0102
  found_at_iter: 7
  area: review
  type: feature
  target: reviews_index_route_404
  problem: "/reviews 경로가 production 에서 404 페이지(영문 'This page could not be found.')를 반환한다. 사이트 메뉴/CTA 어디에도 /reviews 직접 진입은 없지만, 사용자가 거래 신뢰성을 비교하기 위해 '리뷰' 라는 검색/북마크/외부 링크로 진입할 가능성이 있다. 또한 리뷰는 익명 사용자에게 거래 안전성을 보여주는 핵심 신뢰 신호인데 글로벌 진입점 자체가 없어 거래 결정에 활용되지 못한다."
  proposal: "(1) /reviews 라우트를 신설해 최근 작성된 리뷰 N개(좋았어요/아쉬웠어요 비율, 거래 카테고리, 시간) 를 익명에게 공개해 플랫폼 활성도와 신뢰 분위기를 가시화한다. (2) 또는 /reviews → /profile/me/reviews(로그인) / 404+안내(비로그인) 로 명확하게 redirect. (3) 최소 변경으로는 globally not-found.tsx 를 한국어 + 추천 경로(매물/채팅) 안내로 교체."
  effort: medium
  impact: medium
  evidence: "Playwright GET https://giranjt.com/reviews → status 200 이지만 본문은 '404 / This page could not be found.' (h1='404'). bodyTextLength=60. 영문 잔존 + 사용자 다음 액션 가이드 0건."

- id: imp-0103
  found_at_iter: 7
  area: review
  type: feature
  target: rating_dimension_binary_only
  problem: "리뷰 평점이 'positive' / 'negative' 두 단계만 존재한다(backend handlers_review.go L21 binding:oneof=positive negative; ReviewModal '좋았어요' / '아쉬웠어요'). 거래의 핵심 차원 — 약속 시간 준수, 시세 합리성, 의사소통 매너, 매물 상태 일치 — 이 단일 척도에 융합되어 buyer/seller 모두 어디가 좋고 어디가 부족했는지 분리할 수 없다. 5점 척도까지는 아니더라도 차원 분리는 거래 결정에 큰 가치."
  proposal: "스키마 확장: reviews.dimension_punctuality SMALLINT(0~2), dimension_communication SMALLINT, dimension_item_condition SMALLINT(매물 상태가 적용 가능한 경우만). UI: ReviewModal 에서 각 차원별 👍/👎/N/A 3-way 토글 + 종합 한 줄 코멘트. 프로필 페이지에서 dimension 별 누적 그래프(예: '시간 약속 92%, 응답 88%')를 노출해 신뢰 신호를 풍부하게."
  effort: large
  impact: high
  evidence: "코드 backend/cmd/server/handlers_review.go L21 'rating string binding:required,oneof=positive negative'. ReviewModal L41-54 에 두 개 button 외 별도 차원 입력 0건. 프로필 stat 카드(profile/page.tsx L94-97)는 'positiveReviewCount' 단일 숫자만 노출."

- id: imp-0104
  found_at_iter: 7
  area: review
  type: ux
  target: review_modal_not_wired_into_app
  problem: "ReviewModal 컴포넌트가 web/components/forms/review-modal.tsx 에 정의되어 있지만, 코드베이스 전체에서 어디에서도 import / mount 되지 않는다(grep ReviewModal: 정의·테스트만, 사용 0건). 즉 거래 완료(deal_completed) 후에도 사용자가 후기를 작성할 수 있는 UI 진입점이 production 에 없다. 백엔드 POST /completions/:id/reviews 와 모달 컴포넌트는 모두 존재하지만 연결이 끊겨 reviewable 거래의 후기 작성률이 0% 에 가까울 가능성이 높다."
  proposal: "(1) chats/[id] 페이지에서 chatStatus === 'deal_completed' 가 되면 상단에 '거래는 어땠나요? 한 줄 후기 남기기' 배너 + ReviewModal 트리거 추가. (2) /profile/trades 의 '거래 완료' 카드에 '후기 작성' 버튼을 inline 노출(이미 작성한 거래는 '후기 보기' 로 swap). (3) 거래 완료 시점 SSE/notification 으로 후기 작성 모달 자동 prompt(첫 1회만, 사용자가 '나중에' 누르면 24h 후 재안내)."
  effort: medium
  impact: high
  evidence: "Grep 'ReviewModal' across web/app + web/components: 정의 review-modal.tsx, 테스트 review-modal.test.tsx 외 사용 0건. profile/trades/page.tsx 의 '거래 완료' 카드(chatStatusLabel.deal_completed)에 후기 진입 버튼/링크 0개. backend POST /completions/:id/reviews 핸들러는 정의되어 있어 호출만 안 되는 상태."

- id: imp-0105
  found_at_iter: 7
  area: review
  type: feature
  target: review_aggregate_on_listing_card
  problem: "매물 카드에 판매자 신뢰 신호로 'trustBadge'(newcomer 등)·'completedTradeCount' 만 노출되며 후기 비율(positive_ratio) 이나 최근 후기 한 줄 인용은 없다. 같은 'newcomer' 라도 거래 5건 평균 100% 좋았어요인 셀러와 거래 5건 평균 60% 인 셀러는 구매자 의사결정에 결정적 차이지만 카드에서 구별할 방법이 없다."
  proposal: "ListingCard 에 작은 calc 'positiveRatio = positive / (positive + negative)' 배지를 추가(예: '👍 92% (24)'). 데이터 소스는 listings list API 응답에 author.positiveReviewCount + author.negativeReviewCount 추가하거나 캐시 컬럼 review_score_cached 도입. 0 건이면 '아직 후기 없음' 으로 명시(공백 X)."
  effort: medium
  impact: high
  evidence: "API GET /api/v1/listings 응답 author 객체: { nickname, responseBadge, trustBadge, userId } — review 관련 필드 0건. ListingCard 컴포넌트(이전 iter listings 섹션 분석) 도 후기 비율/평균 표시 미존재. trustBadge='newcomer' 가 92% 사용자에 적용되어 변별력 약함."

- id: imp-0106
  found_at_iter: 7
  area: review
  type: ux
  target: reviews_page_no_seller_context
  problem: "/profile/[userId]/reviews 페이지가 '받은 리뷰 (N)' 헤더만 보여주고 어떤 셀러의 리뷰인지(닉네임/아바타/배지/거래 수/매물 링크) 어디에도 표시하지 않는다. 외부 공유 링크나 다른 매물에서 '리뷰 보기 ›' 로 진입하면 사용자가 어느 셀러의 페이지에 있는지 즉시 파악 불가. 빈 상태('아직 리뷰가 없습니다') 도 셀러 컨텍스트 없이 일반 카피만 표시."
  proposal: "페이지 상단에 sticky 헤더로 (아바타·닉네임·trustBadge·거래수·응답뱃지) 카드 + '매물 보기' 버튼을 추가하고, 빈 상태 카피를 '아직 {nickname}님에 대한 리뷰가 없어요. 거래를 시작해 보세요.' 로 변경. /profile/[userId]/reviews 가 SSR/streaming 으로 minimal 메타정보(닉네임)는 즉시 보이게 한다."
  effort: small
  impact: medium
  evidence: "Playwright /profile/da1e387a-fd67-4189-8475-b1a379e9a2a2/reviews — bodyText 60자, sellerName regex 매칭 0건, 셀러 프로필 link 0건, avatar 0건. 페이지 코드 web/app/profile/[userId]/reviews/page.tsx L29-31 헤더는 '받은 리뷰 ({reviews.length})' 만."

- id: imp-0107
  found_at_iter: 7
  area: review
  type: feature
  target: rating_aggregate_visualization
  problem: "받은 리뷰 페이지가 개별 리뷰 카드만 timeline 으로 나열한다. 리뷰가 100건 쌓이면 사용자는 좋았어요/아쉬웠어요 비율, 시간대별 추이, 매물 카테고리별 후기 분포를 파악할 수 없다. 신뢰는 개별 한 줄 코멘트가 아니라 누적 통계에서 온다."
  proposal: "/profile/[userId]/reviews 상단에 (1) 좋았어요 vs 아쉬웠어요 비율 horizontal bar('👍 92% · 👎 8% · 총 24개'), (2) 최근 30일 vs 그 이전 비율 비교(최근 추세 신호), (3) 카테고리별 분포(무기/방어구/소모품) 작은 칩 — 백엔드는 GET /users/{id}/review-stats 신규 엔드포인트로 단일 fetch."
  effort: medium
  impact: high
  evidence: "코드 web/app/profile/[userId]/reviews/page.tsx L33-62: 개별 review 만 map render, 집계 영역 0건. backend handlers_review.go handleGetUserReviews 응답도 raw list 만 반환(L73 'data: reviews'), 통계 필드 0건."

- id: imp-0108
  found_at_iter: 7
  area: review
  type: ux
  target: review_text_length_limit
  problem: "ReviewModal '한줄 코멘트(선택)' textarea 가 maxLength / minLength / 글자수 카운터 모두 없고 height h-24(약 96px)에 placeholder 만 있다. 사용자는 한 줄을 의도했는데 1000자 장문을 쓰거나, 반대로 빈칸/이모지 한 글자로 노이즈 후기를 만들기 쉽다. 백엔드 binding 도 comment 길이 검증 0건이라 DB · 카드 레이아웃이 깨질 수 있다."
  proposal: "textarea 에 maxLength={200} 설정, '한 줄 코멘트 (선택, 최대 200자)' label 명시, 우하단 실시간 카운터('45/200'). 너무 짧은(<5자 또는 emoji-only) 입력은 submit 시 inline 경고 'rating 만 제출하시거나 5자 이상의 코멘트를 입력해 주세요'. 백엔드 binding 에 max=200,omitempty 추가."
  effort: trivial
  impact: medium
  evidence: "코드 review-modal.tsx L56-62 <textarea> 에 maxLength/minLength/카운터 0건. backend handlers_review.go L21-23 req.Comment *string 에 binding 검증 0건."

- id: imp-0109
  found_at_iter: 7
  area: review
  type: feature
  target: review_seller_response
  problem: "negative 후기를 받은 셀러가 자기 입장을 해명/사과할 수 있는 답글(seller response) 기능이 없다. 거래 분쟁/오해에서 'negative=세입자 사정으로 약속 늦음'·'negative=가격 변경 요구' 같은 컨텍스트가 빠지면 후기는 일방적 선언이 되고, 셀러 입장에서 플랫폼에 대한 신뢰가 떨어진다."
  proposal: "reviews 테이블에 seller_response_text TEXT, seller_response_at TIMESTAMPTZ 컬럼 추가. POST /reviews/{id}/response 엔드포인트(셀러 본인만, 1회 한, 7일 이내). UI: review 카드 아래 셀러 답글 inline 표시 ('셀러 답변: ...'). negative 후기에만 답글 prompt notification 발송."
  effort: large
  impact: medium
  evidence: "코드 backend/cmd/server/handlers_review.go 전체 — POST/PATCH 엔드포인트 1개(handleCreateReview)뿐. db/migrations 검색 시 review reply 컬럼 0건(reviews 스키마는 본 task 에서 직접 미확인 — 추정). UI 측 review 카드(profile/[userId]/reviews/page.tsx L34-62) 에 답글 영역 0건."

- id: imp-0110
  found_at_iter: 7
  area: review
  type: ux
  target: review_edit_or_delete_own
  problem: "리뷰 작성 후 작성자가 수정 / 삭제할 방법이 backend handlers_review.go 에 없다(POST createReview + GET listUserReviews 만). 즉 한 번 누른 'positive/negative' 토글이 영구적으로 고정되어 작성 직후 후회/오타/사실관계 정정 시 운영자 개입 외에 정정 불가."
  proposal: "PATCH /reviews/{id} 와 DELETE /reviews/{id} 추가, 작성자 본인만 가능, 작성 후 7일 이내 윈도우(7일 후엔 분쟁 방지 차원 잠금). UI: 본인이 작성한 리뷰 카드 우측에 '수정 / 삭제' 메뉴(이미 7일 지나면 disabled + 안내 툴팁). 모든 변경은 audit table review_audits 에 before/after 기록."
  effort: medium
  impact: medium
  evidence: "코드 backend/cmd/server/handlers_review.go: handler 함수 2개 (handleCreateReview, handleGetUserReviews) 만 정의. PATCH/DELETE 핸들러 0건. UI side review 카드(profile/[userId]/reviews/page.tsx L34-62) 에 수정/삭제 버튼 0건."

- id: imp-0111
  found_at_iter: 7
  area: review
  type: feature
  target: review_report_abuse
  problem: "리뷰는 신원 일부(닉네임)와 함께 공개되는데 모욕/사칭/허위/욕설 후기를 제3자가 신고할 진입점이 UI 에 없다(리뷰 카드에 신고/오류 신고 버튼 0). 매물·사용자 신고 모달은 존재하지만 'review 단위' 신고는 reportable_type='review' 같은 카테고리로 분리되지 않은 듯하다."
  proposal: "review 카드에 ⋯ 메뉴 → '이 리뷰 신고' 추가, 사유 옵션(허위·욕설·개인정보·스팸). reports 테이블에 reportable_type='review', reportable_id=review_id 지원. 운영자 대시보드(admin)에서 review_report 큐 확인 → 후기 hide / 삭제 / 작성자 경고. 클라이언트는 신고된 review 카드를 'hidden' 상태로 즉시 회색 처리."
  effort: medium
  impact: medium
  evidence: "코드 web/app/profile/[userId]/reviews/page.tsx L34-62 리뷰 카드 dom 에 신고 button 0건. 기존 신고 모달은 매물 상세에만 적용(이전 iter listing_detail report_button_anonymous_unreachable 와 별개로 review 단위 신고는 모달 prop type 검토 필요)."

- id: imp-0112
  found_at_iter: 7
  area: review
  type: a11y_mobile
  target: review_modal_keyboard_and_aria
  problem: "ReviewModal 의 'positive/negative' 두 버튼은 aria-pressed 만 있고 group 컨테이너에 role='radiogroup' / aria-label 이 없다. 스크린리더 사용자는 '두 개의 토글이 있다'까지만 인지하고 '하나만 선택하는 거래 후기 평가' 라는 의도를 알 수 없다. 또한 모달 자체의 focus 진입 / Escape / submit 비활성 상태(rating==null)는 시각적으로만 표시되어 스크린리더로는 '버튼 disabled' 만 들린다."
  proposal: "(1) 두 버튼 wrapper 에 role='radiogroup' aria-label='거래 평가' aria-required='true'. (2) 각 button 을 role='radio' 로 변경하거나, 차라리 <input type='radio'> + label 로 네이티브 시맨틱 사용. (3) 모달 open 시 첫 라디오에 자동 focus, ←→ 키로 라디오 이동. (4) submit disabled 상태에 aria-describedby 로 '평가를 먼저 선택해 주세요' 안내 텍스트 연결. (5) 코멘트 textarea 에 aria-describedby 로 글자수 카운터 연결."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/review-modal.tsx L40-55: <button aria-pressed={...}> 두 개를 <div className='flex gap-3'> 으로 감쌈. role='radiogroup'/role='radio' 0건. submit 버튼(L63) disabled 상태에 aria-describedby 0건. 키보드 ←→ 핸들러 0건."

- id: imp-0113
  found_at_iter: 7
  area: review
  type: content
  target: review_rating_label_neutrality
  problem: "rating 라벨 '좋았어요' / '아쉬웠어요' 는 친근하지만, 부정 라벨이 '나빴어요' 가 아닌 '아쉬웠어요' 라 실제 사기/약속 미이행 같은 강한 부정 케이스를 흐리게 만든다. 또한 영문 enum 'positive/negative' 와 한국어 라벨 사이 의미 강도가 어긋나 운영 통계('positive 비율 92%')가 실제 사용자 만족과 괴리될 수 있다."
  proposal: "라벨을 두 단계로 좀 더 구체화: '믿을 만했어요(좋았어요)' / '문제 있었어요(아쉬웠어요)' 로 보조 설명 추가하고, '문제 있었어요' 선택 시 textarea placeholder 가 '예: 약속 시간 30분 늦으셨어요. 그래도 거래는 완료됐습니다.' 같은 구체 가이드로 변경되도록 동적 처리. 또는 영문 enum 을 'satisfied/dissatisfied' 로 정리하고 라벨/문서 일관성 확보."
  effort: trivial
  impact: low
  evidence: "코드 web/components/forms/review-modal.tsx L52 'r === positive ? \"좋았어요\" : \"아쉬웠어요\"'. backend handlers_review.go L21 oneof=positive negative. profile/[userId]/reviews/page.tsx L51 'r.rating === positive ? \"👍 좋아요\" : \"👎 아쉬워요\"' — 같은 enum 에 라벨이 'modal=좋았어요/아쉬웠어요', 'list=좋아요/아쉬워요' 로 두 가지 버전 혼재."

- id: imp-0114
  found_at_iter: 7
  area: review
  type: ux
  target: review_link_to_listing_context
  problem: "받은 리뷰 카드가 (작성자 닉네임, rating, comment, 시간) 만 보여주고 어떤 거래/어떤 매물 에 대한 후기인지 표시하지 않는다. 셀러 입장에서 '동일 닉네임으로 여러 후기가 쌓이면 어느 거래 건이 negative 였는지' 알 수 없고, 잠재 구매자도 '이 셀러는 무기 거래에서 신뢰도가 높다'/'소모품 거래에서 후기 좋다' 같은 카테고리 기반 신뢰 판단이 불가능."
  proposal: "리뷰 카드에 (소형 매물 아이콘 + 매물 제목 + 거래일 + 가격) 인라인 chip 추가, 클릭 시 매물 상세로 이동. 백엔드 ListUserReviews 응답에 listing_id, listing_title, item_icon_url, completed_at 필드 join 추가. 매물이 삭제됐으면 '삭제된 매물' 로 회색 처리."
  effort: small
  impact: medium
  evidence: "코드 backend handlers_review.go L70-72: 응답에 listingId/itemName 0건. profile/[userId]/reviews/page.tsx L34-62 카드에 매물 정보 영역 0건. UserReviewItem 구조체(repository/interfaces.go L581 부근)에도 매물 메타 필드 없음으로 추정."

- id: imp-0115
  found_at_iter: 7
  area: review
  type: performance
  target: reviews_pagination_and_caching
  problem: "GET /users/{id}/reviews 가 단일 SQL 로 전체 리뷰를 한 번에 반환한다(handlers_review.go L62 ListUserReviews). 활발한 셀러가 후기 500건 쌓이면 한 번에 500개 row + reviewerNickname JOIN 결과를 한 응답에 담아 모바일 LCP/메모리/렌더 모두 큰 부담. 클라이언트는 useQuery 로 받지만 페이지네이션도 없고 staleTime/gcTime 설정도 default."
  proposal: "(1) 백엔드: 쿼리 파라미터 ?cursor=...&limit=20 cursor pagination, 응답에 nextCursor 추가. (2) 프론트: useInfiniteQuery 로 변경, 무한 스크롤 + IntersectionObserver. (3) ETag/Last-Modified 또는 review_score_cached 컬럼으로 셀러 단위 캐시 무효화 신호. (4) staleTime 5분 + 새 리뷰 작성 시 invalidateQueries(['user-reviews', userId])."
  effort: medium
  impact: medium
  evidence: "코드 backend handlers_review.go L62: 'items, err := repo.ListUserReviews(ctx, targetUserID)' — limit/cursor 파라미터 0건. ListUserReviews 시그니처(interfaces.go L442) 'ListUserReviews(ctx, targetUserID string) ([]UserReviewItem, error)' — pagination 인자 없음. web/lib/hooks/use-reviews.ts L4-9 useQuery 만 사용, useInfiniteQuery 0건, staleTime 미설정."

- id: imp-0116
  found_at_iter: 7
  area: review
  type: feature
  target: review_helpful_vote
  problem: "쌓인 리뷰 중 '의미 있는' 후기와 '잘 거래했어요(2글자)' 같은 노이즈 후기를 구분하는 시그널이 없다. 결국 새로운 구매자가 가장 최근 후기 N개만 본 뒤 신뢰 판단을 내려야 하는데, 짧은 무내용 후기들이 timeline 을 차지하면 의사결정 비용 증가."
  proposal: "리뷰 카드에 '👍 도움돼요' 버튼 추가(로그인 사용자만, 자기 후기 제외). reviews_helpful_votes 테이블(review_id, user_id, voted_at, UNIQUE(review_id, user_id)). 정렬 옵션 '도움 순' 추가, helpful_count >= 3 인 후기는 카드 상단에 '추천 후기' 배지. 후기 작성자에게 helpful 받았을 때 알림(작성자 동기 부여)."
  effort: large
  impact: low
  evidence: "코드 backend handlers_review.go: helpful/vote 관련 핸들러 0건. db/migrations: helpful 컬럼/테이블 ripgrep 0건(추정). UI: profile/[userId]/reviews/page.tsx 정렬/필터 0건, '도움돼요' 버튼 0건."

- id: imp-0117
  found_at_iter: 7
  area: review
  type: a11y_mobile
  target: review_card_emoji_only_indicator
  problem: "받은 리뷰 카드에서 rating 표시가 '👍 좋아요' / '👎 아쉬워요' 처럼 emoji 가 텍스트 옆에 인라인으로 들어가 있다. 일부 스크린리더는 emoji 를 '엄지 위로 손' 같은 음성으로 읽고 사용자에 따라 비활성화하기도 해서 의미 전달이 일관적이지 않다. 또한 색상(green-400 vs danger) 만으로 긍정/부정 구분되는 경우 색맹 사용자에게 불리하다."
  proposal: "(1) 텍스트 prefix 를 '긍정 평가 — 좋아요' / '부정 평가 — 아쉬워요' 로 명시. (2) emoji 를 <span aria-hidden='true'> 로 감싸 스크린리더 중복 발화 방지. (3) 카드 배경 색상에 더해 좌측 4px border-l-* (green/red) 추가로 색상 외 시각 단서 강화. (4) 카드 자체에 aria-label='유저X 의 긍정 평가, N일 전, 코멘트: ...' 합성 문구 부여."
  effort: trivial
  impact: low
  evidence: "코드 web/app/profile/[userId]/reviews/page.tsx L51: '{r.rating === positive ? \"👍 좋아요\" : \"👎 아쉬워요\"}' — emoji 가 텍스트 노드에 직접 포함되어 aria-hidden 격리 없음. 카드(L37) 'border border-border' 만으로 색상 외 시각 단서 0건."

- id: imp-0118
  found_at_iter: 8
  area: report
  type: bug
  target: chat_report_target_type_mismatch
  problem: "채팅 화면에서 '신고' 버튼을 누르면 ReportModal 이 targetType='chat_room' 으로 POST /reports 를 호출하지만, 백엔드 binding 은 oneof=user listing message 만 허용하므로 모든 채팅 신고가 400 VALIDATION_ERROR 로 즉시 실패한다. 사용자에게는 '신고 접수에 실패했습니다' toast 만 노출되고 운영자는 채팅 단위 신고 데이터를 한 건도 수신할 수 없는 사일런트 결손 상태."
  proposal: "백엔드 oneof 에 chat_room 추가(또는 message 단위로 받게 클라이언트 수정). 추가로 reports.target_type 에 CHECK 제약 또는 enum 도메인 적용해 DB 차원에서 일관성 강제. validation 실패 시 응답 body 의 message 를 toast 에 직접 노출(현재는 generic 'INTERNAL_ERROR' 형태로 디버깅 정보 손실)."
  effort: trivial
  impact: high
  evidence: "코드 web/app/chats/[id]/page.tsx L182-187 <ReportModal targetType=\"chat_room\" />. backend/cmd/server/handlers_report.go L18 binding:\"required,oneof=user listing message\" — chat_room 미포함. Flutter 쪽 frontend/lib/features/report/report_form_sheet.dart L7 주석 'listing, user, chat_room, message, review' 도 백엔드와 불일치(review/chat_room 검증 실패)."

- id: imp-0119
  found_at_iter: 8
  area: report
  type: ux
  target: report_button_anonymous_unreachable
  problem: "비로그인 사용자가 매물 상세에 접속하면 액션 툴바 자체가 availableActions 가 비어 있어 sticky bottom toolbar 가 렌더되지 않고, 그 결과 '신고' 버튼도 함께 사라진다. 즉 사기성 매물을 본 익명 방문자가 신고할 진입점이 0건. (Playwright /listings/b155f6a7… 익명 접속, document.body.innerText 에 '신고' 단어 0회 발견.)"
  proposal: "신고 버튼은 actions.length 분기와 별개로 모든 사용자에게 항상 노출하고, 클릭 시 비로그인 사용자에게는 requireAuth('신고') 로 로그인 모달을 띄운 후 복귀 redirect. 또는 신고는 비로그인 신고도 허용(IP+ja3 fingerprint+캡차)해 진입 장벽을 낮춘다. 추가로 listing-card 에도 우상단 ⋯ 메뉴로 '신고' 항목 노출."
  effort: small
  impact: high
  evidence: "Playwright https://giranjt.com/listings/b155f6a7-a80b-4db5-ac3f-a080b3ad656d 익명 접속 후 querySelectorAll('button').filter(.textContent.includes('신고')).length === 0. 코드 web/app/listings/[id]/page.tsx L207 '{actions.length > 0 && (' 가드로 toolbar 자체가 비로그인 시 0건 렌더, 신고 버튼은 그 안 L263 위치."

- id: imp-0120
  found_at_iter: 8
  area: report
  type: feature
  target: my_reports_history_page
  problem: "백엔드는 GET /me/reports 를 이미 노출(handlers_report.go L45 handleMyReports)하지만, web 프론트에는 이 엔드포인트를 부르는 페이지·hook·UI 가 0건이다. 따라서 사용자는 자기가 어떤 신고를 냈는지·처리 상태(submitted/assigned/resolved)·운영자 회신 여부를 확인할 방법이 전혀 없고, 신고 후 즉시 잊어버리거나 '내 신고가 묻혔다'는 불신만 쌓인다."
  proposal: "/profile/reports 또는 /me/reports 페이지 신설. useMyReports() hook 으로 데이터 fetch, status 별 탭(접수됨/처리중/완료/반려), 각 행에 (대상 타입+ID, 신고 사유, 신고일, 현재 상태, 운영자 회신 메모, '추가 자료 제출' 버튼). 프로필 메뉴에 '내 신고 내역' 링크 추가."
  effort: medium
  impact: medium
  evidence: "코드 backend/cmd/server/main.go L137 'readOnly.GET(\"/me/reports\", handleMyReports(reservationRepo))' 등록됨. ripgrep web/lib/hooks/use-reports* 0건, web/app/me/reports/* 또는 web/app/profile/reports/* 0건. /me/reports API 는 backend 만 살아 있고 web UI 미연결."

- id: imp-0121
  found_at_iter: 8
  area: report
  type: bug
  target: report_reason_options_subset_of_backend
  problem: "ReportModal 의 REPORT_REASONS 배열은 6개(scam_suspicion, fake_listing, harassment, spam, no_show, other)만 노출하지만 백엔드는 8개(prohibited_item, privacy_exposure 추가)를 허용한다. 결과적으로 '금지 품목(현금화 게임머니 등)'이나 '개인정보 노출(폰번호 캡쳐 유포 등)' 같은 핵심 시나리오가 web 사용자에게는 'other' 라는 분류 불가 상태로 묻혀 운영자 큐에서 자동 분류·SLA 적용이 어려워진다."
  proposal: "REPORT_REASONS 에 {value: 'prohibited_item', label: '금지 품목/현금화'}, {value: 'privacy_exposure', label: '개인정보 노출'} 두 항목 추가. 사유별 placeholder 가이드(예: 사기 의심 → '입금 후 잠수, 차단 등 구체 정황을 적어주세요') 동적 노출. Flutter 와 web 의 라벨/icon/순서를 단일 source(shared/report-reasons.json)로 통일."
  effort: trivial
  impact: medium
  evidence: "코드 web/components/forms/report-modal.tsx L8-15 REPORT_REASONS 6 entries. backend/cmd/server/handlers_report.go L20 'oneof=fake_listing scam_suspicion no_show harassment spam prohibited_item privacy_exposure other' 8 entries. frontend/lib/features/report/report_form_sheet.dart L21-30 _reportTypes 8 entries(label/icon 보유) — 모바일은 노출하는데 web 은 누락된 비대칭."

- id: imp-0122
  found_at_iter: 8
  area: report
  type: feature
  target: evidence_attachment_screenshot_link
  problem: "신고 폼이 reportType + free-text description(최대 2000자)만 받고 스크린샷·외부 채팅 캡쳐·거래 인증샷 등 evidence 첨부 채널이 0건이다. 운영자가 사기 의심 신고를 받아도 '카카오톡에서 잠수했어요' 같은 텍스트만 보고 진위를 판정해야 해 false positive/negative 가 늘어난다. 또한 reports 테이블에 evidence_files JSON 컬럼이 없어 사후 추가 자료도 못 받음."
  proposal: "(1) reports 테이블에 evidence_files JSONB(파일 키 배열), evidence_links TEXT[] 컬럼 추가 마이그레이션. (2) ReportModal 에 '증거 자료(스크린샷·링크)' 섹션 추가, 최대 5장까지 이미지 업로드(presigned URL 패턴 재사용 — listing 이미지 업로드와 동일 스택), URL 입력 필드. (3) admin/app/reports/page.tsx 상세 패널에 evidence 갤러리 표시. (4) 1차 신고 후 24시간 이내 추가 자료 첨부 PATCH /me/reports/{id}/evidence 허용."
  effort: large
  impact: high
  evidence: "코드 web/components/forms/report-modal.tsx L46-66 form 안에 file input/url input 0건. backend handlers_report.go L17-22 req struct 에 evidence 필드 0건. db/migrations/001_initial.sql L223-235 reports 스키마에 evidence_* 컬럼 0건. admin/app/reports/page.tsx L164-197 상세 패널에 첨부 표시 영역 0건."

- id: imp-0123
  found_at_iter: 8
  area: report
  type: feature
  target: reporter_resolution_notification
  problem: "운영자가 admin/reports 에서 '완료 처리' 또는 '조치 실행' 을 누르면 reports.status 가 resolved 로 변하고 moderation_actions 에 행이 추가되지만, 신고를 낸 사용자(reporter) 에게 결과를 알려주는 notification 이 0건이다. 사용자는 자기 신고가 받아들여졌는지 반려됐는지·어떤 조치가 취해졌는지 알 수 없어 '신고는 보내봤자 무의미하다'는 학습된 무관심을 만들고 신고 자체가 줄어든다."
  proposal: "handleAdminReportAction / handleAdminUpdateReportStatus 안에서 status 가 resolved 또는 rejected 로 바뀔 때 notifications 테이블에 reporter_user_id 대상 row INSERT (type='report_resolved', body='귀하의 신고가 처리되었습니다: 경고 조치 완료', deep_link='/me/reports/{id}'). 추가로 운영자가 reject 시 사유 선택(증거 부족·중복 신고·정책 외) 후 reporter 에게 안내. SSE 알림 채널로 즉시 푸시."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_admin.go L80-110 handleAdminReportAction 트랜잭션 안에 INSERT INTO notifications 0건. L91 'UPDATE reports SET status = resolved' 직후 reporter notification 디스패치 없음. backend/internal/event/* 에 'report_resolved' 이벤트 정의 0건(grep)."

- id: imp-0124
  found_at_iter: 8
  area: report
  type: bug
  target: self_report_and_duplicate_guard_missing
  problem: "handleCreateReport 에 (1) 자기 자신을 신고하지 못하게 하는 로직, (2) 동일 (reporter_user_id, target_type, target_id) 쌍의 신고를 24시간 이내 중복 차단하는 로직, (3) 일정 시간당 신고 횟수 rate limit 가 모두 0건이다. 결과적으로 누군가 한 사용자/매물을 spam-신고로 폭격해 운영자 큐를 어지럽히거나, 자기 매물 신고로 노이즈 생성이 가능. reports 테이블에도 UNIQUE (reporter_user_id, target_type, target_id) 인덱스 0건이라 같은 reporter 가 같은 대상으로 100건 row 를 만들 수 있다."
  proposal: "(1) handleCreateReport 시작부에 target_type==='user' && target_id===userID → 400 'CANNOT_REPORT_SELF'. listing/chat 의 경우 owner 검증 후 동일하게 차단. (2) reports 테이블에 partial UNIQUE INDEX (reporter_user_id, target_type, target_id) WHERE created_at > NOW() - INTERVAL '24 hours' 또는 INSERT 전 SELECT count check. (3) middleware 로 사용자당 분당 5건/일 30건 신고 rate limit. (4) 어뷰저 식별(같은 reporter 가 동일 사용자에 대한 신고 5회 이상 false positive 누적)에 alignment penalty."
  effort: small
  impact: high
  evidence: "코드 backend/cmd/server/handlers_report.go L14-43 handleCreateReport: self-check, duplicate-check, rate-limit 0건. db/migrations/001_initial.sql L223-235 reports 스키마에 UNIQUE 인덱스 0건. handlers_listing.go·repository 에서도 reporter==target.owner 검증 grep 0건."

- id: imp-0125
  found_at_iter: 8
  area: report
  type: ux
  target: post_submit_no_eta_or_contact_channel
  problem: "신고 제출 후 사용자가 보는 피드백은 toast '신고가 접수되었습니다' 한 줄뿐이다. 처리 예상 기간(SLA), 운영팀 연락 채널, 추가 자료 제출 안내, 신고 ID(추후 문의 시 식별용) 등이 0건이라 '받았다는 거지 처리한다는 보장은 없는' 인상으로 끝난다. ReportModal handleSubmit L36-37 onClose() 후 즉시 닫혀 reportId 참조 불가."
  proposal: "신고 제출 성공 시 modal 을 즉시 닫지 말고 success 단계로 전환: (1) 신고 ID 표시(복사 버튼) (2) '평균 처리 기간 24-48시간' SLA 안내 (3) '추가 자료가 있으면 giranjt@gmail.com 으로 신고 ID 와 함께 보내주세요' (4) '내 신고 내역에서 진행 상황을 확인하세요 →' /me/reports 링크 CTA. 사용자가 직접 X 또는 '확인' 눌러야 닫힘."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/report-modal.tsx L37 'onClose(); addToast(\"success\", \"신고가 접수되었습니다\");' — 응답으로 받은 reportId(API 가 반환하는 값) 미사용. Modal 내부에 success state 분기 0건. SLA 안내 / contact 텍스트 / /me/reports 링크 0건."

- id: imp-0126
  found_at_iter: 8
  area: report
  type: a11y_mobile
  target: report_modal_missing_radiogroup_semantics
  problem: "REPORT_REASONS 6개 옵션이 <label><input type='radio'></label> 로 되어 있지만 wrapper(div.space-y-2) 에 role='radiogroup' / aria-required='true' / aria-labelledby(제목 링크) 가 없어 스크린리더 사용자는 라디오 6개의 그룹 의미와 '필수 선택' 사실을 모른다. 또한 description textarea 에도 글자수 카운터·max=2000 hint·실시간 글자수가 노출되지 않아 백엔드 binding(min=1,max=2000) 위반으로 400 응답을 받아도 사용자에게는 'INTERNAL_ERROR' 로 표시된다."
  proposal: "(1) <fieldset role='radiogroup' aria-labelledby='report-modal-title' aria-required='true'><legend className='sr-only'>신고 사유</legend>...</fieldset> 구조로 감싸기. (2) textarea 옆에 글자수 카운터 '0/2000' 라이브 업데이트, aria-describedby 로 카운터 연결. (3) submit 버튼 disabled 상태에 aria-describedby='report-form-hint' '사유를 선택해 주세요' sr-only 메시지 연결. (4) 에러 toast 대신 modal 내부 inline error region(role='alert') 으로 백엔드 message 그대로 표시."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/report-modal.tsx L48 '<div className=\"space-y-2\">' radio wrapper 에 role='radiogroup' 0건. L56-62 textarea 에 maxLength/카운터/aria-describedby 0건. L63 submit button disabled 에 aria-describedby 0건. L37-39 catch 블록은 toast 만 호출, modal 안 inline error 0건."

- id: imp-0127
  found_at_iter: 8
  area: report
  type: ux
  target: forced_description_silent_default
  problem: "ReportModal handleSubmit L35 '|| \"신고합니다\"' 로 description 이 비어 있으면 클라이언트가 강제로 '신고합니다' 를 채워 보낸다. 백엔드 binding 이 min=1,max=2000 이라 빈 본문을 막으려는 우회지만, 사용자가 의도하지 않은 본문이 운영자 큐에 들어가 분류·검색·중복 식별을 어렵게 만든다. 운영자 admin 화면에서 'description: 신고합니다' 만 있는 무의미 row 가 다수 발생할 위험."
  proposal: "(1) description 을 백엔드에서 진짜 optional 로 바꾸고(omitempty,max=2000) 빈 description 은 NULL 저장. (2) 클라이언트 fallback '신고합니다' 제거. (3) 사유 별 권장 본문 길이(예: '기타' 선택 시 description 50자 이상 강제, 그 외 reasons 는 선택) 차등 검증. (4) 너무 짧거나 의미 없는(반복 문자/이모지 only) 본문은 submit 시 inline 경고."
  effort: trivial
  impact: medium
  evidence: "코드 web/components/forms/report-modal.tsx L35 'description: description || \"신고합니다\"'. backend handlers_report.go L21 binding:\"required,min=1,max=2000\" — required 라서 클라이언트가 fallback 강제 송신. db/migrations/001_initial.sql L229 'description TEXT NOT NULL DEFAULT '''' — NULL 저장 불가."

- id: imp-0128
  found_at_iter: 8
  area: report
  type: feature
  target: admin_queue_priority_sla_auto_classify
  problem: "admin/app/reports/page.tsx 의 운영자 큐는 단순히 created_at DESC LIMIT 50 으로 신고 50건을 한 페이지에 나열한다. (1) 신고 사유별 자동 우선순위(scam_suspicion·privacy_exposure 는 high, spam 은 low) 0건, (2) SLA(접수 후 N시간 미처리 → '지연' 배지) 0건, (3) 동일 대상 다중 신고 자동 클러스터링(같은 listingId 에 5건 신고 → 카드 1개로 묶고 '5건 신고' 카운트) 0건, (4) 자동 모더레이션 후보(같은 reporter 가 동일 대상 5회 이상 → 자기 어뷰즈 의심 표시) 0건. 결과적으로 운영자 1명 체제에서 큐가 쌓이면 critical 신고가 묻힌다."
  proposal: "(1) reports 응답에 priority(높음/중간/낮음, reason 매핑) + sla_due_at 필드 추가. (2) admin 큐 컬럼에 우선순위·잔여 SLA 시계(빨강 ≤ 4h, 주황 ≤ 24h, 회색 > 24h) 표시. (3) 동일 (target_type, target_id) cluster_count 표시, 클릭 시 묶음 펼치기. (4) MVP 자동 분류기: spam keyword 매칭(전화번호 노출, 욕설사전) 으로 prohibited_item·privacy_exposure 자동 태깅 + reviewer 검수 단계 1회 추가. (5) 신고 사유별 LIMIT/페이지네이션."
  effort: large
  impact: medium
  evidence: "코드 backend/cmd/server/handlers_admin.go L18 'SELECT ... FROM reports ORDER BY created_at DESC LIMIT 50' — priority/sla/cluster 컬럼 0건. admin/app/reports/page.tsx L33-83 columns 에 priority/sla 표시 0건. 자동 분류 코드 backend/internal/* grep 'classify\\|priority\\|sla' 0건."

- id: imp-0129
  found_at_iter: 8
  area: report
  type: performance
  target: my_reports_endpoint_no_pagination_or_cache
  problem: "GET /me/reports 가 LIMIT 50 hard-coded 로 단일 페이지만 반환한다(repository L339). 신고를 자주 하는 사용자(전문 모니터링 사용자) 는 50건 이후 자기 신고 이력을 영영 못 보고, 클라이언트 useQuery 측 staleTime/cache 설정도 0건이라 페이지 진입마다 재호출. 또한 응답에 description 필드가 빠져 있어 사용자가 자기 신고 본문을 다시 보려면 불가."
  proposal: "(1) backend ListMyReports(ctx, userID, cursor, limit) 시그니처 변경, cursor pagination 응답에 nextCursor. (2) 응답 row 에 description, evidence_count, last_admin_response_at 추가. (3) web hook useMyReports() 는 useInfiniteQuery + IntersectionObserver. (4) staleTime 60s, on submit 시 invalidateQueries(['my-reports']). (5) 신고 본문 250자 이상이면 'show more' 토글로 truncate."
  effort: medium
  impact: low
  evidence: "코드 backend/internal/repository/postgres_reservation.go L339 'SELECT ... FROM reports WHERE reporter_user_id = $1 ORDER BY created_at DESC LIMIT 50' — cursor/limit param 0건. handlers_report.go L57-58 응답에 description 필드 0건. web 측 useMyReports hook 자체 부재(앞 imp-0120 와 별개로 백엔드 페이지네이션도 미구현)."

- id: imp-0130
  found_at_iter: 9
  area: notification
  type: ux
  target: notification_contract_field_mismatch_blank_rows
  problem: "백엔드 GET /notifications 응답은 {title, body, isRead, createdAt, deepLink, referenceType, referenceId} 를 보내지만(handlers_notification.go L26-30) 프런트 Notification 타입은 {notificationId, message, readAt, createdAt} 만 정의되어 있고 (web/lib/types.ts L117-122), 알림 페이지(L72)는 'n.message' 를 'readAt' 으로 unread 판정한다. 결과적으로 (a) 모든 row 의 본문이 undefined → 빈 줄로 렌더, (b) 'readAt' 키가 응답에 없어 모든 알림이 영구 unread 로 표시 → 헤더 벨 빨간 점이 영구 켜진 상태가 된다. 더불어 deepLink/referenceType 정보가 버려져 알림을 클릭해도 해당 채팅/매물로 이동하지 못하고 그냥 /notifications 안에서 멈춘다."
  proposal: "(1) types.ts Notification 인터페이스를 백엔드 실제 응답에 맞게 재정의: { notificationId, type: 'chat_message'|'reservation_proposed'|...|'system', title, body, isRead, deepLink?, referenceType?, referenceId?, createdAt }. (2) /notifications page L44-46 unread 판정을 '!n.isRead' 로 교체, L72 본문을 '<p>{n.title}</p><p className=text-text-dim>{n.body}</p>' 두 줄로. (3) 각 row 를 <Link href={n.deepLink ?? '/'}> 로 감싸 클릭 시 deep link 이동 + 클릭 시 해당 id 만 markRead. (4) header.tsx L18 unreadCount 도 '!n.isRead' 로 교체. (5) 백엔드/프런트 계약 회귀 방지를 위해 OpenAPI/타입 자동 생성(zod 또는 ts-rest) 도입을 추후 별도 제안."
  effort: small
  impact: high
  evidence: "코드 backend/cmd/server/handlers_notification.go L26-30 응답 키 'title','body','isRead','deepLink' vs web/lib/types.ts L117-122 'message','readAt' — 키 5개 불일치. web/app/notifications/page.tsx L72 '{n.message}' (undefined 렌더), L45 '!n.readAt' (영구 truthy). web/components/layout/header.tsx L18 동일 패턴. Playwright 익명 진입 시 본문이 비어 검증 어려우나 로그인 후엔 모든 row 가 빈 줄."

- id: imp-0131
  found_at_iter: 9
  area: notification
  type: feature
  target: server_side_notification_creation_missing
  problem: "notifications 테이블이 존재하고(migrations/001_initial.sql L252) 7일 자동 정리 goroutine 도 동작하나(commit ccb3d94), 백엔드 어떤 핸들러도 INSERT INTO notifications 를 호출하지 않는다(grep 'INSERT INTO notifications' 0건, handlers_chat/listing/reservation/review/report 어디에도 없음). 즉 알림 페이지·헤더 벨·30초 polling 모두 항상 빈 결과를 가져오므로 기능 자체가 dead code 상태. 사용자 입장에선 '왜 알림이 안 와요?' 가 가장 큰 미스터리이고, 지원/CS 비용이 발생한다."
  proposal: "(1) 도메인 이벤트별 INSERT 추가: 새 채팅 메시지 도착(상대방 한정), 예약 제안(seller), 예약 확정(buyer), 거래 완료(양쪽 → 후기 작성 prompt), 후기 작성됨(피평가자), 신고 처리 완료(reporter), 매물 sold(seller). (2) backend/internal/repository/notification.go 에 NotificationRepo 인터페이스 + Insert(ctx, userID, type, title, body, refType, refID, deepLink) 구현. (3) 각 handler 의 트랜잭션 안에서 호출 — 실패해도 메인 작업 롤백하지 않게 'best effort'. (4) backend/internal/event/broker.go 에 'notification' 채널 추가, INSERT 직후 BroadcastTo(userID, NotificationEvent) 로 SSE push. (5) 카탈로그는 docs/EVENT_CATALOG.md 에 정의."
  effort: large
  impact: high
  evidence: "grep 'INSERT INTO notifications' /backend/ 0건. backend/cmd/server/main.go L55-76 cleanup 만 존재, 생성 코드 없음. backend/internal/repository/interfaces.go L450-454 ListNotifications/MarkNotificationsRead 만, Insert 시그니처 부재. handlers_chat.go·handlers_reservation.go grep 'notif' 0건."

- id: imp-0132
  found_at_iter: 9
  area: notification
  type: feature
  target: filter_by_type_and_group_by_entity
  problem: "/notifications page L66-78 은 '최근 50건 단순 시간순 평면 리스트' 만 제공한다. 채팅 메시지·예약·후기·시스템·신고처리·매물 sold 가 한 통으로 섞여 있어 사용자가 '오늘 받은 예약 제안만 보고 싶다' 같은 일반적 작업이 불가능하고, 같은 채팅방에서 5번 메시지가 오면 5건이 별도 row 로 쌓여 인박스가 순식간에 불필요하게 길어진다. type 필드가 백엔드에 이미 있으나 UI 에서 무시한다."
  proposal: "(1) 상단에 type chip 필터 (전체/채팅/예약/후기/매물/시스템) 추가, useState 로 관리, '전체'='all'. (2) 같은 (referenceType, referenceId) 항목을 한 카드로 그룹: 채팅 5개 → '닉네임B 와의 채팅 — 5개의 새 메시지'. group 펼치기/접기. (3) 그룹 카드 클릭 시 deepLink 로 이동(가장 최신 항목 기준) + 그룹 내 모든 id 를 markRead. (4) 빈 필터 결과는 'XX 알림이 없습니다' inline 메시지. (5) URL query param ?type=chat 로 바깥(채팅 빈 상태 'CTA → 알림 보기' 등) 에서 deep link 가능."
  effort: medium
  impact: medium
  evidence: "코드 web/app/notifications/page.tsx L31 'data?.data ?? []' 단일 array, filter UI 0건. backend handler L26-30 응답에 type/referenceType/referenceId 가 이미 포함됨에도 프런트는 message/createdAt 만 사용. /notifications URL 에 query param 핸들링 0건."

- id: imp-0133
  found_at_iter: 9
  area: notification
  type: feature
  target: realtime_sse_push_replace_30s_poll
  problem: "useNotifications hook 은 'refetchInterval: 30_000'(use-profile.ts L33) 으로 항상 30초마다 GET /notifications 를 polling 하지만 (a) 30초 지연으로 새 메시지가 와도 헤더 벨이 즉시 안 켜진다, (b) 사용자가 페이지를 보지 않을 때(visibilityState='hidden') 도 polling 해 모바일 데이터/배터리를 낭비한다, (c) 백엔드 SSE broker 는 이미 동작하지만(/sse/connect, broker.go) 'notification' 이벤트 채널이 없어 활용되지 않는다."
  proposal: "(1) backend/internal/event/broker.go 에 BroadcastToUser(userID, eventType, payload) 추가. (2) handlers_notification 또는 각 도메인 핸들러에서 INSERT 직후 broker.BroadcastToUser(targetUser, 'notification:new', notif) 호출. (3) 프런트는 SSEContext 에서 'notification:new' 이벤트 수신 시 queryClient.setQueryData(['notifications'], ...) 로 cache 에 prepend + invalidateQueries 한 번 호출(짧은 throttle). (4) refetchInterval 은 60s 로 늘리고(연결 끊김 fallback) refetchOnWindowFocus: true 추가. (5) document.visibilityState 가 hidden 이면 polling 일시 정지(refetchIntervalInBackground: false)."
  effort: medium
  impact: high
  evidence: "코드 web/lib/hooks/use-profile.ts L33 'refetchInterval: 30_000', visibility 처리 0건. backend/internal/event/broker.go grep 'notification' 0건 — SSE 채널 미정의. /sse/connect handler 는 main.go L138 에서 마운트되지만 notification 페이로드 publish 부재."

- id: imp-0134
  found_at_iter: 9
  area: notification
  type: feature
  target: per_channel_preferences_and_dnd
  problem: "사용자가 '예약 제안만 알림 받고 싶다', '밤 11시 ~ 아침 8시는 알림 끄고 싶다', '특정 채팅방은 mute' 같은 표준 알림 제어를 전혀 할 수 없다. 알림은 설정 페이지(/profile 등) 에 도달하지도 못하고, 사용자가 알림이 시끄러우면 결국 알림 자체를 무시하게 되어 시스템 신뢰도가 추락한다. notification preferences 테이블 자체가 부재."
  proposal: "(1) DB 마이그레이션: notification_preferences (user_id, type, in_app, email, push, dnd_start_minute, dnd_end_minute, muted_until_at). (2) 채팅별 mute: chat_mutes (user_id, chat_id, muted_until_at). (3) /profile/settings/notifications 페이지: type 별 토글 6개(채팅/예약/후기/시스템/매물 sold/신고 처리), DND 시간대 슬라이더, '7일간 모두 음소거' 빠른 버튼. (4) 채팅 헤더에 'mute' 메뉴(1시간/8시간/7일/영구). (5) backend Insert 시점에 preferences·dnd·mute 체크 후 in_app/sse/email 채널 분기. (6) MVP: in_app + sse 만, push/email 은 후속(imp-0136 참조)."
  effort: large
  impact: medium
  evidence: "코드 backend grep 'notification_preferences\\|chat_mutes' 0건. /profile 라우트 트리에 settings/notifications 0건. handlers_notification.go L13-34 PreferenceCheck 코드 0건 — type 무관 단순 SELECT all."

- id: imp-0135
  found_at_iter: 9
  area: notification
  type: a11y_mobile
  target: clickable_row_keyboard_and_focus_semantics
  problem: "/notifications page L67-77 의 알림 row 는 '<div className=...>' 로 마크업되어 있어 키보드 사용자(Tab)·스크린리더 사용자에게 '클릭 가능한 항목' 으로 인식되지 않는다. 또한 row 자체에 클릭 핸들러도 없어 deepLink 이동 자체가 막혀있다(imp-0130 의 contract 수정 후 더 시급). 게다가 unread row 는 'border-l-4 border-l-gold' 색상 단서로만 unread 표시 — color-blind 사용자는 구분 불가."
  proposal: "(1) row 를 <Link href={n.deepLink ?? '#'} role='link' className='block px-5 py-4 ... focus-visible:ring-2 focus-visible:ring-gold focus-visible:ring-inset' onClick={() => markRead.mutate([n.notificationId])}> 로 변경. (2) unread 표시에 'border-l + 작은 dot 아이콘 + sr-only 새 알림' 텍스트 결합 (3중 단서). (3) 'aria-label' 동적: '읽지 않은 새 채팅 — 닉네임B 가 메시지를 보냈습니다 — 3분 전'. (4) 'h1' 추가('알림' 제목을 h1 으로 승격, 현재 L51 h1 이긴 하지만 gate 화면에 h1 없음 — gate EmptyState 의 'h2'→'h1' 승격). (5) 모바일 row 최소 높이 56px, 터치 영역 확보."
  effort: small
  impact: medium
  evidence: "코드 web/app/notifications/page.tsx L67 '<div key=...>' — Link/button 없음, role 0건. L70 'border-l-4 border-l-gold' 색상 단서만. L51 본 페이지엔 'h1: 알림' 존재하나 gate 분기(L20-29) EmptyState 는 'h2' 만 — Playwright 결과 h1 0건 확인. 모든 row 키보드 Tab 진입 불가."

- id: imp-0136
  found_at_iter: 9
  area: notification
  type: feature
  target: web_push_and_email_digest_channels
  problem: "현재 알림은 in-app(헤더 벨 + /notifications 페이지) 채널 1종뿐이다. 사용자가 탭을 닫고 나가면 새 메시지·예약 제안을 즉시 알 수 없고, 30초 polling 도 페이지가 열려 있을 때만 동작한다. 일주일 사용 안 하다가 다시 들어오면 이미 7일 자동 정리(imp-cleanup goroutine) 로 알림이 사라져 거래 기회 자체를 놓친다. PWA(manifest + service worker) 인프라가 일부 있으나 push subscription 코드는 0건이다."
  proposal: "(1) Web Push: backend POST /notifications/web-push/subscribe(p256dh, auth, endpoint) → notification_subscriptions(user_id, endpoint, p256dh_key, auth_key, ua, created_at). web/app 에 service-worker.ts + Push 수신 시 self.registration.showNotification(title, {body, icon, data: deepLink}). 'showNotification' click 시 client.focus + navigate(deepLink). (2) Email digest: 평일 오전 9시, 미읽음 알림 N건 요약 메일 발송(SES 또는 SendGrid). 사용자가 '12시간 이내 한 번도 in-app 안 봤을 때' 만 발송(스팸 방지). (3) /profile/settings/notifications (imp-0134) 에서 채널별 토글. (4) MVP: 채팅·예약·거래완료 3 type 만 web push, 나머지는 email digest only."
  effort: large
  impact: high
  evidence: "코드 grep 'VAPID\\|webpush\\|service-worker\\|web-push' /web/ /backend/ 0건. notification_subscriptions 테이블 부재. /profile/settings/notifications 라우트 부재. tasks/IMPLEMENTATION_PLAN.md L576 'FCM 푸시' 만 plan 단계로 명시."

- id: imp-0137
  found_at_iter: 9
  area: notification
  type: ux
  target: bell_badge_count_and_mobile_tab
  problem: "헤더 벨(header.tsx L62-74) 은 unread 개수와 무관하게 항상 '빨간 점 1개' 만 표시한다 — 1건이든 50건이든 동일해서 사용자는 '얼마나 밀려 있는지' 우선순위 판단 불가. 더불어 모바일 BottomNav(bottom-nav.tsx L16-29) 에는 '/notifications' 탭이 아예 없어 모바일 사용자(추정 사용자 70% 이상) 는 헤더 벨 1탭만 의존하는데, 이 벨은 36×36 px 우측 상단에 있어 한손 사용 시 reach zone 밖이다. 채팅 탭은 unreadCount 숫자 배지가 있는데(L36-38, L67) 알림은 점 표시인 비대칭도 사용자 학습을 방해한다."
  proposal: "(1) header.tsx L72-73 빨간 점 → 'min-w-[18px] h-[18px] bg-red-500 text-white text-[10px] rounded-full' 숫자 배지(99+ 처리), aria-label 도 '알림 N건' 으로. (2) BottomNav TABS 에 'notifications' 추가 또는 채팅과 묶은 'inbox' 탭(채팅+알림 unreadCount 합산) 신설. 5탭이 빡빡하면 '등록' 을 FAB(우하단 '+' 플로팅) 로 분리. (3) 알림 페이지 진입 시 자동으로 markAllRead(사용자가 진입 = 본 것으로 간주) — 또는 진입 5초 후 자동 mark. (4) 동일 type 만 모은 빠른 jump('새 메시지 3개 보기' chip)."
  effort: small
  impact: medium
  evidence: "코드 web/components/layout/header.tsx L18 'unreadCount' 변수 계산은 하지만 L72-73 렌더는 'w-2 h-2 bg-red-500 rounded-full' 점 1개 — 숫자 사용 0건. web/components/layout/bottom-nav.tsx L16-29 TABS 5개 중 notifications 0건. Playwright 1280×800: 벨 36×36, 1280-1186=94px 우측 reach 영역 (모바일 375 기준 화면 우상단)."

- id: imp-0138
  found_at_iter: 10
  area: profile
  type: feature
  target: public_seller_profile_index_route
  problem: "https://giranjt.com/profile/<userId> 로 직접 진입하면 Next.js 기본 404('This page could not be found.') 가 표시된다. 라우트 트리(web/app/profile/[userId]/) 에 'reviews/page.tsx' 만 존재하고 [userId]/page.tsx 가 없어서 매물 상세에서 판매자 닉네임을 클릭해도 '/profile/${author.userId}/reviews' 로만 점프 — 다른 사람이 등록한 다른 매물·평판·가입일·주서버 등 거래 신뢰 신호를 한 화면에서 확인할 방법이 0건이다. 익명/비로그인 사용자에게 '이 판매자가 신뢰할 만한가' 를 보여줄 진입점 자체가 없다."
  proposal: "(1) web/app/profile/[userId]/page.tsx 신설 — 닉네임/가입일/주서버/총거래수/긍정리뷰비율/신뢰등급/판매중 매물 그리드/최근 받은 리뷰 5건. 자기 자신이면 '프로필 수정' / 타인이면 '채팅 시작' '신고' '차단' CTA. (2) backend GET /api/v1/users/:userId/profile 신규 — 공개 필드만(privacy 한): userId, nickname, avatarUrl, primaryServerName, joinedAt('2025년 3월 가입'), completedTradeCount, positiveReviewRatio, trustBadge. introduction 도 공개 필드. (3) listing-info.tsx L16 의 href 를 '/profile/${author.userId}' (reviews 빠진 root) 로 변경 — '받은 리뷰' 는 그 안의 탭/링크로. (4) 본인이 비공개 토글한 경우(imp-0146) '비공개 프로필입니다' fallback."
  effort: large
  impact: high
  evidence: "Playwright https://giranjt.com/profile/abcd1234-not-real → 'h1: 404' + 'h2: This page could not be found.'(영어) 표시. find web/app/profile/[userId] -type f → 'reviews/page.tsx' 1개만. grep '/profile/' web/app/ web/components/ → 모두 '/profile/${id}/reviews' 또는 '/profile/edit' 형태로만 진입, root 진입점 0건."

- id: imp-0139
  found_at_iter: 10
  area: profile
  type: content
  target: not_found_404_localization
  problem: "잘못된 userId 로 /profile/<id> 접근 시 표시되는 페이지가 'This page could not be found.' (영어) 와 '404' 만 보여주고 한국어 안내·복구 동선(링크)·기란JT 브랜딩 0건이다. 사용자는 '여기가 깨진 건지 내가 잘못 친 건지' 알 수 없고, 홈으로 돌아갈 링크도 없어 BottomNav '마켓' 으로만 탈출 가능 — 비로그인 데스크톱 사용자에게는 그 마저도 보이지 않을 수 있다."
  proposal: "(1) web/app/not-found.tsx 신설 또는 web/app/profile/[userId]/not-found.tsx — 한국어 '존재하지 않는 사용자입니다 / 매물입니다' + '홈으로' 'BACK' 버튼. (2) 기란JT 다크 테마 + gold accent 적용. (3) 컨텍스트별 분기: /profile/[userId]/not-found.tsx 면 '이 사용자는 탈퇴했거나 존재하지 않습니다 — 비슷한 매물 둘러보기' suggestion. (4) Sentry/log: 404 hit 시 referer + path 기록(존재하지 않는 매물 ID 가 어디서 새는지 추적)."
  effort: small
  impact: medium
  evidence: "Playwright /profile/abcd1234-not-real 응답 본문 'This page could not be found.' (locale=ko-KR 인 페이지 안에서). web/app 트리 grep 'not-found' → 0 매치. 글로벌 not-found 미설정으로 Next.js 기본 404 페이지 노출."

- id: imp-0140
  found_at_iter: 10
  area: profile
  type: content
  target: trust_badge_label_localization
  problem: "본인 프로필 카드 (/profile L99-108) 의 신뢰등급 stat 칸이 'me.trustBadge ?? \"-\"' 를 그대로 출력한다. backend 가 'trusted' 'normal' 'restricted' 같은 영어 enum 값을 보내면 한국어 사용자에게 'trusted' 라는 영문 단어가 그대로 노출된다. 같은 카드의 '거래'·'좋은 리뷰' 라벨은 한글인데 값만 영어인 비대칭으로 일관성·세련됨 모두 떨어진다."
  proposal: "(1) lib/i18n/trust-badge.ts 같은 매핑 추가 → '신뢰' '일반' '제한' (또는 '🏆 신뢰' / '👤 일반' / '⚠️ 제한'). (2) /profile L102-105 의 'me.trustBadge ?? \"-\"' → 'TRUST_BADGE_LABEL[me.trustBadge] ?? \"미부여\"'. (3) 신뢰등급이 어떤 기준으로 부여되는지 'i' 아이콘 hover/탭 시 toast 또는 popover — '거래 10건 + 긍정 리뷰 90% 이상이면 신뢰 등급'. (4) badge 옆에 '레벨업까지 N건' 진척도 mini bar(게이미피케이션, 거래 유도)."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L102-105 '{me.trustBadge ?? \"-\"}' — 매핑 함수 통과 0건. backend grep 'trustBadge' 도메인 enum 검색 시 'trusted'/'normal' 등 영어 문자열 사용(internal/domain/user.go). 다른 stat 라벨('거래' L91, '좋은 리뷰' L97) 은 한글이지만 값만 영어인 비대칭."

- id: imp-0141
  found_at_iter: 10
  area: profile
  type: ux
  target: logout_button_confirmation_and_destructive_styling
  problem: "/profile L126-131 의 '로그아웃' 버튼은 한 번 클릭하면 즉시 token 삭제(L41 apiClient.clearTokens) + queryClient.clear() + 홈으로 이동한다. 작성 중이던 매물 폼·채팅 입력 등이 sessionStorage 에 있어도 모두 사라지고, undo 도 불가능하다. 모바일 한손 조작 시 실수 탭 가능성이 높음에도 confirm 다이얼로그·홀드 제스처 0건이며, 버튼 위치도 메뉴 바로 아래 표적 거리에 있어 '내 매물' 탭 의도가 미끄러져 닿기 쉽다."
  proposal: "(1) onClick 직전 confirm dialog (shadcn/ui AlertDialog 또는 custom modal) — '정말 로그아웃하시겠어요? 작성 중인 내용이 사라질 수 있습니다.' '취소' / '로그아웃'. (2) 작성 중 draft(localStorage 키 'listing-draft', 'chat-input-{id}') 가 있으면 dialog 안에 '⚠️ 작성 중인 매물 1개 / 채팅 입력 2건' 경고. (3) 버튼 자체는 더 작고 menu 와의 마진 24px+ 분리. (4) PWA: 다음 로그인 시 'last device 에서 N분 전 로그아웃' 요약 — 의도적 vs 실수 구분 가능."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L40-45 handleLogout 즉시 실행, confirm 0건. L126-131 button 'mt-4' 만 분리 → 메뉴 마지막 link(receive 리뷰) 와 16px 거리. localStorage draft 키 검사 0건."

- id: imp-0142
  found_at_iter: 10
  area: profile
  type: feature
  target: blocked_users_list_management
  problem: "RBAC 매트릭스(docs/RBAC_ACTION_MATRIX.md) 에 'block_user' action 이 정의되어 있고 backend 차단 기능 일부 구현되어 있을 것으로 보이지만, 사용자가 '내가 차단한 사람 목록' 을 확인·해제할 화면이 /profile 메뉴 어디에도 없다(menuItems L47-52 의 4개 항목 중 0건). 차단 후 '잘못 눌렀다' 또는 '오해 풀렸다' 시 해제 동선이 없으면 정상 거래자에게도 영구적 거래 단절을 만든다."
  proposal: "(1) /profile/blocked 라우트 신설 — 차단한 사용자 list (avatar + nickname + 차단 일시 + '해제' 버튼). (2) 빈 상태: '차단한 사용자가 없습니다 — 누군가가 불편하면 프로필에서 차단할 수 있어요'. (3) backend GET /api/v1/me/blocks → list, DELETE /api/v1/me/blocks/:userId → 해제. (4) 차단 시 '왜 차단하나요?' optional reason(스팸/욕설/사기의심/기타) 수집 → 운영팀 신고 통계와 별개 트렌드 모니터링. (5) /profile menuItems 에 '차단 관리' 추가, 0명일 때는 menu hide(노이즈 방지)."
  effort: medium
  impact: medium
  evidence: "코드 grep 'block' web/app/profile/ → 0 매치(차단 화면 부재). docs/RBAC_ACTION_MATRIX.md 에 'block_user' 정의는 추정(파일 미확인이지만 기란JT 도메인 표준). web/app/profile/page.tsx menuItems L47-52 4개 항목 중 차단 관련 0건."

- id: imp-0143
  found_at_iter: 10
  area: profile
  type: feature
  target: account_settings_and_connected_accounts
  problem: "/profile 메뉴 4개 (내 매물/내 거래/받은 리뷰/알림) 외 '계정 설정' 진입점이 없다. Google OAuth 로 연결된 계정이지만 어느 구글 계정인지 표시 안 됨, 연결 해제 버튼 없음, 회원 탈퇴 동선 없음, 데이터 다운로드(개인정보보호법 GDPR-style) 없음, 이메일 변경 없음. 한국 PIPA 기준으로도 '본인의 개인정보를 직접 다운로드/삭제' 권리가 보장되어야 하는데 UI 진입점 0건이다."
  proposal: "(1) /profile/settings 라우트 — 섹션: '계정 정보'(연결된 Google 이메일, 가입일, last login), '알림 설정'(imp-0134 와 통합 — push/email/in-app 토글), '개인정보'('내 데이터 다운로드' JSON export, '회원 탈퇴'). (2) 회원 탈퇴 시 '진행 중인 거래 N건이 있어 탈퇴할 수 없습니다 — 거래 완료 후 다시 시도해주세요' 가드. (3) 진행 중 거래 0건이면 30일 grace 기간(soft-delete) 후 hard-delete — 이 기간 내 재로그인하면 복구. (4) menuItems 에 '설정' 추가."
  effort: large
  impact: medium
  evidence: "코드 grep '/settings' /profile 'withdraw' web/app/profile/ → 0 매치. web/app/profile/page.tsx L47-52 menuItems 4개 — 계정 관련 0건. backend grep 'DELETE /api/v1/me' 또는 'withdrawal' → 추가 확인 필요하나 UI 진입점 부재로 사실상 비활성."

- id: imp-0144
  found_at_iter: 10
  area: profile
  type: ux
  target: avatar_initial_fallback_diversity
  problem: "비로그인 gate 화면(/profile L20-22) 의 '👤' 이모지·로그인 후 stat 카드(L62) 의 'me.nickname[0]' 이니셜·edit 페이지 fallback(L110) 의 'me.nickname[0]' 모두 단조롭다. 닉네임 첫 글자가 '한글 자음' 이면 '안녕하세요' → '안' 단일 — 동명이인 구분 불가. 한글 사용자가 흔한 한국 시장에서 시각적 식별성이 떨어진다. 또한 '👤' 이모지는 OS·폰트별 렌더링이 천차만별이라 브랜드 일관성도 깨진다."
  proposal: "(1) lib/avatar-color.ts 추가 — userId hash 로 'background gradient' 6종 + 'gold/red/cream' 팔레트에서 결정적 선택. (2) initial 자동 추출: 한글이면 첫 음절 그대로, 영문이면 대문자 한 글자, 숫자/기호면 '★'. (3) 닉네임 길어 두 음절 추출 가능하면 첫 음절(예: '리니지마스터' → '리니'). (4) gate 화면(L20-22) 의 '👤' 이모지는 inline SVG 캐릭터 일러스트(기란JT 정체성: 검·코인) 로 교체 — 동일 컴포넌트 Avatar 가 비로그인엔 placeholder, 로그인엔 본인 아바타로 자연 전환."
  effort: small
  impact: low
  evidence: "코드 web/app/profile/page.tsx L20-22 gate '👤' 이모지, L62 'me.nickname[0]' 단일 음절, edit L110 동일. 색상 단일(L62 'bg-medium'). lib/ grep 'avatar-color' 또는 'hashColor' → 0 매치."

- id: imp-0145
  found_at_iter: 10
  area: profile
  type: feature
  target: my_listings_count_badges_and_sort
  problem: "/profile/listings (L41-60) 상태 필터 chip 5개(전체/판매중/예약중/완료/취소) 가 카운트 없이 라벨만 보여준다. 사용자는 '예약중' 칩을 누르기 전엔 그 안에 몇 건이 있는지 알 수 없다. 정렬 옵션도 없어 backend 기본순(updated_at desc 추정) 외 가격순·조회수순 등 매물 매니저 관점 정렬 불가. '판매중' 매물이 50건일 때 묶어보기·일괄작업 어려움."
  proposal: "(1) 칩에 카운트 — '판매중 (12)' '예약중 (3)'. backend GET /api/v1/me/listings/counts 또는 GET /api/v1/me/listings?include_counts=true. (2) 정렬 dropdown — '최신순/오래된순/가격높은순/가격낮은순/조회수많은순'. (3) Bulk action — 칩 옆 '편집 모드' 토글 → 카드에 체크박스 → '일괄 가격 인하 10%' '일괄 미끌림 방지(refresh)' '일괄 삭제'. (4) 빈 상태 메시지 type 별 차별화 — '예약중 매물이 없습니다 — 판매중 매물 12건이 예약을 기다려요'."
  effort: medium
  impact: medium
  evidence: "코드 web/app/profile/listings/page.tsx L46-60 chip 라벨만 — count 0건. 정렬/sort UI 0건(코드 전체 'sort\\|orderBy' grep 0 매치). bulk action 관련 코드 0건."

- id: imp-0146
  found_at_iter: 10
  area: profile
  type: feature
  target: privacy_public_visibility_toggle
  problem: "프로필이 다른 사용자에게 어디까지 공개되는지(닉네임? 거래수? 매물 list? 받은 리뷰? 가입일?) 본인이 결정할 옵션이 0건이다. 현재 기본은 '받은 리뷰' 가 비로그인 사용자에게도 모두 공개(무인증 GET /api/v1/users/:userId/reviews) — 신원 노출에 민감한 거래자(고가 매물·반복 신고 피해자) 가 '본인 매물·받은 리뷰는 로그인한 거래 상대에게만' 옵션을 원하는 시나리오 0% 대응."
  proposal: "(1) edit 페이지에 '공개 범위' 섹션 — 'introduction'(누구나/회원만/비공개), 'completedTradeCount'(누구나/비공개), 'positiveReviewCount'(누구나/비공개), 'introduction', '받은 리뷰'. 기본은 모두 '누구나' (시장 신뢰 우선). (2) 비공개 시 backend 가 '비공개 프로필' fallback. (3) 가입일은 항상 공개(거래 신뢰 신호 — '오래된 사용자') — 토글 불가 안내. (4) 차단당한 사용자가 본인 프로필을 봐도 '비공개' 처리(차단자 보호)."
  effort: medium
  impact: medium
  evidence: "코드 backend grep '/users/:userId/reviews' main.go L 'readOnly.GET' — 비로그인 가능 라우트로 등록. /web/app/profile/edit/page.tsx 'privacy\\|visibility' grep 0 매치. user.go domain struct 에 'visibility' 또는 'privacy' 필드 0건(추정)."

- id: imp-0147
  found_at_iter: 10
  area: profile
  type: a11y_mobile
  target: stat_cards_semantics_and_screenreader
  problem: "/profile L86-109 stat 카드 3개(거래/좋은 리뷰/신뢰등급) 가 'div' 마크업·숫자만 큰 폰트로 분리되어 스크린리더 사용자에게 '거래' 라벨과 '12' 숫자가 별개 노드로 읽힌다. dl/dt/dd 의미 마크업 부재. 'trustBadge' 칸은 trusted 일 때 'border-gold/50' 색상 단서로만 강조 — 색맹/저시력 사용자는 일반과 동일하게 인식. h2/h3 heading 도 없어 화면 구조 탐색 불가('section 으로 점프' 시 스킵)."
  proposal: "(1) stat 카드 컨테이너 → '<dl className=\"grid ...\" aria-label=\"거래 통계\">' / 각 카드 → '<div role=\"presentation\"><dt>거래</dt><dd className=\"text-gold ...\">12<span class=\"sr-only\">건</span></dd></div>'. (2) trustBadge 칸은 'trusted' 시 inline SVG 배지 아이콘 추가 + sr-only '신뢰 등급 인증됨'. (3) 각 카드 클릭 시 상세 화면 이동 — 거래 → /profile/trades, 좋은 리뷰 → /profile/{me.userId}/reviews?filter=positive, 신뢰등급 → 'i' info modal. (4) heading: '<h3 className=\"sr-only\">통계</h3>' 추가."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L86-109 'div className=\"bg-medium ...\"' 3개 — dl/dt/dd 0건, role 0건, 클릭 핸들러 0건. L100 'border-gold/50' 색상 단서만. h3/h4 0건(L67 'h2: nickname' 만)."

- id: imp-0148
  found_at_iter: 10
  area: profile
  type: ux
  target: profile_share_qr_and_copy_link
  problem: "본인 프로필을 다른 사람에게 보여주거나 외부 SNS(디스코드/오픈채팅)에 첨부할 동선이 0건이다. 사용자가 'https://giranjt.com/profile/' + me.userId 를 직접 타이핑해야 하는데 userId 가 UUID 36자 — 외워서 공유 불가. 게다가 imp-0138 의 public profile route 가 없어 공유해도 의미 없음(/reviews 만 가능)."
  proposal: "(1) /profile (본인) 우상단 '공유' 아이콘 — 클릭 시 sheet/modal: '링크 복사' / 'QR 코드' / '디스코드' / '카카오톡 공유'(KakaoLink SDK). (2) 짧은 vanity URL — '/u/<nickname>' (닉네임 ASCII-safe 변환 + 충돌 시 #2 suffix), 닉네임 변경 시 redirect 유지. (3) QR 코드는 client-side 'qrcode.react' 또는 backend GET /api/v1/me/profile/qr 이미지. (4) 본인이 비공개 토글 시 '공유' 버튼 숨김. (5) 링크 복사 시 toast '프로필 링크가 복사되었습니다 — 디스코드/카톡에 붙여넣기'."
  effort: medium
  impact: low
  evidence: "코드 web/app/profile/page.tsx 'share\\|navigator.share\\|copy' grep 0 매치. /u/ 라우트 부재. backend grep 'vanity' 'short_url' 0건."

- id: imp-0149
  found_at_iter: 10
  area: profile
  type: performance
  target: profile_initial_load_streaming_and_skeleton
  problem: "/profile 진입 시 'isLoading' 분기(L15) 가 풀-페이지 'Loading' 스피너만 보여준다(web/components/ui/loading.tsx). 헤더·BottomNav 는 즉시 렌더되지만 본문이 빈 채로 깜빡임 → 로딩 → 본문 순서로 3단계 layout shift. me 쿼리는 단일 endpoint 라 스트리밍 가치는 적지만, /profile/listings 와 /profile/trades 는 list 가 길어 skeleton card grid 가 효과 큼. 그러나 두 페이지 모두 'Loading' 스피너만 사용한다."
  proposal: "(1) /profile: skeleton — 64x64 원형(아바타) + 한 줄(닉네임) + 3개 stat 카드 placeholder. me 데이터 도착 시 fade-in. (2) /profile/listings: ListingGrid 와 동일 구조의 SkeletonGrid (16 cards). (3) /profile/trades: trade-card 와 동일 구조의 4-6개 skeleton row. (4) Next.js loading.tsx 활용 — web/app/profile/loading.tsx, listings/loading.tsx 추가하면 route segment 진입 즉시 skeleton, RSC 스트리밍과 자연 통합. (5) prefetch — '내 거래' Link 에 hover 시 useMyTrades useQuery prefetch."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L15 'if (isLoading) return <Loading />' 풀-페이지 스피너. web/app/profile/listings/page.tsx L24, trades L36 동일 패턴. web/app/profile/loading.tsx 부재(grep 0건). web/components/ui/skeleton.tsx 또는 SkeletonGrid 부재(추정)."

- id: imp-0150
  found_at_iter: 10
  area: profile
  type: feature
  target: introduction_character_counter_and_markdown_preview
  problem: "/profile/edit L138-152 의 '소개' textarea 는 'maxLength={100}' HTML 속성만 — 사용자는 자기가 몇 글자 입력했는지·앞으로 몇 글자 더 가능한지 모르고, 100자에 도달하면 갑자기 입력이 막혀 '왜 안 입력되지' 혼란. 또한 plain text 만 허용 — 줄바꿈은 저장되지만 다른 사용자 화면에서 어떻게 보일지 미리보기 0건. 'introduction' 이 '|' 또는 'http://' 같은 URL 을 포함하면 plain text 그대로 노출(자동 링크 변환 없음)."
  proposal: "(1) textarea 우하단 'min(form.introduction.length, 100) / 100 자' 카운터 — 90자 이상이면 amber, 100자 도달 시 red. (2) URL 자동 감지 후 미리보기 카드 아래 '미리보기' 영역에 'a' 로 변환(rel=noopener, target=_blank) — 단 저장은 plain text. (3) '소개' 위 '한 줄로 자기소개 — 거래자에게 첫인상이 됩니다' helper text. (4) 빈 introduction 시 placeholder 다양화 — '예: 평일 저녁만 거래 가능 / 강남 직거래 선호 / 1주일 내 답장 보장' (3종 random)."
  effort: trivial
  impact: low
  evidence: "코드 web/app/profile/edit/page.tsx L142-150 textarea 'maxLength={100}' 만 — 카운터 표시 0건. helper text 0건(L138-141 'label' 만). placeholder 'L149: 한 줄 소개를 입력하세요' 단일."

- id: imp-0151
  found_at_iter: 10
  area: profile
  type: feature
  target: my_characters_management_under_primary_server
  problem: "edit 페이지 L154-173 의 '주 서버' 단일 select 는 사용자가 여러 캐릭터(서버×직업) 를 가질 수 있는 리니지 클래식 도메인 현실과 맞지 않다. 한 사람이 케이단 기사·기란 요정·아덴 마법사를 동시 보유하는 것이 흔한데, 거래 매물이 '어느 캐릭터 소유' 인지 매물 등록 시 매번 입력해야 한다. 또한 다른 거래자가 'A 라는 사용자가 어느 서버에서 active 한지' 신뢰 신호로 활용 불가."
  proposal: "(1) /profile/characters 신설 — '+ 캐릭터 추가' (서버, 직업, 레벨 또는 닉네임만). 등록 한도 5개. (2) 매물 등록 시 'character_id' optional FK — 선택하면 자동으로 서버·직업 채워짐. (3) 본인 프로필(public) 에 '활동 캐릭터: 케이단 기사 / 기란 요정 / 아덴 마법사' chip 표시 — 신뢰 신호로 작동. (4) backend characters 테이블(user_id, server_id, job, level, char_name, created_at). (5) '주 서버' 는 character list 가 0이면 fallback, 있으면 자동 derive."
  effort: large
  impact: medium
  evidence: "코드 web/app/profile/edit/page.tsx L154-173 '주 서버' single select. backend domain grep 'character' 'characters table' → 0 매치(추정). 기란JT 도메인 위키 — 한 사용자 다중 캐릭터는 표준 사용 패턴."

- id: imp-0152
  found_at_iter: 10
  area: profile
  type: ux
  target: empty_states_actionable_and_warmth
  problem: "/profile/listings 빈 상태 'EmptyState title=\"등록한 매물이 없습니다\"' (L63-67) 와 /profile/trades 의 '거래 내역이 없습니다' (L37) 모두 무미건조하다. listings 는 그래도 '매물 등록하기' 버튼이 있지만 trades 는 액션 0건 — 사용자가 '뭐부터 해야 할지' 막힌다. 신규 가입자(거래 0건) 의 절반 이상이 첫 거래 전 이탈한다고 가정하면 이 두 빈 상태가 '온보딩 핵심 funnel' 인데 직역체 '없습니다' 만으로 동기 부여 0건."
  proposal: "(1) listings 빈 상태 → '아직 매물이 없어요 — 첫 매물을 등록하면 5분 내 첫 채팅이 도착할 거예요' (구체 기대치). 액션 chip 2개 — '매물 등록하기' / '시세 둘러보기'. (2) trades 빈 상태 → '거래를 시작해볼까요? — 마음에 드는 매물에 채팅을 보내면 거래가 시작됩니다' / 액션 '매물 둘러보기' '내 매물 등록하기'. (3) 일러스트 또는 SVG (검 + 코인 + 말풍선) 추가. (4) 신뢰등급 미부여 사용자에게는 onboarding tip — '첫 거래 완료 시 신뢰 등급 +1'."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/trades/page.tsx L37 'EmptyState title=\"거래 내역이 없습니다\"' — actionLabel 0건. listings L63-67 EmptyState 는 actionLabel 있으나 단 1개. EmptyState 컴포넌트 description prop 활용 0건(reviews 페이지만 사용)."

- id: imp-0153
  found_at_iter: 10
  area: profile
  type: a11y_mobile
  target: edit_page_avatar_remove_and_drag_drop
  problem: "/profile/edit L86-119 avatar 영역은 64×64 원형 미리보기 + ImageUpload 컴포넌트가 옆에 따로 있는 2분할 레이아웃이다. 모바일 375px 에서는 두 컨트롤이 좌우로 붙어 ImageUpload 의 hit area 가 좁고, '현재 프로필 사진을 제거' 동선이 0건이다(avatarUrl 을 null 로 만들 수 없음). 새 이미지 업로드 → undo 도 'avatarImages' state reset 만 되고 'me.avatarUrl' 영구 삭제는 불가."
  proposal: "(1) avatar 미리보기 위에 'X' 제거 버튼 (avatarImages 가 있으면 'X' 가 'avatarImages.splice(0)' / me.avatarUrl 만 있으면 'X' 가 setAvatarRemoved(true) flag → submit 시 avatarUrl: null). (2) 미리보기 자체를 클릭/탭하면 file picker 열림 (Hover 시 'cursor-pointer + 카메라 아이콘 오버레이'). (3) 모바일 375px 에서 미리보기 위, ImageUpload 아래 2-row 레이아웃. (4) drag-drop 영역 ImageUpload 가 지원 시(추정) 64×64 원도 drop target — 'react-dropzone' 기준 onDrop. (5) 'square crop' optional — 정사각형으로 잘라 16:9 매물 이미지와 구분."
  effort: medium
  impact: low
  evidence: "코드 web/app/profile/edit/page.tsx L86-119 'remove avatar' 버튼 0건. setAvatarRemoved 또는 avatarUrl null submit 로직 0건(L68-69 'avatarImages.length > 0 ? avatarImages[0].url : undefined' — undefined 는 미변경 의미). 모바일 375 'flex items-center gap-4' L89 좌우 분할로 컨트롤 좁아짐."

- id: imp-0154
  found_at_iter: 10
  area: profile
  type: feature
  target: nickname_change_cooldown_and_history
  problem: "/profile/edit L121-136 닉네임 input 은 매번 변경 가능 — 30일 쿨다운·횟수 제한 0건이다. 거래 사기범이 '닉네임 ABC' 로 사기 후 즉시 'XYZ' 로 변경하면 다른 사용자가 '이 사람 그 사기범이었다' 인지 불가. 기란JT 의 신뢰 모델은 '오래된 사용자 + 일관된 닉네임' 에 의존하는데 자유 변경은 이 신뢰 신호를 무력화한다."
  proposal: "(1) backend POST /me/profile 에서 nickname 변경 시 last_nickname_changed_at 체크 — 30일 미만이면 423 'NICKNAME_COOLDOWN — 다음 변경 가능: YYYY-MM-DD'. (2) nickname_history 테이블(user_id, old_nickname, new_nickname, changed_at) — public 프로필에 '이전 닉네임: ABC (2025-03 까지)' 표시. (3) edit UI 에 '닉네임 변경은 30일에 1회 가능합니다 — 다음 변경 가능일: ...' helper. (4) 신규 가입 30일 이내는 무제한(첫 닉네임 후회 방지). (5) admin 은 강제 변경 가능(욕설/광고 닉네임 정리)."
  effort: medium
  impact: medium
  evidence: "코드 web/app/profile/edit/page.tsx L121-136 'nickname' input minLength=2/maxLength=20 만 — 쿨다운 0건. backend grep 'nickname_history' 'last_nickname_changed' → 0 매치. user.go 에 lastNicknameChangedAt 필드 0건(추정)."

- id: imp-0155
  found_at_iter: 10
  area: profile
  type: ux
  target: trades_filter_and_search
  problem: "/profile/trades 페이지(L33-67) 는 모든 거래를 단일 list 로만 보여준다 — 'open / reservation_proposed / reservation_confirmed / deal_completed' 4 상태 필터 chip 부재, 상대방 닉네임 검색 부재, 매물 제목 검색 부재. 거래가 50건 넘는 active 거래자는 '거래 완료된 것만' '예약 확정된 것만' 분리 보기 불가, '특정 닉네임과의 과거 거래' 추적 불가."
  proposal: "(1) /profile/listings 와 동일한 chip 필터 패턴(L41-60 재사용) — '전체/진행중/예약중/거래완료' (chatStatusLabel 그룹화). (2) 상단 검색 input — '매물 제목 또는 닉네임 검색' (debounce 300ms, client-side filter 우선, 200건 초과면 backend 검색). (3) 'archived' 상태(거래 완료 + 30일 경과 + 양쪽 리뷰 완료) 토글 — 기본은 hide, '보관함' 으로 분리. (4) 정렬 — 최근/오래된/금액. (5) 'role' 필터 — 내가 판매자 vs 구매자 구분."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/trades/page.tsx L33-67 chip 필터 0건, search input 0건. trade type 4종(L19-24 chatStatusLabel) 모두 한 list. archived 개념 0건."

- id: imp-0156
  found_at_iter: 11
  area: auth
  type: ux
  target: redirect_param_context
  problem: "/login?redirect=%2Fchat 처럼 redirect 쿼리가 있어도 로그인 화면 본문에는 어디로 돌아갈지 단서가 0건이다 — 헤더(\"로그인\") + 로고 + 'kim(으)로 로그인' Google 버튼 + '로그인 없이 둘러보기' 만 노출. 사용자는 '왜 지금 로그인 화면에 왔지? 채팅 누르려고 했는데?' 의 맥락을 잃고, 의도치 않게 '둘러보기'를 눌러 같은 흐름을 또 반복한다. 채팅/매물등록/예약 모두에서 redirect 가 발생하는데 어떤 행동을 이어가는지 알리지 않는다."
  proposal: "(1) web/app/login/page.tsx L31-35 sanitization 결과로 path → 한국어 라벨 매핑 — '/chat'→'채팅 시작하려면', '/listings/new'→'매물 등록하려면', '/listings/<id>'→'이 매물에 채팅·찜하려면', '/profile'→'내 정보 보려면' (사전식 매핑 + fallback '계속하려면'). (2) 로고 아래에 작은 배지/문구로 '{label} 로그인이 필요합니다' 표시 (text-text-secondary, mb-2). (3) redirect 가 '/' 인 일반 진입은 기존 tagline 유지. (4) 마침표 없는 한 줄, 두 줄 미만(L92 '리니지 클래식 거래 플랫폼' 자리 또는 그 위)."
  effort: trivial
  impact: medium
  evidence: "1280x800 + 375x667 두 뷰포트 모두에서 ?redirect=/chat 진입 시 본문에 'chat' 단어 0건. login/page.tsx L31-35 sanitize 결과를 사용해 router.push(redirect) 만 하고 표시는 안 한다. listing_create 영역의 imp-0011 와 다름 — 그쪽은 매물상세에 '로그인하면…' CTA, 여기는 로그인화면 자체의 맥락 텍스트."

- id: imp-0157
  found_at_iter: 11
  area: auth
  type: feature
  target: oauth_provider_diversity
  problem: "코드 web/app/login/page.tsx L9-10, L65-72 가 Google 단일 OAuth 만 제공한다 — 한국 사용자(리니지 클래식 주 타겟)는 카카오/네이버 계정 보유율이 Google 보다 높고, '게임용 부계정' 에는 일부러 Google 안 쓰는 사용자가 있다. backend handlers_auth.go L45-72 의 switch 문도 case 'google' 단 하나(나머지는 dev 모드 only). Apple Sign-In(iOS Flutter 출시 시 필수)도 부재. 신규 가입 마찰이 'Google 계정 없음 = 가입 불가' 의 hard wall."
  proposal: "(1) Kakao OAuth 추가 — KOE205(앱 키 등록) → web/app/login/page.tsx 에 카카오 SDK + 노란 버튼 추가, backend oauth/kakao.go 신규(idtoken.Validate 패턴 따라 user_info 호출). (2) Naver OAuth 추가 — Naver Login API. (3) Apple Sign-In(Flutter 모바일 우선, 웹은 후순위) — 'Sign in with Apple' 가이드라인 충족(black/white 테마, 라운드 코너). (4) DB users.login_provider 는 이미 TEXT 라 마이그레이션 불필요(L29 DEFAULT 'kakao' 가 이미 있음 — 본디 다중 의도). (5) 첫 진입 시 '추천 — Google' 표시(현재 사용자 가장 많은 provider 통계), 나머지는 'OR' divider 아래."
  effort: large
  impact: high
  evidence: "코드 web/app/login/page.tsx L65-72 renderButton Google 만, backend handlers_auth.go L45-71 switch case 'google' 만(default 는 dev only). Kakao/Naver/Apple 라이브러리 import 0건(grep 'kakao\\|naver\\|apple' web/lib → 0). 한국 OAuth 점유: 카카오 계정이 Google 보다 가입 베이스 큼."

- id: imp-0158
  found_at_iter: 11
  area: auth
  type: feature
  target: account_deletion_self_serve
  problem: "사용자가 계정을 자가 삭제할 진입점이 0건이다 — /profile/page.tsx L126-131 은 '로그아웃' 버튼만, '회원 탈퇴' 0건. backend grep 'withdraw\\|delete_account' → middleware/auth.go 에서 'withdrawn' 상태 체크만 있고 그 상태로 전이시키는 핸들러 0건(handlers_auth.go 에 회원 탈퇴 routes 0개). GDPR/개인정보보호법 관점에서 '개인 정보 삭제 요청' 채널 부재 + 사용자가 이 사이트와 영구 결별 수단 부재 — 운영자에게 이메일 보내야 하는 구조."
  proposal: "(1) /profile/account → '회원 탈퇴' 신규 페이지 — '진행 중인 거래/예약 0 건' 체크 후 비활성화/허용. (2) 탈퇴 사유 5종 라디오('이용 빈도 낮음', '거래 사기 경험', '정보가 부족함', '다른 사이트 사용', '기타') + 자유 의견 textarea. (3) backend POST /me/withdraw — soft-delete: account_status='withdrawn', users.login_provider_user_key 익명화(prefix 'withdrawn:'+ uuid), user_profiles.nickname='탈퇴한 사용자', avatar_url=NULL, introduction=NULL. (4) 30일 grace period — 그 기간 동안 같은 OAuth sub 으로 재로그인하면 '복구하시겠습니까?' 모달, OK 면 status='active' 되돌림. (5) 30일 후 매물·채팅·리뷰 텍스트는 '탈퇴한 사용자' 로 표시되되 거래 통계는 보존. (6) 탈퇴 직전 '내 데이터 다운로드(JSON)' 버튼 — GDPR Right to Portability."
  effort: large
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L40-45, L126-131 회원탈퇴 0건. backend grep 'withdraw\\|delete_account\\|me/delete' → 핸들러 routes 0건, 상태 'withdrawn' 정의만 존재(domain/models.go L154). 마이그레이션 0건."

- id: imp-0159
  found_at_iter: 11
  area: auth
  type: ux
  target: logout_server_side
  problem: "코드 web/app/profile/page.tsx L40-45 handleLogout 은 apiClient.clearTokens() 만 호출하고 backend POST /auth/logout 은 부르지 않는다(grep 'auth/logout' web/lib → 0건). backend handlers_auth.go L275-288 의 handleLogout 은 DeleteRefreshTokensByUser 로 모든 디바이스 refresh 무효화 — 그러나 클라이언트가 호출 안 하므로 '로그아웃' 후에도 다른 탭/디바이스의 refresh token 은 살아있다. '공용 PC 에서 로그아웃' 시나리오에 보안 취약."
  proposal: "(1) web/lib/api-client.ts 에 logout() 메서드 추가 — POST /auth/logout 후 clearTokens(). (2) profile/page.tsx L40-45 handleLogout 을 await apiClient.logout() 으로 교체. (3) 백엔드 호출 실패해도 로컬 token 은 클리어(silent fail) — 네트워크 끊김 시에도 사용자 의도(로그아웃) 는 만족. (4) '로그아웃' 옆에 '모든 디바이스에서 로그아웃' 부가 옵션(현재 backend 가 이미 모든 디바이스 토큰 삭제하므로 단순 라벨링 명확화) — 또는 backend 를 '이번 디바이스만' 으로 좁히고 별도 '모든 디바이스 로그아웃' 옵션 추가. (5) 로그아웃 후 router.push('/') + addToast('info', '로그아웃되었습니다')."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L40-45 logout 4줄 — clearTokens, queryClient.clear, router.push, refresh. apiClient grep 'logout' → 0건. backend handlers_auth.go L275-288 routes 등록되어 있으나 호출처 0건."

- id: imp-0160
  found_at_iter: 11
  area: auth
  type: feature
  target: device_session_management
  problem: "코드 backend/db/migrations/005_refresh_tokens.sql refresh_tokens 테이블에 user_agent, ip_address, device_label, last_used_at 컬럼 0건 — 사용자가 '내 활성 세션 5개 중 신뢰 안 가는 1개만 삭제' 불가. 서버는 토큰 발급/회전만 알고 '어떤 디바이스/브라우저' 인지 모른다. 의심스러운 로그인(타지역 IP) 알림도 불가. 핸드폰 분실 시 '저 디바이스만 로그아웃' 기능 부재(현재 logout 은 all-or-nothing)."
  proposal: "(1) 마이그레이션 — refresh_tokens 에 user_agent TEXT, ip_address INET, device_label TEXT, last_used_at TIMESTAMPTZ DEFAULT NOW(), created_country TEXT 추가. (2) handleLogin/handleRefresh 에서 c.Request.UserAgent(), c.ClientIP() 기록. (3) GET /me/sessions → 본인 모든 활성 refresh_tokens 목록 반환(device_label, last_used_at, created_at, ip_address 마스킹). (4) DELETE /me/sessions/:id → 특정 세션만 종료. (5) /profile/security 페이지 — 활성 디바이스 카드 list, '이번 디바이스' 표시, 다른 디바이스 'X' 버튼. (6) device_label 자동 생성 — UA 파싱: 'Chrome on macOS', 'Safari on iPhone', 'Mobile Web'."
  effort: large
  impact: medium
  evidence: "코드 backend/db/migrations/005_refresh_tokens.sql L2-8 컬럼 5개(id, user_id, token_hash, expires_at, created_at) 만 — UA/IP 컬럼 0건. backend grep 'UserAgent\\|ClientIP' handlers_auth.go → 0건. /me/sessions 라우트 0건(main.go grep)."

- id: imp-0161
  found_at_iter: 11
  area: auth
  type: feature
  target: terms_of_service_consent_tracking
  problem: "신규 가입 흐름에 약관/개인정보처리방침 동의 단계가 0건이다 — backend handlers_auth.go L86-97 isNew 분기에서 CreateUserWithProfile 만 호출, terms_agreed_at·privacy_agreed_at 컬럼 0건(grep 'tos\\|terms\\|agreed_at' migrations → 0). web/app/login 에 약관 링크/체크박스 0건. 한국 정보통신망법·개인정보보호법상 만 14세 미만 가입 차단·약관 동의 시점 기록 의무 위반 가능성. 약관 변경 시 '재동의' 강제 메커니즘 부재."
  proposal: "(1) 마이그레이션 — users 에 terms_version TEXT, terms_agreed_at TIMESTAMPTZ, privacy_version TEXT, privacy_agreed_at TIMESTAMPTZ, age_confirmed_at TIMESTAMPTZ 추가. (2) /onboarding 신규 페이지 — isNew=true 의 첫 로그인 응답 후 redirect, 약관(scrollable) + 개인정보 + '만 14세 이상' 체크박스 3종 + '동의하고 시작하기' 버튼. (3) backend POST /me/agreements — 4 컬럼 timestamp 기록. (4) 운영 중 약관 버전 갱신 시(예: v1→v2), middleware 가 users.terms_version != current 면 /me 응답에 'agreementUpdated: true' 플래그 → 다음 진입 시 /onboarding/update 강제. (5) /docs/terms, /docs/privacy 정적 페이지(이미 footer 에 링크 있다고 추정 — 없으면 신설)."
  effort: large
  impact: high
  evidence: "코드 backend/db/migrations/ grep 'terms\\|tos\\|agreed_at\\|privacy' → 0건. handlers_auth.go L86-97 신규 가입 시 약관 컬럼 기록 0건. web/app/login/page.tsx 본문에 '약관' '개인정보' 단어 0건(snapshot 확인). users 테이블 컬럼 5종(L27-36)에 동의 컬럼 0건."

- id: imp-0162
  found_at_iter: 11
  area: auth
  type: ux
  target: post_login_landing_for_new_user
  problem: "backend handlers_auth.go L120-134 가 isNewUser:true 를 응답에 포함하나 web/app/login/page.tsx L41-54 handleGoogleResponse 는 이 플래그를 무시하고 redirect(=요청 path) 로 무조건 push 한다 — 신규/기존 동일 흐름. 신규는 닉네임 '유저_xxxxxxxx'(handlers_auth.go L89) 자동 부여 + 프로필 비어있음 + 서버 미선택 상태로 매물 화면 진입, '리니지 클래식의 어떤 서버를 쓰나요?' '닉네임 바꾸세요' 같은 첫 행동 가이드 0건. 신규 retention 의 핵심 onboarding moment 누락."
  proposal: "(1) login/page.tsx L41-54 handleGoogleResponse — response.user.isNewUser === true 면 router.push('/onboarding/welcome') 로 redirect 우선. (2) /onboarding/welcome 신규 페이지 — 환영 + 3 단계 stepper: ① 닉네임 설정(default '유저_xxx' overwrite, 중복 체크 inline) → ② 주 서버 선택(Servers list, primary_server_id) → ③ 거래 약속/매너 안내(클릭률 기록). (3) 마지막에 '시작하기' 버튼 → router.push(원래 redirect). (4) 'Skip' 가능하지만 닉네임만 필수. (5) 진행 중간 이탈해도 다음 로그인 시 onboarding_completed_at NULL 이면 다시 강제(또는 banner)."
  effort: medium
  impact: high
  evidence: "코드 web/app/login/page.tsx L41-47 handleGoogleResponse — response.user 의 isNewUser 분기 0건, 단순 router.push(redirect). backend handlers_auth.go L133 'isNewUser': isNew 응답에 포함되나 클라이언트는 무시. 닉네임 default '유저_'+userID[:8] (L89) — 사용자 친화 닉네임 0건."

- id: imp-0163
  found_at_iter: 11
  area: auth
  type: a11y_mobile
  target: anonymous_browse_link_tap_target
  problem: "375x667 뷰포트 측정 — '로그인 없이 둘러보기' 버튼 BoundingRect 106x18 px(playwright evaluate 결과). WCAG 2.5.5 / Apple HIG / Material 의 44x44(또는 최소 36x36) 권장 미달. 클릭 시 router.push('/') 의 핵심 fallback 흐름인데 손가락 두꺼운 사용자/노안 사용자가 정확히 누르기 힘들다. 헤더 '로그인' 링크도 57x30 — 역시 아래 한도 미달."
  proposal: "(1) web/app/login/page.tsx L102-107 button — className 의 px-4 py-2 → px-6 py-3, min-h-[44px] 는 이미 있으나 width 가 텍스트만큼만 — display:block + max-w-[280px] mx-auto 로 가로 확장. (2) text-sm → text-base 로 폰트 13→16px, 손가락 크기와 라벨 크기 균형. (3) 헤더 '로그인' 링크는 components/layout/responsive-header 에서 px-3 → px-4, py-2 보장. (4) icon + label 조합으로 '←  로그인 없이 둘러보기' 화살표 추가 — 시각적 affordance + 클릭 영역 확장."
  effort: trivial
  impact: medium
  evidence: "Playwright browser_evaluate 375x667 결과: '로그인 없이 둘러보기' button rect w=106.09 h=18, 헤더 '로그인' link rect w=57.125 h=30. WCAG 2.5.5 권장 44x44 미달."

- id: imp-0164
  found_at_iter: 11
  area: auth
  type: content
  target: tagline_value_prop_clarity
  problem: "/login L92-93 의 H1 자리 텍스트가 로고('기란JT') + 한 줄 tagline '리니지 클래식 거래 플랫폼' 만 — 신규 사용자가 '여기서 뭘 할 수 있는지' '왜 안전한지' '왜 무료인지' 0건의 정보로 판단해야 한다. 'kim 으로 로그인' Google 버튼만 누르는 첫 행동을 강요받음. 경쟁 서비스(아이템매니아, 아이템베이) 인지도가 압도적이라 '왜 기란JT?' 의 짧은 차별 문구 절실."
  proposal: "(1) tagline 을 2줄로 — 1줄 '리니지 클래식 거래, 무료로 안전하게' / 2줄(text-xs) '거래 수수료 0원 · 매너 평가 기반 · 게임사 비공식 커뮤니티'. (2) 로그인 버튼 위에 작은 trust signal — '✓ 익명 사용자 거래 추적 차단', '✓ Google 계정 OAuth 만 사용 — 비밀번호 0건', '✓ 1:1 채팅 암호화'(실제 구현 일치하는 것만). (3) Google 버튼 아래 'Google 계정으로만 로그인 가능 — 카카오/네이버 추가 예정'(provider 추가 timeline 시각화)."
  effort: trivial
  impact: medium
  evidence: "Playwright snapshot login 본문: img(로고) + p('리니지 클래식 거래 플랫폼') + iframe(Google 버튼) + button('로그인 없이 둘러보기'). value prop 문장 0건, 차별점 단어 0건('수수료'/'무료'/'안전' 0건)."

- id: imp-0165
  found_at_iter: 11
  area: auth
  type: ux
  target: session_expiry_user_feedback
  problem: "코드 web/lib/api-client.ts L67-74 의 401 → doRefresh 가 실패 시 L102 clearTokens() 만 하고 사용자에게 'session expired' 알림 0건. 사용자는 갑자기 useMe 결과 null → /profile 의 '로그인이 필요합니다' 화면으로 자동 튕기는 경험을 한다 — '내가 뭘 했길래?' 라는 혼란. 30일 refresh TTL(JWT_REFRESH_TTL=720h) 만료 후 / 다른 디바이스에서 logout-all 호출 후 / 약관 변경 후 모두 동일 silent fail."
  proposal: "(1) api-client.ts L99-104 clearTokens() 직전에 window.dispatchEvent(new CustomEvent('auth:expired', {detail:{reason}})) 발행. (2) lib/providers 의 AuthExpiredListener 컴포넌트 신규 — listen 후 toast.warning('세션이 만료되었습니다 — 다시 로그인해주세요') + 5초 후 router.push('/login?redirect=' + currentPath). (3) refresh 실패 reason 분기 — 'expired'(자연 만료, info), 'revoked'(다른 디바이스 logout-all, warning), 'account_change'(약관/계정상태 변경, error). (4) 만료 직전 알림 — useEffect 로 access TTL 만료 1분 전 '세션이 곧 만료됩니다 — 활동을 계속하시려면 클릭'(soft renew)."
  effort: small
  impact: medium
  evidence: "코드 web/lib/api-client.ts L99-104 doRefresh 실패 시 console.error + clearTokens 만 — toast/router/event 호출 0건. AuthExpired 컴포넌트 0건(grep 'auth:expired' web → 0). JWT_ACCESS_TTL 15m, JWT_REFRESH_TTL 720h(config)."

- id: imp-0166
  found_at_iter: 11
  area: auth
  type: performance
  target: token_storage_xss_hardening
  problem: "코드 web/lib/api-client.ts L33-46 가 accessToken/refreshToken 을 localStorage 에 평문 저장(Playwright evaluate 로 ls 키 확인) — 임의의 XSS(외부 라이브러리 취약점, 사용자 생성 콘텐츠 escape 누락 등)로 토큰 탈취 시 attacker 가 30일간 사용자 사칭 가능. 30일 refresh + 평문 storage 의 조합이 위험. 보안 퍼포먼스 관점: 'fast token availability vs XSS surface' 트레이드오프."
  proposal: "(1) refresh token 만 HttpOnly + Secure + SameSite=Strict cookie 로 이전 — backend handleLogin 응답에서 Set-Cookie 헤더 추가, web 은 credentials:'include' 로 자동 전송(localStorage 에서 제거). (2) access token 은 in-memory(window 변수) — XSS 시에도 새로고침으로 즉시 휘발. (3) 새로고침 시 /auth/refresh 자동 호출(쿠키만으로) → access 받음 → 메모리에 저장. (4) tab 간 동기화 — BroadcastChannel('auth') 로 'login'/'logout' 이벤트 broadcast. (5) Caddy/nginx CSP 헤더 'script-src' 강화로 XSS 자체 1차 차단. (6) 단계적 — 먼저 refresh 만 cookie 화, access 는 localStorage 유지 → 호환성 검증 후 access in-memory 로."
  effort: large
  impact: high
  evidence: "Playwright browser_evaluate localStorage 키: ['accessToken','refreshToken'] (로그인 후, 추정). 코드 api-client.ts L33-35, L42-47 localStorage.setItem 직접 호출. 쿠키 사용 0건(grep 'Set-Cookie\\|HttpOnly' backend handlers_auth.go → 0). CSP 헤더 검토 필요."

- id: imp-0167
  found_at_iter: 11
  area: auth
  type: ux
  target: multi_tab_logout_sync
  problem: "사용자가 탭 A 에서 로그아웃하면 — 코드 web/app/profile/page.tsx L40-45 가 localStorage.removeItem 만, 탭 B 의 apiClient 인스턴스(이미 메모리에 accessToken 보유)는 모르고 계속 작동 — 'storage' 이벤트로 use-auth.ts L7-9 가 isLoggedIn snapshot 만 갱신, 그러나 apiClient.accessToken 멤버는 stale. 다음 fetch 까지 'logged in' 처럼 동작 → 401 받음 → silent refresh 시도 → revoked refresh → clearTokens(이번엔 동기화). 일관성 깨짐 + 보안 윈도우 존재."
  proposal: "(1) api-client.ts 에 syncFromStorage() 추가 — window addEventListener('storage', e => { if (e.key === 'accessToken' && e.newValue === null) { this.accessToken = null; this.refreshToken = null } }). (2) constructor 에서 자동 등록. (3) BroadcastChannel('auth') 도입 시(imp-0166 와 결합) 'logout' 메시지 받으면 즉시 clearTokens + router.push('/'). (4) Service Worker 가 있다면(grep) postMessage 동기화. (5) tab 간 race 제거 — 한 번 로그아웃하면 모든 탭이 0.5초 내 같은 화면(/) 으로."
  effort: small
  impact: medium
  evidence: "코드 web/lib/api-client.ts L23-46 ApiClient 클래스 — storage event listener 0건(grep 'addEventListener\\|storage' api-client.ts → 0). use-auth.ts L6-9 는 useSyncExternalStore 로 isLoggedIn 만 sync, apiClient 멤버 accessToken 은 직접 참조 안 함. logout 시 다른 탭 sync 0건."

- id: imp-0168
  found_at_iter: 11
  area: auth
  type: feature
  target: account_recovery_notification_channel
  problem: "코드 backend/db/migrations/001_initial.sql L27-36 users 테이블에 email/phone 컬럼 0건 — Google OAuth 로 받은 email(oauth/google.go L11-14 GoogleTokenInfo.Email) 은 인증 시 존재하지만 사용자 row 에 저장 안 함(handlers_auth.go L92 CreateUserWithProfile signature 가 nickname 만 받음). 결과: '계정 정지 이메일 통지', '약관 변경 안내', '비활성 계정 곧 삭제' 등 어떤 out-of-band 통신 채널도 부재. 운영자가 사용자에게 직접 닿을 수단 0건."
  proposal: "(1) 마이그레이션 — users 에 email TEXT(nullable, UNIQUE 아님 — 멀티 OAuth 같은 메일 가능), email_verified BOOLEAN DEFAULT FALSE, email_consented_marketing BOOLEAN DEFAULT FALSE 추가. (2) handleLogin L86-97 — info.Email 을 users.email 에 저장(Google OAuth 는 email_verified=true). (3) /profile/notifications-prefs 페이지 — 마케팅 수신 동의 토글, 거래 알림(필수) on, 운영 공지 토글. (4) backend SMTP 설정(env SMTP_HOST 등) + 단순 sendEmail() helper(SES/SendGrid/Mailgun 중 무료 등급). (5) 첫 사용 — 약관 변경 통지 1회 발송 테스트, 정지 알림, 30일 휴면 경고."
  effort: large
  impact: medium
  evidence: "코드 backend/db/migrations/001_initial.sql L27-36 users 컬럼 — id, login_provider, login_provider_user_key, account_status, role, last_login_at, created_at — email/phone 0건. handlers_auth.go L86-97 isNew 분기 info.Email 을 변수에서 dropping(저장 0건). SMTP 라이브러리 import 0건(go.mod grep 'mail' → 0). AccountSuspended 통지 채널 0건."

- id: imp-0169
  found_at_iter: 11
  area: auth
  type: ux
  target: error_message_disambiguation
  problem: "코드 web/app/login/page.tsx L48-53 catch 블록은 backend 에러를 setError(apiErr?.error?.message ?? '로그인에 실패했습니다') 로 일괄 표시 — 'Google 인증에 실패했습니다'(handlers_auth.go L57) 와 '계정이 정지되었습니다' (suspended)(handlers_auth.go L177) 모두 같은 빨간 텍스트로만. 사용자가 '내 잘못인지(토큰 만료/취소) vs 계정 문제(정지/탈퇴) vs 서버 문제(500)' 구분 불가 — 'Google 로 다시 시도' 인지 '고객센터 문의' 인지 다음 행동을 모름."
  proposal: "(1) login/page.tsx L48-53 — apiErr.error.code 별 분기 UI: code='UNAUTHORIZED' → '인증이 거절되었습니다 — 다시 시도하세요' + Google 버튼 강조 / code='ACCOUNT_SUSPENDED' → '계정이 정지되었습니다 — 사유: <message>' + '고객센터 문의(giranjt@gmail.com)' link / code='ACCOUNT_WITHDRAWN' → '탈퇴한 계정입니다 — 새 Google 계정으로 시도하세요' / code='INTERNAL_ERROR' → '서버 오류 — 잠시 후 다시 시도하세요'(retry 버튼 30초 cooldown). (2) 에러 경계 색상 구분 — auth(노랑) vs account(빨강) vs server(회색). (3) 5번 연속 같은 code 면 '문제가 계속되면 giranjt@gmail.com 으로 문의하세요'."
  effort: small
  impact: medium
  evidence: "코드 web/app/login/page.tsx L48-53 catch — apiErr.error.message 만 사용, error.code 분기 0건. backend 가 보내는 code 6종 이상(UNAUTHORIZED/ACCOUNT_SUSPENDED/ACCOUNT_WITHDRAWN/VALIDATION_ERROR/INTERNAL_ERROR/FORBIDDEN). UI 는 단일 빨간 p (text-[#e74c3c])."

- id: imp-0170
  found_at_iter: 11
  area: auth
  type: a11y_mobile
  target: google_oauth_iframe_focus_management
  problem: "Playwright snapshot 의 Google 로그인 iframe(ref=e44) 은 cross-origin → 키보드 Tab 으로 진입은 가능하나 focus ring 이 시각적으로 약하고 뒤로(Shift+Tab) 나가는 경로가 명확하지 않음. /login 본문에는 Google 버튼 외에 fallback 로그인 수단 0건이라 iframe 안에서 막히면 키보드만 쓰는 사용자(스크린리더, 손 부상)가 갇힌다. 또한 iframe focus 시 외부 페이지의 ESC 가 'login 취소' 의도로 동작 안 함."
  proposal: "(1) login/page.tsx L96 ref={googleBtnRef} 컨테이너에 onKeyDown — ESC 누르면 router.push('/'), Tab from outside 진입 시 visible focus ring(focus:ring-2 ring-gold). (2) iframe 위에 sr-only 라벨 — '<span class=sr-only>Google 계정으로 로그인. iframe 안의 버튼을 사용하거나 ESC 로 취소</span>'. (3) iframe 아래 '키보드 사용자: 둘러보기로 돌아가려면 Shift+Tab' helper text(sr-only 또는 detail). (4) fallback CTA(imp-0157 의 다른 provider 도입 시) 키보드 흐름 검증."
  effort: trivial
  impact: low
  evidence: "Playwright snapshot ref=e44 iframe(Google Sign-In) — outer page 에서 ESC handler 0건(login/page.tsx grep 'onKeyDown\\|ESC\\|Escape' → 0). iframe 외 키보드 진입 가능 요소 4종(헤더 로그인, body skip-link, 본문 둘러보기, 하단 nav 4개) 만."

- id: imp-0171
  found_at_iter: 12
  area: home
  type: a11y_mobile
  target: bottom_nav_safe_area_inset
  problem: "하단 고정 네비게이션(nav[aria-label='하단 메뉴'])의 padding-bottom이 0px이라 iPhone X+ 노치/홈 인디케이터 영역과 탭 항목이 겹친다. position:fixed; bottom:0; height:55.5px이며 safe-area-inset-bottom 를 반영하지 않는다."
  proposal: "하단 네비에 padding-bottom: env(safe-area-inset-bottom)를 추가하고 동시에 viewport meta에 viewport-fit=cover를 더한다. 본문 그리드의 마지막 카드가 가려지지 않도록 main에 padding-bottom: calc(56px + env(safe-area-inset-bottom)) 도 함께 부여."
  effort: trivial
  impact: medium
  evidence: "Playwright getComputedStyle: nav[aria-label='하단 메뉴'] { position:fixed, bottom:0, padding-bottom:0px, height:55.5px, bottomFromViewport:0 }. iOS 홈 인디케이터(약 34px)가 탭 영역과 충돌."

- id: imp-0172
  found_at_iter: 12
  area: home
  type: performance
  target: cdn_preconnect_missing
  problem: "매물 카드 아이템 아이콘은 /static/icons/{id}.png 로 노출되며(원본 assets.playnccdn.com), production 페이지의 link[rel=preconnect/dns-prefetch] 헤더에는 fonts.googleapis/gstatic 만 등록되어 있다. 첫 카드 그리드 렌더 시 같은 origin 내라도 Next 이미지 변환 origin과 CDN 사이 추가 RTT 발생."
  proposal: "<head>에 <link rel='preconnect' href='https://assets.playnccdn.com' crossorigin>와 <link rel='dns-prefetch' href='//assets.playnccdn.com'>를 추가한다. Next.js metadata.other 또는 app/layout.tsx의 head 컴포넌트로 주입."
  effort: trivial
  impact: low
  evidence: "Playwright: link[rel=preconnect] = ['fonts.googleapis.com','fonts.gstatic.com']만, link[rel=dns-prefetch] = 0개. 카드 아이콘 src 패턴 https://giranjt.com/static/icons/42.png (CDN 원본은 assets.playnccdn.com)."

- id: imp-0173
  found_at_iter: 12
  area: home
  type: feature
  target: filter_url_state_persistence
  problem: "서버 필터 '데포로쥬' 클릭 후 location.href, location.search 가 그대로 'https://giranjt.com/' 로 유지되어 URL에 어떠한 쿼리 파라미터도 반영되지 않는다. 결과적으로 '데포로쥬 + 무기' 같은 조합 화면을 친구에게 공유할 수 없고, 새로고침/뒤로가기 시 필터가 사라지며 SEO 인덱싱도 단일 URL로 묶인다."
  proposal: "필터 상태를 useSearchParams + replace 로 URL 쿼리에 직렬화한다(?server=depro&category=weapon&min=10000). Next.js App Router의 router.replace를 사용해 history flood 없이 동기화. 서버 컴포넌트에서 searchParams를 받아 SSR 응답에 반영하면 SEO/공유 양쪽 해결."
  effort: medium
  impact: high
  evidence: "Playwright: 데포로쥬 칩 click 후 location.href='https://giranjt.com/' (변화 없음), location.search='', history.length 변화 없음, localStorage/sessionStorage 키 0개. 필터는 클라이언트 상태로만 보존."

- id: imp-0174
  found_at_iter: 12
  area: home
  type: content
  target: meta_description_length
  problem: "meta[name=description].content 가 '리니지 클래식 아이템 거래 중개 플랫폼' (21자)로 Google 권장 50-160자 미달이다. SERP 스니펫이 잘려서 표시되거나 본문에서 자동 발췌되어 키워드 매칭률이 떨어진다."
  proposal: "Next.js metadata.description 을 130-150자로 확장: '리니지 클래식 서버별 무기·방어구·재화·계정 거래를 안전하고 무료로 중개합니다. 데포로쥬·켄라우헬·질리언 서버 매물을 실시간으로 둘러보고 채팅으로 거래하세요.' (검색 키워드 '리니지 클래식 거래', '서버별', '아이템 거래' 포함)."
  effort: trivial
  impact: medium
  evidence: "Playwright: meta[name=description].content.length = 21, 내용 = '리니지 클래식 아이템 거래 중개 플랫폼'. Google 권장 length 50-160."

- id: imp-0175
  found_at_iter: 12
  area: home
  type: content
  target: twitter_image_localhost_url
  problem: "meta[name='twitter:image'].content 가 'http://localhost:3000/images/og-image.png'로 노출되어 Twitter/X 카드 미리보기가 깨진다. og:image와 동일한 패턴이지만 dedup상 다른 target."
  proposal: "Next.js metadata.openGraph.images 와 metadata.twitter.images 를 NEXT_PUBLIC_SITE_URL 기반 절대 URL로 통일한다. 빌드 환경별 검증을 위해 lint 단계에 'localhost:3000' 문자열이 build 결과에 들어있지 않은지 검사."
  effort: trivial
  impact: high
  evidence: "Playwright: meta[name='twitter:image'].content = 'http://localhost:3000/images/og-image.png' (production https://giranjt.com 에서). meta[name='twitter:card'] = 'summary_large_image' 이므로 이미지가 깨지면 카드 노출 자체가 실패."

- id: imp-0176
  found_at_iter: 12
  area: home
  type: content
  target: relative_time_semantic_markup
  problem: "매물 카드의 '1개월 전' 등록 시각이 plain text로 출력되며 <time datetime> 마크업이 없다. 스크린리더는 '일개월전'으로 한 단어처럼 읽고, 정확한 등록일자(YYYY-MM-DD HH:mm)는 호버 툴팁이나 상세 페이지 진입 후에야 알 수 있다."
  proposal: "<time datetime='2026-04-02T13:24:00Z' title='2026년 4월 2일'> 1개월 전 </time> 형태로 마크업한다. 검색엔진의 freshness 시그널, 스크린리더 정확성, hover 툴팁(절대 시각) 세 가지가 동시에 개선됨."
  effort: trivial
  impact: low
  evidence: "Playwright: querySelectorAll('time').length = 0, 카드 본문 텍스트에 '1개월 전' 4건 모두 일반 span 으로 추정(time 태그 부재)."

- id: imp-0177
  found_at_iter: 12
  area: home
  type: a11y_mobile
  target: filter_chip_keyboard_roving_tabindex
  problem: "서버 필터 [aria-label='서버 필터'][role='group'] 의 28개 button 모두 tabindex=null(기본 0)이라 Tab 한 번에 28개 stop을 통과해야 다음 섹션(카테고리)으로 갈 수 있다. role='group' 까지는 있지만 toolbar/radiogroup 패턴 없이 키보드 사용자에게 단조로운 탭 지옥."
  proposal: "role='radiogroup' (단일 선택 의미) 또는 role='toolbar' 로 변경하고 roving tabindex 패턴 적용 — 활성 칩만 tabindex=0, 나머지 -1, ArrowLeft/ArrowRight로 이동. WAI-ARIA Authoring Practices toolbar/radiogroup 예시 따라 구현."
  effort: small
  impact: medium
  evidence: "Playwright: [aria-label='서버 필터'] role='group', 안의 button 5개 sample 모두 tabindex=null. 28+11=39개 칩 모두 Tab stop이라 키보드 사용자가 매물 그리드 도달까지 39+ Tab 입력 필요."

- id: imp-0178
  found_at_iter: 12
  area: home
  type: a11y_mobile
  target: filter_button_aria_pressed_vs_radio
  problem: "서버 필터에서 '전체' 클릭 시 다른 서버는 자동 해제되는 단일 선택(라디오) 의미인데, button[aria-pressed='true|false'] 토글 의미를 사용한다. 스크린리더는 '버튼, 눌림' 으로 읽어 '체크됨' 보다 의미 전달이 약하고 라디오 그룹의 'X / Y 선택됨' 같은 위치 컨텍스트가 없다."
  proposal: "(1) role='radio' + aria-checked + 부모 role='radiogroup' + aria-label='서버 선택' 으로 변경. (2) 또는 단일 select(custom listbox)로 패턴 자체를 통일 — '서버 ▾ 데포로쥬' 드롭다운. 카테고리 칩은 다중 선택일 가능성이 높으므로 aria-pressed 유지 가능."
  effort: small
  impact: medium
  evidence: "Playwright: 서버 필터 button 28개 모두 aria-pressed 사용 ('전체'=true, 나머지=false). 카테고리 11개도 동일 패턴이지만, 클릭 동작 분석상 서버는 단일선택(라디오), 카테고리도 단일선택으로 보여 '하나만 활성' 의미면 라디오가 더 적합."

- id: imp-0179
  found_at_iter: 12
  area: home
  type: a11y_mobile
  target: sort_select_native_height
  problem: "정렬 드롭다운 select 가 height 28px / width 92px 로 노출되며 appearance:auto(브라우저 기본 chevron 사용). 모바일 375px 뷰포트에서 28px 높이는 WCAG 2.5.5 권장 44x44, 최소 32x32 모두 미달이며 한 손 엄지 탭 시 빗나가기 쉽다."
  proposal: "정렬 select 의 min-height 를 44px 로 키우고 padding 을 늘린다. 디자인 통일을 위해 custom dropdown(Radix Select 등) 으로 교체하면 chevron/포커스 링/스타일 일관성도 함께 개선."
  effort: trivial
  impact: medium
  evidence: "Playwright getBoundingClientRect: select height=28, width=92, computed appearance='auto'. 동일 페이지 카테고리 칩 31px, 서버 칩 30px과 함께 탭 타깃 부족 라인업."

- id: imp-0180
  found_at_iter: 12
  area: home
  type: feature
  target: theme_color_scheme_toggle
  problem: "documentElement.style.colorScheme=='', getComputedStyle().colorScheme=='normal' 로 다크 테마가 강제되며 사용자/시스템 라이트 모드 사용자는 어두운 배경(#08080C)을 선택할 수 없다. 야외/주간 사용 시 가독성 저하."
  proposal: "Tailwind 의 darkMode='class' + system + 사용자 토글(헤더 우측)을 도입한다. CSS 변수로 색상 토큰을 정의해 라이트/다크 두 팔레트 모두 지원. localStorage 키 'theme' 에 'system|light|dark' 저장."
  effort: medium
  impact: medium
  evidence: "Playwright: getComputedStyle(html).colorScheme='normal', html.style.colorScheme='', [aria-label*='테마|라이트|다크']=0개, [data-testid*='theme']=0. 모든 사용자에게 강제 다크."

- id: imp-0181
  found_at_iter: 12
  area: home
  type: feature
  target: search_input_no_form_wrapper
  problem: "헤더와 hero 두 곳의 input[type=search] 모두 form 태그로 감싸이지 않아(i.form === false) Enter 입력 시 'submit' 이벤트가 발생하지 않는다. JS hydration 실패 시 검색이 완전히 동작하지 않으며, GET ?q= URL 공유도 불가."
  proposal: "input 을 <form action='/search' method='get'> 으로 감싸고 input.name='q' 를 부여한다. 클라이언트는 onSubmit 으로 router.push 처리, JS 실패 시에도 native submit 으로 /search?q=... 페이지 진입. 동시에 imp-0010 의 enterkeyhint='search' 와 결합."
  effort: small
  impact: medium
  evidence: "Playwright: input[type=search] 2개 모두 i.form===false, name='', form 미부착. /search 페이지 라우트 존재 여부 미확인이지만 어떤 경우든 progressive enhancement 부재."

- id: imp-0182
  found_at_iter: 12
  area: home
  type: content
  target: category_account_legal_clarity
  problem: "카테고리 필터에 '계정' 항목이 노출되며 [aria-label='카테고리 필터'] 안에 button text='계정' 이 존재한다. MMORPG 운영사 약관(NCSOFT 리니지 ToS)은 일반적으로 계정 양도/판매를 금지하며, 마켓플레이스가 이를 적극 노출하면 운영사 분쟁/계정 회수 리스크가 거래 양 당사자에게 발생한다."
  proposal: "(1) '계정' 항목을 제거하고 대체 카테고리(예: '커뮤니티 거래')만 남긴다. 또는 (2) '계정' 칩 클릭 시 약관 안내 모달 — '운영사 약관상 계정 거래는 회수/제재 위험이 있습니다' 노출 후 본인 책임 동의 시에만 진행. 동시에 매물 등록 폼도 동일 가드."
  effort: small
  impact: high
  evidence: "Playwright: [aria-label='카테고리 필터'] 안 button 11개 라벨 = ['전체','무기','방어구','장신구','소모품','재료','재화','주문서','펫/소환수','기타','계정']."

- id: imp-0183
  found_at_iter: 12
  area: home
  type: content
  target: filter_label_all_collision
  problem: "서버 필터와 카테고리 필터 모두 '전체' 라는 동일 라벨의 칩을 보유한다. 스크린리더가 '전체 버튼 눌림 / 전체 버튼' 식으로 두 번 읽어 어느 그룹의 '전체' 인지 사용자가 헷갈린다. 시각 사용자도 처음에는 두 그룹의 차이를 인지하기 어렵다."
  proposal: "두 '전체' 칩에 aria-label 로 컨텍스트를 명시 — 서버 그룹은 aria-label='모든 서버', 카테고리는 aria-label='모든 카테고리'. 시각 라벨은 '전체' 유지하되 SR 만 읽는 라벨을 다르게."
  effort: trivial
  impact: low
  evidence: "Playwright: 서버 그룹 '전체' button 1개(aria-pressed=true), 카테고리 그룹도 '전체' 1개(aria-pressed=true). 동일 가시 텍스트, group aria-label 만 다름('서버 필터' vs '카테고리 필터')."

- id: imp-0184
  found_at_iter: 12
  area: home
  type: performance
  target: pwa_service_worker_offline
  problem: "<link rel='manifest'> 는 존재(manifest=true) 하지만 navigator.serviceWorker.controller 가 null 이라 PWA 설치는 가능해도 오프라인/저사양 네트워크 시 페이지가 깨진다. 매물 카드 메타데이터 캐시도 없다."
  proposal: "Workbox 또는 next-pwa 로 service worker 도입 — (1) precache: HTML shell + 핵심 JS/CSS, (2) runtime cache: /api/v1/listings 30s SWR, /static/icons 30일 cache-first, /_next/image 7일 cache-first. 오프라인 시 마지막 캐시된 리스트 노출 + '오프라인 — 마지막 갱신 N분 전' 배너."
  effort: medium
  impact: medium
  evidence: "Playwright: !!document.querySelector('link[rel=manifest]')===true, navigator.serviceWorker.controller===null. PWA 절반만 구현."

- id: imp-0185
  found_at_iter: 12
  area: home
  type: ux
  target: overscroll_pull_to_refresh
  problem: "body/html 모두 overscroll-behavior:auto 라 모바일에서 위로 당기면 브라우저의 새로고침 제스처가 그대로 트리거된다. 매물 그리드 위쪽에서 의도치 않게 페이지 전체가 새로고침되어 적용한 필터(URL에도 안 남음, imp-0173 참조)도 함께 사라지는 이중 손실."
  proposal: "(1) body { overscroll-behavior-y: contain } 로 새로고침 차단, (2) 그리드 상단에 명시적 'Pull to refresh' 인디케이터 도입 또는 '새로고침' 버튼 + 자동 SSE 갱신. 의도된 새로고침은 명확한 제스처/버튼으로만."
  effort: small
  impact: medium
  evidence: "Playwright getComputedStyle: html.overscrollBehavior='auto', body.overscrollBehavior='auto'. 모바일 Chrome/Safari 기본 pull-to-refresh 활성."

- id: imp-0186
  found_at_iter: 12
  area: home
  type: a11y_mobile
  target: results_live_region_atomic_phrasing
  problem: "결과 섹션에 [aria-live='polite'] '4개 매물' 텍스트가 존재하나 aria-atomic 미설정이고 phrasing 이 무미건조하다. 필터 변경 시 SR 사용자에게 '4개 매물' 만 다시 읽혀 '데포로쥬 서버에서 0개를 찾았습니다' 같은 컨텍스트가 빠진다."
  proposal: "(1) aria-atomic='true' 추가 — 부분 변경이 아닌 전체 메시지로 읽음. (2) live region 텍스트를 '검색 결과: {N}건 ({server} 서버, {category} 카테고리)' 같이 컨텍스트 포함. (3) 0건일 땐 '결과가 없습니다. 필터를 변경해 보세요.' 로 더 친절."
  effort: trivial
  impact: medium
  evidence: "Playwright: [aria-live='polite'] 1개, text='4개 매물', aria-atomic=null. 필터 적용 후에도 라이브리전 phrasing 변화 미관찰."

- id: imp-0187
  found_at_iter: 12
  area: home
  type: ux
  target: long_item_title_truncation
  problem: "매물 카드 제목 '도리깨 +10 판매합니다' / 'ㅇㄴㄹ2ㅈ +3' 등은 짧지만, 사용자가 '데포로쥬 서버 풀강 +10 도리깨 급매 판매합니다 (네고가능)' 같이 60자+ 제목을 입력하면 카드 폭(328px) 안에서 어떻게 잘리는지 시각적 처리가 검증되지 않았다. 현재 카드 마크업에 line-clamp 또는 text-overflow:ellipsis 가 보이지 않는다."
  proposal: "h3 카드 제목에 line-clamp-2 또는 line-clamp-1 + text-overflow:ellipsis 적용. 카드 높이를 일정하게 유지해 그리드 흐트러짐 방지. 모바일 1열, 태블릿 2열, 데스크톱 4열 모두에서 검증."
  effort: trivial
  impact: low
  evidence: "Playwright 카드 텍스트 분석: 제목 '도리깨 +10 판매합니다' 11자, '111233/1234 +1' 등 짧음. line-clamp 클래스명 카드 outerHTML(800자 chunk)에서 미관찰. 매우 긴 제목 입력 케이스 미테스트."

- id: imp-0188
  found_at_iter: 12
  area: home
  type: ux
  target: price_display_compact_format
  problem: "최고가 매물 '43,680,742원' 이 카드에 7자리 그대로 노출되어 가독성이 떨어진다. 한국 사용자에게는 '4,368만원' 또는 '약 4368만원' 같은 만 단위 축약이 즉각 인지되지만 현재는 풀 자리수만 출력."
  proposal: "Intl.NumberFormat('ko-KR', { notation:'compact', maximumFractionDigits:1 }) → '4368.1만' 또는 자체 한국식 만/억 단위 포매터로 변환. hover 툴팁/aria-label 에는 정확한 원 단위(43,680,742원) 보존. 1만원 미만은 그대로 표시."
  effort: trivial
  impact: medium
  evidence: "Playwright: 노출 가격 ['200,000원','43,680,742원','1원','10원']. Intl.compact 테스트 결과 '4368.1만'. 7자리 천 단위 콤마는 인지에 약 1.5x 시간 더 걸림(인지 부하 연구)."

- id: imp-0189
  found_at_iter: 13
  area: listings
  type: performance
  target: sitemap_xml_missing
  problem: "/sitemap.xml 이 404, robots.txt 에도 Sitemap: 디렉티브가 없다. 검색엔진(Google/Naver/Bing)이 매물 상세 URL(/listings/<UUID>) 을 발견할 경로가 카드 href crawling 한 단계뿐이라 인덱싱 커버리지가 매우 낮다. 매물이 100건 이상 쌓여도 첫 페이지 카드 4개만 SEO 노출, 나머지는 영구 미발견."
  proposal: "Next.js App Router의 app/sitemap.ts 동적 라우트로 /sitemap.xml 생성. (1) 정적: /, /create, /login. (2) 동적: 모든 status='available' 매물의 /listings/<id> 를 lastmod=lastActivityAt 으로. (3) /robots.txt 에 'Sitemap: https://giranjt.com/sitemap.xml' 추가. 매물 100건 미만이면 단일 sitemap, 50000건 초과 시 sitemap_index.xml 분할."
  effort: small
  impact: high
  evidence: "curl https://giranjt.com/sitemap.xml → HTTP 404 (Next.js prerender 404 페이지). curl https://giranjt.com/robots.txt → 27라인 content-signal 정책만, Sitemap 디렉티브 0건. 매물 상세 URL 4개 모두 클라이언트 클릭으로만 도달, 정적 sitemap 미존재."

- id: imp-0190
  found_at_iter: 13
  area: listings
  type: performance
  target: jsonld_itemlist_structured_data
  problem: "/ (마켓 메인)에 JSON-LD structured data가 전혀 없다. 검색엔진이 매물 그리드를 ItemList 로, 각 카드를 Product/Offer 로 인식하면 가격/가용성 리치 결과(rich snippet)로 노출 가능한데, 현재는 일반 검색 결과로만 표시되어 CTR 손실."
  proposal: "app/page.tsx 에 server-side fetched listings 로 <script type='application/ld+json'> 주입. ItemList { itemListElement: [{ '@type': 'Offer', name: title, price: priceAmount, priceCurrency: 'KRW', availability: status==='available' ? InStock : OutOfStock, url: '/listings/<id>', seller: { '@type': 'Person', name: nickname } }] }. /listings/<id> 페이지에는 Product 단일 객체."
  effort: medium
  impact: medium
  evidence: "curl https://giranjt.com/ | grep 'ld+json' → 0건. <script type='application/ld+json'> 검색 결과 없음. Google Search Central 가이드: ItemList 마크업 시 SERP CTR 평균 +20~40%."

- id: imp-0191
  found_at_iter: 13
  area: listings
  type: bug
  target: og_image_localhost_url
  problem: "OG/Twitter 메타에 og:image='http://localhost:3000/images/og-image.png' 가 그대로 배포되어 있다. 카카오톡/디스코드/트위터/슬랙 등 어떤 SNS 에 매물 링크를 공유해도 미리보기 이미지가 깨진다(localhost 는 외부에서 도달 불가). 매물 공유 = 마케팅 핵심 채널인데 통째로 망가진 상태."
  proposal: "next.config.ts 에 metadataBase = new URL(process.env.NEXT_PUBLIC_SITE_URL || 'https://giranjt.com') 설정. app/layout.tsx 의 openGraph.images / twitter.images 를 상대경로 '/images/og-image.png' 로 두면 metadataBase 가 자동으로 절대 URL 합성. 빌드 ENV NEXT_PUBLIC_SITE_URL=https://giranjt.com 보장."
  effort: trivial
  impact: high
  evidence: "curl https://giranjt.com/ | grep og:image → '<meta property=\"og:image\" content=\"http://localhost:3000/images/og-image.png\"/>'. 동일하게 twitter:image 도 localhost. 외부 도달 불가 URL → 모든 SNS 미리보기 깨짐."

- id: imp-0192
  found_at_iter: 13
  area: listings
  type: feature
  target: price_range_filter_ui
  problem: "API 는 ?priceMin=&priceMax= 쿼리를 받지만(요청 시 200 OK 반환) 메인 페이지의 ListingFilters 컴포넌트 어디에도 가격 범위 입력 UI 가 없다. 사용자는 '500만원 이하 무기' 같은 가장 자연스러운 필터 의도를 표현할 방법이 없다. 단일 매물의 가격 차이가 1원~4368만원으로 4자리수 차이가 나는 마켓에선 결정적 누락."
  proposal: "ListingFilters 에 가격 범위 슬라이더 또는 두 개 input[type='number'] 'X만원 ~ Y만원' 추가. 빠른 칩으로 '~50만', '50~500만', '500만~' 3-4개 프리셋. URL 동기화(imp-0027 의존). 모바일은 collapsed → '가격 ▾' 아코디언."
  effort: medium
  impact: high
  evidence: "Code web/components/listing/listing-filters.tsx: priceMin/priceMax/priceRange 키워드 0건. API curl ?priceMin=100000 → 200(필터 동작은 됨, 단 현재 데이터 4건 모두 통과). use-listings.ts 훅 시그니처에 priceMin/priceMax 파라미터 미존재."

- id: imp-0193
  found_at_iter: 13
  area: listings
  type: feature
  target: listing_type_filter_buy_sell
  problem: "ListingFilters 컴포넌트에 listingType(판매/구매) 필터 칩이 없다. 카드 좌측 border-l-4 색상으로 sell=gold, buy=blue 구분만 시각적으로 표현될 뿐, '구매 글만 보기' 같은 능동적 필터링이 불가능. 판매자는 자기 아이템에 매칭되는 '구매 요청'만 골라 보고 싶어함."
  proposal: "서버 칩과 같은 패턴으로 [전체 | 판매 | 구매] 3개 button[aria-pressed] 칩 그룹을 결과 헤더 좌측에 둔다. URL 쿼리 ?type=buy|sell. 색상 토큰: 판매=gold, 구매=blue 로 카드 border-l 와 일치."
  effort: small
  impact: medium
  evidence: "Code web/components/listing/listing-filters.tsx: listingType 칩 검색 0건. listing-card.tsx L10 만 'listingType === sell' 체크(border 색). use-listings.ts 훅은 listingType 파라미터를 받음(L9) — 이미 백엔드 지원 완료, UI 만 빠짐."

- id: imp-0194
  found_at_iter: 13
  area: listings
  type: feature
  target: realtime_new_listing_toast
  problem: "사용자가 마켓 페이지에 머무는 동안 새 매물이 등록돼도 화면이 갱신되지 않는다. EventSource/SSE 사용 흔적 0건. 거래는 '먼저 본 사람이 먼저 사는' 시간경쟁이 있는데 사용자는 새로고침을 반복하거나 알림을 영영 못 받음."
  proposal: "Backend 에 GET /api/v1/listings/stream (SSE) 추가. domain.ListingCreated 이벤트를 구독해 push. Web 은 useEffect 로 EventSource 연결, 새 매물 도착 시 화면 상단에 '새 매물 N건 ▾' sticky toast/banner 표시 → 클릭 시 InfiniteQuery refetch + scrollTo top. 30초 idle 시 자동 disconnect 로 NAS 부하 보호."
  effort: large
  impact: high
  evidence: "Code grep web/app/page.tsx 'EventSource|sse|realtime' → 0건. 백엔드 internal/event 디렉토리는 존재하나 listings.stream 핸들러 미존재. 채팅에서는 SSE 사용중(채팅 SSE 재연결 PR 최근 머지) — 인프라는 있으나 listings 에 미적용."

- id: imp-0195
  found_at_iter: 13
  area: listings
  type: feature
  target: saved_search_email_push_alert
  problem: "사용자가 ?server=zillian&category=weapon_mace&priceMax=500000 같은 조건을 자주 검색해도 '저장' 또는 '이 조건으로 새 매물이 나오면 알림' 기능이 없다. 매물이 드물게 등장하는 희귀 아이템(예: 풀강 무기)을 노리는 사용자는 매일 수동 새로고침해야 함."
  proposal: "Backend: saved_searches 테이블(user_id, query_json, alert_channel:none|push|email, created_at). 핸들러 POST /api/v1/saved-searches, GET (내 목록), DELETE. 새 매물 등록 이벤트 시 query_json 매칭하는 saved_searches 의 user 에게 NotificationCreate. Web: 필터 영역 우측에 '이 검색 저장 ★' 버튼 → 비로그인은 useAuthGuard."
  effort: large
  impact: medium
  evidence: "Code grep '/api/v1/saved-searches' → backend 핸들러 0건. curl https://giranjt.com/api/v1/listings/saved-searches → 404. 사용자 페르소나 분석: 풀강 +9 도리깨 매물은 평균 30일에 1회 등록 → 수동 모니터링 비현실적."

- id: imp-0196
  found_at_iter: 13
  area: listings
  type: feature
  target: price_drop_watch_notify
  problem: "사용자가 특정 매물을 보고 '가격이 떨어지면 사겠다'고 결정한 경우, 현재는 매번 매물 상세에 들어가 가격 변경 여부를 확인해야 한다. 찜(favorite) 은 현재 즐겨찾기일 뿐 가격 변경 트리거가 없다."
  proposal: "Backend: listing_watches 테이블(user_id, listing_id, target_price, original_price, created_at). Listing PATCH 핸들러에서 priceAmount 가 감소하면 watches 조회 → 매칭(target_price>=new_price)되는 user 에게 알림('관심 매물 가격이 N% 인하되었습니다'). Web 은 매물 상세 액션바에 '가격 알림 받기 (목표가 입력)' 버튼."
  effort: medium
  impact: medium
  evidence: "Backend grep 'priceAlert|priceDrop|listing_watch' → 0건. listing_favorites 테이블만 존재(웹사이트 상의 찜 토글). 가격 변경 이벤트(ListingPriceChanged) 도메인 이벤트 카탈로그(EVENT_CATALOG.md) 부재 추정 — 별도 검증 필요."

- id: imp-0197
  found_at_iter: 13
  area: listings
  type: feature
  target: listing_compare_side_by_side
  problem: "도리깨 +10 매물이 3개 등장했을 때 사용자는 가격/판매자평판/거래방식을 동시에 비교하고 싶다. 현재는 카드를 클릭→상세 진입→뒤로→다음 카드 클릭의 반복인데 뒤로갈 때마다 필터 상태 초기화(imp-0029)까지 겹쳐 비교가 사실상 불가능."
  proposal: "카드 우상단에 작은 체크박스 '비교 추가'. 최대 3개까지 sticky 하단 비교 트레이('비교 (2/3) ▶'). 클릭 시 /compare?ids=a,b 라우트로 가서 가격/서버/거래방식/판매자뱃지/등록일/조회·찜·채팅수를 표 형태로 비교. 비로그인도 sessionStorage 기반 동작."
  effort: medium
  impact: low
  evidence: "Code grep 'compareListings|ListingCompare' → 0건. /compare 라우트 미존재. 현재 카드 click 패턴: 상세→뒤로 시 React state 손실 = 비교 비효율."

- id: imp-0198
  found_at_iter: 13
  area: listings
  type: feature
  target: rss_atom_feed
  problem: "/api/v1/listings.rss, .atom 모두 404. 매물 알림을 외부 도구(IFTTT, Slack RSS 봇, Feedly)로 받고 싶은 파워유저 니즈를 충족 못 함. 모바일 푸시 알림 인프라가 없는 현재 상태(웹뿐) 에서 RSS 는 가장 저비용의 외부 알림 채널."
  proposal: "Backend handler GET /api/v1/listings.rss → application/rss+xml. <channel> title='기란JT 신규 매물', <item>은 최근 50건 ?serverId=&categoryId= 쿼리도 받아 필터링된 피드 가능. /api/v1/listings.atom 도 함께. robots.txt 와 layout.tsx <link rel='alternate' type='application/rss+xml' href='/api/v1/listings.rss' />."
  effort: small
  impact: low
  evidence: "curl /api/v1/listings.rss → 404, .atom → 404. 매물 등록 알림을 외부로 push할 채널 0개(이메일/푸시 미구현). RSS 는 추가 구독 인프라 0원, gem-stable feature."

- id: imp-0199
  found_at_iter: 13
  area: listings
  type: feature
  target: anti_spam_listing_throttle
  problem: "단일 사용자(da1e387a-fd67-4189-8475-b1a379e9a2a2)가 4개 매물 중 3개를 등록했고, 그중 2개가 명백한 자판 노이즈('123ㄱ1ㄱㅈㄷㄹ', '111233')다. 매물 등록 횟수 제한(throttle) 또는 신규 계정의 시간당/일당 등록 한도가 없어 spam/test 등록이 정상 매물을 밀어낸다."
  proposal: "Backend: trustBadge='newcomer' (가입<7d 또는 거래<3) 인 유저는 시간당 최대 3건, 일당 10건으로 제한. POST /api/v1/listings 응답 429 + Retry-After. 한 시간 내 같은 itemMasterId+priceAmount 중복 등록은 409 + 'similar listing already exists'. middleware 에 user-scoped rate limit Redis (또는 in-memory map for NAS 단일 노드)."
  effort: medium
  impact: high
  evidence: "API 응답: 4매물 중 75%가 한 신규유저(trustBadge='newcomer')의 자판 노이즈 매물. Backend grep 'rateLimit|throttle' → 0건(매물 생성 핸들러에 rate-limit 미적용). 노이즈 매물은 imp-0038 의 '데모트' 만으로는 발생 자체를 못 막음."

- id: imp-0200
  found_at_iter: 13
  area: listings
  type: feature
  target: fraud_signal_price_outlier
  problem: "'머꼬' 매물이 43,680,742원(4,368만원), '도리깨 +10'이 200,000원, '111233'이 10원 — 동일 카테고리/유사 아이템의 시장가 대비 극단치를 자동 감지하는 신호가 없다. 가격 1원/10원 매물은 명백한 테스트/사기 의심이지만 정상 매물처럼 노출."
  proposal: "Backend: 카테고리별 priceAmount 의 P5/P95/median 을 1시간마다 머터리얼라이즈드 뷰로 계산. Listing 응답 또는 별도 필드 priceOutlierFlag: 'too_low' | 'too_high' | 'normal'. Web 카드에 '⚠️ 시세 대비 매우 낮음' 작은 라벨 + admin 모더레이션 큐 자동 진입. 너무 거짓양성 많으면 P1/P99 로 완화."
  effort: large
  impact: medium
  evidence: "API: 카테고리=무기 매물 가격 [200000, 43680742, 1, 10] — std 가 평균보다 높은 long-tail. fraud/anomaly/outlier grep 백엔드 0건. 가격 1원·10원 매물은 100% 사기 또는 테스트일 확률 높지만 일반 사용자에 노출."

- id: imp-0201
  found_at_iter: 14
  area: listing_detail
  type: performance
  target: jsonld_product_offer_schema
  problem: "<script type='application/ld+json'> 가 페이지에 0건이다. Google 의 Product/Offer rich result 가이드에 따르면 매물 페이지는 name/image/description/offers(price, priceCurrency, availability, seller) 를 JSON-LD 로 노출해야 검색 결과에 가격·이미지·재고 상태가 함께 표시된다. 현재는 일반 텍스트 스니펫만 인덱싱되어 검색 CTR 가 낮다."
  proposal: "web/app/listings/[id]/page.tsx 를 server component 로 분리하고 generateMetadata 와 별개로 <Script id='ld-product' type='application/ld+json'> 에 Product schema 를 직렬화한다. 필드: name=title, image=[iconUrl 절대 URL], description=description, sku=listingId, offers={'@type':'Offer', price=priceAmount, priceCurrency='KRW', availability=status==='available'?'InStock':'OutOfStock', seller={'@type':'Person', name=author.nickname}}, brand={'@type':'Brand', name=serverName}. imp-0048(메타) 과 함께 적용."
  effort: small
  impact: high
  evidence: "Playwright: document.querySelectorAll('script[type=\"application/ld+json\"]').length === 0. API 응답에 productSchema 에 매핑 가능한 모든 필드(title, iconUrl, description, listingId, priceAmount, status, author.nickname, serverName) 존재."

- id: imp-0202
  found_at_iter: 14
  area: listing_detail
  type: performance
  target: jsonld_breadcrumb_schema
  problem: "BreadcrumbList JSON-LD 가 없어 검색 결과에 'giranjt.com > 마켓 > 질리언 > 둔기' 같은 계층 표시가 안 된다. 동시에 시각적 breadcrumb UI 도 없다(imp-0057). 둘이 동시 부재라 사용자/검색엔진 모두 매물의 위치 컨텍스트를 잃는다."
  proposal: "imp-0057 의 시각 breadcrumb 와 함께 BreadcrumbList JSON-LD(item: '/', '/listings?server=zillian', '/listings?server=zillian&category=weapon_mace', '/listings/<id>') 를 주입. position 1..4. server component 의 generateMetadata 와 같은 자리에서 함께 처리."
  effort: trivial
  impact: medium
  evidence: "Playwright: breadcrumbSchema=false (JSON-LD 0건), 시각 breadcrumb 도 imp-0057 에서 부재 확인. API 의 serverName/categoryName 이 인덱스 키로 사용 가능."

- id: imp-0203
  found_at_iter: 14
  area: listing_detail
  type: feature
  target: item_master_db_link
  problem: "아이템 아이콘(64x64)을 클릭/호버해도 '도리깨' 아이템의 마스터 정보(공격력, 무게, 직업 제한, 강화 표 등)를 볼 수 없다. 리니지 클래식 사용자는 +10 도리깨가 어느 직업에 어떤 효과인지 알아야 가격 정당성을 판단하는데, 매물 상세에서 도감 정보가 단절되어 외부 위키로 이탈한다."
  proposal: "(1) 아이콘 우측에 작은 'i' 인포 칩을 두고 클릭 시 popover 로 itemMaster 정보(공격력, 무게, 장착 직업, 강화 단계별 효과) 노출. (2) 아이콘 자체는 <a href='/items/{categoryId}/{itemMasterId}'> 로 감싸서 새 도감 페이지(또는 동일 카테고리 매물 목록)로 라우팅. (3) backend item_master 시드에 baseStats JSON 컬럼 추가, GET /api/v1/items/:id 신설."
  effort: large
  impact: medium
  evidence: "Playwright: 아이콘 img alt='', title 없음, 부모 a/button 없음(클릭 불가). itemMasterDb 검색 0건('레벨/장착 직업/아이템 설명' 어떤 텍스트도 페이지에 없음). 사용자가 아이템 정보를 얻으려면 외부 위키 검색 필요."

- id: imp-0204
  found_at_iter: 14
  area: listing_detail
  type: ux
  target: stale_listing_warning
  problem: "createdAt='2026-03-21T01:58:50Z' 와 lastActivityAt 가 정확히 같다 — 게시 후 1개월간 단 한 번도 가격 수정/응답이 없었다. chatCount=0 favoriteCount=0 도 같이 실패 시그널이지만 페이지는 '판매중' 그대로 표시한다. 거래 의사가 살아있는 매물인지 실질적으로 죽은 매물인지 사용자가 판별 어렵다."
  proposal: "lastActivityAt - createdAt 차이가 작고(즉 활동 0) (now - lastActivityAt) > 14일이면 '⚠️ 게시 후 활동 없음 (1개월)' dim 톤 헬퍼 라인을 stats 옆에 노출. 30일 이상이면 '판매자가 응답할 가능성이 낮습니다' 톤으로 강화. 60일 이상이면 backend cron 으로 status='archived' 자동 전환 + 응답 없음 매물 검색에서 후순위. imp-0049(time semantic) 과 같은 영역."
  effort: small
  impact: medium
  evidence: "API: createdAt=lastActivityAt='2026-03-21T01:58:50Z' (정확히 같음 → 활동 0), now=2026-05-03 → 약 43일 경과, viewCount=22 chatCount=0 favoriteCount=0. 페이지에 '활동 없음/장기 미응답' 시그널 0건."

- id: imp-0205
  found_at_iter: 14
  area: listing_detail
  type: content
  target: korean_price_readable_format
  problem: "가격 표시 '200,000원' 은 읽기는 가능하지만 한국어 화자에게 '20만원' 이 더 직관적이다. 큰 숫자(예: 43,680,742원)일수록 인지 부담이 커지는데, '4,368만원' 또는 '약 4,300만원' 보조 라벨이 없다. 또한 스크린리더가 '이백만원' 으로 읽지 못하고 '이백천원' 으로 잘못 읽을 가능성."
  proposal: "(1) 가격 옆 또는 아래 작은 톤으로 '약 20만원' 보조 라벨 추가. 1억 이상이면 '약 1.2억', 1,000만 이상이면 '약 4,300만원' 처럼 만/억 단위. (2) <span aria-label='이십만 원'> 처럼 한국어 음독 aria-label 부여(toKoreanNumeral 헬퍼). (3) 통일된 KRW 포맷 유틸 web/lib/format/krw.ts 신설."
  effort: small
  impact: low
  evidence: "Playwright: 가격 텍스트 '200,000원', 보조 라벨 0건. 같은 카테고리 다른 매물 '43,680,742원' 도 동일 패턴. 가격 element 의 aria-label 없음(스크린리더는 '이만 원' 또는 '이십만 원' 둘 중 무작위)."

- id: imp-0206
  found_at_iter: 14
  area: listing_detail
  type: a11y_mobile
  target: viewport_disable_user_zoom
  problem: "<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1'> 으로 핀치줌이 차단되어 있다. 이는 WCAG 1.4.4 (Resize Text) 와 모바일 a11y 가이드 위반이다. 시력 약한 사용자가 매물 사진/설명/가격을 확대해서 볼 수 없고, 강화 +10 같은 작은 텍스트도 확대 불가."
  proposal: "viewport 메타에서 maximum-scale=1 제거(또는 maximum-scale=5, user-scalable=yes). web/app/layout.tsx 의 viewport 설정 수정. iOS 에서 input focus 시 자동 zoom 우려는 input font-size>=16px 또는 user-scalable 유지로 해결 가능. PWA 일관성을 위해 standalone 모드에서만 maximum-scale 유지하도록 conditional."
  effort: trivial
  impact: medium
  evidence: "Playwright: <meta name='viewport'> content='width=device-width, initial-scale=1, maximum-scale=1', blocksZoom=true. WCAG 1.4.4 / SC 1.4.10 미달 — 시력 약한 사용자 핀치줌 불가."

- id: imp-0207
  found_at_iter: 14
  area: listing_detail
  type: feature
  target: seller_other_listings_preview
  problem: "imp-0046 은 '같은 server+category' 비슷한 매물 추천이고, 이건 다른 각도 — '판매자 da1e387a 의 다른 매물' 미니 카드 3-4개를 판매자 정보 카드 아래에 노출. 사용자가 한 판매자에게 신뢰가 가면 묶음으로 거래하고 싶어하는 패턴(특히 같은 서버/같은 직업의 풀세트 거래)을 지원."
  proposal: "GET /api/v1/listings?authorId=<userId>&excludeId=<listingId>&limit=4&status=available API 추가. 판매자 카드(L99 부근) 아래 SellerOtherListingsPreview 컴포넌트로 가로 스크롤 미니 카드. 0건이면 표시하지 않음. imp-0046 과 컴포넌트 재사용 가능."
  effort: medium
  impact: medium
  evidence: "API author.userId='da1e387a-fd67-4189-8475-b1a379e9a2a2' 존재(이 유저는 listings 페이지에서 4매물 중 3매물 보유 — 다른 매물 표시할 자료 충분). 페이지 검색 '판매자의 다른|이 판매자의' 0건. imp-0046 의 similar 와 별개 축."

- id: imp-0208
  found_at_iter: 14
  area: listing_detail
  type: feature
  target: notify_on_price_change_subscription
  problem: "찜 등록(imp-0042) 이외에 '가격 변경/상태 변경/응답 시 알림' 같은 세분 구독이 없다. 비싸서 못 사는 매물을 '20만원 이하로 떨어지면 알려줘' 같이 가격 트리거로 설정하면 buy-side 컨버전이 살아난다."
  proposal: "찜 버튼 옆 ⋯ 메뉴에 '가격 알림 설정' 추가. 모달에서 'X원 이하 도달 시 알림' 트리거 설정. 백엔드: listing_price_alerts 테이블(user_id, listing_id, threshold_price, channel='in_app|email|kakao'). priceAmount 변경 시 cron 으로 임계값 비교 → 알림 fanout. status='reserved' 변경 시 즉시 '예약됨' 알림."
  effort: large
  impact: medium
  evidence: "Playwright: '가격 변동 알림|상태 알림' 텍스트 0건, statusAlert=false. priceType='fixed' 매물도 협상 가능성 있어 가격 변경은 흔함. 비슷한 기능: 당근/번개장터 모두 가격 알림 제공."

- id: imp-0209
  found_at_iter: 14
  area: listing_detail
  type: ux
  target: counter_offer_inline_form
  problem: "priceType='fixed' 든 'negotiable' 든 buyer 가 직접 '제시 가격'을 보내는 폼이 매물 상세에 없다. 사용자는 채팅을 시작해서 '15만원 어떠세요?' 라고 자유 텍스트로 적어야 하고, 판매자도 협상 진행을 추적하기 어렵다."
  proposal: "negotiable 매물에 '제시 가격 입력' 인라인 폼 + 'X원 으로 제시' 버튼 추가. 클릭 시 채팅방 자동 생성 + 첫 메시지로 'BuyerNickname 님이 150,000원을 제시했습니다' 시스템 메시지 + offerAmount 가 chat_room.proposed_price 에 저장. 판매자 채팅창에 '수락/거절/역제안' 인라인 액션 버튼. priceType='fixed' 도 'fixed_with_negotiation' 옵션 추가 시 활성화."
  effort: large
  impact: high
  evidence: "Playwright: 'offer/제시/입찰/역제안' 텍스트 0건, offerForm=false counterOffer=false. priceType='fixed' 200,000 매물에 협상 진입점 없음. 채팅 기반 자유 협상은 추적/통계/완료율 측정 불가."

- id: imp-0210
  found_at_iter: 14
  area: listing_detail
  type: content
  target: newcomer_safe_trade_warning
  problem: "판매자 trustBadge='newcomer' (가입 신규 + 거래 0회) 인데도 '안전 거래 가이드 / 사기 주의 / 첫 거래 유의사항' 같은 보호 카피가 매물 상세 어디에도 없다. 신규 사용자에게 첫 거래는 사기 위험이 가장 높은 시점인데, 플랫폼이 아무 가이드도 제공하지 않는다."
  proposal: "trustBadge='newcomer' 또는 completedTradeCount<3 일 때 판매자 카드 아래 InfoBox '⚠️ 첫 거래 유의사항' 펼침: (1) 게임 내 거래 인증 후 송금, (2) 외부 카톡 송금 강요 시 신고, (3) 채팅 내 메시지 캡처 보존, (4) 거래 후 리뷰 작성. 카피는 5줄 이내 한국어 친근체. 같은 박스에 '신고하기' 직접 링크."
  effort: small
  impact: high
  evidence: "API: author.trustBadge='newcomer', completedTradeCount=0. Playwright: '안전 거래|사기 주의|유의사항' 텍스트 0건 safeTrade=false. 신규 거래자에게 안내 부재."

- id: imp-0211
  found_at_iter: 14
  area: listing_detail
  type: feature
  target: recently_viewed_strip
  problem: "방문자가 마켓 → 매물 A → 마켓 → 매물 B → 매물 A 패턴으로 비교 탐색하는데, 매물 상세에서 '내가 최근에 본 매물' 가로 스크롤 strip 이 없다. 비교 추적이 불가능해 사용자는 백 버튼 + 메모리에 의존한다."
  proposal: "localStorage 의 recentlyViewedListings 배열(최대 10건, listingId+title+priceAmount+iconUrl+viewedAt) 에 페이지 진입 시 unshift. 매물 상세 하단(혹은 imp-0046 비슷한 매물 위)에 '최근 본 매물' 섹션 렌더. 클릭 시 해당 매물로 이동. 익명/로그인 둘 다 동작 가능, 로그인 시 server 동기화 옵션."
  effort: small
  impact: medium
  evidence: "Playwright: '최근 본|recently viewed' 0건 recentlyViewed=false. localStorage.recentlyViewedListings undefined. 비교 탐색이 흔한 거래 행위인데 도구 부재."

- id: imp-0212
  found_at_iter: 14
  area: listing_detail
  type: feature
  target: kakaotalk_share_intent
  problem: "공유(imp-0056) 가 navigator.share + clipboard 폴백만 있고, 한국 사용자가 가장 많이 쓰는 카카오톡 직접 공유 SDK 가 없다. 카톡으로 매물 링크를 보낼 때 미리보기(og:image, og:title)도 깨져있어(localhost:3000) 시각적 어필 0."
  proposal: "(1) 카카오 JavaScript SDK 의 Kakao.Share.sendDefault({objectType:'feed', content:{title, description, imageUrl, link}}) 호출 버튼 추가('카카오톡으로 공유'). (2) imp-0048(메타 동적 주입) 으로 og:image 를 listing.images[0] 또는 iconUrl 절대 URL 로 fix. (3) 공유 메뉴를 Sheet/Popover 로 묶어 [카톡 / 텔레그램 / URL 복사 / QR 코드] 4 옵션 통합. KAKAO_APP_KEY 환경변수 필요."
  effort: medium
  impact: medium
  evidence: "Playwright: '카카오|kakao' 0건, telegramShare false, qrCode false. og:image='http://localhost:3000/images/og-image.png'(잘못됨). intentLinks=[] (kakaolink:// 또는 다른 앱 인텐트 0건). 카톡 공유는 한국 시장 거래 매물 99% 의 1순위 공유 채널."

- id: imp-0213
  found_at_iter: 15
  area: listing_create
  type: feature
  target: quantity_field_input
  problem: "createListingRequest 백엔드 스펙은 Quantity int min=1 을 요구하지만 web/app/create/page.tsx 는 quantity:1 로 하드코딩(L92)해 사용자가 입력할 UI 가 없다. 결과로 룬/포션/주문서/물약 같은 stack-able 아이템(예: '아데나 포션 ×100', '대장군 룬 ×3') 을 묶어 등록하지 못해, 동일 아이템 100개를 100건의 분리 매물로 올리거나 description 본문에 '×100' 이라고 비정형 표기하는 우회로가 발생한다."
  proposal: "(1) selectedItem.isStackable=true(또는 categoryId in ['potion','rune','scroll']) 일 때만 '수량' number input 노출(default 1, min 1, max 999). (2) 가격 라벨도 '개당 가격' 으로 변경 + total='priceAmount × quantity' 라이브 미리보기. (3) 매물 목록 카드에는 '×수량' chip 표시(imp-0044 와 함께)."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx L92 quantity:1 하드코딩, 폼에 수량 input 0개. backend handlers_listing.go L26 Quantity binding=required,min=1 — DB 스키마는 임의 수량 지원. 룬/포션 카테고리는 본질적으로 stack 거래인데 단일 매물=단일 단위로 강제됨."

- id: imp-0214
  found_at_iter: 15
  area: listing_create
  type: feature
  target: duplicate_listing_action
  problem: "반복 거래자(예: 같은 활성화 룬을 매주 1개씩 파는 회원)가 동일 아이템·서버·조건의 매물을 새로 등록할 때마다 8개 필드를 처음부터 다시 입력해야 한다. 백엔드 GET /listings/:id 응답에는 모든 필드가 있고 / 매물 상세에 '이 매물 복제' 진입점이 0건이다."
  proposal: "(1) 자기 매물 상세 페이지(또는 /profile 의 내 매물 리스트) 에 '복제하여 새로 등록' 메뉴 노출. 클릭 시 router.push(`/create?duplicate=${listingId}`). (2) /create 페이지가 ?duplicate 쿼리를 감지하면 GET /listings/:id 결과로 form/images/selectedItem/enhancementLevel 모두 prefill (단, status/createdAt 제외). (3) 한 줄 알림 '거래 완료된 동일 매물을 다시 등록합니다 — 가격을 확인해주세요' 노출. status=sold 매물도 복제 가능."
  effort: medium
  impact: medium
  evidence: "코드 web/app/create/page.tsx L31-41 form initial state 정적, useSearchParams 사용 0건. listing-detail 우상단 메뉴(?) 본인 매물 진입점에 '복제' 옵션 0건. 거래 데이터(my-listings) 패턴상 동일 카테고리 반복 등록자가 다수일 것."

- id: imp-0215
  found_at_iter: 15
  area: listing_create
  type: feature
  target: listing_template_save
  problem: "imp-0065 의 단일 draft autosave 와 별개로, 다양한 카테고리 거래자(예: 무기·방어구·장신구를 매주 번갈아 파는 회원)가 카테고리별 '기본 양식'을 영구 보관할 수단이 없다. 매번 '거래 가능 시간', '디스코드 ID', '가격 협상 정책' 같은 boilerplate 를 description 에 다시 적어야 한다."
  proposal: "(1) /profile 또는 /create 상단 우측에 '내 템플릿' 드롭다운 추가. 사용자는 등록 화면에서 '템플릿으로 저장' 버튼을 눌러 현재 form snapshot 을 이름붙여 localStorage('listing_templates' 배열, 최대 5개) 에 저장. (2) 새 등록 시 '템플릿 적용' 으로 즉시 prefill (선택 후에도 수정 가능). (3) 백엔드 user_settings.template_blob (JSON) 스키마 추가 시 디바이스 간 동기화. 단순 v1: localStorage."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx 에 template/preset 관련 state/UI 0건. localStorage key 'listing_templates' 미존재. profile 페이지에도 '내 양식' 진입점 0건. 반복 boilerplate 입력은 마찰점."

- id: imp-0216
  found_at_iter: 15
  area: listing_create
  type: feature
  target: bulk_listing_paste
  problem: "거래자가 한꺼번에 여러 매물(예: '활성화 룬 +3 50만, 활성화 룬 +5 80만, 진명검 +6 200만')을 등록하려면 매번 폼을 처음부터 다시 채워야 한다. 폼 1회 작성에 30초~2분 걸리므로 N건 등록 시 N×시간 — 도매 거래자 이탈 요인."
  proposal: "(1) /create 페이지 상단에 '일괄 등록 모드' 토글. 활성 시 textarea 노출 — 한 줄당 1매물 형식 '아이템명 +강화 가격 / 서버 / 거래방식'. 예: '진명검 +6 2000000 / 카인 / either'. (2) Parse 후 미리보기 테이블(N행) 생성, 사용자가 row 별 verify/edit 가능. (3) '한 번에 등록' 버튼 → POST /listings/bulk (N개를 transaction 으로). 실패한 row 만 빨간색 표시. (4) 같은 아이템·강화 데이터 자동 매칭(item_master fuzzy)."
  effort: large
  impact: medium
  evidence: "코드 web/app/create/page.tsx 단일 매물 form 만 존재, '일괄/bulk/multi' 키워드 0건. backend POST /listings 도 단건만. 도매 거래자(예: 룬 100개 보유) 케이스 미지원."

- id: imp-0217
  found_at_iter: 15
  area: listing_create
  type: feature
  target: schedule_publish_at
  problem: "사용자가 '주말 저녁 8시'(거래자 활동 피크) 에 매물을 노출하고 싶어도 등록은 즉시 publish 만 가능하다. status='available' 로 즉시 들어가고 visibility='public' 이라 가시성 컨트롤 0. 결국 사용자는 알람 맞춰서 직접 그 시간에 등록 클릭해야 한다."
  proposal: "(1) 거래 섹션에 '게시 시간' radio: [지금 / 예약]. '예약' 선택 시 datetime-local input(현재 시각+10분 ~ +7일 사이). (2) 백엔드 listings.scheduled_publish_at TIMESTAMPTZ 컬럼 추가, status='scheduled' visibility='hidden' 으로 INSERT, 별도 cron(또는 Listings.ListListings 쿼리에 NOW() >= scheduled_publish_at 조건 추가) 으로 자동 publish. (3) 사용자 매물 목록에 '게시 예정 — 5월 4일 20:00' 뱃지 + 즉시 게시/취소 버튼."
  effort: large
  impact: medium
  evidence: "코드 web/app/create/page.tsx 게시시간 input 0개, backend handlers_listing.go L38 INSERT 시 status='available' visibility='public' 하드코딩. listings 테이블에 scheduled_at/publish_at 컬럼 미존재. UX 적으로 거래 활동 피크 시간 활용 불가."

- id: imp-0218
  found_at_iter: 15
  area: listing_create
  type: feature
  target: image_enhancement_ocr_autofill
  problem: "사용자가 게임 인벤토리 스크린샷을 첫 이미지로 올리면 강화 수치('+6'), 옵션('마력 +1, 명중 +2'), 아이템명을 사람이 다시 타자해야 한다. 이미지에 이미 모든 정보가 있는데 폼은 그것을 활용하지 않는다."
  proposal: "(1) 이미지 업로드 직후 백엔드 POST /images/:id/extract 에서 OCR(Tesseract 한글 또는 Cloud Vision) 실행, '+숫자' 패턴/옵션 텍스트/아이템명 추출. (2) 결과를 응답으로 받아 enhancementLevel/optionsText/itemName 자동 prefill (사용자에게 '이미지 분석 결과를 적용했어요. 확인해주세요' 토스트). (3) MVP 는 Tesseract.js 클라이언트 OCR 로 시작(외부 API 비용 0). 정확도 낮으면 점진적으로 개선. (4) 사용자가 '직접 입력' 모드로 끄는 토글."
  effort: large
  impact: medium
  evidence: "코드 web/components/forms/image-upload.tsx upload() L27-74 단순 multipart POST 만, OCR/EXIF 추출 호출 0건. backend handlers_upload.go(추정) 에 분석 endpoint 미존재. Tesseract.js, Cloud Vision, OpenAI Vision 등 어떤 통합도 없음. 라이프스타일 마켓(당근/번개)도 점진 도입한 핵심 차별화."

- id: imp-0219
  found_at_iter: 15
  area: listing_create
  type: ux
  target: price_amount_won_unit_helper
  problem: "priceAmount 가 type='number' 단일 input 으로 '500000' 같은 long-digit 입력을 강제한다. 한국어 거래자는 자연스럽게 '50만', '200만', '1억', '1.5억' 으로 표현하고, 화면 표시도 '500,000원' 천단위 콤마가 자연스러운데, 폼은 (1) 콤마 미적용, (2) '만/억' 단축 미지원, (3) 라이브 한글 변환('500,000 = 50만원') 미표시."
  proposal: "(1) priceAmount input 을 controlled 컴포넌트로 변경: 표시값=Number(value).toLocaleString('ko-KR'), state=원시 정수. (2) input 옆 라이브 hint '50만원' 표시(value/10000>=10? '50만' : null). (3) chip 그룹 [+1만, +10만, +100만, +1000만, ÷2, ×2] 로 빠른 조정. (4) input mode='numeric' inputMode='numeric' 으로 모바일 숫자 키보드 활성. (5) 50,000 미만 또는 1,000,000,000 초과 시 inline warning."
  effort: small
  impact: high
  evidence: "코드 web/app/create/page.tsx L199-207: <input type='number' value={form.priceAmount} onChange={update} disabled={priceType=='offer'} />. toLocaleString 호출 0건, '만/억' 변환 0건, chip 그룹 0건, inputMode 미지정. 거래 가격은 보통 만~억 단위라 콤마 부재 시 자릿수 오타(0 하나 누락) 위험 큼."

- id: imp-0220
  found_at_iter: 15
  area: listing_create
  type: ux
  target: buy_listing_role_relabel
  problem: "listingType='buy'(구매 매물) 모드를 토글하면 제목 자동생성만 '...구매합니다' 로 바뀌고, 나머지 라벨은 모두 판매 기준 그대로다. (1) '가격(원)' 라벨이 '내가 지불할 최대가' 로 바뀌어야 의미 명확, (2) '거래 방식' 도 '받을 방식' 가깝고, (3) 가격 유형 'offer' 의미가 '제안받음→내가 제안받는다(=판매자 입장)' 라 구매 모드에서는 모순. (4) 이미지도 '내가 사고 싶은 아이템 예시' 로 의미 변동 — 아이콘 1장으로 충분, 5장 강제 필요 없음."
  proposal: "(1) form.listingType==='buy' 일 때 라벨 dynamic: 가격 '최대 지불가', 거래 '수령 방식', 이미지 라벨 '구매하려는 아이템 이미지(선택)'. (2) priceType options 도 buy 모드에서는 [고정가, 협상가능] 만, 'offer' 제거 또는 '제안 받습니다(= 가격 비공개로 채팅)' 로 의미 명시. (3) 이미지 maxImages 를 buy 모드에서는 2장으로 자동 축소. (4) 'BUY' 컬러 액센트(예: 채도 낮은 녹색)로 시각적 구분."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx 라벨 정적, listingType 분기는 L53(typeStr) 1곳만. priceType options L189-193 = [fixed, negotiable, offer] 양 모드 동일. ImageUpload maxImages=5 고정 L250."

- id: imp-0221
  found_at_iter: 15
  area: listing_create
  type: ux
  target: description_overwrite_on_item_change
  problem: "handleItemSelect L70-76 가 selectedItem 변경 시 description 을 항상 `[아이템 옵션]\\n${optionText}\\n\\n` 으로 덮어쓴다(item.optionText 있을 때). 사용자가 이미 description 을 5분간 작성한 후 '아이템 변경' 으로 다른 아이템을 고르면 작성한 거래 조건/디스코드 ID/협상 정책이 모두 사라진다. 토스트나 확인 다이얼로그 0건."
  proposal: "(1) handleItemSelect 에서 description 이 빈 문자열 OR 'auto-generated' 시그널일 때만 덮어쓰기. 사용자가 직접 입력한 텍스트가 있으면 보존(또는 prepend '[아이템 옵션]\\n…\\n\\n' + 기존 텍스트). (2) descriptionAutoFilled state 를 두고 imp-0069 '제목 자동생성' 패턴과 일관 처리. (3) item 변경 시 토스트 '이전 설명을 유지했어요. 옵션을 추가하려면 다시 적어주세요' 안내."
  effort: trivial
  impact: medium
  evidence: "코드 web/app/create/page.tsx L70-76: setForm({...f, description: item.optionText ? `[아이템 옵션]\\n…\\n\\n` : f.description}). 사용자 입력 보존 분기 0건. descriptionAutoFilled state 미존재. confirm()/dialog 0건."

- id: imp-0222
  found_at_iter: 15
  area: listing_create
  type: feature
  target: deep_link_query_prefill
  problem: "외부 커뮤니티(디스코드, 카페)에서 누군가가 '진명검 +6 카인서버 50만 정도에 사세요' 같은 거래 추천 링크를 공유할 때, /create?item=진명검&enchant=6&server=cain&price=500000 같은 deep-link prefill 이 가능하면 conversion 이 큼. 현재 폼은 useSearchParams 호출 0건이라 어떤 쿼리도 prefill 되지 않는다."
  proposal: "(1) /create 마운트 시 useSearchParams 로 [item/enchant/server/price/listingType/categoryId] 추출. (2) ItemAutocomplete 가 item= 쿼리로 자동 fuzzy-search → 첫 결과 selectedItem 으로 자동 선택. (3) 검색 결과 0건이거나 모호하면 prefill 만 하고 사용자가 확정 클릭. (4) prefill 적용 시 상단 dismissable 배너 '외부 링크에서 가져온 정보로 채웠어요'. (5) /listings/[id] 의 '비슷한 매물 등록' CTA 와 결합."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx useSearchParams import 0건, 모든 form state 가 useState 빈값으로 초기화. URL 쿼리 파싱 분기 0건. /listings/[id] 페이지에도 'create with this template' 링크 0건."

- id: imp-0223
  found_at_iter: 15
  area: listing_create
  type: performance
  target: image_upload_per_item_progress
  problem: "ImageUpload.upload (L27-74) 은 validFiles 를 for-of 로 순차 await 하는 동안 단순 'uploading=true' boolean 만 표시한다. 5장 ×10MB 를 모바일 LTE 에서 올리면 30~90초 소요인데 사용자는 '몇 번째 / 몇 장 진행 중' / '실패한 파일이 어느 것' 알 수 없다. 1번째 성공 후에도 다음 4장은 같은 'spinner + 업로드 중...' 만 표시."
  proposal: "(1) 각 파일을 [{file, status:'queued'|'uploading'|'done'|'failed', progress:0-100, errorMsg?}] 배열로 관리. (2) Promise.all + axios onUploadProgress 또는 fetch + ReadableStream 으로 per-file progress 추적. (3) 업로드 그리드의 각 썸네일 자리에 (a) 진행 중: 원형 progress + percent, (b) 실패: '×' + retry 버튼, (c) 완료: 정상 썸네일. (4) 동시 병렬 3개로 throughput 향상."
  effort: medium
  impact: medium
  evidence: "코드 web/components/forms/image-upload.tsx L51-72: setUploading(true) → for-of await uploadImage → setUploading(false). per-file progress 0건, axios/fetch progress 이벤트 0건. UI 단일 'uploading ? 업로드 중...' 분기. 5장 직렬 업로드 시 사용자에게 정보 0."

- id: imp-0224
  found_at_iter: 15
  area: listing_create
  type: content
  target: submit_error_specific_messages
  problem: "handleSubmit L106-108 catch 블록이 서버 응답 무관 항상 '등록에 실패했습니다' 토스트만 표시. backend handlers_listing.go 는 4가지 분기 에러 코드 발생: VALIDATION_ERROR(400), FORBIDDEN(403, 이미지 소유권), 가격 미입력 VALIDATION_ERROR(400, 별도 메시지), INTERNAL_ERROR(500). 사용자는 '뭘 고쳐야 하는지' 모른 채 같은 폼을 다시 submit."
  proposal: "(1) catch (err) 에서 err.response.data.error.code 로 분기: VALIDATION_ERROR → 백엔드 message 그대로 토스트 + 첫 실패 필드로 스크롤(Gin binding 메시지에서 field 추출). FORBIDDEN → '이미지 권한이 없습니다. 다시 업로드해주세요' + images=[] 리셋. INTERNAL_ERROR → '서버 오류 — 1분 후 재시도' + Sentry 로그. (2) error.code 매핑 테이블을 lib/errors.ts 에 통합 — 모든 mutation 이 공통 사용. (3) 토스트 카드에 retry 버튼."
  effort: small
  impact: medium
  evidence: "코드 web/app/create/page.tsx L103-108: try { await mutateAsync } catch { addToast('error', '등록에 실패했습니다') }. err 객체 사용 0건. backend handlers_listing.go 응답 형식 {error:{code, message}} 일관 사용 중인데 클라이언트가 무시."

- id: imp-0225
  found_at_iter: 15
  area: listing_create
  type: ux
  target: form_progress_stepper
  problem: "폼은 8개 섹션(거래 유형 / 아이템 / 서버 / 가격 / 상세 정보(제목·설명) / 이미지 / 거래방식 / 등록 버튼)이 단일 세로 스크롤로 노출된다. 모바일 375px 에서 화면 높이 667px 기준 스크롤 길이가 ~3 화면 분량이라 사용자가 '내가 어디까지 했는지', '얼마나 남았는지' 인지하기 어렵다. 진행률 bar / step indicator 0건."
  proposal: "(1) 상단 sticky 진행률 bar — 필수 필드 7개(itemName, serverId, title, description, priceAmount(offer 제외), images optional, tradeMethod) 중 채워진 비율. (2) 또는 multi-step wizard — [1/4 아이템 → 2/4 가격 → 3/4 상세/이미지 → 4/4 거래방식+제출] 좌우 스와이프 또는 '다음' 버튼. (3) 모바일은 stepper, 데스크탑(lg+)은 단일 스크롤 + 진행률 bar 만. (4) 미입력 단계는 stepper 에 빨간 점 표시."
  effort: medium
  impact: medium
  evidence: "코드 web/app/create/page.tsx form structure: 단일 <form className=space-y-4> 내 div 8개. progressBar/stepper/multi-step 컴포넌트 0건. sticky 헤더 0건. 모바일 폼 길이 측정: 입력 모두 채우면 ~1800px(스크롤 3~4회)."

- id: imp-0226
  found_at_iter: 15
  area: listing_create
  type: ux
  target: post_create_edit_redirect
  problem: "handleSubmit L104-105 가 등록 성공 후 router.push('/') 로 홈으로 리다이렉트한다. 사용자는 '내 매물이 정말 등록됐는지', '오타가 없는지' 즉시 확인 불가능하고, 잘못된 정보(예: 가격 0 하나 더 붙음) 도 detail 페이지를 새로 찾아 들어가서야 발견한다. UpdateListing 백엔드는 있지만 UI 진입점이 거의 없다."
  proposal: "(1) 성공 시 router.push(`/listings/${listingId}?just_created=true`). (2) 매물 상세에서 ?just_created 일 때 상단에 success 카드 — '등록 완료! 이대로 게시할까요?' + [확인 / 수정] 2 버튼. '수정' 클릭 시 인라인 편집 모드로 (제목/설명/가격/거래방식 만 수정 가능). (3) 5분 이내에는 무료 수정, 이후엔 status_history 기록과 함께 수정 가능. (4) 토스트 'X님의 매물이 등록되었습니다' + 본인이면 '편집' 링크."
  effort: small
  impact: high
  evidence: "코드 web/app/create/page.tsx L104-105: await mutateAsync(data); router.push('/'). 성공 시 listingId 사용 0건(응답 무시). backend POST /listings 응답 L104-108 listingId 반환됨. 매물 상세에서 본인 편집 진입점 visible=false(probably hidden)."

- id: imp-0227
  found_at_iter: 15
  area: listing_create
  type: content
  target: char_count_and_max_length
  problem: "백엔드는 title min=2 max=100, description min=10 max=2000 강제. 프런트는 minLength 만 적용하고 maxLength 미적용 — 사용자가 description 에 2001자 작성 후 submit 하면 backend VALIDATION_ERROR 가 토스트로 떠 모든 작업이 무위. 또한 라이브 글자수 카운터도 부재라 '얼마나 남았는지' 모름. imp-0073 placeholder 가이드와 별개 issue."
  proposal: "(1) <input id='title' maxLength={100} aria-describedby='title-counter'> + 우측 작은 'N/100' 표시(>=80% 일 때 amber). (2) <textarea id='description' maxLength={2000} aria-describedby='desc-counter'> + 'N/2000자' 표시. (3) 카운터는 한글 = 1자(서버도 byte 가 아니라 char 기준이라 가정 — 검증 필요). (4) 95% 도달 시 'border-amber', 100% 도달 시 'border-danger' + submit disabled 까진 안 가도 OK(maxLength 가 막음)."
  effort: trivial
  impact: medium
  evidence: "코드 web/app/create/page.tsx L216-228 title <input minLength={2}>, L236-245 textarea <textarea minLength={10}> — maxLength 속성 0건. 글자수 표시 span/div 0건. backend handlers_listing.go L22-23 binding 'min=2,max=100' 'min=10,max=2000' 강제, 위반 시 400 에러."

- id: imp-0228
  found_at_iter: 16
  area: chat
  type: feature
  target: chat_message_search_within_room
  problem: "수십~수백 개의 메시지가 누적된 채팅방에서 '아까 그 좌표', '얼마였지?', '몇 시 약속이었지?' 같은 정보를 다시 찾으려면 사용자가 위로 스크롤하면서 fetchNextPage 가 5~10회 트리거되어야 한다. 백엔드는 cursor pagination 만 지원하고 메시지 본문 검색 API 가 없다. 거래 협상의 핵심 정보가 묻히면 분쟁/오해 발생률이 높아진다."
  proposal: "(1) 채팅방 헤더에 돋보기 아이콘 → 클릭 시 inline search bar 'XX방의 메시지 검색' (active chat 만). (2) 백엔드 GET /api/v1/chats/:id/messages?q=... 신규 엔드포인트, ILIKE '%q%' + 최근 30일 제한. (3) 검색 결과는 highlight 후 클릭 시 해당 메시지로 jump-scroll(앵커 ID = messageId). (4) 모바일에서는 헤더 우측 ⋯ → 검색 메뉴."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_chat.go L108-162 handleListMessages 는 cursor 만 파라미터로 받음. q 파라미터 핸들링 0건. web/app/chats/[id]/page.tsx 헤더 영역 L129-142 에 검색 input/button 0개. fetchNextPage 호출 trigger 는 scrollTop<50 한 가지만(L79)."

- id: imp-0229
  found_at_iter: 16
  area: chat
  type: feature
  target: chat_input_draft_persistence
  problem: "ChatInput 의 text state 가 useState 로만 관리되어 페이지 새로고침/실수로 다른 채팅방 클릭/탭 닫기 시 작성 중이던 긴 메시지가 즉시 휘발된다. 모바일에서 한 손으로 길게 쓴 메시지가 알림 클릭으로 다른 채팅방 진입 시 사라지면 사용자 frustration 이 매우 크다. 거래 협상 중 정성스럽게 작성한 협상 문구도 동일 위험."
  proposal: "(1) ChatInput 의 onChange 시 localStorage `chat_draft_<chatRoomId>` 에 debounce 500ms 로 저장. (2) 채팅방 진입 시 localStorage 에서 draft 복원, 옆에 작은 '임시 저장 복원됨' 토스트 1.5초. (3) submit 성공 시 draft 삭제. (4) 길이 0 일 때도 삭제. (5) 7일 이상 된 draft 는 mount 시 자동 청소(quota 보호)."
  effort: trivial
  impact: medium
  evidence: "코드 web/components/chat/chat-input.tsx L15 const [text, setText] = useState('') — 외부 persist 0건. localStorage/sessionStorage 사용 0건. activeChatId 변경 시 text 가 그대로(잘못된 chat 에 prefill 가능성). useEffect cleanup 0건."

- id: imp-0230
  found_at_iter: 16
  area: chat
  type: feature
  target: browser_native_notification_permission
  problem: "현재 새 메시지가 도착해도 OS-level 알림이 발송되지 않는다. 사용자가 다른 탭/창에 있으면 favicon/title 변화도 없어 도착 사실 자체를 모르고, 거래 응답이 30분~1시간 늦어진다(거래 conversion 에 직결). Web Notifications API 권한 요청도 없음."
  proposal: "(1) 사용자가 처음 채팅방에 진입할 때 일회성 prompt '새 메시지를 데스크탑 알림으로 받을까요?' [허용 / 다음에]. (2) Notification.requestPermission() 결과를 localStorage 에 저장. (3) SSE 의 new_message 수신 시 document.hidden 이거나 활성 chat 이 다르면 new Notification(`${nickname}: ${preview}`, { icon, tag: chatRoomId, body }) 발송. (4) 클릭 시 window.focus() + activeChatId 설정. (5) 안전을 위해 같은 tag 의 새 알림은 기존 알림을 대체."
  effort: small
  impact: high
  evidence: "코드 grep Notification.requestPermission 0건, new Notification 0건. web/lib/hooks/use-sse.ts onmessage 핸들러 (확인 필요) 에는 invalidate/setQueryData 만 있고 OS notification 호출 0건. web/components/chat/chat-panel.tsx L143-155 도 reconnecting/disconnected 만 alert."

- id: imp-0231
  found_at_iter: 16
  area: chat
  type: feature
  target: unread_count_in_document_title_and_favicon
  problem: "데스크탑 사용자가 다른 탭에서 작업하다 채팅 탭으로 돌아오기 전까지는 새 메시지 도착 사실을 알 수 없다. document.title 은 '기란JT' 고정, favicon 도 정적이라 시각적 시그널 0건. SSE 가 메시지를 받아도 useChats invalidate 만 일어남."
  proposal: "(1) 전역 unreadTotal = chats.reduce((s, c) => s + c.unreadCount, 0) 계산. (2) useEffect 로 document.title 을 unreadTotal>0 일 때 `(${unreadTotal}) 기란JT — 새 메시지` 로 변경, =0 이면 원복. (3) favicon 도 canvas 로 redraw — 우상단 빨간 dot+숫자 (badge canvas 패턴, 라이브러리 'favico.js' 또는 직접 구현 ~30 lines). (4) cleanup: 페이지 unmount 시 원래 title/favicon 복원."
  effort: small
  impact: medium
  evidence: "코드 web/app/layout.tsx 또는 globals 에서 document.title 동적 변경 0건. favicon redraw/canvas 0건. useChats unreadCount 합산 컴포넌트 0건. 모바일은 OS 가 처리하지만 데스크탑은 fallback 필요."

- id: imp-0232
  found_at_iter: 16
  area: chat
  type: feature
  target: link_url_auto_detection_and_safety
  problem: "ChatMessage 의 bodyText 는 plain text 로 렌더되어 사용자가 'https://disco...' 또는 'kakao.com' 같은 URL 을 보내도 클릭 가능한 링크가 아니다. 거래 외부 채널 유도(피싱 위험)도 감지/경고 없이 표시. 또한 URL preview 도 없어 신뢰성 판단 어려움."
  proposal: "(1) bodyText 를 정규식 /(https?:\\/\\/[^\\s]+)/g 로 split → 일반 텍스트는 span, URL 은 <a target='_blank' rel='noopener noreferrer nofollow'>. (2) 외부 채널(disco.gg, kakao.com, telegram.org, line.me, naver.com) 도메인은 노란 배경 + '⚠️ 외부 메신저 링크 — 거래 사기 주의' badge. (3) 신뢰 도메인(giranjt.com, nccsoft.co.kr 등)은 일반 링크 색. (4) 향후 OG preview 로 확장 가능."
  effort: small
  impact: medium
  evidence: "코드 web/components/chat/chat-message.tsx L94 {message.bodyText} — 단순 문자열 렌더, parsing 0건. <a> 태그 0개. 외부 메신저 유도는 거래 사기 패턴(상대 사이트 안내 → 환불 분쟁) 의 전형이지만 차단/경고 0건."

- id: imp-0233
  found_at_iter: 16
  area: chat
  type: feature
  target: chat_export_archive_for_dispute
  problem: "거래 분쟁 발생 시 사용자가 채팅 기록을 외부에 증거로 제출하려면 현재는 메시지를 일일이 스크롤+스크린샷해야 한다. 100+ 메시지 채팅을 포착하려면 수십 장의 스크린샷이 필요하고, 시간순서/발신자 식별이 어렵다. 운영팀에게 신고할 때도 동일한 burden."
  proposal: "(1) 채팅방 헤더 ⋯ 메뉴에 '대화 내보내기' 추가. (2) 클릭 시 GET /api/v1/chats/:id/export 호출 → 백엔드가 모든 메시지를 시간순으로 모아 .txt(plain) 또는 .json 파일로 응답. 헤더에 채팅방 ID/매물/참여자/내보낸 시각 포함. (3) 내보낸 파일에 hash(SHA256) 푸터로 위변조 방지 시그널. (4) 신고 모달의 '증거 첨부' 자동 export 옵션 — 신고 시 자동 export → backend admin 에 첨부."
  effort: medium
  impact: medium
  evidence: "코드 backend/cmd/server/handlers_chat.go: export 엔드포인트 0건(GET /chats 만, 또 message list 만). web 앱: download/export/blob.createObjectURL 호출 0건. 신고 모달 web/components/forms/report-modal.tsx 도 첨부 옵션 0개."

- id: imp-0234
  found_at_iter: 16
  area: chat
  type: ux
  target: timestamp_localized_relative_per_message
  problem: "ChatMessage 의 formatMessageTime L4-26 은 '오후 3:42' / 'M월 D일 오후 3:42' 형식만 사용해 절대 시간을 보여준다. 거래 직후 30초 간격 메시지도 같은 시간으로 보여 빠른 응답 vs 늦은 응답 구분이 어렵고, 채팅 목록의 '몇 분 전' 과 채팅 본문의 '오후 3:42' 표기 일관성도 떨어진다."
  proposal: "(1) ChatMessage hover/long-press 시 native title 속성 + 풀 datetime 표시 (이미 sentAt 보유). (2) 마지막 메시지가 5분 이내면 '방금 전', 5~60분이면 'XX분 전', 60분 ~ 24시간이면 'X시간 전', 그 이상이면 현재 absolute. (3) 사용자 설정 절대/상대 토글(setting page). (4) chat-list-item 과 동일한 formatTimeAgo 함수 재사용으로 일관성 유지."
  effort: trivial
  impact: low
  evidence: "코드 web/components/chat/chat-message.tsx formatMessageTime L4-26: 상대 시간 0건, 항상 absolute. utils.ts formatTimeAgo 는 chat-list-item 에서만 사용. title 속성 0건. 가독성 + accessibility(스크린리더) 양쪽 손해."

- id: imp-0235
  found_at_iter: 16
  area: chat
  type: a11y_mobile
  target: chat_no_swipe_back_gesture_handler
  problem: "iOS Safari 의 edge-swipe-back gesture 가 모바일 채팅 상세에서 활성화되어 있어, 사용자가 메시지 영역에서 좌우로 스크롤(이미지 캐러셀이 추가될 경우) 하거나 스크롤 끝에서 손가락을 떼면 의도와 무관하게 /chats 로 빠져나가버리는 사고가 빈번하다. 또한 PWA 모드에서 history.back() 이 의도와 다르게 동작한다."
  proposal: "(1) 채팅 상세 컨테이너에 touch-action: pan-y 또는 onTouchStart 핸들러로 e.touches[0].clientX < 20 인 swipe-from-left 일 때 stopPropagation. (2) 헤더에 명시적 ← back 아이콘 추가(현재 없음 — Header 컴포넌트의 햄버거만). (3) 헤더의 매물 정보 카드 (ListingInfoCard) 좌측에 back arrow 추가. (4) PWA 시 manifest 'standalone' 에서 visible-before-back 효과 보장."
  effort: small
  impact: medium
  evidence: "코드 web/app/chats/[id]/page.tsx L124-189: back button 0개. Header 는 가로 메뉴만. touch-action CSS/onTouchStart 0건. iOS Safari 기본 edge swipe back 활성화 상태(meta apple-mobile-web-app-* 미명시)."

- id: imp-0236
  found_at_iter: 16
  area: chat
  type: ux
  target: scroll_to_bottom_when_user_scrolled_up
  problem: "사용자가 이전 메시지를 보려 위로 스크롤한 상태에서 새 메시지가 도착하면, useEffect L106-108 가 항상 bottomRef.scrollIntoView 를 호출해 스크롤 위치가 강제로 끌려간다. 사용자는 읽고 있던 위치를 잃고 다시 스크롤해야 함 — 협상 중 필요한 정보를 다시 찾을 때 마찰. 또한 새 메시지 도착 인디케이터도 없어 '아래에 새 메시지가 있다'는 신호 0건."
  proposal: "(1) scrollRef 의 scrollTop 과 scrollHeight-clientHeight 차이 < 100px 일 때만 자동 스크롤. (2) 차이가 크면 우하단 floating button '↓ 새 메시지 N' 표시 → 클릭 시 bottomRef 로 점프. (3) 사용자 메시지 전송 시는 항상 강제 점프(intent 명확). (4) 신규 카운터는 마지막 본 messageId 와의 차이로 계산."
  effort: small
  impact: medium
  evidence: "코드 web/app/chats/[id]/page.tsx L106-108: useEffect 가 messages.length 변경 시 무조건 scrollIntoView. floating '새 메시지' 버튼 0개. scrollRef.current.scrollTop 검사 로직 0건. 사용자 의도 보호 가드 부재."

- id: imp-0237
  found_at_iter: 16
  area: chat
  type: feature
  target: copy_message_text_action
  problem: "사용자가 상대방이 보낸 메시지(예: 약속 시간, 좌표, 입금 계좌)를 복사하려면 텍스트를 길게 눌러 OS 텍스트 선택 핸들을 띄워야 하는데, 모바일에서는 메시지 버블 모서리에서 정확히 잡기 어렵고 데스크탑에서도 ChatMessage 가 selectable 여부 보장 0건. 거래에서는 '계좌번호 복사' 가 빈번해 마찰 큰 동작이다."
  proposal: "(1) ChatMessage hover 시(데스크탑) 또는 long-press 시(모바일) 메시지 옆에 '⧉ 복사' 아이콘 노출. (2) 클릭 시 navigator.clipboard.writeText(bodyText) + toast '복사되었어요'. (3) URL 만 포함된 메시지는 '링크 복사', 일반 텍스트는 '메시지 복사' 라벨 분기. (4) ⌘C 가 활성화될 수 있도록 message bubble 에 user-select: text 명시(현재 default)."
  effort: trivial
  impact: low
  evidence: "코드 web/components/chat/chat-message.tsx: hover/long-press 핸들러 0건, copy 아이콘/버튼 0개. navigator.clipboard 호출 web/lib/ 전반 0건(profile/listing 도). 메시지 bubble 의 user-select CSS 미명시(Tailwind default 는 inherit 라 가능하지만 명시적 indicator 없음)."

- id: imp-0238
  found_at_iter: 16
  area: chat
  type: feature
  target: scam_keyword_filter_warning
  problem: "거래 사기 관련 키워드(예: '디스코드로 와서', '카카오톡으로 송금', '먼저 입금', '구글기프트카드', '안전결제 링크')가 채팅에서 발견되어도 어떠한 경고/필터링도 없다. 신규 사용자는 이런 패턴을 모를 가능성 높고, 신고가 누적된 후에야 대응 가능. 백엔드 도메인 검증도 없고 프런트엔드 표시 보호도 없다."
  proposal: "(1) 백엔드에 keyword blocklist (예: ['카톡', '디코', '먼저 입금', '안전결제', '문화상품권']) — domain validation 으로 정의. (2) handleSendMessage 에 키워드 매칭 시 metadataJson.scamFlags=['offsite_channel', 'prepay'] 추가. (3) 프런트는 scamFlags 가 있으면 ChatMessage 의 메시지 버블 상단에 '⚠️ 거래 사기 의심 패턴' 작은 banner. (4) 신고 모달 진입 시 자동으로 의심 메시지 highlight."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_chat.go L164-232 handleSendMessage: text validation 은 'oneof=text image' 만, 본문 키워드 0건 검사. 백엔드 internal/domain 영역에 scam_filter.go 또는 유사 파일 0개. 프런트 ChatMessage 도 scam 분기 0건. 거래 사기 관련 사용자 보호 zero."

- id: imp-0239
  found_at_iter: 16
  area: chat
  type: ux
  target: send_button_loading_disabled_state_visible
  problem: "ChatInput 의 전송 버튼은 useSendMessage.isPending 상태와 무관하게 disabled={disabled || !text.trim()} 만 평가한다. 사용자가 빠르게 두 번 클릭하거나 Enter 를 연속 누르면 같은 텍스트가 두 번 mutate 될 수 있고(clientMessageId 중복 검사가 백엔드에 있지만 Race), 사용자는 첫 클릭이 처리되었는지 시각적 피드백을 얻지 못해 '내가 보낸 거 맞나?' 인지 부조화."
  proposal: "(1) ChatInput 에 isSending: boolean prop 추가 (parent 가 sendMessage.isPending 전달). (2) isSending 일 때 button disabled + 아이콘을 회전 spinner '⟳' 로 교체 + 라벨 '전송중'. (3) 메시지 보낸 직후 즉시 textarea clear(이미 됨) 는 유지하되 placeholder 를 '전송 중...' 로 잠깐 표시. (4) 동일 텍스트를 5초 이내에 보낸 경우 토스트 'X초 전에 보낸 동일 메시지가 있어요. 다시 보낼까요?'."
  effort: trivial
  impact: low
  evidence: "코드 web/components/chat/chat-input.tsx L65-72 button: disabled={disabled || !text.trim()} — pending 상태 0건. 호출처 web/app/chats/[id]/page.tsx L174 onSend={... sendMessage.mutate(...)} 도 isPending 미전달. 같은 텍스트 중복 송신 가드 백엔드 ClientMessageID 만 의존(race window 존재)."

- id: imp-0240
  found_at_iter: 16
  area: chat
  type: feature
  target: chat_room_archive_hide_completed
  problem: "거래 완료(deal_completed) 또는 매물 삭제(listingStatus='deleted') 후에도 채팅방이 영구히 list 에 남는다. 활성 거래 사용자는 수십 개의 끝난 채팅방 사이에서 진행 중인 채팅을 찾기 위해 매번 스크롤. 또한 끝난 채팅방을 다시 확인하고 싶을 때 검색/필터도 없어 결국 '안 사라짐 = 노이즈' 인 상태."
  proposal: "(1) 백엔드 chat_rooms 에 archived_at TIMESTAMP NULL 추가 + POST /api/v1/chats/:id/archive 엔드포인트. (2) deal_completed 또는 listing_deleted + 7일 경과 시 자동 archive(스케줄러). (3) 사용자도 채팅방 메뉴에서 '아카이브' 가능. (4) ChatPanel 의 chats list 는 archive 제외하고 표시, 상단 토글 '아카이브 보기' 클릭 시만 노출. (5) 아카이브된 채팅에 새 메시지 도착 시 자동 unarchive."
  effort: medium
  impact: medium
  evidence: "코드 backend chat_rooms 스키마: archived_at 컬럼 미존재 추정(handlers_chat.go ListChatRooms 가 모든 row 반환). 프런트 web/app/chats/page.tsx 는 chats.map 만 — 필터 toggle 0개. 사용자 chat 정리 entry 0건."

- id: imp-0241
  found_at_iter: 16
  area: chat
  type: performance
  target: optimistic_message_dedupe_on_sse_arrival
  problem: "SSE 의 new_message 이벤트와 useSendMessage 의 onSuccess 응답이 동시에 도착할 수 있어 동일 메시지가 캐시에 두 번 추가될 가능성이 있다. 백엔드는 broker.SendToUser 를 counterpart 한 명에게만 보내지만, sender 가 multi-tab 인 경우(탭 A 송신 → 탭 B 수신) 또는 광범위한 invalidate 결과 동일. 또한 clientMessageId 는 sender 측에서만 매칭되므로 receiver 측에서는 messageId 중복 가드가 필수."
  proposal: "(1) SSE new_message handler 에서 setQueryData 하기 전 기존 pages 를 순회해 messageId 중복이면 무시. (2) optimistic prepend 시 senderUserId='_optimistic_' 마커 사용 중 — onSuccess 시 clientMessageId 매칭으로 교체하지만, 실패 시 이중 노출. (3) Map<messageId, Message> 자료구조로 dedupe 후 array 반환하는 selector 도입. (4) sender 측에도 broker 가 echo 보내면 자기 메시지 dedupe 보장."
  effort: small
  impact: low
  evidence: "코드 web/lib/hooks/use-chats.ts L73-89 onSuccess: messageId === clientMessageId 일 때 교체 — SSE 가 server 응답 도착 전 먼저 도달하면? clientMessageId 가 metadata 에 보관되지 않아 매칭 실패. backend handlers_chat.go L211-218 msgPayload 에 clientMessageId 미포함. SSE handler 에서 dedupe 로직 미확인(use-sse.ts 추가 점검 필요)."

- id: imp-0242
  found_at_iter: 17
  area: reservation
  type: bug
  target: reservation_modal_note_field_silently_dropped
  problem: "ReservationModal 의 메모 textarea(reservation-modal.tsx L17,28,53)는 form.notes 키로 상태를 보관하고 createReservation 페이로드에 'notes: form.notes' 로 전송한다. 그러나 backend handlers_reservation.go L25 가 받는 필드는 'noteToCounterparty' 이며, 'notes' 키는 ShouldBindJSON 에서 무시된다. 결과적으로 사용자가 입력한 메모(예: '아이템 합쳐서 1500만 으로 협의' 등 거래 합의 내용)가 DB note_to_counterparty 컬럼에 NULL 로 저장되고, 상대방 카드에는 표시되지 않는다. 100% 데이터 손실 버그인데 어떤 에러도 없어 발견 불가."
  proposal: "(1) reservation-modal.tsx L28 'notes:' → 'noteToCounterparty:' 로 키 변경(또는 wrapper 에서 매핑). (2) Flutter reservation_form_sheet.dart 도 같은 키 검증. (3) backend req struct 에 binding tag 추가하여 필수 필드 강제(현재 *string 옵션이라 누락 무방). (4) 회귀 방지: __tests__/components/forms/reservation-modal.test.tsx 에 'createReservation 호출 시 noteToCounterparty 키로 전송된다' 단위 테스트 추가. (5) 기존 카드 metadataJson 표시 로직(reservation-card-message.tsx L20)도 'notes' 키만 읽어서 새 데이터(noteToCounterparty)와 정합성 점검 필요."
  effort: trivial
  impact: high
  evidence: "코드 web/components/forms/reservation-modal.tsx L17 useState({...notes:''}), L28 createReservation({notes: form.notes}). backend/cmd/server/handlers_reservation.go L25 req.Note `json:\"noteToCounterparty\"`. JSON unmarshal 시 키 미스매치 → req.Note=nil → DB note_to_counterparty 컬럼 NULL. 사용자/QA 입장에서 어떤 에러도 없어 발견 매우 어려움."

- id: imp-0243
  found_at_iter: 17
  area: reservation
  type: bug
  target: reservation_invalidate_messages_missing
  problem: "예약 제안 성공 시 onCreated 콜백(chats/[id]/page.tsx L180)이 ['chats'] 만 invalidate 하고 ['messages', chatId] 는 invalidate 하지 않는다. 백엔드는 reservation 생성과 동시에 chat_messages 에 message_type='system'(또는 'reservation_card') 행을 INSERT 하지만(postgres_reservation.go L64-67), 사용자 화면에는 SSE 를 통해서만 새 시스템 메시지가 전달된다. SSE 가 끊긴 상태/재연결 중/탭 비활성화이면 사용자가 수동 새로고침 전까지 자기 예약 카드를 볼 수 없다. 또한 모달 닫힘 → 빈 채팅 화면 → 카드 없음 → '예약이 안 됐나?' 혼란."
  proposal: "(1) onCreated 를 () => { qc.invalidateQueries({queryKey:['chats']}); qc.invalidateQueries({queryKey:['messages', id]}); } 로 변경. (2) confirmReservation/cancelReservation 액션 추가 시(imp-0091)도 동일 패턴. (3) 더 견고하게는 mutation 자체에 onSuccess 를 두어 component 단 콜백 의존을 줄임 — useCreateReservation hook 신설. (4) optimistic insert 까지 가면 모달 닫는 순간 자기 카드가 즉시 보여 SSE 의존 제거."
  effort: trivial
  impact: high
  evidence: "코드 web/app/chats/[id]/page.tsx L180 onCreated={() => qc.invalidateQueries({ queryKey: ['chats'] })} — ['messages', id] 누락. backend/internal/repository/postgres_reservation.go L64-67 INSERT chat_messages 시스템 메시지 — useMessages 가 stale cache 유지. SSE off 시 0 reflective UI."

- id: imp-0244
  found_at_iter: 17
  area: reservation
  type: feature
  target: reservation_reschedule_state_machine
  problem: "예약이 confirmed 된 후 시간/장소 변경(상대방과 합의)을 하려면 현재는 1) 예약 취소 → 2) 새 예약 제안 의 2-step 으로만 가능하다. 그 사이 listing.status 가 reserved→available 로 잠시 풀리면서 다른 사용자가 가로채는 race 가 있고, 취소 사유 코드가 'reschedule' 같은 의미를 잃어버리며, 양쪽 알림도 '취소됨' '새 제안' 로 분절되어 거래 맥락이 끊긴다. PRD.md L4303-4308 의 rescheduleState 머신(requested/counter_proposed/accepted/rejected/expired)은 정의되어 있으나 backend/web 어디에도 구현 0건."
  proposal: "(1) 백엔드: POST /reservations/:id/reschedule { newScheduledAt, newMeetingType, newMeetingPoint, scope: 'time_only|location_only|both' } 엔드포인트 신설. 원 예약 status 유지(잠금 유지) + reservation_reschedule_requests 테이블에 row 생성. 상대방 confirm/reject 으로 분기. (2) 프런트: 채팅방 헤더 [예약 변경] 버튼(예약 confirmed 일 때만), 모달은 ReservationModal 재사용 + '변경 사유' 필드. (3) 카드 분리: ReservationRescheduleCard 시스템 메시지로 '변경 요청', accept 시 원 카드 데이터 갱신. (4) listing reserved 상태 유지 → race 차단."
  effort: large
  impact: high
  evidence: "코드 backend grep -r 'reschedule' = 0건(PRD 외). 현재 변경 플로우는 cancel → create 2-step. PRD.md L4303 rescheduleState vocabulary 정의. STATE_SEQUENCE_DIAGRAMS.md L130 confirmed 상태에서 reschedule 전이 미정의. listing reserved_chat_room_id 가 cancel 시 NULL 로 풀려 race window 존재."

- id: imp-0245
  found_at_iter: 17
  area: reservation
  type: feature
  target: reservation_calendar_export_ics
  problem: "예약이 confirmed 되어도 사용자가 자기 캘린더 앱(아이폰 캘린더, Google Calendar)에 약속을 등록할 방법이 없다. ReservationCardMessage 는 readonly 표시일 뿐 .ics 다운로드/구글캘린더 추가 버튼이 없다. 거래 약속을 잊는 no-show(domain.AlignmentTradeExpired = -3 페널티) 의 큰 원인이다. 사용자가 '나중에 카톡 알림으로 떠올리겠지' 라는 외부 의존을 강제 받아 자기 일정관리 시스템과 분리된다."
  proposal: "(1) backend/cmd/server/handlers_reservation.go 에 GET /reservations/:id/ics 추가하여 RFC5545 .ics 파일 stream(BEGIN:VCALENDAR/VEVENT/SUMMARY=거래 약속/DESCRIPTION=listing.title+meetingType/DTSTART=scheduledAt/LOCATION=meetingPoint/END:VEVENT). (2) ReservationCardMessage 에 status='confirmed' 일 때 [📅 캘린더에 추가] 링크(<a download> 으로 .ics 다운로드, 모바일은 자동으로 캘린더 앱 열림). (3) Google Calendar 직접 링크: https://calendar.google.com/calendar/render?action=TEMPLATE&text=...&dates=... 별도 옵션. (4) 알림 설정: 약속 30분 전 사용자 캘린더 자체 알람으로 노쇼 감소."
  effort: small
  impact: medium
  evidence: "코드 web/components/chat/reservation-card-message.tsx L8-28: <a href='*.ics'>/calendar/render 링크 0건. backend handlers_reservation.go GET /reservations/:id 자체가 없음(POST 만 존재). domain/models.go L98 AlignmentTradeExpired=-3 페널티는 강한데 사용자 측 reminder 도구 0개."

- id: imp-0246
  found_at_iter: 17
  area: reservation
  type: feature
  target: reservation_quick_status_chat_macros
  problem: "예약 confirmed 이후 약속 시점 부근에 사용자가 보내는 메시지는 패턴이 매우 좁다 — '5분 늦어요', '도착했어요', '카톡주세요', '캐릭터 어디?', '거래소 앞에서 만나요'. 그러나 ChatInput 은 빈 텍스트 입력창만 제공하고 빠른 응답 매크로/템플릿 0개. 모바일에서 매번 한글 입력 + 오타 수정의 마찰이 있어 '늦었는데 늦었다고 보낼 시간도 없는' 상황 → 무응답 → 노쇼 페널티로 직결."
  proposal: "(1) ChatInput 위에 reservation_confirmed 상태일 때만 노출되는 quick-reply 칩 row: [5분 늦어요] [도착했어요] [거래소 앞] [캐릭터 위치 알려주세요] [잠시만요]. 클릭 시 즉시 send 또는 편집 모드(shift-click). (2) 약속 시간 ±15분 윈도우에서는 칩이 자동으로 더 강조(border-gold). (3) 사용자별 자주 쓰는 매크로 5개 학습(localStorage 'rsv_macros'). (4) 매크로 클릭 시 자동으로 system event 'reservation_eta_update' 메타데이터 첨부 → 상대방 카드에도 '5분 후 도착 예정' 갱신. (5) 시각/거리(meetingPoint) 컨텍스트 기반 매크로 추천(KST timezone+오프라인이면 '거래소 앞', 인게임이면 '서버 어디?')."
  effort: medium
  impact: medium
  evidence: "코드 web/components/chat/chat-input.tsx 가 textarea 단일. quick-reply 칩 0건. 백엔드 message_type 에 system/text/image 만 있고 'eta_update' 같은 카테고리 0건. 약속 시점 30분 전후 상대방 응답 SLA 측정 도구 부재."

- id: imp-0247
  found_at_iter: 17
  area: reservation
  type: feature
  target: reservation_no_show_claim_flow
  problem: "약속 시간이 지났는데 상대방이 안 나타난 경우(no-show)를 처리할 UI가 없다. domain/models.go L65,68 의 ReservationNoShowReported='no_show_reported' 와 AlignmentTradeExpired=-3 는 정의돼 있고, report-modal 의 reportType 에 'no_show'(report-modal.tsx L13) 가 있긴 하지만 이는 일반 신고 폼이라 reservation context(시간/장소/대기시간/증거)가 빠진다. PRD.md L237 'no-show claim 의 생성 조건, 중복 제한, accepted reschedule 와의 관계' 정책은 있으나 web 0건."
  proposal: "(1) ReservationCardMessage 가 status='confirmed' AND scheduledAt + 30min < now 이면 [상대방이 안 와요] 버튼 노출. (2) 클릭 → no-show claim 모달: 대기 시작/종료 시각 자동 기록, 대기 위치 메모, 증거 텍스트('장로 인근에서 30분 대기, 카톡/귓속말 무응답'), 증거 이미지(스샷 1장). (3) backend POST /reservations/:id/no-show-claim 엔드포인트 신설 → reservations.status='no_show_reported' + reports 자동 생성 + 상대방 alignment -3. (4) 24시간 내 상대방이 'reschedule_request' 또는 'reason_explain' 으로 응답 가능(중복 제한). (5) 양쪽 모두 노쇼 클레임이 있으면 admin 큐에 자동 적재."
  effort: large
  impact: high
  evidence: "코드 backend handlers_reservation.go grep 'no_show|noShow' = 0건. domain/models.go L68 ReservationNoShowReported 상태는 정의만 됨. report-modal.tsx L13 no_show 옵션은 일반 신고로 거래 컨텍스트 미포함. PRD.md L237 no-show 정책이 명시되어 있으나 코드 0건."

- id: imp-0248
  found_at_iter: 17
  area: reservation
  type: feature
  target: reservation_listing_busy_indicator_for_buyer
  problem: "구매자가 매물 상세에서 채팅 시작 → 예약 제안 시도 시 backend가 'CONFLICT: 이미 활성 예약이 존재합니다'(handlers_reservation.go L44-47)를 반환할 수 있다. 그런데 이 정보는 사용자가 모달 submit 까지 가서야 알게 된다. 매물 카드/상세 페이지 어디에도 '이 매물은 다른 사용자와 예약 진행 중' 표시가 없다(listing.status='available' 로 표시). 결과적으로 채팅방 시작 → 예약 모달 열기 → 입력 → submit → 실패 → 좌절 의 4단계 마찰."
  proposal: "(1) listing 응답에 has_active_reservation:bool 또는 active_reservation_count 필드 추가(SELECT COUNT(*) FROM reservations WHERE listing_id=... AND status IN ('proposed','confirmed')). (2) 매물 카드에 'reserved' 가 아니라도 '🔒 예약 협의 중' 노란 배지 노출 — 다른 구매자에게 '먼저 진행 중인 거래가 있다' 시그널. (3) 매물 상세 [채팅 시작] 버튼 disabled 까지는 아니더라도 클릭 시 confirm '다른 사용자와 협의 중인 매물입니다. 그래도 채팅을 시작할까요?'. (4) 채팅방 헤더 [예약 제안] 버튼 자체를 disable + tooltip '다른 거래가 진행 중'."
  effort: medium
  impact: medium
  evidence: "코드 backend handlers_listing.go GetListing 응답 필드에 active_reservation 정보 0건. web/components/listing/* 카드 표시에서 reservation 상태 분기 0건. CONFLICT 에러는 backend 에 있으나 프런트는 사후 toast 만(reservation-modal.tsx L33)."

- id: imp-0249
  found_at_iter: 17
  area: reservation
  type: feature
  target: reservation_expires_at_user_visibility
  problem: "DB schema(reservations.expires_at)와 backend req(handlers_reservation.go L26 ExpiresAt *string)는 만료 시각을 받지만 1) ReservationModal 폼에 expires_at 입력 필드가 없어 사용자가 직접 설정 불가, 2) backend는 default 만료 자동 계산 로직도 없음, 3) ReservationCardMessage 는 expires_at 표시 0건이라 '응답 시간이 얼마나 남았는지' 알 수 없음, 4) 만료 발생 시 정리(expired status 갱신)하는 sweep job 도 backend/cmd/server/main.go L39 에 refresh_token 만 정리하고 reservations 0건."
  proposal: "(1) ReservationModal 에 '응답 대기 시간' 드롭다운(1시간/3시간/12시간/24시간/만료없음) — default '24시간'. (2) backend handlers_reservation.go 에서 ExpiresAt 이 없으면 NOW() + interval '24 hour' 자동 설정. (3) ReservationCardMessage 에 status='proposed' 이면 '⏰ 응답 마감 23시간 12분' 라이브 카운터. (4) main.go 에 1시간 주기 goroutine 추가: UPDATE reservations SET status='expired' WHERE status='proposed' AND expires_at < NOW(); → chat_status='open' 복구 + system 메시지 '응답 시간이 지나 자동 만료되었어요'. (5) 만료 시 proposer alignment -1(domain.AlignmentTradeExpired 활용)."
  effort: medium
  impact: high
  evidence: "코드 reservations.expires_at TIMESTAMPTZ NULL(001_initial.sql L_). web reservation-modal.tsx L17 form 객체에 expiresAt 필드 0건. 카드 표시(reservation-card-message.tsx) 카운트다운 0건. backend main.go L39 sweep 은 refresh_token only — reservations expired 정리 cron 0건. domain.AlignmentTradeExpired 정의되어 있으나 트리거 코드 0건."

- id: imp-0250
  found_at_iter: 17
  area: reservation
  type: a11y_mobile
  target: reservation_card_action_44px_and_status_color
  problem: "imp-0091 가 [확정/거절/취소] 버튼 신설을 제안했는데, 그 위에 추가로 모바일 a11y 관점이 있다. 1) ReservationCardMessage 자체가 max-w-[70%] 로 채팅 버블이라 모바일 375px 에서 약 260px 폭 — 그 안에 버튼 2개를 가로 배치하면 각 ~120px 폭으로 thumb tap 어려움. 2) 카드 border 가 status 와 무관하게 항상 border-gold 라서 시각 위계가 없어 confirmed/cancelled/expired 가 같은 색으로 보임. 3) 카드 안 텍스트가 text-xs(12px)/text-sm(14px) 혼재로 약속 시간이 가장 중요한데 가장 작아 보임."
  proposal: "(1) 액션 버튼은 모바일에서 grid-cols-1 stacked, sm:grid-cols-2 에서만 가로(min-h-[44px] 강제). (2) status 별 border/bg 토큰: proposed=border-gold/bg-card, confirmed=border-success/bg-success/5, cancelled=border-danger/bg-danger/5 opacity-70, expired=border-text-dim/bg-card opacity-50. (3) 카드 내 typography 위계 재설계: 약속 시간을 text-base font-semibold(상위), 접선방식 text-sm, 메모 text-xs. (4) 카드 상단 status pill 추가('진행중'/'확정됨'/'취소됨'/'만료'). (5) 다크모드 컨트래스트 4.5:1 검증 — 현재 text-text-secondary 는 확실치 않음."
  effort: small
  impact: medium
  evidence: "코드 web/components/chat/reservation-card-message.tsx L14 className='max-w-[70%] border border-gold' — status 분기 0건, color token 1종. L15 'text-xs' 가 가장 강조될 약속 시간보다 더 큼. L16-21 약속 시간 'text-sm' 상대적으로 작음. button 디자인(아직 미존재) 의 모바일 stacked 정책 0건."

- id: imp-0251
  found_at_iter: 17
  area: reservation
  type: ux
  target: reservation_history_archive_in_profile
  problem: "거래 완료(reservation.status='fulfilled') 또는 취소(cancelled) 후 예약은 채팅 메시지 안에 살아있긴 하지만, 사용자가 '내가 지난 6개월 간 어떤 약속을 잡았는지/노쇼 횟수가 얼마인지' 같은 자기 거래 이력을 한눈에 보는 페이지가 없다. /profile/trades 는 chat_room 단위 list 라 한 채팅방에 여러 예약이 있어도 1개로만 보이고, reservation 상태별 필터(fulfilled/cancelled/expired)도 없다. 자기 평판 관리(예: 노쇼 0회 자랑) 와 분쟁 시 증거 회수에 약점."
  proposal: "(1) /profile/me/reservations 신설 페이지 — 백엔드 GET /me/reservations { status?, from?, to? } 추가. (2) 카드 list: 매물명, 일시(KST), 접선방식, 상대방 닉네임, 상태 배지, 결과(완료/취소/노쇼/만료), [상세] 링크. (3) 상단 탭 [전체][확정][완료][취소/만료] + 기간 필터(이번 달/3개월/6개월/전체). (4) 통계 헤더: 완료 N건, 노쇼 N건, 평균 응답시간. (5) profile.reviewer 관점에서 자기 노쇼 0건이 자랑이라면 공개 프로필에 'verified streak' 배지 노출."
  effort: medium
  impact: medium
  evidence: "코드 web/app/profile/trades/page.tsx 는 chat_rooms 단위만 표시. /profile/me/reservations 라우트 0건. backend GET /me/reservations 0건(POST /chats/:id/reservations 만 존재). 통계 집계(완료/노쇼 카운트) 0건."

- id: imp-0252
  found_at_iter: 18
  area: review
  type: bug
  target: review_create_missing_participant_authorization
  problem: "POST /trade-completions/:compId/reviews 핸들러가 요청 사용자가 거래의 참가자(requested_by 또는 counterpart) 인지 검증하지 않는다. 코드 handlers_review.go L29-38: GetCompletionForReview 로 status='confirmed' 만 확인하고, 그 다음 'if userID == info.CounterpartUserID { targetUser = info.RequestedByUserID } else { targetUser = info.CounterpartUserID }' 분기는 단순 매핑만 한다. 즉 제3자(완전히 다른 사용자) 가 valid completionId 를 알면 누구에게나 후기를 쓸 수 있고, 그 후기가 두 거래 참가자 중 한 명에게 적립된다. 평판 시스템의 무결성을 깨는 보안 결함."
  proposal: "L33 직후 명시적 권한 체크 추가: 'if userID != info.CounterpartUserID && userID != info.RequestedByUserID { c.JSON(http.StatusForbidden, gin.H{...code:\"FORBIDDEN\"...}); return }'. 또한 alignment.Change(reviewer) 가 거래 참가자가 아닌 사용자의 alignment 도 변경할 수 있어 부정 적립 발생 — 권한 체크는 alignment 호출 이전이어야 한다. 추가 방어선으로 reviews INSERT 시 'WHERE EXISTS (SELECT 1 FROM trade_completions WHERE id=$2 AND (requested_by=$3 OR counterpart=$3))' 같은 조건 INSERT 검증."
  effort: trivial
  impact: high
  evidence: "코드 backend/cmd/server/handlers_review.go L29-38: GetCompletionForReview 후 status 만 검사, userID 가 RequestedByUserID/CounterpartUserID 인지 검증 0건. 분기는 단순 'userID == counterpart 면 target=requestedBy, 아니면 target=counterpart' — userID 가 어느 쪽도 아닌 경우 조용히 'else' 가지로 들어가 임의 user 에게 후기 적립. UNIQUE(completion_id, reviewer_user_id) 제약으로 중복은 막히지만 1회 공격은 통과."

- id: imp-0253
  found_at_iter: 18
  area: review
  type: feature
  target: review_received_notification_missing
  problem: "사용자가 후기를 받아도 알림이 발송되지 않는다. CreateReview 트랜잭션(postgres_reservation.go L270-306)이 reviews INSERT + user_profiles 카운터 업데이트 + alignment 적립만 하고, notifications 테이블 insert 또는 SSE event publish 호출 0건. ripgrep 'review' across backend/internal/event 결과 0건. 즉 셀러는 자기 프로필을 직접 새로고침하기 전까진 새 후기가 달렸는지 모르며, negative 후기에 대한 답변·소명 기회도 늦어진다. 후기 작성자는 자기 후기가 카운트됐는지 확인하기 위해 매번 상대 프로필을 봐야 한다."
  proposal: "(1) CreateReview 트랜잭션 끝에 notifications INSERT 추가: 타입 'review_received', target_user_id=reviewee, payload={ reviewerNickname, rating, listingTitle }. (2) SSE /sse/notifications 채널로 즉시 push — 사용자가 알림 벨 클릭 시 'OO님이 좋았어요 후기를 남겼습니다' 카드. (3) negative 후기는 추가로 '답변하기(7일 이내)' CTA 포함(imp-0109 의 seller response 도입과 결합). (4) 7일 후 자동 정리 cron 에 review_received 도 포함."
  effort: small
  impact: high
  evidence: "코드 backend/internal/repository/postgres_reservation.go L270-306 CreateReview 트랜잭션 — notifications INSERT 0건, event publish 0건. backend 디렉토리 grep 'review.*notif|notif.*review' 0건. SSE 핸들러(handlers_sse.go) 에서 review 이벤트 publish/subscribe 0건. 사용자는 프로필 페이지 reload 전까지 새 후기 인지 불가."

- id: imp-0254
  found_at_iter: 18
  area: review
  type: bug
  target: review_list_silent_50_limit_data_loss
  problem: "ListUserReviews SQL 에 'LIMIT 50' 이 하드코딩되어 있다(postgres_reservation.go L312). imp-0115 의 페이지네이션 제안과 별개로, 현 상태는 '데이터 일부가 사일런트 누락' 이라는 더 무거운 문제다. 후기 51개를 받은 셀러의 51번째 후기는 API 응답에서 빠지고, 클라이언트는 'reviews.length === 50' 인 줄만 알고 더 있다는 신호도 받지 못한다. 받은 리뷰 헤더 '받은 리뷰 (50)' 는 거짓말이 되며, 후기 차트/통계(imp-0107)의 모든 ratio 계산이 잘못된 분모로 산출된다."
  proposal: "단기: LIMIT 50 → LIMIT 1000(현실적 상한)으로 즉시 완화 + 응답에 'truncated: true' 플래그 노출해 사용자에게 '최대 N개만 표시 중' 경고. 장기: imp-0115 의 cursor pagination 와 함께 totalCount 별도 SELECT COUNT(*) 응답 추가해 헤더 숫자가 실제 총량과 일치하도록 보장. 또한 user_profiles 에 positive_review_count + negative_review_count 카운터를 read-time 출처로 사용하면 LIMIT 영향 없이 총량 노출 가능."
  effort: trivial
  impact: high
  evidence: "코드 backend/internal/repository/postgres_reservation.go L308-312 SELECT … LIMIT 50 — 51번째 행부터 응답 누락. handlers_review.go L73 'data: reviews' 응답에 totalCount/truncated 필드 0건. 클라이언트 web/app/profile/[userId]/reviews/page.tsx L31 '받은 리뷰 ({reviews.length})' 가 실제 총량과 다를 수 있음."

- id: imp-0255
  found_at_iter: 18
  area: review
  type: feature
  target: review_request_auto_reminder_after_completion
  problem: "거래 완료(deal_completed) 후 후기 작성을 유도하는 reminder 가 0건이다. 사용자가 그 시점에 모달을 닫거나 다른 화면으로 이동하면 후기는 영영 작성되지 않는다. 결과적으로 '후기 작성률(reviews / completed_trades)' 이 매우 낮아 신뢰 시그널 양 부족 — 새 셀러가 신뢰 빌드업까지 가는 시간이 길어진다. 또한 운영 데이터 부족(거래 N% 만 후기로 검증됨) 으로 분쟁 시 알리바이도 약함."
  proposal: "(1) trade_completion 이 confirmed 되면 24h/72h 두 번 reminder 알림 발송(notifications + SSE). 본문 'OO 매물 거래는 어땠나요? 한 줄 후기로 다음 사람을 도와주세요'. (2) 7일 지나면 reminder stop, 14일까지는 작성 가능 윈도우. (3) profile/trades 의 '거래 완료' 카드에 '후기 작성하기 (D-X)' 카운트다운 배지 + 미작성 거래만 필터. (4) 대시보드 메트릭: 후기 작성률 KPI(주 단위)."
  effort: medium
  impact: medium
  evidence: "코드 backend 측 cron/scheduler grep — main.go L39 부근 sweep 은 refresh_token, notification cleanup 만(이전 commit 'feat(backend): 알림 7일 자동 정리'). 후기 reminder 관련 cron 0건. notifications.type 값에 'review_reminder' 없음(현재 도메인 모델 탐색 결과). chat-page 의 deal_completed 이후 reminder UI 0건."

- id: imp-0256
  found_at_iter: 18
  area: review
  type: feature
  target: review_verified_buyer_badge
  problem: "리뷰 카드는 reviewerNickname + rating + comment 만 보여줘 '실제로 거래한 사람의 후기' 라는 신뢰 시그널이 약하다. 외부 플랫폼은 'Verified Purchase' 배지로 마케팅성 가짜 후기와 실거래 후기를 구분한다. 본 시스템은 reviews.completion_id FK 가 있어 모든 후기가 본질적으로 verified buyer 인데, UI 가 그 사실을 사용자에게 전달하지 않는다 — 신뢰의 자원을 그냥 버리는 셈."
  proposal: "(1) 리뷰 카드에 작은 chip '거래 완료 후기' 또는 '✓ 실거래' 추가, 툴팁 '플랫폼이 거래 완료를 확인한 후기입니다'. (2) 카드 좌측 4px 색띠로 시각적 강조. (3) 백엔드 응답에 verified=true 평면 필드 명시(향후 import 된 외부 후기/마이그레이션 후기 구분 시 false 가능). (4) 매물 카드에도 '실거래 후기 N건' 으로 노출(imp-0105 와 결합)."
  effort: trivial
  impact: medium
  evidence: "코드 web/app/profile/[userId]/reviews/page.tsx L34-62 리뷰 카드: verified/실거래 배지 0건. backend handlers_review.go L70-72 응답: verified 필드 0건. 그러나 reviews.completion_id FK 는 존재하므로 (001_initial.sql L211) 모든 후기가 정의상 verified — 시각화만 누락."

- id: imp-0257
  found_at_iter: 18
  area: review
  type: feature
  target: review_with_photo_attachments
  problem: "리뷰는 텍스트 한 줄만 가능해 거래 시 받은 매물 사진/스크린샷 첨부가 불가능하다. 리니지 클래식 거래는 '약속한 옵션이 진짜였는지' 가 핵심 쟁점인데, 후기 사진(가령 인벤토리 스크린샷, 거래 영수증) 없이 텍스트만으로는 분쟁 발생 시 증거력이 약하고, 다른 구매자도 '이 셀러가 약속대로 매물을 줬는지' 시각적으로 검증 불가."
  proposal: "(1) ReviewModal 에 image upload 영역 추가(최대 3장, 5MB). (2) reviews 테이블에 image_urls TEXT[] 컬럼 추가, S3/local 스토리지 같은 ImageController 재사용. (3) 리뷰 카드에 thumbnail 그리드 + 클릭 시 lightbox. (4) 운영자 측 신고된 사진은 blur 처리 + manual review. (5) 사진 첨부 후기는 '사진 첨부 후기' 배지 + 정렬 옵션 '사진 후기만'."
  effort: medium
  impact: medium
  evidence: "코드 web/components/forms/review-modal.tsx L36-67: form 안 image input 0건. backend reviews 스키마(001_initial.sql L209-218): image_urls 컬럼 0건. 도메인 모델 review 에 사진 필드 0건. 매물 listing 은 ImageController + uploaded_images 테이블 활용 중이라 인프라 재사용 가능."

- id: imp-0258
  found_at_iter: 18
  area: review
  type: feature
  target: review_content_moderation_filter
  problem: "리뷰 comment 가 자유 텍스트인데 backend 단에서 욕설/PII(전화번호·계좌번호·디스코드 ID)/외부 링크/스팸 패턴 필터가 0건이다. handlers_review.go L20-23 에 binding 도 string 타입 검증만. 사용자 닉네임 + 후기는 공개 페이지(/profile/[userId]/reviews) 에 영구 노출되므로 악의적 후기 작성자가 '01012345678 사기꾼' '디코 fakeid 사칭함' 같은 PII/명예훼손 텍스트를 박으면, 신고-삭제 워크플로(imp-0111) 가 작동하기 전까지 검색엔진에 인덱싱될 수도 있다."
  proposal: "(1) backend 측 가벼운 필터: 한국어 욕설 사전(공개 OSS) + 정규식(전화번호, 한글 11자 이상 연속숫자, URL) 매치 시 자동 hidden_pending_review 상태로 INSERT, 운영자 검토 후 publish. (2) 또는 즉시 차단 + '욕설/연락처는 후기에 포함할 수 없습니다' inline 안내. (3) 외부 도메인 URL 은 strip 또는 mask. (4) 운영 대시보드에 'pending review' 큐. (5) 셀러 신고 1건 이상 누적된 후기는 자동 hidden + 검토 대기."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_review.go L20-23: comment *string binding 'oneof/regex' 0건. ripgrep 'profanity|badwords|moderation|filter' across backend 결과 0건. 리뷰 작성 후 즉시 LIST API 에 노출(workflow: insert→공개)되어 검토 단계 0개. 도메인 모델 reviews 에 status/visibility 컬럼 0건(001_initial.sql L209-218)."

- id: imp-0259
  found_at_iter: 18
  area: review
  type: ux
  target: review_modal_no_disabled_reason_tooltip
  problem: "ReviewModal 의 submit 버튼이 rating 미선택 시 disabled 되지만, 사용자에게 '왜 비활성화인가' 가 안내되지 않는다(L63 disabled={submitting || !rating}). 시각 단서는 opacity-50 뿐, hover/focus 툴팁/aria-describedby/inline 메시지 0건. 모바일 사용자가 '한 줄 코멘트(선택)' 만 입력하고 [리뷰 제출] 을 탭했을 때 아무 피드백 없이 버튼이 안 눌리면 '버그인가?' 로 오해 후 이탈."
  proposal: "(1) 버튼 아래 inline 안내 문구 — rating === null 이면 '평가를 먼저 선택해 주세요(좋았어요/아쉬웠어요)' 표시, 선택되면 fade-out. (2) submit 버튼에 aria-describedby='submit-helper-text' 로 SR 사용자에게도 사유 노출. (3) 사용자가 disabled 인 버튼을 탭하면 light shake animation + toast '평가 선택 필요'. (4) imp-0112 의 radiogroup 의무화와 결합."
  effort: trivial
  impact: low
  evidence: "코드 web/components/forms/review-modal.tsx L63: 'disabled={submitting || !rating}' — 그 옆/아래에 helper text/error message JSX 0건. aria-describedby 0건. 클릭 시 피드백(toast/shake) 핸들러 0건. opacity-50 시각 단서만 의존."

- id: imp-0260
  found_at_iter: 19
  area: report
  type: feature
  target: admin_report_detail_target_deep_link_and_context
  problem: "admin/app/reports/page.tsx 상세 패널은 'targetType: 매물, targetId: b155f6a7…' 같은 raw ID 만 표시한다. 운영자는 (1) 그 매물·채팅·메시지를 직접 보려면 manually URL 을 조립해 새 탭을 열거나 admin/listings 에서 ID 검색해야 하고, (2) 신고 시점에 매물 제목/가격/이미지·채팅 마지막 N건의 메시지·신고된 메시지 본문·신고자/대상자 닉네임·계정 가입일·alignment_score 등 판단에 필요한 컨텍스트를 한 화면에서 못 본다. 결과적으로 '신고 1건 처리에 5탭을 열어야 한다'는 운영 비효율로 SLA 가 늘어진다."
  proposal: "(1) admin reports 상세 패널에 deep link 자동 생성 — targetType==='listing' → 'admin/listings/:id' 와 'giranjt.com/listings/:id' 두 링크, 'chat_room/message' → 'admin/chats/:id' 모달. (2) 응답에 reporterNickname, reporterAccountStatus, reporterAlignmentScore, reporterPriorReportCount(허위/확정), targetPreview(listing snapshot json: title/price/images[0]·chat last 10 messages·message body) 추가. (3) handleAdminGetReport 가 target_type 별로 JOIN 해 한 번에 컨텍스트 페이로드 반환. (4) 'snapshot' 컬럼을 reports 에 저장(listing 가격/제목 변조 후 신고 회피 방지)."
  effort: large
  impact: high
  evidence: "코드 admin/app/reports/page.tsx L165-197 detail panel 에 targetId 만 raw 노출, <a href> 또는 deep link 0건. backend/cmd/server/handlers_admin.go L42-61 handleAdminGetReport 응답에 reporterNickname/targetPreview/priorReportCount 필드 0건 — JOIN/쿼리 0건. db/migrations/001_initial.sql L223-235 reports 테이블에 snapshot/preview 컬럼 0건."

- id: imp-0261
  found_at_iter: 19
  area: report
  type: feature
  target: auto_hide_listing_after_n_reports_threshold
  problem: "reports 테이블에 같은 listingId 대상 5건·10건 신고가 누적되어도 자동으로 listing.visibility='temp_hidden' 으로 전환하는 로직이 0건이다(handlers_report.go·main.go·repository grep). 결과적으로 (1) 운영자가 알아챌 때까지 사기 의심 매물이 검색·홈 노출에 그대로 남고, (2) 1인 운영 환경에서 신고가 들어오는 속도 > 처리 속도면 피해자가 추가 발생. 임계 신고가 들어와도 사람이 직접 확인할 때까지 자동 보호 장치 0건."
  proposal: "(1) handleCreateReport 트랜잭션 끝에 동일 (target_type='listing', target_id) reports COUNT 가 N(예: 5) 이상이고 타입이 scam_suspicion/fake_listing/prohibited_item 이면 listings.visibility='temp_hidden' + 사유 'auto_hidden_threshold_5' 로 업데이트. (2) 자동 숨김 시 운영자에게 SSE 알림('5건 누적 신고로 자동 숨김 — 검토 필요'). (3) target_type='user' 도 동일 로직(누적 임계 → account_status='auto_review'). (4) 자동 조치는 audit_logs 에 actor_role='system' 으로 기록, 운영자가 24시간 내 검토하지 않으면 매일 리마인드. (5) 매물 작성자에게 '귀하의 매물이 다수 신고로 자동 임시숨김 처리되었습니다 — 이의제기' 알림."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_report.go L29-39 CreateReport 트랜잭션 끝에 SELECT COUNT/UPDATE listings 로직 0건. backend/internal/repository/postgres_reservation.go CreateReport 함수 — 후속 자동 액션 0건. main.go 안 background goroutine 두 개(refresh_token, notification cleanup) 만 존재하고 reports 임계 모니터 0건. listings.visibility 값에 'temp_hidden'/'auto_hidden' 정의 0건(domain/models.go grep)."

- id: imp-0262
  found_at_iter: 19
  area: report
  type: feature
  target: false_report_penalty_and_reporter_alignment
  problem: "신고가 reject 되거나 'false_positive' 로 판정되어도 reporter 의 alignment_score 변동이 0건이다(domain/models.go L93-100 AlignmentReportConfirmed 상수만 존재, AlignmentFalseReport·AlignmentReporterAbuse 정의 0건). 따라서 악의적인 reporter 가 경쟁 매물·싫은 사용자를 무차별 신고해도 본인 페널티 0이고, 운영자 큐만 어지럽혀진다. handleAdminUpdateReportStatus 의 status='resolved' 도 confirmed/rejected 구분 없이 단일 enum 이라 거짓 신고 통계도 못 낸다."
  proposal: "(1) reports.status enum 확장: submitted, assigned, resolved_confirmed(타깃 잘못), resolved_rejected(reporter 잘못), resolved_inconclusive(증거 부족), withdrawn. (2) admin 처리 UI 에서 'reject 사유' 선택(증거 부족·중복·정책 외·악의적 신고). (3) resolved_rejected 면 reporter alignment -2(첫 false report 는 학습 오류 인정), 같은 reporter 의 누적 false report ≥ 3 이면 -10 + 24시간 신고 권한 정지(중간 단계). (4) 누적 false report ≥ 5 이면 신고 권한 영구 박탈 + 운영자 검토 큐로 이동. (5) reporter dashboard 에 'false report 점수' 노출(이의제기 가능)."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_admin_audit.go L223 'oneof=submitted assigned resolved' — confirmed/rejected 구분 0건. backend/internal/domain/models.go L93-100 alignment 상수에 ReportConfirmed(-20) 만 정의, FalseReport/ReporterAbuse 0건. handleAdminReportAction 도 무조건 alignment.Change(req.TargetUserID, AlignmentReportConfirmed) 호출 — reporter 점수 조정 0건(L102)."

- id: imp-0263
  found_at_iter: 19
  area: report
  type: feature
  target: report_appeal_flow_for_target_users
  problem: "운영자가 admin/reports/:reportId/actions 로 'warning/temp_hide/permanent_hide/restrict/suspend' 조치를 내리면 target user 는 alignment -20 + 계정 제한을 받지만, 그 결정을 다투는 appeal 채널이 0건이다. 즉 운영자 1명이 신고 내용만 보고(스크린샷 0건, 컨텍스트 부족) 잘못된 조치를 내려도 사용자가 '저는 그런 짓 안 했습니다' 라고 항변할 공식 경로가 없다. /me/appeals, /api/me/appeals POST/GET 엔드포인트, appeals 테이블, admin/app/appeals 페이지가 모두 0건."
  proposal: "(1) appeals 테이블 추가 (id, target_user_id, related_report_id, related_action_id, body TEXT, evidence_files JSONB, status: pending/under_review/upheld/rejected, admin_response, created_at, resolved_at). (2) handleCreateAppeal POST /me/appeals (target_user_id 가 alignment 가 -20 받은 자거나 moderation_action 의 target_user_id 인 경우만 허용). (3) admin/app/appeals 페이지에 큐 + 'upheld'(원본 조치 취소 + alignment +20 복원 + reporter 페널티) / 'rejected'(원본 조치 유지) 버튼. (4) 사용자 sanction notification 에 '이의제기하기' CTA 포함."
  effort: large
  impact: high
  evidence: "코드 ripgrep 'appeal' across backend/cmd/server/* 결과 0건(grep). db/migrations/*.sql 에 appeals 테이블 0건. admin/app 디렉토리에 appeals/ 폴더 0건(ls). frontend/lib/features/* 에 appeal 모듈 0건. handleAdminReportAction L102 alignment.Change 후 reverse path 0건."

- id: imp-0264
  found_at_iter: 19
  area: report
  type: feature
  target: report_response_time_sla_dashboard
  problem: "운영자/대표가 '평균 신고 처리 시간', '24시간 이상 미처리 건수', 'SLA 위반율' 같은 지표를 한눈에 볼 대시보드가 0건이다. admin/app/page.tsx(메인 대시보드)는 listings/users/trades KPI 만 있고 reports SLA 카드 0건(use-dashboard.ts grep). 결과적으로 운영자는 자기 처리 속도가 느려지고 있는지 자각할 메커니즘이 없고, 사용자에게 SLA 약속(예: '24시간 내 처리')도 못 한다. 외부 transparency report(분기별 처리 통계 공개)도 데이터 백엔드 0건."
  proposal: "(1) admin 대시보드에 'Reports SLA' 카드: 평균 처리시간(median), p90 처리시간, 24h 초과 미처리 건수, 사유별 처리시간 분포(스파크라인). (2) backend GET /admin/reports/stats?range=7d|30d|90d 응답 추가 — reports 테이블에 resolved_at 컬럼(현재 updated_at 만 있음, 명시적 분리 필요) 추가 마이그레이션. (3) /reports/transparency 공개 페이지(인증 불필요) — 분기별 신고 접수/처리 카운트 + 사유별 비율(개인정보·금액 정보는 마스킹) — 사용자 신뢰 향상. (4) SLA 위반(접수 후 48h+) 신고에는 admin 알림 자동 발송."
  effort: large
  impact: medium
  evidence: "코드 admin/lib/hooks/use-dashboard.ts grep 'reports' 0건. admin/app/page.tsx 메인 대시보드에 reports SLA KPI 0건. backend/cmd/server/handlers_admin*.go 에 /admin/reports/stats 엔드포인트 0건. db/migrations/001_initial.sql L233-234 reports 에 resolved_at 별도 컬럼 0건(updated_at 으로 추정)."

- id: imp-0265
  found_at_iter: 19
  area: report
  type: feature
  target: reporter_withdraw_report_endpoint
  problem: "신고 후 사용자가 '오해였다, 잘못된 신고였다, 당사자끼리 해결됐다' 라고 깨달아도 자기 신고를 철회할 채널이 0건이다. /me/reports/:reportId DELETE 또는 PATCH status=withdrawn 엔드포인트 0건(main.go L137, handlers_report.go grep). 결과: (a) 잘못된 신고가 운영자 큐에 영원히 남아 처리 시간 낭비, (b) 운영자가 잘못된 신고 토대로 alignment -20 처리하면 비가역적 피해, (c) 사용자가 신고를 철회하지 못하니 '차라리 신고하지 말걸' 학습 → 경계 사례 신고 위축."
  proposal: "(1) PATCH /me/reports/:reportId/withdraw 엔드포인트 추가, 본인 reporter 만 가능, status==='submitted' 일 때만(처리 시작 후엔 불가). (2) reports 테이블 status enum 에 'withdrawn' 추가 + withdrawn_reason 컬럼. (3) 클라이언트 /me/reports 페이지(imp-0120 으로 신설 시) 각 행에 '신고 철회' 버튼 — 'submitted' 상태 한정. (4) 24시간 골든타임 — 신고 후 24시간 내에는 자유 철회, 이후엔 admin 검토 필요. (5) 잦은 withdraw(같은 reporter 가 7일 내 3회 이상 철회) 는 false report 패턴으로 alignment 페널티 -1."
  effort: small
  impact: medium
  evidence: "코드 backend/cmd/server/main.go L137 readOnly.GET('/me/reports', ...) — DELETE/PATCH 라우트 0건. handlers_report.go L1-63 에 handleWithdrawReport 0건. handlers_admin_audit.go L223 status enum 에 'withdrawn' 0건. web/app/me/reports/* UI 0건(앞 imp-0120 종속)."

- id: imp-0266
  found_at_iter: 19
  area: report
  type: feature
  target: chat_message_inline_report_with_context_capture
  problem: "현재 채팅에서 '신고' 버튼을 누르면 ReportModal 이 targetType='chat_room' 로 호출되는 단순 구조이고(이미 imp-0118 에서 type 불일치 보고), 신고 시점에 신고자가 '어떤 메시지가 문제였는지' 메시지 단위로 핀포인트할 수단이 0건이다. 운영자는 'chat_room 전체' 만 보고 어느 발언이 욕설/사기였는지 추리해야 한다. 메시지 long-press → 컨텍스트 메뉴 '이 메시지 신고' UX 0건(web/app/chats/[id]/page.tsx L78-180 메시지 렌더에 onLongPress/contextmenu 핸들러 0건)."
  proposal: "(1) 메시지 row 에 우상단 ⋯ 아이콘(hover 시 노출) → 메뉴 '메시지 신고/복사/유저 신고/유저 차단'. (2) ReportModal 호출 시 targetType='message', targetId=messageId 전달, 폼 상단에 신고 대상 메시지 카드(보낸이·시각·본문 100자 발췌) 표시 — '이 메시지를 신고합니다' 명확화. (3) 백엔드는 reports 에 message snapshot(body, sender, sent_at)을 함께 저장(수정/삭제 후에도 운영자가 컨텍스트 확보). (4) 모바일은 long-press 350ms → haptic + 컨텍스트 메뉴. (5) ReportModal target preview 영역은 a11y aria-label 'reporting message: ...'."
  effort: medium
  impact: high
  evidence: "코드 web/app/chats/[id]/page.tsx L78-180 메시지 렌더 루프 안에 onContextMenu/onLongPress 0건. 우상단 ⋯ 메뉴 element 0건. L182-187 ReportModal 호출 targetType='chat_room' (imp-0118) — message 단위 0건. handlers_chat.go message-level report endpoint 0건."

- id: imp-0267
  found_at_iter: 19
  area: report
  type: feature
  target: bulk_action_for_repeat_offender_clusters
  problem: "운영자가 같은 user 또는 같은 listing 에 대한 다중 신고를 한 번에 처리할 수단이 0건이다. admin/app/reports/page.tsx 는 신고 1건 단위로 detail panel 을 열고 '조치 실행' 버튼을 눌러야 하므로, 동일 reporter abuser 대상 신고 10건이 들어오면 클릭 10번 + 메모 10번. 결과적으로 cluster(같은 target_id) 처리 일관성도 없고(어떤 건 warn, 어떤 건 suspend), 운영자 손목만 아프다. multi-select + bulk action UI 0건, /admin/reports/bulk-resolve 엔드포인트 0건."
  proposal: "(1) admin reports 표에 row checkbox + '전체 선택' 헤더 추가. (2) bulk action toolbar — '선택된 N건' 표시 후 '동일 사유로 resolve / 동일 조치 적용 / 일괄 reject' 버튼. (3) 백엔드 POST /admin/reports/bulk-resolve {reportIds[], status, actionCode?, memo?} — 트랜잭션 안에서 일괄 처리 + audit_log 일괄 기록. (4) cluster 자동 그룹핑 옵션 — 같은 (targetType, targetId) 신고를 카드 하나로 묶고 'group resolve' 한 번이면 산하 신고 N건 모두 처리. (5) bulk 처리 시 '동일 처분이 정말 적절한가' confirm dialog 로 휴먼 게이트."
  effort: medium
  impact: medium
  evidence: "코드 admin/app/reports/page.tsx L139-145 DataTable 에 row checkbox/selectedRows state 0건. L260 단일 행 onClick=setSelectedReport 패턴만 존재. backend/cmd/server/handlers_admin*.go grep 'bulk\\|BulkResolve' 0건. main.go 에 /admin/reports/bulk* 라우트 0건."

- id: imp-0268
  found_at_iter: 19
  area: report
  type: bug
  target: report_action_no_content_visibility_change
  problem: "handleAdminReportAction 이 ActionCode='temp_hide' 또는 'permanent_hide' 를 받아도 실제 listings/messages 의 visibility 컬럼은 업데이트하지 않는다(handlers_admin.go L86-94: moderation_actions INSERT 와 reports.status='resolved' UPDATE 만 함). 즉 운영자가 '임시 숨김' 조치를 누르면 audit log 에는 'temp_hide' 가 남지만 매물 자체는 그대로 검색/홈에 노출 — 사용자에게는 처리 안 한 거나 마찬가지. ActionCode 와 실제 상태 변경의 분리 불일치."
  proposal: "(1) handleAdminReportAction 에 switch (req.ActionCode) — case 'temp_hide'/'permanent_hide': target_type 분기로 listings UPDATE visibility='hidden' 또는 chat_messages UPDATE deleted_at=NOW() 또는 reviews UPDATE visibility='hidden'. (2) case 'restrict': users UPDATE account_status='restricted' + restriction_scope. (3) case 'suspend': users UPDATE account_status='suspended'. (4) 모든 트랜잭션이 atomically 한 번에 — 부분 실패 방지. (5) 액션 결과 응답에 affectedEntities[] 포함(어떤 row 가 어떻게 바뀌었는지). (6) admin/app/reports 의 select option 라벨/value 와 백엔드 oneof 도 일관성 검증."
  effort: small
  impact: high
  evidence: "코드 backend/cmd/server/handlers_admin.go L86-94 트랜잭션 안: INSERT moderation_actions + UPDATE reports.status — listings/users/chat_messages UPDATE 0건. L68 oneof='warning temp_hide permanent_hide restrict suspend' 5종 받지만 실제 effect 는 alignment.Change 만(L102). admin/app/reports/page.tsx L244-247 select 라벨 'warn_user/restrict_chat/restrict_listing/suspend_account' 와 백엔드 enum 도 불일치."

- id: imp-0269
  found_at_iter: 19
  area: report
  type: a11y_mobile
  target: mobile_long_press_report_haptic_and_focus_trap
  problem: "ReportModal 은 Modal 컴포넌트로 열리지만 (1) 모바일 키보드가 textarea 에 포커스되면 modal 본문이 가려져 submit 버튼이 안 보이고, (2) iOS Safari/Android 에서 modal 외부 body 스크롤이 차단되지 않아 fold 가능, (3) 닫기 버튼 외 영역 탭으로 닫을 때 입력 내용이 경고 없이 손실, (4) 닫힘 시 trigger 버튼으로 focus 가 복귀하지 않아 키보드 사용자가 길을 잃는다. e2e: 모바일 viewport(375x667)에서 ReportModal 열고 textarea focus → 키보드 올라오면 submit 버튼 viewport 밖."
  proposal: "(1) Modal 에 body lock — 'overflow: hidden' 적용 + scrollY 복구. (2) ReportModal max-height: 90vh + body overflow-y: auto, submit 버튼은 sticky bottom 으로 고정해 키보드와 무관하게 항상 노출. (3) 외부 영역 탭(또는 ESC) 시 form dirty 면 confirm dialog '작성 중인 신고가 있습니다. 닫으시겠습니까?'. (4) 닫힐 때 trigger 버튼에 focus 복귀(useRef 패턴). (5) iOS/Android 햅틱(navigator.vibrate(10)) 으로 submit 시 피드백."
  effort: small
  impact: medium
  evidence: "코드 web/components/forms/report-modal.tsx Modal wrapper 사용 — body lock useEffect 0건. submit 버튼 L63 sticky/fixed 0건. onClose 콜백에 dirty 체크 0건. trigger ref 복귀 0건. navigator.vibrate 호출 0건. web/components/ui/modal.tsx 에 max-height/overflow 핸들링 검증 필요(별도 분석)."

- id: imp-0270
  found_at_iter: 19
  area: report
  type: feature
  target: report_status_change_notification_to_target_user
  problem: "신고가 confirmed 되어 alignment -20 와 moderation action(warn/restrict/suspend) 이 가해지는 target user 가 '왜 자기 점수가 떨어졌는지·어떤 조치를 받았는지·어느 신고가 원인이었는지' 알림받는 채널이 0건이다(handlers_admin.go L86-110 에 INSERT INTO notifications 대상 target_user 도 0건, 앞 imp-0123 은 reporter 알림이고 본 건은 target_user 알림으로 별개). 결과: 사용자가 갑자기 채팅 보내기/매물 등록이 거부되어 '버그인가' 로 오해, 항의/이의제기 경로도 모름. 절차 정의의 투명성 부재."
  proposal: "(1) handleAdminReportAction 트랜잭션 안에 INSERT INTO notifications (user_id=req.TargetUserID, type='moderation_action', title='운영팀 조치 안내', body='회원님의 [매물/채팅]에 대한 신고가 검토되어 [경고/임시제한/계정정지] 조치가 적용되었습니다.', deep_link='/me/sanctions/:actionId', reference_type='moderation_action', reference_id=actionID). (2) /me/sanctions 페이지 신설(클라이언트) — 자기 받은 조치 목록 + 사유 + 만료일 + 이의제기 링크. (3) restriction 범위(chat/listing/all) 와 duration_days 명시 + 카운트다운 UI. (4) SSE 실시간 푸시. (5) 이의제기 가이드 — 'evidence 와 함께 giranjt@gmail.com 또는 /me/appeals 로'."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_admin.go L78-113 트랜잭션 안에 INSERT INTO notifications (user_id=TargetUserID) 0건. /me/sanctions 라우트(main.go) 0건. web/app/me/sanctions/* 0건. SSE event 'moderation_action_applied' 카탈로그(internal/event) 정의 0건."

- id: imp-0271
  found_at_iter: 19
  area: report
  type: feature
  target: target_user_alignment_dampening_for_first_offense
  problem: "handleAdminReportAction 이 모든 confirmed 신고에 대해 alignment.Change(target, AlignmentReportConfirmed=-20, ...) 를 일괄 적용한다(handlers_admin.go L102). 이는 (1) 첫 위반자/명백히 의도 없는 실수자(예: 단순 욕설 1회) 와 (2) 5번째 confirmed 신고 받는 상습 어뷰저를 같은 -20 으로 처벌하는 단조성 문제 — 첫 위반자는 grade='caution' 으로 떨어져 추후 회복이 어렵고, 상습범은 -20 으론 부족하다. severity 와 history 에 따른 차등 점수 0건."
  proposal: "(1) action_code 별 alignment delta 차등 — warning -5, temp_hide -10, permanent_hide -20, restrict -25, suspend -40. (2) target user 의 누적 confirmed report 횟수 가중치 — 1st: 1.0x, 2nd: 1.5x, 3rd+: 2.0x(상습 가중). (3) 신고 사유별 가중 — privacy_exposure/prohibited_item 은 기본 +5 추가 페널티(중대 위반). (4) domain/models.go 에 AlignmentDeltaForAction(actionCode, priorOffenses, reportType) 함수 추가, alignment.Change 호출 전에 계산. (5) admin UI 에 '예상 alignment 변동: -10 (warning + 1st offense)' 미리보기 노출."
  effort: medium
  impact: medium
  evidence: "코드 backend/cmd/server/handlers_admin.go L102 'alignment.Change(tx, req.TargetUserID, domain.AlignmentReportConfirmed, ...)' — req.ActionCode/priorOffenses 무시한 단일 -20. backend/internal/domain/models.go L99 AlignmentReportConfirmed = -20 단일 상수, 동적 계산 함수 0건. handleAdminReportAction 에서 SELECT count from reports/moderation_actions WHERE target_user_id=$1 AND status='resolved_confirmed' 0건."

- id: imp-0272
  found_at_iter: 20
  area: notification
  type: feature
  target: cursor_pagination_50_row_hard_cap_history_loss
  problem: "ListNotifications 가 'LIMIT 50' 으로 하드캡(postgres_reservation.go L359), 그 너머 알림은 사용자가 절대 다시 볼 수 없다. 7일 자동 정리(main.go L67) 와 결합되면 더 심각 — 활성 사용자(거래 다수, 채팅 다수) 는 하루 만에 50건을 초과해 어제 받은 '예약 확정' 같은 핵심 알림이 사라진다. cursor 파라미터/load-more 버튼/스크롤 무한 로딩 0건. 'LIMIT 50' 은 client-side 상수도 아닌 SQL 하드코딩이라 운영자 설정도 불가능."
  proposal: "(1) handleListNotifications 에 query param '?cursor=<createdAt|notificationId>&limit=20' 추가. (2) ListNotifications(ctx, userID, cursor *string, limit int) 시그니처 변경, SQL 'WHERE user_id=$1 AND (created_at < $2 OR (created_at = $2 AND id < $3)) ORDER BY created_at DESC, id DESC LIMIT $4'. (3) 응답 'data' 외 'nextCursor' 필드 추가. (4) 프런트 useNotifications → useInfiniteQuery, IntersectionObserver 로 자동 로드. (5) /notifications page 무한 스크롤 + 'X일 이전 알림은 자동으로 정리됩니다' inline 안내. (6) 헤더 벨 dropdown 은 최신 5건만, '모두 보기' 클릭 시 페이지로."
  effort: medium
  impact: medium
  evidence: "코드 backend/internal/repository/postgres_reservation.go L359 'LIMIT 50' 하드코딩, 함수 시그니처(interfaces.go L451) cursor/limit 파라미터 0건. handlers_notification.go L17 'repo.ListNotifications(ctx, userID)' 단일 호출. web/lib/hooks/use-profile.ts L28-35 useQuery(useInfiniteQuery 아님). main.go L67 'INTERVAL 7 days' 정리와 결합 시 활성 사용자 데이터 손실 가속."

- id: imp-0273
  found_at_iter: 20
  area: notification
  type: feature
  target: deduplication_and_rate_limit_per_aggregate
  problem: "imp-0131 이 INSERT 부재를 다루지만 일단 INSERT 가 추가되면 즉시 발생할 새 문제 — 같은 채팅방에서 5초 안에 메시지 5개가 오면 5개의 별도 notifications row 가 쌓이고, 헤더 벨이 5건으로 깜빡이며, 모바일 푸시(imp-0136 후속) 도 5번 울려 사용자를 괴롭힌다. 동일 (user_id, reference_type, reference_id) 에 대해 N분 내 unread 알림이 이미 있으면 새 row 를 만들지 말고 'updated_at + count' 만 증가시키는 디듀프 정책 0건. domain/Notification 에 unread_count 필드도 0건."
  proposal: "(1) DB 마이그레이션 — notifications 에 'unread_count INTEGER NOT NULL DEFAULT 1', 'updated_at TIMESTAMPTZ' 추가 + UNIQUE(user_id, reference_type, reference_id, is_read) WHERE is_read=false partial index. (2) repo.UpsertNotification(ctx, userID, type, refType, refID, deepLink, title, body): 'INSERT ... ON CONFLICT (user_id, reference_type, reference_id) WHERE is_read=false DO UPDATE SET unread_count=unread_count+1, body=EXCLUDED.body, updated_at=NOW()'. (3) 카운트 기반 표시 — '닉네임B (3개의 새 메시지)'. (4) chat 디듀프 윈도 무제한(읽기 전까지), reservation/review 등은 ref 가 변하므로 자연 분리. (5) 사용자가 markRead → 새 알림은 다시 새 row. (6) 필요 시 type 별 N분 cooldown 도 적용 가능."
  effort: medium
  impact: high
  evidence: "코드 backend/db/migrations/001_initial.sql L252-263 notifications 스키마에 unread_count/updated_at 컬럼 0건, partial unique index 0건. domain/models.go L302-313 Notification 구조체에 카운트 필드 0건. UpsertNotification 인터페이스 부재. imp-0131 만 INSERT 부재를 다루며 dedup 은 별개."

- id: imp-0274
  found_at_iter: 20
  area: notification
  type: feature
  target: blocked_user_suppression_in_notifications
  problem: "block_relations 테이블이 존재하고(migrations 001 L52-58) postgres_chat.go L221-232 BlockUser/UnblockUser 가 동작하지만, 차단된 사용자가 (a) 채팅을 보냈을 때 알림이 여전히 차단자에게 도달, (b) 차단자가 등록한 매물에 차단된 사용자가 댓글/문의(향후 기능)할 때도 도달할 우려가 있다. 차단 의미가 '메시지 자체는 막지만 알림은 노출' 이라면 차단의 사용자 효익이 반감 — 진짜 어뷰저는 더미 계정 만들어 다시 보내며, 차단자 인박스에 noise 누적."
  proposal: "(1) repo.UpsertNotification 시 actor_user_id 파라미터 추가 — 'WHERE NOT EXISTS (SELECT 1 FROM block_relations WHERE blocker_user_id=$user_id AND blocked_user_id=$actor_user_id)' guard. (2) 채팅 메시지 INSERT 핸들러(handlers_chat.go) 에서 차단 관계면 INSERT INTO chat_messages 자체를 거부하거나(권장), 또는 메시지는 저장해도 알림만 skip. (3) 매물에 question/inquiry 기능 도입 시 동일 패턴. (4) 차단 해제 시 과거 묻혔던 알림 복구 X(개인정보 노출) — 단순히 차단 시점 이후만 다시 받기. (5) 관리자 broadcast(imp-0276) 는 차단 무관 항상 도달. (6) e2e 테스트 — 'A 가 B 차단 → B 가 메시지 → A 의 GET /notifications 에 0건'."
  effort: small
  impact: medium
  evidence: "코드 backend/internal/repository/postgres_chat.go L221-232 BlockUser 구현. handlers_chat.go SendMessage 핸들러 grep 'block_relations' 0건 — 차단 검사 없이 메시지 저장. 향후 추가될 imp-0131 의 INSERT INTO notifications 도 actor 파라미터 0건이라 차단 필터 자연 누락. backend/internal/repository/interfaces.go BlockedByUser/IsBlocked 헬퍼 0건."

- id: imp-0275
  found_at_iter: 20
  area: notification
  type: feature
  target: marketing_consent_and_korean_privacy_law_compliance
  problem: "한국 정보통신망법 제50조·개인정보보호법 시행령 제17조에 따라 '광고성/마케팅 알림(신규 매물 추천, 이벤트, 프로모션, 가격 인하 알림)' 은 사용자의 명시적 동의(opt-in) 없이 발송 시 3,000만원 이하 과태료 대상이고, 동의했더라도 매 알림 본문에 '수신거부' 1-click 링크를 포함해야 한다. 현재 users 테이블(imp-0131 + 추후 imp-0136 web push) 에 marketing_consent_at/marketing_consent_via 컬럼 0건, 동의 받는 UI 0건, 알림 본문 'unsubscribe' 링크 0건 — 정식 출시 시 즉시 위반."
  proposal: "(1) DB 마이그레이션 — users 에 'marketing_consent_at TIMESTAMPTZ NULL', 'marketing_consent_revoked_at TIMESTAMPTZ NULL', 'marketing_consent_source TEXT NULL'(signup/settings/popup). (2) Google OAuth 콜백 직후 '광고성 알림 수신 동의(선택)' 별도 체크박스 페이지 (transactional 거래 알림은 동의 불필요, marketing 만 분리). (3) /profile/settings/notifications(imp-0134) 에 '광고/이벤트' 토글 추가, 토글 시 marketing_consent_at 갱신. (4) Insert 시점 type='marketing'/'promo'/'price_drop'/'recommendation' 알림은 marketing_consent_at IS NOT NULL AND marketing_consent_revoked_at IS NULL 만 발송. (5) 알림 본문 끝에 '수신거부' deepLink → POST /me/preferences/marketing/revoke 1-click. (6) 매년 동의 갱신 안내(연 1회 재확인)."
  effort: large
  impact: medium
  evidence: "코드 backend/db/migrations/001_initial.sql L7-30 users 테이블 grep 'marketing_consent' 0건. web/app/login/* 'consent\\|동의\\|광고' 0건. handlers_notification.go INSERT(예정) 에 type-specific consent 검사 0건. 한국 정보통신망법 §50 '영리목적 광고성 정보 전송 제한' 미준수 시 과태료 위험."

- id: imp-0276
  found_at_iter: 20
  area: notification
  type: feature
  target: admin_broadcast_system_announcement_endpoint
  problem: "운영자가 '내일 오전 2-4시 점검 예정' '신규 정책 시행' '특정 서버 거래 일시 중단' 같은 공지를 전 사용자에게 보낼 수 있는 채널이 0건이다. handlers_admin.go 에 broadcast/announcement 엔드포인트 0건, /admin 어디에도 'notifications/broadcast' UI 없음. 결과: 운영자가 트위터/디스코드/카페 같은 외부 채널에 공지를 흩뿌리고, 사용자는 인앱에서 '왜 거래가 안 되지' 의문 — 신뢰도·CS 비용 폭증. 점검 시작 직전 갑작스러운 503 도 사전 안내 0건."
  proposal: "(1) admin POST /admin/notifications/broadcast { type: 'maintenance'|'announcement'|'policy', title, body, deepLink?, audienceFilter?: { primaryServer?, role?, joinedBefore? }, scheduledAt?: ISO } 신설. (2) 백엔드 — audienceFilter 매칭 사용자 batch SELECT, INSERT 트랜잭션 chunk 1000건씩, 동시에 SSE 'announcement:new' broadcast. scheduledAt 이 미래면 broadcasts 테이블 enqueue → cron tick 에서 dispatch. (3) /admin/announcements 페이지 — Markdown editor + audience preview('약 12,300명에게 발송') + 발송/예약. (4) 사용자 측 — type='maintenance' 알림은 헤더 상단에 dismissible banner(/notifications 안 가도 보이게). (5) 점검 알림은 marketing consent 무관(시스템 transactional)."
  effort: large
  impact: medium
  evidence: "코드 backend/cmd/server/handlers_admin.go grep 'broadcast\\|announcement' 0건. main.go L194-216 admin route 트리에 '/admin/notifications/*' 0건. broadcasts 테이블 schema 부재. web/components/layout/* 'banner\\|maintenance' 0건. /admin/announcements 라우트 부재."

- id: imp-0277
  found_at_iter: 20
  area: notification
  type: feature
  target: in_app_realtime_toast_for_foreground_users
  problem: "사용자가 /listings/abc 매물 상세를 읽고 있을 때 다른 채팅방에서 새 메시지가 와도, 헤더 벨 빨간 점 1개만 변할 뿐 시각적 피드백이 약하다(움직임 없음, 사운드 없음). 모바일에서는 더더욱 빨간 점이 reach zone 밖이라 거의 인식되지 않음. native push 가 없는 PWA 환경에서 'foreground 사용자에게 즉각적 in-app toast' 는 표준 패턴이지만 ToastContainer 는 mutation success/error 에만 쓰이고 SSE notification:new(imp-0133 도입 후) 와 연결 0건. 결과: 활성 사용자도 알림을 즉시 보지 못함."
  proposal: "(1) SSEContext provider 의 'notification:new' 이벤트 핸들러에서 useToast().addToast('info', `${title} — ${body}`, { actionLabel: '보기', actionHref: deepLink, durationMs: 6000, position: 'bottom-right'(데스크톱)|'top'(모바일) }) 호출. (2) 자기가 현재 보고있는 페이지(pathname) 와 deepLink 가 같으면 toast 생략(이미 보고있는 채팅방의 새 메시지). (3) DND 시간대(imp-0134) 면 toast 도 skip. (4) Toast 컴포넌트에 미디어 쿼리 'prefers-reduced-motion' 시 슬라이드 애니메이션 비활성. (5) 옵션: 사용자 설정에서 '소리'(짧은 비프 또는 시스템 알림음 mp3) on/off — 모바일은 비활성 기본. (6) 다중 toast 큐 (5개 누적 시 그룹화 'X개의 새 알림')."
  effort: small
  impact: medium
  evidence: "코드 web/lib/hooks/use-toast.ts(추정 위치) — useToast 는 onError/onSuccess 에 사용되나 SSE 연동 0건. SSEContext.tsx grep 'addToast' 0건. /chats/* 페이지 외 다른 페이지에서 새 채팅 알림 시각 피드백 0건. window.Notification API(브라우저 native) 호출도 0건. tasks/improvements.md imp-0133 은 cache invalidation 만 다루고 toast UX 별개."

- id: imp-0278
  found_at_iter: 20
  area: notification
  type: feature
  target: snooze_and_remind_me_later_per_notification
  problem: "사용자가 알림을 받았지만 지금 처리할 시간이 없어 '내일 아침에 다시 알려줘' 가 필요한 경우가 빈번(예: 새벽 2시 도착한 거래 제안). 'mark read' 는 영구 처리이고 'unread' 유지는 헤더 벨이 계속 켜져있어 시각적 피로도 높음. 카카오톡/잼/슬랙 같은 표준 메신저는 'snooze' 기능을 제공한다. 현재 /notifications 에 snooze 0건, GET /notifications 에 snoozed_until_at 필터링 0건."
  proposal: "(1) DB 마이그레이션 notifications 에 'snoozed_until_at TIMESTAMPTZ NULL' 컬럼 추가. (2) POST /notifications/:id/snooze { until: '1h'|'8h'|'tomorrow_9am'|'next_week'|customISO } 신설. (3) ListNotifications 쿼리 'WHERE user_id=$1 AND (snoozed_until_at IS NULL OR snoozed_until_at <= NOW())' 필터 — snoozed 알림은 만료 시각 전까지 인박스에서 숨김. (4) /notifications row 에 swipe(모바일)/hover(데스크톱) → '👁 1시간 후 알림' 'Tomorrow 9 AM' 메뉴. (5) cron tick(매 5분) 에서 'WHERE snoozed_until_at < NOW() AND ... ' 매칭 row 들에 SSE notification:resurfaced 푸시. (6) 알림 페이지 별도 'Snoozed' 섹션 — 다가올 시각 미리보기."
  effort: medium
  impact: low
  evidence: "코드 backend/db/migrations/001_initial.sql L252-263 notifications 스키마에 snoozed_until_at 0건. handlers_notification.go grep 'snooze' 0건, 라우트 부재. web/app/notifications/page.tsx swipe 또는 우클릭 메뉴 0건. /notifications 페이지에 단순 mark-read 외 처리 옵션 부재."

- id: imp-0279
  found_at_iter: 20
  area: notification
  type: feature
  target: delivery_open_click_analytics_per_type
  problem: "운영자가 '거래완료 알림은 사용자가 보고 후기 쓸까?' '신고처리 알림 도달 후 사용자가 unblock 했을까?' 같은 효과 측정을 할 데이터 0건이다. notifications 테이블에 opened_at/clicked_at/dismissed_at 컬럼 부재, audit_logs 에도 'notification.delivered' 'notification.opened' 이벤트 0건. 결과: 알림 카피 A/B 테스트 불가, type 별 CTR 비교 불가, 효과 없는 type 을 줄여 noise 줄이는 의사결정 데이터 부재."
  proposal: "(1) notifications 에 'delivered_at TIMESTAMPTZ NOT NULL DEFAULT NOW()' (이미 created_at), 'opened_at TIMESTAMPTZ', 'clicked_at TIMESTAMPTZ' 추가. (2) 사용자가 /notifications 진입 → POST /notifications/seen { ids[] } → opened_at = NOW(). (3) 사용자가 row 클릭(deepLink 이동) → POST /notifications/:id/click → clicked_at = NOW(). (4) 관리자 대시보드 — type 별 CTR(clicked/delivered), 평균 click delay(시간), per-day 추이. (5) A/B 테스트 — variant_key 컬럼 추가, INSERT 시 무작위 A/B 카피 선택 후 기록, 대시보드에 비교. (6) variant_key 도입 후 카피 개선 의사결정 자동화 가능. (7) 7일 정리(main.go L67) 도 cleaned_at 로 logical delete 후 90일 통계용 보관(별도 정리 정책)."
  effort: medium
  impact: low
  evidence: "코드 backend/db/migrations/001_initial.sql L252-263 notifications 컬럼에 opened_at/clicked_at/variant_key 0건. handlers_notification.go grep 'opened\\|clicked\\|seen' 0건. /admin/notifications/stats 라우트 부재. audit_logs L270 grep 'notification' 0건 — 알림 관련 감사 이벤트 미정의."

- id: imp-0280
  found_at_iter: 21
  area: profile
  type: feature
  target: alignment_grade_score_visualization
  problem: "backend domain/models.go L196-197 의 UserProfile 에 'alignmentScore int' 와 'alignmentGrade AlignmentGrade' 필드가 정의되어 있고 GET /me 응답에도 직렬화될 가능성이 높지만(json:\"alignmentScore\"/\"alignmentGrade\"), web/lib/types.ts L1-15 User 타입에는 두 필드가 0건이고 web/app/profile/page.tsx 의 stat 카드 3개(거래/좋은 리뷰/신뢰등급)에도 표시 0건이다. 리니지 클래식의 핵심 신뢰 신호인 '성향(law/chaos/neutral)' 등급이 계산은 되지만 UI 노출 0건 — 사용자가 'X 라는 사람의 성향 점수' 를 한눈에 알 방법이 없고, alignment_score 가 변했어도 본인조차 인지 불가."
  proposal: "(1) web/lib/types.ts User interface 에 'alignmentScore?: number' 'alignmentGrade?: \"law\"|\"chaos\"|\"neutral\"' 추가. (2) /profile stat 카드를 3개에서 4개(또는 carousel)로 늘려 '성향' 카드 추가 — '⚖️ Law +120' / '☠️ Chaos -340' / '⚪ Neutral 0', 색상 차등(gold/red/cream). (3) public profile(/profile/[userId], imp-0138 후속)에도 동일 표시 — 거래자 신뢰 판단에 활용. (4) lib/i18n/alignment.ts: 'law'→'정의', 'chaos'→'무법', 'neutral'→'중립' 한국어 매핑. (5) 점수 변동 트리거(거래완료/리뷰/신고) 발생 시 SSE 'alignment:changed' 이벤트로 toast '⚖️ 성향 점수 +5' 즉시 피드백. (6) /profile 카드 클릭 시 modal — '성향 점수 산정 기준: 거래완료 +10, 긍정 리뷰 +5, 신고 누적 -50, 30일 비활동 ±0' 설명."
  effort: medium
  impact: high
  evidence: "코드 backend/internal/domain/models.go L196-197 'AlignmentScore int' 'AlignmentGrade AlignmentGrade' 필드 정의 + json 태그 직렬화 활성. web/lib/types.ts L1-15 User interface alignment 필드 0건. web/app/profile/page.tsx L86-109 stat 카드 3개에 alignment 0건. backend/internal/alignment/ 디렉토리 존재(CLAUDE.md L40)이나 UI 노출 진입점 0건."

- id: imp-0281
  found_at_iter: 21
  area: profile
  type: feature
  target: response_badge_display_and_hours_chart
  problem: "web/lib/types.ts L11 User 에 'responseBadge?: string' 필드가 있고 backend domain/models.go L194 'ResponseBadge string' 으로 노출되지만, /profile/page.tsx 어디에도 표시 0건이고 매물 카드 작성자 hover preview 에도 미표시(추정). 응답 속도/비율은 거래 결정의 핵심 신호('5분 내 답장' vs '하루에 한 번') 인데 backend 가 계산해도 UI 가 사용 안 함. 결과: 거래자가 '이 판매자에게 채팅 보내면 답장 올까?' 판단 못 하고, 응답률 좋은 사용자도 보상받지 못함."
  proposal: "(1) /profile stat 카드 영역에 responseBadge chip — '⚡ 5분내 답장' / '🐢 하루 이상' (badge 값 한국어 매핑 lib/i18n/response-badge.ts). (2) public profile(/profile/[userId])에도 동일 표시. (3) 매물 카드(/listings) 의 author 영역에 작은 chip — '⚡' 아이콘 단독(공간 절약). (4) 응답 시간 분포 mini chart(최근 30일) — sparkline '평균 12분, 최장 4시간'. (5) backend 가 채팅 메시지 발신 시 '나(seller)→상대(buyer) 첫 답장 지연 시간' 을 collect 하여 user_response_stats(user_id, p50_minutes, p90_minutes, sample_count) 집계 cron(매 시간). (6) 본인 프로필에는 '응답 빠르게 유지하면 거래 성공률 +30%' 동기 부여 카피."
  effort: medium
  impact: medium
  evidence: "코드 backend/internal/domain/models.go L194 'ResponseBadge string' 정의. web/lib/types.ts L11 'responseBadge?: string'. web/app/profile/page.tsx grep 'responseBadge' 0건 — 표시 부재. web/components/listing/ 카드 컴포넌트 grep 'responseBadge' 추정 0건. user_response_stats 같은 집계 테이블 backend/db/migrations 부재."

- id: imp-0282
  found_at_iter: 21
  area: profile
  type: feature
  target: last_active_at_presence_indicator
  problem: "web/lib/types.ts L13 User.lastActiveAt: string 와 backend UserProfile.LastActiveAt 모두 있지만, /profile 본인 화면·public profile·매물 카드·채팅 헤더 어디에도 '최근 접속' 표시 0건이다. 거래자가 '이 사람 1주일 전 마지막 접속 — 답장 기대 안 됨' vs '5분 전 활동 — 지금 답장 가능' 같은 즉각 판단을 못 한다. presence indicator(녹/노/회색 dot)도 부재 — 디스코드/카톡 같은 표준 메신저 패턴 미적용."
  proposal: "(1) avatar 우하단에 색 dot — 5분 이내: green(online), 1시간 이내: yellow(idle), 24시간 이내: gray(recent), 그 이상: hidden. (2) /profile 본인 화면 닉네임 옆 'lastActiveAt' formatTimeAgo — '방금 접속'. (3) public profile(/profile/[userId])과 매물 카드 author 영역에 dot 노출. (4) 채팅방 헤더에 상대방 dot — 'online: 답장 빠를 거예요'. (5) backend lastActiveAt 갱신 정책 — 모든 인증 요청 미들웨어에서 UPDATE users SET last_active_at = NOW() WHERE id = $1 (rate-limit: 1분당 최대 1회 갱신, redis lock 또는 LAST_UPDATED_AT 비교). (6) privacy 옵션(imp-0146 후속)에 'last active 비공개' 토글 — 비활성 시 dot/시간 모두 hide."
  effort: medium
  impact: medium
  evidence: "코드 web/lib/types.ts L13 lastActiveAt 타입 존재, web/app/profile/page.tsx 'lastActiveAt' grep 0건. components/listing/ author hover preview 에 lastActiveAt 표시 0건(추정). 채팅 헤더 web/app/chats/[chatId]/ presence dot 마크업 0건. backend middleware/auth.go 에 last_active_at UPDATE 코드 0건(추정)."

- id: imp-0283
  found_at_iter: 21
  area: profile
  type: ux
  target: profile_completeness_nudge_and_progress
  problem: "신규 가입 사용자가 /profile 진입 시 '닉네임/아바타/소개/주서버' 4 필드 중 일부만 채운 상태에서도 마치 '완료된 프로필' 처럼 보인다. 현재 코드(L77-81)는 introduction 이 있을 때만 '<p>' 표시 — 없으면 그냥 비어있고 '왜 비어있나' '채우면 뭐가 좋나' 단서 0건. 매물 등록·채팅 시작 직전에 '아바타 없으면 신뢰 -10%' 같은 통계도 0건이라 사용자가 채울 동기가 약하다. 결과: 50%+ 사용자가 default state 로 거래에 진입해 신뢰 신호 빈약."
  proposal: "(1) /profile 카드 상단에 '프로필 완성도 75% — 1단계 남음' progress bar (avatarUrl/introduction/primaryServerId 채워졌는지 client-side 계산, nickname 은 가입시 필수). (2) 미완성 필드별 chip — '📷 사진 추가' '✍️ 소개 작성' '🌍 주 서버 선택' (클릭 → /profile/edit#field). (3) 100% 완성 시 '🏆 프로필 완성 — 신뢰 거래자!' 1회 toast 축하 + alignmentScore +5 보상(이펙트). (4) 가입 30일 이내 사용자에게는 헤더 banner '거래 첫 5건은 프로필 완성도 50% 이상이면 평균 30% 더 빠른 매칭' 안내. (5) 닫기 버튼 — 1회 dismiss 후 7일간 다시 표시 안 함(localStorage 'completeness-banner-dismissed-at')."
  effort: small
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L77-81 introduction 없으면 단순 hide — 채우라는 prompt 0건. progress bar UI 0건(grep 'completeness\\|progress' 0건). 가입일 기반 분기(createdAt) 가공 0건. localStorage 'completeness' 키 0건."

- id: imp-0284
  found_at_iter: 21
  area: profile
  type: feature
  target: achievement_badges_milestones
  problem: "기란JT 의 신뢰 모델은 '거래 횟수' 와 '긍정 리뷰' 단순 카운트만 보여주고(/profile L86-98) 점진적 마일스톤 달성을 시각화하지 않는다. 거래 1번/10번/100번/1000번, 긍정 리뷰 90% 유지, 첫 채팅 응답 5분 내 등 게임화 가능한 마일스톤을 메달로 발급하면 retention·engagement 동기가 생긴다. 현재 trustBadge enum 단일 값만 있고 다중 메달/배지 시스템 0건이다. 리니지 클래식 도메인은 RPG-기반 사용자가 '수집 욕구' 에 강하게 반응하는 시장임에도 미활용."
  proposal: "(1) DB 마이그레이션 user_achievements(user_id, achievement_key, earned_at, progress_value) — first_trade/10_trades/100_trades/1000_trades/positive_streak_5/positive_streak_30/responder_5min/early_bird(가입 100명)/community_helper(신고 처리 협조)/anniversary_1y. (2) backend cron(매일 새벽) 또는 이벤트 트리거(거래 완료 시) 에서 조건 평가 → 신규 row INSERT + SSE 'achievement:earned' 이벤트. (3) /profile 본인 화면 메뉴에 '내 메달' 추가 — grid 6×N, 잠긴 것은 회색·진행도 표시('27/100 거래'). (4) public profile 에도 'highlighted 3개' 표시(본인이 선택). (5) 첫 메달 획득 시 modal 축하 + share to discord 동선. (6) 알림 카탈로그에 'achievement.earned' 추가."
  effort: large
  impact: medium
  evidence: "코드 backend/db/migrations/ 'user_achievements' 또는 'badges' 테이블 0건(추정). domain/models.go grep 'Achievement\\|Badge' enum/struct 0건(trustBadge 는 단일 enum). web/app/profile/ 메달/그리드 마크업 0건. SSE 이벤트 카탈로그 docs/EVENT_CATALOG.md grep 'achievement' 추정 0건."

- id: imp-0285
  found_at_iter: 21
  area: profile
  type: ux
  target: joined_date_seniority_label
  problem: "User 타입에 'createdAt: string'(L14)이 노출되지만 /profile 본인 카드·public profile 모두 표시 0건이다. 가입일은 거래 신뢰의 1차 신호(오래된 사용자=낮은 사기 확률)인데 시각화 0건이다. 또한 '회원기간' 동안 milestone 도 표시 안 됨 — 1주년/2주년/100일 같은 의미있는 시점이 사라진다. createdAt 만 있고 fingerprint 없는 사용자는 그저 'UUID' 일 뿐 '커뮤니티 일원' 으로 인식되지 않는다."
  proposal: "(1) /profile 닉네임 옆 회원기간 chip — '🏛 6개월차' / '🌟 1년 회원' (createdAt 으로부터 monthDiff 계산). (2) 1주년·6개월·100일 도달 시 자동 toast 축하('가입 6개월! ⚡') + 'shareto kakao' 동선. (3) 가입일 기반 vintage tier — 'Founder'(가입 1000명 이내, 영구 표시), 'Early'(첫 1만명), 'Member'(나머지). FE 에서 createdAt + tier-cutoff 로 자체 산출. (4) public profile 도 동일 chip — '6개월차 회원' 보고 거래자가 신뢰 판단. (5) 가입일 정확 날짜는 mouse-hover 시 tooltip(2025년 3월 14일)."
  effort: small
  impact: low
  evidence: "코드 web/lib/types.ts L14 'createdAt: string'. web/app/profile/page.tsx grep 'createdAt' 0건 — 표시 부재. monthDiff/vintage tier 산출 함수 0건(lib/utils.ts grep). 회원기간 chip 마크업 부재."

- id: imp-0286
  found_at_iter: 21
  area: profile
  type: feature
  target: favorites_wishlist_quick_access
  problem: "사용자가 마음에 드는 매물을 찜(즐겨찾기)할 가능성이 있으나 — 그래서 매물 상세에 ❤️ 또는 '찜' 버튼이 있다면 — /profile 메뉴 4개(L47-52) 중 '찜한 매물' 진입점 0건이다. 결과: 사용자가 어제 본 매물을 다시 찾을 동선 부재, 가격 떨어졌을 때 알림 받을 동선 부재(price-drop notification 의 source-of-truth 부재). 의도적 retention 메커니즘인 wishlist 가 비활성."
  proposal: "(1) /profile menuItems 에 '찜한 매물' 추가 (L47-52). (2) /profile/favorites 라우트 신설 — favorites 테이블(user_id, listing_id, created_at) 조회 grid. (3) backend POST/DELETE /api/v1/me/favorites/:listingId, GET /api/v1/me/favorites. (4) 매물 상세에 ❤️ toggle 버튼(이미 있다면 활용). (5) price drop notification — 찜한 매물의 priceAmount 가 낮아지면 자동 알림 push(imp-0136 의 web push 채널 활용). (6) /profile/favorites 빈 상태: '찜한 매물이 없습니다 — 매물 상세에서 ❤️ 를 눌러 저장하세요'. (7) 카운트 badge — '내 매물(N) 찜(M)' menu item label."
  effort: medium
  impact: medium
  evidence: "코드 web/app/profile/page.tsx L47-52 menuItems 4개 — '찜' 0건. backend grep 'favorites\\|wishlist' table/route 추정 0건. web/app/listings/[id]/page.tsx 의 ❤️ 버튼 존재 추정 — 그러나 profile 진입점 부재로 사용자 동선 끊김. price drop SSE 이벤트 EVENT_CATALOG.md grep 0건."

- id: imp-0287
  found_at_iter: 21
  area: profile
  type: feature
  target: contact_links_discord_kakao_external
  problem: "리니지 클래식 거래는 in-app 채팅 외에 디스코드/카카오 오픈채팅으로 연속되는 경우가 많다. 거래자가 '디스코드 ID 알려줘' 를 채팅 안에서 plain text 로 주고받는데 — 이는 (a) 사기범이 외부로 빼돌리는 핵심 동선이고 (b) 본인이 명시 등록하면 신뢰 신호가 된다. 현재 /profile/edit 에 'discord/kakao/twitter/youtube' 링크 등록 필드 0건. UserProfile struct 에도 social_links 필드 0건. 결과: 신뢰 거래자가 자기 외부 신원을 검증 가능하게 노출할 수단 부재."
  proposal: "(1) DB 마이그레이션 users 에 'social_links JSONB' 컬럼 — { discord, kakao_id, twitter, youtube } 4종 white-list. (2) /profile/edit 에 '연결된 계정' 섹션 — 4 input + 'discord 연동 인증' optional(OAuth 추가 단계). (3) public profile 에 chip — 'Discord: hyungjoon#1234' (verified 시 ✅, 미인증 시 ⚠️). (4) backend 가 외부 platform 으로 verify 가능한 type(discord OAuth)만 ✅ 부여 — kakao/twitter 는 self-claim 으로 ⚠️. (5) 사기 신고 시 discord_id 도 함께 admin 에 노출되어 cross-platform 추적 가능. (6) 사용자가 본인 채팅 안에서 'discord:' 텍스트 감지 시 '프로필에 등록하면 더 신뢰받습니다' 가벼운 인라인 nudge."
  effort: medium
  impact: medium
  evidence: "코드 backend/internal/domain/models.go L186-199 UserProfile struct grep 'discord\\|kakao\\|social' 0건. /profile/edit/page.tsx L121-173 4 input(nickname/intro/avatar/server) 만 — social link 입력 0건. social_links JSONB 컬럼 backend/db/migrations/ 0건."

- id: imp-0288
  found_at_iter: 21
  area: profile
  type: feature
  target: trade_history_chart_recent_30d
  problem: "/profile 의 stat 카드는 누적 카운트(거래 N건, 좋은 리뷰 M건)만 보여주고, 시간축 데이터(최근 30일/주별 추세)가 0건이다. 본인이 '이번 달 거래량이 줄었나' 자가 인지 불가, 거래 상대방 입장에서도 '최근 활성 거래자' vs '6개월 전엔 활발했지만 지금은 휴면' 구분 불가. 신뢰 평가는 누적 + 최근 활동의 결합인데 후자가 시각화 안 됨."
  proposal: "(1) /profile 본인 화면에 mini sparkline(width 280, height 60) — 최근 30일 일별 거래 완료 건수. recharts/Visx 가벼운 라이브러리. (2) public profile 에도 동일(privacy 옵션으로 hide 가능, imp-0146 후속). (3) backend GET /api/v1/users/:userId/stats?range=30d → daily_trade_count[], daily_review_count[]. (4) hover 시 tooltip — '5월 12일 거래 3건'. (5) 0 데이터 일자도 0 으로 표시(연속성 보장). (6) sparkline 위 'Trend ↑ 23%' 또는 'Trend ↓ 5%' (전월 대비 % 변화) — 본인 motivation 또는 거래자 신호. (7) 6개월 휴면(거래 0) 사용자에게 '최근 활동이 적어요' 회색 dim 처리."
  effort: medium
  impact: low
  evidence: "코드 web/app/profile/page.tsx 'sparkline\\|chart\\|recharts' grep 0건. backend handlers_user.go grep '/stats' '/trade-history' 라우트 0건. domain/models.go UserProfile 에 시계열 데이터 필드 0건."

- id: imp-0289
  found_at_iter: 21
  area: profile
  type: ux
  target: theme_density_preference_per_user
  problem: "기란JT 는 단일 다크 테마(AppColors.gold/blue 변수, CLAUDE.md L13-14)만 제공하여 (a) 라이트 환경(직장/공공장소)에서 어색, (b) 시각 피로 사용자가 라이트로 전환 불가, (c) 정보 밀도 선호(저시력 사용자는 큰 폰트, 파워 거래자는 콤팩트) 모두 단일 강제. /profile/settings(imp-0143 후속) 도 알림 토글만 다루지 시각 환경 옵션 0건이다. 시스템 prefers-color-scheme 도 따르지 않아(추정) 다크 강제."
  proposal: "(1) /profile/settings 에 '화면 설정' 섹션 — '테마: 다크/라이트/시스템 따름', '폰트 크기: 작게/보통/크게', '정보 밀도: 콤팩트/보통/넉넉'. (2) localStorage 'theme-pref' / 'font-size-pref' / 'density-pref' 저장 + Tailwind data-attr 토글 (data-theme=light/dark, data-density=compact). (3) tailwind.config.js 다크/라이트 변수 페어 추가 — Cream/Taupe 는 라이트에서도 동일, Background 만 light=#F5F2EE 추가. (4) 시스템 따름 시 prefers-color-scheme 미디어 쿼리 monitor → 자동 전환. (5) 첫 진입 시 시스템 prefs 자동 감지 prompt — '시스템이 라이트 모드입니다 — 라이트 테마를 사용할까요?' 1회. (6) 모든 컴포넌트가 Tailwind class 기반이라 작업량 작음(직접 hex 사용 0건 가정)."
  effort: large
  impact: low
  evidence: "코드 web/app/layout.tsx 또는 web/app/globals.css 에 'data-theme\\|prefers-color-scheme' grep 0건(추정). /profile/settings 라우트 부재(imp-0143 미구현). lib/theme/ 디렉토리 부재. tailwind.config.js 다크/라이트 페어 컬러 0건(단일 팔레트)."

- id: imp-0290
  found_at_iter: 22
  area: auth
  type: feature
  target: login_throttle_brute_force_lockout
  problem: "코드 backend/cmd/server/handlers_auth.go L33-72 handleLogin 은 IP/사용자별 시도 횟수 제한 0건 — 같은 IP 가 1초당 100번 다른 Google 토큰을 보내도 throttle 0건이다. 사기범이 탈취한 Google ID 토큰 후보 N개를 빠른 brute로 시도하거나, 단순 DDoS 로 /auth/login 을 마비시킬 수 있다. backend grep 'rate.?limit\\|throttle\\|limiter' → 0건(미들웨어 없음). 무인 자동화 차단 부재."
  proposal: "(1) backend/internal/middleware/ratelimit.go 신규 — token bucket 또는 sliding window(github.com/ulule/limiter). (2) /auth/login 에 IP 당 30 req/min, /auth/refresh 에 IP 당 60 req/min 적용. (3) 로그인 실패 5회 누적 시 해당 IP 를 15분 cool-down(409 Too Many Requests). (4) 로그인 성공한 정상 IP 는 화이트리스트 cache(5분), throttle 우회. (5) 응답 헤더 X-RateLimit-Remaining + X-RateLimit-Reset 으로 클라이언트 인지. (6) Redis 가 있다면 분산 카운터, 없으면 in-memory(단일 인스턴스 OK). (7) 정지/탈퇴 계정 관련 logging 도 함께 — 의심 패턴 추적."
  effort: medium
  impact: high
  evidence: "코드 backend/cmd/server/handlers_auth.go L33-72 handleLogin 입구에 limiter 0건. backend/internal/middleware/ ls — auth.go 만(ratelimit.go 부재). go.mod grep 'limiter\\|ratelimit' 0건. /auth/login + /auth/refresh + /auth/logout 모두 무제한 호출 가능."

- id: imp-0291
  found_at_iter: 22
  area: auth
  type: feature
  target: new_device_country_login_alert
  problem: "코드 handlers_auth.go L98-104 기존 사용자 로그인 시 UpdateLastLogin 만 호출 — 이전 로그인 IP/UA 와 비교 0건. 사용자가 한국에서만 쓰던 계정인데 갑자기 미국/중국 IP 에서 로그인되면 — 정상 사용자에게 알림 0건이라 계정 탈취 인지가 늦어진다. Google 자체 알림은 있을 수 있으나 기란JT 의 거래 컨텍스트(매물 N개, 진행중 거래 M건) 손실 위험을 사용자에게 직접 통지 못한다."
  proposal: "(1) handleLogin 에서 c.ClientIP() + UserAgent → 직전 N(=10)회 로그인 이력과 비교(refresh_tokens 테이블에 컬럼 추가, imp-0160 와 결합). (2) 새 country code(MaxMind GeoLite 무료 DB 또는 Cloudflare CF-IPCountry header) 또는 새 device_label 감지 시 — 알림 1건 자동 생성: '🚨 새 디바이스에서 로그인 — 5월 2일 19:35, Chrome on Windows, 서울'. (3) '내가 아니라면 즉시 모든 디바이스 로그아웃' 액션 버튼 → /me/sessions DELETE-all. (4) 알림 채널 imp-0168 의 email 결합 시 critical 통지 가능. (5) 임계 너무 민감하면 거짓양성 — '한국 4G + Wi-Fi 스위칭' 정도는 같은 country 로 묶어 알림 안 보냄."
  effort: large
  impact: high
  evidence: "코드 handlers_auth.go L98-104 existing-user 분기에 새 디바이스 감지 로직 0건. backend grep 'GeoIP\\|country\\|MaxMind\\|CF-IPCountry' → 0건. last_login_ip 컬럼 0건(migrations grep). 알림 type 카탈로그(EVENT_CATALOG.md)에 'new_device_login' 0건."

- id: imp-0292
  found_at_iter: 22
  area: auth
  type: feature
  target: oauth_scope_minimization_no_email_default
  problem: "현재 Google OAuth 통합은 사실상 sub/email/name 모두 받는다 — oauth/google.go L11-14 GoogleTokenInfo 가 Email/Name 필드를 가지나 handlers_auth.go L92 CreateUserWithProfile 은 nickname='유저_' + userID[:8] 만 사용해 email/name 을 그냥 버림(데이터 흘러가도 사용 0건). 그러나 ID token 은 'openid email profile' scope 로 발급되어 — 서버는 받지만 안 쓰는 PII 보유 = 개인정보 최소화 원칙 위반. 추후 leak 시 책임 확대."
  proposal: "(1) login/page.tsx L60-72 initialize 에 'scope' 명시 가능하면 'openid' 만(현재 누락 → 자동 'openid email profile'). Google Sign-In V2 는 scope override 제한이 있어 — 그렇다면 (2) 받은 email/name 즉시 폐기 — handlers_auth.go 에 변수 declare 만 하고 unused 로 처리하지 말고 oauth/google.go L31-32 부터 email/name parse 자체 제거(코드에서 read 안 하도록). (3) 만약 imp-0168 (계정복구 채널)을 위해 email 이 필요하면 — 회원 탈퇴 시 즉시 NULL 처리 가시화 + sha256 해시화로 저장(원문 X). (4) Google 'sub' 만으로도 unique 이므로 충분. (5) Privacy 페이지에 '수집 항목 — Google sub(공개 식별자) 1건' 명시."
  effort: small
  impact: medium
  evidence: "코드 oauth/google.go L13-14 Email + Name 필드 정의. handlers_auth.go L86-97 CreateUserWithProfile signature 에 email/name 미사용 — 받기는 하나 사용 0건. login/page.tsx L60-72 initialize 에 scope 명시 0건(default 'openid email profile'). DB 에 email 저장 0건이지만 서버 메모리는 한 번 통과."

- id: imp-0293
  found_at_iter: 22
  area: auth
  type: feature
  target: passkey_webauthn_support
  problem: "기란JT 는 Google OAuth 단일 인증으로 — 사용자가 Google 계정에 lock-out 되면(Google 정지, 비밀번호 분실, 2단계 인증 분실) 기란JT 진입 0건이다. 또한 Google OAuth 가 외부 종속이라 Google 장애 시 전 사용자 로그인 불가(SPOF). passkey/WebAuthn 은 OS 통합으로 외부 의존성 0건이며, 한 번 등록 후 비밀번호 없이 즉시 로그인 가능 — 모바일 사용자(주 타겟)에게 마찰 최소."
  proposal: "(1) backend/internal/oauth/passkey.go 신규 — github.com/go-webauthn/webauthn 라이브러리. (2) 마이그레이션 — passkeys 테이블(id, user_id, credential_id BYTEA, public_key BYTEA, sign_count BIGINT, transports TEXT[], created_at, last_used_at). (3) /auth/passkey/register/options + /verify, /auth/passkey/login/options + /verify 라우트. (4) /profile/security 에 'Passkey 등록' 버튼 — 'iPhone Face ID/Touch ID, Mac Touch ID, Android 지문' 등으로 등록. (5) /login 에 'Passkey 로 로그인' 버튼(이미 등록된 사용자만 표시 — autofill UI). (6) 신뢰 지표 — Google + Passkey 다중 등록 시 trustBadge '2단계 검증' 부여. (7) 후순위로 Passkey-only 가입(Google 의존 없는 신규 가입 동선)."
  effort: large
  impact: high
  evidence: "코드 backend/internal/oauth/ ls 결과 google.go 만(passkey.go 부재). go.mod grep 'webauthn' → 0건. login/page.tsx 의 fallback 인증 0건. /profile/security 라우트 부재(imp-0160 와 결합 시 같은 페이지에 device + passkey 묶기 자연스러움)."

- id: imp-0294
  found_at_iter: 22
  area: auth
  type: feature
  target: magic_link_email_fallback
  problem: "Google OAuth 가 작동 안 할 때(Google 사이트 차단된 회사망, Google 계정 잠긴 사용자, 부계정 보유자) 기란JT 진입 동선 0건이다. SMS OTP 는 비싸고(Twilio 비용), Magic Link(이메일로 1회용 로그인 링크)는 SMTP 만 있으면 무료 — Google 의존 없는 lightweight fallback. backend 에 SMTP/sendmail 설정 0건(go.mod grep 0)."
  proposal: "(1) backend/internal/auth/magic_link.go — github.com/wneessen/go-mail 또는 net/smtp + Mailgun 무료 등급. (2) POST /auth/magic-link/request — body { email } → magic_links 테이블(id, email, token_hash, expires_at=NOW+15m, used_at) → email 발송 'https://giranjt.com/auth/magic?token=xxx'. (3) GET /auth/magic?token=xxx → server-side verify → 기존 user(email 매칭) 면 login, 없으면 신규(email 만 가진 user 생성) → set cookie + redirect. (4) login/page.tsx — Google 버튼 아래 'OR' divider + email input + '메일로 로그인 링크 받기' 버튼. (5) imp-0168(email 채널) + imp-0292(scope minimization) 와 묶어 데이터 모델 일관성. (6) Rate limit imp-0290 결합 — IP/email 당 분당 3회. (7) 보안 — token 1회용, 발송 후 5분 내 사용 권고, 다른 디바이스 사용 가능(편의)."
  effort: large
  impact: medium
  evidence: "코드 backend/internal/auth/ 디렉토리 부재(grep 'magic_link\\|MagicLink\\|smtp\\|sendmail' 0건). users 테이블에 email 컬럼 0건(imp-0168 dependency). login/page.tsx 의 alternative auth 0건."

- id: imp-0295
  found_at_iter: 22
  area: auth
  type: ux
  target: anonymous_session_preview_before_signup
  problem: "익명 사용자가 '로그인 없이 둘러보기' 클릭 후 매물 / 채팅 / 좋아요 등을 둘러봐도 — 그동안의 행동 기록(찜한 매물, 검색 키워드, 본 매물 list)이 가입 시 0건 carry-over 된다. 사용자가 '아 이거 가입하면 더 좋겠다' 결심해 가입했더니 처음부터 다시 시작 — 동기 손실. localStorage/cookie 에 anonymous_id 부여 + 가입 시 binding 하면 retention 큰 향상."
  proposal: "(1) web/lib/anonymous-session.ts 신규 — 첫 진입 시 cookie 'anon_id'=uuid, 30일 TTL. (2) 익명 사용자의 favoriteListing/searchHistory/recentlyViewedListings 를 localStorage 에 'anon-session.<anon_id>' 키로 저장. (3) 가입(handleGoogleResponse 직후) → POST /me/migrate-anonymous?anonId=xxx → backend 가 favorites/searches 테이블에 user_id 기입(또는 새 매핑 row 추가). (4) /login 에 작은 hint — '회원가입하시면 둘러본 매물과 찜이 그대로 이어집니다' (anon_id 가 비어있지 않을 때만 표시). (5) Privacy — anonymous 사용자에게도 cookie 동의 alert(GDPR/PIPL) 한 번 표시. (6) 가입 안 하고 30일 지나면 anon_id row purge."
  effort: medium
  impact: high
  evidence: "코드 web/lib/ ls — anonymous-session 부재. apiClient 에 favoriteListing 메서드(L169-175) 있으나 익명 시 사용 불가(JWT 필요). cookie 'anon_id' 발급 0건(grep 'anon_id' web → 0). /me/migrate-anonymous 라우트 0건."

- id: imp-0296
  found_at_iter: 22
  area: auth
  type: ux
  target: just_in_time_signup_at_trade_action
  problem: "익명 사용자가 매물 상세에서 '채팅 시작' 누르면 — useAuthGuard.ts L13-25 가 즉시 router.push('/login?redirect=...') 로 페이지 풀 전환 → Google 버튼 클릭 → 로그인 → redirect 의 경로 복귀. 4단계 풀-페이지 전환은 모바일에서 깊은 마찰. 사용자가 '그냥 채팅하고 싶었는데 가입까지?' 라며 이탈. 더 가벼운 흐름(modal 안에서 1탭 로그인 → 즉시 채팅 시작)이 retention 개선."
  proposal: "(1) components/auth/login-modal.tsx 신규 — Google Sign-In One Tap UI(prompt API)를 modal 안에 띄움. (2) useAuthGuard 의 require Auth 가 router.push 대신 useDialog 로 LoginModal 띄우기 — 본문 컨텍스트 유지. (3) 로그인 성공 시 모달 닫히고 원래 액션(handleStartChat) auto-resume — userId 받자마자 createChat 호출. (4) modal 안에 mini onboarding — '닉네임 즉시 변경' inline input(imp-0162 와 결합, 가입 직후 1탭). (5) modal 닫기('취소')도 자연스러움 — 익명 둘러보기 흐름 회복. (6) 페이지 전환 redirect 가 진짜 필요한 케이스(/profile, /chat 직접 진입) 만 router.push 유지."
  effort: medium
  impact: high
  evidence: "코드 web/lib/hooks/use-auth-guard.ts L13-25 requireAuth 가 router.push 만 — modal 옵션 0건. components/auth/login-modal.tsx 부재. apiClient.login 응답 후 콜백 chain 부재. Google One Tap 'prompt()' API 미사용(login/page.tsx 는 renderButton 만)."

- id: imp-0297
  found_at_iter: 22
  area: auth
  type: feature
  target: google_one_tap_auto_prompt
  problem: "Google One Tap 은 페이지 진입 즉시 화면 우상단에 '<email> 로 계속하시겠습니까?' 자동 prompt 를 띄워 1탭 로그인을 가능하게 한다. 코드 login/page.tsx L65-72 은 renderButton 만 사용 — google.accounts.id.prompt() 호출 0건이라 자동 prompt 0건. /listings, /listings/[id] 같은 일반 페이지에도 prompt 0건이라 — 익명 사용자가 Google 로 한 번 로그인했어도 다음 방문 시 매번 /login 가야 한다."
  proposal: "(1) layout.tsx 또는 components/auth/google-one-tap.tsx — useEffect 에서 window.google.accounts.id.initialize + .prompt({notification}). (2) 익명 사용자만 표시(apiClient.isLoggedIn===false). (3) 뷰포트 1280+ 에서만 표시(모바일은 화면 작아 거슬림). (4) cooldown — '취소' 후 24시간 안 띄우기(localStorage 'one-tap-dismissed-at'). (5) FedCM 지원(2024+ Chrome 권장) — use_fedcm_for_prompt: true. (6) 로그인 화면(/login)에서는 prompt 안 보이게(중복) — pathname 체크. (7) 리스트/매물상세에서 prompt → 1탭 로그인 → 페이지 그대로(채팅 시작 버튼 활성)."
  effort: small
  impact: medium
  evidence: "코드 web/app/login/page.tsx L65-72 google.accounts.id.prompt() 호출 0건. layout.tsx L44 의 google gsi script 는 로드되지만 prompt 발화 0건. components/auth/ 디렉토리 부재. FedCM 'use_fedcm_for_prompt' 옵션 사용 0건."

- id: imp-0298
  found_at_iter: 22
  area: auth
  type: performance
  target: google_client_id_env_var_not_hardcoded
  problem: "코드 web/app/login/page.tsx L9-10 GOOGLE_CLIENT_ID 가 소스에 하드코딩(`'1040191360407-...'`). 이는 production 키가 main 브랜치에 영구 노출되어 있다는 것 — 키 회전(rotation) 시 코드 수정+배포 필요, dev/staging/prod 키 분리 불가, 외부 fork 시 그대로 노출. backend 는 cfg.GoogleClientIDs 환경변수로 받아 분리되어 있는데(handlers_auth_test.go L35) 웹 프론트만 하드코드 — 일관성 깨짐."
  proposal: "(1) web/.env.example 에 NEXT_PUBLIC_GOOGLE_CLIENT_ID=... 추가. (2) login/page.tsx L9-10 → const GOOGLE_CLIENT_ID = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID. (3) 빌드 시 누락 검증 — !GOOGLE_CLIENT_ID 면 build error. (4) docker-compose.yml lincle-web 서비스 environment 에 NEXT_PUBLIC_GOOGLE_CLIENT_ID 전달. (5) 키 회전 절차 문서화 — Google Cloud Console 새 client_id 생성 → docker-compose env 변경 → 무중단 rolling 배포 → Old client_id 비활성. (6) backend GoogleClientIDs 도 ',' 로 N개 받게 되어 있어(config.go) 회전 중 dual-validate 가능 — rolling 으로 안전 회전."
  effort: trivial
  impact: medium
  evidence: "코드 web/app/login/page.tsx L9-10 const GOOGLE_CLIENT_ID = '1040191360407-...' 하드코드. .env grep 'NEXT_PUBLIC_GOOGLE' 0건(추정). backend handlers_auth_test.go L35 는 GoogleClientIDs env 사용 — 백엔드 vs 프론트 불일치. docker-compose.yml grep 'GOOGLE_CLIENT' 결과 backend 만(추정)."

- id: imp-0299
  found_at_iter: 22
  area: auth
  type: performance
  target: csp_header_for_google_gsi
  problem: "deploy/Caddyfile 11-33 라인에 reverse_proxy 설정만 — Content-Security-Policy 헤더 0건. layout.tsx L44 가 https://accounts.google.com/gsi/client 외부 스크립트를 로드하는데 CSP 미지정 → XSS 시 attacker 가 임의 외부 스크립트 inject 가능. CSP 'default-src self; script-src self https://accounts.google.com; frame-src https://accounts.google.com' 같은 strict 정책으로 1차 방어막 부재."
  proposal: "(1) deploy/Caddyfile 의 모든 handle 블록에 header Content-Security-Policy 'default-src self; script-src self https://accounts.google.com https://*.gstatic.com; frame-src https://accounts.google.com; img-src self data: https:; connect-src self https://accounts.google.com; font-src self https://fonts.gstatic.com; style-src self unsafe-inline https://fonts.googleapis.com'. (2) 점진 적용 — 먼저 Content-Security-Policy-Report-Only 로 1주 모니터링, 위반 0건 확인 후 enforce. (3) 다른 보안 헤더 동시 추가 — Strict-Transport-Security 'max-age=31536000', X-Content-Type-Options 'nosniff', X-Frame-Options 'DENY', Referrer-Policy 'same-origin', Permissions-Policy 'camera=(), microphone=(), geolocation=()'. (4) backend handlers 가 'unsafe-inline' style 을 안 쓰도록 (Tailwind 만으로 충분 — 가능). (5) report-uri 로 위반 수집(/api/csp-report endpoint)."
  effort: small
  impact: high
  evidence: "코드 deploy/Caddyfile L11-33 grep 'header\\|CSP\\|X-Frame\\|HSTS' → 0건. layout.tsx L44 외부 script src 로드. backend handlers grep 'Content-Security-Policy' → 0건. 보안 헤더 default 0건."

- id: imp-0300
  found_at_iter: 22
  area: auth
  type: feature
  target: jwt_secret_rotation_kid_support
  problem: "코드 middleware/auth.go L19-31 AuthMiddleware 가 단일 secret 만 보유 — JWT_SECRET 환경변수가 leak 또는 정기 회전 시 모든 활성 access/refresh token 즉시 무효화 → 전 사용자 강제 재로그인. 'kid'(key ID) 헤더 미사용으로 — 새 키와 옛 키를 동시 인정하는 grace period 불가. 무중단 secret 회전 부재."
  proposal: "(1) config 에 JWT_SECRETS=[<current>,<previous>] (CSV) 형태로 N개 받기. (2) GenerateTokens 는 첫 번째(current) 로만 sign, header 에 'kid'='current' 추가. (3) ParseToken 은 kid 보고 해당 secret 으로 verify — 'previous' 도 인정(grace, 7일). (4) 회전 절차 — 새 secret 추가 [new, old] → 1주 후 [new] 로 단축 → 옛 키로 발급된 token 자연 만료(15m TTL). (5) jwt.Parse 의 keyfunc 가 token.Header['kid'] 활용. (6) 의도적 키 무효화(전체 logout) 도 같은 메커니즘 — old kid 즉시 제거. (7) 운영자에게 'JWT 키 회전' 명령 (deploy script + admin endpoint)."
  effort: medium
  impact: medium
  evidence: "코드 backend/internal/middleware/auth.go L19-31 AuthMiddleware struct 에 secret []byte 단일 필드. L40-50, L52-65 generate 시 kid header 0건. L72-83 ParseWithClaims keyfunc 가 단일 a.secret 반환 — 다중 키 지원 0건. config grep 'JWT_SECRETS\\|kid\\|key_id' 0건."

- id: imp-0301
  found_at_iter: 22
  area: auth
  type: feature
  target: jti_revocation_blocklist
  problem: "코드 middleware/auth.go L37-66 access token 은 stateless JWT — 한 번 발급되면 15분간 어떤 무효화도 불가. 사용자가 '내 토큰이 탈취된 것 같다' 신고해도 — 다른 디바이스에서 logout-all 해도 — 탈취된 access token 은 만료 시간(15m)까지 살아있다. 탈취 윈도우 15분이 핵심 거래 시점에 걸리면 큰 손실."
  proposal: "(1) GenerateTokens 에서 jwt.RegisteredClaims 에 ID(jti)= uuid 추가. (2) 마이그레이션 — token_blocklist(jti TEXT PRIMARY KEY, user_id TEXT, blocked_at TIMESTAMPTZ, expires_at TIMESTAMPTZ). expires_at 으로 자동 cleanup. (3) ParseToken 안에서 (또는 RequireAuth middleware 에서) repo.IsBlocked(jti) 체크 — true 면 401. (4) 사용자 로그아웃 모든 디바이스 시 — 모든 활성 jti 를 blocklist 에 push(아니면 access 는 그냥 두고 refresh 만 무효화도 정책 가능). (5) 관리자 'force logout user X' 액션 → 해당 user 의 모든 access jti push. (6) 만료 비교 — Redis 또는 DB 인덱스 조회 빠르게(인메모리 cache 30초). (7) imp-0160 의 device sessions 와 결합 — 특정 디바이스만 블록 가능."
  effort: medium
  impact: medium
  evidence: "코드 middleware/auth.go L37-66 jwt.RegisteredClaims 에 ID 필드 미사용(grep 'jti\\|ID:' L37-66 → 0). ParseToken keyfunc 외에 추가 검증 0건(L72-83). token_blocklist 테이블 마이그레이션 부재(grep 'blocklist\\|revoked_token' migrations 0). DeleteRefreshTokensByUser 는 refresh 만 처리, access 는 그대로 살아있음."

- id: imp-0302
  found_at_iter: 22
  area: auth
  type: feature
  target: account_linking_multiple_providers
  problem: "코드 handlers_auth.go L78-97 가 (provider, providerKey) 쌍으로 user 를 lookup — 사용자가 Google 로 가입한 뒤 imp-0157 의 카카오 OAuth 로 로그인하면 — 카카오 sub 으로 새 user row 가 생성되어 두 개의 프로필을 가지게 된다(거래 이력/리뷰 분리). '같은 사람의 다중 OAuth' 통합 동선 부재 — provider 가 1개면 lock-in, N개면 분리 인지."
  proposal: "(1) 마이그레이션 — user_oauth_identities 테이블(user_id, provider, provider_user_key, linked_at) 분리, users 테이블의 (login_provider, login_provider_user_key) 는 deprecate. 기존 row 는 마이그레이션 시 1개 row 로 자동 변환. (2) FindUserByProvider → JOIN user_oauth_identities. (3) /profile/security 에 '연결된 계정' 섹션 — Google ✓ / 카카오 ✗ + '카카오 연결' 버튼 → OAuth flow 후 user_oauth_identities 새 row insert. (4) '연결 해제' — 마지막 1개 provider 는 해제 차단(lock-out 방지). (5) 다른 OAuth 가 같은 email 가지면 — '이미 가입된 계정 발견 — 합치시겠습니까?' modal(우선 잠금 후 사용자 confirm). (6) 합치기 시 거래/리뷰/매물 모두 한 user_id 로 통합."
  effort: large
  impact: medium
  evidence: "코드 handlers_auth.go L78-97 FindUserByProvider 가 (provider, key) 단일 lookup — 다중 identity 모델 0건. backend/db/migrations/ grep 'oauth_identities\\|user_providers' 0건. /profile/security 에 'linked accounts' 섹션 0건. imp-0157 (provider diversity) 와 직접 충돌 가능 — 그쪽 제안만으로는 분리된 계정만 늘어남."

- id: imp-0303
  found_at_iter: 22
  area: auth
  type: ux
  target: oauth_error_code_specific_recovery
  problem: "imp-0169 가 generic 에러 분기를 다뤘으나 — Google OAuth 자체가 보내는 구체 에러(idtoken.Validate 실패, expired_token, invalid_token, audience_mismatch) 는 handlers_auth.go L52-58 에서 모두 'Google 인증에 실패했습니다' 단일 메시지로 묶여 클라이언트에 도달. 사용자는 '내 시계가 어긋난건지' '브라우저 캐시 문제인지' '계정 자체 문제인지' 알 수 없다 — recovery 가이드 0건."
  proposal: "(1) oauth/google.go L25-41 VerifyGoogleIDToken 에러를 wrap — fmt.Errorf('GOOGLE_TOKEN_EXPIRED: %w'), fmt.Errorf('GOOGLE_AUDIENCE_MISMATCH: %w'), fmt.Errorf('GOOGLE_INVALID_SIGNATURE: %w'). (2) handlers_auth.go L53-58 가 errors.Is/As 로 분기 → response code 별도 — 'GOOGLE_TOKEN_EXPIRED'(시계/네트워크 문제, 다시 시도), 'GOOGLE_AUDIENCE_MISMATCH'(앱 client_id 문제, 운영자 알림), 'GOOGLE_INVALID_SIGNATURE'(중간자/캐시, 새로고침 권유). (3) login/page.tsx L48-53 catch — code 별 한국어 메시지: '시계가 어긋났습니다 — 시간을 자동으로 맞추거나 브라우저를 새로고침해주세요' / '앱 설정 문제 — 운영자에게 자동 보고됨' / '인증 토큰 변조 의심 — 시크릿 모드에서 다시 시도'. (4) 'AUDIENCE_MISMATCH' 발생 시 backend 자동 Sentry 에러 보고(임계 alert)."
  effort: small
  impact: medium
  evidence: "코드 oauth/google.go L25-41 idtoken.Validate err 를 그냥 continue 하고 마지막에 generic error 반환(L41 'all client IDs'). errors.Is/As 0건. handlers_auth.go L53-58 단일 'UNAUTHORIZED' code. error code 카탈로그(types.ts grep 'GOOGLE_') 0건."

