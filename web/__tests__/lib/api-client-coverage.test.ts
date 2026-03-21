import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";

let ApiClientModule: typeof import("@/lib/api-client");

/** Mock fetch with a JSON response and re-import the api-client module */
async function setupFetch(responseData: unknown, status = 200) {
  const mockFetch = vi.fn().mockResolvedValue({
    status,
    ok: status >= 200 && status < 400,
    statusText: status === 500 ? "Internal Server Error" : "OK",
    json: () => Promise.resolve(responseData),
  });
  vi.stubGlobal("fetch", mockFetch);
  ApiClientModule = await import("@/lib/api-client");
  return mockFetch;
}

/** setupFetch variant with chained responses (for 401 retry flows) */
async function setupFetchSequence(...responses: Array<{ data: unknown; status?: number }>) {
  let mock = vi.fn();
  for (const r of responses) {
    mock = mock.mockResolvedValueOnce({
      status: r.status ?? 200,
      ok: (r.status ?? 200) < 400,
      json: () => Promise.resolve(r.data),
    });
  }
  vi.stubGlobal("fetch", mock);
  ApiClientModule = await import("@/lib/api-client");
  return mock;
}

beforeEach(() => {
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

// ---------- Listings ----------

describe("ApiClient.getListings", () => {
  it("constructs URL params correctly", async () => {
    const mockFetch = await setupFetch({ data: [], cursor: { hasMore: false } });
    await ApiClientModule.apiClient.getListings({
      serverId: "bartz", categoryId: "weapon", q: "집행검",
      listingType: "sell", sort: "popular", cursor: "abc123", limit: 10,
    });

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("serverId=bartz");
    expect(url).toContain("categoryId=weapon");
    expect(url).toContain("q=%EC%A7%91%ED%96%89%EA%B2%80");
    expect(url).toContain("sort=popular");
    expect(url).toContain("cursor=abc123");
    expect(url).toContain("limit=10");
  });

  it("uses default params when not provided", async () => {
    const mockFetch = await setupFetch({ data: [], cursor: { hasMore: false } });
    await ApiClientModule.apiClient.getListings();

    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("sort=recent");
    expect(url).toContain("limit=20");
  });
});

describe("ApiClient.getListing", () => {
  it("fetches single listing by ID", async () => {
    const listing = { listingId: "test-1", title: "Test Item" };
    const mockFetch = await setupFetch(listing);
    const result = await ApiClientModule.apiClient.getListing("test-1");

    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/listings/test-1"), expect.any(Object));
    expect(result).toEqual(listing);
  });
});

// ---------- Auth ----------

describe("ApiClient.login", () => {
  it("saves tokens to localStorage on success", async () => {
    const authResp = { accessToken: "access-123", refreshToken: "refresh-456", user: { userId: "u1" } };
    await setupFetch(authResp);
    const result = await ApiClientModule.apiClient.login("google", "google-token");

    expect(localStorage.setItem).toHaveBeenCalledWith("accessToken", "access-123");
    expect(localStorage.setItem).toHaveBeenCalledWith("refreshToken", "refresh-456");
    expect(result).toEqual(authResp);
  });
});

// ---------- Profile ----------

describe("ApiClient.getMe", () => {
  it("fetches user profile", async () => {
    const user = { userId: "u1", nickname: "TestUser" };
    const mockFetch = await setupFetch(user);
    const result = await ApiClientModule.apiClient.getMe();

    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/me"), expect.any(Object));
    expect(result).toEqual(user);
  });
});

describe("ApiClient.updateProfile", () => {
  it("sends PATCH request with profile data", async () => {
    const data = { nickname: "NewNick", introduction: "Hello" };
    const mockFetch = await setupFetch({ userId: "u1", ...data });
    await ApiClientModule.apiClient.updateProfile(data);

    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/me/profile"),
      expect.objectContaining({ method: "PATCH", body: JSON.stringify(data) }),
    );
  });
});

// ---------- My Resources ----------

describe("ApiClient.getMyListings", () => {
  it("includes status filter when provided", async () => {
    const mockFetch = await setupFetch({ data: [], cursor: { hasMore: false } });
    await ApiClientModule.apiClient.getMyListings("available");
    expect(mockFetch.mock.calls[0][0]).toContain("/me/listings?status=available");
  });

  it("omits status param when not provided", async () => {
    const mockFetch = await setupFetch({ data: [], cursor: { hasMore: false } });
    await ApiClientModule.apiClient.getMyListings();
    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("/me/listings");
    expect(url).not.toContain("status=");
  });
});

describe("ApiClient.getMyTrades", () => {
  it("fetches user trades", async () => {
    const mockFetch = await setupFetch({ data: [], cursor: { hasMore: false } });
    await ApiClientModule.apiClient.getMyTrades();
    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/me/trades"), expect.any(Object));
  });
});

// ---------- Notifications ----------

describe("ApiClient.getNotifications", () => {
  it("fetches notifications", async () => {
    const mockFetch = await setupFetch({ data: [], cursor: { hasMore: false } });
    await ApiClientModule.apiClient.getNotifications();
    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/notifications"), expect.any(Object));
  });
});

describe("ApiClient.markNotificationsRead", () => {
  it("sends POST with notification IDs", async () => {
    const mockFetch = await setupFetch({});
    await ApiClientModule.apiClient.markNotificationsRead(["n1", "n2"]);
    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/notifications/read"),
      expect.objectContaining({ method: "POST", body: JSON.stringify({ notificationIds: ["n1", "n2"] }) }),
    );
  });
});

// ---------- Transactions ----------

describe("ApiClient.createReservation", () => {
  it("sends POST to chat reservations endpoint", async () => {
    const data = { scheduledAt: "2024-01-01T10:00:00Z" };
    const mockFetch = await setupFetch({ reservationId: "r1" }, 201);
    await ApiClientModule.apiClient.createReservation("chat-1", data);
    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/chats/chat-1/reservations"),
      expect.objectContaining({ method: "POST", body: JSON.stringify(data) }),
    );
  });
});

describe("ApiClient.completeTrade", () => {
  it("sends POST to listing complete endpoint", async () => {
    const data = { rating: 5, feedback: "Great trade!" };
    const mockFetch = await setupFetch({ completionId: "c1" });
    await ApiClientModule.apiClient.completeTrade("listing-1", data);
    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/listings/listing-1/complete"),
      expect.objectContaining({ method: "POST", body: JSON.stringify(data) }),
    );
  });
});

describe("ApiClient.createReview", () => {
  it("sends POST to completion reviews endpoint", async () => {
    const data = { rating: 5, comment: "Excellent trader" };
    const mockFetch = await setupFetch({ reviewId: "rev1" }, 201);
    await ApiClientModule.apiClient.createReview("comp-1", data);
    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/trade-completions/comp-1/reviews"),
      expect.objectContaining({ method: "POST", body: JSON.stringify(data) }),
    );
  });
});

describe("ApiClient.createReport", () => {
  it("sends POST to reports endpoint", async () => {
    const data = { type: "spam", reason: "Posting inappropriate content" };
    const mockFetch = await setupFetch({ reportId: "rep1" }, 201);
    await ApiClientModule.apiClient.createReport(data);
    expect(mockFetch).toHaveBeenCalledWith(
      expect.stringContaining("/reports"),
      expect.objectContaining({ method: "POST", body: JSON.stringify(data) }),
    );
  });
});

// ---------- User Block ----------

describe("ApiClient.blockUser/unblockUser", () => {
  it("blockUser sends POST", async () => {
    const mockFetch = await setupFetch({});
    await ApiClientModule.apiClient.blockUser("user-123");
    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/users/user-123/block"), expect.objectContaining({ method: "POST" }));
  });

  it("unblockUser sends DELETE", async () => {
    const mockFetch = await setupFetch({});
    await ApiClientModule.apiClient.unblockUser("user-123");
    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/users/user-123/block"), expect.objectContaining({ method: "DELETE" }));
  });
});

// ---------- Master Data (null-safe) ----------

describe("ApiClient.getServers", () => {
  it("returns empty array when data is null", async () => {
    await setupFetch({ data: null });
    expect(await ApiClientModule.apiClient.getServers()).toEqual([]);
  });

  it("returns data array when present", async () => {
    const servers = [{ serverId: "bartz", name: "바츠" }];
    await setupFetch({ data: servers });
    expect(await ApiClientModule.apiClient.getServers()).toEqual(servers);
  });
});

describe("ApiClient.getCategories", () => {
  it("returns empty array when data is null", async () => {
    await setupFetch({ data: null });
    expect(await ApiClientModule.apiClient.getCategories()).toEqual([]);
  });

  it("returns data array when present", async () => {
    const cats = [{ categoryId: "weapon", name: "무기" }];
    await setupFetch({ data: cats });
    expect(await ApiClientModule.apiClient.getCategories()).toEqual(cats);
  });
});

// ---------- Upload ----------

describe("ApiClient.uploadImage", () => {
  it("sends FormData without Content-Type header", async () => {
    const mockFetch = await setupFetch({ url: "/uploaded/test.jpg" });
    ApiClientModule.apiClient.saveTokens("token", "refresh");
    await ApiClientModule.apiClient.uploadImage(new File(["test"], "test.jpg", { type: "image/jpeg" }));

    const [url, options] = mockFetch.mock.calls[0];
    expect(url).toContain("/uploads/images");
    expect(options.method).toBe("POST");
    expect(options.body).toBeInstanceOf(FormData);
    expect(options.headers["Content-Type"]).toBeUndefined();
    expect(options.headers["Authorization"]).toBe("Bearer token");
  });
});

// ---------- Search ----------

describe("ApiClient.searchItems", () => {
  it("constructs search query params correctly", async () => {
    const mockFetch = await setupFetch({ data: [] });
    await ApiClientModule.apiClient.searchItems({ q: "집행검", categoryId: "weapon" });
    const url = mockFetch.mock.calls[0][0];
    expect(url).toContain("q=%EC%A7%91%ED%96%89%EA%B2%80");
    expect(url).toContain("categoryId=weapon");
  });

  it("returns empty array when data is null", async () => {
    await setupFetch({ data: null });
    expect(await ApiClientModule.apiClient.searchItems({ q: "test" })).toEqual([]);
  });
});

describe("ApiClient.getUserReviews", () => {
  it("fetches user reviews by ID", async () => {
    const mockFetch = await setupFetch({ data: [] });
    await ApiClientModule.apiClient.getUserReviews("user-123");
    expect(mockFetch).toHaveBeenCalledWith(expect.stringContaining("/users/user-123/reviews"), expect.any(Object));
  });

  it("returns empty array when data is null", async () => {
    await setupFetch({ data: null });
    expect(await ApiClientModule.apiClient.getUserReviews("user-123")).toEqual([]);
  });
});

// ---------- Edge Cases ----------

describe("ApiClient 401 auto-refresh", () => {
  it("retries request after successful token refresh", async () => {
    const mockFetch = await setupFetchSequence(
      { data: { error: { code: "UNAUTHORIZED", message: "Token expired" } }, status: 401 },
      { data: { accessToken: "new-token", refreshToken: "new-refresh" } },
      { data: { userId: "u1", nickname: "test" } },
    );
    ApiClientModule.apiClient.saveTokens("old-token", "old-refresh");
    const result = await ApiClientModule.apiClient.getMe();

    expect(mockFetch).toHaveBeenCalledTimes(3);
    expect(result).toEqual({ userId: "u1", nickname: "test" });
  });
});

describe("ApiClient 204 response", () => {
  it("returns undefined for 204 No Content", async () => {
    const mockFetch = vi.fn().mockResolvedValue({ status: 204, ok: true });
    vi.stubGlobal("fetch", mockFetch);
    ApiClientModule = await import("@/lib/api-client");
    expect(await ApiClientModule.apiClient.markNotificationsRead(["n1"])).toBeUndefined();
  });
});

describe("ApiClient error handling", () => {
  it("handles non-JSON error responses", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 500, ok: false, statusText: "Internal Server Error",
      json: () => Promise.reject(new Error("Not JSON")),
    });
    vi.stubGlobal("fetch", mockFetch);
    ApiClientModule = await import("@/lib/api-client");
    await expect(ApiClientModule.apiClient.getMe()).rejects.toEqual({
      error: { code: "UNKNOWN", message: "Internal Server Error" },
    });
  });
});
