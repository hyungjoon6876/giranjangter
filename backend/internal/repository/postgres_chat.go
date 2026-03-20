package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// PostgresChatRepo implements ChatRepo using PostgreSQL via database/sql.
type PostgresChatRepo struct{ db *sql.DB }

// NewPostgresChatRepo returns a new PostgresChatRepo.
func NewPostgresChatRepo(db *sql.DB) *PostgresChatRepo { return &PostgresChatRepo{db: db} }

func (r *PostgresChatRepo) GetListingAuthor(ctx context.Context, listingID string) (string, error) {
	var authorID string
	err := r.db.QueryRowContext(ctx, "SELECT author_user_id FROM listings WHERE id = $1 AND deleted_at IS NULL", listingID).Scan(&authorID)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return authorID, nil
}

func (r *PostgresChatRepo) FindExistingChatRoom(ctx context.Context, listingID, sellerID, buyerID string) (string, error) {
	var existingID string
	err := r.db.QueryRowContext(ctx,
		"SELECT id FROM chat_rooms WHERE listing_id = $1 AND ((seller_user_id = $2 AND buyer_user_id = $3) OR (seller_user_id = $4 AND buyer_user_id = $5))",
		listingID, sellerID, buyerID, buyerID, sellerID,
	).Scan(&existingID)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return existingID, nil
}

func (r *PostgresChatRepo) CreateChatRoom(ctx context.Context, params *CreateChatRoomParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx,
		"INSERT INTO chat_rooms (id, listing_id, seller_user_id, buyer_user_id, chat_status, created_at, updated_at) VALUES ($1, $2, $3, $4, 'open', $5, $6)",
		params.ID, params.ListingID, params.SellerID, params.BuyerID, params.Now, params.Now,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, "UPDATE listings SET chat_count = chat_count + 1 WHERE id = $1", params.ListingID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresChatRepo) ListChatRooms(ctx context.Context, userID string) ([]ChatRoomListItem, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT cr.id, cr.listing_id, l.title, cr.chat_status, cr.last_message_at, cr.updated_at,
			CASE WHEN cr.seller_user_id = $1 THEN cr.buyer_user_id ELSE cr.seller_user_id END as counterpart_id,
			p.nickname, p.trust_badge,
			(SELECT li.thumbnail_url FROM listing_images li WHERE li.listing_id = l.id ORDER BY li.sort_order ASC LIMIT 1) as thumbnail_url,
			l.status,
			l.price_amount,
			l.server_id,
			s.name as server_name,
			COALESCE(
				(SELECT COUNT(*) FROM chat_messages cm
				 WHERE cm.chat_room_id = cr.id AND cm.deleted_at IS NULL
				 AND cm.sent_at > COALESCE(
					 (SELECT cm_read.sent_at FROM chat_read_cursors crc
					  JOIN chat_messages cm_read ON cm_read.id = crc.last_read_message_id
					  WHERE crc.chat_room_id = cr.id AND crc.user_id = $5), cr.created_at
				 )
				 AND cm.sender_user_id != $6
				), 0
			) as unread_count,
			(SELECT cm2.body_text FROM chat_messages cm2
			 WHERE cm2.chat_room_id = cr.id AND cm2.deleted_at IS NULL
			 ORDER BY cm2.sent_at DESC LIMIT 1) as last_message_body,
			(SELECT cm3.sent_at FROM chat_messages cm3
			 WHERE cm3.chat_room_id = cr.id AND cm3.deleted_at IS NULL
			 ORDER BY cm3.sent_at DESC LIMIT 1) as last_message_sent_at,
			(SELECT crc2.last_read_message_id FROM chat_read_cursors crc2
			 WHERE crc2.chat_room_id = cr.id AND crc2.user_id = $7) as my_last_read_msg_id
		FROM chat_rooms cr
		JOIN listings l ON cr.listing_id = l.id
		LEFT JOIN servers s ON l.server_id = s.id
		JOIN user_profiles p ON p.user_id = CASE WHEN cr.seller_user_id = $2 THEN cr.buyer_user_id ELSE cr.seller_user_id END
		WHERE cr.seller_user_id = $3 OR cr.buyer_user_id = $4
		ORDER BY COALESCE(cr.last_message_at, cr.created_at) DESC
		LIMIT 50`, userID, userID, userID, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ChatRoomListItem
	for rows.Next() {
		var item ChatRoomListItem
		if err := rows.Scan(
			&item.ChatRoomID, &item.ListingID, &item.ListingTitle, &item.ChatStatus,
			&item.LastMessageAt, &item.UpdatedAt,
			&item.CounterpartID, &item.CounterpartNick, &item.CounterpartBadge,
			&item.ListingThumbnail, &item.ListingStatus,
			&item.ListingPrice, &item.ListingServerID, &item.ListingServerName,
			&item.UnreadCount,
			&item.LastMessageBody, &item.LastMessageSentAt, &item.MyLastReadMsgID,
		); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresChatRepo) IsChatParticipant(ctx context.Context, chatRoomID, userID string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM chat_rooms WHERE id = $1 AND (seller_user_id = $2 OR buyer_user_id = $3)",
		chatRoomID, userID, userID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresChatRepo) ListMessages(ctx context.Context, chatRoomID string, limit int, cursor string) ([]ChatMessageItem, error) {
	query := "SELECT id, sender_user_id, message_type, body_text, metadata_json, sent_at FROM chat_messages WHERE chat_room_id = $1 AND deleted_at IS NULL"
	args := []interface{}{chatRoomID}
	paramIdx := 2

	if cursor != "" {
		query += fmt.Sprintf(" AND sent_at < $%d", paramIdx)
		args = append(args, cursor)
		paramIdx++
	}

	query += fmt.Sprintf(" ORDER BY sent_at DESC LIMIT %d", limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ChatMessageItem
	for rows.Next() {
		var item ChatMessageItem
		if err := rows.Scan(&item.MessageID, &item.SenderUserID, &item.MessageType, &item.BodyText, &item.MetadataJSON, &item.SentAt); err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresChatRepo) GetChatRoomParticipants(ctx context.Context, chatRoomID, userID string) (*ChatParticipants, error) {
	var p ChatParticipants
	err := r.db.QueryRowContext(ctx,
		"SELECT seller_user_id, buyer_user_id FROM chat_rooms WHERE id = $1 AND (seller_user_id = $2 OR buyer_user_id = $3)",
		chatRoomID, userID, userID,
	).Scan(&p.SellerID, &p.BuyerID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PostgresChatRepo) CheckDuplicateMessage(ctx context.Context, clientMessageID string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM chat_messages WHERE client_message_id = $1", clientMessageID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresChatRepo) InsertMessage(ctx context.Context, params *InsertMessageParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx,
		"INSERT INTO chat_messages (id, chat_room_id, sender_user_id, message_type, body_text, client_message_id, sent_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		params.ID, params.ChatRoomID, params.SenderUserID, params.MessageType, params.BodyText, params.ClientMessageID, params.Now,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx,
		"UPDATE chat_rooms SET last_message_at = $1, updated_at = $2 WHERE id = $3",
		params.Now, params.Now, params.ChatRoomID,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresChatRepo) UpsertReadCursor(ctx context.Context, chatRoomID, userID, lastReadMessageID string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO chat_read_cursors (chat_room_id, user_id, last_read_message_id, updated_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT(chat_room_id, user_id) DO UPDATE SET last_read_message_id = $4, updated_at = NOW()`,
		chatRoomID, userID, lastReadMessageID, lastReadMessageID)
	return err
}

func (r *PostgresChatRepo) BlockUser(ctx context.Context, id, blockerID, blockedID string) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO block_relations (id, blocker_user_id, blocked_user_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
		id, blockerID, blockedID)
	return err
}

func (r *PostgresChatRepo) UnblockUser(ctx context.Context, blockerID, blockedID string) error {
	_, err := r.db.ExecContext(ctx,
		"DELETE FROM block_relations WHERE blocker_user_id = $1 AND blocked_user_id = $2",
		blockerID, blockedID)
	return err
}
