-- 007: Fix server data — replace incorrect servers with official Lineage Classic server list
-- Old servers (bartz, aden, windawood, etc.) were from Lineage Remastered, not Classic
-- Deactivate old incorrect servers instead of deleting (foreign key safety)
UPDATE servers SET is_active = 0 WHERE id IN ('bartz', 'aden', 'windawood', 'sieghart', 'gustin', 'lionna', 'classic_1', 'classic_2');
