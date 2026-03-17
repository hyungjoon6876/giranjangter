# UX Phase 1: 핵심 사용성 개선 Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix the highest-impact UX issues: auth wall flow, favorite button bug, category/sort filter visibility, login redirect preservation, and auth-aware empty states.

**Architecture:** All changes are in the Next.js web frontend (`web/`). No backend changes needed — all APIs already exist. Each task is independent and can be implemented in any order.

**Tech Stack:** Next.js 16, React 19, TanStack Query, TailwindCSS 3.4

**Spec:** `docs/ux-analysis-2026-03-17.md` (sections 1.1–1.5, issues H1–H5, L1–L4, A1–A3, N1–N2, P1–P3, D2)

---

## File Structure

| File | Responsibility | Action |
|------|---------------|--------|
| `web/lib/hooks/use-auth-guard.ts` | Auth guard with redirect URL | Modify |
| `web/app/login/page.tsx` | Login page reads redirect param | Modify |
| `web/app/page.tsx` | Home page: hero conditional, category filter | Modify |
| `web/components/listing/listing-filters.tsx` | Add category filter chips | Modify |
| `web/app/listings/[id]/page.tsx` | Fix favorite button states | Modify |
| `web/app/notifications/page.tsx` | Auth-aware empty state | Modify |
| `web/app/create/page.tsx` | Use auth guard redirect | Modify |
| `web/app/profile/page.tsx` | Improve non-auth state | Modify |
| `web/app/profile/listings/page.tsx` | Add status filter for my listings | Modify |
| `web/lib/hooks/use-listings.ts` | Add useUpdateListing, useChangeStatus hooks | Modify |
| `web/lib/api-client.ts` | Add updateListing, changeListingStatus methods | Modify |
| `web/lib/hooks/use-auth-guard.test.ts` | Test auth guard redirect | Create |
| `web/components/listing/listing-filters.test.tsx` | Test category filter | Create |

---

## Chunk 1: Auth & Navigation Fixes

### Task 1: Auth Guard — redirect URL 보존

현재 `useAuthGuard`는 `/login`으로 보내지만 원래 URL을 보존하지 않는다. 로그인 후 원래 페이지로 돌아가도록 수정한다.

**Files:**
- Modify: `web/lib/hooks/use-auth-guard.ts`
- Create: `web/lib/hooks/use-auth-guard.test.ts`

- [ ] **Step 1: Write the test**

```typescript
// web/lib/hooks/use-auth-guard.test.ts
import { renderHook } from "@testing-library/react";
import { useAuthGuard } from "./use-auth-guard";
import { apiClient } from "@/lib/api-client";

// Mock dependencies
const mockPush = vi.fn();
const mockAddToast = vi.fn();

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: mockPush }),
  usePathname: () => "/create",
}));
vi.mock("@/lib/hooks/use-toast", () => ({
  useToast: () => ({ addToast: mockAddToast }),
}));
vi.mock("@/lib/api-client", () => ({
  apiClient: { isLoggedIn: false },
}));

describe("useAuthGuard", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("redirects to /login with redirect param when not logged in", () => {
    const { result } = renderHook(() => useAuthGuard());
    const allowed = result.current.requireAuth("채팅");
    expect(allowed).toBe(false);
    expect(mockPush).toHaveBeenCalledWith("/login?redirect=%2Fcreate");
    expect(mockAddToast).toHaveBeenCalledWith(
      "info",
      "채팅은(는) 로그인이 필요합니다",
    );
  });

  it("returns true when logged in", () => {
    (apiClient as { isLoggedIn: boolean }).isLoggedIn = true;
    const { result } = renderHook(() => useAuthGuard());
    expect(result.current.requireAuth()).toBe(true);
    expect(mockPush).not.toHaveBeenCalled();
  });
});
```

- [ ] **Step 2: Run test — expect FAIL**

```bash
cd web && npx vitest run lib/hooks/use-auth-guard.test.ts
```
Expected: FAIL — redirect param not yet implemented.

- [ ] **Step 3: Implement redirect preservation**

```typescript
// web/lib/hooks/use-auth-guard.ts
"use client";

import { useRouter, usePathname } from "next/navigation";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";

export function useAuthGuard() {
  const router = useRouter();
  const pathname = usePathname();
  const { addToast } = useToast();

  const requireAuth = (action?: string): boolean => {
    if (apiClient.isLoggedIn) return true;
    addToast(
      "info",
      action ? `${action}은(는) 로그인이 필요합니다` : "로그인이 필요합니다",
    );
    const redirect = encodeURIComponent(pathname);
    router.push(`/login?redirect=${redirect}`);
    return false;
  };

  return { isLoggedIn: apiClient.isLoggedIn, requireAuth };
}
```

- [ ] **Step 4: Run test — expect PASS**

```bash
cd web && npx vitest run lib/hooks/use-auth-guard.test.ts
```

- [ ] **Step 5: Commit**

```bash
git add web/lib/hooks/use-auth-guard.ts web/lib/hooks/use-auth-guard.test.ts
git commit -m "feat(web): auth guard에 redirect URL 보존 추가"
```

---

### Task 2: Login 페이지 — redirect 파라미터 처리

로그인 성공 후 `?redirect=` 파라미터가 있으면 해당 URL로 이동한다.

**Files:**
- Modify: `web/app/login/page.tsx`

- [ ] **Step 1: Modify login page to read redirect param**

`web/app/login/page.tsx`에서 다음을 변경:

1. `useSearchParams` import 추가
2. `redirect` 파라미터 읽기
3. `handleGoogleResponse`에서 `router.push(redirect || "/")`
4. "둘러보기" 버튼 텍스트 개선

```typescript
// Changes to web/app/login/page.tsx:

// Add import:
import { useRouter, useSearchParams } from "next/navigation";

// Inside LoginPage component, add:
const searchParams = useSearchParams();
const redirect = searchParams.get("redirect") || "/";

// In handleGoogleResponse, change:
//   router.push("/");
// to:
router.push(redirect);

// Change "둘러보기" button text:
//   둘러보기
// to:
//   로그인 없이 둘러보기
```

- [ ] **Step 2: Verify manually**

Navigate to `/login?redirect=%2Fcreate` — after login should go to `/create`.
Navigate to `/login` — after login should go to `/`.

- [ ] **Step 3: Commit**

```bash
git add web/app/login/page.tsx
git commit -m "feat(web): 로그인 후 redirect 복귀 + 둘러보기 문구 개선"
```

---

### Task 3: Create/Chats 페이지 — auth guard 통일

현재 `/create`는 `useEffect`로 직접 redirect하고, auth guard의 toast를 사용하지 않는다. auth guard를 통해 일관된 동작으로 변경한다.

**Files:**
- Modify: `web/app/create/page.tsx`

- [ ] **Step 1: Update create page auth guard**

`web/app/create/page.tsx`에서 다음을 변경:

```typescript
// Before (lines 17-23):
useEffect(() => {
  if (!isLoggedIn) {
    router.replace("/login");
  }
}, [isLoggedIn, router]);

if (!isLoggedIn) return null;

// After:
useEffect(() => {
  if (!isLoggedIn) {
    requireAuth("매물 등록");
  }
}, [isLoggedIn]);

if (!isLoggedIn) return null;
```

이렇게 하면 toast 메시지 + redirect URL 보존이 함께 동작한다.

참고: `useAuthGuard`에서 `requireAuth`를 destructure하고 `router.replace("/login")` 대신 사용.

- [ ] **Step 2: Verify — navigate to `/create` while logged out**

Expected: Toast "매물 등록은(는) 로그인이 필요합니다" + redirect to `/login?redirect=%2Fcreate`.

- [ ] **Step 3: Commit**

```bash
git add web/app/create/page.tsx
git commit -m "fix(web): create 페이지 auth guard 통일 — toast + redirect 보존"
```

---

### Task 4: Notifications 페이지 — auth-aware 빈 상태

비로그인 사용자에게 "로그인하면 알림을 받을 수 있습니다" 안내를 보여준다.

**Files:**
- Modify: `web/app/notifications/page.tsx`

- [ ] **Step 1: Add auth-aware empty state**

```typescript
// web/app/notifications/page.tsx — full replacement:
"use client";

import {
  useNotifications,
  useMarkNotificationsRead,
} from "@/lib/hooks/use-profile";
import { useIsLoggedIn } from "@/lib/hooks/use-auth";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { formatTimeAgo } from "@/lib/utils";

export default function NotificationsPage() {
  const isLoggedIn = useIsLoggedIn();
  const { data, isLoading } = useNotifications();
  const markRead = useMarkNotificationsRead();

  if (!isLoggedIn) {
    return (
      <EmptyState
        icon="🔔"
        title="로그인이 필요합니다"
        description="로그인하면 거래 알림을 받을 수 있습니다"
        actionLabel="로그인하기"
        actionHref="/login?redirect=%2Fnotifications"
      />
    );
  }

  const notifications = data?.data ?? [];

  if (isLoading) return <Loading />;
  if (!notifications.length) {
    return (
      <EmptyState
        icon="🔔"
        title="알림이 없습니다"
        description="새로운 거래 활동이 있으면 여기에 표시됩니다"
      />
    );
  }

  const unreadIds = notifications
    .filter((n) => !n.readAt)
    .map((n) => n.notificationId);

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold text-text-primary">알림</h1>
        {unreadIds.length > 0 && (
          <button
            onClick={() => markRead.mutate(unreadIds)}
            className="text-sm text-gold font-medium"
          >
            모두 읽음
          </button>
        )}
      </div>
      <div className="bg-card rounded-xl border border-border overflow-hidden">
        {notifications.map((n) => (
          <div
            key={n.notificationId}
            className={`px-5 py-4 border-b border-border last:border-0 ${!n.readAt ? "border-l-4 border-l-gold" : ""}`}
          >
            <p className="text-sm text-text-primary">{n.message}</p>
            <p className="text-xs text-text-dim mt-1">
              {formatTimeAgo(n.createdAt)}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}
```

Key changes:
- Import `useIsLoggedIn`
- Non-auth: show `🔔` icon + login prompt with redirect
- Auth + empty: show `🔔` icon (was `🔍`) + helpful description

- [ ] **Step 2: Verify — visit `/notifications` while logged out**

Expected: Bell icon + "로그인이 필요합니다" + "로그인하기" button linking to `/login?redirect=%2Fnotifications`.

- [ ] **Step 3: Commit**

```bash
git add web/app/notifications/page.tsx
git commit -m "fix(web): notifications 페이지 auth-aware 빈 상태 + 아이콘 수정"
```

---

### Task 5: Profile 페이지 — 비로그인 상태 개선

비로그인 시 프로필 페이지가 너무 밋밋하다. 시각적으로 개선하고 로그인 혜택을 안내한다.

**Files:**
- Modify: `web/app/profile/page.tsx`

- [ ] **Step 1: Improve non-auth profile state**

`web/app/profile/page.tsx`에서 비로그인 블록(lines 16-24)을 교체:

```typescript
// Before:
if (!me) {
  return (
    <div className="p-6 text-center">
      <p className="text-text-secondary mb-4">로그인이 필요합니다</p>
      <Link href="/login" className="text-gold font-medium">
        로그인하기
      </Link>
    </div>
  );
}

// After:
if (!me) {
  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <div className="bg-card rounded-xl border border-border p-8 text-center">
        <div className="w-20 h-20 rounded-full bg-medium flex items-center justify-center text-3xl mx-auto mb-4 border-2 border-border">
          👤
        </div>
        <h2 className="text-xl font-bold text-text-primary mb-2">
          로그인이 필요합니다
        </h2>
        <p className="text-text-secondary text-sm mb-6">
          로그인하면 매물 등록, 채팅, 거래 관리를 할 수 있습니다
        </p>
        <Link
          href="/login?redirect=%2Fprofile"
          className="inline-block btn-gold-gradient text-white px-6 py-3 rounded-lg font-medium"
        >
          로그인하기
        </Link>
      </div>
    </div>
  );
}
```

- [ ] **Step 2: Verify — visit `/profile` while logged out**

Expected: Card with avatar placeholder, descriptive text, gold login button.

- [ ] **Step 3: Commit**

```bash
git add web/app/profile/page.tsx
git commit -m "fix(web): profile 비로그인 상태 시각적 개선 + redirect"
```

---

## Chunk 2: Listing UX Fixes

### Task 6: Favorite 버튼 — 상태 구분

현재 찜 버튼이 "관심"/"관심"으로 동일하다. 아이콘으로 구분한다.

**Files:**
- Modify: `web/app/listings/[id]/page.tsx`

- [ ] **Step 1: Fix favorite button display**

`web/app/listings/[id]/page.tsx` lines 142-157에서:

```typescript
// Before:
{actions.includes("favorite") && (
  <button
    onClick={() => {
      if (!requireAuth("찜하기")) return;
      toggleFav.mutate({
        id: l.listingId,
        isFavorited: l.isFavorited ?? false,
      });
    }}
    aria-pressed={l.isFavorited ?? false}
    aria-label={l.isFavorited ? "찜 취소" : "찜하기"}
    className="p-3 bg-card border border-border rounded-lg hover:bg-medium transition-colors text-text-secondary"
  >
    {l.isFavorited ? "관심" : "관심"}
  </button>
)}

// After:
{actions.includes("favorite") && (
  <button
    onClick={() => {
      if (!requireAuth("찜하기")) return;
      toggleFav.mutate({
        id: l.listingId,
        isFavorited: l.isFavorited ?? false,
      });
    }}
    aria-pressed={l.isFavorited ?? false}
    aria-label={l.isFavorited ? "찜 취소" : "찜하기"}
    className={`p-3 border rounded-lg transition-colors ${
      l.isFavorited
        ? "bg-gold/10 border-gold text-gold"
        : "bg-card border-border text-text-secondary hover:bg-medium"
    }`}
  >
    {l.isFavorited ? "♥ 관심" : "♡ 관심"}
  </button>
)}
```

- [ ] **Step 2: Commit**

```bash
git add web/app/listings/[id]/page.tsx
git commit -m "fix(web): 찜 버튼 상태 시각적 구분 — 아이콘 + 배경색"
```

---

### Task 7: Listing 상세 — 조회수/찜/채팅 카운터 표시

상세 페이지에 조회수, 찜 수, 채팅 수 메타 정보를 추가한다.

**Files:**
- Modify: `web/app/listings/[id]/page.tsx`

- [ ] **Step 1: Add counters below title**

`web/app/listings/[id]/page.tsx`의 `{/* Title */}` 섹션 뒤(line 79 이후)에 카운터 추가:

```typescript
{/* Title */}
<h1 className="text-2xl font-bold mb-3 text-text-primary">{l.title}</h1>

{/* Counters */}
<div className="flex items-center gap-4 text-sm text-text-dim mb-4">
  <span>조회 {l.viewCount}</span>
  <span>관심 {l.favoriteCount}</span>
  <span>채팅 {l.chatCount}</span>
  <span>{formatTimeAgo(l.createdAt)}</span>
</div>
```

`formatTimeAgo`는 이미 `@/lib/utils`에서 import되어 있으므로 import 추가 필요. `utils.ts` import 라인에 `formatTimeAgo` 추가:

```typescript
import { formatPrice, statusLabel, statusColor, formatTimeAgo } from "@/lib/utils";
```

- [ ] **Step 2: Commit**

```bash
git add web/app/listings/[id]/page.tsx
git commit -m "feat(web): 매물 상세에 조회수/찜/채팅 카운터 표시"
```

---

### Task 8: Category 필터 추가

홈페이지에 카테고리 필터 칩을 추가한다. 서버 필터와 동일한 패턴.

**Files:**
- Modify: `web/app/page.tsx`
- Modify: `web/components/listing/listing-filters.tsx`

- [ ] **Step 1: Add category state to home page**

`web/app/page.tsx`에서:

```typescript
// Add state (after line 19):
const [categoryId, setCategoryId] = useState<string | null>(null);

// Add categories query (after servers query):
const { data: categories = [] } = useQuery({
  queryKey: ["categories"],
  queryFn: () => apiClient.getCategories(),
});

// Update useListings params (add categoryId):
const { data, isLoading, isError, refetch } = useListings({
  serverId: serverId ?? undefined,
  categoryId: categoryId ?? undefined,
  q: search || undefined,
  sort,
});

// Update ListingFilters props:
<ListingFilters
  servers={servers}
  selectedServer={serverId}
  onServerChange={setServerId}
  categories={categories}
  selectedCategory={categoryId}
  onCategoryChange={setCategoryId}
  searchQuery={search}
  onSearchChange={setSearch}
/>
```

- [ ] **Step 2: Update ListingFilters component**

```typescript
// web/components/listing/listing-filters.tsx — full replacement:
"use client";

import type { Server, Category } from "@/lib/types";

interface ListingFiltersProps {
  servers: Server[];
  selectedServer: string | null;
  onServerChange: (serverId: string | null) => void;
  categories?: Category[];
  selectedCategory?: string | null;
  onCategoryChange?: (categoryId: string | null) => void;
  searchQuery: string;
  onSearchChange: (q: string) => void;
}

export function ListingFilters({
  servers,
  selectedServer,
  onServerChange,
  categories = [],
  selectedCategory = null,
  onCategoryChange,
  searchQuery,
  onSearchChange,
}: ListingFiltersProps) {
  const topCategories = categories.filter((c) => !c.parentId);

  return (
    <div className="flex flex-col gap-3 mb-4">
      {/* Server + Category filter row */}
      <div className="flex flex-col lg:flex-row lg:items-center gap-3">
        <div role="group" aria-label="서버 필터" className="flex flex-wrap gap-2">
          <button
            onClick={() => onServerChange(null)}
            aria-pressed={selectedServer === null}
            className={`px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold ${
              selectedServer === null
                ? "btn-gold-gradient text-white"
                : "bg-medium text-text-secondary hover:bg-light"
            }`}
          >
            전체
          </button>
          {servers.map((s) => (
            <button
              key={s.serverId}
              onClick={() => onServerChange(s.serverId)}
              aria-pressed={selectedServer === s.serverId}
              className={`px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold ${
                selectedServer === s.serverId
                  ? "btn-gold-gradient text-white"
                  : "bg-medium text-text-secondary hover:bg-light"
              }`}
            >
              {s.serverName}
            </button>
          ))}
        </div>
        <search className="lg:ml-auto">
          <input
            type="search"
            aria-label="매물 검색"
            value={searchQuery}
            onChange={(e) => onSearchChange(e.target.value)}
            placeholder="아이템 검색..."
            className="w-full lg:w-60 bg-card border border-border rounded-lg px-3 py-2 text-sm text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim"
          />
        </search>
      </div>

      {/* Category filter chips */}
      {topCategories.length > 0 && onCategoryChange && (
        <div
          role="group"
          aria-label="카테고리 필터"
          className="flex flex-wrap gap-2"
        >
          <button
            onClick={() => onCategoryChange(null)}
            aria-pressed={selectedCategory === null}
            className={`px-3 py-1.5 rounded-lg text-xs whitespace-nowrap transition-colors border ${
              selectedCategory === null
                ? "border-gold text-gold bg-gold/10"
                : "border-border text-text-secondary bg-medium hover:bg-light"
            }`}
          >
            전체
          </button>
          {topCategories.map((c) => (
            <button
              key={c.categoryId}
              onClick={() => onCategoryChange(c.categoryId)}
              aria-pressed={selectedCategory === c.categoryId}
              className={`px-3 py-1.5 rounded-lg text-xs whitespace-nowrap transition-colors border ${
                selectedCategory === c.categoryId
                  ? "border-gold text-gold bg-gold/10"
                  : "border-border text-text-secondary bg-medium hover:bg-light"
              }`}
            >
              {c.categoryName}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
```

- [ ] **Step 3: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 4: Commit**

```bash
git add web/app/page.tsx web/components/listing/listing-filters.tsx
git commit -m "feat(web): 카테고리 필터 칩 추가 — 매물 목록 필터링"
```

---

### Task 9: My Listings — 상태 필터 + 매물 관리 버튼

내 매물 페이지에 상태 필터(판매중/예약중/완료)와 매물 수정/상태 변경 버튼을 추가한다.

**Files:**
- Modify: `web/lib/api-client.ts` — `updateListing`, `changeListingStatus` 메서드 추가
- Modify: `web/lib/hooks/use-listings.ts` — `useUpdateListing`, `useChangeListingStatus` 훅 추가
- Modify: `web/app/profile/listings/page.tsx` — 상태 필터 + 관리 버튼

- [ ] **Step 1: Add API client methods**

`web/lib/api-client.ts`에 메서드 추가 (기존 `createListing` 뒤):

```typescript
async updateListing(
  id: string,
  data: Partial<Listing>,
): Promise<Listing> {
  return this.fetch(`/listings/${id}`, {
    method: "PATCH",
    body: JSON.stringify(data),
  });
}

async changeListingStatus(
  id: string,
  status: string,
): Promise<void> {
  await this.fetch(`/listings/${id}/status`, {
    method: "POST",
    body: JSON.stringify({ status }),
  });
}
```

- [ ] **Step 2: Add hooks**

`web/lib/hooks/use-listings.ts`에 추가:

```typescript
export function useUpdateListing() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Listing> }) =>
      apiClient.updateListing(id, data),
    onSuccess: (_, { id }) => {
      qc.invalidateQueries({ queryKey: ["listing", id] });
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["my-listings"] });
    },
  });
}

export function useChangeListingStatus() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) =>
      apiClient.changeListingStatus(id, status),
    onSuccess: (_, { id }) => {
      qc.invalidateQueries({ queryKey: ["listing", id] });
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["my-listings"] });
    },
  });
}
```

- [ ] **Step 3: Update My Listings page**

```typescript
// web/app/profile/listings/page.tsx — full replacement:
"use client";

import { useState } from "react";
import Link from "next/link";
import { useMyListings } from "@/lib/hooks/use-profile";
import { useChangeListingStatus } from "@/lib/hooks/use-listings";
import { useToast } from "@/lib/hooks/use-toast";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { statusLabel } from "@/lib/utils";

const STATUS_FILTERS = [
  { value: undefined, label: "전체" },
  { value: "available", label: "판매중" },
  { value: "reserved", label: "예약중" },
  { value: "completed", label: "완료" },
  { value: "cancelled", label: "취소" },
] as const;

export default function MyListingsPage() {
  const [statusFilter, setStatusFilter] = useState<string | undefined>(
    undefined,
  );
  const { data, isLoading } = useMyListings(statusFilter);
  const changeStatus = useChangeListingStatus();
  const { addToast } = useToast();

  if (isLoading) return <Loading />;

  const listings = data?.data ?? [];

  return (
    <div className="p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold text-text-primary">내 매물</h1>
        <Link
          href="/create"
          className="btn-gold-gradient text-white px-4 py-2 rounded-lg text-sm"
        >
          + 등록
        </Link>
      </div>

      {/* Status filter */}
      <div
        role="group"
        aria-label="상태 필터"
        className="flex flex-wrap gap-2 mb-4"
      >
        {STATUS_FILTERS.map((f) => (
          <button
            key={f.label}
            onClick={() => setStatusFilter(f.value)}
            aria-pressed={statusFilter === f.value}
            className={`px-3 py-1.5 rounded-lg text-xs transition-colors border ${
              statusFilter === f.value
                ? "border-gold text-gold bg-gold/10"
                : "border-border text-text-secondary bg-medium hover:bg-light"
            }`}
          >
            {f.label}
          </button>
        ))}
      </div>

      {!listings.length ? (
        <EmptyState
          title="등록한 매물이 없습니다"
          actionLabel="매물 등록하기"
          actionHref="/create"
        />
      ) : (
        <ListingGrid listings={listings} />
      )}
    </div>
  );
}
```

참고: `useMyListings` hook이 status 파라미터를 받도록 이미 구현되어 있는지 확인 필요. `web/lib/hooks/use-profile.ts`에서 `useMyListings`의 시그니처를 확인하고, status 파라미터를 넘기지 않으면 수정한다.

- [ ] **Step 4: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 5: Commit**

```bash
git add web/lib/api-client.ts web/lib/hooks/use-listings.ts web/app/profile/listings/page.tsx
git commit -m "feat(web): 내 매물 상태 필터 + 매물 수정/상태 변경 API 연동"
```

---

## Chunk 3: Home Page & Polish

### Task 10: Home — Hero 섹션 개선

Hero 섹션이 매물 목록을 밀어내는 문제를 해결한다. 비로그인 + 첫 방문에만 전체 Hero를 표시하고, 그 외에는 compact 버전을 보여준다.

**Files:**
- Modify: `web/app/page.tsx`

- [ ] **Step 1: Implement collapsible hero**

현재 `{!isLoggedIn && ( ... )}` 조건으로 이미 로그인 사용자에겐 Hero가 숨겨져 있다. 이 부분은 이미 잘 구현되어 있으므로, **비로그인 사용자에게도 compact한 버전**을 제공한다.

`web/app/page.tsx`의 Hero 섹션을 다음과 같이 변경:

```typescript
// Before (line 36-87): 전체 Hero 섹션
// After: Feature 카드를 제거하고 더 compact하게

{!isLoggedIn && (
  <section className="relative overflow-hidden rounded-xl mb-6 p-6 lg:p-8 bg-gradient-to-br from-dark via-card to-medium border border-border">
    <div className="relative z-10 flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
      <div>
        <img src="/logo.png" alt="기란JT" className="h-12 lg:h-14 mb-2" />
        <p className="text-text-secondary">
          리니지 클래식 아이템 거래, 안전하고 무료
        </p>
      </div>
      <div className="flex gap-3">
        <a
          href="#listings"
          className="btn-gold-gradient text-white px-5 py-2.5 rounded-lg font-medium text-sm text-center"
        >
          매물 둘러보기
        </a>
        <Link
          href="/login"
          className="border border-gold text-gold px-5 py-2.5 rounded-lg font-medium text-sm text-center hover:bg-gold/10"
        >
          시작하기
        </Link>
      </div>
    </div>
    <div className="absolute top-0 right-0 w-48 h-48 bg-gold/5 rounded-full blur-3xl" />
  </section>
)}
```

이렇게 하면:
- Hero 높이가 절반 이하로 줄어든다
- Feature 카드 3개를 제거하여 공간 절약
- 데스크톱에서는 가로 배치로 더 compact
- 매물 목록이 above-the-fold에 바로 보인다

- [ ] **Step 2: Verify — home page as non-logged-in user**

Expected: Compact hero + 매물 목록이 스크롤 없이 보인다.

- [ ] **Step 3: Commit**

```bash
git add web/app/page.tsx
git commit -m "fix(web): Hero 섹션 compact화 — 매물 목록 바로 노출"
```

---

### Task 11: Header 검색바 — 매물 목록 검색과 연동

현재 Header의 검색바와 매물 목록의 검색바가 별도로 동작한다. Header 검색바를 글로벌 검색으로 만들어서 홈페이지 매물 목록 검색으로 연결한다.

**Files:**
- Modify: `web/components/layout/header.tsx`

- [ ] **Step 1: Read header.tsx**

먼저 현재 Header 구현을 확인한다.

- [ ] **Step 2: Update header search to navigate to home with query**

Header 검색바에서 Enter를 누르면 `/?q=검색어`로 이동하도록 수정한다.

```typescript
// In header.tsx search input, add onKeyDown handler:
const handleHeaderSearch = (e: React.KeyboardEvent<HTMLInputElement>) => {
  if (e.key === "Enter") {
    const q = (e.target as HTMLInputElement).value.trim();
    if (q) {
      router.push(`/?q=${encodeURIComponent(q)}`);
    }
  }
};
```

참고: page.tsx에서 `useSearchParams`로 `q` 파라미터를 읽어 `search` 상태를 초기화하는 것도 필요. 하지만 현재 구조상 `search`가 state로 관리되므로, URL → state 동기화가 필요하면 추가 작업이 필요하다.

**간단한 접근: Header 검색바를 숨기고 매물 목록 검색바만 사용**

현재 desktop에서 2개의 검색바가 혼란을 줄 수 있다. 가장 간단한 수정은 Header 검색바를 홈이 아닌 다른 페이지에서만 보이게 하거나, 홈에서는 숨기는 것이다.

이 태스크는 Header 구현을 확인한 후 구체적인 접근을 결정한다.

- [ ] **Step 3: Commit**

```bash
git add web/components/layout/header.tsx
git commit -m "fix(web): Header 검색바 개선"
```

---

## Summary

| Task | 파일 | 이슈 | 복잡도 |
|------|------|------|--------|
| 1 | use-auth-guard.ts | Auth redirect URL 보존 | S |
| 2 | login/page.tsx | Login redirect 파라미터 처리 | S |
| 3 | create/page.tsx | Create auth guard 통일 | S |
| 4 | notifications/page.tsx | Auth-aware 빈 상태 | S |
| 5 | profile/page.tsx | 비로그인 상태 개선 | S |
| 6 | listings/[id]/page.tsx | 찜 버튼 상태 구분 | S |
| 7 | listings/[id]/page.tsx | 조회수/찜/채팅 카운터 | S |
| 8 | listing-filters.tsx, page.tsx | 카테고리 필터 | M |
| 9 | profile/listings/page.tsx | 내 매물 상태 필터 + API | M |
| 10 | page.tsx | Hero compact화 | S |
| 11 | header.tsx | 검색바 정리 | S |

**Phase 2 대상 (별도 계획):** 이미지 업로드, 매물 수정 폼, 프로필 수정, 사용자 리뷰, 공유 기능
**Phase 3 대상 (별도 계획):** 카카오/네이버 OAuth, 사용자 차단, 푸터, PWA

---

## Production Data Note

서버/카테고리 API가 `{"data":null}`을 반환한다. `backend/db/seed/seed.sql`의 시드 데이터가 프로덕션 DB에 적용되지 않은 상태이다. 코드 문제가 아닌 DB 시드 문제이므로, 배포 시 다음을 실행해야 한다:

```bash
# NAS에서 실행
psql -U lincle -d lincle -f /path/to/seed.sql
```
