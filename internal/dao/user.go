package dao

import "gorm.io/gorm"

type UserDao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}
