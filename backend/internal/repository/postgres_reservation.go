package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jym/lincle/internal/alignment"
	"github.com/jym/lincle/internal/domain"
)

// PostgresReservationRepo implements ReservationRepo using PostgreSQL via database/sql.
type PostgresReservationRepo struct{ db *sql.DB }

// NewPostgresReservationRepo returns a new PostgresReservationRepo.
func NewPostgresReservationRepo(db *sql.DB) *PostgresReservationRepo {
	return &PostgresReservationRepo{db: db}
}

func (r *PostgresReservationRepo) GetChatRoomForReservation(ctx context.Context, chatRoomID, userID string) (*ChatRoomReservationInfo, error) {
	var info ChatRoomReservationInfo
	err := r.db.QueryRowContext(ctx,
		"SELECT listing_id, seller_user_id, buyer_user_id FROM chat_rooms WHERE id = $1 AND (seller_user_id = $2 OR buyer_user_id = $2)",
		chatRoomID, userID,
	).Scan(&info.ListingID, &info.SellerID, &info.BuyerID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *PostgresReservationRepo) CountActiveReservations(ctx context.Context, listingID string) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM reservations WHERE listing_id = $1 AND status IN ('proposed','confirmed')",
		listingID,
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PostgresReservationRepo) CreateReservation(ctx context.Context, params *CreateReservationParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO reservations (id, listing_id, chat_room_id, proposer_user_id, counterpart_user_id, status, scheduled_at, meeting_type, server_id, meeting_point_text, note_to_counterparty, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, 'proposed', $6, $7, $8, $9, $10, $11, $12)`,
		params.ReservationID, params.ListingID, params.ChatRoomID, params.ProposerUserID, params.CounterpartID,
		params.ScheduledAt, params.MeetingType, params.ServerID, params.MeetingPoint, params.Note, params.ExpiresAt, params.Now,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx,
		"INSERT INTO chat_messages (id, chat_room_id, message_type, body_text, metadata_json, sent_at) VALUES ($1, $2, 'system', '예약이 제안되었습니다.', $3, $4)",
		params.SystemMessageID, params.ChatRoomID, params.SystemMetaJSON, params.Now,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx,
		"UPDATE chat_rooms SET chat_status = 'reservation_proposed', updated_at = $1 WHERE id = $2",
		params.Now, params.ChatRoomID,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresReservationRepo) GetReservationForConfirm(ctx context.Context, reservationID string) (*ReservationConfirmInfo, error) {
	var info ReservationConfirmInfo
	err := r.db.QueryRowContext(ctx,
		"SELECT counterpart_user_id, listing_id, chat_room_id FROM reservations WHERE id = $1 AND status = 'proposed'",
		reservationID,
	).Scan(&info.CounterpartUserID, &info.ListingID, &info.ChatRoomID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *PostgresReservationRepo) ConfirmReservation(ctx context.Context, params *ConfirmReservationParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE reservations SET status = 'confirmed', confirmed_at = $1 WHERE id = $2", params.Now, params.ReservationID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE listings SET status = 'reserved', reserved_chat_room_id = $1, updated_at = $2, last_activity_at = $3 WHERE id = $4",
		params.ChatRoomID, params.Now, params.Now, params.ListingID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE chat_rooms SET chat_status = 'reservation_confirmed', updated_at = $1 WHERE id = $2", params.Now, params.ChatRoomID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO status_history (id, entity_type, entity_id, from_status, to_status, changed_by_user_id, created_at) VALUES ($1, 'listing', $2, 'available', 'reserved', $3, $4)`,
		params.StatusHistoryID, params.ListingID, params.UserID, params.Now); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresReservationRepo) GetReservationForCancel(ctx context.Context, reservationID string) (*ReservationCancelInfo, error) {
	var info ReservationCancelInfo
	err := r.db.QueryRowContext(ctx,
		"SELECT listing_id, chat_room_id, proposer_user_id, counterpart_user_id FROM reservations WHERE id = $1 AND status IN ('proposed','confirmed')",
		reservationID,
	).Scan(&info.ListingID, &info.ChatRoomID, &info.ProposerID, &info.CounterpartID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *PostgresReservationRepo) CancelReservation(ctx context.Context, params *CancelReservationParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE reservations SET status = 'cancelled', cancelled_at = $1, cancellation_reason_code = $2 WHERE id = $3",
		params.Now, params.ReasonCode, params.ReservationID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE listings SET status = 'available', reserved_chat_room_id = NULL, updated_at = $1, last_activity_at = $2 WHERE id = $3",
		params.Now, params.Now, params.ListingID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE chat_rooms SET chat_status = 'open', updated_at = $1 WHERE id = $2",
		params.Now, params.ChatRoomID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresReservationRepo) GetConfirmedReservation(ctx context.Context, reservationID, listingID string) (*ReservationParticipants, error) {
	var p ReservationParticipants
	err := r.db.QueryRowContext(ctx,
		"SELECT proposer_user_id, counterpart_user_id FROM reservations WHERE id = $1 AND listing_id = $2 AND status = 'confirmed'",
		reservationID, listingID,
	).Scan(&p.ProposerID, &p.CounterpartID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PostgresReservationRepo) CreateTradeCompletion(ctx context.Context, params *CreateTradeCompletionParams) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO trade_completions (id, listing_id, reservation_id, requested_by_user_id, counterpart_user_id, status, completion_note, auto_confirm_at, created_at)
		VALUES ($1, $2, $3, $4, $5, 'pending_confirm', $6, $7, $8)`,
		params.ID, params.ListingID, params.ReservationID, params.RequestedByID, params.CounterpartID,
		params.CompletionNote, params.AutoConfirmAt, params.Now,
	)
	return err
}

func (r *PostgresReservationRepo) GetPendingCompletion(ctx context.Context, completionID string) (*PendingCompletionInfo, error) {
	var info PendingCompletionInfo
	err := r.db.QueryRowContext(ctx,
		"SELECT counterpart_user_id, listing_id, reservation_id, requested_by_user_id FROM trade_completions WHERE id = $1 AND status = 'pending_confirm'",
		completionID,
	).Scan(&info.CounterpartUserID, &info.ListingID, &info.ReservationID, &info.RequestedByUserID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *PostgresReservationRepo) ConfirmCompletion(ctx context.Context, params *ConfirmCompletionParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE trade_completions SET status = 'confirmed', confirmed_at = $1 WHERE id = $2",
		params.Now, params.CompletionID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE reservations SET status = 'fulfilled' WHERE id = $1", params.ReservationID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE listings SET status = 'completed', updated_at = $1, last_activity_at = $2 WHERE id = $3",
		params.Now, params.Now, params.ListingID); err != nil {
		return err
	}

	if err := alignment.Change(tx, params.RequestedByID, domain.AlignmentTradeConfirmed, "trade_confirmed", "completion", params.CompletionID); err != nil {
		return err
	}
	if err := alignment.Change(tx, params.CounterpartID, domain.AlignmentTradeConfirmed, "trade_confirmed", "completion", params.CompletionID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresReservationRepo) ListMyTrades(ctx context.Context, userID string) ([]MyTradeItem, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT cr.id, cr.listing_id, l.title, l.status,
			CASE WHEN cr.seller_user_id = $1 THEN cr.buyer_user_id ELSE cr.seller_user_id END as cp_id,
			p.nickname, cr.chat_status, cr.updated_at
		FROM chat_rooms cr
		JOIN listings l ON cr.listing_id = l.id
		JOIN user_profiles p ON p.user_id = CASE WHEN cr.seller_user_id = $2 THEN cr.buyer_user_id ELSE cr.seller_user_id END
		WHERE (cr.seller_user_id = $3 OR cr.buyer_user_id = $4)
		ORDER BY cr.updated_at DESC LIMIT 50`, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MyTradeItem
	for rows.Next() {
		var item MyTradeItem
		if err := rows.Scan(&item.ChatRoomID, &item.ListingID, &item.ListingTitle, &item.ListingStatus,
			&item.CounterpartID, &item.CounterpartNick, &item.ChatStatus, &item.UpdatedAt); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresReservationRepo) GetCompletionForReview(ctx context.Context, completionID string) (*CompletionReviewInfo, error) {
	var info CompletionReviewInfo
	err := r.db.QueryRowContext(ctx,
		"SELECT status, requested_by_user_id, counterpart_user_id FROM trade_completions WHERE id = $1",
		completionID,
	).Scan(&info.Status, &info.RequestedByUserID, &info.CounterpartUserID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *PostgresReservationRepo) CreateReview(ctx context.Context, params *CreateReviewParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx,
		"INSERT INTO reviews (id, completion_id, reviewer_user_id, target_user_id, rating, comment) VALUES ($1, $2, $3, $4, $5, $6)",
		params.ReviewID, params.CompletionID, params.ReviewerID, params.TargetUserID, params.Rating, params.Comment,
	); err != nil {
		return err
	}

	if params.Rating == "positive" {
		if _, err := tx.ExecContext(ctx, "UPDATE user_profiles SET positive_review_count = positive_review_count + 1 WHERE user_id = $1", params.TargetUserID); err != nil {
			return err
		}
		if err := alignment.Change(tx, params.TargetUserID, domain.AlignmentPositiveReview, "positive_review_received", "review", params.ReviewID); err != nil {
			return err
		}
	} else {
		if err := alignment.Change(tx, params.TargetUserID, domain.AlignmentNegativeReview, "negative_review_received", "review", params.ReviewID); err != nil {
			return err
		}
	}

	if _, err := tx.ExecContext(ctx, "UPDATE user_profiles SET completed_trade_count = completed_trade_count + 1 WHERE user_id = $1", params.ReviewerID); err != nil {
		return err
	}

	if err := alignment.Change(tx, params.ReviewerID, domain.AlignmentReviewWritten, "review_written", "review", params.ReviewID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresReservationRepo) ListUserReviews(ctx context.Context, targetUserID string) ([]UserReviewItem, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT r.id, r.rating, r.comment, r.created_at, p.nickname
		FROM reviews r JOIN user_profiles p ON r.reviewer_user_id = p.user_id
		WHERE r.target_user_id = $1 ORDER BY r.created_at DESC LIMIT 50`, targetUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []UserReviewItem
	for rows.Next() {
		var item UserReviewItem
		if err := rows.Scan(&item.ReviewID, &item.Rating, &item.Comment, &item.CreatedAt, &item.ReviewerNickname); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresReservationRepo) CreateReport(ctx context.Context, params *CreateReportParams) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO reports (id, reporter_user_id, target_type, target_id, report_type, description, status) VALUES ($1, $2, $3, $4, $5, $6, 'submitted')",
		params.ID, params.ReporterID, params.TargetType, params.TargetID, params.ReportType, params.Description,
	)
	return err
}

func (r *PostgresReservationRepo) ListMyReports(ctx context.Context, userID string) ([]MyReportItem, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, target_type, target_id, report_type, status, created_at FROM reports WHERE reporter_user_id = $1 ORDER BY created_at DESC LIMIT 50",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MyReportItem
	for rows.Next() {
		var item MyReportItem
		if err := rows.Scan(&item.ReportID, &item.TargetType, &item.TargetID, &item.ReportType, &item.Status, &item.CreatedAt); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresReservationRepo) ListNotifications(ctx context.Context, userID string) ([]NotificationItem, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, type, title, body, reference_type, reference_id, deep_link, is_read, created_at FROM notifications WHERE user_id = $1 ORDER BY created_at DESC LIMIT 50",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []NotificationItem
	for rows.Next() {
		var item NotificationItem
		if err := rows.Scan(&item.NotificationID, &item.Type, &item.Title, &item.Body,
			&item.ReferenceType, &item.ReferenceID, &item.DeepLink,
			&item.IsRead, &item.CreatedAt); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresReservationRepo) MarkNotificationsRead(ctx context.Context, userID string, notificationIDs []string) error {
	if len(notificationIDs) == 0 {
		return nil
	}

	placeholders := make([]string, len(notificationIDs))
	args := make([]interface{}, 0, len(notificationIDs)+1)
	args = append(args, userID)
	for i, nid := range notificationIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args = append(args, nid)
	}

	_, err := r.db.ExecContext(ctx,
		"UPDATE notifications SET is_read = true WHERE user_id = $1 AND id IN ("+strings.Join(placeholders, ",")+")",
		args...)
	return err
}

// Ensure compile-time interface satisfaction.
var _ ReservationRepo = (*PostgresReservationRepo)(nil)

// Ensure compile-time interface satisfaction for other repos.
var _ AuthRepo = (*PostgresAuthRepo)(nil)
var _ ListingRepo = (*PostgresListingRepo)(nil)
var _ ChatRepo = (*PostgresChatRepo)(nil)

