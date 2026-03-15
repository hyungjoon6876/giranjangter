# 린클 도메인 이벤트 카탈로그

> 시스템에서 발행되는 모든 도메인 이벤트 정의
> 작성일: 2026-03-14
> 기반: PRD v2.12

---

## 이벤트 공통 구조

```json
{
  "eventId": "evt_01HXYZ...",
  "eventType": "listing.created",
  "aggregateType": "listing",
  "aggregateId": "01LISTING...",
  "actorId": "01USER...",
  "actorRole": "user",
  "payload": { ... },
  "occurredAt": "2026-03-14T05:30:00Z",
  "version": 1
}
```

| 필드 | 설명 |
|------|------|
| eventId | 이벤트 고유 ID (ULID) |
| eventType | `{domain}.{action}` 형식 |
| aggregateType | 이벤트가 속한 집계 루트 |
| aggregateId | 대상 엔티티 ID |
| actorId | 이벤트를 발생시킨 사용자 (system이면 null) |
| actorRole | user / moderator / admin / system |
| payload | 이벤트별 상세 데이터 |
| occurredAt | 발생 시각 (UTC) |
| version | 이벤트 스키마 버전 |

---

## 1. Listing 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `listing.created` | 매물 등록 완료 | listingId, listingType, serverId, itemName, priceAmount | 알림(저장검색), 분석 |
| `listing.updated` | 매물 수정 | listingId, changedFields[] | 알림(찜 가격변경), 분석 |
| `listing.status_changed` | 상태 전이 | listingId, fromStatus, toStatus, reasonCode | 알림, 채팅 시스템메시지, 분석 |
| `listing.deleted` | 매물 소프트 삭제 | listingId, reasonCode | 검색 인덱스 제거 |
| `listing.hidden_by_admin` | 관리자 강제 숨김 | listingId, adminUserId, reasonCode | 알림(작성자), 감사로그 |
| `listing.favorited` | 찜 추가 | listingId, userId | 분석 |
| `listing.unfavorited` | 찜 해제 | listingId, userId | 분석 |
| `listing.viewed` | 상세 조회 | listingId, viewerUserId | 분석, viewCount 갱신 |

---

## 2. Chat 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `chat.room_created` | 채팅방 생성 | chatRoomId, listingId, sellerUserId, buyerUserId | 알림(판매자), 분석 |
| `chat.message_sent` | 메시지 전송 | chatRoomId, messageId, senderUserId, messageType | SSE push, 알림(오프라인), 분석 |
| `chat.message_read` | 읽음 처리 | chatRoomId, userId, lastReadMessageId | SSE push(읽음 표시), 분석 |
| `chat.status_changed` | 채팅방 상태 변경 | chatRoomId, fromStatus, toStatus | 분석 |
| `chat.report_locked` | 신고로 잠금 | chatRoomId, reportId | 알림(양측), 분석 |

---

## 3. Reservation 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `reservation.proposed` | 예약 제안 | reservationId, chatRoomId, listingId, scheduledAt | 알림(상대방), 채팅 시스템메시지, 분석 |
| `reservation.confirmed` | 예약 확정 | reservationId, confirmedByUserId | 알림(양측), listing 상태 변경, 채팅 시스템메시지, 분석 |
| `reservation.cancelled` | 예약 취소 | reservationId, cancelledByUserId, reasonCode | 알림(상대방), listing 상태 복귀, 채팅 시스템메시지, 분석 |
| `reservation.expired` | 예약 만료 (시스템) | reservationId | 알림(양측), listing 상태 복귀, 분석 |
| `reservation.time_approaching` | 예약 시간 임박 (1시간 전) | reservationId, scheduledAt | 알림(양측) |

---

## 4. Trade Completion 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `trade.completion_requested` | 완료 요청 | completionId, listingId, requestedByUserId | 알림(상대방), 분석 |
| `trade.completion_confirmed` | 상대방 확인 | completionId, confirmedByUserId | 알림(양측), listing→completed, 후기 가능, 분석 |
| `trade.auto_confirmed` | 48시간 자동 확정 | completionId | 알림(양측), listing→completed, 후기 가능, 분석 |
| `trade.completion_disputed` | 이의 제기 | completionId, disputedByUserId, reason | 알림(양측+운영), 운영 큐 적재, 분석 |

---

## 5. Review 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `review.created` | 후기 작성 | reviewId, completionId, rating, reviewerUserId, targetUserId | 알림(대상), 신뢰 점수 갱신, 분석 |
| `review.hidden_by_admin` | 관리자 숨김 | reviewId, adminUserId, reasonCode | 알림(작성자), 신뢰 점수 재계산, 감사로그 |

---

## 6. Report 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `report.submitted` | 신고 접수 | reportId, reporterUserId, targetType, targetId, reportType | 운영 큐, 분석 |
| `report.assigned` | 담당자 배정 | reportId, assigneeUserId | 운영 대시보드 |
| `report.action_taken` | 조치 실행 | reportId, actionCode, targetUserId, restrictionScope | 알림(피신고자), 감사로그, 분석 |
| `report.resolved` | 신고 종결 | reportId, resolution | 알림(신고자), 분석 |

---

## 7. User/Account 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `user.registered` | 회원가입 | userId, provider | 분석, 온보딩 |
| `user.profile_updated` | 프로필 수정 | userId, changedFields[] | 분석 |
| `user.restricted` | 계정 제한 | userId, restrictionScope, durationDays, reasonCode | 알림(본인), 기능 차단, 감사로그 |
| `user.restriction_lifted` | 제한 해제 | userId, restrictionId | 알림(본인), 기능 복구 |
| `user.blocked` | 사용자 차단 | blockerUserId, blockedUserId | 채팅 차단 |
| `user.unblocked` | 차단 해제 | blockerUserId, blockedUserId | |

---

## 8. Notification 이벤트

| eventType | 트리거 | payload 핵심 필드 | 소비자 |
|-----------|--------|------------------|--------|
| `notification.created` | 알림 생성 | notificationId, userId, type, referenceType, referenceId | 푸시 발송, 인박스 |
| `notification.push_sent` | 푸시 발송 완료 | notificationId, platform, success | 분석 |
| `notification.read` | 알림 읽음 | notificationId, userId | 분석 |

---

## 이벤트 → 알림 매핑 (발송 매트릭스)

| 도메인 이벤트 | 알림 수신자 | 채널 | 우선순위 |
|-------------|-----------|------|---------|
| `chat.message_sent` | 상대방 | 푸시 + 인앱 | P2 |
| `chat.room_created` | 매물 작성자 | 푸시 + 인앱 | P2 |
| `reservation.proposed` | 상대방 | 푸시 + 인앱 | P1 |
| `reservation.confirmed` | 양측 | 푸시 + 인앱 | P1 |
| `reservation.cancelled` | 상대방 | 푸시 + 인앱 | P1 |
| `reservation.time_approaching` | 양측 | 푸시 + 인앱 | P1 |
| `trade.completion_requested` | 상대방 | 푸시 + 인앱 | P1 |
| `trade.completion_confirmed` | 양측 | 인앱 | P2 |
| `review.created` | 대상 사용자 | 인앱 | P3 |
| `report.action_taken` | 피신고자 | 인앱 | P4 |
| `report.resolved` | 신고자 | 인앱 | P4 |
| `listing.status_changed` | 찜한 사용자들 | 인앱 | P3 |
| `listing.updated` (가격변경) | 찜한 사용자들 | 푸시 + 인앱 | P3 |

### 알림 억제 규칙
- 현재 해당 채팅방을 보고 있는 경우 → 푸시 억제
- 심야 (23:00~08:00) → P3/P4 푸시 억제, P1/P2만 발송
- 동일 이벤트 5분 이내 중복 → dedup
- 차단한 사용자 → 알림 미발송

---

## MVP 이벤트 발행 방식

### Phase 1: 동기 이벤트 (MVP)
```go
// 서비스 레이어에서 직접 발행
func (s *ListingService) CreateListing(ctx context.Context, req CreateListingRequest) (*Listing, error) {
    listing, err := s.repo.Create(ctx, req)
    if err != nil {
        return nil, err
    }

    // 동기 이벤트 발행
    s.eventBus.Publish(ctx, Event{
        Type:          "listing.created",
        AggregateType: "listing",
        AggregateId:   listing.ID,
        Payload:       listing,
    })

    return listing, nil
}
```

### Phase 2+: Outbox 패턴 (확장)
```
트랜잭션 내:
  1. 비즈니스 데이터 저장
  2. outbox 테이블에 이벤트 저장 (같은 트랜잭션)

별도 워커:
  3. outbox 폴링 → 이벤트 발행 → 발행 완료 마킹
```

---

## 문서 개선 로그
| 날짜 | 변경 |
|------|------|
| 2026-03-14 | 초안: 8개 도메인 40+ 이벤트, 발송 매트릭스, 억제 규칙, 발행 방식 |
