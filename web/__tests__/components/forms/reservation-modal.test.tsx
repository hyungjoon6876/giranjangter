import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ReservationModal } from "@/components/forms/reservation-modal";

const mockAddToast = vi.fn();
const mockCreateReservation = vi.fn();

vi.mock("@/lib/hooks/use-toast", () => ({
  useToast: () => ({ addToast: mockAddToast }),
}));

vi.mock("@/lib/api-client", () => ({
  apiClient: {
    createReservation: (...args: unknown[]) => mockCreateReservation(...args),
  },
}));

// Portal target
beforeEach(() => {
  const portalRoot = document.createElement("div");
  portalRoot.setAttribute("id", "modal-portal");
  document.body.appendChild(portalRoot);
  vi.clearAllMocks();
});

afterEach(() => {
  cleanup();
  document.getElementById("modal-portal")?.remove();
});

describe("ReservationModal", () => {
  const defaultProps = {
    open: true,
    onClose: vi.fn(),
    chatId: "chat-1",
    onCreated: vi.fn(),
  };

  it("renders modal with title and form", () => {
    render(<ReservationModal {...defaultProps} />);
    expect(screen.getByLabelText("거래 날짜")).toBeDefined();
    expect(screen.getByLabelText("거래 시간")).toBeDefined();
    expect(screen.getByLabelText("접선 방식")).toBeDefined();
  });

  it("renders submit button", () => {
    render(<ReservationModal {...defaultProps} />);
    expect(screen.getByRole("button", { name: "예약 제안" })).toBeDefined();
  });

  it("does not render when closed", () => {
    render(<ReservationModal {...defaultProps} open={false} />);
    expect(screen.queryByLabelText("거래 날짜")).toBeNull();
  });

  it("submits form with correct data", async () => {
    mockCreateReservation.mockResolvedValue({});

    render(<ReservationModal {...defaultProps} />);

    fireEvent.change(screen.getByLabelText("거래 날짜"), { target: { value: "2024-06-15" } });
    fireEvent.change(screen.getByLabelText("거래 시간"), { target: { value: "14:00" } });
    fireEvent.change(screen.getByLabelText("접선 장소"), { target: { value: "기란마을" } });

    fireEvent.submit(screen.getByRole("button", { name: "예약 제안" }).closest("form")!);

    // Wait for async operation
    await vi.waitFor(() => {
      expect(mockCreateReservation).toHaveBeenCalledWith("chat-1", expect.objectContaining({
        scheduledAt: "2024-06-15T14:00:00Z",
        meetingType: "in_game",
        meetingPointText: "기란마을",
      }));
    });
  });

  it("shows error toast on submission failure", async () => {
    mockCreateReservation.mockRejectedValue(new Error("fail"));

    render(<ReservationModal {...defaultProps} />);

    fireEvent.change(screen.getByLabelText("거래 날짜"), { target: { value: "2024-06-15" } });
    fireEvent.change(screen.getByLabelText("거래 시간"), { target: { value: "14:00" } });

    fireEvent.submit(screen.getByRole("button", { name: "예약 제안" }).closest("form")!);

    await vi.waitFor(() => {
      expect(mockAddToast).toHaveBeenCalledWith("error", "예약 제안에 실패했습니다");
    });
  });
});
