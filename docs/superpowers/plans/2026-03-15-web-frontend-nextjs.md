# Web Frontend (Next.js) Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a Next.js web frontend (PC + mobile responsive) that replicates all existing Flutter features with desktop-optimized UI, sharing the same Go backend API.

**Architecture:** Next.js 15 (App Router) with Tailwind CSS for styling, TanStack Query for data fetching, NextAuth.js for Google OAuth. Shared design tokens JSON ensures visual consistency with Flutter. The web/ directory sits alongside existing backend/ and frontend/ directories.

**Tech Stack:** Next.js 15, React 19, Tailwind CSS 3.4, TanStack Query v5, React Hook Form, Zod, TypeScript 5

---

## File Structure

```
web/
├── package.json
├── tsconfig.json
├── next.config.ts
├── tailwind.config.ts          # Design tokens from shared/design-tokens.json
├── postcss.config.mjs
├── .env.local                  # API_URL, NEXTAUTH_*, GOOGLE_*
├── app/
│   ├── globals.css             # Tailwind directives + base styles
│   ├── layout.tsx              # Root layout (responsive shell)
│   ├── page.tsx                # Home = listing list
│   ├── login/
│   │   └── page.tsx            # Login page
│   ├── listings/
│   │   └── [id]/
│   │       └── page.tsx        # Listing detail
│   ├── create/
│   │   └── page.tsx            # Create listing form
│   ├── chats/
│   │   ├── page.tsx            # Chat list (mobile) / split panel (desktop)
│   │   └── [id]/
│   │       └── page.tsx        # Chat detail (mobile only, desktop uses panel)
│   ├── profile/
│   │   ├── page.tsx            # Profile overview
│   │   ├── listings/
│   │   │   └── page.tsx        # My listings
│   │   └── trades/
│   │       └── page.tsx        # My trades
│   ├── notifications/
│   │   └── page.tsx            # Notifications
│   └── api/
│       └── auth/
│           └── [...nextauth]/
│               └── route.ts    # NextAuth API route
├── components/
│   ├── layout/
│   │   ├── sidebar.tsx         # Desktop sidebar navigation
│   │   ├── bottom-nav.tsx      # Mobile bottom navigation
│   │   ├── responsive-shell.tsx # Switches sidebar/bottom-nav by breakpoint
│   │   └── header.tsx          # Mobile top header
│   ├── listing/
│   │   ├── listing-card.tsx    # Listing card (grid item)
│   │   ├── listing-filters.tsx # Server/category/type filter bar
│   │   ├── listing-grid.tsx    # Responsive grid container
│   │   └── listing-info.tsx    # Detail page info sections
│   ├── chat/
│   │   ├── chat-list-item.tsx  # Chat room list item
│   │   ├── chat-message.tsx    # Single message bubble
│   │   ├── chat-input.tsx      # Message input bar
│   │   └── chat-panel.tsx      # Desktop split panel (list + messages)
│   ├── ui/
│   │   ├── badge.tsx           # Status/type badge
│   │   ├── modal.tsx           # Desktop modal dialog
│   │   ├── empty-state.tsx     # Empty state placeholder
│   │   └── loading.tsx         # Loading spinner/skeleton
│   └── forms/
│       ├── reservation-modal.tsx
│       ├── review-modal.tsx
│       └── report-modal.tsx
├── lib/
│   ├── api-client.ts           # Fetch wrapper with JWT token management
│   ├── auth.ts                 # NextAuth configuration
│   ├── types.ts                # Auto-generated API types (from openapi-typescript)
│   ├── hooks/
│   │   ├── use-listings.ts     # TanStack Query hooks for listings
│   │   ├── use-chats.ts        # TanStack Query hooks for chats
│   │   ├── use-profile.ts      # TanStack Query hooks for profile
│   │   └── use-sse.ts          # SSE connection hook
│   └── utils.ts                # formatPrice, formatTimeAgo, statusLabel, statusColor
└── __tests__/
    ├── components/
    │   ├── listing-card.test.tsx
    │   ├── chat-message.test.tsx
    │   └── badge.test.tsx
    ├── lib/
    │   ├── api-client.test.ts
    │   └── utils.test.ts
    └── app/
        ├── home.test.tsx
        └── listing-detail.test.tsx

shared/
└── design-tokens.json          # Colors, spacing, typography, border-radius
```

---

## Chunk 1: Foundation

### Task 1: Shared Design Tokens

**Files:**
- Create: `shared/design-tokens.json`

- [ ] **Step 1: Create design tokens JSON**

```json
{
  "colors": {
    "primary": "#2563EB",
    "primaryDark": "#1D4ED8",
    "secondary": "#059669",
    "error": "#DC2626",
    "warning": "#F59E0B",
    "surface": "#F8FAFC",
    "textPrimary": "#1E293B",
    "textSecondary": "#64748B",
    "border": "#E2E8F0",
    "white": "#FFFFFF"
  },
  "borderRadius": {
    "sm": "4px",
    "md": "8px",
    "lg": "10px",
    "xl": "12px",
    "full": "9999px"
  },
  "fontSize": {
    "xs": "11px",
    "sm": "12px",
    "base": "14px",
    "lg": "16px",
    "xl": "18px",
    "2xl": "22px",
    "3xl": "28px"
  },
  "spacing": {
    "1": "4px",
    "2": "8px",
    "3": "12px",
    "4": "16px",
    "6": "24px",
    "8": "32px",
    "14": "56px"
  }
}
```

> Maps 1:1 from Flutter `AppTheme` — primary=#2563EB, secondary=#059669, surface=#F8FAFC, border=#E2E8F0, borderRadius=10/12, fontSize=11-28

- [ ] **Step 2: Commit**

```bash
git add shared/design-tokens.json
git commit -m "feat: add shared design tokens JSON for Flutter/Web consistency"
```

---

### Task 2: Next.js Project Initialization

**Files:**
- Create: `web/` (entire directory via create-next-app)

- [ ] **Step 1: Create Next.js project**

```bash
cd /Users/jym/github-workspace/lincle/.claude/worktrees/rippling-crafting-noodle
npx create-next-app@latest web --typescript --tailwind --eslint --app --src-dir=false --import-alias="@/*" --use-npm
```

Expected: Next.js project created in `web/` directory.

- [ ] **Step 2: Install dependencies**

```bash
cd web
npm install @tanstack/react-query @tanstack/react-query-devtools react-hook-form @hookform/resolvers zod
npm install -D @testing-library/react @testing-library/jest-dom vitest @vitejs/plugin-react jsdom
```

- [ ] **Step 3: Verify project runs**

```bash
cd web && npm run dev
```

Expected: Next.js dev server starts on http://localhost:3000

- [ ] **Step 4: Commit**

```bash
git add web/
git commit -m "feat: initialize Next.js project with dependencies"
```

---

### Task 3: Tailwind Config with Design Tokens

**Files:**
- Modify: `web/tailwind.config.ts`
- Modify: `web/app/globals.css`

- [ ] **Step 1: Configure Tailwind with design tokens**

```typescript
// web/tailwind.config.ts
import type { Config } from "tailwindcss";
import tokens from "../shared/design-tokens.json";

const config: Config = {
  content: [
    "./app/**/*.{ts,tsx}",
    "./components/**/*.{ts,tsx}",
    "./lib/**/*.{ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: { DEFAULT: tokens.colors.primary, dark: tokens.colors.primaryDark },
        secondary: tokens.colors.secondary,
        error: tokens.colors.error,
        warning: tokens.colors.warning,
        surface: tokens.colors.surface,
        "text-primary": tokens.colors.textPrimary,
        "text-secondary": tokens.colors.textSecondary,
        border: tokens.colors.border,
      },
      borderRadius: {
        sm: tokens.borderRadius.sm,
        md: tokens.borderRadius.md,
        lg: tokens.borderRadius.lg,
        xl: tokens.borderRadius.xl,
      },
      fontSize: {
        xs: tokens.fontSize.xs,
        sm: tokens.fontSize.sm,
        base: tokens.fontSize.base,
        lg: tokens.fontSize.lg,
        xl: tokens.fontSize.xl,
        "2xl": tokens.fontSize["2xl"],
        "3xl": tokens.fontSize["3xl"],
      },
    },
  },
  plugins: [],
};
export default config;
```

- [ ] **Step 2: Set up global CSS**

```css
/* web/app/globals.css */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  body {
    @apply bg-surface text-text-primary;
  }
}
```

- [ ] **Step 3: Verify Tailwind works**

```bash
cd web && npm run dev
```

Expected: Page renders with `bg-surface` (#F8FAFC) background.

- [ ] **Step 4: Commit**

```bash
git add web/tailwind.config.ts web/app/globals.css shared/design-tokens.json
git commit -m "feat: configure Tailwind CSS with shared design tokens"
```

---

### Task 4: Utility Functions

**Files:**
- Create: `web/lib/utils.ts`
- Create: `web/__tests__/lib/utils.test.ts`

- [ ] **Step 1: Write failing tests**

```typescript
// web/__tests__/lib/utils.test.ts
import { describe, it, expect } from "vitest";
import { formatPrice, formatTimeAgo, statusLabel, statusColor } from "@/lib/utils";

describe("formatPrice", () => {
  it("formats number with commas", () => {
    expect(formatPrice(500000)).toBe("500,000");
    expect(formatPrice(8000000)).toBe("8,000,000");
    expect(formatPrice(0)).toBe("0");
  });

  it("returns '가격 제안' for null/undefined", () => {
    expect(formatPrice(null)).toBe("가격 제안");
    expect(formatPrice(undefined)).toBe("가격 제안");
  });
});

describe("statusLabel", () => {
  it("maps status to Korean label", () => {
    expect(statusLabel("available")).toBe("판매중");
    expect(statusLabel("reserved")).toBe("예약중");
    expect(statusLabel("completed")).toBe("거래완료");
    expect(statusLabel("cancelled")).toBe("취소됨");
  });
});

describe("statusColor", () => {
  it("returns hex color for status", () => {
    expect(statusColor("available")).toBe("#059669");
    expect(statusColor("reserved")).toBe("#F59E0B");
    expect(statusColor("completed")).toBe("#64748B");
  });
});

describe("formatTimeAgo", () => {
  it("formats recent time as minutes", () => {
    const fiveMinAgo = new Date(Date.now() - 5 * 60_000).toISOString();
    expect(formatTimeAgo(fiveMinAgo)).toBe("5분 전");
  });

  it("formats hours", () => {
    const twoHoursAgo = new Date(Date.now() - 2 * 3600_000).toISOString();
    expect(formatTimeAgo(twoHoursAgo)).toBe("2시간 전");
  });
});
```

- [ ] **Step 2: Run tests to verify they fail**

```bash
cd web && npx vitest run __tests__/lib/utils.test.ts
```

Expected: FAIL — module `@/lib/utils` not found.

- [ ] **Step 3: Implement utility functions**

```typescript
// web/lib/utils.ts
export function formatPrice(amount: number | null | undefined): string {
  if (amount == null) return "가격 제안";
  return amount.toLocaleString("ko-KR");
}

const STATUS_LABELS: Record<string, string> = {
  available: "판매중",
  reserved: "예약중",
  pending_trade: "거래중",
  completed: "거래완료",
  cancelled: "취소됨",
};

export function statusLabel(status: string): string {
  return STATUS_LABELS[status] ?? status;
}

const STATUS_COLORS: Record<string, string> = {
  available: "#059669",
  reserved: "#F59E0B",
  pending_trade: "#2563EB",
  completed: "#64748B",
  cancelled: "#DC2626",
};

export function statusColor(status: string): string {
  return STATUS_COLORS[status] ?? "#64748B";
}

export function formatTimeAgo(isoString: string): string {
  const diff = Date.now() - new Date(isoString).getTime();
  const minutes = Math.floor(diff / 60_000);
  if (minutes < 1) return "방금 전";
  if (minutes < 60) return `${minutes}분 전`;
  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours}시간 전`;
  const days = Math.floor(hours / 24);
  if (days < 30) return `${days}일 전`;
  const months = Math.floor(days / 30);
  return `${months}개월 전`;
}
```

- [ ] **Step 4: Run tests to verify they pass**

```bash
cd web && npx vitest run __tests__/lib/utils.test.ts
```

Expected: All tests PASS.

- [ ] **Step 5: Commit**

```bash
git add web/lib/utils.ts web/__tests__/lib/utils.test.ts
git commit -m "feat: add utility functions (formatPrice, statusLabel, formatTimeAgo)"
```

---

### Task 5: API Client

**Files:**
- Create: `web/lib/api-client.ts`
- Create: `web/lib/types.ts`
- Create: `web/__tests__/lib/api-client.test.ts`

- [ ] **Step 1: Create TypeScript types** (manually for now, openapi-typescript later)

```typescript
// web/lib/types.ts
export interface User {
  userId: string;
  nickname: string;
  avatarUrl?: string;
  introduction?: string;
  primaryServerId?: string;
  role: string;
  accountStatus: string;
  completedTradeCount: number;
  positiveReviewCount: number;
  responseBadge?: string;
  trustBadge?: string;
  lastActiveAt: string;
  createdAt: string;
}

export interface Author {
  userId: string;
  nickname: string;
  avatarUrl?: string;
  trustBadge?: string;
  responseBadge?: string;
  completedTradeCount?: number;
  lastActiveAt?: string;
}

export interface Listing {
  listingId: string;
  listingType: "sell" | "buy";
  title: string;
  itemName: string;
  description?: string;
  priceType: "fixed" | "negotiable" | "offer";
  priceAmount?: number;
  quantity: number;
  enhancementLevel?: number;
  optionsText?: string;
  serverId: string;
  serverName: string;
  categoryId: string;
  categoryName?: string;
  status: string;
  visibility: string;
  tradeMethod: string;
  preferredMeetingAreaText?: string;
  availableTimeText?: string;
  thumbnailUrl?: string;
  iconUrl?: string;
  images?: { imageId: string; url: string; order: number }[];
  author: Author;
  viewCount: number;
  favoriteCount: number;
  chatCount: number;
  isFavorited?: boolean;
  isOwner?: boolean;
  availableActions?: string[];
  lastActivityAt: string;
  createdAt: string;
  updatedAt?: string;
}

export interface ChatRoom {
  chatRoomId: string;
  listingId: string;
  listingTitle: string;
  listingThumbnail?: string;
  listingStatus: string;
  counterparty: Author;
  chatStatus: string;
  lastMessage?: Message;
  unreadCount: number;
  updatedAt: string;
}

export interface Message {
  messageId: string;
  chatRoomId: string;
  senderUserId?: string;
  messageType: "text" | "system" | "reservation_card";
  bodyText?: string;
  metadataJson?: Record<string, unknown>;
  sentAt: string;
}

export interface Server {
  serverId: string;
  serverName: string;
}

export interface Category {
  categoryId: string;
  categoryName: string;
  parentId?: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  cursor: { next?: string; hasMore: boolean };
}

export interface AuthResponse {
  accessToken: string;
  refreshToken: string;
  expiresIn: number;
  user: User;
}

export interface ApiError {
  error: { code: string; message: string; details?: Record<string, unknown> };
}
```

- [ ] **Step 2: Create API client**

```typescript
// web/lib/api-client.ts
import type {
  Listing, ChatRoom, Message, Server, Category,
  PaginatedResponse, AuthResponse, User,
} from "./types";

const API_BASE = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";

class ApiClient {
  private accessToken: string | null = null;
  private refreshToken: string | null = null;

  constructor() {
    if (typeof window !== "undefined") {
      this.accessToken = localStorage.getItem("accessToken");
      this.refreshToken = localStorage.getItem("refreshToken");
    }
  }

  get isLoggedIn(): boolean {
    return this.accessToken !== null;
  }

  saveTokens(access: string, refresh: string): void {
    this.accessToken = access;
    this.refreshToken = refresh;
    localStorage.setItem("accessToken", access);
    localStorage.setItem("refreshToken", refresh);
  }

  clearTokens(): void {
    this.accessToken = null;
    this.refreshToken = null;
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
  }

  private async fetch<T>(path: string, init?: RequestInit): Promise<T> {
    const headers: Record<string, string> = {
      "Content-Type": "application/json",
      ...(init?.headers as Record<string, string>),
    };
    if (this.accessToken) {
      headers["Authorization"] = `Bearer ${this.accessToken}`;
    }

    let res = await fetch(`${API_BASE}${path}`, { ...init, headers });

    // Auto-refresh on 401
    if (res.status === 401 && this.refreshToken) {
      const refreshed = await this.doRefresh();
      if (refreshed) {
        headers["Authorization"] = `Bearer ${this.accessToken}`;
        res = await fetch(`${API_BASE}${path}`, { ...init, headers });
      }
    }

    if (!res.ok) {
      const err = await res.json().catch(() => ({ error: { code: "UNKNOWN", message: res.statusText } }));
      throw err;
    }

    if (res.status === 204) return undefined as T;
    return res.json();
  }

  private async doRefresh(): Promise<boolean> {
    try {
      const res = await fetch(`${API_BASE}/auth/refresh`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ refreshToken: this.refreshToken }),
      });
      if (res.ok) {
        const data = await res.json();
        this.saveTokens(data.accessToken, data.refreshToken);
        return true;
      }
    } catch {}
    return false;
  }

  // Auth
  async login(provider: string, token: string): Promise<AuthResponse> {
    const data = await this.fetch<AuthResponse>("/auth/login", {
      method: "POST",
      body: JSON.stringify({ provider, providerToken: token }),
    });
    this.saveTokens(data.accessToken, data.refreshToken);
    return data;
  }

  async getMe(): Promise<User> {
    return this.fetch("/me");
  }

  // Listings
  async getListings(params?: {
    serverId?: string; categoryId?: string; q?: string;
    listingType?: string; sort?: string; cursor?: string; limit?: number;
  }): Promise<PaginatedResponse<Listing>> {
    const qs = new URLSearchParams();
    if (params?.serverId) qs.set("serverId", params.serverId);
    if (params?.categoryId) qs.set("categoryId", params.categoryId);
    if (params?.q) qs.set("q", params.q);
    if (params?.listingType) qs.set("listingType", params.listingType);
    qs.set("sort", params?.sort ?? "recent");
    qs.set("limit", String(params?.limit ?? 20));
    if (params?.cursor) qs.set("cursor", params.cursor);
    return this.fetch(`/listings?${qs}`);
  }

  async getListing(id: string): Promise<Listing> {
    return this.fetch(`/listings/${id}`);
  }

  async createListing(data: Partial<Listing>): Promise<Listing> {
    return this.fetch("/listings", { method: "POST", body: JSON.stringify(data) });
  }

  async favoriteListing(id: string): Promise<void> {
    return this.fetch(`/listings/${id}/favorite`, { method: "POST" });
  }

  async unfavoriteListing(id: string): Promise<void> {
    return this.fetch(`/listings/${id}/favorite`, { method: "DELETE" });
  }

  // Chat
  async createChat(listingId: string): Promise<{ chatRoomId: string }> {
    return this.fetch(`/listings/${listingId}/chats`, { method: "POST" });
  }

  async getChats(): Promise<PaginatedResponse<ChatRoom>> {
    return this.fetch("/chats");
  }

  async getMessages(chatId: string): Promise<PaginatedResponse<Message>> {
    return this.fetch(`/chats/${chatId}/messages`);
  }

  async sendMessage(chatId: string, text: string): Promise<Message> {
    return this.fetch(`/chats/${chatId}/messages`, {
      method: "POST",
      body: JSON.stringify({ messageType: "text", bodyText: text }),
    });
  }

  // Profile
  async getMyListings(status?: string): Promise<PaginatedResponse<Listing>> {
    const qs = status ? `?status=${status}` : "";
    return this.fetch(`/me/listings${qs}`);
  }

  async getMyTrades(): Promise<PaginatedResponse<Listing>> {
    return this.fetch("/me/trades");
  }

  async getNotifications(): Promise<PaginatedResponse<{ notificationId: string; message: string; readAt?: string; createdAt: string }>> {
    return this.fetch("/notifications");
  }

  async markNotificationsRead(ids: string[]): Promise<void> {
    return this.fetch("/notifications/read", {
      method: "POST",
      body: JSON.stringify({ notificationIds: ids }),
    });
  }

  // Reservation
  async createReservation(chatId: string, data: Record<string, unknown>): Promise<unknown> {
    return this.fetch(`/chats/${chatId}/reservations`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // Trade
  async completeTrade(listingId: string, data: Record<string, unknown>): Promise<unknown> {
    return this.fetch(`/listings/${listingId}/complete`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // Review
  async createReview(completionId: string, data: Record<string, unknown>): Promise<unknown> {
    return this.fetch(`/trade-completions/${completionId}/reviews`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // Report
  async createReport(data: Record<string, unknown>): Promise<unknown> {
    return this.fetch("/reports", { method: "POST", body: JSON.stringify(data) });
  }

  // Master data
  async getServers(): Promise<Server[]> {
    const res = await this.fetch<{ data: Server[] }>("/servers");
    return res.data;
  }

  async getCategories(): Promise<Category[]> {
    const res = await this.fetch<{ data: Category[] }>("/categories");
    return res.data;
  }
}

export const apiClient = new ApiClient();
```

- [ ] **Step 3: Write API client test**

```typescript
// web/__tests__/lib/api-client.test.ts
import { describe, it, expect, vi, beforeEach } from "vitest";

// Test that API client constructs URLs correctly
describe("ApiClient URL construction", () => {
  it("builds listing query params correctly", () => {
    const qs = new URLSearchParams();
    qs.set("serverId", "bartz");
    qs.set("sort", "recent");
    qs.set("limit", "20");
    expect(qs.toString()).toBe("serverId=bartz&sort=recent&limit=20");
  });
});
```

- [ ] **Step 4: Run test**

```bash
cd web && npx vitest run __tests__/lib/api-client.test.ts
```

Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add web/lib/types.ts web/lib/api-client.ts web/__tests__/lib/api-client.test.ts
git commit -m "feat: add API client and TypeScript types matching Flutter ApiClient"
```

---

### Task 6: TanStack Query Provider + Hooks

**Files:**
- Create: `web/lib/providers.tsx`
- Create: `web/lib/hooks/use-listings.ts`
- Create: `web/lib/hooks/use-chats.ts`
- Create: `web/lib/hooks/use-profile.ts`

- [ ] **Step 1: Create query provider**

```tsx
// web/lib/providers.tsx
"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState, type ReactNode } from "react";

export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(
    () => new QueryClient({
      defaultOptions: {
        queries: { staleTime: 30_000, retry: 1 },
      },
    })
  );
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}
```

- [ ] **Step 2: Create listing hooks**

```typescript
// web/lib/hooks/use-listings.ts
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { Listing } from "@/lib/types";

export function useListings(params?: {
  serverId?: string; categoryId?: string; q?: string;
  listingType?: string; sort?: string;
}) {
  return useQuery({
    queryKey: ["listings", params],
    queryFn: () => apiClient.getListings(params),
  });
}

export function useListing(id: string) {
  return useQuery({
    queryKey: ["listing", id],
    queryFn: () => apiClient.getListing(id),
    enabled: !!id,
  });
}

export function useCreateListing() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: Partial<Listing>) => apiClient.createListing(data),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["listings"] }),
  });
}

export function useToggleFavorite() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, isFavorited }: { id: string; isFavorited: boolean }) =>
      isFavorited ? apiClient.unfavoriteListing(id) : apiClient.favoriteListing(id),
    onSuccess: (_, { id }) => {
      qc.invalidateQueries({ queryKey: ["listing", id] });
      qc.invalidateQueries({ queryKey: ["listings"] });
    },
  });
}
```

- [ ] **Step 3: Create chat hooks**

```typescript
// web/lib/hooks/use-chats.ts
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useChats() {
  return useQuery({
    queryKey: ["chats"],
    queryFn: () => apiClient.getChats(),
  });
}

export function useMessages(chatId: string) {
  return useQuery({
    queryKey: ["messages", chatId],
    queryFn: () => apiClient.getMessages(chatId),
    enabled: !!chatId,
    refetchInterval: 5_000, // Poll every 5s as fallback to SSE
  });
}

export function useSendMessage() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ chatId, text }: { chatId: string; text: string }) =>
      apiClient.sendMessage(chatId, text),
    onSuccess: (_, { chatId }) => {
      qc.invalidateQueries({ queryKey: ["messages", chatId] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useCreateChat() {
  return useMutation({
    mutationFn: (listingId: string) => apiClient.createChat(listingId),
  });
}
```

- [ ] **Step 4: Create profile hooks**

```typescript
// web/lib/hooks/use-profile.ts
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useMe() {
  return useQuery({
    queryKey: ["me"],
    queryFn: () => apiClient.getMe(),
    enabled: apiClient.isLoggedIn,
  });
}

export function useMyListings(status?: string) {
  return useQuery({
    queryKey: ["myListings", status],
    queryFn: () => apiClient.getMyListings(status),
    enabled: apiClient.isLoggedIn,
  });
}

export function useMyTrades() {
  return useQuery({
    queryKey: ["myTrades"],
    queryFn: () => apiClient.getMyTrades(),
    enabled: apiClient.isLoggedIn,
  });
}

export function useNotifications() {
  return useQuery({
    queryKey: ["notifications"],
    queryFn: () => apiClient.getNotifications(),
    enabled: apiClient.isLoggedIn,
    refetchInterval: 30_000,
  });
}

export function useMarkNotificationsRead() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (ids: string[]) => apiClient.markNotificationsRead(ids),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["notifications"] }),
  });
}
```

- [ ] **Step 5: Commit**

```bash
git add web/lib/providers.tsx web/lib/hooks/
git commit -m "feat: add TanStack Query provider and data hooks"
```

---

### Task 7: Responsive Layout Shell

**Files:**
- Create: `web/components/layout/sidebar.tsx`
- Create: `web/components/layout/bottom-nav.tsx`
- Create: `web/components/layout/header.tsx`
- Create: `web/components/layout/responsive-shell.tsx`
- Modify: `web/app/layout.tsx`

- [ ] **Step 1: Create sidebar component (desktop)**

```tsx
// web/components/layout/sidebar.tsx
"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const NAV_ITEMS = [
  { href: "/", icon: "🏠", label: "매물" },
  { href: "/chats", icon: "💬", label: "채팅" },
  { href: "/create", icon: "📝", label: "매물 등록" },
  { href: "/profile", icon: "👤", label: "프로필" },
  { href: "/notifications", icon: "🔔", label: "알림" },
];

export function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="hidden lg:flex flex-col w-52 bg-slate-800 text-white min-h-screen flex-shrink-0">
      <div className="px-5 py-4 text-lg font-bold border-b border-slate-700">
        기란장터
      </div>
      <nav className="flex-1 py-2">
        {NAV_ITEMS.map((item) => {
          const isActive = item.href === "/"
            ? pathname === "/"
            : pathname.startsWith(item.href);
          return (
            <Link
              key={item.href}
              href={item.href}
              className={`flex items-center gap-3 px-5 py-3 text-sm transition-colors ${
                isActive
                  ? "bg-slate-700 text-white"
                  : "text-slate-400 hover:bg-slate-700/50 hover:text-white"
              }`}
            >
              <span>{item.icon}</span>
              {item.label}
            </Link>
          );
        })}
      </nav>
    </aside>
  );
}
```

- [ ] **Step 2: Create bottom nav (mobile)**

```tsx
// web/components/layout/bottom-nav.tsx
"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const TABS = [
  { href: "/", icon: "🏠", label: "매물" },
  { href: "/chats", icon: "💬", label: "채팅" },
  { href: "/profile", icon: "👤", label: "프로필" },
];

export function BottomNav() {
  const pathname = usePathname();

  return (
    <nav className="lg:hidden fixed bottom-0 left-0 right-0 bg-white border-t border-border flex z-50">
      {TABS.map((tab) => {
        const isActive = tab.href === "/"
          ? pathname === "/"
          : pathname.startsWith(tab.href);
        return (
          <Link
            key={tab.href}
            href={tab.href}
            className={`flex-1 flex flex-col items-center py-2 text-xs ${
              isActive ? "text-primary" : "text-text-secondary"
            }`}
          >
            <span className="text-lg">{tab.icon}</span>
            {tab.label}
          </Link>
        );
      })}
    </nav>
  );
}
```

- [ ] **Step 3: Create mobile header**

```tsx
// web/components/layout/header.tsx
"use client";

import Link from "next/link";

export function Header() {
  return (
    <header className="lg:hidden flex items-center justify-between px-4 py-3 bg-white border-b border-border">
      <Link href="/" className="text-lg font-bold text-text-primary">
        기란장터
      </Link>
    </header>
  );
}
```

- [ ] **Step 4: Create responsive shell**

```tsx
// web/components/layout/responsive-shell.tsx
"use client";

import { Sidebar } from "./sidebar";
import { BottomNav } from "./bottom-nav";
import { Header } from "./header";
import type { ReactNode } from "react";

export function ResponsiveShell({ children }: { children: ReactNode }) {
  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 pb-16 lg:pb-0">{children}</main>
        <BottomNav />
      </div>
    </div>
  );
}
```

- [ ] **Step 5: Update root layout**

```tsx
// web/app/layout.tsx
import type { Metadata } from "next";
import { Providers } from "@/lib/providers";
import { ResponsiveShell } from "@/components/layout/responsive-shell";
import "./globals.css";

export const metadata: Metadata = {
  title: "기란장터 — 리니지 클래식 거래",
  description: "리니지 클래식 아이템 거래 중개 플랫폼",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ko">
      <body>
        <Providers>
          <ResponsiveShell>{children}</ResponsiveShell>
        </Providers>
      </body>
    </html>
  );
}
```

- [ ] **Step 6: Verify layout renders**

```bash
cd web && npm run dev
```

Expected: Desktop shows sidebar on left, mobile shows bottom tabs. Verify by resizing browser.

- [ ] **Step 7: Commit**

```bash
git add web/components/layout/ web/app/layout.tsx
git commit -m "feat: add responsive layout shell (sidebar for desktop, bottom nav for mobile)"
```

---

### Task 8: Shared UI Components

**Files:**
- Create: `web/components/ui/badge.tsx`
- Create: `web/components/ui/modal.tsx`
- Create: `web/components/ui/empty-state.tsx`
- Create: `web/components/ui/loading.tsx`

- [ ] **Step 1: Create badge component**

```tsx
// web/components/ui/badge.tsx
interface BadgeProps {
  label: string;
  color: string; // hex color
}

export function Badge({ label, color }: BadgeProps) {
  return (
    <span
      className="inline-block px-2 py-0.5 text-xs font-semibold rounded"
      style={{ color, backgroundColor: `${color}1A` }}
    >
      {label}
    </span>
  );
}

export function TypeBadge({ type }: { type: "sell" | "buy" }) {
  return type === "sell"
    ? <Badge label="판매" color="#2563EB" />
    : <Badge label="구매" color="#059669" />;
}
```

- [ ] **Step 2: Create modal, empty-state, loading**

```tsx
// web/components/ui/modal.tsx
"use client";

import { useEffect, type ReactNode } from "react";

interface ModalProps {
  open: boolean;
  onClose: () => void;
  title: string;
  children: ReactNode;
}

export function Modal({ open, onClose, title, children }: ModalProps) {
  useEffect(() => {
    if (open) document.body.style.overflow = "hidden";
    return () => { document.body.style.overflow = ""; };
  }, [open]);

  if (!open) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="absolute inset-0 bg-black/40" onClick={onClose} />
      <div className="relative bg-white rounded-xl shadow-xl w-full max-w-md mx-4 max-h-[90vh] overflow-y-auto">
        <div className="flex items-center justify-between px-5 py-4 border-b border-border">
          <h2 className="text-lg font-semibold">{title}</h2>
          <button onClick={onClose} className="text-text-secondary hover:text-text-primary">✕</button>
        </div>
        <div className="p-5">{children}</div>
      </div>
    </div>
  );
}
```

```tsx
// web/components/ui/empty-state.tsx
interface EmptyStateProps {
  icon?: string;
  title: string;
  description?: string;
}

export function EmptyState({ icon = "🔍", title, description }: EmptyStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-20 text-text-secondary">
      <span className="text-5xl mb-4">{icon}</span>
      <p className="text-lg">{title}</p>
      {description && <p className="text-sm mt-2">{description}</p>}
    </div>
  );
}
```

```tsx
// web/components/ui/loading.tsx
export function Loading() {
  return (
    <div className="flex items-center justify-center py-20">
      <div className="w-8 h-8 border-4 border-primary border-t-transparent rounded-full animate-spin" />
    </div>
  );
}
```

- [ ] **Step 3: Commit**

```bash
git add web/components/ui/
git commit -m "feat: add shared UI components (Badge, Modal, EmptyState, Loading)"
```

---

## Chunk 2: Core Screens

### Task 9: Listing Card Component

**Files:**
- Create: `web/components/listing/listing-card.tsx`
- Create: `web/__tests__/components/listing-card.test.tsx`

- [ ] **Step 1: Write failing test**

```tsx
// web/__tests__/components/listing-card.test.tsx
import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { ListingCard } from "@/components/listing/listing-card";

const mockListing = {
  listingId: "1",
  listingType: "sell" as const,
  title: "집행검 +7 팝니다",
  itemName: "집행검",
  priceType: "negotiable" as const,
  priceAmount: 500000,
  enhancementLevel: 7,
  serverId: "bartz",
  serverName: "바츠",
  status: "available",
  author: { userId: "u1", nickname: "검은기사" },
  viewCount: 12,
  favoriteCount: 3,
  chatCount: 1,
  createdAt: new Date().toISOString(),
  // required fields
  quantity: 1, visibility: "public", tradeMethod: "in_game",
  categoryId: "weapon", lastActivityAt: new Date().toISOString(),
};

describe("ListingCard", () => {
  it("renders title and price", () => {
    render(<ListingCard listing={mockListing} />);
    expect(screen.getByText("집행검 +7 팝니다")).toBeDefined();
    expect(screen.getByText(/500,000/)).toBeDefined();
  });

  it("shows sell badge", () => {
    render(<ListingCard listing={mockListing} />);
    expect(screen.getByText("판매")).toBeDefined();
  });
});
```

- [ ] **Step 2: Run to verify fail**

```bash
cd web && npx vitest run __tests__/components/listing-card.test.tsx
```

- [ ] **Step 3: Implement listing card**

```tsx
// web/components/listing/listing-card.tsx
import Link from "next/link";
import type { Listing } from "@/lib/types";
import { TypeBadge, Badge } from "@/components/ui/badge";
import { formatPrice, formatTimeAgo, statusLabel, statusColor } from "@/lib/utils";

export function ListingCard({ listing }: { listing: Listing }) {
  const l = listing;
  return (
    <Link
      href={`/listings/${l.listingId}`}
      className="block bg-white border border-border rounded-xl p-4 hover:shadow-md transition-shadow"
    >
      <div className="flex items-center gap-2 mb-2">
        <TypeBadge type={l.listingType} />
        <Badge label={statusLabel(l.status)} color={statusColor(l.status)} />
        <span className="ml-auto text-sm text-text-secondary">{l.serverName}</span>
      </div>
      <h3 className="font-semibold text-text-primary truncate">{l.title}</h3>
      <div className="flex items-center gap-1 mt-1 text-sm text-text-secondary">
        {l.iconUrl && (
          <img
            src={`${process.env.NEXT_PUBLIC_API_URL?.replace("/api/v1", "") ?? "http://localhost:8080"}${l.iconUrl}`}
            alt=""
            className="w-5 h-5"
          />
        )}
        <span>{l.itemName}</span>
        {l.enhancementLevel != null && (
          <span className="text-primary font-semibold">+{l.enhancementLevel}</span>
        )}
      </div>
      <div className="flex items-center mt-2">
        <span className="text-lg font-bold">{formatPrice(l.priceAmount)}원</span>
        {l.priceType === "negotiable" && (
          <span className="text-xs text-text-secondary ml-1">(협상가능)</span>
        )}
        <div className="ml-auto flex items-center gap-3 text-xs text-text-secondary">
          <span>👁 {l.viewCount}</span>
          <span>♥ {l.favoriteCount}</span>
          <span>💬 {l.chatCount}</span>
        </div>
      </div>
      <div className="flex items-center justify-between mt-2 text-xs text-text-secondary">
        <span>{l.author.nickname}</span>
        <span>{formatTimeAgo(l.createdAt)}</span>
      </div>
    </Link>
  );
}
```

- [ ] **Step 4: Run test**

```bash
cd web && npx vitest run __tests__/components/listing-card.test.tsx
```

Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add web/components/listing/listing-card.tsx web/__tests__/components/listing-card.test.tsx
git commit -m "feat: add ListingCard component matching Flutter listing card"
```

---

### Task 10: Listing Filters + Grid

**Files:**
- Create: `web/components/listing/listing-filters.tsx`
- Create: `web/components/listing/listing-grid.tsx`

- [ ] **Step 1: Create filters component**

```tsx
// web/components/listing/listing-filters.tsx
"use client";

import type { Server } from "@/lib/types";

interface ListingFiltersProps {
  servers: Server[];
  selectedServer: string | null;
  onServerChange: (serverId: string | null) => void;
  searchQuery: string;
  onSearchChange: (q: string) => void;
}

export function ListingFilters({
  servers, selectedServer, onServerChange, searchQuery, onSearchChange,
}: ListingFiltersProps) {
  return (
    <div className="flex flex-col lg:flex-row lg:items-center gap-3 mb-4">
      <div className="flex gap-2 overflow-x-auto pb-1">
        <button
          onClick={() => onServerChange(null)}
          className={`px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors ${
            selectedServer === null
              ? "bg-primary text-white"
              : "bg-slate-100 text-text-secondary hover:bg-slate-200"
          }`}
        >
          전체
        </button>
        {servers.map((s) => (
          <button
            key={s.serverId}
            onClick={() => onServerChange(s.serverId)}
            className={`px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors ${
              selectedServer === s.serverId
                ? "bg-primary text-white"
                : "bg-slate-100 text-text-secondary hover:bg-slate-200"
            }`}
          >
            {s.serverName}
          </button>
        ))}
      </div>
      <div className="lg:ml-auto">
        <input
          type="text"
          value={searchQuery}
          onChange={(e) => onSearchChange(e.target.value)}
          placeholder="🔍 아이템 검색..."
          className="w-full lg:w-60 bg-white border border-border rounded-lg px-3 py-2 text-sm outline-none focus:border-primary focus:ring-1 focus:ring-primary"
        />
      </div>
    </div>
  );
}
```

- [ ] **Step 2: Create responsive grid**

```tsx
// web/components/listing/listing-grid.tsx
import type { Listing } from "@/lib/types";
import { ListingCard } from "./listing-card";

export function ListingGrid({ listings }: { listings: Listing[] }) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
      {listings.map((l) => (
        <ListingCard key={l.listingId} listing={l} />
      ))}
    </div>
  );
}
```

- [ ] **Step 3: Commit**

```bash
git add web/components/listing/listing-filters.tsx web/components/listing/listing-grid.tsx
git commit -m "feat: add ListingFilters and responsive ListingGrid"
```

---

### Task 11: Home Page (Listing List)

**Files:**
- Modify: `web/app/page.tsx`

- [ ] **Step 1: Implement home page**

```tsx
// web/app/page.tsx
"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useListings } from "@/lib/hooks/use-listings";
import { ListingFilters } from "@/components/listing/listing-filters";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import Link from "next/link";

export default function HomePage() {
  const [serverId, setServerId] = useState<string | null>(null);
  const [search, setSearch] = useState("");

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
  });

  const { data, isLoading } = useListings({
    serverId: serverId ?? undefined,
    q: search || undefined,
  });

  return (
    <div className="p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold hidden lg:block">매물 목록</h1>
        <Link
          href="/create"
          className="hidden lg:inline-flex items-center gap-2 bg-primary text-white px-4 py-2 rounded-lg text-sm hover:bg-primary-dark transition-colors"
        >
          + 매물 등록
        </Link>
      </div>

      <ListingFilters
        servers={servers}
        selectedServer={serverId}
        onServerChange={setServerId}
        searchQuery={search}
        onSearchChange={setSearch}
      />

      {isLoading ? (
        <Loading />
      ) : !data?.data.length ? (
        <EmptyState title="매물이 없습니다" description="첫 매물을 등록해보세요!" />
      ) : (
        <ListingGrid listings={data.data} />
      )}

      {/* Mobile FAB */}
      <Link
        href="/create"
        className="lg:hidden fixed right-4 bottom-20 bg-primary text-white w-14 h-14 rounded-full flex items-center justify-center text-2xl shadow-lg hover:bg-primary-dark z-40"
      >
        +
      </Link>
    </div>
  );
}
```

- [ ] **Step 2: Verify page renders**

```bash
cd web && npm run dev
```

Expected: Home page shows filter chips + grid layout (3 cols on desktop, 1 on mobile). If no backend is running, shows empty state.

- [ ] **Step 3: Commit**

```bash
git add web/app/page.tsx
git commit -m "feat: implement home page with listing grid, filters, and search"
```

---

### Task 12: Listing Detail Page

**Files:**
- Create: `web/components/listing/listing-info.tsx`
- Create: `web/app/listings/[id]/page.tsx`

- [ ] **Step 1: Create listing info sections**

```tsx
// web/components/listing/listing-info.tsx
import type { Author } from "@/lib/types";

export function InfoRow({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex py-2">
      <span className="w-28 text-text-secondary flex-shrink-0">{label}</span>
      <span>{value}</span>
    </div>
  );
}

export function AuthorSection({ author }: { author: Author }) {
  return (
    <div className="flex items-center gap-3">
      <div className="w-10 h-10 rounded-full bg-border flex items-center justify-center font-bold text-text-secondary">
        {author.nickname?.[0] ?? "?"}
      </div>
      <div>
        <div className="font-semibold">{author.nickname}</div>
        <div className="text-sm text-text-secondary">
          거래 {author.completedTradeCount ?? 0}회
          {author.trustBadge && ` · ${author.trustBadge}`}
        </div>
      </div>
    </div>
  );
}

const TRADE_METHOD_LABELS: Record<string, string> = {
  in_game: "인게임",
  offline_pc_bang: "PC방/오프라인",
  either: "무관",
};

export function tradeMethodLabel(method?: string): string {
  return TRADE_METHOD_LABELS[method ?? ""] ?? method ?? "";
}
```

- [ ] **Step 2: Create detail page**

```tsx
// web/app/listings/[id]/page.tsx
"use client";

import { use } from "react";
import { useRouter } from "next/navigation";
import { useListing, useToggleFavorite } from "@/lib/hooks/use-listings";
import { useCreateChat } from "@/lib/hooks/use-chats";
import { TypeBadge, Badge } from "@/components/ui/badge";
import { AuthorSection, InfoRow, tradeMethodLabel } from "@/components/listing/listing-info";
import { Loading } from "@/components/ui/loading";
import { formatPrice, statusLabel, statusColor } from "@/lib/utils";

export default function ListingDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const router = useRouter();
  const { data: listing, isLoading } = useListing(id);
  const toggleFav = useToggleFavorite();
  const createChat = useCreateChat();

  if (isLoading) return <Loading />;
  if (!listing) return <div className="p-6 text-center text-text-secondary">매물을 찾을 수 없습니다</div>;

  const l = listing;
  const actions = l.availableActions ?? [];

  const handleChat = async () => {
    try {
      const chat = await createChat.mutateAsync(l.listingId);
      router.push(`/chats/${chat.chatRoomId}`);
    } catch {
      alert("채팅을 시작할 수 없습니다");
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-4 lg:p-6">
      {/* Header badges */}
      <div className="flex items-center gap-2 mb-4">
        <TypeBadge type={l.listingType} />
        <Badge label={statusLabel(l.status)} color={statusColor(l.status)} />
        {l.tradeMethod && (
          <span className="ml-auto text-sm text-text-secondary">{tradeMethodLabel(l.tradeMethod)}</span>
        )}
      </div>

      {/* Title */}
      <h1 className="text-2xl font-bold mb-3">{l.title}</h1>

      {/* Item info */}
      <div className="flex items-center gap-2 text-lg mb-4">
        {l.iconUrl && (
          <img
            src={`${process.env.NEXT_PUBLIC_API_URL?.replace("/api/v1", "") ?? "http://localhost:8080"}${l.iconUrl}`}
            alt=""
            className="w-8 h-8"
          />
        )}
        <span>{l.itemName}</span>
        {l.enhancementLevel != null && (
          <span className="text-primary font-bold">+{l.enhancementLevel}</span>
        )}
      </div>
      {l.optionsText && <p className="text-text-secondary mb-4">{l.optionsText}</p>}

      {/* Price */}
      <div className="text-3xl font-bold mb-1">{formatPrice(l.priceAmount)}원</div>
      {l.priceType === "negotiable" && <p className="text-text-secondary mb-4">협상 가능</p>}

      <hr className="border-border my-6" />

      {/* Description */}
      <p className="leading-relaxed whitespace-pre-wrap">{l.description}</p>

      <hr className="border-border my-6" />

      {/* Trade info */}
      <InfoRow label="거래 방식" value={tradeMethodLabel(l.tradeMethod)} />
      {l.preferredMeetingAreaText && <InfoRow label="접선 장소" value={l.preferredMeetingAreaText} />}
      {l.availableTimeText && <InfoRow label="거래 가능 시간" value={l.availableTimeText} />}
      <InfoRow label="수량" value={`${l.quantity}개`} />

      <hr className="border-border my-6" />

      {/* Author */}
      {l.author && <AuthorSection author={l.author} />}

      {/* Action bar */}
      {actions.length > 0 && (
        <div className="sticky bottom-0 lg:relative bg-white border-t border-border mt-8 py-4 flex items-center gap-3">
          {actions.includes("favorite") && (
            <button
              onClick={() => toggleFav.mutate({ id: l.listingId, isFavorited: l.isFavorited ?? false })}
              className="p-3 border border-border rounded-lg hover:bg-surface transition-colors"
            >
              {l.isFavorited ? "❤️" : "🤍"}
            </button>
          )}
          {actions.includes("start_chat") && (
            <button
              onClick={handleChat}
              disabled={createChat.isPending}
              className="flex-1 bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition-colors disabled:opacity-50"
            >
              {createChat.isPending ? "연결 중..." : "채팅하기"}
            </button>
          )}
        </div>
      )}
    </div>
  );
}
```

- [ ] **Step 3: Verify page renders**

```bash
cd web && npm run dev
```

Visit http://localhost:3000/listings/test-id — should show loading then error (no backend).

- [ ] **Step 4: Commit**

```bash
git add web/components/listing/listing-info.tsx web/app/listings/
git commit -m "feat: implement listing detail page with responsive layout"
```

---

### Task 13: Login Page

**Files:**
- Create: `web/app/login/page.tsx`

- [ ] **Step 1: Create login page**

```tsx
// web/app/login/page.tsx
"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { apiClient } from "@/lib/api-client";

export default function LoginPage() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const handleDevLogin = async () => {
    setLoading(true);
    try {
      await apiClient.login("google", `dev_user_${Date.now()}`);
      router.push("/");
    } catch (e) {
      alert(`로그인 실패: ${e}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="w-full max-w-sm text-center">
        <div className="text-6xl mb-4">🏪</div>
        <h1 className="text-3xl font-bold text-primary mb-2">기란장터</h1>
        <p className="text-text-secondary mb-10">리니지 클래식 거래 플랫폼</p>

        {/* Google OAuth — TODO: integrate NextAuth.js */}
        <button
          disabled={loading}
          className="w-full flex items-center justify-center gap-3 bg-white border border-border rounded-lg py-3 px-4 text-sm hover:bg-surface transition-colors disabled:opacity-50"
        >
          <span className="text-lg">G</span>
          Google로 시작하기
        </button>

        {/* Dev login */}
        <button
          onClick={handleDevLogin}
          disabled={loading}
          className="w-full mt-3 flex items-center justify-center gap-3 border border-border rounded-lg py-3 px-4 text-sm text-text-secondary hover:bg-surface transition-colors disabled:opacity-50"
        >
          🛠️ 개발자 로그인 (테스트)
        </button>

        <button
          onClick={() => router.push("/")}
          className="mt-4 text-sm text-text-secondary hover:text-primary"
        >
          둘러보기
        </button>
      </div>
    </div>
  );
}
```

- [ ] **Step 2: Commit**

```bash
git add web/app/login/
git commit -m "feat: implement login page with dev login and Google OAuth placeholder"
```

---

### Task 14: Listing Create Page

**Files:**
- Create: `web/app/create/page.tsx`

- [ ] **Step 1: Create listing form page**

```tsx
// web/app/create/page.tsx
"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useCreateListing } from "@/lib/hooks/use-listings";

export default function CreateListingPage() {
  const router = useRouter();
  const createListing = useCreateListing();
  const { data: servers = [] } = useQuery({ queryKey: ["servers"], queryFn: () => apiClient.getServers() });
  const { data: categories = [] } = useQuery({ queryKey: ["categories"], queryFn: () => apiClient.getCategories() });

  const [form, setForm] = useState({
    listingType: "sell",
    serverId: "",
    categoryId: "",
    itemName: "",
    title: "",
    description: "",
    priceType: "fixed",
    priceAmount: "",
    enhancementLevel: "",
    tradeMethod: "either",
  });

  const update = (field: string, value: string) => setForm((f) => ({ ...f, [field]: value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const data: Record<string, unknown> = {
      ...form,
      quantity: 1,
      priceAmount: form.priceType !== "offer" && form.priceAmount ? Number(form.priceAmount) : undefined,
      enhancementLevel: form.enhancementLevel ? Number(form.enhancementLevel) : undefined,
    };
    try {
      await createListing.mutateAsync(data as never);
      router.push("/");
    } catch (e) {
      alert(`등록 실패: ${JSON.stringify(e)}`);
    }
  };

  const inputClass = "w-full border border-border rounded-lg px-4 py-3 text-sm outline-none focus:border-primary focus:ring-1 focus:ring-primary bg-white";
  const labelClass = "block text-sm font-medium text-text-primary mb-1";

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-6">매물 등록</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Type toggle */}
        <div className="flex rounded-lg border border-border overflow-hidden">
          {["sell", "buy"].map((t) => (
            <button
              key={t}
              type="button"
              onClick={() => update("listingType", t)}
              className={`flex-1 py-2.5 text-sm font-medium transition-colors ${
                form.listingType === t ? "bg-primary text-white" : "bg-white text-text-secondary"
              }`}
            >
              {t === "sell" ? "판매" : "구매"}
            </button>
          ))}
        </div>

        <div>
          <label className={labelClass}>서버 *</label>
          <select className={inputClass} value={form.serverId} onChange={(e) => update("serverId", e.target.value)} required>
            <option value="">선택</option>
            {servers.map((s) => <option key={s.serverId} value={s.serverId}>{s.serverName}</option>)}
          </select>
        </div>

        <div>
          <label className={labelClass}>카테고리 *</label>
          <select className={inputClass} value={form.categoryId} onChange={(e) => update("categoryId", e.target.value)} required>
            <option value="">선택</option>
            {categories.filter((c) => !c.parentId).map((c) => <option key={c.categoryId} value={c.categoryId}>{c.categoryName}</option>)}
          </select>
        </div>

        <div>
          <label className={labelClass}>아이템명 *</label>
          <input className={inputClass} value={form.itemName} onChange={(e) => update("itemName", e.target.value)} required />
        </div>

        <div>
          <label className={labelClass}>제목 *</label>
          <input className={inputClass} value={form.title} onChange={(e) => update("title", e.target.value)} placeholder="예: 집행검 +9 급처합니다" required minLength={2} />
        </div>

        <div>
          <label className={labelClass}>설명 *</label>
          <textarea className={`${inputClass} h-28`} value={form.description} onChange={(e) => update("description", e.target.value)} placeholder="아이템 상세 설명" required minLength={10} />
        </div>

        <div className="grid grid-cols-2 gap-3">
          <div>
            <label className={labelClass}>가격 유형</label>
            <select className={inputClass} value={form.priceType} onChange={(e) => update("priceType", e.target.value)}>
              <option value="fixed">고정가</option>
              <option value="negotiable">협상가능</option>
              <option value="offer">제안받음</option>
            </select>
          </div>
          <div>
            <label className={labelClass}>가격 (원)</label>
            <input className={inputClass} type="number" value={form.priceAmount} onChange={(e) => update("priceAmount", e.target.value)} disabled={form.priceType === "offer"} />
          </div>
        </div>

        <div>
          <label className={labelClass}>강화 수치 (선택)</label>
          <input className={inputClass} type="number" value={form.enhancementLevel} onChange={(e) => update("enhancementLevel", e.target.value)} />
        </div>

        <div>
          <label className={labelClass}>거래 방식</label>
          <select className={inputClass} value={form.tradeMethod} onChange={(e) => update("tradeMethod", e.target.value)}>
            <option value="in_game">인게임</option>
            <option value="offline_pc_bang">PC방/오프라인</option>
            <option value="either">무관</option>
          </select>
        </div>

        <button
          type="submit"
          disabled={createListing.isPending}
          className="w-full bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition-colors disabled:opacity-50"
        >
          {createListing.isPending ? "등록 중..." : "등록하기"}
        </button>
      </form>
    </div>
  );
}
```

- [ ] **Step 2: Commit**

```bash
git add web/app/create/
git commit -m "feat: implement listing create form page"
```

---

## Chunk 3: Communication

### Task 15: Chat Components

**Files:**
- Create: `web/components/chat/chat-list-item.tsx`
- Create: `web/components/chat/chat-message.tsx`
- Create: `web/components/chat/chat-input.tsx`
- Create: `web/components/chat/chat-panel.tsx`

- [ ] **Step 1: Create chat components**

```tsx
// web/components/chat/chat-list-item.tsx
import type { ChatRoom } from "@/lib/types";
import { formatTimeAgo } from "@/lib/utils";

interface ChatListItemProps {
  chat: ChatRoom;
  isActive?: boolean;
  onClick: () => void;
}

export function ChatListItem({ chat, isActive, onClick }: ChatListItemProps) {
  return (
    <button
      onClick={onClick}
      className={`w-full text-left p-3 border-b border-border transition-colors ${
        isActive ? "bg-blue-50 border-l-2 border-l-primary" : "hover:bg-surface"
      }`}
    >
      <div className="flex items-center justify-between mb-1">
        <span className="font-semibold text-sm">{chat.counterparty.nickname}</span>
        <span className="text-xs text-text-secondary">
          {chat.lastMessage ? formatTimeAgo(chat.lastMessage.sentAt) : ""}
        </span>
      </div>
      <div className="text-xs text-text-secondary truncate">{chat.listingTitle}</div>
      {chat.lastMessage && (
        <div className="text-sm text-text-secondary truncate mt-1">{chat.lastMessage.bodyText}</div>
      )}
      {chat.unreadCount > 0 && (
        <span className="inline-block mt-1 bg-primary text-white text-xs px-1.5 py-0.5 rounded-full">
          {chat.unreadCount}
        </span>
      )}
    </button>
  );
}
```

```tsx
// web/components/chat/chat-message.tsx
import type { Message } from "@/lib/types";

interface ChatMessageProps {
  message: Message;
  isMine: boolean;
}

export function ChatMessage({ message, isMine }: ChatMessageProps) {
  if (message.messageType === "system") {
    return (
      <div className="flex justify-center my-2">
        <span className="bg-border text-text-secondary text-xs px-3 py-1.5 rounded-full">
          {message.bodyText}
        </span>
      </div>
    );
  }

  return (
    <div className={`flex mb-1 ${isMine ? "justify-end" : "justify-start"}`}>
      <div
        className={`max-w-[70%] px-4 py-2.5 rounded-2xl text-sm ${
          isMine
            ? "bg-primary text-white rounded-br-sm"
            : "bg-white border border-border text-text-primary rounded-bl-sm"
        }`}
      >
        {message.bodyText}
      </div>
    </div>
  );
}
```

```tsx
// web/components/chat/chat-input.tsx
"use client";

import { useState, type FormEvent } from "react";

interface ChatInputProps {
  onSend: (text: string) => void;
  disabled?: boolean;
}

export function ChatInput({ onSend, disabled }: ChatInputProps) {
  const [text, setText] = useState("");

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    const trimmed = text.trim();
    if (!trimmed) return;
    onSend(trimmed);
    setText("");
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2 p-3 bg-white border-t border-border">
      <input
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="메시지를 입력하세요"
        className="flex-1 border border-border rounded-full px-4 py-2 text-sm outline-none focus:border-primary"
        disabled={disabled}
      />
      <button
        type="submit"
        disabled={disabled || !text.trim()}
        className="bg-primary text-white px-4 py-2 rounded-full text-sm font-medium disabled:opacity-50"
      >
        전송
      </button>
    </form>
  );
}
```

```tsx
// web/components/chat/chat-panel.tsx
"use client";

import { useEffect, useRef } from "react";
import type { ChatRoom, Message } from "@/lib/types";
import { ChatListItem } from "./chat-list-item";
import { ChatMessage } from "./chat-message";
import { ChatInput } from "./chat-input";

interface ChatPanelProps {
  chats: ChatRoom[];
  activeChatId: string | null;
  messages: Message[];
  myUserId: string | null;
  onSelectChat: (chatId: string) => void;
  onSendMessage: (text: string) => void;
}

export function ChatPanel({ chats, activeChatId, messages, myUserId, onSelectChat, onSendMessage }: ChatPanelProps) {
  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  const activeChat = chats.find((c) => c.chatRoomId === activeChatId);

  return (
    <div className="flex h-[calc(100vh-64px)] lg:h-[calc(100vh-0px)]">
      {/* Chat list */}
      <div className="w-72 border-r border-border bg-white overflow-y-auto flex-shrink-0">
        <div className="p-3 border-b border-border font-semibold">채팅 목록</div>
        {chats.map((c) => (
          <ChatListItem
            key={c.chatRoomId}
            chat={c}
            isActive={c.chatRoomId === activeChatId}
            onClick={() => onSelectChat(c.chatRoomId)}
          />
        ))}
      </div>

      {/* Messages */}
      <div className="flex-1 flex flex-col bg-surface">
        {activeChat ? (
          <>
            <div className="px-4 py-3 bg-white border-b border-border font-semibold text-sm">
              {activeChat.counterparty.nickname} · {activeChat.listingTitle}
            </div>
            <div className="flex-1 overflow-y-auto p-4">
              {messages.map((m) => (
                <ChatMessage key={m.messageId} message={m} isMine={m.senderUserId === myUserId} />
              ))}
              <div ref={bottomRef} />
            </div>
            <ChatInput onSend={onSendMessage} />
          </>
        ) : (
          <div className="flex-1 flex items-center justify-center text-text-secondary">
            채팅을 선택해주세요
          </div>
        )}
      </div>
    </div>
  );
}
```

- [ ] **Step 2: Commit**

```bash
git add web/components/chat/
git commit -m "feat: add chat components (list, message, input, desktop panel)"
```

---

### Task 16: Chat Pages

**Files:**
- Create: `web/app/chats/page.tsx`
- Create: `web/app/chats/[id]/page.tsx`

- [ ] **Step 1: Create chat list page (desktop=panel, mobile=list)**

```tsx
// web/app/chats/page.tsx
"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useChats, useMessages, useSendMessage } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { ChatPanel } from "@/components/chat/chat-panel";
import { ChatListItem } from "@/components/chat/chat-list-item";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function ChatsPage() {
  const router = useRouter();
  const { data: me } = useMe();
  const { data: chatsData, isLoading } = useChats();
  const [activeChatId, setActiveChatId] = useState<string | null>(null);
  const { data: messagesData } = useMessages(activeChatId ?? "");
  const sendMessage = useSendMessage();

  const chats = chatsData?.data ?? [];
  const messages = messagesData?.data ? [...messagesData.data].reverse() : [];

  if (isLoading) return <Loading />;
  if (!chats.length) return <EmptyState title="채팅이 없습니다" description="매물에서 채팅을 시작해보세요" />;

  // Desktop: split panel
  return (
    <>
      {/* Desktop split panel */}
      <div className="hidden lg:block">
        <ChatPanel
          chats={chats}
          activeChatId={activeChatId}
          messages={messages}
          myUserId={me?.userId ?? null}
          onSelectChat={setActiveChatId}
          onSendMessage={(text) => {
            if (activeChatId) sendMessage.mutate({ chatId: activeChatId, text });
          }}
        />
      </div>

      {/* Mobile: list only */}
      <div className="lg:hidden">
        {chats.map((c) => (
          <ChatListItem
            key={c.chatRoomId}
            chat={c}
            onClick={() => router.push(`/chats/${c.chatRoomId}`)}
          />
        ))}
      </div>
    </>
  );
}
```

- [ ] **Step 2: Create mobile chat detail page**

```tsx
// web/app/chats/[id]/page.tsx
"use client";

import { use, useEffect, useRef } from "react";
import { useMessages, useSendMessage } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { ChatMessage } from "@/components/chat/chat-message";
import { ChatInput } from "@/components/chat/chat-input";
import { Loading } from "@/components/ui/loading";

export default function ChatDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const { data: me } = useMe();
  const { data, isLoading } = useMessages(id);
  const sendMessage = useSendMessage();
  const bottomRef = useRef<HTMLDivElement>(null);

  const messages = data?.data ? [...data.data].reverse() : [];

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  if (isLoading) return <Loading />;

  return (
    <div className="flex flex-col h-[calc(100vh-120px)]">
      <div className="flex-1 overflow-y-auto p-4">
        {messages.map((m) => (
          <ChatMessage key={m.messageId} message={m} isMine={m.senderUserId === me?.userId} />
        ))}
        <div ref={bottomRef} />
      </div>
      <ChatInput onSend={(text) => sendMessage.mutate({ chatId: id, text })} />
    </div>
  );
}
```

- [ ] **Step 3: Commit**

```bash
git add web/app/chats/
git commit -m "feat: implement chat pages (desktop split panel + mobile full screen)"
```

---

### Task 17: Form Modals (Reservation, Review, Report)

**Files:**
- Create: `web/components/forms/reservation-modal.tsx`
- Create: `web/components/forms/review-modal.tsx`
- Create: `web/components/forms/report-modal.tsx`

- [ ] **Step 1: Create reservation modal**

```tsx
// web/components/forms/reservation-modal.tsx
"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";

interface ReservationModalProps {
  open: boolean;
  onClose: () => void;
  chatId: string;
  onCreated: () => void;
}

export function ReservationModal({ open, onClose, chatId, onCreated }: ReservationModalProps) {
  const [form, setForm] = useState({ scheduledDate: "", scheduledTime: "", meetingType: "in_game", meetingPointText: "", notes: "" });
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setSubmitting(true);
    try {
      await apiClient.createReservation(chatId, {
        scheduledAt: `${form.scheduledDate}T${form.scheduledTime}:00Z`,
        meetingType: form.meetingType,
        meetingPointText: form.meetingPointText || undefined,
        notes: form.notes || undefined,
      });
      onCreated();
      onClose();
    } catch (e) {
      alert(`예약 실패: ${JSON.stringify(e)}`);
    } finally {
      setSubmitting(false);
    }
  };

  const inputClass = "w-full border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-primary";

  return (
    <Modal open={open} onClose={onClose} title="예약 제안">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="grid grid-cols-2 gap-3">
          <input type="date" className={inputClass} value={form.scheduledDate} onChange={(e) => setForm({ ...form, scheduledDate: e.target.value })} required />
          <input type="time" className={inputClass} value={form.scheduledTime} onChange={(e) => setForm({ ...form, scheduledTime: e.target.value })} required />
        </div>
        <select className={inputClass} value={form.meetingType} onChange={(e) => setForm({ ...form, meetingType: e.target.value })}>
          <option value="in_game">인게임</option>
          <option value="offline_pc_bang">PC방/오프라인</option>
        </select>
        <input className={inputClass} placeholder="접선 장소" value={form.meetingPointText} onChange={(e) => setForm({ ...form, meetingPointText: e.target.value })} />
        <textarea className={`${inputClass} h-20`} placeholder="메모 (선택)" value={form.notes} onChange={(e) => setForm({ ...form, notes: e.target.value })} />
        <button type="submit" disabled={submitting} className="w-full bg-primary text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제안 중..." : "예약 제안"}
        </button>
      </form>
    </Modal>
  );
}
```

- [ ] **Step 2: Create review modal**

```tsx
// web/components/forms/review-modal.tsx
"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";

interface ReviewModalProps {
  open: boolean;
  onClose: () => void;
  completionId: string;
  onCreated: () => void;
}

export function ReviewModal({ open, onClose, completionId, onCreated }: ReviewModalProps) {
  const [rating, setRating] = useState<"positive" | "negative" | null>(null);
  const [comment, setComment] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!rating) return;
    setSubmitting(true);
    try {
      await apiClient.createReview(completionId, { rating, comment: comment || undefined });
      onCreated();
      onClose();
    } catch (e) {
      alert(`리뷰 실패: ${JSON.stringify(e)}`);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="거래 리뷰">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="flex gap-3">
          {(["positive", "negative"] as const).map((r) => (
            <button
              key={r}
              type="button"
              onClick={() => setRating(r)}
              className={`flex-1 py-3 rounded-lg text-sm font-medium border transition-colors ${
                rating === r
                  ? r === "positive" ? "bg-green-50 border-secondary text-secondary" : "bg-red-50 border-error text-error"
                  : "border-border text-text-secondary hover:bg-surface"
              }`}
            >
              {r === "positive" ? "👍 좋았어요" : "👎 아쉬웠어요"}
            </button>
          ))}
        </div>
        <textarea
          className="w-full border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-primary h-24"
          placeholder="한줄 코멘트 (선택)"
          value={comment}
          onChange={(e) => setComment(e.target.value)}
        />
        <button type="submit" disabled={submitting || !rating} className="w-full bg-primary text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제출 중..." : "리뷰 제출"}
        </button>
      </form>
    </Modal>
  );
}
```

- [ ] **Step 3: Create report modal**

```tsx
// web/components/forms/report-modal.tsx
"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";

const REPORT_REASONS = [
  { value: "scam", label: "사기 의심" },
  { value: "fake_listing", label: "허위 매물" },
  { value: "abuse", label: "욕설/비하" },
  { value: "spam", label: "도배/스팸" },
  { value: "other", label: "기타" },
];

interface ReportModalProps {
  open: boolean;
  onClose: () => void;
  targetType: string;
  targetId: string;
}

export function ReportModal({ open, onClose, targetType, targetId }: ReportModalProps) {
  const [reason, setReason] = useState("");
  const [description, setDescription] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!reason) return;
    setSubmitting(true);
    try {
      await apiClient.createReport({ targetType, targetId, reasonCode: reason, description: description || undefined });
      onClose();
      alert("신고가 접수되었습니다");
    } catch (e) {
      alert(`신고 실패: ${JSON.stringify(e)}`);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="신고하기">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          {REPORT_REASONS.map((r) => (
            <label key={r.value} className={`flex items-center gap-3 p-3 rounded-lg border cursor-pointer transition-colors ${reason === r.value ? "border-primary bg-blue-50" : "border-border hover:bg-surface"}`}>
              <input type="radio" name="reason" value={r.value} checked={reason === r.value} onChange={() => setReason(r.value)} className="accent-primary" />
              {r.label}
            </label>
          ))}
        </div>
        <textarea
          className="w-full border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-primary h-20"
          placeholder="상세 설명 (선택)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button type="submit" disabled={submitting || !reason} className="w-full bg-error text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제출 중..." : "신고하기"}
        </button>
      </form>
    </Modal>
  );
}
```

- [ ] **Step 4: Commit**

```bash
git add web/components/forms/
git commit -m "feat: add form modals (reservation, review, report)"
```

---

## Chunk 4: Profile + Deployment

### Task 18: Profile Pages

**Files:**
- Create: `web/app/profile/page.tsx`
- Create: `web/app/profile/listings/page.tsx`
- Create: `web/app/profile/trades/page.tsx`
- Create: `web/app/notifications/page.tsx`

- [ ] **Step 1: Create profile page**

```tsx
// web/app/profile/page.tsx
"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { useMe } from "@/lib/hooks/use-profile";
import { apiClient } from "@/lib/api-client";
import { Loading } from "@/components/ui/loading";

export default function ProfilePage() {
  const router = useRouter();
  const { data: me, isLoading } = useMe();

  if (isLoading) return <Loading />;
  if (!me) {
    return (
      <div className="p-6 text-center">
        <p className="text-text-secondary mb-4">로그인이 필요합니다</p>
        <Link href="/login" className="text-primary font-medium">로그인하기</Link>
      </div>
    );
  }

  const handleLogout = () => {
    apiClient.clearTokens();
    router.push("/");
    router.refresh();
  };

  const menuItems = [
    { href: "/profile/listings", icon: "📦", label: "내 매물" },
    { href: "/profile/trades", icon: "🤝", label: "내 거래" },
    { href: "/notifications", icon: "🔔", label: "알림" },
  ];

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      {/* User card */}
      <div className="bg-white rounded-xl border border-border p-5 mb-4">
        <div className="flex items-center gap-4">
          <div className="w-16 h-16 rounded-full bg-surface flex items-center justify-center text-2xl font-bold text-text-secondary">
            {me.nickname[0]}
          </div>
          <div>
            <h2 className="text-xl font-bold">{me.nickname}</h2>
            {me.introduction && <p className="text-sm text-text-secondary mt-1">{me.introduction}</p>}
          </div>
        </div>
        <div className="grid grid-cols-3 gap-4 mt-5 text-center">
          <div><div className="text-xl font-bold">{me.completedTradeCount}</div><div className="text-xs text-text-secondary">거래</div></div>
          <div><div className="text-xl font-bold">{me.positiveReviewCount}</div><div className="text-xs text-text-secondary">좋은 리뷰</div></div>
          <div><div className="text-xl font-bold">{me.trustBadge ?? "-"}</div><div className="text-xs text-text-secondary">신뢰등급</div></div>
        </div>
      </div>

      {/* Menu */}
      <div className="bg-white rounded-xl border border-border overflow-hidden">
        {menuItems.map((item) => (
          <Link key={item.href} href={item.href} className="flex items-center gap-3 px-5 py-4 border-b border-border last:border-0 hover:bg-surface transition-colors">
            <span>{item.icon}</span>
            <span className="flex-1">{item.label}</span>
            <span className="text-text-secondary">›</span>
          </Link>
        ))}
      </div>

      <button onClick={handleLogout} className="w-full mt-4 py-3 text-sm text-error border border-border rounded-xl hover:bg-red-50 transition-colors">
        로그아웃
      </button>
    </div>
  );
}
```

- [ ] **Step 2: Create my listings page**

```tsx
// web/app/profile/listings/page.tsx
"use client";

import { useMyListings } from "@/lib/hooks/use-profile";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function MyListingsPage() {
  const { data, isLoading } = useMyListings();

  if (isLoading) return <Loading />;
  if (!data?.data.length) return <EmptyState title="등록한 매물이 없습니다" />;

  return (
    <div className="p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4">내 매물</h1>
      <ListingGrid listings={data.data} />
    </div>
  );
}
```

- [ ] **Step 3: Create my trades and notifications pages**

```tsx
// web/app/profile/trades/page.tsx
"use client";

import { useMyTrades } from "@/lib/hooks/use-profile";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function MyTradesPage() {
  const { data, isLoading } = useMyTrades();

  if (isLoading) return <Loading />;
  if (!data?.data.length) return <EmptyState title="거래 내역이 없습니다" />;

  return (
    <div className="p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4">내 거래</h1>
      <ListingGrid listings={data.data} />
    </div>
  );
}
```

```tsx
// web/app/notifications/page.tsx
"use client";

import { useNotifications, useMarkNotificationsRead } from "@/lib/hooks/use-profile";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { formatTimeAgo } from "@/lib/utils";

export default function NotificationsPage() {
  const { data, isLoading } = useNotifications();
  const markRead = useMarkNotificationsRead();

  const notifications = data?.data ?? [];

  if (isLoading) return <Loading />;
  if (!notifications.length) return <EmptyState title="알림이 없습니다" icon="🔔" />;

  const unreadIds = notifications.filter((n) => !n.readAt).map((n) => n.notificationId);

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold">알림</h1>
        {unreadIds.length > 0 && (
          <button onClick={() => markRead.mutate(unreadIds)} className="text-sm text-primary">모두 읽음</button>
        )}
      </div>
      <div className="bg-white rounded-xl border border-border overflow-hidden">
        {notifications.map((n) => (
          <div key={n.notificationId} className={`px-5 py-4 border-b border-border last:border-0 ${!n.readAt ? "bg-blue-50" : ""}`}>
            <p className="text-sm">{n.message}</p>
            <p className="text-xs text-text-secondary mt-1">{formatTimeAgo(n.createdAt)}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
```

- [ ] **Step 4: Commit**

```bash
git add web/app/profile/ web/app/notifications/
git commit -m "feat: implement profile, my listings, my trades, and notifications pages"
```

---

### Task 19: SSE Hook for Real-time Chat

**Files:**
- Create: `web/lib/hooks/use-sse.ts`

- [ ] **Step 1: Create SSE hook**

```typescript
// web/lib/hooks/use-sse.ts
"use client";

import { useEffect, useRef } from "react";
import { useQueryClient } from "@tanstack/react-query";

const API_BASE = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";

export function useSSE() {
  const qc = useQueryClient();
  const esRef = useRef<EventSource | null>(null);

  useEffect(() => {
    const token = typeof window !== "undefined" ? localStorage.getItem("accessToken") : null;
    if (!token) return;

    const es = new EventSource(`${API_BASE}/sse/connect?token=${token}`);
    esRef.current = es;

    es.addEventListener("new_message", (e) => {
      const data = JSON.parse(e.data);
      qc.invalidateQueries({ queryKey: ["messages", data.chatRoomId] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    });

    es.addEventListener("status_change", () => {
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    });

    es.onerror = () => {
      es.close();
      // Auto-reconnect after 5 seconds
      setTimeout(() => {
        esRef.current = null;
      }, 5_000);
    };

    return () => {
      es.close();
      esRef.current = null;
    };
  }, [qc]);
}
```

- [ ] **Step 2: Add SSE hook to layout provider**

Update `web/lib/providers.tsx` to include SSE initialization:

```tsx
// Add to web/lib/providers.tsx
import { useSSE } from "./hooks/use-sse";

function SSEInitializer() {
  useSSE();
  return null;
}

// Add <SSEInitializer /> inside QueryClientProvider in Providers component
```

- [ ] **Step 3: Commit**

```bash
git add web/lib/hooks/use-sse.ts web/lib/providers.tsx
git commit -m "feat: add SSE hook for real-time chat message updates"
```

---

### Task 20: Vitest Configuration

**Files:**
- Create: `web/vitest.config.ts`

- [ ] **Step 1: Configure Vitest**

```typescript
// web/vitest.config.ts
import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import path from "path";

export default defineConfig({
  plugins: [react()],
  test: {
    environment: "jsdom",
    setupFiles: ["./vitest.setup.ts"],
    globals: true,
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "."),
    },
  },
});
```

```typescript
// web/vitest.setup.ts
import "@testing-library/jest-dom/vitest";
```

- [ ] **Step 2: Run all tests**

```bash
cd web && npx vitest run
```

Expected: All tests pass.

- [ ] **Step 3: Commit**

```bash
git add web/vitest.config.ts web/vitest.setup.ts
git commit -m "feat: configure Vitest with jsdom and React testing library"
```

---

### Task 21: Docker + Caddy Deployment Config

**Files:**
- Create: `web/Dockerfile`
- Modify: `docker-compose.yml` (if exists) or create

- [ ] **Step 1: Create Dockerfile for Next.js**

```dockerfile
# web/Dockerfile
FROM node:20-alpine AS deps
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci --production=false

FROM node:20-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
COPY ../shared ../shared
RUN npm run build

FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
EXPOSE 3000
CMD ["node", "server.js"]
```

- [ ] **Step 2: Update next.config.ts for standalone output**

Add `output: "standalone"` to `web/next.config.ts`.

- [ ] **Step 3: Commit**

```bash
git add web/Dockerfile web/next.config.ts
git commit -m "feat: add Docker deployment config for Next.js"
```

---

### Task 22: End-to-End Verification

- [ ] **Step 1: Start backend + frontend and verify all pages**

```bash
# Terminal 1: Backend
cd backend && go run cmd/server/main.go

# Terminal 2: Web frontend
cd web && npm run dev
```

Verify each page:
- [ ] `http://localhost:3000/` — Listing grid loads with filters
- [ ] `http://localhost:3000/listings/{id}` — Detail page renders
- [ ] `http://localhost:3000/create` — Form works
- [ ] `http://localhost:3000/chats` — Desktop split panel, mobile list
- [ ] `http://localhost:3000/login` — Dev login works
- [ ] `http://localhost:3000/profile` — Profile renders after login
- [ ] `http://localhost:3000/notifications` — Notifications list
- [ ] Responsive: resize browser to verify mobile ↔ desktop layout switch

- [ ] **Step 2: Run all tests**

```bash
cd web && npx vitest run
```

Expected: All tests pass.

- [ ] **Step 3: Build production bundle**

```bash
cd web && npm run build
```

Expected: Build succeeds with no errors.

- [ ] **Step 4: Final commit**

```bash
git add -A
git commit -m "feat: complete Next.js web frontend — all Flutter features replicated"
```
