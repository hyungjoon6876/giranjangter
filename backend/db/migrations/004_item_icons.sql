-- 004: Add icon_id to item_master for item icon images
DO $$ BEGIN
  ALTER TABLE item_master ADD COLUMN icon_id TEXT;
EXCEPTION WHEN duplicate_column THEN NULL;
END $$;
