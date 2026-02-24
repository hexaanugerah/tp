package handlers

import (
	"net/http"
	"strconv"

	"hotel-booking/internal/services"
)

type PaymentHandler struct {
	App            *App
	PaymentService *services.PaymentService
}

func (h *PaymentHandler) PayPage(w http.ResponseWriter, r *http.Request) {
	h.App.Render(w, "payment/payment.html", nil)
}

func (h *PaymentHandler) Pay(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	bookingID, _ := strconv.Atoi(r.FormValue("booking_id"))
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	payment := h.PaymentService.Pay(bookingID, r.FormValue("method"), amount)
	h.App.Render(w, "payment/payment.html", map[string]any{"Payment": payment, "Success": true})
}
