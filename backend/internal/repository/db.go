package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// InitDB opens the PostgreSQL database and runs migrations.
func InitDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("migrations: %w", err)
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	paths := []string{"db/migrations/001_initial.sql", "db/migrations/002_alignment_system.sql", "db/migrations/003_item_master.sql", "db/migrations/004_item_icons.sql", "db/migrations/005_refresh_tokens.sql", "db/migrations/006_image_ownership.sql", "db/migrations/007_fix_servers.sql", "db/migrations/008_favorites_index.sql"}
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("read %s: %w", p, err)
		}
		if _, err := db.Exec(string(data)); err != nil {
			return fmt.Errorf("exec %s: %w", p, err)
		}
		log.Printf("migration applied: %s", p)
	}
	return nil
}

// SeedDB inserts seed data.
func SeedDB(db *sql.DB) error {
	paths := []string{"db/seed/seed.sql", "db/seed/items.sql", "db/seed/item_icons.sql"}
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("read seed %s: %w", p, err)
		}
		if _, err := db.Exec(string(data)); err != nil {
			return fmt.Errorf("exec seed %s: %w", p, err)
		}
		log.Printf("seed applied: %s", p)
	}
	return nil
}
