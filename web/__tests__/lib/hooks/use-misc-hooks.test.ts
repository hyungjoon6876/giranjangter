import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { renderHook, waitFor } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createElement, type ReactNode } from "react";

// Mock the api-client module
vi.mock("@/lib/api-client", () => ({
  apiClient: {
    isLoggedIn: false,
    searchItems: vi.fn(),
    getUserReviews: vi.fn(),
    blockUser: vi.fn(),
    unblockUser: vi.fn(),
  }
}));

const mockPush = vi.fn();
const mockAddToast = vi.fn();

// Mock Next.js navigation
vi.mock("next/navigation", () => ({
  useRouter: vi.fn(() => ({ push: mockPush })),
  usePathname: vi.fn(() => "/current-path"),
}));

// Mock toast hook
vi.mock("@/lib/hooks/use-toast", () => ({
  useToast: vi.fn(() => ({ addToast: mockAddToast })),
}));

import { apiClient } from "@/lib/api-client";
import { useRouter, usePathname } from "next/navigation";
import { useItemSearch } from "@/lib/hooks/use-items";
import { useUserReviews } from "@/lib/hooks/use-reviews";
import { useBlockUser, useUnblockUser } from "@/lib/hooks/use-users";
import { useAuthGuard } from "@/lib/hooks/use-auth-guard";

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
  vi.clearAllMocks();
});

afterEach(() => {
  vi.restoreAllMocks();
});

describe("useItemSearch", () => {
  it("fetches items when query length >= 1", async () => {
    const mockItems = [{ itemId: "item-1", itemName: "집행검", iconUrl: "/icons/sword.png" }];
    vi.mocked(apiClient.searchItems).mockResolvedValue(mockItems);

    const { result } = renderHook(() => useItemSearch("집"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.searchItems).toHaveBeenCalledWith({ q: "집", categoryId: undefined });
    expect(result.current.data).toEqual(mockItems);
  });

  it("fetches items when categoryId is provided even with empty query", async () => {
    const mockItems = [{ itemId: "item-2", itemName: "활", iconUrl: "/icons/bow.png" }];
    vi.mocked(apiClient.searchItems).mockResolvedValue(mockItems);

    const { result } = renderHook(() => useItemSearch("", "weapon"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.searchItems).toHaveBeenCalledWith({ q: undefined, categoryId: "weapon" });
    expect(result.current.data).toEqual(mockItems);
  });

  it("does not fetch when query is empty and no categoryId", () => {
    const { result } = renderHook(() => useItemSearch(""), { wrapper: createWrapper() });

    expect(apiClient.searchItems).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("includes both query and categoryId in search params", async () => {
    const mockItems = [{ itemId: "item-3", itemName: "집행검+7", iconUrl: "/icons/sword.png" }];
    vi.mocked(apiClient.searchItems).mockResolvedValue(mockItems);

    const { result } = renderHook(() => useItemSearch("집행검", "weapon"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.searchItems).toHaveBeenCalledWith({ q: "집행검", categoryId: "weapon" });
  });

  it("uses correct query key with parameters", async () => {
    vi.mocked(apiClient.searchItems).mockResolvedValue([]);

    const { result } = renderHook(() => useItemSearch("test", "armor"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.searchItems).toHaveBeenCalledWith({ q: "test", categoryId: "armor" });
  });
});

describe("useUserReviews", () => {
  it("fetches user reviews when userId is provided", async () => {
    const mockReviews = [
      { reviewId: "r1", rating: 5, comment: "Great trader!", authorId: "u2", targetUserId: "u1" }
    ];
    vi.mocked(apiClient.getUserReviews).mockResolvedValue(mockReviews);

    const { result } = renderHook(() => useUserReviews("user-123"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getUserReviews).toHaveBeenCalledWith("user-123");
    expect(result.current.data).toEqual(mockReviews);
  });

  it("does not fetch when userId is empty", () => {
    const { result } = renderHook(() => useUserReviews(""), { wrapper: createWrapper() });

    expect(apiClient.getUserReviews).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("does not fetch when userId is falsy", () => {
    const { result } = renderHook(() => useUserReviews("" as any), { wrapper: createWrapper() });

    expect(apiClient.getUserReviews).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("uses correct query key with userId", async () => {
    vi.mocked(apiClient.getUserReviews).mockResolvedValue([]);

    const { result } = renderHook(() => useUserReviews("user-456"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.getUserReviews).toHaveBeenCalledWith("user-456");
  });
});

describe("useBlockUser", () => {
  it("calls apiClient.blockUser with userId", async () => {
    vi.mocked(apiClient.blockUser).mockResolvedValue();

    const { result } = renderHook(() => useBlockUser(), { wrapper: createWrapper() });

    result.current.mutate("user-to-block");

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.blockUser).toHaveBeenCalledWith("user-to-block");
  });

  it("invalidates listings and chats cache on success", async () => {
    vi.mocked(apiClient.blockUser).mockResolvedValue();

    const wrapper = createWrapper();
    const queryClient = new QueryClient({
      defaultOptions: { queries: { retry: false }, mutations: { retry: false } }
    });

    const invalidateQueries = vi.spyOn(queryClient, 'invalidateQueries');

    const WrapperWithQueryClient = ({ children }: { children: ReactNode }) =>
      createElement(QueryClientProvider, { client: queryClient }, children);

    const { result } = renderHook(() => useBlockUser(), { wrapper: WrapperWithQueryClient });

    result.current.mutate("user-123");

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(invalidateQueries).toHaveBeenCalledWith({ queryKey: ["listings"] });
    expect(invalidateQueries).toHaveBeenCalledWith({ queryKey: ["chats"] });
  });

  it("handles block user error", async () => {
    const error = { error: { code: "USER_NOT_FOUND", message: "User does not exist" } };
    vi.mocked(apiClient.blockUser).mockRejectedValue(error);

    const { result } = renderHook(() => useBlockUser(), { wrapper: createWrapper() });

    result.current.mutate("invalid-user");

    await waitFor(() => expect(result.current.isError).toBe(true));

    expect(result.current.error).toEqual(error);
  });
});

describe("useUnblockUser", () => {
  it("calls apiClient.unblockUser with userId", async () => {
    vi.mocked(apiClient.unblockUser).mockResolvedValue();

    const { result } = renderHook(() => useUnblockUser(), { wrapper: createWrapper() });

    result.current.mutate("user-to-unblock");

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.unblockUser).toHaveBeenCalledWith("user-to-unblock");
  });

  it("invalidates listings and chats cache on success", async () => {
    vi.mocked(apiClient.unblockUser).mockResolvedValue();

    const wrapper = createWrapper();
    const queryClient = new QueryClient({
      defaultOptions: { queries: { retry: false }, mutations: { retry: false } }
    });

    const invalidateQueries = vi.spyOn(queryClient, 'invalidateQueries');

    const WrapperWithQueryClient = ({ children }: { children: ReactNode }) =>
      createElement(QueryClientProvider, { client: queryClient }, children);

    const { result } = renderHook(() => useUnblockUser(), { wrapper: WrapperWithQueryClient });

    result.current.mutate("user-123");

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(invalidateQueries).toHaveBeenCalledWith({ queryKey: ["listings"] });
    expect(invalidateQueries).toHaveBeenCalledWith({ queryKey: ["chats"] });
  });

  it("handles unblock user error", async () => {
    const error = { error: { code: "NOT_BLOCKED", message: "User is not currently blocked" } };
    vi.mocked(apiClient.unblockUser).mockRejectedValue(error);

    const { result } = renderHook(() => useUnblockUser(), { wrapper: createWrapper() });

    result.current.mutate("not-blocked-user");

    await waitFor(() => expect(result.current.isError).toBe(true));

    expect(result.current.error).toEqual(error);
  });
});

describe("useAuthGuard", () => {
  it("returns true when user is logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });

    const { result } = renderHook(() => useAuthGuard());

    expect(result.current.isLoggedIn).toBe(true);
    expect(result.current.requireAuth()).toBe(true);
  });

  it("redirects to login when user is not logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useAuthGuard());

    expect(result.current.isLoggedIn).toBe(false);

    const authResult = result.current.requireAuth();

    expect(authResult).toBe(false);
    expect(mockAddToast).toHaveBeenCalledWith("info", "로그인이 필요합니다");
    expect(mockPush).toHaveBeenCalledWith("/login?redirect=%2Fcurrent-path");
  });

  it("shows custom action message when action is provided", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useAuthGuard());

    result.current.requireAuth("채팅 전송");

    expect(mockAddToast).toHaveBeenCalledWith("info", "채팅 전송은(는) 로그인이 필요합니다");
  });

  it("encodes redirect URL correctly", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    vi.mocked(usePathname).mockReturnValue("/listings?search=집행검");

    const { result } = renderHook(() => useAuthGuard());

    result.current.requireAuth();

    expect(mockPush).toHaveBeenCalledWith(
      "/login?redirect=%2Flistings%3Fsearch%3D%EC%A7%91%ED%96%89%EA%B2%80"
    );
  });
});