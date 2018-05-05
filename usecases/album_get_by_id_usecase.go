package usecases

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type AlbumGetByIDUsecaseResult struct {
	Album *models.Album
}

type AlbumGetByIDUsecaseExt interface {
	GetByID(c *gin.Context) (*AlbumGetByIDUsecaseResult, error)
}

type albumGetByIDUsecase struct {
	ar repositories.AlbumRepository
}

func NewAlbumGetByIDUsecase(ar repositories.AlbumRepository) AlbumGetByIDUsecaseExt {
	return &albumGetByIDUsecase{
		ar: ar,
	}
}

func NewAlbumGetByIDUsecaseResult(album *models.Album) *AlbumGetByIDUsecaseResult {
	return &AlbumGetByIDUsecaseResult{
		Album: album,
	}
}

func (albumu *albumGetByIDUsecase) GetByID(c *gin.Context) (*AlbumGetByIDUsecaseResult, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	album, err := albumu.ar.GetByID(id)

	return NewAlbumGetByIDUsecaseResult(album), err
}
