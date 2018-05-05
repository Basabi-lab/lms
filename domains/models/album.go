package models

import "github.com/jinzhu/gorm"

type Album struct {
	gorm.Model
	Songs  []*Song `json:songs`
	Artist *Artist `json:artist`
	Title  string  `json:title`
	Genre  string  `json:genre`
	Year   int     `json:year`
}
