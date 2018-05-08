package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type SongAllPresenterExt interface {
	ToByteList(acur *usecases.SongAllUsecaseResult) (*SongAllPresenterResult, error)
}

type SongAllPresenterResult struct {
	JsonByteList []byte
}

type allSongPresenter struct{}

func NewSongAllPresenter() SongAllPresenterExt {
	return &allSongPresenter{}
}

func NewSongAllPresenterResult(json []byte) *SongAllPresenterResult {
	return &SongAllPresenterResult{
		JsonByteList: json,
	}
}

func (songu *allSongPresenter) ToByteList(asur *usecases.SongAllUsecaseResult) (*SongAllPresenterResult, error) {
	json, _ := json.Marshal(asur.Songs)
	return NewSongAllPresenterResult(json), nil
}
