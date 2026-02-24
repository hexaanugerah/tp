package repositories

import "hotel-booking/internal/models"

type BookingRepository struct{ Bookings []models.Booking }

func (r *BookingRepository) Create(b models.Booking) models.Booking {
	b.ID = len(r.Bookings) + 1
	r.Bookings = append(r.Bookings, b)
	return b
}
