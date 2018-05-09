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
	sr repositories.SongRepository
}

func NewSongCreateUsecase(sr repositories.SongRepository) SongCreateUsecaseExt {
	return &songCreateUsecase{
		sr: sr,
	}
}

func NewSongCreateUsecaseResult(id uint) *SongCreateUsecaseResult {
	return &SongCreateUsecaseResult{
		ID: id,
	}
}

func (songu *songCreateUsecase) Create(c *gin.Context) (*SongCreateUsecaseResult, error) {
	var json models.Song
	err := c.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	id, err := songu.sr.Create(&json)

	return NewSongCreateUsecaseResult(id), err
}
