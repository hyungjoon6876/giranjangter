"use client";

import { Header } from "./header";
import { BottomNav } from "./bottom-nav";
import { Footer } from "./footer";
import type { ReactNode } from "react";

export function ResponsiveShell({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen">
      <a
        href="#main-content"
        className="sr-only focus:not-sr-only focus:fixed focus:top-2 focus:left-2 focus:z-[70] focus:bg-gold focus:text-darkest focus:px-4 focus:py-2 focus:rounded-lg focus:font-bold focus:text-sm"
      >
        본문으로 건너뛰기
      </a>
      <Header />
      <main id="main-content" className="pb-16 lg:pb-0">
        {children}
      </main>
      <Footer />
      <BottomNav />
    </div>
  );
}
