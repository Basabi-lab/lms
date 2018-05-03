package dbs

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/gunkan-s/music_lab/domains/models"
)

func TestCocktailCreate(t *testing.T) {
	db, mock := connectDB("test_get_all")
	cr := NewMysqlCocktailRepository(db)
	now := time.Now()
	defer db.Close()

	gm := gorm.Model{
		ID:        0,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: &now,
	}

	album := models.Album{
		Model:  gm,
		Title:  "Album title",
		Artist: "Album Artist",
		Genre:  "Album Genre",
		Year:   2000,
	}

	var cockCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "title", "artist", "genre", "year"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(cockCols).
		AddRow(album.ID, album.CreatedAt, album.UpdatedAt, album.DeletedAt, album.Title, album.Artist, album.Genre, album.Year))

	albums, err := cr.GetAll()

	assert.NoError(t, err)
}
