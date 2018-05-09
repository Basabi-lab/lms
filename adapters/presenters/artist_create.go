package presenters

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/usecases"
)

type ArtistCreatePresenterExt interface {
	Response(acur *usecases.ArtistCreateUsecaseResult) (*ArtistCreatePresenterResult, error)
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

func (artistu *artistCreatePresenter) Response(acur *usecases.ArtistCreateUsecaseResult) (*ArtistCreatePresenterResult, error) {
	res := &gin.H{"message": "success", "id": acur.ID}

	return NewArtistCreatePresenterResult(res), nil
}
