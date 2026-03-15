import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ListingFilters } from "@/components/listing/listing-filters";
import type { Server } from "@/lib/types";

afterEach(() => cleanup());

const servers: Server[] = [
  { serverId: "bartz", serverName: "바츠" },
  { serverId: "ken_rauhel", serverName: "켄라우헬" },
];

const defaultProps = {
  servers,
  selectedServer: null as string | null,
  onServerChange: vi.fn(),
  searchQuery: "",
  onSearchChange: vi.fn(),
};

describe("ListingFilters", () => {
  it("renders 전체 button", () => {
    render(<ListingFilters {...defaultProps} />);
    expect(screen.getByText("전체")).toBeDefined();
  });

  it("renders server buttons", () => {
    render(<ListingFilters {...defaultProps} />);
    expect(screen.getByText("바츠")).toBeDefined();
    expect(screen.getByText("켄라우헬")).toBeDefined();
  });

  it("active server has gold gradient", () => {
    render(<ListingFilters {...defaultProps} selectedServer="bartz" />);
    const bartzBtn = screen.getByText("바츠");
    expect(bartzBtn.className).toContain("btn-gold-gradient");

    // 전체 should NOT have gold gradient when a server is selected
    const allBtn = screen.getByText("전체");
    expect(allBtn.className).not.toContain("btn-gold-gradient");
  });

  it("search input renders with aria-label", () => {
    render(<ListingFilters {...defaultProps} />);
    const input = screen.getByLabelText("매물 검색");
    expect(input).toBeDefined();
    expect(input.getAttribute("type")).toBe("search");
  });

  it("server filter group has role and aria-label", () => {
    render(<ListingFilters {...defaultProps} />);
    const group = screen.getByRole("group", { name: "서버 필터" });
    expect(group).toBeDefined();
  });

  it("server buttons have aria-pressed", () => {
    render(<ListingFilters {...defaultProps} selectedServer="bartz" />);
    const bartzBtn = screen.getByText("바츠");
    expect(bartzBtn.getAttribute("aria-pressed")).toBe("true");

    const allBtn = screen.getByText("전체");
    expect(allBtn.getAttribute("aria-pressed")).toBe("false");

    const kenBtn = screen.getByText("켄라우헬");
    expect(kenBtn.getAttribute("aria-pressed")).toBe("false");
  });

  it("search input is wrapped in search element", () => {
    const { container } = render(<ListingFilters {...defaultProps} />);
    const searchEl = container.querySelector("search");
    expect(searchEl).not.toBeNull();
    expect(searchEl?.querySelector("input[type='search']")).not.toBeNull();
  });
});
