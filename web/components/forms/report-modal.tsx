"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";

const REPORT_REASONS = [
  { value: "scam_suspicion", label: "사기 의심" },
  { value: "fake_listing", label: "허위 매물" },
  { value: "harassment", label: "욕설/비하" },
  { value: "spam", label: "도배/스팸" },
  { value: "no_show", label: "거래 불이행" },
  { value: "other", label: "기타" },
];

interface ReportModalProps {
  open: boolean;
  onClose: () => void;
  targetType: string;
  targetId: string;
}

export function ReportModal({ open, onClose, targetType, targetId }: ReportModalProps) {
  const { addToast } = useToast();
  const [reason, setReason] = useState("");
  const [description, setDescription] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!reason) return;
    setSubmitting(true);
    try {
      await apiClient.createReport({ targetType, targetId, reportType: reason, description: description || "신고합니다" });
      onClose();
      addToast("success", "신고가 접수되었습니다");
    } catch {
      addToast("error", "신고 접수에 실패했습니다");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="신고하기">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          {REPORT_REASONS.map((r) => (
            <label key={r.value} className={`flex items-center gap-3 p-3 rounded-lg border cursor-pointer transition-colors ${reason === r.value ? "border-gold bg-[rgba(196,163,90,0.1)]" : "border-border hover:bg-medium"}`}>
              <input type="radio" name="reason" value={r.value} checked={reason === r.value} onChange={() => setReason(r.value)} className="accent-gold" />
              <span className="text-text-primary">{r.label}</span>
            </label>
          ))}
        </div>
        <textarea
          aria-label="상세 설명"
          className="w-full bg-card border border-border rounded-lg px-3 py-2.5 text-sm text-text-primary outline-none focus:border-gold h-20 placeholder:text-text-dim"
          placeholder="상세 설명 (선택)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button type="submit" disabled={submitting || !reason} className="w-full bg-danger text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제출 중..." : "신고하기"}
        </button>
      </form>
    </Modal>
  );
}
