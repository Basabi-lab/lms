package presenters

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/Basabi-lab/lms/usecases"
)

type AllAlbumPresenter interface {
	ToJson(acur usecases.AllAlbumUsecaseResult) (*AllAlbumPresenterResult, error)
}

type Album struct {
}

type AllAlbumPresenterResult struct {
	Json []byte
}

type allAlbumPresenter struct{}

func NewAllAlbumPresenter(db *gorm.DB) *allAlbumPresenter {
	return &allAlbumPresenter{}
}

func NewAllAlbumPresenterResult(json []byte) *AllAlbumPresenterResult {
	return &AllAlbumPresenterResult{
		Json: json,
	}
}

func (albumu *allAlbumPresenter) ToJson(acur usecases.AllAlbumUsecaseResult) (*AllAlbumPresenterResult, error) {
	json, _ := json.Marshal(acur.Cds)
	return NewAllAlbumPresenterResult(json), nil
}
