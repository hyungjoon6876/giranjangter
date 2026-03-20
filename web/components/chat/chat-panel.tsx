"use client";

import { Fragment, useEffect, useRef } from "react";
import Link from "next/link";
import Image from "next/image";
import type { ChatRoom, Message } from "@/lib/types";
import { useSSEConnectionStatus } from "@/lib/hooks/use-sse";
import { ChatListItem } from "./chat-list-item";
import { ChatMessage, computeGroupFlags } from "./chat-message";
import { ChatInput, type ChatInputHandle } from "./chat-input";

function isSameDay(a: string, b: string): boolean {
  const da = new Date(a);
  const db = new Date(b);
  return da.getFullYear() === db.getFullYear() && da.getMonth() === db.getMonth() && da.getDate() === db.getDate();
}

function formatDateSeparator(dateStr: string): string {
  const date = new Date(dateStr);
  const now = new Date();
  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
  const msgDay = new Date(date.getFullYear(), date.getMonth(), date.getDate());
  const diffDays = Math.floor((today.getTime() - msgDay.getTime()) / 86400000);
  if (diffDays === 0) return "오늘";
  if (diffDays === 1) return "어제";
  if (date.getFullYear() === now.getFullYear()) return `${date.getMonth() + 1}월 ${date.getDate()}일`;
  return `${date.getFullYear()}년 ${date.getMonth() + 1}월 ${date.getDate()}일`;
}

function DateSeparator({ label }: { label: string }) {
  return (
    <div className="flex items-center gap-3 my-3">
      <div className="flex-1 border-t border-border" />
      <span className="text-xs text-text-secondary">{label}</span>
      <div className="flex-1 border-t border-border" />
    </div>
  );
}

function NewMessagesDivider() {
  return (
    <div className="flex items-center gap-3 my-3">
      <div className="flex-1 border-t border-gold" />
      <span className="text-xs text-gold font-medium">새 메시지</span>
      <div className="flex-1 border-t border-gold" />
    </div>
  );
}

interface ChatPanelProps {
  chats: ChatRoom[];
  activeChatId: string | null;
  messages: Message[];
  myUserId: string | null;
  onSelectChat: (chatId: string) => void;
  onSendMessage: (text: string) => void;
  onRetryMessage?: (message: Message) => void;
}

export function ChatPanel({ chats, activeChatId, messages, myUserId, onSelectChat, onSendMessage, onRetryMessage }: ChatPanelProps) {
  const connectionStatus = useSSEConnectionStatus();
  const bottomRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<ChatInputHandle>(null);
  const lastReadIdRef = useRef<string | undefined>(undefined);
  const newMsgRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const chat = chats.find(c => c.chatRoomId === activeChatId);
    if (chat?.myLastReadMessageId) lastReadIdRef.current = chat.myLastReadMessageId;
  }, [activeChatId, chats]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  useEffect(() => {
    if (newMsgRef.current) {
      newMsgRef.current.scrollIntoView({ behavior: "smooth", block: "center" });
    }
  }, [activeChatId]);

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
              <div className="text-xs text-text-dim">{activeChat.listingTitle}</div>
            </div>
            <ListingInfoCard chat={activeChat} />
            {connectionStatus === "reconnecting" && (
              <div role="alert" className="bg-[#e67e22]/10 text-[#e67e22] text-xs text-center py-2 px-4">
                연결이 끊어졌습니다. 재연결 중...
              </div>
            )}
            {connectionStatus === "disconnected" && (
              <div role="alert" className="bg-[#e74c3c]/10 text-[#e74c3c] text-xs text-center py-2 px-4">
                연결이 끊어졌습니다. 페이지를 새로고침해주세요.
              </div>
            )}
            <div role="log" aria-live="polite" className="flex-1 overflow-y-auto p-4">
              {(() => {
                const grouped = computeGroupFlags(messages);
                return grouped.map((m, i) => {
                  const prev = grouped[i - 1];
                  const showDateSep = !prev || !isSameDay(prev.sentAt, m.sentAt);
                  const showNewMsgDivider = lastReadIdRef.current && prev?.messageId === lastReadIdRef.current && m.messageId !== lastReadIdRef.current;

                  return (
                    <Fragment key={m.messageId}>
                      {showDateSep && <DateSeparator label={formatDateSeparator(m.sentAt)} />}
                      {showNewMsgDivider && <div ref={newMsgRef}><NewMessagesDivider /></div>}
                      <ChatMessage
                        message={m}
                        isMine={m.senderUserId === myUserId || m.status === "sending" || m.status === "failed"}
                        isFirstInGroup={m.isFirstInGroup}
                        isLastInGroup={m.isLastInGroup}
                        onRetry={onRetryMessage}
                      />
                    </Fragment>
                  );
                });
              })()}
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

const listingStatusConfig: Record<string, { label: string; color: string; dim?: boolean }> = {
  available: { label: "거래 가능", color: "bg-green-600" },
  reserved: { label: "예약중", color: "bg-yellow-600" },
  sold: { label: "거래 완료", color: "bg-medium", dim: true },
  deleted: { label: "삭제됨", color: "bg-medium", dim: true },
};

export function ListingInfoCard({ chat }: { chat: ChatRoom }) {
  const st = listingStatusConfig[chat.listingStatus] ?? listingStatusConfig.available;

  return (
    <Link
      href={`/listings/${chat.listingId}`}
      className={`flex items-center gap-3 px-4 py-2 border-b border-border hover:bg-medium/50 transition-colors ${st.dim ? "opacity-50" : ""}`}
    >
      {chat.listingThumbnail ? (
        <Image src={chat.listingThumbnail} alt="" width={40} height={40} className="rounded object-cover" unoptimized />
      ) : (
        <div className="w-10 h-10 rounded bg-medium flex items-center justify-center text-text-secondary text-xs">?</div>
      )}
      <div className="flex-1 min-w-0">
        <p className="font-medium text-sm truncate">{chat.listingTitle}</p>
        <p className="text-xs text-text-secondary truncate">
          {chat.listingPrice != null && `${chat.listingPrice.toLocaleString()} 아덴`}
          {chat.listingPrice != null && chat.listingServerName && " · "}
          {chat.listingServerName && chat.listingServerName}
        </p>
      </div>
      <span className={`text-xs px-2 py-0.5 rounded-full text-white whitespace-nowrap ${st.color}`}>
        {st.label}
      </span>
    </Link>
  );
}
