package db

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/vladjong/music_api/internal/config"
)

const (
	COMMAND_UP = "up"
	DIR        = "./migration"
)

func Migrate(cfg *config.DB) error {
	db, err := goose.OpenDBWithDriver("pgx", cfg.DSN)
	if err != nil {
		log.Fatalf("goose: failed to close DB: %v\n", err)
		return fmt.Errorf("[db.Migrate]:%v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()
	arguments := []string{}
	if err := goose.Run(COMMAND_UP, db, DIR, arguments...); err != nil {
		return fmt.Errorf("[db.Migrate]:%v", err)
	}
	return nil
}
