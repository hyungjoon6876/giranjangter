# Admin Dashboard Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a separate Next.js admin dashboard (`admin/`) with bright modern design, backed by new Go admin API endpoints for operations management.

**Architecture:** Next.js 15 admin frontend in `admin/` directory connects to existing Go backend via expanded `/api/v1/admin/*` endpoints. Same JWT auth (role: moderator/admin). Backend gets ~10 new endpoints. Frontend uses shadcn/ui + Tailwind for clean modern UI.

**Tech Stack:** Next.js 15, React 19, Tailwind CSS 3.4, shadcn/ui, Recharts, TanStack Query + Table, TypeScript 5

---

## File Structure

### Backend (new files in `backend/cmd/server/`)

```
backend/cmd/server/
├── handlers_admin.go         # MODIFY: add new admin endpoints
├── handlers_admin_stats.go   # CREATE: dashboard stats endpoint
├── handlers_admin_users.go   # CREATE: user management endpoints
├── handlers_admin_audit.go   # CREATE: audit log + chat inspection
├── main.go                   # MODIFY: register new admin routes
```

### Admin Frontend (new directory)

```
admin/
├── package.json
├── tsconfig.json
├── next.config.ts
├── tailwind.config.ts
├── postcss.config.mjs
├── .env.local                    # NEXT_PUBLIC_API_URL
├── app/
│   ├── globals.css               # Tailwind + Inter/Pretendard fonts
│   ├── layout.tsx                # Root layout with sidebar + header
│   ├── page.tsx                  # Dashboard (KPI cards + chart + recent reports)
│   ├── reports/
│   │   └── page.tsx              # Report management table + detail panel
│   ├── users/
│   │   ├── page.tsx              # User management table
│   │   └── [id]/
│   │       └── page.tsx          # User detail (profile + moderation history)
│   ├── listings/
│   │   └── page.tsx              # Listing management table
│   ├── trades/
│   │   └── page.tsx              # Trade/reservation monitoring
│   ├── audit-logs/
│   │   └── page.tsx              # Audit log viewer
│   └── login/
│       └── page.tsx              # Admin login
├── components/
│   ├── layout/
│   │   ├── admin-sidebar.tsx     # Left sidebar navigation
│   │   ├── admin-header.tsx      # Top header with user info
│   │   └── admin-shell.tsx       # Layout wrapper
│   ├── ui/                       # shadcn/ui components (auto-generated)
│   │   ├── button.tsx
│   │   ├── card.tsx
│   │   ├── table.tsx
│   │   ├── badge.tsx
│   │   ├── input.tsx
│   │   ├── select.tsx
│   │   ├── dialog.tsx
│   │   ├── sheet.tsx             # Side panel (report detail)
│   │   └── tabs.tsx
│   ├── dashboard/
│   │   ├── kpi-card.tsx          # Stat card with icon + trend
│   │   └── daily-chart.tsx       # Line chart (Recharts)
│   └── data-table/
│       └── data-table.tsx        # Reusable TanStack Table wrapper
├── lib/
│   ├── api-client.ts             # Admin API client (JWT, fetch)
│   ├── types.ts                  # Admin-specific TypeScript types
│   ├── hooks/
│   │   ├── use-dashboard.ts      # Dashboard data hooks
│   │   ├── use-reports.ts        # Report management hooks
│   │   ├── use-users.ts          # User management hooks
│   │   └── use-audit.ts          # Audit log hooks
│   └── utils.ts                  # Shared utilities
└── __tests__/
    └── lib/
        └── utils.test.ts
```

---

## Chunk 1: Backend Admin API Expansion

### Task 1: Dashboard Stats Endpoint

**Files:**
- Create: `backend/cmd/server/handlers_admin_stats.go`
- Modify: `backend/cmd/server/main.go`

- [ ] **Step 1: Create stats handler**

```go
// backend/cmd/server/handlers_admin_stats.go
package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handleAdminDashboardStats(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		var stats struct {
			TotalUsers       int `json:"totalUsers"`
			NewUsersToday    int `json:"newUsersToday"`
			ActiveListings   int `json:"activeListings"`
			NewListingsToday int `json:"newListingsToday"`
			PendingReports   int `json:"pendingReports"`
			TradesToday      int `json:"tradesToday"`
			ActiveChats      int `json:"activeChats"`
			RestrictedUsers  int `json:"restrictedUsers"`
		}

		db.QueryRow("SELECT COUNT(*) FROM users").Scan(&stats.TotalUsers)
		db.QueryRow("SELECT COUNT(*) FROM users WHERE created_at >= $1", today).Scan(&stats.NewUsersToday)
		db.QueryRow("SELECT COUNT(*) FROM listings WHERE status = 'available' AND visibility = 'public'").Scan(&stats.ActiveListings)
		db.QueryRow("SELECT COUNT(*) FROM listings WHERE created_at >= $1", today).Scan(&stats.NewListingsToday)
		db.QueryRow("SELECT COUNT(*) FROM reports WHERE status = 'submitted'").Scan(&stats.PendingReports)
		db.QueryRow("SELECT COUNT(*) FROM trade_completions WHERE confirmed_at >= $1", today).Scan(&stats.TradesToday)
		db.QueryRow("SELECT COUNT(*) FROM chat_rooms WHERE chat_status = 'open'").Scan(&stats.ActiveChats)
		db.QueryRow("SELECT COUNT(*) FROM users WHERE account_status IN ('restricted', 'suspended')").Scan(&stats.RestrictedUsers)

		c.JSON(http.StatusOK, stats)
	}
}
```

- [ ] **Step 2: Register in main.go**

Add to admin route group:
```go
admin.GET("/dashboard/stats", handleAdminDashboardStats(db))
```

- [ ] **Step 3: Run `go build ./...`**

- [ ] **Step 4: Commit**
```bash
git add backend/cmd/server/handlers_admin_stats.go backend/cmd/server/main.go
git commit -m "feat(admin-api): add dashboard stats endpoint"
```

---

### Task 2: User Management Endpoints

**Files:**
- Create: `backend/cmd/server/handlers_admin_users.go`
- Modify: `backend/cmd/server/main.go`

- [ ] **Step 1: Create user admin handlers**

```go
// backend/cmd/server/handlers_admin_users.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handleAdminListUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Query("q")
		status := c.Query("status")
		limit := 50

		query := `SELECT u.id, up.nickname, u.account_status, u.role,
			up.completed_trade_count, up.alignment_score, up.alignment_grade,
			u.last_login_at, u.created_at
			FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id
			WHERE 1=1`
		args := []interface{}{}
		argN := 1

		if q != "" {
			query += ` AND (up.nickname ILIKE $` + itoa(argN) + `)`
			args = append(args, "%"+q+"%")
			argN++
		}
		if status != "" {
			query += ` AND u.account_status = $` + itoa(argN)
			args = append(args, status)
			argN++
		}
		query += ` ORDER BY u.created_at DESC LIMIT ` + itoa(limit)

		rows, err := db.Query(query, args...)
		if err != nil {
			log.Printf("admin list users error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var users []gin.H
		for rows.Next() {
			var id, nickname, accountStatus, role, grade string
			var tradeCount int
			var score float64
			var lastLogin, created time.Time
			if err := rows.Scan(&id, &nickname, &accountStatus, &role, &tradeCount, &score, &grade, &lastLogin, &created); err != nil {
				continue
			}
			users = append(users, gin.H{
				"userId": id, "nickname": nickname, "accountStatus": accountStatus, "role": role,
				"completedTradeCount": tradeCount, "alignmentScore": score, "alignmentGrade": grade,
				"lastLoginAt": lastLogin.Format(time.RFC3339), "createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func handleAdminGetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")

		var u struct {
			ID         string  `json:"userId"`
			Nickname   string  `json:"nickname"`
			Status     string  `json:"accountStatus"`
			Role       string  `json:"role"`
			Intro      *string `json:"introduction"`
			Server     *string `json:"primaryServerId"`
			TradeCount int     `json:"completedTradeCount"`
			ReviewPos  int     `json:"positiveReviewCount"`
			Score      float64 `json:"alignmentScore"`
			Grade      string  `json:"alignmentGrade"`
			TrustBadge *string `json:"trustBadge"`
		}
		err := db.QueryRow(`SELECT u.id, up.nickname, u.account_status, u.role,
			up.introduction, up.primary_server_id, up.completed_trade_count,
			up.positive_review_count, up.alignment_score, up.alignment_grade, up.trust_badge
			FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.id = $1`, userID).
			Scan(&u.ID, &u.Nickname, &u.Status, &u.Role, &u.Intro, &u.Server,
				&u.TradeCount, &u.ReviewPos, &u.Score, &u.Grade, &u.TrustBadge)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "사용자를 찾을 수 없습니다."}})
			return
		}
		c.JSON(http.StatusOK, u)
	}
}

func handleAdminUserModerationHistory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")
		rows, err := db.Query(`SELECT id, report_id, actor_user_id, action_code, restriction_scope, duration_days, memo, created_at
			FROM moderation_actions WHERE target_user_id = $1 ORDER BY created_at DESC LIMIT 50`, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var actions []gin.H
		for rows.Next() {
			var id, actorID, actionCode string
			var reportID, scope, memo *string
			var duration *int
			var created time.Time
			if err := rows.Scan(&id, &reportID, &actorID, &actionCode, &scope, &duration, &memo, &created); err != nil {
				continue
			}
			actions = append(actions, gin.H{
				"actionId": id, "reportId": reportID, "actorUserId": actorID,
				"actionCode": actionCode, "restrictionScope": scope,
				"durationDays": duration, "memo": memo, "createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": actions})
	}
}

// simple int-to-string for query building
func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}
```

Note: Add `"fmt"` to imports.

- [ ] **Step 2: Register routes**

```go
admin.GET("/users", handleAdminListUsers(db))
admin.GET("/users/:userId", handleAdminGetUser(db))
admin.GET("/users/:userId/moderation-history", handleAdminUserModerationHistory(db))
```

- [ ] **Step 3: Run `go build ./...`**

- [ ] **Step 4: Commit**
```bash
git commit -m "feat(admin-api): add user management endpoints (list, detail, moderation history)"
```

---

### Task 3: Audit Log + Chat Inspection Endpoints

**Files:**
- Create: `backend/cmd/server/handlers_admin_audit.go`
- Modify: `backend/cmd/server/main.go`

- [ ] **Step 1: Create audit and chat handlers**

```go
// backend/cmd/server/handlers_admin_audit.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handleAdminListAuditLogs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT id, actor_id, actor_role, action, target_type, target_id, details_json, ip_address, created_at
			FROM audit_logs ORDER BY created_at DESC LIMIT 100`)
		if err != nil {
			log.Printf("audit log error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var logs []gin.H
		for rows.Next() {
			var id, actorID, actorRole, action, targetType, targetID string
			var details, ip *string
			var created time.Time
			if err := rows.Scan(&id, &actorID, &actorRole, &action, &targetType, &targetID, &details, &ip, &created); err != nil {
				continue
			}
			logs = append(logs, gin.H{
				"logId": id, "actorId": actorID, "actorRole": actorRole, "action": action,
				"targetType": targetType, "targetId": targetID, "details": details,
				"ipAddress": ip, "createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": logs})
	}
}

func handleAdminChatMessages(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")

		// Verify chat exists
		var exists bool
		db.QueryRow("SELECT EXISTS(SELECT 1 FROM chat_rooms WHERE id = $1)", chatID).Scan(&exists)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "채팅방을 찾을 수 없습니다."}})
			return
		}

		rows, err := db.Query(`SELECT id, sender_user_id, message_type, body_text, metadata_json, sent_at
			FROM chat_messages WHERE chat_room_id = $1 AND deleted_at IS NULL
			ORDER BY sent_at ASC LIMIT 200`, chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var messages []gin.H
		for rows.Next() {
			var id, msgType string
			var senderID, bodyText, metadata *string
			var sentAt time.Time
			if err := rows.Scan(&id, &senderID, &msgType, &bodyText, &metadata, &sentAt); err != nil {
				continue
			}
			messages = append(messages, gin.H{
				"messageId": id, "senderUserId": senderID, "messageType": msgType,
				"bodyText": bodyText, "metadataJson": metadata, "sentAt": sentAt.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": messages})
	}
}

func handleAdminListTrades(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`SELECT tc.id, tc.listing_id, l.title, tc.status,
			tc.requested_by_user_id, tc.counterpart_user_id, tc.auto_confirm_at, tc.created_at
			FROM trade_completions tc
			LEFT JOIN listings l ON tc.listing_id = l.id
			ORDER BY tc.created_at DESC LIMIT 50`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var trades []gin.H
		for rows.Next() {
			var id, listingID, title, status, requestedBy, counterpart string
			var autoConfirm, created time.Time
			if err := rows.Scan(&id, &listingID, &title, &status, &requestedBy, &counterpart, &autoConfirm, &created); err != nil {
				continue
			}
			trades = append(trades, gin.H{
				"completionId": id, "listingId": listingID, "listingTitle": title, "status": status,
				"requestedByUserId": requestedBy, "counterpartUserId": counterpart,
				"autoConfirmAt": autoConfirm.Format(time.RFC3339), "createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": trades})
	}
}

func handleAdminListAllListings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")
		visibility := c.Query("visibility")

		query := `SELECT l.id, l.title, l.item_name, l.status, l.visibility, l.listing_type,
			up.nickname as author_nickname, l.created_at
			FROM listings l
			LEFT JOIN user_profiles up ON l.author_user_id = up.user_id
			WHERE l.deleted_at IS NULL`
		args := []interface{}{}
		argN := 1

		if status != "" {
			query += fmt.Sprintf(" AND l.status = $%d", argN)
			args = append(args, status)
			argN++
		}
		if visibility != "" {
			query += fmt.Sprintf(" AND l.visibility = $%d", argN)
			args = append(args, visibility)
			argN++
		}
		query += " ORDER BY l.created_at DESC LIMIT 50"

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}
		defer rows.Close()

		var listings []gin.H
		for rows.Next() {
			var id, title, itemName, st, vis, lt, author string
			var created time.Time
			if err := rows.Scan(&id, &title, &itemName, &st, &vis, &lt, &author, &created); err != nil {
				continue
			}
			listings = append(listings, gin.H{
				"listingId": id, "title": title, "itemName": itemName, "status": st,
				"visibility": vis, "listingType": lt, "authorNickname": author,
				"createdAt": created.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": listings})
	}
}

func handleAdminRestoreListing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listingID := c.Param("id")
		adminID := middleware.GetUserID(c)

		tx, _ := db.Begin()
		defer tx.Rollback()

		tx.Exec("UPDATE listings SET visibility = 'public', updated_at = NOW() WHERE id = $1", listingID)
		tx.Exec(`INSERT INTO audit_logs (id, actor_id, actor_role, action, target_type, target_id)
			VALUES ($1, $2, 'admin', 'listing.moderation.restore', 'listing', $3)`,
			uuid.New().String(), adminID, listingID)
		tx.Commit()

		c.JSON(http.StatusOK, gin.H{"listingId": listingID, "visibility": "public"})
	}
}

func handleAdminUpdateReportStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		reportID := c.Param("reportId")
		var req struct {
			Status string `json:"status" binding:"required,oneof=submitted assigned resolved"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}
		db.Exec("UPDATE reports SET status = $1, updated_at = NOW() WHERE id = $2", req.Status, reportID)
		c.JSON(http.StatusOK, gin.H{"reportId": reportID, "status": req.Status})
	}
}
```

Note: Add `"fmt"` and `"github.com/google/uuid"`, `"github.com/jym/lincle/internal/middleware"` to imports.

- [ ] **Step 2: Register all new routes in main.go**

```go
// Add to admin group
admin.GET("/dashboard/stats", handleAdminDashboardStats(db))
admin.GET("/users", handleAdminListUsers(db))
admin.GET("/users/:userId", handleAdminGetUser(db))
admin.GET("/users/:userId/moderation-history", handleAdminUserModerationHistory(db))
admin.GET("/audit-logs", handleAdminListAuditLogs(db))
admin.GET("/chats/:chatId/messages", handleAdminChatMessages(db))
admin.GET("/trades", handleAdminListTrades(db))
admin.GET("/listings", handleAdminListAllListings(db))
admin.POST("/listings/:id/restore", handleAdminRestoreListing(db))
admin.PATCH("/reports/:reportId", handleAdminUpdateReportStatus(db))
```

- [ ] **Step 3: Run `go build ./...` and `go test ./...`**

- [ ] **Step 4: Commit**
```bash
git commit -m "feat(admin-api): add audit logs, chat inspection, trades, listings management endpoints"
```

---

## Chunk 2: Admin Frontend Foundation

### Task 4: Initialize Admin Next.js Project

**Files:**
- Create: `admin/` directory

- [ ] **Step 1: Create Next.js project**
```bash
cd /Users/jym/github-workspace/lincle/.claude/worktrees/rippling-crafting-noodle
npx create-next-app@latest admin --typescript --tailwind --eslint --app --src-dir=false --import-alias="@/*" --use-npm
```

- [ ] **Step 2: Install dependencies**
```bash
cd admin
npm install @tanstack/react-query @tanstack/react-table recharts
npm install -D @testing-library/react @testing-library/jest-dom vitest @vitejs/plugin-react jsdom
```

- [ ] **Step 3: Configure Tailwind for modern bright theme**

```typescript
// admin/tailwind.config.ts
import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./app/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}", "./lib/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: { DEFAULT: "#6366f1", light: "#818cf8", dark: "#4f46e5" },
        accent: { DEFAULT: "#f59e0b", light: "#fbbf24" },
        success: "#10b981",
        danger: "#ef4444",
        warning: "#f59e0b",
        surface: "#f8fafc",
        card: "#ffffff",
        sidebar: "#1e1b4b",
        "sidebar-hover": "#312e81",
        "text-primary": "#0f172a",
        "text-secondary": "#64748b",
        "text-dim": "#94a3b8",
        border: "#e2e8f0",
      },
      fontFamily: {
        sans: ["'Inter'", "'Pretendard'", "-apple-system", "sans-serif"],
      },
    },
  },
  plugins: [],
};
export default config;
```

- [ ] **Step 4: Set up globals.css**

```css
/* admin/app/globals.css */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');

@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  body {
    @apply bg-surface text-text-primary font-sans antialiased;
  }
}
```

- [ ] **Step 5: Create .env.local**
```
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

- [ ] **Step 6: Verify `npm run dev`**
- [ ] **Step 7: Commit**
```bash
git commit -m "feat(admin): initialize Next.js admin project with bright modern theme"
```

---

### Task 5: Admin API Client + Types + Hooks

**Files:**
- Create: `admin/lib/api-client.ts`
- Create: `admin/lib/types.ts`
- Create: `admin/lib/utils.ts`
- Create: `admin/lib/providers.tsx`
- Create: `admin/lib/hooks/use-dashboard.ts`
- Create: `admin/lib/hooks/use-reports.ts`
- Create: `admin/lib/hooks/use-users.ts`
- Create: `admin/lib/hooks/use-audit.ts`

- [ ] **Step 1: Create types**

Admin-specific types: DashboardStats, AdminUser, ModerationAction, AuditLog, AdminReport, AdminListing, TradeCompletion, AdminChatMessage. Follow web/lib/types.ts patterns.

- [ ] **Step 2: Create API client** (same pattern as web, JWT via localStorage)

Methods: getDashboardStats, getReports, getReport, executeReportAction, updateReportStatus, getUsers, getUser, getUserModerationHistory, restrictUser, getListings, hideListing, restoreListing, getAuditLogs, getChatMessages, getTrades.

- [ ] **Step 3: Create utils** (formatDate, formatTimeAgo, statusBadgeColor)
- [ ] **Step 4: Create providers** (QueryClientProvider)
- [ ] **Step 5: Create all hooks** (TanStack Query wrappers)
- [ ] **Step 6: Commit**
```bash
git commit -m "feat(admin): add API client, types, hooks for all admin endpoints"
```

---

### Task 6: Admin Layout Shell

**Files:**
- Create: `admin/components/layout/admin-sidebar.tsx`
- Create: `admin/components/layout/admin-header.tsx`
- Create: `admin/components/layout/admin-shell.tsx`
- Modify: `admin/app/layout.tsx`

- [ ] **Step 1: Create sidebar** — dark indigo (#1e1b4b) left sidebar:
  - Logo: "기란장터 Admin" (white text)
  - Nav items: 대시보드, 신고 관리, 사용자 관리, 매물 관리, 거래 모니터링, 감사 로그
  - Active state: white text + indigo-700 bg
  - Inline SVG icons (similar to web/ approach)
  - Width: w-56

- [ ] **Step 2: Create header** — white bg, right-aligned user info, border-bottom

- [ ] **Step 3: Create shell** — sidebar + header + content area

- [ ] **Step 4: Update layout.tsx** with Providers + AdminShell

- [ ] **Step 5: Verify `npm run build`**
- [ ] **Step 6: Commit**
```bash
git commit -m "feat(admin): add layout shell with sidebar navigation"
```

---

## Chunk 3: Admin Pages

### Task 7: Dashboard Page

**Files:**
- Create: `admin/components/dashboard/kpi-card.tsx`
- Create: `admin/components/dashboard/daily-chart.tsx`
- Modify: `admin/app/page.tsx`

- [ ] **Step 1: Create KPI card** — white card with colored left border, big number, label, optional trend

- [ ] **Step 2: Create daily chart** — Recharts AreaChart placeholder (data from stats)

- [ ] **Step 3: Create dashboard page**:
  - 4 KPI cards: 미처리 신고 (danger), 오늘 거래 (success), 활성 매물 (primary), 신규 가입 (accent)
  - Chart section (placeholder)
  - Recent reports table (last 5)

- [ ] **Step 4: Verify**
- [ ] **Step 5: Commit**
```bash
git commit -m "feat(admin): implement dashboard page with KPI cards"
```

---

### Task 8: Report Management Page

**Files:**
- Create: `admin/components/data-table/data-table.tsx`
- Create: `admin/app/reports/page.tsx`

- [ ] **Step 1: Create reusable data table** — TanStack Table with sorting, filtering

- [ ] **Step 2: Create reports page**:
  - Table columns: ID, 유형, 대상, 신고자, 상태, 날짜
  - Status filter tabs: 전체, 접수, 처리중, 완료
  - Row click → slide-out Sheet panel with report detail
  - Action buttons: 할당, 처리 (opens dialog)

- [ ] **Step 3: Verify**
- [ ] **Step 4: Commit**
```bash
git commit -m "feat(admin): implement report management page with data table"
```

---

### Task 9: User Management Pages

**Files:**
- Create: `admin/app/users/page.tsx`
- Create: `admin/app/users/[id]/page.tsx`

- [ ] **Step 1: Create user list page** — table with search, status filter, role filter

- [ ] **Step 2: Create user detail page**:
  - Profile card: nickname, role, status, alignment score/grade
  - Stats: trade count, reviews, trust badge
  - Moderation history table
  - Action buttons: 제한, 정지, 복원

- [ ] **Step 3: Verify**
- [ ] **Step 4: Commit**
```bash
git commit -m "feat(admin): implement user management pages (list + detail)"
```

---

### Task 10: Listing Management + Trade Monitoring + Audit Log Pages

**Files:**
- Create: `admin/app/listings/page.tsx`
- Create: `admin/app/trades/page.tsx`
- Create: `admin/app/audit-logs/page.tsx`

- [ ] **Step 1: Listings page** — table with status/visibility filters, hide/restore buttons

- [ ] **Step 2: Trades page** — table showing trade completions, status, auto-confirm timer

- [ ] **Step 3: Audit log page** — read-only table of all admin actions

- [ ] **Step 4: Verify**
- [ ] **Step 5: Commit**
```bash
git commit -m "feat(admin): implement listings, trades, and audit log pages"
```

---

### Task 11: Admin Login Page

**Files:**
- Create: `admin/app/login/page.tsx`

- [ ] **Step 1: Create login page** — clean centered form, dev login for moderator/admin role

- [ ] **Step 2: Commit**
```bash
git commit -m "feat(admin): implement admin login page"
```

---

## Chunk 4: Verification & Deployment

### Task 12: Vitest Configuration + Tests

**Files:**
- Create: `admin/vitest.config.ts`
- Create: `admin/vitest.setup.ts`
- Create: `admin/__tests__/lib/utils.test.ts`

- [ ] **Step 1: Configure Vitest**
- [ ] **Step 2: Write utils tests**
- [ ] **Step 3: Run `npx vitest run`**
- [ ] **Step 4: Commit**
```bash
git commit -m "test(admin): add Vitest config and utility tests"
```

---

### Task 13: Final Build Verification

- [ ] **Step 1: Backend**
```bash
cd backend && go build ./... && go test ./...
```

- [ ] **Step 2: Admin frontend**
```bash
cd admin && npm run build && npx vitest run
```

- [ ] **Step 3: Verify all admin routes render**

Expected routes: `/`, `/reports`, `/users`, `/users/[id]`, `/listings`, `/trades`, `/audit-logs`, `/login`

- [ ] **Step 4: Commit**
```bash
git commit -m "verify: admin dashboard complete — all builds and tests passing"
```
