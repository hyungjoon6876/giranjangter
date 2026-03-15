import { describe, it, expect, afterEach } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { EmptyState } from "@/components/ui/empty-state";

afterEach(() => cleanup());

describe("EmptyState", () => {
  it("renders title", () => {
    render(<EmptyState title="검색 결과가 없습니다" />);
    expect(screen.getByText("검색 결과가 없습니다")).toBeDefined();
  });

  it("renders description when provided", () => {
    render(<EmptyState title="결과 없음" description="다른 검색어를 시도해보세요" />);
    expect(screen.getByText("결과 없음")).toBeDefined();
    expect(screen.getByText("다른 검색어를 시도해보세요")).toBeDefined();
  });

  it("renders title as h2", () => {
    render(<EmptyState title="검색 결과가 없습니다" />);
    const heading = screen.getByRole("heading", { level: 2 });
    expect(heading.textContent).toBe("검색 결과가 없습니다");
  });

  it("icon is aria-hidden (decorative when heading exists)", () => {
    const { container } = render(<EmptyState title="검색 결과가 없습니다" />);
    const icon = container.querySelector("[aria-hidden='true']");
    expect(icon).not.toBeNull();
    expect(icon!.textContent).toContain("\u{1F50D}");
  });

  it("action link has aria-label", () => {
    render(<EmptyState title="없음" actionLabel="매물 등록" actionHref="/listings/new" />);
    const link = screen.getByLabelText("매물 등록");
    expect(link.tagName).toBe("A");
  });
});
