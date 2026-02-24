package middleware

import "net/http"

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := r.Cookie("role")
		if err != nil || role.Value != "admin" {
			http.Error(w, "admin only", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
