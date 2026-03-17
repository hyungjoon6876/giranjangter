"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useIsLoggedIn } from "@/lib/hooks/use-auth";
import { useListings } from "@/lib/hooks/use-listings";
import { ListingFilters } from "@/components/listing/listing-filters";
import { ListingGrid } from "@/components/listing/listing-grid";
import { ListingSkeleton } from "@/components/ui/skeleton";
import { EmptyState } from "@/components/ui/empty-state";
import { ErrorState } from "@/components/ui/error-state";
import Link from "next/link";

export default function HomePage() {
  const isLoggedIn = useIsLoggedIn();
  const [serverId, setServerId] = useState<string | null>(null);
  const [categoryId, setCategoryId] = useState<string | null>(null);
  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("recent");

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
  });

  const { data: categories = [] } = useQuery({
    queryKey: ["categories"],
    queryFn: () => apiClient.getCategories(),
  });

  const { data, isLoading, isError, refetch } = useListings({
    serverId: serverId ?? undefined,
    categoryId: categoryId ?? undefined,
    q: search || undefined,
    sort,
  });

  return (
    <div className="p-4 lg:p-6">
      {/* Compact hero for non-logged-in users */}
      {!isLoggedIn && (
        <section className="relative overflow-hidden rounded-xl mb-6 p-6 lg:p-8 bg-gradient-to-br from-dark via-card to-medium border border-border">
          <div className="relative z-10 flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
            <div>
              <img
                src="/logo.png"
                alt="기란JT"
                className="h-12 lg:h-14 mb-2"
              />
              <p className="text-text-secondary">
                리니지 클래식 아이템 거래, 안전하고 무료
              </p>
            </div>
            <div className="flex gap-3">
              <a
                href="#listings"
                className="btn-gold-gradient text-white px-5 py-2.5 rounded-lg font-medium text-sm text-center"
              >
                매물 둘러보기
              </a>
              <Link
                href="/login"
                className="border border-gold text-gold px-5 py-2.5 rounded-lg font-medium text-sm text-center hover:bg-gold/10"
              >
                시작하기
              </Link>
            </div>
          </div>
          <div className="absolute top-0 right-0 w-48 h-48 bg-gold/5 rounded-full blur-3xl" />
        </section>
      )}

      <div id="listings">
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-2xl font-bold hidden lg:block text-text-primary">
            매물 목록
          </h2>
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
          categories={categories}
          selectedCategory={categoryId}
          onCategoryChange={setCategoryId}
          searchQuery={search}
          onSearchChange={setSearch}
        />

        {isError ? (
          <ErrorState
            message="매물을 불러올 수 없습니다"
            description="네트워크 연결을 확인해주세요"
            onRetry={() => refetch()}
            autoFocus
          />
        ) : isLoading ? (
          <ListingSkeleton />
        ) : !data?.data?.length ? (
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
                <span className="font-semibold text-text-primary">
                  {data.data?.length ?? 0}
                </span>
                개 매물
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
            <ListingGrid listings={data.data ?? []} />
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
