import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useChats() {
  return useQuery({
    queryKey: ["chats"],
    queryFn: () => apiClient.getChats(),
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
    mutationFn: ({ chatId, text }: { chatId: string; text: string }) =>
      apiClient.sendMessage(chatId, text),
    onSuccess: (_, { chatId }) => {
      qc.invalidateQueries({ queryKey: ["messages", chatId] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useCreateChat() {
  return useMutation({
    mutationFn: (listingId: string) => apiClient.createChat(listingId),
  });
}
