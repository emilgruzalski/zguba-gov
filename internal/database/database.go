package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	db.SetMaxOpenConns(1)

	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, fmt.Errorf("set WAL mode: %w", err)
	}
	if _, err := db.Exec("PRAGMA foreign_keys=ON"); err != nil {
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS found_items (
			id              TEXT PRIMARY KEY,
			municipality_name  TEXT NOT NULL,
			municipality_type  TEXT NOT NULL,
			municipality_email TEXT NOT NULL,
			item_name       TEXT NOT NULL,
			item_category   TEXT NOT NULL,
			item_date       TEXT NOT NULL,
			item_location   TEXT NOT NULL,
			item_status     TEXT NOT NULL DEFAULT 'available',
			item_description TEXT,
			pickup_deadline INTEGER NOT NULL,
			pickup_location TEXT NOT NULL,
			pickup_hours    TEXT,
			pickup_contact  TEXT,
			categories      TEXT,
			created_at      DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at      DATETIME NOT NULL DEFAULT (datetime('now'))
		);

		CREATE INDEX IF NOT EXISTS idx_found_items_municipality ON found_items(municipality_name);
		CREATE INDEX IF NOT EXISTS idx_found_items_category ON found_items(item_category);
		CREATE INDEX IF NOT EXISTS idx_found_items_name ON found_items(item_name);
		CREATE INDEX IF NOT EXISTS idx_found_items_created ON found_items(created_at);
	`)
	return err
}
