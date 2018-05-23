package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type SongCreateUsecaseResult struct {
	Song *models.Song
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

func NewSongCreateUsecaseResult(song *models.Song) *SongCreateUsecaseResult {
	return &SongCreateUsecaseResult{
		Song: song,
	}
}

func (use *songCreateUsecase) Create(c *gin.Context) (*SongCreateUsecaseResult, error) {
	var song *models.Song
	err := c.ShouldBindJSON(&song)
	if err != nil {
		return nil, err
	}
	id, err := use.repo.Create(song)
	song.ID = id

	return NewSongCreateUsecaseResult(song), err
}
