package service

import "firstgo/internal/dao"

type SService struct {
	Dao *dao.UserDao
}

func NewSService(dao *dao.UserDao) *UserService {
	return &UserService{Dao: dao}
}
