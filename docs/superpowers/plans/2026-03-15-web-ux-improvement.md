# Web UX Improvement Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development or superpowers:executing-plans to implement this plan.

**Goal:** Improve the web frontend UX with welcome/guide content, professional PC navigation, and visual variety across all pages.

**Architecture:** Pure UI changes in `web/` — no backend or Flutter changes. Add `lucide-react` for icons, refactor layout components, enhance page components.

**Tech Stack:** Next.js 15, Tailwind CSS 3.4, lucide-react (new)

---

## Task 1: Install lucide-react + Setup

**Files:**
- Modify: `web/package.json`

- [ ] Install lucide-react: `cd web && npm install lucide-react`
- [ ] Commit: "chore: add lucide-react icon library"

---

## Task 2: Sidebar Redesign (PC Navigation)

**Files:**
- Rewrite: `web/components/layout/sidebar.tsx`

- [ ] Replace emoji icons with Lucide icons (Home, MessageCircle, PenSquare, User, Bell, LogOut, LogIn)
- [ ] Change color scheme: dark slate-800 → white bg + border-right + primary accent
- [ ] Add user profile section at bottom (avatar + nickname if logged in, login button if not)
- [ ] Add secondary nav: 알림, 도움말
- [ ] Width: w-52 → w-60
- [ ] Active state: left border primary + bg-primary/5 + text-primary
- [ ] Commit: "feat: redesign PC sidebar with Lucide icons and user profile"

---

## Task 3: Desktop Header

**Files:**
- Rewrite: `web/components/layout/header.tsx`

- [ ] Show on ALL breakpoints (remove lg:hidden)
- [ ] Desktop: page breadcrumb/title area + global search input + notification bell with badge + user avatar dropdown
- [ ] Mobile: app title + notification bell
- [ ] Commit: "feat: add desktop header with search and notifications"

---

## Task 4: Mobile Bottom Nav Enhancement

**Files:**
- Modify: `web/components/layout/bottom-nav.tsx`

- [ ] 3 tabs → 4 tabs: 매물, 채팅, 등록(+), 프로필
- [ ] Replace emoji with Lucide icons
- [ ] Add notification badge on chat tab (unread count)
- [ ] Commit: "feat: enhance mobile bottom nav with Lucide icons and 4 tabs"

---

## Task 5: Home Page Hero Section

**Files:**
- Modify: `web/app/page.tsx`

- [ ] Add hero section for non-logged-in users:
  - Gradient background (primary to primaryDark)
  - Title: "기란장터" + subtitle: "리니지 클래식 아이템 거래, 안전하고 무료"
  - 3 feature cards: 무료 거래, 신뢰 시스템, 실시간 채팅
  - CTA: "매물 둘러보기" (scrolls to listings) + "시작하기" (→ /login)
- [ ] Hide hero for logged-in users (check apiClient.isLoggedIn)
- [ ] Commit: "feat: add hero welcome section for new users on home page"

---

## Task 6: Empty State Improvements

**Files:**
- Modify: `web/components/ui/empty-state.tsx`

- [ ] Add optional `actionLabel` and `actionHref` props
- [ ] Render Link button when action is provided
- [ ] Update all empty state usages:
  - Home: "첫 매물을 등록해보세요!" + [매물 등록] button
  - Chat: "매물에서 채팅을 시작해보세요" + [매물 둘러보기] button
  - Profile listings: [매물 등록하기] button
  - Notifications: "아직 알림이 없습니다" (no action needed)
- [ ] Commit: "feat: improve empty states with action buttons"

---

## Task 7: Listing Card Visual Enhancement

**Files:**
- Modify: `web/components/listing/listing-card.tsx`

- [ ] Add left border color by listingType (sell=primary, buy=secondary) — `border-l-4`
- [ ] Hover effect: `hover:shadow-md hover:-translate-y-0.5 transition-all duration-200`
- [ ] Meta counts: replace text emoji with Lucide icons (Eye, Heart, MessageCircle)
- [ ] Commit: "feat: enhance listing card with colored border and hover effects"

---

## Task 8: Listing Detail Page Improvement

**Files:**
- Modify: `web/app/listings/[id]/page.tsx`
- Modify: `web/components/listing/listing-info.tsx`

- [ ] Move price to hero area (top, prominent)
- [ ] Wrap item info in a light card (bg-surface rounded-xl p-4)
- [ ] Author section: separate card with bg-surface
- [ ] Action bar: `sticky bottom-0` on ALL breakpoints (remove lg:relative)
- [ ] Report button: use Lucide Flag icon
- [ ] Commit: "feat: improve listing detail layout with price hero and card sections"

---

## Task 9: Chat Context Enhancement

**Files:**
- Modify: `web/components/chat/chat-panel.tsx`
- Modify: `web/app/chats/[id]/page.tsx`

- [ ] Desktop chat header: show listing title + price + status badge
- [ ] Desktop: add mini listing card in the right panel header
- [ ] Mobile chat header: show counterparty name + listing title
- [ ] Commit: "feat: add listing context to chat headers"

---

## Task 10: Notification Page Enhancement

**Files:**
- Modify: `web/app/notifications/page.tsx`

- [ ] Unread: left blue border (border-l-4 border-primary) instead of just bg-blue-50
- [ ] Add Lucide icons per notification type (if distinguishable from message text)
- [ ] Make notification clickable (link to relevant page if possible)
- [ ] Commit: "feat: improve notification styling with colored borders"

---

## Task 11: Typography & Spacing Consistency

**Files:**
- Multiple pages

- [ ] Page titles: consistent `text-2xl font-bold` + optional `text-text-secondary text-sm` subtitle
- [ ] Section titles: `text-lg font-semibold`
- [ ] Page padding: consistent `p-4 lg:p-6 max-w-6xl mx-auto` on all content pages
- [ ] Commit: "refactor: normalize typography and spacing across pages"

---

## Task 12: Create Listing Form Guidance

**Files:**
- Modify: `web/app/create/page.tsx`

- [ ] Add helper text below key fields (e.g., "2~100자", "10~2000자")
- [ ] Price grid: `grid-cols-1 sm:grid-cols-2` (responsive)
- [ ] Add section headers: "기본 정보", "가격", "거래 방식"
- [ ] Commit: "feat: add form guidance and section headers to create listing"

---

## Task 13: Final Verification

- [ ] Run `npm run build` — all routes compile
- [ ] Run `npx vitest run` — all tests pass
- [ ] Verify responsive: desktop sidebar, mobile bottom nav
- [ ] Verify hero hides for logged-in users
- [ ] Commit: "verify: all UX improvements complete, tests passing"
