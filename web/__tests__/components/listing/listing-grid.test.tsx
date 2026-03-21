import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ListingGrid } from "@/components/listing/listing-grid";

const mockListing = {
  listingId: "test-1",
  listingType: "sell" as const,
  title: "집행검 +7 팝니다",
  itemName: "집행검",
  priceType: "negotiable" as const,
  priceAmount: 500000,
  enhancementLevel: 7,
  serverId: "bartz",
  serverName: "바츠",
  status: "available",
  author: { userId: "u1", nickname: "검은기사" },
  viewCount: 12,
  favoriteCount: 3,
  chatCount: 1,
  createdAt: new Date().toISOString(),
  quantity: 1,
  visibility: "public",
  tradeMethod: "in_game",
  categoryId: "weapon",
  lastActivityAt: new Date().toISOString(),
};

afterEach(() => cleanup());

describe("ListingGrid", () => {
  it("renders correct number of listing cards", () => {
    const listings = [
      { ...mockListing, listingId: "test-1", title: "Item 1" },
      { ...mockListing, listingId: "test-2", title: "Item 2" },
      { ...mockListing, listingId: "test-3", title: "Item 3" },
    ];

    render(<ListingGrid listings={listings} />);

    expect(screen.getByText("Item 1")).toBeDefined();
    expect(screen.getByText("Item 2")).toBeDefined();
    expect(screen.getByText("Item 3")).toBeDefined();
  });

  it("renders nothing when listings array is empty", () => {
    const { container } = render(<ListingGrid listings={[]} />);

    // Should render the grid container but with no cards
    expect(container.firstChild).toBeDefined();
    expect(container.firstChild?.childNodes).toHaveLength(0);
  });
});