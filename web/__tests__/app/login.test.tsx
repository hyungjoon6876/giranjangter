import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";

// Mock next/navigation
vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: vi.fn() }),
}));

// Mock api-client
vi.mock("@/lib/api-client", () => ({
  apiClient: { login: vi.fn() },
}));

import LoginPage from "@/app/login/page";

afterEach(() => cleanup());

describe("LoginPage", () => {
  it("renders 기란장터 title", () => {
    render(<LoginPage />);
    expect(screen.getByText("기란장터")).toBeDefined();
  });

  it("renders 둘러보기 link", () => {
    render(<LoginPage />);
    expect(screen.getByText("둘러보기")).toBeDefined();
  });

  it("dev login button is present", () => {
    render(<LoginPage />);
    expect(screen.getByText(/개발자 로그인/)).toBeDefined();
  });
});
