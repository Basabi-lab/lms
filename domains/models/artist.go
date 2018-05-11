package models

type Artist struct {
	Model
	Album     []*Album `json:"album"`
	Name      string   `json:"title"`
	Biography string   `json:"biography"`
}
