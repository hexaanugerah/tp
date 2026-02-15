package routes

import (
	"net/http"

	"hotel-booking/controllers"
	"hotel-booking/middleware"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/hotel", controllers.HotelDetail)
	mux.HandleFunc("/recommendations", controllers.Recommendations)

	mux.HandleFunc("/login", methodSwitch(controllers.LoginPage, controllers.Login))
	mux.HandleFunc("/register", methodSwitch(controllers.RegisterPage, controllers.Register))
	mux.HandleFunc("/logout", controllers.Logout)

	mux.HandleFunc("/booking", middleware.RequireLogin(controllers.BookingPage))
	mux.HandleFunc("/booking/create", middleware.RequireLogin(controllers.CreateBooking))
	mux.HandleFunc("/payment", middleware.RequireLogin(controllers.PaymentPage))
	mux.HandleFunc("/payment/webhook", controllers.MidtransWebhook)

	mux.HandleFunc("/admin/dashboard", middleware.RequireLogin(controllers.AdminDashboard))
	mux.HandleFunc("/dashboard", middleware.RequireLogin(controllers.Dashboard))
	mux.HandleFunc("/email/promo", controllers.SendPromoEmail)

	return mux
}

func methodSwitch(getHandler, postHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		getHandler(w, r)
	}
}
