"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { apiClient } from "@/lib/api-client";

export default function LoginPage() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleDevLogin = async () => {
    setLoading(true);
    setError(null);
    try {
      await apiClient.login("google", `dev_user_${Date.now()}`);
      router.push("/");
    } catch {
      setError("로그인에 실패했습니다");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="w-full max-w-sm text-center">
        <img src="/logo.png" alt="기란JT" className="h-16 mx-auto mb-2" />
        <p className="text-text-secondary mb-10">리니지 클래식 거래 플랫폼</p>

        {/* Google OAuth — TODO: integrate NextAuth.js */}
        <button
          disabled={loading}
          className="w-full flex items-center justify-center gap-3 bg-card border border-border rounded-lg py-3 px-4 text-sm text-text-primary hover:bg-medium transition-colors disabled:opacity-50"
        >
          <span className="text-lg">G</span>
          Google로 시작하기
        </button>

        {/* Dev login */}
        <button
          onClick={handleDevLogin}
          disabled={loading}
          className="w-full mt-3 flex items-center justify-center gap-3 border border-border rounded-lg py-3 px-4 text-sm text-text-secondary hover:bg-medium transition-colors disabled:opacity-50"
        >
          개발자 로그인 (테스트)
        </button>

        <button
          onClick={() => router.push("/")}
          className="mt-4 text-sm text-text-secondary hover:text-gold"
        >
          둘러보기
        </button>

        {error && (
          <p role="alert" className="text-sm text-[#e74c3c] text-center mt-4">{error}</p>
        )}
      </div>
    </div>
  );
}
