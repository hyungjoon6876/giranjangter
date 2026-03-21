import { describe, it, expect, afterEach, vi, beforeEach } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createElement, type ReactNode } from "react";
import { BlockConfirmModal } from "@/components/forms/block-confirm-modal";

// Mock the hooks
vi.mock("@/lib/hooks/use-users", () => ({
  useBlockUser: () => ({
    mutateAsync: vi.fn(),
    isPending: false,
  }),
}));

vi.mock("@/lib/hooks/use-toast", () => ({
  useToast: () => ({
    addToast: vi.fn(),
  }),
}));

function createWrapper() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });

  return function Wrapper({ children }: { children: ReactNode }) {
    return createElement(QueryClientProvider, { client: queryClient }, children);
  };
}

beforeEach(() => {
  // Ensure portal container exists for modal
  let portal = document.getElementById("modal-portal");
  if (!portal) {
    portal = document.createElement("div");
    portal.id = "modal-portal";
    document.body.appendChild(portal);
  }
});

afterEach(() => {
  cleanup();
  // Clean up portal container
  const portal = document.getElementById("modal-portal");
  if (portal) {
    portal.remove();
  }
  // Reset body overflow
  document.body.style.overflow = "";
});

describe("BlockConfirmModal", () => {
  it("renders modal with user nickname", () => {
    render(
      <BlockConfirmModal
        open={true}
        onClose={vi.fn()}
        userId="user-123"
        nickname="TestUser"
      />,
      { wrapper: createWrapper() }
    );

    expect(screen.getByText("사용자 차단")).toBeDefined();
    expect(screen.getByText("TestUser")).toBeDefined();
    expect(screen.getByText(/차단하시겠습니까?/)).toBeDefined();
  });

  it("cancel button calls onClose", () => {
    const onClose = vi.fn();
    render(
      <BlockConfirmModal
        open={true}
        onClose={onClose}
        userId="user-123"
        nickname="TestUser"
      />,
      { wrapper: createWrapper() }
    );

    const cancelButton = screen.getByText("취소");
    fireEvent.click(cancelButton);

    expect(onClose).toHaveBeenCalledTimes(1);
  });

  it("does not render when closed", () => {
    render(
      <BlockConfirmModal
        open={false}
        onClose={vi.fn()}
        userId="user-123"
        nickname="TestUser"
      />,
      { wrapper: createWrapper() }
    );

    expect(screen.queryByText("사용자 차단")).toBeNull();
    expect(screen.queryByText("TestUser")).toBeNull();
  });
});