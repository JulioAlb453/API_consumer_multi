package config

import (
	"os"
	"strings"
)

type Config struct {
	BrokerURL     string
	WebbrowserURL string
	Topic         []string
}

func LoadCongif() Config {
	topics := strings.Split(os.Getenv("MQTT_TOPIC"), ",")
	return Config{
		BrokerURL:     os.Getenv("MQTT_BROKER_URL"),
		WebbrowserURL: os.Getenv("WEBBROWSER_URL"),
		Topic:         topics,
	}
}
