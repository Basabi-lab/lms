package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/repositories"
)

type SongClearUsecaseResult struct {
	Error error
}

type SongClearUsecaseExt interface {
	Clear(c *gin.Context) (*SongClearUsecaseResult, error)
}

type songClearUsecase struct {
	repo repositories.SongRepository
}

func NewSongClearUsecase(repo repositories.SongRepository) SongClearUsecaseExt {
	return &songClearUsecase{
		repo: repo,
	}
}

func NewSongClearUsecaseResult(err error) *SongClearUsecaseResult {
	return &SongClearUsecaseResult{
		Error: err,
	}
}

func (use *songClearUsecase) Clear(c *gin.Context) (*SongClearUsecaseResult, error) {
	err := use.repo.Clear()

	return NewSongClearUsecaseResult(err), err
}
