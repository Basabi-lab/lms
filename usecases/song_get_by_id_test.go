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

type songSongGetByIDMysqlMock struct {
	db *gorm.DB
	dbs.MixInSongMysql
}

func newSongGetByIDMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songSongGetByIDMysqlMock{
		db: db,
	}
}

func (mock *songSongGetByIDMysqlMock) GetByID(id uint) (*models.Song, error) {
	song := models.TestSongData()

	return song, nil
}

func TestSongGetByIDUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewSongGetByIDUsecase(newSongGetByIDMysqlMock(db))

	c := &gin.Context{}
	c.Params = gin.Params{gin.Param{Key: "id", Value: "10"}}
	songResult, err := use.GetByID(c)
	assert.NoError(t, err)

	expect := TestSongGetByIDUsecaseResult()
	assert.Equal(t, expect, songResult)
}
