package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/basabi-lab/lms/adapters/dbs"
	"github.com/basabi-lab/lms/domains/models"
)

type AllAlbumsUsecaseResult struct {
	Cds []*models.Album
}

type AllAlbumsUsecase interface {
	AllAlbums(c *gin.Context) (*AllAlbumsUsecaseResult, error)
}

type allAlbumsUsecase struct {
	dbs.MixInMusicRepository
	db *gorm.DB
}

func NewAllAlbumsUsecase(db *gorm.DB) *allAlbumsUsecase {
	return &allAlbumsUsecase{
		db: db,
	}
}

func NewAllAlbumsUsecaseResult(albums []*models.Album) *AllAlbumsUsecaseResult {
	return &AllAlbumsUsecaseResult{
		Cds: albums,
	}
}

func (albumu *allAlbumsUsecase) AllAlbums(c *gin.Context) (*AllAlbumsUsecaseResult, error) {
	mr := albumu.AlbumRepository(albumu.db)
	albums, _ := mr.GetAll()

	return NewAllAlbumsUsecaseResult(albums), nil
}
