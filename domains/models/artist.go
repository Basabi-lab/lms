package models

type Artist struct {
	Model
	Album     []*Album `json:"album"`
	Name      string   `json:"name"`
	Biography string   `json:"biography"`
}
