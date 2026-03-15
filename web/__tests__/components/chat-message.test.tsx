import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ChatMessage } from "@/components/chat/chat-message";
import type { Message } from "@/lib/types";

afterEach(() => cleanup());

const baseMessage: Message = {
  messageId: "m1",
  chatRoomId: "c1",
  senderUserId: "user-1",
  messageType: "text",
  bodyText: "안녕하세요",
  sentAt: new Date().toISOString(),
};

describe("ChatMessage", () => {
  it("my message aligns right with primary bg", () => {
    const { container } = render(<ChatMessage message={baseMessage} isMine={true} />);
    const wrapper = container.firstElementChild as HTMLElement;
    expect(wrapper.className).toContain("justify-end");
    const bubble = wrapper.firstElementChild as HTMLElement;
    expect(bubble.className).toContain("bg-blue-bright");
  });

  it("their message aligns left with card bg", () => {
    const { container } = render(<ChatMessage message={baseMessage} isMine={false} />);
    const wrapper = container.firstElementChild as HTMLElement;
    expect(wrapper.className).toContain("justify-start");
    const bubble = wrapper.firstElementChild as HTMLElement;
    expect(bubble.className).toContain("bg-card");
  });

  it("system message renders centered", () => {
    const sysMsg: Message = {
      ...baseMessage,
      messageType: "system",
      senderUserId: undefined,
      bodyText: "예약이 확정되었습니다",
    };
    const { container } = render(<ChatMessage message={sysMsg} isMine={false} />);
    const wrapper = container.firstElementChild as HTMLElement;
    expect(wrapper.className).toContain("justify-center");
    expect(screen.getByText("예약이 확정되었습니다")).toBeDefined();
  });
});
