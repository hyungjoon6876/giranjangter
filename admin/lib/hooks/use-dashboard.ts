import { useQuery } from "@tanstack/react-query";
import { apiClient } from "../api-client";
import type { DashboardStats } from "../types";

export function useDashboardStats() {
  return useQuery<DashboardStats>({
    queryKey: ["admin", "dashboard-stats"],
    queryFn: () => apiClient.getDashboardStats(),
    refetchInterval: 60_000,
  });
}
