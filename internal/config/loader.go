package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error  when load env %s", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			DbName:   os.Getenv("DATABASE_NAME"),
		},
		JWT: JWT{
			Key: os.Getenv("JWT_KEY"),
		},
		Email: Email{
			Host:      os.Getenv("SMTP_HOST"),
			Port:      os.Getenv("SMTP_PORT"),
			User:      os.Getenv("SMTP_USER"),
			Password:  os.Getenv("SMTP_PASS"),
			EmailFrom: os.Getenv("EMAIL_FROM"),
		},
		Redis: Redis{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASS"),
		},
		Queue: Redis{
			Addr:     os.Getenv("QUEUE_REDIS_ADDR"),
			Password: os.Getenv("QUEUE_REDIS_PASS"),
		},
		Midtrans: Midtrans{
			Key:    os.Getenv("MIDTRANS_KEY"),
			IsProd: os.Getenv("MIDTRANS_ENV") == "production",
		},
	}
}
