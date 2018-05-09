package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AlbumGetByIDPresenterExt interface {
	ToByteList(useResult *usecases.AlbumGetByIDUsecaseResult) (*AlbumGetByIDPresenterResult, error)
}

type AlbumGetByIDPresenterResult struct {
	JsonByteList []byte
}

type albumGetByIDPresenter struct{}

func NewAlbumGetByIDPresenter() AlbumGetByIDPresenterExt {
	return &albumGetByIDPresenter{}
}

func NewAlbumGetByIDPresenterResult(json []byte) *AlbumGetByIDPresenterResult {
	return &AlbumGetByIDPresenterResult{
		JsonByteList: json,
	}
}

func (pre *albumGetByIDPresenter) ToByteList(useResult *usecases.AlbumGetByIDUsecaseResult) (*AlbumGetByIDPresenterResult, error) {
	json, _ := json.Marshal(useResult.Album)
	return NewAlbumGetByIDPresenterResult(json), nil
}
