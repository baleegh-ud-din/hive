package config

import (
	"os"

	"github.com/baleegh-ud-din/hive/utils"
	"github.com/joho/godotenv"
)

var logger = utils.NewLogger()
var Cfg = LoadConfig()

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoadConfig() *Config {
	err := godotenv.Load("./.env")
	logger.Info("🛠️  Loading Config from env file...")
	if err != nil {
		logger.Warning("⚠️  env file not found, relying on local environment variables")
	}

	return &Config{
		AppPort:    getEnv("APPPORT", "8443"),
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
	}
}
