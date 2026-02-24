package handlers

import (
	"net/http"

	"hotel-booking/internal/services"
)

type AuthHandler struct {
	App         *App
	AuthService *services.AuthService
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.App.Render(w, "auth/login.html", nil)
}
func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	h.App.Render(w, "auth/register.html", nil)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	token, user, err := h.AuthService.Login(r.FormValue("email"), r.FormValue("password"), r.FormValue("role"))
	if err != nil {
		http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "token", Value: token, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "role", Value: user.Role, Path: "/"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
