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
        <Link href="/login" className="text-primary font-medium">로그인하기</Link>
      </div>
    );
  }

  const handleLogout = () => {
    apiClient.clearTokens();
    router.push("/");
    router.refresh();
  };

  const menuItems = [
    { href: "/profile/listings", icon: "📦", label: "내 매물" },
    { href: "/profile/trades", icon: "🤝", label: "내 거래" },
    { href: "/notifications", icon: "🔔", label: "알림" },
  ];

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      {/* User card */}
      <div className="bg-white rounded-xl border border-border p-5 mb-4">
        <div className="flex items-center gap-4">
          <div className="w-16 h-16 rounded-full bg-surface flex items-center justify-center text-2xl font-bold text-text-secondary">
            {me.nickname[0]}
          </div>
          <div>
            <h2 className="text-xl font-bold">{me.nickname}</h2>
            {me.introduction && <p className="text-sm text-text-secondary mt-1">{me.introduction}</p>}
          </div>
        </div>
        <div className="grid grid-cols-3 gap-4 mt-5 text-center">
          <div><div className="text-xl font-bold">{me.completedTradeCount}</div><div className="text-xs text-text-secondary">거래</div></div>
          <div><div className="text-xl font-bold">{me.positiveReviewCount}</div><div className="text-xs text-text-secondary">좋은 리뷰</div></div>
          <div><div className="text-xl font-bold">{me.trustBadge ?? "-"}</div><div className="text-xs text-text-secondary">신뢰등급</div></div>
        </div>
      </div>

      {/* Menu */}
      <div className="bg-white rounded-xl border border-border overflow-hidden">
        {menuItems.map((item) => (
          <Link key={item.href} href={item.href} className="flex items-center gap-3 px-5 py-4 border-b border-border last:border-0 hover:bg-surface transition-colors">
            <span>{item.icon}</span>
            <span className="flex-1">{item.label}</span>
            <span className="text-text-secondary">›</span>
          </Link>
        ))}
      </div>

      <button onClick={handleLogout} className="w-full mt-4 py-3 text-sm text-error border border-border rounded-xl hover:bg-red-50 transition-colors">
        로그아웃
      </button>
    </div>
  );
}
