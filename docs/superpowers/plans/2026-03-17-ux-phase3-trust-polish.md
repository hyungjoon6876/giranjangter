# UX Phase 3: 신뢰 시스템 + 완성도 Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build trust/safety features (user blocking, share), add PWA support, and polish the app with footer/about info. Kakao/Naver OAuth는 백엔드 변경이 필요하므로 이 Phase에서는 프론트엔드 준비만 한다.

**Architecture:** Frontend-only changes in `web/`. Backend block API already exists (`POST/DELETE /users/:userId/block`). PWA requires `manifest.json` + meta tags. Share uses Web Share API with clipboard fallback.

**Tech Stack:** Next.js 16, React 19, TanStack Query, TailwindCSS 3.4, Web Share API

**Spec:** `docs/ux-analysis-2026-03-17.md` (Phase 3 features)

---

## File Structure

| File | Responsibility | Action |
|------|---------------|--------|
| `web/lib/api-client.ts` | `blockUser`, `unblockUser` 메서드 | Modify |
| `web/lib/hooks/use-users.ts` | `useBlockUser`, `useUnblockUser` 훅 | Create |
| `web/components/forms/block-confirm-modal.tsx` | 차단 확인 모달 | Create |
| `web/app/listings/[id]/page.tsx` | 공유 버튼 + 차단 버튼 추가 | Modify |
| `web/components/layout/footer.tsx` | 푸터 컴포넌트 | Create |
| `web/components/layout/responsive-shell.tsx` | 푸터 통합 | Modify |
| `web/public/manifest.json` | PWA manifest | Create |
| `web/app/layout.tsx` | PWA meta tags + manifest link | Modify |

---

## Chunk 1: 사용자 차단

### Task 1: 차단 API + 훅

**Files:**
- Modify: `web/lib/api-client.ts`
- Create: `web/lib/hooks/use-users.ts`

- [ ] **Step 1: Add block/unblock API methods**

`web/lib/api-client.ts`의 `ApiClient`에 추가 (Report 섹션 뒤):

```typescript
// Block
async blockUser(userId: string): Promise<void> {
  return this.fetch(`/users/${userId}/block`, { method: "POST" });
}

async unblockUser(userId: string): Promise<void> {
  return this.fetch(`/users/${userId}/block`, { method: "DELETE" });
}
```

- [ ] **Step 2: Create hooks**

```typescript
// web/lib/hooks/use-users.ts
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export function useBlockUser() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (userId: string) => apiClient.blockUser(userId),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useUnblockUser() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (userId: string) => apiClient.unblockUser(userId),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}
```

- [ ] **Step 3: Commit**

```bash
git add web/lib/api-client.ts web/lib/hooks/use-users.ts
git commit -m "feat(web): 사용자 차단/차단해제 API + 훅"
```

---

### Task 2: 차단 확인 모달

차단 전 확인하는 모달 컴포넌트.

**Files:**
- Create: `web/components/forms/block-confirm-modal.tsx`

- [ ] **Step 1: Create BlockConfirmModal**

```typescript
// web/components/forms/block-confirm-modal.tsx
"use client";

import { Modal } from "@/components/ui/modal";
import { useBlockUser } from "@/lib/hooks/use-users";
import { useToast } from "@/lib/hooks/use-toast";

interface BlockConfirmModalProps {
  open: boolean;
  onClose: () => void;
  userId: string;
  nickname: string;
}

export function BlockConfirmModal({
  open,
  onClose,
  userId,
  nickname,
}: BlockConfirmModalProps) {
  const blockUser = useBlockUser();
  const { addToast } = useToast();

  const handleBlock = async () => {
    try {
      await blockUser.mutateAsync(userId);
      addToast("success", `${nickname}님을 차단했습니다`);
      onClose();
    } catch {
      addToast("error", "차단에 실패했습니다");
    }
  };

  return (
    <Modal open={open} onClose={onClose} title="사용자 차단">
      <p className="text-text-secondary text-sm mb-4">
        <span className="font-medium text-text-primary">{nickname}</span>님을
        차단하시겠습니까?
      </p>
      <p className="text-text-dim text-xs mb-6">
        차단하면 해당 사용자의 매물과 채팅이 더 이상 표시되지 않습니다.
        프로필 설정에서 차단을 해제할 수 있습니다.
      </p>
      <div className="flex gap-3">
        <button
          onClick={onClose}
          className="flex-1 py-2.5 border border-border rounded-lg text-text-secondary hover:bg-medium transition-colors"
        >
          취소
        </button>
        <button
          onClick={handleBlock}
          disabled={blockUser.isPending}
          className="flex-1 py-2.5 bg-danger text-white rounded-lg font-medium disabled:opacity-50"
        >
          {blockUser.isPending ? "처리 중..." : "차단"}
        </button>
      </div>
    </Modal>
  );
}
```

- [ ] **Step 2: Commit**

```bash
git add web/components/forms/block-confirm-modal.tsx
git commit -m "feat(web): 사용자 차단 확인 모달"
```

---

### Task 3: 매물 상세에 차단 + 공유 버튼 추가

**Files:**
- Modify: `web/app/listings/[id]/page.tsx`

- [ ] **Step 1: Add imports and state**

`web/app/listings/[id]/page.tsx` 상단에 import 추가:

```typescript
import { BlockConfirmModal } from "@/components/forms/block-confirm-modal";
```

`reportOpen` state 옆에 추가:

```typescript
const [blockOpen, setBlockOpen] = useState(false);
```

- [ ] **Step 2: Add share handler**

`handleChat` 함수 아래에 추가:

```typescript
const handleShare = async () => {
  const shareData = {
    title: l.title,
    text: `${l.itemName} - ${formatPrice(l.priceAmount)}원`,
    url: window.location.href,
  };
  if (navigator.share) {
    try {
      await navigator.share(shareData);
    } catch {
      // User cancelled share — ignore
    }
  } else {
    await navigator.clipboard.writeText(window.location.href);
    addToast("success", "링크가 복사되었습니다");
  }
};
```

- [ ] **Step 3: Add share button to action bar**

신고 버튼 앞에 공유 버튼 추가:

```typescript
<button
  onClick={handleShare}
  className="p-3 border border-border rounded-lg hover:bg-medium transition-colors text-sm text-text-secondary"
  aria-label="공유"
>
  공유
</button>
```

- [ ] **Step 4: Add block button**

신고 버튼 뒤에 차단 버튼 추가 (다른 사용자의 매물일 때만):

```typescript
{!l.isOwner && (
  <button
    onClick={() => {
      if (!requireAuth("차단")) return;
      setBlockOpen(true);
    }}
    className="p-3 border border-border rounded-lg hover:bg-medium transition-colors text-sm text-text-dim"
  >
    차단
  </button>
)}
```

- [ ] **Step 5: Add BlockConfirmModal render**

`ReportModal` 뒤에 추가:

```typescript
{l.author && (
  <BlockConfirmModal
    open={blockOpen}
    onClose={() => setBlockOpen(false)}
    userId={l.author.userId}
    nickname={l.author.nickname}
  />
)}
```

- [ ] **Step 6: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 7: Commit**

```bash
git add web/app/listings/[id]/page.tsx
git commit -m "feat(web): 매물 상세에 공유(Web Share API) + 차단 버튼"
```

---

## Chunk 2: 푸터 + PWA

### Task 4: 푸터 컴포넌트

데스크톱에서 보이는 간결한 푸터. 앱 정보, 약관 링크, 문의 안내.

**Files:**
- Create: `web/components/layout/footer.tsx`

- [ ] **Step 1: Create Footer component**

```typescript
// web/components/layout/footer.tsx
import Link from "next/link";

export function Footer() {
  return (
    <footer className="hidden lg:block border-t border-border bg-dark mt-auto">
      <div className="max-w-6xl mx-auto px-6 py-8">
        <div className="flex flex-col lg:flex-row lg:items-start lg:justify-between gap-6">
          {/* Brand */}
          <div>
            <img src="/logo.png" alt="기란JT" className="h-8 mb-2" />
            <p className="text-text-dim text-sm">
              리니지 클래식 아이템 거래 중개 플랫폼
            </p>
            <p className="text-text-dim text-xs mt-1">무료 · 커뮤니티 기반</p>
          </div>

          {/* Links */}
          <div className="flex gap-8">
            <div>
              <h3 className="text-text-secondary text-xs font-semibold uppercase tracking-wider mb-2">
                서비스
              </h3>
              <ul className="space-y-1.5">
                <li>
                  <Link
                    href="/"
                    className="text-text-dim text-sm hover:text-gold transition-colors"
                  >
                    매물 보기
                  </Link>
                </li>
                <li>
                  <Link
                    href="/create"
                    className="text-text-dim text-sm hover:text-gold transition-colors"
                  >
                    매물 등록
                  </Link>
                </li>
              </ul>
            </div>
            <div>
              <h3 className="text-text-secondary text-xs font-semibold uppercase tracking-wider mb-2">
                정보
              </h3>
              <ul className="space-y-1.5">
                <li>
                  <span className="text-text-dim text-sm">
                    문의: giranjt@gmail.com
                  </span>
                </li>
              </ul>
            </div>
          </div>
        </div>

        <div className="border-t border-border mt-6 pt-4 text-center">
          <p className="text-text-dim text-xs">
            &copy; {new Date().getFullYear()} 기란JT. 이 서비스는 리니지 클래식의 공식 서비스가 아닙니다.
          </p>
        </div>
      </div>
    </footer>
  );
}
```

- [ ] **Step 2: Integrate footer into shell**

`web/components/layout/responsive-shell.tsx`를 수정:

```typescript
// Before:
import { Header } from "./header";
import { BottomNav } from "./bottom-nav";

// After:
import { Header } from "./header";
import { BottomNav } from "./bottom-nav";
import { Footer } from "./footer";
```

그리고 `</main>` 뒤에 `<Footer />` 추가:

```typescript
// Before:
<main id="main-content" className="pb-16 lg:pb-0">
  {children}
</main>
<BottomNav />

// After:
<main id="main-content" className="pb-16 lg:pb-0">
  {children}
</main>
<Footer />
<BottomNav />
```

- [ ] **Step 3: Run tests**

```bash
cd web && npx vitest run
```

- [ ] **Step 4: Commit**

```bash
git add web/components/layout/footer.tsx web/components/layout/responsive-shell.tsx
git commit -m "feat(web): 데스크톱 푸터 — 앱 정보, 링크, 저작권"
```

---

### Task 5: PWA Manifest + Meta Tags

홈 화면 추가를 지원하는 PWA 기본 설정.

**Files:**
- Create: `web/public/manifest.json`
- Modify: `web/app/layout.tsx`

- [ ] **Step 1: Create manifest.json**

```json
{
  "name": "기란JT — 리니지 클래식 거래",
  "short_name": "기란JT",
  "description": "리니지 클래식 아이템 거래 중개 플랫폼",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#08080C",
  "theme_color": "#EBD5C4",
  "icons": [
    {
      "src": "/icons/icon-192.png",
      "sizes": "192x192",
      "type": "image/png"
    },
    {
      "src": "/icons/icon-512.png",
      "sizes": "512x512",
      "type": "image/png"
    }
  ]
}
```

- [ ] **Step 2: Add PWA meta tags to layout**

`web/app/layout.tsx`의 `<head>` 안에 추가:

```typescript
<link rel="manifest" href="/manifest.json" />
<meta name="theme-color" content="#EBD5C4" />
<meta name="apple-mobile-web-app-capable" content="yes" />
<meta name="apple-mobile-web-app-status-bar-style" content="black-translucent" />
<link rel="apple-touch-icon" href="/icons/icon-192.png" />
```

- [ ] **Step 3: Commit**

```bash
git add web/public/manifest.json web/app/layout.tsx
git commit -m "feat(web): PWA manifest + 홈 화면 추가 지원"
```

---

## Summary

| Task | 파일 | 기능 | 복잡도 |
|------|------|------|--------|
| 1 | api-client.ts, use-users.ts | 차단 API + 훅 | S |
| 2 | block-confirm-modal.tsx | 차단 확인 모달 | S |
| 3 | listings/[id]/page.tsx | 공유 + 차단 버튼 | M |
| 4 | footer.tsx, responsive-shell.tsx | 데스크톱 푸터 | S |
| 5 | manifest.json, layout.tsx | PWA 지원 | S |

**Dependencies:**
- Task 1 → Task 2 → Task 3 (순차)
- Task 4, 5 (독립, 병렬 가능)

**Out of scope (백엔드 필요):**
- 카카오/네이버 OAuth — 백엔드에 `internal/oauth/kakao.go`, `naver.go` 추가 필요
- 차단 목록 관리 페이지 — 차단 해제 UI (향후)

---

## 전체 Phase 실행 순서

1. **Phase 1** ✅ 완료 (커밋 필요)
2. **Phase 2** — 12 tasks, ~30분
3. **Phase 3** — 5 tasks, ~15분
