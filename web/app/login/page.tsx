"use client";

import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

const GOOGLE_CLIENT_ID = "1040191360407-2masceiv4vl985m2gfr777qavrd23763.apps.googleusercontent.com";

declare global {
  interface Window {
    google?: {
      accounts: {
        id: {
          initialize: (config: Record<string, unknown>) => void;
          renderButton: (element: HTMLElement, config: Record<string, unknown>) => void;
        };
      };
    };
  }
}

export default function LoginPage() {
  const router = useRouter();
  const queryClient = useQueryClient();
  const googleBtnRef = useRef<HTMLDivElement>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const initGoogle = () => {
      if (!window.google || !googleBtnRef.current) return;

      window.google.accounts.id.initialize({
        client_id: GOOGLE_CLIENT_ID,
        callback: handleGoogleResponse,
      });

      window.google.accounts.id.renderButton(googleBtnRef.current, {
        theme: "filled_black",
        size: "large",
        width: "360",
        text: "signin_with",
        shape: "rectangular",
        logo_alignment: "left",
      });
    };

    // Wait for Google script to load
    if (window.google) {
      initGoogle();
    } else {
      const interval = setInterval(() => {
        if (window.google) {
          clearInterval(interval);
          initGoogle();
        }
      }, 100);
      return () => clearInterval(interval);
    }
  }, []);

  const handleGoogleResponse = async (response: { credential: string }) => {
    setLoading(true);
    setError(null);
    try {
      await apiClient.login("google", response.credential);
      queryClient.invalidateQueries({ queryKey: ["me"] });
      router.push("/");
    } catch (err: unknown) {
      const apiErr = err as { error?: { message?: string } };
      setError(apiErr?.error?.message ?? "로그인에 실패했습니다");
    } finally {
      setLoading(false);
    }
  };

  // Dev login (only show when no GOOGLE_CLIENT_ID or for testing)
  const handleDevLogin = async () => {
    setLoading(true);
    setError(null);
    try {
      await apiClient.login("google", `dev_user_${Date.now()}`);
      queryClient.invalidateQueries({ queryKey: ["me"] });
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
        <h1 className="text-3xl font-bold font-display text-gold mb-2">기란장터</h1>
        <p className="text-text-secondary mb-10">리니지 클래식 거래 플랫폼</p>

        {/* Google official Sign-In button */}
        <div ref={googleBtnRef} className="flex justify-center mb-3" />

        {loading && (
          <p className="text-sm text-text-secondary mt-2">로그인 중...</p>
        )}

        {/* Dev login - for testing */}
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
