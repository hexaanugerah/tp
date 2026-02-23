package main

import (
	"log"
	"net/http"

	"hotel-booking/config"
	"hotel-booking/database/seed"
	"hotel-booking/internal/handlers"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"hotel-booking/internal/routes"
	"hotel-booking/internal/services"
	"hotel-booking/pkg/security"
)

func main() {
	cfg := config.Load()
	hotels, rooms := seed.SeedHotelsAndRooms()

	userRepo := &repositories.UserRepository{Users: []models.User{
		{ID: 1, Name: "Admin", Email: "admin@hotel.local", Password: security.HashPassword("admin123"), Role: "admin"},
		{ID: 2, Name: "Guest", Email: "user@hotel.local", Password: security.HashPassword("user123"), Role: "pengguna"},
		{ID: 3, Name: "Staff", Email: "staff@hotel.local", Password: security.HashPassword("staff123"), Role: "staff"},
	}}
	hotelRepo := &repositories.HotelRepository{Hotels: hotels}
	roomRepo := &repositories.RoomRepository{Rooms: rooms}
	bookingRepo := &repositories.BookingRepository{}
	paymentRepo := &repositories.PaymentRepository{}

	app := &handlers.App{ViewRoot: "views"}
	app.HotelHandler = &handlers.HotelHandler{App: app, HotelService: &services.HotelService{HotelRepo: hotelRepo, RoomRepo: roomRepo}}
	app.AuthHandler = &handlers.AuthHandler{App: app, AuthService: &services.AuthService{UserRepo: userRepo}}
	app.BookingHandler = &handlers.BookingHandler{App: app, BookingService: &services.BookingService{BookingRepo: bookingRepo, RoomRepo: roomRepo}}
	app.PaymentHandler = &handlers.PaymentHandler{App: app, PaymentService: &services.PaymentService{PaymentRepo: paymentRepo}}
	app.AdminHandler = &handlers.AdminHandler{App: app}

	mux := http.NewServeMux()
	routes.Register(mux, app)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("server run on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
