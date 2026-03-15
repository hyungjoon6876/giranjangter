"use client";

import { useMyListings } from "@/lib/hooks/use-profile";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function MyListingsPage() {
  const { data, isLoading } = useMyListings();

  if (isLoading) return <Loading />;
  if (!data?.data?.length) return <EmptyState title="등록한 매물이 없습니다" actionLabel="매물 등록하기" actionHref="/create" />;

  return (
    <div className="p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4">내 매물</h1>
      <ListingGrid listings={data.data ?? []} />
    </div>
  );
}
