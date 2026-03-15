"use client";

import { use, useEffect, useRef, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { useMessages, useSendMessage } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { ChatMessage } from "@/components/chat/chat-message";
import { ChatInput } from "@/components/chat/chat-input";
import { ReservationModal } from "@/components/forms/reservation-modal";
import { ReportModal } from "@/components/forms/report-modal";
import { Loading } from "@/components/ui/loading";

export default function ChatDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const qc = useQueryClient();
  const { data: me } = useMe();
  const { data, isLoading } = useMessages(id);
  const sendMessage = useSendMessage();
  const bottomRef = useRef<HTMLDivElement>(null);
  const [reservationOpen, setReservationOpen] = useState(false);
  const [reportOpen, setReportOpen] = useState(false);

  const messages = data?.data ? [...data.data].reverse() : [];

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  if (isLoading) return <Loading />;

  return (
    <div className="flex flex-col h-[calc(100vh-120px)]">
      {/* Gold accent top bar */}
      <div className="h-0.5 bg-gradient-to-r from-gold/60 via-gold to-gold/60" />
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
      <div className="flex-1 overflow-y-auto p-4">
        {messages.map((m) => (
          <ChatMessage key={m.messageId} message={m} isMine={m.senderUserId === me?.userId} />
        ))}
        <div ref={bottomRef} />
      </div>
      <ChatInput onSend={(text) => sendMessage.mutate({ chatId: id, text })} />

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
