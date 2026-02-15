package controllers

import (
	"net/http"

	"hotel-booking/helpers"
)

func SendPromoEmail(w http.ResponseWriter, r *http.Request) {
	helpers.SendEmail("guest@hotel.com", "Promo Mingguan", "Diskon hingga 30% untuk weekend!")
	_, _ = w.Write([]byte("promo email queued (simulasi)"))
}
