package loader

import (
	"github.com/vbatts/go-taglib/taglib"

	"github.com/Basabi-lab/lms/domains/models"
)

type LoaderExt interface {
	Load() error
}

type loader struct {
	scanner  ScannerExt
	accessor AccessorExt
}

func NewLoader(scanner ScannerExt, accessor AccessorExt) LoaderExt {
	return &loader{
		scanner:  scanner,
		accessor: accessor,
	}
}

func toArtist(info *taglib.Tags) *models.Artist {
	artist := info.Artist
	if artist == "" {
		artist = "Unknown Artist"
	}
	return &models.Artist{
		Name:      artist,
		Biography: "",
	}
}

func toAlbum(info *taglib.Tags) *models.Album {
	album := info.Album
	if album == "" {
		album = "Unknown Album"
	}

	return &models.Album{
		ArtistID: 0,
		Title:    album,
	}
}

func toSong(info *taglib.Tags) *models.Song {
	title := info.Title
	if title == "" {
		title = "Unknown Title"
	}
	genre := info.Genre
	if genre == "" {
		genre = "Unknown Genre"
	}
	year := info.Year
	track := info.Track
	disc := 0

	return &models.Song{
		ArtistID: 0,
		AlbumID:  0,
		Title:    title,
		Genre:    genre,
		Year:     year,
		Track:    track,
		Disc:     disc,
		Path:     "",
	}
}

func (l *loader) Load() error {
	l.accessor.ClearAll()

	songList, err := l.scanner.Scan()
	if err != nil {
		return err
	}

	for _, song := range songList {
		info := taglib.Open(song)
		tags := info.GetTags()

		artist := toArtist(tags)
		artist, err := l.accessor.PostArtist(artist)
		if err != nil {
			return err
		}

		album := toAlbum(tags)
		album.ArtistID = artist.ID
		album, err = l.accessor.PostAlbum(album)
		if err != nil {
			return err
		}

		song := toSong(tags)
		song.ArtistID = artist.ID
		song.AlbumID = album.ID
		_, err = l.accessor.PostSong(song)
		if err != nil {
			return err
		}
	}

	return nil
}
