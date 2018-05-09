package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type SongCreateUsecaseResult struct {
	ID uint
}

type SongCreateUsecaseExt interface {
	Create(c *gin.Context) (*SongCreateUsecaseResult, error)
}

type songCreateUsecase struct {
	repo repositories.SongRepository
}

func NewSongCreateUsecase(repo repositories.SongRepository) SongCreateUsecaseExt {
	return &songCreateUsecase{
		repo: repo,
	}
}

func NewSongCreateUsecaseResult(id uint) *SongCreateUsecaseResult {
	return &SongCreateUsecaseResult{
		ID: id,
	}
}

func (use *songCreateUsecase) Create(c *gin.Context) (*SongCreateUsecaseResult, error) {
	var json models.Song
	err := c.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	id, err := use.repo.Create(&json)

	return NewSongCreateUsecaseResult(id), err
}
