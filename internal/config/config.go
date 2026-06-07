package config

import "os"

type Config struct {
	DB        string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	RedisHost string
	RedisPort string
	BaseUrl   string
	Port      string
}

func Load() *Config {
	return &Config{
		DB:        os.Getenv("POSTGRES_DATABASE"),
		DBHost:    os.Getenv("POSTGRES_HOST"),
		DBPort:    os.Getenv("POSTGRES_PORT"),
		DBUser:    os.Getenv("POSTGRES_USER"),
		DBPass:    os.Getenv("POSTGRES_PASSWORD"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		BaseUrl:   os.Getenv("BASE_URL"),
		Port:      os.Getenv("PORT"),
	}
}
