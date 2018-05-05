package repositories

import "github.com/Basabi-lab/lms/domains/models"

type SongRepository interface {
	GetByID(id uint64) (*models.Song, error)
	Create(song *models.Song) (uint64, error)
}
