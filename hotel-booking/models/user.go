package models

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	Role         UserRole
}
