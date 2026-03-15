"use client";

import {
  useEffect,
  useRef,
  useId,
  useCallback,
  useState,
  type ReactNode,
} from "react";
import { createPortal } from "react-dom";

interface ModalProps {
  open: boolean;
  onClose: () => void;
  title: string;
  children: ReactNode;
}

/**
 * Accessible modal dialog.
 *
 * - Portal-rendered to `#modal-portal` (appended to document.body)
 * - role="dialog", aria-modal="true", aria-labelledby linked to title
 * - Focus trap: Tab / Shift+Tab cycle within the modal
 * - Escape key closes
 * - On open: saves previous focus, applies `inert` to sibling elements
 * - On close: removes `inert`, restores focus to trigger element
 * - Locks body scroll while open
 */
export function Modal({ open, onClose, title, children }: ModalProps) {
  const titleId = useId();
  const dialogRef = useRef<HTMLDivElement>(null);
  const previousFocusRef = useRef<HTMLElement | null>(null);
  const [portalContainer, setPortalContainer] = useState<HTMLElement | null>(
    null,
  );

  // Create / find portal container on mount (client-only)
  useEffect(() => {
    if (typeof window === "undefined") return;
    let el = document.getElementById("modal-portal");
    if (!el) {
      el = document.createElement("div");
      el.id = "modal-portal";
      document.body.appendChild(el);
    }
    setPortalContainer(el);
  }, []);

  // Lock body scroll
  useEffect(() => {
    if (open) {
      document.body.style.overflow = "hidden";
    }
    return () => {
      document.body.style.overflow = "";
    };
  }, [open]);

  // Inert siblings + focus management
  useEffect(() => {
    if (!open) return;

    // Save the currently focused element so we can restore it later
    previousFocusRef.current = document.activeElement as HTMLElement | null;

    // Apply inert to sibling elements of body (excluding portal container)
    const siblings = Array.from(document.body.children).filter(
      (el) =>
        el instanceof HTMLElement &&
        el.id !== "modal-portal" &&
        !el.hasAttribute("data-modal-ignore"),
    ) as HTMLElement[];

    for (const el of siblings) {
      el.setAttribute("inert", "");
    }

    // Focus the dialog itself (or first focusable inside) after paint
    requestAnimationFrame(() => {
      const dialog = dialogRef.current;
      if (!dialog) return;
      const first = getFocusableElements(dialog)[0];
      if (first) {
        first.focus();
      } else {
        dialog.focus();
      }
    });

    return () => {
      // Remove inert from siblings
      for (const el of siblings) {
        el.removeAttribute("inert");
      }

      // Restore focus to previously focused element
      if (previousFocusRef.current && typeof previousFocusRef.current.focus === "function") {
        previousFocusRef.current.focus();
      }
    };
  }, [open]);

  // Escape key handler
  const handleKeyDown = useCallback(
    (e: React.KeyboardEvent) => {
      if (e.key === "Escape") {
        e.stopPropagation();
        onClose();
        return;
      }

      // Focus trap: Tab / Shift+Tab
      if (e.key === "Tab") {
        const dialog = dialogRef.current;
        if (!dialog) return;
        const focusable = getFocusableElements(dialog);
        if (focusable.length === 0) {
          e.preventDefault();
          return;
        }
        const first = focusable[0];
        const last = focusable[focusable.length - 1];

        if (e.shiftKey) {
          if (document.activeElement === first) {
            e.preventDefault();
            last.focus();
          }
        } else {
          if (document.activeElement === last) {
            e.preventDefault();
            first.focus();
          }
        }
      }
    },
    [onClose],
  );

  if (!open || !portalContainer) return null;

  return createPortal(
    <div
      className="fixed inset-0 z-50 flex items-center justify-center"
      onKeyDown={handleKeyDown}
    >
      {/* Backdrop */}
      <div
        className="absolute inset-0 bg-black/40"
        onClick={onClose}
        aria-hidden="true"
      />

      {/* Dialog */}
      <div
        ref={dialogRef}
        role="dialog"
        aria-modal="true"
        aria-labelledby={titleId}
        tabIndex={-1}
        className="relative bg-card rounded-xl shadow-xl w-full max-w-md mx-4 max-h-[90vh] overflow-y-auto border border-border outline-none"
      >
        <div className="flex items-center justify-between px-5 py-4 border-b border-border">
          <h2
            id={titleId}
            className="text-lg font-semibold text-text-primary"
          >
            {title}
          </h2>
          <button
            onClick={onClose}
            aria-label="닫기"
            className="text-text-secondary hover:text-text-primary"
          >
            &times;
          </button>
        </div>
        <div className="p-5">{children}</div>
      </div>
    </div>,
    portalContainer,
  );
}

/** Returns all focusable elements within a container, in DOM order. */
function getFocusableElements(container: HTMLElement): HTMLElement[] {
  const selector =
    'a[href], button:not([disabled]), textarea:not([disabled]), input:not([disabled]), select:not([disabled]), [tabindex]:not([tabindex="-1"])';
  return Array.from(container.querySelectorAll<HTMLElement>(selector));
}
