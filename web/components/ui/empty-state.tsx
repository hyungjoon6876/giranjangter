interface EmptyStateProps {
  icon?: string;
  title: string;
  description?: string;
}

export function EmptyState({ icon = "\u{1F50D}", title, description }: EmptyStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-20 text-text-secondary">
      <span className="text-5xl mb-4">{icon}</span>
      <p className="text-lg">{title}</p>
      {description && <p className="text-sm mt-2">{description}</p>}
    </div>
  );
}
