"use client";

import Link from "next/link";

export function Header() {
  return (
    <header className="lg:hidden flex items-center justify-between px-4 py-3 bg-dark border-b border-border">
      <Link href="/" className="text-lg font-bold font-display text-gold-gradient">
        기란장터
      </Link>
    </header>
  );
}
