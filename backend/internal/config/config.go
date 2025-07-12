package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

var EnvConfig *Config

type Config struct {
	GO_ENV   string `env:"GO_ENV" envDefault:"development"`
	App_Port string `env:"APP_PORT" envDefault:"8080"`
	App_Mode string `env:"APP_MODE" envDefault:"offline"`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"user"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"password"`
	DBName     string `env:"DB_NAME" envDefault:"postgres"`

	JWTAccessSecret  string `env:"JWT_ACCESS_SECRET" envDefault:"your_access_secret"`
	JWTRefreshSecret string `env:"JWT_REFRESH_SECRET" envDefault:"your_refresh_secret"`
	JWTAccessExpiry  string `env:"JWT_ACCESS_EXPIRY" envDefault:"1h"`
	JWTRefreshExpiry string `env:"JWT_REFRESH_EXPIRY" envDefault:"7d"`
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		GO_ENV:   os.Getenv("GO_ENV"),
		App_Port: os.Getenv("APP_PORT"),
		App_Mode: os.Getenv("APP_MODE"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		JWTAccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
		JWTRefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		JWTAccessExpiry:  os.Getenv("JWT_ACCESS_EXPIRY"),
		JWTRefreshExpiry: os.Getenv("JWT_REFRESH_EXPIRY"),
	}

	EnvConfig = config

	fmt.Println("Configuration loaded:")
	fmt.Printf("GO_ENV: %s\n", config.GO_ENV)
	fmt.Printf("APP_MODE: %s\n", config.App_Mode)

	return config, nil
}
