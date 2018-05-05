package repositories

import "github.com/Basabi-lab/lms/domains/models"

type SongRepository interface {
	GetByID(id int64) (*models.Song, error)
	Create(song *models.Song) (int64, error)
}
