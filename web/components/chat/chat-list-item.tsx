import type { ChatRoom } from "@/lib/types";
import { formatTimeAgo } from "@/lib/utils";

interface ChatListItemProps {
  chat: ChatRoom;
  isActive?: boolean;
  onClick: () => void;
}

export function ChatListItem({ chat, isActive, onClick }: ChatListItemProps) {
  const label = chat.unreadCount > 0
    ? `${chat.counterparty.nickname} — ${chat.unreadCount}개 안 읽은 메시지`
    : chat.counterparty.nickname;

  return (
    <button
      onClick={onClick}
      aria-current={isActive ? "true" : undefined}
      aria-label={label}
      className={`w-full text-left p-3 border-b border-border transition-colors ${
        isActive ? "bg-[rgba(74,127,181,0.15)] border-l-2 border-l-gold" : "hover:bg-medium"
      }`}
    >
      <div className="flex items-center justify-between mb-1">
        <span className="font-semibold text-sm text-text-primary">{chat.counterparty.nickname}</span>
        <div className="flex items-center gap-1.5">
          {chat.chatStatus !== "open" && (
            <span className={`text-xs px-1.5 py-0.5 rounded inline-block ${
              chat.chatStatus === "reservation_proposed" ? "bg-yellow-600/20 text-yellow-400" :
              chat.chatStatus === "reservation_confirmed" ? "bg-green-600/20 text-green-400" :
              chat.chatStatus === "deal_completed" ? "bg-gold/20 text-gold" : ""
            }`}>
              {chat.chatStatus === "reservation_proposed" && "예약 제안"}
              {chat.chatStatus === "reservation_confirmed" && "예약 확정"}
              {chat.chatStatus === "deal_completed" && "거래 완료"}
            </span>
          )}
          <span className="text-xs text-text-secondary">
            {chat.lastMessage ? formatTimeAgo(chat.lastMessage.sentAt) : ""}
          </span>
        </div>
      </div>
      <div className="text-xs text-text-secondary truncate">{chat.listingTitle}</div>
      {chat.lastMessage && (
        <div className="text-sm text-text-secondary truncate mt-1">{chat.lastMessage.bodyText}</div>
      )}
      {chat.unreadCount > 0 && (
        <span aria-hidden="true" className="inline-block mt-1 bg-gold text-darkest text-xs px-1.5 py-0.5 rounded-full">
          {chat.unreadCount}
        </span>
      )}
    </button>
  );
}
