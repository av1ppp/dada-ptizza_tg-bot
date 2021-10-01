package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
)

/*
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	chat_id INTEGER UNIQUE,
	price REAL,
	social_network TEXT,
	target_user TEXT,
	label TEXT
*/

type Store struct {
	db *sql.DB
}

func New(conf *config.Config) (*Store, error) {
	db, err := sql.Open("sqlite3", "productsdb.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
