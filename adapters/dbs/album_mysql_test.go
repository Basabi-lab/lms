package dbs

import (
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Basabi-lab/lms/domains/models"
)

func testData() *models.Album {
	return &models.Album{
		Model: newGormModel(0, time.Now()),
		Title: "Album title",
		Genre: "Album Genre",
		Year:  2000,
	}
}

func TestAlbumGetAll(t *testing.T) {
	db, mock := connectDB("test_album_get_all")
	defer db.Close()

	ar := NewAlbumMysql(db)

	album := testData()

	var albumCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "title", "genre", "year"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(albumCols).
		AddRow(album.ID, album.CreatedAt, album.UpdatedAt, album.DeletedAt, album.Title, album.Genre, album.Year))

	albums, err := ar.GetAll()
	assert.NoError(t, err)

	expect := []*models.Album{}
	expect = append(expect, album)
	assert.Equal(t, expect, albums)
}

func TestAlbumGetByID(t *testing.T) {
	db, mock := connectDB("test_album_get_by_id")
	defer db.Close()

	ar := NewAlbumMysql(db)

	expect := testData()

	var albumCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "title", "genre", "year"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(albumCols).
		AddRow(expect.ID, expect.CreatedAt, expect.UpdatedAt, expect.DeletedAt, expect.Title, expect.Genre, expect.Year))

	album, err := ar.GetByID(0)
	assert.NoError(t, err)

	assert.Equal(t, expect, album)
}
