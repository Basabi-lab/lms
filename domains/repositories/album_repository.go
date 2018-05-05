package repositories

import "github.com/Basabi-lab/lms/domains/models"

type AlbumRepository interface {
	GetByID(id uint64) (*models.Album, error)
	GetAll() ([]*models.Album, error)
	Create(cd *models.Album) (uint64, error)
}
