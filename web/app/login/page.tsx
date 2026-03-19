"use client";

import { useEffect, useRef, useState, useCallback, Suspense } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import Image from "next/image";
import { useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

const GOOGLE_CLIENT_ID =
  "1040191360407-2masceiv4vl985m2gfr777qavrd23763.apps.googleusercontent.com";

declare global {
  interface Window {
    google?: {
      accounts: {
        id: {
          initialize: (config: Record<string, unknown>) => void;
          renderButton: (
            element: HTMLElement,
            config: Record<string, unknown>,
          ) => void;
        };
      };
    };
  }
}

function LoginContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const rawRedirect = searchParams.get("redirect") || "/";
  const redirect =
    rawRedirect.startsWith("/") && !rawRedirect.startsWith("//")
      ? rawRedirect
      : "/";
  const queryClient = useQueryClient();
  const googleBtnRef = useRef<HTMLDivElement>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleGoogleResponse = useCallback(async (response: { credential: string }) => {
    setLoading(true);
    setError(null);
    try {
      await apiClient.login("google", response.credential);
      queryClient.invalidateQueries({ queryKey: ["me"] });
      router.push(redirect);
    } catch (err: unknown) {
      const apiErr = err as { error?: { message?: string } };
      setError(apiErr?.error?.message ?? "로그인에 실패했습니다");
    } finally {
      setLoading(false);
    }
  }, [queryClient, router, redirect]);

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
  }, [handleGoogleResponse]);

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="w-full max-w-sm text-center">
        <Image src="/logo.png" alt="기란JT" width={160} height={64} className="h-16 mx-auto mb-2" />
        <p className="text-text-secondary mb-10">리니지 클래식 거래 플랫폼</p>

        {/* Google official Sign-In button */}
        <div ref={googleBtnRef} className="flex justify-center mb-3" />

        {loading && (
          <p className="text-sm text-text-secondary mt-2">로그인 중...</p>
        )}

        <button
          onClick={() => router.push("/")}
          className="mt-6 text-sm text-text-secondary hover:text-gold"
        >
          로그인 없이 둘러보기
        </button>

        {error && (
          <p role="alert" className="text-sm text-[#e74c3c] text-center mt-4">
            {error}
          </p>
        )}
      </div>
    </div>
  );
}

export default function LoginPage() {
  return (
    <Suspense>
      <LoginContent />
    </Suspense>
  );
}
