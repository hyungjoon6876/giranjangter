"use client";

import { useState } from "react";
import Link from "next/link";
import { useMyListings } from "@/lib/hooks/use-profile";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

const STATUS_FILTERS = [
  { value: undefined, label: "전체" },
  { value: "available", label: "판매중" },
  { value: "reserved", label: "예약중" },
  { value: "completed", label: "완료" },
  { value: "cancelled", label: "취소" },
] as const;

export default function MyListingsPage() {
  const [statusFilter, setStatusFilter] = useState<string | undefined>(
    undefined,
  );
  const { data, isLoading } = useMyListings(statusFilter);

  if (isLoading) return <Loading />;

  const listings = data?.data ?? [];

  return (
    <div className="p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold text-text-primary">내 매물</h1>
        <Link
          href="/create"
          className="btn-gold-gradient text-white px-4 py-2 rounded-lg text-sm"
        >
          + 등록
        </Link>
      </div>

      {/* Status filter */}
      <div
        role="group"
        aria-label="상태 필터"
        className="flex flex-wrap gap-2 mb-4"
      >
        {STATUS_FILTERS.map((f) => (
          <button
            key={f.label}
            onClick={() => setStatusFilter(f.value)}
            aria-pressed={statusFilter === f.value}
            className={`px-3 py-1.5 rounded-lg text-xs transition-colors border ${
              statusFilter === f.value
                ? "border-gold text-gold bg-gold/10"
                : "border-border text-text-secondary bg-medium hover:bg-light"
            }`}
          >
            {f.label}
          </button>
        ))}
      </div>

      {!listings.length ? (
        <EmptyState
          title="등록한 매물이 없습니다"
          actionLabel="매물 등록하기"
          actionHref="/create"
        />
      ) : (
        <ListingGrid listings={listings} />
      )}
    </div>
  );
}
