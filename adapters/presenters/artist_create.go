package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type ArtistCreatePresenterExt interface {
	ToByteList(useResult *usecases.ArtistCreateUsecaseResult) (*ArtistCreatePresenterResult, error)
}

type ArtistCreatePresenterResult struct {
	JsonByteList []byte
}

type artistCreatePresenter struct{}

func NewArtistCreatePresenter() ArtistCreatePresenterExt {
	return &artistCreatePresenter{}
}

func NewArtistCreatePresenterResult(json []byte) *ArtistCreatePresenterResult {
	return &ArtistCreatePresenterResult{
		JsonByteList: json,
	}
}

func (use *artistCreatePresenter) ToByteList(useResult *usecases.ArtistCreateUsecaseResult) (*ArtistCreatePresenterResult, error) {
	json, _ := json.Marshal(useResult.Artist)

	return NewArtistCreatePresenterResult(json), nil
}
