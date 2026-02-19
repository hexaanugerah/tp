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

type CityDeal struct {
	Name     string
	Iconic   string
	Price    string
	ImageURL string
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

	cityDeals := []CityDeal{
		{Name: "Bandung", Iconic: "Gedung Sate", Price: "Rp 47.000", ImageURL: "https://images.unsplash.com/photo-1558005530-a7958896ec60?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Yogyakarta", Iconic: "Candi Prambanan", Price: "Rp 40.000", ImageURL: "https://images.unsplash.com/photo-1532186651327-6ac23687d189?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Bali", Iconic: "Pura Ulun Danu", Price: "Rp 56.000", ImageURL: "https://images.unsplash.com/photo-1537953773345-d172ccf13cf1?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Singapore", Iconic: "Merlion Park", Price: "Rp 239.000", ImageURL: "https://images.unsplash.com/photo-1525625293386-3f8f99389edd?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Kuala Lumpur", Iconic: "Menara Kembar Petronas", Price: "Rp 67.000", ImageURL: "https://images.unsplash.com/photo-1596422846543-75c6fc197f07?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Bangkok", Iconic: "Wat Arun", Price: "Rp 47.000", ImageURL: "https://images.unsplash.com/photo-1508009603885-50cf7c579365?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Penang", Iconic: "Kek Lok Si Temple", Price: "Rp 98.000", ImageURL: "https://images.unsplash.com/photo-1527631746610-bca00a040d60?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Seoul", Iconic: "N Seoul Tower", Price: "Rp 280.000", ImageURL: "https://images.unsplash.com/photo-1538485399081-7191377e8241?auto=format&fit=crop&w=1200&q=80"},
		{Name: "Tokyo", Iconic: "Tokyo Tower", Price: "Rp 268.000", ImageURL: "https://images.unsplash.com/photo-1540959733332-eab4deabeeaf?auto=format&fit=crop&w=1200&q=80"},
	}

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
		"Title":          "Beranda",
		"Hotels":         hotels,
		"Comments":       comments,
		"CityDeals":      cityDeals,
		"Query":          query,
		"City":           city,
		"ShowLoginPopup": shouldShowLoginPopup(r),
	})
}

func (h HotelController) Detail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/hotel/"))
	hotel := database.DB.Hotels[id]
	if hotel == nil {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "hotel_detail.html", map[string]any{
		"Title":          hotel.Name,
		"Hotel":          hotel,
		"ShowLoginPopup": shouldShowLoginPopup(r),
	})
}
