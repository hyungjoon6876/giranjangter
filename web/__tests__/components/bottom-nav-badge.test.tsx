import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { createElement, type ReactNode } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import type { ChatRoom } from "@/lib/types";

// Mock next/navigation
vi.mock("next/navigation", () => ({
  usePathname: () => "/",
}));

// Mock api-client
vi.mock("@/lib/api-client", () => ({
  apiClient: { isLoggedIn: true },
  API_BASE: "http://localhost:8080",
}));

// Control chats data via a ref
let mockChatsData: { data: ChatRoom[] } | undefined;

vi.mock("@/lib/hooks/use-chats", () => ({
  useChats: () => ({ data: mockChatsData }),
}));

import { BottomNav } from "@/components/layout/bottom-nav";

afterEach(() => {
  cleanup();
  mockChatsData = undefined;
});

function renderWithQuery(ui: ReactNode) {
  const qc = new QueryClient({ defaultOptions: { queries: { retry: false } } });
  return render(createElement(QueryClientProvider, { client: qc }, ui));
}

const makeChat = (id: string, unreadCount: number): ChatRoom => ({
  chatRoomId: id,
  listingId: "l1",
  listingTitle: "집행검",
  listingStatus: "available",
  counterparty: { userId: "u2", nickname: "흑기사" },
  chatStatus: "active",
  unreadCount,
  updatedAt: new Date().toISOString(),
});

describe("BottomNav chat badge", () => {
  it("shows unread count badge on chat tab", () => {
    mockChatsData = { data: [makeChat("c1", 3), makeChat("c2", 2)] };
    renderWithQuery(createElement(BottomNav));
    const chatLink = screen.getByLabelText("채팅 — 읽지 않은 메시지 5개");
    expect(chatLink).toBeDefined();
    expect(screen.getByText("5")).toBeDefined();
  });

  it("hides badge when unread count is 0", () => {
    mockChatsData = { data: [makeChat("c1", 0)] };
    const { container } = renderWithQuery(createElement(BottomNav));
    const chatLink = screen.getByLabelText("채팅");
    expect(chatLink).toBeDefined();
    const badge = container.querySelector(".bg-gold.rounded-full");
    expect(badge).toBeNull();
  });

  it("shows 99+ for large unread counts", () => {
    mockChatsData = { data: [makeChat("c1", 100)] };
    renderWithQuery(createElement(BottomNav));
    expect(screen.getByText("99+")).toBeDefined();
  });

  it("aria-label updates dynamically with unread count", () => {
    mockChatsData = { data: [makeChat("c1", 7)] };
    renderWithQuery(createElement(BottomNav));
    const chatLink = screen.getByLabelText("채팅 — 읽지 않은 메시지 7개");
    expect(chatLink).toBeDefined();
  });

  it("non-chat tabs have no badge", () => {
    mockChatsData = { data: [makeChat("c1", 5)] };
    renderWithQuery(createElement(BottomNav));
    // Market tab should not have badge
    const marketLink = screen.getByLabelText("마켓");
    expect(marketLink).toBeDefined();
  });
});
