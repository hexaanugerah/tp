package cron

import (
	"fmt"
	"time"

	"hotel-booking/database"
	"hotel-booking/models"
)

func StartReminderJob() {
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for range ticker.C {
			for _, booking := range database.DB.Bookings {
				if booking.PaymentStatus == models.PaymentPaid {
					continue
				}
				if time.Until(booking.CheckInDate) < 48*time.Hour {
					fmt.Printf("[CRON] Reminder pembayaran booking %s\n", booking.BookingCode)
				}
			}
		}
	}()
}
