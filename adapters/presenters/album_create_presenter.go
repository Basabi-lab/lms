package presenters

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/usecases"
)

type AlbumCreatePresenterExt interface {
	Response(acur *usecases.AlbumCreateUsecaseResult) (*AlbumCreatePresenterResult, error)
}

type AlbumCreatePresenterResult struct {
	Res *gin.H
}

type albumCreatePresenter struct{}

func NewAlbumCreatePresenter() AlbumCreatePresenterExt {
	return &albumCreatePresenter{}
}

func NewAlbumCreatePresenterResult(res *gin.H) *AlbumCreatePresenterResult {
	return &AlbumCreatePresenterResult{
		Res: res,
	}
}

func (albumu *albumCreatePresenter) Response(acur *usecases.AlbumCreateUsecaseResult) (*AlbumCreatePresenterResult, error) {
	res := &gin.H{"message": "success", "id": acur.ID}

	return NewAlbumCreatePresenterResult(res), nil
}
