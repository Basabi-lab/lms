package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/kokoax/music_lab/adapters/dbs"
	"github.com/kokoax/music_lab/domains/models"
)

type AllCDsUsecaseResult struct {
	Cds []*models.CD
}

type AllCDsUsecase interface {
	AllCDs(c *gin.Context) (*AllCDsUsecaseResult, error)
}

type allCDsUsecase struct {
	dbs.MixInMusicRepository
	db *gorm.DB
}

func NewAllCDsUsecase(db *gorm.DB) *allCDsUsecase {
	return &allCDsUsecase{
		db: db,
	}
}

func NewAllCDsUsecaseResult(cds []*models.CD) *AllCDsUsecaseResult {
	return &AllCDsUsecaseResult{
		Cds: cds,
	}
}

func (cdu *allCDsUsecase) AllCDs(c *gin.Context) (*AllCDsUsecaseResult, error) {
	mr := cdu.CDRepository(cdu.db)
	cds, _ := mr.GetAll()

	return NewAllCDsUsecaseResult(cds), nil
}
