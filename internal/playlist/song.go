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

type song struct {
	Name     string
	Duration time.Duration
}

func NewSong(name string, duration time.Duration) *song {
	return &song{
		Name:     name,
		Duration: duration,
	}
}
