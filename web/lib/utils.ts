export function formatPrice(amount: number | null | undefined): string {
  if (amount == null) return "가격 제안";
  return amount.toLocaleString("ko-KR");
}

const STATUS_LABELS: Record<string, string> = {
  available: "판매중",
  reserved: "예약중",
  pending_trade: "거래중",
  completed: "거래완료",
  cancelled: "취소됨",
};

export function statusLabel(status: string): string {
  return STATUS_LABELS[status] ?? status;
}

const STATUS_COLORS: Record<string, string> = {
  available: "#059669",
  reserved: "#F59E0B",
  pending_trade: "#2563EB",
  completed: "#64748B",
  cancelled: "#DC2626",
};

export function statusColor(status: string): string {
  return STATUS_COLORS[status] ?? "#64748B";
}

export function formatTimeAgo(isoString: string): string {
  const diff = Date.now() - new Date(isoString).getTime();
  const minutes = Math.floor(diff / 60_000);
  if (minutes < 1) return "방금 전";
  if (minutes < 60) return `${minutes}분 전`;
  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours}시간 전`;
  const days = Math.floor(hours / 24);
  if (days < 30) return `${days}일 전`;
  const months = Math.floor(days / 30);
  return `${months}개월 전`;
}
