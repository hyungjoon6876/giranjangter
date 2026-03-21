import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { renderHook, waitFor } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createElement, type ReactNode } from "react";

// Mock the api-client module
vi.mock("@/lib/api-client", () => ({
  apiClient: {
    getListings: vi.fn(),
    getListing: vi.fn(),
    createListing: vi.fn(),
    updateListing: vi.fn(),
    changeListingStatus: vi.fn(),
    favoriteListing: vi.fn(),
    unfavoriteListing: vi.fn(),
  }
}));

import { apiClient } from "@/lib/api-client";
import {
  useListings,
  useListing,
  useCreateListing,
  useUpdateListing,
  useChangeListingStatus,
  useToggleFavorite,
} from "@/lib/hooks/use-listings";

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

describe("useListings", () => {
  it("calls apiClient.getListings with default params", async () => {
    const mockData = { data: [], cursor: { hasMore: false } };
    vi.mocked(apiClient.getListings).mockResolvedValue(mockData);

    const { result } = renderHook(() => useListings(), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getListings).toHaveBeenCalledWith(undefined);
    expect(result.current.data).toEqual(mockData);
  });

  it("passes filter parameters to API client", async () => {
    const params = {
      serverId: "bartz",
      categoryId: "weapon",
      q: "집행검",
      listingType: "sell",
      sort: "popular",
    };
    const mockData = { data: [], cursor: { hasMore: false } };
    vi.mocked(apiClient.getListings).mockResolvedValue(mockData);

    const { result } = renderHook(() => useListings(params), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getListings).toHaveBeenCalledWith(params);
  });

  it("uses correct query key with params", async () => {
    const params = { serverId: "bartz", sort: "recent" };
    vi.mocked(apiClient.getListings).mockResolvedValue({ data: [], cursor: { hasMore: false } });

    const { result } = renderHook(() => useListings(params), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.getListings).toHaveBeenCalledWith(params);
  });
});

describe("useListing", () => {
  it("fetches single listing when ID is provided", async () => {
    const mockListing = { listingId: "test-1", title: "Test Listing" };
    vi.mocked(apiClient.getListing).mockResolvedValue(mockListing);

    const { result } = renderHook(() => useListing("test-1"), { wrapper: createWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getListing).toHaveBeenCalledWith("test-1");
    expect(result.current.data).toEqual(mockListing);
  });

  it("does not fetch when ID is empty", () => {
    vi.mocked(apiClient.getListing).mockResolvedValue({} as any);

    const { result } = renderHook(() => useListing(""), { wrapper: createWrapper() });

    expect(apiClient.getListing).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });
});

describe("useCreateListing", () => {
  it("calls apiClient.createListing and invalidates listings cache", async () => {
    const listingData = { title: "New Listing", itemName: "Test Item" };
    const mockCreatedListing = { listingId: "new-1", ...listingData };
    vi.mocked(apiClient.createListing).mockResolvedValue(mockCreatedListing);

    const { result } = renderHook(() => useCreateListing(), { wrapper: createWrapper() });

    result.current.mutate(listingData);

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.createListing).toHaveBeenCalledWith(listingData);
    expect(result.current.data).toEqual(mockCreatedListing);
  });

  it("handles mutation error", async () => {
    const error = { error: { code: "INVALID_DATA", message: "Missing required fields" } };
    vi.mocked(apiClient.createListing).mockRejectedValue(error);

    const { result } = renderHook(() => useCreateListing(), { wrapper: createWrapper() });

    result.current.mutate({ title: "Invalid" });

    await waitFor(() => expect(result.current.isError).toBe(true));

    expect(result.current.error).toEqual(error);
  });
});

describe("useUpdateListing", () => {
  it("calls apiClient.updateListing with correct parameters", async () => {
    const updateData = { title: "Updated Title" };
    const mockUpdatedListing = { listingId: "test-1", ...updateData };
    vi.mocked(apiClient.updateListing).mockResolvedValue(mockUpdatedListing);

    const { result } = renderHook(() => useUpdateListing(), { wrapper: createWrapper() });

    result.current.mutate({ id: "test-1", data: updateData });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.updateListing).toHaveBeenCalledWith("test-1", updateData);
    expect(result.current.data).toEqual(mockUpdatedListing);
  });
});

describe("useChangeListingStatus", () => {
  it("calls apiClient.changeListingStatus with ID and status", async () => {
    vi.mocked(apiClient.changeListingStatus).mockResolvedValue();

    const { result } = renderHook(() => useChangeListingStatus(), { wrapper: createWrapper() });

    result.current.mutate({ id: "test-1", status: "sold" });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.changeListingStatus).toHaveBeenCalledWith("test-1", "sold");
  });
});

describe("useToggleFavorite", () => {
  it("calls unfavoriteListing when currently favorited", async () => {
    vi.mocked(apiClient.unfavoriteListing).mockResolvedValue();

    const { result } = renderHook(() => useToggleFavorite(), { wrapper: createWrapper() });

    result.current.mutate({ id: "test-1", isFavorited: true });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.unfavoriteListing).toHaveBeenCalledWith("test-1");
    expect(apiClient.favoriteListing).not.toHaveBeenCalled();
  });

  it("calls favoriteListing when not currently favorited", async () => {
    vi.mocked(apiClient.favoriteListing).mockResolvedValue();

    const { result } = renderHook(() => useToggleFavorite(), { wrapper: createWrapper() });

    result.current.mutate({ id: "test-1", isFavorited: false });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.favoriteListing).toHaveBeenCalledWith("test-1");
    expect(apiClient.unfavoriteListing).not.toHaveBeenCalled();
  });
});