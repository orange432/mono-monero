package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	config := InitConfig("../../config.json")

	fmt.Println(config.SESSION_TIME)
	if config.DSN == "" || config.PORT == "" || config.REDIS_PORT == "" {
		t.Errorf("Failed to initialize variables!")
	}
}
