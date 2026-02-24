package handlers

import "net/http"

type AdminHandler struct{ App *App }

func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	h.App.Render(w, "admin/dashboard.html", nil)
}
