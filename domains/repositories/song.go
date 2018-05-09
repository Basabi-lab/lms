package repositories

import "github.com/Basabi-lab/lms/domains/models"

type SongRepository interface {
	GetByID(id uint) (*models.Song, error)
	GetAll() ([]*models.Song, error)
	Create(song *models.Song) (uint, error)
}
