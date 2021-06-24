package controller

import (
	"fmt"
	"moon-street/internal/model"
	"moon-street/internal/service"
)

func Route(cmd string) {
	switch cmd {
	case "register":
		Register("testuser", "p1", "a@b.com")
	default:
		Register("testuser", "p1", "a@b.com")
		fmt.Println("Unknown cmd!" + cmd)
	}
}
func Register(username, password, email string) {
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
