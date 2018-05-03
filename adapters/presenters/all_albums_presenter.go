package presenters

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/basabi-lab/lms/usecases"
)

type AllAlbumsPresenter interface {
	ToJson(acur usecases.AllAlbumsUsecaseResult) (*AllAlbumsPresenterResult, error)
}

type Albums struct {
}

type AllAlbumsPresenterResult struct {
	Json []byte
}

type allAlbumsPresenter struct{}

func NewAllAlbumsPresenter(db *gorm.DB) *allAlbumsPresenter {
	return &allAlbumsPresenter{}
}

func NewAllAlbumsPresenterResult(json []byte) *AllAlbumsPresenterResult {
	return &AllAlbumsPresenterResult{
		Json: json,
	}
}

func (albumu *allAlbumsPresenter) ToJson(acur usecases.AllAlbumsUsecaseResult) (*AllAlbumsPresenterResult, error) {
	json, _ := json.Marshal(acur.Cds)
	return NewAllAlbumsPresenterResult(json), nil
}
