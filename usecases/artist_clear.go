package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/repositories"
)

type ArtistClearUsecaseResult struct {
	Error error
}

type ArtistClearUsecaseExt interface {
	Clear(c *gin.Context) (*ArtistClearUsecaseResult, error)
}

type artistClearUsecase struct {
	repo repositories.ArtistRepository
}

func NewArtistClearUsecase(repo repositories.ArtistRepository) ArtistClearUsecaseExt {
	return &artistClearUsecase{
		repo: repo,
	}
}

func NewArtistClearUsecaseResult(err error) *ArtistClearUsecaseResult {
	return &ArtistClearUsecaseResult{
		Error: err,
	}
}

func (use *artistClearUsecase) Clear(c *gin.Context) (*ArtistClearUsecaseResult, error) {
	err := use.repo.Clear()

	return NewArtistClearUsecaseResult(err), err
}
