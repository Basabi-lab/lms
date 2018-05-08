package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type AllSongPresenterExt interface {
	ToByteList(acur *usecases.AllSongUsecaseResult) (*AllSongPresenterResult, error)
}

type AllSongPresenterResult struct {
	JsonByteList []byte
}

type allSongPresenter struct{}

func NewAllSongPresenter() AllSongPresenterExt {
	return &allSongPresenter{}
}

func NewAllSongPresenterResult(json []byte) *AllSongPresenterResult {
	return &AllSongPresenterResult{
		JsonByteList: json,
	}
}

func (songu *allSongPresenter) ToByteList(asur *usecases.AllSongUsecaseResult) (*AllSongPresenterResult, error) {
	json, _ := json.Marshal(asur.Songs)
	return NewAllSongPresenterResult(json), nil
}
