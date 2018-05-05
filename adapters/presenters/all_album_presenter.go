package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AllAlbumPresenterExt interface {
	ToByteList(acur usecases.AllAlbumUsecaseResult) (*AllAlbumPresenterResult, error)
}

type AllAlbumPresenterResult struct {
	JsonByteList []byte
}

type allAlbumPresenter struct{}

func NewAllAlbumPresenter() AllAlbumPresenterExt {
	return &allAlbumPresenter{}
}

func NewAllAlbumPresenterResult(json []byte) *AllAlbumPresenterResult {
	return &AllAlbumPresenterResult{
		JsonByteList: json,
	}
}

func (albumu *allAlbumPresenter) ToByteList(acur usecases.AllAlbumUsecaseResult) (*AllAlbumPresenterResult, error) {
	json, _ := json.Marshal(acur.Albums)
	return NewAllAlbumPresenterResult(json), nil
}
