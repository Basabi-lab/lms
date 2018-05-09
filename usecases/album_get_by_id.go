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
	repo repositories.AlbumRepository
}

func NewAlbumGetByIDUsecase(repo repositories.AlbumRepository) AlbumGetByIDUsecaseExt {
	return &albumGetByIDUsecase{
		repo: repo,
	}
}

func NewAlbumGetByIDUsecaseResult(album *models.Album) *AlbumGetByIDUsecaseResult {
	return &AlbumGetByIDUsecaseResult{
		Album: album,
	}
}

func (use *albumGetByIDUsecase) GetByID(c *gin.Context) (*AlbumGetByIDUsecaseResult, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	album, err := use.repo.GetByID(uint(id))

	return NewAlbumGetByIDUsecaseResult(album), err
}
