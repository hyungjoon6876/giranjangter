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
}

export function ChatMessage({ message, isMine }: ChatMessageProps) {
  if (message.messageType === "system") {
    return (
      <div className="flex justify-center my-2" role="status">
        <span className="bg-medium text-text-secondary text-xs px-3 py-1.5 rounded-full">
          {message.bodyText}
        </span>
      </div>
    );
  }

  return (
    <div className={`flex mb-1 ${isMine ? "justify-end" : "justify-start"}`}>
      <div
        className={`max-w-[70%] px-4 py-2.5 rounded-2xl text-sm ${
          isMine
            ? "bg-blue-bright text-white rounded-br-sm"
            : "bg-card border border-border text-text-primary rounded-bl-sm"
        }`}
      >
        {message.bodyText}
        <span
          className={`block text-xs text-text-secondary mt-1 ${isMine ? "text-right" : "text-left"}`}
        >
          {formatMessageTime(message.sentAt)}
        </span>
      </div>
    </div>
  );
}
