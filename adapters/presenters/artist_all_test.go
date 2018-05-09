package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/test"
	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistAllToJson(t *testing.T) {
	pre := NewArtistAllPresenter()
	artist1 := test.TestArtistData()
	artist2 := test.TestArtistData()

	artists := []*models.Artist{}
	artists = append(artists, artist1, artist2)
	usecaseResult := &usecases.ArtistAllUsecaseResult{
		Artists: artists,
	}

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
