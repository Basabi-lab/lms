package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/test"
	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistGetByIDToJson(t *testing.T) {
	agbip := NewArtistGetByIDPresenter()
	artist := test.TestArtistData()
	usecaseResult := &usecases.ArtistGetByIDUsecaseResult{
		Artist: artist,
	}

	_, err := agbip.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
