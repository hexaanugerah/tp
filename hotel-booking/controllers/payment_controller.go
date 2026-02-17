package controllers

import (
	"net/http"
	"strconv"

	"hotel-booking/database"
	"hotel-booking/models"
)

type PaymentController struct{}

func (pc PaymentController) Page(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("booking_id"))
	booking := database.DB.Bookings[id]
	if booking == nil {
		http.NotFound(w, r)
		return
	}
	hotel := database.DB.Hotels[booking.HotelID]
	room := database.DB.Rooms[booking.RoomID]
	renderTemplate(w, "payment.html", map[string]any{"Title": "Pembayaran", "Booking": booking, "Hotel": hotel, "Room": room})
}

func (pc PaymentController) Pay(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("booking_id"))
	booking := database.DB.Bookings[id]
	if booking == nil {
		http.NotFound(w, r)
		return
	}
	booking.PaymentStatus = models.PaymentPaid
	booking.Status = models.BookingPaid
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
