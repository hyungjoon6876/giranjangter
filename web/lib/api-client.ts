import type {
  Listing, ChatRoom, Message, Server, Category,
  PaginatedResponse, AuthResponse, User, Notification,
} from "./types";

export const API_BASE = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";

export function assetUrl(path: string): string {
  return `${API_BASE.replace(/\/api\/v\d+$/, "")}${path}`;
}

class ApiClient {
  private accessToken: string | null = null;
  private refreshToken: string | null = null;

  constructor() {
    if (typeof window !== "undefined" && typeof localStorage !== "undefined" && typeof localStorage.getItem === "function") {
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

    // Auto-refresh on 401
    if (res.status === 401 && this.refreshToken) {
      const refreshed = await this.doRefresh();
      if (refreshed) {
        headers["Authorization"] = `Bearer ${this.accessToken}`;
        res = await fetch(`${API_BASE}${path}`, { ...init, headers });
      }
    }

    if (!res.ok) {
      const err = await res.json().catch(() => ({ error: { code: "UNKNOWN", message: res.statusText } }));
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

  // Auth
  async login(provider: string, token: string): Promise<AuthResponse> {
    const data = await this.fetch<AuthResponse>("/auth/login", {
      method: "POST",
      body: JSON.stringify({ provider, providerToken: token }),
    });
    this.saveTokens(data.accessToken, data.refreshToken);
    return data;
  }

  async getMe(): Promise<User> {
    return this.fetch("/me");
  }

  // Listings
  async getListings(params?: {
    serverId?: string; categoryId?: string; q?: string;
    listingType?: string; sort?: string; cursor?: string; limit?: number;
  }): Promise<PaginatedResponse<Listing>> {
    const qs = new URLSearchParams();
    if (params?.serverId) qs.set("serverId", params.serverId);
    if (params?.categoryId) qs.set("categoryId", params.categoryId);
    if (params?.q) qs.set("q", params.q);
    if (params?.listingType) qs.set("listingType", params.listingType);
    qs.set("sort", params?.sort ?? "recent");
    qs.set("limit", String(params?.limit ?? 20));
    if (params?.cursor) qs.set("cursor", params.cursor);
    return this.fetch(`/listings?${qs}`);
  }

  async getListing(id: string): Promise<Listing> {
    return this.fetch(`/listings/${id}`);
  }

  async createListing(data: Partial<Listing>): Promise<Listing> {
    return this.fetch("/listings", { method: "POST", body: JSON.stringify(data) });
  }

  async favoriteListing(id: string): Promise<void> {
    return this.fetch(`/listings/${id}/favorite`, { method: "POST" });
  }

  async unfavoriteListing(id: string): Promise<void> {
    return this.fetch(`/listings/${id}/favorite`, { method: "DELETE" });
  }

  // Chat
  async createChat(listingId: string): Promise<{ chatRoomId: string }> {
    return this.fetch(`/listings/${listingId}/chats`, { method: "POST" });
  }

  async getChats(): Promise<PaginatedResponse<ChatRoom>> {
    return this.fetch("/chats");
  }

  async getMessages(chatId: string): Promise<PaginatedResponse<Message>> {
    return this.fetch(`/chats/${chatId}/messages`);
  }

  async sendMessage(chatId: string, text: string): Promise<Message> {
    return this.fetch(`/chats/${chatId}/messages`, {
      method: "POST",
      body: JSON.stringify({ messageType: "text", bodyText: text }),
    });
  }

  // Profile
  async getMyListings(status?: string): Promise<PaginatedResponse<Listing>> {
    const qs = status ? `?status=${status}` : "";
    return this.fetch(`/me/listings${qs}`);
  }

  async getMyTrades(): Promise<PaginatedResponse<Listing>> {
    return this.fetch("/me/trades");
  }

  async getNotifications(): Promise<PaginatedResponse<Notification>> {
    return this.fetch("/notifications");
  }

  async markNotificationsRead(ids: string[]): Promise<void> {
    return this.fetch("/notifications/read", {
      method: "POST",
      body: JSON.stringify({ notificationIds: ids }),
    });
  }

  // Reservation
  async createReservation(chatId: string, data: Record<string, unknown>): Promise<unknown> {
    return this.fetch(`/chats/${chatId}/reservations`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // Trade
  async completeTrade(listingId: string, data: Record<string, unknown>): Promise<unknown> {
    return this.fetch(`/listings/${listingId}/complete`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // Review
  async createReview(completionId: string, data: Record<string, unknown>): Promise<unknown> {
    return this.fetch(`/trade-completions/${completionId}/reviews`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  // Report
  async createReport(data: Record<string, unknown>): Promise<unknown> {
    return this.fetch("/reports", { method: "POST", body: JSON.stringify(data) });
  }

  // Master data
  async getServers(): Promise<Server[]> {
    const res = await this.fetch<{ data: Server[] }>("/servers");
    return res.data;
  }

  async getCategories(): Promise<Category[]> {
    const res = await this.fetch<{ data: Category[] }>("/categories");
    return res.data;
  }
}

export const apiClient = new ApiClient();
