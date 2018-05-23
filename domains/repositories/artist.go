package repositories

import "github.com/Basabi-lab/lms/domains/models"

type ArtistRepository interface {
	GetByID(id uint) (*models.Artist, error)
	GetAll() ([]*models.Artist, error)
	Create(cd *models.Artist) (uint, error)
	Clear() error
}
