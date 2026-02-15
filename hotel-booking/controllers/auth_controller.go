package controllers

import (
	"net/http"

	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/models"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	render(w, "login.html", ViewData{"Request": r})
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	render(w, "register.html", ViewData{"Request": r})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	database.AppDB.RLock()
	defer database.AppDB.RUnlock()
	for _, u := range database.AppDB.Users {
		if u.Email == email && helpers.CheckPassword(password, u.Password) {
			token := helpers.GenerateToken(u.ID, u.Role)
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: token, Path: "/", HttpOnly: true})
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/login?err=1", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	newUser := models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: helpers.HashPassword(r.FormValue("password")),
		Role:     "user",
	}

	database.AppDB.Lock()
	newUser.ID = database.AppDB.NextUserID
	database.AppDB.NextUserID++
	database.AppDB.Users = append(database.AppDB.Users, newUser)
	database.AppDB.Unlock()

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
