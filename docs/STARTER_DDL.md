# 린클 초기 데이터베이스 DDL

> PostgreSQL 16 기준 초기 마이그레이션 SQL
>
> PRD 엔티티 정의(섹션 20, 27, 45)와 기술 스택 문서를 기반으로 작성
>
> 작성일: 2026-03-14
> 상태: 초안

---

## 문서 목적

이 문서는 린클 서비스의 PostgreSQL 데이터베이스 초기 스키마를 정의한다.
아래 SQL을 순서대로 실행하면 모든 테이블, 인덱스, 시드 데이터가 생성된다.
`backend/db/migrations/` 디렉토리의 첫 번째 마이그레이션 파일로 사용할 수 있다.

### 포함 범위

- ENUM 타입 정의
- 마스터 데이터 테이블 (servers, categories)
- 사용자 도메인 (users, user_profiles, block_relations, push_tokens)
- 매물 도메인 (listings, listing_images, favorites)
- 채팅 도메인 (chat_rooms, chat_messages, chat_read_cursors)
- 거래 도메인 (reservations, trade_completions, reviews)
- 신고/운영 도메인 (reports, moderation_actions)
- 알림 (notifications)
- 감사/이력 (audit_logs, status_history)
- 인덱스 정의
- 시드 데이터 (게임 서버, 아이템 카테고리)

---

## 1. 확장 기능 활성화 및 ENUM 타입 정의

```sql
-- ============================================================
-- 확장 기능
-- ============================================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";  -- 아이템명 부분 검색용

-- ============================================================
-- ENUM 타입 정의
-- ============================================================

-- 사용자 로그인 수단
CREATE TYPE login_provider AS ENUM (
    'kakao',
    'naver',
    'google',
    'apple',
    'phone'
);

-- 사용자 계정 상태
CREATE TYPE account_status AS ENUM (
    'active',
    'restricted',
    'suspended',
    'withdrawn'
);

-- 사용자 역할
CREATE TYPE user_role AS ENUM (
    'user',
    'moderator',
    'admin'
);

-- 응답 배지
CREATE TYPE response_badge AS ENUM (
    'fast',
    'normal',
    'slow'
);

-- 신뢰 배지
CREATE TYPE trust_badge AS ENUM (
    'newcomer',
    'trading',
    'experienced',
    'trusted'
);

-- 매물 거래 유형
CREATE TYPE listing_type AS ENUM (
    'sell',
    'buy'
);

-- 매물 가격 유형
CREATE TYPE price_type AS ENUM (
    'fixed',
    'negotiable',
    'offer'
);

-- 재화 유형
CREATE TYPE currency_type AS ENUM (
    'adena',
    'krw',
    'mixed'
);

-- 거래 방식
CREATE TYPE trade_method AS ENUM (
    'in_game',
    'offline_pc_bang',
    'either'
);

-- 매물 상태
CREATE TYPE listing_status AS ENUM (
    'available',
    'reserved',
    'pending_trade',
    'completed',
    'cancelled'
);

-- 매물 노출 상태
CREATE TYPE listing_visibility AS ENUM (
    'public',
    'hidden',
    'blocked'
);

-- 채팅방 상태
CREATE TYPE chat_status AS ENUM (
    'open',
    'reservation_proposed',
    'reservation_confirmed',
    'trade_due',
    'deal_completed',
    'deal_cancelled',
    'report_locked'
);

-- 채팅 메시지 유형
CREATE TYPE message_type AS ENUM (
    'text',
    'system',
    'reservation_card',
    'image'
);

-- 예약 상태
CREATE TYPE reservation_status AS ENUM (
    'proposed',
    'confirmed',
    'expired',
    'cancelled',
    'fulfilled',
    'no_show_reported'
);

-- 거래 완료 상태
CREATE TYPE completion_status AS ENUM (
    'requested',
    'confirmed',
    'disputed',
    'closed'
);

-- 후기 추천 여부
CREATE TYPE recommendation_type AS ENUM (
    'recommend',
    'not_recommend'
);

-- 후기 노출 상태
CREATE TYPE review_visibility AS ENUM (
    'visible',
    'hidden',
    'pending_moderation'
);

-- 신고 대상 유형
CREATE TYPE report_target_type AS ENUM (
    'user',
    'listing',
    'chat_room',
    'message',
    'review'
);

-- 신고 사유
CREATE TYPE report_reason_code AS ENUM (
    'fake_listing',
    'scam_suspicion',
    'no_show',
    'abuse',
    'spam',
    'prohibited_item',
    'privacy_exposure',
    'other'
);

-- 신고 처리 상태
CREATE TYPE report_status AS ENUM (
    'submitted',
    'triaged',
    'investigating',
    'resolved',
    'rejected'
);

-- 신고 우선순위
CREATE TYPE report_priority AS ENUM (
    'P1',
    'P2',
    'P3',
    'P4'
);

-- 운영 조치 유형
CREATE TYPE moderation_action_type AS ENUM (
    'warn',
    'hide_listing',
    'restrict_chat',
    'suspend',
    'ban'
);

-- 알림 유형
CREATE TYPE notification_type AS ENUM (
    'chat',
    'reservation',
    'status',
    'review',
    'report',
    'system'
);

-- 상태 이력 대상 엔티티 유형
CREATE TYPE status_entity_type AS ENUM (
    'listing',
    'reservation',
    'chat_room',
    'trade_completion',
    'report'
);

-- 상태 변경 주체 유형
CREATE TYPE actor_type AS ENUM (
    'user',
    'system',
    'admin'
);

-- 디바이스 플랫폼
CREATE TYPE device_platform AS ENUM (
    'ios',
    'android',
    'web'
);
```

---

## 2. 마스터 데이터 테이블

```sql
-- ============================================================
-- 게임 서버 마스터
-- 리니지 클래식 서버 목록. 매물/프로필에서 FK로 참조한다.
-- ============================================================
CREATE TABLE servers (
    server_id   VARCHAR(32)  PRIMARY KEY,           -- 예: 'depolloju', 'ken-raisa'
    name_ko     VARCHAR(64)  NOT NULL,              -- 한국어 서버명
    name_en     VARCHAR(64),                        -- 영문 서버명 (선택)
    region      VARCHAR(32)  NOT NULL DEFAULT 'kr', -- 지역 코드
    sort_order  INTEGER      NOT NULL DEFAULT 0,    -- 정렬 순서
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE, -- 비활성 서버 숨김용
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now()
);

COMMENT ON TABLE  servers IS '게임 서버 마스터 테이블. 리니지 클래식 서버 목록을 관리한다.';
COMMENT ON COLUMN servers.server_id IS '서버 식별 슬러그. URL/필터에 사용한다.';

-- ============================================================
-- 아이템 카테고리 마스터
-- 매물 등록 시 카테고리 선택에 사용한다.
-- ============================================================
CREATE TABLE categories (
    category_id VARCHAR(32)  PRIMARY KEY,           -- 예: 'weapon', 'armor'
    name_ko     VARCHAR(64)  NOT NULL,              -- 한국어 카테고리명
    name_en     VARCHAR(64),                        -- 영문명 (선택)
    parent_id   VARCHAR(32)  REFERENCES categories(category_id), -- 상위 카테고리 (계층 구조)
    sort_order  INTEGER      NOT NULL DEFAULT 0,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now()
);

COMMENT ON TABLE  categories IS '아이템 카테고리 마스터 테이블. 계층 구조를 지원한다.';
COMMENT ON COLUMN categories.parent_id IS 'NULL이면 최상위 카테고리.';
```

---

## 3. 사용자 도메인

```sql
-- ============================================================
-- users: 계정 인증 정보
-- 소셜 로그인 기반 사용자 계정. 탈퇴 시 soft delete (withdrawn_at).
-- ============================================================
CREATE TABLE users (
    user_id                 UUID         PRIMARY KEY DEFAULT uuid_generate_v4(),
    login_provider          login_provider NOT NULL,
    login_provider_user_key VARCHAR(255) NOT NULL,       -- 외부 OAuth 식별자
    account_status          account_status NOT NULL DEFAULT 'active',
    role                    user_role    NOT NULL DEFAULT 'user',
    last_login_at           TIMESTAMPTZ,
    created_at              TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at              TIMESTAMPTZ  NOT NULL DEFAULT now(),
    withdrawn_at            TIMESTAMPTZ,                 -- soft delete 시각

    CONSTRAINT uq_users_provider_key UNIQUE (login_provider, login_provider_user_key)
);

COMMENT ON TABLE  users IS '사용자 계정 테이블. 인증/권한 정보를 관리한다.';
COMMENT ON COLUMN users.login_provider_user_key IS '소셜 로그인 공급자에서 발급한 사용자 고유 키.';
COMMENT ON COLUMN users.withdrawn_at IS '탈퇴 시각. NULL이 아니면 탈퇴 상태.';

-- ============================================================
-- user_profiles: 서비스 프로필 (공개 정보)
-- users와 1:1 관계. 닉네임, 프로필 이미지, 신뢰 지표 캐시를 저장한다.
-- ============================================================
CREATE TABLE user_profiles (
    user_id                UUID         PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    nickname               VARCHAR(30)  NOT NULL,
    avatar_url             VARCHAR(512),
    introduction           VARCHAR(500),
    primary_server_id      VARCHAR(32)  REFERENCES servers(server_id) ON DELETE SET NULL,
    response_badge         response_badge,
    trust_badge            trust_badge  DEFAULT 'newcomer',
    completed_trade_count  INTEGER      NOT NULL DEFAULT 0,
    positive_review_count  INTEGER      NOT NULL DEFAULT 0,
    last_active_at         TIMESTAMPTZ,
    created_at             TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at             TIMESTAMPTZ  NOT NULL DEFAULT now(),

    CONSTRAINT uq_user_profiles_nickname UNIQUE (nickname)
);

COMMENT ON TABLE  user_profiles IS '사용자 공개 프로필. 닉네임/배지/거래 통계 캐시를 포함한다.';
COMMENT ON COLUMN user_profiles.completed_trade_count IS '완료 거래 수 캐시. 원본 데이터에서 재계산 가능해야 한다.';
COMMENT ON COLUMN user_profiles.positive_review_count IS '긍정 후기 수 캐시. 원본 데이터에서 재계산 가능해야 한다.';

-- ============================================================
-- block_relations: 사용자 간 차단 관계
-- blockerUserId + blockedUserId 조합으로 유니크.
-- ============================================================
CREATE TABLE block_relations (
    id               UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    blocker_user_id  UUID        NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    blocked_user_id  UUID        NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_block_relations UNIQUE (blocker_user_id, blocked_user_id),
    CONSTRAINT ck_block_not_self  CHECK  (blocker_user_id != blocked_user_id)
);

COMMENT ON TABLE block_relations IS '사용자 간 차단 관계. 차단자가 피차단자의 매물/채팅을 볼 수 없게 한다.';

-- ============================================================
-- push_tokens: 푸시 알림 디바이스 토큰
-- 사용자당 여러 디바이스 등록 가능 (멀티 디바이스 지원).
-- ============================================================
CREATE TABLE push_tokens (
    id          UUID             PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID             NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    platform    device_platform  NOT NULL,
    token       VARCHAR(512)     NOT NULL,
    is_active   BOOLEAN          NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ      NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ      NOT NULL DEFAULT now(),

    CONSTRAINT uq_push_tokens_token UNIQUE (token)
);

COMMENT ON TABLE push_tokens IS '푸시 알림 디바이스 토큰. FCM/APNs 토큰을 관리한다.';
COMMENT ON COLUMN push_tokens.is_active IS '로그아웃/앱 삭제 시 FALSE로 전환하여 stale 푸시를 방지한다.';
```

---

## 4. 매물 도메인

```sql
-- ============================================================
-- listings: 매물 (거래 게시물)
-- 서비스의 핵심 엔티티. 매물 상태가 거래 흐름 전체를 표현한다.
-- ============================================================
CREATE TABLE listings (
    listing_id              UUID             PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_type            listing_type     NOT NULL,              -- sell / buy
    author_user_id          UUID             NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    server_id               VARCHAR(32)      NOT NULL REFERENCES servers(server_id),
    category_id             VARCHAR(32)      NOT NULL REFERENCES categories(category_id),
    item_name               VARCHAR(100)     NOT NULL,              -- 대표 아이템명
    title                   VARCHAR(100)     NOT NULL,              -- 목록 표시 제목
    description             TEXT             NOT NULL,              -- 상세 설명
    price_type              price_type       NOT NULL,              -- fixed / negotiable / offer
    price_amount            BIGINT,                                 -- 가격 (offer 시 NULL 허용)
    currency_type           currency_type,                          -- adena / krw / mixed
    quantity                NUMERIC(18,4)    NOT NULL DEFAULT 1,    -- 수량
    enhancement_level       INTEGER,                                -- 강화 수치
    options_text            TEXT,                                   -- 자유 옵션 텍스트
    trade_method            trade_method     NOT NULL,              -- in_game / offline_pc_bang / either
    preferred_meeting_area  VARCHAR(200),                           -- 접선 지역 텍스트
    available_time_text     VARCHAR(200),                           -- 가능 시간 자유기입
    status                  listing_status   NOT NULL DEFAULT 'available',
    visibility              listing_visibility NOT NULL DEFAULT 'public',
    reserved_chat_room_id   UUID,                                   -- 현재 우선 진행 채팅방 (FK는 chat_rooms 생성 후 추가)
    last_activity_at        TIMESTAMPTZ      NOT NULL DEFAULT now(),-- 정렬/노출용
    view_count              INTEGER          NOT NULL DEFAULT 0,
    favorite_count          INTEGER          NOT NULL DEFAULT 0,
    chat_count              INTEGER          NOT NULL DEFAULT 0,
    created_at              TIMESTAMPTZ      NOT NULL DEFAULT now(),
    updated_at              TIMESTAMPTZ      NOT NULL DEFAULT now(),
    expires_at              TIMESTAMPTZ,                            -- 자동 만료 정책용
    completed_at            TIMESTAMPTZ,                            -- 거래 완료 시각
    deleted_at              TIMESTAMPTZ                             -- soft delete
);

COMMENT ON TABLE  listings IS '매물 테이블. 거래 게시물의 전체 생애주기를 관리한다.';
COMMENT ON COLUMN listings.price_amount IS '가격 (정수). offer 유형이면 NULL 허용. 아데나 단위는 만 아데나 등 표현 정책에 따라 정한다.';
COMMENT ON COLUMN listings.reserved_chat_room_id IS 'status=reserved 이상일 때 우선 진행 중인 채팅방 참조.';
COMMENT ON COLUMN listings.last_activity_at IS '목록 정렬 및 신선도 판단에 사용. 매물 수정/상태 변경/채팅 시 갱신한다.';

-- ============================================================
-- listing_images: 매물 첨부 이미지
-- 매물당 최대 N장. sort_order로 대표 이미지를 결정한다.
-- ============================================================
CREATE TABLE listing_images (
    id          UUID         PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id  UUID         NOT NULL REFERENCES listings(listing_id) ON DELETE CASCADE,
    image_url   VARCHAR(512) NOT NULL,
    sort_order  INTEGER      NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now()
);

COMMENT ON TABLE listing_images IS '매물 첨부 이미지. sort_order=0이 대표 이미지.';

-- ============================================================
-- favorites: 찜 (사용자-매물 관계)
-- ============================================================
CREATE TABLE favorites (
    id          UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID        NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    listing_id  UUID        NOT NULL REFERENCES listings(listing_id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_favorites UNIQUE (user_id, listing_id)
);

COMMENT ON TABLE favorites IS '찜 테이블. 사용자가 관심 매물을 저장한다.';
```

---

## 5. 채팅 도메인

```sql
-- ============================================================
-- chat_rooms: 1:1 채팅방
-- listingId + sellerUserId + buyerUserId 조합으로 유니크.
-- 하나의 매물에 여러 채팅방이 생길 수 있다.
-- ============================================================
CREATE TABLE chat_rooms (
    chat_room_id           UUID         PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id             UUID         NOT NULL REFERENCES listings(listing_id) ON DELETE CASCADE,
    seller_user_id         UUID         NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    buyer_user_id          UUID         NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    chat_status            chat_status  NOT NULL DEFAULT 'open',
    last_message_id        UUID,                                    -- FK는 chat_messages 생성 후 추가
    last_message_at        TIMESTAMPTZ,
    unread_count_seller    INTEGER      NOT NULL DEFAULT 0,         -- 판매자 미읽음 캐시
    unread_count_buyer     INTEGER      NOT NULL DEFAULT 0,         -- 구매자 미읽음 캐시
    closed_at              TIMESTAMPTZ,
    created_at             TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at             TIMESTAMPTZ  NOT NULL DEFAULT now(),

    CONSTRAINT uq_chat_rooms UNIQUE (listing_id, seller_user_id, buyer_user_id),
    CONSTRAINT ck_chat_rooms_diff_users CHECK (seller_user_id != buyer_user_id)
);

COMMENT ON TABLE  chat_rooms IS '1:1 채팅방. 매물 기준으로 판매자와 구매자를 연결한다.';
COMMENT ON COLUMN chat_rooms.unread_count_seller IS '판매자 미읽음 메시지 수 캐시. 읽음 커서 기반으로 재계산 가능.';
COMMENT ON COLUMN chat_rooms.unread_count_buyer IS '구매자 미읽음 메시지 수 캐시. 읽음 커서 기반으로 재계산 가능.';

-- listings.reserved_chat_room_id FK 추가 (순환 참조 해소)
ALTER TABLE listings
    ADD CONSTRAINT fk_listings_reserved_chat_room
    FOREIGN KEY (reserved_chat_room_id)
    REFERENCES chat_rooms(chat_room_id)
    ON DELETE SET NULL;

-- ============================================================
-- chat_messages: 채팅 메시지
-- text, system, reservation_card, image 등 유형을 지원한다.
-- ============================================================
CREATE TABLE chat_messages (
    message_id      UUID          PRIMARY KEY DEFAULT uuid_generate_v4(),
    chat_room_id    UUID          NOT NULL REFERENCES chat_rooms(chat_room_id) ON DELETE CASCADE,
    sender_user_id  UUID          NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    message_type    message_type  NOT NULL DEFAULT 'text',
    body_text       TEXT,                                           -- 본문 (text/system)
    metadata_json   JSONB,                                         -- 예약/상태 변경 부가정보
    sent_at         TIMESTAMPTZ   NOT NULL DEFAULT now(),
    deleted_at      TIMESTAMPTZ                                    -- soft delete
);

COMMENT ON TABLE  chat_messages IS '채팅 메시지. 텍스트/시스템/예약카드/이미지 등 유형을 지원한다.';
COMMENT ON COLUMN chat_messages.metadata_json IS '예약 카드 정보, 상태 변경 이벤트 등 메시지 유형별 부가 데이터.';

-- chat_rooms.last_message_id FK 추가 (순환 참조 해소)
ALTER TABLE chat_rooms
    ADD CONSTRAINT fk_chat_rooms_last_message
    FOREIGN KEY (last_message_id)
    REFERENCES chat_messages(message_id)
    ON DELETE SET NULL;

-- ============================================================
-- chat_read_cursors: 읽음 커서
-- 사용자별 채팅방에서 마지막으로 읽은 메시지를 추적한다.
-- Redis 캐시와 동기화하여 사용한다.
-- ============================================================
CREATE TABLE chat_read_cursors (
    id                  UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    chat_room_id        UUID        NOT NULL REFERENCES chat_rooms(chat_room_id) ON DELETE CASCADE,
    user_id             UUID        NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    last_read_message_id UUID       REFERENCES chat_messages(message_id) ON DELETE SET NULL,
    last_read_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_chat_read_cursors UNIQUE (chat_room_id, user_id)
);

COMMENT ON TABLE chat_read_cursors IS '채팅 읽음 커서. 사용자별 마지막으로 읽은 메시지를 추적한다. Redis 캐시와 DB를 동기화한다.';
```

---

## 6. 거래 도메인

```sql
-- ============================================================
-- reservations: 거래 예약
-- 채팅방 내에서 거래 시간/장소를 제안하고 확정하는 과정을 관리한다.
-- 한 매물에 동시 활성 예약은 기본적으로 1건만 허용한다.
-- ============================================================
CREATE TABLE reservations (
    reservation_id          UUID               PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id              UUID               NOT NULL REFERENCES listings(listing_id) ON DELETE CASCADE,
    chat_room_id            UUID               NOT NULL REFERENCES chat_rooms(chat_room_id) ON DELETE CASCADE,
    proposer_user_id        UUID               NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    counterpart_user_id     UUID               NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    reservation_status      reservation_status NOT NULL DEFAULT 'proposed',
    scheduled_at            TIMESTAMPTZ        NOT NULL,            -- 약속 기준 시각
    time_window_text        VARCHAR(100),                           -- 시간 범위 보조 텍스트 (예: "21:00~21:30")
    meeting_type            trade_method       NOT NULL,            -- in_game / offline_pc_bang / either
    meeting_point_text      VARCHAR(200)       NOT NULL,            -- 접선 장소 텍스트
    server_id               VARCHAR(32)        REFERENCES servers(server_id), -- 인게임 서버
    character_name_a        VARCHAR(50),                            -- 제안자 캐릭터명
    character_name_b        VARCHAR(50),                            -- 상대방 캐릭터명
    note_to_counterparty    TEXT,                                   -- 전달 메모
    expires_at              TIMESTAMPTZ,                            -- 제안 만료 시각
    confirmed_at            TIMESTAMPTZ,                            -- 확정 시각
    cancelled_at            TIMESTAMPTZ,                            -- 취소 시각
    fulfilled_at            TIMESTAMPTZ,                            -- 이행 완료 시각
    cancellation_reason_code VARCHAR(50),                           -- 취소 사유 코드
    created_at              TIMESTAMPTZ        NOT NULL DEFAULT now(),
    updated_at              TIMESTAMPTZ        NOT NULL DEFAULT now()
);

COMMENT ON TABLE  reservations IS '거래 예약. 시간/장소/방식을 구조화하여 저장한다.';
COMMENT ON COLUMN reservations.meeting_type IS '거래 방식. trade_method ENUM을 재사용한다.';
COMMENT ON COLUMN reservations.expires_at IS '제안 상태에서 응답 없으면 자동 만료되는 기한.';

-- ============================================================
-- trade_completions: 거래 완료 기록
-- 매물 종결 기준 별도 이벤트/엔티티로 보관하여 운영 추적에 활용한다.
-- 매물당 활성 완료 기록(requested/confirmed/disputed)은 최대 1건.
-- ============================================================
CREATE TABLE trade_completions (
    completion_id       UUID              PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id          UUID              NOT NULL REFERENCES listings(listing_id) ON DELETE CASCADE,
    chat_room_id        UUID              NOT NULL REFERENCES chat_rooms(chat_room_id) ON DELETE CASCADE,
    completed_by_user_id UUID             NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    completion_status   completion_status NOT NULL DEFAULT 'requested',
    completed_at        TIMESTAMPTZ       NOT NULL DEFAULT now(),   -- 완료 요청 시각
    confirmed_at        TIMESTAMPTZ,                                -- 최종 확정 시각
    created_at          TIMESTAMPTZ       NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ       NOT NULL DEFAULT now()
);

COMMENT ON TABLE  trade_completions IS '거래 완료 기록. 매물 상태와 별도로 종결 이벤트를 관리한다.';
COMMENT ON COLUMN trade_completions.completed_at IS '완료 요청 시각. confirmed_at은 상대방 확인 후 설정.';

-- ============================================================
-- reviews: 후기
-- 거래 완료 후 양측이 상호 후기를 작성한다.
-- completionId + reviewerUserId 기준 유니크 (거래당 1회).
-- ============================================================
CREATE TABLE reviews (
    review_id           UUID              PRIMARY KEY DEFAULT uuid_generate_v4(),
    completion_id       UUID              NOT NULL REFERENCES trade_completions(completion_id) ON DELETE CASCADE,
    reviewer_user_id    UUID              NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    reviewee_user_id    UUID              NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    recommendation      recommendation_type NOT NULL,              -- recommend / not_recommend
    comment_text        VARCHAR(500),                              -- 후기 본문 (선택)
    visibility_status   review_visibility NOT NULL DEFAULT 'visible',
    created_at          TIMESTAMPTZ       NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ       NOT NULL DEFAULT now(),

    CONSTRAINT uq_reviews_per_completion UNIQUE (completion_id, reviewer_user_id),
    CONSTRAINT ck_reviews_not_self       CHECK  (reviewer_user_id != reviewee_user_id)
);

COMMENT ON TABLE  reviews IS '거래 후기. 완료된 거래 건에 대해 양측이 추천/비추천 + 코멘트를 남긴다.';
COMMENT ON COLUMN reviews.visibility_status IS '운영 숨김, 모더레이션 보류 등 후기 노출 상태를 제어한다.';
```

---

## 7. 신고/운영 도메인

```sql
-- ============================================================
-- reports: 신고
-- 사용자/매물/채팅방/메시지/후기를 대상으로 신고를 접수한다.
-- ============================================================
CREATE TABLE reports (
    report_id          UUID               PRIMARY KEY DEFAULT uuid_generate_v4(),
    reporter_user_id   UUID               NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    target_type        report_target_type NOT NULL,                -- user / listing / chat_room / message / review
    target_id          UUID               NOT NULL,                -- 대상 엔티티 ID (다형성 참조)
    report_reason_code report_reason_code NOT NULL,
    description_text   TEXT               NOT NULL,                -- 상세 설명
    evidence_urls      JSONB,                                      -- 증빙 첨부 URL 목록 (확장)
    report_status      report_status      NOT NULL DEFAULT 'submitted',
    priority           report_priority    NOT NULL DEFAULT 'P3',
    created_at         TIMESTAMPTZ        NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ        NOT NULL DEFAULT now(),
    resolved_at        TIMESTAMPTZ                                 -- 해결 시각
);

COMMENT ON TABLE  reports IS '신고 테이블. 다양한 대상(사용자/매물/채팅방/메시지/후기)에 대한 신고를 관리한다.';
COMMENT ON COLUMN reports.target_id IS '다형성 FK. target_type에 따라 참조 대상이 달라진다. 앱 레벨에서 무결성을 보장한다.';
COMMENT ON COLUMN reports.evidence_urls IS 'JSONB 배열. 증빙 이미지/파일 URL을 저장한다.';

-- ============================================================
-- moderation_actions: 운영 조치
-- 경고, 숨김, 채팅 제한, 기간 정지, 영구 차단 등 제재 이력을 관리한다.
-- ============================================================
CREATE TABLE moderation_actions (
    action_id           UUID                    PRIMARY KEY DEFAULT uuid_generate_v4(),
    report_id           UUID                    REFERENCES reports(report_id) ON DELETE SET NULL,
    target_user_id      UUID                    REFERENCES users(user_id) ON DELETE SET NULL,
    target_listing_id   UUID                    REFERENCES listings(listing_id) ON DELETE SET NULL,
    action_type         moderation_action_type  NOT NULL,
    action_reason       TEXT                    NOT NULL,           -- 조치 사유
    memo                TEXT,                                      -- 운영 내부 메모
    starts_at           TIMESTAMPTZ             NOT NULL DEFAULT now(),
    ends_at             TIMESTAMPTZ,                               -- NULL이면 영구 또는 즉시 완료
    is_active           BOOLEAN                 NOT NULL DEFAULT TRUE,
    created_by_admin_id UUID                    NOT NULL REFERENCES users(user_id),
    created_at          TIMESTAMPTZ             NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ             NOT NULL DEFAULT now()
);

COMMENT ON TABLE  moderation_actions IS '운영 조치 이력. 경고/숨김/제한/정지/차단 등 제재를 기록한다.';
COMMENT ON COLUMN moderation_actions.ends_at IS 'NULL이면 영구 조치 또는 즉시 완료형 조치(경고 등).';
COMMENT ON COLUMN moderation_actions.is_active IS 'FALSE면 해제/만료된 조치. 이력은 보존한다.';
```

---

## 8. 알림

```sql
-- ============================================================
-- notifications: 앱 내 알림
-- 채팅, 예약, 상태 변경, 후기, 신고, 시스템 알림을 통합 관리한다.
-- ============================================================
CREATE TABLE notifications (
    notification_id   UUID              PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id           UUID              NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    notification_type notification_type NOT NULL,
    title             VARCHAR(200)      NOT NULL,
    body              TEXT              NOT NULL,
    deep_link         VARCHAR(512),                                -- 앱 내 이동 경로
    reference_id      UUID,                                        -- 관련 엔티티 ID (선택)
    read_at           TIMESTAMPTZ,                                 -- 읽음 시각
    delivered_at      TIMESTAMPTZ,                                 -- 발송 완료 시각
    created_at        TIMESTAMPTZ       NOT NULL DEFAULT now()
);

COMMENT ON TABLE  notifications IS '앱 내 알림함. 다양한 유형의 알림을 통합 저장한다.';
COMMENT ON COLUMN notifications.deep_link IS '알림 탭 시 이동할 앱 내 경로 (예: /chats/xxx, /listings/xxx).';
COMMENT ON COLUMN notifications.reference_id IS '관련 엔티티 ID. notification_type에 따라 매물/채팅방/예약 등을 참조한다.';
```

---

## 9. 감사/이력

```sql
-- ============================================================
-- audit_logs: 감사 로그
-- 운영자 조치, 민감 데이터 열람, 시스템 이벤트를 추적한다.
-- append-only 테이블이다. UPDATE/DELETE를 금지하는 운영 원칙을 적용한다.
-- ============================================================
CREATE TABLE audit_logs (
    id              UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    actor_type      actor_type  NOT NULL,                          -- user / system / admin
    actor_id        UUID,                                          -- 수행자 ID (시스템이면 NULL)
    action          VARCHAR(100) NOT NULL,                         -- 수행 액션 코드 (예: 'report.resolve', 'listing.force_hide')
    target_type     VARCHAR(50),                                   -- 대상 엔티티 유형
    target_id       UUID,                                          -- 대상 엔티티 ID
    details_json    JSONB,                                         -- 변경 전후 데이터, 사유 등
    ip_address      INET,                                          -- 요청 IP
    user_agent      VARCHAR(512),                                  -- 요청 UA
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT ON TABLE  audit_logs IS '감사 로그. append-only. 운영 조치/민감 데이터 열람/시스템 이벤트를 추적한다.';
COMMENT ON COLUMN audit_logs.details_json IS '변경 전후 데이터(before/after), 사유(reason), 근거 링크 등을 JSONB로 저장한다.';

-- ============================================================
-- status_history: 상태 변경 이력 (append-only)
-- Listing, Reservation, ChatRoom, TradeCompletion, Report의
-- 상태 전이를 시간순으로 기록한다. 운영 추적/디버깅에 활용한다.
-- ============================================================
CREATE TABLE status_history (
    id              UUID              PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_type     status_entity_type NOT NULL,                   -- listing / reservation / chat_room / ...
    entity_id       UUID              NOT NULL,                    -- 대상 엔티티 ID
    from_status     VARCHAR(50),                                   -- 이전 상태 (최초 생성 시 NULL)
    to_status       VARCHAR(50)       NOT NULL,                    -- 변경 후 상태
    changed_by_type actor_type        NOT NULL,                    -- user / system / admin
    changed_by_id   UUID,                                          -- 변경 주체 ID (시스템이면 NULL)
    reason_code     VARCHAR(50),                                   -- 변경 사유 코드
    memo            TEXT,                                          -- 자유 메모
    created_at      TIMESTAMPTZ       NOT NULL DEFAULT now()
);

COMMENT ON TABLE  status_history IS '상태 변경 이력. append-only. 엔티티의 모든 상태 전이를 시간순으로 기록한다.';
COMMENT ON COLUMN status_history.from_status IS '이전 상태. 최초 생성(initial insert)이면 NULL.';
COMMENT ON COLUMN status_history.reason_code IS '변경 사유 코드. 사용자 취소/시스템 만료/운영 강제 등을 구분한다.';
```

---

## 10. 인덱스 정의

```sql
-- ============================================================
-- 인덱스
-- PRD 섹션 45.2 인덱스 후보를 기반으로 정의한다.
-- ============================================================

-- === 사용자 ===
CREATE INDEX idx_users_account_status      ON users (account_status) WHERE withdrawn_at IS NULL;
CREATE INDEX idx_user_profiles_server      ON user_profiles (primary_server_id);
CREATE INDEX idx_block_relations_blocked   ON block_relations (blocked_user_id);

-- === 매물 (검색/필터 핵심 인덱스) ===
-- 메인 매물 목록 쿼리: 상태 + 노출 + 서버 + 카테고리 + 최근활동순
CREATE INDEX idx_listings_main_feed
    ON listings (status, visibility, server_id, category_id, last_activity_at DESC)
    WHERE deleted_at IS NULL;

-- 작성자별 매물 목록
CREATE INDEX idx_listings_author           ON listings (author_user_id, created_at DESC) WHERE deleted_at IS NULL;

-- 아이템명 부분 검색 (pg_trgm)
CREATE INDEX idx_listings_item_name_trgm   ON listings USING gin (item_name gin_trgm_ops) WHERE deleted_at IS NULL;

-- 제목 부분 검색 (pg_trgm)
CREATE INDEX idx_listings_title_trgm       ON listings USING gin (title gin_trgm_ops) WHERE deleted_at IS NULL;

-- 거래유형별 필터
CREATE INDEX idx_listings_type_status      ON listings (listing_type, status, last_activity_at DESC) WHERE deleted_at IS NULL;

-- 만료 대상 조회
CREATE INDEX idx_listings_expires_at       ON listings (expires_at) WHERE expires_at IS NOT NULL AND status = 'available';

-- === 매물 이미지 ===
CREATE INDEX idx_listing_images_listing    ON listing_images (listing_id, sort_order);

-- === 찜 ===
CREATE INDEX idx_favorites_user            ON favorites (user_id, created_at DESC);
CREATE INDEX idx_favorites_listing         ON favorites (listing_id);

-- === 채팅방 ===
-- 판매자 기준 채팅 목록 (최근 메시지순)
CREATE INDEX idx_chat_rooms_seller         ON chat_rooms (seller_user_id, last_message_at DESC);

-- 구매자 기준 채팅 목록 (최근 메시지순)
CREATE INDEX idx_chat_rooms_buyer          ON chat_rooms (buyer_user_id, last_message_at DESC);

-- 매물 기준 채팅방 조회
CREATE INDEX idx_chat_rooms_listing        ON chat_rooms (listing_id);

-- === 채팅 메시지 ===
-- 채팅방 내 메시지 시간순 조회 (페이지네이션)
CREATE INDEX idx_chat_messages_room_time   ON chat_messages (chat_room_id, sent_at DESC) WHERE deleted_at IS NULL;

-- === 예약 ===
-- 매물별 활성 예약 조회
CREATE INDEX idx_reservations_listing_status  ON reservations (listing_id, reservation_status);

-- 채팅방별 예약 시간순
CREATE INDEX idx_reservations_room_schedule   ON reservations (chat_room_id, scheduled_at);

-- 만료 대상 예약 조회 (시스템 배치용)
CREATE INDEX idx_reservations_pending_expiry  ON reservations (expires_at)
    WHERE reservation_status = 'proposed' AND expires_at IS NOT NULL;

-- === 거래 완료 ===
-- 매물별 활성 완료 기록 (중복 방지 체크용)
CREATE INDEX idx_trade_completions_listing    ON trade_completions (listing_id, completion_status);

-- === 후기 ===
-- 대상자별 후기 목록
CREATE INDEX idx_reviews_reviewee            ON reviews (reviewee_user_id, created_at DESC)
    WHERE visibility_status = 'visible';

-- 작성자별 후기 이력
CREATE INDEX idx_reviews_reviewer            ON reviews (reviewer_user_id, created_at DESC);

-- === 신고 ===
-- 운영자 큐: 상태 + 우선순위 + 접수순
CREATE INDEX idx_reports_queue               ON reports (report_status, priority, created_at);

-- 대상별 신고 조회
CREATE INDEX idx_reports_target              ON reports (target_type, target_id);

-- 신고자별 이력
CREATE INDEX idx_reports_reporter            ON reports (reporter_user_id, created_at DESC);

-- === 운영 조치 ===
-- 대상 사용자별 조치 이력
CREATE INDEX idx_mod_actions_target_user     ON moderation_actions (target_user_id, created_at DESC);

-- 처리 운영자별 조치 이력
CREATE INDEX idx_mod_actions_admin           ON moderation_actions (created_by_admin_id, created_at DESC);

-- 활성 조치 조회
CREATE INDEX idx_mod_actions_active          ON moderation_actions (target_user_id, action_type)
    WHERE is_active = TRUE;

-- === 알림 ===
-- 사용자별 알림 목록 (안 읽은 것 우선, 최신순)
CREATE INDEX idx_notifications_user          ON notifications (user_id, read_at NULLS FIRST, created_at DESC);

-- === 감사 로그 ===
CREATE INDEX idx_audit_logs_actor            ON audit_logs (actor_id, created_at DESC) WHERE actor_id IS NOT NULL;
CREATE INDEX idx_audit_logs_target           ON audit_logs (target_type, target_id, created_at DESC);

-- === 상태 이력 ===
CREATE INDEX idx_status_history_entity       ON status_history (entity_type, entity_id, created_at DESC);

-- === 푸시 토큰 ===
CREATE INDEX idx_push_tokens_user            ON push_tokens (user_id) WHERE is_active = TRUE;
```

---

## 11. updated_at 자동 갱신 트리거

```sql
-- ============================================================
-- updated_at 컬럼 자동 갱신 함수 및 트리거
-- updated_at이 있는 모든 테이블에 적용한다.
-- ============================================================
CREATE OR REPLACE FUNCTION trigger_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 트리거 적용 대상 테이블 목록
DO $$
DECLARE
    t TEXT;
BEGIN
    FOREACH t IN ARRAY ARRAY[
        'servers', 'categories',
        'users', 'user_profiles', 'push_tokens',
        'listings',
        'chat_rooms', 'chat_messages', 'chat_read_cursors',
        'reservations', 'trade_completions', 'reviews',
        'reports', 'moderation_actions'
    ]
    LOOP
        EXECUTE format(
            'CREATE TRIGGER trg_%s_updated_at
             BEFORE UPDATE ON %I
             FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at()',
            t, t
        );
    END LOOP;
END;
$$;
```

---

## 12. 시드 데이터

```sql
-- ============================================================
-- 게임 서버 시드 데이터
-- 리니지 클래식 서버 목록 (2026년 3월 기준, 실제 운영 서버에 맞게 조정 필요)
-- ============================================================
INSERT INTO servers (server_id, name_ko, name_en, region, sort_order) VALUES
    ('depolloju',   '데포로쥬',   'Depolloju',   'kr', 1),
    ('ken-raisa',   '켄라우헬',   'Ken Rauhel',  'kr', 2),
    ('bartz',       '바츠',       'Bartz',       'kr', 3),
    ('dion',        '디온',       'Dion',        'kr', 4),
    ('giran',       '기란',       'Giran',       'kr', 5),
    ('oren',        '오렌',       'Oren',        'kr', 6),
    ('aden',        '아덴',       'Aden',        'kr', 7),
    ('gludio',      '글루디오',   'Gludio',      'kr', 8),
    ('heine',       '하이네',     'Heine',       'kr', 9),
    ('innadril',    '인나드릴',   'Innadril',    'kr', 10)
ON CONFLICT (server_id) DO NOTHING;

-- ============================================================
-- 아이템 카테고리 시드 데이터
-- 최상위 카테고리와 하위 카테고리 구분
-- ============================================================

-- 최상위 카테고리
INSERT INTO categories (category_id, name_ko, name_en, parent_id, sort_order) VALUES
    ('weapon',       '무기',         'Weapon',       NULL, 1),
    ('armor',        '방어구',       'Armor',        NULL, 2),
    ('accessory',    '장신구',       'Accessory',    NULL, 3),
    ('consumable',   '소모품',       'Consumable',   NULL, 4),
    ('material',     '재료',         'Material',     NULL, 5),
    ('currency',     '재화',         'Currency',     NULL, 6),
    ('transform',    '변신',         'Transform',    NULL, 7),
    ('magic_doll',   '마법인형',     'Magic Doll',   NULL, 8),
    ('skill_book',   '스킬북/주문서', 'Skill Book',   NULL, 9),
    ('etc',          '기타',         'Etc',          NULL, 10)
ON CONFLICT (category_id) DO NOTHING;

-- 무기 하위 카테고리
INSERT INTO categories (category_id, name_ko, name_en, parent_id, sort_order) VALUES
    ('weapon-sword',     '한손검',     'Sword',        'weapon', 1),
    ('weapon-2h-sword',  '양손검',     '2H Sword',     'weapon', 2),
    ('weapon-dagger',    '단검',       'Dagger',       'weapon', 3),
    ('weapon-bow',       '활',         'Bow',          'weapon', 4),
    ('weapon-staff',     '스태프',     'Staff',        'weapon', 5),
    ('weapon-2h-staff',  '양손스태프', '2H Staff',     'weapon', 6),
    ('weapon-claw',      '클로',       'Claw',         'weapon', 7),
    ('weapon-dual',      '이도류',     'Dual Sword',   'weapon', 8),
    ('weapon-spear',     '창',         'Spear',        'weapon', 9),
    ('weapon-chain',     '체인소드',   'Chain Sword',  'weapon', 10)
ON CONFLICT (category_id) DO NOTHING;

-- 방어구 하위 카테고리
INSERT INTO categories (category_id, name_ko, name_en, parent_id, sort_order) VALUES
    ('armor-helmet',     '투구',       'Helmet',       'armor', 1),
    ('armor-chest',      '상의',       'Chest',        'armor', 2),
    ('armor-legs',       '하의',       'Legs',         'armor', 3),
    ('armor-gloves',     '장갑',       'Gloves',       'armor', 4),
    ('armor-boots',      '부츠',       'Boots',        'armor', 5),
    ('armor-shirt',      '셔츠',       'Shirt',        'armor', 6),
    ('armor-cloak',      '망토',       'Cloak',        'armor', 7),
    ('armor-shield',     '방패',       'Shield',       'armor', 8)
ON CONFLICT (category_id) DO NOTHING;

-- 장신구 하위 카테고리
INSERT INTO categories (category_id, name_ko, name_en, parent_id, sort_order) VALUES
    ('acc-ring',         '반지',       'Ring',         'accessory', 1),
    ('acc-earring',      '귀걸이',     'Earring',      'accessory', 2),
    ('acc-necklace',     '목걸이',     'Necklace',     'accessory', 3),
    ('acc-belt',         '벨트',       'Belt',         'accessory', 4),
    ('acc-rune',         '룬',         'Rune',         'accessory', 5)
ON CONFLICT (category_id) DO NOTHING;

-- 재화 하위 카테고리
INSERT INTO categories (category_id, name_ko, name_en, parent_id, sort_order) VALUES
    ('currency-adena',   '아데나',     'Adena',        'currency', 1),
    ('currency-coin',    '주화/코인',  'Coin',         'currency', 2)
ON CONFLICT (category_id) DO NOTHING;
```

---

## 부록: 테이블 요약

| 테이블 | 도메인 | 설명 |
|--------|--------|------|
| `servers` | 마스터 | 게임 서버 목록 |
| `categories` | 마스터 | 아이템 카테고리 (계층 구조) |
| `users` | 사용자 | 계정 인증/권한 |
| `user_profiles` | 사용자 | 공개 프로필/신뢰 지표 |
| `block_relations` | 사용자 | 사용자 간 차단 관계 |
| `push_tokens` | 사용자 | 푸시 알림 디바이스 토큰 |
| `listings` | 매물 | 거래 게시물 |
| `listing_images` | 매물 | 매물 첨부 이미지 |
| `favorites` | 매물 | 찜 (사용자-매물) |
| `chat_rooms` | 채팅 | 1:1 채팅방 |
| `chat_messages` | 채팅 | 채팅 메시지 |
| `chat_read_cursors` | 채팅 | 읽음 커서 |
| `reservations` | 거래 | 거래 예약 |
| `trade_completions` | 거래 | 거래 완료 기록 |
| `reviews` | 거래 | 후기 |
| `reports` | 신고/운영 | 신고 접수 |
| `moderation_actions` | 신고/운영 | 운영 제재 조치 |
| `notifications` | 알림 | 앱 내 알림 |
| `audit_logs` | 감사 | 감사 로그 (append-only) |
| `status_history` | 감사 | 상태 변경 이력 (append-only) |
