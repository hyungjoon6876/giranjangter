"use client";

import { createContext, useContext, useEffect, useRef, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { API_BASE } from "@/lib/api-client";

export type SSEConnectionStatus = "connected" | "reconnecting" | "disconnected";

export const SSEContext = createContext<SSEConnectionStatus>("disconnected");

export function useSSEConnectionStatus(): SSEConnectionStatus {
  return useContext(SSEContext);
}

const MAX_RETRIES = 10;

export function useSSE() {
  const qc = useQueryClient();
  const esRef = useRef<EventSource | null>(null);
  const retryCountRef = useRef(0);
  const retryTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const [connectionStatus, setConnectionStatus] = useState<SSEConnectionStatus>("disconnected");
  const [reconnectTrigger, setReconnectTrigger] = useState(0);

  useEffect(() => {
    const token = typeof window !== "undefined" ? localStorage.getItem("accessToken") : null;
    if (!token) return;

    const es = new EventSource(`${API_BASE}/sse/connect?token=${token}`);
    esRef.current = es;

    es.onopen = () => {
      setConnectionStatus("connected");
      retryCountRef.current = 0;
    };

    es.addEventListener("new_message", (e) => {
      const data = JSON.parse(e.data);
      qc.invalidateQueries({ queryKey: ["messages", data.chatRoomId] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    });

    es.addEventListener("status_change", () => {
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    });

    es.onerror = () => {
      es.close();
      esRef.current = null;

      if (retryCountRef.current >= MAX_RETRIES) {
        setConnectionStatus("disconnected");
        return;
      }

      setConnectionStatus("reconnecting");
      const delay = Math.min(1000 * Math.pow(2, retryCountRef.current), 30000);
      retryCountRef.current += 1;
      retryTimerRef.current = setTimeout(() => setReconnectTrigger((c) => c + 1), delay);
    };

    return () => {
      es.close();
      esRef.current = null;
      if (retryTimerRef.current) {
        clearTimeout(retryTimerRef.current);
        retryTimerRef.current = null;
      }
    };
  }, [qc, reconnectTrigger]);

  return connectionStatus;
}
