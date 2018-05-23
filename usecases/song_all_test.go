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

type songAllMysqlMock struct {
	db *gorm.DB
	dbs.MixInSongMysql
}

func newSongAllMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songAllMysqlMock{
		db: db,
	}
}

func (mock *songAllMysqlMock) GetAll() ([]*models.Song, error) {
	song := models.TestSongData()

	return []*models.Song{song, song}, nil
}

func TestSongAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewSongAllUsecase(newSongAllMysqlMock(db))

	expect := TestSongAllUsecaseResult()

	songsResult, err := use.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, songsResult)
}
