package presenters

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/usecases"
)

type SongCreatePresenterExt interface {
	Response(useResult *usecases.SongCreateUsecaseResult) (*SongCreatePresenterResult, error)
}

type SongCreatePresenterResult struct {
	Res *gin.H
}

type songCreatePresenter struct{}

func NewSongCreatePresenter() SongCreatePresenterExt {
	return &songCreatePresenter{}
}

func NewSongCreatePresenterResult(res *gin.H) *SongCreatePresenterResult {
	return &SongCreatePresenterResult{
		Res: res,
	}
}

func (pre *songCreatePresenter) Response(useResult *usecases.SongCreateUsecaseResult) (*SongCreatePresenterResult, error) {
	res := &gin.H{"message": "success", "id": useResult.ID}

	return NewSongCreatePresenterResult(res), nil
}
