"use client";

import { useMyTrades } from "@/lib/hooks/use-profile";
import { ListingGrid } from "@/components/listing/listing-grid";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";

export default function MyTradesPage() {
  const { data, isLoading } = useMyTrades();

  if (isLoading) return <Loading />;
  if (!data?.data?.length) return <EmptyState title="거래 내역이 없습니다" />;

  return (
    <div className="p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4">내 거래</h1>
      <ListingGrid listings={data.data ?? []} />
    </div>
  );
}
