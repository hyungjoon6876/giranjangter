/** Format ISO date string to Korean locale date. */
export function formatDate(iso: string): string {
  return new Date(iso).toLocaleDateString("ko-KR", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
}

/** Format ISO date string to relative time ago in Korean. */
export function formatTimeAgo(iso: string): string {
  const diff = Date.now() - new Date(iso).getTime();
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

/** Return Tailwind class string for report status badges. */
export function statusBadgeVariant(
  status: string,
): string {
  switch (status) {
    // Report statuses
    case "submitted":
      return "bg-yellow-100 text-yellow-800";
    case "assigned":
      return "bg-blue-100 text-blue-800";
    case "resolved":
      return "bg-green-100 text-green-800";

    // User account statuses
    case "active":
      return "bg-green-100 text-green-800";
    case "restricted":
      return "bg-orange-100 text-orange-800";
    case "suspended":
      return "bg-red-100 text-red-800";

    // Listing statuses
    case "available":
      return "bg-green-100 text-green-800";
    case "reserved":
      return "bg-yellow-100 text-yellow-800";
    case "pending_trade":
      return "bg-blue-100 text-blue-800";
    case "completed":
      return "bg-gray-100 text-gray-600";
    case "cancelled":
      return "bg-red-100 text-red-800";

    // Listing visibility
    case "public":
      return "bg-green-100 text-green-800";
    case "hidden":
      return "bg-red-100 text-red-800";

    // Trade statuses
    case "pending":
      return "bg-yellow-100 text-yellow-800";
    case "confirmed":
      return "bg-green-100 text-green-800";
    case "disputed":
      return "bg-red-100 text-red-800";

    default:
      return "bg-gray-100 text-gray-600";
  }
}
