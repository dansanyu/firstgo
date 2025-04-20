package service

import "firstgo/internal/dao"

type UserService struct {
	Dao *dao.UserDao
}

func NewService(dao *dao.UserDao) *UserService {
	return &UserService{Dao: dao}
}
