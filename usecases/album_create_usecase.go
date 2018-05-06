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
	ar repositories.AlbumRepository
}

func NewAlbumCreateUsecase(ar repositories.AlbumRepository) AlbumCreateUsecaseExt {
	return &albumCreateUsecase{
		ar: ar,
	}
}

func NewAlbumCreateUsecaseResult(id uint) *AlbumCreateUsecaseResult {
	return &AlbumCreateUsecaseResult{
		ID: id,
	}
}

func (albumu *albumCreateUsecase) Create(c *gin.Context) (*AlbumCreateUsecaseResult, error) {
	var json models.Album
	err := c.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	id, err := albumu.ar.Create(&json)

	return NewAlbumCreateUsecaseResult(id), err
}
