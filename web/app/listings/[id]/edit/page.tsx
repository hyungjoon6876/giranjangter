"use client";

import { use, useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useListing, useUpdateListing } from "@/lib/hooks/use-listings";
import { useToast } from "@/lib/hooks/use-toast";
import { useAuthGuard } from "@/lib/hooks/use-auth-guard";
import { ImageUpload } from "@/components/forms/image-upload";
import { ItemAutocomplete } from "@/components/forms/item-autocomplete";
import { Loading } from "@/components/ui/loading";
import { ErrorState } from "@/components/ui/error-state";
import type { UploadedImage } from "@/lib/types";

export default function EditListingPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  const router = useRouter();
  const { isLoggedIn, requireAuth } = useAuthGuard();
  const { data: listing, isLoading, isError } = useListing(id);
  const updateListing = useUpdateListing();
  const { addToast } = useToast();

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
  });
  const { data: categories = [] } = useQuery({
    queryKey: ["categories"],
    queryFn: () => apiClient.getCategories(),
  });

  const [form, setForm] = useState({
    title: "",
    description: "",
    itemName: "",
    priceType: "fixed",
    priceAmount: "",
    enhancementLevel: "",
    optionsText: "",
    tradeMethod: "either",
    preferredMeetingAreaText: "",
    availableTimeText: "",
  });
  const [images, setImages] = useState<UploadedImage[]>([]);
  const [initialized, setInitialized] = useState(false);

  useEffect(() => {
    if (listing && !initialized) {
      setForm({
        title: listing.title,
        description: listing.description ?? "",
        itemName: listing.itemName,
        priceType: listing.priceType,
        priceAmount: listing.priceAmount?.toString() ?? "",
        enhancementLevel: listing.enhancementLevel?.toString() ?? "",
        optionsText: listing.optionsText ?? "",
        tradeMethod: listing.tradeMethod,
        preferredMeetingAreaText: listing.preferredMeetingAreaText ?? "",
        availableTimeText: listing.availableTimeText ?? "",
      });
      if (listing.images?.length) {
        setImages(
          listing.images.map((img) => ({
            imageId: img.imageId,
            url: img.url,
            thumbnailUrl: img.url,
          })),
        );
      }
      setInitialized(true);
    }
  }, [listing, initialized]);

  useEffect(() => {
    if (!isLoggedIn) {
      requireAuth("매물 수정");
    }
  }, [isLoggedIn]);

  if (!isLoggedIn) return null;
  if (isLoading) return <Loading />;
  if (isError)
    return (
      <ErrorState
        message="매물을 불러올 수 없습니다"
        onRetry={() => router.refresh()}
      />
    );
  if (!listing) return null;

  if (!listing.isOwner) {
    return (
      <ErrorState message="수정 권한이 없습니다" onRetry={() => router.back()} />
    );
  }

  const update = (field: string, value: string) =>
    setForm((f) => ({ ...f, [field]: value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await updateListing.mutateAsync({
        id,
        data: {
          title: form.title,
          description: form.description || undefined,
          itemName: form.itemName,
          priceType: form.priceType as "fixed" | "negotiable" | "offer",
          priceAmount:
            form.priceType !== "offer" && form.priceAmount
              ? Number(form.priceAmount)
              : undefined,
          enhancementLevel: form.enhancementLevel
            ? Number(form.enhancementLevel)
            : undefined,
          optionsText: form.optionsText || undefined,
          tradeMethod: form.tradeMethod,
          preferredMeetingAreaText:
            form.preferredMeetingAreaText || undefined,
          availableTimeText: form.availableTimeText || undefined,
          images: images.map((img, i) => ({
            imageId: img.imageId,
            url: img.url,
            order: i,
          })),
        },
      });
      addToast("success", "매물이 수정되었습니다");
      router.push(`/listings/${id}`);
    } catch {
      addToast("error", "수정에 실패했습니다");
    }
  };

  const inputClass =
    "w-full bg-card border border-border rounded-lg px-4 py-3 text-sm text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim";
  const labelClass = "block text-sm font-medium text-text-primary mb-1";
  const sectionClass =
    "text-gold font-semibold text-sm tracking-wide uppercase mt-6 mb-3";

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-6 text-text-primary">매물 수정</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="bg-medium rounded-lg p-3 text-sm text-text-secondary">
          <span className="font-medium">{listing.serverName}</span>
          <span className="mx-2">·</span>
          <span>{listing.categoryName}</span>
          <span className="mx-2">·</span>
          <span>{listing.listingType === "sell" ? "판매" : "구매"}</span>
        </div>

        <h3 className={sectionClass}>아이템</h3>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <div>
            <label htmlFor="itemName" className={labelClass}>
              아이템명 *
            </label>
            <ItemAutocomplete
              value={form.itemName}
              categoryId={listing.categoryId}
              onChange={(v) => update("itemName", v)}
              required
              className={inputClass}
            />
          </div>
          <div>
            <label htmlFor="enhancementLevel" className={labelClass}>
              강화 수치
            </label>
            <input
              id="enhancementLevel"
              className={inputClass}
              type="number"
              value={form.enhancementLevel}
              onChange={(e) => update("enhancementLevel", e.target.value)}
            />
          </div>
        </div>

        <div>
          <label htmlFor="optionsText" className={labelClass}>
            옵션 설명
          </label>
          <input
            id="optionsText"
            className={inputClass}
            value={form.optionsText}
            onChange={(e) => update("optionsText", e.target.value)}
            placeholder="예: 흑단 +2, 체력 +50"
          />
        </div>

        <h3 className={sectionClass}>상세 정보</h3>

        <div>
          <label htmlFor="title" className={labelClass}>
            제목 *
          </label>
          <input
            id="title"
            className={inputClass}
            value={form.title}
            onChange={(e) => update("title", e.target.value)}
            required
            aria-required="true"
            minLength={2}
          />
        </div>

        <div>
          <label htmlFor="description" className={labelClass}>
            설명
          </label>
          <textarea
            id="description"
            className={`${inputClass} h-28`}
            value={form.description}
            onChange={(e) => update("description", e.target.value)}
          />
        </div>

        <h3 className={sectionClass}>이미지</h3>
        <ImageUpload images={images} onChange={setImages} maxImages={5} />

        <h3 className={sectionClass}>가격</h3>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <div>
            <label htmlFor="priceType" className={labelClass}>
              가격 유형
            </label>
            <select
              id="priceType"
              className={inputClass}
              value={form.priceType}
              onChange={(e) => update("priceType", e.target.value)}
            >
              <option value="fixed">고정가</option>
              <option value="negotiable">협상가능</option>
              <option value="offer">제안받음</option>
            </select>
          </div>
          <div>
            <label htmlFor="priceAmount" className={labelClass}>
              가격 (원)
            </label>
            <input
              id="priceAmount"
              className={inputClass}
              type="number"
              value={form.priceAmount}
              onChange={(e) => update("priceAmount", e.target.value)}
              disabled={form.priceType === "offer"}
            />
          </div>
        </div>

        <h3 className={sectionClass}>거래</h3>

        <div>
          <label htmlFor="tradeMethod" className={labelClass}>
            거래 방식
          </label>
          <select
            id="tradeMethod"
            className={inputClass}
            value={form.tradeMethod}
            onChange={(e) => update("tradeMethod", e.target.value)}
          >
            <option value="in_game">인게임</option>
            <option value="offline_pc_bang">PC방/오프라인</option>
            <option value="either">무관</option>
          </select>
        </div>

        <div>
          <label htmlFor="preferredMeetingAreaText" className={labelClass}>
            접선 장소
          </label>
          <input
            id="preferredMeetingAreaText"
            className={inputClass}
            value={form.preferredMeetingAreaText}
            onChange={(e) =>
              update("preferredMeetingAreaText", e.target.value)
            }
            placeholder="예: 기란마을 2시 방향"
          />
        </div>

        <div>
          <label htmlFor="availableTimeText" className={labelClass}>
            거래 가능 시간
          </label>
          <input
            id="availableTimeText"
            className={inputClass}
            value={form.availableTimeText}
            onChange={(e) => update("availableTimeText", e.target.value)}
            placeholder="예: 평일 저녁 8-11시"
          />
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
            disabled={updateListing.isPending}
            className="flex-1 btn-gold-gradient text-white py-3 rounded-lg font-semibold disabled:opacity-50"
          >
            {updateListing.isPending ? "수정 중..." : "수정하기"}
          </button>
        </div>
      </form>
    </div>
  );
}
