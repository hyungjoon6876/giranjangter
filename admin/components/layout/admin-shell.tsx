"use client";

import { usePathname } from "next/navigation";
import { AdminSidebar } from "./admin-sidebar";
import { AdminHeader } from "./admin-header";

const SHELL_EXCLUDED_PATHS = ["/login"];

export function AdminShell({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const isExcluded = SHELL_EXCLUDED_PATHS.some((p) => pathname.startsWith(p));

  if (isExcluded) {
    return <>{children}</>;
  }

  return (
    <div className="flex min-h-screen">
      <AdminSidebar />
      <div className="ml-56 flex flex-1 flex-col">
        <AdminHeader />
        <main className="flex-1 p-6">{children}</main>
      </div>
    </div>
  );
}
