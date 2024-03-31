package config

import "time"

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
	return nil, nil
}
