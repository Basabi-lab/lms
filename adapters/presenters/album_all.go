package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AlbumAllPresenterExt interface {
	ToByteList(useResult *usecases.AlbumAllUsecaseResult) (*AlbumAllPresenterResult, error)
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

func (pre *albumAllPresenter) ToByteList(useResult *usecases.AlbumAllUsecaseResult) (*AlbumAllPresenterResult, error) {
	json, _ := json.Marshal(useResult.Albums)
	return NewAlbumAllPresenterResult(json), nil
}
