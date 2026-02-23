package services

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type PaymentService struct {
	PaymentRepo *repositories.PaymentRepository
}

func (s *PaymentService) Pay(bookingID int, method string, amount int) models.Payment {
	return s.PaymentRepo.Create(models.Payment{BookingID: bookingID, Method: method, Amount: amount, Status: "paid"})
}
