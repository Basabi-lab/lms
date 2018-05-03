package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type cdMysql struct {
	db *gorm.DB
}

func NewCDMysql(db *gorm.DB) repositories.CDRepository {
	return cdMysql{
		db: db,
	}
}

func (c *CDMysql) GetByID(id int64) (*models.CD, error) {
	cd := *models.CD{}
	err := c.db.Find(&cd, "id = ?", id).Error

	return cd, err
}

func (c *CDMysql) Create(cd *models.CD) (int64, error) {
	err := m.db.Create(&cd).Error

	return err
}
