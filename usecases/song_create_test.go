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

type songCreateMysqlMock struct {
	db *gorm.DB
}

func newSongCreateMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songCreateMysqlMock{
		db: db,
	}
}

func (mock *songCreateMysqlMock) GetByID(id uint) (*models.Song, error) {
	return nil, nil
}

func (mock *songCreateMysqlMock) GetAll() ([]*models.Song, error) {
	return nil, nil
}

func (mock *songCreateMysqlMock) Create(cd *models.Song) (uint, error) {
	return 0, nil
}

func TestSongCreateUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewSongCreateUsecase(newSongCreateMysqlMock(db))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	json := `
		{
			"title":"song title",
			"genre":"song genre",
			"year":2000,
			"artist_id":1,
			"album_id":1,
			"track":1,
			"album_num": 1,
			"dir":"/home/sample/song/dir"
		}
	`
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(json))
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)

	result, err := use.Create(c)
	assert.NoError(t, err)

	expect := TestSongCreateResult()
	assert.Equal(t, expect, result)
}
