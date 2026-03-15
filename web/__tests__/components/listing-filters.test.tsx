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

  it("search input renders", () => {
    render(<ListingFilters {...defaultProps} />);
    expect(screen.getByPlaceholderText(/아이템 검색/)).toBeDefined();
  });
});
