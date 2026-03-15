import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { ErrorState } from "@/components/ui/error-state";

afterEach(() => cleanup());

describe("ErrorState", () => {
  it("has role=alert", () => {
    render(<ErrorState message="오류가 발생했습니다" />);
    expect(screen.getByRole("alert")).toBeDefined();
  });

  it("shows custom message", () => {
    render(<ErrorState message="서버 오류" />);
    expect(screen.getByText("서버 오류")).toBeDefined();
  });

  it("shows description when provided", () => {
    render(<ErrorState message="오류" description="잠시 후 다시 시도하세요" />);
    expect(screen.getByText("잠시 후 다시 시도하세요")).toBeDefined();
  });

  it("warning icon has role=img with aria-label", () => {
    render(<ErrorState message="오류" />);
    const icon = screen.getByRole("img", { name: "경고" });
    expect(icon).toBeDefined();
  });

  it("retry button calls onRetry", () => {
    const onRetry = vi.fn();
    render(<ErrorState message="오류" onRetry={onRetry} />);
    const btn = screen.getByText("다시 시도");
    fireEvent.click(btn);
    expect(onRetry).toHaveBeenCalledOnce();
  });

  it("auto-focuses retry button when autoFocus is true", () => {
    const onRetry = vi.fn();
    render(<ErrorState message="오류" onRetry={onRetry} autoFocus />);
    const btn = screen.getByText("다시 시도");
    expect(document.activeElement).toBe(btn);
  });

  it("does not auto-focus retry button by default", () => {
    const onRetry = vi.fn();
    render(<ErrorState message="오류" onRetry={onRetry} />);
    const btn = screen.getByText("다시 시도");
    expect(document.activeElement).not.toBe(btn);
  });
});
