"use client";

import { useState, useRef, useEffect, useCallback } from "react";
import { useItemSearch } from "@/lib/hooks/use-items";
import { assetUrl } from "@/lib/api-client";
import type { ItemSearchResult } from "@/lib/types";

interface ItemAutocompleteProps {
  value: string;
  categoryId?: string;
  onChange: (value: string) => void;
  onSelect?: (item: ItemSearchResult) => void;
  required?: boolean;
  className?: string;
}

export function ItemAutocomplete({
  value,
  categoryId,
  onChange,
  onSelect,
  required,
  className,
}: ItemAutocompleteProps) {
  const [query, setQuery] = useState(value);
  const [open, setOpen] = useState(false);
  const [selectedIndex, setSelectedIndex] = useState(-1);
  const containerRef = useRef<HTMLDivElement>(null);
  const debounceRef = useRef<ReturnType<typeof setTimeout>>(undefined);

  const [debouncedQuery, setDebouncedQuery] = useState("");
  useEffect(() => {
    debounceRef.current = setTimeout(() => {
      setDebouncedQuery(query);
    }, 300);
    return () => clearTimeout(debounceRef.current);
  }, [query]);

  // Sync external value changes
  useEffect(() => {
    setQuery(value);
  }, [value]);

  const { data: items = [] } = useItemSearch(debouncedQuery, categoryId);

  useEffect(() => {
    const handler = (e: MouseEvent) => {
      if (
        containerRef.current &&
        !containerRef.current.contains(e.target as Node)
      ) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handler);
    return () => document.removeEventListener("mousedown", handler);
  }, []);

  const handleInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const v = e.target.value;
      setQuery(v);
      onChange(v);
      setOpen(v.length >= 1);
      setSelectedIndex(-1);
    },
    [onChange],
  );

  const handleSelect = useCallback(
    (item: ItemSearchResult) => {
      setQuery(item.name);
      onChange(item.name);
      onSelect?.(item);
      setOpen(false);
    },
    [onChange, onSelect],
  );

  const handleKeyDown = useCallback(
    (e: React.KeyboardEvent) => {
      if (!open || !items.length) return;
      if (e.key === "ArrowDown") {
        e.preventDefault();
        setSelectedIndex((i) => Math.min(i + 1, items.length - 1));
      } else if (e.key === "ArrowUp") {
        e.preventDefault();
        setSelectedIndex((i) => Math.max(i - 1, 0));
      } else if (e.key === "Enter" && selectedIndex >= 0) {
        e.preventDefault();
        handleSelect(items[selectedIndex]);
      } else if (e.key === "Escape") {
        setOpen(false);
      }
    },
    [open, items, selectedIndex, handleSelect],
  );

  return (
    <div ref={containerRef} className="relative">
      <input
        type="text"
        role="combobox"
        aria-expanded={open}
        aria-autocomplete="list"
        aria-activedescendant={
          selectedIndex >= 0 ? `item-opt-${selectedIndex}` : undefined
        }
        value={query}
        onChange={handleInput}
        onKeyDown={handleKeyDown}
        onFocus={() => query.length >= 1 && setOpen(true)}
        required={required}
        aria-required={required}
        className={className}
        placeholder="아이템 검색..."
      />

      {open && items.length > 0 && (
        <ul
          role="listbox"
          className="absolute z-20 top-full left-0 right-0 mt-1 bg-card border border-border rounded-lg shadow-lg max-h-60 overflow-y-auto"
        >
          {items.map((item, i) => (
            <li
              key={item.id}
              id={`item-opt-${i}`}
              role="option"
              aria-selected={i === selectedIndex}
              onClick={() => handleSelect(item)}
              className={`flex items-center gap-2 px-3 py-2 cursor-pointer text-sm ${
                i === selectedIndex
                  ? "bg-gold/10 text-gold"
                  : "text-text-primary hover:bg-medium"
              }`}
            >
              {item.iconUrl && (
                <img
                  src={assetUrl(item.iconUrl)}
                  alt=""
                  className="w-6 h-6"
                />
              )}
              <span>{item.name}</span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
