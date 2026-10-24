package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

// InitDB initialise la connexion SQLite en mode WAL et crée le schéma
func InitDB(dbPath string) (*DB, error) {
	dsn := fmt.Sprintf("%s?_journal_mode=WAL&_busy_timeout=5000", dbPath)
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir la BDD: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("impossible de pinger la BDD: %w", err)
	}

	database := &DB{Conn: conn}
	if err := database.createTables(); err != nil {
		return nil, err
	}

	return database, nil
}

func (d *DB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending',
		priority INTEGER DEFAULT 1,
		estimated_duration INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS task_dependencies (
		parent_id INTEGER NOT NULL,
		child_id INTEGER NOT NULL,
		PRIMARY KEY (parent_id, child_id),
		FOREIGN KEY (parent_id) REFERENCES tasks(id) ON DELETE CASCADE,
		FOREIGN KEY (child_id) REFERENCES tasks(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS focus_sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER,
		duration_minutes INTEGER NOT NULL,
		completed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE SET NULL
	);
	`
	_, err := d.Conn.Exec(query)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du schéma: %w", err)
	}
	return nil
}
