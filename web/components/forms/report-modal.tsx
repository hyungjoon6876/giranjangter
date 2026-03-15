"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";

const REPORT_REASONS = [
  { value: "scam", label: "사기 의심" },
  { value: "fake_listing", label: "허위 매물" },
  { value: "abuse", label: "욕설/비하" },
  { value: "spam", label: "도배/스팸" },
  { value: "other", label: "기타" },
];

interface ReportModalProps {
  open: boolean;
  onClose: () => void;
  targetType: string;
  targetId: string;
}

export function ReportModal({ open, onClose, targetType, targetId }: ReportModalProps) {
  const [reason, setReason] = useState("");
  const [description, setDescription] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!reason) return;
    setSubmitting(true);
    try {
      await apiClient.createReport({ targetType, targetId, reasonCode: reason, description: description || undefined });
      onClose();
      alert("신고가 접수되었습니다");
    } catch (err) {
      alert(`신고 실패: ${JSON.stringify(err)}`);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="신고하기">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          {REPORT_REASONS.map((r) => (
            <label key={r.value} className={`flex items-center gap-3 p-3 rounded-lg border cursor-pointer transition-colors ${reason === r.value ? "border-primary bg-blue-50" : "border-border hover:bg-surface"}`}>
              <input type="radio" name="reason" value={r.value} checked={reason === r.value} onChange={() => setReason(r.value)} className="accent-primary" />
              {r.label}
            </label>
          ))}
        </div>
        <textarea
          className="w-full border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-primary h-20"
          placeholder="상세 설명 (선택)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button type="submit" disabled={submitting || !reason} className="w-full bg-error text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제출 중..." : "신고하기"}
        </button>
      </form>
    </Modal>
  );
}
