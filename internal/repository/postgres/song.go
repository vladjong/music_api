package postgres

import (
	"context"
	"fmt"

	"github.com/vladjong/music_api/internal/entity"
)

func (r *repository) GetSongs(ctx context.Context) ([]entity.Song, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetSongs]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	songs := []entity.Song{}

	return songs, nil
}

func (r *repository) GetSong(ctx context.Context, id int) (entity.Song, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return entity.Song{}, fmt.Errorf("[postgres.GetSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	song := entity.Song{}
	return song, nil
}

func (r *repository) AddSong(ctx context.Context, in entity.Song) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	return nil
}

func (r *repository) UpdateSong(ctx context.Context, in entity.Song) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	return nil
}

func (r *repository) DeleteSong(ctx context.Context, id int) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	return nil
}
