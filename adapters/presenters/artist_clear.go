package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type ArtistClearPresenterExt interface {
	Response(useResult *usecases.ArtistClearUsecaseResult) (*ArtistClearPresenterResult, error)
}

type ArtistClearPresenterResult struct {
	Res []byte
}

type artistClearPresenter struct{}

func NewArtistClearPresenter() ArtistClearPresenterExt {
	return &artistClearPresenter{}
}

func NewArtistClearPresenterResult(json []byte) *ArtistClearPresenterResult {
	return &ArtistClearPresenterResult{
		Res: json,
	}
}

func (pre *artistClearPresenter) Response(useResult *usecases.ArtistClearUsecaseResult) (*ArtistClearPresenterResult, error) {
	json, _ := json.Marshal(map[string]error{"message": useResult.Error})
	return NewArtistClearPresenterResult(json), nil
}
