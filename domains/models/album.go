package models

type Album struct {
	Model
	Songs    []*Song `json:"songs"`
	Artist   *Artist `gorm:"foreignkey:ArtistID" json:"artist"`
	ArtistID uint    `json:"artist_id"`
	Title    string  `json:"title"`
}
