package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AlbumClearPresenterExt interface {
	Response(useResult *usecases.AlbumClearUsecaseResult) (*AlbumClearPresenterResult, error)
}

type AlbumClearPresenterResult struct {
	Res []byte
}

type albumClearPresenter struct{}

func NewAlbumClearPresenter() AlbumClearPresenterExt {
	return &albumClearPresenter{}
}

func NewAlbumClearPresenterResult(json []byte) *AlbumClearPresenterResult {
	return &AlbumClearPresenterResult{
		Res: json,
	}
}

func (pre *albumClearPresenter) Response(useResult *usecases.AlbumClearUsecaseResult) (*AlbumClearPresenterResult, error) {
	json, _ := json.Marshal(map[string]error{"message": useResult.Error})
	return NewAlbumClearPresenterResult(json), nil
}
