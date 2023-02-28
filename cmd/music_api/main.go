package main

import (
	"context"
	"log"

	"github.com/vladjong/music_api/internal/app"
	"github.com/vladjong/music_api/internal/config"
	"github.com/vladjong/music_api/internal/db"
	"github.com/vladjong/music_api/internal/repository/postgres"
	"github.com/vladjong/music_api/internal/server/playlist"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init config")

	pgx, err := db.NewPgx(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init pgx driver")

	if err := db.Migrate(cfg.DB); err != nil {
		log.Fatal(err)
	}
	log.Println("completed migrate")

	rep := postgres.New(pgx)

	playlist := playlist.New(rep)

	app := app.New(playlist, *cfg)
	app.Start()
}
