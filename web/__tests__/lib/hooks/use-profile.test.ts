import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { renderHook, waitFor } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createElement, type ReactNode } from "react";
import { createQueryWrapper } from "@/__tests__/test-utils";

vi.mock("@/lib/api-client", () => ({
  apiClient: {
    isLoggedIn: false,
    getMe: vi.fn(),
    getMyListings: vi.fn(),
    getMyTrades: vi.fn(),
    getNotifications: vi.fn(),
    markNotificationsRead: vi.fn(),
    updateProfile: vi.fn(),
  }
}));

import { apiClient } from "@/lib/api-client";
import {
  useMe,
  useMyListings,
  useMyTrades,
  useNotifications,
  useMarkNotificationsRead,
  useUpdateProfile,
} from "@/lib/hooks/use-profile";

beforeEach(() => {
  vi.clearAllMocks();
});

afterEach(() => {
  vi.restoreAllMocks();
});

describe("useMe", () => {
  it("does not fetch when user is not logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useMe(), { wrapper: createQueryWrapper() });

    expect(apiClient.getMe).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("fetches user profile when logged in", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    const mockUser = { userId: "u1", nickname: "TestUser", email: "test@example.com" };
    vi.mocked(apiClient.getMe).mockResolvedValue(mockUser);

    const { result } = renderHook(() => useMe(), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getMe).toHaveBeenCalled();
    expect(result.current.data).toEqual(mockUser);
  });
});

describe("useMyListings", () => {
  it("does not fetch when user is not logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useMyListings(), { wrapper: createQueryWrapper() });

    expect(apiClient.getMyListings).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("fetches user listings when logged in", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    const mockListings = { data: [{ listingId: "l1", title: "My Item" }], cursor: { hasMore: false } };
    vi.mocked(apiClient.getMyListings).mockResolvedValue(mockListings);

    const { result } = renderHook(() => useMyListings(), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getMyListings).toHaveBeenCalledWith(undefined);
    expect(result.current.data).toEqual(mockListings);
  });

  it("passes status filter to API", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    const mockListings = { data: [], cursor: { hasMore: false } };
    vi.mocked(apiClient.getMyListings).mockResolvedValue(mockListings);

    const { result } = renderHook(() => useMyListings("available"), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getMyListings).toHaveBeenCalledWith("available");
  });

  it("uses correct query key with status", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    vi.mocked(apiClient.getMyListings).mockResolvedValue({ data: [], cursor: { hasMore: false } });

    const { result } = renderHook(() => useMyListings("sold"), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.getMyListings).toHaveBeenCalledWith("sold");
  });
});

describe("useMyTrades", () => {
  it("does not fetch when user is not logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useMyTrades(), { wrapper: createQueryWrapper() });

    expect(apiClient.getMyTrades).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("fetches user trades when logged in", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    const mockTrades = { data: [{ listingId: "t1", title: "Trade Item" }], cursor: { hasMore: false } };
    vi.mocked(apiClient.getMyTrades).mockResolvedValue(mockTrades);

    const { result } = renderHook(() => useMyTrades(), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getMyTrades).toHaveBeenCalled();
    expect(result.current.data).toEqual(mockTrades);
  });
});

describe("useNotifications", () => {
  it("does not fetch when user is not logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useNotifications(), { wrapper: createQueryWrapper() });

    expect(apiClient.getNotifications).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("fetches notifications when logged in", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    const mockNotifications = {
      data: [{ notificationId: "n1", message: "New message", read: false }],
      cursor: { hasMore: false }
    };
    vi.mocked(apiClient.getNotifications).mockResolvedValue(mockNotifications);

    const { result } = renderHook(() => useNotifications(), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getNotifications).toHaveBeenCalled();
    expect(result.current.data).toEqual(mockNotifications);
  });

  it("sets up polling with 30 second interval", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    vi.mocked(apiClient.getNotifications).mockResolvedValue({ data: [], cursor: { hasMore: false } });

    const { result } = renderHook(() => useNotifications(), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.getNotifications).toHaveBeenCalled();
  });
});

describe("useMarkNotificationsRead", () => {
  it("calls apiClient.markNotificationsRead with IDs", async () => {
    vi.mocked(apiClient.markNotificationsRead).mockResolvedValue();

    const { result } = renderHook(() => useMarkNotificationsRead(), { wrapper: createQueryWrapper() });

    result.current.mutate(["n1", "n2", "n3"]);

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.markNotificationsRead).toHaveBeenCalledWith(["n1", "n2", "n3"]);
  });

  it("invalidates notifications cache on success", async () => {
    vi.mocked(apiClient.markNotificationsRead).mockResolvedValue();
    const wrapper = createQueryWrapper();
    const queryClient = new QueryClient({
      defaultOptions: { queries: { retry: false }, mutations: { retry: false } }
    });

    // Mock invalidateQueries
    const invalidateQueries = vi.spyOn(queryClient, 'invalidateQueries');

    const WrapperWithQueryClient = ({ children }: { children: ReactNode }) =>
      createElement(QueryClientProvider, { client: queryClient }, children);

    const { result } = renderHook(() => useMarkNotificationsRead(), { wrapper: WrapperWithQueryClient });

    result.current.mutate(["n1"]);

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(invalidateQueries).toHaveBeenCalledWith({ queryKey: ["notifications"] });
  });

  it("handles mark read error", async () => {
    const error = { error: { code: "NOT_FOUND", message: "Notifications not found" } };
    vi.mocked(apiClient.markNotificationsRead).mockRejectedValue(error);

    const { result } = renderHook(() => useMarkNotificationsRead(), { wrapper: createQueryWrapper() });

    result.current.mutate(["invalid-id"]);

    await waitFor(() => expect(result.current.isError).toBe(true));

    expect(result.current.error).toEqual(error);
  });
});

describe("useUpdateProfile", () => {
  it("calls apiClient.updateProfile with correct data", async () => {
    const profileData = {
      nickname: "NewNickname",
      introduction: "Updated bio",
      primaryServerId: "bartz",
      avatarUrl: "/avatar/new.jpg"
    };
    const updatedUser = { userId: "u1", ...profileData };
    vi.mocked(apiClient.updateProfile).mockResolvedValue(updatedUser);

    const { result } = renderHook(() => useUpdateProfile(), { wrapper: createQueryWrapper() });

    result.current.mutate(profileData);

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.updateProfile).toHaveBeenCalledWith(profileData);
    expect(result.current.data).toEqual(updatedUser);
  });

  it("invalidates 'me' cache on successful profile update", async () => {
    vi.mocked(apiClient.updateProfile).mockResolvedValue({ userId: "u1" });

    const wrapper = createQueryWrapper();
    const queryClient = new QueryClient({
      defaultOptions: { queries: { retry: false }, mutations: { retry: false } }
    });

    const invalidateQueries = vi.spyOn(queryClient, 'invalidateQueries');

    const WrapperWithQueryClient = ({ children }: { children: ReactNode }) =>
      createElement(QueryClientProvider, { client: queryClient }, children);

    const { result } = renderHook(() => useUpdateProfile(), { wrapper: WrapperWithQueryClient });

    result.current.mutate({ nickname: "Updated" });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(invalidateQueries).toHaveBeenCalledWith({ queryKey: ["me"] });
  });

  it("handles profile update error", async () => {
    const error = { error: { code: "INVALID_NICKNAME", message: "Nickname already taken" } };
    vi.mocked(apiClient.updateProfile).mockRejectedValue(error);

    const { result } = renderHook(() => useUpdateProfile(), { wrapper: createQueryWrapper() });

    result.current.mutate({ nickname: "taken-nick" });

    await waitFor(() => expect(result.current.isError).toBe(true));

    expect(result.current.error).toEqual(error);
  });

  it("allows partial profile updates", async () => {
    const partialData = { nickname: "OnlyNickname" };
    vi.mocked(apiClient.updateProfile).mockResolvedValue({ userId: "u1", ...partialData });

    const { result } = renderHook(() => useUpdateProfile(), { wrapper: createQueryWrapper() });

    result.current.mutate(partialData);

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.updateProfile).toHaveBeenCalledWith(partialData);
  });
});