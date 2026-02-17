package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	AppName      string
	Port         string
	JWTSecret    string
	MidtransKey  string
	SMTPFrom     string
	ReminderHour int
}

func Load() AppConfig {
	return AppConfig{
		AppName:      getEnv("APP_NAME", "GoStay - Hotel Booking"),
		Port:         getEnv("PORT", "8080"),
		JWTSecret:    getEnv("JWT_SECRET", "gostay-secret"),
		MidtransKey:  getEnv("MIDTRANS_SERVER_KEY", "dummy-midtrans-key"),
		SMTPFrom:     getEnv("SMTP_FROM", "noreply@gostay.local"),
		ReminderHour: getEnvInt("REMINDER_HOUR", 9),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	raw := os.Getenv(key)
	if raw == "" {
		return fallback
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return value
}
