package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type ArtistAllUsecaseResult struct {
	Artists []*models.Artist
}

type ArtistAllUsecaseExt interface {
	All(c *gin.Context) (*ArtistAllUsecaseResult, error)
}

type artistAllUsecase struct {
	repo repositories.ArtistRepository
}

func NewArtistAllUsecase(repo repositories.ArtistRepository) ArtistAllUsecaseExt {
	return &artistAllUsecase{
		repo: repo,
	}
}

func NewArtistAllUsecaseResult(artists []*models.Artist) *ArtistAllUsecaseResult {
	return &ArtistAllUsecaseResult{
		Artists: artists,
	}
}

func (use *artistAllUsecase) All(c *gin.Context) (*ArtistAllUsecaseResult, error) {
	artists, err := use.repo.GetAll()

	return NewArtistAllUsecaseResult(artists), err
}
