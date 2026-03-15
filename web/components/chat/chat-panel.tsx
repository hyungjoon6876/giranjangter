"use client";

import { useEffect, useRef } from "react";
import type { ChatRoom, Message } from "@/lib/types";
import { useSSEConnectionStatus } from "@/lib/hooks/use-sse";
import { ChatListItem } from "./chat-list-item";
import { ChatMessage } from "./chat-message";
import { ChatInput, type ChatInputHandle } from "./chat-input";

interface ChatPanelProps {
  chats: ChatRoom[];
  activeChatId: string | null;
  messages: Message[];
  myUserId: string | null;
  onSelectChat: (chatId: string) => void;
  onSendMessage: (text: string) => void;
}

export function ChatPanel({ chats, activeChatId, messages, myUserId, onSelectChat, onSendMessage }: ChatPanelProps) {
  const connectionStatus = useSSEConnectionStatus();
  const bottomRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<ChatInputHandle>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  useEffect(() => {
    if (activeChatId) {
      inputRef.current?.focus();
    }
  }, [activeChatId]);

  const activeChat = chats.find((c) => c.chatRoomId === activeChatId);

  return (
    <div className="flex h-[calc(100vh-64px)] lg:h-[calc(100vh-0px)]">
      {/* Chat list */}
      <div role="region" aria-label="대화 목록" className="w-72 border-r border-border bg-dark overflow-y-auto flex-shrink-0">
        <div className="p-3 border-b border-border font-semibold text-gold">
          채팅 <span className="text-text-dim font-normal text-sm">({chats.length})</span>
        </div>
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
      <div className="flex-1 flex flex-col bg-darkest">
        {activeChat ? (
          <>
            <div className="px-4 py-3 bg-dark border-b border-border">
              <div className="font-semibold text-sm text-text-primary">{activeChat.counterparty.nickname}</div>
              <div className="text-xs text-text-dim">{activeChat.listingTitle} · {activeChat.listingStatus}</div>
            </div>
            {connectionStatus === "reconnecting" && (
              <div role="alert" className="bg-[#e67e22]/10 text-[#e67e22] text-xs text-center py-2 px-4">
                연결이 끊어졌습니다. 재연결 중...
              </div>
            )}
            <div role="log" aria-live="polite" className="flex-1 overflow-y-auto p-4">
              {messages.map((m) => (
                <ChatMessage key={m.messageId} message={m} isMine={m.senderUserId === myUserId} />
              ))}
              <div ref={bottomRef} />
            </div>
            <ChatInput ref={inputRef} onSend={onSendMessage} />
          </>
        ) : (
          <div className="flex-1 flex flex-col items-center justify-center text-text-secondary gap-3">
            <svg xmlns="http://www.w3.org/2000/svg" className="w-12 h-12 text-text-dim" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.5}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M8.625 12a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm0 0H8.25m4.125 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm0 0H12m4.125 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 0 1-2.555-.337A5.972 5.972 0 0 1 5.41 20.97a5.969 5.969 0 0 1-.474-.065 4.48 4.48 0 0 0 .978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25Z" />
            </svg>
            <p className="text-gold">채팅방을 선택해주세요</p>
          </div>
        )}
      </div>
    </div>
  );
}
