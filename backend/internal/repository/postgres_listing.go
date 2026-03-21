package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jym/lincle/internal/domain"
)

// PostgresListingRepo implements ListingRepo using PostgreSQL via database/sql.
type PostgresListingRepo struct{ db *sql.DB }

// NewPostgresListingRepo returns a new PostgresListingRepo.
func NewPostgresListingRepo(db *sql.DB) *PostgresListingRepo { return &PostgresListingRepo{db: db} }

func (r *PostgresListingRepo) CheckImageOwnership(ctx context.Context, imageID, userID string) (bool, error) {
	var ownerID string
	err := r.db.QueryRowContext(ctx, "SELECT user_id FROM uploaded_images WHERE id = $1", imageID).Scan(&ownerID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return ownerID == userID, nil
}

func (r *PostgresListingRepo) InsertListing(ctx context.Context, p *InsertListingParams) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO listings (id, listing_type, author_user_id, server_id, category_id,
			item_name, title, description, price_type, price_amount, quantity,
			enhancement_level, options_text, trade_method,
			preferred_meeting_area_text, available_time_text,
			status, visibility, last_activity_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, 'available', 'public', $17, $18, $19)`,
		p.ID, p.ListingType, p.AuthorUserID, p.ServerID, p.CategoryID,
		p.ItemName, p.Title, p.Description, p.PriceType, p.PriceAmount, p.Quantity,
		p.Enhancement, p.OptionsText, p.TradeMethod,
		p.MeetingArea, p.TimeText,
		p.Now, p.Now, p.Now,
	)
	return err
}

func (r *PostgresListingRepo) InsertStatusHistory(ctx context.Context, p *InsertStatusHistoryParams) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO status_history (id, entity_type, entity_id, from_status, to_status, changed_by_user_id, reason_code, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		p.ID, p.EntityType, p.EntityID, p.FromStatus, p.ToStatus, p.ChangedByUser, p.ReasonCode, p.CreatedAt)
	return err
}

func (r *PostgresListingRepo) ListListings(ctx context.Context, filter ListingFilter) ([]ListingListItem, error) {
	query := `SELECT l.id, l.listing_type, l.title, l.item_name,
		l.price_type, l.price_amount, l.enhancement_level,
		l.server_id, s.name as server_name,
		l.status, l.trade_method, l.view_count, (SELECT COUNT(*) FROM favorites f WHERE f.listing_id = l.id) as favorite_count, l.chat_count,
		l.last_activity_at, l.created_at,
		p.user_id as author_id, p.nickname, p.trust_badge, p.response_badge,
		im.icon_id
		FROM listings l
		JOIN servers s ON l.server_id = s.id
		JOIN user_profiles p ON l.author_user_id = p.user_id
		LEFT JOIN item_master im ON im.name = l.item_name
		WHERE l.deleted_at IS NULL AND l.visibility = 'public'`

	args := []interface{}{}
	paramIdx := 1

	if filter.Status != "" {
		query += fmt.Sprintf(" AND l.status = $%d", paramIdx)
		args = append(args, filter.Status)
		paramIdx++
	}
	if filter.ServerID != "" {
		query += fmt.Sprintf(" AND l.server_id = $%d", paramIdx)
		args = append(args, filter.ServerID)
		paramIdx++
	}
	if filter.CategoryID != "" {
		query += fmt.Sprintf(" AND (l.category_id = $%d OR l.category_id IN (SELECT id FROM categories WHERE parent_id = $%d))", paramIdx, paramIdx+1)
		args = append(args, filter.CategoryID, filter.CategoryID)
		paramIdx += 2
	}
	if filter.ListingType != "" {
		query += fmt.Sprintf(" AND l.listing_type = $%d", paramIdx)
		args = append(args, filter.ListingType)
		paramIdx++
	}
	if filter.Query != "" {
		query += fmt.Sprintf(" AND (l.title LIKE $%d OR l.item_name LIKE $%d)", paramIdx, paramIdx+1)
		pattern := "%" + filter.Query + "%"
		args = append(args, pattern, pattern)
		paramIdx += 2
	}
	if filter.Cursor != "" {
		query += fmt.Sprintf(" AND l.created_at < $%d", paramIdx)
		args = append(args, filter.Cursor)
		paramIdx++
	}

	switch filter.Sort {
	case "price_asc":
		query += " ORDER BY l.price_amount ASC, l.created_at DESC"
	case "price_desc":
		query += " ORDER BY l.price_amount DESC, l.created_at DESC"
	case "popular":
		query += " ORDER BY favorite_count DESC, l.created_at DESC"
	default:
		query += " ORDER BY l.last_activity_at DESC"
	}

	query += fmt.Sprintf(" LIMIT %d", filter.Limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ListingListItem
	for rows.Next() {
		var item ListingListItem
		err := rows.Scan(&item.ListingID, &item.ListingType, &item.Title, &item.ItemName,
			&item.PriceType, &item.PriceAmount, &item.EnhancementLvl,
			&item.ServerID, &item.ServerName,
			&item.Status, &item.TradeMethod, &item.ViewCount, &item.FavoriteCount, &item.ChatCount,
			&item.LastActivityAt, &item.CreatedAt,
			&item.AuthorID, &item.AuthorNickname, &item.TrustBadge, &item.ResponseBadge,
			&item.IconID)
		if err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresListingRepo) GetListing(ctx context.Context, listingID string) (*ListingDetail, error) {
	var d ListingDetail
	err := r.db.QueryRowContext(ctx, `
		SELECT l.id, l.listing_type, l.title, l.item_name, l.description,
			l.price_type, l.price_amount, l.quantity, l.enhancement_level, l.options_text,
			l.server_id, s.name, l.category_id, c.name,
			l.status, l.visibility, l.trade_method,
			l.preferred_meeting_area_text, l.available_time_text,
			l.author_user_id, p.nickname, p.trust_badge, p.response_badge, p.completed_trade_count,
			l.view_count, (SELECT COUNT(*) FROM favorites f WHERE f.listing_id = l.id) as favorite_count, l.chat_count,
			l.reserved_chat_room_id, l.last_activity_at, l.created_at, l.updated_at,
			im.icon_id
		FROM listings l
		JOIN servers s ON l.server_id = s.id
		JOIN categories c ON l.category_id = c.id
		JOIN user_profiles p ON l.author_user_id = p.user_id
		LEFT JOIN item_master im ON im.name = l.item_name
		WHERE l.id = $1 AND l.deleted_at IS NULL`, listingID,
	).Scan(&d.ID, &d.ListingType, &d.Title, &d.ItemName, &d.Description,
		&d.PriceType, &d.PriceAmount, &d.Quantity, &d.Enhancement, &d.OptionsText,
		&d.ServerID, &d.ServerName, &d.CategoryID, &d.CategoryName,
		&d.Status, &d.Visibility, &d.TradeMethod,
		&d.MeetingArea, &d.TimeText,
		&d.AuthorID, &d.AuthorNickname, &d.TrustBadge, &d.ResponseBadge, &d.TradeCount,
		&d.ViewCount, &d.FavoriteCount, &d.ChatCount,
		&d.ReservedChatID, &d.LastActivityAt, &d.CreatedAt, &d.UpdatedAt,
		&d.IconID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *PostgresListingRepo) IncrementViewCount(ctx context.Context, listingID string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE listings SET view_count = view_count + 1 WHERE id = $1", listingID)
	return err
}

func (r *PostgresListingRepo) IsFavorited(ctx context.Context, userID, listingID string) (bool, error) {
	var fav bool
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM favorites WHERE user_id = $1 AND listing_id = $2", userID, listingID).Scan(&fav)
	if err != nil {
		return false, err
	}
	return fav, nil
}

func (r *PostgresListingRepo) GetListingOwnerAndStatus(ctx context.Context, listingID string) (*ListingOwnerStatus, error) {
	var o ListingOwnerStatus
	err := r.db.QueryRowContext(ctx, "SELECT author_user_id, status FROM listings WHERE id = $1 AND deleted_at IS NULL", listingID).Scan(&o.AuthorUserID, &o.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *PostgresListingRepo) UpdateListing(ctx context.Context, listingID string, fields ListingUpdateFields) error {
	setClauses := []string{}
	args := []interface{}{}
	paramIdx := 1

	type field struct {
		col string
		val interface{}
		set bool
	}
	ff := []field{
		{"title", fields.Title, fields.Title != nil},
		{"description", fields.Description, fields.Description != nil},
		{"price_type", fields.PriceType, fields.PriceType != nil},
		{"price_amount", fields.PriceAmount, fields.PriceAmount != nil},
		{"quantity", fields.Quantity, fields.Quantity != nil},
		{"enhancement_level", fields.Enhancement, fields.Enhancement != nil},
		{"options_text", fields.OptionsText, fields.OptionsText != nil},
		{"trade_method", fields.TradeMethod, fields.TradeMethod != nil},
		{"preferred_meeting_area_text", fields.MeetingArea, fields.MeetingArea != nil},
		{"available_time_text", fields.TimeText, fields.TimeText != nil},
	}
	for _, f := range ff {
		if f.set {
			setClauses = append(setClauses, fmt.Sprintf("%s = $%d", f.col, paramIdx))
			args = append(args, f.val)
			paramIdx++
		}
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE listings SET " + strings.Join(setClauses, ", ") + fmt.Sprintf(", updated_at = NOW() WHERE id = $%d", paramIdx)
	args = append(args, listingID)
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *PostgresListingRepo) UpdateListingStatus(ctx context.Context, listingID string, status domain.ListingStatus, now time.Time) error {
	_, err := r.db.ExecContext(ctx, "UPDATE listings SET status = $1, updated_at = $2, last_activity_at = $3 WHERE id = $4", status, now, now, listingID)
	return err
}

func (r *PostgresListingRepo) ListingExists(ctx context.Context, listingID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM listings WHERE id = $1 AND deleted_at IS NULL)", listingID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *PostgresListingRepo) AddFavorite(ctx context.Context, id, userID, listingID string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO favorites (id, user_id, listing_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
		id, userID, listingID)
	return err
}

func (r *PostgresListingRepo) RemoveFavorite(ctx context.Context, userID, listingID string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM favorites WHERE user_id = $1 AND listing_id = $2", userID, listingID)
	return err
}

func (r *PostgresListingRepo) ListMyListings(ctx context.Context, userID string, status *string) ([]MyListingItem, error) {
	query := `SELECT l.id, l.listing_type, l.title, l.item_name, l.price_type, l.price_amount,
		l.enhancement_level, s.name as server_name,
		l.status, l.view_count, (SELECT COUNT(*) FROM favorites f WHERE f.listing_id = l.id) as favorite_count, l.chat_count, l.created_at,
		p.user_id as author_id, p.nickname,
		im.icon_id
		FROM listings l
		JOIN servers s ON l.server_id = s.id
		JOIN user_profiles p ON l.author_user_id = p.user_id
		LEFT JOIN item_master im ON im.name = l.item_name
		WHERE l.author_user_id = $1 AND l.deleted_at IS NULL`
	args := []interface{}{userID}
	paramIdx := 2

	if status != nil {
		query += fmt.Sprintf(" AND l.status = $%d", paramIdx)
		args = append(args, *status)
		paramIdx++
	}
	query += " ORDER BY l.created_at DESC LIMIT 50"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MyListingItem
	for rows.Next() {
		var item MyListingItem
		if err := rows.Scan(&item.ListingID, &item.ListingType, &item.Title, &item.ItemName,
			&item.PriceType, &item.PriceAmount,
			&item.EnhancementLvl, &item.ServerName,
			&item.Status, &item.ViewCount, &item.FavoriteCount, &item.ChatCount, &item.CreatedAt,
			&item.AuthorID, &item.AuthorNickname,
			&item.IconID); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
