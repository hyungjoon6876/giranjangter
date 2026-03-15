"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { apiClient } from "@/lib/api-client";

export default function LoginPage() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const handleDevLogin = async () => {
    setLoading(true);
    try {
      await apiClient.login("google", `dev_user_${Date.now()}`);
      router.push("/");
    } catch (e) {
      alert(`로그인 실패: ${e}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="w-full max-w-sm text-center">
        <div className="text-6xl mb-4">🏪</div>
        <h1 className="text-3xl font-bold text-primary mb-2">기란장터</h1>
        <p className="text-text-secondary mb-10">리니지 클래식 거래 플랫폼</p>

        {/* Google OAuth — TODO: integrate NextAuth.js */}
        <button
          disabled={loading}
          className="w-full flex items-center justify-center gap-3 bg-white border border-border rounded-lg py-3 px-4 text-sm hover:bg-surface transition-colors disabled:opacity-50"
        >
          <span className="text-lg">G</span>
          Google로 시작하기
        </button>

        {/* Dev login */}
        <button
          onClick={handleDevLogin}
          disabled={loading}
          className="w-full mt-3 flex items-center justify-center gap-3 border border-border rounded-lg py-3 px-4 text-sm text-text-secondary hover:bg-surface transition-colors disabled:opacity-50"
        >
          🛠️ 개발자 로그인 (테스트)
        </button>

        <button
          onClick={() => router.push("/")}
          className="mt-4 text-sm text-text-secondary hover:text-primary"
        >
          둘러보기
        </button>
      </div>
    </div>
  );
}
