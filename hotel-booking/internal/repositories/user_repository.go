package repositories

import "hotel-booking/internal/models"

type UserRepository struct{ Users []models.User }

func (r *UserRepository) FindByEmail(email string) *models.User {
	for i := range r.Users {
		if r.Users[i].Email == email {
			return &r.Users[i]
		}
	}
	return nil
}
