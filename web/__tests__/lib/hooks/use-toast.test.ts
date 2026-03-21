import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { renderHook, act } from "@testing-library/react";
import { useToastState } from "@/lib/hooks/use-toast";

beforeEach(() => {
  vi.useFakeTimers();
});

afterEach(() => {
  vi.useRealTimers();
});

describe("useToastState", () => {
  it("starts with empty toasts", () => {
    const { result } = renderHook(() => useToastState());
    expect(result.current.toasts).toEqual([]);
  });

  it("adds a toast", () => {
    const { result } = renderHook(() => useToastState());

    act(() => {
      result.current.addToast("success", "저장되었습니다");
    });

    expect(result.current.toasts).toHaveLength(1);
    expect(result.current.toasts[0].type).toBe("success");
    expect(result.current.toasts[0].message).toBe("저장되었습니다");
  });

  it("adds multiple toasts", () => {
    const { result } = renderHook(() => useToastState());

    act(() => {
      result.current.addToast("success", "First");
      result.current.addToast("error", "Second");
      result.current.addToast("info", "Third");
    });

    expect(result.current.toasts).toHaveLength(3);
  });

  it("removes a toast by id", () => {
    const { result } = renderHook(() => useToastState());

    act(() => {
      result.current.addToast("info", "to be removed");
    });

    const toastId = result.current.toasts[0].id;

    act(() => {
      result.current.removeToast(toastId);
    });

    expect(result.current.toasts).toHaveLength(0);
  });

  it("auto-removes toast after 5 seconds", () => {
    const { result } = renderHook(() => useToastState());

    act(() => {
      result.current.addToast("info", "auto-remove");
    });

    expect(result.current.toasts).toHaveLength(1);

    act(() => {
      vi.advanceTimersByTime(5000);
    });

    expect(result.current.toasts).toHaveLength(0);
  });

  it("clears timer when toast is manually removed", () => {
    const { result } = renderHook(() => useToastState());

    act(() => {
      result.current.addToast("info", "manual remove");
    });

    const toastId = result.current.toasts[0].id;

    act(() => {
      result.current.removeToast(toastId);
    });

    // Advancing time should not cause issues
    act(() => {
      vi.advanceTimersByTime(5000);
    });

    expect(result.current.toasts).toHaveLength(0);
  });

  it("generates unique ids for each toast", () => {
    const { result } = renderHook(() => useToastState());

    act(() => {
      result.current.addToast("info", "A");
      result.current.addToast("info", "B");
    });

    expect(result.current.toasts[0].id).not.toBe(result.current.toasts[1].id);
  });
});
