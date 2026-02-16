package routes

import (
	"net/http"
	"strings"

	"hotel-booking/config"
	"hotel-booking/controllers"
	"hotel-booking/middleware"
)

func Register(cfg config.AppConfig) *http.ServeMux {
	mux := http.NewServeMux()
	auth := controllers.AuthController{Cfg: cfg}
	hotel := controllers.HotelController{}
	room := controllers.RoomController{}
	booking := controllers.BookingController{}
	payment := controllers.PaymentController{}
	admin := controllers.AdminController{}
	dash := controllers.DashboardController{}
	reco := controllers.RecommendationController{}
	webhook := controllers.WebhookController{}
	email := controllers.EmailController{}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", hotel.Home)
	mux.HandleFunc("/recommendations", reco.TopHotels)

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			auth.Login(w, r)
			return
		}
		auth.LoginPage(w, r)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			auth.Register(w, r)
			return
		}
		auth.RegisterPage(w, r)
	})
	mux.HandleFunc("/logout", auth.Logout)

	mux.HandleFunc("/hotel/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/rooms") {
			room.ListByType(w, r)
			return
		}
		hotel.Detail(w, r)
	})

	mux.HandleFunc("/booking/create", middleware.AuthRequired(cfg, booking.Create))
	mux.HandleFunc("/booking/cancel/", middleware.AuthRequired(cfg, booking.Cancel))
	mux.HandleFunc("/dashboard", middleware.AuthRequired(cfg, dash.Index))

	mux.HandleFunc("/payment", middleware.AuthRequired(cfg, payment.Page))
	mux.HandleFunc("/payment/pay", middleware.AuthRequired(cfg, payment.Pay))

	mux.HandleFunc("/webhook/midtrans", webhook.Midtrans)
	mux.HandleFunc("/admin", middleware.AuthRequired(cfg, admin.Dashboard))
	mux.HandleFunc("/email/test", email.SendTest)

	return mux
}
