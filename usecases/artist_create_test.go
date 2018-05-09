package usecases

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type artistCreateMysqlMock struct {
	db *gorm.DB
}

func newArtistCreateMysqlMock(db *gorm.DB) repositories.ArtistRepository {
	return &artistCreateMysqlMock{
		db: db,
	}
}

func (amm *artistCreateMysqlMock) GetByID(id uint) (*models.Artist, error) {
	return nil, nil
}

func (amm *artistCreateMysqlMock) GetAll() ([]*models.Artist, error) {
	return nil, nil
}

func (amm *artistCreateMysqlMock) Create(cd *models.Artist) (uint, error) {
	return 0, nil
}

func TestArtistCreateUsecase(t *testing.T) {
	db := &gorm.DB{}
	acu := NewArtistCreateUsecase(newArtistCreateMysqlMock(db))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"album_id\": 1, \"name\":\"artist name\", \"biography\":\"artist biography\"}"))
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)

	_, err := acu.Create(c)
	assert.NoError(t, err)
}
