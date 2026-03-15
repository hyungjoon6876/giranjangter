"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useChats, useMessages, useSendMessage } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { ChatPanel } from "@/components/chat/chat-panel";
import { ChatListItem } from "@/components/chat/chat-list-item";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function ChatsPage() {
  const router = useRouter();
  const { data: me } = useMe();
  const { data: chatsData, isLoading } = useChats();
  const [activeChatId, setActiveChatId] = useState<string | null>(null);
  const { data: messagesData } = useMessages(activeChatId ?? "");
  const sendMessage = useSendMessage();

  const chats = chatsData?.data ?? [];
  const messages = messagesData?.data ? [...messagesData.data].reverse() : [];

  if (isLoading) return <Loading />;
  if (!chats.length) return <EmptyState title="채팅이 없습니다" description="매물에서 채팅을 시작해보세요" />;

  // Desktop: split panel
  return (
    <>
      {/* Desktop split panel */}
      <div className="hidden lg:block">
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
