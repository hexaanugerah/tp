package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"hotel-booking/database"
	"hotel-booking/models"
)

type HotelController struct{}

type HotelComment struct {
	HotelName string
	City      string
	Rating    float64
	Message   string
}

func (h HotelController) Home(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	city := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("city")))
	hotels := make([]*models.Hotel, 0, len(database.DB.Hotels))
	for _, hotel := range database.DB.Hotels {
		if query != "" && !strings.Contains(strings.ToLower(hotel.Name), query) {
			continue
		}
		if city != "" && !strings.Contains(strings.ToLower(hotel.City), city) {
			continue
		}
		hotels = append(hotels, hotel)
	}
	sort.Slice(hotels, func(i, j int) bool { return hotels[i].Rating > hotels[j].Rating })

	feedback := []string{
		"Lokasi strategis, dekat kuliner dan transportasi.",
		"Kamarnya nyaman, bersih, dan proses check-in cepat.",
		"Pemandangan bagus, cocok untuk liburan keluarga.",
		"Staff ramah, fasilitas hotel lengkap dan terawat.",
		"Value for money, cocok untuk staycation weekend.",
	}
	comments := make([]HotelComment, 0, len(hotels))
	for i, hotel := range hotels {
		comments = append(comments, HotelComment{
			HotelName: hotel.Name,
			City:      hotel.City,
			Rating:    hotel.Rating,
			Message:   fmt.Sprintf("%s", feedback[i%len(feedback)]),
		})
	}

	renderTemplate(w, "index.html", map[string]any{
		"Title":    "Beranda",
		"Hotels":   hotels,
		"Comments": comments,
		"Query":    query,
		"City":     city,
	})
}

func (h HotelController) Detail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/hotel/"))
	hotel := database.DB.Hotels[id]
	if hotel == nil {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "hotel_detail.html", map[string]any{"Title": hotel.Name, "Hotel": hotel})
}
