package playlist

import (
	"context"
	"fmt"
	"time"

	"github.com/vladjong/music_api/internal/playlist"
)

func (s *server) GetBackup(errorChan chan error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()
	data, err := s.storage.GetSongs(ctx)
	if err != nil {
		errorChan <- fmt.Errorf("[server.GetSongs]:%v", err)
		return
	}
	for _, v := range data {
		song := playlist.NewSong(v.Id, v.Name, time.Duration(v.Duration))
		s.playlist.AddSong(song)
	}
}
