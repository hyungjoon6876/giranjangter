"use client";

import { useState, useRef, useEffect, useCallback, useMemo } from "react";
import Image from "next/image";
import { useQuery } from "@tanstack/react-query";
import { useItemSearch } from "@/lib/hooks/use-items";
import { apiClient, assetUrl } from "@/lib/api-client";
import type { ItemSearchResult } from "@/lib/types";

interface ItemAutocompleteProps {
  value: string;
  onChange: (value: string) => void;
  onSelect?: (item: ItemSearchResult) => void;
  onEnhancementChange?: (level: number) => void;
  enhancementLevel?: number;
  required?: boolean;
  className?: string;
}

export function ItemAutocomplete({
  value,
  onChange,
  onSelect,
  onEnhancementChange,
  enhancementLevel = 0,
  required,
  className,
}: ItemAutocompleteProps) {
  const [selectedItem, setSelectedItem] = useState<ItemSearchResult | null>(
    null,
  );
  const [activeCategoryId, setActiveCategoryId] = useState<string | null>(null);
  const [activeSubCategoryId, setActiveSubCategoryId] = useState<string | null>(
    null,
  );
  const [searchQuery, setSearchQuery] = useState("");
  const [debouncedQuery, setDebouncedQuery] = useState("");
  const debounceRef = useRef<ReturnType<typeof setTimeout>>(undefined);

  // Debounce search input
  useEffect(() => {
    debounceRef.current = setTimeout(() => {
      setDebouncedQuery(searchQuery);
    }, 300);
    return () => clearTimeout(debounceRef.current);
  }, [searchQuery]);

  // Sync with external value reset
  useEffect(() => {
    if (!value && selectedItem) {
      setSelectedItem(null);
    }
  }, [value, selectedItem]);

  // Load categories
  const { data: categories = [] } = useQuery({
    queryKey: ["categories"],
    queryFn: () => apiClient.getCategories(),
  });

  const topCategories = useMemo(
    () => categories.filter((c) => !c.parentId),
    [categories],
  );
  const subCategories = useMemo(
    () => categories.filter((c) => c.parentId === activeCategoryId),
    [categories, activeCategoryId],
  );

  // Effective category for search
  const effectiveCategoryId =
    activeSubCategoryId || activeCategoryId || undefined;
  const { data: items = [] } = useItemSearch(
    debouncedQuery,
    effectiveCategoryId,
  );

  // Resolve category path for selected item
  const categoryPath = useMemo(() => {
    if (!selectedItem) return "";
    const sub = categories.find(
      (c) => c.categoryId === selectedItem.categoryId,
    );
    if (!sub) return "";
    const parent = categories.find((c) => c.categoryId === sub.parentId);
    if (parent) return `${parent.categoryName} > ${sub.categoryName}`;
    return sub.categoryName;
  }, [selectedItem, categories]);

  function handleSelectItem(item: ItemSearchResult) {
    setSelectedItem(item);
    onChange(item.name);
    onSelect?.(item);
    if (item.isEnchantable) {
      onEnhancementChange?.(0);
    }
  }

  function handleClearItem() {
    setSelectedItem(null);
    onChange("");
    setSearchQuery("");
    setDebouncedQuery("");
    setActiveCategoryId(null);
    setActiveSubCategoryId(null);
    onEnhancementChange?.(0);
  }

  const handleSearchInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setSearchQuery(e.target.value);
    },
    [],
  );

  function handleCategoryClick(categoryId: string | null) {
    if (categoryId === activeCategoryId) {
      // Toggle off
      setActiveCategoryId(null);
      setActiveSubCategoryId(null);
    } else {
      setActiveCategoryId(categoryId);
      setActiveSubCategoryId(null);
    }
  }

  function handleSubCategoryClick(categoryId: string) {
    if (categoryId === activeSubCategoryId) {
      setActiveSubCategoryId(null);
    } else {
      setActiveSubCategoryId(categoryId);
    }
  }

  // ── Selected state ──
  if (selectedItem) {
    return (
      <div className={className}>
        <div className="border border-gold/40 rounded-lg bg-card p-3">
          <div className="flex items-start gap-3">
            {selectedItem.iconUrl && (
              <Image
                src={assetUrl(selectedItem.iconUrl)}
                alt={selectedItem.name}
                width={48}
                height={48}
                unoptimized
                className="w-12 h-12 shrink-0"
              />
            )}
            <div className="flex-1 min-w-0">
              <div className="flex items-center justify-between gap-2">
                <span className="font-medium text-text-primary truncate">
                  {selectedItem.name}
                </span>
                <button
                  type="button"
                  onClick={handleClearItem}
                  className="shrink-0 text-xs text-gold hover:text-gold/80 transition-colors"
                  aria-label="아이템 변경"
                >
                  변경
                </button>
              </div>
              {categoryPath && (
                <p className="text-xs text-text-secondary mt-0.5">
                  {categoryPath}
                </p>
              )}
              {selectedItem.optionText && (
                <p className="text-xs text-text-dim mt-0.5">
                  {selectedItem.optionText}
                </p>
              )}
            </div>
          </div>

          {/* Enhancement slider */}
          {selectedItem.isEnchantable && selectedItem.maxEnchantLevel > 0 && (
            <div className="mt-3 pt-3 border-t border-border">
              <div className="flex items-center justify-between mb-1.5">
                <span className="text-xs text-text-secondary">강화</span>
                <span className="text-sm font-medium text-gold">
                  +{enhancementLevel}
                </span>
              </div>
              <input
                type="range"
                min={0}
                max={selectedItem.maxEnchantLevel}
                value={enhancementLevel}
                onChange={(e) =>
                  onEnhancementChange?.(Number(e.target.value))
                }
                className="w-full h-1.5 rounded-full appearance-none cursor-pointer accent-gold bg-border"
                aria-label={`강화 수치: +${enhancementLevel}`}
              />
              <div className="relative text-[10px] text-text-dim mt-0.5">
                <div className="flex justify-between">
                  <span>0</span>
                  <span>{selectedItem.maxEnchantLevel}</span>
                </div>
                {selectedItem.safeEnchantLevel > 0 && (
                  <span
                    className="absolute top-0 text-text-secondary"
                    style={{
                      left: `${(selectedItem.safeEnchantLevel / selectedItem.maxEnchantLevel) * 100}%`,
                      transform: "translateX(-50%)",
                    }}
                  >
                    안전+{selectedItem.safeEnchantLevel}
                  </span>
                )}
              </div>
            </div>
          )}
        </div>
        {/* Hidden input for form validation */}
        {required && (
          <input
            type="text"
            value={value}
            required
            aria-required="true"
            className="sr-only"
            tabIndex={-1}
            readOnly
          />
        )}
      </div>
    );
  }

  // ── Unselected state ──
  return (
    <div className={className}>
      {/* Category tabs */}
      <div
        className="flex gap-1.5 overflow-x-auto pb-2 scrollbar-hide"
        role="tablist"
        aria-label="아이템 카테고리"
      >
        <button
          type="button"
          role="tab"
          aria-selected={activeCategoryId === null}
          onClick={() => handleCategoryClick(null)}
          className={`shrink-0 px-3 py-1.5 rounded-full text-xs font-medium transition-colors ${
            activeCategoryId === null
              ? "bg-gold text-white"
              : "bg-medium text-text-secondary hover:text-text-primary"
          }`}
        >
          전체
        </button>
        {topCategories.map((cat) => (
          <button
            key={cat.categoryId}
            type="button"
            role="tab"
            aria-selected={activeCategoryId === cat.categoryId}
            onClick={() => handleCategoryClick(cat.categoryId)}
            className={`shrink-0 px-3 py-1.5 rounded-full text-xs font-medium transition-colors ${
              activeCategoryId === cat.categoryId
                ? "bg-gold text-white"
                : "bg-medium text-text-secondary hover:text-text-primary"
            }`}
          >
            {cat.categoryName}
          </button>
        ))}
      </div>

      {/* Sub-category chips */}
      {subCategories.length > 0 && (
        <div className="flex gap-1.5 overflow-x-auto pb-2 scrollbar-hide">
          {subCategories.map((sub) => (
            <button
              key={sub.categoryId}
              type="button"
              onClick={() => handleSubCategoryClick(sub.categoryId)}
              className={`shrink-0 px-2.5 py-1 rounded text-[11px] font-medium transition-colors ${
                activeSubCategoryId === sub.categoryId
                  ? "bg-gold/20 text-gold border border-gold/30"
                  : "bg-medium/60 text-text-dim hover:text-text-secondary border border-transparent"
              }`}
            >
              {sub.categoryName}
            </button>
          ))}
        </div>
      )}

      {/* Search input */}
      <div className="relative mb-2">
        <svg
          className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-text-dim pointer-events-none"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          aria-hidden="true"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M21 21l-4.35-4.35M11 19a8 8 0 100-16 8 8 0 000 16z"
          />
        </svg>
        <input
          type="text"
          value={searchQuery}
          onChange={handleSearchInput}
          placeholder="아이템 검색..."
          className="w-full bg-card border border-border rounded-lg pl-9 pr-3 py-2.5 text-base text-text-primary outline-none focus:border-gold focus:ring-1 focus:ring-gold placeholder:text-text-dim"
          aria-label="아이템 검색"
        />
      </div>

      {/* Item grid */}
      {items.length > 0 ? (
        <div
          className="grid grid-cols-4 sm:grid-cols-5 gap-1.5 max-h-52 overflow-y-auto pr-0.5"
          role="listbox"
          aria-label="검색 결과"
        >
          {items.map((item) => (
            <button
              key={item.id}
              type="button"
              role="option"
              aria-selected={false}
              onClick={() => handleSelectItem(item)}
              className="flex flex-col items-center gap-1 p-2 rounded-lg bg-medium/40 hover:bg-gold/10 hover:ring-1 hover:ring-gold/30 transition-all cursor-pointer group"
            >
              {item.iconUrl ? (
                <Image
                  src={assetUrl(item.iconUrl)}
                  alt={item.name}
                  width={40}
                  height={40}
                  unoptimized
                  className="w-10 h-10"
                />
              ) : (
                <div className="w-10 h-10 rounded bg-medium flex items-center justify-center">
                  <span className="text-text-dim text-xs">?</span>
                </div>
              )}
              <span className="text-[11px] text-text-secondary group-hover:text-text-primary text-center leading-tight line-clamp-2">
                {item.name}
              </span>
            </button>
          ))}
        </div>
      ) : (debouncedQuery || effectiveCategoryId) ? (
        <p className="text-center text-xs text-text-dim py-6">
          검색 결과가 없습니다
        </p>
      ) : (
        <p className="text-center text-xs text-text-dim py-6">
          카테고리를 선택하거나 검색어를 입력하세요
        </p>
      )}

      {/* Hidden input for form validation */}
      {required && (
        <input
          type="text"
          value={value}
          required
          aria-required="true"
          className="sr-only"
          tabIndex={-1}
          readOnly
        />
      )}
    </div>
  );
}
