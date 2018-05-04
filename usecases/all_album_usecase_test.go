package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumMysqlMock struct {
	db *gorm.DB
}

func (amm *albumMysqlMock) GetByID(id int64) (*models.Album, error) {
	return nil, nil
}

func (amm *albumMysqlMock) GetAll() ([]*models.Album, error) {
	album := models.Album{
		Title:  "Album title",
		Artist: "Album Artist",
		Genre:  "Album Genre",
		Year:   2000,
	}

	return []*models.Album{&album}, nil
}

func (amm *albumMysqlMock) Create(cd *models.Album) (int64, error) {
	return 0, nil
}

func newAlbumMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumMysqlMock{
		db: db,
	}
}

func TestAllAlbumUsecase(t *testing.T) {
	db := &gorm.DB{}
	aau := NewAllAlbumUsecase(newAlbumMysqlMock(db))

	album := models.Album{
		Title:  "Album title",
		Artist: "Album Artist",
		Genre:  "Album Genre",
		Year:   2000,
	}

	albums := []*models.Album{}
	albums = append(albums, &album)
	expect := &AllAlbumUsecaseResult{
		Albums: albums,
	}

	albumsResult, err := aau.AllAlbum(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumsResult)
}
