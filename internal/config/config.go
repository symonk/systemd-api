package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig  ServerConfig  `mapstructure:"server"`
	LoggingConfig LoggingConfig `mapstructure:"logging"`
	MetricsConfig MetricsConfig `mapstructure:"metrics"`
}

type ServerConfig struct {
	Port             time.Duration `mapstructure:"port"`
	ReadTimeout      time.Duration `mapstructure:"readTimeout"`
	WriteTimeout     time.Duration `mapstructure:"writeTimeout"`
	GracefulShutdown time.Duration `mapstructure:"gracefulShutdown"`
}

type LoggingConfig struct {
	Level       int    `mapstructure:"level"`
	Encoding    string `mapstructure:"encoding"`
	Development bool   `mapstructure:"development"`
}

type MetricsConfig struct {
	Namespace string `mapstructure:"namespace"`
	Subsystem string `mapstructure:"subsystem"`
}

func LoadConfig(configPath string) (*Config, error) {
	var conf Config
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("No config found at %s or %s, %v", configPath, "$HOME/.systemapi/config.yaml", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal("Unable to unmarshal mapstructure into config")
	}
	return &conf, nil
}
