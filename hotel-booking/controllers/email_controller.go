package controllers

import (
	"fmt"
	"net/http"

	"hotel-booking/helpers"
)

type EmailController struct{}

func (ec EmailController) SendTest(w http.ResponseWriter, r *http.Request) {
	_ = helpers.SendEmail("guest@example.com", "Tes Email", "Booking Anda berhasil")
	_, _ = w.Write([]byte(fmt.Sprintf("email sent at %s", r.URL.Path)))
}
