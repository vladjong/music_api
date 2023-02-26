package playlist

import (
	"container/list"
	"fmt"
	"time"

	asynclist "github.com/vladjong/music_api/internal/async_list"
	"github.com/vladjong/music_api/internal/timer"
)

type Playlist interface {
	AddSong(in *song)
	Stop() error
	Next() error
	Prev() error
}

type playlist struct {
	data        asynclist.AsyncList
	currentSong *list.Element
	timer       *timer.Timer
}

func New() *playlist {
	return &playlist{
		data: asynclist.New(),
	}
}

func (p *playlist) AddSong(in *song) {
	p.data.PushBack(in)
}

func (p *playlist) Play() {
	if p.timer == nil {
		p.currentSong = p.data.Front()
		p.timer = timer.NewTimer(time.Second * p.getValue().Duration)
		p.timer.Start()
	} else if p.timer.State == timer.StateIdle {
		p.timer.Start()
	}
	for p.currentSong != nil {
		go func() {
			select {
			case a := <-p.timer.C:
				fmt.Println(p.getValue().Name, "ready", a)
				p.Next()
			}
		}()
	}
}

func (p *playlist) Stop() error {
	if p.timer == nil {
		return fmt.Errorf("[Playlist.Stop]:timer don't init")
	}
	fmt.Println(p.getValue().Name, "stop")
	p.timer.Pause()
	return nil
}

func (p *playlist) Next() error {
	if p.timer == nil {
		return fmt.Errorf("[Playlist.Next]:timer don't init")
	}
	p.timer.Stop()
	p.currentSong = p.data.Next(p.currentSong)
	if p.currentSong == nil {
		p.currentSong = p.data.Front()
	}
	p.timer = timer.NewTimer(time.Second * p.getValue().Duration)
	p.timer.Start()
	return nil
}

func (p *playlist) Prev() error {
	if p.timer == nil {
		return fmt.Errorf("[Playlist.Prev]:timer don't init")
	}
	p.timer.Stop()
	p.currentSong = p.data.Prev(p.currentSong)
	if p.currentSong == nil {
		p.currentSong = p.data.Back()
	}
	p.timer = timer.NewTimer(time.Second * p.currentSong.Value.(*song).Duration)
	p.timer.Start()
	return nil
}

func (p *playlist) Show() {
	for val := p.data.Front(); val != nil; val = p.data.Next(val) {
		fmt.Println(val.Value)
	}
}

func (p *playlist) getValue() *song {
	return p.data.GetValue(p.currentSong).(*song)
}
