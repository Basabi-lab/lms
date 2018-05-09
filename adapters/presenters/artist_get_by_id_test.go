package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistGetByIDToJson(t *testing.T) {
	pre := NewArtistGetByIDPresenter()
	artist := models.TestArtistData()
	usecaseResult := &usecases.ArtistGetByIDUsecaseResult{
		Artist: artist,
	}

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
