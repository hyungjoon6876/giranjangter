import { describe, it, expect, vi, afterEach } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ToastContainer } from "@/components/ui/toast";
import { ToastContext, type ToastContextValue } from "@/lib/hooks/use-toast";

afterEach(() => cleanup());

function renderWithToast(toastValue: ToastContextValue) {
  return render(
    <ToastContext.Provider value={toastValue}>
      <ToastContainer />
    </ToastContext.Provider>,
  );
}

function makeCtx(overrides: Partial<ToastContextValue> = {}): ToastContextValue {
  return {
    toasts: [],
    addToast: () => {},
    removeToast: () => {},
    ...overrides,
  };
}

describe("ToastContainer", () => {
  it("renders nothing when there are no toasts", () => {
    const { container } = renderWithToast(makeCtx());
    expect(container.innerHTML).toBe("");
  });

  it("renders toast inside a status container with aria-live", () => {
    renderWithToast(
      makeCtx({
        toasts: [{ id: "1", type: "success", message: "저장되었습니다" }],
      }),
    );
    const container = screen.getByRole("status");
    expect(container).toBeDefined();
    expect(container.getAttribute("aria-live")).toBe("polite");
    expect(screen.getByText("저장되었습니다")).toBeDefined();
  });

  it("close button has aria-label 닫기", () => {
    renderWithToast(
      makeCtx({
        toasts: [{ id: "1", type: "error", message: "오류" }],
      }),
    );
    const btn = screen.getByLabelText("닫기");
    expect(btn).toBeDefined();
  });

  it("clicking close calls removeToast with the correct id", () => {
    const removeMock = vi.fn();
    renderWithToast(
      makeCtx({
        toasts: [{ id: "t42", type: "info", message: "알림" }],
        removeToast: removeMock,
      }),
    );
    const btn = screen.getByLabelText("닫기");
    fireEvent.click(btn);
    expect(removeMock).toHaveBeenCalledWith("t42");
  });
});
