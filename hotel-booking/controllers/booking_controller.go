package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/middleware"
	"hotel-booking/models"
)

type BookingController struct{}

func (bc BookingController) Create(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	roomID, _ := strconv.Atoi(r.FormValue("room_id"))
	room := database.DB.Rooms[roomID]
	if room == nil || !room.Available {
		http.Error(w, "room unavailable", http.StatusBadRequest)
		return
	}
	checkIn, _ := time.Parse("2006-01-02", r.FormValue("check_in"))
	checkOut, _ := time.Parse("2006-01-02", r.FormValue("check_out"))
	if checkOut.Before(checkIn) || checkOut.Equal(checkIn) {
		http.Error(w, "invalid date range", http.StatusBadRequest)
		return
	}
	nights := int(checkOut.Sub(checkIn).Hours() / 24)
	total := room.PricePerDay * nights

	id := database.DB.NextBookingID
	code := fmt.Sprintf("GST-%05d", id)
	paymentToken := helpers.CreatePaymentToken(code, total)
	booking := &models.Booking{
		ID:            id,
		BookingCode:   code,
		UserID:        user.ID,
		HotelID:       room.HotelID,
		RoomID:        room.ID,
		GuestName:     user.Name,
		GuestEmail:    user.Email,
		CheckInDate:   checkIn,
		CheckOutDate:  checkOut,
		TotalPrice:    total,
		Status:        models.BookingPending,
		PaymentStatus: models.PaymentUnpaid,
		PaymentToken:  paymentToken,
		CreatedAt:     time.Now(),
	}
	database.DB.Bookings[id] = booking
	database.DB.NextBookingID++

	http.Redirect(w, r, "/payment?booking_id="+strconv.Itoa(id), http.StatusFound)
}

func (bc BookingController) ListMine(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r)
	bookings := []*models.Booking{}
	for _, b := range database.DB.Bookings {
		if b.UserID == user.ID {
			bookings = append(bookings, b)
		}
	}
	renderTemplate(w, "dashboard.html", map[string]any{"Title": "Booking Saya", "Bookings": bookings, "Hotels": database.DB.Hotels, "Rooms": database.DB.Rooms})
}

func (bc BookingController) Cancel(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/booking/cancel/"))
	booking := database.DB.Bookings[id]
	if booking != nil {
		booking.Status = models.BookingCanceled
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
