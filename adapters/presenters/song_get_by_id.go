package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type SongGetByIDPresenterExt interface {
	ToByteList(acur *usecases.SongGetByIDUsecaseResult) (*SongGetByIDPresenterResult, error)
}

type SongGetByIDPresenterResult struct {
	JsonByteList []byte
}

type songGetByIDPresenter struct{}

func NewSongGetByIDPresenter() SongGetByIDPresenterExt {
	return &songGetByIDPresenter{}
}

func NewSongGetByIDPresenterResult(json []byte) *SongGetByIDPresenterResult {
	return &SongGetByIDPresenterResult{
		JsonByteList: json,
	}
}

func (songu *songGetByIDPresenter) ToByteList(sgbiu *usecases.SongGetByIDUsecaseResult) (*SongGetByIDPresenterResult, error) {
	json, _ := json.Marshal(sgbiu.Song)
	return NewSongGetByIDPresenterResult(json), nil
}
