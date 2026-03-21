import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createElement, type ReactNode } from "react";

/** Shared QueryClientProvider wrapper for hook tests */
export function createQueryWrapper() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });

  return function Wrapper({ children }: { children: ReactNode }) {
    return createElement(QueryClientProvider, { client: queryClient }, children);
  };
}

/** Create a QueryClient + wrapper pair (useful when you need direct cache access) */
export function createQueryClientAndWrapper() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });

  const wrapper = ({ children }: { children: ReactNode }) =>
    createElement(QueryClientProvider, { client: queryClient }, children);

  return { queryClient, wrapper };
}

/** Setup/cleanup portal root for modal tests */
export function setupPortal(id = "portal-root") {
  const el = document.createElement("div");
  el.setAttribute("id", id);
  document.body.appendChild(el);
  return el;
}

export function cleanupPortal(id = "portal-root") {
  document.getElementById(id)?.remove();
}
