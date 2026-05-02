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

