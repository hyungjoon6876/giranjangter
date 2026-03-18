package repository

import (
	"context"
	"database/sql"
)

// PostgresMasterRepo implements MasterRepo using PostgreSQL via database/sql.
type PostgresMasterRepo struct{ db *sql.DB }

// NewPostgresMasterRepo returns a new PostgresMasterRepo.
func NewPostgresMasterRepo(db *sql.DB) *PostgresMasterRepo { return &PostgresMasterRepo{db: db} }

func (r *PostgresMasterRepo) ListServers(ctx context.Context) ([]ServerItem, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM servers WHERE is_active = 1 ORDER BY sort_order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ServerItem
	for rows.Next() {
		var item ServerItem
		if err := rows.Scan(&item.ServerID, &item.ServerName); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresMasterRepo) ListCategories(ctx context.Context) ([]CategoryItem, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, parent_id FROM categories ORDER BY parent_id NULLS FIRST, sort_order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CategoryItem
	for rows.Next() {
		var item CategoryItem
		if err := rows.Scan(&item.CategoryID, &item.CategoryName, &item.ParentID); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresMasterRepo) SearchItems(ctx context.Context, query string, categoryID *string) ([]ItemSearchResult, error) {
	var rows *sql.Rows
	var err error

	if categoryID == nil || *categoryID == "" {
		rows, err = r.db.QueryContext(ctx,
			"SELECT id, name, category_id, icon_id FROM item_master WHERE name ILIKE $1 ORDER BY name LIMIT 20",
			"%"+query+"%")
	} else {
		rows, err = r.db.QueryContext(ctx,
			"SELECT id, name, category_id, icon_id FROM item_master WHERE name ILIKE $1 AND category_id = $2 ORDER BY name LIMIT 20",
			"%"+query+"%", *categoryID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ItemSearchResult
	for rows.Next() {
		var item ItemSearchResult
		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.IconID); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

// Ensure compile-time interface satisfaction.
var _ MasterRepo = (*PostgresMasterRepo)(nil)
