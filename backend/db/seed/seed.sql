-- Lincle Seed Data: Game Servers & Item Categories

-- ============================================================
-- 리니지 클래식 서버
-- ============================================================
INSERT INTO servers (id, name, region, is_active, sort_order) VALUES
('bartz',       '바츠',       'korea', 1, 1),
('ken_rauhel',  '켄라우헬',    'korea', 1, 2),
('dep_roze',    '데포로쥬',    'korea', 1, 3),
('aden',        '아덴',       'korea', 1, 4),
('windawood',   '윈다우드',    'korea', 1, 5),
('sieghart',    '지그하르트',   'korea', 1, 6),
('gustin',      '거스틴',     'korea', 1, 7),
('lionna',      '리오나',     'korea', 1, 8),
('classic_1',   '클래식1',    'korea', 1, 9),
('classic_2',   '클래식2',    'korea', 1, 10)
ON CONFLICT DO NOTHING;

-- ============================================================
-- 아이템 카테고리 (상위)
-- ============================================================
INSERT INTO categories (id, name, parent_id, sort_order) VALUES
('weapon',      '무기',       NULL, 1),
('armor',       '방어구',     NULL, 2),
('accessory',   '장신구',     NULL, 3),
('consumable',  '소모품',     NULL, 4),
('material',    '재료',       NULL, 5),
('currency',    '재화',       NULL, 6),
('scroll',      '주문서',     NULL, 7),
('pet',         '펫/소환수',   NULL, 8),
('etc',         '기타',       NULL, 9),
('account',     '계정',       NULL, 10)
ON CONFLICT DO NOTHING;

-- ============================================================
-- 아이템 카테고리 (하위 - 무기)
-- ============================================================
INSERT INTO categories (id, name, parent_id, sort_order) VALUES
('weapon_sword',    '한손검',   'weapon', 1),
('weapon_2h_sword', '양손검',   'weapon', 2),
('weapon_dagger',   '단검',    'weapon', 3),
('weapon_bow',      '활',     'weapon', 4),
('weapon_staff',    '지팡이',  'weapon', 5),
('weapon_spear',    '창',     'weapon', 6),
('weapon_dual',     '이도류',  'weapon', 7),
('weapon_claw',     '너클',   'weapon', 8)
ON CONFLICT DO NOTHING;

-- ============================================================
-- 아이템 카테고리 (하위 - 방어구)
-- ============================================================
INSERT INTO categories (id, name, parent_id, sort_order) VALUES
('armor_helmet',  '투구',    'armor', 1),
('armor_top',     '상의',    'armor', 2),
('armor_bottom',  '하의',    'armor', 3),
('armor_gloves',  '장갑',    'armor', 4),
('armor_boots',   '신발',    'armor', 5),
('armor_shield',  '방패',    'armor', 6),
('armor_cloak',   '망토',    'armor', 7)
ON CONFLICT DO NOTHING;

-- ============================================================
-- 아이템 카테고리 (하위 - 장신구)
-- ============================================================
INSERT INTO categories (id, name, parent_id, sort_order) VALUES
('acc_ring',      '반지',    'accessory', 1),
('acc_earring',   '귀걸이',   'accessory', 2),
('acc_necklace',  '목걸이',   'accessory', 3),
('acc_belt',      '벨트',    'accessory', 4)
ON CONFLICT DO NOTHING;

-- ============================================================
-- 아이템 카테고리 (하위 - 재화)
-- ============================================================
INSERT INTO categories (id, name, parent_id, sort_order) VALUES
('currency_adena',   '아데나',  'currency', 1),
('currency_diamond', '다이아',  'currency', 2)
ON CONFLICT DO NOTHING;

-- ============================================================
-- 추가 하위 카테고리 (사이트 데이터에서 발견된 것들)
-- ============================================================
INSERT INTO categories (id, name, parent_id, sort_order) VALUES
('weapon_mace',    '둔기',    'weapon', 9),
('armor_shirt',    '티셔츠',  'armor', 8),
('consumable_potion', '물약', 'consumable', 1),
('spellbook',      '마법서',  'scroll', 1)
ON CONFLICT DO NOTHING;
