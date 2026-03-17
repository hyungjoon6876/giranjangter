"use client";

import { Modal } from "@/components/ui/modal";
import { useBlockUser } from "@/lib/hooks/use-users";
import { useToast } from "@/lib/hooks/use-toast";

interface BlockConfirmModalProps {
  open: boolean;
  onClose: () => void;
  userId: string;
  nickname: string;
}

export function BlockConfirmModal({
  open,
  onClose,
  userId,
  nickname,
}: BlockConfirmModalProps) {
  const blockUser = useBlockUser();
  const { addToast } = useToast();

  const handleBlock = async () => {
    try {
      await blockUser.mutateAsync(userId);
      addToast("success", `${nickname}님을 차단했습니다`);
      onClose();
    } catch {
      addToast("error", "차단에 실패했습니다");
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="사용자 차단">
      <p className="text-text-secondary text-sm mb-4">
        <span className="font-medium text-text-primary">{nickname}</span>님을
        차단하시겠습니까?
      </p>
      <p className="text-text-dim text-xs mb-6">
        차단하면 해당 사용자의 매물과 채팅이 더 이상 표시되지 않습니다.
        프로필 설정에서 차단을 해제할 수 있습니다.
      </p>
      <div className="flex gap-3">
        <button
          onClick={onClose}
          className="flex-1 py-2.5 border border-border rounded-lg text-text-secondary hover:bg-medium transition-colors"
        >
          취소
        </button>
        <button
          onClick={handleBlock}
          disabled={blockUser.isPending}
          className="flex-1 py-2.5 bg-danger text-white rounded-lg font-medium disabled:opacity-50"
        >
          {blockUser.isPending ? "처리 중..." : "차단"}
        </button>
      </div>
    </Modal>
  );
}
