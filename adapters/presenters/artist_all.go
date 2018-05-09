package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type ArtistAllPresenterExt interface {
	ToByteList(useResult *usecases.ArtistAllUsecaseResult) (*ArtistAllPresenterResult, error)
}

type ArtistAllPresenterResult struct {
	JsonByteList []byte
}

type artistAllPresenter struct{}

func NewArtistAllPresenter() ArtistAllPresenterExt {
	return &artistAllPresenter{}
}

func NewArtistAllPresenterResult(json []byte) *ArtistAllPresenterResult {
	return &ArtistAllPresenterResult{
		JsonByteList: json,
	}
}

func (pre *artistAllPresenter) ToByteList(useResult *usecases.ArtistAllUsecaseResult) (*ArtistAllPresenterResult, error) {
	json, _ := json.Marshal(useResult.Artists)
	return NewArtistAllPresenterResult(json), nil
}
