"use client";

import { Header } from "./header";
import { BottomNav } from "./bottom-nav";
import type { ReactNode } from "react";

export function ResponsiveShell({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen">
      <Header />
      <main className="pb-16 lg:pb-0">{children}</main>
      <BottomNav />
    </div>
  );
}
