CREATE TABLE IF NOT EXISTS item_master (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    category_id TEXT NOT NULL REFERENCES categories(id),
    sub_category TEXT NOT NULL,
    material TEXT,
    weight TEXT,
    option_text TEXT,
    is_enchantable INTEGER NOT NULL DEFAULT 0,
    safe_enchant_level INTEGER NOT NULL DEFAULT 0,
    max_enchant_level INTEGER NOT NULL DEFAULT 0,
    is_tradeable INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_item_master_category ON item_master(category_id);
CREATE INDEX IF NOT EXISTS idx_item_master_name ON item_master(name);
