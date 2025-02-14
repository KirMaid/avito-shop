package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App      App      `mapstructure:"app"`
	HTTP     HTTP     `mapstructure:"http"`
	Postgres Postgres `mapstructure:"postgres"`
	Redis    Redis    `mapstructure:"redis"`
	Auth     Auth     `mapstructure:"auth"`
	Logger   Logger   `mapstructure:"logger"`
}

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type HTTP struct {
	Port string `mapstructure:"port"`
}

type Postgres struct {
	URL     string `mapstructure:"url"`
	PoolMax int    `mapstructure:"pool_max"`
}

type Redis struct {
	Port int `mapstructure:"port"`
}

type Auth struct {
	HashSalt   string `mapstructure:"hash_salt"`
	SigningKey string `mapstructure:"signing_key"`
	TokenTTL   string `mapstructure:"token_ttl"`
}

type Logger struct {
	LogLevel string `mapstructure:"log_level"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv()

	viper.BindEnv("http.port", "SERVER_PORT")
	viper.BindEnv("postgres.url", "DATABASE_URL")
	viper.BindEnv("redis.port", "REDIS_PORT")

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	log.Printf("App Name: %s", cfg.App.Name)
	log.Printf("App Version: %s", cfg.App.Version)
	log.Printf("HTTP Port: %s", cfg.HTTP.Port)
	log.Printf("Postgres URL: %s", cfg.Postgres.URL)
	log.Printf("Postgres Pool Max: %d", cfg.Postgres.PoolMax)
	log.Printf("Redis Port: %d", cfg.Redis.Port)
	log.Printf("Hash Salt: %s", cfg.Auth.HashSalt)
	log.Printf("Signing Key: %s", cfg.Auth.SigningKey)
	log.Printf("Token TTL: %s", cfg.Auth.TokenTTL)
	log.Printf("Log Level: %s", cfg.Logger.LogLevel)

	return cfg, nil
}
