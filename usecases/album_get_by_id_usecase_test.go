package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumAlbumGetByIDMysqlMock struct {
	db *gorm.DB
}

func newAlbumGetByIDMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumAlbumGetByIDMysqlMock{
		db: db,
	}
}

func (amm *albumAlbumGetByIDMysqlMock) GetByID(id uint64) (*models.Album, error) {
	album := &models.Album{
		Title: "Album title",
		Genre: "Album Genre",
		Year:  2000,
	}

	return album, nil
}

func (amm *albumAlbumGetByIDMysqlMock) GetAll() ([]*models.Album, error) {
	return nil, nil
}

func (amm *albumAlbumGetByIDMysqlMock) Create(cd *models.Album) (uint64, error) {
	return 0, nil
}

func TestAlbumGetByIDUsecase(t *testing.T) {
	db := &gorm.DB{}
	aau := NewAlbumGetByIDUsecase(newAlbumGetByIDMysqlMock(db))

	album := &models.Album{
		Title: "Album title",
		Genre: "Album Genre",
		Year:  2000,
	}

	expect := &AlbumGetByIDUsecaseResult{
		Album: album,
	}

	albumResult, err := aau.AlbumGetByID(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumResult)
}
