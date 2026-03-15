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

  it("aria-label is sufficient without redundant sr-only text", () => {
    render(<Loading />);
    const status = screen.getByRole("status");
    expect(status.getAttribute("aria-label")).toBe("로딩 중");
    // sr-only span removed — aria-label already provides the accessible name
    expect(status.querySelector(".sr-only")).toBeNull();
  });
});
