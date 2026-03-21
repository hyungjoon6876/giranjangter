import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";

// We need to test the ApiClient class directly, so we re-import after mocking
let ApiClientModule: typeof import("@/lib/api-client");

beforeEach(() => {
  // Reset localStorage mock
  const store: Record<string, string> = {};
  vi.stubGlobal("localStorage", {
    getItem: vi.fn((key: string) => store[key] ?? null),
    setItem: vi.fn((key: string, val: string) => { store[key] = val; }),
    removeItem: vi.fn((key: string) => { delete store[key]; }),
  });
});

afterEach(() => {
  vi.restoreAllMocks();
  vi.resetModules();
});

describe("ApiClient.getListings", () => {
  it("constructs URL params correctly", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getListings({
      serverId: "bartz",
      categoryId: "weapon",
      q: "집행검",
      listingType: "sell",
      sort: "popular",
      cursor: "abc123",
      limit: 10,
    });

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("serverId=bartz");
    expect(url).toContain("categoryId=weapon");
    expect(url).toContain("q=%EC%A7%91%ED%96%89%EA%B2%80"); // URL encoded Korean
    expect(url).toContain("listingType=sell");
    expect(url).toContain("sort=popular");
    expect(url).toContain("cursor=abc123");
    expect(url).toContain("limit=10");
  });

  it("uses default params when not provided", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getListings();

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("sort=recent");
    expect(url).toContain("limit=20");
  });
});

describe("ApiClient.getListing", () => {
  it("fetches single listing by ID", async () => {
    const mockListing = { listingId: "test-1", title: "Test Item" };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve(mockListing),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getListing("test-1");

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/listings/test-1"),
      expect.any(Object)
    );
    expect(result).toEqual(mockListing);
  });
});

describe("ApiClient.login", () => {
  it("saves tokens to localStorage on success", async () => {
    const authResponse = {
      accessToken: "access-123",
      refreshToken: "refresh-456",
      user: { userId: "u1", nickname: "test" },
    };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve(authResponse),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.login("google", "google-token");

    expect(localStorage.setItem).toHaveBeenCalledWith("accessToken", "access-123");
    expect(localStorage.setItem).toHaveBeenCalledWith("refreshToken", "refresh-456");
    expect(result).toEqual(authResponse);
  });
});

describe("ApiClient.getMe", () => {
  it("fetches user profile", async () => {
    const mockUser = { userId: "u1", nickname: "TestUser" };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve(mockUser),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getMe();

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/me"),
      expect.any(Object)
    );
    expect(result).toEqual(mockUser);
  });
});

describe("ApiClient.updateProfile", () => {
  it("sends PATCH request with profile data", async () => {
    const profileData = { nickname: "NewNick", introduction: "Hello" };
    const mockUser = { userId: "u1", ...profileData };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve(mockUser),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.updateProfile(profileData);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/me/profile"),
      expect.objectContaining({
        method: "PATCH",
        body: JSON.stringify(profileData),
      })
    );
    expect(result).toEqual(mockUser);
  });
});

describe("ApiClient.getMyListings", () => {
  it("includes status filter when provided", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getMyListings("available");

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("/me/listings?status=available");
  });

  it("omits status param when not provided", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getMyListings();

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("/me/listings");
    expect(url).not.toContain("status=");
  });
});

describe("ApiClient.getMyTrades", () => {
  it("fetches user trades", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getMyTrades();

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/me/trades"),
      expect.any(Object)
    );
  });
});

describe("ApiClient.getNotifications", () => {
  it("fetches notifications", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getNotifications();

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/notifications"),
      expect.any(Object)
    );
  });
});

describe("ApiClient.markNotificationsRead", () => {
  it("sends POST with notification IDs", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({}),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.markNotificationsRead(["n1", "n2"]);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/notifications/read"),
      expect.objectContaining({
        method: "POST",
        body: JSON.stringify({ notificationIds: ["n1", "n2"] }),
      })
    );
  });
});

describe("ApiClient.createReservation", () => {
  it("sends POST to chat reservations endpoint", async () => {
    const reservationData = { scheduledAt: "2024-01-01T10:00:00Z" };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 201,
      ok: true,
      json: () => Promise.resolve({ reservationId: "r1" }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.createReservation("chat-1", reservationData);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/chats/chat-1/reservations"),
      expect.objectContaining({
        method: "POST",
        body: JSON.stringify(reservationData),
      })
    );
  });
});

describe("ApiClient.completeTrade", () => {
  it("sends POST to listing complete endpoint", async () => {
    const tradeData = { rating: 5, feedback: "Great trade!" };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ completionId: "c1" }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.completeTrade("listing-1", tradeData);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/listings/listing-1/complete"),
      expect.objectContaining({
        method: "POST",
        body: JSON.stringify(tradeData),
      })
    );
  });
});

describe("ApiClient.createReview", () => {
  it("sends POST to completion reviews endpoint", async () => {
    const reviewData = { rating: 5, comment: "Excellent trader" };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 201,
      ok: true,
      json: () => Promise.resolve({ reviewId: "rev1" }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.createReview("comp-1", reviewData);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/trade-completions/comp-1/reviews"),
      expect.objectContaining({
        method: "POST",
        body: JSON.stringify(reviewData),
      })
    );
  });
});

describe("ApiClient.createReport", () => {
  it("sends POST to reports endpoint", async () => {
    const reportData = { type: "spam", reason: "Posting inappropriate content" };
    const mockFetch = vi.fn().mockResolvedValue({
      status: 201,
      ok: true,
      json: () => Promise.resolve({ reportId: "rep1" }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.createReport(reportData);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/reports"),
      expect.objectContaining({
        method: "POST",
        body: JSON.stringify(reportData),
      })
    );
  });
});

describe("ApiClient.blockUser/unblockUser", () => {
  it("blockUser sends POST to user block endpoint", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({}),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.blockUser("user-123");

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/users/user-123/block"),
      expect.objectContaining({ method: "POST" })
    );
  });

  it("unblockUser sends DELETE to user block endpoint", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({}),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.unblockUser("user-123");

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/users/user-123/block"),
      expect.objectContaining({ method: "DELETE" })
    );
  });
});

describe("ApiClient.getServers", () => {
  it("returns empty array when data is null", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: null }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getServers();

    expect(result).toEqual([]);
  });

  it("returns data array when present", async () => {
    const servers = [{ serverId: "bartz", name: "바츠" }];
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: servers }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getServers();

    expect(result).toEqual(servers);
  });
});

describe("ApiClient.getCategories", () => {
  it("returns empty array when data is null", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: null }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getCategories();

    expect(result).toEqual([]);
  });

  it("returns data array when present", async () => {
    const categories = [{ categoryId: "weapon", name: "무기" }];
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: categories }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getCategories();

    expect(result).toEqual(categories);
  });
});

describe("ApiClient.uploadImage", () => {
  it("sends FormData without Content-Type header", async () => {
    const mockFile = new File(["test"], "test.jpg", { type: "image/jpeg" });
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ url: "/uploaded/test.jpg" }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    ApiClientModule.apiClient.saveTokens("token", "refresh");
    await ApiClientModule.apiClient.uploadImage(mockFile);

    const [url, options] = mockFetch.mock.calls[0];
    expect(url).toContain("/uploads/images");
    expect(options.method).toBe("POST");
    expect(options.body).toBeInstanceOf(FormData);
    expect(options.headers["Content-Type"]).toBeUndefined(); // Should NOT be set for FormData
    expect(options.headers["Authorization"]).toBe("Bearer token");
  });
});

describe("ApiClient.searchItems", () => {
  it("constructs search query params correctly", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [] }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.searchItems({
      q: "집행검",
      categoryId: "weapon",
    });

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("q=%EC%A7%91%ED%96%89%EA%B2%80");
    expect(url).toContain("categoryId=weapon");
  });

  it("returns empty array when data is null", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: null }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.searchItems({ q: "test" });

    expect(result).toEqual([]);
  });
});

describe("ApiClient.getUserReviews", () => {
  it("fetches user reviews by ID", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [] }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await ApiClientModule.apiClient.getUserReviews("user-123");

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/users/user-123/reviews"),
      expect.any(Object)
    );
  });

  it("returns empty array when data is null", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: null }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.getUserReviews("user-123");

    expect(result).toEqual([]);
  });
});

describe("ApiClient 401 auto-refresh flow", () => {
  it("retries request after successful token refresh", async () => {
    // First call gets 401, refresh succeeds, retry succeeds
    const mockFetch = vi.fn()
      .mockResolvedValueOnce({
        status: 401,
        ok: false,
        json: () => Promise.resolve({ error: { code: "UNAUTHORIZED", message: "Token expired" } }),
      })
      .mockResolvedValueOnce({
        status: 200,
        ok: true,
        json: () => Promise.resolve({ accessToken: "new-token", refreshToken: "new-refresh" }),
      })
      .mockResolvedValueOnce({
        status: 200,
        ok: true,
        json: () => Promise.resolve({ userId: "u1", nickname: "test" }),
      });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    ApiClientModule.apiClient.saveTokens("old-token", "old-refresh");
    const result = await ApiClientModule.apiClient.getMe();

    expect(mockFetch).toHaveBeenCalledTimes(3); // original + refresh + retry
    expect(result).toEqual({ userId: "u1", nickname: "test" });
  });
});

describe("ApiClient 204 response handling", () => {
  it("returns undefined for 204 No Content", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 204,
      ok: true,
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.markNotificationsRead(["n1"]);

    expect(result).toBeUndefined();
  });
});

describe("ApiClient error handling", () => {
  it("handles non-JSON error responses", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 500,
      ok: false,
      statusText: "Internal Server Error",
      json: () => Promise.reject(new Error("Not JSON")),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    await expect(ApiClientModule.apiClient.getMe()).rejects.toEqual({
      error: { code: "UNKNOWN", message: "Internal Server Error" }
    });
  });
});