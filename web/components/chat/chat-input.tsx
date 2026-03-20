"use client";

import { forwardRef, useCallback, useEffect, useImperativeHandle, useRef, useState, type FormEvent, type KeyboardEvent } from "react";

interface ChatInputProps {
  onSend: (text: string) => void;
  disabled?: boolean;
}

export interface ChatInputHandle {
  focus: () => void;
}

export const ChatInput = forwardRef<ChatInputHandle, ChatInputProps>(function ChatInput({ onSend, disabled }, ref) {
  const [text, setText] = useState("");
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  useImperativeHandle(ref, () => ({
    focus: () => textareaRef.current?.focus(),
  }));

  const resetHeight = useCallback(() => {
    const el = textareaRef.current;
    if (!el) return;
    el.style.height = "auto";
    el.style.height = `${Math.min(el.scrollHeight, 120)}px`;
  }, []);

  useEffect(() => {
    resetHeight();
  }, [text, resetHeight]);

  const submit = () => {
    const trimmed = text.trim();
    if (!trimmed) return;
    onSend(trimmed);
    setText("");
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    submit();
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      submit();
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2 p-3 bg-dark border-t border-border">
      <textarea
        ref={textareaRef}
        value={text}
        onChange={(e) => setText(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder="메시지를 입력하세요"
        rows={1}
        aria-label="메시지 입력"
        className="flex-1 bg-card border border-border rounded-2xl px-4 py-2 text-base text-text-primary outline-none focus:border-gold placeholder:text-text-dim resize-none"
        disabled={disabled}
      />
      <button
        type="submit"
        disabled={disabled || !text.trim()}
        aria-label="메시지 전송"
        className="btn-gold-gradient text-white px-4 py-2 rounded-full text-sm font-medium disabled:opacity-50"
      >
        전송
      </button>
    </form>
  );
});
