package service

import (
	"moon-street/internal/dao"
	"moon-street/internal/model"
)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl { //injection
	serv := &UserServiceImpl{}
	return serv
}

func (s *UserServiceImpl) Save(user model.User) (int64, error) {
	instance := dao.GetDatabaseInstance()
	return instance.Save(user)
}
