"use client";

import type { Server } from "@/lib/types";

interface ListingFiltersProps {
  servers: Server[];
  selectedServer: string | null;
  onServerChange: (serverId: string | null) => void;
  searchQuery: string;
  onSearchChange: (q: string) => void;
}

export function ListingFilters({
  servers, selectedServer, onServerChange, searchQuery, onSearchChange,
}: ListingFiltersProps) {
  return (
    <div className="flex flex-col lg:flex-row lg:items-center gap-3 mb-4">
      <div role="group" aria-label="서버 필터" className="flex gap-2 overflow-x-auto pb-1">
        <button
          onClick={() => onServerChange(null)}
          aria-pressed={selectedServer === null}
          className={`px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold ${
            selectedServer === null
              ? "btn-gold-gradient text-white"
              : "bg-medium text-text-secondary hover:bg-light"
          }`}
        >
          전체
        </button>
        {servers.map((s) => (
          <button
            key={s.serverId}
            onClick={() => onServerChange(s.serverId)}
            aria-pressed={selectedServer === s.serverId}
            className={`px-3 py-1.5 rounded-lg text-sm whitespace-nowrap transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold ${
              selectedServer === s.serverId
                ? "btn-gold-gradient text-white"
                : "bg-medium text-text-secondary hover:bg-light"
            }`}
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
  );
}
