"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const NAV_ITEMS = [
  { href: "/", label: "매물" },
  { href: "/chats", label: "채팅" },
  { href: "/create", label: "매물 등록" },
  { href: "/profile", label: "프로필" },
  { href: "/notifications", label: "알림" },
];

export function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="hidden lg:flex flex-col w-52 bg-dark text-text-primary min-h-screen flex-shrink-0">
      <div className="px-5 py-4 border-b border-border">
        <img src="/logo.png" alt="기란JT" className="h-8" />
      </div>
      <nav className="flex-1 py-2">
        {NAV_ITEMS.map((item) => {
          const isActive = item.href === "/"
            ? pathname === "/"
            : pathname.startsWith(item.href);
          return (
            <Link
              key={item.href}
              href={item.href}
              className={`flex items-center gap-3 px-5 py-3 text-sm transition-colors ${
                isActive
                  ? "bg-medium text-gold border-l-2 border-l-gold"
                  : "text-text-secondary hover:bg-medium/50 hover:text-text-primary"
              }`}
            >
              {item.label}
            </Link>
          );
        })}
      </nav>
    </aside>
  );
}
