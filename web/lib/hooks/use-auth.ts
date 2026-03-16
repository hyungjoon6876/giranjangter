"use client";

import { useSyncExternalStore } from "react";
import { apiClient } from "@/lib/api-client";

function subscribe(callback: () => void) {
  window.addEventListener("storage", callback);
  return () => window.removeEventListener("storage", callback);
}

function getSnapshot() {
  return apiClient.isLoggedIn;
}

function getServerSnapshot() {
  return false;
}

/** SSR-safe hook — returns false on server, actual auth state on client. */
export function useIsLoggedIn() {
  return useSyncExternalStore(subscribe, getSnapshot, getServerSnapshot);
}
