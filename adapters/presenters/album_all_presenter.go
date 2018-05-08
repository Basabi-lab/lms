package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AlbumAllPresenterExt interface {
	ToByteList(acur *usecases.AlbumAllUsecaseResult) (*AlbumAllPresenterResult, error)
}

type AlbumAllPresenterResult struct {
	JsonByteList []byte
}

type albumAllPresenter struct{}

func NewAlbumAllPresenter() AlbumAllPresenterExt {
	return &albumAllPresenter{}
}

func NewAlbumAllPresenterResult(json []byte) *AlbumAllPresenterResult {
	return &AlbumAllPresenterResult{
		JsonByteList: json,
	}
}

func (albumu *albumAllPresenter) ToByteList(acur *usecases.AlbumAllUsecaseResult) (*AlbumAllPresenterResult, error) {
	json, _ := json.Marshal(acur.Albums)
	return NewAlbumAllPresenterResult(json), nil
}
