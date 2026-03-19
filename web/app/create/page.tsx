"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useCreateListing } from "@/lib/hooks/use-listings";
import { useToast } from "@/lib/hooks/use-toast";
import { useAuthGuard } from "@/lib/hooks/use-auth-guard";
import { ImageUpload } from "@/components/forms/image-upload";
import { ItemAutocomplete } from "@/components/forms/item-autocomplete";
import type { UploadedImage } from "@/lib/types";

export default function CreateListingPage() {
  const router = useRouter();
  const { isLoggedIn, requireAuth } = useAuthGuard();
  const createListing = useCreateListing();
  const { addToast } = useToast();

  const { data: servers = [] } = useQuery({
    queryKey: ["servers"],
    queryFn: () => apiClient.getServers(),
    enabled: isLoggedIn,
  });
  const { data: categories = [] } = useQuery({
    queryKey: ["categories"],
    queryFn: () => apiClient.getCategories(),
    enabled: isLoggedIn,
  });

  const [images, setImages] = useState<UploadedImage[]>([]);
  const [form, setForm] = useState({
    listingType: "sell",
    serverId: "",
    categoryId: "",
    itemName: "",
    title: "",
    description: "",
    priceType: "fixed",
    priceAmount: "",
    enhancementLevel: "",
    tradeMethod: "either",
  });

  useEffect(() => {
    if (!isLoggedIn) {
      requireAuth("매물 등록");
    }
  }, [isLoggedIn, requireAuth]);

  if (!isLoggedIn) return null;

  const update = (field: string, value: string) =>
    setForm((f) => ({ ...f, [field]: value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const data: Record<string, unknown> = {
      ...form,
      quantity: 1,
      priceAmount:
        form.priceType !== "offer" && form.priceAmount
          ? Number(form.priceAmount)
          : undefined,
      enhancementLevel: form.enhancementLevel
        ? Number(form.enhancementLevel)
        : undefined,
      images: images.map((img, i) => ({
        imageId: img.imageId,
        order: i,
      })),
    };
    try {
      await createListing.mutateAsync(data);
      router.push("/");
    } catch {
      addToast("error", "등록에 실패했습니다");
    }
  };

  const inputClass =
    "w-full bg-card border border-border rounded-lg px-4 py-3 text-sm text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim";
  const labelClass = "block text-sm font-medium text-text-primary mb-1";
  const sectionClass =
    "text-gold font-semibold text-sm tracking-wide uppercase mt-6 mb-3";

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-6 text-text-primary">매물 등록</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Type toggle */}
        <div className="flex rounded-lg border border-border overflow-hidden">
          {["sell", "buy"].map((t) => (
            <button
              key={t}
              type="button"
              onClick={() => update("listingType", t)}
              className={`flex-1 py-2.5 text-sm font-medium transition-colors ${
                form.listingType === t
                  ? "btn-gold-gradient text-white"
                  : "bg-card text-text-secondary"
              }`}
            >
              {t === "sell" ? "판매" : "구매"}
            </button>
          ))}
        </div>

        {/* Section: 기본 정보 */}
        <h3 className={sectionClass}>기본 정보</h3>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <div>
            <label htmlFor="serverId" className={labelClass}>
              서버 *
            </label>
            <select
              id="serverId"
              className={inputClass}
              value={form.serverId}
              onChange={(e) => update("serverId", e.target.value)}
              required
              aria-required="true"
            >
              <option value="">선택</option>
              {servers.map((s) => (
                <option key={s.serverId} value={s.serverId}>
                  {s.serverName}
                </option>
              ))}
            </select>
          </div>

          <div>
            <label htmlFor="categoryId" className={labelClass}>
              카테고리 *
            </label>
            <select
              id="categoryId"
              className={inputClass}
              value={form.categoryId}
              onChange={(e) => update("categoryId", e.target.value)}
              required
              aria-required="true"
            >
              <option value="">선택</option>
              {categories
                .filter((c) => !c.parentId)
                .map((c) => (
                  <option key={c.categoryId} value={c.categoryId}>
                    {c.categoryName}
                  </option>
                ))}
            </select>
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <div>
            <label htmlFor="itemName" className={labelClass}>
              아이템명 *
            </label>
            <ItemAutocomplete
              value={form.itemName}
              categoryId={form.categoryId || undefined}
              onChange={(v) => update("itemName", v)}
              required
              className={inputClass}
            />
          </div>
          <div>
            <label htmlFor="enhancementLevel" className={labelClass}>
              강화 수치 (선택)
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

        {/* Section: 상세 정보 */}
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
            placeholder="예: 집행검 +9 급처합니다"
            required
            aria-required="true"
            minLength={2}
          />
        </div>

        <div>
          <label htmlFor="description" className={labelClass}>
            설명 *
          </label>
          <textarea
            id="description"
            className={`${inputClass} h-28`}
            value={form.description}
            onChange={(e) => update("description", e.target.value)}
            placeholder="아이템 상세 설명"
            required
            aria-required="true"
            minLength={10}
          />
        </div>

        {/* Section: 가격 */}
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
              aria-required="true"
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

        {/* Section: 이미지 */}
        <h3 className={sectionClass}>이미지</h3>
        <ImageUpload images={images} onChange={setImages} maxImages={5} />

        {/* Section: 거래 */}
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
            aria-required="true"
          >
            <option value="in_game">인게임</option>
            <option value="offline_pc_bang">PC방/오프라인</option>
            <option value="either">무관</option>
          </select>
        </div>

        <button
          type="submit"
          disabled={createListing.isPending}
          className="w-full btn-gold-gradient text-white py-3 rounded-lg font-semibold transition-colors disabled:opacity-50"
        >
          {createListing.isPending ? "등록 중..." : "등록하기"}
        </button>
      </form>
    </div>
  );
}
