"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { useMe } from "@/lib/hooks/use-profile";
import { apiClient } from "@/lib/api-client";
import { Loading } from "@/components/ui/loading";

export default function ProfilePage() {
  const router = useRouter();
  const { data: me, isLoading } = useMe();

  if (isLoading) return <Loading />;
  if (!me) {
    return (
      <div className="p-6 text-center">
        <p className="text-text-secondary mb-4">로그인이 필요합니다</p>
        <Link href="/login" className="text-gold font-medium">로그인하기</Link>
      </div>
    );
  }

  const handleLogout = () => {
    apiClient.clearTokens();
    router.push("/");
    router.refresh();
  };

  const menuItems = [
    { href: "/profile/listings", label: "내 매물" },
    { href: "/profile/trades", label: "내 거래" },
    { href: "/notifications", label: "알림" },
  ];

  const isTrusted = me.trustBadge === "trusted";

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      {/* User card with gradient top */}
      <div className="bg-card rounded-xl border border-border overflow-hidden mb-4">
        <div className="bg-gradient-to-b from-card to-dark p-5">
          <div className="flex items-center gap-4">
            <div className="w-16 h-16 rounded-full bg-medium flex items-center justify-center text-2xl font-bold text-gold border-2 border-gold/30">
              {me.nickname[0]}
            </div>
            <div>
              <h2 className="text-xl font-display font-bold text-gold">{me.nickname}</h2>
              {me.introduction && <p className="text-sm text-text-secondary mt-1">{me.introduction}</p>}
            </div>
          </div>
        </div>
        {/* RPG stat cards */}
        <div className="grid grid-cols-3 gap-4 p-5">
          <div className="bg-medium rounded-lg p-3 text-center border border-border">
            <div className="text-gold font-display text-2xl">{me.completedTradeCount}</div>
            <div className="text-text-dim text-xs mt-1">거래</div>
          </div>
          <div className="bg-medium rounded-lg p-3 text-center border border-border">
            <div className="text-gold font-display text-2xl">{me.positiveReviewCount}</div>
            <div className="text-text-dim text-xs mt-1">좋은 리뷰</div>
          </div>
          <div className={`bg-medium rounded-lg p-3 text-center border ${isTrusted ? "border-gold/50" : "border-border"}`}>
            <div className={`font-display text-2xl ${isTrusted ? "text-gold" : "text-text-secondary"}`}>{me.trustBadge ?? "-"}</div>
            <div className="text-text-dim text-xs mt-1">신뢰등급</div>
          </div>
        </div>
      </div>

      {/* Menu */}
      <div className="bg-card rounded-xl border border-border overflow-hidden">
        {menuItems.map((item) => (
          <Link
            key={item.href}
            href={item.href}
            className="flex items-center gap-3 px-5 py-4 border-b border-border last:border-0 hover:bg-medium hover:text-gold hover:border-l-4 hover:border-l-gold transition-colors"
          >
            <span className="flex-1 text-text-primary">{item.label}</span>
            <span className="text-text-secondary">&rsaquo;</span>
          </Link>
        ))}
      </div>

      <button
        onClick={handleLogout}
        className="w-full mt-4 py-3 text-sm text-danger border border-danger/40 rounded-xl hover:bg-danger/10 transition-colors"
      >
        로그아웃
      </button>
    </div>
  );
}
