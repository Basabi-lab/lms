package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumAllMysqlMock struct {
	db *gorm.DB
	dbs.MixInAlbumMysql
}

func newAlbumAllMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumAllMysqlMock{
		db: db,
	}
}

func (mock *albumAllMysqlMock) GetAll() ([]*models.Album, error) {
	album := models.TestAlbumData()

	return []*models.Album{album, album}, nil
}

func TestAlbumAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewAlbumAllUsecase(newAlbumAllMysqlMock(db))

	expect := TestAlbumAllUsecaseResult()

	albumsResult, err := use.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumsResult)
}
