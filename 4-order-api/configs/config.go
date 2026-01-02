package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db       DbConfig
	Email    EmailConfig
	Password PasswordConfig
	Address  AddressConfig
}

type DbConfig struct {
	Db string
}
type EmailConfig struct {
	Email string
}
type PasswordConfig struct {
	Password string
}
type AddressConfig struct {
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using defaults.")
	}
	return &Config{
		Db: DbConfig{
			Db: os.Getenv("DSN"),
		},
		Email: EmailConfig{
			Email: os.Getenv("EMAIL"),
		},
		Password: PasswordConfig{
			Password: os.Getenv("PASSWORD"),
		},
		Address: AddressConfig{
			Address: os.Getenv("ADDRESS"),
		},
	}
}
