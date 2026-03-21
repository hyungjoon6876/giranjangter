import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ReportModal } from "@/components/forms/report-modal";

const mockAddToast = vi.fn();
const mockCreateReport = vi.fn();

vi.mock("@/lib/hooks/use-toast", () => ({
  useToast: () => ({ addToast: mockAddToast }),
}));

vi.mock("@/lib/api-client", () => ({
  apiClient: {
    createReport: (...args: unknown[]) => mockCreateReport(...args),
  },
}));

beforeEach(() => {
  const portalRoot = document.createElement("div");
  portalRoot.setAttribute("id", "portal-root");
  document.body.appendChild(portalRoot);
  vi.clearAllMocks();
});

afterEach(() => {
  cleanup();
  document.getElementById("portal-root")?.remove();
});

describe("ReportModal", () => {
  const defaultProps = {
    open: true,
    onClose: vi.fn(),
    targetType: "listing",
    targetId: "listing-1",
  };

  it("renders report reasons", () => {
    render(<ReportModal {...defaultProps} />);
    expect(screen.getByRole("button", { name: /신고하기/ })).toBeDefined();
    expect(screen.getByText("사기 의심")).toBeDefined();
    expect(screen.getByText("허위 매물")).toBeDefined();
    expect(screen.getByText("욕설/비하")).toBeDefined();
    expect(screen.getByText("도배/스팸")).toBeDefined();
  });

  it("submit button is disabled without reason", () => {
    render(<ReportModal {...defaultProps} />);
    const button = screen.getByRole("button", { name: /신고하기/ });
    expect(button).toBeDisabled();
  });

  it("enables submit button after selecting reason", () => {
    render(<ReportModal {...defaultProps} />);
    fireEvent.click(screen.getByText("사기 의심"));
    const button = screen.getByRole("button", { name: /신고하기/ });
    expect(button).not.toBeDisabled();
  });

  it("submits with selected reason", async () => {
    mockCreateReport.mockResolvedValue({});

    render(<ReportModal {...defaultProps} />);
    fireEvent.click(screen.getByText("허위 매물"));
    fireEvent.submit(screen.getByRole("button", { name: /신고하기/ }).closest("form")!);

    await vi.waitFor(() => {
      expect(mockCreateReport).toHaveBeenCalledWith(expect.objectContaining({
        targetType: "listing",
        targetId: "listing-1",
        reportType: "fake_listing",
      }));
    });
  });

  it("shows success toast on submission", async () => {
    mockCreateReport.mockResolvedValue({});

    render(<ReportModal {...defaultProps} />);
    fireEvent.click(screen.getByText("사기 의심"));
    fireEvent.submit(screen.getByRole("button", { name: /신고하기/ }).closest("form")!);

    await vi.waitFor(() => {
      expect(mockAddToast).toHaveBeenCalledWith("success", "신고가 접수되었습니다");
    });
  });

  it("shows error toast on failure", async () => {
    mockCreateReport.mockRejectedValue(new Error("fail"));

    render(<ReportModal {...defaultProps} />);
    fireEvent.click(screen.getByText("사기 의심"));
    fireEvent.submit(screen.getByRole("button", { name: /신고하기/ }).closest("form")!);

    await vi.waitFor(() => {
      expect(mockAddToast).toHaveBeenCalledWith("error", "신고 접수에 실패했습니다");
    });
  });
});
