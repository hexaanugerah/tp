package controllers

import "net/http"

type DashboardController struct{}

func (dc DashboardController) Index(w http.ResponseWriter, r *http.Request) {
	BookingController{}.ListMine(w, r)
}
