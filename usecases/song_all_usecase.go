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
	ar repositories.SongRepository
}

func NewSongAllUsecase(ar repositories.SongRepository) SongAllUsecaseExt {
	return &songAllUsecase{
		ar: ar,
	}
}

func NewSongAllUsecaseResult(songs []*models.Song) *SongAllUsecaseResult {
	return &SongAllUsecaseResult{
		Songs: songs,
	}
}

func (songu *songAllUsecase) All(c *gin.Context) (*SongAllUsecaseResult, error) {
	songs, err := songu.ar.GetAll()

	return NewSongAllUsecaseResult(songs), err
}
