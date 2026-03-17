import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";

// Mock next/navigation
vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: vi.fn() }),
  useSearchParams: () => new URLSearchParams(),
}));

// Mock api-client
vi.mock("@/lib/api-client", () => ({
  apiClient: { login: vi.fn(), isLoggedIn: false },
}));

// Mock @tanstack/react-query
vi.mock("@tanstack/react-query", () => ({
  useQueryClient: () => ({ invalidateQueries: vi.fn(), clear: vi.fn() }),
}));

import LoginPage from "@/app/login/page";

afterEach(() => cleanup());

describe("LoginPage", () => {
  it("renders 기란JT logo", () => {
    render(<LoginPage />);
    expect(screen.getByAltText("기란JT")).toBeDefined();
  });

  it("renders 둘러보기 link", () => {
    render(<LoginPage />);
    expect(screen.getByText("로그인 없이 둘러보기")).toBeDefined();
  });

  it("renders Google sign-in container", () => {
    render(<LoginPage />);
    // The Google button container is rendered as an empty div with a ref
    const container = document.querySelector(".flex.justify-center.mb-3");
    expect(container).toBeDefined();
  });
});
