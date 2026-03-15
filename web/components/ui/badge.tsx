interface BadgeProps {
  label: string;
  color: string; // hex color
}

export function Badge({ label, color }: BadgeProps) {
  return (
    <span
      className="inline-block px-2 py-0.5 text-xs font-semibold rounded"
      style={{ color, backgroundColor: `${color}1A` }}
    >
      {label}
    </span>
  );
}

export function TypeBadge({ type }: { type: "sell" | "buy" }) {
  return type === "sell"
    ? <Badge label="판매" color="#c4a35a" />
    : <Badge label="구매" color="#5b9bd5" />;
}
