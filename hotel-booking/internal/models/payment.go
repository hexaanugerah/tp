package models

type Payment struct {
	ID        int
	BookingID int
	Method    string
	Amount    int
	Status    string
}
