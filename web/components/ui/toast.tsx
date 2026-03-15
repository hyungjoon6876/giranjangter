"use client";

import type { Toast, ToastType } from "@/lib/hooks/use-toast";
import { useToast } from "@/lib/hooks/use-toast";

const borderColors: Record<ToastType, string> = {
  success: "#2ecc71",
  error: "#e74c3c",
  info: "#c4a35a",
};

function ToastCard({ toast }: { toast: Toast }) {
  const { removeToast } = useToast();

  return (
    <div
      role="alert"
      className="flex items-start gap-3 bg-card border border-border rounded-lg p-4 shadow-lg"
      style={{ borderLeftWidth: 4, borderLeftColor: borderColors[toast.type] }}
    >
      <p className="flex-1 text-sm text-text-primary">{toast.message}</p>
      <button
        onClick={() => removeToast(toast.id)}
        aria-label="닫기"
        className="text-text-dim hover:text-text-primary transition-colors text-lg leading-none"
      >
        &times;
      </button>
    </div>
  );
}

export function ToastContainer() {
  const { toasts } = useToast();

  if (toasts.length === 0) return null;

  return (
    <div className="fixed top-4 right-4 z-[60] flex flex-col gap-2 max-sm:left-4 max-sm:right-4">
      {toasts.map((toast) => (
        <ToastCard key={toast.id} toast={toast} />
      ))}
    </div>
  );
}
