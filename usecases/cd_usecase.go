package usecases

import (
	"github.com/gin-gonic/gin"

	"github.com/kokoax/music_lab/domains/repositories"
	"github.com/kokoax/music_lab/domains/models"
)

type CDUsecase interface {
	GetById(id int64) (*models.CD, error)
	Create(cd *models.CD) (int64, error)
}

type cdUsecase struct {
	cdr repositories.CDRepository
}

func NewCDUsecase(cdr *repositories.CDRepository) repositories.CDRepository{
	return cdUsecase {
		cdr: cdr,
	}
}

func (cdu *CDUsecase) GetById(c *gin.Context) (*models.CD error) {
}

func (cdu *CDUsecase) Create(c *gin.Context) (int64 error) {
}
