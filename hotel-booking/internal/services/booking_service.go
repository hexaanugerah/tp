package services

import (
	"errors"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type BookingService struct {
	BookingRepo *repositories.BookingRepository
	RoomRepo    *repositories.RoomRepository
}

func (s *BookingService) Create(userID, hotelID, roomID, nights int) (models.Booking, error) {
	room := s.RoomRepo.FindByID(roomID)
	if room == nil {
		return models.Booking{}, errors.New("room not found")
	}
	b := models.Booking{UserID: userID, HotelID: hotelID, RoomID: roomID, Nights: nights, Total: room.PricePerNight * nights, Status: "pending"}
	return s.BookingRepo.Create(b), nil
}
