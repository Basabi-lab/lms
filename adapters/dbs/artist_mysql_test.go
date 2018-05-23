package dbs

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Basabi-lab/lms/domains/models"
)

func TestArtistGetAll(t *testing.T) {
	db, mock := models.ConnectMockDB("test_artist_get_all")
	defer db.Close()

	ar := NewArtistMysql(db)

	artist := models.TestArtistData()

	var artistCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "name", "biography"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(artistCols).
		AddRow(artist.ID, artist.CreatedAt, artist.UpdatedAt, artist.DeletedAt, artist.Name, artist.Biography))

	artists, err := ar.GetAll()
	assert.NoError(t, err)

	expect := []*models.Artist{}
	expect = append(expect, artist)
	assert.Equal(t, expect, artists)
}

func TestArtistGetByID(t *testing.T) {
	db, mock := models.ConnectMockDB("test_artist_get_by_id")
	defer db.Close()

	ar := NewArtistMysql(db)

	expect := models.TestArtistData()

	var artistCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "name", "biography"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(artistCols).
		AddRow(expect.ID, expect.CreatedAt, expect.UpdatedAt, expect.DeletedAt, expect.Name, expect.Biography))

	artist, err := ar.GetByID(0)
	assert.NoError(t, err)

	assert.Equal(t, expect, artist)
}

func TestArtistCreate(t *testing.T) {
	db, mock := models.ConnectMockDB("test_artist_create")
	defer db.Close()

	ar := NewArtistMysql(db)

	expect := models.TestArtistData()

	mock.ExpectExec("INSERT INTO").
		WithArgs(expect.CreatedAt, expect.UpdatedAt, expect.DeletedAt, expect.Name, expect.Biography).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := ar.Create(expect)
	assert.NoError(t, err)
}

func TestArtistClear(t *testing.T) {
	db, mock := models.ConnectMockDB("test_artist_clear")
	defer db.Close()

	// ormのdeleteはsoft deleteなので、updateでdelete_atが更新される
	// このとき、任意のtimeをdelete_atに挿入できないのでテストできない
	// なので、とりあえず、update命令をキャッチはしているが、argは指定していない
	mock.ExpectExec("UPDATE").
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))

	ar := NewArtistMysql(db)

	err := ar.Clear()
	assert.NoError(t, err)
}
