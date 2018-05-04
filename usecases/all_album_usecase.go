package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type AllAlbumUsecaseResult struct {
	Albums []*models.Album
}

type AllAlbumUsecaseExt interface {
	AllAlbum(c *gin.Context) (*AllAlbumUsecaseResult, error)
}

type allAlbumUsecase struct {
	ar repositories.AlbumRepository
}

func NewAllAlbumUsecase(ar repositories.AlbumRepository) AllAlbumUsecaseExt {
	return &allAlbumUsecase{
		ar: ar,
	}
}

func NewAllAlbumUsecaseResult(albums []*models.Album) *AllAlbumUsecaseResult {
	return &AllAlbumUsecaseResult{
		Albums: albums,
	}
}

func (albumu *allAlbumUsecase) AllAlbum(c *gin.Context) (*AllAlbumUsecaseResult, error) {
	albums, err := albumu.ar.GetAll()

	return NewAllAlbumUsecaseResult(albums), err
}
