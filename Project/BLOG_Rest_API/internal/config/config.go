package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configure struct {
	DB_USER        string
	DB_PASSWORD    string
	DB_NAME        string
	DB_PORT        string
	API_PORT       string
	HOST           string
	JWT_SECRET     string
	JWT_EXPIRES_IN string
	OTP_EXPIRES    string
}

func Load() (Configure, error) {
	err := godotenv.Load()
	if err != nil {
		return Configure{}, err
	}

	return Configure{
		DB_USER:        os.Getenv("DB_USER"),
		DB_PASSWORD:    os.Getenv("DB_PASSWORD"),
		DB_NAME:        os.Getenv("DB_NAME"),
		DB_PORT:        os.Getenv("DB_PORT"),
		API_PORT:       os.Getenv("API_PORT"),
		HOST:           os.Getenv("HOST"),
		JWT_SECRET:     os.Getenv("JWT_SECRET"),
		JWT_EXPIRES_IN: os.Getenv("JWT_EXPIRES_IN"),
		OTP_EXPIRES:    os.Getenv("OTP_EXPIRES"),
	}, nil
}
