package repository

import (
	"context"

	"github.com/vladjong/music_api/internal/entity"
)

type Repository interface {
	GetSongs(ctx context.Context) ([]entity.Song, error)
	GetSong(ctx context.Context, id int64) (entity.Song, error)
	AddSong(ctx context.Context, in entity.Song) error
	UpdateSong(ctx context.Context, in entity.Song) error
	DeleteSong(ctx context.Context, id int64) error
}
