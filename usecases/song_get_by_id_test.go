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

type songSongGetByIDMysqlMock struct {
	db *gorm.DB
}

func newSongGetByIDMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songSongGetByIDMysqlMock{
		db: db,
	}
}

func (mock *songSongGetByIDMysqlMock) GetByID(id uint) (*models.Song, error) {
	song := test.TestSongData()

	return song, nil
}

func (mock *songSongGetByIDMysqlMock) GetAll() ([]*models.Song, error) {
	return nil, nil
}

func (mock *songSongGetByIDMysqlMock) Create(cd *models.Song) (uint, error) {
	return 0, nil
}

func TestSongGetByIDUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewSongGetByIDUsecase(newSongGetByIDMysqlMock(db))

	song := test.TestSongData()

	expect := &SongGetByIDUsecaseResult{
		Song: song,
	}

	c := &gin.Context{}
	c.Params = gin.Params{gin.Param{Key: "id", Value: "10"}}
	songResult, err := use.GetByID(c)
	assert.NoError(t, err)

	assert.Equal(t, expect, songResult)
}
