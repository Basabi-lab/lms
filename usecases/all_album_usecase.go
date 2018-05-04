package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/domains/models"
)

type AllAlbumUsecaseResult struct {
	Cds []*models.Album
}

type AllAlbumUsecase interface {
	AllAlbum(c *gin.Context) (*AllAlbumUsecaseResult, error)
}

type allAlbumUsecase struct {
	dbs.MixInAlbumRepository
	db *gorm.DB
}

func NewAllAlbumUsecase(db *gorm.DB) *allAlbumUsecase {
	return &allAlbumUsecase{
		db: db,
	}
}

func NewAllAlbumUsecaseResult(albums []*models.Album) *AllAlbumUsecaseResult {
	return &AllAlbumUsecaseResult{
		Cds: albums,
	}
}

func (albumu *allAlbumUsecase) AllAlbum(c *gin.Context) (*AllAlbumUsecaseResult, error) {
	mr := albumu.AlbumRepository(albumu.db)
	albums, _ := mr.GetAll()

	return NewAllAlbumUsecaseResult(albums), nil
}
