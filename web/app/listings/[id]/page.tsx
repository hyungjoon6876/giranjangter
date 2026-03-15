"use client";

import { use } from "react";
import { useRouter } from "next/navigation";
import { useListing, useToggleFavorite } from "@/lib/hooks/use-listings";
import { useCreateChat } from "@/lib/hooks/use-chats";
import { TypeBadge, Badge } from "@/components/ui/badge";
import { AuthorSection, InfoRow, tradeMethodLabel } from "@/components/listing/listing-info";
import { Loading } from "@/components/ui/loading";
import { formatPrice, statusLabel, statusColor } from "@/lib/utils";

export default function ListingDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const router = useRouter();
  const { data: listing, isLoading } = useListing(id);
  const toggleFav = useToggleFavorite();
  const createChat = useCreateChat();

  if (isLoading) return <Loading />;
  if (!listing) return <div className="p-6 text-center text-text-secondary">매물을 찾을 수 없습니다</div>;

  const l = listing;
  const actions = l.availableActions ?? [];

  const handleChat = async () => {
    try {
      const chat = await createChat.mutateAsync(l.listingId);
      router.push(`/chats/${chat.chatRoomId}`);
    } catch {
      alert("채팅을 시작할 수 없습니다");
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-4 lg:p-6">
      {/* Header badges */}
      <div className="flex items-center gap-2 mb-4">
        <TypeBadge type={l.listingType} />
        <Badge label={statusLabel(l.status)} color={statusColor(l.status)} />
        {l.tradeMethod && (
          <span className="ml-auto text-sm text-text-secondary">{tradeMethodLabel(l.tradeMethod)}</span>
        )}
      </div>

      {/* Title */}
      <h1 className="text-2xl font-bold mb-3">{l.title}</h1>

      {/* Item info */}
      <div className="flex items-center gap-2 text-lg mb-4">
        {l.iconUrl && (
          <img
            src={`${process.env.NEXT_PUBLIC_API_URL?.replace("/api/v1", "") ?? "http://localhost:8080"}${l.iconUrl}`}
            alt=""
            className="w-8 h-8"
          />
        )}
        <span>{l.itemName}</span>
        {l.enhancementLevel != null && (
          <span className="text-primary font-bold">+{l.enhancementLevel}</span>
        )}
      </div>
      {l.optionsText && <p className="text-text-secondary mb-4">{l.optionsText}</p>}

      {/* Price */}
      <div className="text-3xl font-bold mb-1">{formatPrice(l.priceAmount)}원</div>
      {l.priceType === "negotiable" && <p className="text-text-secondary mb-4">협상 가능</p>}

      <hr className="border-border my-6" />

      {/* Description */}
      <p className="leading-relaxed whitespace-pre-wrap">{l.description}</p>

      <hr className="border-border my-6" />

      {/* Trade info */}
      <InfoRow label="거래 방식" value={tradeMethodLabel(l.tradeMethod)} />
      {l.preferredMeetingAreaText && <InfoRow label="접선 장소" value={l.preferredMeetingAreaText} />}
      {l.availableTimeText && <InfoRow label="거래 가능 시간" value={l.availableTimeText} />}
      <InfoRow label="수량" value={`${l.quantity}개`} />

      <hr className="border-border my-6" />

      {/* Author */}
      {l.author && <AuthorSection author={l.author} />}

      {/* Action bar */}
      {actions.length > 0 && (
        <div className="sticky bottom-0 lg:relative bg-white border-t border-border mt-8 py-4 flex items-center gap-3">
          {actions.includes("favorite") && (
            <button
              onClick={() => toggleFav.mutate({ id: l.listingId, isFavorited: l.isFavorited ?? false })}
              className="p-3 border border-border rounded-lg hover:bg-surface transition-colors"
            >
              {l.isFavorited ? "❤️" : "🤍"}
            </button>
          )}
          {actions.includes("start_chat") && (
            <button
              onClick={handleChat}
              disabled={createChat.isPending}
              className="flex-1 bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition-colors disabled:opacity-50"
            >
              {createChat.isPending ? "연결 중..." : "채팅하기"}
            </button>
          )}
        </div>
      )}
    </div>
  );
}
