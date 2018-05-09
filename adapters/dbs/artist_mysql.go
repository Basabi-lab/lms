package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/domains/repositories"
)

type artistMysql struct {
	db *gorm.DB
}

func NewArtistMysql(db *gorm.DB) repositories.ArtistRepository {
	return &artistMysql{
		db: db,
	}
}

func (c *artistMysql) GetByID(id uint) (*models.Artist, error) {
	artist := &models.Artist{}
	err := c.db.First(&artist, "id = ?", id).Error

	return artist, err
}

func (c *artistMysql) GetAll() ([]*models.Artist, error) {
	artists := []*models.Artist{}
	err := c.db.Find(&artists).Error

	return artists, err
}

func (c *artistMysql) Create(artist *models.Artist) (uint, error) {
	err := c.db.Create(&artist).Error

	return artist.ID, err
}
