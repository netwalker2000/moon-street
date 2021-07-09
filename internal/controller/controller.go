package controller

import (
	"errors"
	"log"
	"moon-street/common"
	"moon-street/internal/di"
	"moon-street/internal/model"
	"moon-street/internal/service"
	"moon-street/internal/util"
	"reflect"
)

func Route(rCmd common.RpcData) error { //injection ;  error raise
	maps := make(map[string]reflect.Value)
	maps["register"] = reflect.ValueOf(register)
	maps["login"] = reflect.ValueOf(login)
	fn, ok := maps[rCmd.Name]
	if !ok {
		log.Printf("Unknown cmd, end this conn! [%s]", rCmd.Name)
		return errors.New("unknown cmd")
	}
	inArgs := make([]reflect.Value, len(rCmd.Args))
	for i := range rCmd.Args {
		inArgs[i] = reflect.ValueOf(rCmd.Args[i])
	}
	fn.Call(inArgs)
	return nil
}

func register(username, password, email string) {
	// todo: check params
	user := model.User{
		Name:     username,
		Password: password,
		Email:    email,
	}
	userService := di.InstancesInjection[service.ComponentName].(service.UserService)
	userService.Save(user)
}

func login(username, password string) bool {
	var go_id = util.GoID()
	// todo: check params
	userService := di.InstancesInjection[service.ComponentName].(service.UserService)
	isPass, err := userService.Check(username, password)
	if err != nil {
		log.Printf("error when deal with account check! [%d] %v", go_id, err)
		return false
	} else {
		if isPass {
			log.Printf("Login success with user %s! [%d]", username, go_id)
		} else {
			log.Printf("Cannot Login with user %s! [%d]", username, go_id)
		}
		return isPass
	}
}
