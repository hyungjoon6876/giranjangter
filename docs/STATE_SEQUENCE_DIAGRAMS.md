# 상태 전이 다이어그램

PRD 기반 도메인 객체별 상태 머신 정의.

---

## 1. 매물(Listing) 상태 머신

### 1.1 상태 정의

| 상태 | 설명 |
|---|---|
| `available` | 거래 가능. 새로운 채팅과 제안을 받을 수 있음 |
| `reserved` | 특정 상대와 우선 거래 협의가 확정된 상태 |
| `pending_trade` | 거래 약속이 잡혔고 실제 실행 대기 상태 |
| `completed` | 거래 종결. 신규 문의 불가 |
| `cancelled` | 거래 중단 또는 무산 종결 |

### 1.2 상태 전이 다이어그램

```
                    ┌──────────────────────────────────────────┐
                    │              cancelled                    │
                    └──────────────────────────────────────────┘
                      ^              ^              ^
                      │              │              │
                      │(취소)        │(취소)        │(취소)
                      │              │              │
┌───────────┐  예약확정  ┌───────────┐  거래임박  ┌───────────────┐  완료요청  ┌───────────┐
│ available  │────────>│ reserved  │────────>│ pending_trade │────────>│ completed │
└───────────┘         └───────────┘         └───────────────┘         └───────────┘
      ^                    │                       │
      │                    │                       │
      │   예약해제          │   거래불발 후 재오픈    │
      │<───────────────────┘                       │
      │<───────────────────────────────────────────┘
```

### 1.3 허용 전이 목록

| From | To | 트리거 조건 |
|---|---|---|
| `available` | `reserved` | 채팅방 예약 확정 (예약 시각까지 2시간 초과) |
| `reserved` | `pending_trade` | 예약 시각 임박 (2시간 이내 진입 가정) |
| `pending_trade` | `completed` | 거래완료 요청 + 상대방 확인 또는 자동확정 |
| `reserved` | `available` | 예약 해제/취소/만료 (작성자 종결 처리 안 한 경우) |
| `pending_trade` | `available` | 거래 불발 후 재오픈 |
| `available` | `cancelled` | 작성자 직접 취소 또는 운영 강제 종료 |
| `reserved` | `cancelled` | 작성자 직접 취소 또는 운영 강제 종료 |
| `pending_trade` | `cancelled` | 작성자 직접 취소 또는 운영 강제 종료 |

### 1.4 금지 전이

| From | To | 사유 |
|---|---|---|
| `completed` | `available` | 종결 상태 재오픈 금지 (새 매물 생성 유도) |
| `completed` | `reserved` | 종결 상태에서 예약 진입 불가 |
| `cancelled` | `completed` | 취소된 매물의 완료 처리 금지 |
| `cancelled` | `available` | 직접 복구 불가 (복제 후 재등록 또는 운영 정책 기반 복구) |
| `cancelled` | `reserved` | 직접 복구 불가 |
| `completed` | `pending_trade` | 역방향 전이 불가 |
| `completed` | `cancelled` | 완료 후 취소 불가 (분쟁은 별도 Dispute 객체) |

### 1.5 전이 가드 테이블

| From | To | 가드 조건 | 허용 역할 |
|---|---|---|---|
| `available` | `reserved` | 활성 채팅방/예약 연결값 존재 필수. 확정 예약 객체(`confirmed`) 존재 | 작성자(owner) |
| `reserved` | `pending_trade` | 확정 예약 존재 필수. 예약 시각 2시간 이내 | 시스템(system) |
| `pending_trade` | `completed` | completionStage가 `requested` 이상. 상대방 확인 또는 자동확정(24h) 완료 | 작성자(owner), 시스템(system) |
| `reserved` | `available` | 예약 취소/만료 발생. 작성자가 종결 처리하지 않은 상태 | 작성자(owner), 상대방(counterparty), 시스템(system) |
| `pending_trade` | `available` | 거래 불발 확인. 노쇼/분쟁 미진행 상태 | 작성자(owner) |
| `available` | `cancelled` | 없음 (작성자 자유) | 작성자(owner), 운영자(admin) |
| `reserved` | `cancelled` | 사유 또는 관련 채팅방 지정 필요 | 작성자(owner), 운영자(admin) |
| `pending_trade` | `cancelled` | 사유 또는 관련 채팅방 지정 필요 | 작성자(owner), 운영자(admin) |

### 1.6 운영 예외 전이

| From | To | 조건 | 허용 역할 |
|---|---|---|---|
| `completed` (+ `resolved_not_completed`) | `available` | 운영 판단으로 거래 미완료 판정 후 재오픈 | 운영자(admin) |
| `completed` (+ `resolved_not_completed`) | `cancelled` | 운영 판단 불발/취소 확정 | 운영자(admin) |
| 모든 상태 | 강제 숨김/강제 종료 | 정책 위반 시 | 운영자(admin) |

---

## 2. 예약(Reservation) 상태 머신

### 2.1 상태 정의

| 상태 | 설명 |
|---|---|
| `proposed` | 한쪽이 예약안을 제안한 상태 |
| `confirmed` | 양측 또는 작성자 정책 기준으로 예약 확정 |
| `expired` | 응답 기한 초과로 자동 만료 |
| `cancelled` | 어느 한쪽이 예약을 취소 |
| `fulfilled` | 거래 완료로 예약 이행 완료 |
| `no_show_reported` | 노쇼 신고 접수 |

### 2.2 상태 전이 다이어그램

```
                                     ┌─────────────┐
                                     │  expired    │
                                     └─────────────┘
                                       ^
                                       │(응답기한 초과)
                                       │
┌───────────┐  상대방 수락  ┌───────────┐  거래완료   ┌───────────┐
│ proposed  │────────────>│ confirmed │──────────>│ fulfilled │
└───────────┘             └───────────┘           └───────────┘
      │                        │
      │(제안자/상대방 취소)      │(어느 한쪽 취소)
      v                        v
┌───────────┐             ┌───────────┐
│ cancelled │             │ cancelled │
└───────────┘             └───────────┘
                               │
                               │(예약시각 경과 후 노쇼)
                               v
                          ┌──────────────────┐
                          │ no_show_reported  │
                          └──────────────────┘
```

### 2.3 전이 가드 테이블

| From | To | 가드 조건 | 허용 역할 |
|---|---|---|---|
| `proposed` | `confirmed` | 상대방(counterparty)이 수락. 동일 매물에 다른 활성 예약 없음 | 상대방(counterparty) |
| `proposed` | `expired` | expiresAt 시각 경과. 상대방 무응답 | 시스템(system) |
| `proposed` | `cancelled` | 제안자 또는 상대방이 명시 취소 | 제안자(proposer), 상대방(counterparty) |
| `confirmed` | `fulfilled` | 거래완료 확정 (completionStage가 confirmed/auto_confirmed 이상) | 시스템(system) |
| `confirmed` | `cancelled` | 어느 한쪽 취소 요청. cancellationReasonCode 필수 | 작성자(owner), 상대방(counterparty), 운영자(admin) |
| `confirmed` | `no_show_reported` | 예약 시각 경과 후 어느 한쪽이 노쇼 신고 제출 | 작성자(owner), 상대방(counterparty) |

### 2.4 예약-매물 상태 연동 규칙

| 예약 이벤트 | 매물 상태 변경 |
|---|---|
| `proposed` 생성 | 매물 상태 변경 없음 (즉시 `reserved`로 바꾸지 않음) |
| `proposed` -> `confirmed` (예약시각 2시간 초과) | 매물: `available` -> `reserved` |
| `proposed` -> `confirmed` (예약시각 2시간 이내) | 매물: `available` -> `pending_trade` |
| `confirmed`, 예약시각 2시간 이내 도달 | 매물: `reserved` -> `pending_trade` |
| `confirmed` -> `cancelled` 또는 `expired` | 매물: `available` 복귀 후보 (작성자 종결 처리 안 한 경우) |
| `confirmed` -> `fulfilled` | 매물: `pending_trade` -> `completed` |
| `confirmed` -> `no_show_reported` | 노쇼 사건 큐 진입. 매물 상태는 즉시 변경하지 않음 |

### 2.5 금지 전이

| From | To | 사유 |
|---|---|---|
| `fulfilled` | 모든 상태 | 이행 완료는 종결 상태 |
| `expired` | `confirmed` | 만료 후 직접 확정 불가 (새 예약 생성 필요) |
| `cancelled` | `confirmed` | 취소 후 직접 확정 불가 (새 예약 생성 필요) |
| `no_show_reported` | `fulfilled` | 노쇼 신고와 이행 완료는 상호 배타 |

---

## 3. 채팅방(ChatRoom) 상태 머신

### 3.1 상태 정의

| 상태 | 설명 |
|---|---|
| `open` | 일반 대화 가능 상태 |
| `reservation_proposed` | 한쪽이 예약안을 제안한 상태 |
| `reservation_confirmed` | 예약 확정된 상태 |
| `trade_due` | 예약 시각이 임박했거나 도래한 상태 |
| `deal_completed` | 해당 채팅 기준 거래 완료 |
| `deal_cancelled` | 해당 채팅 기준 협의 종료 |
| `report_locked` | 신고/분쟁 처리로 운영 잠금 |

### 3.2 상태 전이 다이어그램

```
                                                 ┌─────────────────┐
                                            ┌───>│ report_locked    │
                                            │    └─────────────────┘
                                            │(신고/분쟁)    │
                                            │               │(운영 해제)
                                            │               v
┌──────┐  예약제안  ┌─────────────────────┐  │    ┌─────────────────┐
│ open │────────>│ reservation_proposed │──┤    │ open (복귀)      │
└──────┘         └─────────────────────┘  │    └─────────────────┘
   ^  ^                  │                 │
   │  │                  │(상대방 수락)     │
   │  │                  v                 │
   │  │          ┌─────────────────────┐   │
   │  │          │reservation_confirmed│───┤
   │  │          └─────────────────────┘   │
   │  │                  │                 │
   │  │                  │(시각 임박)       │
   │  │                  v                 │
   │  │          ┌─────────────────────┐   │
   │  │          │     trade_due       │───┤
   │  │          └─────────────────────┘   │
   │  │            │              │        │
   │  │   (완료확정)│     (취소/불발)│        │
   │  │            v              v        │
   │  │  ┌────────────────┐ ┌────────────────┐
   │  │  │ deal_completed │ │ deal_cancelled │
   │  │  └────────────────┘ └────────────────┘
   │  │                           │
   │  │(예약취소/만료)              │
   │  └───────────────────────────┘
   │
   │(예약제안 거절/만료)
   └──────────────────────────────────────────────┘
```

### 3.3 전이 가드 테이블

| From | To | 가드 조건 | 허용 역할 |
|---|---|---|---|
| `open` | `reservation_proposed` | 예약 객체 생성 (`proposed` 상태) | 작성자(owner), 상대방(counterparty) |
| `reservation_proposed` | `reservation_confirmed` | 예약 상태가 `confirmed`로 전이 | 상대방(counterparty) |
| `reservation_proposed` | `open` | 예약 거절 또는 만료 | 상대방(counterparty), 시스템(system) |
| `reservation_confirmed` | `trade_due` | 예약 시각 임박 (2시간 이내 가정) | 시스템(system) |
| `trade_due` | `deal_completed` | 거래완료 확정 (completionStage 종결) | 작성자(owner), 시스템(system) |
| `trade_due` | `deal_cancelled` | 거래 취소/불발 처리 | 작성자(owner), 상대방(counterparty) |
| `reservation_confirmed` | `deal_cancelled` | 예약 취소 | 작성자(owner), 상대방(counterparty) |
| `deal_cancelled` | `open` | 재대화 가능 (매물이 `available`로 복귀한 경우) | 시스템(system) |
| 모든 활성 상태 | `report_locked` | 신고 접수 후 운영 잠금 결정 | 운영자(admin) |
| `report_locked` | `open` 또는 이전 상태 | 운영 해제. 이력형 복구 | 운영자(admin) |

### 3.4 채팅방-예약-매물 연동 규칙

| 예약 이벤트 | 채팅방 상태 변경 | 매물 상태 변경 |
|---|---|---|
| 예약 제안 생성 | `open` -> `reservation_proposed` | 변경 없음 |
| 예약 확정 | `reservation_proposed` -> `reservation_confirmed` | `available` -> `reserved` 또는 `pending_trade` |
| 예약 시각 임박 | `reservation_confirmed` -> `trade_due` | `reserved` -> `pending_trade` |
| 예약 취소/만료 | -> `open` 또는 `deal_cancelled` | -> `available` 복귀 후보 |
| 거래 완료 | `trade_due` -> `deal_completed` | `pending_trade` -> `completed` |
| 신고 접수 | -> `report_locked` | 변경 없음 (별도 운영 판단) |

### 3.5 report_locked 상태 제약

- 읽기 전용 또는 운영 안내 메시지만 허용
- typing indicator 비활성화
- 운영 해제 시 `report_locked` -> `open` 또는 적절한 후속 상태로 이력형 복구

---

## 4. 신고(Report) 상태 머신

### 4.1 상태 정의

| 상태 | 설명 |
|---|---|
| `submitted` | 신고 접수 완료, triage 대기 |
| `triaged` | 우선순위 분류 및 담당자 배정 완료 |
| `investigating` | 운영자가 조사/검토 중 |
| `resolved` | 처리 완료 (조치 적용됨) |
| `rejected` | 무근거/오탐으로 기각 |

### 4.2 상태 전이 다이어그램

```
┌───────────┐  1차분류  ┌──────────┐  조사착수  ┌───────────────┐
│ submitted │────────>│ triaged  │────────>│ investigating │
└───────────┘         └──────────┘         └───────────────┘
                                               │         │
                                    (조치 적용) │         │ (무근거 판정)
                                               v         v
                                         ┌──────────┐ ┌──────────┐
                                         │ resolved │ │ rejected │
                                         └──────────┘ └──────────┘
```

### 4.3 전이 가드 테이블

| From | To | 가드 조건 | 허용 역할 |
|---|---|---|---|
| `submitted` | `triaged` | 우선순위(P1-P4) 배정. 담당자 할당 | CS Operator (P3/P4 1차 분류), Moderator (P1/P2) |
| `triaged` | `investigating` | 담당자가 조사 착수. 대상 요약/내부 로그 확인 시작 | Moderator, Senior Moderator |
| `investigating` | `resolved` | 조치 실행 완료 (경고/숨김/잠금/정지 등). 사유 기록 + 통지 완료 | Moderator (P3/P4), Senior Moderator (P1/P2), Admin |
| `investigating` | `rejected` | 무근거/오탐 판정. 사유 기록 완료 | Moderator, Senior Moderator, Admin |

### 4.4 운영 역할별 권한

| 역할 | 허용 범위 | 제한 |
|---|---|---|
| CS Operator | 신고 열람, 기본 메모 작성, P3/P4 1차 분류 | 숨김/제재 확정 불가 |
| Moderator | 매물 임시 숨김, 채팅 잠금, 경고, P1/P2 1차 조치 | 영구 정지, 민감정보 전체 열람 제한 |
| Senior Moderator | 기간 정지, 이의제기 재검토, 고위험 복구 승인 | 시스템 설정 변경 불가 |
| Admin | 모든 운영 조치, 정책 설정, 감사 로그 열람 | 최소 권한 원칙 적용 |

### 4.5 신고 처리 절차

```
1. 접수 (submitted)
   │
   v
2. 대상 요약 확인 (사용자/매물/채팅/메시지 기본 정보)
   │
   v
3. 내부 로그 확인 (관련 채팅, 예약, 상태 변경, 과거 신고/제재 이력)
   │
   v
4. 우선순위 분류 + 담당자 배정 (triaged)
   │
   v
5. 위험도 판정 (즉시 차단 필요 여부, 임시 숨김 필요 여부)
   │
   v
6. 조사/검토 (investigating)
   │
   v
7. 조치 실행 (경고/임시숨김/채팅잠금/기간정지/영구제한)
   │
   v
8. 사유 기록 (코드 + 자유 메모 + 근거 링크)
   │
   v
9. 통지 (대상자 및 신고자에게 결과/상태 통지)
   │
   v
10. resolved 또는 rejected
    │
    v
11. 사후 추적 (재발 여부, 이의제기 기간, 후속 검토 일자)
```

### 4.6 SLA 참고

- unassigned 상태의 P1 사건은 SLA 위반으로 운영 대시보드에 별도 노출
- P1/P2는 개별 검토 기본, P3/P4 범위에서만 대량 작업 허용

---

## 5. 거래 완료(Trade Completion) 흐름

### 5.1 내부 완료 단계(completionStage)

| 단계 | 의미 | 사용자 노출 문구 | 주요 허용 액션 |
|---|---|---|---|
| `none` | 완료 흐름 미시작 | 노출 없음 | 완료 요청 가능 |
| `requested` | 한쪽이 완료 요청 | 거래완료 확인 대기 | 상대 확인, 이의제기 |
| `confirmed_by_counterparty` | 상대가 명시 확인 | 거래완료 | 후기 작성 |
| `auto_confirmed` | 상대 무응답 자동확정 | 거래완료 | 후기 작성, 신고 |
| `disputed` | 완료 주장에 이의 제기 | 분쟁 진행 중 | 소명 제출 |
| `resolved_completed` | 운영이 완료 판정 | 거래완료 | 후기 작성 |
| `resolved_not_completed` | 운영이 미완료 판정 | 거래불발/취소 | 재오픈 후보 |
| `closed_without_review` | 신고 철회/기각 등 종료 | 처리 완료 | 기록 열람 |

### 5.2 정상 완료 흐름

```
[pending_trade / completionStage=none]
        │
        │  작성자(owner)가 거래완료 요청
        v
[completed / completionStage=requested]
        │
        ├──────────────────────────────────────────┐
        │                                          │
        │ 상대방이 24시간 내 확인                     │ 상대방이 24시간 내 무응답
        v                                          v
[completed / confirmed_by_counterparty]    [completed / auto_confirmed]
        │                                          │
        │                                          │
        v                                          v
    후기 작성 가능                              후기 작성 가능
```

### 5.3 자동확정 타이머 규칙

```
완료 요청 시점 (T=0)
    │
    ├── T+0: 상대방에게 확인 요청 알림 발송
    │
    ├── T+12h(가정): 미응답 시 1차 리마인드 + 이의제기 CTA 제공
    │
    ├── T+24h: 자동확정 실행
    │       │
    │       ├── confirmationMethod = auto
    │       ├── Listing.status = completed 최종 반영
    │       ├── 후기 작성 가능 상태 오픈
    │       └── 감사 로그 + 사용자 알림 생성
    │
    └── 자동확정 후에도 신고는 가능하나, 상태 롤백은 운영 개입 없이 불가
```

### 5.4 분쟁(Dispute) 경로

```
[completed / completionStage=requested]
        │
        │  상대방이 이의 제기
        v
[completed / completionStage=disputed]
        │
        │  Dispute 객체 생성
        v
┌─────────────────────────────────────────────────────┐
│  Dispute 상태 머신                                    │
│                                                      │
│  [open] ──> [waiting_statement] ──> [under_review]   │
│                                         │      │     │
│                              (완료 판정) │      │     │
│                                         v      v     │
│                          [resolved]        [closed]  │
│                                                      │
└─────────────────────────────────────────────────────┘
        │                          │
        │ resolutionType=          │ resolutionType=
        │ completed_upheld         │ not_completed_upheld
        v                          v
[resolved_completed]        [resolved_not_completed]
        │                          │
        v                          v
   후기 작성 가능             매물 재오픈 또는 취소 후보
```

### 5.5 분쟁(Dispute) 전이 가드 테이블

| From | To | 가드 조건 | 허용 역할 |
|---|---|---|---|
| `open` | `waiting_statement` | 양측 소명 요청 발송 | 운영자(admin) |
| `waiting_statement` | `under_review` | 양측 소명 수집 완료 또는 기한 경과 | 운영자(admin), 시스템(system) |
| `under_review` | `resolved` | 운영 판정 완료. resolutionType 및 resolutionReasonCode 기록 | Senior Moderator, Admin |
| `under_review` | `closed` | 신고 철회, 기각, 또는 무결함 판정 | Senior Moderator, Admin |

### 5.6 완료 단계-매물 상태 조합 규칙

| Listing.status | completionStage | 허용 여부 | 설명 |
|---|---|---|---|
| `pending_trade` | `none` | 허용 | 거래 직전/직후 기본 상태 |
| `completed` | `requested` | 허용 | 외부 공개상 종결, 내부 확인 대기 |
| `completed` | `confirmed_by_counterparty` | 허용 | 정상 완료 |
| `completed` | `auto_confirmed` | 허용 | 무응답 자동확정 |
| `completed` | `disputed` | 허용 | 공개 목록 종결 유지, 운영은 분쟁 처리 |
| `available` | `resolved_not_completed` | 허용 | 운영 판단 거래 미완료 후 재오픈 |
| `cancelled` | `resolved_not_completed` | 허용 | 불발/취소 확정 |
| `completed` | `none` | **금지** | 완료 상태에는 completion 객체 필수 |

### 5.7 전체 통합 시퀀스

```
[매물등록]
    │
    v
[available] ─────── 채팅방 생성 (open) ─────── 예약 제안 (proposed)
    │                                              │
    │                                              v
    │                                    예약 확정 (confirmed)
    │                                              │
    v                                              v
[reserved] ─────────────────────────── 채팅방 (reservation_confirmed)
    │                                              │
    │  (예약시각 2시간 이내)                          │
    v                                              v
[pending_trade] ────────────────────── 채팅방 (trade_due)
    │                                              │
    │  작성자 완료 요청 (completionStage=requested)  │
    v                                              │
[completed] ────────────────────────────────────────┤
    │                                              │
    ├── 상대방 확인 ──> confirmed_by_counterparty   │
    │                                              │
    ├── 24h 무응답 ──> auto_confirmed               │
    │                                              │
    └── 이의 제기 ──> disputed                      │
         │                                         │
         v                                         v
    Dispute 처리 ──────────────────── 채팅방 (deal_completed 또는 report_locked)
         │
         ├── resolved_completed ──> 후기 작성 가능
         │
         └── resolved_not_completed ──> 매물 재오픈 또는 취소
```

---

## 부록: 상태 변경 공통 규칙

### A. actorType 분리 저장

모든 상태 변경 이벤트에 아래 정보를 기록한다.

| 필드 | 설명 |
|---|---|
| `actorType` | `user` / `system` / `admin` |
| `actorId` | 실행자 식별자 |
| `changedAt` | 변경 시각 |
| `previousStatus` | 이전 상태 |
| `newStatus` | 새 상태 |
| `reasonCode` | 변경 사유 코드 |

### B. 상태 이력 저장 원칙

- 매물/예약/신고/완료 객체는 현재 상태 + 상태 이력 저장 구조를 기본안으로 가정
- 종결 상태(completed, cancelled, fulfilled, resolved, rejected)에서의 역방향 전이는 원칙적으로 금지
- 운영 예외가 필요한 경우 별도 이력형 복구 처리로 대응

### C. 동시성 제약

- 하나의 매물에 여러 채팅방이 생길 수 있음
- 하나의 활성 예약이 존재하면 동일 매물에 다른 예약 확정은 불가
- `reserved` 또는 `pending_trade` 진입 시 어떤 채팅 상대와 진행 중인지 연결 필수
