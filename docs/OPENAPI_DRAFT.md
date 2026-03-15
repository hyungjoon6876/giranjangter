# 린클 OpenAPI 초안

> MVP API 엔드포인트별 상세 스펙
> 작성일: 2026-03-14
> 기반: PRD v2.12, TECH_STACK_ARCHITECTURE.md (Go + PostgreSQL)

---

## 공통 규약

### Base URL
```
Production: https://lincle.yourdomain.com/api/v1
Development: http://localhost:8080/api/v1
```

### 인증
- `Authorization: Bearer {accessToken}` 헤더
- access token: 15분 TTL
- refresh token: 30일 TTL, httpOnly cookie 또는 body

### 공통 에러 응답
```json
{
  "error": {
    "code": "LISTING_NOT_FOUND",
    "message": "매물을 찾을 수 없습니다.",
    "details": {}
  }
}
```

### 공통 에러 코드
| HTTP | code | 설명 |
|------|------|------|
| 400 | VALIDATION_ERROR | 입력값 검증 실패 |
| 401 | UNAUTHORIZED | 인증 필요 |
| 403 | FORBIDDEN | 권한 없음 |
| 404 | NOT_FOUND | 리소스 없음 |
| 409 | CONFLICT | 상태 충돌 (이미 예약됨 등) |
| 422 | INVALID_TRANSITION | 금지된 상태 전이 |
| 429 | RATE_LIMITED | 요청 제한 초과 |
| 500 | INTERNAL_ERROR | 서버 오류 |

### 페이지네이션
커서 기반 페이지네이션 (offset 사용 안 함):
```json
{
  "data": [...],
  "cursor": {
    "next": "eyJpZCI6MTAwfQ==",
    "hasMore": true
  }
}
```
요청 파라미터: `?cursor={next}&limit=20` (limit 기본 20, 최대 100)

### 날짜/시간
- 모든 timestamp는 ISO 8601 UTC: `2026-03-14T05:30:00Z`
- 클라이언트가 로컬 시간으로 변환

---

## 1. 인증/계정 API

### POST /auth/login
소셜 로그인 (OAuth2 콜백 처리)

**Request:**
```json
{
  "provider": "kakao",
  "providerToken": "oauth-token-from-provider"
}
```

**Response 200:**
```json
{
  "accessToken": "eyJ...",
  "refreshToken": "eyJ...",
  "expiresIn": 900,
  "user": {
    "userId": "01HXYZ...",
    "nickname": "검은기사",
    "role": "user",
    "accountStatus": "active",
    "isNewUser": false
  }
}
```

**Response 201** (신규 가입):
- `isNewUser: true`, 프로필 설정으로 유도

### POST /auth/refresh
**Request:**
```json
{
  "refreshToken": "eyJ..."
}
```
**Response 200:** 새 accessToken + refreshToken

### POST /auth/logout
**Headers:** Authorization 필요
**Response:** 204 No Content

### GET /me
현재 로그인 사용자 정보

**Response 200:**
```json
{
  "userId": "01HXYZ...",
  "nickname": "검은기사",
  "avatarUrl": "/uploads/avatars/01HXYZ.jpg",
  "introduction": "바츠서버 8년차",
  "primaryServerId": "bartz",
  "role": "user",
  "accountStatus": "active",
  "completedTradeCount": 42,
  "positiveReviewCount": 38,
  "responseBadge": "fast",
  "trustBadge": "trusted",
  "lastActiveAt": "2026-03-14T05:30:00Z",
  "createdAt": "2025-01-15T00:00:00Z"
}
```

### PATCH /me/profile
**Request:**
```json
{
  "nickname": "새닉네임",
  "introduction": "자기소개 수정",
  "primaryServerId": "ken_rauhel",
  "avatarUrl": "/uploads/avatars/new.jpg"
}
```
**Response 200:** 수정된 프로필 전체

---

## 2. 매물 API

### POST /listings
매물 등록

**Request:**
```json
{
  "listingType": "sell",
  "serverId": "bartz",
  "categoryId": "weapon",
  "itemName": "집행검",
  "title": "집행검 +9 급처합니다",
  "description": "옵션 힘+3 덱+2, 직접 강화한 것입니다",
  "priceType": "negotiable",
  "priceAmount": 8000000,
  "quantity": 1,
  "enhancementLevel": 9,
  "optionsText": "힘+3 덱+2",
  "tradeMethod": "in_game",
  "preferredMeetingAreaText": "기란마을 분수대",
  "availableTimeText": "평일 저녁 8시 이후, 주말 종일",
  "imageIds": ["img_01", "img_02"]
}
```

**Response 201:**
```json
{
  "listingId": "01LISTING...",
  "status": "available",
  "visibility": "public",
  "createdAt": "2026-03-14T05:30:00Z",
  ...전체 필드
}
```

**Validation:**
- title: 2~100자
- description: 10~2000자
- priceAmount: priceType이 fixed/negotiable이면 필수, > 0
- serverId: 유효한 서버 ID
- categoryId: 유효한 카테고리 ID
- imageIds: 최대 5개

### GET /listings
매물 목록 (검색/필터)

**Query Parameters:**
| 파라미터 | 타입 | 설명 | 기본값 |
|----------|------|------|--------|
| q | string | 검색어 (제목+아이템명) | - |
| serverId | string | 서버 필터 | - |
| categoryId | string | 카테고리 필터 | - |
| listingType | enum | sell / buy | 전체 |
| priceMin | int | 최소 가격 | - |
| priceMax | int | 최대 가격 | - |
| status | enum | available / reserved 등 | available |
| tradeMethod | enum | in_game / offline / either | 전체 |
| sort | enum | recent / price_asc / price_desc / popular | recent |
| cursor | string | 페이지네이션 커서 | - |
| limit | int | 페이지 크기 | 20 |

**Response 200:**
```json
{
  "data": [
    {
      "listingId": "01LISTING...",
      "listingType": "sell",
      "title": "집행검 +9 급처합니다",
      "itemName": "집행검",
      "priceType": "negotiable",
      "priceAmount": 8000000,
      "enhancementLevel": 9,
      "serverId": "bartz",
      "serverName": "바츠",
      "status": "available",
      "tradeMethod": "in_game",
      "thumbnailUrl": "/uploads/listings/01.thumb.jpg",
      "author": {
        "userId": "01USER...",
        "nickname": "검은기사",
        "trustBadge": "trusted",
        "responseBadge": "fast"
      },
      "viewCount": 120,
      "favoriteCount": 5,
      "chatCount": 3,
      "lastActivityAt": "2026-03-14T05:00:00Z",
      "createdAt": "2026-03-13T10:00:00Z"
    }
  ],
  "cursor": {
    "next": "eyJ...",
    "hasMore": true
  }
}
```

### GET /listings/{listingId}
매물 상세

**Response 200:**
```json
{
  "listingId": "01LISTING...",
  "listingType": "sell",
  "title": "집행검 +9 급처합니다",
  "itemName": "집행검",
  "description": "옵션 힘+3 덱+2, 직접 강화한 것입니다",
  "priceType": "negotiable",
  "priceAmount": 8000000,
  "quantity": 1,
  "enhancementLevel": 9,
  "optionsText": "힘+3 덱+2",
  "serverId": "bartz",
  "serverName": "바츠",
  "categoryId": "weapon",
  "categoryName": "무기",
  "status": "available",
  "visibility": "public",
  "tradeMethod": "in_game",
  "preferredMeetingAreaText": "기란마을 분수대",
  "availableTimeText": "평일 저녁 8시 이후, 주말 종일",
  "images": [
    {"imageId": "img_01", "url": "/uploads/listings/01_1.jpg", "order": 1},
    {"imageId": "img_02", "url": "/uploads/listings/01_2.jpg", "order": 2}
  ],
  "author": {
    "userId": "01USER...",
    "nickname": "검은기사",
    "avatarUrl": "/uploads/avatars/01.jpg",
    "trustBadge": "trusted",
    "responseBadge": "fast",
    "completedTradeCount": 42,
    "lastActiveAt": "2026-03-14T05:00:00Z"
  },
  "viewCount": 120,
  "favoriteCount": 5,
  "chatCount": 3,
  "isFavorited": false,
  "isOwner": false,
  "availableActions": ["start_chat", "favorite", "report"],
  "reservedChatRoomId": null,
  "lastActivityAt": "2026-03-14T05:00:00Z",
  "createdAt": "2026-03-13T10:00:00Z",
  "updatedAt": "2026-03-14T03:00:00Z"
}
```

**availableActions 규칙:**
- 비로그인: `[]`
- 본인 매물: `["edit", "change_status", "delete"]`
- 타인 + available: `["start_chat", "favorite", "report"]`
- 타인 + reserved: `["start_chat", "favorite", "report"]` (예약중 배지 표시)
- 타인 + completed/cancelled: `["report"]`

### PATCH /listings/{listingId}
매물 수정 (본인만)

**Request:** POST /listings와 동일 필드 (부분 수정 가능)

**제한:**
- `pending_trade` 상태: 가격/수량 수정 불가
- `completed`/`cancelled`: 수정 불가 (403)

### POST /listings/{listingId}/status
매물 상태 변경

**Request:**
```json
{
  "action": "cancel",
  "reasonCode": "changed_mind"
}
```

**action 종류:**
| action | 허용 from 상태 | 결과 상태 |
|--------|--------------|----------|
| reserve | available | reserved |
| unreserve | reserved | available |
| start_trade | reserved | pending_trade |
| complete | pending_trade | completed |
| cancel | available, reserved, pending_trade | cancelled |
| reopen | cancelled | available |

### POST /listings/{listingId}/favorite
### DELETE /listings/{listingId}/favorite
찜 토글. Response 204.

### GET /me/listings
내 매물 목록

**Query:** `?status=available,reserved&cursor=...&limit=20`

---

## 3. 채팅 API

### POST /listings/{listingId}/chats
채팅방 생성 (매물 상세에서 "채팅 시작")

**Response 201:**
```json
{
  "chatRoomId": "01CHAT...",
  "listingId": "01LISTING...",
  "sellerUserId": "01USER_A...",
  "buyerUserId": "01USER_B...",
  "chatStatus": "open",
  "createdAt": "2026-03-14T05:30:00Z"
}
```
**409 Conflict:** 이미 동일 상대와 채팅방이 존재하는 경우 → 기존 chatRoomId 반환

### GET /chats
내 채팅방 목록

**Response 200:**
```json
{
  "data": [
    {
      "chatRoomId": "01CHAT...",
      "listingId": "01LISTING...",
      "listingTitle": "집행검 +9 급처",
      "listingThumbnail": "/uploads/...",
      "listingStatus": "available",
      "counterparty": {
        "userId": "01USER...",
        "nickname": "검은기사",
        "avatarUrl": "...",
        "trustBadge": "trusted"
      },
      "chatStatus": "open",
      "lastMessage": {
        "messageId": "01MSG...",
        "bodyText": "안녕하세요, 가격 협의 가능한가요?",
        "messageType": "text",
        "sentAt": "2026-03-14T05:30:00Z"
      },
      "unreadCount": 2,
      "updatedAt": "2026-03-14T05:30:00Z"
    }
  ],
  "cursor": { "next": "...", "hasMore": true }
}
```

### GET /chats/{chatRoomId}/messages
메시지 목록 (커서 기반, 최신순 역순)

**Query:** `?cursor=...&limit=50&direction=backward`

**Response 200:**
```json
{
  "data": [
    {
      "messageId": "01MSG...",
      "chatRoomId": "01CHAT...",
      "senderUserId": "01USER_B...",
      "messageType": "text",
      "bodyText": "750만원에 가능하신가요?",
      "metadataJson": null,
      "sentAt": "2026-03-14T05:31:00Z"
    },
    {
      "messageId": "01MSG_SYS...",
      "chatRoomId": "01CHAT...",
      "senderUserId": null,
      "messageType": "system",
      "bodyText": "예약이 제안되었습니다.",
      "metadataJson": {
        "eventType": "reservation_proposed",
        "reservationId": "01RES..."
      },
      "sentAt": "2026-03-14T05:35:00Z"
    },
    {
      "messageId": "01MSG_RES...",
      "chatRoomId": "01CHAT...",
      "senderUserId": "01USER_A...",
      "messageType": "reservation_card",
      "bodyText": null,
      "metadataJson": {
        "reservationId": "01RES...",
        "scheduledAt": "2026-03-15T14:00:00Z",
        "meetingType": "in_game",
        "serverText": "바츠",
        "meetingPointText": "기란마을 분수대",
        "status": "proposed"
      },
      "sentAt": "2026-03-14T05:35:00Z"
    }
  ],
  "cursor": { "next": "...", "hasMore": true }
}
```

### POST /chats/{chatRoomId}/messages
메시지 전송

**Request:**
```json
{
  "messageType": "text",
  "bodyText": "네, 750만원 가능합니다!",
  "clientMessageId": "client-uuid-for-dedup"
}
```
**Response 201:** 생성된 메시지 객체

### GET /sse/connect
SSE 실시간 스트림 연결

**Headers:** `Authorization: Bearer {token}`
**Query:** `?lastEventId={messageId}` (재연결 시 누락분 backfill)

**Event format:**
```
event: new_message
data: {"chatRoomId":"01CHAT...","message":{...}}

event: typing
data: {"chatRoomId":"01CHAT...","userId":"01USER..."}

event: status_change
data: {"chatRoomId":"01CHAT...","listingId":"01LIST...","newStatus":"reserved"}
```

### POST /chats/{chatRoomId}/read
읽음 처리

**Request:**
```json
{
  "lastReadMessageId": "01MSG..."
}
```
**Response:** 204

### POST /users/{userId}/block
### DELETE /users/{userId}/block
차단/해제. Response 204.

---

## 4. 예약 API

### POST /chats/{chatRoomId}/reservations
예약 제안

**Request:**
```json
{
  "scheduledAt": "2026-03-15T14:00:00Z",
  "meetingType": "in_game",
  "serverId": "bartz",
  "meetingPointText": "기란마을 분수대",
  "characterNameSeller": "검은기사",
  "characterNameBuyer": "흰마법사",
  "noteToCounterparty": "분수대 왼쪽에서 만나요",
  "expiresAt": "2026-03-14T23:59:00Z"
}
```

**Response 201:**
```json
{
  "reservationId": "01RES...",
  "listingId": "01LISTING...",
  "chatRoomId": "01CHAT...",
  "proposerUserId": "01USER_A...",
  "counterpartUserId": "01USER_B...",
  "reservationStatus": "proposed",
  "scheduledAt": "2026-03-15T14:00:00Z",
  "meetingType": "in_game",
  "serverId": "bartz",
  "meetingPointText": "기란마을 분수대",
  "expiresAt": "2026-03-14T23:59:00Z",
  "createdAt": "2026-03-14T05:35:00Z"
}
```

**409:** 이 매물에 이미 활성 예약이 존재

### POST /reservations/{reservationId}/confirm
예약 확정 (상대방)

**Response 200:** 확정된 예약 + 매물 상태 reserved/pending_trade로 자동 전환

### POST /reservations/{reservationId}/cancel
예약 취소

**Request:**
```json
{
  "reasonCode": "schedule_conflict"
}
```
**Response 200:** 취소된 예약 + 매물 상태 available로 자동 복귀 (조건부)

### GET /me/trades
내 거래 목록 (진행 중인 거래 워크스페이스)

**Query:** `?tab=active,completed&cursor=...`

**Response 200:**
```json
{
  "data": [
    {
      "chatRoomId": "01CHAT...",
      "listingId": "01LISTING...",
      "listingTitle": "집행검 +9",
      "counterparty": { "userId": "...", "nickname": "..." },
      "tradeStatus": "reserved",
      "reservation": {
        "reservationId": "01RES...",
        "scheduledAt": "2026-03-15T14:00:00Z",
        "reservationStatus": "confirmed"
      },
      "unreadCount": 0,
      "nextAction": "wait_for_trade_time",
      "updatedAt": "2026-03-14T06:00:00Z"
    }
  ]
}
```

---

## 5. 거래 완료 API

### POST /listings/{listingId}/complete
거래 완료 요청

**Request:**
```json
{
  "reservationId": "01RES...",
  "completionNote": "정상 거래 완료"
}
```

**Response 201:**
```json
{
  "completionId": "01COMP...",
  "listingId": "01LISTING...",
  "reservationId": "01RES...",
  "requestedByUserId": "01USER_A...",
  "counterpartUserId": "01USER_B...",
  "completionStatus": "pending_confirm",
  "autoConfirmAt": "2026-03-16T06:00:00Z",
  "createdAt": "2026-03-14T14:30:00Z"
}
```

### POST /trade-completions/{completionId}/confirm
상대방 완료 확인

**Response 200:** completionStatus → "confirmed", listing → "completed"

### POST /trade-completions/{completionId}/dispute
이의 제기

**Request:**
```json
{
  "reason": "아이템을 받지 못했습니다",
  "evidenceText": "상세 설명..."
}
```

---

## 6. 후기 API

### POST /trade-completions/{completionId}/reviews
후기 작성

**Request:**
```json
{
  "rating": "positive",
  "comment": "빠른 거래, 친절하심. 추천합니다!"
}
```
- rating: `positive` | `negative`
- comment: 선택, 10~500자

**Validation:**
- 거래 완료(confirmed) 상태에서만 가능
- 완료 후 7일 이내
- 동일 completionId당 1인 1회

### GET /users/{userId}/reviews
사용자 받은 후기 목록

---

## 7. 신고 API

### POST /reports
신고 접수

**Request:**
```json
{
  "targetType": "listing",
  "targetId": "01LISTING...",
  "reportType": "fake_listing",
  "description": "실제로 없는 아이템을 올린 것 같습니다",
  "evidenceImageIds": ["img_ev_01"]
}
```

- targetType: `listing` | `user` | `chat_room` | `message` | `review`
- reportType: `fake_listing` | `scam_suspicion` | `no_show` | `harassment` | `spam` | `prohibited_item` | `privacy_exposure` | `other`

**Response 201:** 신고 접수 확인 + reportId

### GET /me/reports
내 신고 이력

---

## 8. 알림 API

### GET /notifications
알림 목록

**Response 200:**
```json
{
  "data": [
    {
      "notificationId": "01NOTIF...",
      "type": "reservation_confirmed",
      "title": "예약이 확정되었습니다",
      "body": "검은기사님과의 예약이 확정되었습니다.",
      "referenceType": "reservation",
      "referenceId": "01RES...",
      "deepLink": "/trades/01CHAT...",
      "isRead": false,
      "createdAt": "2026-03-14T06:00:00Z"
    }
  ]
}
```

### POST /notifications/read
읽음 처리

**Request:**
```json
{
  "notificationIds": ["01NOTIF_A...", "01NOTIF_B..."]
}
```

### POST /push-tokens
푸시 토큰 등록

**Request:**
```json
{
  "token": "fcm-token...",
  "platform": "ios"
}
```

---

## 9. 관리자 API

모든 관리자 API는 `role: moderator | admin` 인증 필요.

### GET /admin/reports
신고 큐 목록

**Query:** `?status=submitted,assigned&sort=priority&cursor=...`

### GET /admin/reports/{reportId}
신고 상세 (증빙, 대상 이력 포함)

### POST /admin/reports/{reportId}/actions
조치 실행

**Request:**
```json
{
  "actionCode": "report.case.warn",
  "targetUserId": "01USER...",
  "memo": "허위매물 경고 1회",
  "restrictionScope": null
}
```

### POST /admin/listings/{listingId}/hide
매물 강제 숨김

### POST /admin/users/{userId}/restrict
사용자 제한

**Request:**
```json
{
  "restrictionScope": "listing_only",
  "durationDays": 7,
  "reasonCode": "fake_listing_repeat",
  "memo": "허위매물 반복 등록"
}
```

---

## 10. 이미지 업로드 API

### POST /uploads/images
이미지 업로드 (multipart/form-data)

**Request:** `Content-Type: multipart/form-data`, field: `file`

**Response 201:**
```json
{
  "imageId": "img_01HXYZ...",
  "url": "/uploads/images/01HXYZ.jpg",
  "thumbnailUrl": "/uploads/images/01HXYZ.thumb.jpg",
  "width": 1200,
  "height": 800,
  "sizeBytes": 245000
}
```

**제한:**
- 최대 10MB
- 지원 형식: jpg, png, webp
- 업로드 후 매물/프로필에 연결하지 않으면 24시간 후 정리

---

## 문서 개선 로그
| 날짜 | 변경 |
|------|------|
| 2026-03-14 | 초안: MVP 전체 API 39개 엔드포인트, request/response 상세 |
