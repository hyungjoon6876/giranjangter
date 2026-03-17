"use client";

import type { Server, Category } from "@/lib/types";

interface ListingFiltersProps {
  servers: Server[];
  selectedServer: string | null;
  onServerChange: (serverId: string | null) => void;
  categories?: Category[];
  selectedCategory?: string | null;
  onCategoryChange?: (categoryId: string | null) => void;
  searchQuery: string;
  onSearchChange: (q: string) => void;
}

export function ListingFilters({
  servers,
  selectedServer,
  onServerChange,
  categories = [],
  selectedCategory = null,
  onCategoryChange,
  searchQuery,
  onSearchChange,
}: ListingFiltersProps) {
  const topCategories = categories.filter((c) => !c.parentId);

  const chipBase =
    "px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold";
  const chipActive = "btn-gold-gradient text-white";
  const chipInactive = "bg-medium text-text-secondary hover:bg-light";

  const smallChipBase =
    "px-3 py-1.5 rounded-lg text-xs whitespace-nowrap transition-colors border";
  const smallChipActive = "border-gold text-gold bg-gold/10";
  const smallChipInactive =
    "border-border text-text-secondary bg-medium hover:bg-light";

  return (
    <div className="flex flex-col gap-3 mb-4">
      {/* Server filter + search */}
      <div className="flex flex-col lg:flex-row lg:items-center gap-3">
        <div role="group" aria-label="서버 필터" className="flex flex-wrap gap-2">
          <button
            onClick={() => onServerChange(null)}
            aria-pressed={selectedServer === null}
            className={`${chipBase} ${selectedServer === null ? chipActive : chipInactive}`}
          >
            전체
          </button>
          {servers.map((s) => (
            <button
              key={s.serverId}
              onClick={() => onServerChange(s.serverId)}
              aria-pressed={selectedServer === s.serverId}
              className={`${chipBase} ${selectedServer === s.serverId ? chipActive : chipInactive}`}
            >
              {s.serverName}
            </button>
          ))}
        </div>
        <search className="lg:ml-auto">
          <input
            type="search"
            aria-label="매물 검색"
            value={searchQuery}
            onChange={(e) => onSearchChange(e.target.value)}
            placeholder="아이템 검색..."
            className="w-full lg:w-60 bg-card border border-border rounded-lg px-3 py-2 text-sm text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim"
          />
        </search>
      </div>

      {/* Category filter chips */}
      {topCategories.length > 0 && onCategoryChange && (
        <div
          role="group"
          aria-label="카테고리 필터"
          className="flex flex-wrap gap-2"
        >
          <button
            onClick={() => onCategoryChange(null)}
            aria-pressed={selectedCategory === null}
            className={`${smallChipBase} ${selectedCategory === null ? smallChipActive : smallChipInactive}`}
          >
            전체
          </button>
          {topCategories.map((c) => (
            <button
              key={c.categoryId}
              onClick={() => onCategoryChange(c.categoryId)}
              aria-pressed={selectedCategory === c.categoryId}
              className={`${smallChipBase} ${selectedCategory === c.categoryId ? smallChipActive : smallChipInactive}`}
            >
              {c.categoryName}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
