import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ReportModal } from "@/components/forms/report-modal";
import { ToastContext, type ToastContextValue } from "@/lib/hooks/use-toast";

vi.mock("@/lib/api-client", () => ({
  apiClient: { createReport: vi.fn() },
}));

afterEach(() => cleanup());

const toastCtx: ToastContextValue = { toasts: [], addToast: vi.fn(), removeToast: vi.fn() };

function renderModal(props: Parameters<typeof ReportModal>[0]) {
  return render(
    <ToastContext.Provider value={toastCtx}>
      <ReportModal {...props} />
    </ToastContext.Provider>,
  );
}

const defaultProps = {
  open: true,
  onClose: vi.fn(),
  targetType: "listing",
  targetId: "listing-1",
};

describe("ReportModal", () => {
  it("renders all 5 report reasons", () => {
    renderModal(defaultProps);
    expect(screen.getByText("사기 의심")).toBeDefined();
    expect(screen.getByText("허위 매물")).toBeDefined();
    expect(screen.getByText("욕설/비하")).toBeDefined();
    expect(screen.getByText("도배/스팸")).toBeDefined();
    expect(screen.getByText("기타")).toBeDefined();
  });

  it("submit disabled when no reason selected", () => {
    renderModal(defaultProps);
    const submitButton = screen.getByText("신고하기", { selector: "button[type='submit']" });
    expect(submitButton).toBeDisabled();
  });
});
