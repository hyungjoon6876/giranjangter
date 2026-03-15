"use client";

import { use, useEffect, useRef } from "react";
import { useMessages, useSendMessage } from "@/lib/hooks/use-chats";
import { useMe } from "@/lib/hooks/use-profile";
import { ChatMessage } from "@/components/chat/chat-message";
import { ChatInput } from "@/components/chat/chat-input";
import { Loading } from "@/components/ui/loading";

export default function ChatDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const { data: me } = useMe();
  const { data, isLoading } = useMessages(id);
  const sendMessage = useSendMessage();
  const bottomRef = useRef<HTMLDivElement>(null);

  const messages = data?.data ? [...data.data].reverse() : [];

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages.length]);

  if (isLoading) return <Loading />;

  return (
    <div className="flex flex-col h-[calc(100vh-120px)]">
      <div className="flex-1 overflow-y-auto p-4">
        {messages.map((m) => (
          <ChatMessage key={m.messageId} message={m} isMine={m.senderUserId === me?.userId} />
        ))}
        <div ref={bottomRef} />
      </div>
      <ChatInput onSend={(text) => sendMessage.mutate({ chatId: id, text })} />
    </div>
  );
}
