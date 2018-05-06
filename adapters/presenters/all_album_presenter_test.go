package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/usecases"
)

func TestAllAlbumToJson(t *testing.T) {
	aap := NewAllAlbumPresenter()
	album1 := models.Album{
		Title: "Album title",
		Genre: "Album Genre",
		Year:  2000,
	}
	album2 := models.Album{
		Title: "Album title",
		Genre: "Album Genre",
		Year:  2000,
	}
	albums := []*models.Album{}
	albums = append(albums, &album1, &album2)
	usecaseResult := &usecases.AllAlbumUsecaseResult{
		Albums: albums,
	}

	_, err := aap.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
