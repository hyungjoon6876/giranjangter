import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "../api-client";
import type { AdminReport } from "../types";

export function useReports() {
  return useQuery<AdminReport[]>({
    queryKey: ["admin", "reports"],
    queryFn: async () => {
      const res = await apiClient.getReports();
      return res.data;
    },
  });
}

export function useReport(reportId: string) {
  return useQuery<AdminReport>({
    queryKey: ["admin", "reports", reportId],
    queryFn: () => apiClient.getReport(reportId),
    enabled: !!reportId,
  });
}

export function useReportAction() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (vars: {
      reportId: string;
      actionCode: string;
      targetUserId: string;
      memo?: string;
      restrictionScope?: string;
    }) => {
      const { reportId, ...data } = vars;
      return apiClient.reportAction(reportId, data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["admin", "reports"] });
      queryClient.invalidateQueries({ queryKey: ["admin", "dashboard-stats"] });
    },
  });
}

export function useUpdateReportStatus() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (vars: { reportId: string; status: string }) =>
      apiClient.updateReportStatus(vars.reportId, vars.status),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["admin", "reports"] });
    },
  });
}
