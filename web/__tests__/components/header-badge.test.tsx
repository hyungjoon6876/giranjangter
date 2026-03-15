import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup } from "@testing-library/react";
import { createElement, type ReactNode } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

// Mock next/navigation
vi.mock("next/navigation", () => ({
  usePathname: () => "/",
}));

// Mock api-client
vi.mock("@/lib/api-client", () => ({
  apiClient: { isLoggedIn: true },
  API_BASE: "http://localhost:8080",
}));

// Control notifications data via a ref
let mockNotifications: { data: { notificationId: string; message: string; readAt?: string; createdAt: string }[] } | undefined;

vi.mock("@/lib/hooks/use-profile", () => ({
  useNotifications: () => ({ data: mockNotifications }),
}));

import { Header } from "@/components/layout/header";

afterEach(() => {
  cleanup();
  mockNotifications = undefined;
});

function renderWithQuery(ui: ReactNode) {
  const qc = new QueryClient({ defaultOptions: { queries: { retry: false } } });
  return render(createElement(QueryClientProvider, { client: qc }, ui));
}

describe("Header notification badge", () => {
  it("shows red dot when there are unread notifications", () => {
    mockNotifications = {
      data: [
        { notificationId: "n1", message: "새 메시지", createdAt: new Date().toISOString() },
        { notificationId: "n2", message: "예약 확정", readAt: new Date().toISOString(), createdAt: new Date().toISOString() },
      ],
    };
    const { container } = renderWithQuery(createElement(Header));
    const link = screen.getByLabelText("읽지 않은 알림 있음");
    expect(link).toBeDefined();
    // Red dot badge
    const badge = container.querySelector(".bg-red-500.rounded-full");
    expect(badge).not.toBeNull();
  });

  it("hides badge when all notifications are read", () => {
    mockNotifications = {
      data: [
        { notificationId: "n1", message: "새 메시지", readAt: new Date().toISOString(), createdAt: new Date().toISOString() },
      ],
    };
    const { container } = renderWithQuery(createElement(Header));
    const link = screen.getByLabelText("알림");
    expect(link).toBeDefined();
    const badge = container.querySelector(".bg-red-500.rounded-full");
    expect(badge).toBeNull();
  });

  it("hides badge when no notifications exist", () => {
    mockNotifications = undefined;
    const { container } = renderWithQuery(createElement(Header));
    const link = screen.getByLabelText("알림");
    expect(link).toBeDefined();
    const badge = container.querySelector(".bg-red-500.rounded-full");
    expect(badge).toBeNull();
  });
});
