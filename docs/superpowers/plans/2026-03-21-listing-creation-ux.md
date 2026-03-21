# Listing Creation UX Improvement Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 매물 등록 폼의 텍스트 입력을 최소화 — DB 아이템 마스터 데이터(아이콘, 카테고리, 옵션, 강화)를 활용하는 스마트 아이템 피커 + 자동 채움

**Architecture:** 백엔드 검색 API 확장 → 프론트 타입 확장 → 아이템 피커 컴포넌트 재설계 → 등록 폼 자동 채움 통합. 기존 등록 API(handleCreateListing)는 변경 없음.

**Tech Stack:** Next.js 16, React 19, TanStack Query, Go/Gin, PostgreSQL, TailwindCSS

**Spec:** `docs/superpowers/specs/2026-03-21-listing-creation-ux-design.md`

---

## Task 1: 백엔드 아이템 검색 API 확장

**Files:**
- Modify: `backend/internal/repository/interfaces.go:645-651` (ItemSearchResult struct)
- Modify: `backend/internal/repository/postgres_master.go:50-77` (SearchItems SQL)
- Modify: `backend/cmd/server/handlers_master.go:46-85` (searchItems handler response)

**Context:** 현재 검색 API는 `id, name, categoryId, iconUrl`만 반환. `subCategory, optionText, isEnchantable, safeEnchantLevel, maxEnchantLevel` 추가 필요. 또한 빈 검색어 + categoryId로 아이템 목록 브라우징 허용.

- [ ] **Step 1: ItemSearchResult 구조체 확장**

`interfaces.go`의 `ItemSearchResult`에 필드 추가:

```go
type ItemSearchResult struct {
	ID              string
	Name            string
	CategoryID      string
	IconID          *string
	SubCategory    string  // NEW: NOT NULL in DB
	OptionText     *string // NEW: nullable in DB
	IsEnchantable  int     // NEW: 0 or 1, NOT NULL
	SafeEnchantLvl int     // NEW: NOT NULL DEFAULT 0
	MaxEnchantLvl  int     // NEW: NOT NULL DEFAULT 0
}
```

- [ ] **Step 2: SearchItems SQL 확장**

`postgres_master.go`의 두 SQL 쿼리에 컬럼 추가:

```sql
SELECT id, name, category_id, icon_id, sub_category, option_text, is_enchantable, safe_enchant_level, max_enchant_level
FROM item_master
WHERE name ILIKE $1 [AND category_id = $2]
ORDER BY name LIMIT 20
```

`rows.Scan`에 새 필드 추가:
```go
rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.IconID,
    &item.SubCategory, &item.OptionText, &item.IsEnchantable,
    &item.SafeEnchantLvl, &item.MaxEnchantLvl)
```

- [ ] **Step 3: 빈 검색어 + categoryId 브라우징 허용**

`handlers_master.go`의 early return 조건 변경:

현재 (line 52-55):
```go
if query == "" {
    c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
    return
}
```

변경:
```go
if query == "" && categoryID == "" {
    c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
    return
}
```

`postgres_master.go`에 categoryId만으로 검색하는 분기 추가:
```go
if query == "" {
    // categoryId only — browse mode
    rows, err = r.db.QueryContext(ctx,
        "SELECT id, name, category_id, icon_id, sub_category, option_text, is_enchantable, safe_enchant_level, max_enchant_level FROM item_master WHERE category_id = $1 ORDER BY name LIMIT 20",
        *categoryID)
} else if categoryID == nil || *categoryID == "" {
    // query only
    ...
} else {
    // both query + categoryId
    ...
}
```

- [ ] **Step 4: 핸들러 응답에 새 필드 추가**

`handlers_master.go`의 응답 맵에:

```go
items = append(items, gin.H{
    "id":              r.ID,
    "name":            r.Name,
    "categoryId":      r.CategoryID,
    "iconUrl":         iconURL,
    "subCategory":     r.SubCategory,
    "optionText":      r.OptionText,
    "isEnchantable":   r.IsEnchantable != 0,  // int → bool
    "safeEnchantLevel": r.SafeEnchantLvl,
    "maxEnchantLevel":  r.MaxEnchantLvl,
})
```

- [ ] **Step 5: 빌드 + 테스트**

Run: `cd backend && go build ./cmd/server/ && go test ./...`

- [ ] **Step 6: 커밋**

```
feat(backend): 아이템 검색 API 확장 — 옵션/강화 정보 + 카테고리 브라우징
```

---

## Task 2: 프론트엔드 타입 + API 클라이언트 + Hook 업데이트

**Files:**
- Modify: `web/lib/types.ts:124-129` (ItemSearchResult)
- Modify: `web/lib/api-client.ts:348-359` (searchItems)
- Modify: `web/lib/hooks/use-items.ts:4-11` (useItemSearch)

- [ ] **Step 1: ItemSearchResult 타입 확장**

```typescript
export interface ItemSearchResult {
  id: string;
  name: string;
  categoryId: string;
  iconUrl?: string;
  subCategory: string;        // NEW: NOT NULL in DB
  optionText?: string;        // NEW: nullable in DB
  isEnchantable: boolean;     // NEW
  safeEnchantLevel: number;   // NEW: NOT NULL DEFAULT 0
  maxEnchantLevel: number;    // NEW: NOT NULL DEFAULT 0
}
```

- [ ] **Step 2: useItemSearch hook enabled 조건 변경**

```typescript
export function useItemSearch(q: string, categoryId?: string) {
  return useQuery({
    queryKey: ["items-search", q, categoryId],
    queryFn: () => apiClient.searchItems({ q, categoryId }),
    enabled: q.length >= 1 || !!categoryId,  // CHANGED: 카테고리만으로도 검색 허용
    staleTime: 60_000,
  });
}
```

- [ ] **Step 3: 빌드 검증**

Run: `cd web && npx next build`

- [ ] **Step 4: 커밋**

```
feat(web): ItemSearchResult 타입 확장 + 카테고리 브라우징 허용
```

---

## Task 3: 아이템 피커 컴포넌트 재설계

**Files:**
- Modify: `web/components/forms/item-autocomplete.tsx` (전면 재설계)

**Context:** 현재 단순 텍스트 autocomplete → 카테고리 탭 + 검색 + 결과 그리드 + 선택 카드 + 강화 슬라이더로 변경. Props 인터페이스는 유지하되 내부 UI 전면 재설계.

- [ ] **Step 1: Props 인터페이스 확장**

```typescript
interface ItemAutocompleteProps {
  value: string;
  onChange: (value: string) => void;
  onSelect?: (item: ItemSearchResult) => void;
  onEnhancementChange?: (level: number) => void;  // NEW
  enhancementLevel?: number;                       // NEW
  required?: boolean;
  className?: string;
}
```

`categoryId` prop 제거 — 카테고리는 내부에서 관리.

- [ ] **Step 2: 내부 상태 설계**

```typescript
const [selectedItem, setSelectedItem] = useState<ItemSearchResult | null>(null);
const [activeCategoryId, setActiveCategoryId] = useState<string | null>(null);
const [searchQuery, setSearchQuery] = useState("");
const [debouncedQuery, setDebouncedQuery] = useState("");

// 카테고리 데이터 로드
const { data: categories = [] } = useQuery({
  queryKey: ["categories"],
  queryFn: () => apiClient.getCategories(),
});
const topCategories = categories.filter((c: Category) => !c.parentId);
const subCategories = categories.filter((c: Category) => c.parentId === activeCategoryId);

// 아이템 검색 (검색어 또는 카테고리 필터)
const effectiveCategoryId = subCategories.length > 0
  ? /* 선택된 서브카테고리 또는 activeCategoryId */ undefined
  : activeCategoryId;
const { data: items = [] } = useItemSearch(debouncedQuery, effectiveCategoryId ?? undefined);
```

- [ ] **Step 3: 미선택 상태 UI — 카테고리 탭 + 검색 + 그리드**

```tsx
{!selectedItem && (
  <div className="border border-border rounded-xl p-4 bg-card">
    {/* 카테고리 탭 */}
    <div className="flex gap-1 overflow-x-auto pb-2 mb-2 scrollbar-hide">
      <button
        onClick={() => setActiveCategoryId(null)}
        className={`px-3 py-1 rounded-full text-xs whitespace-nowrap ${!activeCategoryId ? 'bg-gold text-darkest' : 'bg-medium text-text-secondary'}`}
      >전체</button>
      {topCategories.map((cat) => (
        <button key={cat.categoryId} onClick={() => setActiveCategoryId(cat.categoryId)}
          className={`px-3 py-1 rounded-full text-xs whitespace-nowrap ${activeCategoryId === cat.categoryId ? 'bg-gold text-darkest' : 'bg-medium text-text-secondary'}`}
        >{cat.categoryName}</button>
      ))}
    </div>

    {/* 서브카테고리 칩 (선택된 상위 카테고리가 있을 때) */}
    {subCategories.length > 0 && (
      <div className="flex gap-1 flex-wrap mb-2">
        {subCategories.map((sub) => (
          <button key={sub.categoryId} onClick={() => setActiveSubCategoryId(sub.categoryId)}
            className={`px-2 py-0.5 rounded text-xs ${activeSubCategoryId === sub.categoryId ? 'bg-gold/20 text-gold' : 'bg-medium/50 text-text-secondary'}`}
          >{sub.categoryName}</button>
        ))}
      </div>
    )}

    {/* 검색 입력 */}
    <input
      type="text"
      value={searchQuery}
      onChange={(e) => setSearchQuery(e.target.value)}
      placeholder="아이템 검색..."
      className="w-full bg-dark border border-border rounded-lg px-3 py-2 text-base mb-3"
    />

    {/* 아이템 그리드 */}
    <div className="grid grid-cols-4 sm:grid-cols-5 gap-2 max-h-60 overflow-y-auto">
      {items.map((item) => (
        <button key={item.id} onClick={() => handleSelectItem(item)}
          className="flex flex-col items-center p-2 rounded-lg hover:bg-medium/50 transition-colors"
        >
          {item.iconUrl ? (
            <Image src={assetUrl(item.iconUrl)} alt={item.name} width={40} height={40} unoptimized />
          ) : (
            <div className="w-10 h-10 bg-medium rounded" />
          )}
          <span className="text-xs text-center mt-1 truncate w-full">{item.name}</span>
        </button>
      ))}
    </div>
  </div>
)}
```

- [ ] **Step 4: 선택됨 상태 UI — 아이템 카드 + 강화 슬라이더**

```tsx
{selectedItem && (
  <div className="border border-gold/30 rounded-xl p-4 bg-card">
    <div className="flex items-start gap-3">
      {selectedItem.iconUrl ? (
        <Image src={assetUrl(selectedItem.iconUrl)} alt={selectedItem.name} width={48} height={48} unoptimized />
      ) : (
        <div className="w-12 h-12 bg-medium rounded" />
      )}
      <div className="flex-1 min-w-0">
        <p className="font-medium">{selectedItem.name}</p>
        <p className="text-xs text-text-secondary">
          {/* 카테고리 경로: 부모 > 자식 */}
          {parentCategoryName} &gt; {selectedItem.subCategory}
        </p>
        {selectedItem.optionText && (
          <p className="text-xs text-gold mt-1">{selectedItem.optionText}</p>
        )}
      </div>
      <button onClick={handleClearItem} className="text-xs text-text-secondary hover:text-text-primary">
        변경
      </button>
    </div>

    {/* 강화 슬라이더 (enchantable일 때만) */}
    {selectedItem.isEnchantable && (
      <div className="mt-3">
        <div className="flex items-center justify-between text-sm mb-1">
          <span className="text-text-secondary">강화</span>
          <span className="font-bold text-gold">+{enhancementLevel ?? 0}</span>
        </div>
        <input
          type="range"
          min={0}
          max={selectedItem.maxEnchantLevel ?? 10}
          value={enhancementLevel ?? 0}
          onChange={(e) => onEnhancementChange?.(Number(e.target.value))}
          className="w-full accent-gold"
        />
        <div className="flex justify-between text-xs text-text-secondary mt-0.5">
          <span>0</span>
          {selectedItem.safeEnchantLevel && (
            <span className="text-green-400">안전 +{selectedItem.safeEnchantLevel}</span>
          )}
          <span>+{selectedItem.maxEnchantLevel ?? 10}</span>
        </div>
      </div>
    )}
  </div>
)}
```

- [ ] **Step 5: 아이템 선택/해제 핸들러**

```typescript
function handleSelectItem(item: ItemSearchResult) {
  setSelectedItem(item);
  onChange(item.name);        // 부모 폼에 이름 전달
  onSelect?.(item);          // 부모 폼에 전체 객체 전달
}

function handleClearItem() {
  setSelectedItem(null);
  onChange("");
  setSearchQuery("");
}
```

- [ ] **Step 6: 기존 value prop 동기화**

외부에서 `value`가 변경되면 (예: 폼 리셋) selectedItem도 초기화:

```typescript
useEffect(() => {
  if (!value && selectedItem) {
    setSelectedItem(null);
  }
}, [value]);
```

- [ ] **Step 7: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 8: 커밋**

```
feat(web): 스마트 아이템 피커 — 카테고리 탭 + 아이콘 그리드 + 강화 슬라이더
```

---

## Task 4: 등록 폼 자동 채움 통합

**Files:**
- Modify: `web/app/create/page.tsx` (자동 채움 + 필드 순서 + 제목 생성)

**Context:** 아이템 선택 시 카테고리 자동 채움, 제목 자동 생성, 설명 프리필. 카테고리 독립 필드 제거. 필드 순서 재배치.

- [ ] **Step 1: 자동 채움 상태 관리**

기존 폼 state에 추가/변경:

```typescript
const [selectedItem, setSelectedItem] = useState<ItemSearchResult | null>(null);
const [titleAutoGenerated, setTitleAutoGenerated] = useState(true);
const [enhancementLevel, setEnhancementLevel] = useState<number>(0);
// 기존 문자열 기반 enhancementLevel을 number로 전환
// 제출 시 String(enhancementLevel) 또는 enhancementLevel || null로 변환
```

기존 독립 강화 수치 `<input type="number">` 필드를 **제거** — 강화 입력은 아이템 피커 내부 슬라이더로 대체됨.

참고: `categoryId`는 아이템 선택 시 자동 결정되므로, 기존 카테고리 드롭다운도 제거. 이전에는 부모 카테고리(예: `weapon`)를 전송했으나, 이제 자식 카테고리(예: `weapon_sword`)를 전송. 백엔드 필터 SQL이 양쪽 모두 처리하므로 호환성 문제 없음.

- [ ] **Step 2: onSelect 콜백 — 자동 채움 로직**

```typescript
function handleItemSelect(item: ItemSearchResult) {
  setSelectedItem(item);
  setCategoryId(item.categoryId);  // 카테고리 자동 채움
  setTitleAutoGenerated(true);     // 제목 자동 생성 리셋

  // 설명 프리필
  if (item.optionText) {
    setDescription(`[아이템 옵션]\n${item.optionText}\n\n`);
  }
}
```

- [ ] **Step 3: 제목 자동 생성**

```typescript
useEffect(() => {
  if (!titleAutoGenerated || !selectedItem) return;
  const enhStr = enhancementLevel ? ` +${enhancementLevel}` : "";
  const typeStr = listingType === "sell" ? "판매" : "구매";
  setTitle(`${selectedItem.name}${enhStr} ${typeStr}합니다`);
}, [selectedItem, enhancementLevel, listingType, titleAutoGenerated]);
```

제목 입력 onChange에서:

```typescript
onChange={(e) => {
  setTitle(e.target.value);
  setTitleAutoGenerated(false);  // 수동 편집 시 자동 생성 끔
}}
```

- [ ] **Step 4: 폼 필드 순서 재배치 + 카테고리 독립 필드 제거**

기존 순서를 변경:

1. 거래 유형 (판매/구매 토글)
2. **아이템 선택** (ItemAutocomplete — 카테고리 탭 + 검색 + 강화 내장)
3. 서버 (드롭다운)
4. 가격 (타입 + 금액)
5. 제목 (자동 생성됨, 편집 가능)
6. 설명 (옵션 프리필, 편집 가능)
7. 이미지 (업로드)
8. 거래 방식 (드롭다운)

카테고리 드롭다운 (`<select>`) 제거 — `categoryId`는 아이템 선택 시 자동 결정.

- [ ] **Step 5: ItemAutocomplete에 새 props 연결**

```tsx
<ItemAutocomplete
  value={itemName}
  onChange={setItemName}
  onSelect={handleItemSelect}
  onEnhancementChange={setEnhancementLevel}
  enhancementLevel={enhancementLevel}
  required
/>
```

기존 `categoryId` prop 제거.

- [ ] **Step 6: 제출 시 categoryId 검증**

카테고리가 아이템 선택으로 자동 결정되므로, 아이템 미선택 시 제출 차단:

```typescript
if (!categoryId) {
  addToast("error", "아이템을 선택해주세요");
  return;
}
```

- [ ] **Step 7: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 8: 커밋**

```
feat(web): 매물 등록 자동 채움 — 아이템 선택 시 카테고리/제목/설명 자동 생성
```

---

## Task 5: 최종 검증 + 배포

**Files:** 없음 (검증만)

- [ ] **Step 1: Go 전체 테스트**

Run: `cd backend && go test ./...`

- [ ] **Step 2: Vitest 전체 테스트**

Run: `cd web && npx vitest run`

- [ ] **Step 3: Next.js 빌드**

Run: `cd web && npx next build`

- [ ] **Step 4: E2E 테스트**

Run: `cd web && npx playwright test`

- [ ] **Step 5: 커밋 + Push**

Run: `git push origin main`

- [ ] **Step 6: NAS 배포**

Run: `bash deploy/deploy.sh`

- [ ] **Step 7: E2E 재검증 (배포된 서버)**

Run: `cd web && npx playwright test`
