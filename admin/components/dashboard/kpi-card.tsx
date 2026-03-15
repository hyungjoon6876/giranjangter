interface KpiCardProps {
  title: string;
  value: number;
  color: "primary" | "success" | "danger" | "warning";
  description?: string;
}

const BORDER_COLORS: Record<KpiCardProps["color"], string> = {
  primary: "border-l-[#6366f1]",
  success: "border-l-[#10b981]",
  danger: "border-l-[#ef4444]",
  warning: "border-l-[#f59e0b]",
};

const VALUE_COLORS: Record<KpiCardProps["color"], string> = {
  primary: "text-[#6366f1]",
  success: "text-[#10b981]",
  danger: "text-[#ef4444]",
  warning: "text-[#f59e0b]",
};

export function KpiCard({ title, value, color, description }: KpiCardProps) {
  return (
    <div
      className={`rounded-xl border border-border bg-white shadow-sm border-l-4 ${BORDER_COLORS[color]} p-5`}
    >
      <p className="text-sm font-medium text-text-secondary">{title}</p>
      <p className={`mt-2 text-3xl font-bold ${VALUE_COLORS[color]}`}>
        {value.toLocaleString("ko-KR")}
      </p>
      {description && (
        <p className="mt-1 text-xs text-text-secondary">{description}</p>
      )}
    </div>
  );
}
