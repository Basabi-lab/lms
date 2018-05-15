package loader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
)

type scannerMock struct {
	Path string
}

func (s *scannerMock) Scan() ([]string, error) {
	return nil, nil
}

func (s *scannerMock) ScanNotUpdated(lastUpdate time.Time) ([]string, error) {
	expect := []string{
		"../tests/song/album_less.mp3",
		"../tests/song/artist_less.mp3",
		"../tests/song/not_updated1.mp3",
		"../tests/song/not_updated2.mp3",
		"../tests/song/other_album.mp3",
		"../tests/song/other_artist.mp3",
		"../tests/song/sample.mp3",
		"../tests/song/title_less.mp3",
		"../tests/song/test_dir/dired.mp3",
		"../tests/song/test_dir/recursive_dir/recursived.mp3",
	}

	return expect, nil
}

type accessorMock struct {
	Host string
}

func (a *accessorMock) LastUpdate() (time.Time, error) {
	return time.Time{}, nil
}
func (a *accessorMock) GetArtistByName(name string) (*models.Artist, error) {
	return nil, nil
}
func (a *accessorMock) GetAlbumByTitle(title string) ([]*models.Album, error) {
	return nil, nil
}
func (a *accessorMock) GetSongByTitle(title string) ([]*models.Song, error) {
	return nil, nil
}
func (a *accessorMock) PostArtist(artist *models.Artist) (*models.Artist, error) {
	artist.ID = 1
	return artist, nil
}
func (a *accessorMock) PostAlbum(album *models.Album) (*models.Album, error) {
	album.ID = 1
	return album, nil
}
func (a *accessorMock) PostSong(song *models.Song) (*models.Song, error) {
	song.ID = 1
	return song, nil
}

func NewScannerMock(path string) ScannerExt {
	return &scannerMock{
		Path: path,
	}
}

func NewAccessorMock(host string) AccessorExt {
	return &accessorMock{
		Host: host,
	}
}

// それぞれ、scanner_test.goとaccessor_test.goで定義しているので、わかりやすいように
// ここにもコメントで書いておく(同パッケージ内なので、参照できる)
// var defPath string = "../tests/song"
// var host string = "http://localhost:8080"

func TestLoad(t *testing.T) {
	scanner := NewScannerMock(defPath)
	accessor := NewAccessorMock(host)

	loader := NewLoader(scanner, accessor)

	err := loader.Load()

	assert.NoError(t, err)
}
