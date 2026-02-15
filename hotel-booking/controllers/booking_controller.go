package controllers

import (
	"net/http"
	"strconv"
	"time"

	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/models"
)

func BookingPage(w http.ResponseWriter, r *http.Request) {
	hotelID, _ := strconv.Atoi(r.URL.Query().Get("hotel_id"))
	roomID, _ := strconv.Atoi(r.URL.Query().Get("room_id"))

	database.AppDB.RLock()
	defer database.AppDB.RUnlock()

	var hotel *models.Hotel
	var room *models.Room
	for _, h := range database.AppDB.Hotels {
		if h.ID == hotelID {
			v := h
			hotel = &v
			break
		}
	}
	for _, rm := range database.AppDB.Rooms {
		if rm.ID == roomID {
			v := rm
			room = &v
			break
		}
	}

	render(w, "booking.html", ViewData{"Request": r, "HotelID": hotelID, "RoomID": roomID, "Hotel": hotel, "Room": room})
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	user := getCurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	hotelID, _ := strconv.Atoi(r.FormValue("hotel_id"))
	roomID, _ := strconv.Atoi(r.FormValue("room_id"))
	checkIn, _ := time.Parse("2006-01-02", r.FormValue("check_in"))
	checkOut, _ := time.Parse("2006-01-02", r.FormValue("check_out"))

	total := 0
	database.AppDB.RLock()
	for _, room := range database.AppDB.Rooms {
		if room.ID == roomID {
			nights := int(checkOut.Sub(checkIn).Hours() / 24)
			if nights < 1 {
				nights = 1
			}
			total = room.Price * nights
			break
		}
	}
	database.AppDB.RUnlock()

	charge := helpers.CreateMidtransCharge(total)
	booking := models.Booking{UserID: user.ID, HotelID: hotelID, RoomID: roomID, CheckIn: checkIn, CheckOut: checkOut, TotalPrice: total, PaymentID: charge.PaymentID, PaymentURL: charge.RedirectURL, PaymentStat: charge.Status, CreatedAt: time.Now()}

	database.AppDB.Lock()
	booking.ID = database.AppDB.NextBookingID
	database.AppDB.NextBookingID++
	database.AppDB.Bookings = append(database.AppDB.Bookings, booking)
	database.AppDB.Unlock()

	helpers.SendEmail(user.Email, "Booking berhasil dibuat", "Silakan selesaikan pembayaran Anda di halaman payment.")
	http.Redirect(w, r, charge.RedirectURL, http.StatusSeeOther)
}
