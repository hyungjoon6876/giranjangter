import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "../api-client";
import type { AdminUser, ModerationAction } from "../types";

export function useUsers(params?: { q?: string; status?: string }) {
  return useQuery<AdminUser[]>({
    queryKey: ["admin", "users", params],
    queryFn: async () => {
      const res = await apiClient.getUsers(params);
      return res.data;
    },
  });
}

export function useUser(userId: string) {
  return useQuery<AdminUser>({
    queryKey: ["admin", "users", userId],
    queryFn: () => apiClient.getUser(userId),
    enabled: !!userId,
  });
}

export function useUserModerationHistory(userId: string) {
  return useQuery<ModerationAction[]>({
    queryKey: ["admin", "users", userId, "moderation-history"],
    queryFn: async () => {
      const res = await apiClient.getUserModerationHistory(userId);
      return res.data;
    },
    enabled: !!userId,
  });
}

export function useRestrictUser() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (vars: {
      userId: string;
      restrictionScope: string;
      durationDays?: number;
      reasonCode: string;
      memo?: string;
    }) => {
      const { userId, ...data } = vars;
      return apiClient.restrictUser(userId, data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["admin", "users"] });
      queryClient.invalidateQueries({ queryKey: ["admin", "dashboard-stats"] });
    },
  });
}
