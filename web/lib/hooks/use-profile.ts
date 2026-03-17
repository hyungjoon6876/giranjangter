import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useMe() {
  return useQuery({
    queryKey: ["me"],
    queryFn: () => apiClient.getMe(),
    enabled: apiClient.isLoggedIn,
  });
}

export function useMyListings(status?: string) {
  return useQuery({
    queryKey: ["myListings", status],
    queryFn: () => apiClient.getMyListings(status),
    enabled: apiClient.isLoggedIn,
  });
}

export function useMyTrades() {
  return useQuery({
    queryKey: ["myTrades"],
    queryFn: () => apiClient.getMyTrades(),
    enabled: apiClient.isLoggedIn,
  });
}

export function useNotifications() {
  return useQuery({
    queryKey: ["notifications"],
    queryFn: () => apiClient.getNotifications(),
    enabled: apiClient.isLoggedIn,
    refetchInterval: 30_000,
  });
}

export function useMarkNotificationsRead() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (ids: string[]) => apiClient.markNotificationsRead(ids),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["notifications"] }),
  });
}

export function useUpdateProfile() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: {
      nickname?: string;
      introduction?: string;
      primaryServerId?: string;
      avatarUrl?: string;
    }) => apiClient.updateProfile(data),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["me"] }),
  });
}
