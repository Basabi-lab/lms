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

type artistCreateMysqlMock struct {
	db *gorm.DB
	dbs.MixInArtistMysql
}

func (mock *artistCreateMysqlMock) Create(cd *models.Artist) (uint, error) {
	return 0, nil
}

func newArtistCreateMysqlMock(db *gorm.DB) repositories.ArtistRepository {
	return &artistCreateMysqlMock{
		db: db,
	}
}

func TestArtistCreateUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewArtistCreateUsecase(newArtistCreateMysqlMock(db))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":\"Artist name\", \"biography\":\"Artist biography\"}"))
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)

	result, err := use.Create(c)
	assert.NoError(t, err)

	expect := TestArtistCreateUsecaseResult()
	expect.Artist.CreatedAt = time.Time{}
	expect.Artist.UpdatedAt = time.Time{}
	expect.Artist.DeletedAt = nil
	assert.Equal(t, expect, result)
}
