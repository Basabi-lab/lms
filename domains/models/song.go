package models

import "github.com/jinzhu/gorm"

type Song struct {
	gorm.Model
	Album    *Album  `gorm:"foreignkey:AlbumID";json:album`
	Artist   *Artist `gorm:"foreignkey:ArtistID";json:artist`
	ArtistID uint64  `json:artist_id`
	Title    string  `json:title`
	Track    int     `json:track`
	AlbumNum int     `json:album_num`
	dir      string  `json:dir`
}
