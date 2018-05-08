package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/test"
	"github.com/Basabi-lab/lms/usecases"
)

func TestSongAllToJson(t *testing.T) {
	aap := NewSongAllPresenter()
	song1 := test.TestSongData()
	song2 := test.TestSongData()

	songs := []*models.Song{}
	songs = append(songs, song1, song2)
	usecaseResult := &usecases.SongAllUsecaseResult{
		Songs: songs,
	}

	_, err := aap.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
