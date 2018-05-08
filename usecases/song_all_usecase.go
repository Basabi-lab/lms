package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type AllSongUsecaseResult struct {
	Songs []*models.Song
}

type AllSongUsecaseExt interface {
	All(c *gin.Context) (*AllSongUsecaseResult, error)
}

type allSongUsecase struct {
	ar repositories.SongRepository
}

func NewAllSongUsecase(ar repositories.SongRepository) AllSongUsecaseExt {
	return &allSongUsecase{
		ar: ar,
	}
}

func NewAllSongUsecaseResult(songs []*models.Song) *AllSongUsecaseResult {
	return &AllSongUsecaseResult{
		Songs: songs,
	}
}

func (songu *allSongUsecase) All(c *gin.Context) (*AllSongUsecaseResult, error) {
	songs, err := songu.ar.GetAll()

	return NewAllSongUsecaseResult(songs), err
}
