package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type artistArtistGetByIDMysqlMock struct {
	db *gorm.DB
}

func newArtistGetByIDMysqlMock(db *gorm.DB) repositories.ArtistRepository {
	return &artistArtistGetByIDMysqlMock{
		db: db,
	}
}

func (mock *artistArtistGetByIDMysqlMock) GetByID(id uint) (*models.Artist, error) {
	artist := models.TestArtistData()

	return artist, nil
}

func (mock *artistArtistGetByIDMysqlMock) GetAll() ([]*models.Artist, error) {
	return nil, nil
}

func (mock *artistArtistGetByIDMysqlMock) Create(cd *models.Artist) (uint, error) {
	return 0, nil
}

func TestArtistGetByIDUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewArtistGetByIDUsecase(newArtistGetByIDMysqlMock(db))

	expect := TestArtistGetByIDResult()

	c := &gin.Context{}
	c.Params = gin.Params{gin.Param{Key: "id", Value: "10"}}
	artistResult, err := use.GetByID(c)
	assert.NoError(t, err)

	assert.Equal(t, expect, artistResult)
}
