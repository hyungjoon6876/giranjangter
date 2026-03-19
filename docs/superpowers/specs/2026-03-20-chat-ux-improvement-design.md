# Chat UX Improvement Design

## Overview

기란JT 채팅 기능의 전면적 UX 개선. 현재 MVP 수준의 채팅을 거래 플랫폼에 걸맞은 수준으로 끌어올린다.

**접근**: Bottom-Up (기반 먼저) — 4 Phase로 나누어 각 Phase 독립 배포 가능.

## Current State Analysis

### 현재 구현

- **채팅 목록**: 데스크톱 split-panel (좌: 목록 w-72, 우: 메시지), 모바일 리스트 전용
- **메시지**: 버블 스타일 (내 메시지 blue, 상대 card), 시스템 메시지 중앙 pill
- **실시간**: SSE + 5초 폴링 fallback, 최대 10회 재시도
- **상태관리**: TanStack Query + SSE 이벤트로 invalidation
- **입력**: textarea auto-height (`text-sm` = 14px), Enter 전송, Shift+Enter 줄바꿈

### 주요 문제점

1. **모바일 줌 버그**: iOS Safari에서 채팅 입력 포커스 시 화면 확대 (font-size 14px < 16px)
2. **메시지 피드백 없음**: 타임스탬프, 전송/읽음 표시 모두 부재
3. **가독성 부족**: 메시지 그루핑 없음, 날짜 구분선 없음
4. **거래 흐름 단절**: 채팅 내 매물 정보/거래 상태 미표시, 예약 카드 미렌더링
5. **실시간 불안정**: SSE 10회 실패 후 영구 끊김, 재연결 액션 없음
6. **Optimistic update 없음**: 메시지 전송 후 서버 응답까지 2-5초 지연
7. **미활용 백엔드 기능**: `markRead` API, `clientMessageId` 중복 방지, cursor 페이지네이션
8. **백엔드 응답 갭**: `handleListChats`가 `unreadCount`, `lastMessage`, `listingThumbnail`, `listingStatus`를 반환하지 않음 (프론트 타입에는 정의되어 있으나 실제 API 미지원)

### Pre-requisite: Backend API Gap Fix

Phase 1 시작 전, `handleListChats` 응답을 프론트엔드 `ChatRoom` 타입에 맞게 확장한다.

**백엔드 변경** (`handlers_chat.go` + `postgres_chat.go`):
- `ListChatRooms` SQL에 JOIN 추가:
  - `chat_read_cursors` 테이블 → `unreadCount` 계산 (현재 유저의 읽지 않은 메시지 수)
  - 서브쿼리 → `lastMessage` (마지막 메시지 body + sentAt)
  - `listings` 테이블 → `listingThumbnail`, `listingStatus` 필드
- `ChatRoomListItem` 구조체 확장
- 응답에 `myLastReadMessageId` 필드 추가 (Phase 2 새 메시지 구분선에서 사용)

**프론트 변경** (`web/lib/types.ts`):
- `ChatRoom` 타입에 `myLastReadMessageId?: string` 추가

이 작업은 Phase 1의 첫 번째 태스크로 처리한다.

---

## Phase 1: Mobile Fix + Message Feedback

### 1-1. Mobile Input Zoom Prevention

**문제**: iOS Safari는 font-size < 16px인 input 포커스 시 자동 줌. 현재 `chat-input.tsx` textarea는 `text-sm` (14px).

**변경**:
- `web/app/layout.tsx`: Next.js App Router의 `viewport` export 사용:
  ```typescript
  export const viewport: Viewport = {
    width: 'device-width',
    initialScale: 1,
    maximumScale: 1,
  }
  ```
- `web/components/chat/chat-input.tsx`: textarea class `text-sm` → `text-base` (16px)

### 1-2. Message Timestamps

**변경 파일**: `web/components/chat/chat-message.tsx`

각 메시지 버블 하단에 시간 표시:
- 내 메시지: 버블 우하단, `text-xs text-text-secondary`
- 상대 메시지: 버블 좌하단, 동일 스타일
- 시스템 메시지: 표시 안 함

포맷 규칙:
- 오늘: `오후 3:42`
- 올해: `3월 15일 오후 3:42`
- 그 외: `2025.3.15 오후 3:42`

유틸 함수 `formatMessageTime(sentAt: string): string` 생성.

참고: 날짜 구분선 (Phase 2-2)은 `2025년 3월 15일` (한글 형식) 사용. 타임스탬프의 compact 형식(`2025.3.15`)과는 의도적으로 구분.

### 1-3. Optimistic Update

**변경 파일**: `web/lib/hooks/use-chats.ts`, `web/lib/api-client.ts`, `web/lib/types.ts`

현재 `useSendMessage`는 서버 응답 후에야 메시지가 표시됨.

**타입 변경** (`web/lib/types.ts`):
- `Message` 인터페이스에 `status?: 'sending' | 'sent' | 'failed'` 필드 추가 (클라이언트 전용, 서버 응답에는 없음)

**API 클라이언트 변경** (`web/lib/api-client.ts`):
- `sendMessage` 시그니처 변경: `sendMessage(chatId: string, text: string, clientMessageId: string): Promise<Message>`
- 요청 body에 `clientMessageId` 포함

**Hook 변경** (`web/lib/hooks/use-chats.ts`):
- `onMutate` 콜백에서 즉시 캐시에 메시지 추가:
  - 임시 `messageId` = `clientMessageId`
  - `chatRoomId` = 현재 채팅방 ID (훅 인자에서 획득)
  - `status: 'sending'`
- `onSuccess`: 서버 응답으로 교체 (`status: 'sent'`)
- `onError`: 메시지 `status: 'failed'` + 이전 캐시 스냅샷 롤백
- `clientMessageId` (UUID) 생성하여 전송 — 백엔드 중복 방지 활용

메시지 상태 표시 (`chat-message.tsx`):
- `sending`: 시간 대신 회색 "전송 중..."
- `sent` 또는 `undefined` (서버 메시지): 타임스탬프 + 체크 1개
- `failed`: 빨간색 "전송 실패 · 재전송" 버튼 (클릭 시 재전송)

### 1-4. Read Receipts (읽음 표시)

**프론트 변경**: `web/lib/hooks/use-chats.ts`, `web/lib/api-client.ts`, `web/components/chat/chat-message.tsx`, `web/lib/hooks/use-sse.ts`
**백엔드 변경**: `backend/cmd/server/handlers_chat.go`, `backend/cmd/server/main.go`

**백엔드 구조 변경**:
- `handleMarkRead` 함수 시그니처에 `broker *event.Broker` 파라미터 추가
- `main.go` 라우트 등록에서 broker 전달
- `handleMarkRead` 내부에서 `repo.GetChatRoomParticipants` 호출하여 상대방 ID 확인
- 상대에게 SSE `read_receipt` 이벤트 브로드캐스트
- 이벤트 페이로드: `{ chatRoomId, readBy, lastReadMessageId, readAt }`

**읽음 상태 결정 로직**:
- 백엔드는 `chat_read_cursors` 테이블에 `last_read_message_id`를 per-room per-user로 저장
- SSE `read_receipt` 페이로드에 `lastReadMessageId` 포함
- 프론트에서 내 메시지의 읽음 상태 판단: `message.messageId <= counterpartyLastReadMessageId` → 읽음

**프론트**:
- 채팅방 진입 시 `POST /chats/:chatId/read` 호출
- 새 메시지 수신 시 채팅방이 열려있으면 자동으로 `markRead` 호출
- 내 메시지 하단에 읽음 상태 아이콘:
  - 체크 1개 (전송됨, 아직 읽지 않음)
  - 체크 2개 (읽음) — 색상: gold

**프론트 SSE 핸들러** (`use-sse.ts`):
- `read_receipt` 이벤트 수신 시 해당 채팅방의 `counterpartyLastReadMessageId` 상태 업데이트
- 메시지 캐시에서 해당 ID 이하 메시지들의 읽음 상태 갱신

---

## Phase 2: Conversation Readability

### 2-1. Message Grouping

**변경 파일**: `web/components/chat/chat-message.tsx`, 부모 컴포넌트 (메시지 목록 렌더링부)

그루핑 규칙:
- 같은 발신자 + 이전 메시지와 1분 이내 간격 → 같은 그룹
- 시스템 메시지가 끼면 그룹 분리

그룹 내 렌더링:
- 첫 메시지만 닉네임 표시 (상대 메시지의 경우)
- 중간 메시지: 버블 간격 2px (일반 8px)
- 마지막 메시지에만 타임스탬프 표시
- 버블 모서리 변형:
  - 그룹 첫: 정상 radius
  - 그룹 중간: 연결부 4px radius (나머지 12px)
  - 그룹 끝: 정상 radius

로직: 메시지 배열을 순회하며 `isFirstInGroup`, `isLastInGroup` 플래그를 계산하여 `ChatMessage` 컴포넌트에 전달.

### 2-2. Date Separators

**변경 파일**: 메시지 목록 렌더링부

메시지 배열 순회 시 이전 메시지와 날짜가 다르면 구분선 삽입:
- 오늘: `오늘`
- 어제: `어제`
- 올해: `3월 15일`
- 그 외: `2025년 3월 15일`

스타일: `flex items-center gap-3` + 좌우 `border-t border-border flex-1` + 중앙 텍스트 `text-xs text-text-secondary`

### 2-3. New Messages Divider

**변경 파일**: 메시지 목록 렌더링부, `web/lib/hooks/use-chats.ts`

**데이터 소스**: Pre-requisite에서 추가한 `ChatRoom.myLastReadMessageId` 사용.
- 채팅방 진입 시, `GET /chats` 응답의 `myLastReadMessageId`를 초기값으로 저장
- 메시지 목록에서 해당 ID 바로 다음에 `새 메시지` 구분선 삽입
- 스타일: gold 색상 라인 + "새 메시지" 텍스트
- 구분선 위치로 자동 스크롤 (맨 아래가 아닌 새 메시지 시작점)
- `markRead` 호출 후에도 현재 세션에서는 구분선 유지; 다음 방문 시 새 위치로 이동

---

## Phase 3: Trade Flow Integration

### 3-1. Listing Info Card (Chat Header)

**변경 파일**: `web/components/chat/chat-panel.tsx`, `web/app/chats/[id]/page.tsx`
**백엔드 변경**: `handlers_chat.go` — `handleListChats` 응답에 `listingPrice`, `listingServerId`, `listingServerName` 필드 추가 (Pre-requisite에서 추가한 JOIN 확장)

채팅방 상단에 컴팩트 매물 카드:
- 왼쪽: 썸네일 이미지 (40x40, rounded). `listingThumbnail`이 없으면 기본 아이콘
- 중앙: 아이템명 (bold) + 가격/서버 (secondary text)
- 우측: 매물 상태 배지
- 전체 영역 클릭 → 매물 상세 페이지 이동 (`/listings/:listingId`)

매물 상태별 표시:
- `available`: 초록 "거래 가능"
- `reserved`: 노란 "예약중"
- `sold`: 회색 "거래 완료" + 카드 opacity 50%
- `deleted`: 회색 "삭제됨" + 카드 opacity 50%

### 3-2. Chat Status Badge

**변경 파일**: `web/components/chat/chat-list-item.tsx`

채팅 리스트 아이템에 거래 상태 배지 추가:
- `open`: 배지 없음
- `reservation_proposed`: 노란 배경 "예약 제안"
- `reservation_confirmed`: 초록 배경 "예약 확정"
- `deal_completed`: gold 배경 + 체크 "거래 완료"

위치: 채팅 리스트 아이템의 시간 아래 (우측 영역)

### 3-3. Reservation Card Message

**새 파일**: `web/components/chat/reservation-card-message.tsx`
**변경 파일**: `web/components/chat/chat-message.tsx`

`messageType === 'reservation_card'`일 때 전용 카드 렌더링.

참고: `reservation_card` 메시지는 사용자가 직접 전송하는 것이 아니라, 예약 생성 API (`POST /chats/:chatId/reservations`)를 통해 시스템이 생성하는 메시지다. `handleSendMessage`의 `messageType` 검증 (`oneof=text image`)과는 별개 경로.

카드 내용 (`metadataJson`에서 추출):
- 날짜/시간
- 거래 방식 (인게임 / 오프라인 PC방)
- 장소
- 메모 (있으면)

카드 하단 액션:
- 상대에게만 "수락" / "거절" 버튼 표시
- 이미 처리됨: "수락됨" / "거절됨" 상태 텍스트 (버튼 대신)
- 본인이 보낸 제안: "예약 제안을 보냈습니다" + 대기 상태

스타일: border-gold, rounded-lg, 일반 버블과 구분되는 카드형 레이아웃

---

## Phase 4: Real-time Stability

### 4-1. Infinite SSE Reconnection

**변경 파일**: `web/lib/hooks/use-sse.ts`

현재 문제: `maxRetries = 10` 이후 영구 끊김.

변경:
- 최대 재시도 제한 제거 (무한 재시도)
- backoff 상한: 60초
- `document.visibilitychange` 핸들러: 탭 복귀(visible) 시 즉시 재연결 시도
- `window.online` 이벤트: 네트워크 복귀 시 즉시 재연결
- 연결 성공 시 backoff 카운터 및 delay 리셋

**Heartbeat timeout 메커니즘 추가**:
- 현재 클라이언트에는 heartbeat 타임아웃 감지가 없음 (EventSource.onerror에만 의존)
- 새로 추가: 마지막 이벤트 수신 시각 추적, 45초간 이벤트가 없으면 연결 stale로 판단하고 EventSource를 닫고 재연결 (서버 heartbeat 간격 30초 + 15초 여유)

**멀티 탭 주의**: SSE broker는 유저당 1개 연결만 지원 (이전 연결 자동 해제). 두 탭이 동시에 재연결하면 루프 발생 가능. 이번 스코프에서는 별도 처리 안 함 (Out of Scope에 멀티 탭 기재). 단, reconnect backoff가 있으므로 무한 루프는 아님.

### 4-2. Connection Status UI

**변경 파일**: `web/components/chat/chat-panel.tsx`, `web/components/chat/chat-input.tsx`

현재: 채팅 패널 상단에 alert 바 (텍스트만).

변경 — 상태 표시를 입력창 바로 위로 이동:
- `connected`: 표시 없음
- `reconnecting`: 노란 배경 바 "연결 중..." + 스피너 (입력 가능 유지)
- `disconnected`: 빨간 배경 바 "연결 끊김" + "다시 연결" 버튼
  - 입력 textarea disabled + placeholder "연결이 끊겼습니다"
  - "다시 연결" 클릭 시 즉시 SSE 재연결 시도

### 4-3. Message Pagination

**변경 파일**: `web/lib/hooks/use-chats.ts`

현재: 메시지 전체 한 번에 로드 + 5초 폴링.

변경:
- `useMessages` → `useInfiniteQuery` 전환
- 초기 50건 로드 (백엔드 기본값, `ORDER BY sent_at DESC` → 최신 50건)
- 스크롤 상단 도달 시 `fetchNextPage` (이전 메시지 로드)
- cursor: 가장 오래된 메시지의 `sentAt`

**메시지 정렬 처리**:
- 백엔드는 `sent_at DESC` (최신순)으로 반환
- 클라이언트에서 각 페이지를 reverse하여 시간순(오래된→최신)으로 표시
- `useInfiniteQuery`의 `pages` 배열: `[page3(oldest), page2, page1(newest)]` 순서로 `flatMap` + reverse

**스크롤 위치 보존**:
- 이전 메시지 로드 시 스크롤 점프 방지: 로드 전 `scrollTop`과 `scrollHeight` 저장, 로드 후 `scrollTop = newScrollHeight - savedScrollHeight + savedScrollTop`으로 복원

- 상단 로딩 중 스피너 표시
- SSE 연결 중에는 5초 폴링 비활성화 (불필요한 요청 제거)
- SSE 끊김 시에만 폴링 fallback 활성화

---

## File Change Summary

### Pre-requisite (Phase 1 첫 번째 태스크)
| File | Change |
|------|--------|
| `backend/cmd/server/handlers_chat.go` | `handleListChats` SQL 확장 — unreadCount, lastMessage, listingThumbnail, listingStatus, myLastReadMessageId |
| `backend/internal/repository/postgres_chat.go` | `ListChatRooms` SQL JOIN 추가, `ChatRoomListItem` 구조체 확장 |
| `web/lib/types.ts` | `ChatRoom`에 `myLastReadMessageId` 추가, `Message`에 `status?` 추가 |

### Phase 1
| File | Change |
|------|--------|
| `web/app/layout.tsx` | `export const viewport: Viewport` 추가 (`maximumScale: 1`) |
| `web/components/chat/chat-input.tsx` | font-size `text-sm` → `text-base`, disabled 상태 지원 |
| `web/components/chat/chat-message.tsx` | 타임스탬프, 전송 상태 아이콘, 읽음 표시 |
| `web/lib/hooks/use-chats.ts` | optimistic update (onMutate/onSuccess/onError), markRead 호출 |
| `web/lib/api-client.ts` | `markRead` 메서드 추가, `sendMessage` 시그니처에 `clientMessageId` 추가 |
| `web/lib/hooks/use-sse.ts` | `read_receipt` 이벤트 핸들러 추가 |
| `backend/cmd/server/handlers_chat.go` | `handleMarkRead`에 `broker` 파라미터 추가, `read_receipt` SSE 브로드캐스트 |
| `backend/cmd/server/main.go` | `handleMarkRead` 라우트에 broker 전달 |

### Phase 2
| File | Change |
|------|--------|
| `web/components/chat/chat-message.tsx` | 그룹 props (`isFirstInGroup`, `isLastInGroup`), 모서리 변형 |
| `web/app/chats/[id]/page.tsx` 또는 `chat-panel.tsx` | 그루핑 로직, 날짜 구분선, 새 메시지 구분선 |

### Phase 3
| File | Change |
|------|--------|
| `web/components/chat/chat-panel.tsx` | 매물 정보 카드 헤더 |
| `web/app/chats/[id]/page.tsx` | 모바일 매물 카드 헤더 |
| `web/components/chat/chat-list-item.tsx` | 거래 상태 배지 |
| `web/components/chat/reservation-card-message.tsx` | 새 파일 — 예약 카드 메시지 컴포넌트 |
| `web/components/chat/chat-message.tsx` | `reservation_card` 타입 분기 |
| `backend/cmd/server/handlers_chat.go` | 채팅 목록 응답에 `listingPrice`, `listingServerId`, `listingServerName` 필드 추가 |
| `web/lib/types.ts` | `ChatRoom` 타입에 가격/서버 필드 추가 |

### Phase 4
| File | Change |
|------|--------|
| `web/lib/hooks/use-sse.ts` | 무한 재연결, visibility/online 이벤트, heartbeat 타임아웃 추가, backoff 리셋 |
| `web/components/chat/chat-input.tsx` | disconnected 시 disabled 상태 |
| `web/components/chat/chat-panel.tsx` | 연결 상태 UI를 입력창 상단으로 이동 |
| `web/lib/hooks/use-chats.ts` | `useInfiniteQuery` 전환, 스크롤 위치 보존, SSE 연결 시 폴링 비활성화 |

---

## Out of Scope

이번 개선에 포함하지 않는 항목:
- 파일/이미지 전송 (별도 피처)
- 타이핑 인디케이터 (백엔드 WebSocket 전환 필요)
- 메시지 수정/삭제 (별도 피처)
- 멀티 디바이스/멀티 탭 동기화 (현재 단일 SSE 구조 한계 — broker가 유저당 1개 연결만 지원)
- Flutter 앱 동기화 (현재 Web 우선 전략)
- 메시지 검색 (별도 피처)

## Dependencies

- **Pre-requisite** → Phase 1: 백엔드 API 갭 수정이 Phase 1 전체에 필요
- Phase 1의 `markRead` 호출 → Phase 2의 새 메시지 구분선에서 `myLastReadMessageId` 사용
- Phase 3의 예약 카드 렌더링은 Phase 1~2 완료 후 자연스럽게 통합
- Phase 4는 독립적이나, Phase 1의 optimistic update와 함께 적용 시 최적
