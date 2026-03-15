"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { apiClient, API_BASE } from "@/lib/api-client";

export default function LoginPage() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  async function handleDevLogin(role: "admin" | "moderator") {
    setLoading(true);
    setError("");
    try {
      const res = await fetch(`${API_BASE}/auth/dev-login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ role }),
      });
      if (!res.ok) {
        throw new Error("로그인에 실패했습니다.");
      }
      const data = await res.json();
      apiClient.saveTokens(data.accessToken, data.refreshToken);
      router.push("/");
    } catch (err) {
      setError(err instanceof Error ? err.message : "알 수 없는 오류");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-[#f8fafc]">
      <div className="w-full max-w-sm rounded-xl border border-border bg-white p-8 shadow-sm">
        {/* Logo */}
        <div className="mb-8 text-center">
          <div className="mx-auto mb-3 flex h-12 w-12 items-center justify-center rounded-xl bg-primary">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <rect x="3" y="3" width="7" height="7" rx="1" />
              <rect x="14" y="3" width="7" height="7" rx="1" />
              <rect x="3" y="14" width="7" height="7" rx="1" />
              <rect x="14" y="14" width="7" height="7" rx="1" />
            </svg>
          </div>
          <h1 className="text-xl font-bold text-text-primary">
            기란장터 Admin
          </h1>
          <p className="mt-1 text-sm text-text-secondary">
            관리자 대시보드
          </p>
        </div>

        {/* Error */}
        {error && (
          <div className="mb-4 rounded-lg bg-red-50 px-4 py-3 text-sm text-red-700">
            {error}
          </div>
        )}

        {/* Dev Login Buttons */}
        <div className="space-y-3">
          <button
            onClick={() => handleDevLogin("admin")}
            disabled={loading}
            className="flex w-full items-center justify-center gap-2 rounded-lg bg-primary px-4 py-3 text-sm font-medium text-white transition-colors hover:bg-primary-600 disabled:opacity-50"
          >
            {loading ? "로그인 중..." : "Admin으로 로그인 (Dev)"}
          </button>
          <button
            onClick={() => handleDevLogin("moderator")}
            disabled={loading}
            className="flex w-full items-center justify-center gap-2 rounded-lg border border-border bg-white px-4 py-3 text-sm font-medium text-text-primary transition-colors hover:bg-slate-50 disabled:opacity-50"
          >
            {loading ? "로그인 중..." : "Moderator로 로그인 (Dev)"}
          </button>
        </div>

        <p className="mt-6 text-center text-xs text-text-secondary">
          개발 환경 전용 로그인
        </p>
      </div>
    </div>
  );
}
