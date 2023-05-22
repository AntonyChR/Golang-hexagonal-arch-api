package lib

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DBHost           string   `toml:"DB_HOST"`
	DBPort           string   `toml:"DB_PORT"`
	DBUser           string   `toml:"DB_USER"`
	DBPassword       string   `toml:"DB_PASSWORD"`
	DBName           string   `toml:"DB_NAME"`
	AllowedMethods   []string `toml:"ALLOWED_METHODS"`
	AllowedOrigins   []string `toml:"ALLOWED_ORIGINS"`
	ServerPort       string   `toml:"PORT"`
	AllowCredentials bool     `toml:"ALLOW_CREDENTIALS"`
}

func ReadConfig(filename string) (*Config, error) {
	config := &Config{}

	content, err := os.ReadFile(filename)

	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(content), config)

	if err != nil {
		return config, err
	}

	return config, nil
}
