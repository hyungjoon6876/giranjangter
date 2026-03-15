"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

interface NavItem {
  label: string;
  href: string;
  icon: React.ReactNode;
}

const navItems: NavItem[] = [
  {
    label: "대시보드",
    href: "/",
    icon: (
      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
        <rect x="2" y="2" width="7" height="7" rx="1" />
        <rect x="11" y="2" width="7" height="7" rx="1" />
        <rect x="2" y="11" width="7" height="7" rx="1" />
        <rect x="11" y="11" width="7" height="7" rx="1" />
      </svg>
    ),
  },
  {
    label: "신고관리",
    href: "/reports",
    icon: (
      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
        <path d="M4 3v14" />
        <path d="M4 3l12 4-12 4" />
      </svg>
    ),
  },
  {
    label: "사용자",
    href: "/users",
    icon: (
      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
        <circle cx="7" cy="6" r="3" />
        <path d="M2 17c0-3 2.5-5 5-5s5 2 5 5" />
        <circle cx="14" cy="6" r="2.5" />
        <path d="M14 11c2 0 4 1.5 4 4" />
      </svg>
    ),
  },
  {
    label: "매물",
    href: "/listings",
    icon: (
      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
        <rect x="3" y="5" width="14" height="12" rx="1" />
        <path d="M7 5V3a1 1 0 011-1h4a1 1 0 011 1v2" />
        <path d="M3 9h14" />
      </svg>
    ),
  },
  {
    label: "거래",
    href: "/trades",
    icon: (
      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
        <path d="M14 3l-4 4h3v6h2V7h3l-4-4z" />
        <path d="M6 17l4-4H7V7H5v6H2l4 4z" />
      </svg>
    ),
  },
  {
    label: "감사로그",
    href: "/audit",
    icon: (
      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
        <rect x="4" y="2" width="12" height="16" rx="1" />
        <path d="M7 6h6" />
        <path d="M7 9h6" />
        <path d="M7 12h4" />
      </svg>
    ),
  },
];

export function AdminSidebar() {
  const pathname = usePathname();

  function isActive(href: string) {
    if (href === "/") return pathname === "/";
    return pathname.startsWith(href);
  }

  return (
    <aside className="fixed inset-y-0 left-0 z-30 flex w-56 flex-col bg-sidebar text-white">
      {/* Logo */}
      <div className="flex h-14 items-center px-5 border-b border-white/10">
        <span className="text-lg font-bold tracking-tight">기란장터 Admin</span>
      </div>

      {/* Navigation */}
      <nav className="flex-1 overflow-y-auto px-3 py-4">
        <ul className="space-y-1">
          {navItems.map((item) => (
            <li key={item.href}>
              <Link
                href={item.href}
                className={`flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors ${
                  isActive(item.href)
                    ? "bg-indigo-700/50 text-white"
                    : "text-indigo-200 hover:bg-indigo-700/30 hover:text-white"
                }`}
              >
                {item.icon}
                {item.label}
              </Link>
            </li>
          ))}
        </ul>
      </nav>

      {/* Bottom: Logout */}
      <div className="border-t border-white/10 px-3 py-4">
        <button
          onClick={() => {
            localStorage.removeItem("accessToken");
            localStorage.removeItem("refreshToken");
            window.location.href = "/login";
          }}
          className="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium text-indigo-300 transition-colors hover:bg-indigo-700/30 hover:text-white"
        >
          <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
            <path d="M7 17H4a1 1 0 01-1-1V4a1 1 0 011-1h3" />
            <path d="M14 14l4-4-4-4" />
            <path d="M18 10H8" />
          </svg>
          로그아웃
        </button>
      </div>
    </aside>
  );
}
