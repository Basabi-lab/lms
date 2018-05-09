package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type artistAllMysqlMock struct {
	db *gorm.DB
}

func newArtistAllMysqlMock(db *gorm.DB) repositories.ArtistRepository {
	return &artistAllMysqlMock{
		db: db,
	}
}

func (mock *artistAllMysqlMock) GetByID(id uint) (*models.Artist, error) {
	return nil, nil
}

func (mock *artistAllMysqlMock) GetAll() ([]*models.Artist, error) {
	artist := models.TestArtistData()

	return []*models.Artist{artist}, nil
}

func (mock *artistAllMysqlMock) Create(cd *models.Artist) (uint, error) {
	return 0, nil
}

func TestArtistAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewArtistAllUsecase(newArtistAllMysqlMock(db))

	artist := models.TestArtistData()

	artists := []*models.Artist{}
	artists = append(artists, artist)
	expect := &ArtistAllUsecaseResult{
		Artists: artists,
	}

	artistsResult, err := use.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, artistsResult)
}
