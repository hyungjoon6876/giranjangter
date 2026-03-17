"use client";

import { useRouter, usePathname } from "next/navigation";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";

export function useAuthGuard() {
  const router = useRouter();
  const pathname = usePathname();
  const { addToast } = useToast();

  const requireAuth = (action?: string): boolean => {
    if (apiClient.isLoggedIn) return true;
    addToast(
      "info",
      action ? `${action}은(는) 로그인이 필요합니다` : "로그인이 필요합니다",
    );
    const redirect = encodeURIComponent(pathname);
    router.push(`/login?redirect=${redirect}`);
    return false;
  };

  return { isLoggedIn: apiClient.isLoggedIn, requireAuth };
}
