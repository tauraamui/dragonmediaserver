package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("dragon.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
