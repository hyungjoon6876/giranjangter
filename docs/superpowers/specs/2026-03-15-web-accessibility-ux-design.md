# 웹 접근성 + UX 개선 설계 — 사용자 흐름 기반

## 배경

기란장터 웹 프론트엔드(Next.js 16 + TailwindCSS)에 다크 판타지 테마가 적용된 상태. 그러나 첫 진입 시 "뚝 떨어지는 느낌"이 있고, WCAG AA 수준의 접근성이 미비함. PC 웹 우선(리니지 클래식이 PC 게임), 모바일도 잘 동작해야 함.

## 접근 방식

**사용자 흐름 퍼스트 (Top-Down)**: 첫 진입 → 탐색 → 상호작용 → 네비게이션 → 에러 순서로, 각 단계에서 접근성과 UX를 동시에 개선.

## 범위

- **대상**: `web/` 디렉토리 (Next.js)
- **포함**: 접근성(ARIA, 키보드, 시맨틱), UX(로딩, 에러, 검색, 네비게이션), 반응형 미세 조정
- **제외**: 디자인(색상/폰트/테마) 변경 — 이미 적용 완료
- **제외**: 키보드 단축키 — 혼란을 줄 수 있어 제외
- **제외**: 백엔드 변경 — 프론트엔드만
- **인증 정책**: 비로그인 사용자도 매물 탐색 자유. 로그인은 액션(찜, 채팅, 등록) 시점에 자연스럽게 유도.

---

## 섹션 1: 첫 진입 경험

### 1-A. 스켈레톤 로딩

**현재**: `<Loading />` 컴포넌트가 CSS border 스피너(`animate-spin`)만 표시. 콘텐츠 등장 시 화면 점프.

**개선**:
- `components/ui/skeleton.tsx` 신규 생성
- 매물 목록: 3~5개의 카드 형태 스켈레톤 (필터 칩 + 카드 레이아웃 예약)
- `globals.css`에 `.skeleton` 클래스 추가 (pulse 애니메이션)
- `role="status"` + `aria-label="매물 목록을 불러오는 중"` 적용
- `aria-busy="true"`를 부모 컨테이너에 적용

### 1-B. 자연스러운 로그인 진입점

**현재**: 헤더에 알림/프로필 아이콘이 있지만 비로그인 사용자에게 명시적인 로그인 유도 UI가 없음.

**개선**:
- 헤더 우측: 비로그인 시 **"로그인 / 시작하기"** 텍스트 버튼 표시
- 히어로 섹션 CTA "시작하기" 버튼 강화
- 찜하기 등 인증 필요 액션 시 → Toast로 로그인 유도 ("로그인이 필요합니다")
- 로그인 페이지의 `alert()` → 인라인 에러 메시지로 교체

### 1-C. Loading 컴포넌트 접근성

**현재**: `role`, `aria-label` 모두 없음. 스크린 리더에 정보 전달 안 됨.

**개선**:
- `role="status"` + `aria-label="로딩 중"` 추가
- `prefers-reduced-motion` 대응 (애니메이션 비활성화 시 "로딩 중..." 텍스트)
- sr-only 텍스트로 로딩 상태 명시

### 1-D. Toast 컴포넌트 (alert() 전면 교체)

**현재**: `login/page.tsx`, `listings/[id]/page.tsx`, `create/page.tsx`에서 `alert()` 사용.

**개선**:
- `components/ui/toast.tsx` 신규 생성
- `lib/hooks/use-toast.ts` + `ToastProvider` 전역 관리
- 3가지 변형: success (green), error (red), info (gold)
- `role="alert"` (암묵적으로 `aria-live="assertive"` 포함)
- 5초 자동 사라짐 + 수동 닫기 (`aria-label="닫기"`)
- 화면 우상단 고정 (`z-[60]`, 모달 z-50보다 위), 모바일은 상단 전체 너비
- 모든 `alert()` 호출을 toast로 교체

---

## 섹션 2: 메인 탐색 경험

### 2-A. 검색 접근성

**현재**: 검색 input에 `aria-label` 없음. placeholder만 의존. 결과 카운트 안내 없음.

**개선**:
- 검색 input을 `<search>` 랜드마크 요소(또는 `<form role="search">`)로 감싸기
- input `type="search"` + `aria-label="매물 검색"` 추가
- 헤더 검색과 필터 검색 두 곳 모두 동일하게 적용 (헤더 검색이 기능적이면 필터 state와 연동)
- 검색 결과 카운트를 `aria-live="polite"` 영역으로 표시 ("23개 매물")
- 검색 변경 시 스크린 리더에 결과 수 자동 안내

### 2-B. 필터 칩 키보드 접근성

**현재**: 서버 필터 버튼에 선택 상태를 프로그래밍 방식으로 전달하지 않음. 시각적으로만 표현.

**개선**:
- 버튼 그룹: `role="group"` + `aria-label="서버 필터"`
- 선택된 필터: `aria-pressed="true"` (토글 버튼 패턴, 필터에 적합)
- `focus-visible:ring-2 focus-visible:ring-gold` 포커스 링
- 필터 변경 시 결과 카운트 `aria-live` 업데이트

### 2-C. 결과 카운트 + 정렬 컨트롤

**현재**: 결과 수, 정렬 옵션이 없음.

**개선**:
- 필터 아래 정보 바: "N개 매물" + 정렬 select
- 네이티브 `<select>` 사용 (키보드/스크린 리더 완벽 지원)
- `aria-label="정렬 방식"` 연결
- 정렬 옵션: 최신순, 가격 낮은순, 가격 높은순, 인기순
- 결과 카운트: `aria-live="polite"`

### 2-D. 매물 카드 접근성

**현재**: Link에 aria-label 없음, truncated 텍스트 접근 불가, focus-visible 없음.

**개선**:
- Link에 `aria-label="{아이템명} - {가격} - {서버명}"` 조합 라벨
- truncated 텍스트에 `title` 속성
- `focus-visible:ring-2 focus-visible:ring-gold focus-visible:ring-offset-2`
- 호버/포커스 동일 시각적 피드백
- 아이콘 이미지: `alt="{아이템명} 아이콘"`

### 2-E. 컨텍스트 빈 상태

**현재**: 항상 동일한 "매물이 없습니다" 메시지.

**개선**:
- 검색 결과 없음: "'{검색어}'에 대한 매물이 없습니다"
- 필터 결과 없음: "선택한 서버에 매물이 없습니다. 다른 서버를 선택해보세요."
- 전체 빈 상태: "아직 매물이 없습니다. 첫 매물을 등록해보세요!"
- EmptyState 타이틀: `<p>` → `<h2>` (시맨틱 계층)
- 아이콘 emoji: `role="img"` + `aria-label`

---

## 섹션 3: 상세 페이지 + 상호작용

### 3-A. 모달 접근성 (CRITICAL)

**현재**: focus trap 없음, role="dialog" 없음, Escape 닫기 없음, 포커스 복귀 없음.

**개선** (`components/ui/modal.tsx` 전면 개선):
- `role="dialog"` + `aria-modal="true"`
- `aria-labelledby` → 모달 타이틀 id 연결
- 포커스 트랩: Tab/Shift+Tab이 모달 내부에서만 순환
- Escape 키로 닫기
- 닫기 버튼: `aria-label="닫기"`
- 닫힌 후 트리거 요소로 포커스 복귀
- Modal을 `createPortal`로 `document.body`에 렌더링
- 열릴 때 `<main>`, `<header>`, `<nav>` 등 형제 요소에 `inert` 속성 추가, 닫힐 때 제거
- → 예약/리뷰/신고 모달 모두 동일 Modal 사용하므로 한 번 수정으로 전체 적용

### 3-B. 매물 상세 페이지 접근성

**개선**:
- 아이템 정보: `<dl>/<dt>/<dd>` 정의 목록 구조
- 액션 바: `role="toolbar"` + `aria-label="매물 액션"`
- 찜 버튼: `aria-pressed="true/false"` + `aria-label="찜하기"/"찜 취소"`
- 채팅 시작 `alert()` → Toast 교체
- 로딩 실패 시 ErrorState + "다시 시도" 버튼

### 3-C. 채팅 실시간 UX + 접근성

**개선**:
- 메시지 영역: `role="log"` + `aria-live="polite"` (새 메시지 도착 안내)
- 시스템 메시지: `role="status"`
- 입력 필드: `aria-label="메시지 입력"`
- 전송 버튼: `aria-label="메시지 전송"`
- PC: Shift+Enter = 줄바꿈, Enter = 전송
- 전송 실패 시 Toast + 재전송 안내
- 활성 채팅: `aria-current="true"`
- 안 읽은 뱃지: `aria-label="N개 안 읽은 메시지"`
- 채팅방 선택 시 메시지 입력으로 auto-focus

### 3-D. SSE 연결 상태 표시

**현재**: 연결 끊어져도 사용자에게 피드백 없음. 무한 재시도.

**개선**:
- `useSSE` 훅에 `isConnected` / `isReconnecting` 상태 노출
- 연결 끊김 시 채팅 상단 배너: "연결이 끊어졌습니다. 재연결 중..."
- 지수 백오프: 1s → 2s → 4s → 8s...
- 최대 재시도 10회 제한
- 배너: `role="alert"` 스크린 리더 안내

---

## 섹션 4: 네비게이션 인프라

### 4-A. Skip-to-Content 링크

**개선** (`components/layout/responsive-shell.tsx`):
- `<a href="#main-content" class="sr-only focus:not-sr-only ...">본문으로 건너뛰기</a>`
- Tab 시 화면 상단에 금색 배경 링크 표시
- `<main id="main-content">` 연결
- WCAG 2.4.1 필수 요구사항

### 4-B. 헤더 접근성

**개선** (`components/layout/header.tsx`):
- `<header>` (최상위 header는 암묵적 banner role 보유)
- 데스크탑 네비: `<nav aria-label="메인 메뉴">`
- 활성 링크: `aria-current="page"`
- 알림 아이콘: `aria-label="알림"` (뱃지 있으면 `aria-label="알림 3건"`)
- 프로필 아이콘: `aria-label="내 프로필"`
- SVG 아이콘: `aria-hidden="true"`
- 검색 input: `aria-label="매물 검색"`

### 4-C. 하단 탭바 접근성

**개선** (`components/layout/bottom-nav.tsx`):
- `<nav aria-label="하단 메뉴">`
- 활성 탭: `aria-current="page"`
- 각 탭: `aria-label` 포함
- SVG: `aria-hidden="true"`
- `focus-visible:ring-2 focus-visible:ring-gold`

### 4-D. 네비게이션 뱃지

**개선**:
- 헤더 알림 아이콘: 읽지 않은 알림 있을 때 빨간 dot 뱃지
- 모바일 채팅 탭: 읽지 않은 메시지 수 숫자 뱃지
- SSE `new_message` 이벤트로 실시간 업데이트
- 뱃지 숫자를 `aria-label`에 포함

---

## 섹션 5: 에러 처리 + 엣지 케이스

### 5-A. Toast 시스템

섹션 1-D에서 정의. `components/ui/toast.tsx` + `lib/hooks/use-toast.ts`.

### 5-B. ErrorState 컴포넌트

**신규**: `components/ui/error-state.tsx`
- 아이콘 + 에러 메시지 + "다시 시도" 버튼
- `role="alert"` 즉시 안내
- 재시도 버튼 auto-focus
- 적용 대상: 매물 목록, 상세, 채팅, 알림, 프로필

### 5-C. 폼 검증 접근성

**개선** (`app/create/page.tsx` + 모든 폼):
- 모든 `<label>`에 `htmlFor` 연결, `<input>`/`<select>`에 `id` 추가 (WCAG 1.3.1 기본)
- `aria-describedby` → 에러 메시지 연결
- 에러 메시지: `role="alert"`
- 에러 발생 시 첫 번째 에러 필드로 포커스 이동
- 필수 필드: `aria-required="true"`
- 에러 필드: `aria-invalid="true"`
- `create/page.tsx`의 `alert()` → Toast로 교체

### 5-D. prefers-reduced-motion 대응

**개선** (`globals.css`):
- `@media (prefers-reduced-motion: reduce)` 추가
- 스피너 → 정적 텍스트 대체
- 카드 hover 애니메이션 → 배경색 변경만
- 트랜지션 duration → 0

### 5-E. 접근성 CSS 유틸리티

**추가** (`globals.css`):
- `.sr-only` — 시각적 숨김, 스크린 리더 노출
- `.not-sr-only` — sr-only 해제 (skip link 포커스)
- `.skeleton` — 스켈레톤 펄스 애니메이션
- `.focus-gold` — 통일 focus-visible 스타일

---

## 변경 대상 파일 요약

### 신규 생성
| 파일 | 목적 |
|------|------|
| `components/ui/skeleton.tsx` | 스켈레톤 로딩 컴포넌트 |
| `components/ui/toast.tsx` | Toast 알림 컴포넌트 |
| `components/ui/error-state.tsx` | 에러 상태 컴포넌트 |
| `lib/hooks/use-toast.ts` | Toast 상태 관리 훅 |

### 수정
| 파일 | 변경 내용 |
|------|-----------|
| `components/ui/modal.tsx` | focus trap, role="dialog", Escape, 포커스 복귀 |
| `components/ui/loading.tsx` | role="status", aria-label, reduced-motion |
| `components/ui/empty-state.tsx` | 시맨틱 계층, role="img", 컨텍스트별 메시지 |
| `components/layout/responsive-shell.tsx` | skip-to-content, main id |
| `components/layout/header.tsx` | aria-label, aria-current, 로그인 버튼 |
| `components/layout/bottom-nav.tsx` | aria-label, aria-current, focus-visible, 뱃지 |
| `components/listing/listing-filters.tsx` | role="group", aria-pressed, 결과 카운트, 정렬 |
| `components/listing/listing-card.tsx` | aria-label 조합, title, focus-visible, alt |
| `components/listing/listing-info.tsx` | dl/dt/dd 시맨틱 구조 |
| `components/chat/chat-panel.tsx` | role="log", aria-live, aria-current |
| `components/chat/chat-input.tsx` | aria-label, Shift+Enter |
| `components/chat/chat-message.tsx` | role="status" (시스템), 시간 정보 |
| `components/chat/chat-list-item.tsx` | aria-current, 뱃지 aria-label |
| `app/page.tsx` | 스켈레톤, 결과 카운트, 정렬, 컨텍스트 빈 상태, FAB `aria-label="매물 등록"` |
| `app/login/page.tsx` | alert() 제거, 인라인 에러 |
| `app/listings/[id]/page.tsx` | alert() 제거, ErrorState, aria-pressed, dl 구조 |
| `app/chats/page.tsx` | 채팅 접근성, auto-focus |
| `app/chats/[id]/page.tsx` | 모바일 채팅 상세 — 동일한 채팅 접근성 패턴 적용 |
| `app/create/page.tsx` | aria-describedby, aria-required, aria-invalid, label htmlFor, alert() 제거 |
| `app/globals.css` | sr-only, skeleton, focus-gold, reduced-motion |
| `lib/hooks/use-sse.ts` | isConnected, 지수 백오프, 최대 재시도 |
| `lib/providers.tsx` | ToastProvider 래핑 |

---

## 비범위 (명시적 제외)

- 색상/폰트/테마 변경 (다른 세션에서 완료)
- 키보드 단축키 (혼란 유발 가능)
- 백엔드 API 변경
- 새로운 페이지/라우트 추가
- 이미지 업로드 기능
- 웹 푸시 알림
