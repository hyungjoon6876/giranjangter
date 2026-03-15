-- 006: Track uploaded image ownership
CREATE TABLE IF NOT EXISTS uploaded_images (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    filename TEXT NOT NULL,
    url TEXT NOT NULL,
    content_type TEXT,
    size_bytes BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_uploaded_images_user ON uploaded_images(user_id);
