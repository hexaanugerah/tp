package controllers

import (
	"net/http"
	"strings"

	"hotel-booking/config"
	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/models"
)

type AuthController struct {
	Cfg config.AppConfig
}

func (a AuthController) LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", map[string]any{"Title": "Login"})
}

func (a AuthController) RegisterPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html", map[string]any{"Title": "Register"})
}

func (a AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))
	password := r.FormValue("password")
	if name == "" || email == "" || password == "" {
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}
	for _, user := range database.DB.Users {
		if user.Email == email {
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}
	}
	hash, _ := helpers.HashPassword(password)
	id := database.DB.NextUserID
	database.DB.Users[id] = &models.User{ID: id, Name: name, Email: email, PasswordHash: hash, Role: models.RoleUser}
	database.DB.NextUserID++
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	email := strings.TrimSpace(r.FormValue("email"))
	password := r.FormValue("password")
	for _, user := range database.DB.Users {
		valid := helpers.CheckPasswordHash(password, user.PasswordHash) || password == user.PasswordHash
		if user.Email == email && valid {
			token := helpers.GenerateToken(user.ID, a.Cfg.JWTSecret)
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: token, Path: "/", HttpOnly: true})
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusFound)
}
