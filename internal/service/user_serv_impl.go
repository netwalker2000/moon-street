package service

import (
	"log"
	"moon-street/internal/dao"
	"moon-street/internal/model"
	"moon-street/internal/util"
)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl { //injection
	serv := &UserServiceImpl{}
	return serv
}

func (s *UserServiceImpl) Save(user model.User) (int64, error) {
	instance := dao.GetDatabaseInstance()
	//special treat password
	ePassword := util.EncryptWithSalt(user.Password)
	user.Password = ePassword
	return instance.Save(user)
}

func (s *UserServiceImpl) Check(name string, password string) (bool, error) {
	instance := dao.GetDatabaseInstance()
	retUser, err := instance.GetByName(name)
	if err != nil {
		log.Fatal("error when check!")
		return false, err
	}
	ePassword := util.EncryptWithSalt(password)
	if retUser.Password != ePassword {
		log.Fatal("cannot login, not match!")
		return false, nil
	}
	return true, nil
}
