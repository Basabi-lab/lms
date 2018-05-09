package presenters

import (
	"github.com/gin-gonic/gin"
)

func TestAlbumCreatePresenterResult() *AlbumCreatePresenterResult {
	return NewAlbumCreatePresenterResult(&gin.H{"message": "success", "id": uint(0)})
}

func TestSongCreatePresenterResult() *SongCreatePresenterResult {
	return NewSongCreatePresenterResult(&gin.H{"message": "success", "id": uint(0)})
}

func TestArtistCreatePresenterResult() *ArtistCreatePresenterResult {
	return NewArtistCreatePresenterResult(&gin.H{"message": "success", "id": uint(0)})
}
