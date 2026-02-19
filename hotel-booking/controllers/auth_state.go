package controllers

import "net/http"

func shouldShowLoginPopup(r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	return err != nil || cookie.Value == ""
}
