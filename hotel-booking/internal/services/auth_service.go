package services

import (
	"errors"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"hotel-booking/pkg/security"
)

type AuthService struct{ UserRepo *repositories.UserRepository }

func (s *AuthService) Login(email, password, role string) (string, *models.User, error) {
	u := s.UserRepo.FindByEmail(email)
	if u == nil || u.Password != security.HashPassword(password) || (role != "" && u.Role != role) {
		return "", nil, errors.New("invalid credentials")
	}
	return security.GenerateToken(u.Email, u.Role), u, nil
}
