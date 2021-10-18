package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func New() (*Store, error) {
	db, err := sql.Open("sqlite3", "tgbot.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,

			first_name     TEXT,
			last_name      TEXT,
			sex            TEXT CHECK(sex IN ('male', 'female')),
			social_network TEXT CHECK(social_network IN ('vk', 'insta')),
			url            TEXT UNIQUE,
			picture        BLOB,

			count_private_photos INTEGER,
			count_private_videos INTEGER,
			count_hidden_data    INTEGER,
			count_hidden_friends INTEGER
		);

		CREATE TABLE IF NOT EXISTS purchases (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,

			chat_id        INTEGER UNIQUE,
			price          REAL,
			target_user_id INTEGER,
			label          TEXT,

			FOREIGN KEY (target_user_id) REFERENCES users(id)
				ON UPDATE CASCADE
				ON DELETE CASCADE
		);`,
	)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
