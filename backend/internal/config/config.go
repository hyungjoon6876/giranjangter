package config

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Env                  string
	Port                 string
	DatabaseURL          string
	RedisURL             string
	JWTSecret            string
	JWTAccessTTL         time.Duration
	JWTRefreshTTL        time.Duration
	UploadDir            string
	MaxUploadSize        int64
	AllowedOrigins       []string
	GoogleClientIDs      []string // Web + iOS + Android client IDs for token verification
}

func Load() *Config {
	cfg := &Config{
		Env:             getEnv("ENV", "production"),
		Port:            getEnv("PORT", "8080"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		RedisURL:        getEnv("REDIS_URL", ""),
		JWTSecret:       getEnv("JWT_SECRET", "dev-secret-change-me"),
		JWTAccessTTL:    parseDuration(getEnv("JWT_ACCESS_TTL", "15m")),
		JWTRefreshTTL:   parseDuration(getEnv("JWT_REFRESH_TTL", "720h")),
		UploadDir:       getEnv("UPLOAD_DIR", "./uploads"),
		MaxUploadSize:   10 << 20, // 10MB
		AllowedOrigins: func() []string {
			if origins := parseCSV(getEnv("ALLOWED_ORIGINS", "")); len(origins) > 0 {
				return origins
			}
			return []string{"http://localhost:3000", "http://localhost:8081"}
		}(),
		GoogleClientIDs: parseCSV(getEnv("GOOGLE_CLIENT_IDS", "")),
	}
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	if cfg.JWTSecret == "dev-secret-change-me" || len(cfg.JWTSecret) < 16 {
		if !cfg.IsDev() {
			log.Fatal("JWT_SECRET must be set to a secure value (min 16 chars) in production")
		}
	}
	return cfg
}

func (c *Config) IsDev() bool {
	return c.Env == "development"
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseCSV(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			result = append(result, t)
		}
	}
	return result
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 15 * time.Minute
	}
	return d
}
