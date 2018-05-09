package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type AlbumCreateUsecaseResult struct {
	ID uint
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

func NewAlbumCreateUsecaseResult(id uint) *AlbumCreateUsecaseResult {
	return &AlbumCreateUsecaseResult{
		ID: id,
	}
}

func (use *albumCreateUsecase) Create(c *gin.Context) (*AlbumCreateUsecaseResult, error) {
	var json models.Album
	err := c.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	id, err := use.repo.Create(&json)

	return NewAlbumCreateUsecaseResult(id), err
}
