import type { Message } from "@/lib/types";

interface ReservationCardProps {
  message: Message;
  isMine: boolean;
}

export function ReservationCardMessage({ message, isMine }: ReservationCardProps) {
  const meta = message.metadataJson ?? {};
  const meetingType = meta.meetingType === "offline_pc_bang" ? "오프라인 PC방" : "인게임 거래";

  return (
    <div className={`flex mb-2 ${isMine ? "justify-end" : "justify-start"}`}>
      <div className="max-w-[70%] border border-gold rounded-lg p-4 bg-card">
        <p className="text-xs text-gold font-medium mb-2">예약 제안</p>
        <div className="space-y-1 text-sm">
          {meta.date ? <p>{String(meta.date)} {meta.time ? String(meta.time) : ""}</p> : null}
          <p>{meetingType}</p>
          {meta.meetingPoint ? <p>{String(meta.meetingPoint)}</p> : null}
          {meta.notes ? <p className="text-text-secondary text-xs mt-2">{String(meta.notes)}</p> : null}
        </div>
        <p className="text-xs text-text-secondary mt-3">
          {isMine ? "예약 제안을 보냈습니다" : "예약 제안을 받았습니다"}
        </p>
      </div>
    </div>
  );
}
