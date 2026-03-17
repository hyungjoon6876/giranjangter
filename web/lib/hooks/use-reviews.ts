import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useUserReviews(userId: string) {
  return useQuery({
    queryKey: ["user-reviews", userId],
    queryFn: () => apiClient.getUserReviews(userId),
    enabled: !!userId,
  });
}
