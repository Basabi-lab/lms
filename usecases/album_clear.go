package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/repositories"
)

type AlbumClearUsecaseResult struct {
	Error error
}

type AlbumClearUsecaseExt interface {
	Clear(c *gin.Context) (*AlbumClearUsecaseResult, error)
}

type albumClearUsecase struct {
	repo repositories.AlbumRepository
}

func NewAlbumClearUsecase(repo repositories.AlbumRepository) AlbumClearUsecaseExt {
	return &albumClearUsecase{
		repo: repo,
	}
}

func NewAlbumClearUsecaseResult(err error) *AlbumClearUsecaseResult {
	return &AlbumClearUsecaseResult{
		Error: err,
	}
}

func (use *albumClearUsecase) Clear(c *gin.Context) (*AlbumClearUsecaseResult, error) {
	err := use.repo.Clear()

	return NewAlbumClearUsecaseResult(err), err
}
