"use client";

import { useState } from "react";
import { Modal } from "@/components/ui/modal";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";

interface ReviewModalProps {
  open: boolean;
  onClose: () => void;
  completionId: string;
  onCreated: () => void;
}

export function ReviewModal({ open, onClose, completionId, onCreated }: ReviewModalProps) {
  const { addToast } = useToast();
  const [rating, setRating] = useState<"positive" | "negative" | null>(null);
  const [comment, setComment] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!rating) return;
    setSubmitting(true);
    try {
      await apiClient.createReview(completionId, { rating, comment: comment || undefined });
      onCreated();
      onClose();
    } catch (err) {
      addToast("error", "리뷰 제출에 실패했습니다");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="거래 리뷰">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="flex gap-3">
          {(["positive", "negative"] as const).map((r) => (
            <button
              key={r}
              type="button"
              onClick={() => setRating(r)}
              className={`flex-1 py-3 rounded-lg text-sm font-medium border transition-colors ${
                rating === r
                  ? r === "positive" ? "bg-success/10 border-success text-success" : "bg-danger/10 border-danger text-danger"
                  : "border-border text-text-secondary hover:bg-medium"
              }`}
            >
              {r === "positive" ? "좋았어요" : "아쉬웠어요"}
            </button>
          ))}
        </div>
        <textarea
          className="w-full bg-card border border-border rounded-lg px-3 py-2.5 text-sm text-text-primary outline-none focus:border-gold h-24 placeholder:text-text-dim"
          placeholder="한줄 코멘트 (선택)"
          value={comment}
          onChange={(e) => setComment(e.target.value)}
        />
        <button type="submit" disabled={submitting || !rating} className="w-full btn-gold-gradient text-white py-3 rounded-lg font-semibold disabled:opacity-50">
          {submitting ? "제출 중..." : "리뷰 제출"}
        </button>
      </form>
    </Modal>
  );
}
