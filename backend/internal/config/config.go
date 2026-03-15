package config

import (
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
	return &Config{
		Env:            getEnv("ENV", "development"),
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://lincle:lincle_db_2026@192.168.50.222:15433/lincle?sslmode=disable"),
		RedisURL:       getEnv("REDIS_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", "dev-secret-change-me"),
		JWTAccessTTL:   parseDuration(getEnv("JWT_ACCESS_TTL", "15m")),
		JWTRefreshTTL:  parseDuration(getEnv("JWT_REFRESH_TTL", "720h")),
		UploadDir:      getEnv("UPLOAD_DIR", "./uploads"),
		MaxUploadSize:  10 << 20, // 10MB
		AllowedOrigins:  []string{"http://localhost:3000", "http://localhost:8081"},
		GoogleClientIDs: parseCSV(getEnv("GOOGLE_CLIENT_IDS", "")),
	}
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
