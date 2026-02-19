package models

type Hotel struct {
	ID             int
	Name           string
	City           string
	Address        string
	Description    string
	Rating         float64
	ImageURL       string
	Amenities      []string
	FloorMapURL    string
	LocationMapURL string
	RoomIDs        []int
}
