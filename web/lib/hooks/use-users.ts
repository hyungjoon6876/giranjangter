import { useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useBlockUser() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (userId: string) => apiClient.blockUser(userId),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useUnblockUser() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (userId: string) => apiClient.unblockUser(userId),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}
