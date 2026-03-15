"use client";

import { useState, useMemo } from "react";
import { type ColumnDef } from "@tanstack/react-table";
import { DataTable } from "@/components/data-table/data-table";
import { useReports, useReportAction, useUpdateReportStatus } from "@/lib/hooks/use-reports";
import { statusBadgeVariant, formatDate } from "@/lib/utils";
import type { AdminReport } from "@/lib/types";

const STATUS_TABS = [
  { label: "전체", value: "" },
  { label: "접수", value: "submitted" },
  { label: "처리중", value: "assigned" },
  { label: "완료", value: "resolved" },
] as const;

export default function ReportsPage() {
  const [statusFilter, setStatusFilter] = useState("");
  const [selectedReport, setSelectedReport] = useState<AdminReport | null>(null);
  const [actionMemo, setActionMemo] = useState("");
  const [actionCode, setActionCode] = useState("warn_user");

  const { data: reports, isLoading } = useReports();
  const reportAction = useReportAction();
  const updateStatus = useUpdateReportStatus();

  const filteredReports = useMemo(() => {
    if (!reports) return [];
    if (!statusFilter) return reports;
    return reports.filter((r) => r.status === statusFilter);
  }, [reports, statusFilter]);

  const columns: ColumnDef<AdminReport, unknown>[] = useMemo(
    () => [
      {
        accessorKey: "reportId",
        header: "신고 ID",
        cell: ({ getValue }) => (
          <span className="font-mono text-xs">
            {(getValue() as string).slice(0, 8)}
          </span>
        ),
      },
      { accessorKey: "reportType", header: "유형" },
      {
        accessorKey: "targetType",
        header: "대상",
        cell: ({ row }) => (
          <span className="text-xs">
            {row.original.targetType}:{row.original.targetId.slice(0, 8)}
          </span>
        ),
      },
      {
        accessorKey: "reporterUserId",
        header: "신고자",
        cell: ({ getValue }) => (
          <span className="font-mono text-xs">
            {(getValue() as string).slice(0, 8)}
          </span>
        ),
      },
      {
        accessorKey: "status",
        header: "상태",
        cell: ({ getValue }) => {
          const status = getValue() as string;
          return (
            <span
              className={`inline-block rounded-full px-2 py-1 text-xs font-medium ${statusBadgeVariant(status)}`}
            >
              {status}
            </span>
          );
        },
      },
      {
        accessorKey: "createdAt",
        header: "날짜",
        cell: ({ getValue }) => formatDate(getValue() as string),
      },
    ],
    [],
  );

  function handleAction() {
    if (!selectedReport) return;
    reportAction.mutate(
      {
        reportId: selectedReport.reportId,
        actionCode,
        targetUserId: selectedReport.targetId,
        memo: actionMemo || undefined,
      },
      {
        onSuccess: () => {
          setSelectedReport(null);
          setActionMemo("");
        },
      },
    );
  }

  function handleStatusChange(reportId: string, status: string) {
    updateStatus.mutate({ reportId, status });
  }

  return (
    <div className="space-y-4">
      {/* Status filter tabs */}
      <div className="flex gap-1 rounded-lg border border-border bg-white p-1">
        {STATUS_TABS.map((tab) => (
          <button
            key={tab.value}
            onClick={() => setStatusFilter(tab.value)}
            className={`rounded-md px-4 py-2 text-sm font-medium transition-colors ${
              statusFilter === tab.value
                ? "bg-primary text-white"
                : "text-text-secondary hover:bg-slate-100"
            }`}
          >
            {tab.label}
          </button>
        ))}
      </div>

      <div className="flex gap-4">
        {/* Table */}
        <div
          className={`rounded-xl border border-border bg-white shadow-sm ${
            selectedReport ? "flex-1" : "w-full"
          }`}
        >
          {isLoading ? (
            <div className="flex h-40 items-center justify-center text-text-secondary">
              로딩 중...
            </div>
          ) : (
            <DataTable
              data={filteredReports}
              columns={columns}
              onRowClick={(row) => setSelectedReport(row)}
              emptyMessage="신고 내역이 없습니다."
            />
          )}
        </div>

        {/* Detail panel */}
        {selectedReport && (
          <div className="w-96 shrink-0 rounded-xl border border-border bg-white shadow-sm">
            <div className="flex items-center justify-between border-b border-border px-5 py-4">
              <h3 className="text-sm font-semibold text-text-primary">
                신고 상세
              </h3>
              <button
                onClick={() => setSelectedReport(null)}
                className="text-text-secondary hover:text-text-primary"
              >
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round">
                  <path d="M4 4l8 8M12 4l-8 8" />
                </svg>
              </button>
            </div>
            <div className="space-y-4 p-5">
              <div>
                <p className="text-xs text-text-secondary">신고 ID</p>
                <p className="font-mono text-sm">{selectedReport.reportId}</p>
              </div>
              <div>
                <p className="text-xs text-text-secondary">유형</p>
                <p className="text-sm">{selectedReport.reportType}</p>
              </div>
              <div>
                <p className="text-xs text-text-secondary">대상</p>
                <p className="text-sm">
                  {selectedReport.targetType}: {selectedReport.targetId}
                </p>
              </div>
              <div>
                <p className="text-xs text-text-secondary">상태</p>
                <span
                  className={`inline-block rounded-full px-2 py-1 text-xs font-medium ${statusBadgeVariant(selectedReport.status)}`}
                >
                  {selectedReport.status}
                </span>
              </div>
              {selectedReport.description && (
                <div>
                  <p className="text-xs text-text-secondary">설명</p>
                  <p className="text-sm">{selectedReport.description}</p>
                </div>
              )}
              <div>
                <p className="text-xs text-text-secondary">신고일</p>
                <p className="text-sm">{formatDate(selectedReport.createdAt)}</p>
              </div>

              {/* Status change buttons */}
              {selectedReport.status !== "resolved" && (
                <div className="border-t border-border pt-4">
                  <p className="mb-2 text-xs font-medium text-text-secondary">
                    상태 변경
                  </p>
                  <div className="flex gap-2">
                    {selectedReport.status === "submitted" && (
                      <button
                        onClick={() =>
                          handleStatusChange(
                            selectedReport.reportId,
                            "assigned",
                          )
                        }
                        className="rounded-lg bg-primary px-3 py-1.5 text-xs font-medium text-white hover:bg-primary-600"
                      >
                        처리 시작
                      </button>
                    )}
                    <button
                      onClick={() =>
                        handleStatusChange(
                          selectedReport.reportId,
                          "resolved",
                        )
                      }
                      className="rounded-lg bg-green-600 px-3 py-1.5 text-xs font-medium text-white hover:bg-green-700"
                    >
                      완료 처리
                    </button>
                  </div>
                </div>
              )}

              {/* Action form */}
              {selectedReport.status !== "resolved" && (
                <div className="border-t border-border pt-4">
                  <p className="mb-2 text-xs font-medium text-text-secondary">
                    조치
                  </p>
                  <select
                    value={actionCode}
                    onChange={(e) => setActionCode(e.target.value)}
                    className="mb-2 w-full rounded-lg border border-border px-3 py-2 text-sm focus:border-primary focus:outline-none"
                  >
                    <option value="warn_user">경고</option>
                    <option value="restrict_chat">채팅 제한</option>
                    <option value="restrict_listing">매물 등록 제한</option>
                    <option value="suspend_account">계정 정지</option>
                  </select>
                  <textarea
                    value={actionMemo}
                    onChange={(e) => setActionMemo(e.target.value)}
                    placeholder="조치 사유를 입력하세요"
                    className="mb-2 w-full rounded-lg border border-border px-3 py-2 text-sm focus:border-primary focus:outline-none"
                    rows={3}
                  />
                  <button
                    onClick={handleAction}
                    disabled={reportAction.isPending}
                    className="w-full rounded-lg bg-danger px-3 py-2 text-sm font-medium text-white hover:bg-red-700 disabled:opacity-50"
                  >
                    {reportAction.isPending ? "처리 중..." : "조치 실행"}
                  </button>
                </div>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
