package dbs

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Basabi-lab/lms/domains/models"
)

func TestAlbumGetAll(t *testing.T) {
	db, mock := models.ConnectMockDB("test_album_get_all")
	defer db.Close()

	ar := NewAlbumMysql(db)

	album := models.TestAlbumData()

	var albumCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "artist_id", "title"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(albumCols).
		AddRow(album.ID, album.CreatedAt, album.UpdatedAt, album.DeletedAt, album.ArtistID, album.Title))

	albums, err := ar.GetAll()
	assert.NoError(t, err)

	expect := []*models.Album{}
	expect = append(expect, album)
	assert.Equal(t, expect, albums)
}

func TestAlbumGetByID(t *testing.T) {
	db, mock := models.ConnectMockDB("test_album_get_by_id")
	defer db.Close()

	ar := NewAlbumMysql(db)

	expect := models.TestAlbumData()

	var albumCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "artist_id", "title"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(albumCols).
		AddRow(expect.ID, expect.CreatedAt, expect.UpdatedAt, expect.DeletedAt, expect.ArtistID, expect.Title))

	album, err := ar.GetByID(0)
	assert.NoError(t, err)

	assert.Equal(t, expect, album)
}
