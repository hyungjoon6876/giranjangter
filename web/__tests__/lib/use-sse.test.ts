import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";

// Mock EventSource before importing the hook
let esInstances: MockEventSource[] = [];

class MockEventSource {
  url: string;
  onopen: (() => void) | null = null;
  onerror: (() => void) | null = null;
  listeners: Record<string, ((e: { data: string }) => void)[]> = {};
  closed = false;

  constructor(url: string) {
    this.url = url;
    esInstances.push(this);
  }

  addEventListener(type: string, cb: (e: { data: string }) => void) {
    this.listeners[type] = this.listeners[type] || [];
    this.listeners[type].push(cb);
  }

  close() {
    this.closed = true;
  }
}

vi.stubGlobal("EventSource", MockEventSource);

// Mock localStorage
const store: Record<string, string> = {};
const mockLocalStorage = {
  getItem: vi.fn((key: string) => store[key] ?? null),
  setItem: vi.fn((key: string, value: string) => { store[key] = value; }),
  removeItem: vi.fn((key: string) => { delete store[key]; }),
  clear: vi.fn(() => { Object.keys(store).forEach((k) => delete store[k]); }),
  get length() { return Object.keys(store).length; },
  key: vi.fn(() => null),
};
vi.stubGlobal("localStorage", mockLocalStorage);

// Must import after stubbing
import { renderHook, act, cleanup } from "@testing-library/react";
import { createElement, type ReactNode } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useSSE } from "@/lib/hooks/use-sse";

function createWrapper() {
  const qc = new QueryClient({ defaultOptions: { queries: { retry: false } } });
  return function Wrapper({ children }: { children: ReactNode }) {
    return createElement(QueryClientProvider, { client: qc }, children);
  };
}

beforeEach(() => {
  esInstances = [];
  vi.useFakeTimers();
  mockLocalStorage.setItem("accessToken", "test-token");
});

afterEach(() => {
  cleanup();
  vi.useRealTimers();
  mockLocalStorage.clear();
});

describe("useSSE", () => {
  it("starts as disconnected before token exists", () => {
    mockLocalStorage.clear();
    const { result } = renderHook(() => useSSE(), { wrapper: createWrapper() });
    expect(result.current).toBe("disconnected");
  });

  it("transitions to connected on EventSource open", () => {
    const { result } = renderHook(() => useSSE(), { wrapper: createWrapper() });
    expect(esInstances).toHaveLength(1);

    act(() => { esInstances[0].onopen?.(); });
    expect(result.current).toBe("connected");
  });

  it("transitions to reconnecting on error with exponential backoff", () => {
    const { result } = renderHook(() => useSSE(), { wrapper: createWrapper() });
    const firstEs = esInstances[0];

    // Simulate open then error
    act(() => { firstEs.onopen?.(); });
    expect(result.current).toBe("connected");

    act(() => { firstEs.onerror?.(); });
    expect(result.current).toBe("reconnecting");
    expect(firstEs.closed).toBe(true);

    // First retry: 1000ms delay (1000 * 2^0)
    act(() => { vi.advanceTimersByTime(1000); });
    expect(esInstances).toHaveLength(2);
  });

  it("applies increasing backoff delays", () => {
    const { result } = renderHook(() => useSSE(), { wrapper: createWrapper() });

    // First error — should reconnect after 1s
    act(() => { esInstances[0].onopen?.(); });
    act(() => { esInstances[0].onerror?.(); });
    act(() => { vi.advanceTimersByTime(999); });
    expect(esInstances).toHaveLength(1); // not yet reconnected
    act(() => { vi.advanceTimersByTime(1); });
    expect(esInstances).toHaveLength(2); // reconnected at 1000ms

    // Second error — should reconnect after 2s
    act(() => { esInstances[1].onerror?.(); });
    act(() => { vi.advanceTimersByTime(1999); });
    expect(esInstances).toHaveLength(2);
    act(() => { vi.advanceTimersByTime(1); });
    expect(esInstances).toHaveLength(3); // reconnected at 2000ms

    expect(result.current).toBe("reconnecting");
  });

  it("gives up after MAX_RETRIES (10) and sets disconnected", () => {
    const { result } = renderHook(() => useSSE(), { wrapper: createWrapper() });

    // Exhaust all retries
    for (let i = 0; i < 10; i++) {
      const es = esInstances[esInstances.length - 1];
      act(() => { es.onerror?.(); });
      const delay = Math.min(1000 * Math.pow(2, i), 30000);
      act(() => { vi.advanceTimersByTime(delay); });
    }

    // 11th error should set disconnected (retryCount is now 10)
    const lastEs = esInstances[esInstances.length - 1];
    act(() => { lastEs.onerror?.(); });
    expect(result.current).toBe("disconnected");
  });

  it("resets retry count on successful reconnect", () => {
    const { result } = renderHook(() => useSSE(), { wrapper: createWrapper() });

    // Error then reconnect
    act(() => { esInstances[0].onerror?.(); });
    act(() => { vi.advanceTimersByTime(1000); });
    expect(esInstances).toHaveLength(2);

    // Open succeeds — reset retry count
    act(() => { esInstances[1].onopen?.(); });
    expect(result.current).toBe("connected");

    // Error again — delay should be back to 1s (not 2s)
    act(() => { esInstances[1].onerror?.(); });
    act(() => { vi.advanceTimersByTime(1000); });
    expect(esInstances).toHaveLength(3);
  });
});
