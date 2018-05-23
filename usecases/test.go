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
	album := models.TestAlbumData()
	return NewAlbumCreateUsecaseResult(album)
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
	song := models.TestSongData()
	return NewSongCreateUsecaseResult(song)
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
	artist := models.TestArtistData()
	return NewArtistCreateUsecaseResult(artist)
}

func TestArtistClearUsecaseResult() *ArtistClearUsecaseResult {
	var err error = nil
	return NewArtistClearUsecaseResult(err)
}
