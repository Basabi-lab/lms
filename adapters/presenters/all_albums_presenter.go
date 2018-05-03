package presenters

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/kokoax/music_lab/usecases"
)

type AllCDsPresenter interface {
	ToJson(acur usecases.AllCDsUsecaseResult) (*AllCDsPresenterResult, error)
}

type CDs struct {
}

type AllCDsPresenterResult struct {
	Json []byte
}

type allCDsPresenter struct{}

func NewAllCDsPresenter(db *gorm.DB) *allCDsPresenter {
	return &allCDsPresenter{}
}

func NewAllCDsPresenterResult(json []byte) *AllCDsPresenterResult {
	return &AllCDsPresenterResult{
		Json: json,
	}
}

func (cdu *allCDsPresenter) ToJson(acur usecases.AllCDsUsecaseResult) (*AllCDsPresenterResult, error) {
	json, _ := json.Marshal(acur.Cds)
	return NewAllCDsPresenterResult(json), nil
}
