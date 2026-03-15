"use client";

import { Sidebar } from "./sidebar";
import { BottomNav } from "./bottom-nav";
import { Header } from "./header";
import type { ReactNode } from "react";

export function ResponsiveShell({ children }: { children: ReactNode }) {
  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 pb-16 lg:pb-0">{children}</main>
        <BottomNav />
      </div>
    </div>
  );
}
