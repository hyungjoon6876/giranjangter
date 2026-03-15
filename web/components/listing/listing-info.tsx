import type { Author } from "@/lib/types";

export function InfoRow({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex py-2">
      <span className="w-28 text-text-secondary flex-shrink-0">{label}</span>
      <span>{value}</span>
    </div>
  );
}

export function AuthorSection({ author }: { author: Author }) {
  return (
    <div className="flex items-center gap-3">
      <div className="w-10 h-10 rounded-full bg-border flex items-center justify-center font-bold text-text-secondary">
        {author.nickname?.[0] ?? "?"}
      </div>
      <div>
        <div className="font-semibold">{author.nickname}</div>
        <div className="text-sm text-text-secondary">
          거래 {author.completedTradeCount ?? 0}회
          {author.trustBadge && ` · ${author.trustBadge}`}
        </div>
      </div>
    </div>
  );
}

const TRADE_METHOD_LABELS: Record<string, string> = {
  in_game: "인게임",
  offline_pc_bang: "PC방/오프라인",
  either: "무관",
};

export function tradeMethodLabel(method?: string): string {
  return TRADE_METHOD_LABELS[method ?? ""] ?? method ?? "";
}
