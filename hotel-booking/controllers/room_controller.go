package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"hotel-booking/models"
)

type RoomData struct {
	Title   string
	Hotel   models.Hotel
	VIP     []models.Room
	Deluxe  []models.Room
	Regular []models.Room
	Rooms   []models.Room
}

func (a *App) BookingPage(w http.ResponseWriter, r *http.Request) {
	hotelID, _ := strconv.Atoi(r.URL.Query().Get("hotel"))
	var selected models.Hotel
	for _, h := range a.DB.Hotels {
		if h.ID == hotelID {
			selected = h
		}
	}
	data := RoomData{Title: "Pilih Kamar", Hotel: selected}
	for _, room := range a.DB.Rooms {
		if room.HotelID != hotelID {
			continue
		}
		switch room.Type {
		case models.RoomVIP:
			data.VIP = append(data.VIP, room)
		case models.RoomDeluxe:
			data.Deluxe = append(data.Deluxe, room)
		default:
			data.Regular = append(data.Regular, room)
		}
	}

	if len(data.Deluxe) > 0 {
	base := data.Deluxe[0]

	newRooms := make([]models.Room, 0, 5)

	for i := 0; i < 5; i++ {
		clone := base
		clone.Name = fmt.Sprintf("DELUXE-%02d", i+1)
		clone.PricePerNight = base.PricePerNight + (i * 35000)
		clone.Stock = 2 + (i % 5)
		clone.Beds = 1 + (i % 3)
		clone.Capacity = 2 + (i % 4)

		newRooms = append(newRooms, clone)
	}

	data.Deluxe = newRooms
}

// 🔥 GABUNG SEMUA KE Rooms
data.Rooms = append(data.Rooms, data.Regular...)
data.Rooms = append(data.Rooms, data.Deluxe...)
data.Rooms = append(data.Rooms, data.VIP...)

	render(w, "booking.html", data)
}
