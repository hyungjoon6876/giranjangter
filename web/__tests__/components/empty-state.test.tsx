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
});
