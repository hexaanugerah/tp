package controllers

import (
	"net/http"

	"hotel-booking/database"
)

func PaymentPage(w http.ResponseWriter, r *http.Request) {
	trx := r.URL.Query().Get("trx")
	database.AppDB.RLock()
	defer database.AppDB.RUnlock()
	for _, b := range database.AppDB.Bookings {
		if b.PaymentID == trx {
			render(w, "payment.html", ViewData{"Request": r, "Booking": b})
			return
		}
	}
	render(w, "payment.html", ViewData{"Request": r})
}
