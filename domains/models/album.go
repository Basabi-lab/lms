package models

import "github.com/jinzhu/gorm"

type Album struct {
	gorm.Model
	Songs    []*Song `json:songs`
	Artist   *Artist `gorm:"foreignkey:ArtistID";json:artist`
	ArtistID uint    `json:artist_id`
	Title    string  `json:title`
}
