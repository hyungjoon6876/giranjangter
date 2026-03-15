"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const NAV_ITEMS = [
  { href: "/", icon: "\u{1F3E0}", label: "\uB9E4\uBB3C" },
  { href: "/chats", icon: "\u{1F4AC}", label: "\uCC44\uD305" },
  { href: "/create", icon: "\u{1F4DD}", label: "\uB9E4\uBB3C \uB4F1\uB85D" },
  { href: "/profile", icon: "\u{1F464}", label: "\uD504\uB85C\uD544" },
  { href: "/notifications", icon: "\u{1F514}", label: "\uC54C\uB9BC" },
];

export function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="hidden lg:flex flex-col w-52 bg-slate-800 text-white min-h-screen flex-shrink-0">
      <div className="px-5 py-4 text-lg font-bold border-b border-slate-700">
        기란장터
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
                  ? "bg-slate-700 text-white"
                  : "text-slate-400 hover:bg-slate-700/50 hover:text-white"
              }`}
            >
              <span>{item.icon}</span>
              {item.label}
            </Link>
          );
        })}
      </nav>
    </aside>
  );
}
