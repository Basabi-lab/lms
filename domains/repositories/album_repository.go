package repositories

import "github.com/basabi-lab/lms/domains/models"

type AlbumRepository interface {
	GetByID(id int64) (*models.Album, error)
	GetAll() ([]*models.Album, error)
	Create(cd *models.Album) (int64, error)
}
