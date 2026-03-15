import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ChatInput, type ChatInputHandle } from "@/components/chat/chat-input";
import { createRef } from "react";

afterEach(() => cleanup());

describe("ChatInput", () => {
  it("renders a textarea with aria-label", () => {
    render(<ChatInput onSend={vi.fn()} />);
    const textarea = screen.getByLabelText("메시지 입력");
    expect(textarea).toBeDefined();
    expect(textarea.tagName).toBe("TEXTAREA");
  });

  it("send button has aria-label", () => {
    render(<ChatInput onSend={vi.fn()} />);
    const btn = screen.getByLabelText("메시지 전송");
    expect(btn).toBeDefined();
    expect(btn.tagName).toBe("BUTTON");
  });

  it("calls onSend on Enter (without Shift)", async () => {
    const onSend = vi.fn();
    render(<ChatInput onSend={onSend} />);
    const textarea = screen.getByLabelText("메시지 입력");

    fireEvent.change(textarea, { target: { value: "hello" } });
    fireEvent.keyDown(textarea, { key: "Enter", shiftKey: false });

    expect(onSend).toHaveBeenCalledWith("hello");
  });

  it("does not submit on Shift+Enter", () => {
    const onSend = vi.fn();
    render(<ChatInput onSend={onSend} />);
    const textarea = screen.getByLabelText("메시지 입력");

    fireEvent.change(textarea, { target: { value: "hello" } });
    fireEvent.keyDown(textarea, { key: "Enter", shiftKey: true });

    expect(onSend).not.toHaveBeenCalled();
  });

  it("does not submit empty text", () => {
    const onSend = vi.fn();
    render(<ChatInput onSend={onSend} />);
    const textarea = screen.getByLabelText("메시지 입력");

    fireEvent.keyDown(textarea, { key: "Enter", shiftKey: false });
    expect(onSend).not.toHaveBeenCalled();
  });

  it("clears input after submit", () => {
    const onSend = vi.fn();
    render(<ChatInput onSend={onSend} />);
    const textarea = screen.getByLabelText("메시지 입력") as HTMLTextAreaElement;

    fireEvent.change(textarea, { target: { value: "hello" } });
    fireEvent.keyDown(textarea, { key: "Enter", shiftKey: false });

    expect(textarea.value).toBe("");
  });

  it("exposes focus via ref", () => {
    const ref = createRef<ChatInputHandle>();
    render(<ChatInput ref={ref} onSend={vi.fn()} />);

    const textarea = screen.getByLabelText("메시지 입력");
    ref.current?.focus();
    expect(document.activeElement).toBe(textarea);
  });

  it("send button is disabled when input is empty", () => {
    render(<ChatInput onSend={vi.fn()} />);
    const btn = screen.getByLabelText("메시지 전송") as HTMLButtonElement;
    expect(btn.disabled).toBe(true);
  });
});
