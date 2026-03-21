import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { InfoRow, AuthorSection, tradeMethodLabel } from "@/components/listing/listing-info";

afterEach(() => cleanup());

describe("InfoRow", () => {
  it("renders label and value", () => {
    render(<InfoRow label="가격" value="500,000원" />);

    expect(screen.getByText("가격")).toBeDefined();
    expect(screen.getByText("500,000원")).toBeDefined();
  });

  it("displays label in first column and value in second", () => {
    const { container } = render(<InfoRow label="서버" value="바츠" />);

    const dt = container.querySelector("dt");
    const dd = container.querySelector("dd");

    expect(dt?.textContent).toBe("서버");
    expect(dd?.textContent).toBe("바츠");
  });
});

describe("AuthorSection", () => {
  const mockAuthor = {
    userId: "user-123",
    nickname: "검은기사",
    completedTradeCount: 5,
    trustBadge: "신뢰회원",
  };

  it("renders author nickname", () => {
    render(<AuthorSection author={mockAuthor} />);

    expect(screen.getByText("검은기사")).toBeDefined();
  });

  it("displays completed trade count", () => {
    render(<AuthorSection author={mockAuthor} />);

    expect(screen.getByText(/거래.*5.*회/)).toBeDefined();
  });

  it("shows trust badge when available", () => {
    render(<AuthorSection author={mockAuthor} />);

    expect(screen.getByText(/신뢰회원/)).toBeDefined();
  });

  it("displays first letter of nickname as avatar", () => {
    render(<AuthorSection author={mockAuthor} />);

    expect(screen.getByText("검")).toBeDefined();
  });

  it("links to user reviews page", () => {
    render(<AuthorSection author={mockAuthor} />);

    const link = screen.getByRole("link");
    expect(link.getAttribute("href")).toBe("/profile/user-123/reviews");
  });

  it("handles missing author data gracefully", () => {
    const incompleteAuthor = {
      userId: "user-456",
      nickname: "",
    };

    render(<AuthorSection author={incompleteAuthor} />);

    expect(screen.getByText("?")).toBeDefined(); // fallback avatar
    expect(screen.getByText(/거래.*0.*회/)).toBeDefined(); // default trade count
  });
});

describe("tradeMethodLabel", () => {
  it("maps correct trade methods to Korean labels", () => {
    expect(tradeMethodLabel("in_game")).toBe("인게임");
    expect(tradeMethodLabel("offline_pc_bang")).toBe("PC방/오프라인");
    expect(tradeMethodLabel("either")).toBe("무관");
  });

  it("returns original method for unknown values", () => {
    expect(tradeMethodLabel("unknown_method")).toBe("unknown_method");
  });

  it("returns empty string for null/undefined", () => {
    expect(tradeMethodLabel(null as any)).toBe("");
    expect(tradeMethodLabel(undefined)).toBe("");
  });

  it("returns empty string for empty string", () => {
    expect(tradeMethodLabel("")).toBe("");
  });
});