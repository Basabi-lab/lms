package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumGetByIDToJson(t *testing.T) {
	agbip := NewAlbumGetByIDPresenter()
	album := &models.Album{
		Title: "Album title",
	}
	usecaseResult := &usecases.AlbumGetByIDUsecaseResult{
		Album: album,
	}

	_, err := agbip.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
