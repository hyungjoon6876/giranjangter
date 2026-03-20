import type { Message } from "@/lib/types";

function formatMessageTime(sentAt: string): string {
  const date = new Date(sentAt);
  const now = new Date();
  const timeStr = date.toLocaleTimeString("ko-KR", {
    hour: "numeric",
    minute: "2-digit",
    hour12: true,
  });

  if (
    date.getFullYear() === now.getFullYear() &&
    date.getMonth() === now.getMonth() &&
    date.getDate() === now.getDate()
  ) {
    return timeStr; // Today: "오후 3:42"
  }

  if (date.getFullYear() === now.getFullYear()) {
    return `${date.getMonth() + 1}월 ${date.getDate()}일 ${timeStr}`;
  }

  return `${date.getFullYear()}.${date.getMonth() + 1}.${date.getDate()} ${timeStr}`;
}

interface ChatMessageProps {
  message: Message;
  isMine: boolean;
  onRetry?: (message: Message) => void;
}

export function ChatMessage({ message, isMine, onRetry }: ChatMessageProps) {
  if (message.messageType === "system") {
    return (
      <div className="flex justify-center my-2" role="status">
        <span className="bg-medium text-text-secondary text-xs px-3 py-1.5 rounded-full">
          {message.bodyText}
        </span>
      </div>
    );
  }

  const isSending = message.status === "sending";
  const isFailed = message.status === "failed";

  return (
    <div className={`flex mb-1 ${isMine ? "justify-end" : "justify-start"} ${isSending ? "opacity-60" : ""}`}>
      <div
        className={`max-w-[70%] px-4 py-2.5 rounded-2xl text-sm ${
          isMine
            ? "bg-blue-bright text-white rounded-br-sm"
            : "bg-card border border-border text-text-primary rounded-bl-sm"
        }`}
      >
        {message.bodyText}
        <span
          className={`block text-xs mt-1 ${isMine ? "text-right" : "text-left"} ${isFailed ? "text-danger" : "text-text-secondary"}`}
        >
          {isSending ? "전송 중..." : isFailed ? (
            <button
              onClick={() => onRetry?.(message)}
              className="hover:underline"
            >
              전송 실패 · 재전송
            </button>
          ) : isMine ? (
            `${formatMessageTime(message.sentAt)} ✓`
          ) : formatMessageTime(message.sentAt)}
        </span>
      </div>
    </div>
  );
}
