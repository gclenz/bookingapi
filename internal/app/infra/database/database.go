package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetDatabaseConnection() *sql.DB {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Unable to connect to databse", "error", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Unable to ping databse", "error", err)
		os.Exit(1)
	}

	return db
}
