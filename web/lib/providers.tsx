"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState, type ReactNode } from "react";
import { useSSE, SSEContext } from "./hooks/use-sse";
import { ToastContext, useToastState } from "./hooks/use-toast";
import { ToastContainer } from "@/components/ui/toast";

function SSEProvider({ children }: { children: ReactNode }) {
  const connectionStatus = useSSE();
  return (
    <SSEContext.Provider value={connectionStatus}>
      {children}
    </SSEContext.Provider>
  );
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
        <SSEProvider>
          {children}
        </SSEProvider>
        <ToastContainer />
      </ToastContext.Provider>
    </QueryClientProvider>
  );
}
