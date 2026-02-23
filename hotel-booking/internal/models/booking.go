package models

type Booking struct {
	ID      int
	UserID  int
	HotelID int
	RoomID  int
	Nights  int
	Total   int
	Status  string
}
