package playlist

import (
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
)

type Song struct {
	Id       int64
	Name     string
	Duration time.Duration
}

func NewSong(id int64, name string, duration time.Duration) *Song {
	return &Song{
		Id:       id,
		Name:     name,
		Duration: duration,
	}
}
