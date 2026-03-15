"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";

const NAV_LINKS = [
  { href: "/", label: "마켓" },
  { href: "/chats", label: "채팅" },
  { href: "/create", label: "매물 등록" },
];

export function Header() {
  const pathname = usePathname();

  return (
    <header className="sticky top-0 z-50 h-16 bg-dark/95 backdrop-blur border-b border-border flex items-center px-4 lg:px-6">
      {/* Logo */}
      <Link href="/" className="font-display text-xl text-gold mr-8 flex-shrink-0">
        기란장터
      </Link>

      {/* Desktop nav links */}
      <nav className="hidden lg:flex items-center gap-1">
        {NAV_LINKS.map((link) => {
          const isActive = link.href === "/" ? pathname === "/" : pathname.startsWith(link.href);
          return (
            <Link
              key={link.href}
              href={link.href}
              className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors ${
                isActive ? "text-gold bg-[rgba(196,163,90,0.1)]" : "text-text-secondary hover:text-gold hover:bg-medium"
              }`}
            >
              {link.label}
            </Link>
          );
        })}
      </nav>

      {/* Spacer */}
      <div className="flex-1" />

      {/* Right: Search (desktop) + Notifications + Profile */}
      <div className="flex items-center gap-3">
        {/* Search - desktop only */}
        <div className="hidden lg:block">
          <input
            type="text"
            placeholder="아이템 검색..."
            className="bg-card border border-border rounded-lg px-3 py-1.5 text-sm text-text-primary placeholder:text-text-dim w-48 focus:border-gold focus:outline-none focus:ring-1 focus:ring-gold/30"
          />
        </div>
        {/* Notifications */}
        <Link href="/notifications" className="relative text-text-secondary hover:text-gold transition-colors p-2">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" />
            <path d="M13.73 21a2 2 0 0 1-3.46 0" />
          </svg>
        </Link>
        {/* Profile */}
        <Link href="/profile" className="text-text-secondary hover:text-gold transition-colors p-2">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
        </Link>
      </div>
    </header>
  );
}
