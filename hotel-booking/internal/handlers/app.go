package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type App struct {
	HotelHandler   *HotelHandler
	AuthHandler    *AuthHandler
	BookingHandler *BookingHandler
	PaymentHandler *PaymentHandler
	AdminHandler   *AdminHandler
	ViewRoot       string
}

func (a *App) Render(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles(
		filepath.Join(a.ViewRoot, "layouts/header.html"),
		filepath.Join(a.ViewRoot, "layouts/navbar.html"),
		filepath.Join(a.ViewRoot, "layouts/footer.html"),
		filepath.Join(a.ViewRoot, page),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "header", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
