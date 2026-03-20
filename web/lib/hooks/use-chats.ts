import { useEffect } from "react";
import { useQuery, useMutation, useQueryClient, useInfiniteQuery, type InfiniteData } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { PaginatedResponse, Message } from "@/lib/types";

export function useChats() {
  return useQuery({
    queryKey: ["chats"],
    queryFn: () => apiClient.getChats(),
    enabled: apiClient.isLoggedIn,
  });
}

export function useMessages(chatId: string, sseConnected?: boolean) {
  return useInfiniteQuery({
    queryKey: ["messages", chatId],
    queryFn: ({ pageParam }: { pageParam?: string }) =>
      apiClient.getMessages(chatId, pageParam),
    initialPageParam: undefined as string | undefined,
    getNextPageParam: (lastPage) =>
      lastPage.cursor?.hasMore ? lastPage.cursor.next : undefined,
    enabled: !!chatId,
    refetchInterval: sseConnected ? false : 5_000,
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
      const previous = qc.getQueryData<InfiniteData<PaginatedResponse<Message>>>(["messages", chatId]);

      const optimisticMsg: Message = {
        messageId: clientMessageId,
        chatRoomId: chatId,
        senderUserId: "_optimistic_",
        messageType: "text",
        bodyText: text,
        sentAt: new Date().toISOString(),
        status: "sending",
      };

      // PREPEND to first page because data is in DESC order (newest first)
      qc.setQueryData<InfiniteData<PaginatedResponse<Message>>>(
        ["messages", chatId],
        (old) => {
          if (!old) return old;
          const firstPage = old.pages[0];
          return {
            ...old,
            pages: [
              { ...firstPage, data: [optimisticMsg, ...firstPage.data] },
              ...old.pages.slice(1),
            ],
          };
        }
      );

      return { previous, chatId };
    },

    onSuccess: (serverMsg, { chatId, clientMessageId }) => {
      qc.setQueryData<InfiniteData<PaginatedResponse<Message>>>(
        ["messages", chatId],
        (old) => {
          if (!old) return old;
          return {
            ...old,
            pages: old.pages.map((page) => ({
              ...page,
              data: page.data.map((m) =>
                m.messageId === clientMessageId ? { ...serverMsg, status: "sent" as const } : m
              ),
            })),
          };
        }
      );
      qc.invalidateQueries({ queryKey: ["chats"] });
    },

    onError: (_err, { chatId, clientMessageId }) => {
      qc.setQueryData<InfiniteData<PaginatedResponse<Message>>>(
        ["messages", chatId],
        (old) => {
          if (!old) return old;
          return {
            ...old,
            pages: old.pages.map((page) => ({
              ...page,
              data: page.data.map((m) =>
                m.messageId === clientMessageId ? { ...m, status: "failed" as const } : m
              ),
            })),
          };
        }
      );
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
