# UX Phase 2: 기능 완성 Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Complete the core transactional features: image upload in listing creation/detail, listing edit page, item autocomplete search, profile edit, and user review display.

**Architecture:** Frontend-only changes in `web/`. All backend APIs already exist and have been verified:
- `POST /uploads/images` — multipart/form-data, `file` field, max 10MB, returns `{imageId, url, thumbnailUrl}`
- `PATCH /listings/:id` — partial update, owner check, terminal state check
- `GET /items/search?q=&categoryId=` — public, ILIKE, limit 20
- `PATCH /me/profile` — optional nickname/introduction/primaryServerId/avatarUrl
- `GET /users/:userId/reviews` — limit 50, DESC

**Tech Stack:** Next.js 16, React 19, TanStack Query, TailwindCSS 3.4

**Spec:** `docs/ux-analysis-2026-03-17.md` (Phase 2 features)

---

## File Structure

| File | Responsibility | Action |
|------|---------------|--------|
| `web/lib/api-client.ts` | `uploadImage`, `searchItems`, `updateProfile`, `getUserReviews` 메서드 | Modify |
| `web/lib/types.ts` | `ItemSearchResult`, `Review` 타입 추가 | Modify |
| `web/lib/hooks/use-listings.ts` | `useUpdateListing` 이미 존재, 확인만 | — |
| `web/lib/hooks/use-profile.ts` | `useUpdateProfile` 훅 추가 | Modify |
| `web/lib/hooks/use-items.ts` | `useItemSearch` 훅 (debounced autocomplete) | Create |
| `web/lib/hooks/use-reviews.ts` | `useUserReviews` 훅 | Create |
| `web/components/forms/image-upload.tsx` | 이미지 업로드 컴포넌트 (드래그&드롭 + 프리뷰) | Create |
| `web/components/forms/item-autocomplete.tsx` | 아이템 자동완성 검색 입력 | Create |
| `web/app/listings/[id]/edit/page.tsx` | 매물 수정 페이지 | Create |
| `web/app/listings/[id]/page.tsx` | 이미지 갤러리 추가 | Modify |
| `web/app/create/page.tsx` | 이미지 업로드 + 아이템 자동완성 통합 | Modify |
| `web/app/profile/edit/page.tsx` | 프로필 수정 페이지 | Create |
| `web/app/profile/page.tsx` | "수정" 버튼 + 리뷰 링크 추가 | Modify |
| `web/app/profile/[userId]/reviews/page.tsx` | 사용자 리뷰 목록 페이지 | Create |
| `web/components/listing/listing-info.tsx` | `AuthorSection`에 리뷰 링크 추가 | Modify |

---

## Chunk 1: API 레이어 + 타입 확장

### Task 1: 타입 정의 추가

Phase 2에서 필요한 새 타입들을 추가한다.

**Files:**
- Modify: `web/lib/types.ts`

- [ ] **Step 1: Add new types**

`web/lib/types.ts` 끝에 추가:

```typescript
export interface ItemSearchResult {
  id: string;
  name: string;
  categoryId: string;
  iconUrl?: string;
}

export interface Review {
  reviewId: string;
  rating: "positive" | "negative";
  comment: string;
  reviewerNickname: string;
  createdAt: string;
}

export interface UploadedImage {
  imageId: string;
  url: string;
  thumbnailUrl: string;
  sizeBytes?: number;
}
```

- [ ] **Step 2: Commit**

```bash
git add web/lib/types.ts
git commit -m "feat(web): Phase 2 타입 추가 — ItemSearchResult, Review, UploadedImage"
```

---

### Task 2: API 클라이언트 메서드 추가

`ApiClient`에 이미지 업로드, 아이템 검색, 프로필 수정, 리뷰 조회 메서드를 추가한다.

**Files:**
- Modify: `web/lib/api-client.ts`

- [ ] **Step 1: Add new methods to ApiClient**

`web/lib/api-client.ts`의 `ApiClient` 클래스에 다음 메서드들을 추가한다. 각 메서드의 위치는 관련 섹션 근처에 배치.

`// Profile` 섹션 뒤에:

```typescript
async updateProfile(data: {
  nickname?: string;
  introduction?: string;
  primaryServerId?: string;
  avatarUrl?: string;
}): Promise<User> {
  return this.fetch("/me/profile", {
    method: "PATCH",
    body: JSON.stringify(data),
  });
}
```

`// Master data` 섹션 뒤에:

```typescript
// Upload
async uploadImage(file: File): Promise<UploadedImage> {
  const formData = new FormData();
  formData.append("file", file);

  const headers: Record<string, string> = {};
  if (this.accessToken) {
    headers["Authorization"] = `Bearer ${this.accessToken}`;
  }
  // Do NOT set Content-Type — browser sets multipart boundary automatically

  let res = await fetch(`${API_BASE}/uploads/images`, {
    method: "POST",
    headers,
    body: formData,
  });

  // Auto-refresh on 401 (same pattern as this.fetch)
  if (res.status === 401 && this.refreshToken) {
    const refreshed = await this.doRefresh();
    if (refreshed) {
      headers["Authorization"] = `Bearer ${this.accessToken}`;
      res = await fetch(`${API_BASE}/uploads/images`, {
        method: "POST",
        headers,
        body: formData,
      });
    }
  }

  if (!res.ok) {
    const err = await res
      .json()
      .catch(() => ({ error: { code: "UNKNOWN", message: res.statusText } }));
    throw err;
  }
  return res.json();
}

// Item search
async searchItems(params: {
  q?: string;
  categoryId?: string;
}): Promise<ItemSearchResult[]> {
  const qs = new URLSearchParams();
  if (params.q) qs.set("q", params.q);
  if (params.categoryId) qs.set("categoryId", params.categoryId);
  const res = await this.fetch<{ data: ItemSearchResult[] | null }>(
    `/items/search?${qs}`,
  );
  return res.data ?? [];
}

// User reviews
async getUserReviews(userId: string): Promise<Review[]> {
  const res = await this.fetch<{ data: Review[] | null }>(
    `/users/${userId}/reviews`,
  );
  return res.data ?? [];
}
```

주의: `uploadImage`는 `this.fetch`를 사용하지 않는다 — `Content-Type`을 자동 설정하면 multipart boundary가 깨지므로 직접 `fetch`를 호출한다.

타입 import도 추가:

```typescript
import type {
  Listing,
  ChatRoom,
  Message,
  Server,
  Category,
  PaginatedResponse,
  AuthResponse,
  User,
  Notification,
  ItemSearchResult,
  Review,
  UploadedImage,
} from "./types";
```

- [ ] **Step 2: Run type check**

```bash
cd web && npx tsc --noEmit
```

- [ ] **Step 3: Commit**

```bash
git add web/lib/api-client.ts
git commit -m "feat(web): API 클라이언트에 업로드/아이템검색/프로필수정/리뷰 메서드 추가"
```

---

### Task 3: React Query 훅 추가

새 API 메서드에 대한 TanStack Query 훅을 만든다.

**Files:**
- Create: `web/lib/hooks/use-items.ts`
- Create: `web/lib/hooks/use-reviews.ts`
- Modify: `web/lib/hooks/use-profile.ts`

- [ ] **Step 1: Create useItemSearch hook**

```typescript
// web/lib/hooks/use-items.ts
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useItemSearch(q: string, categoryId?: string) {
  return useQuery({
    queryKey: ["items-search", q, categoryId],
    queryFn: () => apiClient.searchItems({ q, categoryId }),
    enabled: q.length >= 1,
    staleTime: 60_000,
  });
}
```

- [ ] **Step 2: Create useUserReviews hook**

```typescript
// web/lib/hooks/use-reviews.ts
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useUserReviews(userId: string) {
  return useQuery({
    queryKey: ["user-reviews", userId],
    queryFn: () => apiClient.getUserReviews(userId),
    enabled: !!userId,
  });
}
```

- [ ] **Step 3: Add useUpdateProfile hook**

`web/lib/hooks/use-profile.ts`에 추가:

```typescript
export function useUpdateProfile() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: {
      nickname?: string;
      introduction?: string;
      primaryServerId?: string;
      avatarUrl?: string;
    }) => apiClient.updateProfile(data),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["me"] }),
  });
}
```

- [ ] **Step 4: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 5: Commit**

```bash
git add web/lib/hooks/use-items.ts web/lib/hooks/use-reviews.ts web/lib/hooks/use-profile.ts
git commit -m "feat(web): Phase 2 훅 추가 — useItemSearch, useUserReviews, useUpdateProfile"
```

---

## Chunk 2: 이미지 업로드 + 아이템 자동완성 컴포넌트

### Task 4: 이미지 업로드 컴포넌트

드래그&드롭 + 클릭으로 이미지를 업로드하고 프리뷰를 보여주는 컴포넌트. 최대 5장, 순서 변경 가능.

**Files:**
- Create: `web/components/forms/image-upload.tsx`

- [ ] **Step 1: Create ImageUpload component**

```typescript
// web/components/forms/image-upload.tsx
"use client";

import { useState, useRef, useCallback } from "react";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";
import type { UploadedImage } from "@/lib/types";

interface ImageUploadProps {
  images: UploadedImage[];
  onChange: (images: UploadedImage[]) => void;
  maxImages?: number;
}

const ALLOWED_TYPES = ["image/jpeg", "image/png", "image/webp"];
const MAX_SIZE = 10 * 1024 * 1024; // 10MB

export function ImageUpload({
  images,
  onChange,
  maxImages = 5,
}: ImageUploadProps) {
  const [uploading, setUploading] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);
  const { addToast } = useToast();

  const upload = useCallback(
    async (files: FileList | File[]) => {
      const remaining = maxImages - images.length;
      if (remaining <= 0) {
        addToast("error", `최대 ${maxImages}장까지 업로드할 수 있습니다`);
        return;
      }

      const validFiles = Array.from(files)
        .slice(0, remaining)
        .filter((f) => {
          if (!ALLOWED_TYPES.includes(f.type)) {
            addToast("error", `${f.name}: JPG, PNG, WebP만 가능합니다`);
            return false;
          }
          if (f.size > MAX_SIZE) {
            addToast("error", `${f.name}: 10MB 이하만 가능합니다`);
            return false;
          }
          return true;
        });

      if (!validFiles.length) return;

      setUploading(true);
      try {
        const results: UploadedImage[] = [];
        for (const file of validFiles) {
          const result = await apiClient.uploadImage(file);
          results.push(result);
        }
        onChange([...images, ...results]);
      } catch {
        addToast("error", "이미지 업로드에 실패했습니다");
      } finally {
        setUploading(false);
      }
    },
    [images, maxImages, onChange, addToast],
  );

  const handleDrop = useCallback(
    (e: React.DragEvent) => {
      e.preventDefault();
      upload(e.dataTransfer.files);
    },
    [upload],
  );

  const handleRemove = (index: number) => {
    onChange(images.filter((_, i) => i !== index));
  };

  return (
    <div>
      <div
        onDrop={handleDrop}
        onDragOver={(e) => e.preventDefault()}
        onClick={() => inputRef.current?.click()}
        className="border-2 border-dashed border-border rounded-lg p-6 text-center cursor-pointer hover:border-gold/50 transition-colors"
      >
        <input
          ref={inputRef}
          type="file"
          accept={ALLOWED_TYPES.join(",")}
          multiple
          className="hidden"
          onChange={(e) => e.target.files && upload(e.target.files)}
        />
        {uploading ? (
          <p className="text-text-secondary text-sm">업로드 중...</p>
        ) : (
          <>
            <p className="text-text-secondary text-sm">
              클릭 또는 드래그하여 이미지 추가
            </p>
            <p className="text-text-dim text-xs mt-1">
              JPG, PNG, WebP · 최대 10MB · {images.length}/{maxImages}장
            </p>
          </>
        )}
      </div>

      {/* Preview grid */}
      {images.length > 0 && (
        <div className="grid grid-cols-5 gap-2 mt-3">
          {images.map((img, i) => (
            <div key={img.imageId} className="relative group">
              <img
                src={img.thumbnailUrl || img.url}
                alt={`업로드 이미지 ${i + 1}`}
                className="w-full aspect-square object-cover rounded-lg border border-border"
              />
              <button
                type="button"
                onClick={() => handleRemove(i)}
                className="absolute -top-1 -right-1 w-5 h-5 bg-danger text-white text-xs rounded-full opacity-0 group-hover:opacity-100 transition-opacity"
                aria-label={`이미지 ${i + 1} 삭제`}
              >
                ×
              </button>
              {i === 0 && (
                <span className="absolute bottom-1 left-1 bg-gold text-dark text-[10px] px-1.5 py-0.5 rounded font-medium">
                  대표
                </span>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
```

- [ ] **Step 2: Run type check**

```bash
cd web && npx tsc --noEmit
```

- [ ] **Step 3: Commit**

```bash
git add web/components/forms/image-upload.tsx
git commit -m "feat(web): 이미지 업로드 컴포넌트 — 드래그&드롭, 프리뷰, 삭제"
```

---

### Task 5: 아이템 자동완성 검색 컴포넌트

아이템명을 입력하면 `GET /items/search`로 자동완성 결과를 보여주는 combobox.

**Files:**
- Create: `web/components/forms/item-autocomplete.tsx`

- [ ] **Step 1: Create ItemAutocomplete component**

```typescript
// web/components/forms/item-autocomplete.tsx
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
  const debounceRef = useRef<ReturnType<typeof setTimeout>>();

  // Debounce search query
  const [debouncedQuery, setDebouncedQuery] = useState("");
  useEffect(() => {
    debounceRef.current = setTimeout(() => {
      setDebouncedQuery(query);
    }, 300);
    return () => clearTimeout(debounceRef.current);
  }, [query]);

  const { data: items = [] } = useItemSearch(debouncedQuery, categoryId);

  // Close on outside click
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
```

- [ ] **Step 2: Run type check**

```bash
cd web && npx tsc --noEmit
```

- [ ] **Step 3: Commit**

```bash
git add web/components/forms/item-autocomplete.tsx
git commit -m "feat(web): 아이템 자동완성 검색 컴포넌트 — debounce, keyboard nav, a11y"
```

---

## Chunk 3: 매물 등록 폼 개선 (이미지 + 자동완성)

### Task 6: Create 페이지에 이미지 업로드 + 자동완성 통합

매물 등록 폼에 이미지 업로드와 아이템 자동완성을 통합한다.

**Files:**
- Modify: `web/app/create/page.tsx`

- [ ] **Step 1: Add imports and state**

`web/app/create/page.tsx` 상단에 import 추가:

```typescript
import { ImageUpload } from "@/components/forms/image-upload";
import { ItemAutocomplete } from "@/components/forms/item-autocomplete";
import type { UploadedImage } from "@/lib/types";
```

`form` state 아래에 이미지 state 추가:

```typescript
const [images, setImages] = useState<UploadedImage[]>([]);
```

- [ ] **Step 2: Replace itemName input with ItemAutocomplete**

현재 아이템명 input (lines 148-159)을 교체:

```typescript
// Before:
<input
  id="itemName"
  className={inputClass}
  value={form.itemName}
  onChange={(e) => update("itemName", e.target.value)}
  required
  aria-required="true"
/>

// After:
<ItemAutocomplete
  value={form.itemName}
  categoryId={form.categoryId || undefined}
  onChange={(v) => update("itemName", v)}
  required
  className={inputClass}
/>
```

- [ ] **Step 3: Add ImageUpload section**

`{/* Section: 거래 */}` 바로 위에 이미지 업로드 섹션 추가:

```typescript
{/* Section: 이미지 */}
<h3 className={sectionClass}>이미지</h3>
<ImageUpload images={images} onChange={setImages} maxImages={5} />
```

- [ ] **Step 4: Include images in submit payload**

`handleSubmit`의 `data` 객체에 이미지 추가:

```typescript
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
```

- [ ] **Step 5: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 6: Commit**

```bash
git add web/app/create/page.tsx
git commit -m "feat(web): 매물 등록에 이미지 업로드 + 아이템 자동완성 통합"
```

---

## Chunk 4: 매물 수정 페이지

### Task 7: 매물 수정 페이지 생성

기존 매물 데이터를 로드하여 수정할 수 있는 페이지. create 폼과 동일한 구조이되, 기존 데이터로 초기화한다.

**Files:**
- Create: `web/app/listings/[id]/edit/page.tsx`

- [ ] **Step 1: Create edit page**

```typescript
// web/app/listings/[id]/edit/page.tsx
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

  // Initialize form with listing data
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
            sizeBytes: 0,
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

  // Check ownership
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
        {/* 기본 정보 (읽기 전용 표시) */}
        <div className="bg-medium rounded-lg p-3 text-sm text-text-secondary">
          <span className="font-medium">{listing.serverName}</span>
          <span className="mx-2">·</span>
          <span>{listing.categoryName}</span>
          <span className="mx-2">·</span>
          <span>{listing.listingType === "sell" ? "판매" : "구매"}</span>
        </div>

        {/* Section: 아이템 */}
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

        {/* Section: 이미지 */}
        <h3 className={sectionClass}>이미지</h3>
        <ImageUpload images={images} onChange={setImages} maxImages={5} />

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
```

- [ ] **Step 2: Run type check**

```bash
cd web && npx tsc --noEmit
```

- [ ] **Step 3: Commit**

```bash
git add web/app/listings/[id]/edit/page.tsx
git commit -m "feat(web): 매물 수정 페이지 — 기존 데이터 로드, 이미지/자동완성 통합"
```

---

### Task 8: 매물 상세에 이미지 갤러리 + 수정 버튼 추가

상세 페이지에 이미지 갤러리를 표시하고, 소유자에게 수정 버튼을 보여준다.

**Files:**
- Modify: `web/app/listings/[id]/page.tsx`

- [ ] **Step 1: Add image gallery**

`web/app/listings/[id]/page.tsx`의 `{/* Item info */}` 섹션 바로 위에 이미지 갤러리 추가:

```typescript
// Import at top:
import Link from "next/link";

// After {/* Counters */} div and before {/* Item info */}:
{/* Image gallery */}
{l.images && l.images.length > 0 && (
  <div className="mb-6">
    <div className="grid grid-cols-3 sm:grid-cols-4 gap-2">
      {l.images
        .sort((a, b) => a.order - b.order)
        .map((img, i) => (
          <img
            key={img.imageId}
            src={img.url}
            alt={`${l.itemName} 이미지 ${i + 1}`}
            className={`w-full rounded-lg border border-border object-cover ${
              i === 0 ? "col-span-2 row-span-2 aspect-square" : "aspect-square"
            }`}
          />
        ))}
    </div>
  </div>
)}
```

- [ ] **Step 2: Add edit button for owner**

Action bar 안에 수정 버튼 추가. `{actions.includes("favorite") && (` 바로 위에:

```typescript
{l.isOwner && l.status === "available" && (
  <Link
    href={`/listings/${l.listingId}/edit`}
    className="p-3 border border-border rounded-lg hover:bg-medium transition-colors text-sm text-text-secondary"
  >
    수정
  </Link>
)}
```

- [ ] **Step 3: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 4: Commit**

```bash
git add web/app/listings/[id]/page.tsx
git commit -m "feat(web): 매물 상세에 이미지 갤러리 + 소유자 수정 버튼"
```

---

## Chunk 5: 프로필 수정 + 리뷰

### Task 9: 프로필 수정 페이지

닉네임, 소개글, 주 서버, 아바타를 수정할 수 있는 페이지.

**Files:**
- Create: `web/app/profile/edit/page.tsx`

- [ ] **Step 1: Create profile edit page**

```typescript
// web/app/profile/edit/page.tsx
"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
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

  const [form, setForm] = useState({
    nickname: "",
    introduction: "",
    primaryServerId: "",
  });
  const [avatarImages, setAvatarImages] = useState<UploadedImage[]>([]);
  const [initialized, setInitialized] = useState(false);

  useEffect(() => {
    if (me && !initialized) {
      setForm({
        nickname: me.nickname,
        introduction: me.introduction ?? "",
        primaryServerId: me.primaryServerId ?? "",
      });
      setInitialized(true);
    }
  }, [me, initialized]);

  useEffect(() => {
    if (!isLoggedIn) {
      requireAuth("프로필 수정");
    }
  }, [isLoggedIn]);

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
                <img
                  src={avatarImages[0].thumbnailUrl || avatarImages[0].url}
                  alt="프로필"
                  className="w-full h-full object-cover"
                />
              ) : me.avatarUrl ? (
                <img
                  src={me.avatarUrl}
                  alt="프로필"
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
```

- [ ] **Step 2: Commit**

```bash
git add web/app/profile/edit/page.tsx
git commit -m "feat(web): 프로필 수정 페이지 — 닉네임/소개/서버/아바타"
```

---

### Task 10: 프로필 페이지에 수정 버튼 추가

프로필 페이지에 "수정" 링크와 리뷰 메뉴 항목을 추가한다.

**Files:**
- Modify: `web/app/profile/page.tsx`

- [ ] **Step 1: Add edit link to profile header**

`web/app/profile/page.tsx`의 닉네임/소개 블록(lines 64-73)을 수정하여 수정 링크를 추가:

```typescript
// Before (lines 64-73):
<div>
  <h2 className="text-xl font-display font-bold text-gold">
    {me.nickname}
  </h2>
  {me.introduction && (
    <p className="text-sm text-text-secondary mt-1">
      {me.introduction}
    </p>
  )}
</div>

// After:
<div className="flex-1">
  <div className="flex items-center gap-2">
    <h2 className="text-xl font-display font-bold text-gold">
      {me.nickname}
    </h2>
    <Link
      href="/profile/edit"
      className="text-xs text-text-secondary hover:text-gold transition-colors"
    >
      수정
    </Link>
  </div>
  {me.introduction && (
    <p className="text-sm text-text-secondary mt-1">
      {me.introduction}
    </p>
  )}
</div>
```

- [ ] **Step 2: Add reviews menu item**

`menuItems` 배열에 리뷰 항목 추가:

```typescript
const menuItems = [
  { href: "/profile/listings", label: "내 매물" },
  { href: "/profile/trades", label: "내 거래" },
  { href: `/profile/${me.userId}/reviews`, label: "받은 리뷰" },
  { href: "/notifications", label: "알림" },
];
```

- [ ] **Step 3: Commit**

```bash
git add web/app/profile/page.tsx
git commit -m "feat(web): 프로필에 수정 링크 + 받은 리뷰 메뉴 추가"
```

---

### Task 11: 사용자 리뷰 목록 페이지

특정 사용자의 리뷰를 보여주는 페이지.

**Files:**
- Create: `web/app/profile/[userId]/reviews/page.tsx`

- [ ] **Step 1: Create reviews page**

```typescript
// web/app/profile/[userId]/reviews/page.tsx
"use client";

import { use } from "react";
import { useUserReviews } from "@/lib/hooks/use-reviews";
import { Loading } from "@/components/ui/loading";
import { EmptyState } from "@/components/ui/empty-state";
import { formatTimeAgo } from "@/lib/utils";

export default function UserReviewsPage({
  params,
}: {
  params: Promise<{ userId: string }>;
}) {
  const { userId } = use(params);
  const { data: reviews = [], isLoading } = useUserReviews(userId);

  if (isLoading) return <Loading />;

  if (!reviews.length) {
    return (
      <EmptyState
        title="아직 리뷰가 없습니다"
        description="거래를 완료하면 리뷰를 받을 수 있습니다"
      />
    );
  }

  return (
    <div className="max-w-lg mx-auto p-4 lg:p-6">
      <h1 className="text-2xl font-bold mb-4 text-text-primary">
        받은 리뷰 ({reviews.length})
      </h1>
      <div className="space-y-3">
        {reviews.map((r) => (
          <div
            key={r.reviewId}
            className="bg-card rounded-xl border border-border p-4"
          >
            <div className="flex items-center justify-between mb-2">
              <div className="flex items-center gap-2">
                <span className="font-medium text-text-primary">
                  {r.reviewerNickname}
                </span>
                <span
                  className={`text-xs font-medium px-2 py-0.5 rounded-full ${
                    r.rating === "positive"
                      ? "bg-green-500/10 text-green-400"
                      : "bg-danger/10 text-danger"
                  }`}
                >
                  {r.rating === "positive" ? "👍 좋아요" : "👎 아쉬워요"}
                </span>
              </div>
              <span className="text-xs text-text-dim">
                {formatTimeAgo(r.createdAt)}
              </span>
            </div>
            {r.comment && (
              <p className="text-sm text-text-secondary">{r.comment}</p>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
```

- [ ] **Step 2: Commit**

```bash
git add web/app/profile/[userId]/reviews/page.tsx
git commit -m "feat(web): 사용자 리뷰 목록 페이지 — 별점, 시간, 코멘트 표시"
```

---

### Task 12: 매물 상세 판매자 영역에 리뷰 링크 추가

`AuthorSection`에 리뷰 수와 링크를 추가한다.

**Files:**
- Modify: `web/components/listing/listing-info.tsx`

- [ ] **Step 1: Add reviews link to AuthorSection**

`web/components/listing/listing-info.tsx`의 `AuthorSection`을 수정:

```typescript
import Link from "next/link";
import type { Author } from "@/lib/types";

// ... InfoRow stays the same ...

export function AuthorSection({ author }: { author: Author }) {
  return (
    <Link
      href={`/profile/${author.userId}/reviews`}
      className="flex items-center gap-3 hover:opacity-80 transition-opacity"
    >
      <div className="w-10 h-10 rounded-full bg-border flex items-center justify-center font-bold text-text-secondary">
        {author.nickname?.[0] ?? "?"}
      </div>
      <div className="flex-1">
        <div className="font-semibold">{author.nickname}</div>
        <div className="text-sm text-text-secondary">
          거래 {author.completedTradeCount ?? 0}회
          {author.trustBadge && ` · ${author.trustBadge}`}
        </div>
      </div>
      <span className="text-xs text-text-dim">리뷰 보기 ›</span>
    </Link>
  );
}
```

- [ ] **Step 2: Run full tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 3: Commit**

```bash
git add web/components/listing/listing-info.tsx
git commit -m "feat(web): 판매자 영역에 리뷰 페이지 링크 추가"
```

---

## Summary

| Task | 파일 | 기능 | 복잡도 |
|------|------|------|--------|
| 1 | types.ts | 새 타입 정의 | S |
| 2 | api-client.ts | API 메서드 4개 추가 | M |
| 3 | hooks/*.ts | React Query 훅 3개 | S |
| 4 | image-upload.tsx | 이미지 업로드 컴포넌트 | M |
| 5 | item-autocomplete.tsx | 자동완성 combobox | M |
| 6 | create/page.tsx | 등록 폼에 이미지+자동완성 통합 | M |
| 7 | listings/[id]/edit/page.tsx | 매물 수정 페이지 | L |
| 8 | listings/[id]/page.tsx | 이미지 갤러리 + 수정 버튼 | S |
| 9 | profile/edit/page.tsx | 프로필 수정 페이지 | M |
| 10 | profile/page.tsx | 수정 링크 + 리뷰 메뉴 | S |
| 11 | profile/[userId]/reviews/page.tsx | 리뷰 목록 페이지 | S |
| 12 | listing-info.tsx | 판매자 리뷰 링크 | S |

**Dependencies:**
- Task 1 → Task 2 → Task 3 (순차 — 타입 → API → 훅)
- Task 4 depends on Task 1 (UploadedImage 타입) + Task 2 (uploadImage API)
- Task 5 depends on Task 1 (ItemSearchResult 타입) + Task 3 (useItemSearch 훅)
- Task 4, 5 (Task 3 이후 병렬 가능)
- Task 6 depends on Task 4 + 5
- Task 7 depends on Task 4 + 5 + Task 2 (useUpdateListing API)
- Task 8 depends on Task 7 (수정 버튼 Link), 이미지 갤러리는 독립
- Task 9 depends on Task 3 (useUpdateProfile) + Task 1 (UploadedImage) + Task 4 (ImageUpload)
- Task 10, 11, 12 (독립, 병렬 가능, Task 3 이후)

**Phase 3 대상 (별도 계획):** 카카오/네이버 OAuth, 공유 기능, 사용자 차단, 푸터/앱 정보, PWA
