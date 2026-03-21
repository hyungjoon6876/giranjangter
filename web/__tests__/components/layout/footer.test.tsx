import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { Footer } from "@/components/layout/footer";

afterEach(() => cleanup());

describe("Footer", () => {
  it("renders brand name", () => {
    render(<Footer />);
    expect(screen.getByText("기란JT")).toBeDefined();
  });

  it("shows copyright year", () => {
    render(<Footer />);
    const currentYear = new Date().getFullYear();
    expect(screen.getByText(new RegExp(`© ${currentYear}`))).toBeDefined();
  });
});