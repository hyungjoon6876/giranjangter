function SkeletonCard() {
  return (
    <div className="bg-card border border-border rounded-xl p-4 space-y-3">
      <div className="skeleton h-5 w-3/4 rounded" />
      <div className="skeleton h-4 w-1/2 rounded" />
      <div className="skeleton h-4 w-1/3 rounded" />
      <div className="flex justify-between items-center pt-2">
        <div className="skeleton h-6 w-20 rounded" />
        <div className="skeleton h-4 w-16 rounded" />
      </div>
    </div>
  );
}

interface ListingSkeletonProps {
  count?: number;
}

export function ListingSkeleton({ count = 6 }: ListingSkeletonProps) {
  return (
    <div role="status" aria-label="매물 목록을 불러오는 중" aria-busy="true">
      {/* Filter chip placeholders */}
      <div className="flex gap-2 mb-6">
        {Array.from({ length: 3 }).map((_, i) => (
          <div key={i} className="skeleton h-8 w-16 rounded-full" />
        ))}
      </div>

      {/* Skeleton cards grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {Array.from({ length: count }).map((_, i) => (
          <SkeletonCard key={i} />
        ))}
      </div>

    </div>
  );
}
