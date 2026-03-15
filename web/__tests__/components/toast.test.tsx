import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
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

  it("renders toast with role=alert", () => {
    renderWithToast(
      makeCtx({
        toasts: [{ id: "1", type: "success", message: "저장되었습니다" }],
      }),
    );
    const alert = screen.getByRole("alert");
    expect(alert).toBeDefined();
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
});
