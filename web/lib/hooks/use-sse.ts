"use client";

import { useEffect, useRef, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { API_BASE } from "@/lib/api-client";

export function useSSE() {
  const qc = useQueryClient();
  const esRef = useRef<EventSource | null>(null);
  const [reconnectCount, setReconnectCount] = useState(0);

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
      esRef.current = null;
      setTimeout(() => setReconnectCount((c) => c + 1), 5_000);
    };

    return () => { es.close(); esRef.current = null; };
  }, [qc, reconnectCount]);
}
