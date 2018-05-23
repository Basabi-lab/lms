package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type AlbumCreateUsecaseResult struct {
	Album *models.Album
}

type AlbumCreateUsecaseExt interface {
	Create(c *gin.Context) (*AlbumCreateUsecaseResult, error)
}

type albumCreateUsecase struct {
	repo repositories.AlbumRepository
}

func NewAlbumCreateUsecase(repo repositories.AlbumRepository) AlbumCreateUsecaseExt {
	return &albumCreateUsecase{
		repo: repo,
	}
}

func NewAlbumCreateUsecaseResult(album *models.Album) *AlbumCreateUsecaseResult {
	return &AlbumCreateUsecaseResult{
		Album: album,
	}
}

func (use *albumCreateUsecase) Create(c *gin.Context) (*AlbumCreateUsecaseResult, error) {
	var album *models.Album
	err := c.ShouldBindJSON(&album)
	if err != nil {
		return nil, err
	}
	id, err := use.repo.Create(album)
	album.ID = id

	return NewAlbumCreateUsecaseResult(album), err
}
