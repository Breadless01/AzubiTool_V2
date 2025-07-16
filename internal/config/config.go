package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey      string
	DBUser      string
	DBPassword  string
	SESSION_KEY string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		APIKey:      os.Getenv("DOTSOURCE_API_KEY"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		SESSION_KEY: os.Getenv("SESSION_KEY"),
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("DOTSOURCE_API_KEY nicht gesetzt")
	}
	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DB_USER nicht gesetzt")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD nicht gesetzt")
	}
	if cfg.SESSION_KEY == "" {
		return nil, fmt.Errorf("SESSION_KEY nicht gesetzt")
	}

	return cfg, nil
}
