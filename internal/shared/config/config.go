package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

// Load membaca konfigurasi aplikasi dari .env atau Environment Variable
func Load() (*Config, error) {

	// Tidak error jika .env tidak ada.
	// Production biasanya menggunakan Environment Variable langsung.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	redisDB, err := getIntEnv("REDIS_DB", 0)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App: AppConfig{
			Name: getEnvOrDefault("APP_NAME", "Restaurant API"),
			Env:  getEnvOrDefault("APP_ENV", "development"),
			Port: getEnvOrDefault("APP_PORT", "8080"),
		},

		Database: DatabaseConfig{
			Host:     mustGetEnv("DB_HOST"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
			Name:     mustGetEnv("DB_NAME"),
			SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
		},

		Redis: RedisConfig{
			Host:     mustGetEnv("REDIS_HOST"),
			Port:     getEnvOrDefault("REDIS_PORT", "6379"),
			Password: getEnvOrDefault("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
	}

	return cfg, nil
}

// ----------------------------------------------------------------------
// Helper
// ----------------------------------------------------------------------

func mustGetEnv(key string) string {

	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	return value
}

func getEnvOrDefault(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func getIntEnv(key string, defaultValue int) (int, error) {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue, nil
	}

	number, err := strconv.Atoi(value)

	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", key)
	}

	return number, nil
}