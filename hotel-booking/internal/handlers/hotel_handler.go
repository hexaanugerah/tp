package handlers

import (
	"net/http"
	"strconv"

	"hotel-booking/internal/services"
)

type HotelHandler struct {
	App          *App
	HotelService *services.HotelService
}

type CityCard struct {
	Name      string
	ImageURL  string
	PriceFrom string
}

var cityImageMap = map[string]string{
	"Jakarta":     "https://picsum.photos/seed/monas-jakarta/900/400",
	"Bandung":     "https://picsum.photos/seed/gedungsate-bandung/900/400",
	"Yogyakarta":  "https://picsum.photos/seed/tugu-yogya/900/400",
	"Bali":        "https://picsum.photos/seed/pura-bali/900/400",
	"Surabaya":    "https://picsum.photos/seed/suroboyo/900/400",
	"Medan":       "https://picsum.photos/seed/istana-maimun/900/400",
	"Semarang":    "https://picsum.photos/seed/lawang-sewu/900/400",
	"Malang":      "https://picsum.photos/seed/bromo-malang/900/400",
	"Solo":        "https://picsum.photos/seed/keraton-solo/900/400",
	"Bogor":       "https://picsum.photos/seed/kebun-raya-bogor/900/400",
	"Makassar":    "https://picsum.photos/seed/pantai-losari/900/400",
	"Balikpapan":  "https://picsum.photos/seed/balikpapan/900/400",
	"Lombok":      "https://picsum.photos/seed/mandalika-lombok/900/400",
	"Labuan Bajo": "https://picsum.photos/seed/labuanbajo-komodo/900/400",
	"Palembang":   "https://picsum.photos/seed/ampera-palembang/900/400",
	"Padang":      "https://picsum.photos/seed/jam-gadang-padang/900/400",
	"Manado":      "https://picsum.photos/seed/bunaken-manado/900/400",
	"Pontianak":   "https://picsum.photos/seed/tugu-khatulistiwa/900/400",
	"Banjarmasin": "https://picsum.photos/seed/pasar-terapung/900/400",
	"Jayapura":    "https://picsum.photos/seed/jayapura-papua/900/400",
}

func (h *HotelHandler) Home(w http.ResponseWriter, r *http.Request) {
	cities := h.HotelService.Cities()
	cityCards := make([]CityCard, 0, len(cities))
	for i, c := range cities {
		cityCards = append(cityCards, CityCard{Name: c, ImageURL: cityImageMap[c], PriceFrom: strconv.Itoa(45000 + (i * 3000))})
	}
	h.App.Render(w, "hotels/index.html", map[string]any{"Cities": cities, "CityCards": cityCards})
}

func (h *HotelHandler) CityHotels(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("name")
	if city == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	h.App.Render(w, "hotels/city_hotels.html", map[string]any{"City": city, "Hotels": h.HotelService.HotelsByCity(city)})
}

func (h *HotelHandler) Detail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	hotel := h.HotelService.HotelByID(id)
	if hotel == nil {
		http.NotFound(w, r)
		return
	}
	h.App.Render(w, "hotels/detail.html", map[string]any{"Hotel": hotel})
}

func (h *HotelHandler) Rooms(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("hotel_id"))
	roomType := r.URL.Query().Get("type")
	hotel := h.HotelService.HotelByID(id)
	if hotel == nil {
		http.NotFound(w, r)
		return
	}
	h.App.Render(w, "booking/booking.html", map[string]any{"Hotel": hotel, "Rooms": h.HotelService.Rooms(id, roomType), "Type": roomType})
}
