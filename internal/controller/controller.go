package controller

import (
	"log"
	"moon-street/internal/model"
	"moon-street/internal/service"
)

func Route(cmd string) { //injection ;  error raise
	switch cmd {
	case "register":
		Register("testuser", "p1", "a@b.com")
	default:
		log.Printf("Unknown cmd! [%s]", cmd)
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
