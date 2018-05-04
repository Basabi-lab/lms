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

type MixInAlbumRepository struct{}

func (mmr *MixInAlbumRepository) AlbumRepository(db *gorm.DB) repositories.AlbumRepository {
	return NewAlbumMysql(db)
}

func NewAlbumMysql(db *gorm.DB) repositories.AlbumRepository {
	return &albumMysql{
		db: db,
	}
}

func (c *albumMysql) GetByID(id int64) (*models.Album, error) {
	album := &models.Album{}
	err := c.db.Find(&album, "id = ?", id).Error

	return album, err
}

func (c *albumMysql) GetAll() ([]*models.Album, error) {
	albums := []*models.Album{}
	err := c.db.Find(albums).Error

	return albums, err
}

func (c *albumMysql) Create(album *models.Album) (int64, error) {
	err := c.db.Create(&album).Error

	return 0, err
}

type AlbumRepository interface {
	GetByID(id int64) (*models.Album, error)
	GetAll() ([]*models.Album, error)
	Create(album *models.Album) (int64, error)
}
