package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MusicMysql struct {
	db *gorm.DB
}

func NewMusicMysql(db *gorm.DB) repositories.MusicRepository {
	return MusicMysql{
		db: db,
	}
}

func (m *MusicMysql) GetByID(id int64) (*models.Music, error) {
	music := *models.Music{}
	err := m.db.Find(&music).Error

	return music, err
}

func (m *MusicMysql) Create(music *models.Music) error {
	err := m.db.Create(&music).Error

	return err
}
