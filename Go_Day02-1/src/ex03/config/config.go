package config

import (
	"log"
)

type Config struct {
}

var cfg Config

func LoadConfig() {
	log.Println("Loading configuration...")

}

func GetConfig() *Config {
	return &cfg
}
