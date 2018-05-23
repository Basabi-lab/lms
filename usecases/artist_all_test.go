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

type artistAllMysqlMock struct {
	db *gorm.DB
	dbs.MixInArtistMysql
}

func newArtistAllMysqlMock(db *gorm.DB) repositories.ArtistRepository {
	return &artistAllMysqlMock{
		db: db,
	}
}

func (mock *artistAllMysqlMock) GetAll() ([]*models.Artist, error) {
	artist := models.TestArtistData()

	return []*models.Artist{artist, artist}, nil
}

func TestArtistAllUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewArtistAllUsecase(newArtistAllMysqlMock(db))

	expect := TestArtistAllUsecaseResult()

	artistsResult, err := use.All(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, artistsResult)
}
