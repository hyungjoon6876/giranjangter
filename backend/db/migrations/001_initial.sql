-- Lincle Initial Migration (PostgreSQL)
-- Based on docs/STARTER_DDL.md

-- ============================================================
-- Master Data
-- ============================================================
CREATE TABLE IF NOT EXISTS servers (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    region TEXT,
    is_active INTEGER NOT NULL DEFAULT 1,
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    parent_id TEXT REFERENCES categories(id),
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- Users
-- ============================================================
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    login_provider TEXT NOT NULL DEFAULT 'kakao',
    login_provider_user_key TEXT NOT NULL,
    account_status TEXT NOT NULL DEFAULT 'active',
    role TEXT NOT NULL DEFAULT 'user',
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(login_provider, login_provider_user_key)
);

CREATE TABLE IF NOT EXISTS user_profiles (
    user_id TEXT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    nickname TEXT NOT NULL,
    avatar_url TEXT,
    introduction TEXT,
    primary_server_id TEXT REFERENCES servers(id),
    response_badge TEXT NOT NULL DEFAULT 'normal',
    trust_badge TEXT NOT NULL DEFAULT 'newcomer',
    completed_trade_count INTEGER NOT NULL DEFAULT 0,
    positive_review_count INTEGER NOT NULL DEFAULT 0,
    last_active_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS block_relations (
    id TEXT PRIMARY KEY,
    blocker_user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    blocked_user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(blocker_user_id, blocked_user_id)
);

CREATE TABLE IF NOT EXISTS push_tokens (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    platform TEXT NOT NULL DEFAULT 'ios',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, token)
);

-- ============================================================
-- Listings
-- ============================================================
CREATE TABLE IF NOT EXISTS listings (
    id TEXT PRIMARY KEY,
    listing_type TEXT NOT NULL DEFAULT 'sell',
    author_user_id TEXT NOT NULL REFERENCES users(id),
    server_id TEXT NOT NULL REFERENCES servers(id),
    category_id TEXT NOT NULL REFERENCES categories(id),
    item_name TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    price_type TEXT NOT NULL DEFAULT 'fixed',
    price_amount INTEGER,
    quantity INTEGER NOT NULL DEFAULT 1,
    enhancement_level INTEGER,
    options_text TEXT,
    trade_method TEXT NOT NULL DEFAULT 'either',
    preferred_meeting_area_text TEXT,
    available_time_text TEXT,
    status TEXT NOT NULL DEFAULT 'available',
    visibility TEXT NOT NULL DEFAULT 'public',
    reserved_chat_room_id TEXT,
    view_count INTEGER NOT NULL DEFAULT 0,
    favorite_count INTEGER NOT NULL DEFAULT 0,
    chat_count INTEGER NOT NULL DEFAULT 0,
    last_activity_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_listings_status_activity ON listings(status, last_activity_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_listings_server ON listings(server_id, status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_listings_author ON listings(author_user_id, status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_listings_category ON listings(category_id, status) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS listing_images (
    id TEXT PRIMARY KEY,
    listing_id TEXT NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    thumbnail_url TEXT,
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS favorites (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    listing_id TEXT NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, listing_id)
);

-- ============================================================
-- Chat
-- ============================================================
CREATE TABLE IF NOT EXISTS chat_rooms (
    id TEXT PRIMARY KEY,
    listing_id TEXT NOT NULL REFERENCES listings(id),
    seller_user_id TEXT NOT NULL REFERENCES users(id),
    buyer_user_id TEXT NOT NULL REFERENCES users(id),
    chat_status TEXT NOT NULL DEFAULT 'open',
    last_message_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(listing_id, seller_user_id, buyer_user_id)
);

CREATE INDEX IF NOT EXISTS idx_chat_rooms_seller ON chat_rooms(seller_user_id, updated_at DESC);
CREATE INDEX IF NOT EXISTS idx_chat_rooms_buyer ON chat_rooms(buyer_user_id, updated_at DESC);

CREATE TABLE IF NOT EXISTS chat_messages (
    id TEXT PRIMARY KEY,
    chat_room_id TEXT NOT NULL REFERENCES chat_rooms(id),
    sender_user_id TEXT,
    message_type TEXT NOT NULL DEFAULT 'text',
    body_text TEXT,
    metadata_json TEXT,
    client_message_id TEXT,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_chat_messages_room ON chat_messages(chat_room_id, sent_at DESC) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS chat_read_cursors (
    chat_room_id TEXT NOT NULL REFERENCES chat_rooms(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    last_read_message_id TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (chat_room_id, user_id)
);

-- ============================================================
-- Reservations
-- ============================================================
CREATE TABLE IF NOT EXISTS reservations (
    id TEXT PRIMARY KEY,
    listing_id TEXT NOT NULL REFERENCES listings(id),
    chat_room_id TEXT NOT NULL REFERENCES chat_rooms(id),
    proposer_user_id TEXT NOT NULL REFERENCES users(id),
    counterpart_user_id TEXT NOT NULL REFERENCES users(id),
    status TEXT NOT NULL DEFAULT 'proposed',
    scheduled_at TIMESTAMPTZ NOT NULL,
    meeting_type TEXT NOT NULL DEFAULT 'either',
    server_id TEXT REFERENCES servers(id),
    meeting_point_text TEXT,
    character_name_seller TEXT,
    character_name_buyer TEXT,
    note_to_counterparty TEXT,
    expires_at TIMESTAMPTZ,
    confirmed_at TIMESTAMPTZ,
    cancelled_at TIMESTAMPTZ,
    cancellation_reason_code TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_reservations_listing ON reservations(listing_id, status);

-- ============================================================
-- Trade Completion
-- ============================================================
CREATE TABLE IF NOT EXISTS trade_completions (
    id TEXT PRIMARY KEY,
    listing_id TEXT NOT NULL REFERENCES listings(id),
    reservation_id TEXT NOT NULL REFERENCES reservations(id),
    requested_by_user_id TEXT NOT NULL REFERENCES users(id),
    counterpart_user_id TEXT NOT NULL REFERENCES users(id),
    status TEXT NOT NULL DEFAULT 'pending_confirm',
    completion_note TEXT,
    auto_confirm_at TIMESTAMPTZ,
    confirmed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- Reviews
-- ============================================================
CREATE TABLE IF NOT EXISTS reviews (
    id TEXT PRIMARY KEY,
    completion_id TEXT NOT NULL REFERENCES trade_completions(id),
    reviewer_user_id TEXT NOT NULL REFERENCES users(id),
    target_user_id TEXT NOT NULL REFERENCES users(id),
    rating TEXT NOT NULL,
    comment TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(completion_id, reviewer_user_id)
);

-- ============================================================
-- Reports
-- ============================================================
CREATE TABLE IF NOT EXISTS reports (
    id TEXT PRIMARY KEY,
    reporter_user_id TEXT NOT NULL REFERENCES users(id),
    target_type TEXT NOT NULL,
    target_id TEXT NOT NULL,
    report_type TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'submitted',
    assigned_to_user_id TEXT REFERENCES users(id),
    resolution_note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS moderation_actions (
    id TEXT PRIMARY KEY,
    report_id TEXT REFERENCES reports(id),
    actor_user_id TEXT NOT NULL REFERENCES users(id),
    target_user_id TEXT NOT NULL REFERENCES users(id),
    action_code TEXT NOT NULL,
    restriction_scope TEXT,
    duration_days INTEGER,
    memo TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- Notifications
-- ============================================================
CREATE TABLE IF NOT EXISTS notifications (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL DEFAULT '',
    reference_type TEXT,
    reference_id TEXT,
    deep_link TEXT,
    is_read INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_notifications_user ON notifications(user_id, is_read, created_at DESC);

-- ============================================================
-- Audit / History
-- ============================================================
CREATE TABLE IF NOT EXISTS audit_logs (
    id TEXT PRIMARY KEY,
    actor_id TEXT,
    actor_role TEXT,
    action TEXT NOT NULL,
    target_type TEXT NOT NULL,
    target_id TEXT NOT NULL,
    details_json TEXT,
    ip_address TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS status_history (
    id TEXT PRIMARY KEY,
    entity_type TEXT NOT NULL,
    entity_id TEXT NOT NULL,
    from_status TEXT,
    to_status TEXT NOT NULL,
    changed_by_user_id TEXT,
    reason_code TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_status_history_entity ON status_history(entity_type, entity_id, created_at DESC);
