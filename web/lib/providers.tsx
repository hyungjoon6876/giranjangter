"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState, type ReactNode } from "react";
import { useSSE } from "./hooks/use-sse";
import { ToastContext, useToastState } from "./hooks/use-toast";
import { ToastContainer } from "@/components/ui/toast";

function SSEInitializer() {
  useSSE();
  return null;
}

export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(
    () => new QueryClient({
      defaultOptions: {
        queries: { staleTime: 30_000, retry: 1 },
      },
    })
  );
  const toastState = useToastState();

  return (
    <QueryClientProvider client={queryClient}>
      <ToastContext.Provider value={toastState}>
        <SSEInitializer />
        {children}
        <ToastContainer />
      </ToastContext.Provider>
    </QueryClientProvider>
  );
}
