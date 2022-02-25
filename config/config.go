package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DSN          string `json:"dsn"`
	PORT         string `json:"port"`
	REDIS_PORT   string `json:"redis_port"`
	SESSION_TIME int    `json:"session_time"`
}

var DSN string
var PORT string
var REDIS_PORT string
var SESSION_TIME int

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
	return config
}
