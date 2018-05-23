package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumClearMysqlMock struct {
	db *gorm.DB
	dbs.MixInAlbumMysql
}

func newAlbumClearMysqlMock(db *gorm.DB) repositories.AlbumRepository {
	return &albumClearMysqlMock{
		db: db,
	}
}

func (mock *albumClearMysqlMock) Clear() error {
	return nil
}

func TestAlbumClearUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewAlbumClearUsecase(newAlbumClearMysqlMock(db))

	expect := TestAlbumClearUsecaseResult()

	albumsResult, err := use.Clear(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, albumsResult)
}
