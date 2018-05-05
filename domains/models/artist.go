package models

import "github.com/jinzhu/gorm"

type Artist struct {
	gorm.Model
	Album     *Album `gorm:"foreignkey:AlbumID";json:album`
	AlbumID   uint   `json:album_id`
	Name      string `json:title`
	Biography string `json:biography`
}
