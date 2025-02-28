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
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
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
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return cfg, nil
}
