"use client";

import { useMemo } from "react";
import { type ColumnDef } from "@tanstack/react-table";
import { DataTable } from "@/components/data-table/data-table";
import { useTrades } from "@/lib/hooks/use-audit";
import { statusBadgeVariant, formatDate } from "@/lib/utils";
import type { TradeCompletion } from "@/lib/types";

export default function TradesPage() {
  const { data: trades, isLoading } = useTrades();

  const columns: ColumnDef<TradeCompletion, unknown>[] = useMemo(
    () => [
      {
        accessorKey: "listingTitle",
        header: "매물 제목",
        cell: ({ getValue }) => (
          <span className="max-w-xs truncate font-medium">
            {getValue() as string}
          </span>
        ),
      },
      {
        accessorKey: "status",
        header: "상태",
        cell: ({ getValue }) => {
          const status = getValue() as string;
          const label =
            status === "pending_confirm"
              ? "확인대기"
              : status === "confirmed"
                ? "확인완료"
                : status;
          return (
            <span
              className={`inline-block rounded-full px-2 py-1 text-xs font-medium ${statusBadgeVariant(status)}`}
            >
              {label}
            </span>
          );
        },
      },
      {
        accessorKey: "requestedByUserId",
        header: "요청자",
        cell: ({ getValue }) => (
          <span className="font-mono text-xs">
            {(getValue() as string).slice(0, 8)}
          </span>
        ),
      },
      {
        accessorKey: "counterpartUserId",
        header: "상대방",
        cell: ({ getValue }) => (
          <span className="font-mono text-xs">
            {(getValue() as string).slice(0, 8)}
          </span>
        ),
      },
      {
        accessorKey: "autoConfirmAt",
        header: "자동확인 시한",
        cell: ({ getValue }) => {
          const v = getValue() as string | undefined;
          if (!v) return "-";
          const deadline = new Date(v);
          const now = new Date();
          const remaining = deadline.getTime() - now.getTime();
          if (remaining <= 0) {
            return <span className="text-danger text-xs font-medium">만료</span>;
          }
          const hours = Math.floor(remaining / 3_600_000);
          const minutes = Math.floor((remaining % 3_600_000) / 60_000);
          return (
            <span className="text-xs">
              {hours > 0 ? `${hours}시간 ` : ""}
              {minutes}분 남음
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

  return (
    <div className="space-y-4">
      <div className="rounded-xl border border-border bg-white shadow-sm">
        {isLoading ? (
          <div className="flex h-40 items-center justify-center text-text-secondary">
            로딩 중...
          </div>
        ) : (
          <DataTable
            data={trades ?? []}
            columns={columns}
            emptyMessage="거래 내역이 없습니다."
          />
        )}
      </div>
    </div>
  );
}
