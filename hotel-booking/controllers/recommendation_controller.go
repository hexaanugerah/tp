package controllers

import (
	"net/http"
	"sort"

	"hotel-booking/database"
	"hotel-booking/models"
)

type RecommendationController struct{}

func (rc RecommendationController) TopHotels(w http.ResponseWriter, r *http.Request) {
	items := make([]*models.Hotel, 0, len(database.DB.Hotels))
	for _, h := range database.DB.Hotels {
		items = append(items, h)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Rating > items[j].Rating })
	if len(items) > 6 {
		items = items[:6]
	}
	renderTemplate(w, "recommendation.html", map[string]any{"Title": "Rekomendasi", "Hotels": items})
}
