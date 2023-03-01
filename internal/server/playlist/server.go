package playlist

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/vladjong/music_api/internal/entity"
	"github.com/vladjong/music_api/internal/playlist"
	"github.com/vladjong/music_api/internal/repository"
)

const (
	TIME_SLEEP = 300
)

type Server interface {
	AddSong(w http.ResponseWriter, r *http.Request)
	Play(w http.ResponseWriter, r *http.Request)
	Stop(w http.ResponseWriter, r *http.Request)
	Next(w http.ResponseWriter, r *http.Request)
	Prev(w http.ResponseWriter, r *http.Request)
	GetSongs(w http.ResponseWriter, r *http.Request)
	GetSong(w http.ResponseWriter, r *http.Request)
	DeleteSong(w http.ResponseWriter, r *http.Request)
	UpdateSong(w http.ResponseWriter, r *http.Request)
	GetPlaySong(w http.ResponseWriter, r *http.Request)
	GetBackup(errorChan chan error)
}

type server struct {
	storage  repository.Repository
	playlist playlist.Playlist
}

func New(storage repository.Repository) *server {
	return &server{
		storage:  storage,
		playlist: playlist.New(),
	}
}

func (s *server) GetPlaySong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errors := make(chan error, 1)
	song := make(chan *playlist.Song, 1)
	go s.playlist.GetSong(errors, song)
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.Play]:%v", err), http.StatusBadRequest)
	case data := <-song:
		writeResponseJson(w, data)
	}
}

func (s *server) AddSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := &entity.Song{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		writeError(w, fmt.Errorf("[server.AddSong]:%v", err), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()
	if err := s.storage.AddSong(ctx, *data); err != nil {
		writeError(w, fmt.Errorf("[server.AddSong]:%v", err), http.StatusBadRequest)
		return
	}
	song := playlist.NewSong(data.Id, data.Name, time.Duration(data.Duration))
	s.playlist.AddSong(song)
	response := Response{Status: "Song add in playlist"}
	writeResponseJson(w, response)
}

func (s *server) GetSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()
	data, err := s.storage.GetSongs(ctx)
	if err != nil {
		writeError(w, fmt.Errorf("[server.GetSongs]:%v", err), http.StatusBadRequest)
		return
	}
	writeResponseJson(w, data)
}

func (s *server) GetSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	in := chi.URLParam(r, "id")
	id, err := strconv.Atoi(in)
	if err != nil {
		writeError(w, fmt.Errorf("[server.GetSong]:%v", err), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()
	data, err := s.storage.GetSong(ctx, int64(id))
	if err != nil {
		writeError(w, fmt.Errorf("[server.GetSong]:%v", err), http.StatusBadRequest)
		return
	}
	writeResponseJson(w, data)
}

func (s *server) DeleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	in := chi.URLParam(r, "id")
	id, err := strconv.Atoi(in)
	if err != nil {
		writeError(w, fmt.Errorf("[server.DeleteSong]:%v", err), http.StatusBadRequest)
		return
	}
	errors := make(chan error, 1)
	go s.playlist.DeleteSong(int64(id), errors)
	response := Response{Status: fmt.Sprintf("Song id=%v delete in playlist", id)}
	time.Sleep(time.Millisecond * TIME_SLEEP)
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.DeleteSong]:%v", err), http.StatusBadRequest)
	default:
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
		defer cancel()
		if err := s.storage.DeleteSong(ctx, int64(id)); err != nil {
			writeError(w, fmt.Errorf("[server.DeleteSong]:%v", err), http.StatusBadRequest)
			return
		}
		writeResponseJson(w, response)
	}
}

func (s *server) UpdateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := &entity.Song{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		writeError(w, fmt.Errorf("[server.UpdateSong]:%v", err), http.StatusBadRequest)
		return
	}
	in := playlist.NewSong(data.Id, data.Name, time.Duration(data.Duration))
	errors := make(chan error, 1)
	go s.playlist.UpdateSong(in, errors)
	time.Sleep(time.Millisecond * TIME_SLEEP)
	response := Response{Status: fmt.Sprintf("Song id=%v update in playlist", in.Id)}
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.UpdateSong]:%v", err), http.StatusBadRequest)
	default:
		writeResponseJson(w, response)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
		defer cancel()
		if err := s.storage.UpdateSong(ctx, *data); err != nil {
			writeError(w, fmt.Errorf("[server.UpdateSong]:%v", err), http.StatusBadRequest)
			return
		}
	}
}

func (s *server) Play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errors := make(chan error, 1)
	go s.playlist.Play(errors)
	response := Response{Status: "Play apply"}
	time.Sleep(time.Millisecond * TIME_SLEEP)
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.Play]:%v", err), http.StatusBadRequest)
	default:
		writeResponseJson(w, response)
	}
}

func (s *server) Stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errors := make(chan error, 1)
	go s.playlist.Stop(errors)
	response := Response{Status: "Stop apply"}
	time.Sleep(time.Millisecond * TIME_SLEEP)
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.Stop]:%v", err), http.StatusBadRequest)
	default:
		writeResponseJson(w, response)
	}
}

func (s *server) Next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errors := make(chan error, 1)
	go s.playlist.Next(errors)
	response := Response{Status: "Next apply"}
	time.Sleep(time.Millisecond * TIME_SLEEP)
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.Next]:%v", err), http.StatusBadRequest)
	default:
		writeResponseJson(w, response)
	}
}

func (s *server) Prev(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errors := make(chan error, 1)
	go s.playlist.Prev(errors)
	response := Response{Status: "Prev apply"}
	time.Sleep(time.Millisecond * TIME_SLEEP)
	select {
	case err := <-errors:
		writeError(w, fmt.Errorf("[server.Prev]:%v", err), http.StatusBadRequest)
	default:
		writeResponseJson(w, response)
	}
}
