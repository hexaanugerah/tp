package dto

type CreateBookingRequest struct {
	HotelID int
	RoomID  int
	Nights  int
}
