package controllers

import "net/http"

func Dashboard(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
