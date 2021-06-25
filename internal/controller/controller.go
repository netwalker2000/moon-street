package controller

import (
	"log"
	"moon-street/common"
	"moon-street/internal/model"
	"moon-street/internal/service"
	"reflect"
)

func Route(rCmd common.RpcData) { //injection ;  error raise
	log.Printf("args %v", rCmd.Args...)
	maps := make(map[string]reflect.Value)
	maps["register"] = reflect.ValueOf(register)
	fn, ok := maps[rCmd.Name]
	if !ok {
		log.Printf("Unknown cmd, end this conn! [%s]", rCmd.Name)
		return
	}
	inArgs := make([]reflect.Value, len(rCmd.Args))
	for i := range rCmd.Args {
		inArgs[i] = reflect.ValueOf(rCmd.Args[i])
	}
	fn.Call(inArgs)
}

func register(username, password, email string) {
	// todo: check params
	// todo: salt + hash password
	user := model.User{
		Name:     username,
		Password: password,
		Email:    email,
	}
	userService := service.NewUserServiceImpl()
	userService.Save(user)
}
