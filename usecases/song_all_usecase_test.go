package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type allSongMysqlMock struct {
	db *gorm.DB
}

func newAllSongMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &allSongMysqlMock{
		db: db,
	}
}

func (amm *allSongMysqlMock) GetByID(id uint) (*models.Song, error) {
	return nil, nil
}

func (amm *allSongMysqlMock) GetAll() ([]*models.Song, error) {
	song := models.Song{
		Title: "Song title",
		Genre: "Song Genre",
		Year:  2000,
	}

	return []*models.Song{&song}, nil
}

func (amm *allSongMysqlMock) Create(cd *models.Song) (uint, error) {
	return 0, nil
}

func TestAllSongUsecase(t *testing.T) {
	db := &gorm.DB{}
	aau := NewAllSongUsecase(newAllSongMysqlMock(db))

	song := models.Song{
		Title: "Song title",
		Genre: "Song Genre",
		Year:  2000,
	}

	songs := []*models.Song{}
	songs = append(songs, &song)
	expect := &AllSongUsecaseResult{
		Songs: songs,
	}

	songsResult, err := aau.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, songsResult)
}
