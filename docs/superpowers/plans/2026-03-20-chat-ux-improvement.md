# Chat UX Improvement Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 채팅 UX를 MVP에서 거래 플랫폼 수준으로 전면 개선 — 모바일 줌 수정, 메시지 피드백, 가독성, 거래 연동, 실시간 안정성

**Architecture:** 4 Phase Bottom-Up 접근. 각 Phase 독립 배포 가능. 백엔드 API 갭 수정을 Pre-requisite로 먼저 처리한 후, 프론트엔드 개선을 순차 적용.

**Tech Stack:** Next.js 16, React 19, TanStack Query, SSE (EventSource), Go/Gin, PostgreSQL, TailwindCSS

**Spec:** `docs/superpowers/specs/2026-03-20-chat-ux-improvement-design.md`

---

## Pre-requisite: Backend API Gap Fix

### Task 1: ListChatRooms SQL 확장 — unreadCount, lastMessage, listing 필드

**Files:**
- Modify: `backend/internal/repository/postgres_chat.go:63-90` (ListChatRooms SQL)
- Modify: `backend/internal/repository/interfaces.go:328-338` (ChatRoomListItem struct — 구조체는 interfaces.go에만 정의)

**Context:** 현재 `ListChatRooms`는 `chatRoomId`, `listingId`, `listingTitle`, `chatStatus`, `counterparty`, `updatedAt`만 반환. 프론트 `ChatRoom` 타입이 요구하는 `unreadCount`, `lastMessage`, `listingThumbnail`, `listingStatus`, `myLastReadMessageId`가 누락.

- [ ] **Step 1: ChatRoomListItem 구조체 확장**

`interfaces.go`의 `ChatRoomListItem` struct에 필드 추가 (구조체는 interfaces.go에만 정의됨):

```go
type ChatRoomListItem struct {
	ChatRoomID       string
	ListingID        string
	ListingTitle     string
	ListingThumbnail *string   // NEW
	ListingStatus    string    // NEW
	ChatStatus       string
	LastMessageAt    *time.Time
	UpdatedAt        *time.Time
	CounterpartID    string
	CounterpartNick  string
	CounterpartBadge string
	UnreadCount      int       // NEW
	LastMessageBody  *string   // NEW
	LastMessageSentAt *time.Time // NEW
	MyLastReadMsgID  *string   // NEW
}
```

`postgres_chat.go`의 `rows.Scan`도 새 필드에 맞게 수정.

- [ ] **Step 2: ListChatRooms SQL 확장**

기존 SQL에 LEFT JOIN 및 서브쿼리 추가:

```sql
SELECT cr.id, cr.listing_id, l.title,
  (SELECT li.thumbnail_url FROM listing_images li WHERE li.listing_id = l.id ORDER BY li.sort_order ASC LIMIT 1) as thumbnail_url,
  l.status,
  cr.chat_status, cr.last_message_at, cr.updated_at,
  CASE WHEN cr.seller_user_id = $1 THEN cr.buyer_user_id ELSE cr.seller_user_id END as counterpart_id,
  p.nickname, p.trust_badge,
  COALESCE(
    (SELECT COUNT(*) FROM chat_messages cm
     WHERE cm.chat_room_id = cr.id AND cm.deleted_at IS NULL
     AND cm.sent_at > COALESCE(
       (SELECT crc.updated_at FROM chat_read_cursors crc
        WHERE crc.chat_room_id = cr.id AND crc.user_id = $2), cr.created_at
     )
     AND cm.sender_user_id != $3
    ), 0
  ) as unread_count,
  (SELECT cm2.body_text FROM chat_messages cm2
   WHERE cm2.chat_room_id = cr.id AND cm2.deleted_at IS NULL
   ORDER BY cm2.sent_at DESC LIMIT 1) as last_message_body,
  (SELECT cm3.sent_at FROM chat_messages cm3
   WHERE cm3.chat_room_id = cr.id AND cm3.deleted_at IS NULL
   ORDER BY cm3.sent_at DESC LIMIT 1) as last_message_sent_at,
  (SELECT crc2.last_read_message_id FROM chat_read_cursors crc2
   WHERE crc2.chat_room_id = cr.id AND crc2.user_id = $4) as my_last_read_msg_id
FROM chat_rooms cr
JOIN listings l ON cr.listing_id = l.id
JOIN user_profiles p ON p.user_id = CASE WHEN cr.seller_user_id = $5 THEN cr.buyer_user_id ELSE cr.seller_user_id END
WHERE cr.seller_user_id = $6 OR cr.buyer_user_id = $7
ORDER BY COALESCE(cr.last_message_at, cr.created_at) DESC
LIMIT 50
```

`rows.Scan`에 새 필드 추가. `$1`~`$7` 모두 `userID` 값.

- [ ] **Step 3: Go 빌드 검증**

Run: `cd backend && go build ./cmd/server/`
Expected: 성공 (컴파일 에러 없음)

---

### Task 2: handleListChats 응답에 새 필드 포함

**Files:**
- Modify: `backend/cmd/server/handlers_chat.go:68-90` (handleListChats)

- [ ] **Step 1: 핸들러 응답에 새 필드 추가**

`handleListChats` 내 `gin.H` 맵에 추가:

```go
chat := gin.H{
	"chatRoomId":   item.ChatRoomID,
	"listingId":    item.ListingID,
	"listingTitle": item.ListingTitle,
	"listingThumbnail": item.ListingThumbnail,  // NEW
	"listingStatus":    item.ListingStatus,      // NEW
	"chatStatus":  item.ChatStatus,
	"counterparty": gin.H{
		"userId":     item.CounterpartID,
		"nickname":   item.CounterpartNick,
		"trustBadge": item.CounterpartBadge,
	},
	"unreadCount": item.UnreadCount,  // NEW
	"updatedAt":   item.UpdatedAt,
	"myLastReadMessageId": item.MyLastReadMsgID,  // NEW
}
// lastMessage 추가
if item.LastMessageBody != nil {
	chat["lastMessage"] = gin.H{
		"bodyText": *item.LastMessageBody,
		"sentAt":   item.LastMessageSentAt,
	}
}
```

- [ ] **Step 2: 빌드 + 테스트**

Run: `cd backend && go build ./cmd/server/ && go test ./...`

- [ ] **Step 3: 커밋**

```
fix(backend): handleListChats 응답 확장 — unreadCount, lastMessage, listing 상세, myLastReadMessageId
```

---

### Task 3: 프론트엔드 타입 업데이트

**Files:**
- Modify: `web/lib/types.ts:62-83` (ChatRoom, Message)

- [ ] **Step 1: ChatRoom 타입에 새 필드 추가**

```typescript
export interface ChatRoom {
  chatRoomId: string;
  listingId: string;
  listingTitle: string;
  listingThumbnail?: string;
  listingStatus: string;
  counterparty: Author;
  chatStatus: string;
  lastMessage?: Message;
  unreadCount: number;
  updatedAt: string;
  myLastReadMessageId?: string;  // NEW
}
```

- [ ] **Step 2: Message 타입에 status 필드 추가**

```typescript
export interface Message {
  messageId: string;
  chatRoomId: string;
  senderUserId?: string;
  messageType: "text" | "system" | "reservation_card";
  bodyText?: string;
  metadataJson?: Record<string, unknown>;
  sentAt: string;
  status?: "sending" | "sent" | "failed";  // NEW: client-only
}
```

- [ ] **Step 3: 빌드 검증**

Run: `cd web && npx next build`

- [ ] **Step 4: 커밋**

```
feat(web): ChatRoom/Message 타입 확장 — myLastReadMessageId, status
```

---

## Phase 1: Mobile Fix + Message Feedback

### Task 4: 모바일 입력 줌 방지

**Files:**
- Modify: `web/app/layout.tsx` (viewport export 추가)
- Modify: `web/components/chat/chat-input.tsx:62` (font-size)

- [ ] **Step 1: layout.tsx에 viewport export 추가**

파일 상단 import에 `Viewport` 추가, `metadata` export 옆에:

```typescript
import type { Metadata, Viewport } from "next";

export const viewport: Viewport = {
  width: "device-width",
  initialScale: 1,
  maximumScale: 1,
};
```

- [ ] **Step 2: chat-input.tsx textarea font-size 변경**

`text-sm` → `text-base` (14px → 16px). textarea의 className에서 변경.

- [ ] **Step 3: 빌드 검증**

Run: `cd web && npx next build`

- [ ] **Step 4: 커밋**

```
fix(web): 모바일 입력 줌 방지 — viewport maximumScale + textarea 16px
```

---

### Task 5: 메시지 타임스탬프

**Files:**
- Modify: `web/components/chat/chat-message.tsx` (타임스탬프 렌더링)

- [ ] **Step 1: formatMessageTime 유틸 함수 작성**

`chat-message.tsx` 상단 또는 별도 util에:

```typescript
function formatMessageTime(sentAt: string): string {
  const date = new Date(sentAt);
  const now = new Date();
  const timeStr = date.toLocaleTimeString("ko-KR", {
    hour: "numeric",
    minute: "2-digit",
    hour12: true,
  });

  if (
    date.getFullYear() === now.getFullYear() &&
    date.getMonth() === now.getMonth() &&
    date.getDate() === now.getDate()
  ) {
    return timeStr; // 오늘: "오후 3:42"
  }

  if (date.getFullYear() === now.getFullYear()) {
    return `${date.getMonth() + 1}월 ${date.getDate()}일 ${timeStr}`;
  }

  return `${date.getFullYear()}.${date.getMonth() + 1}.${date.getDate()} ${timeStr}`;
}
```

- [ ] **Step 2: ChatMessage에 타임스탬프 추가**

text 메시지 버블 내부, `{message.bodyText}` 아래에:

```tsx
<span className="block text-xs text-text-secondary mt-1">
  {formatMessageTime(message.sentAt)}
</span>
```

시스템 메시지에는 추가하지 않음.

- [ ] **Step 3: 빌드 검증**

Run: `cd web && npx next build`

- [ ] **Step 4: 커밋**

```
feat(web): 채팅 메시지 타임스탬프 추가
```

---

### Task 6: Optimistic Update + clientMessageId

**Files:**
- Modify: `web/lib/api-client.ts:201-206` (sendMessage 시그니처)
- Modify: `web/lib/hooks/use-chats.ts:21-31` (useSendMessage)
- Modify: `web/app/chats/page.tsx` (sendMessage 호출부)
- Modify: `web/app/chats/[id]/page.tsx` (sendMessage 호출부)
- Modify: `web/components/chat/chat-panel.tsx` (sendMessage 호출부)

- [ ] **Step 1: apiClient.sendMessage에 clientMessageId 추가**

```typescript
async sendMessage(
  chatId: string,
  text: string,
  clientMessageId: string
): Promise<Message> {
  return this.fetch(`/chats/${chatId}/messages`, {
    method: "POST",
    body: JSON.stringify({
      messageType: "text",
      bodyText: text,
      clientMessageId,
    }),
  });
}
```

- [ ] **Step 2: useSendMessage에 optimistic update 구현**

```typescript
export function useSendMessage() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({
      chatId,
      text,
      clientMessageId,
    }: {
      chatId: string;
      text: string;
      clientMessageId: string;
    }) => apiClient.sendMessage(chatId, text, clientMessageId),

    onMutate: async ({ chatId, text, clientMessageId }) => {
      await qc.cancelQueries({ queryKey: ["messages", chatId] });
      const previous = qc.getQueryData<PaginatedResponse<Message>>([
        "messages",
        chatId,
      ]);

      // myUserId는 호출측에서 mutate 인자로 전달받거나 apiClient에서 획득
      const myUserId = apiClient.getUserId?.() ?? "";

      const optimisticMsg: Message = {
        messageId: clientMessageId,
        chatRoomId: chatId,
        senderUserId: myUserId, // isMine 판별에 실제 userId 필요
        messageType: "text",
        bodyText: text,
        sentAt: new Date().toISOString(),
        status: "sending",
      };

      // 백엔드는 sent_at DESC (최신 먼저) 반환 → optimistic 메시지는 맨 앞에 prepend
      qc.setQueryData<PaginatedResponse<Message>>(
        ["messages", chatId],
        (old) => ({
          data: [optimisticMsg, ...(old?.data ?? [])],
          cursor: old?.cursor ?? { hasMore: false },
        })
      );

      return { previous, chatId };
    },

    onSuccess: (serverMsg, { chatId, clientMessageId }) => {
      qc.setQueryData<PaginatedResponse<Message>>(
        ["messages", chatId],
        (old) => ({
          data: (old?.data ?? []).map((m) =>
            m.messageId === clientMessageId
              ? { ...serverMsg, status: "sent" as const }
              : m
          ),
          cursor: old?.cursor ?? { hasMore: false },
        })
      );
      qc.invalidateQueries({ queryKey: ["chats"] });
    },

    onError: (_err, { chatId, clientMessageId }, context) => {
      if (context?.previous) {
        // 롤백 대신 failed 상태로 표시
        qc.setQueryData<PaginatedResponse<Message>>(
          ["messages", chatId],
          (old) => ({
            data: (old?.data ?? []).map((m) =>
              m.messageId === clientMessageId
                ? { ...m, status: "failed" as const }
                : m
            ),
            cursor: old?.cursor ?? { hasMore: false },
          })
        );
      }
    },
  });
}
```

`PaginatedResponse` import 필요: `import type { PaginatedResponse, Message } from "@/lib/types";`

- [ ] **Step 3: 호출부 업데이트**

`chats/page.tsx`, `chats/[id]/page.tsx`, `chat-panel.tsx`에서 sendMessage 호출 시 `clientMessageId` 전달:

```typescript
import { v4 as uuidv4 } from "uuid";
// 또는 crypto.randomUUID() 사용

sendMessage.mutate({
  chatId: id,
  text,
  clientMessageId: crypto.randomUUID(),
});
```

기존 `mutate({ chatId, text })` 형태를 모두 업데이트.

- [ ] **Step 4: 빌드 검증**

Run: `cd web && npx next build`

- [ ] **Step 5: Vitest 테스트**

Run: `cd web && npx vitest run`

- [ ] **Step 6: 커밋**

```
feat(web): 메시지 optimistic update + clientMessageId 중복 방지
```

---

### Task 7: ChatMessage 전송 상태 표시

**Files:**
- Modify: `web/components/chat/chat-message.tsx` (status 기반 UI)

- [ ] **Step 1: ChatMessage에 전송 상태 표시 추가**

타임스탬프 영역을 상태 기반으로 분기:

```tsx
{/* 내 메시지 하단 상태 표시 */}
{isMine && (
  <span className="block text-xs mt-1 text-right">
    {message.status === "sending" ? (
      <span className="text-text-secondary">전송 중...</span>
    ) : message.status === "failed" ? (
      <button
        onClick={() => onRetry?.(message)}
        className="text-red-400 hover:underline"
      >
        전송 실패 · 재전송
      </button>
    ) : (
      <span className="text-text-secondary">
        {formatMessageTime(message.sentAt)} ✓
      </span>
    )}
  </span>
)}
{/* 상대 메시지 하단 타임스탬프 */}
{!isMine && message.messageType !== "system" && (
  <span className="block text-xs text-text-secondary mt-1">
    {formatMessageTime(message.sentAt)}
  </span>
)}
```

- [ ] **Step 2: ChatMessageProps에 onRetry 추가**

```typescript
interface ChatMessageProps {
  message: Message;
  isMine: boolean;
  onRetry?: (message: Message) => void;
}
```

- [ ] **Step 3: 부모 컴포넌트에서 onRetry 연결**

`chat-panel.tsx`와 `chats/[id]/page.tsx`의 ChatMessage 렌더링에:

```tsx
<ChatMessage
  message={msg}
  isMine={msg.senderUserId === myUserId}
  onRetry={(m) =>
    sendMessage.mutate({
      chatId: m.chatRoomId,
      text: m.bodyText ?? "",
      clientMessageId: crypto.randomUUID(),
    })
  }
/>
```

- [ ] **Step 4: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 5: 커밋**

```
feat(web): 메시지 전송 상태 표시 — sending/sent/failed + 재전송
```

---

### Task 8: 백엔드 read_receipt SSE 브로드캐스트

**Files:**
- Modify: `backend/cmd/server/handlers_chat.go:218-241` (handleMarkRead)
- Modify: `backend/cmd/server/main.go:137` (라우트 등록)

- [ ] **Step 1: handleMarkRead에 broker 파라미터 추가**

시그니처 변경 + SSE 브로드캐스트 로직 추가:

```go
func handleMarkRead(repo repository.ChatRepo, broker *event.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chatId")
		userID := middleware.GetUserID(c)

		ok, _ := repo.IsChatParticipant(c.Request.Context(), chatID, userID)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "채팅 참여자가 아닙니다."}})
			return
		}

		var req struct {
			LastReadMessageID string `json:"lastReadMessageId" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
			return
		}

		if err := repo.UpsertReadCursor(c.Request.Context(), chatID, userID, req.LastReadMessageID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "서버 오류"}})
			return
		}

		// SSE broadcast read_receipt to counterpart
		participants, err := repo.GetChatRoomParticipants(c.Request.Context(), chatID, userID)
		if err == nil && participants != nil {
			counterpart := participants.BuyerID
			if userID == participants.BuyerID {
				counterpart = participants.SellerID
			}
			broker.SendToUser(counterpart, event.SSEEvent{
				EventType: "read_receipt",
				Data: map[string]interface{}{
					"chatRoomId":        chatID,
					"readBy":            userID,
					"lastReadMessageId": req.LastReadMessageID,
					"readAt":            time.Now().UTC().Format(time.RFC3339),
				},
			})
		}

		c.Status(http.StatusNoContent)
	}
}
```

- [ ] **Step 2: main.go 라우트 등록 수정**

```go
write.POST("/chats/:chatId/read", handleMarkRead(chatRepo, sseBroker))
```

- [ ] **Step 3: 빌드 + 테스트**

Run: `cd backend && go build ./cmd/server/ && go test ./...`

- [ ] **Step 4: 커밋**

```
feat(backend): handleMarkRead에 read_receipt SSE 브로드캐스트 추가
```

---

### Task 9: 프론트엔드 읽음 표시

**Files:**
- Modify: `web/lib/api-client.ts` (markRead 메서드)
- Modify: `web/lib/hooks/use-sse.ts` (read_receipt 이벤트)
- Modify: `web/lib/hooks/use-chats.ts` (markRead 호출)
- Modify: `web/components/chat/chat-message.tsx` (읽음 체크 표시)

- [ ] **Step 1: apiClient에 markRead 메서드 추가**

```typescript
async markRead(chatId: string, lastReadMessageId: string): Promise<void> {
  await this.fetch(`/chats/${chatId}/read`, {
    method: "POST",
    body: JSON.stringify({ lastReadMessageId }),
  });
}
```

- [ ] **Step 2: use-sse.ts에 read_receipt 이벤트 핸들러 추가**

기존 `new_message`, `status_change` 핸들러 옆에:

```typescript
es.addEventListener("read_receipt", (e) => {
  const data = JSON.parse(e.data);
  qc.invalidateQueries({ queryKey: ["messages", data.chatRoomId] });
  qc.invalidateQueries({ queryKey: ["chats"] });
});
```

- [ ] **Step 3: 채팅방 진입 시 markRead 호출 훅**

`use-chats.ts`에 `useMarkRead` 훅 추가:

```typescript
export function useMarkRead(chatId: string, messages: Message[]) {
  const lastMsg = messages.filter((m) => m.senderUserId && m.status !== "sending").at(-1);
  useEffect(() => {
    if (chatId && lastMsg?.messageId) {
      apiClient.markRead(chatId, lastMsg.messageId).catch(() => {});
    }
  }, [chatId, lastMsg?.messageId]);
}
```

채팅방 페이지 (`chats/[id]/page.tsx`, `chats/page.tsx`)에서 호출.

- [ ] **Step 4: ChatMessage에 읽음 표시 추가**

내 메시지 중 `status !== 'sending' && status !== 'failed'`인 경우, `counterpartyLastReadMessageId` 비교:

이 단계에서는 간단하게 체크마크만 표시. counterparty 읽음 상태는 chats 목록의 캐시 또는 별도 state에서 관리.

현재 단계에서는 `✓` (sent)만 표시, 읽음(`✓✓`)은 SSE read_receipt 이벤트 수신 시 UI 갱신으로 처리.

- [ ] **Step 5: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 6: Phase 1 전체 커밋**

```
feat(web): 읽음 표시 — markRead API 호출 + SSE read_receipt + ✓ 아이콘
```

- [ ] **Step 7: Phase 1 배포**

Run: `cd /Users/jym/github-workspace/lincle && bash deploy/deploy.sh`

---

## Phase 2: Conversation Readability

### Task 10: 메시지 그루핑 로직 + ChatMessage 그룹 props

**Files:**
- Modify: `web/components/chat/chat-message.tsx` (그룹 props + 모서리 변형)
- Modify: `web/app/chats/[id]/page.tsx` (그루핑 계산)
- Modify: `web/components/chat/chat-panel.tsx` (그루핑 계산)

- [ ] **Step 1: ChatMessageProps에 그룹 플래그 추가**

```typescript
interface ChatMessageProps {
  message: Message;
  isMine: boolean;
  onRetry?: (message: Message) => void;
  isFirstInGroup?: boolean;  // NEW
  isLastInGroup?: boolean;   // NEW
}
```

- [ ] **Step 2: 그루핑 계산 유틸 함수**

```typescript
function computeGroupFlags(messages: Message[]) {
  return messages.map((msg, i) => {
    const prev = messages[i - 1];
    const next = messages[i + 1];

    const sameAsPrev =
      prev &&
      prev.senderUserId === msg.senderUserId &&
      prev.messageType !== "system" &&
      msg.messageType !== "system" &&
      new Date(msg.sentAt).getTime() - new Date(prev.sentAt).getTime() < 60_000;

    const sameAsNext =
      next &&
      next.senderUserId === msg.senderUserId &&
      next.messageType !== "system" &&
      msg.messageType !== "system" &&
      new Date(next.sentAt).getTime() - new Date(msg.sentAt).getTime() < 60_000;

    return {
      ...msg,
      isFirstInGroup: !sameAsPrev,
      isLastInGroup: !sameAsNext,
    };
  });
}
```

- [ ] **Step 3: ChatMessage 렌더링에 그루핑 적용**

- 닉네임: `isFirstInGroup && !isMine`일 때만 표시
- 간격: `isFirstInGroup ? 'mt-2' : 'mt-0.5'`
- 타임스탬프: `isLastInGroup`일 때만 표시
- 버블 모서리: 그룹 중간 메시지는 연결부 4px

```tsx
const bubbleRadius = isMine
  ? `rounded-xl ${!isLastInGroup ? 'rounded-br' : 'rounded-br-none'} ${!isFirstInGroup ? 'rounded-tr' : ''}`
  : `rounded-xl ${!isLastInGroup ? 'rounded-bl' : 'rounded-bl-none'} ${!isFirstInGroup ? 'rounded-tl' : ''}`;
```

- [ ] **Step 4: 부모 컴포넌트에서 computeGroupFlags 적용**

`chats/[id]/page.tsx`와 `chat-panel.tsx`의 메시지 렌더링부에서:

```tsx
const groupedMessages = computeGroupFlags(messages);
// ...
{groupedMessages.map((msg) => (
  <ChatMessage
    key={msg.messageId}
    message={msg}
    isMine={msg.senderUserId === myUserId}
    isFirstInGroup={msg.isFirstInGroup}
    isLastInGroup={msg.isLastInGroup}
    onRetry={...}
  />
))}
```

- [ ] **Step 5: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 6: 커밋**

```
feat(web): 메시지 그루핑 — 연속 버블 합치기, 간격 축소, 모서리 변형
```

---

### Task 11: 날짜 구분선

**Files:**
- Modify: `web/app/chats/[id]/page.tsx` (날짜 구분선 삽입)
- Modify: `web/components/chat/chat-panel.tsx` (날짜 구분선 삽입)

- [ ] **Step 1: formatDateSeparator 유틸 함수**

```typescript
function formatDateSeparator(dateStr: string): string {
  const date = new Date(dateStr);
  const now = new Date();
  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
  const msgDay = new Date(date.getFullYear(), date.getMonth(), date.getDate());
  const diffDays = Math.floor((today.getTime() - msgDay.getTime()) / 86400000);

  if (diffDays === 0) return "오늘";
  if (diffDays === 1) return "어제";
  if (date.getFullYear() === now.getFullYear()) {
    return `${date.getMonth() + 1}월 ${date.getDate()}일`;
  }
  return `${date.getFullYear()}년 ${date.getMonth() + 1}월 ${date.getDate()}일`;
}

function isSameDay(a: string, b: string): boolean {
  const da = new Date(a);
  const db = new Date(b);
  return da.getFullYear() === db.getFullYear()
    && da.getMonth() === db.getMonth()
    && da.getDate() === db.getDate();
}
```

- [ ] **Step 2: DateSeparator 컴포넌트 (인라인)**

```tsx
function DateSeparator({ label }: { label: string }) {
  return (
    <div className="flex items-center gap-3 my-3">
      <div className="flex-1 border-t border-border" />
      <span className="text-xs text-text-secondary">{label}</span>
      <div className="flex-1 border-t border-border" />
    </div>
  );
}
```

- [ ] **Step 3: 메시지 목록 렌더링에 날짜 구분선 삽입**

```tsx
{groupedMessages.map((msg, i) => {
  const prev = groupedMessages[i - 1];
  const showDateSep = !prev || !isSameDay(prev.sentAt, msg.sentAt);
  return (
    <Fragment key={msg.messageId}>
      {showDateSep && <DateSeparator label={formatDateSeparator(msg.sentAt)} />}
      <ChatMessage ... />
    </Fragment>
  );
})}
```

- [ ] **Step 4: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 5: 커밋**

```
feat(web): 채팅 날짜 구분선 — 오늘/어제/날짜 표시
```

---

### Task 12: 새 메시지 구분선

**Files:**
- Modify: `web/app/chats/[id]/page.tsx` (새 메시지 구분선)
- Modify: `web/components/chat/chat-panel.tsx` (새 메시지 구분선)

- [ ] **Step 1: 새 메시지 구분선 컴포넌트**

```tsx
function NewMessagesDivider() {
  return (
    <div className="flex items-center gap-3 my-3">
      <div className="flex-1 border-t border-gold" />
      <span className="text-xs text-gold font-medium">새 메시지</span>
      <div className="flex-1 border-t border-gold" />
    </div>
  );
}
```

- [ ] **Step 2: myLastReadMessageId 기반 구분선 삽입**

채팅방 컴포넌트에서:

```tsx
// chats 목록에서 현재 방의 myLastReadMessageId 획득
const currentChat = chats.find((c) => c.chatRoomId === activeChatId);
const lastReadId = useRef(currentChat?.myLastReadMessageId);

// 렌더링 시:
{groupedMessages.map((msg, i) => {
  const showNewMsgDivider =
    lastReadId.current &&
    msg.messageId !== lastReadId.current &&
    groupedMessages[i - 1]?.messageId === lastReadId.current;

  return (
    <Fragment key={msg.messageId}>
      {showDateSep && <DateSeparator ... />}
      {showNewMsgDivider && <NewMessagesDivider />}
      <ChatMessage ... />
    </Fragment>
  );
})}
```

- [ ] **Step 3: 새 메시지 위치로 스크롤**

```tsx
const newMsgRef = useRef<HTMLDivElement>(null);
useEffect(() => {
  if (newMsgRef.current) {
    newMsgRef.current.scrollIntoView({ behavior: "smooth", block: "center" });
  } else {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }
}, [activeChatId]);
```

NewMessagesDivider에 ref 연결: `<div ref={newMsgRef}><NewMessagesDivider /></div>`

- [ ] **Step 4: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 5: Phase 2 커밋 + 배포**

```
feat(web): 새 메시지 구분선 — gold 라인 + 자동 스크롤
```

Run: `cd /Users/jym/github-workspace/lincle && bash deploy/deploy.sh`

---

## Phase 3: Trade Flow Integration

### Task 13: 백엔드 채팅 목록에 매물 가격/서버 필드 추가

**Files:**
- Modify: `backend/internal/repository/postgres_chat.go` (SQL + struct)
- Modify: `backend/internal/repository/interfaces.go` (struct)
- Modify: `backend/cmd/server/handlers_chat.go` (응답)

- [ ] **Step 1: ChatRoomListItem에 가격/서버 필드 추가**

```go
ListingPrice      *int64  // NEW
ListingServerID   *string // NEW
ListingServerName *string // NEW
```

- [ ] **Step 2: SQL에 listings JOIN 필드 추가**

기존 `JOIN listings l` 에서 추가 필드 SELECT:
```sql
l.price_amount, l.server_id,
(SELECT s.name FROM servers s WHERE s.id = l.server_id) as server_name
```

- [ ] **Step 3: 핸들러 응답에 필드 추가**

```go
"listingPrice":      item.ListingPrice,
"listingServerId":   item.ListingServerID,
"listingServerName": item.ListingServerName,
```

- [ ] **Step 4: 프론트 ChatRoom 타입 확장**

```typescript
listingPrice?: number;
listingServerId?: string;
listingServerName?: string;
```

- [ ] **Step 5: 빌드 + 테스트**

Run: `cd backend && go build ./cmd/server/ && go test ./...`
Run: `cd web && npx next build`

- [ ] **Step 6: 커밋**

```
feat: 채팅 목록에 매물 가격/서버 정보 추가
```

---

### Task 14: 매물 정보 카드 (채팅 헤더)

**Files:**
- Modify: `web/components/chat/chat-panel.tsx` (데스크톱 헤더)
- Modify: `web/app/chats/[id]/page.tsx` (모바일 헤더)

- [ ] **Step 1: ListingInfoCard 컴포넌트 작성**

`chat-panel.tsx` 내 인라인 또는 별도 파일:

```tsx
function ListingInfoCard({ chat }: { chat: ChatRoom }) {
  const statusConfig: Record<string, { label: string; color: string; dim?: boolean }> = {
    available: { label: "거래 가능", color: "bg-green-600" },
    reserved: { label: "예약중", color: "bg-yellow-600" },
    sold: { label: "거래 완료", color: "bg-medium", dim: true },
    deleted: { label: "삭제됨", color: "bg-medium", dim: true },
  };
  const status = statusConfig[chat.listingStatus] ?? statusConfig.available;

  return (
    <Link
      href={`/listings/${chat.listingId}`}
      className={`flex items-center gap-3 px-4 py-2 border-b border-border hover:bg-medium/50 ${status.dim ? "opacity-50" : ""}`}
    >
      {chat.listingThumbnail ? (
        <Image src={chat.listingThumbnail} alt="" width={40} height={40} className="rounded object-cover" unoptimized />
      ) : (
        <div className="w-10 h-10 rounded bg-medium flex items-center justify-center text-text-secondary text-xs">?</div>
      )}
      <div className="flex-1 min-w-0">
        <p className="font-medium text-sm truncate">{chat.listingTitle}</p>
        <p className="text-xs text-text-secondary truncate">
          {chat.listingPrice != null && `${chat.listingPrice.toLocaleString()} 아덴`}
          {chat.listingServerName && ` · ${chat.listingServerName}`}
        </p>
      </div>
      <span className={`text-xs px-2 py-0.5 rounded-full text-white ${status.color}`}>
        {status.label}
      </span>
    </Link>
  );
}
```

- [ ] **Step 2: 데스크톱 chat-panel.tsx에 ListingInfoCard 삽입**

메시지 영역 상단, 기존 gold accent bar 대신:

```tsx
{activeChat && <ListingInfoCard chat={activeChat} />}
```

- [ ] **Step 3: 모바일 chats/[id]/page.tsx에도 삽입**

chats 목록에서 현재 채팅방 찾아서 전달. 또는 별도 API 호출.

- [ ] **Step 4: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 5: 커밋**

```
feat(web): 채팅 헤더에 매물 정보 카드 추가
```

---

### Task 15: 채팅 리스트 거래 상태 배지

**Files:**
- Modify: `web/components/chat/chat-list-item.tsx` (상태 배지)

- [ ] **Step 1: 상태 배지 추가**

`ChatListItem` 내부, 시간 아래 영역에:

```tsx
{chat.chatStatus !== "open" && (
  <span className={`text-xs px-1.5 py-0.5 rounded mt-1 inline-block ${
    chat.chatStatus === "reservation_proposed" ? "bg-yellow-600/20 text-yellow-400" :
    chat.chatStatus === "reservation_confirmed" ? "bg-green-600/20 text-green-400" :
    chat.chatStatus === "deal_completed" ? "bg-gold/20 text-gold" : ""
  }`}>
    {chat.chatStatus === "reservation_proposed" && "예약 제안"}
    {chat.chatStatus === "reservation_confirmed" && "예약 확정"}
    {chat.chatStatus === "deal_completed" && "거래 완료"}
  </span>
)}
```

- [ ] **Step 2: 빌드 검증**

Run: `cd web && npx next build`

- [ ] **Step 3: 커밋**

```
feat(web): 채팅 리스트에 거래 상태 배지 추가
```

---

### Task 16: 예약 카드 메시지 렌더링

**Files:**
- Create: `web/components/chat/reservation-card-message.tsx`
- Modify: `web/components/chat/chat-message.tsx` (reservation_card 분기)

- [ ] **Step 1: ReservationCardMessage 컴포넌트 생성**

```tsx
interface ReservationCardProps {
  message: Message;
  isMine: boolean;
}

export function ReservationCardMessage({ message, isMine }: ReservationCardProps) {
  const meta = message.metadataJson ?? {};
  const meetingType = meta.meetingType === "offline_pc_bang" ? "오프라인 PC방" : "인게임 거래";

  return (
    <div className={`flex mb-2 ${isMine ? "justify-end" : "justify-start"}`}>
      <div className="max-w-[70%] border border-gold rounded-lg p-4 bg-card">
        <p className="text-xs text-gold font-medium mb-2">예약 제안</p>
        <div className="space-y-1 text-sm">
          {meta.date && <p>📅 {String(meta.date)} {meta.time && String(meta.time)}</p>}
          <p>📍 {meetingType}</p>
          {meta.meetingPoint && <p>💬 {String(meta.meetingPoint)}</p>}
          {meta.notes && <p className="text-text-secondary text-xs mt-2">{String(meta.notes)}</p>}
        </div>
        {/* 상태 표시: 향후 수락/거절 버튼은 예약 상태 API 연동 후 추가 */}
        <p className="text-xs text-text-secondary mt-3">
          {isMine ? "예약 제안을 보냈습니다" : "예약 제안을 받았습니다"}
        </p>
      </div>
    </div>
  );
}
```

- [ ] **Step 2: ChatMessage에서 reservation_card 분기**

```tsx
if (message.messageType === "reservation_card") {
  return <ReservationCardMessage message={message} isMine={isMine} />;
}
```

system 메시지 분기 바로 아래에 추가.

- [ ] **Step 3: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 4: Phase 3 커밋 + 배포**

```
feat(web): 예약 카드 메시지 렌더링 — metadataJson 기반
```

Run: `cd /Users/jym/github-workspace/lincle && bash deploy/deploy.sh`

---

## Phase 4: Real-time Stability

### Task 17: SSE 무한 재연결 + heartbeat 타임아웃

**Files:**
- Modify: `web/lib/hooks/use-sse.ts` (전면 개선)

- [ ] **Step 1: 최대 재시도 제한 제거 + backoff 상한 60초**

기존 `maxRetries` 관련 로직 제거. backoff 계산:

```typescript
const delay = Math.min(1000 * Math.pow(2, retryCountRef.current), 60_000);
```

연결 성공 시 리셋:

```typescript
es.onopen = () => {
  retryCountRef.current = 0;
  setStatus("connected");
};
```

- [ ] **Step 2: heartbeat 타임아웃 메커니즘 추가**

```typescript
const lastEventRef = useRef<number>(Date.now());
const heartbeatCheckRef = useRef<ReturnType<typeof setInterval>>();

// 모든 이벤트 수신 시:
lastEventRef.current = Date.now();

// 45초 타임아웃 체크 (서버 heartbeat 30초 + 15초 여유):
heartbeatCheckRef.current = setInterval(() => {
  if (Date.now() - lastEventRef.current > 45_000) {
    es.close();
    reconnect();
  }
}, 10_000);
```

- [ ] **Step 3: visibilitychange + online 이벤트 핸들러**

```typescript
useEffect(() => {
  const handleVisibility = () => {
    if (document.visibilityState === "visible" && statusRef.current !== "connected") {
      reconnect();
    }
  };
  const handleOnline = () => {
    if (statusRef.current !== "connected") {
      reconnect();
    }
  };

  document.addEventListener("visibilitychange", handleVisibility);
  window.addEventListener("online", handleOnline);
  return () => {
    document.removeEventListener("visibilitychange", handleVisibility);
    window.removeEventListener("online", handleOnline);
  };
}, []);
```

- [ ] **Step 4: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 5: 커밋**

```
feat(web): SSE 무한 재연결 + heartbeat 타임아웃 + 탭 복귀 감지
```

---

### Task 18: 연결 상태 UI 개선

**Files:**
- Modify: `web/components/chat/chat-panel.tsx` (상태바 이동)
- Modify: `web/components/chat/chat-input.tsx` (disabled 상태)

- [ ] **Step 1: chat-panel.tsx 상단 alert 제거, 입력 상단으로 이동**

기존 SSE status alert 코드 제거. 입력 영역 바로 위에 새 상태 바:

```tsx
{connectionStatus === "reconnecting" && (
  <div className="px-4 py-1.5 bg-yellow-600/20 text-yellow-400 text-xs flex items-center gap-2">
    <span className="animate-spin">⟳</span> 연결 중...
  </div>
)}
{connectionStatus === "disconnected" && (
  <div className="px-4 py-1.5 bg-red-600/20 text-red-400 text-xs flex items-center justify-between">
    <span>연결 끊김</span>
    <button onClick={() => window.location.reload()} className="underline">다시 연결</button>
  </div>
)}
```

- [ ] **Step 2: ChatInput에 disabled prop 지원**

기존 `disabled` prop이 이미 있으면 활용. 없으면 추가:

```tsx
<ChatInput
  onSend={...}
  disabled={connectionStatus === "disconnected"}
/>
```

ChatInput 내부에서 disabled일 때 placeholder 변경:

```tsx
placeholder={disabled ? "연결이 끊겼습니다" : "메시지를 입력하세요..."}
```

- [ ] **Step 3: 빌드 + 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 4: 커밋**

```
feat(web): 채팅 연결 상태 UI — 입력창 상단 바 + disabled 상태
```

---

### Task 19: 메시지 페이지네이션 (useInfiniteQuery)

**Files:**
- Modify: `web/lib/hooks/use-chats.ts` (useMessages → useInfiniteQuery)
- Modify: `web/app/chats/[id]/page.tsx` (스크롤 + 데이터 소비)
- Modify: `web/components/chat/chat-panel.tsx` (스크롤 + 데이터 소비)
- Modify: `web/lib/api-client.ts` (getMessages cursor 파라미터)

- [ ] **Step 1: apiClient.getMessages에 cursor 파라미터 추가**

```typescript
async getMessages(
  chatId: string,
  cursor?: string
): Promise<PaginatedResponse<Message>> {
  const params = cursor ? `?cursor=${cursor}` : "";
  return this.fetch(`/chats/${chatId}/messages${params}`);
}
```

- [ ] **Step 2: useMessages를 useInfiniteQuery로 전환**

```typescript
export function useMessages(chatId: string, sseConnected: boolean) {
  return useInfiniteQuery({
    queryKey: ["messages", chatId],
    queryFn: ({ pageParam }: { pageParam?: string }) =>
      apiClient.getMessages(chatId, pageParam),
    initialPageParam: undefined as string | undefined,
    getNextPageParam: (lastPage) =>
      lastPage.cursor.hasMore ? lastPage.cursor.next : undefined,
    enabled: !!chatId,
    refetchInterval: sseConnected ? false : 5_000, // SSE 연결 시 폴링 OFF
  });
}
```

- [ ] **Step 3: 데이터 소비부 업데이트**

```typescript
const { data, fetchNextPage, hasNextPage, isFetchingNextPage } = useMessages(chatId, sseConnected);

// 페이지를 reverse하여 시간순 정렬
const messages = data?.pages
  ? data.pages.flatMap((p) => [...p.data]).reverse()
  : [];
```

- [ ] **Step 4: 스크롤 상단 도달 시 이전 메시지 로드**

```typescript
const scrollRef = useRef<HTMLDivElement>(null);

const handleScroll = () => {
  const el = scrollRef.current;
  if (!el) return;
  if (el.scrollTop < 50 && hasNextPage && !isFetchingNextPage) {
    const prevHeight = el.scrollHeight;
    const prevTop = el.scrollTop;
    fetchNextPage().then(() => {
      requestAnimationFrame(() => {
        if (el) {
          el.scrollTop = el.scrollHeight - prevHeight + prevTop;
        }
      });
    });
  }
};
```

```tsx
<div ref={scrollRef} onScroll={handleScroll} className="flex-1 overflow-y-auto ...">
  {isFetchingNextPage && (
    <div className="text-center py-2 text-text-secondary text-xs">이전 메시지 불러오는 중...</div>
  )}
  {/* messages */}
</div>
```

- [ ] **Step 5: optimistic update를 infiniteQuery 캐시 구조에 맞게 수정**

`useSendMessage`의 `onMutate`, `onSuccess`, `onError`에서 `setQueryData`를 `InfiniteData` 구조에 맞게 수정:

onMutate, onSuccess, onError 모두 `InfiniteData` 캐시 구조에 맞게 수정:

```typescript
// onMutate: 최신 페이지(첫 번째)에 prepend
qc.setQueryData<InfiniteData<PaginatedResponse<Message>>>(
  ["messages", chatId],
  (old) => {
    if (!old) return old;
    const firstPage = old.pages[0];
    return {
      ...old,
      pages: [
        { ...firstPage, data: [optimisticMsg, ...firstPage.data] },
        ...old.pages.slice(1),
      ],
    };
  }
);

// onSuccess: 서버 응답으로 교체
qc.setQueryData<InfiniteData<PaginatedResponse<Message>>>(
  ["messages", chatId],
  (old) => {
    if (!old) return old;
    return {
      ...old,
      pages: old.pages.map((page) => ({
        ...page,
        data: page.data.map((m) =>
          m.messageId === clientMessageId
            ? { ...serverMsg, status: "sent" as const }
            : m
        ),
      })),
    };
  }
);

// onError: failed 상태로 표시
qc.setQueryData<InfiniteData<PaginatedResponse<Message>>>(
  ["messages", chatId],
  (old) => {
    if (!old) return old;
    return {
      ...old,
      pages: old.pages.map((page) => ({
        ...page,
        data: page.data.map((m) =>
          m.messageId === clientMessageId
            ? { ...m, status: "failed" as const }
            : m
        ),
      })),
    };
  }
);
```

- [ ] **Step 6: 빌드 + 전체 테스트**

Run: `cd web && npx next build && npx vitest run`

- [ ] **Step 7: Phase 4 커밋 + 최종 배포**

```
feat(web): 메시지 페이지네이션 — useInfiniteQuery + 스크롤 위치 보존
```

- [ ] **Step 8: E2E 테스트**

Run: `cd web && npx playwright test`

- [ ] **Step 9: 최종 배포**

Run: `cd /Users/jym/github-workspace/lincle && bash deploy/deploy.sh`
