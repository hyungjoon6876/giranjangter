import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ReviewModal } from "@/components/forms/review-modal";

vi.mock("@/lib/api-client", () => ({
  apiClient: { createReview: vi.fn() },
}));

afterEach(() => cleanup());

const defaultProps = {
  open: true,
  onClose: vi.fn(),
  completionId: "comp-1",
  onCreated: vi.fn(),
};

describe("ReviewModal", () => {
  it("renders positive/negative buttons", () => {
    render(<ReviewModal {...defaultProps} />);
    expect(screen.getByText(/좋았어요/)).toBeDefined();
    expect(screen.getByText(/아쉬웠어요/)).toBeDefined();
  });

  it("submit disabled when no rating selected", () => {
    render(<ReviewModal {...defaultProps} />);
    const submitButton = screen.getByText("리뷰 제출");
    expect(submitButton).toBeDisabled();
  });

  it("shows selected rating state", () => {
    render(<ReviewModal {...defaultProps} />);
    const positiveBtn = screen.getByText(/좋았어요/);
    fireEvent.click(positiveBtn);

    // After clicking, the submit button should now be enabled
    const submitButton = screen.getByText("리뷰 제출");
    expect(submitButton).not.toBeDisabled();
  });
});
