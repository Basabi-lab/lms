package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AlbumCreatePresenterExt interface {
	ToByteList(useResult *usecases.AlbumCreateUsecaseResult) (*AlbumCreatePresenterResult, error)
}

type AlbumCreatePresenterResult struct {
	JsonByteList []byte
}

type albumCreatePresenter struct{}

func NewAlbumCreatePresenter() AlbumCreatePresenterExt {
	return &albumCreatePresenter{}
}

func NewAlbumCreatePresenterResult(json []byte) *AlbumCreatePresenterResult {
	return &AlbumCreatePresenterResult{
		JsonByteList: json,
	}
}

func (pre *albumCreatePresenter) ToByteList(useResult *usecases.AlbumCreateUsecaseResult) (*AlbumCreatePresenterResult, error) {
	json, _ := json.Marshal(useResult.Album)
	return NewAlbumCreatePresenterResult(json), nil
}
