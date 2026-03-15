import Link from "next/link";
import type { Listing } from "@/lib/types";
import { TypeBadge, Badge } from "@/components/ui/badge";
import { formatPrice, formatTimeAgo, statusLabel, statusColor } from "@/lib/utils";

export function ListingCard({ listing }: { listing: Listing }) {
  const l = listing;
  return (
    <Link
      href={`/listings/${l.listingId}`}
      className="block bg-white border border-border rounded-xl p-4 hover:shadow-md transition-shadow"
    >
      <div className="flex items-center gap-2 mb-2">
        <TypeBadge type={l.listingType} />
        <Badge label={statusLabel(l.status)} color={statusColor(l.status)} />
        <span className="ml-auto text-sm text-text-secondary">{l.serverName}</span>
      </div>
      <h3 className="font-semibold text-text-primary truncate">{l.title}</h3>
      <div className="flex items-center gap-1 mt-1 text-sm text-text-secondary">
        {l.iconUrl && (
          <img
            src={`${process.env.NEXT_PUBLIC_API_URL?.replace("/api/v1", "") ?? "http://localhost:8080"}${l.iconUrl}`}
            alt=""
            className="w-5 h-5"
          />
        )}
        <span>{l.itemName}</span>
        {l.enhancementLevel != null && (
          <span className="text-primary font-semibold">+{l.enhancementLevel}</span>
        )}
      </div>
      <div className="flex items-center mt-2">
        <span className="text-lg font-bold">{formatPrice(l.priceAmount)}원</span>
        {l.priceType === "negotiable" && (
          <span className="text-xs text-text-secondary ml-1">(협상가능)</span>
        )}
        <div className="ml-auto flex items-center gap-3 text-xs text-text-secondary">
          <span>👁 {l.viewCount}</span>
          <span>♥ {l.favoriteCount}</span>
          <span>💬 {l.chatCount}</span>
        </div>
      </div>
      <div className="flex items-center justify-between mt-2 text-xs text-text-secondary">
        <span>{l.author.nickname}</span>
        <span>{formatTimeAgo(l.createdAt)}</span>
      </div>
    </Link>
  );
}
