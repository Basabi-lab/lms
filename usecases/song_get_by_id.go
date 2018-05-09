package usecases

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type SongGetByIDUsecaseResult struct {
	Song *models.Song
}

type SongGetByIDUsecaseExt interface {
	GetByID(c *gin.Context) (*SongGetByIDUsecaseResult, error)
}

type songGetByIDUsecase struct {
	sr repositories.SongRepository
}

func NewSongGetByIDUsecase(sr repositories.SongRepository) SongGetByIDUsecaseExt {
	return &songGetByIDUsecase{
		sr: sr,
	}
}

func NewSongGetByIDUsecaseResult(song *models.Song) *SongGetByIDUsecaseResult {
	return &SongGetByIDUsecaseResult{
		Song: song,
	}
}

func (songu *songGetByIDUsecase) GetByID(c *gin.Context) (*SongGetByIDUsecaseResult, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	song, err := songu.sr.GetByID(uint(id))

	return NewSongGetByIDUsecaseResult(song), err
}
