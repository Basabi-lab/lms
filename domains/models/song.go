package models

type Song struct {
	Model
	Album    *Album  `gorm:"foreignkey:AlbumID" json:"album"`
	AlbumID  uint    `json:"album_id"`
	Artist   *Artist `gorm:"foreignkey:ArtistID" json:"artist"`
	ArtistID uint    `json:"artist_id"`
	Title    string  `json:"title"`
	Genre    string  `json:"genre"`
	Year     int     `json:"year"`
	Track    int     `json:"track"`
	Disc     int     `json:"disc"`
	Path     string  `json:"path"`
}
