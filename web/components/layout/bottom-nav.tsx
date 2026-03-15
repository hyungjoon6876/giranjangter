"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const TABS = [
  { href: "/", icon: "\u{1F3E0}", label: "\uB9E4\uBB3C" },
  { href: "/chats", icon: "\u{1F4AC}", label: "\uCC44\uD305" },
  { href: "/profile", icon: "\u{1F464}", label: "\uD504\uB85C\uD544" },
];

export function BottomNav() {
  const pathname = usePathname();

  return (
    <nav className="lg:hidden fixed bottom-0 left-0 right-0 bg-white border-t border-border flex z-50">
      {TABS.map((tab) => {
        const isActive = tab.href === "/"
          ? pathname === "/"
          : pathname.startsWith(tab.href);
        return (
          <Link
            key={tab.href}
            href={tab.href}
            className={`flex-1 flex flex-col items-center py-2 text-xs ${
              isActive ? "text-primary" : "text-text-secondary"
            }`}
          >
            <span className="text-lg">{tab.icon}</span>
            {tab.label}
          </Link>
        );
      })}
    </nav>
  );
}
