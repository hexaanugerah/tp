package main

import (
	"log"
	"net/http"

	"hotel-booking/config"
	"hotel-booking/cron"
	"hotel-booking/database"
	"hotel-booking/routes"
)

func main() {
	cfg := config.Load()
	database.Init()
	cron.StartBookingReminderJob()

	mux := routes.SetupRoutes()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("server started at http://localhost:%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
