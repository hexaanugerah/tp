package models

type Room struct {
	ID            int
	HotelID       int
	Type          string
	Name          string
	PricePerNight int
	Beds          int
	Capacity      int
	Facilities    []string
	FloorPlanURL  string
}
