package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type songAllMysqlMock struct {
	db *gorm.DB
}

func newSongAllMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songAllMysqlMock{
		db: db,
	}
}

func (amm *songAllMysqlMock) GetByID(id uint) (*models.Song, error) {
	return nil, nil
}

func (amm *songAllMysqlMock) GetAll() ([]*models.Song, error) {
	song := models.Song{
		Title: "Song title",
		Genre: "Song Genre",
		Year:  2000,
	}

	return []*models.Song{&song}, nil
}

func (amm *songAllMysqlMock) Create(cd *models.Song) (uint, error) {
	return 0, nil
}

func TestSongAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	aau := NewSongAllUsecase(newSongAllMysqlMock(db))

	song := models.Song{
		Title: "Song title",
		Genre: "Song Genre",
		Year:  2000,
	}

	songs := []*models.Song{}
	songs = append(songs, &song)
	expect := &SongAllUsecaseResult{
		Songs: songs,
	}

	songsResult, err := aau.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, songsResult)
}
