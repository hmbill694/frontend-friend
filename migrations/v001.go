package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		_, err := app.DB().NewQuery(`
		CREATE TABLE IF NOT EXISTS chats (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			project_id INTEGER NOT NULL,
			message TEXT NOT NULL,
			sender TEXT NOT NULL,
			created_at INTEGER NOT NULL,
			FOREIGN KEY (project_id) REFERENCES projects(id) ON UPDATE NO ACTION ON DELETE NO ACTION
		);

		CREATE TABLE IF NOT EXISTS components (
			id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
			project_id integer NOT NULL,
			name text NOT NULL,
			definition text NOT NULL,
			created_at integer NOT NULL,
			updated_at integer NOT NULL,
			FOREIGN KEY (project_id) REFERENCES projects(id) ON UPDATE no action ON DELETE no action
		);

		CREATE TABLE IF NOT EXISTS projects (
			id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
			name text NOT NULL,
			created_at integer NOT NULL,
			updated_at integer NOT NULL
		);`).Execute()

		if err != nil {
			panic("Could not run migration up script")
		}

		return nil
	}, func(app core.App) error {
		_, err := app.DB().NewQuery(`
		DROP TABLE IF EXISTS chats;
		DROP TABLE IF EXISTS components;
		DROP TABLE IF EXISTS projects;
		`).Execute()

		if err != nil {
			panic("Could not run migration down script")
		}

		return nil
	})
}
