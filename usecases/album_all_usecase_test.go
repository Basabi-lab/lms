package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumAllMysqlMock struct {
	db *gorm.DB
}

func newAlbumAllMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumAllMysqlMock{
		db: db,
	}
}

func (amm *albumAllMysqlMock) GetByID(id uint) (*models.Album, error) {
	return nil, nil
}

func (amm *albumAllMysqlMock) GetAll() ([]*models.Album, error) {
	album := models.Album{
		Title: "Album title",
	}

	return []*models.Album{&album}, nil
}

func (amm *albumAllMysqlMock) Create(cd *models.Album) (uint, error) {
	return 0, nil
}

func TestAlbumAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	aau := NewAlbumAllUsecase(newAlbumAllMysqlMock(db))

	album := models.Album{
		Title: "Album title",
	}

	albums := []*models.Album{}
	albums = append(albums, &album)
	expect := &AlbumAllUsecaseResult{
		Albums: albums,
	}

	albumsResult, err := aau.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumsResult)
}
