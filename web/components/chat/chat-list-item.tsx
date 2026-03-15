import type { ChatRoom } from "@/lib/types";
import { formatTimeAgo } from "@/lib/utils";

interface ChatListItemProps {
  chat: ChatRoom;
  isActive?: boolean;
  onClick: () => void;
}

export function ChatListItem({ chat, isActive, onClick }: ChatListItemProps) {
  return (
    <button
      onClick={onClick}
      className={`w-full text-left p-3 border-b border-border transition-colors ${
        isActive ? "bg-blue-50 border-l-2 border-l-primary" : "hover:bg-surface"
      }`}
    >
      <div className="flex items-center justify-between mb-1">
        <span className="font-semibold text-sm">{chat.counterparty.nickname}</span>
        <span className="text-xs text-text-secondary">
          {chat.lastMessage ? formatTimeAgo(chat.lastMessage.sentAt) : ""}
        </span>
      </div>
      <div className="text-xs text-text-secondary truncate">{chat.listingTitle}</div>
      {chat.lastMessage && (
        <div className="text-sm text-text-secondary truncate mt-1">{chat.lastMessage.bodyText}</div>
      )}
      {chat.unreadCount > 0 && (
        <span className="inline-block mt-1 bg-primary text-white text-xs px-1.5 py-0.5 rounded-full">
          {chat.unreadCount}
        </span>
      )}
    </button>
  );
}
