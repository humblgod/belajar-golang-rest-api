package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost			string
	DBPort 			string
	DBUser			string
	DBPassword		string
	DBName			string
}


var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBHost: getEnv("DB_HOST", "127.0.0.1"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "tangerang12"),
		DBName: getEnv("DB_NAME", "belajargolangrestapi"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

