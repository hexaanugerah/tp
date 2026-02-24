package handlers

import (
	"net/http"
	"strconv"

	"hotel-booking/internal/services"
)

type BookingHandler struct {
	App            *App
	BookingService *services.BookingService
}

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	hotelID, _ := strconv.Atoi(r.FormValue("hotel_id"))
	roomID, _ := strconv.Atoi(r.FormValue("room_id"))
	nights, _ := strconv.Atoi(r.FormValue("nights"))
	b, err := h.BookingService.Create(1, hotelID, roomID, nights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.App.Render(w, "booking/history.html", map[string]any{"Booking": b})
}
