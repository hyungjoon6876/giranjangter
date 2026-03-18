package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// PostgresAuthRepo implements AuthRepo using PostgreSQL via database/sql.
type PostgresAuthRepo struct{ db *sql.DB }

// NewPostgresAuthRepo returns a new PostgresAuthRepo.
func NewPostgresAuthRepo(db *sql.DB) *PostgresAuthRepo { return &PostgresAuthRepo{db: db} }

func (r *PostgresAuthRepo) FindUserByProvider(ctx context.Context, provider, providerKey string) (*UserWithNickname, error) {
	var u UserWithNickname
	err := r.db.QueryRowContext(ctx,
		"SELECT u.id, u.role, COALESCE(p.nickname, '') FROM users u LEFT JOIN user_profiles p ON u.id = p.user_id WHERE u.login_provider = $1 AND u.login_provider_user_key = $2",
		provider, providerKey,
	).Scan(&u.UserID, &u.Role, &u.Nickname)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PostgresAuthRepo) CreateUserWithProfile(ctx context.Context, userID, provider, providerKey, nickname string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.ExecContext(ctx,
		"INSERT INTO users (id, login_provider, login_provider_user_key, account_status, role) VALUES ($1, $2, $3, 'active', 'user')",
		userID, provider, providerKey,
	); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx,
		"INSERT INTO user_profiles (user_id, nickname) VALUES ($1, $2)",
		userID, nickname,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresAuthRepo) UpdateLastLogin(ctx context.Context, userID string, at time.Time) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET last_login_at = $1 WHERE id = $2", at, userID)
	return err
}

func (r *PostgresAuthRepo) StoreRefreshToken(ctx context.Context, id, userID, tokenHash string, expiresAt time.Time) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, $4)",
		id, userID, tokenHash, expiresAt,
	)
	return err
}

func (r *PostgresAuthRepo) FindRefreshToken(ctx context.Context, tokenHash string) (string, error) {
	var tokenID string
	err := r.db.QueryRowContext(ctx,
		"SELECT id FROM refresh_tokens WHERE token_hash = $1 AND expires_at > NOW()",
		tokenHash,
	).Scan(&tokenID)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return tokenID, nil
}

func (r *PostgresAuthRepo) GetAccountStatus(ctx context.Context, userID string) (string, error) {
	var status string
	err := r.db.QueryRowContext(ctx, "SELECT account_status FROM users WHERE id = $1", userID).Scan(&status)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return status, nil
}

func (r *PostgresAuthRepo) RotateRefreshToken(ctx context.Context, oldTokenID, newTokenID, userID, newTokenHash string, expiresAt time.Time) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE id = $1", oldTokenID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx,
		"INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, $4)",
		newTokenID, userID, newTokenHash, expiresAt,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresAuthRepo) DeleteRefreshTokensByUser(ctx context.Context, userID string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE user_id = $1", userID)
	return err
}

func (r *PostgresAuthRepo) GetUserProfile(ctx context.Context, userID string) (*FullUserProfile, error) {
	var p FullUserProfile
	err := r.db.QueryRowContext(ctx, `
		SELECT u.id, u.role, u.account_status,
			p.nickname, p.avatar_url, p.introduction, p.primary_server_id,
			p.completed_trade_count, p.positive_review_count, p.response_badge, p.trust_badge,
			p.alignment_score, p.alignment_grade
		FROM users u JOIN user_profiles p ON u.id = p.user_id
		WHERE u.id = $1`, userID,
	).Scan(&p.UserID, &p.Role, &p.AccountStatus,
		&p.Nickname, &p.AvatarURL, &p.Introduction, &p.PrimaryServerID,
		&p.TradeCount, &p.ReviewCount, &p.ResponseBadge, &p.TrustBadge,
		&p.AlignmentScore, &p.AlignmentGrade)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PostgresAuthRepo) UpdateProfile(ctx context.Context, userID string, fields ProfileUpdateFields) error {
	setClauses := []string{}
	args := []interface{}{}
	paramIdx := 1

	if fields.Nickname != nil {
		setClauses = append(setClauses, fmt.Sprintf("nickname = $%d", paramIdx))
		args = append(args, *fields.Nickname)
		paramIdx++
	}
	if fields.Introduction != nil {
		setClauses = append(setClauses, fmt.Sprintf("introduction = $%d", paramIdx))
		args = append(args, *fields.Introduction)
		paramIdx++
	}
	if fields.PrimaryServer != nil {
		setClauses = append(setClauses, fmt.Sprintf("primary_server_id = $%d", paramIdx))
		args = append(args, *fields.PrimaryServer)
		paramIdx++
	}
	if fields.AvatarURL != nil {
		setClauses = append(setClauses, fmt.Sprintf("avatar_url = $%d", paramIdx))
		args = append(args, *fields.AvatarURL)
		paramIdx++
	}

	if len(setClauses) == 0 {
		return nil
	}

	args = append(args, userID)
	query := "UPDATE user_profiles SET " + strings.Join(setClauses, ", ") + fmt.Sprintf(", updated_at = NOW() WHERE user_id = $%d", paramIdx)
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}
