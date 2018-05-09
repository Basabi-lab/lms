package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type ArtistCreateUsecaseResult struct {
	ID uint
}

type ArtistCreateUsecaseExt interface {
	Create(c *gin.Context) (*ArtistCreateUsecaseResult, error)
}

type artistCreateUsecase struct {
	repo repositories.ArtistRepository
}

func NewArtistCreateUsecase(repo repositories.ArtistRepository) ArtistCreateUsecaseExt {
	return &artistCreateUsecase{
		repo: repo,
	}
}

func NewArtistCreateUsecaseResult(id uint) *ArtistCreateUsecaseResult {
	return &ArtistCreateUsecaseResult{
		ID: id,
	}
}

func (use *artistCreateUsecase) Create(c *gin.Context) (*ArtistCreateUsecaseResult, error) {
	var json models.Artist
	err := c.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	id, err := use.repo.Create(&json)

	return NewArtistCreateUsecaseResult(id), err
}
