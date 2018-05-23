package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type songClearMysqlMock struct {
	db *gorm.DB
	dbs.MixInSongMysql
}

func newSongClearMysqlMock(db *gorm.DB) repositories.SongRepository {
	return &songClearMysqlMock{
		db: db,
	}
}

func (mock *songClearMysqlMock) Clear() error {
	return nil
}

func TestSongClearUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewSongClearUsecase(newSongClearMysqlMock(db))

	expect := TestSongClearUsecaseResult()

	songsResult, err := use.Clear(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, songsResult)
}
