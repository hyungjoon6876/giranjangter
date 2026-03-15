"use client";

import { KpiCard } from "@/components/dashboard/kpi-card";
import { useDashboardStats } from "@/lib/hooks/use-dashboard";
import { useReports } from "@/lib/hooks/use-reports";
import { statusBadgeVariant, formatTimeAgo } from "@/lib/utils";

export default function DashboardPage() {
  const { data: stats, isLoading: statsLoading } = useDashboardStats();
  const { data: reports } = useReports();

  const recentReports = (reports ?? []).slice(0, 5);

  return (
    <div className="space-y-6">
      {/* KPI Cards */}
      {statsLoading ? (
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {[0, 1, 2, 3].map((i) => (
            <div
              key={i}
              className="h-28 animate-pulse rounded-xl border border-border bg-white"
            />
          ))}
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <KpiCard
            title="미처리 신고"
            value={stats?.pendingReports ?? 0}
            color="danger"
            description="접수된 신고 중 미처리"
          />
          <KpiCard
            title="오늘 거래"
            value={stats?.tradesToday ?? 0}
            color="success"
            description="오늘 완료된 거래"
          />
          <KpiCard
            title="활성 매물"
            value={stats?.activeListings ?? 0}
            color="primary"
            description="현재 공개 중인 매물"
          />
          <KpiCard
            title="신규 가입"
            value={stats?.newUsersToday ?? 0}
            color="warning"
            description="오늘 신규 가입자"
          />
        </div>
      )}

      {/* Recent Reports */}
      <div className="rounded-xl border border-border bg-white shadow-sm">
        <div className="border-b border-border px-5 py-4">
          <h2 className="text-base font-semibold text-text-primary">
            최근 신고
          </h2>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-border text-left text-text-secondary">
                <th className="px-5 py-3 font-medium">유형</th>
                <th className="px-5 py-3 font-medium">대상</th>
                <th className="px-5 py-3 font-medium">상태</th>
                <th className="px-5 py-3 font-medium">시간</th>
              </tr>
            </thead>
            <tbody>
              {recentReports.length === 0 ? (
                <tr>
                  <td
                    colSpan={4}
                    className="px-5 py-8 text-center text-text-secondary"
                  >
                    신고 내역이 없습니다.
                  </td>
                </tr>
              ) : (
                recentReports.map((r) => (
                  <tr
                    key={r.reportId}
                    className="border-b border-border last:border-b-0 hover:bg-slate-50"
                  >
                    <td className="px-5 py-3">{r.reportType}</td>
                    <td className="px-5 py-3 font-mono text-xs">
                      {r.targetType}:{r.targetId.slice(0, 8)}
                    </td>
                    <td className="px-5 py-3">
                      <span
                        className={`inline-block rounded-full px-2 py-1 text-xs font-medium ${statusBadgeVariant(r.status)}`}
                      >
                        {r.status}
                      </span>
                    </td>
                    <td className="px-5 py-3 text-text-secondary">
                      {formatTimeAgo(r.createdAt)}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
