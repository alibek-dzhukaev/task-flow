package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	ServerPort string
	ServerEnv  string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Redis
	RedisHost string
	RedisPort string

	// JWT
	JWTSecret             string
	JWTAccessTokenExpiry  string
	JWTRefreshTokenExpiry string

	// Rate Limiter
	RateLimitRequests int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerEnv:  getEnv("SERVER_ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "task_flow"),

		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),

		JWTSecret:             getEnv("JWT_SECRET", "secret"),
		JWTAccessTokenExpiry:  getEnv("JWT_ACCESS_TOKEN_EXPIRY", "15m"),
		JWTRefreshTokenExpiry: getEnv("JWT_REFRESH_TOKEN_EXPIRY", "7d"),

		RateLimitRequests: getEnvInt("RATE_LIMIT_REQUESTS", 100),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func (c *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func (c *Config) IsProd() bool {
	return c.ServerEnv == "production"
}

func (c *Config) validate() error {
	if c.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	if c.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	return nil
}

// getEnv - analogue to process.env.KEY || defaultValue in JS
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if n, err := strconv.Atoi(value); err == nil {
			return n
		}
	}
	return defaultValue
}
