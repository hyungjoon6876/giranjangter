"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import { useCreateListing } from "@/lib/hooks/use-listings";

export default function CreateListingPage() {
  const router = useRouter();
  const createListing = useCreateListing();
  const { data: servers = [] } = useQuery({ queryKey: ["servers"], queryFn: () => apiClient.getServers() });
  const { data: categories = [] } = useQuery({ queryKey: ["categories"], queryFn: () => apiClient.getCategories() });

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

  const update = (field: string, value: string) => setForm((f) => ({ ...f, [field]: value }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const data: Record<string, unknown> = {
      ...form,
      quantity: 1,
      priceAmount: form.priceType !== "offer" && form.priceAmount ? Number(form.priceAmount) : undefined,
      enhancementLevel: form.enhancementLevel ? Number(form.enhancementLevel) : undefined,
    };
    try {
      await createListing.mutateAsync(data);
      router.push("/");
    } catch (e) {
      alert(`등록 실패: ${JSON.stringify(e)}`);
    }
  };

  const inputClass = "w-full bg-card border border-border rounded-lg px-4 py-3 text-sm text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim";
  const labelClass = "block text-sm font-medium text-text-primary mb-1";
  const sectionClass = "text-gold font-semibold text-sm tracking-wide uppercase mt-6 mb-3";

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
                form.listingType === t ? "btn-gold-gradient text-white" : "bg-card text-text-secondary"
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
            <label className={labelClass}>서버 *</label>
            <select className={inputClass} value={form.serverId} onChange={(e) => update("serverId", e.target.value)} required>
              <option value="">선택</option>
              {servers.map((s) => <option key={s.serverId} value={s.serverId}>{s.serverName}</option>)}
            </select>
          </div>

          <div>
            <label className={labelClass}>카테고리 *</label>
            <select className={inputClass} value={form.categoryId} onChange={(e) => update("categoryId", e.target.value)} required>
              <option value="">선택</option>
              {categories.filter((c) => !c.parentId).map((c) => <option key={c.categoryId} value={c.categoryId}>{c.categoryName}</option>)}
            </select>
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <div>
            <label className={labelClass}>아이템명 *</label>
            <input className={inputClass} value={form.itemName} onChange={(e) => update("itemName", e.target.value)} required />
          </div>
          <div>
            <label className={labelClass}>강화 수치 (선택)</label>
            <input className={inputClass} type="number" value={form.enhancementLevel} onChange={(e) => update("enhancementLevel", e.target.value)} />
          </div>
        </div>

        {/* Section: 상세 정보 */}
        <h3 className={sectionClass}>상세 정보</h3>

        <div>
          <label className={labelClass}>제목 *</label>
          <input className={inputClass} value={form.title} onChange={(e) => update("title", e.target.value)} placeholder="예: 집행검 +9 급처합니다" required minLength={2} />
        </div>

        <div>
          <label className={labelClass}>설명 *</label>
          <textarea className={`${inputClass} h-28`} value={form.description} onChange={(e) => update("description", e.target.value)} placeholder="아이템 상세 설명" required minLength={10} />
        </div>

        {/* Section: 가격 */}
        <h3 className={sectionClass}>가격</h3>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <div>
            <label className={labelClass}>가격 유형</label>
            <select className={inputClass} value={form.priceType} onChange={(e) => update("priceType", e.target.value)}>
              <option value="fixed">고정가</option>
              <option value="negotiable">협상가능</option>
              <option value="offer">제안받음</option>
            </select>
          </div>
          <div>
            <label className={labelClass}>가격 (원)</label>
            <input className={inputClass} type="number" value={form.priceAmount} onChange={(e) => update("priceAmount", e.target.value)} disabled={form.priceType === "offer"} />
          </div>
        </div>

        {/* Section: 거래 */}
        <h3 className={sectionClass}>거래</h3>

        <div>
          <label className={labelClass}>거래 방식</label>
          <select className={inputClass} value={form.tradeMethod} onChange={(e) => update("tradeMethod", e.target.value)}>
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
