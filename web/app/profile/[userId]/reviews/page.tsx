"use client";

import { use } from "react";
import { useUserReviews } from "@/lib/hooks/use-reviews";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { formatTimeAgo } from "@/lib/utils";

export default function UserReviewsPage({
  params,
}: {
  params: Promise<{ userId: string }>;
}) {
  const { userId } = use(params);
  const { data: reviews = [], isLoading } = useUserReviews(userId);

  if (isLoading) return <Loading />;

  if (!reviews.length) {
    return (
      <EmptyState
        title="아직 리뷰가 없습니다"
        description="거래를 완료하면 리뷰를 받을 수 있습니다"
      />
    );
  }

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4 text-text-primary">
        받은 리뷰 ({reviews.length})
      </h1>
      <div className="space-y-3">
        {reviews.map((r) => (
          <div
            key={r.reviewId}
            className="bg-card rounded-xl border border-border p-4"
          >
            <div className="flex items-center justify-between mb-2">
              <div className="flex items-center gap-2">
                <span className="font-medium text-text-primary">
                  {r.reviewerNickname}
                </span>
                <span
                  className={`text-xs font-medium px-2 py-0.5 rounded-full ${
                    r.rating === "positive"
                      ? "bg-green-500/10 text-green-400"
                      : "bg-danger/10 text-danger"
                  }`}
                >
                  {r.rating === "positive" ? "👍 좋아요" : "👎 아쉬워요"}
                </span>
              </div>
              <span className="text-xs text-text-dim">
                {formatTimeAgo(r.createdAt)}
              </span>
            </div>
            {r.comment && (
              <p className="text-sm text-text-secondary">{r.comment}</p>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
