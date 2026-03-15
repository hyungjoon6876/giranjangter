import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ChatListItem } from "@/components/chat/chat-list-item";
import type { ChatRoom } from "@/lib/types";

afterEach(() => cleanup());

const baseChat: ChatRoom = {
  chatRoomId: "c1",
  listingId: "l1",
  listingTitle: "집행검 +7 팝니다",
  listingStatus: "available",
  counterparty: { userId: "u2", nickname: "흑기사" },
  chatStatus: "active",
  unreadCount: 0,
  updatedAt: new Date().toISOString(),
};

describe("ChatListItem", () => {
  it("sets aria-current=true when active", () => {
    render(<ChatListItem chat={baseChat} isActive={true} onClick={vi.fn()} />);
    const btn = screen.getByRole("button");
    expect(btn.getAttribute("aria-current")).toBe("true");
  });

  it("omits aria-current when not active", () => {
    render(<ChatListItem chat={baseChat} isActive={false} onClick={vi.fn()} />);
    const btn = screen.getByRole("button");
    expect(btn.getAttribute("aria-current")).toBeNull();
  });

  it("aria-label includes unread count when present", () => {
    const chatWithUnread: ChatRoom = { ...baseChat, unreadCount: 5 };
    render(<ChatListItem chat={chatWithUnread} onClick={vi.fn()} />);
    const btn = screen.getByRole("button");
    expect(btn.getAttribute("aria-label")).toBe("흑기사 — 5개 안 읽은 메시지");
  });

  it("aria-label is just nickname when no unread", () => {
    render(<ChatListItem chat={baseChat} onClick={vi.fn()} />);
    const btn = screen.getByRole("button");
    expect(btn.getAttribute("aria-label")).toBe("흑기사");
  });

  it("unread badge has aria-hidden", () => {
    const chatWithUnread: ChatRoom = { ...baseChat, unreadCount: 3 };
    const { container } = render(<ChatListItem chat={chatWithUnread} onClick={vi.fn()} />);
    const badge = container.querySelector("[aria-hidden='true']");
    expect(badge).not.toBeNull();
    expect(badge?.textContent).toBe("3");
  });
});
