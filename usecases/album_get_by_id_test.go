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

type albumAlbumGetByIDMysqlMock struct {
	db *gorm.DB
	dbs.MixInAlbumMysql
}

func newAlbumGetByIDMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumAlbumGetByIDMysqlMock{
		db: db,
	}
}

func (mock *albumAlbumGetByIDMysqlMock) GetByID(id uint) (*models.Album, error) {
	album := models.TestAlbumData()

	return album, nil
}

func TestAlbumGetByIDUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewAlbumGetByIDUsecase(newAlbumGetByIDMysqlMock(db))

	expect := TestAlbumGetByIDUsecaseResult()

	c := &gin.Context{}
	c.Params = gin.Params{gin.Param{Key: "id", Value: "10"}}
	albumResult, err := use.GetByID(c)
	assert.NoError(t, err)

	assert.Equal(t, expect, albumResult)
}
