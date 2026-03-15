import { describe, it, expect } from "vitest";
import { formatPrice, formatTimeAgo, statusLabel, statusColor } from "@/lib/utils";

describe("formatPrice", () => {
  it("formats number with commas", () => {
    expect(formatPrice(500000)).toBe("500,000");
    expect(formatPrice(8000000)).toBe("8,000,000");
    expect(formatPrice(0)).toBe("0");
  });

  it("returns '가격 제안' for null/undefined", () => {
    expect(formatPrice(null)).toBe("가격 제안");
    expect(formatPrice(undefined)).toBe("가격 제안");
  });
});

describe("statusLabel", () => {
  it("maps status to Korean label", () => {
    expect(statusLabel("available")).toBe("판매중");
    expect(statusLabel("reserved")).toBe("예약중");
    expect(statusLabel("pending_trade")).toBe("거래중");
    expect(statusLabel("completed")).toBe("거래완료");
    expect(statusLabel("cancelled")).toBe("취소됨");
  });

  it("returns raw status for unknown values", () => {
    expect(statusLabel("unknown_status")).toBe("unknown_status");
  });
});

describe("statusColor", () => {
  it("returns hex color for status", () => {
    expect(statusColor("available")).toBe("#059669");
    expect(statusColor("reserved")).toBe("#F59E0B");
    expect(statusColor("pending_trade")).toBe("#2563EB");
    expect(statusColor("completed")).toBe("#64748B");
    expect(statusColor("cancelled")).toBe("#DC2626");
  });

  it("returns default gray for unknown status", () => {
    expect(statusColor("unknown")).toBe("#64748B");
  });
});

describe("formatTimeAgo", () => {
  it("formats just now", () => {
    const now = new Date().toISOString();
    expect(formatTimeAgo(now)).toBe("방금 전");
  });

  it("formats recent time as minutes", () => {
    const fiveMinAgo = new Date(Date.now() - 5 * 60_000).toISOString();
    expect(formatTimeAgo(fiveMinAgo)).toBe("5분 전");
  });

  it("formats hours", () => {
    const twoHoursAgo = new Date(Date.now() - 2 * 3600_000).toISOString();
    expect(formatTimeAgo(twoHoursAgo)).toBe("2시간 전");
  });

  it("formats days", () => {
    const threeDaysAgo = new Date(Date.now() - 3 * 86400_000).toISOString();
    expect(formatTimeAgo(threeDaysAgo)).toBe("3일 전");
  });

  it("formats months", () => {
    const twoMonthsAgo = new Date(Date.now() - 65 * 86400_000).toISOString();
    expect(formatTimeAgo(twoMonthsAgo)).toBe("2개월 전");
  });
});
