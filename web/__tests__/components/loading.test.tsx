import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { Loading } from "@/components/ui/loading";

afterEach(() => cleanup());

describe("Loading", () => {
  it("has role=status", () => {
    render(<Loading />);
    expect(screen.getByRole("status")).toBeDefined();
  });

  it("has accessible label", () => {
    render(<Loading />);
    const status = screen.getByRole("status");
    expect(status.getAttribute("aria-label")).toBe("로딩 중");
  });

  it("has sr-only text for screen readers", () => {
    render(<Loading />);
    expect(screen.getByText("로딩 중")).toBeDefined();
  });
});
