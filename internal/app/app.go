package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vladjong/music_api/internal/config"
	"github.com/vladjong/music_api/internal/server/playlist"
)

type App struct {
	handler playlist.Server
	cfg     config.Config
}

func New(handler playlist.Server, cfg config.Config) *App {
	return &App{
		handler: handler,
		cfg:     cfg,
	}
}

func (a *App) Start() {
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0%s", a.cfg.Listen.Port),
		Handler: a.initHandler(),
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Println("context deadline")
			}
		}()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf(err.Error())
		}
		serverStopCtx()
	}()

	log.Printf("start service at %s", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}

	<-serverCtx.Done()

}

func (a *App) initHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	if err := a.getBackup(); err != nil {
		log.Fatal(err)
	}

	r.Route("/api/v1/playlist", func(router chi.Router) {
		router.Post("/", a.handler.AddSong)
		router.Get("/play_song", a.handler.GetPlaySong)
		router.Get("/play", a.handler.Play)
		router.Get("/stop", a.handler.Stop)
		router.Get("/next", a.handler.Next)
		router.Get("/prev", a.handler.Prev)
		router.Get("/song", a.handler.GetSongs)
		router.Get("/song/{id}", a.handler.GetSong)
		router.Delete("/song/{id}", a.handler.DeleteSong)
		router.Put("/song", a.handler.UpdateSong)
	})
	return r
}

func (a *App) getBackup() error {
	errors := make(chan error, 1)
	go a.handler.GetBackup(errors)
	select {
	case err := <-errors:
		return fmt.Errorf("[App.getBackup]%v", err)
	default:
		return nil
	}
}
