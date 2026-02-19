package controllers

import "net/http"

type WebhookController struct{}

func (wc WebhookController) Midtrans(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
