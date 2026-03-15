# 웹 UX 전면 개선 — 다크 판타지 테마 적용 (design-proposals.html P1~P9)

## 배경

`docs/design-proposals.html`에 9개 디자인 제안이 작성되어 있으나 현재 웹 프론트엔드에 미적용 상태.
밝은 Material 3 테마 → 리니지 클래식 세계관 기반 다크 판타지 테마로 전면 전환.

## 적용할 디자인 제안 (P1~P9)

### P1: Dark Theme Color Palette
- 배경: `#0a0e17` (darkest) → `#222940` (light)
- 카드: `#1e2538`
- 악센트: gold `#c4a35a` / `#d4af37`, blue `#4a7fb5` / `#5b9bd5`
- 텍스트: `#e0e0e0` (primary), `#a0a0a0` (secondary), `#6a7080` (dim)
- 테두리: `#2a3045`
- 상태: success `#2ecc71`, danger `#e74c3c`
- 글로우: `rgba(212,175,55,0.15)` (gold), `rgba(74,127,181,0.15)` (blue)

### P2: Typography System
- Display: Black Han Sans (제목, 히어로)
- Body: Noto Sans KR (본문)
- Serif: Noto Serif KR (아이템 강조)
- 스케일: Display 48/900, H1 32/700, H2 24/700, H3 18/600, Body 15/400, Caption 12/300

### P3: Item Card Components
- 등급별 글로우 (normal/rare/unique/hero)
- 호버 시 금빛 글로우 + 상승 효과
- 아이콘: **리니지 클래식 아이콘만 사용** (`/static/icons/{icon_id}.png`)
- 이모지/외부 아이콘 라이브러리 사용 금지

### P4: Navigation & Header
- 고정 헤더 64px: 로고(Black Han Sans, gold) + 검색 + 알림(뱃지) + 프로필 아바타
- 서브 네비: 카테고리 칩 (pill shape, scrollable on mobile)
- 데스크탑: 헤더 네비 (마켓, 내 거래, 채팅)
- 모바일: 하단 탭바 유지

### P5: Item Detail Page
- 게임 툴팁 스타일 레이아웃
- 아이템 이미지 영역 (리니지 아이콘 크게 표시)
- 스탯 테이블 형식
- 가격 블록 강조
- 액션 버튼: 금빛 그라데이션

### P6: Chat Interface
- 다크 테마 채팅
- 좌측 사이드바 대화 목록 (데스크탑)
- 상단에 아이템 카드 미니뷰
- 내 메시지: blue 버블, 상대: card 배경
- 시스템: 골드 텍스트 중앙

### P7: Profile & Trust System
- RPG 캐릭터 프로필 스타일
- 스탯 카드 (거래수, 신뢰도, 응답속도)
- 뱃지 시스템 시각화

### P8: Buttons & Interactive Elements
- Primary: 금빛 그라데이션 (#8a7340 → #c4a35a)
- Secondary: 투명 + gold 테두리
- Ghost: 텍스트만, hover 시 gold
- 호버 글로우 + 상승 애니메이션

### P9: Loading & Animation
- 금빛 스피너
- 스켈레톤 로딩 (카드 배경색)
- 페이드인 애니메이션

## 아이콘 정책

**리니지 클래식 아이콘만 사용:**
- 아이템 아이콘: `/static/icons/{icon_id}.png` (331개)
- 네비게이션 아이콘: 이모지/lucide 대신 SVG 인라인 또는 CSS 심볼 사용
- 외부 아이콘 라이브러리 설치 금지

## 기술 범위

- **변경**: `web/` 디렉토리 (Next.js) — Tailwind config, 모든 컴포넌트/페이지
- **변경**: `shared/design-tokens.json` — 다크 테마 토큰으로 교체
- **변경 없음**: `backend/`, `frontend/`
- **추가 의존성**: Google Fonts (Black Han Sans, Noto Serif KR) — CSS import만
- **제거**: lucide-react 사용 안 함
