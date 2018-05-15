package models

import (
	"time"

	"github.com/jinzhu/gorm"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func NewModel(id uint, now time.Time) Model {
	loc, _ := time.LoadLocation("UTC")
	date := time.Date(2000, 1, 1, 1, 0, 0, 0, loc)
	return Model{
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

func TestAlbumData() *Album {
	return &Album{
		Model:    NewModel(0, time.Now()),
		Title:    "Album title",
		ArtistID: 1,
	}
}

func TestArtistData() *Artist {
	return &Artist{
		Model:     NewModel(0, time.Now()),
		Name:      "Artist name",
		Biography: "Artist biography",
	}
}

func TestSongData() *Song {
	return &Song{
		Model:    NewModel(0, time.Now()),
		AlbumID:  1,
		ArtistID: 1,
		Title:    "Song Title",
		Track:    1,
		Genre:    "Song Genre",
		Year:     2000,
		Disc:     1,
		Path:     "/home/sample/Music/dir",
	}
}
