package playlist

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSong(t *testing.T) {
	tests := []struct {
		name       string
		idIn       int64
		nameIn     string
		durationIn time.Duration
		out        *Song
	}{
		{
			name:       "number",
			idIn:       1,
			nameIn:     "test",
			durationIn: time.Duration(10),
			out:        &Song{1, "test", time.Duration(10)},
		},
	}

	for _, test := range tests {
		out := NewSong(test.idIn, test.nameIn, test.durationIn)
		assert.Equal(t, test.out, out)
	}
}
