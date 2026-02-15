package controllers

import (
	"net/http"

	"hotel-booking/database"
)

func Recommendations(w http.ResponseWriter, r *http.Request) {
	database.AppDB.RLock()
	hotels := database.AppDB.Hotels
	database.AppDB.RUnlock()
	if len(hotels) > 4 {
		hotels = hotels[:4]
	}
	render(w, "index.html", ViewData{"Request": r, "Hotels": hotels, "Message": "Rekomendasi terbaik untuk kamu minggu ini"})
}
