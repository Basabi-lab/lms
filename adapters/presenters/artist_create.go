package presenters

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/usecases"
)

type ArtistCreatePresenterExt interface {
	Response(useResult *usecases.ArtistCreateUsecaseResult) (*ArtistCreatePresenterResult, error)
}

type ArtistCreatePresenterResult struct {
	Res *gin.H
}

type artistCreatePresenter struct{}

func NewArtistCreatePresenter() ArtistCreatePresenterExt {
	return &artistCreatePresenter{}
}

func NewArtistCreatePresenterResult(res *gin.H) *ArtistCreatePresenterResult {
	return &ArtistCreatePresenterResult{
		Res: res,
	}
}

func (use *artistCreatePresenter) Response(useResult *usecases.ArtistCreateUsecaseResult) (*ArtistCreatePresenterResult, error) {
	res := &gin.H{"message": "success", "id": useResult.ID}

	return NewArtistCreatePresenterResult(res), nil
}
