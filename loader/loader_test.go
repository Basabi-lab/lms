package loader

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vbatts/go-taglib/taglib"

	"github.com/Basabi-lab/lms/domains/models"
)

type scannerMock struct {
	Path string
}

func (s *scannerMock) Scan() ([]string, error) {
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

func (s *scannerMock) ScanNotUpdated(lastUpdate time.Time) ([]string, error) {
	expect := []string{
		"../tests/song/not_updated1.mp3",
		"../tests/song/not_updated2.mp3",
	}

	return expect, nil
}

type scannerMockScanError struct {
	Path string
	scannerMock
}

func (s *scannerMockScanError) Scan() ([]string, error) {
	return nil, fmt.Errorf("Error: scanner Scan")
}

type accessorMock struct {
	Host string
}

func (a *accessorMock) ClearAll() error {
	return nil
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

type accessorMockPostArtistError struct {
	Host string
	accessorMock
}

func (a *accessorMockPostArtistError) PostArtist(artist *models.Artist) (*models.Artist, error) {
	return nil, fmt.Errorf("Error: accessor PostArtist")
}

type accessorMockPostAlbumError struct {
	Host string
	accessorMock
}

func (a *accessorMockPostAlbumError) PostAlbum(artist *models.Album) (*models.Album, error) {
	return nil, fmt.Errorf("Error: accessor PostAlbum")
}

type accessorMockPostSongError struct {
	Host string
	accessorMock
}

func (a *accessorMockPostSongError) PostSong(artist *models.Song) (*models.Song, error) {
	return nil, fmt.Errorf("Error: accessor PostSong")
}

func NewScannerMock(path string) ScannerExt {
	return &scannerMock{
		Path: path,
	}
}

func NewScannerMockScanError(path string) ScannerExt {
	return &scannerMockScanError{
		Path: path,
	}
}

func NewAccessorMock(host string) AccessorExt {
	return &accessorMock{
		Host: host,
	}
}

func NewAccessorMockPostArtistError(host string) AccessorExt {
	return &accessorMockPostArtistError{
		Host: host,
	}
}

func NewAccessorMockPostAlbumError(host string) AccessorExt {
	return &accessorMockPostAlbumError{
		Host: host,
	}
}

func NewAccessorMockPostSongError(host string) AccessorExt {
	return &accessorMockPostSongError{
		Host: host,
	}
}

// それぞれ、scanner_test.goとaccessor_test.goで定義しているので、わかりやすいように
// ここにもコメントで書いておく(同パッケージ内なので、参照できる)
// var defPath string = "../tests/song"
// var host string = "http://localhost:8080"

func TestLoad(t *testing.T) {
	scanner := NewScannerMock(defPath)
	scannerScanErr := NewScannerMockScanError(defPath)

	accessor := NewAccessorMock(host)
	accessorPostArtistErr := NewAccessorMockPostArtistError(host)
	accessorPostAlbumErr := NewAccessorMockPostAlbumError(host)
	accessorPostSongErr := NewAccessorMockPostSongError(host)

	loader := NewLoader(scanner, accessor)

	loaderScanErr := NewLoader(scannerScanErr, accessor)

	loaderPostArtistErr := NewLoader(scanner, accessorPostArtistErr)
	loaderPostAlbumErr := NewLoader(scanner, accessorPostAlbumErr)
	loaderPostSongErr := NewLoader(scanner, accessorPostSongErr)

	err := loader.Load()
	assert.NoError(t, err)

	err = loaderScanErr.Load()
	assert.Error(t, err)

	err = loaderPostArtistErr.Load()
	assert.Error(t, err)

	err = loaderPostAlbumErr.Load()
	assert.Error(t, err)

	err = loaderPostSongErr.Load()
	assert.Error(t, err)
}

func TestToArtist(t *testing.T) {
	name := "sample"

	named := &taglib.Tags{Artist: name}
	named_expect := &models.Artist{
		Name:      name,
		Biography: "",
	}

	unnamed := &taglib.Tags{}
	unnamed_expect := &models.Artist{
		Name:      "Unknown Artist",
		Biography: "",
	}
	resp := toArtist(named)
	assert.Equal(t, resp, named_expect)

	resp = toArtist(unnamed)
	assert.Equal(t, resp, unnamed_expect)
}

func TestToAlbum(t *testing.T) {
	title := "sample"

	named := &taglib.Tags{Album: title}
	named_expect := &models.Album{
		ArtistID: 0,
		Title:    title,
	}

	unnamed := &taglib.Tags{}
	unnamed_expect := &models.Album{
		ArtistID: 0,
		Title:    "Unknown Album",
	}
	resp := toAlbum(named)
	assert.Equal(t, resp, named_expect)

	resp = toAlbum(unnamed)
	assert.Equal(t, resp, unnamed_expect)
}

func TestToSong(t *testing.T) {
	title := "sample"
	genre := "sample"
	year := 2000
	track := 1

	named := &taglib.Tags{
		Title: title,
		Genre: genre,
		Year:  year,
		Track: track,
	}
	named_expect := &models.Song{
		ArtistID: 0,
		AlbumID:  0,
		Title:    title,
		Genre:    genre,
		Year:     year,
		Track:    track,
		Disc:     0,
		Path:     "",
	}

	unnamed := &taglib.Tags{}
	unnamed_expect := &models.Song{
		ArtistID: 0,
		AlbumID:  0,
		Title:    "Unknown Title",
		Genre:    "Unknown Genre",
		Year:     0,
		Track:    0,
		Disc:     0,
		Path:     "",
	}
	resp := toSong(named)
	assert.Equal(t, resp, named_expect)

	resp = toSong(unnamed)
	assert.Equal(t, resp, unnamed_expect)
}
