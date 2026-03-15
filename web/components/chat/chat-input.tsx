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
