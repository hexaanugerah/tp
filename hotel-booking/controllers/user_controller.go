package controllers

import "net/http"

func Profile(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", ViewData{"Request": r, "Message": "Profil user masih sederhana untuk demo."})
}
