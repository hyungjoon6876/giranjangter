"use client";

import Link from "next/link";
import { useMyTrades } from "@/lib/hooks/use-profile";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { formatTimeAgo } from "@/lib/utils";

interface Trade {
  chatRoomId: string;
  listingId: string;
  listingTitle: string;
  tradeStatus: string;
  chatStatus: string;
  counterparty: { userId: string; nickname: string };
  updatedAt: string;
}

const chatStatusLabel: Record<string, string> = {
  open: "진행중",
  reservation_proposed: "예약 제안",
  reservation_confirmed: "예약 확정",
  deal_completed: "거래 완료",
};

const chatStatusColor: Record<string, string> = {
  open: "bg-medium text-text-secondary",
  reservation_proposed: "bg-yellow-600/20 text-yellow-400",
  reservation_confirmed: "bg-green-600/20 text-green-400",
  deal_completed: "bg-gold/20 text-gold",
};

export default function MyTradesPage() {
  const { data, isLoading } = useMyTrades();

  if (isLoading) return <Loading />;
  if (!data?.data?.length) return <EmptyState title="거래 내역이 없습니다" />;

  const trades = data.data as unknown as Trade[];

  return (
    <div className="p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4">내 거래</h1>
      <div className="space-y-3">
        {trades.map((t) => (
          <Link
            key={t.chatRoomId}
            href={`/chats/${t.chatRoomId}`}
            className="block bg-card border border-border rounded-xl p-4 hover:border-gold/30 transition-colors"
          >
            <div className="flex items-center justify-between mb-2">
              <h3 className="font-medium truncate flex-1">{t.listingTitle}</h3>
              <span
                className={`text-xs px-2 py-0.5 rounded-full ml-2 whitespace-nowrap ${chatStatusColor[t.chatStatus] ?? chatStatusColor.open}`}
              >
                {chatStatusLabel[t.chatStatus] ?? t.chatStatus}
              </span>
            </div>
            <div className="flex items-center justify-between text-sm text-text-secondary">
              <span>{t.counterparty.nickname}</span>
              <span>{formatTimeAgo(t.updatedAt)}</span>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
}
