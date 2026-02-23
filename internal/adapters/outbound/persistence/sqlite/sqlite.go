package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Enable WAL mode for better concurrency (allows concurrent reads and writes)
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, err
	}

	// Set busy timeout to 5 seconds to handle lock contention
	if _, err := db.Exec("PRAGMA busy_timeout = 5000"); err != nil {
		return nil, err
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	dbInstance := &SQLiteDB{db: db}
	if err := dbInstance.init(); err != nil {
		return nil, err
	}

	return dbInstance, nil
}

func (r *SQLiteDB) init() error {
	query := `
  CREATE TABLE IF NOT EXISTS templates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    quantity REAL,
    unit TEXT,
    difficulty INTEGER
  );

  CREATE TABLE IF NOT EXISTS template_resources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    template_id INTEGER NOT NULL,
    resource_id INTEGER NOT NULL,
    quantity REAL,
    unit TEXT,
    FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE,
    FOREIGN KEY (resource_id) REFERENCES resources(id) ON DELETE CASCADE
  );

  CREATE TABLE IF NOT EXISTS steps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    template_id INTEGER NOT NULL,
    step_order INTEGER NOT NULL,
    name TEXT NOT NULL,
    instructions TEXT,
    FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE
  );

  CREATE TABLE IF NOT EXISTS resources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    price INTEGER,
    base_quantity REAL,
    base_unit TEXT
  );

  CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    price INTEGER,
    quantity REAL,
    unit TEXT
  );
  `
	_, err := r.db.Exec(query)
	return err
}
