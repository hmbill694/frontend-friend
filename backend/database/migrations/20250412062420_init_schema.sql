-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	name text NOT NULL,
	created_at integer NOT NULL,
	updated_at integer NOT NULL
);

CREATE TABLE IF NOT EXISTS chats (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	project_id integer NOT NULL,
	message text NOT NULL,
	sender text NOT NULL,
	created_at integer NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects(id) ON UPDATE no action ON DELETE no action
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


CREATE TABLE IF NOT EXISTS generated_pages (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	name text NOT NULL,
	content text NOT NULL,
	created_at integer NOT NULL,
	updated_at integer,
	project_id integer NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects(id) ON UPDATE no action ON DELETE no action
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;

DROP TABLE IF EXISTS components;

DROP TABLE IF EXISTS generated_pages;

DROP TABLE IF EXISTS projects;
-- +goose StatementEnd
