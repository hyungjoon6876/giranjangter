package repository

import (
	"context"
	"database/sql"
)

// PostgresUploadRepo implements UploadRepo using PostgreSQL via database/sql.
type PostgresUploadRepo struct{ db *sql.DB }

// NewPostgresUploadRepo returns a new PostgresUploadRepo.
func NewPostgresUploadRepo(db *sql.DB) *PostgresUploadRepo { return &PostgresUploadRepo{db: db} }

func (r *PostgresUploadRepo) InsertImage(ctx context.Context, params *InsertImageParams) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO uploaded_images (id, user_id, filename, url, content_type, size_bytes) VALUES ($1, $2, $3, $4, $5, $6)",
		params.ID, params.UserID, params.Filename, params.URL, params.ContentType, params.SizeBytes,
	)
	return err
}

// Ensure compile-time interface satisfaction.
var _ UploadRepo = (*PostgresUploadRepo)(nil)
