package service

import (
	"log"
	"moon-street/internal/dao"
	"moon-street/internal/di"
	"moon-street/internal/model"
	"moon-street/internal/util"
	"reflect"
)

type UserServiceImpl struct {
}

const ComponentName = "serviceComponent"

var cache = make(map[string]string)

func init() {
	di.Dependencies[ComponentName] = []string{dao.ComponentName}
	di.Factories[ComponentName] = reflect.ValueOf(newUserServiceImpl)
}

func newUserServiceImpl() *UserServiceImpl { //injection
	serv := &UserServiceImpl{}
	return serv
}

func (s *UserServiceImpl) Save(user model.User) (int64, error) {
	instance := di.InstancesInjection[dao.ComponentName].(dao.UserRepo)
	//special treat password
	ePassword := util.EncryptWithSalt(user.Password)
	user.Password = ePassword
	return instance.Save(user)
}

func (s *UserServiceImpl) Check(name string, password string) (bool, error) {
	if hit, ok := cache[name]; ok {
		if hit == password {
			return true, nil
		} else {
			return false, nil
		}
	}
	instance := di.InstancesInjection[dao.ComponentName].(dao.UserRepo)
	retUser, err := instance.GetByName(name)
	if err != nil {
		log.Printf("error when check! %v", err)
		return false, err
	}
	ePassword := util.EncryptWithSalt(password)
	if retUser.Password != ePassword {
		log.Printf("cannot login, not match!")
		return false, nil
	}
	cache[name] = password
	return true, nil
}
