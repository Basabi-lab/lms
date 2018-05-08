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
	repo repositories.AlbumRepository
}

func NewAlbumAllUsecase(repo repositories.AlbumRepository) AlbumAllUsecaseExt {
	return &albumAllUsecase{
		repo: repo,
	}
}

func NewAlbumAllUsecaseResult(albums []*models.Album) *AlbumAllUsecaseResult {
	return &AlbumAllUsecaseResult{
		Albums: albums,
	}
}

func (use *albumAllUsecase) All(c *gin.Context) (*AlbumAllUsecaseResult, error) {
	albums, err := use.repo.GetAll()

	return NewAlbumAllUsecaseResult(albums), err
}
