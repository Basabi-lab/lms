package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type musicMysql struct {
	db *gorm.DB
}

type MixInMusicRepository struct{}

func (mmr *MixInMusicRepository) MusicRepository(db *gorm.DB) repositories.MusicRepository {
	return NewMusicMysql(db)
}

func NewMusicMysql(db *gorm.DB) repositories.MusicRepository {
	return &musicMysql{
		db: db,
	}
}

func (m *musicMysql) GetByID(id int64) (*models.Music, error) {
	music := &models.Music{}
	err := m.db.Find(&music).Error

	return music, err
}

func (m *musicMysql) Create(music *models.Music) (int64, error) {
	err := m.db.Create(&music).Error

	return 0, err
}
