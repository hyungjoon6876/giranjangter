"use client";

import { createContext, useCallback, useContext, useEffect, useRef, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { API_BASE } from "@/lib/api-client";

export type SSEConnectionStatus = "connected" | "reconnecting" | "disconnected";

export const SSEContext = createContext<SSEConnectionStatus>("disconnected");

export function useSSEConnectionStatus(): SSEConnectionStatus {
  return useContext(SSEContext);
}

export function useSSE() {
  const qc = useQueryClient();
  const esRef = useRef<EventSource | null>(null);
  const retryCountRef = useRef(0);
  const retryTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const lastEventRef = useRef<number>(Date.now());
  const heartbeatIntervalRef = useRef<ReturnType<typeof setInterval> | null>(null);
  const [connectionStatus, setConnectionStatus] = useState<SSEConnectionStatus>("disconnected");
  const statusRef = useRef<SSEConnectionStatus>("disconnected");
  const [reconnectTrigger, setReconnectTrigger] = useState(0);

  // Keep statusRef in sync with state
  useEffect(() => {
    statusRef.current = connectionStatus;
  }, [connectionStatus]);

  const reconnect = useCallback(() => {
    // Close existing connection if any
    if (esRef.current) {
      esRef.current.close();
      esRef.current = null;
    }
    setReconnectTrigger((c) => c + 1);
  }, []);

  // SSE connection effect
  useEffect(() => {
    const token = typeof window !== "undefined" ? localStorage.getItem("accessToken") : null;
    if (!token) return;

    const es = new EventSource(`${API_BASE}/sse/connect?token=${token}`);
    esRef.current = es;

    es.onopen = () => {
      retryCountRef.current = 0;
      lastEventRef.current = Date.now();
      setConnectionStatus("connected");
    };

    // Heartbeat listener — server sends heartbeat every 30s
    es.addEventListener("heartbeat", () => {
      lastEventRef.current = Date.now();
    });

    es.addEventListener("new_message", (e) => {
      lastEventRef.current = Date.now();
      try {
        const data = JSON.parse(e.data);
        qc.invalidateQueries({ queryKey: ["messages", data.chatRoomId] });
        qc.invalidateQueries({ queryKey: ["chats"] });
      } catch (err) {
        console.error("[SSE] Failed to parse new_message:", err);
      }
    });

    es.addEventListener("status_change", (e) => {
      lastEventRef.current = Date.now();
      try {
        if (e.data) JSON.parse(e.data);
        qc.invalidateQueries({ queryKey: ["listings"] });
        qc.invalidateQueries({ queryKey: ["chats"] });
      } catch (err) {
        console.error("[SSE] Failed to parse status_change:", err);
      }
    });

    es.addEventListener("read_receipt", (e) => {
      lastEventRef.current = Date.now();
      try {
        const data = JSON.parse(e.data);
        qc.invalidateQueries({ queryKey: ["messages", data.chatRoomId] });
        qc.invalidateQueries({ queryKey: ["chats"] });
      } catch (err) {
        console.error("[SSE] Failed to parse read_receipt:", err);
      }
    });

    es.onerror = () => {
      es.close();
      esRef.current = null;

      setConnectionStatus("reconnecting");
      const delay = Math.min(1000 * Math.pow(2, retryCountRef.current), 60_000);
      retryCountRef.current += 1;
      retryTimerRef.current = setTimeout(() => setReconnectTrigger((c) => c + 1), delay);
    };

    // Heartbeat timeout check — if no event for 45s, connection is stale
    heartbeatIntervalRef.current = setInterval(() => {
      if (Date.now() - lastEventRef.current > 45_000) {
        es.close();
        esRef.current = null;
        setConnectionStatus("reconnecting");
        setReconnectTrigger((c) => c + 1);
      }
    }, 15_000);

    return () => {
      es.close();
      esRef.current = null;
      if (retryTimerRef.current) {
        clearTimeout(retryTimerRef.current);
        retryTimerRef.current = null;
      }
      if (heartbeatIntervalRef.current) {
        clearInterval(heartbeatIntervalRef.current);
        heartbeatIntervalRef.current = null;
      }
    };
  }, [qc, reconnectTrigger]);

  // Tab visibility + network recovery
  useEffect(() => {
    const handleVisibility = () => {
      if (document.visibilityState === "visible" && statusRef.current !== "connected") {
        reconnect();
      }
    };
    const handleOnline = () => {
      if (statusRef.current !== "connected") {
        reconnect();
      }
    };

    document.addEventListener("visibilitychange", handleVisibility);
    window.addEventListener("online", handleOnline);
    return () => {
      document.removeEventListener("visibilitychange", handleVisibility);
      window.removeEventListener("online", handleOnline);
    };
  }, [reconnect]);

  return connectionStatus;
}
