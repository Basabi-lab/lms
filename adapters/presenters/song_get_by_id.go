package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type SongGetByIDPresenterExt interface {
	ToByteList(useResult *usecases.SongGetByIDUsecaseResult) (*SongGetByIDPresenterResult, error)
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

func (pre *songGetByIDPresenter) ToByteList(useResult *usecases.SongGetByIDUsecaseResult) (*SongGetByIDPresenterResult, error) {
	json, _ := json.Marshal(useResult.Song)
	return NewSongGetByIDPresenterResult(json), nil
}
