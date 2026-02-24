package config

type Config struct{ Port string }

func Load() Config { return Config{Port: GetEnv("PORT", "8080")} }
