"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useChats } from "@/lib/hooks/use-chats";
import { useIsLoggedIn } from "@/lib/hooks/use-auth";
import type { ReactNode } from "react";

interface Tab {
  href: string;
  label: string;
  ariaLabel: string;
  icon: ReactNode;
  badgeKey?: "chat";
}

const TABS: Tab[] = [
  { href: "/", label: "마켓", ariaLabel: "마켓", icon: (
    <svg aria-hidden="true" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/></svg>
  )},
  { href: "/chats", label: "채팅", ariaLabel: "채팅", badgeKey: "chat", icon: (
    <svg aria-hidden="true" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
  )},
  { href: "/create", label: "등록", ariaLabel: "매물 등록", icon: (
    <svg aria-hidden="true" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
  )},
  { href: "/profile", label: "프로필", ariaLabel: "프로필", icon: (
    <svg aria-hidden="true" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
  )},
];

export function BottomNav() {
  const pathname = usePathname();
  const isLoggedIn = useIsLoggedIn();
  const { data: chatsData } = useChats();

  const chatUnreadCount = isLoggedIn
    ? (chatsData?.data ?? []).reduce((sum, c) => sum + c.unreadCount, 0)
    : 0;

  function getBadgeCount(badgeKey?: string): number {
    if (badgeKey === "chat") return chatUnreadCount;
    return 0;
  }

  return (
    <nav aria-label="하단 메뉴" className="lg:hidden fixed bottom-0 left-0 right-0 bg-dark/95 backdrop-blur border-t border-border flex z-50">
      {TABS.map((tab) => {
        const isActive = tab.href === "/" ? pathname === "/" : pathname.startsWith(tab.href);
        const badgeCount = getBadgeCount(tab.badgeKey);
        const ariaLabel = badgeCount > 0
          ? `${tab.ariaLabel} — 읽지 않은 메시지 ${badgeCount}개`
          : tab.ariaLabel;
        return (
          <Link
            key={tab.href}
            href={tab.href}
            aria-label={ariaLabel}
            aria-current={isActive ? "page" : undefined}
            className={`flex-1 flex flex-col items-center py-2 text-xs transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold focus-visible:rounded ${isActive ? "text-gold" : "text-text-dim"}`}
          >
            <span className="relative">
              {tab.icon}
              {badgeCount > 0 && (
                <span
                  aria-hidden="true"
                  className="absolute -top-1.5 -right-2.5 min-w-[18px] h-[18px] flex items-center justify-center bg-gold text-darkest text-[10px] font-bold rounded-full px-1"
                >
                  {badgeCount > 99 ? "99+" : badgeCount}
                </span>
              )}
            </span>
            <span className="mt-0.5">{tab.label}</span>
          </Link>
        );
      })}
    </nav>
  );
}
