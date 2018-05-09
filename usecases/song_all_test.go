package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
	"github.com/Basabi-lab/lms/test"
)

type songAllMysqlMock struct {
	db *gorm.DB
}

func newSongAllMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songAllMysqlMock{
		db: db,
	}
}

func (mock *songAllMysqlMock) GetByID(id uint) (*models.Song, error) {
	return nil, nil
}

func (mock *songAllMysqlMock) GetAll() ([]*models.Song, error) {
	song := test.TestSongData()

	return []*models.Song{song}, nil
}

func (mock *songAllMysqlMock) Create(cd *models.Song) (uint, error) {
	return 0, nil
}

func TestSongAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewSongAllUsecase(newSongAllMysqlMock(db))

	song := test.TestSongData()

	songs := []*models.Song{}
	songs = append(songs, song)
	expect := &SongAllUsecaseResult{
		Songs: songs,
	}

	songsResult, err := use.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, songsResult)
}
