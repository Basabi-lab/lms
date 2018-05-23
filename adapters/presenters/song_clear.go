package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type SongClearPresenterExt interface {
	Response(useResult *usecases.SongClearUsecaseResult) (*SongClearPresenterResult, error)
}

type SongClearPresenterResult struct {
	Res []byte
}

type songClearPresenter struct{}

func NewSongClearPresenter() SongClearPresenterExt {
	return &songClearPresenter{}
}

func NewSongClearPresenterResult(json []byte) *SongClearPresenterResult {
	return &SongClearPresenterResult{
		Res: json,
	}
}

func (pre *songClearPresenter) Response(useResult *usecases.SongClearUsecaseResult) (*SongClearPresenterResult, error) {
	json, _ := json.Marshal(map[string]error{"message": useResult.Error})
	return NewSongClearPresenterResult(json), nil
}
