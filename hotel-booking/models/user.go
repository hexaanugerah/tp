package models

type UserRole string

const (
	RoleUser       UserRole = "user"
	RoleAdmin      UserRole = "admin"
	RoleHotelStaff UserRole = "hotel_staff"
)

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	Role         UserRole
}
