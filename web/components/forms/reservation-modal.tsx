"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";

interface ReservationModalProps {
  open: boolean;
  onClose: () => void;
  chatId: string;
  onCreated: () => void;
}

export function ReservationModal({ open, onClose, chatId, onCreated }: ReservationModalProps) {
  const { addToast } = useToast();
  const [form, setForm] = useState({ scheduledDate: "", scheduledTime: "", meetingType: "in_game", meetingPointText: "", notes: "" });
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setSubmitting(true);
    try {
      await apiClient.createReservation(chatId, {
        scheduledAt: `${form.scheduledDate}T${form.scheduledTime}:00Z`,
        meetingType: form.meetingType,
        meetingPointText: form.meetingPointText || undefined,
        notes: form.notes || undefined,
      });
      onCreated();
      onClose();
    } catch {
      addToast("error", "예약 제안에 실패했습니다");
    } finally {
      setSubmitting(false);
    }
  };

  const inputClass = "w-full bg-card border border-border rounded-lg px-3 py-2.5 text-sm text-text-primary outline-none focus:border-gold placeholder:text-text-dim";

  return (
    <Modal open={open} onClose={onClose} title="예약 제안">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="grid grid-cols-2 gap-3">
          <input type="date" aria-label="거래 날짜" className={inputClass} value={form.scheduledDate} onChange={(e) => setForm({ ...form, scheduledDate: e.target.value })} required />
          <input type="time" aria-label="거래 시간" className={inputClass} value={form.scheduledTime} onChange={(e) => setForm({ ...form, scheduledTime: e.target.value })} required />
        </div>
        <select aria-label="접선 방식" className={inputClass} value={form.meetingType} onChange={(e) => setForm({ ...form, meetingType: e.target.value })}>
          <option value="in_game">인게임</option>
          <option value="offline_pc_bang">PC방/오프라인</option>
        </select>
        <input aria-label="접선 장소" className={inputClass} placeholder="접선 장소" value={form.meetingPointText} onChange={(e) => setForm({ ...form, meetingPointText: e.target.value })} />
        <textarea aria-label="메모" className={`${inputClass} h-20`} placeholder="메모 (선택)" value={form.notes} onChange={(e) => setForm({ ...form, notes: e.target.value })} />
        <button type="submit" disabled={submitting} className="w-full btn-gold-gradient text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제안 중..." : "예약 제안"}
        </button>
      </form>
    </Modal>
  );
}
