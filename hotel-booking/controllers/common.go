package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/models"
)

type ViewData map[string]any

func render(w http.ResponseWriter, tmpl string, data ViewData) {
	if data == nil {
		data = ViewData{}
	}
	data["AppName"] = "TravelBook"
	if req, ok := data["Request"].(*http.Request); ok {
		data["CurrentUser"] = getCurrentUser(req)
	} else {
		data["CurrentUser"] = nil
	}

	base := filepath.Join("views", "layout.html")
	page := filepath.Join("views", tmpl)
	t, err := template.ParseFiles(base, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getCurrentUser(r *http.Request) *models.User {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil
	}
	id, _, ok := helpers.ParseToken(cookie.Value)
	if !ok {
		return nil
	}
	database.AppDB.RLock()
	defer database.AppDB.RUnlock()
	for _, u := range database.AppDB.Users {
		if u.ID == id {
			user := u
			return &user
		}
	}
	return nil
}
