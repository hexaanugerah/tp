package main

import (
	"fmt"
	"net/http"

	"hotel-booking/config"
	"hotel-booking/cron"
	"hotel-booking/database"
	"hotel-booking/routes"
)

func main() {
	cfg := config.Load()
	database.Init()
	cron.StartReminderJob()

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: routes.Register(cfg),
	}

	fmt.Printf("%s running at http://localhost:%s\n", cfg.AppName, cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
