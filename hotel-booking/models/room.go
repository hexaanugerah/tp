package models

type Room struct {
	ID         int
	HotelID    int
	Name       string
	Type       string
	Price      int
	Stock      int
	Capacity   int
	Beds       int
	Facilities []string
	ImageURL   string
}
