package playlist

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func initPlaylist() *playlist {
	playlist := New()
	for i := 0; i < 5; i++ {
		song := NewSong(int64(i), fmt.Sprintf("test_%d", i), time.Duration(3))
		playlist.AddSong(song)
	}
	return playlist
}

func TestGetSong(t *testing.T) {
	type testCase struct {
		name     string
		result   *Song
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct",
			result:   NewSong(0, "test_0", time.Duration(3)),
			hasError: false,
		},
		{
			name:     "incorrect",
			hasError: true,
		},
	}

	playlistError := New()
	playlist := initPlaylist()

	for _, test := range testTable {
		errorChan := make(chan error, 1)
		resultChan := make(chan *Song, 1)
		if test.hasError {
			playlistError.GetSong(errorChan, resultChan)
			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
			}
		} else {
			go playlist.Play(errorChan)
			time.Sleep(50 * Millisecond)
			playlist.GetSong(errorChan, resultChan)
			select {
			case err := <-errorChan:
				log.Fatal(err)
			default:
				assert.Equal(t, test.result, <-resultChan)
			}
			time.Sleep(1 * Second)
		}
	}
}

func TestDeleteSong(t *testing.T) {
	type testCase struct {
		name     string
		input    int64
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct",
			input:    1,
			hasError: false,
		},
		{
			name:     "incorrect",
			input:    0,
			hasError: true,
		},
	}

	playlistError := New()
	playlist := initPlaylist()

	for _, test := range testTable {
		errorChan := make(chan error, 1)
		resultChan := make(chan *Song, 1)
		if test.hasError {
			playlistError.DeleteSong(test.input, errorChan)
			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
			}
		} else {
			go playlist.Play(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.DeleteSong(test.input, errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.GetSong(errorChan, resultChan)
			time.Sleep(50 * Millisecond)
			select {
			case err := <-errorChan:
				log.Fatal(err)
			case song := <-resultChan:
				assert.NotEqual(t, test.input, song.Id)
			}
			time.Sleep(1 * Second)
		}
	}
}

func TestUpdateSong(t *testing.T) {
	type testCase struct {
		name     string
		input    *Song
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct",
			input:    NewSong(0, "test_100", time.Duration(3)),
			hasError: false,
		},
		{
			name:     "incorrect",
			input:    NewSong(0, "test_100", time.Duration(3)),
			hasError: true,
		},
	}

	playlistError := New()
	playlist := initPlaylist()

	for _, test := range testTable {
		errorChan := make(chan error, 1)
		resultChan := make(chan *Song, 1)
		if test.hasError {
			playlistError.UpdateSong(test.input, errorChan)
			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
			}
		} else {
			go playlist.Play(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.UpdateSong(test.input, errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.GetSong(errorChan, resultChan)
			time.Sleep(50 * Millisecond)
			select {
			case err := <-errorChan:
				log.Fatal(err)
			case song := <-resultChan:
				assert.NotEqual(t, test.input.Name, song.Name)
			}
			time.Sleep(1 * Second)
		}
	}
}

func TestStopSong(t *testing.T) {
	type testCase struct {
		name     string
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct",
			hasError: false,
		},
		{
			name:     "incorrect",
			hasError: true,
		},
	}

	playlistError := New()
	playlist := initPlaylist()

	for _, test := range testTable {
		errorChan := make(chan error, 1)
		if test.hasError {
			playlistError.Stop(errorChan)
			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
			}
		} else {
			go playlist.Play(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.Stop(errorChan)
			time.Sleep(50 * Millisecond)
			select {
			case err := <-errorChan:
				log.Fatal(err)
			default:
				assert.True(t, true)
			}
			time.Sleep(1 * Second)
		}
	}
}

func TestNextSong(t *testing.T) {
	type testCase struct {
		name     string
		result   int64
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct",
			result:   1,
			hasError: false,
		},
		{
			name:     "incorrect",
			hasError: true,
		},
	}

	playlistError := New()
	playlist := initPlaylist()

	for _, test := range testTable {
		errorChan := make(chan error, 1)
		resultChan := make(chan *Song, 1)
		if test.hasError {
			playlistError.Next(errorChan)
			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
			}
		} else {
			go playlist.Play(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.Next(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.GetSong(errorChan, resultChan)
			time.Sleep(50 * Millisecond)

			select {
			case err := <-errorChan:
				log.Fatal(err)
			case song := <-resultChan:
				assert.Equal(t, test.result, song.Id)
			}
			time.Sleep(1 * Second)
		}
	}
}

func TestPrevtSong(t *testing.T) {
	type testCase struct {
		name     string
		result   int64
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct",
			result:   4,
			hasError: false,
		},
		{
			name:     "incorrect",
			hasError: true,
		},
	}

	playlistError := New()
	playlist := initPlaylist()

	for _, test := range testTable {
		errorChan := make(chan error, 1)
		resultChan := make(chan *Song, 1)
		if test.hasError {
			playlistError.Prev(errorChan)
			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
			}
		} else {
			go playlist.Play(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.Prev(errorChan)
			time.Sleep(50 * Millisecond)
			go playlist.GetSong(errorChan, resultChan)
			time.Sleep(50 * Millisecond)
			select {
			case err := <-errorChan:
				log.Fatal(err)
			case song := <-resultChan:
				assert.Equal(t, test.result, song.Id)
			}
			time.Sleep(1 * Second)
		}
	}
}
