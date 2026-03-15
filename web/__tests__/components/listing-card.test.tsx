import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { ListingCard } from "@/components/listing/listing-card";

const mockListing = {
  listingId: "1",
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
  // required fields
  quantity: 1, visibility: "public", tradeMethod: "in_game",
  categoryId: "weapon", lastActivityAt: new Date().toISOString(),
};

afterEach(() => cleanup());

describe("ListingCard", () => {
  it("renders title and price", () => {
    render(<ListingCard listing={mockListing} />);
    expect(screen.getByText("집행검 +7 팝니다")).toBeDefined();
    expect(screen.getByText(/500,000/)).toBeDefined();
  });

  it("shows sell badge", () => {
    render(<ListingCard listing={mockListing} />);
    expect(screen.getByText("판매")).toBeDefined();
  });
});
