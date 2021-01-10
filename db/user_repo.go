package db

import "gorm.io/gorm"

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		conn: db,
	}
}

func (r *UserRepository) Create(user *User) error {
	result := r.conn.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
