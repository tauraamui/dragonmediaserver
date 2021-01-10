package db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string
	Username string
	Email    string
	AuthHash []byte
	Admin    bool
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	hash, err := plainToHash(u.AuthHash)
	if err != nil {
		return err
	}

	u.AuthHash = hash
	return nil
}

func plainToHash(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	return hash, err
}
