package loader

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"github.com/Basabi-lab/lms/domains/models"
)

type testAccessor struct{}

var host string = "http://localhost:8080"

func TestClearAll(t *testing.T) {
	acc := NewAccessor(host)

	artistClear := host + "/artist/clear"
	albumClear := host + "/album/clear"
	songClear := host + "/song/clear"

	data := map[string]string{"message": "success"}
	json, _ := json.Marshal(data)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PATCH", artistClear,
		httpmock.NewStringResponder(200, string(json)))
	httpmock.RegisterResponder("PATCH", albumClear,
		httpmock.NewStringResponder(200, string(json)))
	httpmock.RegisterResponder("PATCH", songClear,
		httpmock.NewStringResponder(200, string(json)))

	err := acc.ClearAll()
	assert.NoError(t, err)
}

func TestGetArtistByName(t *testing.T) {
	acc := NewAccessor(host)
	name := "Artist name"
	expect := models.TestArtistData()
	json, _ := json.Marshal(expect)

	values := url.Values{}
	values.Add("name", name)
	url := host + "/api/artist" + "?" + values.Encode()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, string(json)))

	artist, err := acc.GetArtistByName(name)
	assert.NoError(t, err)

	assert.Equal(t, expect, artist)
}

func TestGetAlbumByTitle(t *testing.T) {
	acc := NewAccessor(host)
	title := "Album title"
	expect := []*models.Album{models.TestAlbumData()}
	json, _ := json.Marshal(expect)

	values := url.Values{}
	values.Add("title", title)
	url := host + "/api/album" + "?" + values.Encode()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, string(json)))

	albums, err := acc.GetAlbumByTitle(title)
	assert.NoError(t, err)

	assert.Equal(t, expect, albums)
}

func TestGetSongByTitle(t *testing.T) {
	acc := NewAccessor(host)
	title := "Song title"
	expect := []*models.Song{models.TestSongData()}
	json, _ := json.Marshal(expect)

	values := url.Values{}
	values.Add("title", title)
	url := host + "/api/song" + "?" + values.Encode()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, string(json)))

	songs, err := acc.GetSongByTitle(title)
	assert.NoError(t, err)

	assert.Equal(t, expect, songs)
}

func TestPostArtist(t *testing.T) {
	acc := NewAccessor(host)
	artist := models.TestArtistData()
	url := host + "/api/artist"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", url,
		func(req *http.Request) (*http.Response, error) {
			var artist *models.Artist
			body, _ := ioutil.ReadAll(req.Body)
			_ = json.Unmarshal(body, &artist)
			artist.ID = uint(1)

			resp, err := httpmock.NewJsonResponse(200, artist)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, err
		},
	)

	resp, err := acc.PostArtist(artist)
	assert.NoError(t, err)

	assert.NotEqual(t, artist.ID, resp.ID)
	artist.ID = resp.ID
	assert.Equal(t, artist, resp)
}

func TestPostAlbum(t *testing.T) {
	acc := NewAccessor(host)
	album := models.TestAlbumData()
	url := host + "/api/album"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", url,
		func(req *http.Request) (*http.Response, error) {
			var album *models.Album
			body, _ := ioutil.ReadAll(req.Body)
			_ = json.Unmarshal(body, &album)
			album.ID = uint(1)

			resp, err := httpmock.NewJsonResponse(200, album)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, err
		},
	)

	resp, err := acc.PostAlbum(album)
	assert.NoError(t, err)

	assert.NotEqual(t, album.ID, resp.ID)
	album.ID = resp.ID
	assert.Equal(t, album, resp)
}

func TestPostSong(t *testing.T) {
	acc := NewAccessor(host)
	song := models.TestSongData()
	url := host + "/api/song"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", url,
		func(req *http.Request) (*http.Response, error) {
			var song *models.Song
			body, _ := ioutil.ReadAll(req.Body)
			_ = json.Unmarshal(body, &song)
			song.ID = uint(1)

			resp, err := httpmock.NewJsonResponse(200, song)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, err
		},
	)

	resp, err := acc.PostSong(song)
	assert.NoError(t, err)

	assert.NotEqual(t, song.ID, resp.ID)
	song.ID = resp.ID
	assert.Equal(t, song, resp)
}
