"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useListings } from "@/lib/hooks/use-listings";
import { ListingFilters } from "@/components/listing/listing-filters";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import Link from "next/link";

export default function HomePage() {
  const [serverId, setServerId] = useState<string | null>(null);
  const [search, setSearch] = useState("");

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
  });

  const { data, isLoading } = useListings({
    serverId: serverId ?? undefined,
    q: search || undefined,
  });

  return (
    <div className="p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold hidden lg:block">매물 목록</h1>
        <Link
          href="/create"
          className="hidden lg:inline-flex items-center gap-2 bg-primary text-white px-4 py-2 rounded-lg text-sm hover:bg-primary-dark transition-colors"
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
        <Loading />
      ) : !data?.data.length ? (
        <EmptyState title="매물이 없습니다" description="첫 매물을 등록해보세요!" />
      ) : (
        <ListingGrid listings={data.data} />
      )}

      {/* Mobile FAB */}
      <Link
        href="/create"
        className="lg:hidden fixed right-4 bottom-20 bg-primary text-white w-14 h-14 rounded-full flex items-center justify-center text-2xl shadow-lg hover:bg-primary-dark z-40"
      >
        +
      </Link>
    </div>
  );
}
