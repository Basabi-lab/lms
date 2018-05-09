package usecases

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type ArtistGetByIDUsecaseResult struct {
	Artist *models.Artist
}

type ArtistGetByIDUsecaseExt interface {
	GetByID(c *gin.Context) (*ArtistGetByIDUsecaseResult, error)
}

type artistGetByIDUsecase struct {
	repo repositories.ArtistRepository
}

func NewArtistGetByIDUsecase(repo repositories.ArtistRepository) ArtistGetByIDUsecaseExt {
	return &artistGetByIDUsecase{
		repo: repo,
	}
}

func NewArtistGetByIDUsecaseResult(artist *models.Artist) *ArtistGetByIDUsecaseResult {
	return &ArtistGetByIDUsecaseResult{
		Artist: artist,
	}
}

func (use *artistGetByIDUsecase) GetByID(c *gin.Context) (*ArtistGetByIDUsecaseResult, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	artist, err := use.repo.GetByID(uint(id))

	return NewArtistGetByIDUsecaseResult(artist), err
}
