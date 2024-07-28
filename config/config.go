package config

import (
	"encoding/json"
	"os"

	"gorm.io/gorm"
)

type App struct {
	DB        *gorm.DB
	Port      string `json:"port"`
	SecretKey string `json:"secret_key"`
}

func LoadConfig(filePath string) (*App, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config App
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
