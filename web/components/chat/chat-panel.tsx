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
      <div className="w-72 border-r border-border bg-dark overflow-y-auto flex-shrink-0">
        <div className="p-3 border-b border-border font-semibold text-text-primary">채팅 목록</div>
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
            <div className="px-4 py-3 bg-dark border-b border-border font-semibold text-sm text-text-primary">
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
