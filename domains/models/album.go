package models

import "github.com/jinzhu/gorm"

type Album struct {
	gorm.Model
	Song   Song   `json:song`
	Title  string `json:title`
	Artist string `json:artist`
	Genre  string `json:genre`
	Year   int    `json:year`
}
