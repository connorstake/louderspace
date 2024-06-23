package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DatabaseURL string
	JwtSecret   string
}

func LoadConfig() (*Config, error) {
	// Get the absolute path to the .env file
	rootPath, err := filepath.Abs("../../")
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(rootPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}

	return config, nil
}
