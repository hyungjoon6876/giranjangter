"use client";

import { use, useEffect, useRef, useState } from "react";
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

export default function ChatDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const router = useRouter();
  const { isLoggedIn } = useAuthGuard();
  const qc = useQueryClient();
  const { data: me } = useMe();
  const { data: chatsData } = useChats();
  const { data, isLoading } = useMessages(id);
  const sendMessage = useSendMessage();
  const bottomRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<ChatInputHandle>(null);
  const [reservationOpen, setReservationOpen] = useState(false);
  const [reportOpen, setReportOpen] = useState(false);

  const messages = data?.data ? [...data.data].reverse() : [];
  const groupedMessages = computeGroupFlags(messages);
  const activeChat = chatsData?.data?.find((c: { chatRoomId: string }) => c.chatRoomId === id);
  useMarkRead(id, messages);

  useEffect(() => {
    if (!isLoggedIn) {
      router.replace("/login");
    }
  }, [isLoggedIn, router]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

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
      <div role="log" aria-live="polite" className="flex-1 overflow-y-auto p-4">
        {groupedMessages.map((m) => (
          <ChatMessage
            key={m.messageId}
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
        ))}
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
