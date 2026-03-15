"use client";

import { useState, useMemo } from "react";
import { type ColumnDef } from "@tanstack/react-table";
import { DataTable } from "@/components/data-table/data-table";
import { useListings, useHideListing, useRestoreListing } from "@/lib/hooks/use-audit";
import { statusBadgeVariant, formatDate } from "@/lib/utils";
import type { AdminListing } from "@/lib/types";

const STATUS_OPTIONS = [
  { label: "전체", value: "" },
  { label: "판매중", value: "available" },
  { label: "예약중", value: "reserved" },
  { label: "완료", value: "completed" },
  { label: "취소", value: "cancelled" },
] as const;

const VISIBILITY_OPTIONS = [
  { label: "전체", value: "" },
  { label: "공개", value: "public" },
  { label: "숨김", value: "hidden" },
] as const;

export default function ListingsPage() {
  const [statusFilter, setStatusFilter] = useState("");
  const [visibilityFilter, setVisibilityFilter] = useState("");

  const { data: listings, isLoading } = useListings({
    status: statusFilter || undefined,
    visibility: visibilityFilter || undefined,
  });

  const hideListing = useHideListing();
  const restoreListing = useRestoreListing();

  const columns: ColumnDef<AdminListing, unknown>[] = useMemo(
    () => [
      {
        accessorKey: "title",
        header: "제목",
        cell: ({ getValue }) => (
          <span className="max-w-xs truncate font-medium">
            {getValue() as string}
          </span>
        ),
      },
      { accessorKey: "itemName", header: "아이템" },
      {
        accessorKey: "listingType",
        header: "유형",
        cell: ({ getValue }) => (
          <span className="text-xs">{getValue() as string}</span>
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
        accessorKey: "visibility",
        header: "공개여부",
        cell: ({ getValue }) => {
          const v = getValue() as string;
          return (
            <span
              className={`inline-block rounded-full px-2 py-1 text-xs font-medium ${statusBadgeVariant(v)}`}
            >
              {v}
            </span>
          );
        },
      },
      { accessorKey: "authorNickname", header: "판매자" },
      {
        accessorKey: "createdAt",
        header: "등록일",
        cell: ({ getValue }) => formatDate(getValue() as string),
      },
      {
        id: "actions",
        header: "",
        cell: ({ row }) => {
          const listing = row.original;
          const isHidden = listing.visibility === "hidden";
          return (
            <button
              onClick={(e) => {
                e.stopPropagation();
                if (isHidden) {
                  restoreListing.mutate(listing.listingId);
                } else {
                  hideListing.mutate(listing.listingId);
                }
              }}
              className={`rounded-lg px-3 py-1.5 text-xs font-medium ${
                isHidden
                  ? "border border-green-300 text-green-700 hover:bg-green-50"
                  : "border border-red-300 text-red-700 hover:bg-red-50"
              }`}
            >
              {isHidden ? "복원" : "숨기기"}
            </button>
          );
        },
      },
    ],
    [hideListing, restoreListing],
  );

  return (
    <div className="space-y-4">
      {/* Filters */}
      <div className="flex flex-wrap items-center gap-3">
        <div className="flex gap-1 rounded-lg border border-border bg-white p-1">
          {STATUS_OPTIONS.map((opt) => (
            <button
              key={opt.value}
              onClick={() => setStatusFilter(opt.value)}
              className={`rounded-md px-3 py-2 text-sm font-medium transition-colors ${
                statusFilter === opt.value
                  ? "bg-primary text-white"
                  : "text-text-secondary hover:bg-slate-100"
              }`}
            >
              {opt.label}
            </button>
          ))}
        </div>
        <div className="flex gap-1 rounded-lg border border-border bg-white p-1">
          {VISIBILITY_OPTIONS.map((opt) => (
            <button
              key={opt.value}
              onClick={() => setVisibilityFilter(opt.value)}
              className={`rounded-md px-3 py-2 text-sm font-medium transition-colors ${
                visibilityFilter === opt.value
                  ? "bg-primary text-white"
                  : "text-text-secondary hover:bg-slate-100"
              }`}
            >
              {opt.label}
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
            data={listings ?? []}
            columns={columns}
            emptyMessage="매물이 없습니다."
          />
        )}
      </div>
    </div>
  );
}
