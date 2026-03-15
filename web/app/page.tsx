"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useListings } from "@/lib/hooks/use-listings";
import { ListingFilters } from "@/components/listing/listing-filters";
import { ListingGrid } from "@/components/listing/listing-grid";
import { ListingSkeleton } from "@/components/ui/skeleton";
import { EmptyState } from "@/components/ui/empty-state";
import Link from "next/link";

export default function HomePage() {
  const [serverId, setServerId] = useState<string | null>(null);
  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("recent");

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
  });

  const { data, isLoading } = useListings({
    serverId: serverId ?? undefined,
    q: search || undefined,
    sort,
  });

  return (
    <div className="p-4 lg:p-6">
      {/* Hero section for non-logged-in users */}
      {!apiClient.isLoggedIn && (
        <section className="relative overflow-hidden rounded-xl mb-8 p-8 lg:p-12 bg-gradient-to-br from-dark via-card to-medium border border-border">
          <div className="relative z-10">
            <h1 className="font-display text-3xl lg:text-[48px] text-gold mb-3">기란장터</h1>
            <p className="text-text-secondary text-lg mb-6">리니지 클래식 아이템 거래, 안전하고 무료</p>
            <div className="flex flex-col sm:flex-row gap-3 mb-8">
              <a href="#listings" className="btn-gold-gradient text-white px-6 py-3 rounded-lg font-medium text-center">매물 둘러보기</a>
              <Link href="/login" className="border border-gold text-gold px-6 py-3 rounded-lg font-medium text-center hover:bg-gold/10">시작하기</Link>
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div className="bg-medium/50 rounded-lg p-4 border border-border">
                <div className="text-gold font-bold text-lg mb-1">무료 거래</div>
                <div className="text-text-dim text-sm">수수료 없이 안전한 아이템 거래</div>
              </div>
              <div className="bg-medium/50 rounded-lg p-4 border border-border">
                <div className="text-gold font-bold text-lg mb-1">신뢰 시스템</div>
                <div className="text-text-dim text-sm">거래 평가와 신뢰도 뱃지</div>
              </div>
              <div className="bg-medium/50 rounded-lg p-4 border border-border">
                <div className="text-gold font-bold text-lg mb-1">실시간 채팅</div>
                <div className="text-text-dim text-sm">판매자와 바로 대화하세요</div>
              </div>
            </div>
          </div>
          {/* Decorative glow */}
          <div className="absolute top-0 right-0 w-64 h-64 bg-gold/5 rounded-full blur-3xl" />
          <div className="absolute bottom-0 left-0 w-48 h-48 bg-blue/5 rounded-full blur-3xl" />
        </section>
      )}

      <div id="listings">
        <div className="flex items-center justify-between mb-4">
          <h1 className="text-2xl font-bold hidden lg:block text-text-primary">매물 목록</h1>
          <Link
            href="/create"
            className="hidden lg:inline-flex items-center gap-2 btn-gold-gradient text-white px-4 py-2 rounded-lg text-sm transition-colors"
          >
            + 매물 등록
          </Link>
        </div>

        <ListingFilters
          servers={servers}
          selectedServer={serverId}
          onServerChange={setServerId}
          searchQuery={search}
          onSearchChange={setSearch}
        />

        {isLoading ? (
          <ListingSkeleton />
        ) : !data?.data.length ? (
          search ? (
            <EmptyState
              title="검색 결과가 없습니다"
              description={`"${search}"에 대한 매물을 찾을 수 없습니다`}
            />
          ) : serverId ? (
            <EmptyState
              title="해당 서버에 매물이 없습니다"
              description="다른 서버를 선택하거나 첫 매물을 등록해보세요!"
              actionLabel="매물 등록"
              actionHref="/create"
            />
          ) : (
            <EmptyState
              title="매물이 없습니다"
              description="첫 매물을 등록해보세요!"
              actionLabel="매물 등록"
              actionHref="/create"
            />
          )
        ) : (
          <>
            <div className="flex items-center justify-between px-4 lg:px-6 py-2">
              <p className="text-sm text-text-secondary" aria-live="polite">
                <span className="font-semibold text-text-primary">{data.data.length}</span>개 매물
              </p>
              <select
                aria-label="정렬 방식"
                value={sort}
                onChange={(e) => setSort(e.target.value)}
                className="bg-medium text-text-secondary text-xs border border-border rounded-md px-2 py-1"
              >
                <option value="recent">최신순</option>
                <option value="price_asc">가격 낮은순</option>
                <option value="price_desc">가격 높은순</option>
                <option value="popular">인기순</option>
              </select>
            </div>
            <ListingGrid listings={data.data} />
          </>
        )}
      </div>

      {/* Mobile FAB */}
      <Link
        href="/create"
        aria-label="매물 등록"
        className="lg:hidden fixed right-4 bottom-20 btn-gold-gradient text-white w-14 h-14 rounded-full flex items-center justify-center text-2xl shadow-lg z-40"
      >
        +
      </Link>
    </div>
  );
}
