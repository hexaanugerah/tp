package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"hotel-booking/database"
	"hotel-booking/models"
)

type RoomController struct{}

func (rc RoomController) ListByType(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/hotel/")
	path = strings.TrimSuffix(path, "/rooms")
	hotelID, _ := strconv.Atoi(strings.Trim(path, "/"))
	hotel := database.DB.Hotels[hotelID]
	if hotel == nil {
		http.NotFound(w, r)
		return
	}
	roomType := models.RoomType(strings.TrimSpace(strings.ToLower(r.URL.Query().Get("type"))))
	if roomType == "" {
		roomType = models.RoomTypeVIP
	}
	if roomType != models.RoomTypeVIP && roomType != models.RoomTypeDeluxe && roomType != models.RoomTypeStandard {
		roomType = models.RoomTypeVIP
	}

	rooms := []*models.Room{}
	for _, roomID := range hotel.RoomIDs {
		room := database.DB.Rooms[roomID]
		if room.Type == roomType {
			rooms = append(rooms, room)
		}
	}

	renderTemplate(w, "booking.html", map[string]any{
		"Title":          "Pilih Kamar",
		"Hotel":          hotel,
		"Rooms":          rooms,
		"RoomType":       roomType,
		"CurrentType":    string(roomType),
		"ShowLoginPopup": shouldShowLoginPopup(r),
	})
}
