package loader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var defPath = "../tests/song"

func TestScan(t *testing.T) {
	scanner := NewScanner(defPath)

	songList, err := scanner.Scan()
	assert.NoError(t, err)

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
	assert.Equal(t, expect, songList)
}

func TestScanNotUpdated(t *testing.T) {
	// TODO: run set_timestamp_song.sh
	scanner := NewScanner(defPath)

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	lastUpdateUTC := time.Date(2018, time.May, 12, 12-9, 0, 0, 0, time.UTC)

	lastUpdateJST := lastUpdateUTC.In(jst)

	songList, err := scanner.ScanNotUpdated(lastUpdateJST)
	assert.NoError(t, err)

	expect := []string{
		"../tests/song/not_updated1.mp3",
		"../tests/song/not_updated2.mp3",
	}
	assert.Equal(t, expect, songList)
}
