package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DSN              string `json:"dsn"`
	PORT             string `json:"port"`
	REDIS_PORT       string `json:"redis_port"`
	SESSION_TIME     int    `json:"session_time"`     // session time in minutes
	TEMPLATE_REFRESH int64  `json:"template_refresh"` //Seconds per template refresh
	SECRET           string `json:"secret"`           // Used for encryption
}

var DSN string
var PORT string
var REDIS_PORT string
var SESSION_TIME int
var TEMPLATE_REFRESH int64
var SECRET string

// InitConfig intializes the configuration of this go app
func InitConfig(fileName string) Config {
	var config Config

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		return config
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal(err)
		return config
	}

	// Doing it this way because it is easier to handle outside config.go
	DSN = config.DSN
	PORT = config.PORT
	REDIS_PORT = config.REDIS_PORT
	SESSION_TIME = config.SESSION_TIME
	TEMPLATE_REFRESH = config.TEMPLATE_REFRESH
	SECRET = config.SECRET
	return config
}
