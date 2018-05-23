package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type ArtistCreateUsecaseResult struct {
	Artist *models.Artist
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

func NewArtistCreateUsecaseResult(artist *models.Artist) *ArtistCreateUsecaseResult {
	return &ArtistCreateUsecaseResult{
		Artist: artist,
	}
}

func (use *artistCreateUsecase) Create(c *gin.Context) (*ArtistCreateUsecaseResult, error) {
	var artist *models.Artist
	err := c.ShouldBindJSON(&artist)
	if err != nil {
		return nil, err
	}
	id, err := use.repo.Create(artist)
	artist.ID = id

	return NewArtistCreateUsecaseResult(artist), err
}
