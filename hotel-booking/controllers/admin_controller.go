package controllers

import (
	"net/http"

	"hotel-booking/database"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	database.AppDB.RLock()
	data := ViewData{"Request": r, "TotalHotels": len(database.AppDB.Hotels), "TotalRooms": len(database.AppDB.Rooms), "TotalBookings": len(database.AppDB.Bookings)}
	database.AppDB.RUnlock()
	render(w, "admin_dashboard.html", data)
}
