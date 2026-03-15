"use client";

import { useRouter } from "next/navigation";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";

export function useAuthGuard() {
  const router = useRouter();
  const { addToast } = useToast();

  const requireAuth = (action?: string): boolean => {
    if (apiClient.isLoggedIn) return true;
    addToast("info", action ? `${action}은(는) 로그인이 필요합니다` : "로그인이 필요합니다");
    router.push("/login");
    return false;
  };

  return { isLoggedIn: apiClient.isLoggedIn, requireAuth };
}
