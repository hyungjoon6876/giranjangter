import Link from "next/link";

interface EmptyStateProps {
  icon?: string;
  title: string;
  description?: string;
  actionLabel?: string;
  actionHref?: string;
}

export function EmptyState({ icon = "\u{1F50D}", title, description, actionLabel, actionHref }: EmptyStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-20 text-text-secondary">
      <div aria-hidden="true" className="text-5xl mb-4">{icon}</div>
      <h2 className="text-lg">{title}</h2>
      {description && <p className="text-sm mt-2 text-text-dim">{description}</p>}
      {actionLabel && actionHref && (
        <Link href={actionHref} aria-label={actionLabel} className="mt-4 btn-gold-gradient text-white px-4 py-2 rounded-lg text-sm font-medium">
          {actionLabel}
        </Link>
      )}
    </div>
  );
}
