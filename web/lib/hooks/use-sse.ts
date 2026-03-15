"use client";

import { useEffect, useRef } from "react";
import { useQueryClient } from "@tanstack/react-query";

const API_BASE = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";

export function useSSE() {
  const qc = useQueryClient();
  const esRef = useRef<EventSource | null>(null);

  useEffect(() => {
    const token = typeof window !== "undefined" ? localStorage.getItem("accessToken") : null;
    if (!token) return;

    const es = new EventSource(`${API_BASE}/sse/connect?token=${token}`);
    esRef.current = es;

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
      // Auto-reconnect after 5 seconds
      setTimeout(() => {
        esRef.current = null;
      }, 5_000);
    };

    return () => {
      es.close();
      esRef.current = null;
    };
  }, [qc]);
}
