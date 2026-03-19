"use client";

import { useState, useEffect, useMemo } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useMe, useUpdateProfile } from "@/lib/hooks/use-profile";
import { useAuthGuard } from "@/lib/hooks/use-auth-guard";
import { useToast } from "@/lib/hooks/use-toast";
import { ImageUpload } from "@/components/forms/image-upload";
import { Loading } from "@/components/ui/loading";
import type { UploadedImage } from "@/lib/types";

export default function ProfileEditPage() {
  const router = useRouter();
  const { isLoggedIn, requireAuth } = useAuthGuard();
  const { data: me, isLoading } = useMe();
  const updateProfile = useUpdateProfile();
  const { addToast } = useToast();

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
  });

  // Initialize form with user data
  const initialForm = useMemo(() => {
    if (!me) {
      return {
        nickname: "",
        introduction: "",
        primaryServerId: "",
      };
    }
    return {
      nickname: me.nickname,
      introduction: me.introduction ?? "",
      primaryServerId: me.primaryServerId ?? "",
    };
  }, [me]);

  const [form, setForm] = useState(initialForm);
  const [avatarImages, setAvatarImages] = useState<UploadedImage[]>([]);

  // Update form when user data changes
  useEffect(() => {
    setForm(initialForm);
  }, [initialForm]);

  useEffect(() => {
    if (!isLoggedIn) {
      requireAuth("프로필 수정");
    }
  }, [isLoggedIn, requireAuth]);

  if (!isLoggedIn) return null;
  if (isLoading) return <Loading />;
  if (!me) return null;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await updateProfile.mutateAsync({
        nickname: form.nickname,
        introduction: form.introduction || undefined,
        primaryServerId: form.primaryServerId || undefined,
        avatarUrl:
          avatarImages.length > 0 ? avatarImages[0].url : undefined,
      });
      addToast("success", "프로필이 수정되었습니다");
      router.push("/profile");
    } catch {
      addToast("error", "프로필 수정에 실패했습니다");
    }
  };

  const inputClass =
    "w-full bg-card border border-border rounded-lg px-4 py-3 text-sm text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim";
  const labelClass = "block text-sm font-medium text-text-primary mb-1";

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-6 text-text-primary">프로필 수정</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Avatar */}
        <div>
          <label className={labelClass}>프로필 사진</label>
          <div className="flex items-center gap-4">
            <div className="w-16 h-16 rounded-full bg-medium flex items-center justify-center text-2xl font-bold text-gold border-2 border-gold/30 overflow-hidden">
              {avatarImages.length > 0 ? (
                <Image
                  src={avatarImages[0].thumbnailUrl || avatarImages[0].url}
                  alt="프로필"
                  width={64}
                  height={64}
                  unoptimized
                  className="w-full h-full object-cover"
                />
              ) : me.avatarUrl ? (
                <Image
                  src={me.avatarUrl}
                  alt="프로필"
                  width={64}
                  height={64}
                  unoptimized
                  className="w-full h-full object-cover"
                />
              ) : (
                me.nickname[0]
              )}
            </div>
            <ImageUpload
              images={avatarImages}
              onChange={setAvatarImages}
              maxImages={1}
            />
          </div>
        </div>

        <div>
          <label htmlFor="nickname" className={labelClass}>
            닉네임 *
          </label>
          <input
            id="nickname"
            className={inputClass}
            value={form.nickname}
            onChange={(e) =>
              setForm((f) => ({ ...f, nickname: e.target.value }))
            }
            required
            minLength={2}
            maxLength={20}
          />
        </div>

        <div>
          <label htmlFor="introduction" className={labelClass}>
            소개
          </label>
          <textarea
            id="introduction"
            className={`${inputClass} h-20`}
            value={form.introduction}
            onChange={(e) =>
              setForm((f) => ({ ...f, introduction: e.target.value }))
            }
            placeholder="한 줄 소개를 입력하세요"
            maxLength={100}
          />
        </div>

        <div>
          <label htmlFor="primaryServerId" className={labelClass}>
            주 서버
          </label>
          <select
            id="primaryServerId"
            className={inputClass}
            value={form.primaryServerId}
            onChange={(e) =>
              setForm((f) => ({ ...f, primaryServerId: e.target.value }))
            }
          >
            <option value="">선택 안 함</option>
            {servers.map((s) => (
              <option key={s.serverId} value={s.serverId}>
                {s.serverName}
              </option>
            ))}
          </select>
        </div>

        <div className="flex gap-3 pt-2">
          <button
            type="button"
            onClick={() => router.back()}
            className="flex-1 py-3 border border-border rounded-lg text-text-secondary hover:bg-medium transition-colors"
          >
            취소
          </button>
          <button
            type="submit"
            disabled={updateProfile.isPending}
            className="flex-1 btn-gold-gradient text-white py-3 rounded-lg font-semibold disabled:opacity-50"
          >
            {updateProfile.isPending ? "저장 중..." : "저장"}
          </button>
        </div>
      </form>
    </div>
  );
}
