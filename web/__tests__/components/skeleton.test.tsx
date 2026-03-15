import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ListingSkeleton } from "@/components/ui/skeleton";

afterEach(() => cleanup());

describe("ListingSkeleton", () => {
  it("has role=status", () => {
    render(<ListingSkeleton />);
    expect(screen.getByRole("status")).toBeDefined();
  });

  it("has aria-label describing loading state", () => {
    render(<ListingSkeleton />);
    const status = screen.getByRole("status");
    expect(status.getAttribute("aria-label")).toBe("매물 목록을 불러오는 중");
  });

  it("has aria-busy=true", () => {
    render(<ListingSkeleton />);
    const status = screen.getByRole("status");
    expect(status.getAttribute("aria-busy")).toBe("true");
  });

  it("renders default 6 skeleton cards", () => {
    const { container } = render(<ListingSkeleton />);
    const grid = container.querySelector(".grid");
    expect(grid?.children.length).toBe(6);
  });

  it("renders custom count of skeleton cards", () => {
    const { container } = render(<ListingSkeleton count={3} />);
    const grid = container.querySelector(".grid");
    expect(grid?.children.length).toBe(3);
  });
});
