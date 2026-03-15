import type { Message } from "@/lib/types";

interface ChatMessageProps {
  message: Message;
  isMine: boolean;
}

export function ChatMessage({ message, isMine }: ChatMessageProps) {
  if (message.messageType === "system") {
    return (
      <div className="flex justify-center my-2">
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
      </div>
    </div>
  );
}
