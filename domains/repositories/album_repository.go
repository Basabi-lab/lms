package repositories

import "github.com/kokoax/music_lab/domains/models"

type AlbumRepository interface {
	GetByID(id int64) (*models.Album, error)
	GetAll() ([]*models.Album, error)
	Create(cd *models.Album) (int64, error)
}
