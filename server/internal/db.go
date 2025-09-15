package internal

import (
	"database/sql"
	"log"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/database"
	sqlc "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/database/sqlc"
)

type DB struct {
	Conn    *sql.DB
	Queries *sqlc.Queries
}

func InitDB(databaseURL string) (*DB, error) {
	conn, err := database.Connect(databaseURL)
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(conn)

	log.Println("Database initialized successfully")

	return &DB{
		Conn:    conn,
		Queries: queries,
	}, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}
