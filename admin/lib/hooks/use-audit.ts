import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "../api-client";
import type {
  AuditLog,
  AdminChatMessage,
  TradeCompletion,
  AdminListing,
} from "../types";

export function useAuditLogs() {
  return useQuery<AuditLog[]>({
    queryKey: ["admin", "audit-logs"],
    queryFn: async () => {
      const res = await apiClient.getAuditLogs();
      return res.data;
    },
  });
}

export function useChatMessages(chatId: string) {
  return useQuery<AdminChatMessage[]>({
    queryKey: ["admin", "chat-messages", chatId],
    queryFn: async () => {
      const res = await apiClient.getChatMessages(chatId);
      return res.data;
    },
    enabled: !!chatId,
  });
}

export function useTrades() {
  return useQuery<TradeCompletion[]>({
    queryKey: ["admin", "trades"],
    queryFn: async () => {
      const res = await apiClient.getTrades();
      return res.data;
    },
  });
}

export function useListings(params?: {
  status?: string;
  visibility?: string;
}) {
  return useQuery<AdminListing[]>({
    queryKey: ["admin", "listings", params],
    queryFn: async () => {
      const res = await apiClient.getListings(params);
      return res.data;
    },
  });
}

export function useHideListing() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (listingId: string) => apiClient.hideListing(listingId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["admin", "listings"] });
    },
  });
}

export function useRestoreListing() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (listingId: string) => apiClient.restoreListing(listingId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["admin", "listings"] });
    },
  });
}
