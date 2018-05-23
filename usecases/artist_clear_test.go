package usecases

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type artistClearMysqlMock struct {
	db *gorm.DB
	dbs.MixInArtistMysql
}

func newArtistClearMysqlMock(db *gorm.DB) repositories.ArtistRepository {
	return &artistClearMysqlMock{
		db: db,
	}
}

func (mock *artistClearMysqlMock) Clear() error {
	return nil
}

func TestArtistClearUsecase(t *testing.T) {
	db := &gorm.DB{}
	use := NewArtistClearUsecase(newArtistClearMysqlMock(db))

	expect := TestArtistClearUsecaseResult()

	artistsResult, err := use.Clear(&gin.Context{})
	assert.NoError(t, err)

	assert.Equal(t, expect, artistsResult)
}
