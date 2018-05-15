package loader

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Basabi-lab/lms/domains/models"
)

type AccessorExt interface {
	GetArtistByName(name string) (*models.Artist, error)
	GetAlbumByTitle(title string) ([]*models.Album, error)
	GetSongByTitle(title string) ([]*models.Song, error)
	PostArtist(artist *models.Artist) (*models.Artist, error)
	PostAlbum(album *models.Album) (*models.Album, error)
	PostSong(song *models.Song) (*models.Song, error)
}

type accessor struct {
	Host string
}

func NewAccessor(host string) AccessorExt {
	return &accessor{
		Host: host,
	}
}

func execute(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func (a *accessor) LastUpdate() (time.Time, error) {
	return time.Time{}, nil
}

func (a *accessor) GetArtistByName(name string) (*models.Artist, error) {
	values := url.Values{}
	values.Add("name", name)
	url := a.Host + "/api/artist" + "?" + values.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := execute(resp)
	if err != nil {
		return nil, err
	}

	var artist *models.Artist
	err = json.Unmarshal(body, &artist)

	return artist, err
}

func (a *accessor) GetAlbumByTitle(title string) ([]*models.Album, error) {
	values := url.Values{}
	values.Add("title", title)
	url := a.Host + "/api/album" + "?" + values.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := execute(resp)
	if err != nil {
		return nil, err
	}

	var albums []*models.Album
	err = json.Unmarshal(body, &albums)

	return albums, err
}

func (a *accessor) GetSongByTitle(title string) ([]*models.Song, error) {
	values := url.Values{}
	values.Add("title", title)
	url := a.Host + "/api/song" + "?" + values.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := execute(resp)
	if err != nil {
		return nil, err
	}

	var songs []*models.Song
	err = json.Unmarshal(body, &songs)

	return songs, err
}

func (a *accessor) PostArtist(artist *models.Artist) (*models.Artist, error) {
	url := a.Host + "/api/artist"
	ctype := "application/json"
	content, _ := json.Marshal(artist)

	client := &http.Client{}

	resp, _ := client.Post(
		url,
		ctype,
		strings.NewReader(string(content)),
	)
	defer resp.Body.Close()

	body, err := execute(resp)
	if err != nil {
		return nil, err
	}
	var artistResp *models.Artist
	err = json.Unmarshal(body, &artistResp)

	return artistResp, nil
}

func (a *accessor) PostAlbum(album *models.Album) (*models.Album, error) {
	url := a.Host + "/api/album"
	ctype := "application/json"
	content, _ := json.Marshal(album)

	client := &http.Client{}

	resp, _ := client.Post(
		url,
		ctype,
		strings.NewReader(string(content)),
	)
	defer resp.Body.Close()

	body, err := execute(resp)
	if err != nil {
		return nil, err
	}
	var albumResp *models.Album
	err = json.Unmarshal(body, &albumResp)

	return albumResp, nil
}

func (a *accessor) PostSong(song *models.Song) (*models.Song, error) {
	url := a.Host + "/api/song"
	ctype := "application/json"
	content, _ := json.Marshal(song)

	client := &http.Client{}

	resp, _ := client.Post(
		url,
		ctype,
		strings.NewReader(string(content)),
	)
	defer resp.Body.Close()

	body, err := execute(resp)
	if err != nil {
		return nil, err
	}
	var songResp *models.Song
	err = json.Unmarshal(body, &songResp)

	return songResp, nil
}
