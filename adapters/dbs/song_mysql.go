package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type songMysql struct {
	db *gorm.DB
}

type MixInSongRepository struct{}

func (mmr *MixInSongRepository) SongRepository(db *gorm.DB) repositories.SongRepository {
	return NewSongMysql(db)
}

func NewSongMysql(db *gorm.DB) repositories.SongRepository {
	return &songMysql{
		db: db,
	}
}

func (m *songMysql) GetByID(id uint) (*models.Song, error) {
	song := &models.Song{}
	err := m.db.First(&song, "id = ?", id).Error

	return song, err
}

func (c *songMysql) GetAll() ([]*models.Song, error) {
	songs := []*models.Song{}
	err := c.db.Find(&songs).Error

	return songs, err
}

func (m *songMysql) Create(song *models.Song) (uint, error) {
	err := m.db.Create(&song).Error

	return 0, err
}
