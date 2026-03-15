"use client";

import { useEffect, useRef } from "react";

interface ErrorStateProps {
  message: string;
  description?: string;
  onRetry?: () => void;
  autoFocus?: boolean;
}

export function ErrorState({ message, description, onRetry, autoFocus }: ErrorStateProps) {
  const retryRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    if (autoFocus) retryRef.current?.focus();
  }, [autoFocus]);

  return (
    <div role="alert" className="flex flex-col items-center justify-center py-20 text-text-secondary">
      <span role="img" aria-label="경고" className="text-5xl mb-4">
        ⚠️
      </span>
      <p className="text-lg text-text-primary">{message}</p>
      {description && <p className="text-sm mt-2 text-text-dim">{description}</p>}
      {onRetry && (
        <button
          ref={retryRef}
          onClick={onRetry}
          className="mt-6 btn-gold-gradient text-white px-6 py-2 rounded-lg text-sm font-medium"
        >
          다시 시도
        </button>
      )}
    </div>
  );
}
