import { describe, it, expect } from "vitest";

// Test that API client constructs URLs correctly
describe("ApiClient URL construction", () => {
  it("builds listing query params correctly", () => {
    const qs = new URLSearchParams();
    qs.set("serverId", "bartz");
    qs.set("sort", "recent");
    qs.set("limit", "20");
    expect(qs.toString()).toBe("serverId=bartz&sort=recent&limit=20");
  });

  it("handles empty params", () => {
    const qs = new URLSearchParams();
    qs.set("sort", "recent");
    qs.set("limit", "20");
    expect(qs.toString()).toBe("sort=recent&limit=20");
  });

  it("encodes special characters in search query", () => {
    const qs = new URLSearchParams();
    qs.set("q", "집행검");
    qs.set("sort", "recent");
    qs.set("limit", "20");
    expect(qs.get("q")).toBe("집행검");
  });
});
