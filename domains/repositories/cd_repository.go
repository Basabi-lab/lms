package repositories

import "github.com/kokoax/music_lab/domains/models"

type CDRepository interface {
	GetByID(id int64) (*models.CD, error)
	GetAll() ([]*models.CD, error)
	Create(cd *models.CD) (int64, error)
}
