package routes

import (
	"net/http"

	"hotel-booking/internal/handlers"
	"hotel-booking/internal/middleware"
)

func Register(mux *http.ServeMux, app *handlers.App) {
	mux.HandleFunc("/", app.HotelHandler.Home)
	mux.HandleFunc("/hotel", app.HotelHandler.Detail)
	mux.HandleFunc("/city-hotels", app.HotelHandler.CityHotels)
	mux.HandleFunc("/rooms", app.HotelHandler.Rooms)
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			app.AuthHandler.Login(w, r)
			return
		}
		app.AuthHandler.LoginPage(w, r)
	})
	mux.HandleFunc("/register", app.AuthHandler.RegisterPage)
	mux.HandleFunc("/booking/create", middleware.RequireAuth(app.BookingHandler.Create))
	mux.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			app.PaymentHandler.Pay(w, r)
			return
		}
		app.PaymentHandler.PayPage(w, r)
	})
	mux.HandleFunc("/admin", middleware.RequireAuth(middleware.RequireAdmin(app.AdminHandler.Dashboard)))
}
