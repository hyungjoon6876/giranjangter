import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { Badge, TypeBadge } from "@/components/ui/badge";

afterEach(() => cleanup());

describe("Badge", () => {
  it("renders label text", () => {
    render(<Badge label="테스트" color="#FF0000" />);
    expect(screen.getByText("테스트")).toBeDefined();
  });
});

describe("TypeBadge", () => {
  it("sell shows 판매", () => {
    render(<TypeBadge type="sell" />);
    expect(screen.getByText("판매")).toBeDefined();
  });

  it("buy shows 구매", () => {
    render(<TypeBadge type="buy" />);
    expect(screen.getByText("구매")).toBeDefined();
  });
});
