-- 008: Add index on favorites(listing_id) for efficient COUNT subquery
CREATE INDEX IF NOT EXISTS idx_favorites_listing ON favorites(listing_id);
