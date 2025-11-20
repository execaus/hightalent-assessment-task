package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

const configPath = "config/config.yaml"
const testConfigPath = "../../config/config.test.yaml"

func LoadConfig() *Config {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config into struct: %v", err)
	}

	cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")

	return &cfg
}

func LoadTestConfig() *Config {
	var cfg Config

	v := viper.New()
	v.SetConfigFile(testConfigPath)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config into struct: %v", err)
	}

	return &cfg
}
