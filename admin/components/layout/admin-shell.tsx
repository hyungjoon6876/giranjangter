import { AdminSidebar } from "./admin-sidebar";
import { AdminHeader } from "./admin-header";

export function AdminShell({ children }: { children: React.ReactNode }) {
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
