package main

import (
	"context"
	"database/sql"
	"fmt"
	"frontend-friend/backend/database"
	"log/slog"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	_, err := database.CreateDB()

	if err != nil {
		panic(err)
	}

	db, err := database.CreateConnection()

	if err != nil {
		panic(err)
	}

	a.db = db

	err = database.RunMigrations(a.db)

	if err != nil {
		slog.Error(fmt.Sprintf("%v", err))
	}
}

func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
