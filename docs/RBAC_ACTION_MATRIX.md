# RBAC 액션 매트릭스

> PRD 섹션 5, 24, 36, 50, 67, 69에서 파생된 역할 기반 접근 제어(RBAC) 정의서

---

## 1. 역할(Role) 정의

### 1.1 일반 사용자 역할

| 역할 | 코드 | 설명 |
|---|---|---|
| 비회원 | `guest` | 로그인 이전 탐색 사용자. 매물 목록/상세 일부 조회만 가능 |
| 회원 | `user` | 로그인 완료 일반 사용자. 매물 등록, 채팅, 찜, 신고 가능 |
| 거래참여자 | `participant` | 특정 매물의 채팅/예약/완료 흐름에 실제 참여한 사용자 |
| 작성자 | `owner` | 해당 매물의 등록자. 본인 매물에 대한 관리 권한 보유 |

### 1.2 운영 역할

| 역할 | 코드 | 설명 | 기본 범위 |
|---|---|---|---|
| CS 오퍼레이터 | `cs_operator` | 접수/분류/기본 응답 담당 | 조회, triage, 비파괴 메모/할당 |
| 모더레이터 | `moderator` | 일반 정책 집행 담당 | 경고, 임시 숨김, 저위험 제한, 채팅 잠금 |
| 시니어 모더레이터 | `senior_moderator` | 고위험 사건/복구/민감 열람 담당 | 기간 정지, 복구 승인, 원본 열람 일부 |
| 관리자 | `admin` | 정책 owner 및 최고 권한 | 영구 제재, export 승인, break-glass 승인 |
| 시스템 | `system` | 배치/자동화 actor | 자동 만료, 자동확정, 알림/큐 처리 |

### 1.3 역할 판단 원칙

- 역할은 사람의 직급이 아니라 **기본 권한 묶음**이다.
- 하나의 사용자는 객체마다 다른 역할을 가질 수 있다.
- 실제 실행 가능 여부는 `role + actionCode + case scope + approval requirement`의 교집합으로 판단한다.
- 인가 판단 우선순위:
  1. 법적/운영 강제 제한 (`blocked`, `report_locked`, 계정 정지)
  2. 객체 소유/참여 여부 (작성자, 거래참여자, 신고 당사자)
  3. 공개 범위 (`public`, `hidden`, `blocked`)
  4. 기능별 추가 상태 조건 (`reserved`, `pending_trade`, `completed` 등)
  5. 사용자 개인 설정 (알림, 뮤트 등)

---

## 2. 도메인별 리소스-액션 매트릭스

### 2.1 매물(Listing) 액션

| 액션 | actionCode | guest | user | participant | owner | moderator | senior_moderator | admin | system |
|---|---|---|---|---|---|---|---|---|---|
| 공개 목록 조회 | `listing.list.read` | O (정책 범위) | O | O | O | O | O | O | O |
| 비공개/숨김 매물 조회 | `listing.hidden.read` | X | X | 제한적 | 본인 것만 | O | O | O | O |
| 매물 상세 조회 | `listing.detail.read` | 부분 | O | O | O | O | O | O | O |
| 매물 등록 | `listing.item.create` | X | O | O | - | X | X | X | X |
| 매물 수정 | `listing.item.update` | X | X | X | 상태 조건 충족 시 | 제한적 | 제한적 | O | X |
| 매물 상태 변경 | `listing.status.update` | X | X | X | O | 강제 변경 | 강제 변경 | 강제 변경 | 자동 변경 |
| 매물 삭제(소프트) | `listing.item.delete` | X | X | X | O | X | 복구/강제종료 | O | X |
| 매물 임시 숨김 | `listing.moderation.hide_temp` | X | X | X | 자가 숨김 | O | O | O | X |
| 매물 정책 숨김 | `listing.moderation.hide_policy` | X | X | X | X | O | O | O | O |
| 매물 차단 | `listing.moderation.block` | X | X | X | X | O | O | O | O |
| 매물 복구 | `listing.moderation.restore` | X | X | X | X | X | O | O | X |

**소유권 규칙:**
- 작성자(owner)는 본인 매물에 채팅을 시작할 수 없다.
- `hidden`, `blocked` 매물은 공개 검색 결과에서 제외되며, 직접 URL 접근도 제한한다.
- 거래참여자는 종결 후에도 자신이 참여한 거래 기록으로서 제한 조회 권한을 가진다.

### 2.2 채팅(Chat) 액션

| 액션 | actionCode | guest | user | participant | owner | moderator | senior_moderator | admin | system |
|---|---|---|---|---|---|---|---|---|---|
| 채팅방 목록 조회 | `chat.room.list` | X | 본인 참여방만 | O | O | 권한 범위 내 | O | O | O |
| 채팅 시작 | `chat.room.create` | X | 상태 조건 충족 시 | 재진입 | X (본인 매물) | X | X | X | X |
| 메시지 조회 | `chat.message.read` | X | 본인 참여방만 | O | O | 신고/운영 범위 | O | O | O |
| 메시지 발송 | `chat.message.send` | X | 참여 중 + 상태 허용 | O | O | X | X | X | 시스템 메시지 |
| 사용자 차단 | `chat.user.block` | X | 참여 중 | O | O | 운영 액션 별도 | 운영 액션 별도 | 운영 액션 별도 | X |
| 채팅 잠금 | `chat.moderation.lock` | X | X | X | X | O | O | O | O |
| 채팅 잠금 해제 | `chat.moderation.unlock` | X | X | X | X | X | O | O | X |
| 메시지 마스킹 | `chat.moderation.mask_message` | X | X | X | X | O | O | O | X |

**소유권 규칙:**
- 채팅 메시지 원문은 참여자와 운영자 외 공개 금지.
- 운영자도 기본은 마스킹/필요 최소 열람 원칙을 적용한다.
- `report_locked` 상태에서는 참여자도 신규 메시지 발송 불가.
- 자기 잠금 해제 단독 처리 금지 (senior_moderator 이상 + 본인이 잠근 건은 다른 사람이 해제).

### 2.3 예약(Reservation) 액션

| 액션 | actionCode | guest | user | participant | owner | moderator | senior_moderator | admin | system |
|---|---|---|---|---|---|---|---|---|---|
| 예약 제안 | `reservation.item.propose` | X | 참여 중 | O | O | X | X | X | X |
| 예약 수정 | `reservation.item.update` | X | X | 제안자만 | O | X | X | X | X |
| 예약 확정 | `reservation.item.confirm` | X | X | 상대방만 | O | X | X | X | X |
| 예약 취소 | `reservation.item.cancel` | X | X | 양측 가능 | O | X | 강제 취소 | 강제 취소 | 자동 만료 |
| 예약 거절 | `reservation.item.reject` | X | X | 상대방만 | O | X | X | X | X |
| 예약 상세 조회 | `reservation.item.read` | X | X | O | O | O | O | O | O |

**소유권 규칙:**
- 생성/수정/확정/취소는 해당 채팅방 참여자만 가능.
- 운영자는 강제 만료/강제 취소 가능하나, 사용자 예약 내용을 임의 수정하지 않는다.
- 예약 메타데이터 중 캐릭터명/장소/시간은 공개 매물 영역에 노출하지 않는다.
- 제3자가 상대방이 제안한 예약을 조작할 수 없어야 한다.

### 2.4 거래(Trade) 액션

| 액션 | actionCode | guest | user | participant | owner | moderator | senior_moderator | admin | system |
|---|---|---|---|---|---|---|---|---|---|
| 거래 완료 처리 | `trade.completion.complete` | X | X | O | O | X | X | X | 자동확정 |
| 거래 취소 처리 | `trade.completion.cancel` | X | X | O | O | X | O | O | X |
| 거래 완료 이의제기 | `trade.completion.dispute` | X | X | 상대방 | 상대방 | X | O | O | X |
| 거래 재오픈 | `trade.completion.reopen` | X | X | X | X | X | O | O | X |
| 거래 기록 조회 | `trade.completion.read` | X | X | O | O | O | O | O | O |
| 노쇼 보고 | `trade.noshow.report` | X | X | O | O | X | X | X | X |

**소유권 규칙:**
- 완료 요청은 실제 진행 채팅방 참여자만 가능.
- 완료 확정/이의제기는 상대방 또는 운영자만 가능.
- 완료 기록은 작성자/상대방/운영자만 전체 조회 가능.

### 2.5 후기(Review) 액션

| 액션 | actionCode | guest | user | participant | owner | moderator | senior_moderator | admin | system |
|---|---|---|---|---|---|---|---|---|---|
| 후기 작성 | `review.item.create` | X | X | 해당 completion 당사자만 | 해당 completion 당사자만 | X | X | X | X |
| 후기 조회 | `review.item.read` | 부분 공개 | O | O | O | O | O | O | O |
| 후기 숨김 | `review.moderation.hide` | X | X | X | X | O | O | O | O |
| 후기 복구 | `review.moderation.restore` | X | X | X | X | X | O | O | X |
| 후기 신고 | `review.report.create` | X | O | O | O | X | X | X | X |

**소유권 규칙:**
- 후기 작성은 해당 completion의 당사자만 가능.
- 공개 후기는 전체 회원/비회원에게 일부 또는 전체 공개 가능하나, 작성자 식별/운영 숨김 여부 정책을 따른다.
- 후기 신고는 로그인 사용자라면 누구나 가능.

### 2.6 신고(Report) 액션

| 액션 | actionCode | guest | user | reporter | target | cs_operator | moderator | senior_moderator | admin |
|---|---|---|---|---|---|---|---|---|---|
| 신고 접수 | `report.case.create` | X | O | - | - | X | X | X | X |
| 내 신고 상태 확인 | `report.case.read_status` | X | X | 접수/처리중/완료 | X | O | O | O | O |
| 신고 상세 조회 | `report.case.read_detail` | X | X | 본인 작성분 일부 | X | O | O | O | O |
| 사건 할당 | `report.case.assign` | X | X | X | X | O | O | O | O |
| 사건 분류 | `report.case.triage` | X | X | X | X | O | O | O | O |
| 추가 정보 요청 | `report.case.request_info` | X | X | X | X | O | O | O | O |
| 사건 처리(판정) | `report.case.process` | X | X | X | X | X | X | O | O |
| 사건 종결 | `report.case.close` | X | X | X | X | X | X | O | O |
| 증빙 열람(마스킹) | `evidence.asset.view_masked` | X | X | 본인 제출분만 | X | O | O | O | O |
| 증빙 열람(원본) | `evidence.asset.view_raw` | X | X | X | X | X | X | O | O |

**소유권 규칙:**
- 피신고자(target)에게는 신고자 신원과 내부 메모를 공개하지 않는다.
- 신고자(reporter)는 "조치 완료/미조치" 등 결과 수준만 확인 가능.
- 운영 내부 메모, 위험 점수, 탐지 룰 적중 여부는 외부 비공개.
- 제재 이력 전체 조회는 운영자만 가능하며, 피신고자는 본인 계정 관련 일부만 확인 가능.

### 2.7 사용자 제재(User Restriction) 액션

| 액션 | actionCode | cs_operator | moderator | senior_moderator | admin | system |
|---|---|---|---|---|---|---|
| 경고 | `user.restriction.warn` | X | O | O | O | X |
| 매물 등록 제한 | `user.restriction.limit_listing` | X | O | O | O | O |
| 채팅 제한 | `user.restriction.limit_chat` | X | O | O | O | O |
| 임시 정지 | `user.restriction.suspend_temp` | X | X | O | O | X |
| 영구 이용 제한 | `user.restriction.ban_permanent` | X | X | X | O | X |
| 제재 복구 | `user.restriction.restore` | X | X | O | O | X |

**제재 단계 (scope 분류):**

| scope | 의미 | 사용자 영향 |
|---|---|---|
| `warning_only` | 기능 차단 없는 경고 | 배너/알림만 노출 |
| `listing_only` | 매물 등록/수정/끌어올리기 제한 | 채팅/기록 조회는 가능 |
| `chat_only` | 신규 채팅/메시지/예약 제한 | 매물 조회/관리 일부 가능 |
| `trade_execution_only` | 예약/완료/후기 등 거래 실행 단계 제한 | 탐색/문의는 가능할 수 있음 |
| `trust_limited` | 노출/랭킹/상대 신뢰 표시 제한 | 검색 노출 감점, 배지 하향 |
| `read_only_account` | 대부분 쓰기 차단, 기록/이의제기만 허용 | 읽기/증빙 제출만 가능 |
| `temporary_suspend` | 기간 정지 | 로그인 또는 쓰기 전면 제한 |
| `permanent_ban` | 영구 이용 제한 | 복구는 이의제기 경로로만 |

**우선순위:** `정지 > 기능 제한 > 차단 > 일반 상태`

### 2.8 관리/운영(Admin) 액션

| 액션 | actionCode | cs_operator | moderator | senior_moderator | admin |
|---|---|---|---|---|---|
| 감사 로그 조회 | `admin.audit.view_logs` | X | X | 제한적 | O |
| 데이터 export | `evidence.asset.export` | X | X | X | O |
| 원본 다운로드 | `evidence.asset.download_raw` | X | X | O | O |
| 사용자 관리 | `admin.user.manage` | X | X | X | O |
| 긴급 우회(break-glass) | `admin.break_glass.execute` | X | X | X | O |
| 정책 override | `policy.override.apply` | X | X | X | O |

**분쟁 판정 액션:**

| 액션 | actionCode | cs_operator | moderator | senior_moderator | admin |
|---|---|---|---|---|---|
| 거래완료 판정 | `dispute.case.resolve_completed` | X | X | O | O |
| 거래불발 판정 | `dispute.case.resolve_not_completed` | X | X | O | O |
| 무조치 종결 | `dispute.case.close_no_fault` | X | X | O | O |

---

## 3. 소유권(Ownership) 규칙

### 3.1 매물 기준 관계 분류

| 관계 | 코드 | 설명 | 예시 |
|---|---|---|---|
| 매물 작성자 | `owner` | 해당 매물을 등록한 사용자 | 매물 수정, 상태 변경, 자가 숨김 |
| 상대방(거래참여자) | `counterparty` | 채팅/예약/거래에 참여한 상대 사용자 | 예약 확정/거절, 거래 완료 동의 |
| 제3자 | `third_party` | 해당 매물/거래와 무관한 일반 사용자 | 조회, 찜, 신고만 가능 |

### 3.2 관계별 핵심 권한 차이

| 액션 | 작성자(owner) | 상대방(counterparty) | 제3자(third_party) |
|---|---|---|---|
| 매물 수정 | O (상태 조건 충족 시) | X | X |
| 매물 상태 변경 | O | X | X |
| 매물 자가 숨김 | O | X | X |
| 채팅 시작 | X (본인 매물) | 재진입 | O (상태 조건 충족 시) |
| 예약 제안 | O | O | X |
| 예약 확정 | 제안자 아닌 경우 | 제안자 아닌 경우 | X |
| 거래 완료 제안 | O | O | X |
| 거래 완료 확정 | 상대방 확인 필요 | 상대방 확인 필요 | X |
| 후기 작성 | completion 당사자 시 | completion 당사자 시 | X |
| 신고 접수 | O | O | O |
| 비공개 매물 조회 | 본인 것만 | 참여 기록 범위 | X |

### 3.3 예약 내 역할 구분

| 역할 | 가능 액션 (proposed 상태) | 가능 액션 (confirmed 상태) |
|---|---|---|
| 제안자 | 수정, 취소 | 취소 요청, 시간조정 요청 |
| 상대방 | 수락, 거절, 대안 제시 | 취소 요청, 시간조정 요청 |

---

## 4. 액션 코드(Action Code) 네이밍 컨벤션

### 4.1 형식

```
{domain}.{object}.{verb}
```

### 4.2 도메인(domain) 목록

| domain | 설명 |
|---|---|
| `listing` | 매물 |
| `chat` | 채팅 |
| `reservation` | 예약 |
| `trade` | 거래 |
| `review` | 후기 |
| `report` | 신고 |
| `user` | 사용자/계정 |
| `evidence` | 증빙 |
| `dispute` | 분쟁 |
| `admin` | 관리 |
| `policy` | 정책 |

### 4.3 대표 액션 코드 전체 목록

| actionCode | 설명 |
|---|---|
| **매물** | |
| `listing.item.create` | 매물 등록 |
| `listing.item.update` | 매물 수정 |
| `listing.item.delete` | 매물 삭제(소프트) |
| `listing.status.update` | 매물 상태 변경 |
| `listing.moderation.hide_temp` | 매물 임시 숨김 |
| `listing.moderation.hide_policy` | 매물 정책 숨김 |
| `listing.moderation.restore` | 매물 복구 |
| `listing.moderation.block` | 매물 차단 |
| **채팅** | |
| `chat.room.create` | 채팅방 생성 |
| `chat.message.send` | 메시지 발송 |
| `chat.message.read` | 메시지 조회 |
| `chat.user.block` | 사용자 차단 |
| `chat.moderation.lock` | 채팅 잠금 |
| `chat.moderation.unlock` | 채팅 잠금 해제 |
| `chat.moderation.mask_message` | 메시지 마스킹 |
| **예약** | |
| `reservation.item.propose` | 예약 제안 |
| `reservation.item.confirm` | 예약 확정 |
| `reservation.item.cancel` | 예약 취소 |
| `reservation.item.update` | 예약 수정 |
| `reservation.item.reject` | 예약 거절 |
| **거래** | |
| `trade.completion.complete` | 거래 완료 |
| `trade.completion.cancel` | 거래 취소 |
| `trade.completion.reopen` | 거래 재오픈 |
| `trade.completion.dispute` | 거래 이의제기 |
| `trade.noshow.report` | 노쇼 보고 |
| **후기** | |
| `review.item.create` | 후기 작성 |
| `review.item.read` | 후기 조회 |
| `review.moderation.hide` | 후기 숨김 |
| `review.moderation.restore` | 후기 복구 |
| `review.report.create` | 후기 신고 |
| **신고** | |
| `report.case.create` | 신고 접수 |
| `report.case.assign` | 사건 할당 |
| `report.case.triage` | 사건 분류 |
| `report.case.request_info` | 추가 정보 요청 |
| `report.case.process` | 사건 처리 |
| `report.case.close` | 사건 종결 |
| **사용자 제재** | |
| `user.restriction.warn` | 경고 |
| `user.restriction.limit_listing` | 매물 등록 제한 |
| `user.restriction.limit_chat` | 채팅 제한 |
| `user.restriction.suspend_temp` | 임시 정지 |
| `user.restriction.ban_permanent` | 영구 이용 제한 |
| `user.restriction.restore` | 제재 복구 |
| **증빙** | |
| `evidence.asset.view_masked` | 마스킹 증빙 조회 |
| `evidence.asset.view_raw` | 원본 증빙 조회 |
| `evidence.asset.download_raw` | 원본 다운로드 |
| `evidence.asset.export` | 데이터 외부 반출 |
| **분쟁** | |
| `dispute.case.resolve_completed` | 거래완료 판정 |
| `dispute.case.resolve_not_completed` | 거래불발 판정 |
| `dispute.case.close_no_fault` | 무조치 종결 |
| **관리** | |
| `admin.break_glass.execute` | 긴급 우회 실행 |
| `admin.audit.view_logs` | 감사 로그 조회 |
| `admin.user.manage` | 사용자 관리 |
| `policy.override.apply` | 정책 override |

### 4.4 네이밍 원칙

- actionCode는 UI 문구가 아니라 **불변에 가까운 시스템 식별자**로 관리한다.
- 감사 로그, 권한 테이블, OpenAPI, 운영 매뉴얼에서 동일한 문자열을 재사용한다.
- 새로운 액션 추가 시 반드시 `{domain}.{object}.{verb}` 형식을 준수한다.

---

## 5. 승인 정책(Approval Policy)

### 5.1 정책 유형

| 정책 코드 | 설명 | 승인 요건 |
|---|---|---|
| `single_actor` | 단독 처리 가능 | 역할 조건만 충족하면 즉시 실행 |
| `single_actor_with_reason` | 단독 처리 + 사유 필수 | reasonCode 입력 필수 |
| `single_actor_senior` | senior_moderator 이상 단독 처리 | senior_moderator 이상만 실행 가능 |
| `single_actor_senior_with_justification` | senior 이상 + 열람 사유 | 민감 원본 열람 시 적용 |
| `dual_review_recommended` | 1인 실행 가능, 사후 승인 권장 | P1/P2는 사후 승인 권장 |
| `dual_review_required` | 최소 2인 승인 필수 | 실행 전 secondary reviewer 승인 필요 |
| `dual_review_posthoc_required` | 즉시 실행 + 사후 승인 필수 | 긴급 실행 후 24시간 내 사후 검토 |

### 5.2 액션별 승인 정책 매핑

| actionCode 패밀리 | approvalPolicy | 비고 |
|---|---|---|
| `report.case.*` | `single_actor` | 일반 triage/할당은 단독 처리 |
| `listing.moderation.hide_*` | `single_actor_with_reason` | 빠른 조치 우선 |
| `listing.moderation.restore` | `single_actor_senior` | 오조치 복구는 senior 이상 |
| `chat.moderation.lock` | `single_actor_with_reason` | 사건 연결 필수 |
| `chat.moderation.unlock` | `single_actor_senior` | 재노출 리스크 고려 |
| `user.restriction.warn` | `single_actor_with_reason` | 저위험 제재 |
| `user.restriction.limit_*` | `single_actor_with_reason` | duration/사유 필수 |
| `user.restriction.suspend_temp` | `dual_review_recommended` | P1/P2는 사후 승인 권장 |
| `user.restriction.ban_permanent` | `dual_review_required` | **최소 2인 승인 필수** |
| `user.restriction.restore` | `single_actor_senior` | 원조치 참조 필수 |
| `evidence.asset.view_raw` | `single_actor_senior_with_justification` | 민감 원본 열람 |
| `evidence.asset.export` | `dual_review_required` | 외부 반출은 가장 엄격 |
| `dispute.case.resolve_*` | `single_actor_senior` | 증빙/판정 메모 필수 |
| `admin.break_glass.execute` | `dual_review_posthoc_required` | 즉시 실행 후 사후 승인 필수 |
| `policy.override.apply` | `dual_review_required` | 만료시각/영향범위 필수 |

### 5.3 승인 필수 입력 필드

모든 승인 대상 액션은 아래 필드를 포함해야 한다:

| 필드 | 설명 | 필수 여부 |
|---|---|---|
| `reasonCode` | 구조화된 사유 코드 | 필수 |
| `noteText` | 자유 메모 | 선택 (reasonCode 없이 단독 사용 금지) |
| `caseRefId` | 연결 사건 식별자 | 관련 사건 있을 시 필수 |
| `approvedByUserIds` | 승인자 목록 | dual_review 시 필수 |
| `postReviewDueAt` | 사후 검토 기한 | posthoc 정책 시 필수 |

---

## 6. 자기사건(Self-Case) 방지 규칙

### 6.1 금지 원칙

1. 운영자는 자신이 **신고자/피신고자/거래당사자**인 사건을 직접 판정할 수 없다.
2. 자신의 이전 액션을 스스로 최종 승인할 수 없다.
3. 자기 제재를 스스로 복구할 수 없다.

### 6.2 이해충돌(Conflict) 플래그

시스템은 아래 conflict flag를 자동 계산해야 한다:

| 플래그 | 설명 | 발생 시 조치 |
|---|---|---|
| `isReporterConflict` | 현재 운영자가 해당 사건의 신고자인 경우 | 실행 차단, 다른 reviewer 할당 요구 |
| `isTargetConflict` | 현재 운영자가 해당 사건의 피신고자/대상인 경우 | 실행 차단 |
| `isPriorActorConflict` | 현재 운영자가 이전 조치를 실행한 당사자인 경우 | 최종 승인/복구 차단 |
| `isSecondaryReviewerConflict` | 현재 운영자가 이미 secondary reviewer로 지정된 경우 | 중복 승인 차단 |

### 6.3 conflict 발생 시 처리 흐름

1. 시스템이 사건 상세 조회 시 conflict flag를 계산한다.
2. conflict 발생 시 UI에서 해당 액션 버튼을 비활성화한다.
3. API에서도 동일하게 실행을 차단하고 `ADMIN_CONFLICT_OF_INTEREST` 에러를 반환한다.
4. 다른 reviewer 할당을 자동 요구한다.

### 6.4 break-glass 예외

- 긴급 상황(대규모 개인정보 노출, 명백한 사기 파급, 법적/보안 긴급상황)에서는 self-case 방지 규칙을 우회할 수 있다.
- 단, 아래 조건을 모두 충족해야 한다:
  - `admin` 역할 이상
  - `emergencyReasonCode` 입력
  - `impactScope` 명시
  - `expectedRollbackPlan` 기록
  - `postReviewDueAt` 설정 (24시간 내)
- break-glass 실행 후 24시간 내 사후 검토 기록이 없으면 운영 경보를 발생시킨다.

---

## 7. 운영 역할별 허용 액션 통합 매트릭스

> O = 가능, X = 불가, 조건 = 추가 조건 충족 시 가능

| actionCode | cs_operator | moderator | senior_moderator | admin | 추가 조건 |
|---|---|---|---|---|---|
| `report.case.assign` | O | O | O | O | 자기 자신에게 재할당 허용 |
| `report.case.triage` | O | O | O | O | reasonCode 필수 |
| `report.case.request_info` | O | O | O | O | dueAt 권장 |
| `listing.moderation.hide_temp` | X | O | O | O | reportId 또는 caseRef 필요 |
| `listing.moderation.hide_policy` | X | O | O | O | policyCode 필수 |
| `listing.moderation.restore` | X | X | O | O | 복구 사유 필수 |
| `chat.moderation.lock` | X | O | O | O | 사건 연결 필수 |
| `chat.moderation.unlock` | X | X | O | O | 자기 잠금 해제 단독 처리 금지 |
| `chat.moderation.mask_message` | X | O | O | O | 원문 보관 전제 |
| `user.restriction.warn` | X | O | O | O | expiry 없음 |
| `user.restriction.limit_listing` | X | O | O | O | duration/사유 필수 |
| `user.restriction.limit_chat` | X | O | O | O | duration/사유 필수 |
| `user.restriction.suspend_temp` | X | X | O | O | 기간 제한 필수 |
| `user.restriction.ban_permanent` | X | X | X | O | **2인 승인 필수** |
| `user.restriction.restore` | X | X | O | O | 원조치 참조 필수 |
| `dispute.case.resolve_completed` | X | X | O | O | 증빙/판정 메모 필수 |
| `dispute.case.resolve_not_completed` | X | X | O | O | 재오픈/취소 후속액션 명시 |
| `dispute.case.close_no_fault` | X | X | O | O | 양측 통지 필요 |
| `evidence.asset.view_masked` | O | O | O | O | 사건 컨텍스트 내 |
| `evidence.asset.view_raw` | X | X | O | O | 열람 사유 + 감사 태그 |
| `evidence.asset.download_raw` | X | X | O | O | 시간제한 URL/승인 정책 |
| `evidence.asset.export` | X | X | X | O | **2인 승인 + export 목적 필수** |
| `admin.break_glass.execute` | X | X | X | O | **사후 리뷰 강제** |
| `policy.override.apply` | X | X | X | O | 만료시각/영향범위 필수 |

---

## 8. 감사 로그 최소 필드

모든 운영 액션은 아래 필드를 기록해야 한다:

| 필드 | 설명 |
|---|---|
| `auditLogId` | 감사 로그 식별자 |
| `actionCode` | 실행 액션 코드 |
| `actorUserId` | 실행 운영자 |
| `actorRole` | 실행 시점 역할 |
| `targetType` / `targetId` | 대상 객체 |
| `caseRefType` / `caseRefId` | 연결 사건 |
| `reasonCode` | 구조화 사유 |
| `noteText` | 자유 메모 |
| `approvalPolicy` | 적용된 승인 정책 |
| `approvedByUserIds` | 승인자 목록 |
| `conflictCheckResult` | self-case 검사 결과 |
| `beforeSnapshot` / `afterSnapshot` | 핵심 diff |
| `createdAt` | 실행 시각 |
| `breakGlass` | 긴급 우회 여부 |

---

## 9. API 응답 계약 시사점

백오피스 상세 API는 아래 구조를 함께 반환해야 한다:

```json
{
  "caseId": "rep_123",
  "viewerRole": "senior_moderator",
  "allowedAdminActions": [
    "chat.moderation.lock",
    "user.restriction.suspend_temp",
    "evidence.asset.view_masked"
  ],
  "approvalRequirements": {
    "user.restriction.suspend_temp": "dual_review_recommended",
    "evidence.asset.export": "not_allowed_for_viewer"
  },
  "conflictFlags": {
    "isReporterConflict": false,
    "isTargetConflict": false,
    "isPriorActorConflict": true
  }
}
```

- 프론트는 역할명만 보고 버튼을 노출하지 않고, `allowedAdminActions`와 `approvalRequirements`를 우선 사용한다.
- 같은 역할이어도 사건 종류/소유 여부/conflict에 따라 허용 액션이 달라질 수 있다.
- 403 응답은 `ADMIN_CONFLICT_OF_INTEREST`, `ADMIN_APPROVAL_REQUIRED`, `ADMIN_ROLE_NOT_GRANTED` 같은 typed error를 제공한다.
