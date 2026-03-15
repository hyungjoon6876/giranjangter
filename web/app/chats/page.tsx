"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { useChats, useMessages, useSendMessage } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { ChatPanel } from "@/components/chat/chat-panel";
import { ChatListItem } from "@/components/chat/chat-list-item";
import { ReservationModal } from "@/components/forms/reservation-modal";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function ChatsPage() {
  const router = useRouter();
  const { data: me } = useMe();
  const { data: chatsData, isLoading } = useChats();
  const qc = useQueryClient();
  const [activeChatId, setActiveChatId] = useState<string | null>(null);
  const [reservationOpen, setReservationOpen] = useState(false);
  const { data: messagesData } = useMessages(activeChatId ?? "");
  const sendMessage = useSendMessage();

  const chats = chatsData?.data ?? [];
  const messages = messagesData?.data ? [...messagesData.data].reverse() : [];

  if (isLoading) return <Loading />;
  if (!chats.length) return <EmptyState title="채팅이 없습니다" description="매물에서 채팅을 시작해보세요" actionLabel="매물 둘러보기" actionHref="/" />;

  // Desktop: split panel
  return (
    <>
      {/* Desktop split panel */}
      <div className="hidden lg:block relative">
        {activeChatId && (
          <div className="absolute top-0 right-0 z-10 p-2">
            <button
              onClick={() => setReservationOpen(true)}
              className="px-3 py-1.5 text-xs font-medium btn-gold-gradient text-white rounded-lg transition-colors"
            >
              예약 제안
            </button>
          </div>
        )}
        <ChatPanel
          chats={chats}
          activeChatId={activeChatId}
          messages={messages}
          myUserId={me?.userId ?? null}
          onSelectChat={setActiveChatId}
          onSendMessage={(text) => {
            if (activeChatId) sendMessage.mutate({ chatId: activeChatId, text });
          }}
        />
        {activeChatId && (
          <ReservationModal
            open={reservationOpen}
            onClose={() => setReservationOpen(false)}
            chatId={activeChatId}
            onCreated={() => qc.invalidateQueries({ queryKey: ["chats"] })}
          />
        )}
      </div>

      {/* Mobile: list only */}
      <div className="lg:hidden">
        {chats.map((c) => (
          <ChatListItem
            key={c.chatRoomId}
            chat={c}
            onClick={() => router.push(`/chats/${c.chatRoomId}`)}
          />
        ))}
      </div>
    </>
  );
}
