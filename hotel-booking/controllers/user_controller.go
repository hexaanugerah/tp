package controllers

import "net/http"

type UserController struct{}

func (uc UserController) Profile(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "profile.html", map[string]any{"Title": "Profil"})
}
