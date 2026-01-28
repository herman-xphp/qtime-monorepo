package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	HTTP     HTTPConfig
	Redis    RedisConfig
	Redpanda RedpandaConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type HTTPConfig struct {
	Port string
}

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type RedpandaConfig struct {
	Brokers []string
	Topic   string
}

func LoadConfig() *Config {
	viper.SetDefault("APP_NAME", "Q-Time Queue Engine")
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("HTTP_PORT", ":3000")
	viper.SetDefault("REDIS_HOST", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("REDPANDA_BROKERS", "localhost:9092")
	viper.SetDefault("REDPANDA_TOPIC", "queue-events")

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
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		Redpanda: RedpandaConfig{
			Brokers: strings.Split(viper.GetString("REDPANDA_BROKERS"), ","),
			Topic:   viper.GetString("REDPANDA_TOPIC"),
		},
	}

	log.Printf("Config loaded: Env=%s Port=%s Redpanda=%v", config.App.Env, config.HTTP.Port, config.Redpanda.Brokers)
	return config
}
