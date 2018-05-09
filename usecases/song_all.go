package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type SongAllUsecaseResult struct {
	Songs []*models.Song
}

type SongAllUsecaseExt interface {
	All(c *gin.Context) (*SongAllUsecaseResult, error)
}

type songAllUsecase struct {
	repo repositories.SongRepository
}

func NewSongAllUsecase(repo repositories.SongRepository) SongAllUsecaseExt {
	return &songAllUsecase{
		repo: repo,
	}
}

func NewSongAllUsecaseResult(songs []*models.Song) *SongAllUsecaseResult {
	return &SongAllUsecaseResult{
		Songs: songs,
	}
}

func (use *songAllUsecase) All(c *gin.Context) (*SongAllUsecaseResult, error) {
	songs, err := use.repo.GetAll()

	return NewSongAllUsecaseResult(songs), err
}
