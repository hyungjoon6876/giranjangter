"use client";

import { useNotifications, useMarkNotificationsRead } from "@/lib/hooks/use-profile";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { formatTimeAgo } from "@/lib/utils";

export default function NotificationsPage() {
  const { data, isLoading } = useNotifications();
  const markRead = useMarkNotificationsRead();

  const notifications = data?.data ?? [];

  if (isLoading) return <Loading />;
  if (!notifications.length) return <EmptyState title="알림이 없습니다" icon="🔔" />;

  const unreadIds = notifications.filter((n) => !n.readAt).map((n) => n.notificationId);

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold">알림</h1>
        {unreadIds.length > 0 && (
          <button onClick={() => markRead.mutate(unreadIds)} className="text-sm text-primary">모두 읽음</button>
        )}
      </div>
      <div className="bg-white rounded-xl border border-border overflow-hidden">
        {notifications.map((n) => (
          <div key={n.notificationId} className={`px-5 py-4 border-b border-border last:border-0 ${!n.readAt ? "bg-blue-50" : ""}`}>
            <p className="text-sm">{n.message}</p>
            <p className="text-xs text-text-secondary mt-1">{formatTimeAgo(n.createdAt)}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
