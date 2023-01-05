package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Redis    Redis    `json:"redis"`
	Mongo    Mongo    `json:"mongo"`
	Mysql    MySQL    `json:"mysql"`
	Rabbitmq RabbitMQ `json:"rabbitmq"`
	Metric   Metric   `json:"metric"`
	Grpc     GRPC     `json:"grpc"`
	Server   Server   `json:"server"`
}

func New(filePath string) (Config, error) {
	
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	
	var c Config
	err = json.Unmarshal(content, &c)
	if err != nil {
		return Config{}, err
	}
	
	return c, nil
}
