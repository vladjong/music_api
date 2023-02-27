package playlist

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	asynclist "github.com/vladjong/music_api/internal/async_list"
	"github.com/vladjong/music_api/internal/timer"
)

type Playlist interface {
	Play(errorsChan chan error)
	AddSong(in *Song)
	Stop(errorsChan chan error)
	Next(errorsChan chan error)
	Prev(errorsChan chan error)
	DeleteSong(id int64, errorsChan chan error)
	UpdateSong(in *Song, errorsChan chan error)
	GetSong(errorsChan chan error, resultChan chan *Song)
}

type playlist struct {
	data        asynclist.AsyncList
	currentSong *list.Element
	timer       *timer.Timer
	mu          sync.Mutex
}

func New() *playlist {
	return &playlist{
		data: asynclist.New(),
	}
}

func (p *playlist) GetSong(errorsChan chan error, resultChan chan *Song) {
	if p.currentSong == nil {
		errorsChan <- fmt.Errorf("[Playlist.GetSong]:playlist don't play")
		return
	}
	resultChan <- p.getValue()
}

func (p *playlist) DeleteSong(id int64, errorsChan chan error) {
	for i := p.data.Front(); i != nil; p.Next(errorsChan) {
		if id == p.data.GetValue(i).(*Song).Id {
			p.data.Remove(i)
		}
	}
	errorsChan <- fmt.Errorf("[Playlist.DeleteSong]:don't exist element id=%v", id)
}

func (p *playlist) UpdateSong(in *Song, errorsChan chan error) {
	for i := p.data.Front(); i != nil; p.Next(errorsChan) {
		if in.Id == p.data.GetValue(i).(*Song).Id {
			p.update(i, in)
		}
	}
	errorsChan <- fmt.Errorf("[Playlist.UpdateSong]:don't exist element id=%v", in.Id)
}

func (p *playlist) AddSong(in *Song) {
	p.data.PushBack(in)
}

func (p *playlist) Play(errorsChan chan error) {
	if p.data.Len() == 0 {
		errorsChan <- fmt.Errorf("[Playlist.Play]:empty playlist")
		return
	}
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
			case <-p.timer.C:
				p.Next(errorsChan)
			}
		}()
	}
}

func (p *playlist) Stop(errorsChan chan error) {
	if p.timer == nil {
		errorsChan <- fmt.Errorf("[Playlist.Stop]:timer don't init")
		return
	}
	p.timer.Pause()
}

func (p *playlist) Next(errorsChan chan error) {
	if p.timer == nil {
		errorsChan <- fmt.Errorf("[Playlist.Next]:timer don't init")
		return
	}
	p.timer.Stop()
	p.currentSong = p.data.Next(p.currentSong)
	if p.currentSong == nil {
		p.currentSong = p.data.Front()
	}
	p.timer = timer.NewTimer(time.Second * p.getValue().Duration)
	p.timer.Start()
}

func (p *playlist) Prev(errorsChan chan error) {
	if p.timer == nil {
		errorsChan <- fmt.Errorf("[Playlist.Prev]:timer don't init")
		return
	}
	p.timer.Stop()
	p.currentSong = p.data.Prev(p.currentSong)
	if p.currentSong == nil {
		p.currentSong = p.data.Back()
	}
	p.timer = timer.NewTimer(time.Second * p.currentSong.Value.(*Song).Duration)
	p.timer.Start()
}

func (p *playlist) getValue() *Song {
	return p.data.GetValue(p.currentSong).(*Song)
}

func (p *playlist) update(in *list.Element, out *Song) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data.GetValue(in).(*Song).Name = out.Name
	p.data.GetValue(in).(*Song).Duration = out.Duration
}
