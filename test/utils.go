package test

import (
	"time"

	"github.com/jinzhu/gorm"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Basabi-lab/lms/domains/models"
)

func NewGormModel(id uint, now time.Time) gorm.Model {
	date := time.Date(2000, 1, 1, 1, 0, 0, 0, &time.Location{})
	return gorm.Model{
		ID:        id,
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: &date,
	}
}

func ConnectMockDB(dns string) (*gorm.DB, sqlmock.Sqlmock) {
	_, mock, _ := sqlmock.NewWithDSN(dns)
	db, _ := gorm.Open("sqlmock", dns)

	return db, mock
}

func TestAlbumData() *models.Album {
	return &models.Album{
		Model: NewGormModel(0, time.Now()),
		Title: "Album title",
	}
}

func TestArtistData() *models.Artist {
	return &models.Artist{
		Model:     NewGormModel(0, time.Now()),
		Name:      "Artist name",
		Biography: "Artist biography",
	}
}

func TestSongData() *models.Song {
	return &models.Song{
		Model:    NewGormModel(0, time.Now()),
		AlbumID:  1,
		ArtistID: 1,
		Title:    "Song Title",
		Track:    1,
		Genre:    "Song Genre",
		Year:     2000,
		AlbumNum: 1,
		Dir:      "/home/sample/Music/dir",
	}
}
