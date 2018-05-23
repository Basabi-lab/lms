package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type SongCreatePresenterExt interface {
	ToByteList(useResult *usecases.SongCreateUsecaseResult) (*SongCreatePresenterResult, error)
}

type SongCreatePresenterResult struct {
	JsonByteList []byte
}

type songCreatePresenter struct{}

func NewSongCreatePresenter() SongCreatePresenterExt {
	return &songCreatePresenter{}
}

func NewSongCreatePresenterResult(json []byte) *SongCreatePresenterResult {
	return &SongCreatePresenterResult{
		JsonByteList: json,
	}
}

func (pre *songCreatePresenter) ToByteList(useResult *usecases.SongCreateUsecaseResult) (*SongCreatePresenterResult, error) {
	json, _ := json.Marshal(useResult.Song)
	return NewSongCreatePresenterResult(json), nil
}
