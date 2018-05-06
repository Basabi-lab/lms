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

type albumCreateMysqlMock struct {
	db *gorm.DB
}

func newAlbumCreateMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumCreateMysqlMock{
		db: db,
	}
}

func (amm *albumCreateMysqlMock) GetByID(id uint) (*models.Album, error) {
	return nil, nil
}

func (amm *albumCreateMysqlMock) GetAll() ([]*models.Album, error) {
	return nil, nil
}

func (amm *albumCreateMysqlMock) Create(cd *models.Album) (uint, error) {
	return 0, nil
}

func TestAlbumCreateUsecase(t *testing.T) {
	db := &gorm.DB{}
	acu := NewAlbumCreateUsecase(newAlbumCreateMysqlMock(db))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"title\":\"bar\", \"genre\":\"foo\", \"year\":2000}"))
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)

	_, err := acu.Create(c)
	assert.NoError(t, err)
}
