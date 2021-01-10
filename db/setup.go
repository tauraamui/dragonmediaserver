package db

import "gorm.io/gorm"

func Setup(dbConn *gorm.DB) error {
	dbConn.AutoMigrate(&User{})
	userRepository := NewUserRepository(dbConn)
	return userRepository.Create(&User{
		Username: "admin",
		AuthHash: []byte("admin"),
	})
}
