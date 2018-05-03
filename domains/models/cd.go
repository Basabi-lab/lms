package models

import "github.com/jinzhu/gorm"

type CD struct {
	gorm.Model
	Title  string `json:title`
	Artist string `json:artist`
	Genre  string `json:genre`
	Year   int    `json:year`
	dir    string `json:dir`
}