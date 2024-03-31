package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig ServerConfig `json:"server"`
}

type ServerConfig struct {
	Port             time.Duration `json:"port"`
	ReadTimeout      time.Duration `json:"readTimeout"`
	WriteTimeout     time.Duration `json:"writeTimeout"`
	GracefulShutdown time.Duration `json:"gracefulShutdown"`
}

type LoggingConfig struct {
	Level       int    `json:"level"`
	Encoding    string `json:"encoding"`
	Development bool   `json:"development"`
}

type MetricsConfig struct {
	Namespace string `json:"namespace"`
	Subsystem string `json:"subsystem"`
}

func LoadConfig(configPath string) (*Config, error) {
	var conf Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.systemdapi")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("No config found at %s or %s", configPath, "$HOME/.systemapi/")
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal("Unable to unmarshal json into config")
	}
	return &conf, nil
}
