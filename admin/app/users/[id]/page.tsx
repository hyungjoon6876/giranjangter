"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { useUser, useUserModerationHistory, useRestrictUser } from "@/lib/hooks/use-users";
import { statusBadgeVariant, formatDate } from "@/lib/utils";

export default function UserDetailPage() {
  const params = useParams();
  const userId = params.id as string;

  const { data: user, isLoading } = useUser(userId);
  const { data: history } = useUserModerationHistory(userId);
  const restrictUser = useRestrictUser();

  const [showActionDialog, setShowActionDialog] = useState(false);
  const [restrictionScope, setRestrictionScope] = useState("chat");
  const [durationDays, setDurationDays] = useState(7);
  const [memo, setMemo] = useState("");
  const [reasonCode, setReasonCode] = useState("inappropriate_behavior");

  if (isLoading) {
    return (
      <div className="flex h-40 items-center justify-center text-text-secondary">
        로딩 중...
      </div>
    );
  }

  if (!user) {
    return (
      <div className="flex h-40 items-center justify-center text-text-secondary">
        사용자를 찾을 수 없습니다.
      </div>
    );
  }

  function handleRestrict() {
    restrictUser.mutate(
      {
        userId,
        restrictionScope,
        durationDays,
        reasonCode,
        memo: memo || undefined,
      },
      {
        onSuccess: () => {
          setShowActionDialog(false);
          setMemo("");
        },
      },
    );
  }

  return (
    <div className="space-y-6">
      {/* Profile Card */}
      <div className="rounded-xl border border-border bg-white p-6 shadow-sm">
        <div className="flex items-start justify-between">
          <div className="flex items-center gap-4">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-primary-100 text-xl font-bold text-primary-700">
              {user.nickname?.charAt(0)?.toUpperCase() ?? "?"}
            </div>
            <div>
              <h2 className="text-lg font-semibold text-text-primary">
                {user.nickname}
              </h2>
              <div className="mt-1 flex items-center gap-2">
                <span className="text-xs font-medium uppercase text-text-secondary">
                  {user.role}
                </span>
                <span
                  className={`inline-block rounded-full px-2 py-0.5 text-xs font-medium ${statusBadgeVariant(user.accountStatus)}`}
                >
                  {user.accountStatus}
                </span>
              </div>
            </div>
          </div>
          {/* Action buttons */}
          <div className="flex gap-2">
            <button
              onClick={() => {
                setRestrictionScope("chat");
                setShowActionDialog(true);
              }}
              className="rounded-lg border border-warning bg-white px-4 py-2 text-sm font-medium text-warning hover:bg-yellow-50"
            >
              제한
            </button>
            <button
              onClick={() => {
                setRestrictionScope("account");
                setShowActionDialog(true);
              }}
              className="rounded-lg bg-danger px-4 py-2 text-sm font-medium text-white hover:bg-red-700"
            >
              정지
            </button>
          </div>
        </div>

        {/* Stats row */}
        <div className="mt-6 grid grid-cols-2 gap-4 border-t border-border pt-6 sm:grid-cols-4">
          <div>
            <p className="text-xs text-text-secondary">성향점수</p>
            <p className="mt-1 text-lg font-semibold">
              {user.alignmentScore.toFixed(1)}{" "}
              <span className="text-sm font-normal text-text-secondary">
                ({user.alignmentGrade})
              </span>
            </p>
          </div>
          <div>
            <p className="text-xs text-text-secondary">완료 거래</p>
            <p className="mt-1 text-lg font-semibold">
              {user.completedTradeCount}
            </p>
          </div>
          <div>
            <p className="text-xs text-text-secondary">긍정 리뷰</p>
            <p className="mt-1 text-lg font-semibold">
              {user.positiveReviewCount ?? 0}
            </p>
          </div>
          <div>
            <p className="text-xs text-text-secondary">신뢰 뱃지</p>
            <p className="mt-1 text-lg font-semibold">
              {user.trustBadge ?? "없음"}
            </p>
          </div>
        </div>

        {/* Additional info */}
        <div className="mt-4 grid grid-cols-2 gap-4 border-t border-border pt-4 sm:grid-cols-3">
          <div>
            <p className="text-xs text-text-secondary">가입일</p>
            <p className="mt-1 text-sm">{formatDate(user.createdAt)}</p>
          </div>
          {user.lastLoginAt && (
            <div>
              <p className="text-xs text-text-secondary">마지막 로그인</p>
              <p className="mt-1 text-sm">{formatDate(user.lastLoginAt)}</p>
            </div>
          )}
          {user.introduction && (
            <div>
              <p className="text-xs text-text-secondary">소개</p>
              <p className="mt-1 text-sm">{user.introduction}</p>
            </div>
          )}
        </div>
      </div>

      {/* Action Dialog */}
      {showActionDialog && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/30">
          <div className="w-full max-w-md rounded-xl border border-border bg-white p-6 shadow-lg">
            <h3 className="text-base font-semibold text-text-primary">
              {restrictionScope === "account" ? "계정 정지" : "이용 제한"}
            </h3>
            <div className="mt-4 space-y-3">
              <div>
                <label className="text-xs text-text-secondary">제한 범위</label>
                <select
                  value={restrictionScope}
                  onChange={(e) => setRestrictionScope(e.target.value)}
                  className="mt-1 w-full rounded-lg border border-border px-3 py-2 text-sm focus:border-primary focus:outline-none"
                >
                  <option value="chat">채팅</option>
                  <option value="listing">매물 등록</option>
                  <option value="trade">거래</option>
                  <option value="account">계정 전체</option>
                </select>
              </div>
              <div>
                <label className="text-xs text-text-secondary">기간 (일)</label>
                <input
                  type="number"
                  value={durationDays}
                  onChange={(e) => setDurationDays(Number(e.target.value))}
                  min={1}
                  max={365}
                  className="mt-1 w-full rounded-lg border border-border px-3 py-2 text-sm focus:border-primary focus:outline-none"
                />
              </div>
              <div>
                <label className="text-xs text-text-secondary">사유 코드</label>
                <select
                  value={reasonCode}
                  onChange={(e) => setReasonCode(e.target.value)}
                  className="mt-1 w-full rounded-lg border border-border px-3 py-2 text-sm focus:border-primary focus:outline-none"
                >
                  <option value="inappropriate_behavior">부적절한 행동</option>
                  <option value="scam_attempt">사기 시도</option>
                  <option value="spam">스팸</option>
                  <option value="harassment">괴롭힘</option>
                  <option value="other">기타</option>
                </select>
              </div>
              <div>
                <label className="text-xs text-text-secondary">메모</label>
                <textarea
                  value={memo}
                  onChange={(e) => setMemo(e.target.value)}
                  placeholder="조치 사유를 입력하세요"
                  className="mt-1 w-full rounded-lg border border-border px-3 py-2 text-sm focus:border-primary focus:outline-none"
                  rows={3}
                />
              </div>
            </div>
            <div className="mt-5 flex justify-end gap-2">
              <button
                onClick={() => setShowActionDialog(false)}
                className="rounded-lg border border-border px-4 py-2 text-sm font-medium text-text-secondary hover:bg-slate-50"
              >
                취소
              </button>
              <button
                onClick={handleRestrict}
                disabled={restrictUser.isPending}
                className="rounded-lg bg-danger px-4 py-2 text-sm font-medium text-white hover:bg-red-700 disabled:opacity-50"
              >
                {restrictUser.isPending ? "처리 중..." : "확인"}
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Moderation History */}
      <div className="rounded-xl border border-border bg-white shadow-sm">
        <div className="border-b border-border px-5 py-4">
          <h3 className="text-base font-semibold text-text-primary">
            제재 이력
          </h3>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-border text-left text-text-secondary">
                <th className="px-5 py-3 font-medium">조치</th>
                <th className="px-5 py-3 font-medium">범위</th>
                <th className="px-5 py-3 font-medium">기간</th>
                <th className="px-5 py-3 font-medium">메모</th>
                <th className="px-5 py-3 font-medium">처리자</th>
                <th className="px-5 py-3 font-medium">날짜</th>
              </tr>
            </thead>
            <tbody>
              {!history || history.length === 0 ? (
                <tr>
                  <td
                    colSpan={6}
                    className="px-5 py-8 text-center text-text-secondary"
                  >
                    제재 이력이 없습니다.
                  </td>
                </tr>
              ) : (
                history.map((action) => (
                  <tr
                    key={action.actionId}
                    className="border-b border-border last:border-b-0 hover:bg-slate-50"
                  >
                    <td className="px-5 py-3">{action.actionCode}</td>
                    <td className="px-5 py-3">
                      {action.restrictionScope ?? "-"}
                    </td>
                    <td className="px-5 py-3">
                      {action.durationDays ? `${action.durationDays}일` : "-"}
                    </td>
                    <td className="max-w-xs truncate px-5 py-3">
                      {action.memo ?? "-"}
                    </td>
                    <td className="px-5 py-3 font-mono text-xs">
                      {action.actorUserId.slice(0, 8)}
                    </td>
                    <td className="px-5 py-3">
                      {formatDate(action.createdAt)}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
