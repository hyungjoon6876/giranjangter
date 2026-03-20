"use client";

import { Fragment, use, useCallback, useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { useChats, useMessages, useSendMessage, useMarkRead } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { useAuthGuard } from "@/lib/hooks/use-auth-guard";
import { ChatMessage, computeGroupFlags } from "@/components/chat/chat-message";
import { ChatInput, type ChatInputHandle } from "@/components/chat/chat-input";
import { ReservationModal } from "@/components/forms/reservation-modal";
import { ReportModal } from "@/components/forms/report-modal";
import { Loading } from "@/components/ui/loading";
import { ListingInfoCard } from "@/components/chat/chat-panel";
import { useSSEConnectionStatus } from "@/lib/hooks/use-sse";

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

export default function ChatDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const router = useRouter();
  const { isLoggedIn } = useAuthGuard();
  const qc = useQueryClient();
  const { data: me } = useMe();
  const { data: chatsData } = useChats();
  const connectionStatus = useSSEConnectionStatus();
  const sseConnected = connectionStatus === "connected";
  const { data, isLoading, fetchNextPage, hasNextPage, isFetchingNextPage } = useMessages(id, sseConnected);
  const sendMessage = useSendMessage();
  const bottomRef = useRef<HTMLDivElement>(null);
  const scrollRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<ChatInputHandle>(null);
  const lastReadIdRef = useRef<string | undefined>(undefined);
  const newMsgRef = useRef<HTMLDivElement>(null);
  const [reservationOpen, setReservationOpen] = useState(false);
  const [reportOpen, setReportOpen] = useState(false);

  const messages = data?.pages ? data.pages.flatMap((p) => [...p.data]).reverse() : [];

  const handleScroll = useCallback(() => {
    const el = scrollRef.current;
    if (!el) return;
    if (el.scrollTop < 50 && hasNextPage && !isFetchingNextPage) {
      const prevHeight = el.scrollHeight;
      const prevTop = el.scrollTop;
      fetchNextPage().then(() => {
        requestAnimationFrame(() => {
          if (scrollRef.current) {
            scrollRef.current.scrollTop = scrollRef.current.scrollHeight - prevHeight + prevTop;
          }
        });
      });
    }
  }, [hasNextPage, isFetchingNextPage, fetchNextPage]);
  const groupedMessages = computeGroupFlags(messages);
  const activeChat = chatsData?.data?.find((c: { chatRoomId: string }) => c.chatRoomId === id);
  useMarkRead(id, messages);

  useEffect(() => {
    if (!isLoggedIn) {
      router.replace("/login");
    }
  }, [isLoggedIn, router]);

  useEffect(() => {
    const chat = chatsData?.data?.find((c: { chatRoomId: string }) => c.chatRoomId === id);
    if (chat?.myLastReadMessageId) lastReadIdRef.current = chat.myLastReadMessageId;
  }, [id, chatsData]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  useEffect(() => {
    if (newMsgRef.current) {
      newMsgRef.current.scrollIntoView({ behavior: "smooth", block: "center" });
    }
  }, [id]);

  useEffect(() => {
    inputRef.current?.focus();
  }, []);

  if (!isLoggedIn) return null;

  if (isLoading) return <Loading />;

  return (
    <div className="flex flex-col h-[calc(100vh-120px)]">
      {/* Gold accent top bar */}
      <div className="h-0.5 bg-gradient-to-r from-gold/60 via-gold to-gold/60" />
      {activeChat && <ListingInfoCard chat={activeChat} />}
      <div className="flex items-center gap-2 px-4 py-2 bg-dark border-b border-border">
        <button
          onClick={() => setReservationOpen(true)}
          className="px-3 py-1.5 text-xs font-medium btn-gold-gradient text-white rounded-lg transition-colors"
        >
          예약 제안
        </button>
        <button
          onClick={() => setReportOpen(true)}
          className="px-3 py-1.5 text-xs font-medium border border-border text-danger rounded-lg hover:bg-medium transition-colors"
        >
          신고
        </button>
      </div>
      <div ref={scrollRef} onScroll={handleScroll} role="log" aria-live="polite" className="flex-1 overflow-y-auto p-4">
        {isFetchingNextPage && (
          <div className="text-center py-2 text-text-secondary text-xs">이전 메시지 불러오는 중...</div>
        )}
        {groupedMessages.map((m, i) => {
          const prev = groupedMessages[i - 1];
          const showDateSep = !prev || !isSameDay(prev.sentAt, m.sentAt);
          const showNewMsgDivider = lastReadIdRef.current && prev?.messageId === lastReadIdRef.current && m.messageId !== lastReadIdRef.current;

          return (
            <Fragment key={m.messageId}>
              {showDateSep && <DateSeparator label={formatDateSeparator(m.sentAt)} />}
              {showNewMsgDivider && <div ref={newMsgRef}><NewMessagesDivider /></div>}
              <ChatMessage
                message={m}
                isMine={m.senderUserId === me?.userId || m.status === "sending" || m.status === "failed"}
                isFirstInGroup={m.isFirstInGroup}
                isLastInGroup={m.isLastInGroup}
                onRetry={(failedMsg) => {
                  sendMessage.mutate({
                    chatId: failedMsg.chatRoomId || id,
                    text: failedMsg.bodyText ?? "",
                    clientMessageId: crypto.randomUUID(),
                  });
                }}
              />
            </Fragment>
          );
        })}
        <div ref={bottomRef} />
      </div>
      <ChatInput ref={inputRef} onSend={(text) => sendMessage.mutate({ chatId: id, text, clientMessageId: crypto.randomUUID() })} />

      <ReservationModal
        open={reservationOpen}
        onClose={() => setReservationOpen(false)}
        chatId={id}
        onCreated={() => qc.invalidateQueries({ queryKey: ["chats"] })}
      />
      <ReportModal
        open={reportOpen}
        onClose={() => setReportOpen(false)}
        targetType="chat_room"
        targetId={id}
      />
    </div>
  );
}
