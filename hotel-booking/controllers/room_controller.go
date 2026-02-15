package controllers

import "net/http"

func RoomsByHotel(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
