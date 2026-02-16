package middleware

import (
	"context"
	"net/http"

	"hotel-booking/config"
	"hotel-booking/database"
	"hotel-booking/helpers"
	"hotel-booking/models"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthRequired(cfg config.AppConfig, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		uid, err := helpers.ParseToken(cookie.Value, cfg.JWTSecret)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		user := database.DB.Users[uid]
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func CurrentUser(r *http.Request) *models.User {
	if user, ok := r.Context().Value(UserContextKey).(*models.User); ok {
		return user
	}
	return nil
}
