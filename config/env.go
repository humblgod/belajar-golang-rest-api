package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost				string
	DBPort 				string
	DBUser				string
	DBPassword			string
	DBName				string
	JWTExpirationTime 	int64
	JWTSecret 			string
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
		JWTExpirationTime: getEnvsInteger("JWT_EXP", 3600*24*7), // a week time
		JWTSecret: getEnv("JWT_SECRET", "something-secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}


func getEnvsInteger(key string, fallback int64) int64{ 
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64) 
			if err != nil {
				return fallback
			}
		return i
	}
	return fallback
}

