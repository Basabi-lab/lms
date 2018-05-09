package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/test"
	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumAllToJson(t *testing.T) {
	pre := NewAlbumAllPresenter()
	album1 := test.TestAlbumData()
	album2 := test.TestAlbumData()

	albums := []*models.Album{}
	albums = append(albums, album1, album2)
	usecaseResult := &usecases.AlbumAllUsecaseResult{
		Albums: albums,
	}

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
