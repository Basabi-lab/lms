package usecases

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumCreateMysqlMock struct {
	db *gorm.DB
	dbs.MixInAlbumMysql
}

func (mock *albumCreateMysqlMock) Create(cd *models.Album) (uint, error) {
	return 0, nil
}

func newAlbumCreateMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumCreateMysqlMock{
		db: db,
	}
}

func TestAlbumCreateUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewAlbumCreateUsecase(newAlbumCreateMysqlMock(db))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"title\":\"Album title\", \"artist_id\":1}"))
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)

	result, err := use.Create(c)
	assert.NoError(t, err)

	expect := TestAlbumCreateUsecaseResult()
	expect.Album.CreatedAt = time.Time{}
	expect.Album.UpdatedAt = time.Time{}
	expect.Album.DeletedAt = nil
	assert.Equal(t, expect, result)
}
