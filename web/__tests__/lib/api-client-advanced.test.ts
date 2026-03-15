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

describe("ApiClient.createChat", () => {
  it("returns chatRoomId on 201 Created", async () => {
    vi.stubGlobal("fetch", vi.fn().mockResolvedValue({
      status: 201,
      ok: true,
      json: () => Promise.resolve({ chatRoomId: "chat-new-1" }),
    }));

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.createChat("listing-1");
    expect(result).toEqual({ chatRoomId: "chat-new-1" });
  });

  it("returns chatRoomId on 409 Conflict (existing room)", async () => {
    vi.stubGlobal("fetch", vi.fn().mockResolvedValue({
      status: 409,
      ok: false,
      json: () => Promise.resolve({ chatRoomId: "chat-existing-1" }),
    }));

    ApiClientModule = await import("@/lib/api-client");
    const result = await ApiClientModule.apiClient.createChat("listing-1");
    expect(result).toEqual({ chatRoomId: "chat-existing-1" });
  });

  it("throws on other errors (500)", async () => {
    vi.stubGlobal("fetch", vi.fn().mockResolvedValue({
      status: 500,
      ok: false,
      json: () => Promise.resolve({ error: { code: "INTERNAL", message: "Server error" } }),
    }));

    ApiClientModule = await import("@/lib/api-client");
    await expect(ApiClientModule.apiClient.createChat("listing-1"))
      .rejects.toEqual({ error: { code: "INTERNAL", message: "Server error" } });
  });
});

describe("Token management", () => {
  it("saves tokens to localStorage", async () => {
    ApiClientModule = await import("@/lib/api-client");
    ApiClientModule.apiClient.saveTokens("access-123", "refresh-456");

    expect(localStorage.setItem).toHaveBeenCalledWith("accessToken", "access-123");
    expect(localStorage.setItem).toHaveBeenCalledWith("refreshToken", "refresh-456");
  });

  it("clears tokens from localStorage", async () => {
    ApiClientModule = await import("@/lib/api-client");
    ApiClientModule.apiClient.saveTokens("access-123", "refresh-456");
    ApiClientModule.apiClient.clearTokens();

    expect(localStorage.removeItem).toHaveBeenCalledWith("accessToken");
    expect(localStorage.removeItem).toHaveBeenCalledWith("refreshToken");
    expect(ApiClientModule.apiClient.isLoggedIn).toBe(false);
  });

  it("includes Authorization header when token exists", async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      status: 200,
      ok: true,
      json: () => Promise.resolve({ data: [], cursor: { hasMore: false } }),
    });
    vi.stubGlobal("fetch", mockFetch);

    ApiClientModule = await import("@/lib/api-client");
    ApiClientModule.apiClient.saveTokens("my-token", "my-refresh");
    await ApiClientModule.apiClient.getChats();

    const callHeaders = mockFetch.mock.calls[0][1].headers;
    expect(callHeaders["Authorization"]).toBe("Bearer my-token");
  });
});
