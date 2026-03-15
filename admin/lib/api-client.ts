import type {
  DashboardStats,
  AdminUser,
  ModerationAction,
  AuditLog,
  AdminReport,
  AdminListing,
  TradeCompletion,
  AdminChatMessage,
  DataResponse,
} from "./types";

export const API_BASE =
  process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";

class ApiClient {
  private accessToken: string | null = null;
  private refreshToken: string | null = null;

  constructor() {
    if (
      typeof window !== "undefined" &&
      typeof localStorage !== "undefined" &&
      typeof localStorage.getItem === "function"
    ) {
      this.accessToken = localStorage.getItem("accessToken");
      this.refreshToken = localStorage.getItem("refreshToken");
    }
  }

  get isLoggedIn(): boolean {
    return this.accessToken !== null;
  }

  saveTokens(access: string, refresh: string): void {
    this.accessToken = access;
    this.refreshToken = refresh;
    localStorage.setItem("accessToken", access);
    localStorage.setItem("refreshToken", refresh);
  }

  clearTokens(): void {
    this.accessToken = null;
    this.refreshToken = null;
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
  }

  private async fetch<T>(path: string, init?: RequestInit): Promise<T> {
    const headers: Record<string, string> = {
      "Content-Type": "application/json",
      ...(init?.headers as Record<string, string>),
    };
    if (this.accessToken) {
      headers["Authorization"] = `Bearer ${this.accessToken}`;
    }

    let res = await fetch(`${API_BASE}${path}`, { ...init, headers });

    if (res.status === 401 && this.refreshToken) {
      const refreshed = await this.doRefresh();
      if (refreshed) {
        headers["Authorization"] = `Bearer ${this.accessToken}`;
        res = await fetch(`${API_BASE}${path}`, { ...init, headers });
      }
    }

    if (!res.ok) {
      const err = await res
        .json()
        .catch(() => ({
          error: { code: "UNKNOWN", message: res.statusText },
        }));
      throw err;
    }

    if (res.status === 204) return undefined as T;
    return res.json();
  }

  private async doRefresh(): Promise<boolean> {
    try {
      const res = await fetch(`${API_BASE}/auth/refresh`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ refreshToken: this.refreshToken }),
      });
      if (res.ok) {
        const data = await res.json();
        this.saveTokens(data.accessToken, data.refreshToken);
        return true;
      }
    } catch {
      // refresh failed silently
    }
    return false;
  }

  // ---- Dashboard ----
  async getDashboardStats(): Promise<DashboardStats> {
    return this.fetch("/admin/stats");
  }

  // ---- Reports ----
  async getReports(): Promise<DataResponse<AdminReport>> {
    return this.fetch("/admin/reports");
  }

  async getReport(reportId: string): Promise<AdminReport> {
    return this.fetch(`/admin/reports/${reportId}`);
  }

  async reportAction(
    reportId: string,
    data: {
      actionCode: string;
      targetUserId: string;
      memo?: string;
      restrictionScope?: string;
    },
  ): Promise<{ actionId: string; status: string }> {
    return this.fetch(`/admin/reports/${reportId}/actions`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  async updateReportStatus(
    reportId: string,
    status: string,
  ): Promise<{ reportId: string; status: string }> {
    return this.fetch(`/admin/reports/${reportId}/status`, {
      method: "PATCH",
      body: JSON.stringify({ status }),
    });
  }

  // ---- Users ----
  async getUsers(params?: {
    q?: string;
    status?: string;
  }): Promise<DataResponse<AdminUser>> {
    const qs = new URLSearchParams();
    if (params?.q) qs.set("q", params.q);
    if (params?.status) qs.set("status", params.status);
    const query = qs.toString();
    return this.fetch(`/admin/users${query ? `?${query}` : ""}`);
  }

  async getUser(userId: string): Promise<AdminUser> {
    return this.fetch(`/admin/users/${userId}`);
  }

  async getUserModerationHistory(
    userId: string,
  ): Promise<DataResponse<ModerationAction>> {
    return this.fetch(`/admin/users/${userId}/moderation-history`);
  }

  async restrictUser(
    userId: string,
    data: {
      restrictionScope: string;
      durationDays?: number;
      reasonCode: string;
      memo?: string;
    },
  ): Promise<{
    actionId: string;
    targetUserId: string;
    restrictionScope: string;
  }> {
    return this.fetch(`/admin/users/${userId}/restrict`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // ---- Listings ----
  async getListings(params?: {
    status?: string;
    visibility?: string;
  }): Promise<DataResponse<AdminListing>> {
    const qs = new URLSearchParams();
    if (params?.status) qs.set("status", params.status);
    if (params?.visibility) qs.set("visibility", params.visibility);
    const query = qs.toString();
    return this.fetch(`/admin/listings${query ? `?${query}` : ""}`);
  }

  async hideListing(
    listingId: string,
  ): Promise<{ listingId: string; visibility: string }> {
    return this.fetch(`/admin/listings/${listingId}/hide`, { method: "POST" });
  }

  async restoreListing(
    listingId: string,
  ): Promise<{ listingId: string; visibility: string }> {
    return this.fetch(`/admin/listings/${listingId}/restore`, {
      method: "POST",
    });
  }

  // ---- Trades ----
  async getTrades(): Promise<DataResponse<TradeCompletion>> {
    return this.fetch("/admin/trades");
  }

  // ---- Audit Logs ----
  async getAuditLogs(): Promise<DataResponse<AuditLog>> {
    return this.fetch("/admin/audit-logs");
  }

  // ---- Chat Messages (admin review) ----
  async getChatMessages(
    chatId: string,
  ): Promise<DataResponse<AdminChatMessage>> {
    return this.fetch(`/admin/chats/${chatId}/messages`);
  }
}

export const apiClient = new ApiClient();
