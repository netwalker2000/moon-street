package service

import "moon-street/internal/model"

type UserService interface {
	Save(model.User) (int64, error)
}
