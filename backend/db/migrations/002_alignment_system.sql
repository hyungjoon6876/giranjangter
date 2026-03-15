-- 002: Alignment (성향치) System + Trade Completion Flow Changes
--
-- 성향치: 리니지 클래식의 성향 시스템에서 착안
-- 거래 완료 시 양측 확인 필수, 미응답 시 expired 처리 (자동확정 제거)

-- Add alignment_score to user_profiles (idempotent)
DO $$ BEGIN
  ALTER TABLE user_profiles ADD COLUMN alignment_score INTEGER NOT NULL DEFAULT 0;
EXCEPTION WHEN duplicate_column THEN NULL;
END $$;

DO $$ BEGIN
  ALTER TABLE user_profiles ADD COLUMN alignment_grade TEXT NOT NULL DEFAULT 'neutral';
EXCEPTION WHEN duplicate_column THEN NULL;
END $$;

-- Track alignment changes for auditing
CREATE TABLE IF NOT EXISTS alignment_history (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    delta INTEGER NOT NULL,
    reason TEXT NOT NULL,
    reference_type TEXT,
    reference_id TEXT,
    score_after INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_alignment_history_user ON alignment_history(user_id, created_at DESC);
