package config

import "os"

type AppConfig struct {
	Port              string
	JWTSecret         string
	MidtransServerKey string
	MidtransClientKey string
}

func Load() AppConfig {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super-secret-local"
	}

	return AppConfig{
		Port:              port,
		JWTSecret:         jwtSecret,
		MidtransServerKey: getOrDefault("MIDTRANS_SERVER_KEY", "SB-Mid-server-xxx"),
		MidtransClientKey: getOrDefault("MIDTRANS_CLIENT_KEY", "SB-Mid-client-xxx"),
	}
}

func getOrDefault(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
