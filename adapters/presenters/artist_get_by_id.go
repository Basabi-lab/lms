package presenters

import (
	"encoding/json"

	"github.com/Basabi-lab/lms/usecases"
)

type ArtistGetByIDPresenterExt interface {
	ToByteList(useResult *usecases.ArtistGetByIDUsecaseResult) (*ArtistGetByIDPresenterResult, error)
}

type ArtistGetByIDPresenterResult struct {
	JsonByteList []byte
}

type artistGetByIDPresenter struct{}

func NewArtistGetByIDPresenter() ArtistGetByIDPresenterExt {
	return &artistGetByIDPresenter{}
}

func NewArtistGetByIDPresenterResult(json []byte) *ArtistGetByIDPresenterResult {
	return &ArtistGetByIDPresenterResult{
		JsonByteList: json,
	}
}

func (pre *artistGetByIDPresenter) ToByteList(useResult *usecases.ArtistGetByIDUsecaseResult) (*ArtistGetByIDPresenterResult, error) {
	json, _ := json.Marshal(useResult.Artist)
	return NewArtistGetByIDPresenterResult(json), nil
}
