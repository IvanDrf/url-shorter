package config

import (
	"log"
	"os"
	"url-shorter/internal/errs"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	LoggerLevel string
}

func InitCFG() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(errs.InvalidEnv())
	}

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),

		LoggerLevel: os.Getenv("LOGGER_LEVEL"),
	}
}
