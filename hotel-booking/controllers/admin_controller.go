package controllers

import (
	"net/http"

	"hotel-booking/database"
)

type AdminController struct{}

func (ac AdminController) Dashboard(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "admin_dashboard.html", map[string]any{
		"Title":    "Admin Dashboard",
		"Users":    database.DB.Users,
		"Bookings": database.DB.Bookings,
		"Hotels":   database.DB.Hotels,
	})
}
