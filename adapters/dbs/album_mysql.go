package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type albumMysql struct {
	db *gorm.DB
}

func NewAlbumMysql(db *gorm.DB) repositories.AlbumRepository {
	return &albumMysql{
		db: db,
	}
}

func (c *albumMysql) GetByID(id uint) (*models.Album, error) {
	album := &models.Album{}
	err := c.db.First(&album, "id = ?", id).Error

	return album, err
}

func (c *albumMysql) GetAll() ([]*models.Album, error) {
	albums := []*models.Album{}
	err := c.db.Find(&albums).Error

	return albums, err
}

func (c *albumMysql) Create(album *models.Album) (uint, error) {
	err := c.db.Create(&album).Error

	return album.ID, err
}
