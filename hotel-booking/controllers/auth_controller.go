package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"hotel-booking/config"
	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/models"
)

type AuthController struct {
	Cfg config.AppConfig
}

type googleUserInfo struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	VerifiedEmail bool   `json:"verified_email"`
}

func (a AuthController) oauthConfig() *oauth2.Config {
	if a.Cfg.GoogleClientID == "" || a.Cfg.GoogleClientSecret == "" {
		return nil
	}
	return &oauth2.Config{
		ClientID:     a.Cfg.GoogleClientID,
		ClientSecret: a.Cfg.GoogleClientSecret,
		RedirectURL:  a.Cfg.GoogleRedirectURL,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func (a AuthController) LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", map[string]any{
		"Title":            "Login",
		"GoogleLoginReady": a.oauthConfig() != nil,
	})
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
	roleWanted := models.UserRole(strings.TrimSpace(r.FormValue("role")))
	if roleWanted == "" {
		roleWanted = models.RoleUser
	}
	for _, user := range database.DB.Users {
		valid := helpers.CheckPasswordHash(password, user.PasswordHash) || password == user.PasswordHash
		if user.Email == email && valid {
			if roleWanted != "" && user.Role != roleWanted {
				continue
			}
			token := helpers.GenerateToken(user.ID, a.Cfg.JWTSecret)
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: token, Path: "/", HttpOnly: true})
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (a AuthController) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthCfg := a.oauthConfig()
	if oauthCfg == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	state := "gostay_google_state"
	http.SetCookie(w, &http.Cookie{Name: "oauth_state", Value: state, Path: "/", HttpOnly: true})
	http.Redirect(w, r, oauthCfg.AuthCodeURL(state, oauth2.AccessTypeOffline), http.StatusTemporaryRedirect)
}

func (a AuthController) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthCfg := a.oauthConfig()
	if oauthCfg == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	cookie, err := r.Cookie("oauth_state")
	if err != nil || cookie.Value != r.URL.Query().Get("state") {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	token, err := oauthCfg.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	client := oauthCfg.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	defer resp.Body.Close()

	var gu googleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&gu); err != nil || gu.Email == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var found *models.User
	for _, u := range database.DB.Users {
		if strings.EqualFold(u.Email, gu.Email) {
			found = u
			break
		}
	}
	if found == nil {
		id := database.DB.NextUserID
		found = &models.User{ID: id, Name: gu.Name, Email: gu.Email, PasswordHash: "google-oauth", Role: models.RoleUser}
		database.DB.Users[id] = found
		database.DB.NextUserID++
	}

	session := helpers.GenerateToken(found.ID, a.Cfg.JWTSecret)
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: session, Path: "/", HttpOnly: true})
	http.Redirect(w, r, "/", http.StatusFound)
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusFound)
}
