package db

import (
	"log"

	"gorm.io/gorm"
)

func Setup(dbConn *gorm.DB, stdlog *log.Logger) error {
	stdlog.Print("Running DB auto migrations... ")
	err := dbConn.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	stdlog.Println("Finished...")

	stdlog.Print("Creating default admin user... ")

	userRepository := NewUserRepository(dbConn)
	err = userRepository.Create(&User{
		Username: "admin",
		AuthHash: []byte("admin"),
		Admin:    true,
	})

	stdlog.Println("Finished...")
	return err
}
