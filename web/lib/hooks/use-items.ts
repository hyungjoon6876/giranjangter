import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useItemSearch(q: string, categoryId?: string) {
  return useQuery({
    queryKey: ["items-search", q, categoryId],
    queryFn: () => apiClient.searchItems({ q, categoryId }),
    enabled: q.length >= 1 || !!categoryId,
    staleTime: 60_000,
  });
}
