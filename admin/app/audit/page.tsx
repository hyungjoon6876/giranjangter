"use client";

import { useMemo } from "react";
import { type ColumnDef } from "@tanstack/react-table";
import { DataTable } from "@/components/data-table/data-table";
import { useAuditLogs } from "@/lib/hooks/use-audit";
import { formatDate } from "@/lib/utils";
import type { AuditLog } from "@/lib/types";

export default function AuditLogsPage() {
  const { data: logs, isLoading } = useAuditLogs();

  const columns: ColumnDef<AuditLog, unknown>[] = useMemo(
    () => [
      {
        accessorKey: "createdAt",
        header: "시간",
        cell: ({ getValue }) => (
          <span className="text-xs">{formatDate(getValue() as string)}</span>
        ),
      },
      {
        accessorKey: "actorId",
        header: "관리자",
        cell: ({ getValue }) => {
          const v = getValue() as string | null;
          return (
            <span className="font-mono text-xs">
              {v ? v.slice(0, 8) : "-"}
            </span>
          );
        },
      },
      {
        accessorKey: "actorRole",
        header: "역할",
        cell: ({ getValue }) => (
          <span className="text-xs font-medium uppercase text-text-secondary">
            {(getValue() as string | null) ?? "-"}
          </span>
        ),
      },
      {
        accessorKey: "action",
        header: "액션",
        cell: ({ getValue }) => (
          <span className="rounded bg-slate-100 px-2 py-0.5 font-mono text-xs">
            {getValue() as string}
          </span>
        ),
      },
      {
        id: "target",
        header: "대상",
        cell: ({ row }) => (
          <span className="text-xs">
            {row.original.targetType}:{row.original.targetId.slice(0, 8)}
          </span>
        ),
      },
      {
        accessorKey: "details",
        header: "상세",
        cell: ({ getValue }) => {
          const v = getValue() as string | null;
          return (
            <span className="max-w-xs truncate text-xs text-text-secondary">
              {v ?? "-"}
            </span>
          );
        },
      },
    ],
    [],
  );

  return (
    <div className="space-y-4">
      <div className="rounded-xl border border-border bg-white shadow-sm">
        {isLoading ? (
          <div className="flex h-40 items-center justify-center text-text-secondary">
            로딩 중...
          </div>
        ) : (
          <DataTable
            data={logs ?? []}
            columns={columns}
            emptyMessage="감사 로그가 없습니다."
          />
        )}
      </div>
    </div>
  );
}
