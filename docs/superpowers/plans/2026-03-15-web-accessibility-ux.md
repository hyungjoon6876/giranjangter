# 웹 접근성 + UX 개선 구현 계획

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 기란장터 웹 프론트엔드의 WCAG AA 접근성과 첫 진입 UX를 사용자 흐름 순서대로 개선한다.

**Architecture:** Next.js 16 App Router 기반 web/ 디렉토리만 수정. 신규 UI 컴포넌트 4개(skeleton, toast, error-state, 접근성 유틸리티) 생성 후, 기존 22개 파일에 ARIA/시맨틱/키보드/에러 처리를 적용. 백엔드 변경 없음.

**Tech Stack:** Next.js 16, React 19, TailwindCSS 3.4, Vitest 4, @testing-library/react

**Spec:** `docs/superpowers/specs/2026-03-15-web-accessibility-ux-design.md`

---

## Chunk 1: Foundation — CSS 유틸리티 + 공통 컴포넌트

이 청크에서 모든 후속 작업이 의존하는 기반 컴포넌트를 만든다.

---

### Task 1: 접근성 CSS 유틸리티 추가

**Files:**
- Modify: `web/app/globals.css`

- [ ] **Step 1: sr-only, not-sr-only, skeleton, focus-gold 클래스 추가**

`globals.css`의 `@layer utilities` 블록 끝에 추가:

```css
/* Accessibility utilities */
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}

.not-sr-only {
  position: static;
  width: auto;
  height: auto;
  padding: 0;
  margin: 0;
  overflow: visible;
  clip: auto;
  white-space: normal;
}

.skeleton {
  background: linear-gradient(90deg, #1e2538 25%, #2a3045 50%, #1e2538 75%);
  background-size: 200% 100%;
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}

@keyframes skeleton-pulse {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

.focus-gold {
  outline: 2px solid #c4a35a;
  outline-offset: 2px;
}

/* Reduced motion */
@media (prefers-reduced-motion: reduce) {
  .skeleton {
    animation: none;
    background: #1e2538;
  }
  .animate-spin {
    animation: none;
  }
  * {
    transition-duration: 0.01ms !important;
    animation-duration: 0.01ms !important;
  }
}
```

- [ ] **Step 2: 빌드 확인**

Run: `cd web && npx next build 2>&1 | tail -5`
Expected: Build succeeds

- [ ] **Step 3: Commit**

```bash
git add web/app/globals.css
git commit -m "feat(a11y): add sr-only, skeleton, focus-gold, reduced-motion CSS utilities"
```

---

### Task 2: Toast 시스템 구현

**Files:**
- Create: `web/lib/hooks/use-toast.ts`
- Create: `web/components/ui/toast.tsx`
- Modify: `web/lib/providers.tsx`
- Create: `web/__tests__/components/toast.test.tsx`

- [ ] **Step 1: useToast 훅 작성**

```typescript
// web/lib/hooks/use-toast.ts
"use client";

import { createContext, useContext, useCallback, useState } from "react";

export type ToastType = "success" | "error" | "info";

export interface Toast {
  id: string;
  type: ToastType;
  title: string;
  description?: string;
}

interface ToastContextValue {
  toasts: Toast[];
  addToast: (type: ToastType, title: string, description?: string) => void;
  removeToast: (id: string) => void;
}

export const ToastContext = createContext<ToastContextValue>({
  toasts: [],
  addToast: () => {},
  removeToast: () => {},
});

export function useToast() {
  return useContext(ToastContext);
}

export function useToastState() {
  const [toasts, setToasts] = useState<Toast[]>([]);

  const addToast = useCallback(
    (type: ToastType, title: string, description?: string) => {
      const id = `toast-${Date.now()}`;
      setToasts((prev) => [...prev, { id, type, title, description }]);
      setTimeout(() => {
        setToasts((prev) => prev.filter((t) => t.id !== id));
      }, 5000);
    },
    []
  );

  const removeToast = useCallback((id: string) => {
    setToasts((prev) => prev.filter((t) => t.id !== id));
  }, []);

  return { toasts, addToast, removeToast };
}
```

- [ ] **Step 2: Toast 컴포넌트 작성**

```tsx
// web/components/ui/toast.tsx
"use client";

import { useToast, type Toast as ToastItem } from "@/lib/hooks/use-toast";

const borderColors = {
  success: "border-l-[#2ecc71]",
  error: "border-l-[#e74c3c]",
  info: "border-l-gold",
};

function ToastCard({ toast }: { toast: ToastItem }) {
  const { removeToast } = useToast();

  return (
    <div
      role="alert"
      className={`bg-card border-l-4 ${borderColors[toast.type]} rounded-lg p-3 pr-8 shadow-lg relative max-w-sm w-full`}
    >
      <p className="text-sm font-semibold text-text-primary">{toast.title}</p>
      {toast.description && (
        <p className="text-xs text-text-secondary mt-1">{toast.description}</p>
      )}
      <button
        onClick={() => removeToast(toast.id)}
        aria-label="닫기"
        className="absolute top-2 right-2 text-text-dim hover:text-text-primary text-lg leading-none"
      >
        ×
      </button>
    </div>
  );
}

export function ToastContainer() {
  const { toasts } = useToast();

  if (toasts.length === 0) return null;

  return (
    <div className="fixed top-4 right-4 z-[60] flex flex-col gap-2 max-sm:left-4 max-sm:right-4">
      {toasts.map((toast) => (
        <ToastCard key={toast.id} toast={toast} />
      ))}
    </div>
  );
}
```

- [ ] **Step 3: providers.tsx에 ToastProvider 추가**

`web/lib/providers.tsx` 수정 — QueryClientProvider 내부에 ToastContext.Provider 추가:

```tsx
// providers.tsx에 import 추가
import { ToastContext, useToastState } from "@/lib/hooks/use-toast";
import { ToastContainer } from "@/components/ui/toast";

// return 문 내부 수정:
return (
  <QueryClientProvider client={queryClient}>
    <ToastProviderInner>
      <SSEInitializer />
      {children}
    </ToastProviderInner>
  </QueryClientProvider>
);

// 파일 하단에 내부 컴포넌트 추가:
function ToastProviderInner({ children }: { children: React.ReactNode }) {
  const state = useToastState();
  return (
    <ToastContext.Provider value={state}>
      {children}
      <ToastContainer />
    </ToastContext.Provider>
  );
}
```

- [ ] **Step 4: Toast 테스트 작성**

```tsx
// web/__tests__/components/toast.test.tsx
import { render, screen, act } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { ToastContext, type Toast } from "@/lib/hooks/use-toast";
import { ToastContainer } from "@/components/ui/toast";

function renderWithToast(toasts: Toast[], removeToast = vi.fn()) {
  return render(
    <ToastContext.Provider value={{ toasts, addToast: vi.fn(), removeToast }}>
      <ToastContainer />
    </ToastContext.Provider>
  );
}

describe("ToastContainer", () => {
  it("renders nothing when no toasts", () => {
    renderWithToast([]);
    expect(screen.queryByRole("alert")).toBeNull();
  });

  it("renders toast with role=alert", () => {
    renderWithToast([{ id: "1", type: "error", title: "에러 발생" }]);
    expect(screen.getByRole("alert")).toHaveTextContent("에러 발생");
  });

  it("close button has aria-label", () => {
    renderWithToast([{ id: "1", type: "info", title: "알림" }]);
    expect(screen.getByLabelText("닫기")).toBeInTheDocument();
  });
});
```

- [ ] **Step 5: 테스트 실행**

Run: `cd web && npx vitest run __tests__/components/toast.test.tsx`
Expected: All 3 tests pass

- [ ] **Step 6: Commit**

```bash
git add web/lib/hooks/use-toast.ts web/components/ui/toast.tsx web/lib/providers.tsx web/__tests__/components/toast.test.tsx
git commit -m "feat(a11y): add Toast system with role=alert and auto-dismiss"
```

---

### Task 3: Loading 컴포넌트 접근성 개선

**Files:**
- Modify: `web/components/ui/loading.tsx`
- Create: `web/__tests__/components/loading.test.tsx`

- [ ] **Step 1: 테스트 작성**

```tsx
// web/__tests__/components/loading.test.tsx
import { render, screen } from "@testing-library/react";
import { Loading } from "@/components/ui/loading";

describe("Loading", () => {
  it("has role=status", () => {
    render(<Loading />);
    expect(screen.getByRole("status")).toBeInTheDocument();
  });

  it("has accessible label", () => {
    render(<Loading />);
    expect(screen.getByLabelText("로딩 중")).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: 테스트 실패 확인**

Run: `cd web && npx vitest run __tests__/components/loading.test.tsx`
Expected: FAIL (no role="status")

- [ ] **Step 3: Loading 컴포넌트 수정**

```tsx
// web/components/ui/loading.tsx
export function Loading() {
  return (
    <div className="flex justify-center py-20" role="status" aria-label="로딩 중">
      <div className="w-8 h-8 border-4 border-gold border-t-transparent rounded-full animate-spin" />
      <span className="sr-only">로딩 중</span>
    </div>
  );
}
```

- [ ] **Step 4: 테스트 통과 확인**

Run: `cd web && npx vitest run __tests__/components/loading.test.tsx`
Expected: All 2 tests pass

- [ ] **Step 5: Commit**

```bash
git add web/components/ui/loading.tsx web/__tests__/components/loading.test.tsx
git commit -m "feat(a11y): add role=status and aria-label to Loading component"
```

---

### Task 4: Skeleton 컴포넌트 생성

**Files:**
- Create: `web/components/ui/skeleton.tsx`
- Create: `web/__tests__/components/skeleton.test.tsx`

- [ ] **Step 1: 테스트 작성**

```tsx
// web/__tests__/components/skeleton.test.tsx
import { render, screen } from "@testing-library/react";
import { ListingSkeleton } from "@/components/ui/skeleton";

describe("ListingSkeleton", () => {
  it("has role=status with accessible label", () => {
    render(<ListingSkeleton />);
    expect(screen.getByRole("status")).toBeInTheDocument();
    expect(screen.getByLabelText("매물 목록을 불러오는 중")).toBeInTheDocument();
  });

  it("renders multiple skeleton cards", () => {
    const { container } = render(<ListingSkeleton count={3} />);
    const cards = container.querySelectorAll("[data-testid='skeleton-card']");
    expect(cards).toHaveLength(3);
  });
});
```

- [ ] **Step 2: Skeleton 컴포넌트 구현**

```tsx
// web/components/ui/skeleton.tsx
export function SkeletonCard() {
  return (
    <div data-testid="skeleton-card" className="bg-card rounded-xl p-4 border border-border">
      <div className="flex gap-3">
        <div className="w-10 h-10 rounded-lg skeleton" />
        <div className="flex-1 space-y-2">
          <div className="h-4 w-3/5 rounded skeleton" />
          <div className="h-3 w-2/5 rounded skeleton" />
          <div className="h-5 w-1/4 rounded skeleton" />
        </div>
      </div>
    </div>
  );
}

export function ListingSkeleton({ count = 5 }: { count?: number }) {
  return (
    <div role="status" aria-label="매물 목록을 불러오는 중" aria-busy="true">
      <div className="flex gap-2 mb-4">
        {[1, 2, 3].map((i) => (
          <div key={i} className="h-8 w-16 rounded-full skeleton" />
        ))}
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
        {Array.from({ length: count }, (_, i) => (
          <SkeletonCard key={i} />
        ))}
      </div>
      <span className="sr-only">매물 목록을 불러오는 중</span>
    </div>
  );
}
```

- [ ] **Step 3: 테스트 통과 확인**

Run: `cd web && npx vitest run __tests__/components/skeleton.test.tsx`
Expected: All 2 tests pass

- [ ] **Step 4: Commit**

```bash
git add web/components/ui/skeleton.tsx web/__tests__/components/skeleton.test.tsx
git commit -m "feat(a11y): add Skeleton loading component with aria-label"
```

---

### Task 5: ErrorState 컴포넌트 생성

**Files:**
- Create: `web/components/ui/error-state.tsx`
- Create: `web/__tests__/components/error-state.test.tsx`

- [ ] **Step 1: 테스트 작성**

```tsx
// web/__tests__/components/error-state.test.tsx
import { render, screen } from "@testing-library/react";
import { ErrorState } from "@/components/ui/error-state";

describe("ErrorState", () => {
  it("has role=alert", () => {
    render(<ErrorState message="에러 발생" onRetry={() => {}} />);
    expect(screen.getByRole("alert")).toBeInTheDocument();
  });

  it("shows retry button", () => {
    const onRetry = vi.fn();
    render(<ErrorState message="에러 발생" onRetry={onRetry} />);
    screen.getByText("다시 시도").click();
    expect(onRetry).toHaveBeenCalled();
  });

  it("shows custom message", () => {
    render(<ErrorState message="매물을 불러올 수 없습니다" onRetry={() => {}} />);
    expect(screen.getByText("매물을 불러올 수 없습니다")).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: ErrorState 구현**

```tsx
// web/components/ui/error-state.tsx
"use client";

import { useEffect, useRef } from "react";

interface ErrorStateProps {
  message: string;
  description?: string;
  onRetry: () => void;
}

export function ErrorState({ message, description, onRetry }: ErrorStateProps) {
  const retryRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    retryRef.current?.focus();
  }, []);

  return (
    <div role="alert" className="flex flex-col items-center justify-center py-20 px-4 text-center">
      <div className="text-4xl mb-4" role="img" aria-label="경고">⚠️</div>
      <p className="text-base font-semibold text-text-primary mb-2">{message}</p>
      {description && (
        <p className="text-sm text-text-secondary mb-6">{description}</p>
      )}
      <button
        ref={retryRef}
        onClick={onRetry}
        className="btn-gold-gradient text-darkest font-bold px-5 py-2 rounded-lg text-sm"
      >
        다시 시도
      </button>
    </div>
  );
}
```

- [ ] **Step 3: 테스트 통과 확인**

Run: `cd web && npx vitest run __tests__/components/error-state.test.tsx`
Expected: All 3 tests pass

- [ ] **Step 4: Commit**

```bash
git add web/components/ui/error-state.tsx web/__tests__/components/error-state.test.tsx
git commit -m "feat(a11y): add ErrorState component with role=alert and auto-focus retry"
```

---

### Task 6: EmptyState 접근성 개선

**Files:**
- Modify: `web/components/ui/empty-state.tsx`

- [ ] **Step 1: EmptyState 시맨틱 개선**

현재 `<p>` 태그를 `<h2>`로, emoji에 `role="img"` 추가:

```tsx
// web/components/ui/empty-state.tsx — 전체 교체
import Link from "next/link";

interface EmptyStateProps {
  icon: string;
  title: string;
  description?: string;
  actionLabel?: string;
  actionHref?: string;
}

export function EmptyState({ icon = "🔍", title, description, actionLabel, actionHref }: EmptyStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-20 px-4 text-center">
      <div className="text-5xl mb-4" role="img" aria-label={title}>{icon}</div>
      <h2 className="text-lg font-semibold text-text-primary mb-2">{title}</h2>
      {description && (
        <p className="text-sm text-text-secondary mb-4">{description}</p>
      )}
      {actionLabel && actionHref && (
        <Link
          href={actionHref}
          className="btn-gold-gradient text-darkest font-bold px-5 py-2 rounded-lg text-sm inline-block"
          aria-label={actionLabel}
        >
          {actionLabel}
        </Link>
      )}
    </div>
  );
}
```

- [ ] **Step 2: 기존 테스트 통과 확인**

Run: `cd web && npx vitest run`
Expected: All tests pass

- [ ] **Step 3: Commit**

```bash
git add web/components/ui/empty-state.tsx
git commit -m "feat(a11y): improve EmptyState with h2 heading and role=img on icon"
```

---

## Chunk 2: 모달 접근성 (CRITICAL)

가장 심각한 접근성 위반인 모달을 전면 개선한다.

---

### Task 7: Modal 전면 개선 — Portal + Focus Trap + ARIA

**Files:**
- Modify: `web/components/ui/modal.tsx`
- Create: `web/__tests__/components/modal.test.tsx`

- [ ] **Step 1: 모달 테스트 작성**

```tsx
// web/__tests__/components/modal.test.tsx
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { Modal } from "@/components/ui/modal";

describe("Modal accessibility", () => {
  it("has role=dialog and aria-modal", () => {
    render(<Modal open onClose={() => {}} title="테스트"><p>내용</p></Modal>);
    const dialog = screen.getByRole("dialog");
    expect(dialog).toHaveAttribute("aria-modal", "true");
  });

  it("has aria-labelledby linked to title", () => {
    render(<Modal open onClose={() => {}} title="제목"><p>내용</p></Modal>);
    const dialog = screen.getByRole("dialog");
    const titleId = dialog.getAttribute("aria-labelledby");
    expect(titleId).toBeTruthy();
    expect(screen.getByText("제목").id).toBe(titleId);
  });

  it("close button has aria-label", () => {
    render(<Modal open onClose={() => {}} title="제목"><p>내용</p></Modal>);
    expect(screen.getByLabelText("닫기")).toBeInTheDocument();
  });

  it("calls onClose on Escape key", async () => {
    const onClose = vi.fn();
    render(<Modal open onClose={onClose} title="제목"><p>내용</p></Modal>);
    await userEvent.keyboard("{Escape}");
    expect(onClose).toHaveBeenCalled();
  });

  it("does not render when closed", () => {
    render(<Modal open={false} onClose={() => {}} title="제목"><p>내용</p></Modal>);
    expect(screen.queryByRole("dialog")).toBeNull();
  });
});
```

- [ ] **Step 2: 테스트 실패 확인**

Run: `cd web && npx vitest run __tests__/components/modal.test.tsx`
Expected: FAIL (no role="dialog")

- [ ] **Step 3: Modal 전면 재작성**

```tsx
// web/components/ui/modal.tsx — 전체 교체
"use client";

import { useEffect, useRef, useId, type ReactNode } from "react";
import { createPortal } from "react-dom";

interface ModalProps {
  open: boolean;
  onClose: () => void;
  title: string;
  children: ReactNode;
}

export function Modal({ open, onClose, title, children }: ModalProps) {
  const titleId = useId();
  const modalRef = useRef<HTMLDivElement>(null);
  const previousFocusRef = useRef<HTMLElement | null>(null);

  useEffect(() => {
    if (!open) return;

    previousFocusRef.current = document.activeElement as HTMLElement;
    document.body.style.overflow = "hidden";

    // Apply inert to siblings
    const siblings = Array.from(document.body.children).filter(
      (el) => el.id !== "modal-portal"
    );
    siblings.forEach((el) => el.setAttribute("inert", ""));

    // Focus first focusable element
    requestAnimationFrame(() => {
      const focusable = modalRef.current?.querySelector<HTMLElement>(
        'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
      );
      focusable?.focus();
    });

    return () => {
      document.body.style.overflow = "";
      siblings.forEach((el) => el.removeAttribute("inert"));
      previousFocusRef.current?.focus();
    };
  }, [open]);

  // Escape key handler
  useEffect(() => {
    if (!open) return;
    const handler = (e: KeyboardEvent) => {
      if (e.key === "Escape") onClose();
    };
    document.addEventListener("keydown", handler);
    return () => document.removeEventListener("keydown", handler);
  }, [open, onClose]);

  // Focus trap
  useEffect(() => {
    if (!open) return;
    const handler = (e: KeyboardEvent) => {
      if (e.key !== "Tab" || !modalRef.current) return;
      const focusable = modalRef.current.querySelectorAll<HTMLElement>(
        'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
      );
      if (focusable.length === 0) return;
      const first = focusable[0];
      const last = focusable[focusable.length - 1];

      if (e.shiftKey && document.activeElement === first) {
        e.preventDefault();
        last.focus();
      } else if (!e.shiftKey && document.activeElement === last) {
        e.preventDefault();
        first.focus();
      }
    };
    document.addEventListener("keydown", handler);
    return () => document.removeEventListener("keydown", handler);
  }, [open]);

  if (!open || typeof window === "undefined") return null;

  // Ensure portal container exists
  let portal = document.getElementById("modal-portal");
  if (!portal) {
    portal = document.createElement("div");
    portal.id = "modal-portal";
    document.body.appendChild(portal);
  }

  return createPortal(
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div
        className="fixed inset-0 bg-black/60"
        onClick={onClose}
        aria-hidden="true"
      />
      <div
        ref={modalRef}
        role="dialog"
        aria-modal="true"
        aria-labelledby={titleId}
        className="relative bg-[#1a2035] border border-border rounded-xl w-full max-w-md max-h-[85vh] overflow-y-auto shadow-xl"
      >
        <div className="flex items-center justify-between p-4 border-b border-border">
          <h2 id={titleId} className="text-lg font-bold text-text-primary">
            {title}
          </h2>
          <button
            onClick={onClose}
            aria-label="닫기"
            className="text-text-dim hover:text-text-primary text-xl leading-none p-1"
          >
            ×
          </button>
        </div>
        <div className="p-4">{children}</div>
      </div>
    </div>,
    portal
  );
}
```

- [ ] **Step 4: 테스트 통과 확인**

Run: `cd web && npx vitest run __tests__/components/modal.test.tsx`
Expected: All 5 tests pass

- [ ] **Step 5: 전체 테스트 통과 확인**

Run: `cd web && npx vitest run`
Expected: All tests pass (기존 모달 사용처도 props 호환)

- [ ] **Step 6: Commit**

```bash
git add web/components/ui/modal.tsx web/__tests__/components/modal.test.tsx
git commit -m "feat(a11y): rewrite Modal with Portal, focus trap, Escape, ARIA dialog"
```

---

## Chunk 3: 첫 진입 + 네비게이션

사용자가 처음 도착했을 때의 경험과 앱 전체 네비게이션 접근성.

---

### Task 8: Skip-to-Content 링크

**Files:**
- Modify: `web/components/layout/responsive-shell.tsx`

- [ ] **Step 1: skip link + main id 추가**

```tsx
// responsive-shell.tsx — 전체 교체
import { Header } from "./header";
import { BottomNav } from "./bottom-nav";

export function ResponsiveShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-darkest text-text-primary">
      <a
        href="#main-content"
        className="sr-only focus:not-sr-only focus:fixed focus:top-2 focus:left-2 focus:z-[70] focus:bg-gold focus:text-darkest focus:px-4 focus:py-2 focus:rounded-lg focus:font-bold focus:text-sm"
      >
        본문으로 건너뛰기
      </a>
      <Header />
      <main id="main-content" className="pb-16 lg:pb-0">
        {children}
      </main>
      <BottomNav />
    </div>
  );
}
```

- [ ] **Step 2: Commit**

```bash
git add web/components/layout/responsive-shell.tsx
git commit -m "feat(a11y): add skip-to-content link and main landmark id"
```

---

### Task 9: 헤더 접근성 + 로그인 버튼

**Files:**
- Modify: `web/components/layout/header.tsx`

- [ ] **Step 1: 헤더에 ARIA 속성 + 로그인 버튼 추가**

수정 사항:
- `<nav aria-label="메인 메뉴">` 래핑
- 활성 링크에 `aria-current="page"`
- 알림/프로필 아이콘에 `aria-label`
- SVG에 `aria-hidden="true"`
- 검색 input을 `<search>` 요소로 감싸기 + `type="search"` + `aria-label="매물 검색"`
- 비로그인 시 "로그인" 텍스트 버튼 표시

`header.tsx`의 네비게이션 링크 부분 수정:
- `<nav>` → `<nav aria-label="메인 메뉴">`
- 각 Link에 `aria-current={isActive ? "page" : undefined}`
- 알림 Link: `aria-label="알림"`, SVG에 `aria-hidden="true"`
- 프로필 Link: `aria-label="내 프로필"`, SVG에 `aria-hidden="true"`
- 검색 input 래핑: `<search>` + `type="search"` + `aria-label="매물 검색"`
- 프로필 영역에 비로그인 체크 추가 (apiClient.isLoggedIn 확인 후 "로그인" 링크 표시)

- [ ] **Step 2: 빌드 확인**

Run: `cd web && npx next build 2>&1 | tail -5`
Expected: Build succeeds

- [ ] **Step 3: Commit**

```bash
git add web/components/layout/header.tsx
git commit -m "feat(a11y): add ARIA to header — nav landmark, aria-current, icon labels, login button"
```

---

### Task 10: 하단 탭바 접근성

**Files:**
- Modify: `web/components/layout/bottom-nav.tsx`

- [ ] **Step 1: ARIA 속성 추가**

수정 사항:
- 외부 `<nav>` → `<nav aria-label="하단 메뉴">`
- 각 Link에 `aria-current={isActive ? "page" : undefined}`
- 각 Link에 `aria-label` (예: "마켓", "채팅", "매물 등록", "프로필")
- SVG에 `aria-hidden="true"`
- `focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gold focus-visible:rounded` 추가

- [ ] **Step 2: Commit**

```bash
git add web/components/layout/bottom-nav.tsx
git commit -m "feat(a11y): add ARIA to bottom nav — nav landmark, aria-current, focus-visible"
```

---

### Task 11: 홈페이지 스켈레톤 + 결과 카운트 + 정렬

**Files:**
- Modify: `web/app/page.tsx`

- [ ] **Step 1: Loading → ListingSkeleton 교체**

`app/page.tsx`에서:
- `import { Loading }` → `import { ListingSkeleton }` 교체
- `<Loading />` → `<ListingSkeleton />` 교체
- FAB 버튼에 `aria-label="매물 등록"` 추가
- 정렬 상태 추가: `const [sort, setSort] = useState("recent");`

- [ ] **Step 2: 결과 카운트 + 정렬 추가**

리스팅 그리드 위에 정보 바 추가:

```tsx
{/* 필터 아래, 그리드 위에 삽입 */}
<div className="flex items-center justify-between px-4 lg:px-6 py-2">
  <p className="text-sm text-text-secondary" aria-live="polite">
    <span className="font-semibold text-text-primary">{listings.length}</span>개 매물
  </p>
  <select
    aria-label="정렬 방식"
    value={sort}
    onChange={(e) => setSort(e.target.value)}
    className="bg-medium text-text-secondary text-xs border border-border rounded-md px-2 py-1"
  >
    <option value="recent">최신순</option>
    <option value="price_asc">가격 낮은순</option>
    <option value="price_desc">가격 높은순</option>
    <option value="popular">인기순</option>
  </select>
</div>
```

- [ ] **Step 3: 컨텍스트 빈 상태**

EmptyState 호출을 조건별로 분기:

```tsx
// 기존 EmptyState 호출 교체
{searchQuery ? (
  <EmptyState
    icon="🔍"
    title={`'${searchQuery}'에 대한 매물이 없습니다`}
    description="다른 검색어를 시도해보세요"
  />
) : selectedServer ? (
  <EmptyState
    icon="🗺️"
    title="선택한 서버에 매물이 없습니다"
    description="다른 서버를 선택해보세요"
  />
) : (
  <EmptyState
    icon="📦"
    title="아직 매물이 없습니다"
    description="첫 매물을 등록해보세요!"
    actionLabel="매물 등록"
    actionHref="/create"
  />
)}
```

- [ ] **Step 4: 빌드 확인**

Run: `cd web && npx next build 2>&1 | tail -5`
Expected: Build succeeds

- [ ] **Step 5: Commit**

```bash
git add web/app/page.tsx
git commit -m "feat(a11y): home page — skeleton loading, result count, sort, contextual empty states"
```

---

## Chunk 4: 검색 + 필터 + 카드 접근성

메인 탐색 컴포넌트들의 접근성 개선.

---

### Task 12: 필터 접근성 + 검색 랜드마크

**Files:**
- Modify: `web/components/listing/listing-filters.tsx`

- [ ] **Step 1: 필터에 ARIA 속성 추가**

수정 사항:
- 서버 버튼 그룹: `role="group"` + `aria-label="서버 필터"`
- 선택 버튼: `aria-pressed={isSelected}`
- 각 버튼에 `focus-visible:ring-2 focus-visible:ring-gold`
- 검색 input을 `<search>` 요소로 감싸기
- input `type="search"` + `aria-label="매물 검색"`

- [ ] **Step 2: Commit**

```bash
git add web/components/listing/listing-filters.tsx
git commit -m "feat(a11y): filter chips aria-pressed, search landmark, focus-visible"
```

---

### Task 13: 매물 카드 접근성

**Files:**
- Modify: `web/components/listing/listing-card.tsx`

- [ ] **Step 1: 카드에 ARIA 속성 추가**

수정 사항:
- Link에 `aria-label={`${listing.itemName} - ${formatPrice(listing.priceAmount)}원 - ${listing.serverName}`}`
- truncated 텍스트에 `title` 속성
- `focus-visible:ring-2 focus-visible:ring-gold focus-visible:ring-offset-2 focus-visible:ring-offset-darkest`
- 아이콘 이미지에 의미 있는 `alt={`${listing.itemName} 아이콘`}`

- [ ] **Step 2: Commit**

```bash
git add web/components/listing/listing-card.tsx
git commit -m "feat(a11y): listing card — aria-label, title, focus-visible, meaningful alt"
```

---

### Task 14: listing-info 시맨틱 구조

**Files:**
- Modify: `web/components/listing/listing-info.tsx`

- [ ] **Step 1: InfoRow를 dl/dt/dd로 변환**

현재 `<div><span><span>` → `<div className="flex ..."><dt>label</dt><dd>value</dd></div>` 변경.
부모에서 `<dl>` 래핑 필요 — `listings/[id]/page.tsx`에서 InfoRow 사용부를 `<dl>` 감싸기.

- [ ] **Step 2: Commit**

```bash
git add web/components/listing/listing-info.tsx web/app/listings/[id]/page.tsx
git commit -m "feat(a11y): listing info — dl/dt/dd semantic structure"
```

---

## Chunk 5: 상세 페이지 + 채팅 + alert() 제거

alert()를 전부 Toast로 교체하고, 채팅/상세 페이지의 접근성 개선.

---

### Task 15: 매물 상세 페이지 접근성 + alert() 제거

**Files:**
- Modify: `web/app/listings/[id]/page.tsx`

- [ ] **Step 1: alert() → useToast() 교체**

Line 33의 `alert("채팅을 시작할 수 없습니다")` 제거 후:

```tsx
const { addToast } = useToast();
// ...
catch (e) {
  addToast("error", "채팅을 시작할 수 없습니다", "네트워크 연결을 확인해주세요");
}
```

- [ ] **Step 2: 접근성 속성 추가**

- 하단 액션 바: `role="toolbar"` + `aria-label="매물 액션"`
- 찜 버튼: `aria-pressed={isFavorited}` + `aria-label={isFavorited ? "찜 취소" : "찜하기"}`
- 로딩 실패 시 `<ErrorState>` 사용

- [ ] **Step 3: Commit**

```bash
git add web/app/listings/[id]/page.tsx
git commit -m "feat(a11y): listing detail — toast errors, toolbar role, aria-pressed favorite"
```

---

### Task 16: 로그인 페이지 alert() 제거

**Files:**
- Modify: `web/app/login/page.tsx`

- [ ] **Step 1: alert() → 인라인 에러 메시지**

Line 17의 `alert(...)` 제거. 대신 상태 변수 `error`를 두고 인라인 표시:

```tsx
const [error, setError] = useState<string | null>(null);
// ...
catch (e) {
  setError(`로그인에 실패했습니다: ${e}`);
}
// JSX에서:
{error && (
  <p role="alert" className="text-sm text-[#e74c3c] text-center">{error}</p>
)}
```

- [ ] **Step 2: Commit**

```bash
git add web/app/login/page.tsx
git commit -m "feat(a11y): login page — replace alert() with inline error message"
```

---

### Task 17: 매물 등록 폼 접근성 + alert() 제거

**Files:**
- Modify: `web/app/create/page.tsx`

- [ ] **Step 1: alert() → useToast()**

Line 42의 `alert(...)` 제거 후 Toast 사용:

```tsx
const { addToast } = useToast();
// catch 블록:
addToast("error", "등록에 실패했습니다", "잠시 후 다시 시도해주세요");
```

- [ ] **Step 2: label htmlFor + input id 연결**

모든 `<label>` + `<input>`/`<select>` 쌍에 `htmlFor`/`id` 추가. 예시:

```tsx
<label htmlFor="item-name" className="...">아이템명 *</label>
<input id="item-name" aria-required="true" ... />
```

필수 필드 목록: `serverId`, `categoryId`, `itemName`, `title`, `description`, `priceType`, `tradeMethod`
— 각각에 `aria-required="true"` 추가.

- [ ] **Step 3: 에러 상태 접근성**

폼 제출 실패 시 에러 필드에 접근성 속성 추가:

```tsx
// 에러 상태 관리:
const [errors, setErrors] = useState<Record<string, string>>({});

// 각 input에:
<input
  id="item-name"
  aria-required="true"
  aria-invalid={!!errors.itemName}
  aria-describedby={errors.itemName ? "item-name-error" : undefined}
/>
{errors.itemName && (
  <p id="item-name-error" role="alert" className="text-xs text-[#e74c3c] mt-1">
    {errors.itemName}
  </p>
)}

// 제출 실패 시 첫 에러 필드로 포커스:
const firstErrorField = document.getElementById(Object.keys(newErrors)[0]);
firstErrorField?.focus();
```

- [ ] **Step 4: Commit**

```bash
git add web/app/create/page.tsx
git commit -m "feat(a11y): create listing — toast errors, label associations, aria-required, aria-invalid"
```

---

### Task 18: 채팅 컴포넌트 접근성

**Files:**
- Modify: `web/components/chat/chat-panel.tsx`
- Modify: `web/components/chat/chat-input.tsx`
- Modify: `web/components/chat/chat-message.tsx`
- Modify: `web/components/chat/chat-list-item.tsx`
- Modify: `web/app/chats/page.tsx`
- Modify: `web/app/chats/[id]/page.tsx`

- [ ] **Step 1: chat-panel.tsx**

- 메시지 영역: `role="log"` + `aria-live="polite"`
- 채팅 목록 영역: `role="region"` + `aria-label="대화 목록"`

- [ ] **Step 2: chat-input.tsx**

- input: `aria-label="메시지 입력"`
- 전송 버튼: `aria-label="메시지 전송"`
- PC에서 Shift+Enter = 줄바꿈 지원 (onKeyDown 핸들러)

- [ ] **Step 3: chat-message.tsx**

- 시스템 메시지: `role="status"`
- 일반 메시지: 시간 정보를 sr-only span으로 추가 (있는 경우)

- [ ] **Step 4: chat-list-item.tsx**

- 활성 채팅: `aria-current="true"`
- 안 읽은 뱃지: `aria-label={`${name} — ${unreadCount}개 안 읽은 메시지`}`

- [ ] **Step 5: chats/page.tsx + chats/[id]/page.tsx**

- 채팅방 선택 시 input auto-focus

- [ ] **Step 6: Commit**

```bash
git add web/components/chat/ web/app/chats/
git commit -m "feat(a11y): chat — role=log, aria-live, aria-current, Shift+Enter, input labels"
```

---

## Chunk 6: SSE + 네비게이션 뱃지 + 최종 검증

---

### Task 19: SSE 훅 개선 — 연결 상태 + 지수 백오프

**Files:**
- Modify: `web/lib/hooks/use-sse.ts`

- [ ] **Step 1: isConnected, isReconnecting 상태 노출 + 지수 백오프**

```typescript
// use-sse.ts에 추가할 상태:
const [connectionStatus, setConnectionStatus] = useState<"connected" | "reconnecting" | "disconnected">("disconnected");
const retryCountRef = useRef(0);
const MAX_RETRIES = 10;

// onerror 핸들러 수정:
es.onerror = () => {
  es.close();
  setConnectionStatus("reconnecting");
  if (retryCountRef.current < MAX_RETRIES) {
    const delay = Math.min(1000 * Math.pow(2, retryCountRef.current), 30000);
    retryCountRef.current++;
    setTimeout(() => setReconnectCount((c) => c + 1), delay);
  } else {
    setConnectionStatus("disconnected");
  }
};

// onopen 핸들러 추가:
es.onopen = () => {
  setConnectionStatus("connected");
  retryCountRef.current = 0;
};

// 훅 반환값에 추가:
return connectionStatus;
```

- [ ] **Step 2: 채팅에 연결 상태 배너 표시**

`chat-panel.tsx`에서 connectionStatus가 "reconnecting"이면 상단 배너:

```tsx
{connectionStatus === "reconnecting" && (
  <div role="alert" className="bg-[#e67e22]/10 text-[#e67e22] text-xs text-center py-2">
    연결이 끊어졌습니다. 재연결 중...
  </div>
)}
```

- [ ] **Step 3: Commit**

```bash
git add web/lib/hooks/use-sse.ts web/components/chat/chat-panel.tsx
git commit -m "feat(a11y): SSE — connection status, exponential backoff, reconnection banner"
```

---

### Task 20: 네비게이션 뱃지 (읽지 않은 알림/채팅)

**Files:**
- Modify: `web/components/layout/header.tsx`
- Modify: `web/components/layout/bottom-nav.tsx`

- [ ] **Step 1: 헤더 알림 뱃지**

알림 아이콘 옆에 dot 뱃지:

```tsx
{hasUnread && (
  <span className="absolute -top-1 -right-1 w-2.5 h-2.5 bg-[#e74c3c] rounded-full" />
)}
```

`aria-label`을 동적으로: `aria-label={hasUnread ? "읽지 않은 알림 있음" : "알림"}`

- [ ] **Step 2: 모바일 채팅 탭 뱃지**

채팅 탭에 안 읽은 수 표시. `aria-label`에 건수 포함.

- [ ] **Step 3: Commit**

```bash
git add web/components/layout/header.tsx web/components/layout/bottom-nav.tsx
git commit -m "feat(a11y): navigation badges — unread dot on notifications, count on chat tab"
```

---

### Task 21: 최종 검증

- [ ] **Step 1: 전체 테스트 실행**

Run: `cd web && npx vitest run`
Expected: All tests pass

- [ ] **Step 2: 빌드 확인**

Run: `cd web && npx next build`
Expected: Build succeeds without errors

- [ ] **Step 3: Lint 확인**

Run: `cd web && npx eslint . --max-warnings=0 2>&1 | tail -10`
Expected: No errors

- [ ] **Step 4: 접근성 체크리스트 최종 확인**

수동 확인:
- [ ] Tab 키로 모든 인터랙티브 요소 순회 가능
- [ ] 모달 열기 → Tab이 모달 내에서만 순환 → Escape 닫기 → 포커스 복귀
- [ ] Skip-to-content 링크 동작
- [ ] 스크린 리더로 주요 흐름 테스트 (VoiceOver)
- [ ] prefers-reduced-motion 활성화 시 애니메이션 제거 확인

- [ ] **Step 5: 변경 사항 있으면 Commit**

변경이 있는 경우에만:
```bash
git status
# 변경 파일이 있으면 개별 파일 지정하여 commit
```
