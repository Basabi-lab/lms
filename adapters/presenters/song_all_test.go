package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/usecases"
)

func TestSongAllToJson(t *testing.T) {
	pre := NewSongAllPresenter()
	song1 := models.TestSongData()
	song2 := models.TestSongData()

	songs := []*models.Song{}
	songs = append(songs, song1, song2)
	usecaseResult := &usecases.SongAllUsecaseResult{
		Songs: songs,
	}

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
