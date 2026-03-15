package alignment

import (
	"database/sql"

	"github.com/google/uuid"
)

// Execer abstracts *sql.DB and *sql.Tx for shared use.
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Change adjusts a user's alignment score atomically and records the change in history.
func Change(ex Execer, userID string, delta int, reason, refType, refID string) error {
	// Single UPDATE with inline score calculation — no SELECT needed
	if _, err := ex.Exec(`
		UPDATE user_profiles SET
			alignment_score = alignment_score + $1,
			alignment_grade = CASE
				WHEN alignment_score + $2 >= 100 THEN 'royal_knight'
				WHEN alignment_score + $3 >= 50  THEN 'lawful'
				WHEN alignment_score + $4 >= 0   THEN 'neutral'
				WHEN alignment_score + $5 >= -30 THEN 'caution'
				ELSE 'chaotic'
			END,
			trust_badge = CASE
				WHEN alignment_score + $6 >= 100 THEN 'royal_knight'
				WHEN alignment_score + $7 >= 50  THEN 'lawful'
				WHEN alignment_score + $8 >= 0   THEN 'neutral'
				WHEN alignment_score + $9 >= -30 THEN 'caution'
				ELSE 'chaotic'
			END,
			updated_at = NOW()
		WHERE user_id = $10`,
		delta, delta, delta, delta, delta, delta, delta, delta, delta, userID,
	); err != nil {
		return err
	}

	// Read back the new score for history record
	var newScore int
	if err := ex.QueryRow("SELECT alignment_score FROM user_profiles WHERE user_id = $1", userID).Scan(&newScore); err != nil {
		return err
	}

	if _, err := ex.Exec(
		"INSERT INTO alignment_history (id, user_id, delta, reason, reference_type, reference_id, score_after) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		uuid.New().String(), userID, delta, reason, refType, refID, newScore,
	); err != nil {
		return err
	}

	return nil
}
