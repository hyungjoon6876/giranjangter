import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ReviewModal } from "@/components/forms/review-modal";
import { ToastContext, type ToastContextValue } from "@/lib/hooks/use-toast";

vi.mock("@/lib/api-client", () => ({
  apiClient: { createReview: vi.fn() },
}));

afterEach(() => cleanup());

const toastCtx: ToastContextValue = { toasts: [], addToast: vi.fn(), removeToast: vi.fn() };

function renderModal(props: Parameters<typeof ReviewModal>[0]) {
  return render(
    <ToastContext.Provider value={toastCtx}>
      <ReviewModal {...props} />
    </ToastContext.Provider>,
  );
}

const defaultProps = {
  open: true,
  onClose: vi.fn(),
  completionId: "comp-1",
  onCreated: vi.fn(),
};

describe("ReviewModal", () => {
  it("renders positive/negative buttons", () => {
    renderModal(defaultProps);
    expect(screen.getByText(/좋았어요/)).toBeDefined();
    expect(screen.getByText(/아쉬웠어요/)).toBeDefined();
  });

  it("submit disabled when no rating selected", () => {
    renderModal(defaultProps);
    const submitButton = screen.getByText("리뷰 제출");
    expect(submitButton).toBeDisabled();
  });

  it("shows selected rating state", () => {
    renderModal(defaultProps);
    const positiveBtn = screen.getByText(/좋았어요/);
    fireEvent.click(positiveBtn);

    // After clicking, the submit button should now be enabled
    const submitButton = screen.getByText("리뷰 제출");
    expect(submitButton).not.toBeDisabled();
  });
});
