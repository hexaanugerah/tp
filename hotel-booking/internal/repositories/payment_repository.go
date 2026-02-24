package repositories

import "hotel-booking/internal/models"

type PaymentRepository struct{ Payments []models.Payment }

func (r *PaymentRepository) Create(p models.Payment) models.Payment {
	p.ID = len(r.Payments) + 1
	r.Payments = append(r.Payments, p)
	return p
}
