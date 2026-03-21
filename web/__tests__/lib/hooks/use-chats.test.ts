import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { renderHook, waitFor, act } from "@testing-library/react";
import { createQueryWrapper } from "@/__tests__/test-utils";

vi.mock("@/lib/api-client", () => ({
  apiClient: {
    isLoggedIn: false,
    getChats: vi.fn(),
    getMessages: vi.fn(),
    sendMessage: vi.fn(),
    markRead: vi.fn(),
    createChat: vi.fn(),
  }
}));

import { apiClient } from "@/lib/api-client";
import {
  useChats,
  useMessages,
  useSendMessage,
  useCreateChat,
  useMarkRead,
} from "@/lib/hooks/use-chats";

beforeEach(() => {
  vi.clearAllMocks();
});

afterEach(() => {
  vi.restoreAllMocks();
});

describe("useChats", () => {
  it("does not fetch when user is not logged in", () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: false, configurable: true });

    const { result } = renderHook(() => useChats(), { wrapper: createQueryWrapper() });

    expect(apiClient.getChats).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("fetches chats when user is logged in", async () => {
    Object.defineProperty(apiClient, 'isLoggedIn', { value: true, configurable: true });
    const mockChats = { data: [{ chatRoomId: "chat-1", lastMessage: "Hello" }], cursor: { hasMore: false } };
    vi.mocked(apiClient.getChats).mockResolvedValue(mockChats);

    const { result } = renderHook(() => useChats(), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getChats).toHaveBeenCalled();
    expect(result.current.data).toEqual(mockChats);
  });
});

describe("useMessages", () => {
  it("fetches messages when chatId is provided", async () => {
    const mockMessages = {
      data: [{ messageId: "msg-1", bodyText: "Hello", sentAt: "2024-01-01T10:00:00Z" }],
      cursor: { hasMore: false }
    };
    vi.mocked(apiClient.getMessages).mockResolvedValue(mockMessages);

    const { result } = renderHook(() => useMessages("chat-1"), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.getMessages).toHaveBeenCalledWith("chat-1", undefined);
    expect(result.current.data?.pages[0]).toEqual(mockMessages);
  });

  it("does not fetch when chatId is empty", () => {
    const { result } = renderHook(() => useMessages(""), { wrapper: createQueryWrapper() });

    expect(apiClient.getMessages).not.toHaveBeenCalled();
    expect(result.current.fetchStatus).toBe("idle");
  });

  it("returns hasNextPage based on cursor", async () => {
    const page = {
      data: [{ messageId: "msg-1", bodyText: "Hello", sentAt: "2024-01-01T10:00:00Z" }],
      cursor: { hasMore: true, next: "cursor-1" }
    };
    vi.mocked(apiClient.getMessages).mockResolvedValue(page);

    const { result } = renderHook(() => useMessages("chat-1"), { wrapper: createQueryWrapper() });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(result.current.hasNextPage).toBe(true);
  });

  it("fetches with SSE connected flag", async () => {
    vi.mocked(apiClient.getMessages).mockResolvedValue({ data: [], cursor: { hasMore: false } });

    const { result } = renderHook(
      () => useMessages("chat-1", true),
      { wrapper: createQueryWrapper() }
    );

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.getMessages).toHaveBeenCalledWith("chat-1", undefined);
  });
});

describe("useSendMessage", () => {
  it("calls apiClient.sendMessage with correct parameters", async () => {
    const mockMessage = {
      messageId: "server-msg-1",
      chatRoomId: "chat-1",
      senderUserId: "user-1",
      messageType: "text" as const,
      bodyText: "Hello world",
      sentAt: "2024-01-01T10:00:00Z",
      status: "sent" as const,
    };
    vi.mocked(apiClient.sendMessage).mockResolvedValue(mockMessage);

    const { result } = renderHook(() => useSendMessage(), { wrapper: createQueryWrapper() });

    result.current.mutate({
      chatId: "chat-1",
      text: "Hello world",
      clientMessageId: "client-1",
    });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.sendMessage).toHaveBeenCalledWith("chat-1", "Hello world", "client-1");
    expect(result.current.data).toEqual(mockMessage);
  });

  it("calls sendMessage with correct parameters", async () => {
    const mockMessage = {
      messageId: "server-1",
      chatRoomId: "chat-1",
      senderUserId: "user-1",
      messageType: "text" as const,
      bodyText: "Hello world",
      sentAt: "2024-01-01T10:00:00Z",
      status: "sent" as const,
    };
    vi.mocked(apiClient.sendMessage).mockResolvedValue(mockMessage);

    const { result } = renderHook(() => useSendMessage(), { wrapper: createQueryWrapper() });

    result.current.mutate({
      chatId: "chat-1",
      text: "Hello world",
      clientMessageId: "client-1",
    });

    await waitFor(() => expect(result.current.isSuccess).toBe(true));
    expect(apiClient.sendMessage).toHaveBeenCalledWith("chat-1", "Hello world", "client-1");
  });
});

describe("useCreateChat", () => {
  it("calls apiClient.createChat with listing ID", async () => {
    const mockResult = { chatRoomId: "new-chat-1" };
    vi.mocked(apiClient.createChat).mockResolvedValue(mockResult);

    const { result } = renderHook(() => useCreateChat(), { wrapper: createQueryWrapper() });

    result.current.mutate("listing-123");

    await waitFor(() => expect(result.current.isSuccess).toBe(true));

    expect(apiClient.createChat).toHaveBeenCalledWith("listing-123");
    expect(result.current.data).toEqual(mockResult);
  });

  it("handles chat creation error", async () => {
    const error = { error: { code: "FORBIDDEN", message: "Cannot create chat with own listing" } };
    vi.mocked(apiClient.createChat).mockRejectedValue(error);

    const { result } = renderHook(() => useCreateChat(), { wrapper: createQueryWrapper() });

    result.current.mutate("listing-123");

    await waitFor(() => expect(result.current.isError).toBe(true));

    expect(result.current.error).toEqual(error);
  });
});

describe("useMarkRead", () => {
  it("calls markRead with last valid message ID", async () => {
    vi.mocked(apiClient.markRead).mockResolvedValue();

    const messages = [
      {
        messageId: "msg-1",
        senderUserId: "user-1",
        status: "sent" as const,
        bodyText: "First",
        chatRoomId: "chat-1",
        messageType: "text" as const,
        sentAt: "2024-01-01T10:00:00Z",
      },
      {
        messageId: "msg-2",
        senderUserId: "user-2",
        status: "sent" as const,
        bodyText: "Second",
        chatRoomId: "chat-1",
        messageType: "text" as const,
        sentAt: "2024-01-01T10:01:00Z",
      },
    ];

    renderHook(() => useMarkRead("chat-1", messages), { wrapper: createQueryWrapper() });

    await waitFor(() => {
      expect(apiClient.markRead).toHaveBeenCalledWith("chat-1", "msg-2");
    });
  });

  it("ignores messages with 'sending' status", async () => {
    vi.mocked(apiClient.markRead).mockResolvedValue();

    const messages = [
      {
        messageId: "msg-1",
        senderUserId: "user-1",
        status: "sent" as const,
        bodyText: "Sent message",
        chatRoomId: "chat-1",
        messageType: "text" as const,
        sentAt: "2024-01-01T10:00:00Z",
      },
      {
        messageId: "msg-2",
        senderUserId: "user-1",
        status: "sending" as const,
        bodyText: "Sending message",
        chatRoomId: "chat-1",
        messageType: "text" as const,
        sentAt: "2024-01-01T10:01:00Z",
      },
    ];

    renderHook(() => useMarkRead("chat-1", messages), { wrapper: createQueryWrapper() });

    await waitFor(() => {
      expect(apiClient.markRead).toHaveBeenCalledWith("chat-1", "msg-1");
    });
  });

  it("does not call markRead when no valid messages", () => {
    const messages = [
      {
        messageId: "msg-1",
        senderUserId: "",
        status: "sending" as const,
        bodyText: "Invalid",
        chatRoomId: "chat-1",
        messageType: "text" as const,
        sentAt: "2024-01-01T10:00:00Z",
      },
    ];

    renderHook(() => useMarkRead("chat-1", messages), { wrapper: createQueryWrapper() });

    expect(apiClient.markRead).not.toHaveBeenCalled();
  });

  it("handles markRead API errors gracefully", async () => {
    vi.mocked(apiClient.markRead).mockRejectedValue(new Error("Network error"));

    const messages = [
      {
        messageId: "msg-1",
        senderUserId: "user-1",
        status: "sent" as const,
        bodyText: "Message",
        chatRoomId: "chat-1",
        messageType: "text" as const,
        sentAt: "2024-01-01T10:00:00Z",
      },
    ];

    // Should not throw error
    expect(() => {
      renderHook(() => useMarkRead("chat-1", messages), { wrapper: createQueryWrapper() });
    }).not.toThrow();
  });
});