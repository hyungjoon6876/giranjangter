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

export function computeGroupFlags(messages: Message[]) {
  return messages.map((msg, i) => {
    const prev = messages[i - 1];
    const next = messages[i + 1];

    const sameAsPrev =
      prev &&
      prev.senderUserId === msg.senderUserId &&
      prev.messageType !== "system" &&
      msg.messageType !== "system" &&
      new Date(msg.sentAt).getTime() - new Date(prev.sentAt).getTime() < 60_000;

    const sameAsNext =
      next &&
      next.senderUserId === msg.senderUserId &&
      next.messageType !== "system" &&
      msg.messageType !== "system" &&
      new Date(next.sentAt).getTime() - new Date(msg.sentAt).getTime() < 60_000;

    return {
      ...msg,
      isFirstInGroup: !sameAsPrev,
      isLastInGroup: !sameAsNext,
    };
  });
}

interface ChatMessageProps {
  message: Message;
  isMine: boolean;
  onRetry?: (message: Message) => void;
  isFirstInGroup?: boolean;
  isLastInGroup?: boolean;
}

export function ChatMessage({ message, isMine, onRetry, isFirstInGroup = true, isLastInGroup = true }: ChatMessageProps) {
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

  const cornerClass = isMine
    ? `${isFirstInGroup ? "rounded-tr-xl" : "rounded-tr"} ${isLastInGroup ? "rounded-br-sm" : "rounded-br"} rounded-tl-xl rounded-bl-xl`
    : `${isFirstInGroup ? "rounded-tl-xl" : "rounded-tl"} ${isLastInGroup ? "rounded-bl-sm" : "rounded-bl"} rounded-tr-xl rounded-br-xl`;

  return (
    <div className={`flex ${isFirstInGroup ? "mt-2" : "mt-0.5"} ${isMine ? "justify-end" : "justify-start"} ${isSending ? "opacity-60" : ""}`}>
      <div
        className={`max-w-[70%] px-4 py-2.5 text-sm ${cornerClass} ${
          isMine
            ? "bg-blue-bright text-white"
            : "bg-card border border-border text-text-primary"
        }`}
      >
        {message.bodyText}
        {(isSending || isFailed) && (
          <span
            className={`block text-xs mt-1 ${isMine ? "text-right" : "text-left"} ${isFailed ? "text-danger" : "text-text-secondary"}`}
          >
            {isSending ? "전송 중..." : (
              <button
                onClick={() => onRetry?.(message)}
                className="hover:underline"
              >
                전송 실패 · 재전송
              </button>
            )}
          </span>
        )}
        {isLastInGroup && !isSending && !isFailed && (
          <span
            className={`block text-xs mt-1 ${isMine ? "text-right" : "text-left"} text-text-secondary`}
          >
            {isMine
              ? `${formatMessageTime(message.sentAt)} ✓`
              : formatMessageTime(message.sentAt)}
          </span>
        )}
      </div>
    </div>
  );
}
