package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBsslmode  string

	LoggerLevel string
}

func InitCFG() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't read env file")
	}

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBsslmode:  os.Getenv("DB_MODE"),

		LoggerLevel: os.Getenv("LOGGER_LEVEL"),
	}
}
