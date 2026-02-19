package models

import "time"

type BookingStatus string

type PaymentStatus string

const (
	BookingPending  BookingStatus = "pending"
	BookingPaid     BookingStatus = "paid"
	BookingCanceled BookingStatus = "canceled"

	PaymentUnpaid PaymentStatus = "unpaid"
	PaymentPaid   PaymentStatus = "paid"
)

type Booking struct {
	ID            int
	BookingCode   string
	UserID        int
	HotelID       int
	RoomID        int
	CheckInDate   time.Time
	CheckOutDate  time.Time
	GuestName     string
	GuestEmail    string
	TotalPrice    int
	Status        BookingStatus
	PaymentStatus PaymentStatus
	PaymentToken  string
	CreatedAt     time.Time
}
