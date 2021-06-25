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

func (s *UserServiceImpl) Save(model.User) (int64, error) {
	instance := dao.NewDatabaseInstance()
	instance.Save()
	return 1, nil
}
