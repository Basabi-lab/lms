package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type allAlbumMysqlMock struct {
	db *gorm.DB
}

func newAllAlbumMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &allAlbumMysqlMock{
		db: db,
	}
}

func (amm *allAlbumMysqlMock) GetByID(id uint) (*models.Album, error) {
	return nil, nil
}

func (amm *allAlbumMysqlMock) GetAll() ([]*models.Album, error) {
	album := models.Album{
		Title: "Album title",
	}

	return []*models.Album{&album}, nil
}

func (amm *allAlbumMysqlMock) Create(cd *models.Album) (uint, error) {
	return 0, nil
}

func TestAllAlbumUsecase(t *testing.T) {
	db := &gorm.DB{}
	aau := NewAllAlbumUsecase(newAllAlbumMysqlMock(db))

	album := models.Album{
		Title: "Album title",
	}

	albums := []*models.Album{}
	albums = append(albums, &album)
	expect := &AllAlbumUsecaseResult{
		Albums: albums,
	}

	albumsResult, err := aau.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumsResult)
}
