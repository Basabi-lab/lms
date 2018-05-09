package usecases

import (
	"github.com/Basabi-lab/lms/domains/models"
)

func TestAlbumAllResult() *AlbumAllUsecaseResult {
	album := models.TestAlbumData()
	albums := []*models.Album{album, album}
	return NewAlbumAllUsecaseResult(albums)
}

func TestAlbumGetByIDResult() *AlbumGetByIDUsecaseResult {
	album := models.TestAlbumData()
	return NewAlbumGetByIDUsecaseResult(album)
}

func TestAlbumCreateResult() *AlbumCreateUsecaseResult {
	id := uint(0)
	return NewAlbumCreateUsecaseResult(id)
}

func TestSongAllResult() *SongAllUsecaseResult {
	song := models.TestSongData()
	songs := []*models.Song{song, song}
	return NewSongAllUsecaseResult(songs)
}

func TestSongGetByIDResult() *SongGetByIDUsecaseResult {
	song := models.TestSongData()
	return NewSongGetByIDUsecaseResult(song)
}

func TestSongCreateResult() *SongCreateUsecaseResult {
	id := uint(0)
	return NewSongCreateUsecaseResult(id)
}

func TestArtistAllResult() *ArtistAllUsecaseResult {
	artist := models.TestArtistData()
	artists := []*models.Artist{artist, artist}

	return NewArtistAllUsecaseResult(artists)
}

func TestArtistGetByIDResult() *ArtistGetByIDUsecaseResult {
	artist := models.TestArtistData()
	return NewArtistGetByIDUsecaseResult(artist)
}

func TestArtistCreateResult() *ArtistCreateUsecaseResult {
	id := uint(0)
	return NewArtistCreateUsecaseResult(id)
}
