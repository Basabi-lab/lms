package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kokoax/music_lab/domains/models"
	"github.com/kokoax/music_lab/domains/repositories"
)

type cdMysql struct {
	db *gorm.DB
}

type MixInCDRepository struct{}

func (mmr *MixInMusicRepository) CDRepository(db *gorm.DB) repositories.CDRepository {
	return NewCDMysql(db)
}

func NewCDMysql(db *gorm.DB) repositories.CDRepository {
	return &cdMysql{
		db: db,
	}
}

func (c *cdMysql) GetByID(id int64) (*models.CD, error) {
	cd := &models.CD{}
	err := c.db.Find(&cd, "id = ?", id).Error

	return cd, err
}

func (c *cdMysql) GetAll() ([]*models.CD, error) {
	cds := []*models.CD{}
	err := c.db.Find(cds).Error

	return cds, err
}

func (c *cdMysql) Create(cd *models.CD) (int64, error) {
	err := c.db.Create(&cd).Error

	return 0, err
}

type CDRepository interface {
	GetByID(id int64) (*models.CD, error)
	GetAll() ([]*models.CD, error)
	Create(cd *models.CD) (int64, error)
}
