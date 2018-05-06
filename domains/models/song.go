package models

import "github.com/jinzhu/gorm"

type Song struct {
	gorm.Model
	Album    *Album  `gorm:"foreignkey:AlbumID";json:album`
	AlbumID  uint    `json:album_id`
	Artist   *Artist `gorm:"foreignkey:ArtistID";json:artist`
	ArtistID uint    `json:artist_id`
	Title    string  `json:title`
	Track    int     `json:track`
	AlbumNum int     `json:album_num`
	dir      string  `json:dir`
}
