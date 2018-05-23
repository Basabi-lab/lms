package usecases

import (
	"github.com/Basabi-lab/lms/domains/models"
)

func TestAlbumAllUsecaseResult() *AlbumAllUsecaseResult {
	album := models.TestAlbumData()
	albums := []*models.Album{album, album}
	return NewAlbumAllUsecaseResult(albums)
}

func TestAlbumGetByIDUsecaseResult() *AlbumGetByIDUsecaseResult {
	album := models.TestAlbumData()
	return NewAlbumGetByIDUsecaseResult(album)
}

func TestAlbumCreateUsecaseResult() *AlbumCreateUsecaseResult {
	id := uint(0)
	return NewAlbumCreateUsecaseResult(id)
}

func TestAlbumClearUsecaseResult() *AlbumClearUsecaseResult {
	var err error = nil
	return NewAlbumClearUsecaseResult(err)
}

func TestSongAllUsecaseResult() *SongAllUsecaseResult {
	song := models.TestSongData()
	songs := []*models.Song{song, song}
	return NewSongAllUsecaseResult(songs)
}

func TestSongGetByIDUsecaseResult() *SongGetByIDUsecaseResult {
	song := models.TestSongData()
	return NewSongGetByIDUsecaseResult(song)
}

func TestSongCreateUsecaseResult() *SongCreateUsecaseResult {
	id := uint(0)
	return NewSongCreateUsecaseResult(id)
}

func TestSongClearUsecaseResult() *SongClearUsecaseResult {
	var err error = nil
	return NewSongClearUsecaseResult(err)
}

func TestArtistAllUsecaseResult() *ArtistAllUsecaseResult {
	artist := models.TestArtistData()
	artists := []*models.Artist{artist, artist}

	return NewArtistAllUsecaseResult(artists)
}

func TestArtistGetByIDUsecaseResult() *ArtistGetByIDUsecaseResult {
	artist := models.TestArtistData()
	return NewArtistGetByIDUsecaseResult(artist)
}

func TestArtistCreateUsecaseResult() *ArtistCreateUsecaseResult {
	id := uint(0)
	return NewArtistCreateUsecaseResult(id)
}

func TestArtistClearUsecaseResult() *ArtistClearUsecaseResult {
	var err error = nil
	return NewArtistClearUsecaseResult(err)
}
