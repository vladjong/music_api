package main

import (
	"context"
	"log"

	"github.com/vladjong/music_api/internal/config"
	"github.com/vladjong/music_api/internal/db"
	"github.com/vladjong/music_api/internal/repository/postgres"
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

	postgres.New(pgx)

	// p := playlist.New()
	// for i := 0; i < 10; i++ {
	// 	name := fmt.Sprintf("song_%d", i)
	// 	s := playlist.NewSong(name, 1)
	// 	p.AddSong(s)
	// }
	// go p.Play()
	// time.Sleep(2 * time.Second)
	// p.Stop()
	// time.Sleep(2 * time.Second)
	// go p.Play()
	// p.Next()
	// p.AddSong(playlist.NewSong("test_test", 4))
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// p.Prev()
	// time.Sleep(30 * time.Second)
}
