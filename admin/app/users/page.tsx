"use client";

import { useState, useMemo } from "react";
import { useRouter } from "next/navigation";
import { type ColumnDef } from "@tanstack/react-table";
import { DataTable } from "@/components/data-table/data-table";
import { useUsers } from "@/lib/hooks/use-users";
import { statusBadgeVariant, formatDate } from "@/lib/utils";
import type { AdminUser } from "@/lib/types";

const STATUS_TABS = [
  { label: "전체", value: "" },
  { label: "활성", value: "active" },
  { label: "제한", value: "restricted" },
  { label: "정지", value: "suspended" },
] as const;

export default function UsersPage() {
  const router = useRouter();
  const [search, setSearch] = useState("");
  const [statusFilter, setStatusFilter] = useState("");
  const [debouncedSearch, setDebouncedSearch] = useState("");

  const { data: users, isLoading } = useUsers({
    q: debouncedSearch || undefined,
    status: statusFilter || undefined,
  });

  // Simple debounce for search
  const handleSearchChange = (value: string) => {
    setSearch(value);
    const timeout = setTimeout(() => setDebouncedSearch(value), 300);
    return () => clearTimeout(timeout);
  };

  const columns: ColumnDef<AdminUser, unknown>[] = useMemo(
    () => [
      { accessorKey: "nickname", header: "닉네임" },
      {
        accessorKey: "role",
        header: "역할",
        cell: ({ getValue }) => (
          <span className="text-xs font-medium uppercase text-text-secondary">
            {getValue() as string}
          </span>
        ),
      },
      {
        accessorKey: "accountStatus",
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
      { accessorKey: "completedTradeCount", header: "거래수" },
      {
        accessorKey: "alignmentScore",
        header: "성향점수",
        cell: ({ row }) => (
          <span>
            {row.original.alignmentScore.toFixed(1)}{" "}
            <span className="text-xs text-text-secondary">
              ({row.original.alignmentGrade})
            </span>
          </span>
        ),
      },
      {
        accessorKey: "lastLoginAt",
        header: "마지막 로그인",
        cell: ({ getValue }) => {
          const v = getValue() as string | undefined;
          return v ? formatDate(v) : "-";
        },
      },
      {
        accessorKey: "createdAt",
        header: "가입일",
        cell: ({ getValue }) => formatDate(getValue() as string),
      },
    ],
    [],
  );

  return (
    <div className="space-y-4">
      {/* Search + Filters */}
      <div className="flex flex-wrap items-center gap-3">
        <input
          type="text"
          value={search}
          onChange={(e) => handleSearchChange(e.target.value)}
          placeholder="닉네임 검색..."
          className="rounded-lg border border-border bg-white px-4 py-2 text-sm focus:border-primary focus:outline-none"
        />
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
      </div>

      {/* Table */}
      <div className="rounded-xl border border-border bg-white shadow-sm">
        {isLoading ? (
          <div className="flex h-40 items-center justify-center text-text-secondary">
            로딩 중...
          </div>
        ) : (
          <DataTable
            data={users ?? []}
            columns={columns}
            onRowClick={(user) => router.push(`/users/${user.userId}`)}
            emptyMessage="사용자가 없습니다."
          />
        )}
      </div>
    </div>
  );
}
