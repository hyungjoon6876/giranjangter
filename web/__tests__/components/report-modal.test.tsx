import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ReportModal } from "@/components/forms/report-modal";

vi.mock("@/lib/api-client", () => ({
  apiClient: { createReport: vi.fn() },
}));

afterEach(() => cleanup());

const defaultProps = {
  open: true,
  onClose: vi.fn(),
  targetType: "listing",
  targetId: "listing-1",
};

describe("ReportModal", () => {
  it("renders all 5 report reasons", () => {
    render(<ReportModal {...defaultProps} />);
    expect(screen.getByText("사기 의심")).toBeDefined();
    expect(screen.getByText("허위 매물")).toBeDefined();
    expect(screen.getByText("욕설/비하")).toBeDefined();
    expect(screen.getByText("도배/스팸")).toBeDefined();
    expect(screen.getByText("기타")).toBeDefined();
  });

  it("submit disabled when no reason selected", () => {
    render(<ReportModal {...defaultProps} />);
    const submitButton = screen.getByText("신고하기", { selector: "button[type='submit']" });
    expect(submitButton).toBeDisabled();
  });
});
