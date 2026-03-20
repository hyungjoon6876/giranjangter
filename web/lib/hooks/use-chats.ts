import { useEffect } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { PaginatedResponse, Message } from "@/lib/types";

export function useChats() {
  return useQuery({
    queryKey: ["chats"],
    queryFn: () => apiClient.getChats(),
    enabled: apiClient.isLoggedIn,
  });
}

export function useMessages(chatId: string) {
  return useQuery({
    queryKey: ["messages", chatId],
    queryFn: () => apiClient.getMessages(chatId),
    enabled: !!chatId,
    refetchInterval: 5_000, // Poll every 5s as fallback to SSE
  });
}

export function useSendMessage() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({
      chatId,
      text,
      clientMessageId,
    }: {
      chatId: string;
      text: string;
      clientMessageId: string;
    }) => apiClient.sendMessage(chatId, text, clientMessageId),

    onMutate: async ({ chatId, text, clientMessageId }) => {
      await qc.cancelQueries({ queryKey: ["messages", chatId] });
      const previous = qc.getQueryData<PaginatedResponse<Message>>(["messages", chatId]);

      const optimisticMsg: Message = {
        messageId: clientMessageId,
        chatRoomId: chatId,
        senderUserId: "_optimistic_",
        messageType: "text",
        bodyText: text,
        sentAt: new Date().toISOString(),
        status: "sending",
      };

      // PREPEND because data is in DESC order (newest first)
      qc.setQueryData<PaginatedResponse<Message>>(["messages", chatId], (old) => ({
        data: [optimisticMsg, ...(old?.data ?? [])],
        cursor: old?.cursor ?? { hasMore: false },
      }));

      return { previous, chatId };
    },

    onSuccess: (serverMsg, { chatId, clientMessageId }) => {
      qc.setQueryData<PaginatedResponse<Message>>(["messages", chatId], (old) => ({
        data: (old?.data ?? []).map((m) =>
          m.messageId === clientMessageId ? { ...serverMsg, status: "sent" as const } : m
        ),
        cursor: old?.cursor ?? { hasMore: false },
      }));
      qc.invalidateQueries({ queryKey: ["chats"] });
    },

    onError: (_err, { chatId, clientMessageId }) => {
      qc.setQueryData<PaginatedResponse<Message>>(["messages", chatId], (old) => ({
        data: (old?.data ?? []).map((m) =>
          m.messageId === clientMessageId ? { ...m, status: "failed" as const } : m
        ),
        cursor: old?.cursor ?? { hasMore: false },
      }));
    },
  });
}

export function useCreateChat() {
  return useMutation({
    mutationFn: (listingId: string) => apiClient.createChat(listingId),
  });
}

export function useMarkRead(chatId: string, messages: Message[]) {
  const lastMsg = messages.filter(
    (m) => m.senderUserId && m.status !== "sending"
  ).at(-1);

  useEffect(() => {
    if (chatId && lastMsg?.messageId) {
      apiClient.markRead(chatId, lastMsg.messageId).catch(() => {});
    }
  }, [chatId, lastMsg?.messageId]);
}
