package models

import "time"

type Booking struct {
	ID          int
	UserID      int
	HotelID     int
	RoomID      int
	CheckIn     time.Time
	CheckOut    time.Time
	TotalPrice  int
	PaymentID   string
	PaymentURL  string
	PaymentStat string
	CreatedAt   time.Time
}
