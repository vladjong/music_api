package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
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

	data := []entity.Song{}
	query, _, err := goqu.From(PLAYLIST_TABLE).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetSongs]:%v", err)
	}
	if err := tx.SelectContext(ctx, &data, query); err != nil {
		return nil, fmt.Errorf("[postgres.GetSongs]:%v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("[postgres.GetSongs]:%v", err)
	}
	return data, nil
}

func (r *repository) GetSong(ctx context.Context, id int64) (entity.Song, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return entity.Song{}, fmt.Errorf("[postgres.GetSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()
	data := entity.Song{}
	query, _, err := goqu.From(PLAYLIST_TABLE).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return data, fmt.Errorf("[postgres.GetSong]:%v", err)
	}
	if err := tx.GetContext(ctx, &data, query); err != nil {
		return data, fmt.Errorf("[postgres.GetSong]:%v", err)
	}
	if err := tx.Commit(); err != nil {
		return data, fmt.Errorf("[postgres.GetSong]:%v", err)
	}
	return data, nil
}

func (r *repository) AddSong(ctx context.Context, in entity.Song) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query, _, err := goqu.Insert(PLAYLIST_TABLE).Rows(in).OnConflict(goqu.DoNothing()).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.AddSong]:%v", err)
	}
	return nil
}

func (r *repository) UpdateSong(ctx context.Context, in entity.Song) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.UpdateSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query, _, err := goqu.Update(PLAYLIST_TABLE).Set(in).Where(goqu.C("id").Eq(in.Id)).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.UpdateSong]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("[postgres.UpdateSong]:%v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.UpdateSong]:%v", err)
	}

	return nil
}

func (r *repository) DeleteSong(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.DeleteSong]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()
	query, _, err := goqu.Delete(PLAYLIST_TABLE).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.DeleteSong]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("[postgres.DeleteSong]:%v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.DeleteSong]:%v", err)
	}

	return nil
}
