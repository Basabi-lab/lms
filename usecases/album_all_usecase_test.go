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

func (mock *albumAllMysqlMock) GetByID(id uint) (*models.Album, error) {
	return nil, nil
}

func (mock *albumAllMysqlMock) GetAll() ([]*models.Album, error) {
	album := models.Album{
		Title: "Album title",
	}

	return []*models.Album{&album}, nil
}

func (mock *albumAllMysqlMock) Create(cd *models.Album) (uint, error) {
	return 0, nil
}

func TestAlbumAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewAlbumAllUsecase(newAlbumAllMysqlMock(db))

	album := models.Album{
		Title: "Album title",
	}

	albums := []*models.Album{}
	albums = append(albums, &album)
	expect := &AlbumAllUsecaseResult{
		Albums: albums,
	}

	albumsResult, err := use.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumsResult)
}
