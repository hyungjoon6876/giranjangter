import { describe, it, expect, afterEach, beforeAll, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ChatPanel } from "@/components/chat/chat-panel";
import { SSEContext, type SSEConnectionStatus } from "@/lib/hooks/use-sse";
import type { ChatRoom, Message } from "@/lib/types";

beforeAll(() => {
  Element.prototype.scrollIntoView = vi.fn();
});

afterEach(() => cleanup());

const baseChat: ChatRoom = {
  chatRoomId: "c1",
  listingId: "l1",
  listingTitle: "집행검 +7",
  listingStatus: "available",
  counterparty: { userId: "u2", nickname: "흑기사" },
  chatStatus: "active",
  unreadCount: 0,
  updatedAt: new Date().toISOString(),
};

const baseMessage: Message = {
  messageId: "m1",
  chatRoomId: "c1",
  senderUserId: "u1",
  messageType: "text",
  bodyText: "안녕하세요",
  sentAt: new Date().toISOString(),
};

function renderWithSSE(status: SSEConnectionStatus) {
  return render(
    <SSEContext.Provider value={status}>
      <ChatPanel
        chats={[baseChat]}
        activeChatId="c1"
        messages={[baseMessage]}
        myUserId="u1"
        onSelectChat={vi.fn()}
        onSendMessage={vi.fn()}
      />
    </SSEContext.Provider>,
  );
}

describe("ChatPanel — SSE banner", () => {
  it("shows reconnecting banner when SSE is reconnecting", () => {
    renderWithSSE("reconnecting");
    const alert = screen.getByRole("alert");
    expect(alert).toBeDefined();
    expect(alert.textContent).toContain("재연결 중");
  });

  it("does not show banner when SSE is connected", () => {
    renderWithSSE("connected");
    expect(screen.queryByRole("alert")).toBeNull();
  });

  it("does not show banner when SSE is disconnected", () => {
    renderWithSSE("disconnected");
    expect(screen.queryByRole("alert")).toBeNull();
  });

  it("does not show banner when no active chat", () => {
    render(
      <SSEContext.Provider value="reconnecting">
        <ChatPanel
          chats={[baseChat]}
          activeChatId={null}
          messages={[]}
          myUserId="u1"
          onSelectChat={vi.fn()}
          onSendMessage={vi.fn()}
        />
      </SSEContext.Provider>,
    );
    expect(screen.queryByRole("alert")).toBeNull();
  });
});
