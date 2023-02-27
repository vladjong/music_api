package postgres

import "github.com/jmoiron/sqlx"

const (
	PLAYLIST_TABLE = "playlist"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}
