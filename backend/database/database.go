package database

import (
	"database/sql"
	"embed"
	"log/slog"
	"os"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

const (
	DIRECTORY     = "database"
	DATABASE_NAME = "frontend-friend.db"
	FULL_PATH     = DIRECTORY + "/" + DATABASE_NAME
)

func checkPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // Path exists
	}
	if os.IsNotExist(err) {
		return false, nil // Path does not exist
	}
	return false, err // Error occurred
}

func CreateDB() (string, error) {
	pathExists, err := checkPathExists(FULL_PATH)

	if err != nil {
		return "", err
	}

	if !pathExists {
		err := os.Mkdir(DIRECTORY, 0755)

		if err != nil {
			return "", err

		}

		file, err := os.Create(FULL_PATH)

		if err != nil {
			return "", err
		}

		defer file.Close()

		slog.Info("Successfully created DB file.")
	} else {
		slog.Info("Skipping database creation.")
	}

	return FULL_PATH, nil
}

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("libsql", "file:"+FULL_PATH)
	if err != nil {
		slog.Error("Failed to create connection to database.")
		return nil, err
	}

	slog.Info("Successfully connect to database.")

	return db, nil
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(db *sql.DB) error {
	slog.Info("Running Migrations...")

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
