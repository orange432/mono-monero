package database

import (
	"github.com/orange432/mono-monero/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Opens a connection to the database
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	return db, err
}
