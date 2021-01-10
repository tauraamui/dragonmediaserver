package db

import "gorm.io/gorm"

func Setup(dbConn *gorm.DB) error {
	err := dbConn.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	userRepository := NewUserRepository(dbConn)
	return userRepository.Create(&User{
		Username: "admin",
		AuthHash: []byte("admin"),
		Admin:    true,
	})
}
