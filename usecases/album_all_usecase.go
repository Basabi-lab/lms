package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type AlbumAllUsecaseResult struct {
	Albums []*models.Album
}

type AlbumAllUsecaseExt interface {
	All(c *gin.Context) (*AlbumAllUsecaseResult, error)
}

type albumAllUsecase struct {
	ar repositories.AlbumRepository
}

func NewAlbumAllUsecase(ar repositories.AlbumRepository) AlbumAllUsecaseExt {
	return &albumAllUsecase{
		ar: ar,
	}
}

func NewAlbumAllUsecaseResult(albums []*models.Album) *AlbumAllUsecaseResult {
	return &AlbumAllUsecaseResult{
		Albums: albums,
	}
}

func (albumu *albumAllUsecase) All(c *gin.Context) (*AlbumAllUsecaseResult, error) {
	albums, err := albumu.ar.GetAll()

	return NewAlbumAllUsecaseResult(albums), err
}
