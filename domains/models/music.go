package models

import "github.com/jinzhu/gorm"

type Music struct {
	gorm.Model
	Title  string `json:title`
	Artist string `json:artist`
	Track  int    `json:track`
	CDNum  int    `json:cd_num`
	dir    string `json:dir`
	length int    `json:length`
}
