# 웹 프론트엔드 분리 설계 — Next.js (PC+모바일웹) + Flutter (네이티브 모바일)

## 배경

기란장터의 Flutter Web UI가 PC에서 모바일 전용 UX를 그대로 노출하는 문제. PC 사용자에게 데스크탑에 적합한 경험을 제공하기 위해 웹 프론트엔드를 Next.js로 분리한다.

## 결정 사항

| 항목 | 결정 |
|------|------|
| 웹 전체 (PC + 모바일 브라우저) | Next.js 15 (App Router) + Tailwind CSS 3.4 |
| 네이티브 모바일 (Android/iOS) | Flutter (기존 코드 유지) |
| 백엔드 | Go + Gin (변경 없음, API 공유) |
| 디자인 일관성 | 공유 디자인 토큰 (JSON → Tailwind config + Flutter AppTheme) |
| Flutter Web | 폐기 — Next.js가 웹 전담 |

## 아키텍처

```
┌───────────────────────────────────────────┐
│              공유 레이어                    │
│  Design Tokens (JSON)                     │
│  → tailwind.config.ts / Flutter AppTheme  │
│  API 계약 (OpenAPI → TypeScript 타입생성)  │
├───────────────────────────────────────────┤
│                                           │
│  Next.js Web          Flutter Native      │
│  (PC + 모바일웹)       (Android / iOS)     │
│  Tailwind CSS          Riverpod           │
│  TanStack Query        GoRouter           │
│                                           │
│         └──────┬──────────┘               │
│                ▼                          │
│  Go + Gin Backend                         │
│  REST API + SSE + PostgreSQL              │
└───────────────────────────────────────────┘
```

## 프로젝트 구조

```
프로젝트 루트/
├── backend/          # Go 백엔드 (변경 없음)
├── frontend/         # Flutter (네이티브 모바일 전용)
├── web/              # Next.js 프로젝트 (신규)
│   ├── app/              # App Router
│   │   ├── layout.tsx        # 루트 레이아웃
│   │   ├── page.tsx          # 매물 목록 (홈)
│   │   ├── listings/[id]/    # 매물 상세
│   │   ├── create/           # 매물 등록
│   │   ├── chats/            # 채팅 목록 + 상세
│   │   ├── profile/          # 프로필
│   │   └── login/            # 로그인
│   ├── components/       # 재사용 UI 컴포넌트
│   ├── lib/              # API 클라이언트, 유틸리티
│   ├── tailwind.config.ts
│   └── package.json
└── shared/           # 공유 디자인 토큰
    └── design-tokens.json
```

## 화면 매핑

| Flutter 화면 | Next.js 라우트 | PC 레이아웃 |
|---|---|---|
| listing_list_screen | `/` | 사이드바 + 그리드 카드 |
| listing_detail_screen | `/listings/[id]` | 2컬럼 (정보 + 사이드패널) |
| listing_create_screen | `/create` | 센터 폼 |
| chat_list_screen + chat_detail_screen | `/chats`, `/chats/[id]` | 분할 패널 (목록 + 대화) |
| profile_screen | `/profile` | 사이드바 내 패널 |
| login_screen | `/login` | 중앙 카드 |
| 바텀시트 (예약/리뷰/신고) | 모달 다이얼로그 | 센터 모달 |

## 반응형 Breakpoint

| 구간 | 범위 | 네비게이션 | 레이아웃 |
|------|------|-----------|---------|
| Mobile | < 768px | 하단 탭바 | 단일 컬럼, 페이지 전환 |
| Tablet | 768~1024px | 접힌 사이드바 | 2열 그리드 |
| Desktop | > 1024px | 고정 사이드바 | 3열 그리드, 분할 패널 |

## 기술 스택 상세

| 영역 | 선택 | 이유 |
|------|------|------|
| 프레임워크 | Next.js 15 (App Router) | SSR/SSG, SEO |
| 스타일링 | Tailwind CSS 3.4 | 반응형 내장, 디자인 토큰 매핑, 생태계 안정 |
| 상태/데이터 | TanStack Query | API 캐싱, 자동 리페치 |
| 인증 | 직접 JWT 관리 (POST /auth/login) | 백엔드가 JWT 발급, NextAuth 불필요 |
| 실시간 | EventSource (SSE) | query param 토큰 방식 (`?token=`) |
| 폼 | React Hook Form + Zod | 타입 안전한 폼 검증 |
| 타입생성 | 수동 TypeScript 타입 (OPENAPI_DRAFT.md 기반) | YAML 파일 미존재, 수동 정의 |

## 구현 순서 (Approach A — 점진적 구축)

**Phase 1: 기반 구축**
- 디자인 토큰 JSON 정의 → Tailwind config 매핑
- Next.js 프로젝트 초기화
- 공통 레이아웃 (사이드바 + 반응형 셸)
- API 클라이언트 + 타입 생성

**Phase 2: 핵심 화면**
- 매물 목록 (홈) — 그리드 + 필터 + 검색
- 매물 상세 — 2컬럼 레이아웃
- 로그인 (Google OAuth)
- 매물 등록 폼

**Phase 3: 커뮤니케이션**
- 채팅 목록 + 상세 (SSE 연동)
- 예약/거래완료 모달
- 리뷰/신고 모달

**Phase 4: 프로필 + 마무리**
- 프로필, 내 매물, 내 거래
- 알림
- 회귀 테스트 + 크로스브라우저
- 배포 설정 (Docker + Caddy)

## 인증 전략

백엔드(`POST /auth/login`)가 자체 JWT를 발급하므로 NextAuth.js 불필요.
- Google ID Token을 클라이언트에서 획득 → `POST /auth/login`으로 전송
- 반환된 accessToken/refreshToken을 localStorage에 저장
- 401 시 refreshToken으로 자동 갱신
- SSE 연결: `EventSource`는 커스텀 헤더 불가 → `?token={accessToken}` query param 방식

## 배포 및 메모리

NAS 2GB RAM 제약 사항:
- Next.js standalone 빌드로 최소 메모리 사용 (~80-120MB)
- 또는 `next export` 정적 빌드 후 Caddy에서 직접 서빙 (Node.js 프로세스 불필요, 0MB 추가)
- 정적 빌드 시 SSR 불가하지만, 기란장터는 클라이언트 렌더링으로 충분 (인증 필요한 기능 위주)

## 제약 사항

- 백엔드 API 변경 없음
- Flutter 기존 기능 유지 (네이티브 모바일용)
- 2GB RAM NAS 배포 환경 고려 — 정적 빌드 우선 검토
- 디자인 토큰으로 Flutter ↔ Web 룩앤필 일관성 보장
- PC UI 디자인은 구현 과정에서 구체화
- CORS: 프로덕션 origin을 backend `.env`의 `ALLOWED_ORIGINS`에 추가 필요
