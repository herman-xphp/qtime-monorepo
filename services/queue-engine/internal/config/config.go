package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App  AppConfig
	HTTP HTTPConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type HTTPConfig struct {
	Port string
}

func LoadConfig() *Config {
	viper.SetDefault("APP_NAME", "Q-Time Queue Engine")
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("HTTP_PORT", ":3000")

	// 1. Read from .env file (if exists) used for local dev
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// 2. Read from ENV variables (override .env) used for Docker/Prod
	viper.AutomaticEnv()

	config := &Config{
		App: AppConfig{
			Name: viper.GetString("APP_NAME"),
			Env:  viper.GetString("APP_ENV"),
		},
		HTTP: HTTPConfig{
			Port: viper.GetString("HTTP_PORT"),
		},
	}

	log.Printf("Config loaded: Env=%s Port=%s", config.App.Env, config.HTTP.Port)
	return config
}
